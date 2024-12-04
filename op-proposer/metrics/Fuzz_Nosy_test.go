package metrics

import (
	"testing"

	"github.com/ethereum-optimism/optimism/op-service/eth"
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

func Fuzz_Nosy_Metrics_Document__(f *testing.F) {
	f.Fuzz(func(t *testing.T, procName string) {
		m := NewMetrics(procName)
		m.Document()
	})
}

func Fuzz_Nosy_Metrics_RecordInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, procName string, version string) {
		m := NewMetrics(procName)
		m.RecordInfo(version)
	})
}

func Fuzz_Nosy_Metrics_RecordL2BlocksProposed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var procName string
		fill_err = tp.Fill(&procName)
		if fill_err != nil {
			return
		}
		var l2ref eth.L2BlockRef
		fill_err = tp.Fill(&l2ref)
		if fill_err != nil {
			return
		}

		m := NewMetrics(procName)
		m.RecordL2BlocksProposed(l2ref)
	})
}

func Fuzz_Nosy_Metrics_RecordUp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, procName string) {
		m := NewMetrics(procName)
		m.RecordUp()
	})
}

func Fuzz_Nosy_Metrics_Registry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, procName string) {
		m := NewMetrics(procName)
		m.Registry()
	})
}

// skipping Fuzz_Nosy_Metrics_StartBalanceMetrics__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/log.Logger

func Fuzz_Nosy_noopMetrics_RecordInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *noopMetrics
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

func Fuzz_Nosy_noopMetrics_RecordL2BlocksProposed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *noopMetrics
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var l2ref eth.L2BlockRef
		fill_err = tp.Fill(&l2ref)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.RecordL2BlocksProposed(l2ref)
	})
}

func Fuzz_Nosy_noopMetrics_RecordUp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *noopMetrics
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

// skipping Fuzz_Nosy_noopMetrics_StartBalanceMetrics__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/log.Logger

// skipping Fuzz_Nosy_Metricer_RecordInfo__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-proposer/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordL2BlocksProposed__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-proposer/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordUp__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-proposer/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_StartBalanceMetrics__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-proposer/metrics.Metricer
