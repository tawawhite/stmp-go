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
		action, _ = murmur3.StringSum128(method)
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

type PreHandlerFunc func(inCtx context.Context) (outCtx context.Context, err error)
type PostHandlerFunc func(ctx context.Context, status Status, header Header, payload []byte) error
type InterceptFunc func(inCtx context.Context) (outCtx context.Context, done bool, out []byte, err error)
type HandlerFunc func(ctx context.Context, in interface{}, inst interface{}) (out interface{}, err error)
type ListenerFunc func(ctx context.Context, in interface{}, inst interface{})

type handlerOptions struct {
	fn   HandlerFunc
	inst interface{}
}

type listenerOptions struct {
	fn   ListenerFunc
	inst interface{}
}

// Core router for dispatch a request
//
// The process order is: pre -> interceptor -> handler -> post -> response -> listener
//
// Please note the handler is not routine-safe
type Router struct {
	sync.RWMutex
	preHandlers  []PreHandlerFunc
	interceptors []InterceptFunc
	listeners    map[uint64][]listenerOptions
	handlers     map[uint64]handlerOptions
	postHandlers []PostHandlerFunc
	host         interface{}
}

// Create a router
//
// The host is a *Server or *Client, which is used for create a stmp context
// you can use SelectServer(ctx) or SelectClient(ctx) to get it in handler
func NewRouter(host interface{}) *Router {
	return &Router{
		listeners: map[uint64][]listenerOptions{},
		handlers:  map[uint64]handlerOptions{},
		host:      host,
	}
}

// Add a pre handler, which is executed at fist,
// and could update the context or stop the dispatch chain by emit an error
func (r *Router) Pre(handlers ...PreHandlerFunc) {
	r.preHandlers = append(r.preHandlers, handlers...)
}

// Add a post handler, which is executed just before response the packet,
// it could emit an error to change the response status and payload, or update the response header
func (r *Router) Post(handlers ...PostHandlerFunc) {
	r.postHandlers = append(r.postHandlers, handlers...)
}

// Add an interceptor, which is executed after pre, and before handler
func (r *Router) Intercept(handlers ...InterceptFunc) {
	r.interceptors = append(r.interceptors, handlers...)
}

// Add a handler, which is executed after interceptor, and before post handler,
// it is the core handler to generate the response packet.
//
// Please note the handlers of pre, interceptor, handler, post, listener is executed serially in the read channel,
// which means it will block the read from the connection, so it should run A.S.A.P.
func (r *Router) Handle(method string, inst interface{}, fn HandlerFunc) {
	action, ok := ms.methods[method]
	if !ok {
		panic("method " + method + " is not registered")
	}
	if _, ok := r.handlers[action]; ok {
		panic("method " + method + " is handled already")
	}
	r.handlers[action] = handlerOptions{fn: fn, inst: inst}
}

// Add a listener, which is executed after response
func (r *Router) AddListener(method string, inst interface{}, fn ListenerFunc) {
	action, ok := ms.methods[method]
	if !ok {
		panic("method " + method + " is not registered")
	}
	if r.listeners[action] == nil {
		r.listeners[action] = []listenerOptions{{fn: fn, inst: inst}}
	} else {
		for _, v := range r.listeners[action] {
			if v.inst == inst {
				panic("register on method " + method + " multi-times with same instance")
			}
		}
		r.listeners[action] = append(r.listeners[action], listenerOptions{fn: fn, inst: inst})
	}
}

// Remove a listener
func (r *Router) RemoveListener(method string, inst interface{}) {
	action, ok := ms.methods[method]
	if !ok {
		panic("method " + method + " is not registered")
	}
	hs := r.listeners[action]
	if hs == nil {
		return
	}
	for i, hc := range hs {
		if hc.inst == inst {
			copy(hs[i:], hs[i+1:])
			r.listeners[action] = hs[:len(hs)-1]
			break
		}
	}
}

func (r *Router) res(ctx context.Context, c *Conn, q *Packet, in interface{}, status Status, payload []byte) {
	header := selectStmp(ctx).header
	for _, p := range r.postHandlers {
		if err := p(ctx, status, header, payload); err != nil {
			status, payload = DetectError(err, StatusInternalServerError).Spread()
			break
		}
	}
	if q.Kind != MessageKindResponse {
		c.send(ctx, &Packet{
			Kind:    MessageKindResponse,
			Status:  status,
			Mid:     q.Mid,
			Header:  header,
			Payload: payload,
		}, false)
	}
	r.RLock()
	defer r.RUnlock()
	if in == nil && q.Payload != nil {
		m := ms.actions[q.Action]
		if m == nil {
			return
		}
		in = m.Input()
		if err := c.Media.Unmarshal(q.Payload, in); err != nil {
			return
		}
	}
	for _, hc := range r.listeners[q.Action] {
		hc.fn(ctx, in, hc.inst)
	}
}

func (r *Router) err(ctx context.Context, c *Conn, q *Packet, in interface{}, err error, rs Status) {
	se := DetectError(err, rs)
	r.res(ctx, c, q, in, se.Code(), []byte(se.Message()))
}

// dispatch a request to handlers
func (r *Router) dispatch(c *Conn, q *Packet) {
	m := ms.actions[q.Action]
	ctx := withStmp(context.Background(), q, c, m, r.host)
	for _, f := range r.preHandlers {
		newCtx, err := f(ctx)
		if err != nil {
			r.err(ctx, c, q, nil, err, StatusInternalServerError)
			return
		}
		if newCtx != nil {
			ctx = newCtx
		}
	}
	for _, f := range r.interceptors {
		newCtx, done, ret, err := f(ctx)
		if err != nil {
			r.err(ctx, c, q, nil, err, StatusInternalServerError)
			return
		}
		if done {
			r.res(ctx, c, q, nil, StatusOk, ret)
			return
		}
		if newCtx != nil {
			ctx = newCtx
		}
	}
	h, ok := r.handlers[q.Action]
	if m == nil || !ok {
		r.err(ctx, c, q, nil, StatusNotFound, StatusNotFound)
		return
	}
	in := m.Input()
	if q.Payload != nil {
		err := c.Media.Unmarshal(q.Payload, in)
		if err != nil {
			r.err(ctx, c, q, nil, err, StatusBadRequest)
			return
		}
	}
	out, err := h.fn(ctx, in, h.inst)
	if err != nil {
		r.err(ctx, c, q, in, err, StatusInternalServerError)
		return
	}
	if isNil(out) {
		r.res(ctx, c, q, in, StatusOk, nil)
		return
	} else if !isNil(out) {
		ret, err := c.Media.Marshal(out)
		if err != nil {
			r.err(ctx, c, q, in, err, StatusInternalServerError)
		} else {
			r.res(ctx, c, q, in, StatusOk, ret)
		}
	}
}
