package txmgr

import (
	"testing"
	"github.com/ethereum-optimism/optimism/op-service"
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
	
	func Fuzz_Nosy_Config_Check__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var m *Config
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
	if m == nil {
		return
	}

	m.Check()
	})
}

// skipping Fuzz_Nosy_Queue[T any]_Send__ because parameters include func, chan, or unsupported interface: T

// skipping Fuzz_Nosy_Queue[T any]_TrySend__ because parameters include func, chan, or unsupported interface: T

func Fuzz_Nosy_Queue[T any]_Wait__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var q *Queue[T]
		fill_err = tp.Fill(&q)
		if fill_err != nil {
			return
		}
	if q == nil {
		return
	}

	q.Wait()
	})
}

func Fuzz_Nosy_Queue[T any]_groupContext__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var q *Queue[T]
		fill_err = tp.Fill(&q)
		if fill_err != nil {
			return
		}
	if q == nil {
		return
	}

	q.groupContext()
	})
}

// skipping Fuzz_Nosy_Queue[T any]_sendTx__ because parameters include func, chan, or unsupported interface: T

func Fuzz_Nosy_SendState_CriticalError__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var safeAbortNonceTooLowCount uint64
		fill_err = tp.Fill(&safeAbortNonceTooLowCount)
		if fill_err != nil {
			return
		}
		var unableToSendTimeout time.Duration
		fill_err = tp.Fill(&unableToSendTimeout)
		if fill_err != nil {
			return
		}

	s := NewSendState(safeAbortNonceTooLowCount, unableToSendTimeout)
	s.CriticalError()
	})
}

func Fuzz_Nosy_SendState_IsWaitingForConfirmation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var safeAbortNonceTooLowCount uint64
		fill_err = tp.Fill(&safeAbortNonceTooLowCount)
		if fill_err != nil {
			return
		}
		var unableToSendTimeout time.Duration
		fill_err = tp.Fill(&unableToSendTimeout)
		if fill_err != nil {
			return
		}

	s := NewSendState(safeAbortNonceTooLowCount, unableToSendTimeout)
	s.IsWaitingForConfirmation()
	})
}

// skipping Fuzz_Nosy_SendState_ProcessSendError__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_SendState_TxMined__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var safeAbortNonceTooLowCount uint64
		fill_err = tp.Fill(&safeAbortNonceTooLowCount)
		if fill_err != nil {
			return
		}
		var unableToSendTimeout time.Duration
		fill_err = tp.Fill(&unableToSendTimeout)
		if fill_err != nil {
			return
		}
		var txHash common.Hash
		fill_err = tp.Fill(&txHash)
		if fill_err != nil {
			return
		}

	s := NewSendState(safeAbortNonceTooLowCount, unableToSendTimeout)
	s.TxMined(txHash)
	})
}

func Fuzz_Nosy_SendState_TxNotMined__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var safeAbortNonceTooLowCount uint64
		fill_err = tp.Fill(&safeAbortNonceTooLowCount)
		if fill_err != nil {
			return
		}
		var unableToSendTimeout time.Duration
		fill_err = tp.Fill(&unableToSendTimeout)
		if fill_err != nil {
			return
		}
		var txHash common.Hash
		fill_err = tp.Fill(&txHash)
		if fill_err != nil {
			return
		}

	s := NewSendState(safeAbortNonceTooLowCount, unableToSendTimeout)
	s.TxNotMined(txHash)
	})
}

func Fuzz_Nosy_SimpleTxManager_API__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var m *SimpleTxManager
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
	if m == nil {
		return
	}

	m.API()
	})
}

func Fuzz_Nosy_SimpleTxManager_BlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var m *SimpleTxManager
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
	if m == nil {
		return
	}

	m.BlockNumber(ctx)
	})
}

func Fuzz_Nosy_SimpleTxManager_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var m *SimpleTxManager
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
	if m == nil {
		return
	}

	m.Close()
	})
}

