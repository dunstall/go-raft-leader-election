package server

import (
	"github.com/dunstall/goraft/pkg/pb"
)

type VoteRequest interface {
	Term() uint32
	CandidateID() uint32
	Grant()
	Deny()
}

type VoteCallback struct {
	Request  *pb.RequestVoteRequest
	RespChan chan *pb.RequestVoteResponse
}

func (cb *VoteCallback) Term() uint32 {
	return cb.Request.Term
}

func (cb *VoteCallback) CandidateID() uint32 {
	return cb.Request.CandidateId
}

func (cb *VoteCallback) Grant() {
	cb.RespChan <- &pb.RequestVoteResponse{
		Term:        cb.Request.Term,
		VoteGranted: true,
	}
}

func (cb *VoteCallback) Deny() {
	cb.RespChan <- &pb.RequestVoteResponse{
		Term:        cb.Request.Term,
		VoteGranted: false,
	}
}
