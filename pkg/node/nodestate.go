package node

type nodeState interface {
	Expire()
	Elect()
	// TODO(AD) Use callback with resp channel
	ReceiveVoteRequest()
	// TODO(AD) Use callback with resp channel
	ReceiveAppendEntriesRequest()
}
