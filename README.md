# Icepop

<img alt="icepop" src="icepop.png" width="200em" />


An SSH routing library compatible with
[charmbracelet/wish](https://github.com/charmbracelet/wish) and
[gliderlabs/ssh](https://github.com/gliderlabs/ssh)!

Extend your SSH applications by routing connections to unique handlers and
middleware based on various characteristics such as the login username, or
entrypoint.

## Compatible with [charmbracelet/wish](https://github.com/charmbracelet/wish)

Utilizing an Icepop Router with Wish is super simple. Icepop routers implement
`ssh.Middleware`, which means that it fits right into any Wish application.

```go
// source: ./examples/minimal-wish-demo/main.go
package main

import (
	"fmt"

	"github.com/charmbracelet/wish"
	"github.com/gliderlabs/ssh"

	"github.com/concerthall/icepop"
)

func main() {
	rtr := icepop.NewUsernameRouter()
	rtr.HandleFunc(
		// the expected username
		"itshotoutside", 
		// The handler to be used.
		func(s ssh.Session) {
			wish.Println(s, "I love Ice pops!")
			s.Exit(0)
			s.Close()
		}),
	)

	s, _ := wish.NewServer(
		wish.WithAddress("localhost:23234"),
		wish.WithHostKeyPath(".ssh/term_info_ed25519"),
		// Here's our router! 
		wish.WithMiddleware(rtr.Middleware),
	)

	fmt.Println("Listening!")
	if err := s.ListenAndServe(); err != nil {
		panic(err) // handle how you see fit!
	}
}
```

And then SSH to the demo server!

```
$ ssh itshotoutside@localhost -p 23234
I love Ice pops!
Connection to localhost closed.
```

Unregistered routes are handled automatically!

```
$ ssh itscoldoutside@localhost -p 23234
this route does not exist
Connection to localhost closed.
```

## Compatible with [gliderlab/ssh](https://github.com/gliderlabs/ssh)

Because Wish extends the Gliderlabs SSH library, Icepop also fits nicely with the gliderlabs/ssh library.

```go
// source: ./examples/minimal-gliderlabs-demo/main.go
package main

import (
	"io"
	"log"

	"github.com/concerthall/icepop"
	"github.com/gliderlabs/ssh"
)

func main() {
	rtr := icepop.NewCommandRouter()
	rtr.HandleFunc(
		"favorite-flavor",
		func(s ssh.Session) {
			io.WriteString(s, "I like cherry!")
			s.Exit(0)
		})

	ssh.Handle(rtr.Handler())

	log.Println("starting the gliderlabs demo")
	log.Fatal(ssh.ListenAndServe(":23234", nil))
}
```

Output:

```
$ ssh localhost -p 23234 favorite-flavor
I like cherry!
```

## Concepts & Constructs

### Routers

Routers are just `ssh.Handlers` that perform some kind of logical routing based on some factor.

This library provides:

- a **Command Router** which allows you to bind different handlers, or routes, to each registered command the user passes into their session.

- A **Username Router** which allows you to bind different handlers, or routes, to each registered username the user passes into their session.

### SessionHandlers

A Session Handler is an abstraction based on the `ssh.Handler` definition. It binds a `ssh.Handler` to a method `ServeSSH`, as you might find in the [net/http](https://pkg.go.dev/net/http#Handler) package.

For compatibility, you can convert an existing `ssh.Handler` to a `SessionHandler` using `icepop.NewSessionHandlerFrom`. A note: the `SessionHandler` interface doesn't replace `ssh.Handler`, but simply exists to allow us to do some creative things!

### Routes

Routes are components that represent a base handler along with its middleware. Routes can also be treated as `ssh.Handler`s by themselves.

## Examples

Several annotated examples exist in the [examples](./examples/) directory, along with those already defined above.

Take a look at our [nested demo](./examples/nested-demo/main.go), where we take advantage of the flexibility the `SessionHandler` interface affords us by nesting a command router within a specific endpoint of a username router.

## Docs

The source is commented throughout, and feel free to open an issue if anything is unclear.

[Go Doc](https://pkg.go.dev/github.com/concerthall/icepop)
## Attribution

Image Credit: "Ice cream stickers" by [Gohsantosadrive - Flaticon](https://www.flaticon.com/free-stickers/ice-cream)
