package rpc

import (
	"context"
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

func Fuzz_Nosy_adminAPI_StartBatcher__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a *adminAPI
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if a == nil {
			return
		}

		a.StartBatcher(_x2)
	})
}

func Fuzz_Nosy_adminAPI_StopBatcher__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a *adminAPI
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

		a.StopBatcher(ctx)
	})
}

// skipping Fuzz_Nosy_BatcherDriver_StartBatchSubmitting__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-batcher/rpc.BatcherDriver

// skipping Fuzz_Nosy_BatcherDriver_StopBatchSubmitting__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-batcher/rpc.BatcherDriver
