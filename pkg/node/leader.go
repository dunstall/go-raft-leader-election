package node

type leader struct {
}

func NewLeader() nodeState {
	return &leader{}
}

func (l *leader) Expire(node *Node) {
	// TODO(AD)
}

func (l *leader) Elect(node *Node) {
	// TODO(AD)
}

func (l *leader) ReceiveVoteRequest(node *Node) {
	// TODO(AD)
}

func (l *leader) ReceiveAppendEntriesRequest(node *Node) {
	// TODO(AD)
}
