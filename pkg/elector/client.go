package elector

type Client interface {
	// Dial creates a connection to the server at the given address. Note this
	// connection in the background.
	Dial(addr string) Connection
}
