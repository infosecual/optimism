package mocks

import (
	"testing"

	"github.com/ethereum-optimism/optimism/op-conductor/consensus"
	eth "github.com/ethereum-optimism/optimism/op-service/eth"
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

func Fuzz_Nosy_Consensus_AddNonVoter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *Consensus
		fill_err = tp.Fill(&_m)
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
		if _m == nil {
			return
		}

		_m.AddNonVoter(id, addr, version)
	})
}

func Fuzz_Nosy_Consensus_AddVoter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *Consensus
		fill_err = tp.Fill(&_m)
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
		if _m == nil {
			return
		}

		_m.AddVoter(id, addr, version)
	})
}

func Fuzz_Nosy_Consensus_Addr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *Consensus
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.Addr()
	})
}

func Fuzz_Nosy_Consensus_ClusterMembership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *Consensus
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.ClusterMembership()
	})
}

func Fuzz_Nosy_Consensus_CommitUnsafePayload__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *Consensus
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var payload *eth.ExecutionPayloadEnvelope
		fill_err = tp.Fill(&payload)
		if fill_err != nil {
			return
		}
		if _m == nil || payload == nil {
			return
		}

		_m.CommitUnsafePayload(payload)
	})
}

func Fuzz_Nosy_Consensus_DemoteVoter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *Consensus
		fill_err = tp.Fill(&_m)
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
		if _m == nil {
			return
		}

		_m.DemoteVoter(id, version)
	})
}

func Fuzz_Nosy_Consensus_EXPECT__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *Consensus
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.EXPECT()
	})
}

func Fuzz_Nosy_Consensus_LatestUnsafePayload__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *Consensus
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.LatestUnsafePayload()
	})
}

func Fuzz_Nosy_Consensus_Leader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *Consensus
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.Leader()
	})
}

func Fuzz_Nosy_Consensus_LeaderCh__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *Consensus
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.LeaderCh()
	})
}

func Fuzz_Nosy_Consensus_LeaderWithID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *Consensus
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.LeaderWithID()
	})
}

func Fuzz_Nosy_Consensus_RemoveServer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *Consensus
		fill_err = tp.Fill(&_m)
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
		if _m == nil {
			return
		}

		_m.RemoveServer(id, version)
	})
}

func Fuzz_Nosy_Consensus_ServerID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *Consensus
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.ServerID()
	})
}

func Fuzz_Nosy_Consensus_Shutdown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *Consensus
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.Shutdown()
	})
}

func Fuzz_Nosy_Consensus_TransferLeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *Consensus
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.TransferLeader()
	})
}

func Fuzz_Nosy_Consensus_TransferLeaderTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *Consensus
		fill_err = tp.Fill(&_m)
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
		if _m == nil {
			return
		}

		_m.TransferLeaderTo(id, addr)
	})
}

// skipping Fuzz_Nosy_Consensus_AddNonVoter_Call_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_Consensus_AddNonVoter_Call_Run__ because parameters include func, chan, or unsupported interface: func(id string, addr string, version uint64)

// skipping Fuzz_Nosy_Consensus_AddNonVoter_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(string, string, uint64) error

// skipping Fuzz_Nosy_Consensus_AddVoter_Call_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_Consensus_AddVoter_Call_Run__ because parameters include func, chan, or unsupported interface: func(id string, addr string, version uint64)

// skipping Fuzz_Nosy_Consensus_AddVoter_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(string, string, uint64) error

func Fuzz_Nosy_Consensus_Addr_Call_Return__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _c *Consensus_Addr_Call
		fill_err = tp.Fill(&_c)
		if fill_err != nil {
			return
		}
		var _a0 string
		fill_err = tp.Fill(&_a0)
		if fill_err != nil {
			return
		}
		if _c == nil {
			return
		}

		_c.Return(_a0)
	})
}

// skipping Fuzz_Nosy_Consensus_Addr_Call_Run__ because parameters include func, chan, or unsupported interface: func()