func Fuzz_Nosy_SimpleTxManager_From__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var m *SimpleTxManager
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
	if m == nil {
		return
	}

	m.From()
	})
}

func Fuzz_Nosy_SimpleTxManager_GetBumpFeeRetryTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var m *SimpleTxManager
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
	if m == nil {
		return
	}

	m.GetBumpFeeRetryTime()
	})
}

func Fuzz_Nosy_SimpleTxManager_GetFeeLimitMultiplier__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var m *SimpleTxManager
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
	if m == nil {
		return
	}

	m.GetFeeLimitMultiplier()
	})
}

func Fuzz_Nosy_SimpleTxManager_GetFeeThreshold__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var m *SimpleTxManager
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
	if m == nil {
		return
	}

	m.GetFeeThreshold()
	})
}

func Fuzz_Nosy_SimpleTxManager_GetMinBaseFee__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var m *SimpleTxManager
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
	if m == nil {
		return
	}

	m.GetMinBaseFee()
	})
}

func Fuzz_Nosy_SimpleTxManager_GetMinBlobFee__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var m *SimpleTxManager
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
	if m == nil {
		return
	}

	m.GetMinBlobFee()
	})
}

func Fuzz_Nosy_SimpleTxManager_GetMinPriorityFee__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var m *SimpleTxManager
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
	if m == nil {
		return
	}

	m.GetMinPriorityFee()
	})
}

func Fuzz_Nosy_SimpleTxManager_IsClosed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var m *SimpleTxManager
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
	if m == nil {
		return
	}

	m.IsClosed()
	})
}

func Fuzz_Nosy_SimpleTxManager_Send__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var m *SimpleTxManager
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var candidate TxCandidate
		fill_err = tp.Fill(&candidate)
		if fill_err != nil {
			return
		}
	if m == nil {
		return
	}

	m.Send(ctx, candidate)
	})
}

// skipping Fuzz_Nosy_SimpleTxManager_SendAsync__ because parameters include func, chan, or unsupported interface: chan github.com/ethereum-optimism/optimism/op-service/txmgr.SendResponse

func Fuzz_Nosy_SimpleTxManager_SetBumpFeeRetryTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var m *SimpleTxManager
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var val time.Duration
		fill_err = tp.Fill(&val)
		if fill_err != nil {
			return
		}
	if m == nil {
		return
	}

	m.SetBumpFeeRetryTime(val)
	})
}

func Fuzz_Nosy_SimpleTxManager_SetFeeLimitMultiplier__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var m *SimpleTxManager
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var val uint64
		fill_err = tp.Fill(&val)
		if fill_err != nil {
			return
		}
	if m == nil {
		return
	}

	m.SetFeeLimitMultiplier(val)
	})
}

func Fuzz_Nosy_SimpleTxManager_SetFeeThreshold__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var m *SimpleTxManager
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var val *big.Int
		fill_err = tp.Fill(&val)
		if fill_err != nil {
			return
		}
	if m == nil|| val == nil {
		return
	}

	m.SetFeeThreshold(val)
	})
}

func Fuzz_Nosy_SimpleTxManager_SetMinBaseFee__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var m *SimpleTxManager
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var val *big.Int
		fill_err = tp.Fill(&val)
		if fill_err != nil {
			return
		}
	if m == nil|| val == nil {
		return
	}

	m.SetMinBaseFee(val)
	})
}

func Fuzz_Nosy_SimpleTxManager_SetMinBlobFee__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var m *SimpleTxManager
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var val *big.Int
		fill_err = tp.Fill(&val)
		if fill_err != nil {
			return
		}
	if m == nil|| val == nil {
		return
	}

	m.SetMinBlobFee(val)
	})
}

func Fuzz_Nosy_SimpleTxManager_SetMinPriorityFee__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var m *SimpleTxManager
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var val *big.Int
		fill_err = tp.Fill(&val)
		if fill_err != nil {
			return
		}
	if m == nil|| val == nil {
		return
	}

	m.SetMinPriorityFee(val)
	})
}

