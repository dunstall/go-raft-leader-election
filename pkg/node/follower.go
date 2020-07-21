package node

import (
	"log"

	"github.com/dunstall/goraft/pkg/server"
)

type Follower struct {
}

func NewFollower() nodeState {
	return &Follower{}
}

func (f *Follower) Expire(node *Node) {
	log.Println("follower: node timed out")
	node.setState(node.candidateState())
}

func (f *Follower) Elect(node *Node) {
	log.Println("follower: cannot elect a follower")
}

func (f *Follower) VoteRequest(node *Node, cb server.Callback) {
	// TODO(AD)
	log.Println("follower: received vote request")
}

func (f *Follower) AppendEntriesRequest(node *Node) {
	// TODO(AD)
	log.Println("follower: received append entries request")
}
