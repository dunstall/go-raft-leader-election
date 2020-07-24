package node

import (
	"testing"

	"github.com/dunstall/goraft/pkg/elector/mock_elector"
	"github.com/dunstall/goraft/pkg/server/mock_server"
	"github.com/golang/mock/gomock"
)

func TestCandidateExpire(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var expectedTerm uint32 = 2
	elector := mock_elector.NewMockElector(ctrl)
	elector.EXPECT().Elect(expectedTerm)

	node := NewNode(0xfa, elector)
	node.setState(NewCandidate())

	node.Expire()
	if node.state.name() != candidateName {
		t.Error("expected node to be in candidate state")
	}

	// The node should have entered a new term.
	actual := node.Term()
	if actual != expectedTerm {
		t.Errorf("node.Term() != %d, actual %d", expectedTerm, actual)
	}
}

func TestCandidateElect(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	node := NewNode(0xfa, mock_elector.NewMockElector(ctrl))
	node.setState(NewCandidate())

	node.Elect()
	if node.state.name() != leaderName {
		t.Error("expected node to be in leader state")
	}

	// The term should not have changed.
	var expected uint32 = 1
	actual := node.Term()
	if actual != expected {
		t.Errorf("node.Term() != %d, actual %d", expected, actual)
	}
}

func TestCandidateVoteRequestTermGreater(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	node := NewNode(0xfa, mock_elector.NewMockElector(ctrl))
	node.setState(NewCandidate())

	var newTerm uint32 = node.Term() + 1
	var candidateID uint32 = 0xff

	mockreq := mock_server.NewMockVoteRequest(ctrl)
	mockreq.EXPECT().CandidateID().AnyTimes().Return(candidateID)
	mockreq.EXPECT().Term().AnyTimes().Return(newTerm)

	// As the term is greater than the node should grant and call back to
	// follower in new term.
	mockreq.EXPECT().Grant()
	node.VoteRequest(mockreq)

	if node.state.name() != followerName {
		t.Error("expected node to be in follower state")
	}

	term := node.Term()
	if term != newTerm {
		t.Errorf("node.Term() != %d, actual %d", newTerm, term)
	}
}

func TestCandidateVoteRequestTermEqual(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	node := NewNode(0xfa, mock_elector.NewMockElector(ctrl))
	node.setState(NewCandidate())

	var newTerm uint32 = node.Term()
	var candidateID uint32 = 0xff

	mockreq := mock_server.NewMockVoteRequest(ctrl)
	mockreq.EXPECT().CandidateID().AnyTimes().Return(candidateID)
	mockreq.EXPECT().Term().AnyTimes().Return(newTerm)

	// As the term is equal deny the request and dont change the current state.
	mockreq.EXPECT().Deny()
	node.VoteRequest(mockreq)

	if node.state.name() != candidateName {
		t.Error("expected node to be in candidate state")
	}
	term := node.Term()
	if term != newTerm {
		t.Errorf("node.Term() != %d, actual %d", newTerm, term)
	}
}