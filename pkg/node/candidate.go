package node

type candidate struct {
}

func NewCandidate() nodeState {
	return &candidate{}
}

func (c *candidate) Expire(node *Node) {
	// TODO(AD) -> Candidate
}

func (c *candidate) Elect(node *Node) {
	// TODO(AD) -> Leader
}

func (c *candidate) ReceiveVoteRequest(node *Node) {
	// TODO(AD)
}

func (c *candidate) ReceiveAppendEntriesRequest(node *Node) {
	// TODO(AD)
}