// skipping Fuzz_Nosy_Consensus_Addr_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func() string

// skipping Fuzz_Nosy_Consensus_ClusterMembership_Call_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_Consensus_ClusterMembership_Call_Run__ because parameters include func, chan, or unsupported interface: func()

// skipping Fuzz_Nosy_Consensus_ClusterMembership_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func() (*github.com/ethereum-optimism/optimism/op-conductor/consensus.ClusterMembership, error)

// skipping Fuzz_Nosy_Consensus_CommitUnsafePayload_Call_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_Consensus_CommitUnsafePayload_Call_Run__ because parameters include func, chan, or unsupported interface: func(payload *github.com/ethereum-optimism/optimism/op-service/eth.ExecutionPayloadEnvelope)

// skipping Fuzz_Nosy_Consensus_CommitUnsafePayload_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(*github.com/ethereum-optimism/optimism/op-service/eth.ExecutionPayloadEnvelope) error

// skipping Fuzz_Nosy_Consensus_DemoteVoter_Call_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_Consensus_DemoteVoter_Call_Run__ because parameters include func, chan, or unsupported interface: func(id string, version uint64)

// skipping Fuzz_Nosy_Consensus_DemoteVoter_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(string, uint64) error

// skipping Fuzz_Nosy_Consensus_Expecter_AddNonVoter__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_Consensus_Expecter_AddVoter__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_Consensus_Expecter_Addr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _e *Consensus_Expecter
		fill_err = tp.Fill(&_e)
		if fill_err != nil {
			return
		}
		if _e == nil {
			return
		}

		_e.Addr()
	})
}

func Fuzz_Nosy_Consensus_Expecter_ClusterMembership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _e *Consensus_Expecter
		fill_err = tp.Fill(&_e)
		if fill_err != nil {
			return
		}
		if _e == nil {
			return
		}

		_e.ClusterMembership()
	})
}

// skipping Fuzz_Nosy_Consensus_Expecter_CommitUnsafePayload__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_Consensus_Expecter_DemoteVoter__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_Consensus_Expecter_LatestUnsafePayload__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _e *Consensus_Expecter
		fill_err = tp.Fill(&_e)
		if fill_err != nil {
			return
		}
		if _e == nil {
			return
		}

		_e.LatestUnsafePayload()
	})
}

func Fuzz_Nosy_Consensus_Expecter_Leader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _e *Consensus_Expecter
		fill_err = tp.Fill(&_e)
		if fill_err != nil {
			return
		}
		if _e == nil {
			return
		}

		_e.Leader()
	})
}

func Fuzz_Nosy_Consensus_Expecter_LeaderCh__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _e *Consensus_Expecter
		fill_err = tp.Fill(&_e)
		if fill_err != nil {
			return
		}
		if _e == nil {
			return
		}

		_e.LeaderCh()
	})
}

func Fuzz_Nosy_Consensus_Expecter_LeaderWithID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _e *Consensus_Expecter
		fill_err = tp.Fill(&_e)
		if fill_err != nil {
			return
		}
		if _e == nil {
			return
		}

		_e.LeaderWithID()
	})
}

// skipping Fuzz_Nosy_Consensus_Expecter_RemoveServer__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_Consensus_Expecter_ServerID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _e *Consensus_Expecter
		fill_err = tp.Fill(&_e)
		if fill_err != nil {
			return
		}
		if _e == nil {
			return
		}

		_e.ServerID()
	})
}

func Fuzz_Nosy_Consensus_Expecter_Shutdown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _e *Consensus_Expecter
		fill_err = tp.Fill(&_e)
		if fill_err != nil {
			return
		}
		if _e == nil {
			return
		}

		_e.Shutdown()
	})
}

func Fuzz_Nosy_Consensus_Expecter_TransferLeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _e *Consensus_Expecter
		fill_err = tp.Fill(&_e)
		if fill_err != nil {
			return
		}
		if _e == nil {
			return
		}

		_e.TransferLeader()
	})
}

