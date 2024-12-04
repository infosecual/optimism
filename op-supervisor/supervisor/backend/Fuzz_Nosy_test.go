package backend

import (
	"context"
	"testing"

	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum-optimism/optimism/op-supervisor/config"
	"github.com/ethereum-optimism/optimism/op-supervisor/supervisor/types"
	"github.com/ethereum/go-ethereum/common"
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

func Fuzz_Nosy_MockBackend_AddL2RPC__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var rpc string
		fill_err = tp.Fill(&rpc)
		if fill_err != nil {
			return
		}

		m := NewMockBackend()
		m.AddL2RPC(ctx, rpc)
	})
}

func Fuzz_Nosy_MockBackend_CheckMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var identifier types.Identifier
		fill_err = tp.Fill(&identifier)
		if fill_err != nil {
			return
		}
		var payloadHash common.Hash
		fill_err = tp.Fill(&payloadHash)
		if fill_err != nil {
			return
		}

		m := NewMockBackend()
		m.CheckMessage(identifier, payloadHash)
	})
}

func Fuzz_Nosy_MockBackend_CheckMessages__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var messages []types.Message
		fill_err = tp.Fill(&messages)
		if fill_err != nil {
			return
		}
		var minSafety types.SafetyLevel
		fill_err = tp.Fill(&minSafety)
		if fill_err != nil {
			return
		}

		m := NewMockBackend()
		m.CheckMessages(messages, minSafety)
	})
}

func Fuzz_Nosy_MockBackend_CrossDerivedFrom__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var chainID types.ChainID
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var derived eth.BlockID
		fill_err = tp.Fill(&derived)
		if fill_err != nil {
			return
		}

		m := NewMockBackend()
		m.CrossDerivedFrom(ctx, chainID, derived)
	})
}

func Fuzz_Nosy_MockBackend_Finalized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var chainID types.ChainID
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}

		m := NewMockBackend()
		m.Finalized(ctx, chainID)
	})
}

func Fuzz_Nosy_MockBackend_SafeView__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var chainID types.ChainID
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var safe types.ReferenceView
		fill_err = tp.Fill(&safe)
		if fill_err != nil {
			return
		}

		m := NewMockBackend()
		m.SafeView(ctx, chainID, safe)
	})
}

func Fuzz_Nosy_MockBackend_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		m := NewMockBackend()
		m.Start(ctx)
	})
}

func Fuzz_Nosy_MockBackend_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		m := NewMockBackend()
		m.Stop(ctx)
	})
}

func Fuzz_Nosy_MockBackend_UnsafeView__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var chainID types.ChainID
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var unsafe types.ReferenceView
		fill_err = tp.Fill(&unsafe)
		if fill_err != nil {
			return
		}

		m := NewMockBackend()
		m.UnsafeView(ctx, chainID, unsafe)
	})
}

func Fuzz_Nosy_MockBackend_UpdateFinalizedL1__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var chainID types.ChainID
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var finalized eth.L1BlockRef
		fill_err = tp.Fill(&finalized)
		if fill_err != nil {
			return
		}

		m := NewMockBackend()
		m.UpdateFinalizedL1(ctx, chainID, finalized)
	})
}

func Fuzz_Nosy_MockBackend_UpdateLocalSafe__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var chainID types.ChainID
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var derivedFrom eth.L1BlockRef
		fill_err = tp.Fill(&derivedFrom)
		if fill_err != nil {
			return
		}
		var lastDerived eth.L1BlockRef
		fill_err = tp.Fill(&lastDerived)
		if fill_err != nil {
			return
		}

		m := NewMockBackend()
		m.UpdateLocalSafe(ctx, chainID, derivedFrom, lastDerived)
	})
}

func Fuzz_Nosy_MockBackend_UpdateLocalUnsafe__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var chainID types.ChainID
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var head eth.L1BlockRef
		fill_err = tp.Fill(&head)
		if fill_err != nil {
			return
		}

		m := NewMockBackend()
		m.UpdateLocalUnsafe(ctx, chainID, head)
	})
}

func Fuzz_Nosy_SupervisorBackend_AddL2RPC__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var su *SupervisorBackend
		fill_err = tp.Fill(&su)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var rpc string
		fill_err = tp.Fill(&rpc)
		if fill_err != nil {
			return
		}
		if su == nil {
			return
		}

		su.AddL2RPC(ctx, rpc)
	})
}

// skipping Fuzz_Nosy_SupervisorBackend_AttachProcessorSource__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/processors.Source

