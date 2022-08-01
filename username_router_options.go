package icepop

// UsernameRouterOption configures a CommandRouter
type UsernameRouterOption func(r *usernameRouter)

// OptUsernameNotFoundMessage configures the message sent to the user
// when the username provided was not registered.
func OptUsernameNotFoundMessage(msg string) UsernameRouterOption {
	return func(r *usernameRouter) {
		r.notFoundMsg = msg
	}
}

// OptUsernameNotProvidedMessage configures the message sent to the user
// when no command was provided.
func OptUsernameNotProvidedMessage(msg string) UsernameRouterOption {
	return func(r *usernameRouter) {
		r.notProvidedMsg = msg
	}
}
