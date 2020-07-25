package elector_test

import (
	"testing"

	"github.com/dunstall/goraft/pkg/elector"
	"github.com/dunstall/goraft/pkg/elector/conn/mock_conn"
	"github.com/golang/mock/gomock"
)

const (
	exampleID   uint32 = 0xfa
	exampleTerm uint32 = 0xaa
)

func TestNodeElectorOnlyNodeIsSelf(t *testing.T) {
	nodes := map[uint32]string{
		exampleID: "1.1.1.1",
	}

	client, ctrl := newMockClient(t)
	defer ctrl.Finish()

	e := elector.NewNodeElector(exampleID, client, nodes)
	e.Elect(exampleTerm)
	if !<-e.Elected() {
		t.Errorf("e.Elect() != true")
	}
}

func TestNodeElectorAllGrant(t *testing.T) {
	// Expect all nodes to grant.
	client, nodes, ctrl := registerClient(4, t)
	defer ctrl.Finish()

	e := elector.NewNodeElector(exampleID, client, nodes)
	e.Elect(exampleTerm)
	if !<-e.Elected() {
		t.Errorf("e.Elect() != true")
	}
}

func TestNodeElectorMajorityGrant(t *testing.T) {
	// Expect 2 out of 4 external nodes to grant. As the node votes for itself
	// this is a majority.
	client, nodes, ctrl := registerClient(2, t)
	defer ctrl.Finish()

	e := elector.NewNodeElector(exampleID, client, nodes)
	e.Elect(exampleTerm)
	if !<-e.Elected() {
		t.Errorf("e.Elect() != true")
	}
}

func TestNodeElectorMajorityDeny(t *testing.T) {
	// Expect all nodes deny.
	client, nodes, ctrl := registerClient(0, t)
	defer ctrl.Finish()

	e := elector.NewNodeElector(exampleID, client, nodes)
	e.Elect(exampleTerm)
	if <-e.Elected() {
		t.Errorf("e.Elect() != false")
	}
}

func registerClient(votes uint32, t *testing.T) (*mock_conn.MockClient, map[uint32]string, *gomock.Controller) {
	nodes := map[uint32]string{
		exampleID: "1.1.1.1",
		0xfb:      "2.2.2.2",
		0xfc:      "3.3.3.3",
		0xfd:      "3.3.3.3",
		0xfe:      "5.5.5.5",
	}

	client, ctrl := newMockClient(t)
	for nodeID, addr := range nodes {
		// Only connect to nodes that are not self.
		if nodeID == exampleID {
			continue
		}
		conn := mock_conn.NewMockConnection(ctrl)
		conn.EXPECT().RequestVote(exampleTerm).Return(votes > 0)
		client.EXPECT().Dial(addr).Return(conn)
		if votes > 0 {
			votes--
		}
	}
	return client, nodes, ctrl
}

func newMockClient(t *testing.T) (*mock_conn.MockClient, *gomock.Controller) {
	ctrl := gomock.NewController(t)
	return mock_conn.NewMockClient(ctrl), ctrl
}
