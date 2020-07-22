package node

import (
	"fmt"

	"github.com/dunstall/goraft/pkg/elector"
	"github.com/dunstall/goraft/pkg/server"
	"github.com/golang/glog"
)

const (
	initialTerm = 1
)

type Node struct {
	id   uint32
	term uint32

	state nodeState

	elector elector.Elector
}

func NewNode(id uint32, elector elector.Elector) Node {
	return Node{
		id:      id,
		term:    initialTerm,
		state:   NewFollower(),
		elector: elector,
	}
}

func (n *Node) Term() uint32 {
	return n.term
}

func (n *Node) SetTerm(term uint32) {
	n.term = term
}

func (n *Node) IncTerm() {
	n.term++
}

func (n *Node) Expire() {
	glog.Info(n.logFormat("node expired"))
	n.state.Expire(n)
}

func (n *Node) Elect() {
	glog.Info(n.logFormat("node elected"))
	n.state.Elect(n)
}

func (n *Node) VoteRequest(req server.VoteRequest) {
	glog.Infof(n.logFormat("received vote request from candidate %d"), req.CandidateID())
	n.state.VoteRequest(n, req)
}

func (n *Node) AppendEntriesRequest() {
	glog.Info(n.logFormat("received append entries request"))
	n.state.AppendEntriesRequest(n)
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
