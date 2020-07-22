package node

import (
	"github.com/dunstall/goraft/pkg/server"
	"github.com/golang/glog"
)

const (
	candidateName = "candidate"
)

type candidate struct{}

func NewCandidate() nodeState {
	return &candidate{}
}

func (c *candidate) Expire(node *Node) {
	node.IncTerm()
	node.Elector().Elect(node.Term())
}

func (c *candidate) Elect(node *Node) {
	node.setState(NewLeader())
}

func (c *candidate) VoteRequest(node *Node, req server.VoteRequest) {
	if req.Term() > node.Term() {
		glog.Infof(node.logFormat("reverting to follower"))

		node.setState(NewFollower())
		node.VoteRequest(req)
	} else {
		req.Deny()

		glog.Infof(node.logFormat("denied vote request with term %d"), req.Term())
	}
}

func (c *candidate) AppendEntriesRequest(node *Node) {
	// TODO(AD)
}

func (c *candidate) name() string {
	return candidateName
}
