package metrics

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

// skipping Fuzz_Nosy_Metrics_RecordHealthCheck__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_Metrics_RecordInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, version string) {
		m := NewMetrics()
		m.RecordInfo(version)
	})
}

func Fuzz_Nosy_Metrics_RecordLeaderTransfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, success bool) {
		m := NewMetrics()
		m.RecordLeaderTransfer(success)
	})
}

func Fuzz_Nosy_Metrics_RecordLoopExecutionTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, duration float64) {
		m := NewMetrics()
		m.RecordLoopExecutionTime(duration)
	})
}

func Fuzz_Nosy_Metrics_RecordStartSequencer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, success bool) {
		m := NewMetrics()
		m.RecordStartSequencer(success)
	})
}

func Fuzz_Nosy_Metrics_RecordStateChange__(f *testing.F) {
	f.Fuzz(func(t *testing.T, leader bool, healthy bool, active bool) {
		m := NewMetrics()
		m.RecordStateChange(leader, healthy, active)
	})
}

func Fuzz_Nosy_Metrics_RecordStopSequencer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, success bool) {
		m := NewMetrics()
		m.RecordStopSequencer(success)
	})
}

func Fuzz_Nosy_Metrics_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, host string, port int) {
		m := NewMetrics()
		m.Start(host, port)
	})
}

// skipping Fuzz_Nosy_NoopMetricsImpl_RecordHealthCheck__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_NoopMetricsImpl_RecordInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *NoopMetricsImpl
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var version string
		fill_err = tp.Fill(&version)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.RecordInfo(version)
	})
}

func Fuzz_Nosy_NoopMetricsImpl_RecordLeaderTransfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *NoopMetricsImpl
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var success bool
		fill_err = tp.Fill(&success)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.RecordLeaderTransfer(success)
	})
}

func Fuzz_Nosy_NoopMetricsImpl_RecordLoopExecutionTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *NoopMetricsImpl
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var duration float64
		fill_err = tp.Fill(&duration)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.RecordLoopExecutionTime(duration)
	})
}

func Fuzz_Nosy_NoopMetricsImpl_RecordStartSequencer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *NoopMetricsImpl
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var success bool
		fill_err = tp.Fill(&success)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.RecordStartSequencer(success)
	})
}

func Fuzz_Nosy_NoopMetricsImpl_RecordStateChange__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *NoopMetricsImpl
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var leader bool
		fill_err = tp.Fill(&leader)
		if fill_err != nil {
			return
		}
		var healthy bool
		fill_err = tp.Fill(&healthy)
		if fill_err != nil {
			return
		}
		var active bool
		fill_err = tp.Fill(&active)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.RecordStateChange(leader, healthy, active)
	})
}

func Fuzz_Nosy_NoopMetricsImpl_RecordStopSequencer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *NoopMetricsImpl
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var success bool
		fill_err = tp.Fill(&success)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.RecordStopSequencer(success)
	})
}

func Fuzz_Nosy_NoopMetricsImpl_RecordUp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *NoopMetricsImpl
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.RecordUp()
	})
}

// skipping Fuzz_Nosy_Metricer_RecordHealthCheck__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-conductor/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordInfo__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-conductor/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordLeaderTransfer__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-conductor/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordLoopExecutionTime__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-conductor/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordStartSequencer__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-conductor/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordStateChange__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-conductor/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordStopSequencer__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-conductor/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordUp__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-conductor/metrics.Metricer
