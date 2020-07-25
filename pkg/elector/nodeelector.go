package elector

import (
	"sync"
	"sync/atomic"

  "github.com/dunstall/goraft/pkg/elector/conn"
)

type NodeElector struct {
	elected chan bool
	id      uint32
	conns   map[uint32]conn.Connection
}

func NewNodeElector(id uint32, client conn.Client, nodes map[uint32]string) Elector {
	conns := make(map[uint32]conn.Connection)
	for nodeID, addr := range nodes {
		if nodeID != id {
			conns[nodeID] = client.Dial(addr)
		}
	}
	return &NodeElector{elected: make(chan bool, 1), id: id, conns: conns}
}

func (e *NodeElector) Elect(term uint32) {
	var votes uint32 = 0

	var wg sync.WaitGroup
	for _, c := range e.conns {
		wg.Add(1)
		go func(c conn.Connection) {
			defer wg.Done()
			if c.RequestVote(term) {
				atomic.AddUint32(&votes, 1)
			}
		}(c)
	}

	wg.Wait()

	e.elected <- e.isMajority(int(votes))
}

func (e *NodeElector) Elected() <-chan bool {
	return e.elected
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
