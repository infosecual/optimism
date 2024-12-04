package metrics

import (
	"testing"

	"github.com/ethereum-optimism/optimism/op-node/rollup/derive"
	"github.com/ethereum-optimism/optimism/op-service/eth"
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

func Fuzz_Nosy_Metrics_Document__(f *testing.F) {
	f.Fuzz(func(t *testing.T, procName string) {
		m := NewMetrics(procName)
		m.Document()
	})
}

func Fuzz_Nosy_Metrics_PendingDABytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, procName string) {
		m := NewMetrics(procName)
		m.PendingDABytes()
	})
}

func Fuzz_Nosy_Metrics_RecordBatchTxFailed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, procName string) {
		m := NewMetrics(procName)
		m.RecordBatchTxFailed()
	})
}

func Fuzz_Nosy_Metrics_RecordBatchTxSubmitted__(f *testing.F) {
	f.Fuzz(func(t *testing.T, procName string) {
		m := NewMetrics(procName)
		m.RecordBatchTxSubmitted()
	})
}

func Fuzz_Nosy_Metrics_RecordBatchTxSuccess__(f *testing.F) {
	f.Fuzz(func(t *testing.T, procName string) {
		m := NewMetrics(procName)
		m.RecordBatchTxSuccess()
	})
}

func Fuzz_Nosy_Metrics_RecordBlobUsedBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, procName string, num int) {
		m := NewMetrics(procName)
		m.RecordBlobUsedBytes(num)
	})
}

// skipping Fuzz_Nosy_Metrics_RecordChannelClosed__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_Metrics_RecordChannelFullySubmitted__(f *testing.F) {
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
		var id derive.ChannelID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}

		m := NewMetrics(procName)
		m.RecordChannelFullySubmitted(id)
	})
}

func Fuzz_Nosy_Metrics_RecordChannelOpened__(f *testing.F) {
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
		var id derive.ChannelID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var numPendingBlocks int
		fill_err = tp.Fill(&numPendingBlocks)
		if fill_err != nil {
			return
		}

		m := NewMetrics(procName)
		m.RecordChannelOpened(id, numPendingBlocks)
	})
}

func Fuzz_Nosy_Metrics_RecordChannelTimedOut__(f *testing.F) {
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
		var id derive.ChannelID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}

		m := NewMetrics(procName)
		m.RecordChannelTimedOut(id)
	})
}

func Fuzz_Nosy_Metrics_RecordInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, procName string, version string) {
		m := NewMetrics(procName)
		m.RecordInfo(version)
	})
}

func Fuzz_Nosy_Metrics_RecordL2BlockInChannel__(f *testing.F) {
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
		var block *types.Block
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		if block == nil {
			return
		}

		m := NewMetrics(procName)
		m.RecordL2BlockInChannel(block)
	})
}

func Fuzz_Nosy_Metrics_RecordL2BlockInPendingQueue__(f *testing.F) {
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
		var block *types.Block
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		if block == nil {
			return
		}

		m := NewMetrics(procName)
		m.RecordL2BlockInPendingQueue(block)
	})
}

func Fuzz_Nosy_Metrics_RecordL2BlocksAdded__(f *testing.F) {
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
		var numBlocksAdded int
		fill_err = tp.Fill(&numBlocksAdded)
		if fill_err != nil {
			return
		}
		var numPendingBlocks int
		fill_err = tp.Fill(&numPendingBlocks)
		if fill_err != nil {
			return
		}
		var inputBytes int
		fill_err = tp.Fill(&inputBytes)
		if fill_err != nil {
			return
		}
		var outputComprBytes int
		fill_err = tp.Fill(&outputComprBytes)
		if fill_err != nil {
			return
		}

		m := NewMetrics(procName)
		m.RecordL2BlocksAdded(l2ref, numBlocksAdded, numPendingBlocks, inputBytes, outputComprBytes)
	})
}

func Fuzz_Nosy_Metrics_RecordL2BlocksLoaded__(f *testing.F) {
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
		m.RecordL2BlocksLoaded(l2ref)
	})
}

func Fuzz_Nosy_Metrics_RecordLatestL1Block__(f *testing.F) {
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
		var l1ref eth.L1BlockRef
		fill_err = tp.Fill(&l1ref)
		if fill_err != nil {
			return
		}

		m := NewMetrics(procName)
		m.RecordLatestL1Block(l1ref)
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

func Fuzz_Nosy_TestMetrics_RecordL2BlockInChannel__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *TestMetrics
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var block *types.Block
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		if m == nil || block == nil {
			return
		}

		m.RecordL2BlockInChannel(block)
	})
}

func Fuzz_Nosy_TestMetrics_RecordL2BlockInPendingQueue__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *TestMetrics
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var block *types.Block
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		if m == nil || block == nil {
			return
		}

		m.RecordL2BlockInPendingQueue(block)
	})
}

