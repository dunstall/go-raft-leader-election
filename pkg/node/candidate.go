package node

import (
	"log"

	"github.com/dunstall/goraft/pkg/server"
)

type candidate struct {
}

func NewCandidate() nodeState {
	return &candidate{}
}

func (c *candidate) Expire(node *Node) {
	// TODO(AD) -> Candidate
	log.Println("candidate: node timed out")
}

func (c *candidate) Elect(node *Node) {
	// TODO(AD) -> Leader
	log.Println("candidate: node elected")
	node.setState(node.leaderState())
}

func (c *candidate) VoteRequest(node *Node, cb server.VoteRequest) {
	// TODO(AD)
	log.Println("candidate: received vote request")
}

func (c *candidate) AppendEntriesRequest(node *Node) {
	// TODO(AD)
	log.Println("candidate: received append entries request")
}
