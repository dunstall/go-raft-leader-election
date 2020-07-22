package elector

import (
	"context"
	"log"

	"github.com/dunstall/goraft/pkg/pb"
)

type GRPCElector struct {
	// Cluster servers mapping ID to address.
	servers map[uint32]string
}

func (e *GRPCElector) Elect(id uint32, term uint32) (bool, error) {
	results := make(chan bool)
	// Fire of all requests concurrently.
	for serverID, _ := range e.servers {
		go e.requestVote(id, serverID, term, results)
	}

	votes := 0
	for n := 0; n < len(e.servers); n++ {
		select {
		// TODO(AD) Also include timeout case?
		case r := <-results:
			if r {
				votes++
			}
			// If have a majority return early.
			if e.isMajority(votes) {
				return true, nil
			}
		}
	}

	return false, nil
}

func (e *GRPCElector) requestVote(id uint32, serverID uint32, term uint32, results chan<- bool) {
	if id == serverID {
		// Vote for self.
		results <- true
		return
	}

	// TODO(AD) Should have a poll of connections for each peer.
	// TODO(AD) context.WithTimeout
	conn, err := DialContext(context.Background(), e.servers[serverID])
	if err != nil {
		log.Println(err)
		// Did not receive a vote.
		results <- false
	}
	defer conn.Close()

	resp, err := conn.RequestVote(context.Background(), &pb.RequestVoteRequest{
		Term:        term,
		CandidateId: id,
	})
	if err != nil {
		log.Println(err)
		// Did not receive a vote.
		results <- false
	}

	results <- resp.VoteGranted
}

func (e *GRPCElector) isMajority(votes int) bool {
	return votes > len(e.servers)/2
}
