package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/charmbracelet/wish"
	wishlogging "github.com/charmbracelet/wish/logging"
	"github.com/gliderlabs/ssh"

	"github.com/concerthall/icepop"
)

const (
	host = "localhost"
	port = 23234
)

func main() {
	// Set up the router. We'll use the username router for this
	// demo.
	rtr := icepop.NewUsernameRouter()

	// Register our two routes.
	rtr.Handle("register", RegisterRoute)
	rtr.Handle("stream", StreamRoute)

	s, err := wish.NewServer(
		wish.WithAddress(fmt.Sprintf("%s:%d", host, port)),
		wish.WithHostKeyPath(".ssh/term_info_ed25519"),
		wish.WithMiddleware(
			// Add our router as a middleware. This is to comply
			// with the Wish architecture. It must be the top-most
			// definition.
			rtr.AsMiddleware(),
			// These middleware will run before the router is ever
			// called.
			MWPrintGlobal,
			wishlogging.Middleware(),
		),
	)

	if err != nil {
		log.Fatalln(err)
	}

	run(s)
}

func run(s *ssh.Server) {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	log.Printf("Starting SSH server on %s:%d", host, port)
	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Fatalln(err)
		}
	}()

	<-done
	log.Println("Stopping SSH server")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer func() { cancel() }()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatalln(err)
	}
}
