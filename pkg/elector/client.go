package elector

type Client interface {
	// Dial creates a connection to the server at the given address. Note this
	// connection in the background.
	// TODO(AD) Connection must handle reconnect (grpc can do this)
	Dial(addr string) Connection
}
