package node

import (
	"github.com/dunstall/goraft/pkg/server"
  "github.com/golang/glog"
)

type follower struct {
}

func NewFollower() nodeState {
	return &follower{}
}

func (f *follower) Expire(node *Node) {
	glog.Info("follower: node timed out")
	node.setState(node.candidateState())
}

func (f *follower) Elect(node *Node) {
	glog.Warning("follower: cannot elect a follower")
}

func (f *follower) VoteRequest(node *Node, cb server.Callback) {
	// TODO(AD)
	glog.Info("follower: received vote request")
}

func (f *follower) AppendEntriesRequest(node *Node) {
	// TODO(AD)
	glog.Info("follower: received append entries request")
}
