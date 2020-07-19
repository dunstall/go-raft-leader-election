package raft

import ()

type Node struct {
	ID   uint32
	Term Term
}

func NewNode(id uint32) Node {
	return Node{ID: id, Term: NewTerm()}
}
