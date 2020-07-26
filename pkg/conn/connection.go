package conn

// Connection represents a connection to a node.
//
// If the underlying transport connection is lost this will reconnect
// automatically so theres no need to create a new connection to reconnect.
type Connection interface {
	// RequestVote requests a vote in the given term to the connected node.
	// Returns true if the vote was granted, false otherwise. Note failing to
	// reach the node is treated as a denied request so returns false.
	RequestVote(term uint32) bool

	// RequestAppend sends a heartbeat in the given term to the connected node.
	// Returns true if the request was accepted, false otherwise. Note failing to
	// connect to the node is treated as a failure so returns false.
	// TODO(AD) This is only used as heartbeat for now.
	RequestAppend(term uint32) bool

	// Close closes the connection to the node.
	Close()
}
