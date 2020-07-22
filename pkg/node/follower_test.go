package node

import (
	"testing"

	"github.com/dunstall/goraft/pkg/elector/mock_elector"
	"github.com/dunstall/goraft/pkg/server/mock_server"
	"github.com/golang/mock/gomock"
)

func TestFollowerInitialTerm(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	node := NewNode(mock_elector.NewMockElector(ctrl))
	var expected uint32 = 1
	actual := node.Term()
	if actual != expected {
		t.Errorf("node.Term() != %d, actual %d", expected, actual)
	}
}

func TestFollowerExpire(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var expectedTerm uint32 = 2
	elector := mock_elector.NewMockElector(ctrl)
	elector.EXPECT().Elect(expectedTerm)

	node := NewNode(elector)
	node.Expire()
	if node.state != node.candidateState() {
		t.Error("expected node to be in candidate state")
	}

	// The node should have entered a new term.
	actual := node.Term()
	if actual != expectedTerm {
		t.Errorf("node.Term() != %d, actual %d", expectedTerm, actual)
	}
}

func TestFollowerElect(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	node := NewNode(mock_elector.NewMockElector(ctrl))
	node.Elect()
	if node.state != node.followerState() {
		t.Error("expected node to be in follower state")
	}
}

func TestFollowerVoteRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var newTerm uint32 = 2
	var candidateID uint32 = 0xff
	mockreq := mock_server.NewMockVoteRequest(ctrl)
	mockreq.EXPECT().CandidateID().AnyTimes().Return(candidateID)
	mockreq.EXPECT().Term().AnyTimes().Return(newTerm)

	node := NewNode(mock_elector.NewMockElector(ctrl))

	// As the term is greater the request should be granted.
	mockreq.EXPECT().Grant()
	node.VoteRequest(mockreq)

	// Nodes term should have updated.
	term := node.Term()
	if term != newTerm {
		t.Errorf("node.Term() != %d, actual %d", newTerm, term)
	}

	// The follower has given its vote for this term so should deny the request.
	mockreq.EXPECT().Deny()
	node.VoteRequest(mockreq)
}
