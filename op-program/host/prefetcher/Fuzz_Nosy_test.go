package prefetcher

import (
	"context"
	"testing"

	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
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

func Fuzz_Nosy_Prefetcher_GetPreimage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *Prefetcher
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var key common.Hash
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.GetPreimage(ctx, key)
	})
}

func Fuzz_Nosy_Prefetcher_Hint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *Prefetcher
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var hint string
		fill_err = tp.Fill(&hint)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.Hint(hint)
	})
}

func Fuzz_Nosy_Prefetcher_prefetch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *Prefetcher
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var hint string
		fill_err = tp.Fill(&hint)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.prefetch(ctx, hint)
	})
}

func Fuzz_Nosy_Prefetcher_storeReceipts__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *Prefetcher
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var receipts types.Receipts
		fill_err = tp.Fill(&receipts)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.storeReceipts(receipts)
	})
}

func Fuzz_Nosy_Prefetcher_storeTransactions__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *Prefetcher
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var txs types.Transactions
		fill_err = tp.Fill(&txs)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.storeTransactions(txs)
	})
}

func Fuzz_Nosy_Prefetcher_storeTrieNodes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *Prefetcher
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var values []hexutil.Bytes
		fill_err = tp.Fill(&values)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.storeTrieNodes(values)
	})
}

func Fuzz_Nosy_RetryingL1BlobSource_GetBlobSidecars__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *RetryingL1BlobSource
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var ref eth.L1BlockRef
		fill_err = tp.Fill(&ref)
		if fill_err != nil {
			return
		}
		var hashes []eth.IndexedBlobHash
		fill_err = tp.Fill(&hashes)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.GetBlobSidecars(ctx, ref, hashes)
	})
}

func Fuzz_Nosy_RetryingL1BlobSource_GetBlobs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *RetryingL1BlobSource
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var ref eth.L1BlockRef
		fill_err = tp.Fill(&ref)
		if fill_err != nil {
			return
		}
		var hashes []eth.IndexedBlobHash
		fill_err = tp.Fill(&hashes)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.GetBlobs(ctx, ref, hashes)
	})
}

func Fuzz_Nosy_RetryingL1Source_FetchReceipts__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *RetryingL1Source
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var blockHash common.Hash
		fill_err = tp.Fill(&blockHash)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.FetchReceipts(ctx, blockHash)
	})
}

func Fuzz_Nosy_RetryingL1Source_InfoAndTxsByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *RetryingL1Source
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var blockHash common.Hash
		fill_err = tp.Fill(&blockHash)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.InfoAndTxsByHash(ctx, blockHash)
	})
}

func Fuzz_Nosy_RetryingL1Source_InfoByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *RetryingL1Source
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var blockHash common.Hash
		fill_err = tp.Fill(&blockHash)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.InfoByHash(ctx, blockHash)
	})
}

func Fuzz_Nosy_RetryingL2Source_CodeByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *RetryingL2Source
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.CodeByHash(ctx, hash)
	})
}

func Fuzz_Nosy_RetryingL2Source_InfoAndTxsByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *RetryingL2Source
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var blockHash common.Hash
		fill_err = tp.Fill(&blockHash)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.InfoAndTxsByHash(ctx, blockHash)
	})
}

func Fuzz_Nosy_RetryingL2Source_NodeByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *RetryingL2Source
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.NodeByHash(ctx, hash)
	})
}

func Fuzz_Nosy_RetryingL2Source_OutputByRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *RetryingL2Source
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var root common.Hash
		fill_err = tp.Fill(&root)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.OutputByRoot(ctx, root)
	})
}

// skipping Fuzz_Nosy_L1BlobSource_GetBlobSidecars__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-program/host/prefetcher.L1BlobSource

// skipping Fuzz_Nosy_L1BlobSource_GetBlobs__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-program/host/prefetcher.L1BlobSource

// skipping Fuzz_Nosy_L1Source_FetchReceipts__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-program/host/prefetcher.L1Source

// skipping Fuzz_Nosy_L1Source_InfoAndTxsByHash__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-program/host/prefetcher.L1Source

// skipping Fuzz_Nosy_L1Source_InfoByHash__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-program/host/prefetcher.L1Source

// skipping Fuzz_Nosy_L2Source_CodeByHash__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-program/host/prefetcher.L2Source

// skipping Fuzz_Nosy_L2Source_InfoAndTxsByHash__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-program/host/prefetcher.L2Source

// skipping Fuzz_Nosy_L2Source_NodeByHash__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-program/host/prefetcher.L2Source

// skipping Fuzz_Nosy_L2Source_OutputByRoot__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-program/host/prefetcher.L2Source

func Fuzz_Nosy_parseHint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, hint string) {
		parseHint(hint)
	})
}
