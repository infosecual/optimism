package testutils

import (
	"context"
	"math/big"
	"math/rand"
	"testing"
	"time"

	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
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

func Fuzz_Nosy_FakeChainSource_AdvanceL1__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *FakeChainSource
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.AdvanceL1()
	})
}

func Fuzz_Nosy_FakeChainSource_ForkchoiceUpdate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *FakeChainSource
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var state *eth.ForkchoiceState
		fill_err = tp.Fill(&state)
		if fill_err != nil {
			return
		}
		var attr *eth.PayloadAttributes
		fill_err = tp.Fill(&attr)
		if fill_err != nil {
			return
		}
		if m == nil || state == nil || attr == nil {
			return
		}

		m.ForkchoiceUpdate(ctx, state, attr)
	})
}

func Fuzz_Nosy_FakeChainSource_L1BlockRefByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *FakeChainSource
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var l1Hash common.Hash
		fill_err = tp.Fill(&l1Hash)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.L1BlockRefByHash(ctx, l1Hash)
	})
}

func Fuzz_Nosy_FakeChainSource_L1BlockRefByLabel__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *FakeChainSource
		fill_err = tp.Fill(&m)
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
		if m == nil {
			return
		}

		m.L1BlockRefByLabel(ctx, label)
	})
}

func Fuzz_Nosy_FakeChainSource_L1BlockRefByNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *FakeChainSource
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var l1Num uint64
		fill_err = tp.Fill(&l1Num)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.L1BlockRefByNumber(ctx, l1Num)
	})
}

func Fuzz_Nosy_FakeChainSource_L1Head__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *FakeChainSource
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.L1Head()
	})
}

func Fuzz_Nosy_FakeChainSource_L1Range__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *FakeChainSource
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var base eth.BlockID
		fill_err = tp.Fill(&base)
		if fill_err != nil {
			return
		}
		var max uint64
		fill_err = tp.Fill(&max)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.L1Range(ctx, base, max)
	})
}

func Fuzz_Nosy_FakeChainSource_L2BlockRefByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *FakeChainSource
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var l2Hash common.Hash
		fill_err = tp.Fill(&l2Hash)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.L2BlockRefByHash(ctx, l2Hash)
	})
}

func Fuzz_Nosy_FakeChainSource_L2BlockRefByLabel__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *FakeChainSource
		fill_err = tp.Fill(&m)
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
		if m == nil {
			return
		}

		m.L2BlockRefByLabel(ctx, label)
	})
}

func Fuzz_Nosy_FakeChainSource_L2BlockRefByNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *FakeChainSource
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var l2Num *big.Int
		fill_err = tp.Fill(&l2Num)
		if fill_err != nil {
			return
		}
		if m == nil || l2Num == nil {
			return
		}

		m.L2BlockRefByNumber(ctx, l2Num)
	})
}

func Fuzz_Nosy_FakeChainSource_ReorgL1__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *FakeChainSource
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.ReorgL1()
	})
}

func Fuzz_Nosy_FakeChainSource_SetL2Finalized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *FakeChainSource
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var finalized common.Hash
		fill_err = tp.Fill(&finalized)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.SetL2Finalized(finalized)
	})
}

func Fuzz_Nosy_FakeChainSource_SetL2Head__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *FakeChainSource
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var head int
		fill_err = tp.Fill(&head)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.SetL2Head(head)
	})
}

func Fuzz_Nosy_FakeChainSource_SetL2Safe__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *FakeChainSource
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var safe common.Hash
		fill_err = tp.Fill(&safe)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.SetL2Safe(safe)
	})
}

func Fuzz_Nosy_FakeInteropBackend_CrossDerivedFrom__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *FakeInteropBackend
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var chainID types.ChainID
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var derived eth.BlockID
		fill_err = tp.Fill(&derived)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.CrossDerivedFrom(ctx, chainID, derived)
	})
}

func Fuzz_Nosy_FakeInteropBackend_Finalized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *FakeInteropBackend
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var chainID types.ChainID
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Finalized(ctx, chainID)
	})
}

func Fuzz_Nosy_FakeInteropBackend_SafeView__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *FakeInteropBackend
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var chainID types.ChainID
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var safe types.ReferenceView
		fill_err = tp.Fill(&safe)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.SafeView(ctx, chainID, safe)
	})
}

func Fuzz_Nosy_FakeInteropBackend_UnsafeView__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *FakeInteropBackend
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var chainID types.ChainID
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var unsafe types.ReferenceView
		fill_err = tp.Fill(&unsafe)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.UnsafeView(ctx, chainID, unsafe)
	})
}

func Fuzz_Nosy_FakeInteropBackend_UpdateFinalizedL1__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *FakeInteropBackend
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var chainID types.ChainID
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var finalized eth.L1BlockRef
		fill_err = tp.Fill(&finalized)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.UpdateFinalizedL1(ctx, chainID, finalized)
	})
}

func Fuzz_Nosy_FakeInteropBackend_UpdateLocalSafe__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *FakeInteropBackend
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var chainID types.ChainID
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var derivedFrom eth.L1BlockRef
		fill_err = tp.Fill(&derivedFrom)
		if fill_err != nil {
			return
		}
		var lastDerived eth.L1BlockRef
		fill_err = tp.Fill(&lastDerived)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.UpdateLocalSafe(ctx, chainID, derivedFrom, lastDerived)
	})
}

func Fuzz_Nosy_FakeInteropBackend_UpdateLocalUnsafe__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *FakeInteropBackend
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var chainID types.ChainID
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var head eth.L1BlockRef
		fill_err = tp.Fill(&head)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.UpdateLocalUnsafe(ctx, chainID, head)
	})
}

