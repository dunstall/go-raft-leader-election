package conn

type Connection interface {
	RequestVote(term uint32) bool
	Close()
}
