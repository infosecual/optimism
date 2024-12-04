package rpc

import (
	"context"
	"testing"

	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"
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

func Fuzz_Nosy_APIBackend_Active__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *APIBackend
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if api == nil {
			return
		}

		api.Active(_x2)
	})
}

func Fuzz_Nosy_APIBackend_AddServerAsNonvoter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *APIBackend
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
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
		if api == nil {
			return
		}

		api.AddServerAsNonvoter(ctx, id, addr, version)
	})
}

func Fuzz_Nosy_APIBackend_AddServerAsVoter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *APIBackend
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
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
		if api == nil {
			return
		}

		api.AddServerAsVoter(ctx, id, addr, version)
	})
}

func Fuzz_Nosy_APIBackend_ClusterMembership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *APIBackend
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if api == nil {
			return
		}

		api.ClusterMembership(ctx)
	})
}

func Fuzz_Nosy_APIBackend_CommitUnsafePayload__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *APIBackend
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var payload *eth.ExecutionPayloadEnvelope
		fill_err = tp.Fill(&payload)
		if fill_err != nil {
			return
		}
		if api == nil || payload == nil {
			return
		}

		api.CommitUnsafePayload(ctx, payload)
	})
}

func Fuzz_Nosy_APIBackend_Leader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *APIBackend
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if api == nil {
			return
		}

		api.Leader(ctx)
	})
}

func Fuzz_Nosy_APIBackend_LeaderOverridden__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *APIBackend
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if api == nil {
			return
		}

		api.LeaderOverridden(_x2)
	})
}

func Fuzz_Nosy_APIBackend_LeaderWithID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *APIBackend
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if api == nil {
			return
		}

		api.LeaderWithID(ctx)
	})
}

func Fuzz_Nosy_APIBackend_OverrideLeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *APIBackend
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var override bool
		fill_err = tp.Fill(&override)
		if fill_err != nil {
			return
		}
		if api == nil {
			return
		}

		api.OverrideLeader(_x2, override)
	})
}

func Fuzz_Nosy_APIBackend_Pause__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *APIBackend
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if api == nil {
			return
		}

		api.Pause(ctx)
	})
}

func Fuzz_Nosy_APIBackend_Paused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *APIBackend
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if api == nil {
			return
		}

		api.Paused(ctx)
	})
}

func Fuzz_Nosy_APIBackend_RemoveServer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *APIBackend
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
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
		if api == nil {
			return
		}

		api.RemoveServer(ctx, id, version)
	})
}

func Fuzz_Nosy_APIBackend_Resume__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *APIBackend
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if api == nil {
			return
		}

		api.Resume(ctx)
	})
}

func Fuzz_Nosy_APIBackend_SequencerHealthy__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *APIBackend
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if api == nil {
			return
		}

		api.SequencerHealthy(ctx)
	})
}

func Fuzz_Nosy_APIBackend_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *APIBackend
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if api == nil {
			return
		}

		api.Stop(ctx)
	})
}

func Fuzz_Nosy_APIBackend_Stopped__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *APIBackend
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if api == nil {
			return
		}

		api.Stopped(ctx)
	})
}

func Fuzz_Nosy_APIBackend_TransferLeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *APIBackend
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if api == nil {
			return
		}

		api.TransferLeader(ctx)
	})
}

func Fuzz_Nosy_APIBackend_TransferLeaderToServer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *APIBackend
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
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
		if api == nil {
			return
		}

		api.TransferLeaderToServer(ctx, id, addr)
	})
}

func Fuzz_Nosy_APIClient_Active__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
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

		c1 := NewAPIClient(c)
		c1.Active(ctx)
	})
}

func Fuzz_Nosy_APIClient_AddServerAsNonvoter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
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
		if c == nil {
			return
		}

		c1 := NewAPIClient(c)
		c1.AddServerAsNonvoter(ctx, id, addr, version)
	})
}

func Fuzz_Nosy_APIClient_AddServerAsVoter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
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
		if c == nil {
			return
		}

		c1 := NewAPIClient(c)
		c1.AddServerAsVoter(ctx, id, addr, version)
	})
}

func Fuzz_Nosy_APIClient_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c1 := NewAPIClient(c)
		c1.Close()
	})
}

func Fuzz_Nosy_APIClient_ClusterMembership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
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

		c1 := NewAPIClient(c)
		c1.ClusterMembership(ctx)
	})
}