// skipping Fuzz_Nosy_MockBlobsFetcher_ExpectOnGetBlobSidecars__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_MockBlobsFetcher_ExpectOnGetBlobs__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_MockBlobsFetcher_GetBlobSidecars__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cl *MockBlobsFetcher
		fill_err = tp.Fill(&cl)
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
		if cl == nil {
			return
		}

		cl.GetBlobSidecars(ctx, ref, hashes)
	})
}

func Fuzz_Nosy_MockBlobsFetcher_GetBlobs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cl *MockBlobsFetcher
		fill_err = tp.Fill(&cl)
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
		if cl == nil {
			return
		}

		cl.GetBlobs(ctx, ref, hashes)
	})
}

func Fuzz_Nosy_MockBlockInfo_BaseFee__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rng *rand.Rand
		fill_err = tp.Fill(&rng)
		if fill_err != nil {
			return
		}
		if rng == nil {
			return
		}

		l := RandomBlockInfo(rng)
		l.BaseFee()
	})
}

func Fuzz_Nosy_MockBlockInfo_BlobBaseFee__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rng *rand.Rand
		fill_err = tp.Fill(&rng)
		if fill_err != nil {
			return
		}
		if rng == nil {
			return
		}

		l := RandomBlockInfo(rng)
		l.BlobBaseFee()
	})
}

func Fuzz_Nosy_MockBlockInfo_BlockRef__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rng *rand.Rand
		fill_err = tp.Fill(&rng)
		if fill_err != nil {
			return
		}
		if rng == nil {
			return
		}

		l := RandomBlockInfo(rng)
		l.BlockRef()
	})
}

func Fuzz_Nosy_MockBlockInfo_Coinbase__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rng *rand.Rand
		fill_err = tp.Fill(&rng)
		if fill_err != nil {
			return
		}
		if rng == nil {
			return
		}

		l := RandomBlockInfo(rng)
		l.Coinbase()
	})
}

func Fuzz_Nosy_MockBlockInfo_GasLimit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rng *rand.Rand
		fill_err = tp.Fill(&rng)
		if fill_err != nil {
			return
		}
		if rng == nil {
			return
		}

		l := RandomBlockInfo(rng)
		l.GasLimit()
	})
}

func Fuzz_Nosy_MockBlockInfo_GasUsed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rng *rand.Rand
		fill_err = tp.Fill(&rng)
		if fill_err != nil {
			return
		}
		if rng == nil {
			return
		}

		l := RandomBlockInfo(rng)
		l.GasUsed()
	})
}

func Fuzz_Nosy_MockBlockInfo_Hash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rng *rand.Rand
		fill_err = tp.Fill(&rng)
		if fill_err != nil {
			return
		}
		if rng == nil {
			return
		}

		l := RandomBlockInfo(rng)
		l.Hash()
	})
}

func Fuzz_Nosy_MockBlockInfo_HeaderRLP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rng *rand.Rand
		fill_err = tp.Fill(&rng)
		if fill_err != nil {
			return
		}
		if rng == nil {
			return
		}

		l := RandomBlockInfo(rng)
		l.HeaderRLP()
	})
}

func Fuzz_Nosy_MockBlockInfo_ID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rng *rand.Rand
		fill_err = tp.Fill(&rng)
		if fill_err != nil {
			return
		}
		if rng == nil {
			return
		}

		l := RandomBlockInfo(rng)
		l.ID()
	})
}

func Fuzz_Nosy_MockBlockInfo_MixDigest__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rng *rand.Rand
		fill_err = tp.Fill(&rng)
		if fill_err != nil {
			return
		}
		if rng == nil {
			return
		}

		l := RandomBlockInfo(rng)
		l.MixDigest()
	})
}

func Fuzz_Nosy_MockBlockInfo_NumberU64__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rng *rand.Rand
		fill_err = tp.Fill(&rng)
		if fill_err != nil {
			return
		}
		if rng == nil {
			return
		}

		l := RandomBlockInfo(rng)
		l.NumberU64()
	})
}

func Fuzz_Nosy_MockBlockInfo_ParentBeaconRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rng *rand.Rand
		fill_err = tp.Fill(&rng)
		if fill_err != nil {
			return
		}
		if rng == nil {
			return
		}

		l := RandomBlockInfo(rng)
		l.ParentBeaconRoot()
	})
}

func Fuzz_Nosy_MockBlockInfo_ParentHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rng *rand.Rand
		fill_err = tp.Fill(&rng)
		if fill_err != nil {
			return
		}
		if rng == nil {
			return
		}

		l := RandomBlockInfo(rng)
		l.ParentHash()
	})
}

func Fuzz_Nosy_MockBlockInfo_ReceiptHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rng *rand.Rand
		fill_err = tp.Fill(&rng)
		if fill_err != nil {
			return
		}
		if rng == nil {
			return
		}

		l := RandomBlockInfo(rng)
		l.ReceiptHash()
	})
}

func Fuzz_Nosy_MockBlockInfo_Root__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rng *rand.Rand
		fill_err = tp.Fill(&rng)
		if fill_err != nil {
			return
		}
		if rng == nil {
			return
		}

		l := RandomBlockInfo(rng)
		l.Root()
	})
}

func Fuzz_Nosy_MockBlockInfo_Time__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rng *rand.Rand
		fill_err = tp.Fill(&rng)
		if fill_err != nil {
			return
		}
		if rng == nil {
			return
		}

		l := RandomBlockInfo(rng)
		l.Time()
	})
}

