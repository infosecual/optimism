package disputegame

import (
	"context"
	"testing"
	"time"

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

func Fuzz_Nosy_ClaimHelper_AgreesWithOutputRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var game *OutputGameHelper
		fill_err = tp.Fill(&game)
		if fill_err != nil {
			return
		}
		var idx int64
		fill_err = tp.Fill(&idx)
		if fill_err != nil {
			return
		}
		var claim types.Claim
		fill_err = tp.Fill(&claim)
		if fill_err != nil {
			return
		}
		if game == nil {
			return
		}

		c := newClaimHelper(game, idx, claim)
		c.AgreesWithOutputRoot()
	})
}

// skipping Fuzz_Nosy_ClaimHelper_Attack__ because parameters include func, chan, or unsupported interface: []github.com/ethereum-optimism/optimism/op-e2e/e2eutils/disputegame.MoveOpt

// skipping Fuzz_Nosy_ClaimHelper_Defend__ because parameters include func, chan, or unsupported interface: []github.com/ethereum-optimism/optimism/op-e2e/e2eutils/disputegame.MoveOpt

func Fuzz_Nosy_ClaimHelper_Depth__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var game *OutputGameHelper
		fill_err = tp.Fill(&game)
		if fill_err != nil {
			return
		}
		var idx int64
		fill_err = tp.Fill(&idx)
		if fill_err != nil {
			return
		}
		var claim types.Claim
		fill_err = tp.Fill(&claim)
		if fill_err != nil {
			return
		}
		if game == nil {
			return
		}

		c := newClaimHelper(game, idx, claim)
		c.Depth()
	})
}

func Fuzz_Nosy_ClaimHelper_IsBottomGameRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var game *OutputGameHelper
		fill_err = tp.Fill(&game)
		if fill_err != nil {
			return
		}
		var idx int64
		fill_err = tp.Fill(&idx)
		if fill_err != nil {
			return
		}
		var claim types.Claim
		fill_err = tp.Fill(&claim)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if game == nil {
			return
		}

		c := newClaimHelper(game, idx, claim)
		c.IsBottomGameRoot(ctx)
	})
}

func Fuzz_Nosy_ClaimHelper_IsMaxDepth__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var game *OutputGameHelper
		fill_err = tp.Fill(&game)
		if fill_err != nil {
			return
		}
		var idx int64
		fill_err = tp.Fill(&idx)
		if fill_err != nil {
			return
		}
		var claim types.Claim
		fill_err = tp.Fill(&claim)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if game == nil {
			return
		}

		c := newClaimHelper(game, idx, claim)
		c.IsMaxDepth(ctx)
	})
}

func Fuzz_Nosy_ClaimHelper_IsOutputRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var game *OutputGameHelper
		fill_err = tp.Fill(&game)
		if fill_err != nil {
			return
		}
		var idx int64
		fill_err = tp.Fill(&idx)
		if fill_err != nil {
			return
		}
		var claim types.Claim
		fill_err = tp.Fill(&claim)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if game == nil {
			return
		}

		c := newClaimHelper(game, idx, claim)
		c.IsOutputRoot(ctx)
	})
}

func Fuzz_Nosy_ClaimHelper_IsOutputRootLeaf__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var game *OutputGameHelper
		fill_err = tp.Fill(&game)
		if fill_err != nil {
			return
		}
		var idx int64
		fill_err = tp.Fill(&idx)
		if fill_err != nil {
			return
		}
		var claim types.Claim
		fill_err = tp.Fill(&claim)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if game == nil {
			return
		}

		c := newClaimHelper(game, idx, claim)
		c.IsOutputRootLeaf(ctx)
	})
}

func Fuzz_Nosy_ClaimHelper_IsRootClaim__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var game *OutputGameHelper
		fill_err = tp.Fill(&game)
		if fill_err != nil {
			return
		}
		var idx int64
		fill_err = tp.Fill(&idx)
		if fill_err != nil {
			return
		}
		var claim types.Claim
		fill_err = tp.Fill(&claim)
		if fill_err != nil {
			return
		}
		if game == nil {
			return
		}

		c := newClaimHelper(game, idx, claim)
		c.IsRootClaim()
	})
}

