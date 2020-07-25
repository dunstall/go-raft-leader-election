package elector

// Elector handles electing nodes.
type Elector interface {
	// Elect attempts to elect the current node by requesting votes from peers.
	// The result of the election is written to Elected() so this can be run
	// in the background and watch Elected() for results.
	Elect(term uint32)

	// Elected returns the results of elections.
	Elected() <-chan bool

	// Close closes the elector.
	Close()
}
