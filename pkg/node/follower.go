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
	glog.Info(node.logFormat("node timed out"))
	node.SetTerm(node.Term() + 1)
	node.setState(node.candidateState())
	node.Elector().Elect(node.Term())
}

func (f *follower) Elect(node *Node) {
	glog.Warning(node.logFormat("cannot elect a follower"))
}

func (f *follower) VoteRequest(node *Node, req server.VoteRequest) {
	glog.Infof(node.logFormat("received vote request from candidate %d"), req.CandidateID())

	if req.Term() > node.Term() {
		node.SetTerm(req.Term())
		req.Grant()

		glog.Infof(node.logFormat("granted vote request with term %d"), req.Term())
	} else {
		req.Deny()

		glog.Infof(node.logFormat("denied vote request with term %d"), req.Term())
	}
}

func (f *follower) AppendEntriesRequest(node *Node) {
	// TODO(AD)
	glog.Info(node.logFormat("received append entries request"))
}

func (f *follower) name() string {
	return "follower"
}