func Fuzz_Nosy_ClaimHelper_RequireCorrectOutputRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var game *OutputGameHelper
		fill_err = tp.Fill(&game)
		if fill_err != nil {
			return
		}
		var idx int64
		fill_err = tp.Fill(&idx)
		if fill_err != nil {
			return
		}
		var claim types.Claim
		fill_err = tp.Fill(&claim)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if game == nil {
			return
		}

		c := newClaimHelper(game, idx, claim)
		c.RequireCorrectOutputRoot(ctx)
	})
}

func Fuzz_Nosy_ClaimHelper_RequireDifferentClaimValue__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var game *OutputGameHelper
		fill_err = tp.Fill(&game)
		if fill_err != nil {
			return
		}
		var idx int64
		fill_err = tp.Fill(&idx)
		if fill_err != nil {
			return
		}
		var claim types.Claim
		fill_err = tp.Fill(&claim)
		if fill_err != nil {
			return
		}
		var other *ClaimHelper
		fill_err = tp.Fill(&other)
		if fill_err != nil {
			return
		}
		if game == nil || other == nil {
			return
		}

		c := newClaimHelper(game, idx, claim)
		c.RequireDifferentClaimValue(other)
	})
}

func Fuzz_Nosy_ClaimHelper_RequireInvalidStatusCode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var game *OutputGameHelper
		fill_err = tp.Fill(&game)
		if fill_err != nil {
			return
		}
		var idx int64
		fill_err = tp.Fill(&idx)
		if fill_err != nil {
			return
		}
		var claim types.Claim
		fill_err = tp.Fill(&claim)
		if fill_err != nil {
			return
		}
		if game == nil {
			return
		}

		c := newClaimHelper(game, idx, claim)
		c.RequireInvalidStatusCode()
	})
}

func Fuzz_Nosy_ClaimHelper_RequireOnlyCounteredBy__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var game *OutputGameHelper
		fill_err = tp.Fill(&game)
		if fill_err != nil {
			return
		}
		var idx int64
		fill_err = tp.Fill(&idx)
		if fill_err != nil {
			return
		}
		var claim types.Claim
		fill_err = tp.Fill(&claim)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var expected []*ClaimHelper
		fill_err = tp.Fill(&expected)
		if fill_err != nil {
			return
		}
		if game == nil {
			return
		}

		c := newClaimHelper(game, idx, claim)
		c.RequireOnlyCounteredBy(ctx, expected...)
	})
}

func Fuzz_Nosy_ClaimHelper_WaitForCounterClaim__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var game *OutputGameHelper
		fill_err = tp.Fill(&game)
		if fill_err != nil {
			return
		}
		var idx int64
		fill_err = tp.Fill(&idx)
		if fill_err != nil {
			return
		}
		var claim types.Claim
		fill_err = tp.Fill(&claim)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var ignoreClaims []*ClaimHelper
		fill_err = tp.Fill(&ignoreClaims)
		if fill_err != nil {
			return
		}
		if game == nil {
			return
		}

		c := newClaimHelper(game, idx, claim)
		c.WaitForCounterClaim(ctx, ignoreClaims...)
	})
}

func Fuzz_Nosy_ClaimHelper_WaitForCountered__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var game *OutputGameHelper
		fill_err = tp.Fill(&game)
		if fill_err != nil {
			return
		}
		var idx int64
		fill_err = tp.Fill(&idx)
		if fill_err != nil {
			return
		}
		var claim types.Claim
		fill_err = tp.Fill(&claim)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if game == nil {
			return
		}

		c := newClaimHelper(game, idx, claim)
		c.WaitForCountered(ctx)
	})
}

func Fuzz_Nosy_DishonestHelper_ExhaustDishonestClaims__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *OutputGameHelper
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		var correctTrace *OutputHonestHelper
		fill_err = tp.Fill(&correctTrace)
		if fill_err != nil {
			return
		}
		var defender bool
		fill_err = tp.Fill(&defender)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var rootClaim *ClaimHelper
		fill_err = tp.Fill(&rootClaim)
		if fill_err != nil {
			return
		}
		if g == nil || correctTrace == nil || rootClaim == nil {
			return
		}

		d := newDishonestHelper(g, correctTrace, defender)
		d.ExhaustDishonestClaims(ctx, rootClaim)
	})
}

