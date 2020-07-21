package server

import (
	"context"
	"log"
	"net"

	"github.com/dunstall/goraft/pkg/pb"
	"google.golang.org/grpc"
)

type Callback struct {
	Request  *pb.RequestVoteRequest
	RespChan chan *pb.RequestVoteResponse
}

type Server struct {
	VoteRequests chan Callback
}

func NewServer() Server {
	return Server{
		VoteRequests: make(chan Callback),
	}
}

func (s *Server) RequestVote(ctx context.Context, in *pb.RequestVoteRequest) (*pb.RequestVoteResponse, error) {
	respChan := make(chan *pb.RequestVoteResponse)
	s.VoteRequests <- Callback{in, respChan}

	// TODO(AD) Timeout with select?
	resp := <-respChan
	return resp, nil
}

func (s *Server) ListenAndServe(addr string) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterRaftServer(grpcServer, s)
	if err := grpcServer.Serve(lis); err != nil {
		return err
	}
	return nil
}
