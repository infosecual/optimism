package l1

import (
	"context"
	"testing"

	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum/go-ethereum/common"
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

func Fuzz_Nosy_BlobFetcher_GetBlobs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BlobFetcher
		fill_err = tp.Fill(&b)
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
		if b == nil {
			return
		}

		b.GetBlobs(ctx, ref, hashes)
	})
}

func Fuzz_Nosy_CachingOracle_GetBlob__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o *CachingOracle
		fill_err = tp.Fill(&o)
		if fill_err != nil {
			return
		}
		var ref eth.L1BlockRef
		fill_err = tp.Fill(&ref)
		if fill_err != nil {
			return
		}
		var blobHash eth.IndexedBlobHash
		fill_err = tp.Fill(&blobHash)
		if fill_err != nil {
			return
		}
		if o == nil {
			return
		}

		o.GetBlob(ref, blobHash)
	})
}

func Fuzz_Nosy_CachingOracle_HeaderByBlockHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o *CachingOracle
		fill_err = tp.Fill(&o)
		if fill_err != nil {
			return
		}
		var blockHash common.Hash
		fill_err = tp.Fill(&blockHash)
		if fill_err != nil {
			return
		}
		if o == nil {
			return
		}

		o.HeaderByBlockHash(blockHash)
	})
}

func Fuzz_Nosy_CachingOracle_Precompile__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o *CachingOracle
		fill_err = tp.Fill(&o)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		var requiredGas uint64
		fill_err = tp.Fill(&requiredGas)
		if fill_err != nil {
			return
		}
		if o == nil {
			return
		}

		o.Precompile(address, input, requiredGas)
	})
}

func Fuzz_Nosy_CachingOracle_ReceiptsByBlockHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o *CachingOracle
		fill_err = tp.Fill(&o)
		if fill_err != nil {
			return
		}
		var blockHash common.Hash
		fill_err = tp.Fill(&blockHash)
		if fill_err != nil {
			return
		}
		if o == nil {
			return
		}

		o.ReceiptsByBlockHash(blockHash)
	})
}

func Fuzz_Nosy_CachingOracle_TransactionsByBlockHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o *CachingOracle
		fill_err = tp.Fill(&o)
		if fill_err != nil {
			return
		}
		var blockHash common.Hash
		fill_err = tp.Fill(&blockHash)
		if fill_err != nil {
			return
		}
		if o == nil {
			return
		}

		o.TransactionsByBlockHash(blockHash)
	})
}

func Fuzz_Nosy_OracleL1Client_FetchReceipts__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o *OracleL1Client
		fill_err = tp.Fill(&o)
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
		if o == nil {
			return
		}

		o.FetchReceipts(ctx, blockHash)
	})
}

func Fuzz_Nosy_OracleL1Client_InfoAndTxsByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o *OracleL1Client
		fill_err = tp.Fill(&o)
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
		if o == nil {
			return
		}

		o.InfoAndTxsByHash(ctx, hash)
	})
}

func Fuzz_Nosy_OracleL1Client_InfoByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o *OracleL1Client
		fill_err = tp.Fill(&o)
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
		if o == nil {
			return
		}

		o.InfoByHash(ctx, hash)
	})
}

func Fuzz_Nosy_OracleL1Client_L1BlockRefByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o *OracleL1Client
		fill_err = tp.Fill(&o)
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
		if o == nil {
			return
		}

		o.L1BlockRefByHash(ctx, hash)
	})
}

func Fuzz_Nosy_OracleL1Client_L1BlockRefByLabel__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o *OracleL1Client
		fill_err = tp.Fill(&o)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var label eth.BlockLabel
		fill_err = tp.Fill(&label)
		if fill_err != nil {
			return
		}
		if o == nil {
			return
		}

		o.L1BlockRefByLabel(ctx, label)
	})
}

func Fuzz_Nosy_OracleL1Client_L1BlockRefByNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o *OracleL1Client
		fill_err = tp.Fill(&o)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var number uint64
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		if o == nil {
			return
		}

		o.L1BlockRefByNumber(ctx, number)
	})
}

func Fuzz_Nosy_PreimageOracle_GetBlob__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *PreimageOracle
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var ref eth.L1BlockRef
		fill_err = tp.Fill(&ref)
		if fill_err != nil {
			return
		}
		var blobHash eth.IndexedBlobHash
		fill_err = tp.Fill(&blobHash)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.GetBlob(ref, blobHash)
	})
}

func Fuzz_Nosy_PreimageOracle_HeaderByBlockHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *PreimageOracle
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var blockHash common.Hash
		fill_err = tp.Fill(&blockHash)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.HeaderByBlockHash(blockHash)
	})
}

func Fuzz_Nosy_PreimageOracle_Precompile__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *PreimageOracle
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		var requiredGas uint64
		fill_err = tp.Fill(&requiredGas)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.Precompile(address, input, requiredGas)
	})
}

func Fuzz_Nosy_PreimageOracle_ReceiptsByBlockHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *PreimageOracle
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var blockHash common.Hash
		fill_err = tp.Fill(&blockHash)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.ReceiptsByBlockHash(blockHash)
	})
}

func Fuzz_Nosy_PreimageOracle_TransactionsByBlockHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *PreimageOracle
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var blockHash common.Hash
		fill_err = tp.Fill(&blockHash)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.TransactionsByBlockHash(blockHash)
	})
}

func Fuzz_Nosy_PreimageOracle_headerByBlockHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *PreimageOracle
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var blockHash common.Hash
		fill_err = tp.Fill(&blockHash)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.headerByBlockHash(blockHash)
	})
}

func Fuzz_Nosy_BlobHint_Hint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l BlobHint
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}

		l.Hint()
	})
}

func Fuzz_Nosy_BlockHeaderHint_Hint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l BlockHeaderHint
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}

		l.Hint()
	})
}

// skipping Fuzz_Nosy_Oracle_GetBlob__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-program/client/l1.Oracle

// skipping Fuzz_Nosy_Oracle_HeaderByBlockHash__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-program/client/l1.Oracle

// skipping Fuzz_Nosy_Oracle_Precompile__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-program/client/l1.Oracle

// skipping Fuzz_Nosy_Oracle_ReceiptsByBlockHash__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-program/client/l1.Oracle

// skipping Fuzz_Nosy_Oracle_TransactionsByBlockHash__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-program/client/l1.Oracle

func Fuzz_Nosy_PrecompileHint_Hint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l PrecompileHint
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}

		l.Hint()
	})
}

func Fuzz_Nosy_PrecompileHintV2_Hint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l PrecompileHintV2
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}

		l.Hint()
	})
}

func Fuzz_Nosy_ReceiptsHint_Hint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l ReceiptsHint
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}

		l.Hint()
	})
}

func Fuzz_Nosy_TransactionsHint_Hint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l TransactionsHint
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}

		l.Hint()
	})
}
