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
	node.IncTerm()
	node.setState(NewFollower())
}

func (l *leader) Elect(node *Node) {}

func (l *leader) VoteRequest(node *Node, req server.VoteRequest) {
	if req.Term() > node.Term() {
		glog.Infof(node.logFormat("granted vote request with term %d"), req.Term())

		node.setState(NewFollower())
		node.VoteRequest(req)
	} else {
		req.Deny()

		glog.Infof(node.logFormat("denied vote request with term %d"), req.Term())
	}
}

func (l *leader) AppendEntriesRequest(node *Node) {
	// TODO(AD)
}

func (l *leader) name() string {
	return leaderName
}
