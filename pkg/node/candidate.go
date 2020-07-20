package node

type candidate struct {
	node Node
}

func NewCandidate(node Node) nodeState {
	return &candidate{node: node}
}

func (c *candidate) Expire() {
	// TODO(AD) -> Candidate
}

func (c *candidate) Elect() {
	// TODO(AD) -> Leader
}

func (c *candidate) ReceiveVoteRequest() {
	// TODO(AD)
}

func (c *candidate) ReceiveAppendEntriesRequest() {
	// TODO(AD)
}
