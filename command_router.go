package icepop

import (
	"fmt"

	"github.com/charmbracelet/wish"
	"github.com/gliderlabs/ssh"
)

// Ensure CommandRouter implements SessionHandler.
var _ SessionHandler = &commandRouter{}

// NewCommandRouter returns a ssh.Handler that routes to registered
// SessionHandlers based on the ssh.Session.Command value. Only the
// entrypoint, or the first entry of a given command is parsed.
// All other parameters are ignored.
//
// If the command is not registered, then a configurable message
// is printed to the user and the session is closed.
func NewCommandRouter(opts ...CommandRouterOption) *commandRouter {
	r := &commandRouter{
		Router: Router{
			routes: map[string]SessionHandler{},
		},
		notFoundMsg:    "not found",
		notProvidedMsg: "you need to provide a command",
	}

	// apply options
	for _, opt := range opts {
		opt(r)
	}

	return r
}

type commandRouter struct {
	Router

	// notFoundMsg is the message sent to the user
	// when the command they provided is not registered
	// as a handler.
	notFoundMsg string

	// notProvidedMsg is the message sent to the user
	// when they don't provide a command.
	notProvidedMsg string
}

// AsMiddleware returns the entire router handler chain
// as a Middleware. This must be the final handler executed
// as it does not pass control to handlers provided as input.
func (r *commandRouter) AsMiddleware() wish.Middleware {
	return func(_ ssh.Handler) ssh.Handler {
		return func(s ssh.Session) {
			r.ServeSSH(s)
		}
	}
}

// ServeSSH handles a ssh.Session
func (r *commandRouter) ServeSSH(s ssh.Session) {
	if len(s.Command()) == 0 {
		wish.Fatalln(s, fmt.Sprintf("%s\r", r.notProvidedMsg))
		return
	}

	var sh SessionHandler
	var err error
	if sh, err = r.HandlerFor(s.Command()[0]); err != nil {
		wish.Fatalln(s, fmt.Sprintf("%s\r", r.notFoundMsg))
		return
	}

	sh.ServeSSH(s)
}

// Handler returns the ssh.Handler function defined
// for this router. This method exists for compatibility and
// terminology consistency purposes, and simply returns r.ServeSSH.
func (r *commandRouter) Handler() ssh.Handler {
	return r.ServeSSH
}
