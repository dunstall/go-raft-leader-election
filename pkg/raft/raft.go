package raft

import (
	"log"

	"github.com/dunstall/goraft/pkg/pb"
)

type Raft struct {
}

func (r *Raft) Election() {
	req := pb.RequestVoteRequest{
		Term:        0,
		CandidateId: 0,
	}
	log.Println(req.Term)
}
