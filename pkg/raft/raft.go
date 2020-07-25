package raft

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/dunstall/goraft/pkg/conn"
	"github.com/dunstall/goraft/pkg/elector"
	"github.com/dunstall/goraft/pkg/node"
	"github.com/dunstall/goraft/pkg/server"
)

const (
	basePort = 4110
)

func Run(id uint32) {
	rand.Seed(time.Now().UnixNano())

	nodes := map[uint32]string{
		1: ":4111",
		2: ":4112",
		3: ":4113",
	}

	e := elector.NewNodeElector(id, conn.NewGRPCClient(id), nodes)
	node := node.NewNode(id, e)

	server := server.NewServer()
	// TODO(AD) if starting goroutine myst also handler closing
	addr := ":" + strconv.Itoa(basePort+int(id))
	go server.ListenAndServe(addr)

	<-time.After(time.Second * 10)
	for {
		d := time.Duration(time.Duration(rand.Intn(150)+150)) * time.Millisecond
		select {
		case <-time.After(d):
			node.Expire()
		case granted := <-e.Elected():
			// TODO(AD) pass vote to node so it can log and update (revert to follower if
			// denied
			if granted {
				node.Elect()
			}
		case req := <-server.VoteRequests:
			node.VoteRequest(&req)
		case req := <-server.AppendRequests:
			node.AppendRequest(&req)
		}
	}
}
