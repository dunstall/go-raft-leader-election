package elector

import (
	"context"
	"time"

	"github.com/dunstall/goraft/pkg/pb"
	"github.com/golang/glog"
	"google.golang.org/grpc"
)

const (
	// TODO(AD) Configuration option
	voteRequestTimeoutMS = 100
)

type GRPCConnection struct {
	client pb.RaftClient
	conn   *grpc.ClientConn
}

func NewGRPCConnection(conn *grpc.ClientConn) GRPCConnection {
	return GRPCConnection{client: pb.NewRaftClient(conn), conn: conn}
}

func (conn *GRPCConnection) RequestVote(term uint32) bool {
	req := &pb.RequestVoteRequest{
		Term:        term,
		CandidateId: 0,
	}

	ctx, cancel := context.WithTimeout(context.Background(), voteRequestTimeoutMS*time.Millisecond)
	defer cancel()
	resp, err := conn.client.RequestVote(ctx, req)
	if err != nil {
		glog.Warning("error response from connection request vote %s", err)
		return false
	}

	if resp.Term != term {
		glog.Warning("vote response included an invalid term: %d, expected %d", resp.Term, term)
		return false
	}

	return resp.VoteGranted
}

func (conn *GRPCConnection) Close() {
	conn.conn.Close()
}
