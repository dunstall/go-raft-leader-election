package koala

import (
	"context"
	"net"

	"github.com/dunstall/goraft/pkg/pb"
	"google.golang.org/grpc"
)

/* Server handles incoming requests and routes to the outgoing channels. */

// Each incoming request is pushed to a channel along with a response callback
// channel. The server then waits for the response. This blocking is ok as each
/* handler runs in its own goroutine. */
type Server struct {
	leaderLookupRequests chan LeaderLookupCallback
}

func NewServer() Server {
	return Server{
		leaderLookupRequests: make(chan LeaderLookupCallback),
	}
}

func (s *Server) LeaderLookupRequests() <-chan LeaderLookupCallback {
	return s.leaderLookupRequests
}

func (s *Server) LeaderLookup(ctx context.Context, req *pb.LeaderLookupRequest) (*pb.LeaderLookupResponse, error) {
	respChan := make(chan *pb.LeaderLookupResponse)
	s.leaderLookupRequests <- LeaderLookupCallback{req, respChan}
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
		pb.RegisterKoalaServer(grpcServer, s)
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
