# Basic Routing Demo

This demo server shows how a single server can serve different routes each with
unique middleware layers. This reuses the Wish server but allows the Command and
Username routers to act as middleware that then hands routing over to them
exclusively.
