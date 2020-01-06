package main

import (
	"bytes"
	"github.com/golang/protobuf/protoc-gen-go/plugin"
	"strings"
)

func (g *generator) js() []*plugin_go.CodeGeneratorResponse_File {
	var out []*plugin_go.CodeGeneratorResponse_File
	data := &JsRenderData{
		Files:         g.request.GetFileToGenerate(),
		PbFile:        g.options.jspb,
		RootNamespace: g.options.jsroot,
		Module:        g.options.jsmodule,
		TopNamespaces: map[string]struct{}{},
		Namespaces:    map[string][]*ServiceOptions{},
	}

	for _, filename := range g.request.GetFileToGenerate() {
		req := g.lookupFile(filename)
		ns := req.GetPackage()
		sep := strings.IndexByte(ns, '.')
		if sep == -1 {
			sep = len(ns)
		}
		data.TopNamespaces[ns[:sep]] = struct{}{}
		for _, s := range req.Service {
			service := g.parseService(s)
			data.Namespaces[ns] = append(data.Namespaces[ns], service)
			for _, m := range s.GetMethod() {
				g.parseMethod(req, m, service)
			}
		}
	}

	jsOut := &plugin_go.CodeGeneratorResponse_File{}
	out = append(out, jsOut)
	jsOut.Name = &g.options.jsout
	jsBuf := bytes.NewBuffer(nil)
	err := templateJs.Execute(jsBuf, data)
	if err != nil {
		g.error(err, "building js")
	}
	content := jsBuf.String()
	jsOut.Content = &content
	if g.options.jsdts {
		dtsOut := &plugin_go.CodeGeneratorResponse_File{}
		out = append(out, dtsOut)
		dtsName := strings.TrimSuffix(g.options.jsout, ".js") + ".d.ts"
		dtsOut.Name = &dtsName
		dtsBuf := bytes.NewBuffer(nil)
		err = templateDts.Execute(dtsBuf, data)
		if err != nil {
			g.error(err, "building dts")
		}
		content := dtsBuf.String()
		dtsOut.Content = &content
	}
	return out
}