func Fuzz_Nosy_FactoryHelper_CreateBisectionGameExtraData__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *FactoryHelper
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var l2Node string
		fill_err = tp.Fill(&l2Node)
		if fill_err != nil {
			return
		}
		var l2BlockNumber uint64
		fill_err = tp.Fill(&l2BlockNumber)
		if fill_err != nil {
			return
		}
		var cfg *GameCfg
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if h == nil || cfg == nil {
			return
		}

		h.CreateBisectionGameExtraData(l2Node, l2BlockNumber, cfg)
	})
}

// skipping Fuzz_Nosy_FactoryHelper_GetL1Head__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/contracts.FaultDisputeGameContract

func Fuzz_Nosy_FactoryHelper_PreimageHelper__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *FactoryHelper
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.PreimageHelper(ctx)
	})
}

// skipping Fuzz_Nosy_FactoryHelper_StartChallenger__ because parameters include func, chan, or unsupported interface: []github.com/ethereum-optimism/optimism/op-e2e/e2eutils/challenger.Option

// skipping Fuzz_Nosy_FactoryHelper_StartOutputAlphabetGame__ because parameters include func, chan, or unsupported interface: []github.com/ethereum-optimism/optimism/op-e2e/e2eutils/disputegame.GameOpt

// skipping Fuzz_Nosy_FactoryHelper_StartOutputAlphabetGameWithCorrectRoot__ because parameters include func, chan, or unsupported interface: []github.com/ethereum-optimism/optimism/op-e2e/e2eutils/disputegame.GameOpt

// skipping Fuzz_Nosy_FactoryHelper_StartOutputCannonGame__ because parameters include func, chan, or unsupported interface: []github.com/ethereum-optimism/optimism/op-e2e/e2eutils/disputegame.GameOpt

// skipping Fuzz_Nosy_FactoryHelper_StartOutputCannonGameWithCorrectRoot__ because parameters include func, chan, or unsupported interface: []github.com/ethereum-optimism/optimism/op-e2e/e2eutils/disputegame.GameOpt

// skipping Fuzz_Nosy_FactoryHelper_StartPermissionedGame__ because parameters include func, chan, or unsupported interface: []github.com/ethereum-optimism/optimism/op-e2e/e2eutils/disputegame.GameOpt

func Fuzz_Nosy_FactoryHelper_WaitForBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *FactoryHelper
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var l2Node string
		fill_err = tp.Fill(&l2Node)
		if fill_err != nil {
			return
		}
		var l2BlockNumber uint64
		fill_err = tp.Fill(&l2BlockNumber)
		if fill_err != nil {
			return
		}
		var cfg *GameCfg
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if h == nil || cfg == nil {
			return
		}

		h.WaitForBlock(l2Node, l2BlockNumber, cfg)
	})
}

// skipping Fuzz_Nosy_FactoryHelper_startOutputCannonGameOfType__ because parameters include func, chan, or unsupported interface: []github.com/ethereum-optimism/optimism/op-e2e/e2eutils/disputegame.GameOpt

func Fuzz_Nosy_OutputAlphabetGameHelper_CreateDishonestHelper__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *OutputAlphabetGameHelper
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var l2Node string
		fill_err = tp.Fill(&l2Node)
		if fill_err != nil {
			return
		}
		var defender bool
		fill_err = tp.Fill(&defender)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.CreateDishonestHelper(ctx, l2Node, defender)
	})
}

func Fuzz_Nosy_OutputAlphabetGameHelper_CreateHonestActor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *OutputAlphabetGameHelper
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var l2Node string
		fill_err = tp.Fill(&l2Node)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.CreateHonestActor(ctx, l2Node)
	})
}

// skipping Fuzz_Nosy_OutputAlphabetGameHelper_StartChallenger__ because parameters include func, chan, or unsupported interface: []github.com/ethereum-optimism/optimism/op-e2e/e2eutils/challenger.Option

// skipping Fuzz_Nosy_OutputCannonGameHelper_ChallengeToPreimageLoad__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/trace/utils.PreimageOpt

// skipping Fuzz_Nosy_OutputCannonGameHelper_CreateHonestActor__ because parameters include func, chan, or unsupported interface: []github.com/ethereum-optimism/optimism/op-e2e/e2eutils/disputegame.HonestActorOpt

