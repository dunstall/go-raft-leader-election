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

type ClusterConfig struct {
	Nodes map[uint32]string
}

type Raft struct {
	id    uint32
	conns map[uint32]conn.Connection
}

func NewRaft(id uint32, config ClusterConfig) Raft {
	client := conn.NewGRPCClient(id)
	conns := make(map[uint32]conn.Connection)
	for nodeID, addr := range config.Nodes {
		if nodeID != id {
			conns[nodeID] = client.Dial(addr)
		}
	}
	return Raft{id: id, conns: conns}
}

func (r *Raft) Run() {
	rand.Seed(time.Now().UnixNano())

	e := elector.NewNodeElector(r.id, r.conns)
	hb := heartbeat.NewNodeHeartbeat(r.id, r.conns)
	node := node.NewNode(r.id, e, hb)

	server := server.NewServer()
	addr := ":" + strconv.Itoa(basePort+int(r.id))
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

func (r *Raft) Close() {
	for _, conn := range r.conns {
		conn.Close()
	}
}
