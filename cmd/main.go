package main

import (
	"flag"
	"log"

	"github.com/dunstall/goraft/pkg/raft"
)

func main() {
	id := flag.Uint("id", 0, "node ID")
	flag.Parse()

	if *id == 0 {
		log.Println("bad id")
	}

	raft.Run(uint32(*id))
}
