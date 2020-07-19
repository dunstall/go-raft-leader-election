package raft

type Term struct {
	N        uint32
	VotedFor uint32
}

func NewTerm() Term {
	return Term{
		N:        0,
		VotedFor: 0,
	}
}

func (t *Term) Next() {
	t.N++
	t.VotedFor = 0
}
