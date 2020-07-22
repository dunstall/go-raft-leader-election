package node

import (
	"testing"

	"github.com/dunstall/goraft/pkg/elector/mock_elector"
	"github.com/dunstall/goraft/pkg/server/mock_server"
	"github.com/golang/mock/gomock"
)

func TestLeaderExpire(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	node := NewNode(mock_elector.NewMockElector(ctrl))
	node.setState(node.leaderState())

	node.Expire()
	if node.state != node.followerState() {
		t.Error("expected node to be in follower state")
	}

	// The node should have entered a new term.
	var expected uint32 = 2
	actual := node.Term()
	if actual != expected {
		t.Errorf("node.Term() != %d, actual %d", expected, actual)
	}
}

func TestLeaderElect(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	node := NewNode(mock_elector.NewMockElector(ctrl))
	node.setState(node.leaderState())

	node.Elect()
	if node.state != node.leaderState() {
		t.Error("expected node to be in leader state")
	}

	// The term should not have changed.
	var expected uint32 = 1
	actual := node.Term()
	if actual != expected {
		t.Errorf("node.Term() != %d, actual %d", expected, actual)
	}
}

func TestLeaderVoteRequestTermGreater(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	node := NewNode(mock_elector.NewMockElector(ctrl))
	node.setState(node.leaderState())

	var newTerm uint32 = node.Term() + 1
	var candidateID uint32 = 0xff

	mockreq := mock_server.NewMockVoteRequest(ctrl)
	mockreq.EXPECT().CandidateID().AnyTimes().Return(candidateID)
	mockreq.EXPECT().Term().AnyTimes().Return(newTerm)

	// As the term is greater than the node should grant and call back to
	// follower in new term.
	mockreq.EXPECT().Grant()
	node.VoteRequest(mockreq)

	if node.state != node.followerState() {
		t.Error("expected node to be in follower state")
	}

	term := node.Term()
	if term != newTerm {
		t.Errorf("node.Term() != %d, actual %d", newTerm, term)
	}
}

func TestLeaderVoteRequestTermEqual(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	node := NewNode(mock_elector.NewMockElector(ctrl))
	node.setState(node.leaderState())

	var newTerm uint32 = node.Term()
	var candidateID uint32 = 0xff

	mockreq := mock_server.NewMockVoteRequest(ctrl)
	mockreq.EXPECT().CandidateID().AnyTimes().Return(candidateID)
	mockreq.EXPECT().Term().AnyTimes().Return(newTerm)

	// As the term is equal deny the request and dont change the current state.
	mockreq.EXPECT().Deny()
	node.VoteRequest(mockreq)

	if node.state != node.leaderState() {
		t.Error("expected node to be in leader state")
	}
	term := node.Term()
	if term != newTerm {
		t.Errorf("node.Term() != %d, actual %d", newTerm, term)
	}
}
