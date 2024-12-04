package retry

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

func Fuzz_Nosy_ErrFailedPermanently_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *ErrFailedPermanently
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.Error()
	})
}

func Fuzz_Nosy_ErrFailedPermanently_Unwrap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *ErrFailedPermanently
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.Unwrap()
	})
}

func Fuzz_Nosy_ExponentialStrategy_Duration__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *ExponentialStrategy
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var attempt int
		fill_err = tp.Fill(&attempt)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.Duration(attempt)
	})
}

func Fuzz_Nosy_FixedStrategy_Duration__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FixedStrategy
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var attempt int
		fill_err = tp.Fill(&attempt)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.Duration(attempt)
	})
}

func Fuzz_Nosy_Strategy_Duration__(f *testing.F) {
	f.Fuzz(func(t *testing.T, attempt int) {
		_x1 := Exponential()
		_x1.Duration(attempt)
	})
}

// skipping Fuzz_Nosy_Do2__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/retry.Strategy

// skipping Fuzz_Nosy_Do__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/retry.Strategy