func Fuzz_Nosy_OutputCannonGameHelper_CreateStepLargePreimageLoadCheck__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *OutputCannonGameHelper
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var sender common.Address
		fill_err = tp.Fill(&sender)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.CreateStepLargePreimageLoadCheck(ctx, sender)
	})
}

func Fuzz_Nosy_OutputCannonGameHelper_CreateStepPreimageLoadCheck__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *OutputCannonGameHelper
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

		g.CreateStepPreimageLoadCheck(ctx)
	})
}

// skipping Fuzz_Nosy_OutputCannonGameHelper_StartChallenger__ because parameters include func, chan, or unsupported interface: []github.com/ethereum-optimism/optimism/op-e2e/e2eutils/challenger.Option

// skipping Fuzz_Nosy_OutputCannonGameHelper_VerifyPreimage__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-preimage.Key

// skipping Fuzz_Nosy_OutputCannonGameHelper_createCannonTraceProvider__ because parameters include func, chan, or unsupported interface: []github.com/ethereum-optimism/optimism/op-e2e/e2eutils/challenger.Option

func Fuzz_Nosy_OutputCannonGameHelper_defaultChallengerOptions__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *OutputCannonGameHelper
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.defaultChallengerOptions()
	})
}

// skipping Fuzz_Nosy_OutputGameHelper_Attack__ because parameters include func, chan, or unsupported interface: []github.com/ethereum-optimism/optimism/op-e2e/e2eutils/disputegame.MoveOpt

func Fuzz_Nosy_OutputGameHelper_AvailableCredit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *OutputGameHelper
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.AvailableCredit(ctx, addr)
	})
}

// skipping Fuzz_Nosy_OutputGameHelper_ChallengeClaim__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/e2eutils/disputegame.Mover

func Fuzz_Nosy_OutputGameHelper_ChallengePeriod__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *OutputGameHelper
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

		g.ChallengePeriod(ctx)
	})
}

func Fuzz_Nosy_OutputGameHelper_ChallengePeriodStartTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *OutputGameHelper
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var sender common.Address
		fill_err = tp.Fill(&sender)
		if fill_err != nil {
			return
		}
		var d4 *types.PreimageOracleData
		fill_err = tp.Fill(&d4)
		if fill_err != nil {
			return
		}
		if g == nil || d4 == nil {
			return
		}

		g.ChallengePeriodStartTime(ctx, sender, d4)
	})
}

func Fuzz_Nosy_OutputGameHelper_Credit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *OutputGameHelper
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.Credit(ctx, addr)
	})
}

func Fuzz_Nosy_OutputGameHelper_CreditUnlockDuration__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *OutputGameHelper
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

		g.CreditUnlockDuration(ctx)
	})
}

// skipping Fuzz_Nosy_OutputGameHelper_Defend__ because parameters include func, chan, or unsupported interface: []github.com/ethereum-optimism/optimism/op-e2e/e2eutils/disputegame.MoveOpt

// skipping Fuzz_Nosy_OutputGameHelper_DefendClaim__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/e2eutils/disputegame.Mover

func Fuzz_Nosy_OutputGameHelper_DisputeBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *OutputGameHelper
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var disputeBlockNum uint64
		fill_err = tp.Fill(&disputeBlockNum)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.DisputeBlock(ctx, disputeBlockNum)
	})
}

func Fuzz_Nosy_OutputGameHelper_DisputeLastBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *OutputGameHelper
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

		g.DisputeLastBlock(ctx)
	})
}

func Fuzz_Nosy_OutputGameHelper_ExecDepth__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *OutputGameHelper
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

		g.ExecDepth(ctx)
	})
}

func Fuzz_Nosy_OutputGameHelper_GameData__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *OutputGameHelper
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

		g.GameData(ctx)
	})
}

func Fuzz_Nosy_OutputGameHelper_GetClaimValue__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *OutputGameHelper
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var claimIdx int64
		fill_err = tp.Fill(&claimIdx)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.GetClaimValue(ctx, claimIdx)
	})
}

func Fuzz_Nosy_OutputGameHelper_GetL1Head__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *OutputGameHelper
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

		g.GetL1Head(ctx)
	})
}

