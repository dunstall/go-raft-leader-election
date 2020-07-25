package node

import (
	"github.com/dunstall/goraft/pkg/server"
	"github.com/golang/glog"
)

const (
	followerName = "follower"
)

type follower struct{}

func NewFollower() nodeState {
	return &follower{}
}

func (f *follower) Expire(node *Node) {
	node.IncTerm()
	node.setState(NewCandidate())
	node.Elector().Elect(node.Term())
}

func (f *follower) Elect(node *Node) {}

func (f *follower) VoteRequest(node *Node, req server.VoteRequest) {
	if req.Term() > node.Term() {
		f.grantVoteRequest(node, req)
	} else {
		f.denyVoteRequest(node, req)
	}
}

func (f *follower) AppendRequest(node *Node) {
	// TODO(AD)
}

func (f *follower) name() string {
	return followerName
}

func (f *follower) grantVoteRequest(node *Node, req server.VoteRequest) {
	node.SetTerm(req.Term())
	req.Grant()

	glog.Infof(node.logFormat("granted vote request with term %d"), req.Term())
}

func (f *follower) denyVoteRequest(node *Node, req server.VoteRequest) {
	req.Deny()

	glog.Infof(node.logFormat("denied vote request with term %d"), req.Term())
}
