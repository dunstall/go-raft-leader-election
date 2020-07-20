package node

type Node struct {
	follower  nodeState
	candidate nodeState
	leader    nodeState

	state nodeState
}

func NewNode() Node {
	node := Node{}

	node.follower = NewFollower(node)
	node.candidate = NewCandidate(node)
	node.leader = NewLeader(node)

	node.state = node.follower

	return node
}

func (n *Node) Expire() {
	n.state.Expire()
}

func (n *Node) Elect() {
	n.state.Elect()
}

func (n *Node) ReceiveVoteRequest() {
	n.state.ReceiveVoteRequest()
}

func (n *Node) ReceiveAppendEntriesRequest() {
	n.state.ReceiveAppendEntriesRequest()
}

func (n *Node) followerState() nodeState {
	return n.follower
}

func (n *Node) candidateState() nodeState {
	return n.candidate
}

func (n *Node) leaderState() nodeState {
	return n.leader
}

func (n *Node) setState(state nodeState) {
	n.state = state
}