func Fuzz_Nosy_OutputGameHelper_L2BlockNum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *OutputGameHelper
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

		g.L2BlockNum(ctx)
	})
}

func Fuzz_Nosy_OutputGameHelper_LogGameData__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *OutputGameHelper
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

		g.LogGameData(ctx)
	})
}

func Fuzz_Nosy_OutputGameHelper_MaxClockDuration__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *OutputGameHelper
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

		g.MaxClockDuration(ctx)
	})
}

func Fuzz_Nosy_OutputGameHelper_MaxDepth__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *OutputGameHelper
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

		g.MaxDepth(ctx)
	})
}

func Fuzz_Nosy_OutputGameHelper_Resolve__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *OutputGameHelper
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

		g.Resolve(ctx)
	})
}

func Fuzz_Nosy_OutputGameHelper_ResolveClaim__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *OutputGameHelper
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var claimIdx int64
		fill_err = tp.Fill(&claimIdx)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.ResolveClaim(ctx, claimIdx)
	})
}

func Fuzz_Nosy_OutputGameHelper_RootClaim__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *OutputGameHelper
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

		g.RootClaim(ctx)
	})
}

func Fuzz_Nosy_OutputGameHelper_SplitDepth__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *OutputGameHelper
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

		g.SplitDepth(ctx)
	})
}

func Fuzz_Nosy_OutputGameHelper_StartingBlockNum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *OutputGameHelper
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

		g.StartingBlockNum(ctx)
	})
}

func Fuzz_Nosy_OutputGameHelper_Status__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *OutputGameHelper
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

		g.Status(ctx)
	})
}

func Fuzz_Nosy_OutputGameHelper_StepFails__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *OutputGameHelper
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var claimIdx int64
		fill_err = tp.Fill(&claimIdx)
		if fill_err != nil {
			return
		}
		var isAttack bool
		fill_err = tp.Fill(&isAttack)
		if fill_err != nil {
			return
		}
		var stateData []byte
		fill_err = tp.Fill(&stateData)
		if fill_err != nil {
			return
		}
		var proof []byte
		fill_err = tp.Fill(&proof)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.StepFails(ctx, claimIdx, isAttack, stateData, proof)
	})
}

func Fuzz_Nosy_OutputGameHelper_UploadPreimage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *OutputGameHelper
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var d3 *types.PreimageOracleData
		fill_err = tp.Fill(&d3)
		if fill_err != nil {
			return
		}
		if g == nil || d3 == nil {
			return
		}

		g.UploadPreimage(ctx, d3)
	})
}

func Fuzz_Nosy_OutputGameHelper_WaitForAllClaimsCountered__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *OutputGameHelper
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

		g.WaitForAllClaimsCountered(ctx)
	})
}

func Fuzz_Nosy_OutputGameHelper_WaitForChallengePeriodStart__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *OutputGameHelper
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var sender common.Address
		fill_err = tp.Fill(&sender)
		if fill_err != nil {
			return
		}
		var d4 *types.PreimageOracleData
		fill_err = tp.Fill(&d4)
		if fill_err != nil {
			return
		}
		if g == nil || d4 == nil {
			return
		}

		g.WaitForChallengePeriodStart(ctx, sender, d4)
	})
}

func Fuzz_Nosy_OutputGameHelper_WaitForClaimAtDepth__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *OutputGameHelper
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var depth types.Depth
		fill_err = tp.Fill(&depth)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.WaitForClaimAtDepth(ctx, depth)
	})
}

func Fuzz_Nosy_OutputGameHelper_WaitForClaimAtMaxDepth__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *OutputGameHelper
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var countered bool
		fill_err = tp.Fill(&countered)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.WaitForClaimAtMaxDepth(ctx, countered)
	})
}

func Fuzz_Nosy_OutputGameHelper_WaitForClaimCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *OutputGameHelper
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var count int64
		fill_err = tp.Fill(&count)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.WaitForClaimCount(ctx, count)
	})
}

func Fuzz_Nosy_OutputGameHelper_WaitForCorrectOutputRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *OutputGameHelper
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var claimIdx int64
		fill_err = tp.Fill(&claimIdx)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.WaitForCorrectOutputRoot(ctx, claimIdx)
	})
}

