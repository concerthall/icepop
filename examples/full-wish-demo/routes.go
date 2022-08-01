package main

import (
	"github.com/charmbracelet/wish"
	"github.com/concerthall/icepop"
	"github.com/gliderlabs/ssh"
)

var (
	RegisterRoute = icepop.NewRoute(
		registerHandler,
		[]wish.Middleware{
			MWPrintFoo,
			MWPrintBar,
		},
	)

	StreamRoute = icepop.NewRoute(
		streamHandler,
		[]wish.Middleware{
			MWPrintFoo,
			MWPrintBar,
			MWPrintBaz,
		},
		// This executes the middleware in reverse order compared
		// to how they're provided above, similar to Wish.
		icepop.OptExecuteMiddlewareInReverse(),
	)
)

func registerHandler(s ssh.Session) {
	wish.Println(s, "You've reached the register Handler")
}

func streamHandler(s ssh.Session) {
	wish.Println(s, "You've reached the stream Handler")
}
