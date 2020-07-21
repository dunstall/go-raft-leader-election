package node

import (
	"github.com/dunstall/goraft/pkg/server"
)

const (
	initialTerm = 1
)

type Node struct {
	term uint32

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
		term:      initialTerm,
		follower:  follower,
		candidate: candidate,
		leader:    leader,
		state:     follower,
	}
}

func (n *Node) Term() uint32 {
	return n.term
}

func (n *Node) SetTerm(term uint32) {
	n.term = term
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
