package conn

import (
	"context"
	"time"

	"github.com/dunstall/goraft/pkg/pb"
	"github.com/golang/glog"
	"google.golang.org/grpc"
)

const (
	// TODO(AD) Configuration option
	requestTimeoutMS = 100
)

type GRPCConnection struct {
	client pb.RaftClient
	conn   *grpc.ClientConn
	id     uint32
}

func NewGRPCConnection(conn *grpc.ClientConn, id uint32) Connection {
	return &GRPCConnection{client: pb.NewRaftClient(conn), conn: conn, id: id}
}

// TODO(AD) This must handle reconnect
func (conn *GRPCConnection) RequestVote(term uint32) bool {
	req := &pb.RequestVoteRequest{
		Term:        term,
		CandidateId: conn.id,
	}

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeoutMS*time.Millisecond)
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

func (conn *GRPCConnection) RequestAppend(term uint32) bool {
	req := &pb.AppendEntriesRequest{
		Term:     term,
		LeaderId: conn.id,
	}

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeoutMS*time.Millisecond)
	defer cancel()
	resp, err := conn.client.AppendEntries(ctx, req)
	if err != nil {
		glog.Warning("error response from connection append entries %s", err)
		return false
	}

	if resp.Term != term {
		glog.Warning("append entries response included an invalid term: %d, expected %d", resp.Term, term)
		return false
	}

	return resp.Success
}

func (conn *GRPCConnection) Close() {
	conn.conn.Close()
}
