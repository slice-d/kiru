package core

import (
	"github.com/slice-d/genzai/common/raft"
	"github.com/slice-d/genzai/common/service"
)

type RaftGroup struct {
	service.BaseService

	// Raft
	raft      *raft.Raft // The consensus mechanism
	snapshots raft.SnapshotStore
	transport *RaftTransport
	//transport  *raft.NetworkTransport
	store      IRaftStore
	observerCh chan raft.Observation
	observer   *raft.Observer
}

func newRaftGroup(store IRaftStore) *RaftGroup {
	r := &RaftGroup{}

	return r
}

func (r *RaftGroup) OnReset() error {
	return nil
}

func (r *RaftGroup) startRaft() error {
	return nil
}
