package solver

import (
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

// skipping Fuzz_Nosy_GameSolver_AgreeWithRootClaim__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/types.Game

// skipping Fuzz_Nosy_GameSolver_CalculateNextActions__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/types.Game

// skipping Fuzz_Nosy_GameSolver_calculateMove__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/types.Game

// skipping Fuzz_Nosy_GameSolver_calculateStep__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/types.Game

// skipping Fuzz_Nosy_claimSolver_AttemptStep__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/types.Game

// skipping Fuzz_Nosy_claimSolver_NextMove__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/types.Game

// skipping Fuzz_Nosy_claimSolver_agreeWithClaim__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/types.Game

// skipping Fuzz_Nosy_claimSolver_attack__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/types.Game

// skipping Fuzz_Nosy_claimSolver_defend__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/types.Game

// skipping Fuzz_Nosy_claimSolver_shouldCounter__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/types.Game

func Fuzz_Nosy_honestClaimTracker_AddHonestClaim__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var parent types.Claim
		fill_err = tp.Fill(&parent)
		if fill_err != nil {
			return
		}
		var claim types.Claim
		fill_err = tp.Fill(&claim)
		if fill_err != nil {
			return
		}

		a := newHonestClaimTracker()
		a.AddHonestClaim(parent, claim)
	})
}

func Fuzz_Nosy_honestClaimTracker_HonestCounter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var parent types.Claim
		fill_err = tp.Fill(&parent)
		if fill_err != nil {
			return
		}

		a := newHonestClaimTracker()
		a.HonestCounter(parent)
	})
}

func Fuzz_Nosy_honestClaimTracker_IsHonest__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var claim types.Claim
		fill_err = tp.Fill(&claim)
		if fill_err != nil {
			return
		}

		a := newHonestClaimTracker()
		a.IsHonest(claim)
	})
}

// skipping Fuzz_Nosy_actor_Apply__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/solver.actor

// skipping Fuzz_Nosy_actorFn_Apply__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/solver.actorFn

// skipping Fuzz_Nosy_builderFn_Apply__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/solver.builderFn

// skipping Fuzz_Nosy_printClaim__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/types.Game
