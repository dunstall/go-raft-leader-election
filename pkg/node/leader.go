package node

import (
	"github.com/dunstall/goraft/pkg/server"
	"github.com/golang/glog"
)

const (
	leaderName = "leader"
)

type leader struct{}

func NewLeader() nodeState {
	return &leader{}
}

func (l *leader) Expire(node *Node) {
	// TODO(AD) Send heartbeat - leader cannot timeout?
	node.IncTerm()
	node.setState(NewFollower())
}

func (l *leader) Elect(node *Node) {}

func (l *leader) VoteRequest(node *Node, req server.VoteRequest) {
	if req.Term() > node.Term() {
		l.grantVoteRequest(node, req)
	} else {
		l.denyVoteRequest(node, req)
	}
}

func (l *leader) AppendRequest(node *Node) {
	// TODO(AD)
}

func (l *leader) name() string {
	return leaderName
}

func (l *leader) grantVoteRequest(node *Node, req server.VoteRequest) {
	glog.Infof(node.logFormat("granted vote request with term %d"), req.Term())

	node.setState(NewFollower())
	node.VoteRequest(req)
}

func (l *leader) denyVoteRequest(node *Node, req server.VoteRequest) {
	req.Deny()

	glog.Infof(node.logFormat("denied vote request with term %d"), req.Term())
}
