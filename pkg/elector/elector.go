package elector

type Elector interface {
	// TODO run in background and write to Election clannel for raft to read
	Elect(term uint32)
	Elected() <-chan bool
	Close()
}
