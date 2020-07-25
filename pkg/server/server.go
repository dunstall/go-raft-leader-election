package server

import (
	"context"
	"log"
	"net"

	"github.com/dunstall/goraft/pkg/pb"
	"google.golang.org/grpc"
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

type Server struct {
	VoteRequests   chan VoteCallback
	AppendRequests chan AppendCallback
}

func NewServer() Server {
	return Server{
		VoteRequests: make(chan VoteCallback),
	}
}

func (s *Server) RequestVote(ctx context.Context, in *pb.RequestVoteRequest) (*pb.RequestVoteResponse, error) {
	respChan := make(chan *pb.RequestVoteResponse)
	s.VoteRequests <- VoteCallback{in, respChan}

	// TODO(AD) Timeout with select?
	resp := <-respChan
	return resp, nil
}

func (s *Server) AppendEntries(ctx context.Context, in *pb.AppendEntriesRequest) (*pb.AppendEntriesResponse, error) {
	respChan := make(chan *pb.AppendEntriesResponse)
	s.AppendRequests <- AppendCallback{in, respChan}

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