func Fuzz_Nosy_SimpleTxManager_SuggestGasPriceCaps__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var m *SimpleTxManager
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
	if m == nil {
		return
	}

	m.SuggestGasPriceCaps(ctx)
	})
}

func Fuzz_Nosy_SimpleTxManager_calcBlobFeeCap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var m *SimpleTxManager
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var blobBaseFee *big.Int
		fill_err = tp.Fill(&blobBaseFee)
		if fill_err != nil {
			return
		}
	if m == nil|| blobBaseFee == nil {
		return
	}

	m.calcBlobFeeCap(blobBaseFee)
	})
}

func Fuzz_Nosy_SimpleTxManager_checkBlobFeeLimits__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var m *SimpleTxManager
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var blobBaseFee *big.Int
		fill_err = tp.Fill(&blobBaseFee)
		if fill_err != nil {
			return
		}
		var bumpedBlobFee *big.Int
		fill_err = tp.Fill(&bumpedBlobFee)
		if fill_err != nil {
			return
		}
	if m == nil|| blobBaseFee == nil|| bumpedBlobFee == nil {
		return
	}

	m.checkBlobFeeLimits(blobBaseFee, bumpedBlobFee)
	})
}

func Fuzz_Nosy_SimpleTxManager_checkLimits__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var m *SimpleTxManager
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var tip *big.Int
		fill_err = tp.Fill(&tip)
		if fill_err != nil {
			return
		}
		var baseFee *big.Int
		fill_err = tp.Fill(&baseFee)
		if fill_err != nil {
			return
		}
		var bumpedTip *big.Int
		fill_err = tp.Fill(&bumpedTip)
		if fill_err != nil {
			return
		}
		var bumpedFee *big.Int
		fill_err = tp.Fill(&bumpedFee)
		if fill_err != nil {
			return
		}
	if m == nil|| tip == nil|| baseFee == nil|| bumpedTip == nil|| bumpedFee == nil {
		return
	}

	m.checkLimits(tip, baseFee, bumpedTip, bumpedFee)
	})
}

func Fuzz_Nosy_SimpleTxManager_craftTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var m *SimpleTxManager
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var candidate TxCandidate
		fill_err = tp.Fill(&candidate)
		if fill_err != nil {
			return
		}
	if m == nil {
		return
	}

	m.craftTx(ctx, candidate)
	})
}

func Fuzz_Nosy_SimpleTxManager_increaseGasPrice__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var m *SimpleTxManager
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var tx *types.Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
	if m == nil|| tx == nil {
		return
	}

	m.increaseGasPrice(ctx, tx)
	})
}

func Fuzz_Nosy_SimpleTxManager_prepare__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var m *SimpleTxManager
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var candidate TxCandidate
		fill_err = tp.Fill(&candidate)
		if fill_err != nil {
			return
		}
	if m == nil {
		return
	}

	m.prepare(ctx, candidate)
	})
}

func Fuzz_Nosy_SimpleTxManager_publishTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var m *SimpleTxManager
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var tx *types.Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		var sendState *SendState
		fill_err = tp.Fill(&sendState)
		if fill_err != nil {
			return
		}
	if m == nil|| tx == nil|| sendState == nil {
		return
	}

	m.publishTx(ctx, tx, sendState)
	})
}

func Fuzz_Nosy_SimpleTxManager_queryReceipt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var m *SimpleTxManager
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var txHash common.Hash
		fill_err = tp.Fill(&txHash)
		if fill_err != nil {
			return
		}
		var sendState *SendState
		fill_err = tp.Fill(&sendState)
		if fill_err != nil {
			return
		}
	if m == nil|| sendState == nil {
		return
	}

	m.queryReceipt(ctx, txHash, sendState)
	})
}

func Fuzz_Nosy_SimpleTxManager_resetNonce__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var m *SimpleTxManager
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
	if m == nil {
		return
	}

	m.resetNonce()
	})
}

