// Copyright 2019 acrazing <joking.young@gmail.com>. All rights reserved.
// Since 2019-12-28 15:21:22
package main

import (
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"os"
)

func main() {
	g := newGenerator()

	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		g.error(err, "reading Input")
	}

	if err := proto.Unmarshal(data, g.request); err != nil {
		g.error(err, "parsing Input proto")
	}

	if len(g.request.FileToGenerate) == 0 {
		g.fail("no files to generate")
	}
	if err = g.parseOptions(g.request.GetParameter()); err != nil {
		g.error(err, "invalid option")
	}
	if g.options.js {
		g.response.File = append(g.response.File, g.js()...)
	}
	if g.options.golang {
		g.response.File = append(g.response.File, g.golang()...)
	}
	buf, err := proto.Marshal(g.response)
	if err != nil {
		g.error(err, "marshal response")
	}
	_, err = os.Stdout.Write(buf)
	if err != nil {
		g.error(err, "write response")
	}
}
