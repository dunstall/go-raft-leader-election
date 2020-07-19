package raft

import (
	"context"
	"fmt"

	"github.com/dunstall/goraft/pkg/pb"
)

type server struct{}

func (s *server) AppendEntries(ctx context.Context, in *pb.AppendEntriesRequest) (*pb.AppendEntriesResponse, error) {
	return nil, fmt.Errorf("err")
}

func (s *server) RequestVote(ctx context.Context, in *pb.RequestVoteRequest) (*pb.RequestVoteResponse, error) {
	fmt.Println("vote")
	return &pb.RequestVoteResponse{
		Term:        0,
		VoteGranted: true,
	}, nil
}
