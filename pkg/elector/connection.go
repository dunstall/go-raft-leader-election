package elector

type Connection interface {
	// TODO(AD) Must handle timeouts by returning false
	RequestVote(term uint32) bool
}
