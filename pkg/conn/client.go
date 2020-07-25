package conn

// Client handles connecting to other nodes.
type Client interface {
	// Dial creates a connection to the node at the given address. Note this
	// connection created in the background.
	Dial(addr string) Connection
}
