package icepop

import (
	"github.com/charmbracelet/wish"
	"github.com/gliderlabs/ssh"
)

// Ensure route implements SessionHandler
var _ SessionHandler = &route{}

// NewRoute creates an instance of a route, representing a handler and
// its associated middleware.
func NewRoute(h ssh.Handler, mw []wish.Middleware, opts ...RouteOption) *route {
	var r route

	r.baseHandler = h
	r.middleware = mw

	for _, opt := range opts {
		opt(&r)
	}

	r.buildHandler()
	return &r
}

type route struct {
	baseHandler ssh.Handler
	middleware  []wish.Middleware

	compiledHandler ssh.Handler

	reversedMiddlewareOrder bool
}

// ServeSSH implements the SessionHandler interface
func (r *route) ServeSSH(s ssh.Session) {
	r.compiledHandler(s)
}

func (r *route) buildHandler() {
	h := r.baseHandler

	if r.reversedMiddlewareOrder {
		// The user asked us to reverse the order, similar to how
		// middleware is parsed by Wish.
		for _, mw := range r.middleware {
			h = mw(h)
		}
	} else {
		for idx := len(r.middleware) - 1; idx >= 0; idx-- {
			h = r.middleware[idx](h)
		}
	}

	r.compiledHandler = h
}

// Build returns the ssh.Handler, marrying route.Middleware to
// r.Handler.
func (r *route) Handler() ssh.Handler {
	return r.compiledHandler
}