func Fuzz_Nosy_MockClient_BalanceAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockClient
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var account common.Address
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		var block *big.Int
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		if m == nil || block == nil {
			return
		}

		m.BalanceAt(ctx, account, block)
	})
}

func Fuzz_Nosy_MockClient_BlockByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockClient
		fill_err = tp.Fill(&m)
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
		if m == nil {
			return
		}

		m.BlockByHash(ctx, hash)
	})
}

func Fuzz_Nosy_MockClient_BlockByNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockClient
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var number *big.Int
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		if m == nil || number == nil {
			return
		}

		m.BlockByNumber(ctx, number)
	})
}

func Fuzz_Nosy_MockClient_BlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockClient
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

func Fuzz_Nosy_MockClient_CallContract__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockClient
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var msg ethereum.CallMsg
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		var blockNumber *big.Int
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		if m == nil || blockNumber == nil {
			return
		}

		m.CallContract(ctx, msg, blockNumber)
	})
}

func Fuzz_Nosy_MockClient_CallContractAtHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockClient
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var msg ethereum.CallMsg
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		var blockHash common.Hash
		fill_err = tp.Fill(&blockHash)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.CallContractAtHash(ctx, msg, blockHash)
	})
}

func Fuzz_Nosy_MockClient_ChainID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockClient
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

		m.ChainID(ctx)
	})
}

func Fuzz_Nosy_MockClient_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockClient
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

func Fuzz_Nosy_MockClient_CodeAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockClient
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var account common.Address
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		var block *big.Int
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		if m == nil || block == nil {
			return
		}

		m.CodeAt(ctx, account, block)
	})
}

func Fuzz_Nosy_MockClient_EstimateGas__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockClient
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var msg ethereum.CallMsg
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.EstimateGas(ctx, msg)
	})
}

// skipping Fuzz_Nosy_MockClient_ExpectBalanceAt__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_MockClient_ExpectBlockByHash__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_MockClient_ExpectBlockByNumber__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_MockClient_ExpectBlockNumber__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_MockClient_ExpectCallContract__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_MockClient_ExpectCallContractAtHash__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_MockClient_ExpectChainID__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_MockClient_ExpectClose__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockClient
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.ExpectClose()
	})
}

// skipping Fuzz_Nosy_MockClient_ExpectCodeAt__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_MockClient_ExpectEstimateGas__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_MockClient_ExpectFilterLogs__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_MockClient_ExpectHeaderByHash__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_MockClient_ExpectHeaderByNumber__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_MockClient_ExpectNetworkID__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_MockClient_ExpectNonceAt__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_MockClient_ExpectPeerCount__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_MockClient_ExpectPendingBalanceAt__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_MockClient_ExpectPendingCallContract__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_MockClient_ExpectPendingCodeAt__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_MockClient_ExpectPendingNonceAt__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_MockClient_ExpectPendingStorageAt__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_MockClient_ExpectPendingTransactionCount__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_MockClient_ExpectRPC__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/client.RPC

// skipping Fuzz_Nosy_MockClient_ExpectSendTransaction__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_MockClient_ExpectStorageAt__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_MockClient_ExpectSubscribeFilterLogs__ because parameters include func, chan, or unsupported interface: chan<- github.com/ethereum/go-ethereum/core/types.Log

// skipping Fuzz_Nosy_MockClient_ExpectSubscribeNewHead__ because parameters include func, chan, or unsupported interface: chan<- *github.com/ethereum/go-ethereum/core/types.Header

// skipping Fuzz_Nosy_MockClient_ExpectSuggestGasPrice__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_MockClient_ExpectSuggestGasTipCap__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_MockClient_ExpectSyncProgress__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_MockClient_ExpectTransactionByHash__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_MockClient_ExpectTransactionCount__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_MockClient_ExpectTransactionInBlock__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_MockClient_ExpectTransactionReceipt__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_MockClient_ExpectTransactionSender__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_MockClient_FilterLogs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockClient
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var q ethereum.FilterQuery
		fill_err = tp.Fill(&q)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.FilterLogs(ctx, q)
	})
}

func Fuzz_Nosy_MockClient_HeaderByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockClient
		fill_err = tp.Fill(&m)
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
		if m == nil {
			return
		}

		m.HeaderByHash(ctx, hash)
	})
}

func Fuzz_Nosy_MockClient_HeaderByNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockClient
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var number *big.Int
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		if m == nil || number == nil {
			return
		}

		m.HeaderByNumber(ctx, number)
	})
}

func Fuzz_Nosy_MockClient_NetworkID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockClient
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

		m.NetworkID(ctx)
	})
}

func Fuzz_Nosy_MockClient_NonceAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockClient
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var account common.Address
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		var block *big.Int
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		if m == nil || block == nil {
			return
		}

		m.NonceAt(ctx, account, block)
	})
}

func Fuzz_Nosy_MockClient_PeerCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockClient
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

		m.PeerCount(ctx)
	})
}

func Fuzz_Nosy_MockClient_PendingBalanceAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockClient
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var account common.Address
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.PendingBalanceAt(ctx, account)
	})
}

func Fuzz_Nosy_MockClient_PendingCallContract__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockClient
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var msg ethereum.CallMsg
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.PendingCallContract(ctx, msg)
	})
}

func Fuzz_Nosy_MockClient_PendingCodeAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockClient
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var account common.Address
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.PendingCodeAt(ctx, account)
	})
}

