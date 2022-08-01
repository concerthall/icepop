package icepop

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommandRouterCustomNotFoundMsg(t *testing.T) {
	r := NewCommandRouter(OptCommandNotFoundMessage("foo"))

	assert.Equal(t, "foo", r.notFoundMsg, "the not found message should be set")
}

func TestCommandRouterCustomNotProvideMsg(t *testing.T) {
	r := NewCommandRouter(OptCommandNotProvidedMessage("bar"))

	assert.Equal(t, "bar", r.notProvidedMsg, "the not provided message should be set")
}

func TestCommandRouterNoOptions(t *testing.T) {
	r := NewCommandRouter()

	assert.Equal(t, "not found", r.notFoundMsg, "the not found message should be the default")
	assert.Equal(t, "you need to provide a command", r.notProvidedMsg, "the not provided message should be the default")
}
