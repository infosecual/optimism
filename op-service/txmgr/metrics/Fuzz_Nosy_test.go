package metrics

import (
	"math/big"
	"testing"

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

func Fuzz_Nosy_NoopTxMetrics_RPCError__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *NoopTxMetrics
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.RPCError()
	})
}

func Fuzz_Nosy_NoopTxMetrics_RecordBaseFee__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *NoopTxMetrics
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var _x2 *big.Int
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if _x1 == nil || _x2 == nil {
			return
		}

		_x1.RecordBaseFee(_x2)
	})
}

func Fuzz_Nosy_NoopTxMetrics_RecordBlobBaseFee__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *NoopTxMetrics
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var _x2 *big.Int
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if _x1 == nil || _x2 == nil {
			return
		}

		_x1.RecordBlobBaseFee(_x2)
	})
}

func Fuzz_Nosy_NoopTxMetrics_RecordGasBumpCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *NoopTxMetrics
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

		_x1.RecordGasBumpCount(_x2)
	})
}

func Fuzz_Nosy_NoopTxMetrics_RecordNonce__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *NoopTxMetrics
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

		_x1.RecordNonce(_x2)
	})
}

func Fuzz_Nosy_NoopTxMetrics_RecordPendingTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *NoopTxMetrics
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var _x2 int64
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.RecordPendingTx(_x2)
	})
}

func Fuzz_Nosy_NoopTxMetrics_RecordTipCap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *NoopTxMetrics
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var _x2 *big.Int
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if _x1 == nil || _x2 == nil {
			return
		}

		_x1.RecordTipCap(_x2)
	})
}

func Fuzz_Nosy_NoopTxMetrics_RecordTxConfirmationLatency__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *NoopTxMetrics
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var _x2 int64
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.RecordTxConfirmationLatency(_x2)
	})
}

func Fuzz_Nosy_NoopTxMetrics_TxConfirmed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *NoopTxMetrics
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var _x2 *types.Receipt
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if _x1 == nil || _x2 == nil {
			return
		}

		_x1.TxConfirmed(_x2)
	})
}

func Fuzz_Nosy_NoopTxMetrics_TxPublished__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *NoopTxMetrics
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var _x2 string
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.TxPublished(_x2)
	})
}

func Fuzz_Nosy_TxMetrics_RPCError__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *TxMetrics
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.RPCError()
	})
}

func Fuzz_Nosy_TxMetrics_RecordBaseFee__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *TxMetrics
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var baseFee *big.Int
		fill_err = tp.Fill(&baseFee)
		if fill_err != nil {
			return
		}
		if t1 == nil || baseFee == nil {
			return
		}

		t1.RecordBaseFee(baseFee)
	})
}

func Fuzz_Nosy_TxMetrics_RecordBlobBaseFee__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *TxMetrics
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var blobBaseFee *big.Int
		fill_err = tp.Fill(&blobBaseFee)
		if fill_err != nil {
			return
		}
		if t1 == nil || blobBaseFee == nil {
			return
		}

		t1.RecordBlobBaseFee(blobBaseFee)
	})
}

func Fuzz_Nosy_TxMetrics_RecordGasBumpCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *TxMetrics
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var times int
		fill_err = tp.Fill(&times)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.RecordGasBumpCount(times)
	})
}

func Fuzz_Nosy_TxMetrics_RecordNonce__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *TxMetrics
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var nonce uint64
		fill_err = tp.Fill(&nonce)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.RecordNonce(nonce)
	})
}

func Fuzz_Nosy_TxMetrics_RecordPendingTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *TxMetrics
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var pending int64
		fill_err = tp.Fill(&pending)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.RecordPendingTx(pending)
	})
}

func Fuzz_Nosy_TxMetrics_RecordTipCap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *TxMetrics
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var tipcap *big.Int
		fill_err = tp.Fill(&tipcap)
		if fill_err != nil {
			return
		}
		if t1 == nil || tipcap == nil {
			return
		}

		t1.RecordTipCap(tipcap)
	})
}

func Fuzz_Nosy_TxMetrics_RecordTxConfirmationLatency__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *TxMetrics
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var latency int64
		fill_err = tp.Fill(&latency)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.RecordTxConfirmationLatency(latency)
	})
}

func Fuzz_Nosy_TxMetrics_TxConfirmed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *TxMetrics
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var receipt *types.Receipt
		fill_err = tp.Fill(&receipt)
		if fill_err != nil {
			return
		}
		if t1 == nil || receipt == nil {
			return
		}

		t1.TxConfirmed(receipt)
	})
}

func Fuzz_Nosy_TxMetrics_TxPublished__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *TxMetrics
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var errString string
		fill_err = tp.Fill(&errString)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.TxPublished(errString)
	})
}

// skipping Fuzz_Nosy_TxMetricer_RPCError__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/txmgr/metrics.TxMetricer

// skipping Fuzz_Nosy_TxMetricer_RecordBaseFee__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/txmgr/metrics.TxMetricer

// skipping Fuzz_Nosy_TxMetricer_RecordBlobBaseFee__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/txmgr/metrics.TxMetricer

// skipping Fuzz_Nosy_TxMetricer_RecordGasBumpCount__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/txmgr/metrics.TxMetricer

// skipping Fuzz_Nosy_TxMetricer_RecordNonce__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/txmgr/metrics.TxMetricer

// skipping Fuzz_Nosy_TxMetricer_RecordPendingTx__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/txmgr/metrics.TxMetricer

// skipping Fuzz_Nosy_TxMetricer_RecordTipCap__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/txmgr/metrics.TxMetricer

// skipping Fuzz_Nosy_TxMetricer_RecordTxConfirmationLatency__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/txmgr/metrics.TxMetricer

// skipping Fuzz_Nosy_TxMetricer_TxConfirmed__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/txmgr/metrics.TxMetricer

// skipping Fuzz_Nosy_TxMetricer_TxPublished__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/txmgr/metrics.TxMetricer

func Fuzz_Nosy_receiptStatusString__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var receipt *types.Receipt
		fill_err = tp.Fill(&receipt)
		if fill_err != nil {
			return
		}
		if receipt == nil {
			return
		}

		receiptStatusString(receipt)
	})
}
