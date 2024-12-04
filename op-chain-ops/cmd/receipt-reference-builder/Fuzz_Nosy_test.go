package main

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
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

func Fuzz_Nosy_ByFirst_Len__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a ByFirst
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}

		a.Len()
	})
}

func Fuzz_Nosy_ByFirst_Less__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a ByFirst
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var i int
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
		var j int
		fill_err = tp.Fill(&j)
		if fill_err != nil {
			return
		}

		a.Less(i, j)
	})
}

func Fuzz_Nosy_ByFirst_Swap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a ByFirst
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var i int
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
		var j int
		fill_err = tp.Fill(&j)
		if fill_err != nil {
			return
		}

		a.Swap(i, j)
	})
}

// skipping Fuzz_Nosy_aggregateReaderWriter_readAggregate__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-chain-ops/cmd/receipt-reference-builder.aggregateReaderWriter

// skipping Fuzz_Nosy_aggregateReaderWriter_writeAggregate__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-chain-ops/cmd/receipt-reference-builder.aggregateReaderWriter

func Fuzz_Nosy_gobAggregateReaderWriter_readAggregate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w gobAggregateReaderWriter
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var f2 string
		fill_err = tp.Fill(&f2)
		if fill_err != nil {
			return
		}

		w.readAggregate(f2)
	})
}

func Fuzz_Nosy_gobAggregateReaderWriter_writeAggregate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w gobAggregateReaderWriter
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var a aggregate
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var o string
		fill_err = tp.Fill(&o)
		if fill_err != nil {
			return
		}

		w.writeAggregate(a, o)
	})
}

func Fuzz_Nosy_jsonAggregateReaderWriter_readAggregate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w jsonAggregateReaderWriter
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var f2 string
		fill_err = tp.Fill(&f2)
		if fill_err != nil {
			return
		}

		w.readAggregate(f2)
	})
}

func Fuzz_Nosy_jsonAggregateReaderWriter_writeAggregate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w jsonAggregateReaderWriter
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var a aggregate
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var o string
		fill_err = tp.Fill(&o)
		if fill_err != nil {
			return
		}

		w.writeAggregate(a, o)
	})
}

func Fuzz_Nosy_batchBlockByNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var c *ethclient.Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var blockNumbers []rpc.BlockNumber
		fill_err = tp.Fill(&blockNumbers)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		batchBlockByNumber(ctx, c, blockNumbers)
	})
}

// skipping Fuzz_Nosy_checkTransaction__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/log.Logger

// skipping Fuzz_Nosy_doWork__ because parameters include func, chan, or unsupported interface: chan github.com/ethereum-optimism/optimism/op-chain-ops/cmd/receipt-reference-builder.result

// skipping Fuzz_Nosy_processBlockRange__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/log.Logger

func Fuzz_Nosy_splitBatchRange__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b batchRange
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		splitBatchRange(b)
	})
}

// skipping Fuzz_Nosy_startWorker__ because parameters include func, chan, or unsupported interface: chan github.com/ethereum-optimism/optimism/op-chain-ops/cmd/receipt-reference-builder.batchRange

func Fuzz_Nosy_toBatches__(f *testing.F) {
	f.Fuzz(func(t *testing.T, start uint64, end uint64, size uint64) {
		toBatches(start, end, size)
	})
}
