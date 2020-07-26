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

func (hb *NodeHeartbeat) Beat(term uint32) {
	go func(term uint32) {
		for _, c := range hb.conns {
			go func(c conn.Connection) {
				// TODO(AD) Not checking result as only used as heartbeat.
				c.RequestAppend(term)
			}(c)
		}
	}(term)
}
