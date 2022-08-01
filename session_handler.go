package icepop

import "github.com/gliderlabs/ssh"

// ensure sessionHandler always implements SessionHandler
var _ SessionHandler = &sessionHandler{}

// SessionHandler ∂escribes a ssh.Handler implementation bound to
// a fixed/known method name.
type SessionHandler interface {
	ServeSSH(ssh.Session)
}

// NewSessionHandlerFrom returns a SessionHandler from an existing Handler.
// This has a narrow scope of responsibility, aiming to remain compatible
// with existing ssh.Handler definitions.
func NewSessionHandlerFrom(handler ssh.Handler) SessionHandler {
	return &sessionHandler{handle: handler}
}

type sessionHandler struct {
	handle ssh.Handler
}

func (sh *sessionHandler) ServeSSH(s ssh.Session) {
	sh.handle(s)
}