// skipping Fuzz_Nosy_Consensus_Expecter_TransferLeaderTo__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_Consensus_LatestUnsafePayload_Call_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_Consensus_LatestUnsafePayload_Call_Run__ because parameters include func, chan, or unsupported interface: func()

// skipping Fuzz_Nosy_Consensus_LatestUnsafePayload_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func() (*github.com/ethereum-optimism/optimism/op-service/eth.ExecutionPayloadEnvelope, error)

// skipping Fuzz_Nosy_Consensus_LeaderCh_Call_Return__ because parameters include func, chan, or unsupported interface: <-chan bool

// skipping Fuzz_Nosy_Consensus_LeaderCh_Call_Run__ because parameters include func, chan, or unsupported interface: func()

// skipping Fuzz_Nosy_Consensus_LeaderCh_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func() <-chan bool

func Fuzz_Nosy_Consensus_LeaderWithID_Call_Return__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _c *Consensus_LeaderWithID_Call
		fill_err = tp.Fill(&_c)
		if fill_err != nil {
			return
		}
		var _a0 *consensus.ServerInfo
		fill_err = tp.Fill(&_a0)
		if fill_err != nil {
			return
		}
		if _c == nil || _a0 == nil {
			return
		}

		_c.Return(_a0)
	})
}

// skipping Fuzz_Nosy_Consensus_LeaderWithID_Call_Run__ because parameters include func, chan, or unsupported interface: func()

// skipping Fuzz_Nosy_Consensus_LeaderWithID_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func() *github.com/ethereum-optimism/optimism/op-conductor/consensus.ServerInfo

func Fuzz_Nosy_Consensus_Leader_Call_Return__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _c *Consensus_Leader_Call
		fill_err = tp.Fill(&_c)
		if fill_err != nil {
			return
		}
		var _a0 bool
		fill_err = tp.Fill(&_a0)
		if fill_err != nil {
			return
		}
		if _c == nil {
			return
		}

		_c.Return(_a0)
	})
}

// skipping Fuzz_Nosy_Consensus_Leader_Call_Run__ because parameters include func, chan, or unsupported interface: func()

// skipping Fuzz_Nosy_Consensus_Leader_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func() bool

// skipping Fuzz_Nosy_Consensus_RemoveServer_Call_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_Consensus_RemoveServer_Call_Run__ because parameters include func, chan, or unsupported interface: func(id string, version uint64)

// skipping Fuzz_Nosy_Consensus_RemoveServer_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(string, uint64) error

func Fuzz_Nosy_Consensus_ServerID_Call_Return__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _c *Consensus_ServerID_Call
		fill_err = tp.Fill(&_c)
		if fill_err != nil {
			return
		}
		var _a0 string
		fill_err = tp.Fill(&_a0)
		if fill_err != nil {
			return
		}
		if _c == nil {
			return
		}

		_c.Return(_a0)
	})
}

// skipping Fuzz_Nosy_Consensus_ServerID_Call_Run__ because parameters include func, chan, or unsupported interface: func()

// skipping Fuzz_Nosy_Consensus_ServerID_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func() string

// skipping Fuzz_Nosy_Consensus_Shutdown_Call_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_Consensus_Shutdown_Call_Run__ because parameters include func, chan, or unsupported interface: func()

// skipping Fuzz_Nosy_Consensus_Shutdown_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func() error

// skipping Fuzz_Nosy_Consensus_TransferLeaderTo_Call_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_Consensus_TransferLeaderTo_Call_Run__ because parameters include func, chan, or unsupported interface: func(id string, addr string)

// skipping Fuzz_Nosy_Consensus_TransferLeaderTo_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(string, string) error

// skipping Fuzz_Nosy_Consensus_TransferLeader_Call_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_Consensus_TransferLeader_Call_Run__ because parameters include func, chan, or unsupported interface: func()

// skipping Fuzz_Nosy_Consensus_TransferLeader_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func() error
