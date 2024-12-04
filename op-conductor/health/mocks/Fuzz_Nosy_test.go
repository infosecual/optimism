package mocks

import (
	context "context"
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

func Fuzz_Nosy_HealthMonitor_EXPECT__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *HealthMonitor
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.EXPECT()
	})
}

func Fuzz_Nosy_HealthMonitor_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *HealthMonitor
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.Start(ctx)
	})
}

func Fuzz_Nosy_HealthMonitor_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *HealthMonitor
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.Stop()
	})
}

func Fuzz_Nosy_HealthMonitor_Subscribe__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *HealthMonitor
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.Subscribe()
	})
}

// skipping Fuzz_Nosy_HealthMonitor_Expecter_Start__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_HealthMonitor_Expecter_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _e *HealthMonitor_Expecter
		fill_err = tp.Fill(&_e)
		if fill_err != nil {
			return
		}
		if _e == nil {
			return
		}

		_e.Stop()
	})
}

func Fuzz_Nosy_HealthMonitor_Expecter_Subscribe__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _e *HealthMonitor_Expecter
		fill_err = tp.Fill(&_e)
		if fill_err != nil {
			return
		}
		if _e == nil {
			return
		}

		_e.Subscribe()
	})
}

// skipping Fuzz_Nosy_HealthMonitor_Start_Call_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_HealthMonitor_Start_Call_Run__ because parameters include func, chan, or unsupported interface: func(ctx context.Context)

// skipping Fuzz_Nosy_HealthMonitor_Start_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(context.Context) error

// skipping Fuzz_Nosy_HealthMonitor_Stop_Call_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_HealthMonitor_Stop_Call_Run__ because parameters include func, chan, or unsupported interface: func()

// skipping Fuzz_Nosy_HealthMonitor_Stop_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func() error

// skipping Fuzz_Nosy_HealthMonitor_Subscribe_Call_Return__ because parameters include func, chan, or unsupported interface: <-chan error

// skipping Fuzz_Nosy_HealthMonitor_Subscribe_Call_Run__ because parameters include func, chan, or unsupported interface: func()

// skipping Fuzz_Nosy_HealthMonitor_Subscribe_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func() <-chan error
