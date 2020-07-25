package conn

import (
	"github.com/golang/glog"
	"google.golang.org/grpc"
)

// GRPCClient implements Client using gRPC.
type GRPCClient struct {
	id uint32
}

// NewGRPCClient returns a new client with the given node ID.
func NewGRPCClient(id uint32) Client {
	return &GRPCClient{id: id}
}

func (c *GRPCClient) Dial(addr string) Connection {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		// As the connection is in the background this should never happen so crash
		// early.
		glog.Fatalf("failed to Dial the node at %s: %s", addr, err)
	}
	return NewGRPCConnection(conn, c.id)
}
