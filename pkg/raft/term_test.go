package raft

import (
	"testing"
)

func TestNewTerm(t *testing.T) {
	expected := Term{
		N:        0,
		VotedFor: 0,
	}

	term := NewTerm()
	if term != expected {
		t.Errorf("term == %#v, expected %#v", term, expected)
	}
}

func TestTermNext(t *testing.T) {
	var n uint32 = 0xffaa
	expected := Term{
		N:        n + 1,
		VotedFor: 0,
	}

	term := Term{
		N:        n,
		VotedFor: 0xab,
	}
	term.Next()

	if term != expected {
		t.Errorf("term.Next() == %#v, expected %#v", term, expected)
	}
}
