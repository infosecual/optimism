package test

import (
	"testing"

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

func Fuzz_Nosy_KvStateOracle_CodeByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o *KvStateOracle
		fill_err = tp.Fill(&o)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if o == nil {
			return
		}

		o.CodeByHash(hash)
	})
}

func Fuzz_Nosy_KvStateOracle_NodeByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o *KvStateOracle
		fill_err = tp.Fill(&o)
		if fill_err != nil {
			return
		}
		var nodeHash common.Hash
		fill_err = tp.Fill(&nodeHash)
		if fill_err != nil {
			return
		}
		if o == nil {
			return
		}

		o.NodeByHash(nodeHash)
	})
}

func Fuzz_Nosy_StubPrecompileOracle_Precompile__(f *testing.F) {
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
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		var requiredGas uint64
		fill_err = tp.Fill(&requiredGas)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		o := NewStubPrecompileOracle(t1)
		o.Precompile(address, input, requiredGas)
	})
}

func Fuzz_Nosy_StubStateOracle_CodeByHash__(f *testing.F) {
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
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		o := NewStubStateOracle(t1)
		o.CodeByHash(hash)
	})
}

func Fuzz_Nosy_StubStateOracle_NodeByHash__(f *testing.F) {
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
		var nodeHash common.Hash
		fill_err = tp.Fill(&nodeHash)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		o := NewStubStateOracle(t1)
		o.NodeByHash(nodeHash)
	})
}

func Fuzz_Nosy_StubBlockOracle_BlockByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o StubBlockOracle
		fill_err = tp.Fill(&o)
		if fill_err != nil {
			return
		}
		var blockHash common.Hash
		fill_err = tp.Fill(&blockHash)
		if fill_err != nil {
			return
		}

		o.BlockByHash(blockHash)
	})
}

func Fuzz_Nosy_StubBlockOracle_OutputByRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o StubBlockOracle
		fill_err = tp.Fill(&o)
		if fill_err != nil {
			return
		}
		var root common.Hash
		fill_err = tp.Fill(&root)
		if fill_err != nil {
			return
		}

		o.OutputByRoot(root)
	})
}

// skipping Fuzz_Nosy_stateOracle_CodeByHash__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-program/client/l2/test.stateOracle

// skipping Fuzz_Nosy_stateOracle_NodeByHash__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-program/client/l2/test.stateOracle

func Fuzz_Nosy_NewStubOracle__(f *testing.F) {
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
		if t1 == nil {
			return
		}

		NewStubOracle(t1)
	})
}
