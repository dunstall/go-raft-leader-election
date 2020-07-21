package node

import (
	"github.com/dunstall/goraft/pkg/server"
	"github.com/golang/glog"
)

type follower struct{}

func NewFollower() nodeState {
	return &follower{}
}

func (f *follower) Expire(node *Node) {
	// TODO(AD) Logs should have format <state> <node ID>: msg. Can make this
	// a method of node.
	glog.Infof("follower: node timed out in term %d", node.Term())
	node.SetTerm(node.Term() + 1)
	node.setState(node.candidateState())
}

func (f *follower) Elect(node *Node) {
	glog.Warning("follower: cannot elect a follower")
}

func (f *follower) VoteRequest(node *Node, req server.VoteRequest) {
	glog.Infof("follower: received vote request from candidate %d", req.CandidateID())

	if req.Term() > node.Term() {
		node.SetTerm(req.Term())
		req.Grant()

		glog.Infof("follower: granted vote request with term %d", req.Term())
	} else {
		req.Deny()

		glog.Infof("follower: denied vote request with term %d", req.Term())
	}
}

func (f *follower) AppendEntriesRequest(node *Node) {
	// TODO(AD)
	glog.Info("follower: received append entries request")
}
