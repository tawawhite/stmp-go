// Copyright 2019 acrazing <joking.young@gmail.com>. All rights reserved.
// Since 2019-12-30 22:20:00
// most of the code comes from {@link https://github.com/golang/protobuf/protoc-gen-go}
package main

import (
	"errors"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/golang/protobuf/protoc-gen-go/plugin"
	"log"
	"os"
	"strings"
)

type generatorOptions struct {
	lang   []string
	js     bool
	golang bool

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
			lang:     []string{},
			jsmodule: "cjs",
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
			g.options.lang = strings.Split(value, "+")
			for _, l := range g.options.lang {
				switch l {
				case "js", "javascript":
					g.options.js = true
				case "go", "golang":
					g.options.golang = true
				default:
					return errors.New("unsupported language: " + l)
				}
			}
		case "jspb", "js.pb":
			g.options.jspb = value
		case "jsout", "js.out":
			g.options.jsout = value
		case "jsdts", "js.dts":
			g.options.jsdts = value != "" && value != "0"
		case "jsmodule", "js.module":
			g.options.jsmodule = value
		default:
			return errors.New("unknown option: " + key)
		}
	}
	if g.options.js {
		if g.options.jspb == "" {
			return errors.New("js.pb is required for js language")
		}
		if g.options.jsout == "" {
			return errors.New("js.out is required for js language")
		}
		if g.options.jsmodule != "esm" && g.options.jsmodule != "cjs" {
			return errors.New("unsupported js module: " + g.options.jsmodule)
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
