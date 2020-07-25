package heartbeat

import (
	"github.com/dunstall/goraft/pkg/conn"
)

type NodeHeartbeat struct {
	id    uint32
	conns map[uint32]conn.Connection
}

// TODO(AD) Need a success/failure channel
func NewNodeHeartbeat(id uint32, conns map[uint32]conn.Connection) Heartbeat {
	return &NodeHeartbeat{id: id, conns: conns}
}

func (e *NodeHeartbeat) Beat(term uint32) {
	// TODO(AD) Run in background.
	// go func() {
	// }()

	for _, c := range e.conns {
		go func(c conn.Connection) {
			// TODO(AD) Not checking result as only used as heartbeat.
			c.RequestAppend(term)
		}(c)
	}
}
