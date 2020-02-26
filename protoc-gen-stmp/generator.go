package main

import (
	"errors"
	"github.com/acrazing/stmp-go/stmp"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/twmb/murmur3"
	"log"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

type generatorOptions struct {
	lang string
	// the js pb file to import
	// for example, if you use protobufjs to generate pb with command `pbjs -o ./foo.pb.js ./*.proto`
	// then you can set this option as ./foo.pb.js
	// the path is relative to current directory
	jspb string
	// js Output file name, all the Input files will be composed to one single file(same to pbjs)
	// the path is relative to current directory, that means you must ensure it under protoc Output path
	jsout string
	// generate .d.ts for js out, if Output lang include js, the is true in default,
	// you can set as empty string or "0" to disable it, this is dependent on pbts generated .d.ts file
	// the file name is same to jspb(replace .js to .d.ts)
	jsdts bool
	// js module mode, could be "cjs" or "esm"
	jsmodule string
	// the root namespace name, useful for .d.ts auto import, default is stmp
	jsroot string
}

type generator struct {
	request    *plugin_go.CodeGeneratorRequest  // The Input.
	response   *plugin_go.CodeGeneratorResponse // The Output.
	options    *generatorOptions
	typesCache map[string][2]interface{}
}

func newGenerator() *generator {
	return &generator{
		request:  new(plugin_go.CodeGeneratorRequest),
		response: new(plugin_go.CodeGeneratorResponse),
		options: &generatorOptions{
			lang:     "",
			jsmodule: "esm",
			jsdts:    true,
		},
		typesCache: map[string][2]interface{}{},
	}
}

func (g *generator) parseOptions(argv string) error {
	for _, item := range strings.Split(argv, ",") {
		sep := strings.IndexByte(item, '=')
		if sep == -1 {
			return errors.New("invalid option format: " + item)
		}
		key := strings.TrimSpace(item[0:sep])
		value := strings.TrimSpace(item[sep+1:])
		switch key {
		case "lang":
			switch value {
			case "js", "javascript":
				g.options.lang = "js"
			case "go", "golang":
				g.options.lang = "golang"
			}
		case "jspb", "js.pb":
			g.options.jspb = value
		case "jsout", "js.out":
			g.options.jsout = value
		case "jsdts", "js.dts":
			g.options.jsdts = value != "" && value != "0"
		case "jsmodule", "js.module":
			g.options.jsmodule = value
		case "jsroot", "js.root":
			g.options.jsroot = value
		default:
			return errors.New("unknown option: " + key)
		}
	}
	if g.options.lang == "js" {
		if g.options.jspb == "" {
			return errors.New("js.pb is required for js language")
		}
		if g.options.jsout == "" {
			return errors.New("js.out is required for js language")
		}
		if g.options.jsmodule != "esm" && g.options.jsmodule != "cjs" {
			return errors.New("unsupported js module: " + g.options.jsmodule)
		}
		if g.options.jsroot == "" {
			g.options.jsroot = "stmp"
		}
		g.options.jsout = strings.TrimSuffix(g.options.jsout, ".js") + ".js"
		g.options.jsout = strings.TrimPrefix(g.options.jsout, "./")
		if strings.HasPrefix(g.options.jspb, ".") {
			var err error
			g.options.jspb, err = filepath.Rel(path.Dir(g.options.jsout), g.options.jspb)
			if err != nil {
				return err
			}
			g.options.jspb = strings.TrimSuffix(g.options.jspb, ".js")
			if !strings.HasPrefix(g.options.jspb, "./") && !strings.HasPrefix(g.options.jspb, "../") {
				g.options.jspb = "./" + g.options.jspb
			}
		}
	}
	return nil
}

// error reports a problem, including an error, and exits the program.
func (g *generator) error(err error, msgs ...string) {
	s := strings.Join(msgs, " ") + ": " + err.Error()
	log.Print("protoc-gen-stmp: error: ", s)
	os.Exit(1)
}

// Fail reports a problem and exits the program.
func (g *generator) fail(msgs ...string) {
	s := strings.Join(msgs, " ")
	log.Print("protoc-gen-stmp: error:", s)
	os.Exit(1)
}

func (g *generator) lookupFile(filename string) *descriptor.FileDescriptorProto {
	for _, f := range g.request.ProtoFile {
		if f.GetName() == filename {
			return f
		}
	}
	return nil
}

func enumIs(pkg string, name string, enumTypes []*descriptor.EnumDescriptorProto) (bool, interface{}) {
	if len(enumTypes) == 0 {
		return false, nil
	}
	for _, e := range enumTypes {
		if pkg+"."+e.GetName() == name {
			return true, e
		}
	}
	return false, nil
}

func messageIs(pkg string, name string, messageTypes []*descriptor.DescriptorProto) (bool, interface{}) {
	if len(messageTypes) == 0 {
		return false, nil
	}
	for _, m := range messageTypes {
		subPkg := pkg + "." + m.GetName()
		if subPkg == name {
			return true, m
		}
		if !strings.HasPrefix(name, subPkg) {
			continue
		}
		if ok, v := messageIs(subPkg, name, m.GetNestedType()); ok {
			return ok, v
		}
		if ok, v := enumIs(subPkg, name, m.GetEnumType()); ok {
			return ok, v
		}
	}
	return false, nil
}

func (g *generator) lookupType(name string) (*descriptor.FileDescriptorProto, interface{}) {
	_, ok := g.typesCache[name]
	if !ok {
		for _, file := range g.request.ProtoFile {
			pkg := file.GetPackage()
			if !strings.HasPrefix(name, pkg) {
				continue
			}
			if ok, v := messageIs(pkg, name, file.GetMessageType()); ok {
				g.typesCache[name] = [2]interface{}{file, v}
				break
			}
			if ok, v := enumIs(pkg, name, file.GetEnumType()); ok {
				g.typesCache[name] = [2]interface{}{file, v}
				break
			}
		}
		_, ok := g.typesCache[name]
		if !ok {
			g.typesCache[name] = [2]interface{}{nil, nil}
		}
	}
	if g.typesCache[name][0] == nil {
		return nil, nil
	} else {
		return g.typesCache[name][0].(*descriptor.FileDescriptorProto), g.typesCache[name][1]
	}
}

func upFirst(str string) string {
	if len(str) == 0 {
		return str
	}
	c0 := str[0]
	if c0 >= 'a' && c0 <= 'z' {
		return string(c0-('a'-'A')) + str[1:]
	}
	return str
}

type MethodOptions struct {
	Id         *uint64
	MethodName string
	FullMethod string
	Action     uint64
	ActionHex  string
	IInput     string
	Input      string
	Output     string
}

type ServiceOptions struct {
	Id          *uint64
	ServiceName string
	IsService   bool
	IsEvents    bool
	Methods     []*MethodOptions
}

func (g *generator) parseService(s *descriptor.ServiceDescriptorProto) *ServiceOptions {
	service := new(ServiceOptions)
	service.ServiceName = upFirst(s.GetName())
	serviceOption, err := proto.GetExtension(s.GetOptions(), stmp.E_Service)
	if err == nil && serviceOption != nil {
		service.Id = serviceOption.(*uint64)
	}
	kindOption, err := proto.GetExtension(s.GetOptions(), stmp.E_Kind)
	if err == nil && kindOption != nil {
		service.IsService = kindOption.(*stmp.ServiceKind).Service
		service.IsEvents = kindOption.(*stmp.ServiceKind).Events
	} else if strings.HasSuffix(service.ServiceName, "Service") {
		service.IsService = true
	} else if strings.HasSuffix(service.ServiceName, "Events") {
		service.IsEvents = true
	} else {
		service.IsService = true
		service.IsEvents = true
	}
	return service
}

func (g *generator) parseMethod(req *descriptor.FileDescriptorProto, m *descriptor.MethodDescriptorProto, service *ServiceOptions) *MethodOptions {
	method := new(MethodOptions)
	method.MethodName = upFirst(m.GetName())
	method.FullMethod = req.GetPackage() + "." + service.ServiceName + "." + method.MethodName
	methodOption, err := proto.GetExtension(m.GetOptions(), stmp.E_Method)
	if err == nil && methodOption != nil {
		method.Id = methodOption.(*uint64)
	}
	if method.Id != nil && service.Id != nil {
		method.Action = *service.Id<<8 | *method.Id
	} else {
		method.Action = uint64(murmur3.StringSum32(method.FullMethod))
	}
	method.ActionHex = strings.ToUpper(strconv.FormatUint(method.Action, 16))
	method.Input = m.GetInputType()[1:]
	method.Output = m.GetOutputType()[1:]
	sep := strings.LastIndexByte(method.Input, '.')
	method.IInput = method.Input[0:sep] + ".I" + method.Input[sep+1:]
	service.Methods = append(service.Methods, method)
	return method
}
