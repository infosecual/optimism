package pipeline

import (
	"testing"

	"github.com/ethereum-optimism/optimism/op-deployer/pkg/deployer/state"
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

func Fuzz_Nosy_StateWriter_WriteState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var st *state.State
		fill_err = tp.Fill(&st)
		if fill_err != nil {
			return
		}
		if st == nil {
			return
		}

		_x1 := NoopStateWriter()
		_x1.WriteState(st)
	})
}

// skipping Fuzz_Nosy_stateWriterFunc_WriteState__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-deployer/pkg/deployer/pipeline.stateWriterFunc

func Fuzz_Nosy_IsSupportedStateVersion__(f *testing.F) {
	f.Fuzz(func(t *testing.T, version int) {
		IsSupportedStateVersion(version)
	})
}

func Fuzz_Nosy_shouldDeployAltDA__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainIntent *state.ChainIntent
		fill_err = tp.Fill(&chainIntent)
		if fill_err != nil {
			return
		}
		var chainState *state.ChainState
		fill_err = tp.Fill(&chainState)
		if fill_err != nil {
			return
		}
		if chainIntent == nil || chainState == nil {
			return
		}

		shouldDeployAltDA(chainIntent, chainState)
	})
}

func Fuzz_Nosy_shouldDeployImplementations__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var intent *state.Intent
		fill_err = tp.Fill(&intent)
		if fill_err != nil {
			return
		}
		var st *state.State
		fill_err = tp.Fill(&st)
		if fill_err != nil {
			return
		}
		if intent == nil || st == nil {
			return
		}

		shouldDeployImplementations(intent, st)
	})
}

func Fuzz_Nosy_shouldDeployOPChain__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var st *state.State
		fill_err = tp.Fill(&st)
		if fill_err != nil {
			return
		}
		var chainID common.Hash
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		if st == nil {
			return
		}

		shouldDeployOPChain(st, chainID)
	})
}

func Fuzz_Nosy_shouldDeploySuperchain__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var intent *state.Intent
		fill_err = tp.Fill(&intent)
		if fill_err != nil {
			return
		}
		var st *state.State
		fill_err = tp.Fill(&st)
		if fill_err != nil {
			return
		}
		if intent == nil || st == nil {
			return
		}

		shouldDeploySuperchain(intent, st)
	})
}

func Fuzz_Nosy_shouldGenerateL2Genesis__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var thisChainState *state.ChainState
		fill_err = tp.Fill(&thisChainState)
		if fill_err != nil {
			return
		}
		if thisChainState == nil {
			return
		}

		shouldGenerateL2Genesis(thisChainState)
	})
}
