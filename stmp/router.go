// Copyright 2019 acrazing <joking.young@gmail.com>. All rights reserved.
// Since 2019-12-28 19:40:24
package stmp

import (
	"context"
	"github.com/twmb/murmur3"
	"strconv"
)

type MethodMetadata struct {
	method string
	action uint64
	input  ModelFactory
	output ModelFactory
}

type methodStore = struct {
	methods map[string]uint64
	actions map[uint64]*MethodMetadata
}

var ms = methodStore{
	methods: map[string]uint64{},
	actions: map[uint64]*MethodMetadata{},
}

func RegisterMethodAction(method string, action uint64, input ModelFactory, output ModelFactory) {
	if action == 0 {
		action, _ = murmur3.Sum128([]byte(method))
		action |= 1 << 63
	}
	if ms.actions[action] != nil && ms.actions[action].method != method {
		panic("duplicated Action " + strconv.FormatUint(action, 16) + "for method " + ms.actions[action].method + " and " + method)
	}
	ms.actions[action] = &MethodMetadata{
		method: method,
		action: action,
		input:  input,
		output: output,
	}
	ms.methods[method] = action
}

type MiddlewareFunc func(inCtx context.Context, method *MethodMetadata, in []byte) (outCtx context.Context, err error)
type InterceptFunc func(inCtx context.Context, method *MethodMetadata, in []byte) (outCtx context.Context, done bool, out []byte, err error)
type HandlerFunc func(ctx context.Context, in interface{}, inst interface{}) (out interface{}, err error)
type CloseHandlerFunc func(status Status, message string)
type ModelFactory func() interface{}

type handlerOptions struct {
	fn   HandlerFunc
	inst interface{}
}

type Router interface {
	Middleware(handlers ...MiddlewareFunc)
	Intercept(handlers ...InterceptFunc)
	Register(inst interface{}, method string, fn HandlerFunc)
	Unregister(inst interface{}, method string)
}

type router struct {
	middlewares  []MiddlewareFunc
	interceptors []InterceptFunc
	handlers     map[uint64][]handlerOptions
}

var _ Router = (*router)(nil)

func NewRouter() *router {
	return &router{
		middlewares:  []MiddlewareFunc{},
		interceptors: []InterceptFunc{},
		handlers:     map[uint64][]handlerOptions{},
	}
}

// add bypass handler, the handler will not intercept the request
// unless the handler returns error
func (r *router) Middleware(handlers ...MiddlewareFunc) {
	r.middlewares = append(r.middlewares, handlers...)
}

// intercept will intercept request, if done == true, will not pass
// the request to next handlers
func (r *router) Intercept(handlers ...InterceptFunc) {
	r.interceptors = append(r.interceptors, handlers...)
}

// Action bound handler
func (r *router) Register(inst interface{}, method string, fn HandlerFunc) {
	action, ok := ms.methods[method]
	if !ok {
		panic("method " + method + " is not registered")
	}
	if r.handlers[action] == nil {
		r.handlers[action] = []handlerOptions{{fn: fn, inst: inst}}
	} else {
		for _, v := range r.handlers[action] {
			if v.inst == inst {
				panic("register on method " + method + " multi-times with same instance")
			}
		}
		r.handlers[action] = append(r.handlers[action], handlerOptions{fn: fn, inst: inst})
	}
}

func (r *router) Unregister(inst interface{}, method string) {
	action, ok := ms.methods[method]
	if !ok {
		panic("method " + method + " is not registered")
	}
	hs := r.handlers[action]
	if hs == nil {
		return
	}
	for i, hc := range hs {
		if hc.inst == inst {
			copy(hs[i:], hs[i+1:])
			r.handlers[action] = hs[:len(hs)-1]
			break
		}
	}
}

// dispatch a request to handlers
func (r *router) dispatch(ctx context.Context, action uint64, payload []byte, codec MediaCodec) (status Status, ret []byte) {
	method := ms.actions[action]
	ctx = WithAction(ctx, action)
	var err error
	for _, f := range r.middlewares {
		ctx, err = f(ctx, method, payload)
		if err != nil {
			return DetectError(err, StatusInternalServerError)
		}
	}
	var done bool
	for _, f := range r.interceptors {
		ctx, done, ret, err = f(ctx, method, payload)
		if err != nil {
			return DetectError(err, StatusInternalServerError)
		}
		if done {
			return StatusOk, ret
		}
	}
	h, ok := r.handlers[action]
	if method == nil || !ok {
		return DetectError(StatusNotFound, StatusNotFound)
	}
	in := method.input()
	if payload != nil {
		err = codec.Unmarshal(payload, in)
		if err != nil {
			return DetectError(err, StatusBadRequest)
		}
	}
	var out interface{}
	for _, hc := range h {
		out, err = hc.fn(ctx, in, hc.inst)
		if err != nil {
			return DetectError(err, StatusInternalServerError)
		}
	}
	if out != nil {
		ret, err = codec.Marshal(out)
		if err != nil {
			return DetectError(err, StatusInternalServerError)
		}
	}
	return StatusOk, ret
}
