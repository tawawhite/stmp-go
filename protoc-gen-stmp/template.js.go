// Copyright 2020 acrazing <joking.young@gmail.com>. All rights reserved.
// Since 2020-01-03 22:59:02
package main

import "text/template"

type JsRenderData struct {
	Files         []string
	PbFile        string
	RootNamespace string
	Module        string
	TopNamespaces map[string]struct{}
	Namespaces    map[string][]*RenderService
}

var templateJs *template.Template
var templateDts *template.Template

func init() {
	var err error
	templateJs, err = template.New("js").Parse(`// Code generated by protoc-gen-stmp. DO NOT EDIT.
{{range $index, $file := .Files}}// source: {{$file}}
{{end}}{{if eq .Module "esm"}}import pb from "{{.PbFile}}";
import { PayloadMap, registerMethodAction, notifyOptions } from "stmp";
{{else}}const pb = require("{{.PbFile}}");
const {} = require("stmp");
{{end}}
const {{.RootNamespace}} = Object.create(null);

{{if eq .Module "esm"}}export default {{.RootNamespace}};
{{else}}module.exports = exports.default = {{.RootNamespace}};
{{end}}
function initNamespace(root, ns, factory) {
    for (const item of ns.split(".")) {
        root = (root[item] = root[item] || Object.create(null))
    }
	factory(root)
}

{{range $ns, $services := .Namespaces}}initNamespace({{$.RootNamespace}}, "{{$ns}}", (ns) => {
{{range $i1, $service := $services}}

{{range $i2, $method := $service.Methods}}  registerMethodAction("{{$method.FullMethod}}", "{{$method.ActionHex}}", pb.{{$method.Input}}, pb.{{$method.Output}});
{{end}}
  ns.{{$service.ServiceName}}Server = class {{$service.ServiceName}}Server {
    static register(srv, inst) {
{{range $i2, $method := $service.Methods}}      srv.register(inst, "{{$method.FullMethod}}", inst.{{$method.MethodName}});
{{end}}    }

    static unregister(srv, inst) {
{{range $i2, $method := $service.Methods}}      srv.unregister(inst, "{{$method.FullMethod}}");
{{end}}    }

{{range $i2, $method := $service.Methods}}    {{$method.MethodName}}(ctx, input, output) { throw new Error("not implemented") }
{{end}}  };
  
  ns.{{$service.ServiceName}}Broadcaster = class {{$service.ServiceName}}Broadcaster {
{{range $i2, $method := $service.Methods}}    {{$method.MethodName}}ToOne(input, conn, options) { return conn.invoke("{{$method.FullMethod}}", input, options) }
    {{$method.MethodName}}ToSet(input, conns) { const pm = new PayloadMap(input); for (const conn of conns) conn.call("{{$method.FullMethod}}", pm.get(conn), notifyOptions) }
    {{$method.MethodName}}ToAll(input, srv, filter) { return srv.broadcast("{{$method.FullMethod}}", input, filter) }
{{end}}  };
  
  ns.{{$service.ServiceName}}Client = class {{$service.ServiceName}}Client {
    constructor(conn) { this.conn = conn }
{{range $i2, $method := $service.Methods}}    {{$method.MethodName}}(input, options) { return this.conn.invoke("{{$method.FullMethod}}", input, options) }
{{end}}  };{{end}}
});{{end}}
`)
	if err != nil {
		panic(err)
	}
	templateDts, err = template.New("dts").Parse(`// Code generated by protoc-gen-stmp. DO NOT EDIT.
{{range $index, $file := .Files}}// source: {{$file}}
{{end}}import pb from "{{.PbFile}}";
import { CallOptions, Connection, ConnFilter, Context, Server } from "stmp";

export default {{.RootNamespace}};

declare namespace {{.RootNamespace}} {
{{range $ns, $services := .Namespaces}}  namespace {{$ns}} {
{{range $i1, $service := $services}}

    class {{$service.ServiceName}}Server {
      static register(srv: Server, inst: {{$service.ServiceName}}Server): void
      static unregister(srv: Server, inst: {{$service.ServiceName}}Server): void
{{range $i1, $method := $service.Methods}}      {{$method.MethodName}}(ctx: Context, input: pb.{{$method.Input}}, output: pb.{{$method.Output}}): void | Promise<void>
{{end}}    }

    class {{$service.ServiceName}}Broadcaster {
      constructor()
{{range $i1, $method := $service.Methods}}      {{$method.MethodName}}ToOne(input: pb.{{$method.IInput}}, conn: Connection, options?: Partial<CallOptions>): Promise<pb.{{$method.Output}}>
      {{$method.MethodName}}ToSet(input: pb.{{$method.IInput}}, conns: Set<Connection>): void
      {{$method.MethodName}}ToAll(input: pb.{{$method.IInput}}, srv: Server, filter?: ConnFilter): void
{{end}}    }

    class {{$service.ServiceName}}Client {
      private conn: Connection;
      constructor(conn: Connection)
{{range $i1, $method := $service.Methods}}      {{$method.MethodName}}(data: pb.{{$method.IInput}}, options?: Partial<CallOptions>): Promise<pb.{{$method.Output}}>
{{end}}    }{{end}}
  }
{{end}}}
`)
	if err != nil {
		panic(err)
	}
}
