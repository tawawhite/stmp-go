// Copyright 2020 acrazing <joking.young@gmail.com>. All rights reserved.
// Since 2020-01-03 19:39:13
package main

import "text/template"

type GolangRenderData struct {
	Filename string
	Package  string
	Deps     map[string]string
	Services []*ServiceOptions
}

var templateGolang *template.Template

func init() {
	var err error
	templateGolang, err = template.New("golang").Parse(`// Code generated by protoc-gen-stmp. DO NOT EDIT.
// source: {{.Filename}}
package {{.Package}}

import (
	"context"
	"github.com/acrazing/stmp-go/stmp"
{{range $name, $path := .Deps}}	{{$name}} "{{$path}}"
{{end}})

{{range $index, $service := .Services}}
func init() {
{{range $index, $method := $service.Methods}}	stmp.RegisterMethodAction("{{$method.FullMethod}}", 0x{{$method.ActionHex}}, func() interface{} { return &{{$method.Input}}{} }, func() interface{} { return &{{$method.Output}}{} })
{{end}}}
{{if $service.IsService}}
type STMP{{$service.ServiceName}}Server interface {
{{range $index, $method := $service.Methods}}	{{$method.MethodName}}(ctx context.Context, in *{{$method.Input}}) (out *{{$method.Output}}, err error)
{{end}}}

func STMPRegister{{$service.ServiceName}}Server(srv *stmp.Server, s STMP{{$service.ServiceName}}Server) {
{{range $index, $method := $service.Methods}}	srv.Register(s, "{{$method.FullMethod}}", func(ctx context.Context, in interface{}, inst interface{}) (out interface{}, err error) { return inst.(STMP{{$service.ServiceName}}Server).{{$method.MethodName}}(ctx, in.(*{{$method.Input}})) })
{{end}}}

func STMPUnregister{{$service.ServiceName}}Server(srv *stmp.Server, s STMP{{$service.ServiceName}}Server) {
{{range $index, $method := $service.Methods}}	srv.Unregister(s, "{{$method.FullMethod}}")
{{end}}}

type STMP{{$service.ServiceName}}Client interface {
{{range $index, $method := $service.Methods}}	{{$method.MethodName}}(ctx context.Context, in *{{$method.Input}}, opts ...*stmp.CallOptions) (*{{$method.Output}}, error)
{{end}}}

type stmp{{$service.ServiceName}}Client struct {
	c *stmp.ClientConn
}
{{range $index, $method := $service.Methods}}
func (s *stmp{{$service.ServiceName}}Client) {{$method.MethodName}}(ctx context.Context, in *{{$method.Input}}, opts ...*stmp.CallOptions) (*{{$method.Output}}, error) {
	out, err := s.c.Invoke(ctx, "{{$method.FullMethod}}", in, stmp.BuildCallOptions(opts...))
	return out.(*{{$method.Output}}), err
}
{{end}}
func STMPNew{{$service.ServiceName}}Client(c *stmp.ClientConn) STMP{{$service.ServiceName}}Client {
	return &stmp{{$service.ServiceName}}Client{c: c}
}
{{end}}
{{if $service.IsEvents}}
type STMP{{$service.ServiceName}}Listener interface {
{{range $index, $method := $service.Methods}}	Handle{{$method.MethodName}}(ctx context.Context, in *{{$method.Input}}) (out *{{$method.Output}}, err error)
{{end}}}

func STMPRegister{{$service.ServiceName}}Listener(cc *stmp.ClientConn, s STMP{{$service.ServiceName}}Listener) {
{{range $index, $method := $service.Methods}}	cc.Register(s, "{{$method.FullMethod}}", func(ctx context.Context, in interface{}, inst interface{}) (out interface{}, err error) { return inst.(STMP{{$service.ServiceName}}Listener).Handle{{$method.MethodName}}(ctx, in.(*{{$method.Input}})) })
{{end}}}

func STMPUnregister{{$service.ServiceName}}Listener(cc *stmp.ClientConn, s STMP{{$service.ServiceName}}Listener) {
{{range $index, $method := $service.Methods}}	cc.Unregister(s, "{{$method.FullMethod}}")
{{end}}}

type STMP{{$service.ServiceName}}Broadcaster struct{}
{{range $index, $method := $service.Methods}}
func (s STMP{{$service.ServiceName}}Broadcaster) {{$method.MethodName}}(ctx context.Context, in *{{$method.Input}}, conn *stmp.Conn, opts ...*stmp.CallOptions) (*{{$method.Output}}, error) {
	out, err := conn.Invoke(ctx, "{{$method.FullMethod}}", in, stmp.BuildCallOptions(opts...))
	return out.(*{{$method.Output}}), err
}

func (s STMP{{$service.ServiceName}}Broadcaster) {{$method.MethodName}}ToList(ctx context.Context, in *{{$method.Input}}, conns ...*stmp.Conn) error {
	payloads := stmp.NewPayloadMap(in)
	for _, conn := range conns {
		payload, err := payloads.Marshal(conn)
		if err != nil {
			return err
		}
		_, err = conn.Call(ctx, "{{$method.FullMethod}}", payload, stmp.NotifyOptions)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s STMP{{$service.ServiceName}}Broadcaster) {{$method.MethodName}}ToSet(ctx context.Context, in *{{$method.Input}}, conns stmp.ConnSet, exclude ...*stmp.Conn) error {
	payloads := stmp.NewPayloadMap(in)
	for conn := range conns {
		for _, e := range exclude {
			if e == conn {
				conn = nil
				break
			}
		}
		if conn == nil {
			continue
		}
		payload, err := payloads.Marshal(conn)
		if err != nil {
			return err
		}
		_, err = conn.Call(ctx, "{{$method.FullMethod}}", payload, stmp.NotifyOptions)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s STMP{{$service.ServiceName}}Broadcaster) {{$method.MethodName}}ToAll(ctx context.Context, in *{{$method.Input}}, srv *stmp.Server, filter stmp.ConnFilter) error {
	return srv.Broadcast(ctx, "{{$method.FullMethod}}", in, filter)
}{{end}}{{/* broadcaster.methods */}}{{end}}{{/* service.isEvents */}}{{end}}{{/* range $services */}}
`)
	if err != nil {
		panic(err)
	}
}
