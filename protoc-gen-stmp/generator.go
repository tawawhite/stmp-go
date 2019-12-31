// Copyright 2019 acrazing <joking.young@gmail.com>. All rights reserved.
// Since 2019-12-30 22:20:00
// most of the code comes from {@link https://github.com/golang/protobuf/protoc-gen-go}
package main

import (
	"errors"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"log"
	"os"
	"strings"
)

type generatorOptions struct {
	lang []string
}

type generator struct {
	request  *plugin.CodeGeneratorRequest  // The input.
	response *plugin.CodeGeneratorResponse // The output.
	options  *generatorOptions
}

func newGenerator() *generator {
	return &generator{
		request:  new(plugin.CodeGeneratorRequest),
		response: new(plugin.CodeGeneratorResponse),
		options:  &generatorOptions{lang: []string{}},
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
		default:
			return errors.New("unknown option: " + key)
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
