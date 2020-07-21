package node

import (
	"log"

	"github.com/dunstall/goraft/pkg/server"
)

type leader struct {
}

func NewLeader() nodeState {
	return &leader{}
}

func (l *leader) Expire(node *Node) {
	// TODO(AD)
	log.Println("leader: node timed out")
}

func (l *leader) Elect(node *Node) {
	// TODO(AD)
	log.Println("leader: leader already elected")
}

func (l *leader) VoteRequest(node *Node, cb server.Callback) {
	// TODO(AD)
	log.Println("leader: leader received vote request")
}

func (l *leader) AppendEntriesRequest(node *Node) {
	// TODO(AD)
	log.Println("leader: leader received append entries request")
}
