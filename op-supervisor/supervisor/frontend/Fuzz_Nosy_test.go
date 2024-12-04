package frontend

import (
	"context"
	"testing"

	"github.com/ethereum-optimism/optimism/op-service/eth"
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

func Fuzz_Nosy_AdminFrontend_AddL2RPC__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a *AdminFrontend
		fill_err = tp.Fill(&a)
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
		if a == nil {
			return
		}

		a.AddL2RPC(ctx, rpc)
	})
}

func Fuzz_Nosy_AdminFrontend_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a *AdminFrontend
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if a == nil {
			return
		}

		a.Start(ctx)
	})
}

func Fuzz_Nosy_AdminFrontend_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a *AdminFrontend
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if a == nil {
			return
		}

		a.Stop(ctx)
	})
}

func Fuzz_Nosy_QueryFrontend_CheckMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var q *QueryFrontend
		fill_err = tp.Fill(&q)
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
		if q == nil {
			return
		}

		q.CheckMessage(identifier, payloadHash)
	})
}

func Fuzz_Nosy_QueryFrontend_CheckMessages__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var q *QueryFrontend
		fill_err = tp.Fill(&q)
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
		if q == nil {
			return
		}

		q.CheckMessages(messages, minSafety)
	})
}

func Fuzz_Nosy_QueryFrontend_CrossDerivedFrom__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var q *QueryFrontend
		fill_err = tp.Fill(&q)
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
		if q == nil {
			return
		}

		q.CrossDerivedFrom(ctx, chainID, derived)
	})
}

func Fuzz_Nosy_QueryFrontend_Finalized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var q *QueryFrontend
		fill_err = tp.Fill(&q)
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
		if q == nil {
			return
		}

		q.Finalized(ctx, chainID)
	})
}

func Fuzz_Nosy_QueryFrontend_SafeView__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var q *QueryFrontend
		fill_err = tp.Fill(&q)
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
		if q == nil {
			return
		}

		q.SafeView(ctx, chainID, safe)
	})
}

func Fuzz_Nosy_QueryFrontend_UnsafeView__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var q *QueryFrontend
		fill_err = tp.Fill(&q)
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
		if q == nil {
			return
		}

		q.UnsafeView(ctx, chainID, unsafe)
	})
}

func Fuzz_Nosy_UpdatesFrontend_UpdateFinalizedL1__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var u *UpdatesFrontend
		fill_err = tp.Fill(&u)
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
		if u == nil {
			return
		}

		u.UpdateFinalizedL1(ctx, chainID, finalized)
	})
}

func Fuzz_Nosy_UpdatesFrontend_UpdateLocalSafe__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var u *UpdatesFrontend
		fill_err = tp.Fill(&u)
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
		if u == nil {
			return
		}

		u.UpdateLocalSafe(ctx, chainID, derivedFrom, lastDerived)
	})
}

func Fuzz_Nosy_UpdatesFrontend_UpdateLocalUnsafe__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var u *UpdatesFrontend
		fill_err = tp.Fill(&u)
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
		if u == nil {
			return
		}

		u.UpdateLocalUnsafe(ctx, chainID, head)
	})
}

// skipping Fuzz_Nosy_AdminBackend_AddL2RPC__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/frontend.AdminBackend

// skipping Fuzz_Nosy_AdminBackend_Start__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/frontend.AdminBackend

// skipping Fuzz_Nosy_AdminBackend_Stop__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/frontend.AdminBackend

// skipping Fuzz_Nosy_QueryBackend_CheckMessage__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/frontend.QueryBackend

// skipping Fuzz_Nosy_QueryBackend_CheckMessages__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/frontend.QueryBackend

// skipping Fuzz_Nosy_QueryBackend_CrossDerivedFrom__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/frontend.QueryBackend

// skipping Fuzz_Nosy_QueryBackend_Finalized__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/frontend.QueryBackend

// skipping Fuzz_Nosy_QueryBackend_SafeView__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/frontend.QueryBackend

// skipping Fuzz_Nosy_QueryBackend_UnsafeView__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/frontend.QueryBackend

// skipping Fuzz_Nosy_UpdatesBackend_UpdateFinalizedL1__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/frontend.UpdatesBackend

// skipping Fuzz_Nosy_UpdatesBackend_UpdateLocalSafe__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/frontend.UpdatesBackend

// skipping Fuzz_Nosy_UpdatesBackend_UpdateLocalUnsafe__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/frontend.UpdatesBackend
