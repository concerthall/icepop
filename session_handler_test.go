package icepop

import (
	"testing"

	"github.com/gliderlabs/ssh"
	"github.com/stretchr/testify/assert"
)

func TestNewSessionHandlerCreation(t *testing.T) {
	h := func(s ssh.Session) {
		s.Exit(0)
	}
	x := NewSessionHandlerFrom(h)
	assert.ObjectsAreEqual(h, x)
}