func Fuzz_Nosy_MockClient_PendingNonceAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockClient
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var account common.Address
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.PendingNonceAt(ctx, account)
	})
}

func Fuzz_Nosy_MockClient_PendingStorageAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockClient
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var account common.Address
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		var key common.Hash
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.PendingStorageAt(ctx, account, key)
	})
}

func Fuzz_Nosy_MockClient_PendingTransactionCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockClient
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

		m.PendingTransactionCount(ctx)
	})
}

func Fuzz_Nosy_MockClient_RPC__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockClient
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.RPC()
	})
}

func Fuzz_Nosy_MockClient_SendTransaction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockClient
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
		if m == nil || tx == nil {
			return
		}

		m.SendTransaction(ctx, tx)
	})
}

func Fuzz_Nosy_MockClient_StorageAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockClient
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var account common.Address
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		var key common.Hash
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var block *big.Int
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		if m == nil || block == nil {
			return
		}

		m.StorageAt(ctx, account, key, block)
	})
}

// skipping Fuzz_Nosy_MockClient_SubscribeFilterLogs__ because parameters include func, chan, or unsupported interface: chan<- github.com/ethereum/go-ethereum/core/types.Log

// skipping Fuzz_Nosy_MockClient_SubscribeNewHead__ because parameters include func, chan, or unsupported interface: chan<- *github.com/ethereum/go-ethereum/core/types.Header

func Fuzz_Nosy_MockClient_SuggestGasPrice__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockClient
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

		m.SuggestGasPrice(ctx)
	})
}

func Fuzz_Nosy_MockClient_SuggestGasTipCap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockClient
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

		m.SuggestGasTipCap(ctx)
	})
}

func Fuzz_Nosy_MockClient_SyncProgress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockClient
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

		m.SyncProgress(ctx)
	})
}

func Fuzz_Nosy_MockClient_TransactionByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockClient
		fill_err = tp.Fill(&m)
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
		if m == nil {
			return
		}

		m.TransactionByHash(ctx, hash)
	})
}

func Fuzz_Nosy_MockClient_TransactionCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockClient
		fill_err = tp.Fill(&m)
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
		if m == nil {
			return
		}

		m.TransactionCount(ctx, hash)
	})
}

func Fuzz_Nosy_MockClient_TransactionInBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockClient
		fill_err = tp.Fill(&m)
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
		var index uint
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.TransactionInBlock(ctx, hash, index)
	})
}

func Fuzz_Nosy_MockClient_TransactionReceipt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockClient
		fill_err = tp.Fill(&m)
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
		if m == nil {
			return
		}

		m.TransactionReceipt(ctx, hash)
	})
}

func Fuzz_Nosy_MockClient_TransactionSender__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockClient
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
		var block common.Hash
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		var index uint
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		if m == nil || tx == nil {
			return
		}

		m.TransactionSender(ctx, tx, block, index)
	})
}

func Fuzz_Nosy_MockDebugClient_CodeByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockDebugClient
		fill_err = tp.Fill(&m)
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
		if m == nil {
			return
		}

		m.CodeByHash(ctx, hash)
	})
}

// skipping Fuzz_Nosy_MockDebugClient_ExpectCodeByHash__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_MockDebugClient_ExpectNodeByHash__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_MockDebugClient_NodeByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockDebugClient
		fill_err = tp.Fill(&m)
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
		if m == nil {
			return
		}

		m.NodeByHash(ctx, hash)
	})
}

// skipping Fuzz_Nosy_MockEmitter_AssertExpectations__ because parameters include func, chan, or unsupported interface: github.com/stretchr/testify/mock.TestingT

// skipping Fuzz_Nosy_MockEmitter_Emit__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/event.Event

// skipping Fuzz_Nosy_MockEmitter_ExpectMaybeRun__ because parameters include func, chan, or unsupported interface: func(ev github.com/ethereum-optimism/optimism/op-node/rollup/event.Event)

// skipping Fuzz_Nosy_MockEmitter_ExpectOnce__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/event.Event

// skipping Fuzz_Nosy_MockEmitter_ExpectOnceRun__ because parameters include func, chan, or unsupported interface: func(ev github.com/ethereum-optimism/optimism/op-node/rollup/event.Event)

func Fuzz_Nosy_MockEmitter_ExpectOnceType__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockEmitter
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var typ string
		fill_err = tp.Fill(&typ)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.ExpectOnceType(typ)
	})
}

// skipping Fuzz_Nosy_MockEngine_ExpectForkchoiceUpdate__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_MockEngine_ExpectGetPayload__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_MockEngine_ExpectNewPayload__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_MockEngine_ForkchoiceUpdate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockEngine
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var state *eth.ForkchoiceState
		fill_err = tp.Fill(&state)
		if fill_err != nil {
			return
		}
		var attr *eth.PayloadAttributes
		fill_err = tp.Fill(&attr)
		if fill_err != nil {
			return
		}
		if m == nil || state == nil || attr == nil {
			return
		}

		m.ForkchoiceUpdate(ctx, state, attr)
	})
}

func Fuzz_Nosy_MockEngine_GetPayload__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockEngine
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var payloadInfo eth.PayloadInfo
		fill_err = tp.Fill(&payloadInfo)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetPayload(ctx, payloadInfo)
	})
}

