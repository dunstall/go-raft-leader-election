package node

import (
	"fmt"
)

type Follower struct {
	node Node
}

func NewFollower(node Node) nodeState {
	return &Follower{node: node}
}

func (f *Follower) Expire() {
	// TODO(AD)
	fmt.Println("follower timed out")
	f.node.setState(f.node.candidateState())
}

func (f *Follower) Elect() {
	// TODO(AD)
	fmt.Println("cannot elect a follower")
}

func (f *Follower) ReceiveVoteRequest() {
	// TODO(AD)
}

func (f *Follower) ReceiveAppendEntriesRequest() {
	// TODO(AD)
}
