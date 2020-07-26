package main

import (
	"flag"
	"log"

	"github.com/dunstall/goraft/pkg/raft"
)

func parseNodeID() uint32 {
	id := flag.Uint("id", 0, "node ID")
	flag.Parse()
	if *id == 0 || *id > 0xFFFFFFFF {
		log.Fatalf("bad node ID: %d", *id)
	}
	return uint32(*id)
}

func main() {
	nodes := map[uint32]string{
		1: "localhost:4111",
		2: "localhost:4112",
		3: "localhost:4113",
	}
	config := raft.ClusterConfig{
		Nodes: nodes,
	}

	raft := raft.NewRaft(parseNodeID(), config)
	defer raft.Close()

	// TODO(AD) Accept context
	raft.Run()
}