func Fuzz_Nosy_MockEngine_NewPayload__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockEngine
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var payload *eth.ExecutionPayload
		fill_err = tp.Fill(&payload)
		if fill_err != nil {
			return
		}
		var parentBeaconBlockRoot *common.Hash
		fill_err = tp.Fill(&parentBeaconBlockRoot)
		if fill_err != nil {
			return
		}
		if m == nil || payload == nil || parentBeaconBlockRoot == nil {
			return
		}

		m.NewPayload(ctx, payload, parentBeaconBlockRoot)
	})
}

func Fuzz_Nosy_MockEthClient_BlockByNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockEthClient
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var number *big.Int
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		if m == nil || number == nil {
			return
		}

		m.BlockByNumber(ctx, number)
	})
}

func Fuzz_Nosy_MockEthClient_Client__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockEthClient
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Client()
	})
}

func Fuzz_Nosy_MockEthClient_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockEthClient
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

// skipping Fuzz_Nosy_MockEthClient_ExpectBlockByNumber__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_MockEthClient_ExpectClose__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockEthClient
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.ExpectClose()
	})
}

// skipping Fuzz_Nosy_MockEthClient_ExpectFetchReceipts__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/eth.BlockInfo

// skipping Fuzz_Nosy_MockEthClient_ExpectGetProof__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_MockEthClient_ExpectGetStorageAt__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_MockEthClient_ExpectInfoAndTxsByHash__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/eth.BlockInfo

// skipping Fuzz_Nosy_MockEthClient_ExpectInfoAndTxsByLabel__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/eth.BlockInfo

// skipping Fuzz_Nosy_MockEthClient_ExpectInfoAndTxsByNumber__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/eth.BlockInfo

// skipping Fuzz_Nosy_MockEthClient_ExpectInfoByHash__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/eth.BlockInfo

// skipping Fuzz_Nosy_MockEthClient_ExpectInfoByLabel__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/eth.BlockInfo

// skipping Fuzz_Nosy_MockEthClient_ExpectInfoByNumber__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/eth.BlockInfo

// skipping Fuzz_Nosy_MockEthClient_ExpectPayloadByHash__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_MockEthClient_ExpectPayloadByLabel__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_MockEthClient_ExpectPayloadByNumber__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_MockEthClient_ExpectReadStorageAt__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_MockEthClient_FetchReceipts__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockEthClient
		fill_err = tp.Fill(&m)
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
		if m == nil {
			return
		}

		m.FetchReceipts(ctx, blockHash)
	})
}

func Fuzz_Nosy_MockEthClient_GetProof__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockEthClient
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		var storage []common.Hash
		fill_err = tp.Fill(&storage)
		if fill_err != nil {
			return
		}
		var blockTag string
		fill_err = tp.Fill(&blockTag)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetProof(ctx, address, storage, blockTag)
	})
}

func Fuzz_Nosy_MockEthClient_GetStorageAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockEthClient
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		var storageSlot common.Hash
		fill_err = tp.Fill(&storageSlot)
		if fill_err != nil {
			return
		}
		var blockTag string
		fill_err = tp.Fill(&blockTag)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetStorageAt(ctx, address, storageSlot, blockTag)
	})
}

func Fuzz_Nosy_MockEthClient_InfoAndTxsByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockEthClient
		fill_err = tp.Fill(&m)
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
		if m == nil {
			return
		}

		m.InfoAndTxsByHash(ctx, hash)
	})
}

func Fuzz_Nosy_MockEthClient_InfoAndTxsByLabel__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockEthClient
		fill_err = tp.Fill(&m)
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
		if m == nil {
			return
		}

		m.InfoAndTxsByLabel(ctx, label)
	})
}

func Fuzz_Nosy_MockEthClient_InfoAndTxsByNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockEthClient
		fill_err = tp.Fill(&m)
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
		if m == nil {
			return
		}

		m.InfoAndTxsByNumber(ctx, number)
	})
}

func Fuzz_Nosy_MockEthClient_InfoByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockEthClient
		fill_err = tp.Fill(&m)
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
		if m == nil {
			return
		}

		m.InfoByHash(ctx, hash)
	})
}

func Fuzz_Nosy_MockEthClient_InfoByLabel__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockEthClient
		fill_err = tp.Fill(&m)
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
		if m == nil {
			return
		}

		m.InfoByLabel(ctx, label)
	})
}

func Fuzz_Nosy_MockEthClient_InfoByNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockEthClient
		fill_err = tp.Fill(&m)
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
		if m == nil {
			return
		}

		m.InfoByNumber(ctx, number)
	})
}

func Fuzz_Nosy_MockEthClient_MaybeClose__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockEthClient
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MaybeClose()
	})
}

func Fuzz_Nosy_MockEthClient_PayloadByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockEthClient
		fill_err = tp.Fill(&m)
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
		if m == nil {
			return
		}

		m.PayloadByHash(ctx, hash)
	})
}

func Fuzz_Nosy_MockEthClient_PayloadByLabel__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockEthClient
		fill_err = tp.Fill(&m)
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
		if m == nil {
			return
		}

		m.PayloadByLabel(ctx, label)
	})
}

func Fuzz_Nosy_MockEthClient_PayloadByNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockEthClient
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var n uint64
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.PayloadByNumber(ctx, n)
	})
}

func Fuzz_Nosy_MockEthClient_ReadStorageAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockEthClient
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		var storageSlot common.Hash
		fill_err = tp.Fill(&storageSlot)
		if fill_err != nil {
			return
		}
		var blockHash common.Hash
		fill_err = tp.Fill(&blockHash)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.ReadStorageAt(ctx, address, storageSlot, blockHash)
	})
}