func Fuzz_Nosy_OutputGameHelper_WaitForGameStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *OutputGameHelper
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var expected types.GameStatus
		fill_err = tp.Fill(&expected)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.WaitForGameStatus(ctx, expected)
	})
}

func Fuzz_Nosy_OutputGameHelper_WaitForInactivity__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *OutputGameHelper
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var numInactiveBlocks int
		fill_err = tp.Fill(&numInactiveBlocks)
		if fill_err != nil {
			return
		}
		var untilGameEnds bool
		fill_err = tp.Fill(&untilGameEnds)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.WaitForInactivity(ctx, numInactiveBlocks, untilGameEnds)
	})
}

func Fuzz_Nosy_OutputGameHelper_WaitForL2BlockNumberChallenged__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *OutputGameHelper
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

		g.WaitForL2BlockNumberChallenged(ctx)
	})
}

func Fuzz_Nosy_OutputGameHelper_WaitForNewClaim__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *OutputGameHelper
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var checkPoint int64
		fill_err = tp.Fill(&checkPoint)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.WaitForNewClaim(ctx, checkPoint)
	})
}

func Fuzz_Nosy_OutputGameHelper_WaitForNoAvailableCredit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *OutputGameHelper
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.WaitForNoAvailableCredit(ctx, addr)
	})
}

func Fuzz_Nosy_OutputGameHelper_WaitForPreimageInOracle__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *OutputGameHelper
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var d3 *types.PreimageOracleData
		fill_err = tp.Fill(&d3)
		if fill_err != nil {
			return
		}
		if g == nil || d3 == nil {
			return
		}

		g.WaitForPreimageInOracle(ctx, d3)
	})
}

func Fuzz_Nosy_OutputGameHelper_WethBalance__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *OutputGameHelper
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.WethBalance(ctx, addr)
	})
}

func Fuzz_Nosy_OutputGameHelper_correctOutputRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *OutputGameHelper
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var pos types.Position
		fill_err = tp.Fill(&pos)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.correctOutputRoot(ctx, pos)
	})
}

func Fuzz_Nosy_OutputGameHelper_getAllClaims__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *OutputGameHelper
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

		g.getAllClaims(ctx)
	})
}

func Fuzz_Nosy_OutputGameHelper_getClaim__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *OutputGameHelper
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var claimIdx int64
		fill_err = tp.Fill(&claimIdx)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.getClaim(ctx, claimIdx)
	})
}

func Fuzz_Nosy_OutputGameHelper_hasClaim__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *OutputGameHelper
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var parentIdx int64
		fill_err = tp.Fill(&parentIdx)
		if fill_err != nil {
			return
		}
		var pos types.Position
		fill_err = tp.Fill(&pos)
		if fill_err != nil {
			return
		}
		var value common.Hash
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.hasClaim(ctx, parentIdx, pos, value)
	})
}

// skipping Fuzz_Nosy_OutputGameHelper_moveCfg__ because parameters include func, chan, or unsupported interface: []github.com/ethereum-optimism/optimism/op-e2e/e2eutils/disputegame.MoveOpt

func Fuzz_Nosy_OutputGameHelper_oracle__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *OutputGameHelper
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

		g.oracle(ctx)
	})
}

// skipping Fuzz_Nosy_OutputGameHelper_waitForClaim__ because parameters include func, chan, or unsupported interface: func(claimIdx int64, claim github.com/ethereum-optimism/optimism/op-challenger/game/fault/types.Claim) bool

func Fuzz_Nosy_OutputGameHelper_waitForNewClaim__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *OutputGameHelper
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var checkPoint int64
		fill_err = tp.Fill(&checkPoint)
		if fill_err != nil {
			return
		}
		var timeout time.Duration
		fill_err = tp.Fill(&timeout)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.waitForNewClaim(ctx, checkPoint, timeout)
	})
}

// skipping Fuzz_Nosy_OutputGameHelper_waitForNoClaim__ because parameters include func, chan, or unsupported interface: func(claim github.com/ethereum-optimism/optimism/op-challenger/game/fault/types.Claim) bool

// skipping Fuzz_Nosy_OutputHonestHelper_Attack__ because parameters include func, chan, or unsupported interface: []github.com/ethereum-optimism/optimism/op-e2e/e2eutils/disputegame.MoveOpt

