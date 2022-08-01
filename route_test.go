package icepop

import (
	"testing"

	"github.com/charmbracelet/wish"
	"github.com/gliderlabs/ssh"
	"github.com/stretchr/testify/assert"
)

var (
	// TODO: use test ssh.Session and test middleware actions.
	h  = func(s ssh.Session) { s.Context().SetValue("handler", "handler") }
	mw = []wish.Middleware{
		func(h ssh.Handler) ssh.Handler {
			return func(s ssh.Session) {
				s.Context().SetValue("first", "first")
				h(s)
			}
		},
		func(h ssh.Handler) ssh.Handler {
			return func(s ssh.Session) {
				s.Context().SetValue("second", "second")
				h(s)
			}
		},
	}
)

func TestRouteOptionReverseMiddlewareOrder(t *testing.T) {
	r := NewRoute(h, mw, OptExecuteMiddlewareInReverse())
	assert.Equal(t, true, r.reversedMiddlewareOrder, "the middleware order should be reversed")

	assert.ObjectsAreEqual(h, r.baseHandler)
	assert.ObjectsAreEqual(r.middleware, mw)
	comp := mw[1](mw[0](h))
	assert.ObjectsAreEqual(r.compiledHandler, comp)
}

func TestRouteNoOpti(t *testing.T) {
	r := NewRoute(h, mw, OptExecuteMiddlewareInReverse())
	assert.Equal(t, true, r.reversedMiddlewareOrder, "the middleware order should be as declared")

	assert.ObjectsAreEqual(h, r.baseHandler)
	assert.ObjectsAreEqual(r.middleware, mw)
	comp := mw[0](mw[1](h))
	assert.ObjectsAreEqual(r.compiledHandler, comp)
}
