package elector_test

import (
	"testing"

	"github.com/dunstall/goraft/pkg/conn"
	"github.com/dunstall/goraft/pkg/conn/mock_conn"
	"github.com/dunstall/goraft/pkg/elector"
	"github.com/golang/mock/gomock"
)

const (
	exampleID   uint32 = 0xfa
	exampleTerm uint32 = 0xaa
)

func TestNodeElectorOnlyNodeIsSelf(t *testing.T) {
	e := elector.NewNodeElector(exampleID, make(map[uint32]conn.Connection))
	e.Elect(exampleTerm)
	if !<-e.Elected() {
		t.Errorf("e.Elect() != true")
	}
}

func TestNodeElectorAllGrant(t *testing.T) {
	// Expect all nodes to grant.
	conns, ctrl := registerClient(4, 4, t)
	defer ctrl.Finish()

	e := elector.NewNodeElector(exampleID, conns)
	e.Elect(exampleTerm)
	if !<-e.Elected() {
		t.Errorf("e.Elect() != true")
	}
}

func TestNodeElectorMajorityGrant(t *testing.T) {
	// Expect 2 out of 4 external nodes to grant. As the node votes for itself
	// this is a majority.
	conns, ctrl := registerClient(4, 2, t)
	defer ctrl.Finish()

	e := elector.NewNodeElector(exampleID, conns)
	e.Elect(exampleTerm)
	if !<-e.Elected() {
		t.Errorf("e.Elect() != true")
	}
}

func TestNodeElectorMajorityDeny(t *testing.T) {
	// Expect all nodes deny.
	conns, ctrl := registerClient(4, 0, t)
	defer ctrl.Finish()

	e := elector.NewNodeElector(exampleID, conns)
	e.Elect(exampleTerm)
	if <-e.Elected() {
		t.Errorf("e.Elect() != false")
	}
}

func registerClient(n uint32, votes uint32, t *testing.T) (map[uint32]conn.Connection, *gomock.Controller) {
	conns := make(map[uint32]conn.Connection)
	ctrl := gomock.NewController(t)

	var nodeID uint32
	for nodeID = 1; nodeID <= n; nodeID++ {
		conn := mock_conn.NewMockConnection(ctrl)
		conn.EXPECT().RequestVote(exampleTerm).Return(votes > 0)
		conns[nodeID] = conn
		if votes > 0 {
			votes--
		}
	}
	return conns, ctrl
}
