package node

import (
	"time"

	"github.com/dunstall/goraft/pkg/server"
)

type nodeState interface {
	Expire(node *Node)
	Elect(node *Node)
	Timeout() time.Duration
	VoteRequest(node *Node, req server.VoteRequest)
	AppendRequest(node *Node, req server.AppendRequest)

	name() string
}
