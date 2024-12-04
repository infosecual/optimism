package metrics

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

func Fuzz_Nosy_Metrics_RecordActedL1Block__(f *testing.F) {
	f.Fuzz(func(t *testing.T, n uint64) {
		m := NewMetrics()
		m.RecordActedL1Block(n)
	})
}

func Fuzz_Nosy_Metrics_RecordBondClaimed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, amount uint64) {
		m := NewMetrics()
		m.RecordBondClaimed(amount)
	})
}

func Fuzz_Nosy_Metrics_RecordClaimResolutionTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, t1 float64) {
		m := NewMetrics()
		m.RecordClaimResolutionTime(t1)
	})
}

func Fuzz_Nosy_Metrics_RecordGameActTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, t1 float64) {
		m := NewMetrics()
		m.RecordGameActTime(t1)
	})
}

func Fuzz_Nosy_Metrics_RecordGamesStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, inProgress int, defenderWon int, challengerWon int) {
		m := NewMetrics()
		m.RecordGamesStatus(inProgress, defenderWon, challengerWon)
	})
}

func Fuzz_Nosy_Metrics_RecordInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, version string) {
		m := NewMetrics()
		m.RecordInfo(version)
	})
}

func Fuzz_Nosy_Metrics_RecordLargePreimageCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, count int) {
		m := NewMetrics()
		m.RecordLargePreimageCount(count)
	})
}

func Fuzz_Nosy_Metrics_RecordVmExecutionTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var vmType string
		fill_err = tp.Fill(&vmType)
		if fill_err != nil {
			return
		}
		var dur time.Duration
		fill_err = tp.Fill(&dur)
		if fill_err != nil {
			return
		}

		m := NewMetrics()
		m.RecordVmExecutionTime(vmType, dur)
	})
}

func Fuzz_Nosy_Metrics_RecordVmMemoryUsed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, vmType string, memoryUsed uint64) {
		m := NewMetrics()
		m.RecordVmMemoryUsed(vmType, memoryUsed)
	})
}

func Fuzz_Nosy_Metrics_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, host string, port int) {
		m := NewMetrics()
		m.Start(host, port)
	})
}

// skipping Fuzz_Nosy_Metrics_StartBalanceMetrics__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/log.Logger

func Fuzz_Nosy_Metrics_VmMetrics__(f *testing.F) {
	f.Fuzz(func(t *testing.T, vmType string) {
		m := NewMetrics()
		m.VmMetrics(vmType)
	})
}

func Fuzz_Nosy_NoopMetricsImpl_CacheAdd__(f *testing.F) {
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
		var _x2 string
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var _x3 int
		fill_err = tp.Fill(&_x3)
		if fill_err != nil {
			return
		}
		var _x4 bool
		fill_err = tp.Fill(&_x4)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.CacheAdd(_x2, _x3, _x4)
	})
}

func Fuzz_Nosy_NoopMetricsImpl_CacheGet__(f *testing.F) {
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
		var _x2 string
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var _x3 bool
		fill_err = tp.Fill(&_x3)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.CacheGet(_x2, _x3)
	})
}

func Fuzz_Nosy_NoopMetricsImpl_DecActiveExecutors__(f *testing.F) {
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

		_x1.DecActiveExecutors()
	})
}

func Fuzz_Nosy_NoopMetricsImpl_DecIdleExecutors__(f *testing.F) {
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

		_x1.DecIdleExecutors()
	})
}

func Fuzz_Nosy_NoopMetricsImpl_IncActiveExecutors__(f *testing.F) {
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

		_x1.IncActiveExecutors()
	})
}

func Fuzz_Nosy_NoopMetricsImpl_IncIdleExecutors__(f *testing.F) {
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

		_x1.IncIdleExecutors()
	})
}

func Fuzz_Nosy_NoopMetricsImpl_RecordActedL1Block__(f *testing.F) {
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
		var _x2 uint64
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.RecordActedL1Block(_x2)
	})
}

func Fuzz_Nosy_NoopMetricsImpl_RecordBondClaimFailed__(f *testing.F) {
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

		_x1.RecordBondClaimFailed()
	})
}

func Fuzz_Nosy_NoopMetricsImpl_RecordBondClaimed__(f *testing.F) {
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
		var _x2 uint64
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.RecordBondClaimed(_x2)
	})
}

func Fuzz_Nosy_NoopMetricsImpl_RecordClaimResolutionTime__(f *testing.F) {
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
		var t2 float64
		fill_err = tp.Fill(&t2)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.RecordClaimResolutionTime(t2)
	})
}

