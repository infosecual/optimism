package test

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum-optimism/optimism/op-challenger/game/fault/types"
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

func Fuzz_Nosy_AlphabetWithProofProvider_GetL2BlockNumberChallenge__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *AlphabetWithProofProvider
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.GetL2BlockNumberChallenge(_x2)
	})
}

func Fuzz_Nosy_AlphabetWithProofProvider_GetStepData__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a *AlphabetWithProofProvider
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var i types.Position
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
		if a == nil {
			return
		}

		a.GetStepData(ctx, i)
	})
}

// skipping Fuzz_Nosy_ClaimBuilder_AttackClaim__ because parameters include func, chan, or unsupported interface: []github.com/ethereum-optimism/optimism/op-challenger/game/fault/test.ClaimOpt

func Fuzz_Nosy_ClaimBuilder_CorrectClaimAtPosition__(f *testing.F) {
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
		var startingL2BlockNumber *big.Int
		fill_err = tp.Fill(&startingL2BlockNumber)
		if fill_err != nil {
			return
		}
		var maxDepth types.Depth
		fill_err = tp.Fill(&maxDepth)
		if fill_err != nil {
			return
		}
		var pos types.Position
		fill_err = tp.Fill(&pos)
		if fill_err != nil {
			return
		}
		if t1 == nil || startingL2BlockNumber == nil {
			return
		}

		c := NewAlphabetClaimBuilder(t1, startingL2BlockNumber, maxDepth)
		c.CorrectClaimAtPosition(pos)
	})
}

func Fuzz_Nosy_ClaimBuilder_CorrectOracleData__(f *testing.F) {
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
		var startingL2BlockNumber *big.Int
		fill_err = tp.Fill(&startingL2BlockNumber)
		if fill_err != nil {
			return
		}
		var maxDepth types.Depth
		fill_err = tp.Fill(&maxDepth)
		if fill_err != nil {
			return
		}
		var idx *big.Int
		fill_err = tp.Fill(&idx)
		if fill_err != nil {
			return
		}
		if t1 == nil || startingL2BlockNumber == nil || idx == nil {
			return
		}

		c := NewAlphabetClaimBuilder(t1, startingL2BlockNumber, maxDepth)
		c.CorrectOracleData(idx)
	})
}

func Fuzz_Nosy_ClaimBuilder_CorrectPreState__(f *testing.F) {
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
		var startingL2BlockNumber *big.Int
		fill_err = tp.Fill(&startingL2BlockNumber)
		if fill_err != nil {
			return
		}
		var maxDepth types.Depth
		fill_err = tp.Fill(&maxDepth)
		if fill_err != nil {
			return
		}
		var idx *big.Int
		fill_err = tp.Fill(&idx)
		if fill_err != nil {
			return
		}
		if t1 == nil || startingL2BlockNumber == nil || idx == nil {
			return
		}

		c := NewAlphabetClaimBuilder(t1, startingL2BlockNumber, maxDepth)
		c.CorrectPreState(idx)
	})
}

func Fuzz_Nosy_ClaimBuilder_CorrectProofData__(f *testing.F) {
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
		var startingL2BlockNumber *big.Int
		fill_err = tp.Fill(&startingL2BlockNumber)
		if fill_err != nil {
			return
		}
		var maxDepth types.Depth
		fill_err = tp.Fill(&maxDepth)
		if fill_err != nil {
			return
		}
		var idx *big.Int
		fill_err = tp.Fill(&idx)
		if fill_err != nil {
			return
		}
		if t1 == nil || startingL2BlockNumber == nil || idx == nil {
			return
		}

		c := NewAlphabetClaimBuilder(t1, startingL2BlockNumber, maxDepth)
		c.CorrectProofData(idx)
	})
}

func Fuzz_Nosy_ClaimBuilder_CorrectTraceProvider__(f *testing.F) {
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
		var startingL2BlockNumber *big.Int
		fill_err = tp.Fill(&startingL2BlockNumber)
		if fill_err != nil {
			return
		}
		var maxDepth types.Depth
		fill_err = tp.Fill(&maxDepth)
		if fill_err != nil {
			return
		}
		if t1 == nil || startingL2BlockNumber == nil {
			return
		}

		c := NewAlphabetClaimBuilder(t1, startingL2BlockNumber, maxDepth)
		c.CorrectTraceProvider()
	})
}

// skipping Fuzz_Nosy_ClaimBuilder_CreateLeafClaim__ because parameters include func, chan, or unsupported interface: []github.com/ethereum-optimism/optimism/op-challenger/game/fault/test.ClaimOpt

// skipping Fuzz_Nosy_ClaimBuilder_CreateRootClaim__ because parameters include func, chan, or unsupported interface: []github.com/ethereum-optimism/optimism/op-challenger/game/fault/test.ClaimOpt

