package main

import (
	"github.com/charmbracelet/wish"
	"github.com/gliderlabs/ssh"
)

// MWPrintFoo prints a statement indicating that the user
// has reached it to the ssh.Session.
func MWPrintFoo(next ssh.Handler) ssh.Handler {
	return func(s ssh.Session) {
		wish.Println(s, "reached the FOO log middleware")
		next(s)
	}
}

// MWPrintBar prints a statement indicating that the user
// has reached it to the ssh.Session.
func MWPrintBar(next ssh.Handler) ssh.Handler {
	return func(s ssh.Session) {
		wish.Println(s, "reached the BAR log middleware")
		next(s)
	}
}

// MWPrintBaz prints a statement indicating that the user
// has reached it to the ssh.Session.
func MWPrintBaz(next ssh.Handler) ssh.Handler {
	return func(s ssh.Session) {
		wish.Println(s, "reached the BAZ log middleware")
		next(s)
	}
}

// MWPrintGlobal prints a statement indicating that the user
// has reached it to the ssh.Session.
func MWPrintGlobal(next ssh.Handler) ssh.Handler {
	return func(s ssh.Session) {
		wish.Println(s, "reached the GLOBAL log middleware")
		next(s)
	}
}
