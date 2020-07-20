package node

type nodeState interface {
	Expire(node *Node)
	Elect(node *Node)
	// TODO(AD) Use callback with resp channel
	ReceiveVoteRequest(node *Node)
	// TODO(AD) Use callback with resp channel
	ReceiveAppendEntriesRequest(node *Node)
}
