package node

import (
	"github.com/dunstall/goraft/pkg/server"
)

type nodeState interface {
	Expire(node *Node)
	Elect(node *Node)
	VoteRequest(node *Node, req server.VoteRequest)
	AppendRequest(node *Node, req server.AppendRequest)

	name() string
}
