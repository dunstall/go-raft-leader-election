package raft

import (
	"context"

	"github.com/dunstall/goraft/pkg/pb"
	"google.golang.org/grpc"
)

const (
	port = ":3110"
)

type Connection struct {
	client pb.RaftClient
	conn   *grpc.ClientConn
}

func DialContext(ctx context.Context, addr string) (Connection, error) {
	// TODO(AD) DialWithContext?
	conn, err := grpc.DialContext(ctx, addr, grpc.WithInsecure())
	if err != nil {
		return Connection{}, err
	}
	return Connection{client: pb.NewRaftClient(conn), conn: conn}, nil
}

func (conn *Connection) RequestVote(ctx context.Context, in *pb.RequestVoteRequest) (*pb.RequestVoteResponse, error) {
	return conn.client.RequestVote(ctx, in)
}

func (conn *Connection) Close() {
	defer conn.conn.Close()
}
