package main

import (
	"github.com/charmbracelet/wish"
	"github.com/gliderlabs/ssh"
)

func MWPrintFoo(h ssh.Handler) ssh.Handler {
	return func(s ssh.Session) {
		wish.Println(s, "reached the FOO log middleware")
		h(s)
	}
}

func MWPrintBar(h ssh.Handler) ssh.Handler {
	return func(s ssh.Session) {
		wish.Println(s, "reached the BAR log middleware")
		h(s)
	}
}

func MWPrintBaz(h ssh.Handler) ssh.Handler {
	return func(s ssh.Session) {
		wish.Println(s, "reached the BAZ log middleware")
		h(s)
	}
}

func MWPrintOne(h ssh.Handler) ssh.Handler {
	return func(s ssh.Session) {
		wish.Println(s, "reached the ONE log middleware")
		h(s)
	}
}

func MWPrintTwo(h ssh.Handler) ssh.Handler {
	return func(s ssh.Session) {
		wish.Println(s, "reached the TWO log middleware")
		h(s)
	}
}
