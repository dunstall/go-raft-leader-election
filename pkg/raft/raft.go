package raft

import (
	"math/rand"
	"time"

	"github.com/dunstall/goraft/pkg/node"
	"github.com/dunstall/goraft/pkg/server"
)

func Run() {
	rand.Seed(time.Now().UnixNano())

	node := node.NewNode()

	server := server.NewServer()
	go server.ListenAndServe(":3110")

	for {
		d := time.Duration(time.Duration(rand.Intn(150) + 150)) * time.Millisecond
		select {
		case <-time.After(d):
			node.Expire()
		// case <-election
		// node.Elect()
		case cb := <-server.VoteRequests:
			node.VoteRequest(cb)
		}
	}
}