// skipping Fuzz_Nosy_MockInteropBackend_AssertExpectations__ because parameters include func, chan, or unsupported interface: github.com/stretchr/testify/mock.TestingT

func Fuzz_Nosy_MockInteropBackend_CrossDerivedFrom__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockInteropBackend
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var chainID types.ChainID
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var derived eth.BlockID
		fill_err = tp.Fill(&derived)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.CrossDerivedFrom(ctx, chainID, derived)
	})
}

// skipping Fuzz_Nosy_MockInteropBackend_ExpectAnyUpdateLocalUnsafe__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_MockInteropBackend_ExpectDerivedFrom__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_MockInteropBackend_ExpectFinalized__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_MockInteropBackend_ExpectSafeView__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_MockInteropBackend_ExpectUnsafeView__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_MockInteropBackend_ExpectUpdateFinalizedL1__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_MockInteropBackend_ExpectUpdateLocalSafe__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_MockInteropBackend_ExpectUpdateLocalUnsafe__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_MockInteropBackend_Finalized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockInteropBackend
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var chainID types.ChainID
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Finalized(ctx, chainID)
	})
}

// skipping Fuzz_Nosy_MockInteropBackend_OnSafeView__ because parameters include func, chan, or unsupported interface: func(request github.com/ethereum-optimism/optimism/op-supervisor/supervisor/types.ReferenceView) (result github.com/ethereum-optimism/optimism/op-supervisor/supervisor/types.ReferenceView, err error)

// skipping Fuzz_Nosy_MockInteropBackend_OnUnsafeView__ because parameters include func, chan, or unsupported interface: func(request github.com/ethereum-optimism/optimism/op-supervisor/supervisor/types.ReferenceView) (result github.com/ethereum-optimism/optimism/op-supervisor/supervisor/types.ReferenceView, err error)

func Fuzz_Nosy_MockInteropBackend_SafeView__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockInteropBackend
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var chainID types.ChainID
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var safe types.ReferenceView
		fill_err = tp.Fill(&safe)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.SafeView(ctx, chainID, safe)
	})
}

func Fuzz_Nosy_MockInteropBackend_UnsafeView__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockInteropBackend
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var chainID types.ChainID
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var unsafe types.ReferenceView
		fill_err = tp.Fill(&unsafe)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.UnsafeView(ctx, chainID, unsafe)
	})
}

func Fuzz_Nosy_MockInteropBackend_UpdateFinalizedL1__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockInteropBackend
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var chainID types.ChainID
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var finalized eth.L1BlockRef
		fill_err = tp.Fill(&finalized)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.UpdateFinalizedL1(ctx, chainID, finalized)
	})
}

func Fuzz_Nosy_MockInteropBackend_UpdateLocalSafe__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockInteropBackend
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var chainID types.ChainID
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var derivedFrom eth.L1BlockRef
		fill_err = tp.Fill(&derivedFrom)
		if fill_err != nil {
			return
		}
		var lastDerived eth.L1BlockRef
		fill_err = tp.Fill(&lastDerived)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.UpdateLocalSafe(ctx, chainID, derivedFrom, lastDerived)
	})
}

func Fuzz_Nosy_MockInteropBackend_UpdateLocalUnsafe__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockInteropBackend
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var chainID types.ChainID
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var head eth.L1BlockRef
		fill_err = tp.Fill(&head)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.UpdateLocalUnsafe(ctx, chainID, head)
	})
}

// skipping Fuzz_Nosy_MockL1Source_ExpectL1BlockRefByHash__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_MockL1Source_ExpectL1BlockRefByLabel__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_MockL1Source_ExpectL1BlockRefByNumber__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_MockL1Source_L1BlockRefByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockL1Source
		fill_err = tp.Fill(&m)
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
		if m == nil {
			return
		}

		m.L1BlockRefByHash(ctx, hash)
	})
}

func Fuzz_Nosy_MockL1Source_L1BlockRefByLabel__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockL1Source
		fill_err = tp.Fill(&m)
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
		if m == nil {
			return
		}

		m.L1BlockRefByLabel(ctx, label)
	})
}

func Fuzz_Nosy_MockL1Source_L1BlockRefByNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockL1Source
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var num uint64
		fill_err = tp.Fill(&num)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.L1BlockRefByNumber(ctx, num)
	})
}

// skipping Fuzz_Nosy_MockL2Client_ExpectL2BlockRefByHash__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_MockL2Client_ExpectL2BlockRefByLabel__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_MockL2Client_ExpectL2BlockRefByNumber__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_MockL2Client_ExpectOutputV0AtBlock__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_MockL2Client_ExpectSystemConfigByL2Hash__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_MockL2Client_L2BlockRefByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *MockL2Client
		fill_err = tp.Fill(&c)
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
		if c == nil {
			return
		}

		c.L2BlockRefByHash(ctx, hash)
	})
}

func Fuzz_Nosy_MockL2Client_L2BlockRefByLabel__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *MockL2Client
		fill_err = tp.Fill(&c)
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
		if c == nil {
			return
		}

		c.L2BlockRefByLabel(ctx, label)
	})
}

func Fuzz_Nosy_MockL2Client_L2BlockRefByNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *MockL2Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var num uint64
		fill_err = tp.Fill(&num)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.L2BlockRefByNumber(ctx, num)
	})
}

func Fuzz_Nosy_MockL2Client_OutputV0AtBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockL2Client
		fill_err = tp.Fill(&m)
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
		if m == nil {
			return
		}

		m.OutputV0AtBlock(ctx, blockHash)
	})
}