func Fuzz_Nosy_SupervisorBackend_CheckMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var su *SupervisorBackend
		fill_err = tp.Fill(&su)
		if fill_err != nil {
			return
		}
		var identifier types.Identifier
		fill_err = tp.Fill(&identifier)
		if fill_err != nil {
			return
		}
		var payloadHash common.Hash
		fill_err = tp.Fill(&payloadHash)
		if fill_err != nil {
			return
		}
		if su == nil {
			return
		}

		su.CheckMessage(identifier, payloadHash)
	})
}

func Fuzz_Nosy_SupervisorBackend_CheckMessages__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var su *SupervisorBackend
		fill_err = tp.Fill(&su)
		if fill_err != nil {
			return
		}
		var messages []types.Message
		fill_err = tp.Fill(&messages)
		if fill_err != nil {
			return
		}
		var minSafety types.SafetyLevel
		fill_err = tp.Fill(&minSafety)
		if fill_err != nil {
			return
		}
		if su == nil {
			return
		}

		su.CheckMessages(messages, minSafety)
	})
}

func Fuzz_Nosy_SupervisorBackend_CrossDerivedFrom__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var su *SupervisorBackend
		fill_err = tp.Fill(&su)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var chainID types.ChainID
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var derived eth.BlockID
		fill_err = tp.Fill(&derived)
		if fill_err != nil {
			return
		}
		if su == nil {
			return
		}

		su.CrossDerivedFrom(ctx, chainID, derived)
	})
}

func Fuzz_Nosy_SupervisorBackend_DependencySet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var su *SupervisorBackend
		fill_err = tp.Fill(&su)
		if fill_err != nil {
			return
		}
		if su == nil {
			return
		}

		su.DependencySet()
	})
}

func Fuzz_Nosy_SupervisorBackend_Finalized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var su *SupervisorBackend
		fill_err = tp.Fill(&su)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var chainID types.ChainID
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		if su == nil {
			return
		}

		su.Finalized(ctx, chainID)
	})
}

func Fuzz_Nosy_SupervisorBackend_SafeView__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var su *SupervisorBackend
		fill_err = tp.Fill(&su)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var chainID types.ChainID
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var safe types.ReferenceView
		fill_err = tp.Fill(&safe)
		if fill_err != nil {
			return
		}
		if su == nil {
			return
		}

		su.SafeView(ctx, chainID, safe)
	})
}

func Fuzz_Nosy_SupervisorBackend_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var su *SupervisorBackend
		fill_err = tp.Fill(&su)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if su == nil {
			return
		}

		su.Start(ctx)
	})
}

func Fuzz_Nosy_SupervisorBackend_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var su *SupervisorBackend
		fill_err = tp.Fill(&su)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if su == nil {
			return
		}

		su.Stop(ctx)
	})
}

func Fuzz_Nosy_SupervisorBackend_SyncCrossSafe__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var su *SupervisorBackend
		fill_err = tp.Fill(&su)
		if fill_err != nil {
			return
		}
		var chainID types.ChainID
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		if su == nil {
			return
		}

		su.SyncCrossSafe(chainID)
	})
}

func Fuzz_Nosy_SupervisorBackend_SyncCrossUnsafe__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var su *SupervisorBackend
		fill_err = tp.Fill(&su)
		if fill_err != nil {
			return
		}
		var chainID types.ChainID
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		if su == nil {
			return
		}

		su.SyncCrossUnsafe(chainID)
	})
}

func Fuzz_Nosy_SupervisorBackend_SyncEvents__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var su *SupervisorBackend
		fill_err = tp.Fill(&su)
		if fill_err != nil {
			return
		}
		var chainID types.ChainID
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		if su == nil {
			return
		}

		su.SyncEvents(chainID)
	})
}

func Fuzz_Nosy_SupervisorBackend_UnsafeView__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var su *SupervisorBackend
		fill_err = tp.Fill(&su)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var chainID types.ChainID
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var unsafe types.ReferenceView
		fill_err = tp.Fill(&unsafe)
		if fill_err != nil {
			return
		}
		if su == nil {
			return
		}

		su.UnsafeView(ctx, chainID, unsafe)
	})
}

func Fuzz_Nosy_SupervisorBackend_UpdateFinalizedL1__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var su *SupervisorBackend
		fill_err = tp.Fill(&su)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var chainID types.ChainID
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var finalized eth.L1BlockRef
		fill_err = tp.Fill(&finalized)
		if fill_err != nil {
			return
		}
		if su == nil {
			return
		}

		su.UpdateFinalizedL1(ctx, chainID, finalized)
	})
}

