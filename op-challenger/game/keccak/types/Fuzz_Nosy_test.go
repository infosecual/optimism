package types

import (
	"testing"
	"time"

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

func Fuzz_Nosy_LargePreimageMetaData_ShouldVerify__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m LargePreimageMetaData
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var now time.Time
		fill_err = tp.Fill(&now)
		if fill_err != nil {
			return
		}
		var ignoreAfter time.Duration
		fill_err = tp.Fill(&ignoreAfter)
		if fill_err != nil {
			return
		}

		m.ShouldVerify(now, ignoreAfter)
	})
}

// skipping Fuzz_Nosy_LargePreimageOracle_Addr__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/keccak/types.LargePreimageOracle

// skipping Fuzz_Nosy_LargePreimageOracle_ChallengePeriod__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/keccak/types.LargePreimageOracle

// skipping Fuzz_Nosy_LargePreimageOracle_ChallengeTx__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/keccak/types.LargePreimageOracle

// skipping Fuzz_Nosy_LargePreimageOracle_DecodeInputData__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/keccak/types.LargePreimageOracle

// skipping Fuzz_Nosy_LargePreimageOracle_GetActivePreimages__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/keccak/types.LargePreimageOracle

// skipping Fuzz_Nosy_LargePreimageOracle_GetInputDataBlocks__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/keccak/types.LargePreimageOracle

// skipping Fuzz_Nosy_LargePreimageOracle_GetProposalTreeRoot__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/keccak/types.LargePreimageOracle

func Fuzz_Nosy_Leaf_Hash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l Leaf
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}

		l.Hash()
	})
}

func Fuzz_Nosy_StateSnapshot_Pack__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s StateSnapshot
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}

		s.Pack()
	})
}
