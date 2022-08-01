package icepop

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUsernameRouterCustomNotFoundMsg(t *testing.T) {
	r := NewUsernameRouter(OptUsernameNotFoundMessage("foo"))

	assert.Equal(t, "foo", r.notFoundMsg, "the not found message should be set")
}

func TestUsernameRouterCustomNotProvideMsg(t *testing.T) {
	r := NewUsernameRouter(OptUsernameNotProvidedMessage("bar"))

	assert.Equal(t, "bar", r.notProvidedMsg, "the not provided message should be set")
}

func TestUsernameRouterNoOptions(t *testing.T) {
	r := NewUsernameRouter()

	assert.Equal(t, "not found", r.notFoundMsg, "the not found message should be the default")
	assert.Equal(t, "you need to provide a username", r.notProvidedMsg, "the not provided message should be the default")
}