func Fuzz_Nosy_SupervisorBackend_UpdateLocalSafe__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var su *SupervisorBackend
		fill_err = tp.Fill(&su)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var chainID types.ChainID
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var derivedFrom eth.L1BlockRef
		fill_err = tp.Fill(&derivedFrom)
		if fill_err != nil {
			return
		}
		var lastDerived eth.L1BlockRef
		fill_err = tp.Fill(&lastDerived)
		if fill_err != nil {
			return
		}
		if su == nil {
			return
		}

		su.UpdateLocalSafe(ctx, chainID, derivedFrom, lastDerived)
	})
}

func Fuzz_Nosy_SupervisorBackend_UpdateLocalUnsafe__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var su *SupervisorBackend
		fill_err = tp.Fill(&su)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var chainID types.ChainID
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var head eth.L1BlockRef
		fill_err = tp.Fill(&head)
		if fill_err != nil {
			return
		}
		if su == nil {
			return
		}

		su.UpdateLocalUnsafe(ctx, chainID, head)
	})
}

func Fuzz_Nosy_SupervisorBackend_attachRPC__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var su *SupervisorBackend
		fill_err = tp.Fill(&su)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var rpc string
		fill_err = tp.Fill(&rpc)
		if fill_err != nil {
			return
		}
		if su == nil {
			return
		}

		su.attachRPC(ctx, rpc)
	})
}

func Fuzz_Nosy_SupervisorBackend_initResources__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var su *SupervisorBackend
		fill_err = tp.Fill(&su)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var cfg *config.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if su == nil || cfg == nil {
			return
		}

		su.initResources(ctx, cfg)
	})
}

func Fuzz_Nosy_SupervisorBackend_onIndexedLocalUnsafeData__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var su *SupervisorBackend
		fill_err = tp.Fill(&su)
		if fill_err != nil {
			return
		}
		if su == nil {
			return
		}

		su.onIndexedLocalUnsafeData()
	})
}

func Fuzz_Nosy_SupervisorBackend_onNewLocalSafeData__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var su *SupervisorBackend
		fill_err = tp.Fill(&su)
		if fill_err != nil {
			return
		}
		if su == nil {
			return
		}

		su.onNewLocalSafeData()
	})
}

func Fuzz_Nosy_SupervisorBackend_openChainDBs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var su *SupervisorBackend
		fill_err = tp.Fill(&su)
		if fill_err != nil {
			return
		}
		var chainID types.ChainID
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		if su == nil {
			return
		}

		su.openChainDBs(chainID)
	})
}

func Fuzz_Nosy_chainMetrics_CacheAdd__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *chainMetrics
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var label string
		fill_err = tp.Fill(&label)
		if fill_err != nil {
			return
		}
		var cacheSize int
		fill_err = tp.Fill(&cacheSize)
		if fill_err != nil {
			return
		}
		var evicted bool
		fill_err = tp.Fill(&evicted)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.CacheAdd(label, cacheSize, evicted)
	})
}

func Fuzz_Nosy_chainMetrics_CacheGet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *chainMetrics
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var label string
		fill_err = tp.Fill(&label)
		if fill_err != nil {
			return
		}
		var hit bool
		fill_err = tp.Fill(&hit)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.CacheGet(label, hit)
	})
}

func Fuzz_Nosy_chainMetrics_RecordDBEntryCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *chainMetrics
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var kind string
		fill_err = tp.Fill(&kind)
		if fill_err != nil {
			return
		}
		var count int64
		fill_err = tp.Fill(&count)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.RecordDBEntryCount(kind, count)
	})
}

func Fuzz_Nosy_chainMetrics_RecordDBSearchEntriesRead__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *chainMetrics
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var count int64
		fill_err = tp.Fill(&count)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.RecordDBSearchEntriesRead(count)
	})
}

// skipping Fuzz_Nosy_Metrics_CacheAdd__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend.Metrics

// skipping Fuzz_Nosy_Metrics_CacheGet__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend.Metrics

// skipping Fuzz_Nosy_Metrics_RecordDBEntryCount__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend.Metrics

// skipping Fuzz_Nosy_Metrics_RecordDBSearchEntriesRead__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend.Metrics

// skipping Fuzz_Nosy_clientForL2__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/log.Logger