func Fuzz_Nosy_APIClient_CommitUnsafePayload__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var payload *eth.ExecutionPayloadEnvelope
		fill_err = tp.Fill(&payload)
		if fill_err != nil {
			return
		}
		if c == nil || payload == nil {
			return
		}

		c1 := NewAPIClient(c)
		c1.CommitUnsafePayload(ctx, payload)
	})
}

func Fuzz_Nosy_APIClient_Leader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
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

		c1 := NewAPIClient(c)
		c1.Leader(ctx)
	})
}

func Fuzz_Nosy_APIClient_LeaderOverridden__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
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

		c1 := NewAPIClient(c)
		c1.LeaderOverridden(ctx)
	})
}

func Fuzz_Nosy_APIClient_LeaderWithID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
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

		c1 := NewAPIClient(c)
		c1.LeaderWithID(ctx)
	})
}

func Fuzz_Nosy_APIClient_OverrideLeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var override bool
		fill_err = tp.Fill(&override)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c1 := NewAPIClient(c)
		c1.OverrideLeader(ctx, override)
	})
}

func Fuzz_Nosy_APIClient_Pause__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
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

		c1 := NewAPIClient(c)
		c1.Pause(ctx)
	})
}

func Fuzz_Nosy_APIClient_Paused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
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

		c1 := NewAPIClient(c)
		c1.Paused(ctx)
	})
}

func Fuzz_Nosy_APIClient_RemoveServer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
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
		if c == nil {
			return
		}

		c1 := NewAPIClient(c)
		c1.RemoveServer(ctx, id, version)
	})
}

func Fuzz_Nosy_APIClient_Resume__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
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

		c1 := NewAPIClient(c)
		c1.Resume(ctx)
	})
}

func Fuzz_Nosy_APIClient_SequencerHealthy__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
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

		c1 := NewAPIClient(c)
		c1.SequencerHealthy(ctx)
	})
}

func Fuzz_Nosy_APIClient_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
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

		c1 := NewAPIClient(c)
		c1.Stop(ctx)
	})
}

func Fuzz_Nosy_APIClient_Stopped__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
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

		c1 := NewAPIClient(c)
		c1.Stopped(ctx)
	})
}

func Fuzz_Nosy_APIClient_TransferLeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
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

		c1 := NewAPIClient(c)
		c1.TransferLeader(ctx)
	})
}

func Fuzz_Nosy_APIClient_TransferLeaderToServer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
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
		if c == nil {
			return
		}

		c1 := NewAPIClient(c)
		c1.TransferLeaderToServer(ctx, id, addr)
	})
}

func Fuzz_Nosy_ExecutionMinerProxyBackend_SetMaxDASize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *ExecutionMinerProxyBackend
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var maxTxSize hexutil.Big
		fill_err = tp.Fill(&maxTxSize)
		if fill_err != nil {
			return
		}
		var maxBlockSize hexutil.Big
		fill_err = tp.Fill(&maxBlockSize)
		if fill_err != nil {
			return
		}
		if api == nil {
			return
		}

		api.SetMaxDASize(ctx, maxTxSize, maxBlockSize)
	})
}

func Fuzz_Nosy_ExecutionProxyBackend_GetBlockByNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *ExecutionProxyBackend
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var number rpc.BlockNumber
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		var fullTx bool
		fill_err = tp.Fill(&fullTx)
		if fill_err != nil {
			return
		}
		if api == nil {
			return
		}

		api.GetBlockByNumber(ctx, number, fullTx)
	})
}

func Fuzz_Nosy_NodeAdminProxyBackend_SequencerActive__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *NodeAdminProxyBackend
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if api == nil {
			return
		}

		api.SequencerActive(ctx)
	})
}

func Fuzz_Nosy_NodeProxyBackend_OutputAtBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *NodeProxyBackend
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var blockNumString string
		fill_err = tp.Fill(&blockNumString)
		if fill_err != nil {
			return
		}
		if api == nil {
			return
		}

		api.OutputAtBlock(ctx, blockNumString)
	})
}

func Fuzz_Nosy_NodeProxyBackend_RollupConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *NodeProxyBackend
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if api == nil {
			return
		}

		api.RollupConfig(ctx)
	})
}

func Fuzz_Nosy_NodeProxyBackend_SyncStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *NodeProxyBackend
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if api == nil {
			return
		}

		api.SyncStatus(ctx)
	})
}

