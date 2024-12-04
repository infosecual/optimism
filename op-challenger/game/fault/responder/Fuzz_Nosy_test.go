package responder

import (
	"context"
	"testing"

	"github.com/ethereum-optimism/optimism/op-challenger/game/fault/types"
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

func Fuzz_Nosy_FaultResponder_CallResolve__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *FaultResponder
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.CallResolve(ctx)
	})
}

func Fuzz_Nosy_FaultResponder_CallResolveClaim__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *FaultResponder
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var claimIdx uint64
		fill_err = tp.Fill(&claimIdx)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.CallResolveClaim(ctx, claimIdx)
	})
}

func Fuzz_Nosy_FaultResponder_PerformAction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *FaultResponder
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var action types.Action
		fill_err = tp.Fill(&action)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.PerformAction(ctx, action)
	})
}

func Fuzz_Nosy_FaultResponder_Resolve__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *FaultResponder
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.Resolve()
	})
}

func Fuzz_Nosy_FaultResponder_ResolveClaims__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *FaultResponder
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var claimIdxs []uint64
		fill_err = tp.Fill(&claimIdxs)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.ResolveClaims(claimIdxs...)
	})
}

// skipping Fuzz_Nosy_GameContract_AttackTx__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/responder.GameContract

// skipping Fuzz_Nosy_GameContract_CallResolve__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/responder.GameContract

// skipping Fuzz_Nosy_GameContract_CallResolveClaim__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/responder.GameContract

// skipping Fuzz_Nosy_GameContract_ChallengeL2BlockNumberTx__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/responder.GameContract

// skipping Fuzz_Nosy_GameContract_DefendTx__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/responder.GameContract

// skipping Fuzz_Nosy_GameContract_ResolveClaimTx__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/responder.GameContract

// skipping Fuzz_Nosy_GameContract_ResolveTx__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/responder.GameContract

// skipping Fuzz_Nosy_GameContract_StepTx__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/responder.GameContract

// skipping Fuzz_Nosy_Oracle_GlobalDataExists__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/responder.Oracle

// skipping Fuzz_Nosy_TxSender_SendAndWaitSimple__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/responder.TxSender
