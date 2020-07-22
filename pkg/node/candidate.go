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
	glog.Infof("candidate: node timed out in term %d", node.Term())
	node.SetTerm(node.Term() + 1)
	node.Elector().Elect(node.Term())
}

func (c *candidate) Elect(node *Node) {
	glog.Infof("candidate: node elected in term %d", node.Term())
	node.setState(node.leaderState())
}

func (c *candidate) VoteRequest(node *Node, req server.VoteRequest) {
	glog.Infof("candidate: received vote request from candidate %d", req.CandidateID())

	if req.Term() > node.Term() {
		glog.Infof("candidate: vote request has greater term %d - reverting to follower", req.Term())

		node.setState(node.followerState())
		node.VoteRequest(req)
	} else {
		req.Deny()

		glog.Infof("candidate: denied vote request with term %d", req.Term())
	}
}

func (c *candidate) AppendEntriesRequest(node *Node) {
	// TODO(AD)
	glog.Info("candidate: received append entries request")
}
