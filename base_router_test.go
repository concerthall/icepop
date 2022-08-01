package icepop

import (
	"testing"

	"github.com/gliderlabs/ssh"
	"github.com/stretchr/testify/assert"
)

func TestBaseRouterSuccessfulPathRegistration(t *testing.T) {
	r := Router{map[string]SessionHandler{}}
	expected := NewSessionHandlerFrom(func(s ssh.Session) { s.Close() })
	r.Handle("foo", expected)

	actual, err := r.HandlerFor("foo")
	assert.Equal(t, expected, actual, "bound handler should be accessible at input path")
	assert.Equal(t, nil, err, "should have no errors")
	assert.ObjectsAreEqual(expected.ServeSSH, actual.ServeSSH)

}

func TestBaseRouterUnregisteredPath(t *testing.T) {
	r := Router{map[string]SessionHandler{}}

	_, err := r.HandlerFor("bar")
	assert.NotNil(t, err, "unregistered routes should return an error")
}
