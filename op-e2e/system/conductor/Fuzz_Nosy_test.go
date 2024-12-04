package conductor

import (
	"context"
	"testing"

	"github.com/ethereum-optimism/optimism/op-conductor/consensus"
	"github.com/ethereum-optimism/optimism/op-e2e/system/e2esys"
	"github.com/ethereum-optimism/optimism/op-node/rollup"
	"github.com/ethereum-optimism/optimism/op-service/sources"
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

func Fuzz_Nosy_conductor_ConsensusEndpoint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var serverID string
		fill_err = tp.Fill(&serverID)
		if fill_err != nil {
			return
		}
		var dir string
		fill_err = tp.Fill(&dir)
		if fill_err != nil {
			return
		}
		var nodeRPC string
		fill_err = tp.Fill(&nodeRPC)
		if fill_err != nil {
			return
		}
		var engineRPC string
		fill_err = tp.Fill(&engineRPC)
		if fill_err != nil {
			return
		}
		var bootstrap bool
		fill_err = tp.Fill(&bootstrap)
		if fill_err != nil {
			return
		}
		var rollupCfg rollup.Config
		fill_err = tp.Fill(&rollupCfg)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		c, err := setupConductor(t1, serverID, dir, nodeRPC, engineRPC, bootstrap, rollupCfg)
		if err != nil {
			return
		}
		c.ConsensusEndpoint()
	})
}

func Fuzz_Nosy_conductor_RPCEndpoint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var serverID string
		fill_err = tp.Fill(&serverID)
		if fill_err != nil {
			return
		}
		var dir string
		fill_err = tp.Fill(&dir)
		if fill_err != nil {
			return
		}
		var nodeRPC string
		fill_err = tp.Fill(&nodeRPC)
		if fill_err != nil {
			return
		}
		var engineRPC string
		fill_err = tp.Fill(&engineRPC)
		if fill_err != nil {
			return
		}
		var bootstrap bool
		fill_err = tp.Fill(&bootstrap)
		if fill_err != nil {
			return
		}
		var rollupCfg rollup.Config
		fill_err = tp.Fill(&rollupCfg)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		c, err := setupConductor(t1, serverID, dir, nodeRPC, engineRPC, bootstrap, rollupCfg)
		if err != nil {
			return
		}
		c.RPCEndpoint()
	})
}

func Fuzz_Nosy_conductorActive__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var con *conductor
		fill_err = tp.Fill(&con)
		if fill_err != nil {
			return
		}
		if t1 == nil || con == nil {
			return
		}

		conductorActive(t1, ctx, con)
	})
}

func Fuzz_Nosy_conductorResumed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var con *conductor
		fill_err = tp.Fill(&con)
		if fill_err != nil {
			return
		}
		if t1 == nil || con == nil {
			return
		}

		conductorResumed(t1, ctx, con)
	})
}

func Fuzz_Nosy_conductorStopped__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var con *conductor
		fill_err = tp.Fill(&con)
		if fill_err != nil {
			return
		}
		if t1 == nil || con == nil {
			return
		}

		conductorStopped(t1, ctx, con)
	})
}

func Fuzz_Nosy_ensureOnlyOneLeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var sys *e2esys.System
		fill_err = tp.Fill(&sys)
		if fill_err != nil {
			return
		}
		var conductors map[string]*conductor
		fill_err = tp.Fill(&conductors)
		if fill_err != nil {
			return
		}
		if t1 == nil || sys == nil {
			return
		}

		ensureOnlyOneLeader(t1, sys, conductors)
	})
}

func Fuzz_Nosy_findFollower__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var conductors map[string]*conductor
		fill_err = tp.Fill(&conductors)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		findFollower(t1, conductors)
	})
}

func Fuzz_Nosy_findLeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var conductors map[string]*conductor
		fill_err = tp.Fill(&conductors)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		findLeader(t1, conductors)
	})
}

func Fuzz_Nosy_healthy__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var con *conductor
		fill_err = tp.Fill(&con)
		if fill_err != nil {
			return
		}
		if t1 == nil || con == nil {
			return
		}

		healthy(t1, ctx, con)
	})
}

func Fuzz_Nosy_leader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var con *conductor
		fill_err = tp.Fill(&con)
		if fill_err != nil {
			return
		}
		if t1 == nil || con == nil {
			return
		}

		leader(t1, ctx, con)
	})
}

func Fuzz_Nosy_memberIDs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var membership *consensus.ClusterMembership
		fill_err = tp.Fill(&membership)
		if fill_err != nil {
			return
		}
		if membership == nil {
			return
		}

		memberIDs(membership)
	})
}

func Fuzz_Nosy_sequencerActive__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var rollupClient *sources.RollupClient
		fill_err = tp.Fill(&rollupClient)
		if fill_err != nil {
			return
		}
		if t1 == nil || rollupClient == nil {
			return
		}

		sequencerActive(t1, ctx, rollupClient)
	})
}

func Fuzz_Nosy_setupBatcher__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var sys *e2esys.System
		fill_err = tp.Fill(&sys)
		if fill_err != nil {
			return
		}
		var conductors map[string]*conductor
		fill_err = tp.Fill(&conductors)
		if fill_err != nil {
			return
		}
		if t1 == nil || sys == nil {
			return
		}

		setupBatcher(t1, sys, conductors)
	})
}

func Fuzz_Nosy_setupHAInfra__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		setupHAInfra(t1, ctx)
	})
}

func Fuzz_Nosy_setupSequencerFailoverTest__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		setupSequencerFailoverTest(t1)
	})
}

func Fuzz_Nosy_waitForLeadershipChange__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var prev *conductor
		fill_err = tp.Fill(&prev)
		if fill_err != nil {
			return
		}
		var prevID string
		fill_err = tp.Fill(&prevID)
		if fill_err != nil {
			return
		}
		var conductors map[string]*conductor
		fill_err = tp.Fill(&conductors)
		if fill_err != nil {
			return
		}
		var sys *e2esys.System
		fill_err = tp.Fill(&sys)
		if fill_err != nil {
			return
		}
		if t1 == nil || prev == nil || sys == nil {
			return
		}

		waitForLeadershipChange(t1, prev, prevID, conductors, sys)
	})
}
