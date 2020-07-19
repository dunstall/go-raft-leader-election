package raft

import (
	"context"
	"log"
	"net"

	"github.com/dunstall/goraft/pkg/pb"
	"google.golang.org/grpc"
)

func ListenAndServe() error {
	server := server{}

	lis, err := net.Listen("tcp", ":3110")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterRaftServer(s, &server)
	if err := s.Serve(lis); err != nil {
		return err
	}
	return nil
}

func RequestVote() {
	conn, err := DialContext(context.Background(), "localhost:3110")
	if err != nil {
		log.Fatal(err)
	}

	resp, err := conn.RequestVote(context.Background(), &pb.RequestVoteRequest{
		Term:        0,
		CandidateId: 1,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.VoteGranted)
}
