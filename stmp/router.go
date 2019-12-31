// Copyright 2019 acrazing <joking.young@gmail.com>. All rights reserved.
// Since 2019-12-28 19:40:24
package stmp

import (
	"context"
	"strconv"
)

type MiddlewareFunc func(inCtx context.Context, method string, in []byte) (outCtx context.Context, err error)
type InterceptFunc func(inCtx context.Context, method string, in []byte) (outCtx context.Context, done bool, out []byte, err error)
type HandlerFunc func(ctx context.Context, in interface{}) (out interface{}, err error)
type CloseHandlerFunc func(status Status, message string)
type ModelFactory func() interface{}

var mapMethodAction = map[string]uint64{}
var mapActionMethod = map[uint64]string{}
var dynamicActions []uint64

func RegisterMethodAction(method string, action uint64) {
	if action == 0 {
		action = uint64(len(dynamicActions))
		if action > 0x0fff {
			panic("dynamic method id should range from 0 to 0x0fff")
		}
		dynamicActions = append(dynamicActions, action)
	} else if action < 0x1000 {
		panic("static method id should range from 0x1000 to max")
	}
	if m, ok := mapActionMethod[action]; ok {
		panic("duplicated action 0x" + strconv.FormatUint(action, 16) + " for " + m + " and " + method)
	}
	mapMethodAction[method] = action
	mapActionMethod[action] = method
}

func RegisterMethodActions(pairs ...interface{}) {
	for i := 0; i < len(pairs); i += 2 {
		method := pairs[i].(string)
		action := pairs[i+1].(uint64)
		RegisterMethodAction(method, action)
	}
}

type HandlerOptions struct {
	factory  ModelFactory
	handlers []HandlerFunc
}

type Router struct {
	middlewares  []MiddlewareFunc
	interceptors []InterceptFunc
	handlers     map[uint64]*HandlerOptions
	closeHandler CloseHandlerFunc
}

var noopCloseHandler CloseHandlerFunc = func(status Status, message string) {
}

func NewRouter() *Router {
	return &Router{
		middlewares:  []MiddlewareFunc{},
		interceptors: []InterceptFunc{},
		handlers:     map[uint64]*HandlerOptions{},
		closeHandler: noopCloseHandler,
	}
}

// add bypass handler, the handler will not intercept the request
// unless the handler returns error
func (r *Router) Middleware(handlers ...MiddlewareFunc) {
	r.middlewares = append(r.middlewares, handlers...)
}

// intercept will intercept request, if done == true, will not pass
// the request to next handlers
func (r *Router) Intercept(handlers ...InterceptFunc) {
	r.interceptors = append(r.interceptors, handlers...)
}

// action bound handler
func (r *Router) Register(method string, factory ModelFactory, handlers ...HandlerFunc) {
	action, ok := mapMethodAction[method]
	if !ok {
		panic("method " + method + " is not registered")
	}
	if r.handlers[action] == nil {
		r.handlers[action] = &HandlerOptions{factory: factory, handlers: handlers}
	} else {
		r.handlers[action].handlers = append(r.handlers[action].handlers, handlers...)
	}
}

func (r *Router) SetCloseHandler(h CloseHandlerFunc) {
	r.closeHandler = h
}

// dispatch a request to handlers
func (r *Router) Dispatch(ctx context.Context, action uint64, payload []byte, codec MediaCodec) (Status, []byte) {
	method, ok := mapActionMethod[action]
	if !ok {
		return StatusNotFound, nil
	}
	var err error
	for _, f := range r.middlewares {
		ctx, err = f(ctx, method, payload)
		if err != nil {
			return DetectError(err, StatusInternalServerError)
		}
	}
	var done bool
	var out []byte
	for _, f := range r.interceptors {
		ctx, done, out, err = f(ctx, method, payload)
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
