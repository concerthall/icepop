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

	HelloCommandRoute = icepop.NewRoute(
		helloHandler,
		[]wish.Middleware{
			MWPrintOne,
			MWPrintTwo,
		},
	)

	GoodbyeCommandRoute = icepop.NewRoute(
		goodbyeHandler,
		[]wish.Middleware{
			MWPrintOne,
			MWPrintTwo,
		},
		icepop.OptExecuteMiddlewareInReverse(),
	)
)

func registerHandler(s ssh.Session) {
	wish.Println(s, "You've reached the register Handler")
}

func streamHandler(s ssh.Session) {
	wish.Println(s, "You've reached the stream Handler")
}

func helloHandler(s ssh.Session) {
	wish.Println(s, "You've reached the hello Handler")
}

func goodbyeHandler(s ssh.Session) {
	wish.Println(s, "You've reached the goodbye Handler")
}
