package node

import (
	"testing"
	"time"

	"github.com/dunstall/goraft/pkg/pb"
	"github.com/dunstall/goraft/pkg/server"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
)

func TestFollowerExpire(t *testing.T) {
	node := NewNode()
	node.Expire()
	if node.state != node.candidateState() {
		t.Error("expected node to be in candidate state")
	}
}

func TestFollowerElect(t *testing.T) {
	node := NewNode()
	node.Elect()
	if node.state != node.followerState() {
		t.Error("expected node to be in follower state")
	}
}

func TestFollowerInitialTerm(t *testing.T) {
	node := NewNode()
	var expected uint32 = 1
	actual := node.Term()
	if actual != expected {
		t.Error("node.Term() != %d, actual %d", expected, actual)
	}
}

func TestFollowerVoteRequestGreaterTerm(t *testing.T) {
	var newTerm uint32 = 2

	respChan := make(chan *pb.RequestVoteResponse)
	cb := server.Callback{
		&pb.RequestVoteRequest{
			Term:        newTerm,
			CandidateId: 0xff,
		},
		respChan,
	}

	node := NewNode()
	go node.VoteRequest(cb)

	expected := &pb.RequestVoteResponse{
		Term:        newTerm,
		VoteGranted: true,
	}

	select {
	case resp := <-cb.RespChan:
		if !proto.Equal(resp, expected) {
			t.Errorf("%s != %s", prototext.Format(resp), prototext.Format(expected))
		}
	case <-time.After(10 * time.Millisecond):
		t.Error("no vote response")
	}

	term := node.Term()
	if term != newTerm {
		t.Errorf("node.Term() != %d, actual %d", newTerm, term)
	}
}

func TestFollowerVoteRequestLessTerm(t *testing.T) {
	var newTerm uint32 = 0

	respChan := make(chan *pb.RequestVoteResponse)
	cb := server.Callback{
		&pb.RequestVoteRequest{
			Term:        newTerm,
			CandidateId: 0xff,
		},
		respChan,
	}

	node := NewNode()
	go node.VoteRequest(cb)

	expected := &pb.RequestVoteResponse{
		Term:        newTerm,
		VoteGranted: false,
	}

	select {
	case resp := <-cb.RespChan:
		if !proto.Equal(resp, expected) {
			t.Errorf("%s != %s", prototext.Format(resp), prototext.Format(expected))
		}
	case <-time.After(10 * time.Millisecond):
		t.Error("no vote response")
	}

	term := node.Term()
	if term != 1 {
		t.Errorf("node.Term() != %d, actual %d", 1, term)
	}
}
