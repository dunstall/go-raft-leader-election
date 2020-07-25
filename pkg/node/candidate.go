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
		c.grantVoteRequest(node, req)
	} else {
		c.denyVoteRequest(node, req)
	}
}

func (c *candidate) AppendRequest(node *Node, req server.AppendRequest) {
	// TODO(AD)
}

func (c *candidate) name() string {
	return candidateName
}

func (c *candidate) grantVoteRequest(node *Node, req server.VoteRequest) {
	glog.Infof(node.logFormat("reverting to follower"))

	node.setState(NewFollower())
	node.VoteRequest(req)
}

func (c *candidate) denyVoteRequest(node *Node, req server.VoteRequest) {
	req.Deny()

	glog.Infof(node.logFormat("denied vote request with term %d"), req.Term())
}
