package node

import (
	"github.com/dunstall/goraft/pkg/server"
	"github.com/golang/glog"
)

type leader struct{}

func NewLeader() nodeState {
	return &leader{}
}

func (l *leader) Expire(node *Node) {
	glog.Info(node.logFormat("node timed out"))
	node.SetTerm(node.Term() + 1)
	node.setState(node.followerState())
}

func (l *leader) Elect(node *Node) {
	glog.Warning(node.logFormat("node already elected"))
}

func (l *leader) VoteRequest(node *Node, req server.VoteRequest) {
	glog.Infof(node.logFormat("received vote request from candidate %d"), req.CandidateID())

	if req.Term() > node.Term() {
		glog.Infof(node.logFormat("granted vote request with term %d"), req.Term())

		node.setState(node.followerState())
		node.VoteRequest(req)
	} else {
		req.Deny()

		glog.Infof(node.logFormat("denied vote request with term %d"), req.Term())
	}
}

func (l *leader) AppendEntriesRequest(node *Node) {
	// TODO(AD)
	glog.Info(node.logFormat("received append entries request"))
}

func (l *leader) name() string {
	return "leader"
}
