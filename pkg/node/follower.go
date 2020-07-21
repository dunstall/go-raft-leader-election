package node

import (
	"github.com/dunstall/goraft/pkg/server"
  "github.com/golang/glog"
)

type Follower struct {
}

func NewFollower() nodeState {
	return &Follower{}
}

func (f *Follower) Expire(node *Node) {
	glog.Info("follower: node timed out")
	node.setState(node.candidateState())
}

func (f *Follower) Elect(node *Node) {
	glog.Warning("follower: cannot elect a follower")
}

func (f *Follower) VoteRequest(node *Node, cb server.Callback) {
	// TODO(AD)
	glog.Info("follower: received vote request")
}

func (f *Follower) AppendEntriesRequest(node *Node) {
	// TODO(AD)
	glog.Info("follower: received append entries request")
}
