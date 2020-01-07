package stmp

import (
	"context"
	"github.com/twmb/murmur3"
	"strconv"
	"strings"
	"sync"
)

// the model factory should create instance of input/output, it is used for protoc-gen-stmp
type ModelFactory func() interface{}

// method information, for middleware and interceptors
type Method struct {
	Method    string
	Action    uint64
	ActionHex string
	Input     ModelFactory
	Output    ModelFactory
}

type methodStore = struct {
	methods map[string]uint64
	actions map[uint64]*Method
}

var ms = methodStore{
	methods: map[string]uint64{},
	actions: map[uint64]*Method{},
}

// register a method, this is used for protoc-gen-stmp
func RegisterMethodAction(method string, action uint64, input ModelFactory, output ModelFactory) {
	if action == 0 {
		action, _ = murmur3.Sum128([]byte(method))
		action |= 1 << 63
	}
	if ms.actions[action] != nil && ms.actions[action].Method != method {
		panic("duplicated Action " + strconv.FormatUint(action, 16) + "for method " + ms.actions[action].Method + " and " + method)
	}
	ms.actions[action] = &Method{
		Method:    method,
		Action:    action,
		ActionHex: strings.ToUpper(strconv.FormatUint(action, 16)),
		Input:     input,
		Output:    output,
	}
	ms.methods[method] = action
}

// middleware, which will not intercept the request, just update the context, or do arbitrary check
// if the err is not nil, will intercept the request.
type MiddlewareFunc func(inCtx context.Context) (outCtx context.Context, err error)

// the arbitrary interceptor
//
// if done is true, means the request is intercepted, then if the err is nil, the out will be treated as
// the output content, and status is OK.
// if err is not nil, the request will always be intercepted.
type InterceptFunc func(inCtx context.Context) (outCtx context.Context, done bool, out []byte, err error)

// the handler for protoc-gen-stmp to register specific action handler
type HandlerFunc func(ctx context.Context, in interface{}, inst interface{}) (out interface{}, err error)

type handlerOptions struct {
	fn   HandlerFunc
	inst interface{}
}

type Router struct {
	sync.RWMutex
	middlewares  []MiddlewareFunc
	interceptors []InterceptFunc
	handlers     map[uint64][]handlerOptions
	host         interface{}
}

// create a new router
//
// the host could be *Server or *Client, which is use for dispatch action to hold it in context
func NewRouter(host interface{}) *Router {
	return &Router{
		middlewares:  []MiddlewareFunc{},
		interceptors: []InterceptFunc{},
		handlers:     map[uint64][]handlerOptions{},
		host:         host,
	}
}

// add bypass handler, the handler will not intercept the request
// unless the handler returns error
func (r *Router) Middleware(handlers ...MiddlewareFunc) {
	r.Lock()
	r.middlewares = append(r.middlewares, handlers...)
	r.Unlock()
}

// intercept will intercept request, if done == true, will not pass
// the request to next handlers
func (r *Router) Intercept(handlers ...InterceptFunc) {
	r.Lock()
	r.interceptors = append(r.interceptors, handlers...)
	r.Unlock()
}

// register handler for specified action
//
// the inst is the server/client instance, which is use for unregister check equals, for func is not comparable in go.
// if registered multiple handler for same action, the last one's response will be used as the final response.
func (r *Router) Register(inst interface{}, method string, fn HandlerFunc) {
	r.Lock()
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
	r.Unlock()
}

// unregister handler
func (r *Router) Unregister(inst interface{}, method string) {
	r.Lock()
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
	r.Unlock()
}

// dispatch a request to handlers
func (r *Router) dispatch(c *Conn, p *Packet) (status Status, ret []byte) {
	m := ms.actions[p.Action]
	ctx := withStmp(context.Background(), p, c, m, r)
	for _, f := range r.middlewares {
		newCtx, err := f(ctx)
		if err != nil {
			return DetectError(err, StatusInternalServerError).Spread()
		}
		if newCtx != nil {
			ctx = newCtx
		}
	}
	for _, f := range r.interceptors {
		newCtx, done, ret, err := f(ctx)
		if err != nil {
			return DetectError(err, StatusInternalServerError).Spread()
		}
		if done {
			return StatusOk, ret
		}
		if newCtx != nil {
			ctx = newCtx
		}
	}
	h, ok := r.handlers[p.Action]
	if m == nil || !ok {
		return StatusNotFound.Spread()
	}
	// thinking: should this be nil when input payload is nil?
	in := m.Input()
	if p.Payload != nil {
		err := c.Media.Unmarshal(p.Payload, in)
		if err != nil {
			return DetectError(err, StatusBadRequest).Spread()
		}
	}
	var err error
	var out interface{}
	for _, hc := range h {
		out, err = hc.fn(ctx, in, hc.inst)
		if err != nil {
			return DetectError(err, StatusInternalServerError).Spread()
		}
	}
	if !isNil(out) {
		ret, err = c.Media.Marshal(out)
		if err != nil {
			return DetectError(err, StatusInternalServerError).Spread()
		}
	}
	return StatusOk, ret
}
