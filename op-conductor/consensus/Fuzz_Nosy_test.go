package consensus

import (
	"io"
	"testing"

	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/hashicorp/raft"
	go_fuzz_utils "github.com/trailofbits/go-fuzz-utils"
)

func GetTypeProvider(data []byte) (*go_fuzz_utils.TypeProvider, error) {
	tp, err := go_fuzz_utils.NewTypeProvider(data)
	if err != nil {
		return nil, err
	}
	err = tp.SetParamsStringBounds(0, 1024)
	if err != nil {
		return nil, err
	}
	err = tp.SetParamsSliceBounds(0, 4096)
	if err != nil {
		return nil, err
	}
	err = tp.SetParamsBiases(0, 0, 0, 0)
	if err != nil {
		return nil, err
	}
	return tp, nil
}

func Fuzz_Nosy_RaftConsensus_AddNonVoter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rc *RaftConsensus
		fill_err = tp.Fill(&rc)
		if fill_err != nil {
			return
		}
		var id string
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var addr string
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var version uint64
		fill_err = tp.Fill(&version)
		if fill_err != nil {
			return
		}
		if rc == nil {
			return
		}

		rc.AddNonVoter(id, addr, version)
	})
}

func Fuzz_Nosy_RaftConsensus_AddVoter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rc *RaftConsensus
		fill_err = tp.Fill(&rc)
		if fill_err != nil {
			return
		}
		var id string
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var addr string
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var version uint64
		fill_err = tp.Fill(&version)
		if fill_err != nil {
			return
		}
		if rc == nil {
			return
		}

		rc.AddVoter(id, addr, version)
	})
}

func Fuzz_Nosy_RaftConsensus_Addr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rc *RaftConsensus
		fill_err = tp.Fill(&rc)
		if fill_err != nil {
			return
		}
		if rc == nil {
			return
		}

		rc.Addr()
	})
}

func Fuzz_Nosy_RaftConsensus_ClusterMembership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rc *RaftConsensus
		fill_err = tp.Fill(&rc)
		if fill_err != nil {
			return
		}
		if rc == nil {
			return
		}

		rc.ClusterMembership()
	})
}

func Fuzz_Nosy_RaftConsensus_CommitUnsafePayload__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rc *RaftConsensus
		fill_err = tp.Fill(&rc)
		if fill_err != nil {
			return
		}
		var payload *eth.ExecutionPayloadEnvelope
		fill_err = tp.Fill(&payload)
		if fill_err != nil {
			return
		}
		if rc == nil || payload == nil {
			return
		}

		rc.CommitUnsafePayload(payload)
	})
}

func Fuzz_Nosy_RaftConsensus_DemoteVoter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rc *RaftConsensus
		fill_err = tp.Fill(&rc)
		if fill_err != nil {
			return
		}
		var id string
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var version uint64
		fill_err = tp.Fill(&version)
		if fill_err != nil {
			return
		}
		if rc == nil {
			return
		}

		rc.DemoteVoter(id, version)
	})
}

func Fuzz_Nosy_RaftConsensus_LatestUnsafePayload__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rc *RaftConsensus
		fill_err = tp.Fill(&rc)
		if fill_err != nil {
			return
		}
		if rc == nil {
			return
		}

		rc.LatestUnsafePayload()
	})
}

func Fuzz_Nosy_RaftConsensus_Leader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rc *RaftConsensus
		fill_err = tp.Fill(&rc)
		if fill_err != nil {
			return
		}
		if rc == nil {
			return
		}

		rc.Leader()
	})
}

func Fuzz_Nosy_RaftConsensus_LeaderCh__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rc *RaftConsensus
		fill_err = tp.Fill(&rc)
		if fill_err != nil {
			return
		}
		if rc == nil {
			return
		}

		rc.LeaderCh()
	})
}

func Fuzz_Nosy_RaftConsensus_LeaderWithID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rc *RaftConsensus
		fill_err = tp.Fill(&rc)
		if fill_err != nil {
			return
		}
		if rc == nil {
			return
		}

		rc.LeaderWithID()
	})
}

