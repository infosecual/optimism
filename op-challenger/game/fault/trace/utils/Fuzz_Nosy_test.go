package utils

import (
	"testing"

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

func Fuzz_Nosy_PreimageLoader_LoadPreimage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *PreimageLoader
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		var proof *ProofData
		fill_err = tp.Fill(&proof)
		if fill_err != nil {
			return
		}
		if l == nil || proof == nil {
			return
		}

		l.LoadPreimage(proof)
	})
}

func Fuzz_Nosy_PreimageLoader_loadBlobPreimage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *PreimageLoader
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		var proof *ProofData
		fill_err = tp.Fill(&proof)
		if fill_err != nil {
			return
		}
		if l == nil || proof == nil {
			return
		}

		l.loadBlobPreimage(proof)
	})
}

func Fuzz_Nosy_PreimageLoader_loadPrecompilePreimage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *PreimageLoader
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		var proof *ProofData
		fill_err = tp.Fill(&proof)
		if fill_err != nil {
			return
		}
		if l == nil || proof == nil {
			return
		}

		l.loadPrecompilePreimage(proof)
	})
}

// skipping Fuzz_Nosy_GameInputsSource_GetProposals__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/trace/utils.GameInputsSource

// skipping Fuzz_Nosy_L1HeadSource_GetL1Head__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/trace/utils.L1HeadSource

// skipping Fuzz_Nosy_L2HeaderSource_HeaderByNumber__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/trace/utils.L2HeaderSource

// skipping Fuzz_Nosy_PreimageSource_Close__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/trace/utils.PreimageSource

// skipping Fuzz_Nosy_PreimageSource_Get__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/trace/utils.PreimageSource

// skipping Fuzz_Nosy_ProofGenerator_GenerateProof__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/trace/utils.ProofGenerator

func Fuzz_Nosy_ReadLastStep__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dir string) {
		ReadLastStep(dir)
	})
}

func Fuzz_Nosy_lengthPrefixed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d1 []byte) {
		lengthPrefixed(d1)
	})
}