// skipping Fuzz_Nosy_ClaimBuilder_DefendClaim__ because parameters include func, chan, or unsupported interface: []github.com/ethereum-optimism/optimism/op-challenger/game/fault/test.ClaimOpt

// skipping Fuzz_Nosy_ClaimBuilder_GameBuilder__ because parameters include func, chan, or unsupported interface: []github.com/ethereum-optimism/optimism/op-challenger/game/fault/test.ClaimOpt

// skipping Fuzz_Nosy_ClaimBuilder_Seq__ because parameters include func, chan, or unsupported interface: []github.com/ethereum-optimism/optimism/op-challenger/game/fault/test.ClaimOpt

// skipping Fuzz_Nosy_ClaimBuilder_claim__ because parameters include func, chan, or unsupported interface: []github.com/ethereum-optimism/optimism/op-challenger/game/fault/test.ClaimOpt

func Fuzz_Nosy_ClaimBuilder_incorrectClaim__(f *testing.F) {
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
		var startingL2BlockNumber *big.Int
		fill_err = tp.Fill(&startingL2BlockNumber)
		if fill_err != nil {
			return
		}
		var maxDepth types.Depth
		fill_err = tp.Fill(&maxDepth)
		if fill_err != nil {
			return
		}
		var pos types.Position
		fill_err = tp.Fill(&pos)
		if fill_err != nil {
			return
		}
		if t1 == nil || startingL2BlockNumber == nil {
			return
		}

		c := NewAlphabetClaimBuilder(t1, startingL2BlockNumber, maxDepth)
		c.incorrectClaim(pos)
	})
}

func Fuzz_Nosy_GameBuilder_Seq__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *GameBuilder
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.Seq()
	})
}

func Fuzz_Nosy_GameBuilder_SeqFrom__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *GameBuilder
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		var claim types.Claim
		fill_err = tp.Fill(&claim)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.SeqFrom(claim)
	})
}

// skipping Fuzz_Nosy_GameBuilderSeq_Attack__ because parameters include func, chan, or unsupported interface: []github.com/ethereum-optimism/optimism/op-challenger/game/fault/test.ClaimOpt

// skipping Fuzz_Nosy_GameBuilderSeq_Defend__ because parameters include func, chan, or unsupported interface: []github.com/ethereum-optimism/optimism/op-challenger/game/fault/test.ClaimOpt

func Fuzz_Nosy_GameBuilderSeq_ExpectAttack__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *GameBuilderSeq
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.ExpectAttack()
	})
}

func Fuzz_Nosy_GameBuilderSeq_ExpectDefend__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *GameBuilderSeq
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.ExpectDefend()
	})
}

func Fuzz_Nosy_GameBuilderSeq_ExpectStepAttack__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *GameBuilderSeq
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.ExpectStepAttack()
	})
}

func Fuzz_Nosy_GameBuilderSeq_ExpectStepDefend__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *GameBuilderSeq
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.ExpectStepDefend()
	})
}

func Fuzz_Nosy_GameBuilderSeq_IsMaxDepth__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *GameBuilderSeq
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.IsMaxDepth()
	})
}

func Fuzz_Nosy_GameBuilderSeq_IsRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *GameBuilderSeq
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.IsRoot()
	})
}

// skipping Fuzz_Nosy_GameBuilderSeq_Step__ because parameters include func, chan, or unsupported interface: []github.com/ethereum-optimism/optimism/op-challenger/game/fault/test.ClaimOpt

func Fuzz_Nosy_GameBuilderSeq_addClaimToGame__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *GameBuilderSeq
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var claim *types.Claim
		fill_err = tp.Fill(&claim)
		if fill_err != nil {
			return
		}
		if s == nil || claim == nil {
			return
		}

		s.addClaimToGame(claim)
	})
}

// skipping Fuzz_Nosy_SequenceBuilder_Attack__ because parameters include func, chan, or unsupported interface: []github.com/ethereum-optimism/optimism/op-challenger/game/fault/test.ClaimOpt

// skipping Fuzz_Nosy_SequenceBuilder_Defend__ because parameters include func, chan, or unsupported interface: []github.com/ethereum-optimism/optimism/op-challenger/game/fault/test.ClaimOpt

func Fuzz_Nosy_SequenceBuilder_Get__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *SequenceBuilder
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Get()
	})
}

func Fuzz_Nosy_ClaimOpt_Apply__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var value common.Hash
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
		var cfg *claimCfg
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		_x1 := WithValue(value)
		_x1.Apply(cfg)
	})
}

// skipping Fuzz_Nosy_claimOptFn_Apply__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/test.claimOptFn
