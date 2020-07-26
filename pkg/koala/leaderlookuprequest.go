package koala

import (
	"github.com/dunstall/goraft/pkg/pb"
)

type LeaderLookupRequest interface {
	Respond(leaderID uint32)
}

type LeaderLookupCallback struct {
	Request  *pb.LeaderLookupRequest
	RespChan chan *pb.LeaderLookupResponse
}

func (cb *LeaderLookupCallback) Respond(leaderID uint32) {
	cb.RespChan <- &pb.LeaderLookupResponse{
		LeaderId: leaderID,
	}
}
