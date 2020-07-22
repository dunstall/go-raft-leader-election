package node

import (
	"github.com/dunstall/goraft/pkg/server"
	"github.com/golang/glog"
)

type leader struct {
}

func NewLeader() nodeState {
	return &leader{}
}

func (l *leader) Expire(node *Node) {
	glog.Infof("leader: node timed out in term %d", node.Term())
	node.SetTerm(node.Term() + 1)
	node.setState(node.followerState())
}

func (l *leader) Elect(node *Node) {
	glog.Warning("leader: leader already elected in term %d", node.Term())
}

func (l *leader) VoteRequest(node *Node, req server.VoteRequest) {
	glog.Infof("leader: received vote request from candidate %d", req.CandidateID())

	if req.Term() > node.Term() {
		glog.Infof("leader: vote request has greater term %d - reverting to follower", req.Term())

		node.setState(node.followerState())
		node.VoteRequest(req)
	} else {
		req.Deny()

		glog.Infof("leader: denied vote request with term %d", req.Term())
	}
}

func (l *leader) AppendEntriesRequest(node *Node) {
	// TODO(AD)
	glog.Infof("leader: leader received append entries request")
}
