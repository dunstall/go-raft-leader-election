package heartbeat

// Heartbeat handles sending a heartbeat to the nodes in the cluster.
type Heartbeat interface {
	// Beat sends the heartbeat to the nodes. This runs in the background (in
	// another goroutine) so returns immediately.
	Beat(term uint32)
}
