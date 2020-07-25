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
	// TODO(AD) Ensure this is all background so should never return an error.
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		glog.Fatalf("")
	}
	return NewGRPCConnection(conn, c.id)
}
