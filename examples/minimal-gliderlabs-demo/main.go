package main

import (
	"io"
	"log"

	"github.com/concerthall/icepop"
	"github.com/gliderlabs/ssh"
)

func main() {
	rtr := icepop.NewCommandRouter()
	rtr.HandleFunc(
		"favorite-flavor",
		func(s ssh.Session) {
			io.WriteString(s, "I like cherry!")
			s.Exit(0)
		})

	ssh.Handle(rtr.Handler())

	log.Println("starting the gliderlabs demo")
	log.Fatal(ssh.ListenAndServe(":23234", nil))
}
