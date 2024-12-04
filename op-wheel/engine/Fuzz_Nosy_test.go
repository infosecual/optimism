package engine

import (
	"testing"

	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum/go-ethereum/common"
	"github.com/prometheus/client_golang/prometheus"
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

func Fuzz_Nosy_Metrics_RecordBlockFail__(f *testing.F) {
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
		var registry *prometheus.Registry
		fill_err = tp.Fill(&registry)
		if fill_err != nil {
			return
		}
		if registry == nil {
			return
		}

		r := NewMetrics(procName, registry)
		r.RecordBlockFail()
	})
}

func Fuzz_Nosy_Metrics_RecordBlockStats__(f *testing.F) {
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
		var registry *prometheus.Registry
		fill_err = tp.Fill(&registry)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		var num uint64
		fill_err = tp.Fill(&num)
		if fill_err != nil {
			return
		}
		var time uint64
		fill_err = tp.Fill(&time)
		if fill_err != nil {
			return
		}
		var txs uint64
		fill_err = tp.Fill(&txs)
		if fill_err != nil {
			return
		}
		var gas uint64
		fill_err = tp.Fill(&gas)
		if fill_err != nil {
			return
		}
		var baseFee float64
		fill_err = tp.Fill(&baseFee)
		if fill_err != nil {
			return
		}
		if registry == nil {
			return
		}

		r := NewMetrics(procName, registry)
		r.RecordBlockStats(hash, num, time, txs, gas, baseFee)
	})
}

// skipping Fuzz_Nosy_Metricer_RecordBlockFail__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-wheel/engine.Metricer

// skipping Fuzz_Nosy_Metricer_RecordBlockStats__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-wheel/engine.Metricer

func Fuzz_Nosy_StaticVersionProvider_ForkchoiceUpdatedVersion__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v StaticVersionProvider
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		var _x2 *eth.PayloadAttributes
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if _x2 == nil {
			return
		}

		v.ForkchoiceUpdatedVersion(_x2)
	})
}

func Fuzz_Nosy_StaticVersionProvider_GetPayloadVersion__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v StaticVersionProvider
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		var _x2 uint64
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}

		v.GetPayloadVersion(_x2)
	})
}

func Fuzz_Nosy_StaticVersionProvider_NewPayloadVersion__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v StaticVersionProvider
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		var _x2 uint64
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}

		v.NewPayloadVersion(_x2)
	})
}

// skipping Fuzz_Nosy_headBlockSafeFinalized__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/client.RPC

// skipping Fuzz_Nosy_headSafeFinalized__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/client.RPC

func Fuzz_Nosy_isUnquotedJsonString__(f *testing.F) {
	f.Fuzz(func(t *testing.T, v string) {
		isUnquotedJsonString(v)
	})
}

// skipping Fuzz_Nosy_safeFinalized__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/client.RPC
