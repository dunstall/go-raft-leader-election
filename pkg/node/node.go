package node

import (
	"github.com/dunstall/goraft/pkg/server"
)

type Node struct {
	follower  nodeState
	candidate nodeState
	leader    nodeState

	state nodeState
}

func NewNode() Node {
	follower := NewFollower()
	candidate := NewCandidate()
	leader := NewLeader()

	return Node{
		follower:  follower,
		candidate: candidate,
		leader:    leader,
		state:     follower,
	}
}

func (n *Node) Expire() {
	n.state.Expire(n)
}

func (n *Node) Elect() {
	n.state.Elect(n)
}

func (n *Node) VoteRequest(cb server.Callback) {
	n.state.VoteRequest(n, cb)
}

func (n *Node) AppendEntriesRequest() {
	n.state.AppendEntriesRequest(n)
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
