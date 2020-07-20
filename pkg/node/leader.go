package node

type leader struct {
	node Node
}

func NewLeader(node Node) nodeState {
	return &leader{node: node}
}

func (l *leader) Expire() {
	// TODO(AD)
}

func (l *leader) Elect() {
	// TODO(AD)
}

func (l *leader) ReceiveVoteRequest() {
	// TODO(AD)
}

func (l *leader) ReceiveAppendEntriesRequest() {
	// TODO(AD)
}
