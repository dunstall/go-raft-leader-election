package node

import (
	"github.com/dunstall/goraft/pkg/server"
)

type nodeState interface {
	Expire(node *Node)
	Elect(node *Node)
	VoteRequest(node *Node, cb server.VoteRequest)
	AppendEntriesRequest(node *Node)
}
