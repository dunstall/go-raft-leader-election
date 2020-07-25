package node

import (
	"fmt"
	"time"

	"github.com/dunstall/goraft/pkg/elector"
	"github.com/dunstall/goraft/pkg/heartbeat"
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

	elector   elector.Elector
	heartbeat heartbeat.Heartbeat
}

func NewNode(id uint32, elector elector.Elector, heartbeat heartbeat.Heartbeat) Node {
	return Node{
		id:        id,
		term:      initialTerm,
		state:     NewFollower(),
		elector:   elector,
		heartbeat: heartbeat,
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

func (n *Node) Timeout() time.Duration {
	return n.state.Timeout()
}

func (n *Node) VoteRequest(req server.VoteRequest) {
	glog.Infof(n.logFormat("received vote request from candidate %d"), req.CandidateID())
	n.state.VoteRequest(n, req)
}

func (n *Node) AppendRequest(req server.AppendRequest) {
	glog.Info(n.logFormat("received append entries request"))
	n.state.AppendRequest(n, req)
}

func (n *Node) Elector() elector.Elector {
	return n.elector
}

func (n *Node) Heartbeat() heartbeat.Heartbeat {
	return n.heartbeat
}

func (n *Node) setState(state nodeState) {
	n.state = state
}

func (n *Node) logFormat(msg string) string {
	return fmt.Sprintf("%s %d %d: %s", n.state.name(), n.id, n.Term(), msg)
}
