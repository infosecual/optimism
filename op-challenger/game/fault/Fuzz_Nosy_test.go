package fault

import (
	"context"
	"sync"
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

func Fuzz_Nosy_Agent_Act__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a *Agent
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

		a.Act(ctx)
	})
}

func Fuzz_Nosy_Agent_newGameFromContracts__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a *Agent
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

		a.newGameFromContracts(ctx)
	})
}

func Fuzz_Nosy_Agent_performAction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a *Agent
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var wg *sync.WaitGroup
		fill_err = tp.Fill(&wg)
		if fill_err != nil {
			return
		}
		var action types.Action
		fill_err = tp.Fill(&action)
		if fill_err != nil {
			return
		}
		if a == nil || wg == nil {
			return
		}

		a.performAction(ctx, wg, action)
	})
}

func Fuzz_Nosy_Agent_resolveClaims__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a *Agent
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

		a.resolveClaims(ctx)
	})
}

func Fuzz_Nosy_Agent_tryResolve__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a *Agent
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

		a.tryResolve(ctx)
	})
}

func Fuzz_Nosy_Agent_tryResolveClaims__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a *Agent
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

		a.tryResolveClaims(ctx)
	})
}

func Fuzz_Nosy_GamePlayer_ProgressGame__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *GamePlayer
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.ProgressGame(ctx)
	})
}

func Fuzz_Nosy_GamePlayer_Status__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *GamePlayer
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.Status()
	})
}

func Fuzz_Nosy_GamePlayer_ValidatePrestate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *GamePlayer
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.ValidatePrestate(ctx)
	})
}

func Fuzz_Nosy_GamePlayer_logGameStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *GamePlayer
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var status types.GameStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.logGameStatus(ctx, status)
	})
}

func Fuzz_Nosy_PrestateValidator_Validate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v *PrestateValidator
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if v == nil {
			return
		}

		v.Validate(ctx)
	})
}

// skipping Fuzz_Nosy_RegisterTask_Register__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault.Registry

func Fuzz_Nosy_syncStatusValidator_ValidateNodeSynced__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *syncStatusValidator
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var gameL1Head eth.BlockID
		fill_err = tp.Fill(&gameL1Head)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.ValidateNodeSynced(ctx, gameL1Head)
	})
}

// skipping Fuzz_Nosy_ClaimLoader_GetAllClaims__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault.ClaimLoader

// skipping Fuzz_Nosy_ClaimLoader_IsL2BlockNumberChallenged__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault.ClaimLoader

// skipping Fuzz_Nosy_GameContract_GetL1Head__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault.GameContract

// skipping Fuzz_Nosy_GameContract_GetMaxClockDuration__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault.GameContract

// skipping Fuzz_Nosy_GameContract_GetMaxGameDepth__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault.GameContract

// skipping Fuzz_Nosy_GameContract_GetOracle__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault.GameContract

// skipping Fuzz_Nosy_GameContract_GetStatus__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault.GameContract

// skipping Fuzz_Nosy_GameInfo_GetClaimCount__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault.GameInfo

// skipping Fuzz_Nosy_GameInfo_GetStatus__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault.GameInfo

// skipping Fuzz_Nosy_L1HeaderSource_HeaderByHash__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault.L1HeaderSource

// skipping Fuzz_Nosy_OracleRegistry_RegisterOracle__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault.OracleRegistry

// skipping Fuzz_Nosy_PrestateSource_PrestatePath__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault.PrestateSource

// skipping Fuzz_Nosy_Registry_RegisterBondContract__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault.Registry

// skipping Fuzz_Nosy_Registry_RegisterGameType__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault.Registry

// skipping Fuzz_Nosy_Responder_CallResolve__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault.Responder

// skipping Fuzz_Nosy_Responder_CallResolveClaim__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault.Responder

// skipping Fuzz_Nosy_Responder_PerformAction__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault.Responder

// skipping Fuzz_Nosy_Responder_Resolve__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault.Responder

// skipping Fuzz_Nosy_Responder_ResolveClaims__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault.Responder

// skipping Fuzz_Nosy_SyncStatusProvider_SyncStatus__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault.SyncStatusProvider

// skipping Fuzz_Nosy_SyncValidator_ValidateNodeSynced__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault.SyncValidator

// skipping Fuzz_Nosy_TxSender_From__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault.TxSender

// skipping Fuzz_Nosy_TxSender_SendAndWaitSimple__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault.TxSender

// skipping Fuzz_Nosy_Validator_Validate__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault.Validator

// skipping Fuzz_Nosy_cachePrestates__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/trace/vm.StateConverter
