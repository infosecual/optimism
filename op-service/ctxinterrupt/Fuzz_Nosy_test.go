package ctxinterrupt

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

// skipping Fuzz_Nosy_WaiterFunc_waitForInterrupt__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/ctxinterrupt.WaiterFunc

func Fuzz_Nosy_signalWaiter_waitForInterrupt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		me := newSignalWaiter()
		me.waitForInterrupt(ctx)
	})
}

func Fuzz_Nosy_waitResult_Cause__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var me waitResult
		fill_err = tp.Fill(&me)
		if fill_err != nil {
			return
		}

		me.Cause()
	})
}

func Fuzz_Nosy_waiter_waitForInterrupt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var c2 context.Context
		fill_err = tp.Fill(&c2)
		if fill_err != nil {
			return
		}

		_x1 := contextInterruptWaiter(c1)
		_x1.waitForInterrupt(c2)
	})
}

func Fuzz_Nosy_WithSignalWaiter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		WithSignalWaiter(ctx)
	})
}
