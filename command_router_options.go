package icepop

// CommandRouterOption configures a CommandRouter
type CommandRouterOption func(r *commandRouter)

// OptCommandNotFoundMessage configures the message sent to the user
// when the command provided was not registered.
func OptCommandNotFoundMessage(msg string) CommandRouterOption {
	return func(r *commandRouter) {
		r.notFoundMsg = msg
	}
}

// OptCommandNotProvidedMessage configures the message sent to the user
// when no command was provided.
func OptCommandNotProvidedMessage(msg string) CommandRouterOption {
	return func(r *commandRouter) {
		r.notProvidedMsg = msg
	}
}
