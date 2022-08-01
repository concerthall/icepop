package icepop

// RouteOption configures a Route
type RouteOption func(*route)

// OptExecuteMiddlewareInReverse indicates that a Route should
// execute middleware in reverse order relative to how it was
// defined in source code.
//
// This is a compatibility option for developers who are familiar
// with how charmbraclete/wish handles middleware.
func OptExecuteMiddlewareInReverse() RouteOption {
	return func(r *route) {
		r.reversedMiddlewareOrder = true
	}
}
