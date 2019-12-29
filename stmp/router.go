// Copyright 2019 acrazing <joking.young@gmail.com>. All rights reserved.
// Since 2019-12-28 19:40:24
package stmp

import (
	"context"
)

type MiddlewareFunc func(inCtx context.Context, action uint64, in []byte) (outCtx context.Context, err error)
type InterceptFunc func(inCtx context.Context, action uint64, in []byte) (outCtx context.Context, done bool, out []byte, err error)
type HandlerFunc func(ctx context.Context, in interface{}) (out interface{}, err error)
type ModelFactory func() interface{}

type HandlerOptions struct {
	factory  ModelFactory
	handlers []HandlerFunc
}

type Router struct {
	middlewares  []MiddlewareFunc
	interceptors []InterceptFunc
	handlers     map[uint64]*HandlerOptions
}

func NewRouter() *Router {
	return &Router{
		middlewares:  []MiddlewareFunc{},
		interceptors: []InterceptFunc{},
		handlers:     map[uint64]*HandlerOptions{},
	}
}

func (r *Router) Before(handlers ...MiddlewareFunc) {
	r.middlewares = append(r.middlewares, handlers...)
}

func (r *Router) Intercept(handlers ...InterceptFunc) {
	r.interceptors = append(r.interceptors, handlers...)
}

func (r *Router) Register(action uint64, factory ModelFactory, handlers ...HandlerFunc) {
	if r.handlers[action] == nil {
		r.handlers[action] = &HandlerOptions{factory: factory, handlers: handlers}
	} else {
		r.handlers[action].handlers = append(r.handlers[action].handlers, handlers...)
	}
}

func (r *Router) Dispatch(ctx context.Context, action uint64, payload []byte, codec MediaCodec) (Status, []byte) {
	var err error
	for _, f := range r.middlewares {
		ctx, err = f(ctx, action, payload)
		if err != nil {
			return DetectError(err, StatusInternalServerError)
		}
	}
	var done bool
	var out []byte
	for _, f := range r.interceptors {
		ctx, done, out, err = f(ctx, action, payload)
		if err != nil {
			return DetectError(err, StatusInternalServerError)
		}
		if done {
			return StatusOk, out
		}
	}
	var in interface{}
	h, ok := r.handlers[action]
	if !ok {
		return DetectError(StatusNotFound, StatusNotFound)
	}
	if payload != nil {
		in = h.factory()
		err = codec.Unmarshal(payload, in)
		if err != nil {
			return DetectError(err, StatusBadRequest)
		}
	}
	var v interface{}
	for _, f := range h.handlers {
		v, err = f(ctx, in)
		if err != nil {
			return DetectError(err, StatusInternalServerError)
		}
	}
	if v != nil {
		out, err = codec.Marshal(v)
		if err != nil {
			return DetectError(err, StatusInternalServerError)
		}
	}
	return StatusOk, out
}
