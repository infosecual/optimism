package conductor

import (
	"context"
	"testing"

	"github.com/ethereum-optimism/optimism/op-service/eth"
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

func Fuzz_Nosy_Config_Check__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *Config
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.Check()
	})
}

func Fuzz_Nosy_HealthCheckConfig_Check__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *HealthCheckConfig
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.Check()
	})
}

func Fuzz_Nosy_OpConductor_AddServerAsNonvoter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var oc *OpConductor
		fill_err = tp.Fill(&oc)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
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
		if oc == nil {
			return
		}

		oc.AddServerAsNonvoter(_x2, id, addr, version)
	})
}

func Fuzz_Nosy_OpConductor_AddServerAsVoter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var oc *OpConductor
		fill_err = tp.Fill(&oc)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
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
		if oc == nil {
			return
		}

		oc.AddServerAsVoter(_x2, id, addr, version)
	})
}

func Fuzz_Nosy_OpConductor_ClusterMembership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var oc *OpConductor
		fill_err = tp.Fill(&oc)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if oc == nil {
			return
		}

		oc.ClusterMembership(_x2)
	})
}

func Fuzz_Nosy_OpConductor_CommitUnsafePayload__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var oc *OpConductor
		fill_err = tp.Fill(&oc)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var payload *eth.ExecutionPayloadEnvelope
		fill_err = tp.Fill(&payload)
		if fill_err != nil {
			return
		}
		if oc == nil || payload == nil {
			return
		}

		oc.CommitUnsafePayload(_x2, payload)
	})
}

func Fuzz_Nosy_OpConductor_ConsensusEndpoint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var oc *OpConductor
		fill_err = tp.Fill(&oc)
		if fill_err != nil {
			return
		}
		if oc == nil {
			return
		}

		oc.ConsensusEndpoint()
	})
}

func Fuzz_Nosy_OpConductor_HTTPEndpoint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var oc *OpConductor
		fill_err = tp.Fill(&oc)
		if fill_err != nil {
			return
		}
		if oc == nil {
			return
		}

		oc.HTTPEndpoint()
	})
}

func Fuzz_Nosy_OpConductor_LatestUnsafePayload__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var oc *OpConductor
		fill_err = tp.Fill(&oc)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if oc == nil {
			return
		}

		oc.LatestUnsafePayload(_x2)
	})
}

func Fuzz_Nosy_OpConductor_Leader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var oc *OpConductor
		fill_err = tp.Fill(&oc)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if oc == nil {
			return
		}

		oc.Leader(ctx)
	})
}

func Fuzz_Nosy_OpConductor_LeaderOverridden__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var oc *OpConductor
		fill_err = tp.Fill(&oc)
		if fill_err != nil {
			return
		}
		if oc == nil {
			return
		}

		oc.LeaderOverridden()
	})
}

func Fuzz_Nosy_OpConductor_LeaderWithID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var oc *OpConductor
		fill_err = tp.Fill(&oc)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if oc == nil {
			return
		}

		oc.LeaderWithID(ctx)
	})
}

func Fuzz_Nosy_OpConductor_OverrideLeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var oc *OpConductor
		fill_err = tp.Fill(&oc)
		if fill_err != nil {
			return
		}
		var override bool
		fill_err = tp.Fill(&override)
		if fill_err != nil {
			return
		}
		if oc == nil {
			return
		}

		oc.OverrideLeader(override)
	})
}

func Fuzz_Nosy_OpConductor_Pause__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var oc *OpConductor
		fill_err = tp.Fill(&oc)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if oc == nil {
			return
		}

		oc.Pause(ctx)
	})
}

func Fuzz_Nosy_OpConductor_Paused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var oc *OpConductor
		fill_err = tp.Fill(&oc)
		if fill_err != nil {
			return
		}
		if oc == nil {
			return
		}

		oc.Paused()
	})
}

func Fuzz_Nosy_OpConductor_RemoveServer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var oc *OpConductor
		fill_err = tp.Fill(&oc)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
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
		if oc == nil {
			return
		}

		oc.RemoveServer(_x2, id, version)
	})
}

func Fuzz_Nosy_OpConductor_Resume__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var oc *OpConductor
		fill_err = tp.Fill(&oc)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if oc == nil {
			return
		}

		oc.Resume(ctx)
	})
}

func Fuzz_Nosy_OpConductor_SequencerHealthy__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var oc *OpConductor
		fill_err = tp.Fill(&oc)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if oc == nil {
			return
		}

		oc.SequencerHealthy(_x2)
	})
}

func Fuzz_Nosy_OpConductor_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var oc *OpConductor
		fill_err = tp.Fill(&oc)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if oc == nil {
			return
		}

		oc.Start(ctx)
	})
}

func Fuzz_Nosy_OpConductor_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var oc *OpConductor
		fill_err = tp.Fill(&oc)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if oc == nil {
			return
		}

		oc.Stop(ctx)
	})
}

