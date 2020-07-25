package node

import (
	"time"

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
	node.Heartbeat().Beat(node.Term())
}

func (l *leader) Elect(node *Node) {}

func (l *leader) Timeout() time.Duration {
	// TODO(AD) RTT
	return 50 * time.Millisecond * 10
}

func (l *leader) VoteRequest(node *Node, req server.VoteRequest) {
	if req.Term() > node.Term() {
		l.grantVoteRequest(node, req)
	} else {
		l.denyVoteRequest(node, req)
	}
}

func (l *leader) AppendRequest(node *Node, req server.AppendRequest) {
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
