package conn

// Connection represents a connection to a node.
type Connection interface {
	RequestVote(term uint32) bool
	Close()
}
