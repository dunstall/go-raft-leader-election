package heartbeat

type Heartbeat interface {
	Beat(term uint32)
}
