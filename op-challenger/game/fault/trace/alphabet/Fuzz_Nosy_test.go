package alphabet

import (
	"context"
	"math/big"
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

func Fuzz_Nosy_AlphabetTraceProvider_Get__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var startingBlockNumber *big.Int
		fill_err = tp.Fill(&startingBlockNumber)
		if fill_err != nil {
			return
		}
		var depth types.Depth
		fill_err = tp.Fill(&depth)
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
		if startingBlockNumber == nil {
			return
		}

		ap := NewTraceProvider(startingBlockNumber, depth)
		ap.Get(ctx, i)
	})
}

func Fuzz_Nosy_AlphabetTraceProvider_GetL2BlockNumberChallenge__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var startingBlockNumber *big.Int
		fill_err = tp.Fill(&startingBlockNumber)
		if fill_err != nil {
			return
		}
		var depth types.Depth
		fill_err = tp.Fill(&depth)
		if fill_err != nil {
			return
		}
		var _x3 context.Context
		fill_err = tp.Fill(&_x3)
		if fill_err != nil {
			return
		}
		if startingBlockNumber == nil {
			return
		}

		ap := NewTraceProvider(startingBlockNumber, depth)
		ap.GetL2BlockNumberChallenge(_x3)
	})
}

func Fuzz_Nosy_AlphabetTraceProvider_GetStepData__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var startingBlockNumber *big.Int
		fill_err = tp.Fill(&startingBlockNumber)
		if fill_err != nil {
			return
		}
		var depth types.Depth
		fill_err = tp.Fill(&depth)
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
		if startingBlockNumber == nil {
			return
		}

		ap := NewTraceProvider(startingBlockNumber, depth)
		ap.GetStepData(ctx, pos)
	})
}

func Fuzz_Nosy_alphabetPrestateProvider_AbsolutePreStateCommitment__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ap *alphabetPrestateProvider
		fill_err = tp.Fill(&ap)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if ap == nil {
			return
		}

		ap.AbsolutePreStateCommitment(_x2)
	})
}

func Fuzz_Nosy_BuildAlphabetPreimage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var traceIndex *big.Int
		fill_err = tp.Fill(&traceIndex)
		if fill_err != nil {
			return
		}
		var claim *big.Int
		fill_err = tp.Fill(&claim)
		if fill_err != nil {
			return
		}
		if traceIndex == nil || claim == nil {
			return
		}

		BuildAlphabetPreimage(traceIndex, claim)
	})
}