func Fuzz_Nosy_SimpleTxManager_sendTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var m *SimpleTxManager
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var tx *types.Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
	if m == nil|| tx == nil {
		return
	}

	m.sendTx(ctx, tx)
	})
}

// skipping Fuzz_Nosy_SimpleTxManager_signWithNextNonce__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/types.TxData

func Fuzz_Nosy_SimpleTxManager_txLogger__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var m *SimpleTxManager
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var tx *types.Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		var logGas bool
		fill_err = tp.Fill(&logGas)
		if fill_err != nil {
			return
		}
	if m == nil|| tx == nil {
		return
	}

	m.txLogger(tx, logGas)
	})
}

// skipping Fuzz_Nosy_SimpleTxManager_waitForTx__ because parameters include func, chan, or unsupported interface: chan *github.com/ethereum/go-ethereum/core/types.Receipt

func Fuzz_Nosy_SimpleTxManager_waitMined__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var m *SimpleTxManager
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var tx *types.Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		var sendState *SendState
		fill_err = tp.Fill(&sendState)
		if fill_err != nil {
			return
		}
	if m == nil|| tx == nil|| sendState == nil {
		return
	}

	m.waitMined(ctx, tx, sendState)
	})
}

func Fuzz_Nosy_SimpleTxmgrAPI_GetBumpFeeRetryTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var a *SimpleTxmgrAPI
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
	if a == nil {
		return
	}

	a.GetBumpFeeRetryTime(_x2)
	})
}

func Fuzz_Nosy_SimpleTxmgrAPI_GetFeeThreshold__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var a *SimpleTxmgrAPI
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
	if a == nil {
		return
	}

	a.GetFeeThreshold(_x2)
	})
}

func Fuzz_Nosy_SimpleTxmgrAPI_GetMinBaseFee__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var a *SimpleTxmgrAPI
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
	if a == nil {
		return
	}

	a.GetMinBaseFee(_x2)
	})
}

func Fuzz_Nosy_SimpleTxmgrAPI_GetMinBlobFee__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var a *SimpleTxmgrAPI
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
	if a == nil {
		return
	}

	a.GetMinBlobFee(_x2)
	})
}

func Fuzz_Nosy_SimpleTxmgrAPI_GetMinPriorityFee__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var a *SimpleTxmgrAPI
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
	if a == nil {
		return
	}

	a.GetMinPriorityFee(_x2)
	})
}

func Fuzz_Nosy_SimpleTxmgrAPI_SetBumpFeeRetryTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var a *SimpleTxmgrAPI
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var val time.Duration
		fill_err = tp.Fill(&val)
		if fill_err != nil {
			return
		}
	if a == nil {
		return
	}

	a.SetBumpFeeRetryTime(_x2, val)
	})
}

func Fuzz_Nosy_SimpleTxmgrAPI_SetFeeThreshold__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var a *SimpleTxmgrAPI
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var val *big.Int
		fill_err = tp.Fill(&val)
		if fill_err != nil {
			return
		}
	if a == nil|| val == nil {
		return
	}

	a.SetFeeThreshold(_x2, val)
	})
}

func Fuzz_Nosy_SimpleTxmgrAPI_SetMinBaseFee__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var a *SimpleTxmgrAPI
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var val *big.Int
		fill_err = tp.Fill(&val)
		if fill_err != nil {
			return
		}
	if a == nil|| val == nil {
		return
	}

	a.SetMinBaseFee(_x2, val)
	})
}

func Fuzz_Nosy_SimpleTxmgrAPI_SetMinBlobFee__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var a *SimpleTxmgrAPI
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var val *big.Int
		fill_err = tp.Fill(&val)
		if fill_err != nil {
			return
		}
	if a == nil|| val == nil {
		return
	}

	a.SetMinBlobFee(_x2, val)
	})
}

func Fuzz_Nosy_SimpleTxmgrAPI_SetMinPriorityFee__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var a *SimpleTxmgrAPI
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var val *big.Int
		fill_err = tp.Fill(&val)
		if fill_err != nil {
			return
		}
	if a == nil|| val == nil {
		return
	}

	a.SetMinPriorityFee(_x2, val)
	})
}

