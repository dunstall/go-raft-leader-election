package raft

import (
	"testing"
)

func TestNewNode(t *testing.T) {
	var id uint32 = 0xffaa
	expected := Node{
		ID:   id,
		Term: NewTerm(),
	}

	node := NewNode(id)
	if node != expected {
		t.Errorf("node == %#v, expected %#v", node, expected)
	}
}
