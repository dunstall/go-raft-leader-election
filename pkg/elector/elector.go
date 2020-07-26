package elector

// Elector handles electing nodes.
type Elector interface {
	// Elect attempts to elect the current node by requesting votes from peers.
	// This runs in the background (in a new goroutine and writes the result
	// to Elected().
	Elect(term uint32)

	// Elected returns the results of elections.
	Elected() <-chan bool
}
