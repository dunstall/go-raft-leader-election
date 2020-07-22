package node

import (
	"fmt"

	"github.com/dunstall/goraft/pkg/elector"
	"github.com/dunstall/goraft/pkg/server"
)

const (
	initialTerm = 1
)

type Node struct {
	id   uint32
	term uint32

	follower  nodeState
	candidate nodeState
	leader    nodeState

	state nodeState

	elector elector.Elector
}

func NewNode(id uint32, elector elector.Elector) Node {
	follower := NewFollower()
	candidate := NewCandidate()
	leader := NewLeader()

	return Node{
		id:        id,
		term:      initialTerm,
		follower:  follower,
		candidate: candidate,
		leader:    leader,
		state:     follower,
		elector:   elector,
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

func (n *Node) VoteRequest(cb server.VoteRequest) {
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

func (n *Node) Elector() elector.Elector {
	return n.elector
}

func (n *Node) setState(state nodeState) {
	n.state = state
}

func (n *Node) logFormat(msg string) string {
	return fmt.Sprintf("%s %d %d: %s", n.state.name(), n.id, n.Term(), msg)
}
