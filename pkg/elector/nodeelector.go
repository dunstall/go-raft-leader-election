package elector

import (
	"sync"
	"sync/atomic"
)

type NodeElector struct {
	id    uint32
	conns map[uint32]Connection
}

func NewNodeElector(id uint32, client Client, nodes map[uint32]string) Elector {
	conns := make(map[uint32]Connection)
	for nodeID, addr := range nodes {
		if nodeID != id {
			conns[nodeID] = client.Dial(addr)
		}
	}
	return &NodeElector{id: id, conns: conns}
}

func (e *NodeElector) Elect(term uint32) bool {
	var votes uint32 = 0

	var wg sync.WaitGroup
	for _, conn := range e.conns {
		wg.Add(1)
		go func(conn Connection) {
			defer wg.Done()
			if conn.RequestVote(term) {
				atomic.AddUint32(&votes, 1)
			}
		}(conn)
	}

	wg.Wait()

	return e.isMajority(int(votes))
}

func (e *NodeElector) Close() {
	for _, conn := range e.conns {
		conn.Close()
	}
}

func (e *NodeElector) isMajority(votes int) bool {
	// As the node always votes for itself need at least floor(n/2) other nodes.
	return votes >= (len(e.conns) / 2)
}
