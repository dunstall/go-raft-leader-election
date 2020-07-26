package server

import (
	"context"
	"net"

	"github.com/dunstall/goraft/pkg/pb"
	"google.golang.org/grpc"
)

// Server handles incoming requests and routes to the outgoing channels.
//
// Each incoming request is pushed to a channel along with a response callback
// channel. The server then waits for the response. This blocking is ok as each
// handler runs in its own goroutine.
type Server struct {
	voteRequests   chan VoteCallback
	appendRequests chan AppendCallback
}

func NewServer() Server {
	return Server{
		voteRequests:   make(chan VoteCallback),
		appendRequests: make(chan AppendCallback),
	}
}

func (s *Server) VoteRequests() <-chan VoteCallback {
	return s.voteRequests
}

func (s *Server) AppendRequests() <-chan AppendCallback {
	return s.appendRequests
}

func (s *Server) RequestVote(ctx context.Context, req *pb.RequestVoteRequest) (*pb.RequestVoteResponse, error) {
	respChan := make(chan *pb.RequestVoteResponse)
	s.voteRequests <- VoteCallback{req, respChan}
	return <-respChan, nil
}

func (s *Server) AppendEntries(ctx context.Context, req *pb.AppendEntriesRequest) (*pb.AppendEntriesResponse, error) {
	respChan := make(chan *pb.AppendEntriesResponse)
	s.appendRequests <- AppendCallback{req, respChan}
	return <-respChan, nil
}

func (s *Server) ListenAndServe(ctx context.Context, addr string) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	errChan := make(chan error)
	go func(lis net.Listener, errChan chan<- error) {
		grpcServer := grpc.NewServer()
		pb.RegisterRaftServer(grpcServer, s)
		if err := grpcServer.Serve(lis); err != nil {
			errChan <- err
		}
	}(lis, errChan)

	select {
	case e := <-errChan:
		return e
	case <-ctx.Done():
		return nil
	}
}
