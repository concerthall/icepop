package icepop

import (
	"errors"

	"github.com/gliderlabs/ssh"
)

// Router is basic Router structure. It can be embedded in custom
// routers to provide easy access to the Handle method. Custom
// routers will need to implement methods that use these routes.
type Router struct {
	routes map[string]SessionHandler
}

// Handle binds a path to a SessionHandler. If a handler already
// exists, this panics.
func (r *Router) Handle(path string, sh SessionHandler) {
	if _, exists := r.routes[path]; exists {
		panic("Attempted to rebind a pre-existing handler: %s")
	}

	r.routes[path] = sh
}

// HandleFunc binds a ssh.Handler to a SessionHandler and binds it to path.
// If a handler already exists, this panics.
func (r *Router) HandleFunc(path string, hf ssh.Handler) {
	if _, exists := r.routes[path]; exists {
		panic("Attempted to rebind a pre-existing handler: %s")
	}

	r.routes[path] = NewSessionHandlerFrom(hf)
}

// HandlerFor returns the handler for the given path value, or an error
// if it does not exist.
func (r *Router) HandlerFor(path string) (SessionHandler, error) {
	if _, exists := r.routes[path]; !exists {
		return nil, errors.New("unregistered route")
	}

	return r.routes[path], nil
}
