package metrics

import (
	"testing"

	"github.com/ethereum-optimism/optimism/op-supervisor/supervisor/types"
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

func Fuzz_Nosy_Metrics_CacheAdd__(f *testing.F) {
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
		var chainID types.ChainID
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var label string
		fill_err = tp.Fill(&label)
		if fill_err != nil {
			return
		}
		var cacheSize int
		fill_err = tp.Fill(&cacheSize)
		if fill_err != nil {
			return
		}
		var evicted bool
		fill_err = tp.Fill(&evicted)
		if fill_err != nil {
			return
		}

		m := NewMetrics(procName)
		m.CacheAdd(chainID, label, cacheSize, evicted)
	})
}

func Fuzz_Nosy_Metrics_CacheGet__(f *testing.F) {
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
		var chainID types.ChainID
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var label string
		fill_err = tp.Fill(&label)
		if fill_err != nil {
			return
		}
		var hit bool
		fill_err = tp.Fill(&hit)
		if fill_err != nil {
			return
		}

		m := NewMetrics(procName)
		m.CacheGet(chainID, label, hit)
	})
}

func Fuzz_Nosy_Metrics_Document__(f *testing.F) {
	f.Fuzz(func(t *testing.T, procName string) {
		m := NewMetrics(procName)
		m.Document()
	})
}

func Fuzz_Nosy_Metrics_RecordDBEntryCount__(f *testing.F) {
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
		var chainID types.ChainID
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var kind string
		fill_err = tp.Fill(&kind)
		if fill_err != nil {
			return
		}
		var count int64
		fill_err = tp.Fill(&count)
		if fill_err != nil {
			return
		}

		m := NewMetrics(procName)
		m.RecordDBEntryCount(chainID, kind, count)
	})
}

func Fuzz_Nosy_Metrics_RecordDBSearchEntriesRead__(f *testing.F) {
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
		var chainID types.ChainID
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var count int64
		fill_err = tp.Fill(&count)
		if fill_err != nil {
			return
		}

		m := NewMetrics(procName)
		m.RecordDBSearchEntriesRead(chainID, count)
	})
}

func Fuzz_Nosy_Metrics_RecordInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, procName string, version string) {
		m := NewMetrics(procName)
		m.RecordInfo(version)
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

func Fuzz_Nosy_noopMetrics_CacheAdd__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *noopMetrics
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var _x2 types.ChainID
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var _x3 string
		fill_err = tp.Fill(&_x3)
		if fill_err != nil {
			return
		}
		var _x4 int
		fill_err = tp.Fill(&_x4)
		if fill_err != nil {
			return
		}
		var _x5 bool
		fill_err = tp.Fill(&_x5)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.CacheAdd(_x2, _x3, _x4, _x5)
	})
}

func Fuzz_Nosy_noopMetrics_CacheGet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *noopMetrics
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var _x2 types.ChainID
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var _x3 string
		fill_err = tp.Fill(&_x3)
		if fill_err != nil {
			return
		}
		var _x4 bool
		fill_err = tp.Fill(&_x4)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.CacheGet(_x2, _x3, _x4)
	})
}

func Fuzz_Nosy_noopMetrics_Document__(f *testing.F) {
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

		_x1.Document()
	})
}

func Fuzz_Nosy_noopMetrics_RecordDBEntryCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *noopMetrics
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var _x2 types.ChainID
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var _x3 string
		fill_err = tp.Fill(&_x3)
		if fill_err != nil {
			return
		}
		var _x4 int64
		fill_err = tp.Fill(&_x4)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.RecordDBEntryCount(_x2, _x3, _x4)
	})
}

func Fuzz_Nosy_noopMetrics_RecordDBSearchEntriesRead__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *noopMetrics
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var _x2 types.ChainID
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var _x3 int64
		fill_err = tp.Fill(&_x3)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.RecordDBSearchEntriesRead(_x2, _x3)
	})
}

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

// skipping Fuzz_Nosy_Metricer_CacheAdd__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_CacheGet__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_Document__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordDBEntryCount__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordDBSearchEntriesRead__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordInfo__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordUp__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/metrics.Metricer

func Fuzz_Nosy_chainIDLabel__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainID types.ChainID
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}

		chainIDLabel(chainID)
	})
}
