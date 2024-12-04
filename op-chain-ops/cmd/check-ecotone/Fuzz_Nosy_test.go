package main

import (
	"testing"

	"github.com/ethereum/go-ethereum/core/types"
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

func Fuzz_Nosy_actionEnv_RecordGasUsed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ae *actionEnv
		fill_err = tp.Fill(&ae)
		if fill_err != nil {
			return
		}
		var rec *types.Receipt
		fill_err = tp.Fill(&rec)
		if fill_err != nil {
			return
		}
		if ae == nil || rec == nil {
			return
		}

		ae.RecordGasUsed(rec)
	})
}

func Fuzz_Nosy_conditionalCode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d1 []byte) {
		conditionalCode(d1)
	})
}

// skipping Fuzz_Nosy_makeCommandAction__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-chain-ops/cmd/check-ecotone.CheckAction