func Fuzz_Nosy_MockL2Client_SystemConfigByL2Hash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockL2Client
		fill_err = tp.Fill(&m)
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
		if m == nil {
			return
		}

		m.SystemConfigByL2Hash(ctx, hash)
	})
}

func Fuzz_Nosy_MockRPC_BatchCallContext__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockRPC
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var b []rpc.BatchElem
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.BatchCallContext(ctx, b)
	})
}

// skipping Fuzz_Nosy_MockRPC_CallContext__ because parameters include func, chan, or unsupported interface: any

func Fuzz_Nosy_MockRPC_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockRPC
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

// skipping Fuzz_Nosy_MockRPC_EthSubscribe__ because parameters include func, chan, or unsupported interface: any

// skipping Fuzz_Nosy_MockRPC_ExpectBatchCallContext__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_MockRPC_ExpectCallContext__ because parameters include func, chan, or unsupported interface: any

func Fuzz_Nosy_MockRPC_ExpectClose__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockRPC
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.ExpectClose()
	})
}

// skipping Fuzz_Nosy_MockRPC_ExpectEthSubscribe__ because parameters include func, chan, or unsupported interface: any

func Fuzz_Nosy_MockRollupClient_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockRollupClient
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

func Fuzz_Nosy_MockRollupClient_ExpectClose__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockRollupClient
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.ExpectClose()
	})
}

// skipping Fuzz_Nosy_MockRollupClient_ExpectOutputAtBlock__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_MockRollupClient_ExpectRollupConfig__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_MockRollupClient_ExpectSequencerActive__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_MockRollupClient_ExpectStartSequencer__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_MockRollupClient_ExpectSyncStatus__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_MockRollupClient_MaybeClose__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockRollupClient
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MaybeClose()
	})
}

func Fuzz_Nosy_MockRollupClient_OutputAtBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockRollupClient
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var blockNum uint64
		fill_err = tp.Fill(&blockNum)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.OutputAtBlock(ctx, blockNum)
	})
}

func Fuzz_Nosy_MockRollupClient_RollupConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockRollupClient
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

		m.RollupConfig(ctx)
	})
}

func Fuzz_Nosy_MockRollupClient_SequencerActive__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockRollupClient
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

		m.SequencerActive(ctx)
	})
}

func Fuzz_Nosy_MockRollupClient_StartSequencer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockRollupClient
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var unsafeHead common.Hash
		fill_err = tp.Fill(&unsafeHead)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.StartSequencer(ctx, unsafeHead)
	})
}

func Fuzz_Nosy_MockRollupClient_SyncStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockRollupClient
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

		m.SyncStatus(ctx)
	})
}

func Fuzz_Nosy_MockRuntimeConfig_P2PSequencerAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MockRuntimeConfig
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.P2PSequencerAddress()
	})
}

func Fuzz_Nosy_TestDerivationMetrics_CountSequencedTxsInBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *TestDerivationMetrics
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var txns int
		fill_err = tp.Fill(&txns)
		if fill_err != nil {
			return
		}
		var deposits int
		fill_err = tp.Fill(&deposits)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.CountSequencedTxsInBlock(txns, deposits)
	})
}

func Fuzz_Nosy_TestDerivationMetrics_RecordChannelInputBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *TestDerivationMetrics
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var inputCompressedBytes int
		fill_err = tp.Fill(&inputCompressedBytes)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.RecordChannelInputBytes(inputCompressedBytes)
	})
}

func Fuzz_Nosy_TestDerivationMetrics_RecordChannelTimedOut__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *TestDerivationMetrics
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.RecordChannelTimedOut()
	})
}

func Fuzz_Nosy_TestDerivationMetrics_RecordDerivedBatches__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *TestDerivationMetrics
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var batchType string
		fill_err = tp.Fill(&batchType)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.RecordDerivedBatches(batchType)
	})
}

func Fuzz_Nosy_TestDerivationMetrics_RecordFrame__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *TestDerivationMetrics
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.RecordFrame()
	})
}

func Fuzz_Nosy_TestDerivationMetrics_RecordHeadChannelOpened__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *TestDerivationMetrics
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.RecordHeadChannelOpened()
	})
}

func Fuzz_Nosy_TestDerivationMetrics_RecordL1Ref__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *TestDerivationMetrics
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var ref eth.L1BlockRef
		fill_err = tp.Fill(&ref)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.RecordL1Ref(name, ref)
	})
}

func Fuzz_Nosy_TestDerivationMetrics_RecordL1ReorgDepth__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *TestDerivationMetrics
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var d uint64
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.RecordL1ReorgDepth(d)
	})
}

func Fuzz_Nosy_TestDerivationMetrics_RecordL2Ref__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *TestDerivationMetrics
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var ref eth.L2BlockRef
		fill_err = tp.Fill(&ref)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.RecordL2Ref(name, ref)
	})
}

func Fuzz_Nosy_TestDerivationMetrics_RecordPipelineReset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *TestDerivationMetrics
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.RecordPipelineReset()
	})
}

func Fuzz_Nosy_TestDerivationMetrics_RecordSequencerBuildingDiffTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *TestDerivationMetrics
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var duration time.Duration
		fill_err = tp.Fill(&duration)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.RecordSequencerBuildingDiffTime(duration)
	})
}

func Fuzz_Nosy_TestDerivationMetrics_RecordSequencerSealingTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *TestDerivationMetrics
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var duration time.Duration
		fill_err = tp.Fill(&duration)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.RecordSequencerSealingTime(duration)
	})
}