func Fuzz_Nosy_ThrottlingMetrics_PendingDABytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var nm *ThrottlingMetrics
		fill_err = tp.Fill(&nm)
		if fill_err != nil {
			return
		}
		if nm == nil {
			return
		}

		nm.PendingDABytes()
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

func Fuzz_Nosy_noopMetrics_PendingDABytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var nm *noopMetrics
		fill_err = tp.Fill(&nm)
		if fill_err != nil {
			return
		}
		if nm == nil {
			return
		}

		nm.PendingDABytes()
	})
}

func Fuzz_Nosy_noopMetrics_RecordBatchTxFailed__(f *testing.F) {
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

		_x1.RecordBatchTxFailed()
	})
}

func Fuzz_Nosy_noopMetrics_RecordBatchTxSubmitted__(f *testing.F) {
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

		_x1.RecordBatchTxSubmitted()
	})
}

func Fuzz_Nosy_noopMetrics_RecordBatchTxSuccess__(f *testing.F) {
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

		_x1.RecordBatchTxSuccess()
	})
}

func Fuzz_Nosy_noopMetrics_RecordBlobUsedBytes__(f *testing.F) {
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
		var _x2 int
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.RecordBlobUsedBytes(_x2)
	})
}

// skipping Fuzz_Nosy_noopMetrics_RecordChannelClosed__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_noopMetrics_RecordChannelFullySubmitted__(f *testing.F) {
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
		var _x2 derive.ChannelID
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.RecordChannelFullySubmitted(_x2)
	})
}

func Fuzz_Nosy_noopMetrics_RecordChannelOpened__(f *testing.F) {
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
		var _x2 derive.ChannelID
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var _x3 int
		fill_err = tp.Fill(&_x3)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.RecordChannelOpened(_x2, _x3)
	})
}

func Fuzz_Nosy_noopMetrics_RecordChannelTimedOut__(f *testing.F) {
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
		var _x2 derive.ChannelID
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.RecordChannelTimedOut(_x2)
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

func Fuzz_Nosy_noopMetrics_RecordL2BlockInChannel__(f *testing.F) {
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
		var _x2 *types.Block
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if _x1 == nil || _x2 == nil {
			return
		}

		_x1.RecordL2BlockInChannel(_x2)
	})
}

func Fuzz_Nosy_noopMetrics_RecordL2BlockInPendingQueue__(f *testing.F) {
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
		var _x2 *types.Block
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if _x1 == nil || _x2 == nil {
			return
		}

		_x1.RecordL2BlockInPendingQueue(_x2)
	})
}

func Fuzz_Nosy_noopMetrics_RecordL2BlocksAdded__(f *testing.F) {
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
		var _x2 eth.L2BlockRef
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var _x3 int
		fill_err = tp.Fill(&_x3)
		if fill_err != nil {
			return
		}
		var _x4 int
		fill_err = tp.Fill(&_x4)
		if fill_err != nil {
			return
		}
		var _x5 int
		fill_err = tp.Fill(&_x5)
		if fill_err != nil {
			return
		}
		var _x6 int
		fill_err = tp.Fill(&_x6)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.RecordL2BlocksAdded(_x2, _x3, _x4, _x5, _x6)
	})
}

func Fuzz_Nosy_noopMetrics_RecordL2BlocksLoaded__(f *testing.F) {
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
		var _x2 eth.L2BlockRef
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.RecordL2BlocksLoaded(_x2)
	})
}

func Fuzz_Nosy_noopMetrics_RecordLatestL1Block__(f *testing.F) {
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
		var l1ref eth.L1BlockRef
		fill_err = tp.Fill(&l1ref)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.RecordLatestL1Block(l1ref)
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

// skipping Fuzz_Nosy_Metricer_Document__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-batcher/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_PendingDABytes__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-batcher/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordBatchTxFailed__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-batcher/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordBatchTxSubmitted__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-batcher/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordBatchTxSuccess__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-batcher/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordBlobUsedBytes__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-batcher/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordChannelClosed__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-batcher/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordChannelFullySubmitted__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-batcher/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordChannelOpened__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-batcher/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordChannelTimedOut__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-batcher/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordInfo__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-batcher/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordL2BlockInChannel__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-batcher/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordL2BlockInPendingQueue__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-batcher/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordL2BlocksAdded__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-batcher/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordL2BlocksLoaded__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-batcher/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordLatestL1Block__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-batcher/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordUp__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-batcher/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_StartBalanceMetrics__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-batcher/metrics.Metricer

// skipping Fuzz_Nosy_ClosedReasonToNum__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_estimateBatchSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var block *types.Block
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		if block == nil {
			return
		}

		estimateBatchSize(block)
	})
}
