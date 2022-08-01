package main

import (
	"fmt"

	"github.com/charmbracelet/wish"
	"github.com/gliderlabs/ssh"

	"github.com/concerthall/icepop"
)

func main() {
	// Registered handlers check for the username!
	rtr := icepop.NewUsernameRouter()
	// Our first handler!
	rtr.Handle(
		// the expected username
		"itshotoutside",
		// The handler to be used.
		icepop.NewSessionHandlerFrom(func(s ssh.Session) {
			wish.Println(s, "I love Ice pops!")
			s.Exit(0)
			s.Close()
		}),
	)

	// Create the Wish server
	s, _ := wish.NewServer(
		wish.WithAddress("localhost:23234"),
		wish.WithHostKeyPath(".ssh/term_info_ed25519"),
		// Here's our router!
		wish.WithMiddleware(rtr.AsMiddleware()),
	)

	// Start the server!
	fmt.Println("Listening!")
	if err := s.ListenAndServe(); err != nil {
		panic(err) // handle how you see fit!
	}
}