func Fuzz_Nosy_TestDerivationMetrics_RecordUnsafePayloadsBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *TestDerivationMetrics
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var length uint64
		fill_err = tp.Fill(&length)
		if fill_err != nil {
			return
		}
		var memSize uint64
		fill_err = tp.Fill(&memSize)
		if fill_err != nil {
			return
		}
		var next eth.BlockID
		fill_err = tp.Fill(&next)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.RecordUnsafePayloadsBuffer(length, memSize, next)
	})
}

func Fuzz_Nosy_TestDerivationMetrics_SetDerivationIdle__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *TestDerivationMetrics
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var idle bool
		fill_err = tp.Fill(&idle)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.SetDerivationIdle(idle)
	})
}

func Fuzz_Nosy_TestRPCMetrics_RecordRPCClientRequest__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *TestRPCMetrics
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var method string
		fill_err = tp.Fill(&method)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.RecordRPCClientRequest(method)
	})
}

// skipping Fuzz_Nosy_TestRPCMetrics_RecordRPCClientResponse__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_TestRPCMetrics_RecordRPCServerRequest__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *TestRPCMetrics
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var method string
		fill_err = tp.Fill(&method)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.RecordRPCServerRequest(method)
	})
}

func Fuzz_Nosy_RPCErrFaker_BatchCallContext__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r RPCErrFaker
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var b []rpc.BatchElem
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		r.BatchCallContext(ctx, b)
	})
}

// skipping Fuzz_Nosy_RPCErrFaker_CallContext__ because parameters include func, chan, or unsupported interface: any

func Fuzz_Nosy_RPCErrFaker_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r RPCErrFaker
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}

		r.Close()
	})
}

// skipping Fuzz_Nosy_RPCErrFaker_EthSubscribe__ because parameters include func, chan, or unsupported interface: any

func Fuzz_Nosy_TestID_ID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id TestID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}

		id.ID()
	})
}

func Fuzz_Nosy_BigEqual__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a *big.Int
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var b *big.Int
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if a == nil || b == nil {
			return
		}

		BigEqual(a, b)
	})
}

// skipping Fuzz_Nosy_MakeBlockInfo__ because parameters include func, chan, or unsupported interface: func(l *github.com/ethereum-optimism/optimism/op-service/testutils.MockBlockInfo)

func Fuzz_Nosy_RandomBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rng *rand.Rand
		fill_err = tp.Fill(&rng)
		if fill_err != nil {
			return
		}
		var txCount uint64
		fill_err = tp.Fill(&txCount)
		if fill_err != nil {
			return
		}
		if rng == nil {
			return
		}

		RandomBlock(rng, txCount)
	})
}

func Fuzz_Nosy_RandomBlockPrependTxs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rng *rand.Rand
		fill_err = tp.Fill(&rng)
		if fill_err != nil {
			return
		}
		var txCount int
		fill_err = tp.Fill(&txCount)
		if fill_err != nil {
			return
		}
		var ptxs []*types.Transaction
		fill_err = tp.Fill(&ptxs)
		if fill_err != nil {
			return
		}
		if rng == nil {
			return
		}

		RandomBlockPrependTxs(rng, txCount, ptxs...)
	})
}

func Fuzz_Nosy_RandomBlockPrependTxsWithTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rng *rand.Rand
		fill_err = tp.Fill(&rng)
		if fill_err != nil {
			return
		}
		var txCount int
		fill_err = tp.Fill(&txCount)
		if fill_err != nil {
			return
		}
		var t3 uint64
		fill_err = tp.Fill(&t3)
		if fill_err != nil {
			return
		}
		var ptxs []*types.Transaction
		fill_err = tp.Fill(&ptxs)
		if fill_err != nil {
			return
		}
		if rng == nil {
			return
		}

		RandomBlockPrependTxsWithTime(rng, txCount, t3, ptxs...)
	})
}

func Fuzz_Nosy_RandomBool__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rng *rand.Rand
		fill_err = tp.Fill(&rng)
		if fill_err != nil {
			return
		}
		if rng == nil {
			return
		}

		RandomBool(rng)
	})
}

func Fuzz_Nosy_RandomData__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rng *rand.Rand
		fill_err = tp.Fill(&rng)
		if fill_err != nil {
			return
		}
		var size int
		fill_err = tp.Fill(&size)
		if fill_err != nil {
			return
		}
		if rng == nil {
			return
		}

		RandomData(rng, size)
	})
}

func Fuzz_Nosy_RequireBigEqual__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var exp *big.Int
		fill_err = tp.Fill(&exp)
		if fill_err != nil {
			return
		}
		var actual *big.Int
		fill_err = tp.Fill(&actual)
		if fill_err != nil {
			return
		}
		if t1 == nil || exp == nil || actual == nil {
			return
		}

		RequireBigEqual(t1, exp, actual)
	})
}

func Fuzz_Nosy_chainL1__(f *testing.F) {
	f.Fuzz(func(t *testing.T, offset uint64, ids string) {
		chainL1(offset, ids)
	})
}

func Fuzz_Nosy_chainL2__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var offset int
		fill_err = tp.Fill(&offset)
		if fill_err != nil {
			return
		}
		var l1 []eth.L1BlockRef
		fill_err = tp.Fill(&l1)
		if fill_err != nil {
			return
		}
		var ids string
		fill_err = tp.Fill(&ids)
		if fill_err != nil {
			return
		}

		chainL2(offset, l1, ids)
	})
}

// skipping Fuzz_Nosy_mustJson__ because parameters include func, chan, or unsupported interface: E