func Fuzz_Nosy_RaftConsensus_RemoveServer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rc *RaftConsensus
		fill_err = tp.Fill(&rc)
		if fill_err != nil {
			return
		}
		var id string
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var version uint64
		fill_err = tp.Fill(&version)
		if fill_err != nil {
			return
		}
		if rc == nil {
			return
		}

		rc.RemoveServer(id, version)
	})
}

func Fuzz_Nosy_RaftConsensus_ServerID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rc *RaftConsensus
		fill_err = tp.Fill(&rc)
		if fill_err != nil {
			return
		}
		if rc == nil {
			return
		}

		rc.ServerID()
	})
}

func Fuzz_Nosy_RaftConsensus_Shutdown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rc *RaftConsensus
		fill_err = tp.Fill(&rc)
		if fill_err != nil {
			return
		}
		if rc == nil {
			return
		}

		rc.Shutdown()
	})
}

func Fuzz_Nosy_RaftConsensus_TransferLeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rc *RaftConsensus
		fill_err = tp.Fill(&rc)
		if fill_err != nil {
			return
		}
		if rc == nil {
			return
		}

		rc.TransferLeader()
	})
}

func Fuzz_Nosy_RaftConsensus_TransferLeaderTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rc *RaftConsensus
		fill_err = tp.Fill(&rc)
		if fill_err != nil {
			return
		}
		var id string
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var addr string
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if rc == nil {
			return
		}

		rc.TransferLeaderTo(id, addr)
	})
}

// skipping Fuzz_Nosy_snapshot_Persist__ because parameters include func, chan, or unsupported interface: github.com/hashicorp/raft.SnapshotSink

func Fuzz_Nosy_snapshot_Release__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *snapshot
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Release()
	})
}

func Fuzz_Nosy_unsafeHeadTracker_Apply__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *unsafeHeadTracker
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var l *raft.Log
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		if t1 == nil || l == nil {
			return
		}

		t1.Apply(l)
	})
}

func Fuzz_Nosy_unsafeHeadTracker_Restore__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *unsafeHeadTracker
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var snapshot io.ReadCloser
		fill_err = tp.Fill(&snapshot)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Restore(snapshot)
	})
}

func Fuzz_Nosy_unsafeHeadTracker_Snapshot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *unsafeHeadTracker
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Snapshot()
	})
}

func Fuzz_Nosy_unsafeHeadTracker_UnsafeHead__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *unsafeHeadTracker
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.UnsafeHead()
	})
}

// skipping Fuzz_Nosy_Consensus_AddNonVoter__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-conductor/consensus.Consensus

// skipping Fuzz_Nosy_Consensus_AddVoter__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-conductor/consensus.Consensus

// skipping Fuzz_Nosy_Consensus_Addr__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-conductor/consensus.Consensus

// skipping Fuzz_Nosy_Consensus_ClusterMembership__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-conductor/consensus.Consensus

// skipping Fuzz_Nosy_Consensus_CommitUnsafePayload__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-conductor/consensus.Consensus

// skipping Fuzz_Nosy_Consensus_DemoteVoter__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-conductor/consensus.Consensus

// skipping Fuzz_Nosy_Consensus_LatestUnsafePayload__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-conductor/consensus.Consensus

// skipping Fuzz_Nosy_Consensus_Leader__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-conductor/consensus.Consensus

// skipping Fuzz_Nosy_Consensus_LeaderCh__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-conductor/consensus.Consensus

// skipping Fuzz_Nosy_Consensus_LeaderWithID__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-conductor/consensus.Consensus

// skipping Fuzz_Nosy_Consensus_RemoveServer__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-conductor/consensus.Consensus

// skipping Fuzz_Nosy_Consensus_ServerID__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-conductor/consensus.Consensus

// skipping Fuzz_Nosy_Consensus_Shutdown__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-conductor/consensus.Consensus

// skipping Fuzz_Nosy_Consensus_TransferLeader__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-conductor/consensus.Consensus

// skipping Fuzz_Nosy_Consensus_TransferLeaderTo__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-conductor/consensus.Consensus

func Fuzz_Nosy_ServerSuffrage_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s ServerSuffrage
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}

		s.String()
	})
}
