package server

import (
	"github.com/dunstall/goraft/pkg/pb"
)

type AppendRequest interface {
	Term() uint32
	LeaderID() uint32
	Ok()
	Failure()
}

type AppendCallback struct {
	Request  *pb.AppendEntriesRequest
	RespChan chan *pb.AppendEntriesResponse
}

func (cb *AppendCallback) Term() uint32 {
	return cb.Request.Term
}

func (cb *AppendCallback) LeaderID() uint32 {
	return cb.Request.LeaderId
}

func (cb *AppendCallback) Ok() {
	cb.RespChan <- &pb.AppendEntriesResponse{
		Term:    cb.Request.Term,
		Success: true,
	}
}

func (cb *AppendCallback) Failure() {
	cb.RespChan <- &pb.AppendEntriesResponse{
		Term:    cb.Request.Term,
		Success: false,
	}
}