func Fuzz_Nosy_OpConductor_Stopped__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var oc *OpConductor
		fill_err = tp.Fill(&oc)
		if fill_err != nil {
			return
		}
		if oc == nil {
			return
		}

		oc.Stopped()
	})
}

func Fuzz_Nosy_OpConductor_TransferLeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var oc *OpConductor
		fill_err = tp.Fill(&oc)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if oc == nil {
			return
		}

		oc.TransferLeader(_x2)
	})
}

func Fuzz_Nosy_OpConductor_TransferLeaderToServer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var oc *OpConductor
		fill_err = tp.Fill(&oc)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
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
		if oc == nil {
			return
		}

		oc.TransferLeaderToServer(_x2, id, addr)
	})
}

func Fuzz_Nosy_OpConductor_action__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var oc *OpConductor
		fill_err = tp.Fill(&oc)
		if fill_err != nil {
			return
		}
		if oc == nil {
			return
		}

		oc.action()
	})
}

func Fuzz_Nosy_OpConductor_compareUnsafeHead__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var oc *OpConductor
		fill_err = tp.Fill(&oc)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if oc == nil {
			return
		}

		oc.compareUnsafeHead(ctx)
	})
}

// skipping Fuzz_Nosy_OpConductor_handleHealthUpdate__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_OpConductor_handleLeaderUpdate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var oc *OpConductor
		fill_err = tp.Fill(&oc)
		if fill_err != nil {
			return
		}
		var leader bool
		fill_err = tp.Fill(&leader)
		if fill_err != nil {
			return
		}
		if oc == nil {
			return
		}

		oc.handleLeaderUpdate(leader)
	})
}

func Fuzz_Nosy_OpConductor_init__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *OpConductor
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.init(ctx)
	})
}

func Fuzz_Nosy_OpConductor_initConsensus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *OpConductor
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.initConsensus(ctx)
	})
}

func Fuzz_Nosy_OpConductor_initHealthMonitor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *OpConductor
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.initHealthMonitor(ctx)
	})
}

func Fuzz_Nosy_OpConductor_initRPCServer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var oc *OpConductor
		fill_err = tp.Fill(&oc)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if oc == nil {
			return
		}

		oc.initRPCServer(ctx)
	})
}

func Fuzz_Nosy_OpConductor_initSequencerControl__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *OpConductor
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.initSequencerControl(ctx)
	})
}

func Fuzz_Nosy_OpConductor_loop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var oc *OpConductor
		fill_err = tp.Fill(&oc)
		if fill_err != nil {
			return
		}
		if oc == nil {
			return
		}

		oc.loop()
	})
}

func Fuzz_Nosy_OpConductor_loopAction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var oc *OpConductor
		fill_err = tp.Fill(&oc)
		if fill_err != nil {
			return
		}
		if oc == nil {
			return
		}

		oc.loopAction()
	})
}

func Fuzz_Nosy_OpConductor_queueAction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var oc *OpConductor
		fill_err = tp.Fill(&oc)
		if fill_err != nil {
			return
		}
		if oc == nil {
			return
		}

		oc.queueAction()
	})
}

func Fuzz_Nosy_OpConductor_startSequencer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var oc *OpConductor
		fill_err = tp.Fill(&oc)
		if fill_err != nil {
			return
		}
		if oc == nil {
			return
		}

		oc.startSequencer()
	})
}

func Fuzz_Nosy_OpConductor_stopSequencer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var oc *OpConductor
		fill_err = tp.Fill(&oc)
		if fill_err != nil {
			return
		}
		if oc == nil {
			return
		}

		oc.stopSequencer()
	})
}

func Fuzz_Nosy_OpConductor_transferLeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var oc *OpConductor
		fill_err = tp.Fill(&oc)
		if fill_err != nil {
			return
		}
		if oc == nil {
			return
		}

		oc.transferLeader()
	})
}

func Fuzz_Nosy_OpConductor_updateSequencerActiveStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var oc *OpConductor
		fill_err = tp.Fill(&oc)
		if fill_err != nil {
			return
		}
		if oc == nil {
			return
		}

		oc.updateSequencerActiveStatus()
	})
}

func Fuzz_Nosy_state_Equal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var leader bool
		fill_err = tp.Fill(&leader)
		if fill_err != nil {
			return
		}
		var healthy bool
		fill_err = tp.Fill(&healthy)
		if fill_err != nil {
			return
		}
		var active bool
		fill_err = tp.Fill(&active)
		if fill_err != nil {
			return
		}
		var other *state
		fill_err = tp.Fill(&other)
		if fill_err != nil {
			return
		}
		if other == nil {
			return
		}

		s := NewState(leader, healthy, active)
		s.Equal(other)
	})
}

func Fuzz_Nosy_state_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, leader bool, healthy bool, active bool) {
		s := NewState(leader, healthy, active)
		s.String()
	})
}