// skipping Fuzz_Nosy_API_Active__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-conductor/rpc.API

// skipping Fuzz_Nosy_API_AddServerAsNonvoter__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-conductor/rpc.API

// skipping Fuzz_Nosy_API_AddServerAsVoter__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-conductor/rpc.API

// skipping Fuzz_Nosy_API_ClusterMembership__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-conductor/rpc.API

// skipping Fuzz_Nosy_API_CommitUnsafePayload__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-conductor/rpc.API

// skipping Fuzz_Nosy_API_Leader__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-conductor/rpc.API

// skipping Fuzz_Nosy_API_LeaderOverridden__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-conductor/rpc.API

// skipping Fuzz_Nosy_API_LeaderWithID__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-conductor/rpc.API

// skipping Fuzz_Nosy_API_OverrideLeader__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-conductor/rpc.API

// skipping Fuzz_Nosy_API_Pause__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-conductor/rpc.API

// skipping Fuzz_Nosy_API_Paused__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-conductor/rpc.API

// skipping Fuzz_Nosy_API_RemoveServer__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-conductor/rpc.API

// skipping Fuzz_Nosy_API_Resume__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-conductor/rpc.API

// skipping Fuzz_Nosy_API_SequencerHealthy__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-conductor/rpc.API

// skipping Fuzz_Nosy_API_Stop__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-conductor/rpc.API

// skipping Fuzz_Nosy_API_Stopped__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-conductor/rpc.API

// skipping Fuzz_Nosy_API_TransferLeader__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-conductor/rpc.API

// skipping Fuzz_Nosy_API_TransferLeaderToServer__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-conductor/rpc.API

// skipping Fuzz_Nosy_ExecutionMinerProxyAPI_SetMaxDASize__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-conductor/rpc.ExecutionMinerProxyAPI

// skipping Fuzz_Nosy_ExecutionProxyAPI_GetBlockByNumber__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-conductor/rpc.ExecutionProxyAPI

// skipping Fuzz_Nosy_NodeAdminProxyAPI_SequencerActive__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-conductor/rpc.NodeAdminProxyAPI

// skipping Fuzz_Nosy_NodeProxyAPI_OutputAtBlock__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-conductor/rpc.NodeProxyAPI

// skipping Fuzz_Nosy_NodeProxyAPI_RollupConfig__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-conductor/rpc.NodeProxyAPI

// skipping Fuzz_Nosy_NodeProxyAPI_SyncStatus__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-conductor/rpc.NodeProxyAPI

// skipping Fuzz_Nosy_conductor_AddServerAsNonvoter__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-conductor/rpc.conductor

// skipping Fuzz_Nosy_conductor_AddServerAsVoter__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-conductor/rpc.conductor

// skipping Fuzz_Nosy_conductor_ClusterMembership__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-conductor/rpc.conductor

// skipping Fuzz_Nosy_conductor_CommitUnsafePayload__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-conductor/rpc.conductor

// skipping Fuzz_Nosy_conductor_Leader__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-conductor/rpc.conductor

// skipping Fuzz_Nosy_conductor_LeaderOverridden__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-conductor/rpc.conductor

// skipping Fuzz_Nosy_conductor_LeaderWithID__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-conductor/rpc.conductor

// skipping Fuzz_Nosy_conductor_OverrideLeader__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-conductor/rpc.conductor

// skipping Fuzz_Nosy_conductor_Pause__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-conductor/rpc.conductor

// skipping Fuzz_Nosy_conductor_Paused__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-conductor/rpc.conductor

// skipping Fuzz_Nosy_conductor_RemoveServer__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-conductor/rpc.conductor

// skipping Fuzz_Nosy_conductor_Resume__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-conductor/rpc.conductor

// skipping Fuzz_Nosy_conductor_SequencerHealthy__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-conductor/rpc.conductor

// skipping Fuzz_Nosy_conductor_Stop__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-conductor/rpc.conductor

// skipping Fuzz_Nosy_conductor_Stopped__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-conductor/rpc.conductor

// skipping Fuzz_Nosy_conductor_TransferLeader__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-conductor/rpc.conductor

// skipping Fuzz_Nosy_conductor_TransferLeaderToServer__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-conductor/rpc.conductor

func Fuzz_Nosy_prefixRPC__(f *testing.F) {
	f.Fuzz(func(t *testing.T, method string) {
		prefixRPC(method)
	})
}