func Fuzz_Nosy_TestTxManager_JamTxPool__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var m *TestTxManager
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var candidate TxCandidate
		fill_err = tp.Fill(&candidate)
		if fill_err != nil {
			return
		}
	if m == nil {
		return
	}

	m.JamTxPool(ctx, candidate)
	})
}

func Fuzz_Nosy_TestTxManager_WaitOnJammingTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var m *TestTxManager
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
	if m == nil {
		return
	}

	m.WaitOnJammingTx(ctx)
	})
}

func Fuzz_Nosy_TestTxManager_makeStuckTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var m *TestTxManager
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var candidate TxCandidate
		fill_err = tp.Fill(&candidate)
		if fill_err != nil {
			return
		}
	if m == nil {
		return
	}

	m.makeStuckTx(ctx, candidate)
	})
}

func Fuzz_Nosy_CLIConfig_Check__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var ctx *cli.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
	if ctx == nil {
		return
	}

	m := ReadCLIConfig(ctx)
	m.Check()
	})
}

// skipping Fuzz_Nosy_ETHBackend_BlockNumber__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/txmgr.ETHBackend

// skipping Fuzz_Nosy_ETHBackend_CallContract__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/txmgr.ETHBackend

// skipping Fuzz_Nosy_ETHBackend_Close__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/txmgr.ETHBackend

// skipping Fuzz_Nosy_ETHBackend_EstimateGas__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/txmgr.ETHBackend

// skipping Fuzz_Nosy_ETHBackend_HeaderByNumber__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/txmgr.ETHBackend

// skipping Fuzz_Nosy_ETHBackend_NonceAt__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/txmgr.ETHBackend

// skipping Fuzz_Nosy_ETHBackend_PendingNonceAt__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/txmgr.ETHBackend

// skipping Fuzz_Nosy_ETHBackend_SendTransaction__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/txmgr.ETHBackend

// skipping Fuzz_Nosy_ETHBackend_SuggestGasTipCap__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/txmgr.ETHBackend

// skipping Fuzz_Nosy_ETHBackend_TransactionReceipt__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/txmgr.ETHBackend

// skipping Fuzz_Nosy_TxManager_API__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/txmgr.TxManager

// skipping Fuzz_Nosy_TxManager_BlockNumber__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/txmgr.TxManager

// skipping Fuzz_Nosy_TxManager_Close__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/txmgr.TxManager

// skipping Fuzz_Nosy_TxManager_From__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/txmgr.TxManager

// skipping Fuzz_Nosy_TxManager_IsClosed__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/txmgr.TxManager

// skipping Fuzz_Nosy_TxManager_Send__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/txmgr.TxManager

// skipping Fuzz_Nosy_TxManager_SendAsync__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/txmgr.TxManager

// skipping Fuzz_Nosy_TxManager_SuggestGasPriceCaps__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/txmgr.TxManager

func Fuzz_Nosy_CLIFlags__(f *testing.F) {
	f.Fuzz(func(t *testing.T, envPrefix string) {
	CLIFlags(envPrefix)
	})
}

func Fuzz_Nosy_CLIFlagsWithDefaults__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var envPrefix string
		fill_err = tp.Fill(&envPrefix)
		if fill_err != nil {
			return
		}
		var defaults DefaultFlagValues
		fill_err = tp.Fill(&defaults)
		if fill_err != nil {
			return
		}

	CLIFlagsWithDefaults(envPrefix, defaults)
	})
}

// skipping Fuzz_Nosy_DefaultGasPriceEstimatorFn__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/txmgr.ETHBackend

func Fuzz_Nosy_MakeSidecar__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var blobs []*eth.Blob
		fill_err = tp.Fill(&blobs)
		if fill_err != nil {
			return
		}

	MakeSidecar(blobs)
	})
}

// skipping Fuzz_Nosy_errStringMatch__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_updateFees__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/log.Logger

