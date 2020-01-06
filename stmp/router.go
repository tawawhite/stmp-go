package stmp

import (
	"context"
	"github.com/twmb/murmur3"
	"strconv"
	"sync"
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
type CloseHandlerFunc func(conn *Conn, status Status, message string)
type ModelFactory func() interface{}

type handlerOptions struct {
	fn   HandlerFunc
	inst interface{}
}

type Router struct {
	mu           sync.RWMutex
	middlewares  []MiddlewareFunc
	interceptors []InterceptFunc
	handlers     map[uint64][]handlerOptions
}

func NewRouter() *Router {
	return &Router{
		middlewares:  []MiddlewareFunc{},
		interceptors: []InterceptFunc{},
		handlers:     map[uint64][]handlerOptions{},
	}
}

// add bypass handler, the handler will not intercept the request
// unless the handler returns error
func (r *Router) Middleware(handlers ...MiddlewareFunc) {
	r.mu.Lock()
	r.middlewares = append(r.middlewares, handlers...)
	r.mu.Unlock()
}

// intercept will intercept request, if done == true, will not pass
// the request to next handlers
func (r *Router) Intercept(handlers ...InterceptFunc) {
	r.mu.Lock()
	r.interceptors = append(r.interceptors, handlers...)
	r.mu.Unlock()
}

// Action bound handler
func (r *Router) Register(inst interface{}, method string, fn HandlerFunc) {
	r.mu.Lock()
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
	r.mu.Unlock()
}

func (r *Router) Unregister(inst interface{}, method string) {
	r.mu.Lock()
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
	r.mu.Unlock()
}

// dispatch a request to handlers
func (r *Router) dispatch(ctx context.Context, action uint64, payload []byte, codec MediaCodec) (status Status, ret []byte) {
	method := ms.actions[action]
	var err error
	for _, f := range r.middlewares {
		ctx, err = f(ctx, method, payload)
		if err != nil {
			return DetectError(err, StatusInternalServerError).Spread()
		}
	}
	var done bool
	for _, f := range r.interceptors {
		ctx, done, ret, err = f(ctx, method, payload)
		if err != nil {
			return DetectError(err, StatusInternalServerError).Spread()
		}
		if done {
			return StatusOk, ret
		}
	}
	h, ok := r.handlers[action]
	if method == nil || !ok {
		return StatusNotFound.Spread()
	}
	in := method.input()
	if payload != nil {
		err = codec.Unmarshal(payload, in)
		if err != nil {
			return DetectError(err, StatusBadRequest).Spread()
		}
	}
	var out interface{}
	for _, hc := range h {
		out, err = hc.fn(ctx, in, hc.inst)
		if err != nil {
			return DetectError(err, StatusInternalServerError).Spread()
		}
	}
	if out != nil {
		ret, err = codec.Marshal(out)
		if err != nil {
			return DetectError(err, StatusInternalServerError).Spread()
		}
	}
	return StatusOk, ret
}
