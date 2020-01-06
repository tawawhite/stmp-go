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
	switch g.options.lang {
	case "js":
		g.response.File = append(g.response.File, g.js()...)
	case "golang":
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
