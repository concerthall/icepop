# Nested Routing Demo

This demo server shows that you can also nest another router at a specific route
of an existing server. In this demo, we take advantage of a the SessionHandler
interface, which allows us to bind any type that implements the `ServeSSH`
method. The `ServeSSH` method is designed to match the
[ssh.Handler](https://pkg.go.dev/github.com/gliderlabs/ssh#Handler) function
definition.