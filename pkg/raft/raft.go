package raft

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/dunstall/goraft/pkg/conn"
	"github.com/dunstall/goraft/pkg/elector"
	"github.com/dunstall/goraft/pkg/heartbeat"
	"github.com/dunstall/goraft/pkg/node"
	"github.com/dunstall/goraft/pkg/server"
)

const (
	basePort = 4110
)

func Run(id uint32) {
	rand.Seed(time.Now().UnixNano())

	nodes := map[uint32]string{
		1: "localhost:4111",
		2: "localhost:4112",
		3: "localhost:4113",
	}

	client := conn.NewGRPCClient(id)
	conns := make(map[uint32]conn.Connection)
	for nodeID, addr := range nodes {
		if nodeID != id {
			conns[nodeID] = client.Dial(addr)
			defer conns[nodeID].Close()
		}
	}

	e := elector.NewNodeElector(id, conns)
	hb := heartbeat.NewNodeHeartbeat(id, conns)
	node := node.NewNode(id, e, hb)

	server := server.NewServer()
	// TODO(AD) if starting goroutine myst also handler closing
	addr := ":" + strconv.Itoa(basePort+int(id))
	go server.ListenAndServe(addr)

	for {
		select {
		case <-time.After(node.Timeout()):
			node.Expire()
		case granted := <-e.Elected():
			// TODO(AD) pass vote to node so it can log and update (revert to follower if
			// denied
			if granted {
				node.Elect()
			}
		case req := <-server.VoteRequests():
			node.VoteRequest(&req)
		case req := <-server.AppendRequests():
			node.AppendRequest(&req)
		}
	}
}
