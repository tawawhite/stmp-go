package main

import (
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/golang/protobuf/protoc-gen-go/plugin"
	"path"
	"strconv"
	"strings"
)

func resolveGolangFile(file *descriptor.FileDescriptorProto) (pkg string, name string) {
	pkg = file.GetOptions().GetGoPackage()
	name = strings.TrimSuffix(file.GetName(), ".proto") + ".stmp.go"
	if pkg == "" {
		pkg = strings.ReplaceAll(file.GetPackage(), ".", "_")
		return
	}
	force := strings.LastIndexByte(pkg, '/')
	if force == -1 {
		return
	} else {
		sep := strings.IndexByte(pkg, ';')
		base := strings.TrimSuffix(path.Base(file.GetName()), ".proto") + ".stmp.go"
		if sep > -1 {
			name = pkg[:sep] + "/" + base
			pkg = pkg[sep+1:]
			return
		} else {
			name = pkg + "/" + base
			pkg = pkg[force+1:]
			return
		}
	}
}

func (g *generator) golangFile(filename string) (res *plugin_go.CodeGeneratorResponse_File, err error) {
	req := g.lookupFile(filename)
	if len(req.GetService()) == 0 {
		return
	}
	res = new(plugin_go.CodeGeneratorResponse_File)
	pkg, name := resolveGolangFile(req)
	res.Name = &name
	ds := newDepSet(g, path.Dir(req.GetName()))
	data := new(GolangRenderData)
	data.Filename = req.GetName()
	data.Package = pkg
	data.Deps = ds.bases
	builder := new(strings.Builder)
	for _, s := range req.GetService() {
		service := g.parseService(s)
		data.Services = append(data.Services, service)
		for _, m := range s.GetMethod() {
			method := g.parseMethod(req, m, service)
			method.Input = ds.Resolve(method.Input)
			method.Output = ds.Resolve(method.Output)
		}
	}
	err = templateGolang.Execute(builder, data)
	if err != nil {
		return
	}
	content := builder.String()
	res.Content = &content
	return
}

type GoDepSet struct {
	g     *generator
	dir   string
	bases map[string]string
	paths map[string]string
}

func newDepSet(g *generator, dir string) *GoDepSet {
	return &GoDepSet{
		g:     g,
		dir:   dir,
		bases: map[string]string{},
		paths: map[string]string{},
	}
}

// stmp.examples.room.%s.Status -> %s_Status
// google.protobuf.Empty -> empty.Empty
// foo.bar.Empty -> empty2.Empty
func (s *GoDepSet) Resolve(typ string) string {
	file, _ := s.g.lookupType(typ)
	if file == nil {
		s.g.fail("cannot find file for type " + typ)
		return ""
	}
	typ = upFirst(strings.ReplaceAll(strings.TrimPrefix(typ, file.GetPackage()+"."), ".", "_"))
	if path.Dir(file.GetName()) == s.dir {
		return typ
	}
	pkg, name := resolveGolangFile(file)
	ident := path.Dir(name)
	if old, ok := s.paths[ident]; ok {
		return old + "." + typ
	} else {
		if _, ok := s.bases[pkg]; ok {
			for i := 0; ; i++ {
				is := pkg + "_" + strconv.Itoa(i)
				if _, ok := s.bases[is]; !ok {
					pkg = is
					break
				}
			}
		}
		s.bases[pkg] = ident
		s.paths[ident] = pkg
		return pkg + "." + typ
	}
}

func (g *generator) golang() (files []*plugin_go.CodeGeneratorResponse_File) {
	for _, filename := range g.request.FileToGenerate {
		file, err := g.golangFile(filename)
		if err != nil {
			g.error(err, "processing "+filename)
		}
		if file != nil {
			files = append(files, file)
		}
	}
	return
}
