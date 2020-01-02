// Copyright 2019 acrazing <joking.young@gmail.com>. All rights reserved.
// Since 2019-12-30 22:20:00
// most of the code comes from {@link https://github.com/golang/protobuf/protoc-gen-go}
package main

import (
	"errors"
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
	// js output file name, all the input files will be composed to one single file(same to pbjs)
	// the path is relative to current directory, that means you must ensure it under protoc output path
	jsout string
	// generate .d.ts for js out, if output lang include js, the is true in default,
	// you can set as empty string or "0" to disable it, this is dependent on pbts generated .d.ts file
	// the file name is same to jspb(replace .js to .d.ts)
	jsdts bool
	// js module mode, could be "cjs" or "esm"
	jsmodule string
}

type generator struct {
	request  *plugin_go.CodeGeneratorRequest  // The input.
	response *plugin_go.CodeGeneratorResponse // The output.
	options  *generatorOptions
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
	s := strings.Join(msgs, " ") + ":" + err.Error()
	log.Print("protoc-gen-stmp: error:", s)
	os.Exit(1)
}

// Fail reports a problem and exits the program.
func (g *generator) fail(msgs ...string) {
	s := strings.Join(msgs, " ")
	log.Print("protoc-gen-stmp: error:", s)
	os.Exit(1)
}
