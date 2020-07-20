package node

import (
	"fmt"
)

type Follower struct {
}

func NewFollower() nodeState {
	return &Follower{}
}

func (f *Follower) Expire(node *Node) {
	// TODO(AD)
	fmt.Println("follower timed out")
	node.setState(node.candidateState())
}

func (f *Follower) Elect(node *Node) {
	// TODO(AD)
	fmt.Println("cannot elect a follower")
}

func (f *Follower) ReceiveVoteRequest(node *Node) {
	// TODO(AD)
}

func (f *Follower) ReceiveAppendEntriesRequest(node *Node) {
	// TODO(AD)
}