func Fuzz_Nosy_NoopMetricsImpl_RecordGameActTime__(f *testing.F) {
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
		var t2 float64
		fill_err = tp.Fill(&t2)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.RecordGameActTime(t2)
	})
}

func Fuzz_Nosy_NoopMetricsImpl_RecordGameL2Challenge__(f *testing.F) {
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

		_x1.RecordGameL2Challenge()
	})
}

func Fuzz_Nosy_NoopMetricsImpl_RecordGameMove__(f *testing.F) {
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

		_x1.RecordGameMove()
	})
}

func Fuzz_Nosy_NoopMetricsImpl_RecordGameStep__(f *testing.F) {
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

		_x1.RecordGameStep()
	})
}

func Fuzz_Nosy_NoopMetricsImpl_RecordGameUpdateCompleted__(f *testing.F) {
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

		_x1.RecordGameUpdateCompleted()
	})
}

func Fuzz_Nosy_NoopMetricsImpl_RecordGameUpdateScheduled__(f *testing.F) {
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

		_x1.RecordGameUpdateScheduled()
	})
}

func Fuzz_Nosy_NoopMetricsImpl_RecordGamesStatus__(f *testing.F) {
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
		var inProgress int
		fill_err = tp.Fill(&inProgress)
		if fill_err != nil {
			return
		}
		var defenderWon int
		fill_err = tp.Fill(&defenderWon)
		if fill_err != nil {
			return
		}
		var challengerWon int
		fill_err = tp.Fill(&challengerWon)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.RecordGamesStatus(inProgress, defenderWon, challengerWon)
	})
}

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

func Fuzz_Nosy_NoopMetricsImpl_RecordLargePreimageCount__(f *testing.F) {
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
		var _x2 int
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.RecordLargePreimageCount(_x2)
	})
}

func Fuzz_Nosy_NoopMetricsImpl_RecordPreimageChallengeFailed__(f *testing.F) {
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

		_x1.RecordPreimageChallengeFailed()
	})
}

func Fuzz_Nosy_NoopMetricsImpl_RecordPreimageChallenged__(f *testing.F) {
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

		_x1.RecordPreimageChallenged()
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

func Fuzz_Nosy_NoopMetricsImpl_RecordVmExecutionTime__(f *testing.F) {
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
		var _x2 string
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var _x3 time.Duration
		fill_err = tp.Fill(&_x3)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.RecordVmExecutionTime(_x2, _x3)
	})
}

func Fuzz_Nosy_NoopMetricsImpl_RecordVmMemoryUsed__(f *testing.F) {
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
		var _x2 string
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var _x3 uint64
		fill_err = tp.Fill(&_x3)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.RecordVmMemoryUsed(_x2, _x3)
	})
}

// skipping Fuzz_Nosy_NoopMetricsImpl_StartBalanceMetrics__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/log.Logger

func Fuzz_Nosy_NoopMetricsImpl_VmMetrics__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *NoopMetricsImpl
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var vmType string
		fill_err = tp.Fill(&vmType)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.VmMetrics(vmType)
	})
}

func Fuzz_Nosy_VmMetrics_RecordExecutionTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *VmMetrics
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dur time.Duration
		fill_err = tp.Fill(&dur)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.RecordExecutionTime(dur)
	})
}

func Fuzz_Nosy_VmMetrics_RecordMemoryUsed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *VmMetrics
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var memoryUsed uint64
		fill_err = tp.Fill(&memoryUsed)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.RecordMemoryUsed(memoryUsed)
	})
}

// skipping Fuzz_Nosy_Metricer_DecActiveExecutors__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_DecIdleExecutors__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_IncActiveExecutors__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_IncIdleExecutors__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordActedL1Block__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordBondClaimFailed__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordBondClaimed__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordClaimResolutionTime__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordGameActTime__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordGameL2Challenge__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordGameMove__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordGameStep__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordGameUpdateCompleted__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordGameUpdateScheduled__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordGamesStatus__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordInfo__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordLargePreimageCount__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordPreimageChallengeFailed__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordPreimageChallenged__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordUp__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_StartBalanceMetrics__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_VmMetrics__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/metrics.Metricer

// skipping Fuzz_Nosy_VmMetricer_RecordVmExecutionTime__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/metrics.VmMetricer

// skipping Fuzz_Nosy_VmMetricer_RecordVmMemoryUsed__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/metrics.VmMetricer
