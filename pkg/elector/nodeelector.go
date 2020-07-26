package elector

import (
	"sync"
	"sync/atomic"

	"github.com/dunstall/goraft/pkg/conn"
)

// NodeElector implements Elector by requesting votes from all nodes in the
// cluster.
type NodeElector struct {
	elected chan bool
	id      uint32
	conns   map[uint32]conn.Connection
}

// NewNodeElector returns an elector that uses the given connections to contact
// the cluster nodes. ID is the candidate ID of the node being elected.
func NewNodeElector(id uint32, conns map[uint32]conn.Connection) Elector {
	return &NodeElector{elected: make(chan bool, 1), id: id, conns: conns}
}

func (e *NodeElector) Elect(term uint32) {
	go func() {
		var votes uint32

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
	}()
}

func (e *NodeElector) Elected() <-chan bool {
	return e.elected
}

func (e *NodeElector) isMajority(votes int) bool {
	// As the node always votes for itself need at least floor(n/2) other nodes.
	return votes >= (len(e.conns) / 2)
}