// skipping Fuzz_Nosy_OutputHonestHelper_AttackClaim__ because parameters include func, chan, or unsupported interface: []github.com/ethereum-optimism/optimism/op-e2e/e2eutils/disputegame.MoveOpt

// skipping Fuzz_Nosy_OutputHonestHelper_CounterClaim__ because parameters include func, chan, or unsupported interface: []github.com/ethereum-optimism/optimism/op-e2e/e2eutils/disputegame.MoveOpt

// skipping Fuzz_Nosy_OutputHonestHelper_Defend__ because parameters include func, chan, or unsupported interface: []github.com/ethereum-optimism/optimism/op-e2e/e2eutils/disputegame.MoveOpt

// skipping Fuzz_Nosy_OutputHonestHelper_DefendClaim__ because parameters include func, chan, or unsupported interface: []github.com/ethereum-optimism/optimism/op-e2e/e2eutils/disputegame.MoveOpt

func Fuzz_Nosy_OutputHonestHelper_StepClaimFails__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *OutputHonestHelper
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var claim *ClaimHelper
		fill_err = tp.Fill(&claim)
		if fill_err != nil {
			return
		}
		var isAttack bool
		fill_err = tp.Fill(&isAttack)
		if fill_err != nil {
			return
		}
		if h == nil || claim == nil {
			return
		}

		h.StepClaimFails(ctx, claim, isAttack)
	})
}

func Fuzz_Nosy_OutputHonestHelper_StepFails__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *OutputHonestHelper
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var claimIdx int64
		fill_err = tp.Fill(&claimIdx)
		if fill_err != nil {
			return
		}
		var isAttack bool
		fill_err = tp.Fill(&isAttack)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.StepFails(ctx, claimIdx, isAttack)
	})
}

func Fuzz_Nosy_OutputHonestHelper_loadState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *OutputHonestHelper
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var claimIdx int64
		fill_err = tp.Fill(&claimIdx)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.loadState(ctx, claimIdx)
	})
}

// skipping Fuzz_Nosy_DisputeSystem_AdvanceTime__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/e2eutils/disputegame.DisputeSystem

// skipping Fuzz_Nosy_DisputeSystem_AllocType__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/e2eutils/disputegame.DisputeSystem

// skipping Fuzz_Nosy_DisputeSystem_L1BeaconEndpoint__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/e2eutils/disputegame.DisputeSystem

// skipping Fuzz_Nosy_DisputeSystem_L1Deployments__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/e2eutils/disputegame.DisputeSystem

// skipping Fuzz_Nosy_DisputeSystem_L2Genesis__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/e2eutils/disputegame.DisputeSystem

// skipping Fuzz_Nosy_DisputeSystem_NodeClient__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/e2eutils/disputegame.DisputeSystem

// skipping Fuzz_Nosy_DisputeSystem_NodeEndpoint__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/e2eutils/disputegame.DisputeSystem

// skipping Fuzz_Nosy_DisputeSystem_RollupCfg__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/e2eutils/disputegame.DisputeSystem

// skipping Fuzz_Nosy_DisputeSystem_RollupClient__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/e2eutils/disputegame.DisputeSystem

// skipping Fuzz_Nosy_DisputeSystem_RollupEndpoint__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/e2eutils/disputegame.DisputeSystem

func Fuzz_Nosy_GameOpt_Apply__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *GameCfg
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		_x1 := WithUnsafeProposal()
		_x1.Apply(cfg)
	})
}

func Fuzz_Nosy_MoveOpt_Apply__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *moveCfg
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		_x1 := WithIgnoreDuplicates()
		_x1.Apply(cfg)
	})
}

// skipping Fuzz_Nosy_gameOptFn_Apply__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/e2eutils/disputegame.gameOptFn

// skipping Fuzz_Nosy_moveOptFn_Apply__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/e2eutils/disputegame.moveOptFn

func Fuzz_Nosy_containsClaim__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var claimIdx int64
		fill_err = tp.Fill(&claimIdx)
		if fill_err != nil {
			return
		}
		var haystack []*ClaimHelper
		fill_err = tp.Fill(&haystack)
		if fill_err != nil {
			return
		}

		containsClaim(claimIdx, haystack)
	})
}
