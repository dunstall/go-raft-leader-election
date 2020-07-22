package node

import (
	"github.com/dunstall/goraft/pkg/server"
	"github.com/golang/glog"
)

type candidate struct{}

func NewCandidate() nodeState {
	return &candidate{}
}

func (c *candidate) Expire(node *Node) {
	glog.Info(node.logFormat("node timed out"))
	node.SetTerm(node.Term() + 1)
	node.Elector().Elect(node.Term())
}

func (c *candidate) Elect(node *Node) {
	glog.Warning(node.logFormat("node elected"))
	node.setState(node.leaderState())
}

func (c *candidate) VoteRequest(node *Node, req server.VoteRequest) {
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

func (c *candidate) AppendEntriesRequest(node *Node) {
	// TODO(AD)
	glog.Infof(node.logFormat("received append entries request"))
}

func (c *candidate) name() string {
	return "candidate"
}
