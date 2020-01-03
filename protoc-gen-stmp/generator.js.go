// Copyright 2020 acrazing <joking.young@gmail.com>. All rights reserved.
// Since 2020-01-02 15:42:43
package main

import (
	"bytes"
	"github.com/acrazing/stmp-go/stmp"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/twmb/murmur3"
	"strconv"
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
		Namespaces:    map[string][]*RenderService{},
	}

	var action uint64
	var stmpService uint64
	var stmpMethod uint64
	for _, filename := range g.request.GetFileToGenerate() {
		req := g.lookupFile(filename)
		ns := req.GetPackage()
		sep := strings.IndexByte(ns, '.')
		if sep == -1 {
			sep = len(ns)
		}
		data.TopNamespaces[ns[:sep]] = struct{}{}
		for _, s := range req.Service {
			service := new(RenderService)
			data.Namespaces[ns] = append(data.Namespaces[ns], service)
			service.ServiceName = upFirst(s.GetName())
			serviceOption, err := proto.GetExtension(s.GetOptions(), stmp.E_Service)
			if err == nil && serviceOption != nil {
				stmpService = *(serviceOption.(*uint64))
			}
			for _, m := range s.GetMethod() {
				method := new(RenderMethod)
				service.Methods = append(service.Methods, method)
				method.MethodName = upFirst(m.GetName())
				method.FullMethod = req.GetPackage() + "." + service.ServiceName + "." + method.MethodName
				methodOption, err := proto.GetExtension(m.GetOptions(), stmp.E_Method)
				if err == nil && serviceOption != nil {
					stmpMethod = *(methodOption.(*uint64))
				}
				if serviceOption != nil && methodOption != nil {
					action = stmpService<<8 | stmpMethod
				} else {
					action = murmur3.Sum64([]byte(method.FullMethod)) | (1 << 63)
				}
				method.ActionHex = strings.ToUpper(strconv.FormatUint(action, 16))
				method.Input = m.GetInputType()[1:]
				method.Output = m.GetOutputType()[1:]
				dot := strings.LastIndexByte(method.Input, '.')
				method.IInput = method.Input[:dot+1] + "I" + method.Input[dot+1:]
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
