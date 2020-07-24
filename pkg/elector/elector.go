package elector

type Elector interface {
	Elect(term uint32) bool
	Close()
}
