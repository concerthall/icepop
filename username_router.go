package icepop

import (
	"github.com/charmbracelet/wish"
	"github.com/gliderlabs/ssh"
)

// Ensure usernameRouter implements SessionHandler
var _ SessionHandler = &usernameRouter{}

// NewCommandRouter returns a ssh.Handler that routes to registered
// SessionHandlers based on the ssh.Session.Username value. If the
// username is not registered, then a configurable message is printed
// to the user and the session is closed.
func NewUsernameRouter(opts ...UsernameRouterOption) *usernameRouter {
	r := &usernameRouter{
		Router: Router{
			routes: map[string]SessionHandler{},
		},
		notFoundMsg:    "not found",
		notProvidedMsg: "you need to provide a username",
	}

	// apply options
	for _, opt := range opts {
		opt(r)
	}

	return r
}

// usernameRouter is a router that parses the ssh.Session username
// and routes to to handlers accordingly.
type usernameRouter struct {
	Router

	// notFoundMsg is the message sent to the user
	// when the username they provided is not registered
	// as a handler.
	notFoundMsg string

	// notProvidedMsg is the message sent to the user
	// when they don't provide a username.
	notProvidedMsg string
}

// AsMiddleware returns the entire router handler chain
// as a Middleware. This must be the final handler executed
// as it does not pass control to handlers provided as input.
func (r *usernameRouter) AsMiddleware() wish.Middleware {
	return func(_ ssh.Handler) ssh.Handler {
		return func(s ssh.Session) {
			r.ServeSSH(s)
		}
	}
}

// ServeSSH handles a ssh.Session
func (r *usernameRouter) ServeSSH(s ssh.Session) {
	var sh SessionHandler
	var err error
	if sh, err = r.HandlerFor(s.User()); err != nil {
		wish.Fatalln(s, "this route does not exist\r")
		return
	}

	// the route is valid, so run the handler
	sh.ServeSSH(s)
}

// Handler returns the ssh.Handler function defined
// for this router. This method exists for compatibility and
// terminology consistency purposes, and simply returns r.ServeSSH.
func (r *usernameRouter) Handler() ssh.Handler {
	return r.ServeSSH
}
