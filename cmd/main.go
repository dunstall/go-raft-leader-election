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
		log.Fatal("bad node ID: %d", *id)
	}
	return uint32(*id)
}

func main() {
	raft.Run(parseNodeID())
}
