package raft

import (
	"math/rand"
	"time"

	"github.com/dunstall/goraft/pkg/elector"
	"github.com/dunstall/goraft/pkg/node"
	"github.com/dunstall/goraft/pkg/server"
)

func Run(id uint32) {
	rand.Seed(time.Now().UnixNano())

	nodes := make(map[uint32]string)
	node := node.NewNode(id, elector.NewNodeElector(id, elector.NewGRPCClient(), nodes))

	server := server.NewServer()
	go server.ListenAndServe(":3110")

	for {
		d := time.Duration(time.Duration(rand.Intn(150)+150)) * time.Millisecond
		select {
		case <-time.After(d):
			node.Expire()
		// case <-election
		// node.Elect()
		case req := <-server.VoteRequests:
			node.VoteRequest(&req)
		}
	}
}
