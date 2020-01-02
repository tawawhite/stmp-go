// Copyright 2019 acrazing <joking.young@gmail.com>. All rights reserved.
// Since 2019-12-28 15:21:22
package main

import (
	"encoding/json"
	"github.com/acrazing/stmp-go/stmp"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"os"
)

func main() {
	g := newGenerator()

	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		g.error(err, "reading input")
	}

	if err := proto.Unmarshal(data, g.request); err != nil {
		g.error(err, "parsing input proto")
	}

	if len(g.request.FileToGenerate) == 0 {
		g.fail("no files to generate")
	}
	if err = g.parseOptions(g.request.GetParameter()); err != nil {
		g.error(err, "invalid option")
	}
	serviceOptions := map[string]map[string]interface{}{}
	for _, file := range g.request.ProtoFile {
		serviceOptions[file.GetName()] = map[string]interface{}{}
		for _, s := range file.GetService() {
			ext, err := proto.GetExtension(s.Options, stmp.E_Service)
			if err != nil {
				serviceOptions[file.GetName()][s.GetName()] = err.Error()
			} else {
				serviceOptions[file.GetName()][s.GetName()] = *(ext.(*uint64))
			}
			for _, m := range s.GetMethod() {
				ext, err := proto.GetExtension(m.Options, stmp.E_Method)
				if err != nil {
					serviceOptions[file.GetName()][s.GetName()+"."+m.GetName()] = err.Error()
				} else {
					serviceOptions[file.GetName()][s.GetName()+"."+m.GetName()] = *(ext.(*uint64))
				}
			}
		}
	}
	str, _ := json.MarshalIndent(g.request, "", "  ")
	ioutil.WriteFile("out/debug_protoc_plugin_input.json", str, os.ModePerm)
	str, _ = json.MarshalIndent(g.response, "", "  ")
	ioutil.WriteFile("out/debug_protoc_plugin_output.json", str, os.ModePerm)
	ioutil.WriteFile("out/debug_protoc_input.pb", data, os.ModePerm)
	str, _ = json.MarshalIndent(serviceOptions, "", "  ")
	ioutil.WriteFile("out/debug_protoc_options.json", str, os.ModePerm)
}
