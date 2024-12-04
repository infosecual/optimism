package derive

import (
	"bytes"
	"context"
	"io"
	"math/big"
	"testing"

	"github.com/ethereum-optimism/optimism/op-node/rollup"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
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

func Fuzz_Nosy_AfterForceIncludeSource_SourceHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dep *AfterForceIncludeSource
		fill_err = tp.Fill(&dep)
		if fill_err != nil {
			return
		}
		if dep == nil {
			return
		}

		dep.SourceHash()
	})
}

func Fuzz_Nosy_AltDADataSource_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *AltDADataSource
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Next(ctx)
	})
}

func Fuzz_Nosy_AttributesQueue_DepositsOnlyAttributes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var aq *AttributesQueue
		fill_err = tp.Fill(&aq)
		if fill_err != nil {
			return
		}
		var parent eth.BlockID
		fill_err = tp.Fill(&parent)
		if fill_err != nil {
			return
		}
		var derivedFrom eth.L1BlockRef
		fill_err = tp.Fill(&derivedFrom)
		if fill_err != nil {
			return
		}
		if aq == nil {
			return
		}

		aq.DepositsOnlyAttributes(parent, derivedFrom)
	})
}

func Fuzz_Nosy_AttributesQueue_NextAttributes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var aq *AttributesQueue
		fill_err = tp.Fill(&aq)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var parent eth.L2BlockRef
		fill_err = tp.Fill(&parent)
		if fill_err != nil {
			return
		}
		if aq == nil {
			return
		}

		aq.NextAttributes(ctx, parent)
	})
}

func Fuzz_Nosy_AttributesQueue_Origin__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var aq *AttributesQueue
		fill_err = tp.Fill(&aq)
		if fill_err != nil {
			return
		}
		if aq == nil {
			return
		}

		aq.Origin()
	})
}

func Fuzz_Nosy_AttributesQueue_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var aq *AttributesQueue
		fill_err = tp.Fill(&aq)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var _x3 eth.L1BlockRef
		fill_err = tp.Fill(&_x3)
		if fill_err != nil {
			return
		}
		var _x4 eth.SystemConfig
		fill_err = tp.Fill(&_x4)
		if fill_err != nil {
			return
		}
		if aq == nil {
			return
		}

		aq.Reset(ctx, _x3, _x4)
	})
}

func Fuzz_Nosy_AttributesQueue_createNextAttributes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var aq *AttributesQueue
		fill_err = tp.Fill(&aq)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var batch *SingularBatch
		fill_err = tp.Fill(&batch)
		if fill_err != nil {
			return
		}
		var l2SafeHead eth.L2BlockRef
		fill_err = tp.Fill(&l2SafeHead)
		if fill_err != nil {
			return
		}
		if aq == nil || batch == nil {
			return
		}

		aq.createNextAttributes(ctx, batch, l2SafeHead)
	})
}

func Fuzz_Nosy_AttributesQueue_reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var aq *AttributesQueue
		fill_err = tp.Fill(&aq)
		if fill_err != nil {
			return
		}
		if aq == nil {
			return
		}

		aq.reset()
	})
}

func Fuzz_Nosy_AttributesWithParent_IsDerived__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a *AttributesWithParent
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if a == nil {
			return
		}

		a.IsDerived()
	})
}

func Fuzz_Nosy_AttributesWithParent_WithDepositsOnly__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a *AttributesWithParent
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if a == nil {
			return
		}

		a.WithDepositsOnly()
	})
}

func Fuzz_Nosy_BaseChannelCompressor_GetCompressed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bcc *BaseChannelCompressor
		fill_err = tp.Fill(&bcc)
		if fill_err != nil {
			return
		}
		if bcc == nil {
			return
		}

		bcc.GetCompressed()
	})
}

func Fuzz_Nosy_BaseChannelCompressor_Len__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bcc *BaseChannelCompressor
		fill_err = tp.Fill(&bcc)
		if fill_err != nil {
			return
		}
		if bcc == nil {
			return
		}

		bcc.Len()
	})
}

func Fuzz_Nosy_BaseChannelCompressor_Read__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bcc *BaseChannelCompressor
		fill_err = tp.Fill(&bcc)
		if fill_err != nil {
			return
		}
		var p []byte
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if bcc == nil {
			return
		}

		bcc.Read(p)
	})
}

func Fuzz_Nosy_BatchData_DecodeRLP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BatchData
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var s *rlp.Stream
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if b == nil || s == nil {
			return
		}

		b.DecodeRLP(s)
	})
}

func Fuzz_Nosy_BatchData_EncodeRLP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BatchData
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var w io.Writer
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.EncodeRLP(w)
	})
}

func Fuzz_Nosy_BatchData_GetBatchType__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bd *BatchData
		fill_err = tp.Fill(&bd)
		if fill_err != nil {
			return
		}
		if bd == nil {
			return
		}

		bd.GetBatchType()
	})
}

func Fuzz_Nosy_BatchData_MarshalBinary__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BatchData
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.MarshalBinary()
	})
}

func Fuzz_Nosy_BatchData_UnmarshalBinary__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BatchData
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var d2 []byte
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.UnmarshalBinary(d2)
	})
}

func Fuzz_Nosy_BatchData_decodeTyped__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BatchData
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var d2 []byte
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.decodeTyped(d2)
	})
}

func Fuzz_Nosy_BatchData_encodeTyped__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BatchData
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var buf *bytes.Buffer
		fill_err = tp.Fill(&buf)
		if fill_err != nil {
			return
		}
		if b == nil || buf == nil {
			return
		}

		b.encodeTyped(buf)
	})
}

func Fuzz_Nosy_BatchMux_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BatchMux
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var base eth.L1BlockRef
		fill_err = tp.Fill(&base)
		if fill_err != nil {
			return
		}
		var sysCfg eth.SystemConfig
		fill_err = tp.Fill(&sysCfg)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.Reset(ctx, base, sysCfg)
	})
}

func Fuzz_Nosy_BatchMux_Transform__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BatchMux
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var f2 rollup.ForkName
		fill_err = tp.Fill(&f2)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.Transform(f2)
	})
}

func Fuzz_Nosy_BatchMux_TransformHolocene__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BatchMux
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.TransformHolocene()
	})
}

// skipping Fuzz_Nosy_BatchQueue_AddBatch__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.Batch

func Fuzz_Nosy_BatchQueue_FlushChannel__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bq *BatchQueue
		fill_err = tp.Fill(&bq)
		if fill_err != nil {
			return
		}
		if bq == nil {
			return
		}

		bq.FlushChannel()
	})
}

func Fuzz_Nosy_BatchQueue_NextBatch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bq *BatchQueue
		fill_err = tp.Fill(&bq)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var parent eth.L2BlockRef
		fill_err = tp.Fill(&parent)
		if fill_err != nil {
			return
		}
		if bq == nil {
			return
		}

		bq.NextBatch(ctx, parent)
	})
}

func Fuzz_Nosy_BatchQueue_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bq *BatchQueue
		fill_err = tp.Fill(&bq)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var base eth.L1BlockRef
		fill_err = tp.Fill(&base)
		if fill_err != nil {
			return
		}
		var _x4 eth.SystemConfig
		fill_err = tp.Fill(&_x4)
		if fill_err != nil {
			return
		}
		if bq == nil {
			return
		}

		bq.Reset(_x2, base, _x4)
	})
}

func Fuzz_Nosy_BatchQueue_deriveNextBatch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bq *BatchQueue
		fill_err = tp.Fill(&bq)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var outOfData bool
		fill_err = tp.Fill(&outOfData)
		if fill_err != nil {
			return
		}
		var parent eth.L2BlockRef
		fill_err = tp.Fill(&parent)
		if fill_err != nil {
			return
		}
		if bq == nil {
			return
		}

		bq.deriveNextBatch(ctx, outOfData, parent)
	})
}

func Fuzz_Nosy_BatchStage_FlushChannel__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bs *BatchStage
		fill_err = tp.Fill(&bs)
		if fill_err != nil {
			return
		}
		if bs == nil {
			return
		}

		bs.FlushChannel()
	})
}

func Fuzz_Nosy_BatchStage_NextBatch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bs *BatchStage
		fill_err = tp.Fill(&bs)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var parent eth.L2BlockRef
		fill_err = tp.Fill(&parent)
		if fill_err != nil {
			return
		}
		if bs == nil {
			return
		}

		bs.NextBatch(ctx, parent)
	})
}

func Fuzz_Nosy_BatchStage_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bs *BatchStage
		fill_err = tp.Fill(&bs)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var base eth.L1BlockRef
		fill_err = tp.Fill(&base)
		if fill_err != nil {
			return
		}
		var _x4 eth.SystemConfig
		fill_err = tp.Fill(&_x4)
		if fill_err != nil {
			return
		}
		if bs == nil {
			return
		}

		bs.Reset(_x2, base, _x4)
	})
}

func Fuzz_Nosy_BatchStage_nextSingularBatchCandidate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bs *BatchStage
		fill_err = tp.Fill(&bs)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var parent eth.L2BlockRef
		fill_err = tp.Fill(&parent)
		if fill_err != nil {
			return
		}
		if bs == nil {
			return
		}

		bs.nextSingularBatchCandidate(ctx, parent)
	})
}

func Fuzz_Nosy_BlobDataSource_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ds *BlobDataSource
		fill_err = tp.Fill(&ds)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if ds == nil {
			return
		}

		ds.Next(ctx)
	})
}

func Fuzz_Nosy_BlobDataSource_open__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ds *BlobDataSource
		fill_err = tp.Fill(&ds)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if ds == nil {
			return
		}

		ds.open(ctx)
	})
}

func Fuzz_Nosy_BrotliCompressor_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BrotliCompressor
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		if bc == nil {
			return
		}

		bc.Reset()
	})
}

func Fuzz_Nosy_CalldataSource_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ds *CalldataSource
		fill_err = tp.Fill(&ds)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if ds == nil {
			return
		}

		ds.Next(ctx)
	})
}

func Fuzz_Nosy_Channel_AddFrame__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id ChannelID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var openBlock eth.L1BlockRef
		fill_err = tp.Fill(&openBlock)
		if fill_err != nil {
			return
		}
		var requireInOrder bool
		fill_err = tp.Fill(&requireInOrder)
		if fill_err != nil {
			return
		}
		var frame Frame
		fill_err = tp.Fill(&frame)
		if fill_err != nil {
			return
		}
		var l1InclusionBlock eth.L1BlockRef
		fill_err = tp.Fill(&l1InclusionBlock)
		if fill_err != nil {
			return
		}

		ch := NewChannel(id, openBlock, requireInOrder)
		ch.AddFrame(frame, l1InclusionBlock)
	})
}

func Fuzz_Nosy_Channel_HighestBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id ChannelID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var openBlock eth.L1BlockRef
		fill_err = tp.Fill(&openBlock)
		if fill_err != nil {
			return
		}
		var requireInOrder bool
		fill_err = tp.Fill(&requireInOrder)
		if fill_err != nil {
			return
		}

		ch := NewChannel(id, openBlock, requireInOrder)
		ch.HighestBlock()
	})
}

func Fuzz_Nosy_Channel_ID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id ChannelID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var openBlock eth.L1BlockRef
		fill_err = tp.Fill(&openBlock)
		if fill_err != nil {
			return
		}
		var requireInOrder bool
		fill_err = tp.Fill(&requireInOrder)
		if fill_err != nil {
			return
		}

		ch := NewChannel(id, openBlock, requireInOrder)
		ch.ID()
	})
}

func Fuzz_Nosy_Channel_IsReady__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id ChannelID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var openBlock eth.L1BlockRef
		fill_err = tp.Fill(&openBlock)
		if fill_err != nil {
			return
		}
		var requireInOrder bool
		fill_err = tp.Fill(&requireInOrder)
		if fill_err != nil {
			return
		}

		ch := NewChannel(id, openBlock, requireInOrder)
		ch.IsReady()
	})
}

func Fuzz_Nosy_Channel_OpenBlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id ChannelID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var openBlock eth.L1BlockRef
		fill_err = tp.Fill(&openBlock)
		if fill_err != nil {
			return
		}
		var requireInOrder bool
		fill_err = tp.Fill(&requireInOrder)
		if fill_err != nil {
			return
		}

		ch := NewChannel(id, openBlock, requireInOrder)
		ch.OpenBlockNumber()
	})
}

func Fuzz_Nosy_Channel_Reader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id ChannelID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var openBlock eth.L1BlockRef
		fill_err = tp.Fill(&openBlock)
		if fill_err != nil {
			return
		}
		var requireInOrder bool
		fill_err = tp.Fill(&requireInOrder)
		if fill_err != nil {
			return
		}

		ch := NewChannel(id, openBlock, requireInOrder)
		ch.Reader()
	})
}

func Fuzz_Nosy_Channel_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id ChannelID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var openBlock eth.L1BlockRef
		fill_err = tp.Fill(&openBlock)
		if fill_err != nil {
			return
		}
		var requireInOrder bool
		fill_err = tp.Fill(&requireInOrder)
		if fill_err != nil {
			return
		}

		ch := NewChannel(id, openBlock, requireInOrder)
		ch.Size()
	})
}

func Fuzz_Nosy_ChannelAssembler_FlushChannel__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ca *ChannelAssembler
		fill_err = tp.Fill(&ca)
		if fill_err != nil {
			return
		}
		if ca == nil {
			return
		}

		ca.FlushChannel()
	})
}

func Fuzz_Nosy_ChannelAssembler_Log__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ca *ChannelAssembler
		fill_err = tp.Fill(&ca)
		if fill_err != nil {
			return
		}
		if ca == nil {
			return
		}

		ca.Log()
	})
}

func Fuzz_Nosy_ChannelAssembler_NextRawChannel__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ca *ChannelAssembler
		fill_err = tp.Fill(&ca)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if ca == nil {
			return
		}

		ca.NextRawChannel(ctx)
	})
}

func Fuzz_Nosy_ChannelAssembler_Origin__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ca *ChannelAssembler
		fill_err = tp.Fill(&ca)
		if fill_err != nil {
			return
		}
		if ca == nil {
			return
		}

		ca.Origin()
	})
}

func Fuzz_Nosy_ChannelAssembler_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ca *ChannelAssembler
		fill_err = tp.Fill(&ca)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var _x3 eth.L1BlockRef
		fill_err = tp.Fill(&_x3)
		if fill_err != nil {
			return
		}
		var _x4 eth.SystemConfig
		fill_err = tp.Fill(&_x4)
		if fill_err != nil {
			return
		}
		if ca == nil {
			return
		}

		ca.Reset(_x2, _x3, _x4)
	})
}

func Fuzz_Nosy_ChannelAssembler_channelTimedOut__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ca *ChannelAssembler
		fill_err = tp.Fill(&ca)
		if fill_err != nil {
			return
		}
		if ca == nil {
			return
		}

		ca.channelTimedOut()
	})
}

func Fuzz_Nosy_ChannelAssembler_resetChannel__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ca *ChannelAssembler
		fill_err = tp.Fill(&ca)
		if fill_err != nil {
			return
		}
		if ca == nil {
			return
		}

		ca.resetChannel()
	})
}

func Fuzz_Nosy_ChannelBank_FlushChannel__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bq *ChannelBank
		fill_err = tp.Fill(&bq)
		if fill_err != nil {
			return
		}
		if bq == nil {
			return
		}

		bq.FlushChannel()
	})
}

func Fuzz_Nosy_ChannelBank_IngestFrame__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cb *ChannelBank
		fill_err = tp.Fill(&cb)
		if fill_err != nil {
			return
		}
		var f2 Frame
		fill_err = tp.Fill(&f2)
		if fill_err != nil {
			return
		}
		if cb == nil {
			return
		}

		cb.IngestFrame(f2)
	})
}

func Fuzz_Nosy_ChannelBank_NextRawChannel__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cb *ChannelBank
		fill_err = tp.Fill(&cb)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if cb == nil {
			return
		}

		cb.NextRawChannel(ctx)
	})
}

func Fuzz_Nosy_ChannelBank_Origin__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cb *ChannelBank
		fill_err = tp.Fill(&cb)
		if fill_err != nil {
			return
		}
		if cb == nil {
			return
		}

		cb.Origin()
	})
}

func Fuzz_Nosy_ChannelBank_Read__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cb *ChannelBank
		fill_err = tp.Fill(&cb)
		if fill_err != nil {
			return
		}
		if cb == nil {
			return
		}

		cb.Read()
	})
}

func Fuzz_Nosy_ChannelBank_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cb *ChannelBank
		fill_err = tp.Fill(&cb)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var base eth.L1BlockRef
		fill_err = tp.Fill(&base)
		if fill_err != nil {
			return
		}
		var _x4 eth.SystemConfig
		fill_err = tp.Fill(&_x4)
		if fill_err != nil {
			return
		}
		if cb == nil {
			return
		}

		cb.Reset(ctx, base, _x4)
	})
}

func Fuzz_Nosy_ChannelBank_prune__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cb *ChannelBank
		fill_err = tp.Fill(&cb)
		if fill_err != nil {
			return
		}
		if cb == nil {
			return
		}

		cb.prune()
	})
}

func Fuzz_Nosy_ChannelBank_tryReadChannelAtIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cb *ChannelBank
		fill_err = tp.Fill(&cb)
		if fill_err != nil {
			return
		}
		var i int
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
		if cb == nil {
			return
		}

		cb.tryReadChannelAtIndex(i)
	})
}

func Fuzz_Nosy_ChannelID_UnmarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id *ChannelID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var text []byte
		fill_err = tp.Fill(&text)
		if fill_err != nil {
			return
		}
		if id == nil {
			return
		}

		id.UnmarshalText(text)
	})
}

func Fuzz_Nosy_ChannelInReader_FlushChannel__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cr *ChannelInReader
		fill_err = tp.Fill(&cr)
		if fill_err != nil {
			return
		}
		if cr == nil {
			return
		}

		cr.FlushChannel()
	})
}

func Fuzz_Nosy_ChannelInReader_NextBatch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cr *ChannelInReader
		fill_err = tp.Fill(&cr)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if cr == nil {
			return
		}

		cr.NextBatch(ctx)
	})
}

func Fuzz_Nosy_ChannelInReader_NextChannel__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cr *ChannelInReader
		fill_err = tp.Fill(&cr)
		if fill_err != nil {
			return
		}
		if cr == nil {
			return
		}

		cr.NextChannel()
	})
}

func Fuzz_Nosy_ChannelInReader_Origin__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cr *ChannelInReader
		fill_err = tp.Fill(&cr)
		if fill_err != nil {
			return
		}
		if cr == nil {
			return
		}

		cr.Origin()
	})
}

func Fuzz_Nosy_ChannelInReader_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cr *ChannelInReader
		fill_err = tp.Fill(&cr)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var _x3 eth.L1BlockRef
		fill_err = tp.Fill(&_x3)
		if fill_err != nil {
			return
		}
		var _x4 eth.SystemConfig
		fill_err = tp.Fill(&_x4)
		if fill_err != nil {
			return
		}
		if cr == nil {
			return
		}

		cr.Reset(ctx, _x3, _x4)
	})
}

func Fuzz_Nosy_ChannelInReader_WriteChannel__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cr *ChannelInReader
		fill_err = tp.Fill(&cr)
		if fill_err != nil {
			return
		}
		var d2 []byte
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}
		if cr == nil {
			return
		}

		cr.WriteChannel(d2)
	})
}

func Fuzz_Nosy_ChannelMux_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *ChannelMux
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var base eth.L1BlockRef
		fill_err = tp.Fill(&base)
		if fill_err != nil {
			return
		}
		var sysCfg eth.SystemConfig
		fill_err = tp.Fill(&sysCfg)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.Reset(ctx, base, sysCfg)
	})
}

func Fuzz_Nosy_ChannelMux_Transform__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *ChannelMux
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var f2 rollup.ForkName
		fill_err = tp.Fill(&f2)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.Transform(f2)
	})
}

func Fuzz_Nosy_ChannelMux_TransformHolocene__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *ChannelMux
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.TransformHolocene()
	})
}

func Fuzz_Nosy_CompressionAlgo_Clone__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var algo *CompressionAlgo
		fill_err = tp.Fill(&algo)
		if fill_err != nil {
			return
		}
		if algo == nil {
			return
		}

		algo.Clone()
	})
}

func Fuzz_Nosy_CompressionAlgo_IsBrotli__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var algo *CompressionAlgo
		fill_err = tp.Fill(&algo)
		if fill_err != nil {
			return
		}
		if algo == nil {
			return
		}

		algo.IsBrotli()
	})
}

func Fuzz_Nosy_CompressionAlgo_Set__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var algo *CompressionAlgo
		fill_err = tp.Fill(&algo)
		if fill_err != nil {
			return
		}
		var value string
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
		if algo == nil {
			return
		}

		algo.Set(value)
	})
}

func Fuzz_Nosy_DataSourceFactory_OpenData__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ds *DataSourceFactory
		fill_err = tp.Fill(&ds)
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
		var batcherAddr common.Address
		fill_err = tp.Fill(&batcherAddr)
		if fill_err != nil {
			return
		}
		if ds == nil {
			return
		}

		ds.OpenData(ctx, ref, batcherAddr)
	})
}

func Fuzz_Nosy_DerivationPipeline_ConfirmEngineReset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dp *DerivationPipeline
		fill_err = tp.Fill(&dp)
		if fill_err != nil {
			return
		}
		if dp == nil {
			return
		}

		dp.ConfirmEngineReset()
	})
}

func Fuzz_Nosy_DerivationPipeline_DepositsOnlyAttributes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dp *DerivationPipeline
		fill_err = tp.Fill(&dp)
		if fill_err != nil {
			return
		}
		var parent eth.BlockID
		fill_err = tp.Fill(&parent)
		if fill_err != nil {
			return
		}
		var derivedFrom eth.L1BlockRef
		fill_err = tp.Fill(&derivedFrom)
		if fill_err != nil {
			return
		}
		if dp == nil {
			return
		}

		dp.DepositsOnlyAttributes(parent, derivedFrom)
	})
}

func Fuzz_Nosy_DerivationPipeline_DerivationReady__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dp *DerivationPipeline
		fill_err = tp.Fill(&dp)
		if fill_err != nil {
			return
		}
		if dp == nil {
			return
		}

		dp.DerivationReady()
	})
}

func Fuzz_Nosy_DerivationPipeline_Origin__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dp *DerivationPipeline
		fill_err = tp.Fill(&dp)
		if fill_err != nil {
			return
		}
		if dp == nil {
			return
		}

		dp.Origin()
	})
}

func Fuzz_Nosy_DerivationPipeline_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dp *DerivationPipeline
		fill_err = tp.Fill(&dp)
		if fill_err != nil {
			return
		}
		if dp == nil {
			return
		}

		dp.Reset()
	})
}

func Fuzz_Nosy_DerivationPipeline_Step__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dp *DerivationPipeline
		fill_err = tp.Fill(&dp)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var pendingSafeHead eth.L2BlockRef
		fill_err = tp.Fill(&pendingSafeHead)
		if fill_err != nil {
			return
		}
		if dp == nil {
			return
		}

		dp.Step(ctx, pendingSafeHead)
	})
}

func Fuzz_Nosy_DerivationPipeline_initialReset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dp *DerivationPipeline
		fill_err = tp.Fill(&dp)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var resetL2Safe eth.L2BlockRef
		fill_err = tp.Fill(&resetL2Safe)
		if fill_err != nil {
			return
		}
		if dp == nil {
			return
		}

		dp.initialReset(ctx, resetL2Safe)
	})
}

func Fuzz_Nosy_DerivationPipeline_transformStages__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *DerivationPipeline
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var oldOrigin eth.L1BlockRef
		fill_err = tp.Fill(&oldOrigin)
		if fill_err != nil {
			return
		}
		var newOrigin eth.L1BlockRef
		fill_err = tp.Fill(&newOrigin)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		db.transformStages(oldOrigin, newOrigin)
	})
}

func Fuzz_Nosy_FetchingAttributesBuilder_PreparePayloadAttributes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ba *FetchingAttributesBuilder
		fill_err = tp.Fill(&ba)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var l2Parent eth.L2BlockRef
		fill_err = tp.Fill(&l2Parent)
		if fill_err != nil {
			return
		}
		var epoch eth.BlockID
		fill_err = tp.Fill(&epoch)
		if fill_err != nil {
			return
		}
		if ba == nil {
			return
		}

		ba.PreparePayloadAttributes(ctx, l2Parent, epoch)
	})
}

func Fuzz_Nosy_FetchingAttributesBuilder_TestSkipL1OriginCheck__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ba *FetchingAttributesBuilder
		fill_err = tp.Fill(&ba)
		if fill_err != nil {
			return
		}
		if ba == nil {
			return
		}

		ba.TestSkipL1OriginCheck()
	})
}

func Fuzz_Nosy_Frame_MarshalBinary__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id ChannelID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var frame uint64
		fill_err = tp.Fill(&frame)
		if fill_err != nil {
			return
		}
		var readyBytes int
		fill_err = tp.Fill(&readyBytes)
		if fill_err != nil {
			return
		}
		var closed bool
		fill_err = tp.Fill(&closed)
		if fill_err != nil {
			return
		}
		var maxSize uint64
		fill_err = tp.Fill(&maxSize)
		if fill_err != nil {
			return
		}
		var w io.Writer
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}

		f1 := createEmptyFrame(id, frame, readyBytes, closed, maxSize)
		f1.MarshalBinary(w)
	})
}

// skipping Fuzz_Nosy_Frame_UnmarshalBinary__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.ByteReader

func Fuzz_Nosy_FrameQueue_NextFrame__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fq *FrameQueue
		fill_err = tp.Fill(&fq)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if fq == nil {
			return
		}

		fq.NextFrame(ctx)
	})
}

func Fuzz_Nosy_FrameQueue_Origin__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fq *FrameQueue
		fill_err = tp.Fill(&fq)
		if fill_err != nil {
			return
		}
		if fq == nil {
			return
		}

		fq.Origin()
	})
}

func Fuzz_Nosy_FrameQueue_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fq *FrameQueue
		fill_err = tp.Fill(&fq)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var _x3 eth.L1BlockRef
		fill_err = tp.Fill(&_x3)
		if fill_err != nil {
			return
		}
		var _x4 eth.SystemConfig
		fill_err = tp.Fill(&_x4)
		if fill_err != nil {
			return
		}
		if fq == nil {
			return
		}

		fq.Reset(_x2, _x3, _x4)
	})
}

func Fuzz_Nosy_FrameQueue_Transform__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fq *FrameQueue
		fill_err = tp.Fill(&fq)
		if fill_err != nil {
			return
		}
		var f2 rollup.ForkName
		fill_err = tp.Fill(&f2)
		if fill_err != nil {
			return
		}
		if fq == nil {
			return
		}

		fq.Transform(f2)
	})
}

func Fuzz_Nosy_FrameQueue_loadNextFrames__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fq *FrameQueue
		fill_err = tp.Fill(&fq)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if fq == nil {
			return
		}

		fq.loadNextFrames(ctx)
	})
}

func Fuzz_Nosy_FrameQueue_prune__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fq *FrameQueue
		fill_err = tp.Fill(&fq)
		if fill_err != nil {
			return
		}
		if fq == nil {
			return
		}

		fq.prune()
	})
}

func Fuzz_Nosy_FrameQueue_reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fq *FrameQueue
		fill_err = tp.Fill(&fq)
		if fill_err != nil {
			return
		}
		if fq == nil {
			return
		}

		fq.reset()
	})
}

func Fuzz_Nosy_L1BlockInfo_marshalBinaryBedrock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rollupCfg *rollup.Config
		fill_err = tp.Fill(&rollupCfg)
		if fill_err != nil {
			return
		}
		var l2BlockTime uint64
		fill_err = tp.Fill(&l2BlockTime)
		if fill_err != nil {
			return
		}
		var d3 []byte
		fill_err = tp.Fill(&d3)
		if fill_err != nil {
			return
		}
		if rollupCfg == nil {
			return
		}

		info, err := L1BlockInfoFromBytes(rollupCfg, l2BlockTime, d3)
		if err != nil {
			return
		}
		info.marshalBinaryBedrock()
	})
}

func Fuzz_Nosy_L1BlockInfo_marshalBinaryEcotone__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rollupCfg *rollup.Config
		fill_err = tp.Fill(&rollupCfg)
		if fill_err != nil {
			return
		}
		var l2BlockTime uint64
		fill_err = tp.Fill(&l2BlockTime)
		if fill_err != nil {
			return
		}
		var d3 []byte
		fill_err = tp.Fill(&d3)
		if fill_err != nil {
			return
		}
		if rollupCfg == nil {
			return
		}

		info, err := L1BlockInfoFromBytes(rollupCfg, l2BlockTime, d3)
		if err != nil {
			return
		}
		info.marshalBinaryEcotone()
	})
}

func Fuzz_Nosy_L1BlockInfo_marshalBinaryInterop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rollupCfg *rollup.Config
		fill_err = tp.Fill(&rollupCfg)
		if fill_err != nil {
			return
		}
		var l2BlockTime uint64
		fill_err = tp.Fill(&l2BlockTime)
		if fill_err != nil {
			return
		}
		var d3 []byte
		fill_err = tp.Fill(&d3)
		if fill_err != nil {
			return
		}
		if rollupCfg == nil {
			return
		}

		info, err := L1BlockInfoFromBytes(rollupCfg, l2BlockTime, d3)
		if err != nil {
			return
		}
		info.marshalBinaryInterop()
	})
}

func Fuzz_Nosy_L1BlockInfo_unmarshalBinaryBedrock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rollupCfg *rollup.Config
		fill_err = tp.Fill(&rollupCfg)
		if fill_err != nil {
			return
		}
		var l2BlockTime uint64
		fill_err = tp.Fill(&l2BlockTime)
		if fill_err != nil {
			return
		}
		var d3 []byte
		fill_err = tp.Fill(&d3)
		if fill_err != nil {
			return
		}
		var d4 []byte
		fill_err = tp.Fill(&d4)
		if fill_err != nil {
			return
		}
		if rollupCfg == nil {
			return
		}

		info, err := L1BlockInfoFromBytes(rollupCfg, l2BlockTime, d3)
		if err != nil {
			return
		}
		info.unmarshalBinaryBedrock(d4)
	})
}

func Fuzz_Nosy_L1BlockInfo_unmarshalBinaryEcotone__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rollupCfg *rollup.Config
		fill_err = tp.Fill(&rollupCfg)
		if fill_err != nil {
			return
		}
		var l2BlockTime uint64
		fill_err = tp.Fill(&l2BlockTime)
		if fill_err != nil {
			return
		}
		var d3 []byte
		fill_err = tp.Fill(&d3)
		if fill_err != nil {
			return
		}
		var d4 []byte
		fill_err = tp.Fill(&d4)
		if fill_err != nil {
			return
		}
		if rollupCfg == nil {
			return
		}

		info, err := L1BlockInfoFromBytes(rollupCfg, l2BlockTime, d3)
		if err != nil {
			return
		}
		info.unmarshalBinaryEcotone(d4)
	})
}

func Fuzz_Nosy_L1BlockInfo_unmarshalBinaryInterop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rollupCfg *rollup.Config
		fill_err = tp.Fill(&rollupCfg)
		if fill_err != nil {
			return
		}
		var l2BlockTime uint64
		fill_err = tp.Fill(&l2BlockTime)
		if fill_err != nil {
			return
		}
		var d3 []byte
		fill_err = tp.Fill(&d3)
		if fill_err != nil {
			return
		}
		var d4 []byte
		fill_err = tp.Fill(&d4)
		if fill_err != nil {
			return
		}
		if rollupCfg == nil {
			return
		}

		info, err := L1BlockInfoFromBytes(rollupCfg, l2BlockTime, d3)
		if err != nil {
			return
		}
		info.unmarshalBinaryInterop(d4)
	})
}

func Fuzz_Nosy_L1InfoDepositSource_SourceHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dep *L1InfoDepositSource
		fill_err = tp.Fill(&dep)
		if fill_err != nil {
			return
		}
		if dep == nil {
			return
		}

		dep.SourceHash()
	})
}

func Fuzz_Nosy_L1Retrieval_NextData__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l1r *L1Retrieval
		fill_err = tp.Fill(&l1r)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if l1r == nil {
			return
		}

		l1r.NextData(ctx)
	})
}

func Fuzz_Nosy_L1Retrieval_Origin__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l1r *L1Retrieval
		fill_err = tp.Fill(&l1r)
		if fill_err != nil {
			return
		}
		if l1r == nil {
			return
		}

		l1r.Origin()
	})
}

func Fuzz_Nosy_L1Retrieval_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l1r *L1Retrieval
		fill_err = tp.Fill(&l1r)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var base eth.L1BlockRef
		fill_err = tp.Fill(&base)
		if fill_err != nil {
			return
		}
		var sysCfg eth.SystemConfig
		fill_err = tp.Fill(&sysCfg)
		if fill_err != nil {
			return
		}
		if l1r == nil {
			return
		}

		l1r.Reset(ctx, base, sysCfg)
	})
}

func Fuzz_Nosy_L1Traversal_AdvanceL1Block__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l1t *L1Traversal
		fill_err = tp.Fill(&l1t)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if l1t == nil {
			return
		}

		l1t.AdvanceL1Block(ctx)
	})
}

func Fuzz_Nosy_L1Traversal_NextL1Block__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l1t *L1Traversal
		fill_err = tp.Fill(&l1t)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if l1t == nil {
			return
		}

		l1t.NextL1Block(_x2)
	})
}

func Fuzz_Nosy_L1Traversal_Origin__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l1t *L1Traversal
		fill_err = tp.Fill(&l1t)
		if fill_err != nil {
			return
		}
		if l1t == nil {
			return
		}

		l1t.Origin()
	})
}

func Fuzz_Nosy_L1Traversal_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l1t *L1Traversal
		fill_err = tp.Fill(&l1t)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var base eth.L1BlockRef
		fill_err = tp.Fill(&base)
		if fill_err != nil {
			return
		}
		var cfg eth.SystemConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if l1t == nil {
			return
		}

		l1t.Reset(ctx, base, cfg)
	})
}

func Fuzz_Nosy_L1Traversal_SystemConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l1c *L1Traversal
		fill_err = tp.Fill(&l1c)
		if fill_err != nil {
			return
		}
		if l1c == nil {
			return
		}

		l1c.SystemConfig()
	})
}

// skipping Fuzz_Nosy_PipelineDeriver_AttachEmitter__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/event.Emitter

// skipping Fuzz_Nosy_PipelineDeriver_OnEvent__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/event.Event

func Fuzz_Nosy_PipelineDeriver_emitDerivedAttributesEvent__(f *testing.F) {
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
		var pipeline *DerivationPipeline
		fill_err = tp.Fill(&pipeline)
		if fill_err != nil {
			return
		}
		var attrib *AttributesWithParent
		fill_err = tp.Fill(&attrib)
		if fill_err != nil {
			return
		}
		if pipeline == nil || attrib == nil {
			return
		}

		d := NewPipelineDeriver(ctx, pipeline)
		d.emitDerivedAttributesEvent(attrib)
	})
}

func Fuzz_Nosy_RawSpanBatch_GetBatchType__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *RawSpanBatch
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.GetBatchType()
	})
}

func Fuzz_Nosy_RawSpanBatch_ToSpanBatch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *RawSpanBatch
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var blockTime uint64
		fill_err = tp.Fill(&blockTime)
		if fill_err != nil {
			return
		}
		var genesisTimestamp uint64
		fill_err = tp.Fill(&genesisTimestamp)
		if fill_err != nil {
			return
		}
		var chainID *big.Int
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		if b == nil || chainID == nil {
			return
		}

		b.ToSpanBatch(blockTime, genesisTimestamp, chainID)
	})
}

func Fuzz_Nosy_RawSpanBatch_decode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *RawSpanBatch
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var r *bytes.Reader
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if b == nil || r == nil {
			return
		}

		b.decode(r)
	})
}

func Fuzz_Nosy_RawSpanBatch_derive__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *RawSpanBatch
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var blockTime uint64
		fill_err = tp.Fill(&blockTime)
		if fill_err != nil {
			return
		}
		var genesisTimestamp uint64
		fill_err = tp.Fill(&genesisTimestamp)
		if fill_err != nil {
			return
		}
		var chainID *big.Int
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		if b == nil || chainID == nil {
			return
		}

		b.derive(blockTime, genesisTimestamp, chainID)
	})
}

func Fuzz_Nosy_RawSpanBatch_encode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *RawSpanBatch
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var w io.Writer
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.encode(w)
	})
}

func Fuzz_Nosy_SingularBatch_AsSingularBatch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var batchData *BatchData
		fill_err = tp.Fill(&batchData)
		if fill_err != nil {
			return
		}
		if batchData == nil {
			return
		}

		b, err := GetSingularBatch(batchData)
		if err != nil {
			return
		}
		b.AsSingularBatch()
	})
}

func Fuzz_Nosy_SingularBatch_AsSpanBatch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var batchData *BatchData
		fill_err = tp.Fill(&batchData)
		if fill_err != nil {
			return
		}
		if batchData == nil {
			return
		}

		b, err := GetSingularBatch(batchData)
		if err != nil {
			return
		}
		b.AsSpanBatch()
	})
}

func Fuzz_Nosy_SingularBatch_Epoch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var batchData *BatchData
		fill_err = tp.Fill(&batchData)
		if fill_err != nil {
			return
		}
		if batchData == nil {
			return
		}

		b, err := GetSingularBatch(batchData)
		if err != nil {
			return
		}
		b.Epoch()
	})
}

func Fuzz_Nosy_SingularBatch_GetBatchType__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var batchData *BatchData
		fill_err = tp.Fill(&batchData)
		if fill_err != nil {
			return
		}
		if batchData == nil {
			return
		}

		b, err := GetSingularBatch(batchData)
		if err != nil {
			return
		}
		b.GetBatchType()
	})
}

func Fuzz_Nosy_SingularBatch_GetEpochNum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var batchData *BatchData
		fill_err = tp.Fill(&batchData)
		if fill_err != nil {
			return
		}
		if batchData == nil {
			return
		}

		b, err := GetSingularBatch(batchData)
		if err != nil {
			return
		}
		b.GetEpochNum()
	})
}

func Fuzz_Nosy_SingularBatch_GetTimestamp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var batchData *BatchData
		fill_err = tp.Fill(&batchData)
		if fill_err != nil {
			return
		}
		if batchData == nil {
			return
		}

		b, err := GetSingularBatch(batchData)
		if err != nil {
			return
		}
		b.GetTimestamp()
	})
}

// skipping Fuzz_Nosy_SingularBatch_LogContext__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/log.Logger

func Fuzz_Nosy_SingularBatch_decode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var batchData *BatchData
		fill_err = tp.Fill(&batchData)
		if fill_err != nil {
			return
		}
		var r *bytes.Reader
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if batchData == nil || r == nil {
			return
		}

		b, err := GetSingularBatch(batchData)
		if err != nil {
			return
		}
		b.decode(r)
	})
}

func Fuzz_Nosy_SingularBatch_encode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var batchData *BatchData
		fill_err = tp.Fill(&batchData)
		if fill_err != nil {
			return
		}
		var w io.Writer
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if batchData == nil {
			return
		}

		b, err := GetSingularBatch(batchData)
		if err != nil {
			return
		}
		b.encode(w)
	})
}

func Fuzz_Nosy_SingularChannelOut_AddBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var co *SingularChannelOut
		fill_err = tp.Fill(&co)
		if fill_err != nil {
			return
		}
		var rollupCfg *rollup.Config
		fill_err = tp.Fill(&rollupCfg)
		if fill_err != nil {
			return
		}
		var block *types.Block
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		if co == nil || rollupCfg == nil || block == nil {
			return
		}

		co.AddBlock(rollupCfg, block)
	})
}

func Fuzz_Nosy_SingularChannelOut_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var co *SingularChannelOut
		fill_err = tp.Fill(&co)
		if fill_err != nil {
			return
		}
		if co == nil {
			return
		}

		co.Close()
	})
}

func Fuzz_Nosy_SingularChannelOut_Flush__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var co *SingularChannelOut
		fill_err = tp.Fill(&co)
		if fill_err != nil {
			return
		}
		if co == nil {
			return
		}

		co.Flush()
	})
}

func Fuzz_Nosy_SingularChannelOut_FullErr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var co *SingularChannelOut
		fill_err = tp.Fill(&co)
		if fill_err != nil {
			return
		}
		if co == nil {
			return
		}

		co.FullErr()
	})
}

func Fuzz_Nosy_SingularChannelOut_ID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var co *SingularChannelOut
		fill_err = tp.Fill(&co)
		if fill_err != nil {
			return
		}
		if co == nil {
			return
		}

		co.ID()
	})
}

func Fuzz_Nosy_SingularChannelOut_InputBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var co *SingularChannelOut
		fill_err = tp.Fill(&co)
		if fill_err != nil {
			return
		}
		if co == nil {
			return
		}

		co.InputBytes()
	})
}

func Fuzz_Nosy_SingularChannelOut_OutputFrame__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var co *SingularChannelOut
		fill_err = tp.Fill(&co)
		if fill_err != nil {
			return
		}
		var w *bytes.Buffer
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var maxSize uint64
		fill_err = tp.Fill(&maxSize)
		if fill_err != nil {
			return
		}
		if co == nil || w == nil {
			return
		}

		co.OutputFrame(w, maxSize)
	})
}

func Fuzz_Nosy_SingularChannelOut_ReadyBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var co *SingularChannelOut
		fill_err = tp.Fill(&co)
		if fill_err != nil {
			return
		}
		if co == nil {
			return
		}

		co.ReadyBytes()
	})
}

func Fuzz_Nosy_SingularChannelOut_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var co *SingularChannelOut
		fill_err = tp.Fill(&co)
		if fill_err != nil {
			return
		}
		if co == nil {
			return
		}

		co.Reset()
	})
}

func Fuzz_Nosy_SingularChannelOut_addSingularBatch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var co *SingularChannelOut
		fill_err = tp.Fill(&co)
		if fill_err != nil {
			return
		}
		var batch *SingularBatch
		fill_err = tp.Fill(&batch)
		if fill_err != nil {
			return
		}
		var _x3 uint64
		fill_err = tp.Fill(&_x3)
		if fill_err != nil {
			return
		}
		if co == nil || batch == nil {
			return
		}

		co.addSingularBatch(batch, _x3)
	})
}

func Fuzz_Nosy_SpanBatch_AppendSingularBatch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var batchData *BatchData
		fill_err = tp.Fill(&batchData)
		if fill_err != nil {
			return
		}
		var blockTime uint64
		fill_err = tp.Fill(&blockTime)
		if fill_err != nil {
			return
		}
		var genesisTimestamp uint64
		fill_err = tp.Fill(&genesisTimestamp)
		if fill_err != nil {
			return
		}
		var chainID *big.Int
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var singularBatch *SingularBatch
		fill_err = tp.Fill(&singularBatch)
		if fill_err != nil {
			return
		}
		var seqNum uint64
		fill_err = tp.Fill(&seqNum)
		if fill_err != nil {
			return
		}
		if batchData == nil || chainID == nil || singularBatch == nil {
			return
		}

		b, err := DeriveSpanBatch(batchData, blockTime, genesisTimestamp, chainID)
		if err != nil {
			return
		}
		b.AppendSingularBatch(singularBatch, seqNum)
	})
}

func Fuzz_Nosy_SpanBatch_AsSingularBatch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var batchData *BatchData
		fill_err = tp.Fill(&batchData)
		if fill_err != nil {
			return
		}
		var blockTime uint64
		fill_err = tp.Fill(&blockTime)
		if fill_err != nil {
			return
		}
		var genesisTimestamp uint64
		fill_err = tp.Fill(&genesisTimestamp)
		if fill_err != nil {
			return
		}
		var chainID *big.Int
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		if batchData == nil || chainID == nil {
			return
		}

		b, err := DeriveSpanBatch(batchData, blockTime, genesisTimestamp, chainID)
		if err != nil {
			return
		}
		b.AsSingularBatch()
	})
}

func Fuzz_Nosy_SpanBatch_AsSpanBatch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var batchData *BatchData
		fill_err = tp.Fill(&batchData)
		if fill_err != nil {
			return
		}
		var blockTime uint64
		fill_err = tp.Fill(&blockTime)
		if fill_err != nil {
			return
		}
		var genesisTimestamp uint64
		fill_err = tp.Fill(&genesisTimestamp)
		if fill_err != nil {
			return
		}
		var chainID *big.Int
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		if batchData == nil || chainID == nil {
			return
		}

		b, err := DeriveSpanBatch(batchData, blockTime, genesisTimestamp, chainID)
		if err != nil {
			return
		}
		b.AsSpanBatch()
	})
}

func Fuzz_Nosy_SpanBatch_CheckOriginHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var batchData *BatchData
		fill_err = tp.Fill(&batchData)
		if fill_err != nil {
			return
		}
		var blockTime uint64
		fill_err = tp.Fill(&blockTime)
		if fill_err != nil {
			return
		}
		var genesisTimestamp uint64
		fill_err = tp.Fill(&genesisTimestamp)
		if fill_err != nil {
			return
		}
		var chainID *big.Int
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if batchData == nil || chainID == nil {
			return
		}

		b, err := DeriveSpanBatch(batchData, blockTime, genesisTimestamp, chainID)
		if err != nil {
			return
		}
		b.CheckOriginHash(hash)
	})
}

func Fuzz_Nosy_SpanBatch_CheckParentHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var batchData *BatchData
		fill_err = tp.Fill(&batchData)
		if fill_err != nil {
			return
		}
		var blockTime uint64
		fill_err = tp.Fill(&blockTime)
		if fill_err != nil {
			return
		}
		var genesisTimestamp uint64
		fill_err = tp.Fill(&genesisTimestamp)
		if fill_err != nil {
			return
		}
		var chainID *big.Int
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if batchData == nil || chainID == nil {
			return
		}

		b, err := DeriveSpanBatch(batchData, blockTime, genesisTimestamp, chainID)
		if err != nil {
			return
		}
		b.CheckParentHash(hash)
	})
}

func Fuzz_Nosy_SpanBatch_GetBatchType__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var batchData *BatchData
		fill_err = tp.Fill(&batchData)
		if fill_err != nil {
			return
		}
		var blockTime uint64
		fill_err = tp.Fill(&blockTime)
		if fill_err != nil {
			return
		}
		var genesisTimestamp uint64
		fill_err = tp.Fill(&genesisTimestamp)
		if fill_err != nil {
			return
		}
		var chainID *big.Int
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		if batchData == nil || chainID == nil {
			return
		}

		b, err := DeriveSpanBatch(batchData, blockTime, genesisTimestamp, chainID)
		if err != nil {
			return
		}
		b.GetBatchType()
	})
}

func Fuzz_Nosy_SpanBatch_GetBlockCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var batchData *BatchData
		fill_err = tp.Fill(&batchData)
		if fill_err != nil {
			return
		}
		var blockTime uint64
		fill_err = tp.Fill(&blockTime)
		if fill_err != nil {
			return
		}
		var genesisTimestamp uint64
		fill_err = tp.Fill(&genesisTimestamp)
		if fill_err != nil {
			return
		}
		var chainID *big.Int
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		if batchData == nil || chainID == nil {
			return
		}

		b, err := DeriveSpanBatch(batchData, blockTime, genesisTimestamp, chainID)
		if err != nil {
			return
		}
		b.GetBlockCount()
	})
}

func Fuzz_Nosy_SpanBatch_GetBlockEpochNum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var batchData *BatchData
		fill_err = tp.Fill(&batchData)
		if fill_err != nil {
			return
		}
		var blockTime uint64
		fill_err = tp.Fill(&blockTime)
		if fill_err != nil {
			return
		}
		var genesisTimestamp uint64
		fill_err = tp.Fill(&genesisTimestamp)
		if fill_err != nil {
			return
		}
		var chainID *big.Int
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var i int
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
		if batchData == nil || chainID == nil {
			return
		}

		b, err := DeriveSpanBatch(batchData, blockTime, genesisTimestamp, chainID)
		if err != nil {
			return
		}
		b.GetBlockEpochNum(i)
	})
}

func Fuzz_Nosy_SpanBatch_GetBlockTimestamp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var batchData *BatchData
		fill_err = tp.Fill(&batchData)
		if fill_err != nil {
			return
		}
		var blockTime uint64
		fill_err = tp.Fill(&blockTime)
		if fill_err != nil {
			return
		}
		var genesisTimestamp uint64
		fill_err = tp.Fill(&genesisTimestamp)
		if fill_err != nil {
			return
		}
		var chainID *big.Int
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var i int
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
		if batchData == nil || chainID == nil {
			return
		}

		b, err := DeriveSpanBatch(batchData, blockTime, genesisTimestamp, chainID)
		if err != nil {
			return
		}
		b.GetBlockTimestamp(i)
	})
}

func Fuzz_Nosy_SpanBatch_GetBlockTransactions__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var batchData *BatchData
		fill_err = tp.Fill(&batchData)
		if fill_err != nil {
			return
		}
		var blockTime uint64
		fill_err = tp.Fill(&blockTime)
		if fill_err != nil {
			return
		}
		var genesisTimestamp uint64
		fill_err = tp.Fill(&genesisTimestamp)
		if fill_err != nil {
			return
		}
		var chainID *big.Int
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var i int
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
		if batchData == nil || chainID == nil {
			return
		}

		b, err := DeriveSpanBatch(batchData, blockTime, genesisTimestamp, chainID)
		if err != nil {
			return
		}
		b.GetBlockTransactions(i)
	})
}

func Fuzz_Nosy_SpanBatch_GetSingularBatches__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var batchData *BatchData
		fill_err = tp.Fill(&batchData)
		if fill_err != nil {
			return
		}
		var blockTime uint64
		fill_err = tp.Fill(&blockTime)
		if fill_err != nil {
			return
		}
		var genesisTimestamp uint64
		fill_err = tp.Fill(&genesisTimestamp)
		if fill_err != nil {
			return
		}
		var chainID *big.Int
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var l1Origins []eth.L1BlockRef
		fill_err = tp.Fill(&l1Origins)
		if fill_err != nil {
			return
		}
		var l2SafeHead eth.L2BlockRef
		fill_err = tp.Fill(&l2SafeHead)
		if fill_err != nil {
			return
		}
		if batchData == nil || chainID == nil {
			return
		}

		b, err := DeriveSpanBatch(batchData, blockTime, genesisTimestamp, chainID)
		if err != nil {
			return
		}
		b.GetSingularBatches(l1Origins, l2SafeHead)
	})
}

func Fuzz_Nosy_SpanBatch_GetStartEpochNum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var batchData *BatchData
		fill_err = tp.Fill(&batchData)
		if fill_err != nil {
			return
		}
		var blockTime uint64
		fill_err = tp.Fill(&blockTime)
		if fill_err != nil {
			return
		}
		var genesisTimestamp uint64
		fill_err = tp.Fill(&genesisTimestamp)
		if fill_err != nil {
			return
		}
		var chainID *big.Int
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		if batchData == nil || chainID == nil {
			return
		}

		b, err := DeriveSpanBatch(batchData, blockTime, genesisTimestamp, chainID)
		if err != nil {
			return
		}
		b.GetStartEpochNum()
	})
}

func Fuzz_Nosy_SpanBatch_GetTimestamp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var batchData *BatchData
		fill_err = tp.Fill(&batchData)
		if fill_err != nil {
			return
		}
		var blockTime uint64
		fill_err = tp.Fill(&blockTime)
		if fill_err != nil {
			return
		}
		var genesisTimestamp uint64
		fill_err = tp.Fill(&genesisTimestamp)
		if fill_err != nil {
			return
		}
		var chainID *big.Int
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		if batchData == nil || chainID == nil {
			return
		}

		b, err := DeriveSpanBatch(batchData, blockTime, genesisTimestamp, chainID)
		if err != nil {
			return
		}
		b.GetTimestamp()
	})
}

// skipping Fuzz_Nosy_SpanBatch_LogContext__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/log.Logger

func Fuzz_Nosy_SpanBatch_MarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var batchData *BatchData
		fill_err = tp.Fill(&batchData)
		if fill_err != nil {
			return
		}
		var blockTime uint64
		fill_err = tp.Fill(&blockTime)
		if fill_err != nil {
			return
		}
		var genesisTimestamp uint64
		fill_err = tp.Fill(&genesisTimestamp)
		if fill_err != nil {
			return
		}
		var chainID *big.Int
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		if batchData == nil || chainID == nil {
			return
		}

		b, err := DeriveSpanBatch(batchData, blockTime, genesisTimestamp, chainID)
		if err != nil {
			return
		}
		b.MarshalJSON()
	})
}

func Fuzz_Nosy_SpanBatch_ToRawSpanBatch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var batchData *BatchData
		fill_err = tp.Fill(&batchData)
		if fill_err != nil {
			return
		}
		var blockTime uint64
		fill_err = tp.Fill(&blockTime)
		if fill_err != nil {
			return
		}
		var genesisTimestamp uint64
		fill_err = tp.Fill(&genesisTimestamp)
		if fill_err != nil {
			return
		}
		var chainID *big.Int
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		if batchData == nil || chainID == nil {
			return
		}

		b, err := DeriveSpanBatch(batchData, blockTime, genesisTimestamp, chainID)
		if err != nil {
			return
		}
		b.ToRawSpanBatch()
	})
}

func Fuzz_Nosy_SpanBatch_TxCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var batchData *BatchData
		fill_err = tp.Fill(&batchData)
		if fill_err != nil {
			return
		}
		var blockTime uint64
		fill_err = tp.Fill(&blockTime)
		if fill_err != nil {
			return
		}
		var genesisTimestamp uint64
		fill_err = tp.Fill(&genesisTimestamp)
		if fill_err != nil {
			return
		}
		var chainID *big.Int
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		if batchData == nil || chainID == nil {
			return
		}

		b, err := DeriveSpanBatch(batchData, blockTime, genesisTimestamp, chainID)
		if err != nil {
			return
		}
		b.TxCount()
	})
}

func Fuzz_Nosy_SpanBatch_peek__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var batchData *BatchData
		fill_err = tp.Fill(&batchData)
		if fill_err != nil {
			return
		}
		var blockTime uint64
		fill_err = tp.Fill(&blockTime)
		if fill_err != nil {
			return
		}
		var genesisTimestamp uint64
		fill_err = tp.Fill(&genesisTimestamp)
		if fill_err != nil {
			return
		}
		var chainID *big.Int
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var n int
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if batchData == nil || chainID == nil {
			return
		}

		b, err := DeriveSpanBatch(batchData, blockTime, genesisTimestamp, chainID)
		if err != nil {
			return
		}
		b.peek(n)
	})
}

func Fuzz_Nosy_SpanChannelOut_AddBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var co *SpanChannelOut
		fill_err = tp.Fill(&co)
		if fill_err != nil {
			return
		}
		var rollupCfg *rollup.Config
		fill_err = tp.Fill(&rollupCfg)
		if fill_err != nil {
			return
		}
		var block *types.Block
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		if co == nil || rollupCfg == nil || block == nil {
			return
		}

		co.AddBlock(rollupCfg, block)
	})
}

func Fuzz_Nosy_SpanChannelOut_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var co *SpanChannelOut
		fill_err = tp.Fill(&co)
		if fill_err != nil {
			return
		}
		if co == nil {
			return
		}

		co.Close()
	})
}

func Fuzz_Nosy_SpanChannelOut_Flush__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var co *SpanChannelOut
		fill_err = tp.Fill(&co)
		if fill_err != nil {
			return
		}
		if co == nil {
			return
		}

		co.Flush()
	})
}

func Fuzz_Nosy_SpanChannelOut_FullErr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var co *SpanChannelOut
		fill_err = tp.Fill(&co)
		if fill_err != nil {
			return
		}
		if co == nil {
			return
		}

		co.FullErr()
	})
}

func Fuzz_Nosy_SpanChannelOut_ID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var co *SpanChannelOut
		fill_err = tp.Fill(&co)
		if fill_err != nil {
			return
		}
		if co == nil {
			return
		}

		co.ID()
	})
}

func Fuzz_Nosy_SpanChannelOut_InputBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var co *SpanChannelOut
		fill_err = tp.Fill(&co)
		if fill_err != nil {
			return
		}
		if co == nil {
			return
		}

		co.InputBytes()
	})
}

func Fuzz_Nosy_SpanChannelOut_OutputFrame__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var co *SpanChannelOut
		fill_err = tp.Fill(&co)
		if fill_err != nil {
			return
		}
		var w *bytes.Buffer
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var maxSize uint64
		fill_err = tp.Fill(&maxSize)
		if fill_err != nil {
			return
		}
		if co == nil || w == nil {
			return
		}

		co.OutputFrame(w, maxSize)
	})
}

func Fuzz_Nosy_SpanChannelOut_ReadyBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var co *SpanChannelOut
		fill_err = tp.Fill(&co)
		if fill_err != nil {
			return
		}
		if co == nil {
			return
		}

		co.ReadyBytes()
	})
}

func Fuzz_Nosy_SpanChannelOut_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var co *SpanChannelOut
		fill_err = tp.Fill(&co)
		if fill_err != nil {
			return
		}
		if co == nil {
			return
		}

		co.Reset()
	})
}

func Fuzz_Nosy_SpanChannelOut_activeRLP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var co *SpanChannelOut
		fill_err = tp.Fill(&co)
		if fill_err != nil {
			return
		}
		if co == nil {
			return
		}

		co.activeRLP()
	})
}

func Fuzz_Nosy_SpanChannelOut_addSingularBatch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var co *SpanChannelOut
		fill_err = tp.Fill(&co)
		if fill_err != nil {
			return
		}
		var batch *SingularBatch
		fill_err = tp.Fill(&batch)
		if fill_err != nil {
			return
		}
		var seqNum uint64
		fill_err = tp.Fill(&seqNum)
		if fill_err != nil {
			return
		}
		if co == nil || batch == nil {
			return
		}

		co.addSingularBatch(batch, seqNum)
	})
}

func Fuzz_Nosy_SpanChannelOut_checkFull__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var co *SpanChannelOut
		fill_err = tp.Fill(&co)
		if fill_err != nil {
			return
		}
		if co == nil {
			return
		}

		co.checkFull()
	})
}

func Fuzz_Nosy_SpanChannelOut_compress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var co *SpanChannelOut
		fill_err = tp.Fill(&co)
		if fill_err != nil {
			return
		}
		if co == nil {
			return
		}

		co.compress()
	})
}

func Fuzz_Nosy_SpanChannelOut_ensureOpenSpanBatch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var co *SpanChannelOut
		fill_err = tp.Fill(&co)
		if fill_err != nil {
			return
		}
		if co == nil {
			return
		}

		co.ensureOpenSpanBatch()
	})
}

func Fuzz_Nosy_SpanChannelOut_inactiveRLP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var co *SpanChannelOut
		fill_err = tp.Fill(&co)
		if fill_err != nil {
			return
		}
		if co == nil {
			return
		}

		co.inactiveRLP()
	})
}

func Fuzz_Nosy_SpanChannelOut_resetSpanBatch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var co *SpanChannelOut
		fill_err = tp.Fill(&co)
		if fill_err != nil {
			return
		}
		if co == nil {
			return
		}

		co.resetSpanBatch()
	})
}

func Fuzz_Nosy_SpanChannelOut_setRandomID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var co *SpanChannelOut
		fill_err = tp.Fill(&co)
		if fill_err != nil {
			return
		}
		if co == nil {
			return
		}

		co.setRandomID()
	})
}

func Fuzz_Nosy_SpanChannelOut_swapRLP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var co *SpanChannelOut
		fill_err = tp.Fill(&co)
		if fill_err != nil {
			return
		}
		if co == nil {
			return
		}

		co.swapRLP()
	})
}

func Fuzz_Nosy_UpgradeDepositSource_SourceHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dep *UpgradeDepositSource
		fill_err = tp.Fill(&dep)
		if fill_err != nil {
			return
		}
		if dep == nil {
			return
		}

		dep.SourceHash()
	})
}

func Fuzz_Nosy_UserDepositSource_SourceHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dep *UserDepositSource
		fill_err = tp.Fill(&dep)
		if fill_err != nil {
			return
		}
		if dep == nil {
			return
		}

		dep.SourceHash()
	})
}

func Fuzz_Nosy_ZlibCompressor_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var zc *ZlibCompressor
		fill_err = tp.Fill(&zc)
		if fill_err != nil {
			return
		}
		if zc == nil {
			return
		}

		zc.Reset()
	})
}

func Fuzz_Nosy_baseBatchStage_Log__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bs *baseBatchStage
		fill_err = tp.Fill(&bs)
		if fill_err != nil {
			return
		}
		if bs == nil {
			return
		}

		bs.Log()
	})
}

func Fuzz_Nosy_baseBatchStage_Origin__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bs *baseBatchStage
		fill_err = tp.Fill(&bs)
		if fill_err != nil {
			return
		}
		if bs == nil {
			return
		}

		bs.Origin()
	})
}

func Fuzz_Nosy_baseBatchStage_base__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bs *baseBatchStage
		fill_err = tp.Fill(&bs)
		if fill_err != nil {
			return
		}
		if bs == nil {
			return
		}

		bs.base()
	})
}

func Fuzz_Nosy_baseBatchStage_deriveNextEmptyBatch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bs *baseBatchStage
		fill_err = tp.Fill(&bs)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var outOfData bool
		fill_err = tp.Fill(&outOfData)
		if fill_err != nil {
			return
		}
		var parent eth.L2BlockRef
		fill_err = tp.Fill(&parent)
		if fill_err != nil {
			return
		}
		if bs == nil {
			return
		}

		bs.deriveNextEmptyBatch(ctx, outOfData, parent)
	})
}

func Fuzz_Nosy_baseBatchStage_nextFromSpanBatch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bs *baseBatchStage
		fill_err = tp.Fill(&bs)
		if fill_err != nil {
			return
		}
		var parent eth.L2BlockRef
		fill_err = tp.Fill(&parent)
		if fill_err != nil {
			return
		}
		if bs == nil {
			return
		}

		bs.nextFromSpanBatch(parent)
	})
}

func Fuzz_Nosy_baseBatchStage_originBehind__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bs *baseBatchStage
		fill_err = tp.Fill(&bs)
		if fill_err != nil {
			return
		}
		var parent eth.L2BlockRef
		fill_err = tp.Fill(&parent)
		if fill_err != nil {
			return
		}
		if bs == nil {
			return
		}

		bs.originBehind(parent)
	})
}

func Fuzz_Nosy_baseBatchStage_popNextBatch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bs *baseBatchStage
		fill_err = tp.Fill(&bs)
		if fill_err != nil {
			return
		}
		var parent eth.L2BlockRef
		fill_err = tp.Fill(&parent)
		if fill_err != nil {
			return
		}
		if bs == nil {
			return
		}

		bs.popNextBatch(parent)
	})
}

func Fuzz_Nosy_baseBatchStage_reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bs *baseBatchStage
		fill_err = tp.Fill(&bs)
		if fill_err != nil {
			return
		}
		var base eth.L1BlockRef
		fill_err = tp.Fill(&base)
		if fill_err != nil {
			return
		}
		if bs == nil {
			return
		}

		bs.reset(base)
	})
}

func Fuzz_Nosy_baseBatchStage_updateOrigins__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bs *baseBatchStage
		fill_err = tp.Fill(&bs)
		if fill_err != nil {
			return
		}
		var parent eth.L2BlockRef
		fill_err = tp.Fill(&parent)
		if fill_err != nil {
			return
		}
		if bs == nil {
			return
		}

		bs.updateOrigins(parent)
	})
}

func Fuzz_Nosy_spanBatchAccessListTxData_txType__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var txData *spanBatchAccessListTxData
		fill_err = tp.Fill(&txData)
		if fill_err != nil {
			return
		}
		if txData == nil {
			return
		}

		txData.txType()
	})
}

func Fuzz_Nosy_spanBatchDynamicFeeTxData_txType__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var txData *spanBatchDynamicFeeTxData
		fill_err = tp.Fill(&txData)
		if fill_err != nil {
			return
		}
		if txData == nil {
			return
		}

		txData.txType()
	})
}

func Fuzz_Nosy_spanBatchLegacyTxData_txType__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var txData *spanBatchLegacyTxData
		fill_err = tp.Fill(&txData)
		if fill_err != nil {
			return
		}
		if txData == nil {
			return
		}

		txData.txType()
	})
}

func Fuzz_Nosy_spanBatchPayload_decodeBlockCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bp *spanBatchPayload
		fill_err = tp.Fill(&bp)
		if fill_err != nil {
			return
		}
		var r *bytes.Reader
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if bp == nil || r == nil {
			return
		}

		bp.decodeBlockCount(r)
	})
}

func Fuzz_Nosy_spanBatchPayload_decodeBlockTxCounts__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bp *spanBatchPayload
		fill_err = tp.Fill(&bp)
		if fill_err != nil {
			return
		}
		var r *bytes.Reader
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if bp == nil || r == nil {
			return
		}

		bp.decodeBlockTxCounts(r)
	})
}

func Fuzz_Nosy_spanBatchPayload_decodeOriginBits__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bp *spanBatchPayload
		fill_err = tp.Fill(&bp)
		if fill_err != nil {
			return
		}
		var r *bytes.Reader
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if bp == nil || r == nil {
			return
		}

		bp.decodeOriginBits(r)
	})
}

func Fuzz_Nosy_spanBatchPayload_decodePayload__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bp *spanBatchPayload
		fill_err = tp.Fill(&bp)
		if fill_err != nil {
			return
		}
		var r *bytes.Reader
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if bp == nil || r == nil {
			return
		}

		bp.decodePayload(r)
	})
}

func Fuzz_Nosy_spanBatchPayload_decodeTxs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bp *spanBatchPayload
		fill_err = tp.Fill(&bp)
		if fill_err != nil {
			return
		}
		var r *bytes.Reader
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if bp == nil || r == nil {
			return
		}

		bp.decodeTxs(r)
	})
}

func Fuzz_Nosy_spanBatchPayload_encodeBlockCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bp *spanBatchPayload
		fill_err = tp.Fill(&bp)
		if fill_err != nil {
			return
		}
		var w io.Writer
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if bp == nil {
			return
		}

		bp.encodeBlockCount(w)
	})
}

func Fuzz_Nosy_spanBatchPayload_encodeBlockTxCounts__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bp *spanBatchPayload
		fill_err = tp.Fill(&bp)
		if fill_err != nil {
			return
		}
		var w io.Writer
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if bp == nil {
			return
		}

		bp.encodeBlockTxCounts(w)
	})
}

func Fuzz_Nosy_spanBatchPayload_encodeOriginBits__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bp *spanBatchPayload
		fill_err = tp.Fill(&bp)
		if fill_err != nil {
			return
		}
		var w io.Writer
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if bp == nil {
			return
		}

		bp.encodeOriginBits(w)
	})
}

func Fuzz_Nosy_spanBatchPayload_encodePayload__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bp *spanBatchPayload
		fill_err = tp.Fill(&bp)
		if fill_err != nil {
			return
		}
		var w io.Writer
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if bp == nil {
			return
		}

		bp.encodePayload(w)
	})
}

func Fuzz_Nosy_spanBatchPayload_encodeTxs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bp *spanBatchPayload
		fill_err = tp.Fill(&bp)
		if fill_err != nil {
			return
		}
		var w io.Writer
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if bp == nil {
			return
		}

		bp.encodeTxs(w)
	})
}

func Fuzz_Nosy_spanBatchPrefix_decodeL1OriginCheck__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bp *spanBatchPrefix
		fill_err = tp.Fill(&bp)
		if fill_err != nil {
			return
		}
		var r *bytes.Reader
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if bp == nil || r == nil {
			return
		}

		bp.decodeL1OriginCheck(r)
	})
}

func Fuzz_Nosy_spanBatchPrefix_decodeL1OriginNum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bp *spanBatchPrefix
		fill_err = tp.Fill(&bp)
		if fill_err != nil {
			return
		}
		var r *bytes.Reader
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if bp == nil || r == nil {
			return
		}

		bp.decodeL1OriginNum(r)
	})
}

func Fuzz_Nosy_spanBatchPrefix_decodeParentCheck__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bp *spanBatchPrefix
		fill_err = tp.Fill(&bp)
		if fill_err != nil {
			return
		}
		var r *bytes.Reader
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if bp == nil || r == nil {
			return
		}

		bp.decodeParentCheck(r)
	})
}

func Fuzz_Nosy_spanBatchPrefix_decodePrefix__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bp *spanBatchPrefix
		fill_err = tp.Fill(&bp)
		if fill_err != nil {
			return
		}
		var r *bytes.Reader
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if bp == nil || r == nil {
			return
		}

		bp.decodePrefix(r)
	})
}

func Fuzz_Nosy_spanBatchPrefix_decodeRelTimestamp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bp *spanBatchPrefix
		fill_err = tp.Fill(&bp)
		if fill_err != nil {
			return
		}
		var r *bytes.Reader
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if bp == nil || r == nil {
			return
		}

		bp.decodeRelTimestamp(r)
	})
}

func Fuzz_Nosy_spanBatchPrefix_encodeL1OriginCheck__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bp *spanBatchPrefix
		fill_err = tp.Fill(&bp)
		if fill_err != nil {
			return
		}
		var w io.Writer
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if bp == nil {
			return
		}

		bp.encodeL1OriginCheck(w)
	})
}

func Fuzz_Nosy_spanBatchPrefix_encodeL1OriginNum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bp *spanBatchPrefix
		fill_err = tp.Fill(&bp)
		if fill_err != nil {
			return
		}
		var w io.Writer
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if bp == nil {
			return
		}

		bp.encodeL1OriginNum(w)
	})
}

func Fuzz_Nosy_spanBatchPrefix_encodeParentCheck__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bp *spanBatchPrefix
		fill_err = tp.Fill(&bp)
		if fill_err != nil {
			return
		}
		var w io.Writer
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if bp == nil {
			return
		}

		bp.encodeParentCheck(w)
	})
}

func Fuzz_Nosy_spanBatchPrefix_encodePrefix__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bp *spanBatchPrefix
		fill_err = tp.Fill(&bp)
		if fill_err != nil {
			return
		}
		var w io.Writer
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if bp == nil {
			return
		}

		bp.encodePrefix(w)
	})
}

func Fuzz_Nosy_spanBatchPrefix_encodeRelTimestamp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bp *spanBatchPrefix
		fill_err = tp.Fill(&bp)
		if fill_err != nil {
			return
		}
		var w io.Writer
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if bp == nil {
			return
		}

		bp.encodeRelTimestamp(w)
	})
}

func Fuzz_Nosy_spanBatchTx_MarshalBinary__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *types.Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		t1, err := newSpanBatchTx(tx)
		if err != nil {
			return
		}
		t1.MarshalBinary()
	})
}

func Fuzz_Nosy_spanBatchTx_Type__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *types.Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		t1, err := newSpanBatchTx(tx)
		if err != nil {
			return
		}
		t1.Type()
	})
}

func Fuzz_Nosy_spanBatchTx_UnmarshalBinary__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *types.Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		t1, err := newSpanBatchTx(tx)
		if err != nil {
			return
		}
		t1.UnmarshalBinary(b)
	})
}

func Fuzz_Nosy_spanBatchTx_convertToFullTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *types.Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		var nonce uint64
		fill_err = tp.Fill(&nonce)
		if fill_err != nil {
			return
		}
		var gas uint64
		fill_err = tp.Fill(&gas)
		if fill_err != nil {
			return
		}
		var to *common.Address
		fill_err = tp.Fill(&to)
		if fill_err != nil {
			return
		}
		var chainID *big.Int
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var V *big.Int
		fill_err = tp.Fill(&V)
		if fill_err != nil {
			return
		}
		var R *big.Int
		fill_err = tp.Fill(&R)
		if fill_err != nil {
			return
		}
		var S *big.Int
		fill_err = tp.Fill(&S)
		if fill_err != nil {
			return
		}
		if tx == nil || to == nil || chainID == nil || V == nil || R == nil || S == nil {
			return
		}

		t1, err := newSpanBatchTx(tx)
		if err != nil {
			return
		}
		t1.convertToFullTx(nonce, gas, to, chainID, V, R, S)
	})
}

func Fuzz_Nosy_spanBatchTx_decodeTyped__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *types.Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		t1, err := newSpanBatchTx(tx)
		if err != nil {
			return
		}
		t1.decodeTyped(b)
	})
}

func Fuzz_Nosy_spanBatchTx_encodeTyped__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *types.Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		var w *bytes.Buffer
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if tx == nil || w == nil {
			return
		}

		t1, err := newSpanBatchTx(tx)
		if err != nil {
			return
		}
		t1.encodeTyped(w)
	})
}

// skipping Fuzz_Nosy_spanBatchTx_setDecoded__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.spanBatchTxData

func Fuzz_Nosy_spanBatchTxs_AddTxs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 [][]byte
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var c2 *big.Int
		fill_err = tp.Fill(&c2)
		if fill_err != nil {
			return
		}
		var t3 [][]byte
		fill_err = tp.Fill(&t3)
		if fill_err != nil {
			return
		}
		var c4 *big.Int
		fill_err = tp.Fill(&c4)
		if fill_err != nil {
			return
		}
		if c2 == nil || c4 == nil {
			return
		}

		sbtx, err := newSpanBatchTxs(t1, c2)
		if err != nil {
			return
		}
		sbtx.AddTxs(t3, c4)
	})
}

func Fuzz_Nosy_spanBatchTxs_contractCreationCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var txs [][]byte
		fill_err = tp.Fill(&txs)
		if fill_err != nil {
			return
		}
		var chainID *big.Int
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		if chainID == nil {
			return
		}

		btx, err := newSpanBatchTxs(txs, chainID)
		if err != nil {
			return
		}
		btx.contractCreationCount()
	})
}

func Fuzz_Nosy_spanBatchTxs_decode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var txs [][]byte
		fill_err = tp.Fill(&txs)
		if fill_err != nil {
			return
		}
		var chainID *big.Int
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var r *bytes.Reader
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if chainID == nil || r == nil {
			return
		}

		btx, err := newSpanBatchTxs(txs, chainID)
		if err != nil {
			return
		}
		btx.decode(r)
	})
}

func Fuzz_Nosy_spanBatchTxs_decodeContractCreationBits__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var txs [][]byte
		fill_err = tp.Fill(&txs)
		if fill_err != nil {
			return
		}
		var chainID *big.Int
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var r *bytes.Reader
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if chainID == nil || r == nil {
			return
		}

		btx, err := newSpanBatchTxs(txs, chainID)
		if err != nil {
			return
		}
		btx.decodeContractCreationBits(r)
	})
}

func Fuzz_Nosy_spanBatchTxs_decodeProtectedBits__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var txs [][]byte
		fill_err = tp.Fill(&txs)
		if fill_err != nil {
			return
		}
		var chainID *big.Int
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var r *bytes.Reader
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if chainID == nil || r == nil {
			return
		}

		btx, err := newSpanBatchTxs(txs, chainID)
		if err != nil {
			return
		}
		btx.decodeProtectedBits(r)
	})
}

func Fuzz_Nosy_spanBatchTxs_decodeTxDatas__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var txs [][]byte
		fill_err = tp.Fill(&txs)
		if fill_err != nil {
			return
		}
		var chainID *big.Int
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var r *bytes.Reader
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if chainID == nil || r == nil {
			return
		}

		btx, err := newSpanBatchTxs(txs, chainID)
		if err != nil {
			return
		}
		btx.decodeTxDatas(r)
	})
}

func Fuzz_Nosy_spanBatchTxs_decodeTxGases__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var txs [][]byte
		fill_err = tp.Fill(&txs)
		if fill_err != nil {
			return
		}
		var chainID *big.Int
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var r *bytes.Reader
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if chainID == nil || r == nil {
			return
		}

		btx, err := newSpanBatchTxs(txs, chainID)
		if err != nil {
			return
		}
		btx.decodeTxGases(r)
	})
}

func Fuzz_Nosy_spanBatchTxs_decodeTxNonces__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var txs [][]byte
		fill_err = tp.Fill(&txs)
		if fill_err != nil {
			return
		}
		var chainID *big.Int
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var r *bytes.Reader
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if chainID == nil || r == nil {
			return
		}

		btx, err := newSpanBatchTxs(txs, chainID)
		if err != nil {
			return
		}
		btx.decodeTxNonces(r)
	})
}

func Fuzz_Nosy_spanBatchTxs_decodeTxSigsRS__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var txs [][]byte
		fill_err = tp.Fill(&txs)
		if fill_err != nil {
			return
		}
		var chainID *big.Int
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var r *bytes.Reader
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if chainID == nil || r == nil {
			return
		}

		btx, err := newSpanBatchTxs(txs, chainID)
		if err != nil {
			return
		}
		btx.decodeTxSigsRS(r)
	})
}

func Fuzz_Nosy_spanBatchTxs_decodeTxTos__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var txs [][]byte
		fill_err = tp.Fill(&txs)
		if fill_err != nil {
			return
		}
		var chainID *big.Int
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var r *bytes.Reader
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if chainID == nil || r == nil {
			return
		}

		btx, err := newSpanBatchTxs(txs, chainID)
		if err != nil {
			return
		}
		btx.decodeTxTos(r)
	})
}

func Fuzz_Nosy_spanBatchTxs_decodeYParityBits__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var txs [][]byte
		fill_err = tp.Fill(&txs)
		if fill_err != nil {
			return
		}
		var chainID *big.Int
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var r *bytes.Reader
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if chainID == nil || r == nil {
			return
		}

		btx, err := newSpanBatchTxs(txs, chainID)
		if err != nil {
			return
		}
		btx.decodeYParityBits(r)
	})
}

func Fuzz_Nosy_spanBatchTxs_encode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var txs [][]byte
		fill_err = tp.Fill(&txs)
		if fill_err != nil {
			return
		}
		var chainID *big.Int
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var w io.Writer
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if chainID == nil {
			return
		}

		btx, err := newSpanBatchTxs(txs, chainID)
		if err != nil {
			return
		}
		btx.encode(w)
	})
}

func Fuzz_Nosy_spanBatchTxs_encodeContractCreationBits__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var txs [][]byte
		fill_err = tp.Fill(&txs)
		if fill_err != nil {
			return
		}
		var chainID *big.Int
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var w io.Writer
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if chainID == nil {
			return
		}

		btx, err := newSpanBatchTxs(txs, chainID)
		if err != nil {
			return
		}
		btx.encodeContractCreationBits(w)
	})
}

func Fuzz_Nosy_spanBatchTxs_encodeProtectedBits__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var txs [][]byte
		fill_err = tp.Fill(&txs)
		if fill_err != nil {
			return
		}
		var chainID *big.Int
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var w io.Writer
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if chainID == nil {
			return
		}

		btx, err := newSpanBatchTxs(txs, chainID)
		if err != nil {
			return
		}
		btx.encodeProtectedBits(w)
	})
}

func Fuzz_Nosy_spanBatchTxs_encodeTxDatas__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var txs [][]byte
		fill_err = tp.Fill(&txs)
		if fill_err != nil {
			return
		}
		var chainID *big.Int
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var w io.Writer
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if chainID == nil {
			return
		}

		btx, err := newSpanBatchTxs(txs, chainID)
		if err != nil {
			return
		}
		btx.encodeTxDatas(w)
	})
}

func Fuzz_Nosy_spanBatchTxs_encodeTxGases__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var txs [][]byte
		fill_err = tp.Fill(&txs)
		if fill_err != nil {
			return
		}
		var chainID *big.Int
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var w io.Writer
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if chainID == nil {
			return
		}

		btx, err := newSpanBatchTxs(txs, chainID)
		if err != nil {
			return
		}
		btx.encodeTxGases(w)
	})
}

func Fuzz_Nosy_spanBatchTxs_encodeTxNonces__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var txs [][]byte
		fill_err = tp.Fill(&txs)
		if fill_err != nil {
			return
		}
		var chainID *big.Int
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var w io.Writer
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if chainID == nil {
			return
		}

		btx, err := newSpanBatchTxs(txs, chainID)
		if err != nil {
			return
		}
		btx.encodeTxNonces(w)
	})
}

func Fuzz_Nosy_spanBatchTxs_encodeTxSigsRS__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var txs [][]byte
		fill_err = tp.Fill(&txs)
		if fill_err != nil {
			return
		}
		var chainID *big.Int
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var w io.Writer
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if chainID == nil {
			return
		}

		btx, err := newSpanBatchTxs(txs, chainID)
		if err != nil {
			return
		}
		btx.encodeTxSigsRS(w)
	})
}

func Fuzz_Nosy_spanBatchTxs_encodeTxTos__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var txs [][]byte
		fill_err = tp.Fill(&txs)
		if fill_err != nil {
			return
		}
		var chainID *big.Int
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var w io.Writer
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if chainID == nil {
			return
		}

		btx, err := newSpanBatchTxs(txs, chainID)
		if err != nil {
			return
		}
		btx.encodeTxTos(w)
	})
}

func Fuzz_Nosy_spanBatchTxs_encodeYParityBits__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var txs [][]byte
		fill_err = tp.Fill(&txs)
		if fill_err != nil {
			return
		}
		var chainID *big.Int
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var w io.Writer
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if chainID == nil {
			return
		}

		btx, err := newSpanBatchTxs(txs, chainID)
		if err != nil {
			return
		}
		btx.encodeYParityBits(w)
	})
}

func Fuzz_Nosy_spanBatchTxs_fullTxs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var txs [][]byte
		fill_err = tp.Fill(&txs)
		if fill_err != nil {
			return
		}
		var c2 *big.Int
		fill_err = tp.Fill(&c2)
		if fill_err != nil {
			return
		}
		var c3 *big.Int
		fill_err = tp.Fill(&c3)
		if fill_err != nil {
			return
		}
		if c2 == nil || c3 == nil {
			return
		}

		btx, err := newSpanBatchTxs(txs, c2)
		if err != nil {
			return
		}
		btx.fullTxs(c3)
	})
}

func Fuzz_Nosy_spanBatchTxs_recoverV__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var txs [][]byte
		fill_err = tp.Fill(&txs)
		if fill_err != nil {
			return
		}
		var c2 *big.Int
		fill_err = tp.Fill(&c2)
		if fill_err != nil {
			return
		}
		var c3 *big.Int
		fill_err = tp.Fill(&c3)
		if fill_err != nil {
			return
		}
		if c2 == nil || c3 == nil {
			return
		}

		btx, err := newSpanBatchTxs(txs, c2)
		if err != nil {
			return
		}
		btx.recoverV(c3)
	})
}

// skipping Fuzz_Nosy_AltDAInputFetcher_AdvanceL1Origin__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.AltDAInputFetcher

// skipping Fuzz_Nosy_AltDAInputFetcher_GetInput__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.AltDAInputFetcher

// skipping Fuzz_Nosy_AltDAInputFetcher_Reset__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.AltDAInputFetcher

// skipping Fuzz_Nosy_AttributesBuilder_PreparePayloadAttributes__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.AttributesBuilder

// skipping Fuzz_Nosy_Batch_AsSingularBatch__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.Batch

// skipping Fuzz_Nosy_Batch_AsSpanBatch__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.Batch

// skipping Fuzz_Nosy_Batch_GetBatchType__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.Batch

// skipping Fuzz_Nosy_Batch_GetTimestamp__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.Batch

// skipping Fuzz_Nosy_Batch_LogContext__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.Batch

func Fuzz_Nosy_ChannelCompressor_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var algo CompressionAlgo
		fill_err = tp.Fill(&algo)
		if fill_err != nil {
			return
		}

		_x1, err := NewChannelCompressor(algo)
		if err != nil {
			return
		}
		_x1.Close()
	})
}

func Fuzz_Nosy_ChannelCompressor_Flush__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var algo CompressionAlgo
		fill_err = tp.Fill(&algo)
		if fill_err != nil {
			return
		}

		_x1, err := NewChannelCompressor(algo)
		if err != nil {
			return
		}
		_x1.Flush()
	})
}

func Fuzz_Nosy_ChannelCompressor_GetCompressed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var algo CompressionAlgo
		fill_err = tp.Fill(&algo)
		if fill_err != nil {
			return
		}

		_x1, err := NewChannelCompressor(algo)
		if err != nil {
			return
		}
		_x1.GetCompressed()
	})
}

func Fuzz_Nosy_ChannelCompressor_Len__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var algo CompressionAlgo
		fill_err = tp.Fill(&algo)
		if fill_err != nil {
			return
		}

		_x1, err := NewChannelCompressor(algo)
		if err != nil {
			return
		}
		_x1.Len()
	})
}

func Fuzz_Nosy_ChannelCompressor_Read__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var algo CompressionAlgo
		fill_err = tp.Fill(&algo)
		if fill_err != nil {
			return
		}
		var _x2 []byte
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}

		_x1, err := NewChannelCompressor(algo)
		if err != nil {
			return
		}
		_x1.Read(_x2)
	})
}

func Fuzz_Nosy_ChannelCompressor_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var algo CompressionAlgo
		fill_err = tp.Fill(&algo)
		if fill_err != nil {
			return
		}

		_x1, err := NewChannelCompressor(algo)
		if err != nil {
			return
		}
		_x1.Reset()
	})
}

func Fuzz_Nosy_ChannelCompressor_Write__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var algo CompressionAlgo
		fill_err = tp.Fill(&algo)
		if fill_err != nil {
			return
		}
		var _x2 []byte
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}

		_x1, err := NewChannelCompressor(algo)
		if err != nil {
			return
		}
		_x1.Write(_x2)
	})
}

// skipping Fuzz_Nosy_ChannelFlusher_FlushChannel__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.ChannelFlusher

func Fuzz_Nosy_ChannelID_MarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id ChannelID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}

		id.MarshalText()
	})
}

func Fuzz_Nosy_ChannelID_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id ChannelID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}

		id.String()
	})
}

func Fuzz_Nosy_ChannelID_TerminalString__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id ChannelID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}

		id.TerminalString()
	})
}

// skipping Fuzz_Nosy_ChannelOut_AddBlock__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.ChannelOut

// skipping Fuzz_Nosy_ChannelOut_Close__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.ChannelOut

// skipping Fuzz_Nosy_ChannelOut_Flush__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.ChannelOut

// skipping Fuzz_Nosy_ChannelOut_FullErr__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.ChannelOut

// skipping Fuzz_Nosy_ChannelOut_ID__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.ChannelOut

// skipping Fuzz_Nosy_ChannelOut_InputBytes__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.ChannelOut

// skipping Fuzz_Nosy_ChannelOut_OutputFrame__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.ChannelOut

// skipping Fuzz_Nosy_ChannelOut_ReadyBytes__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.ChannelOut

// skipping Fuzz_Nosy_ChannelOut_Reset__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.ChannelOut

// skipping Fuzz_Nosy_ChannelStageSpec_ChannelTimeout__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.ChannelStageSpec

// skipping Fuzz_Nosy_ChannelStageSpec_MaxRLPBytesPerChannel__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.ChannelStageSpec

func Fuzz_Nosy_CompressionAlgo_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var algo CompressionAlgo
		fill_err = tp.Fill(&algo)
		if fill_err != nil {
			return
		}

		algo.String()
	})
}

// skipping Fuzz_Nosy_Compressor_Flush__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.Compressor

// skipping Fuzz_Nosy_Compressor_FullErr__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.Compressor

// skipping Fuzz_Nosy_Compressor_Len__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.Compressor

// skipping Fuzz_Nosy_Compressor_Reset__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.Compressor

// skipping Fuzz_Nosy_CompressorWriter_Close__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.CompressorWriter

// skipping Fuzz_Nosy_CompressorWriter_Flush__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.CompressorWriter

// skipping Fuzz_Nosy_CompressorWriter_Reset__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.CompressorWriter

// skipping Fuzz_Nosy_CompressorWriter_Write__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.CompressorWriter

func Fuzz_Nosy_ConfirmPipelineResetEvent_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d ConfirmPipelineResetEvent
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}

		d.String()
	})
}

func Fuzz_Nosy_ConfirmReceivedAttributesEvent_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d ConfirmReceivedAttributesEvent
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}

		d.String()
	})
}

// skipping Fuzz_Nosy_DataAvailabilitySource_OpenData__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.DataAvailabilitySource

// skipping Fuzz_Nosy_DataIter_Next__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.DataIter

func Fuzz_Nosy_DepositsOnlyPayloadAttributesRequestEvent_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ev DepositsOnlyPayloadAttributesRequestEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}

		ev.String()
	})
}

func Fuzz_Nosy_DerivedAttributesEvent_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ev DerivedAttributesEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}

		ev.String()
	})
}

func Fuzz_Nosy_DeriverIdleEvent_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d DeriverIdleEvent
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}

		d.String()
	})
}

func Fuzz_Nosy_DeriverL1StatusEvent_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d DeriverL1StatusEvent
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}

		d.String()
	})
}

func Fuzz_Nosy_DeriverMoreEvent_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d DeriverMoreEvent
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}

		d.String()
	})
}

func Fuzz_Nosy_Error_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e Error
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}

		e.Error()
	})
}

// skipping Fuzz_Nosy_Error_Is__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_Error_Unwrap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e Error
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}

		e.Unwrap()
	})
}

// skipping Fuzz_Nosy_ForkTransformer_Transform__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.ForkTransformer

// skipping Fuzz_Nosy_InnerBatchData_GetBatchType__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.InnerBatchData

// skipping Fuzz_Nosy_InnerBatchData_decode__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.InnerBatchData

// skipping Fuzz_Nosy_InnerBatchData_encode__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.InnerBatchData

// skipping Fuzz_Nosy_L1BlobsFetcher_GetBlobs__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.L1BlobsFetcher

// skipping Fuzz_Nosy_L1BlockRefByHashFetcher_L1BlockRefByHash__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.L1BlockRefByHashFetcher

// skipping Fuzz_Nosy_L1BlockRefByNumber_L1BlockRefByNumber__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.L1BlockRefByNumber

// skipping Fuzz_Nosy_L1BlockRefByNumberFetcher_FetchReceipts__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.L1BlockRefByNumberFetcher

// skipping Fuzz_Nosy_L1BlockRefByNumberFetcher_L1BlockRefByNumber__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.L1BlockRefByNumberFetcher

// skipping Fuzz_Nosy_L1Fetcher_L1BlockRefByLabel__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.L1Fetcher

// skipping Fuzz_Nosy_L1ReceiptsFetcher_FetchReceipts__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.L1ReceiptsFetcher

// skipping Fuzz_Nosy_L1ReceiptsFetcher_InfoByHash__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.L1ReceiptsFetcher

// skipping Fuzz_Nosy_L1TransactionFetcher_InfoAndTxsByHash__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.L1TransactionFetcher

// skipping Fuzz_Nosy_L2BlockRefSource_Hash__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.L2BlockRefSource

// skipping Fuzz_Nosy_L2BlockRefSource_NumberU64__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.L2BlockRefSource

// skipping Fuzz_Nosy_L2BlockRefSource_ParentHash__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.L2BlockRefSource

// skipping Fuzz_Nosy_L2BlockRefSource_Time__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.L2BlockRefSource

// skipping Fuzz_Nosy_L2BlockRefSource_Transactions__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.L2BlockRefSource

// skipping Fuzz_Nosy_L2Source_L2BlockRefByHash__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.L2Source

// skipping Fuzz_Nosy_L2Source_L2BlockRefByLabel__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.L2Source

// skipping Fuzz_Nosy_L2Source_L2BlockRefByNumber__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.L2Source

// skipping Fuzz_Nosy_L2Source_PayloadByHash__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.L2Source

// skipping Fuzz_Nosy_L2Source_PayloadByNumber__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.L2Source

func Fuzz_Nosy_Level_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var lvl Level
		fill_err = tp.Fill(&lvl)
		if fill_err != nil {
			return
		}

		lvl.String()
	})
}

// skipping Fuzz_Nosy_Metrics_RecordChannelInputBytes__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.Metrics

// skipping Fuzz_Nosy_Metrics_RecordChannelTimedOut__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.Metrics

// skipping Fuzz_Nosy_Metrics_RecordDerivedBatches__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.Metrics

// skipping Fuzz_Nosy_Metrics_RecordFrame__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.Metrics

// skipping Fuzz_Nosy_Metrics_RecordHeadChannelOpened__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.Metrics

// skipping Fuzz_Nosy_Metrics_RecordL1Ref__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.Metrics

// skipping Fuzz_Nosy_Metrics_RecordL2Ref__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.Metrics

// skipping Fuzz_Nosy_Metrics_RecordPipelineReset__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.Metrics

// skipping Fuzz_Nosy_Metrics_SetDerivationIdle__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.Metrics

// skipping Fuzz_Nosy_NextBatchProvider_NextBatch__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.NextBatchProvider

// skipping Fuzz_Nosy_NextBatchProvider_Origin__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.NextBatchProvider

// skipping Fuzz_Nosy_NextBlockProvider_NextL1Block__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.NextBlockProvider

// skipping Fuzz_Nosy_NextBlockProvider_Origin__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.NextBlockProvider

// skipping Fuzz_Nosy_NextBlockProvider_SystemConfig__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.NextBlockProvider

// skipping Fuzz_Nosy_NextDataProvider_NextData__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.NextDataProvider

// skipping Fuzz_Nosy_NextDataProvider_Origin__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.NextDataProvider

// skipping Fuzz_Nosy_NextFrameProvider_NextFrame__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.NextFrameProvider

// skipping Fuzz_Nosy_NextFrameProvider_Origin__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.NextFrameProvider

func Fuzz_Nosy_PipelineStepEvent_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ev PipelineStepEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}

		ev.String()
	})
}

// skipping Fuzz_Nosy_RawChannelProvider_NextRawChannel__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.RawChannelProvider

// skipping Fuzz_Nosy_RawChannelProvider_Origin__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.RawChannelProvider

// skipping Fuzz_Nosy_ResettableStage_Reset__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.ResettableStage

// skipping Fuzz_Nosy_SafeBlockFetcher_L2BlockRefByNumber__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.SafeBlockFetcher

// skipping Fuzz_Nosy_SafeBlockFetcher_PayloadByNumber__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.SafeBlockFetcher

// skipping Fuzz_Nosy_SingularBatchProvider_NextBatch__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.SingularBatchProvider

// skipping Fuzz_Nosy_SingularBatchProvider_Origin__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.SingularBatchProvider

// skipping Fuzz_Nosy_SystemConfigL2Fetcher_SystemConfigByL2Hash__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.SystemConfigL2Fetcher

// skipping Fuzz_Nosy_batchWithMetadata_LogContext__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/log.Logger

// skipping Fuzz_Nosy_spanBatchTxData_txType__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/derive.spanBatchTxData

func Fuzz_Nosy_BatchReader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r io.Reader
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var maxRLPBytesPerChannel uint64
		fill_err = tp.Fill(&maxRLPBytesPerChannel)
		if fill_err != nil {
			return
		}
		var isFjord bool
		fill_err = tp.Fill(&isFjord)
		if fill_err != nil {
			return
		}

		BatchReader(r, maxRLPBytesPerChannel, isFjord)
	})
}

func Fuzz_Nosy_BlockToSingularBatch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rollupCfg *rollup.Config
		fill_err = tp.Fill(&rollupCfg)
		if fill_err != nil {
			return
		}
		var block *types.Block
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		if rollupCfg == nil || block == nil {
			return
		}

		BlockToSingularBatch(rollupCfg, block)
	})
}

// skipping Fuzz_Nosy_DataFromEVMTransactions__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/log.Logger

// skipping Fuzz_Nosy_DepositsCompleteBytes__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/eth.BlockInfo

func Fuzz_Nosy_DeriveDeposits__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var receipts []*types.Receipt
		fill_err = tp.Fill(&receipts)
		if fill_err != nil {
			return
		}
		var depositContractAddr common.Address
		fill_err = tp.Fill(&depositContractAddr)
		if fill_err != nil {
			return
		}

		DeriveDeposits(receipts, depositContractAddr)
	})
}

func Fuzz_Nosy_ForceCloseTxData__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var frames []Frame
		fill_err = tp.Fill(&frames)
		if fill_err != nil {
			return
		}

		ForceCloseTxData(frames)
	})
}

func Fuzz_Nosy_GetBrotliLevel__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var algo CompressionAlgo
		fill_err = tp.Fill(&algo)
		if fill_err != nil {
			return
		}

		GetBrotliLevel(algo)
	})
}

// skipping Fuzz_Nosy_L1InfoDepositBytes__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/eth.BlockInfo

func Fuzz_Nosy_ParseFrames__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d1 []byte) {
		ParseFrames(d1)
	})
}

func Fuzz_Nosy_ReadTxData__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *bytes.Reader
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		ReadTxData(r)
	})
}

func Fuzz_Nosy_UserDeposits__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var receipts []*types.Receipt
		fill_err = tp.Fill(&receipts)
		if fill_err != nil {
			return
		}
		var depositContractAddr common.Address
		fill_err = tp.Fill(&depositContractAddr)
		if fill_err != nil {
			return
		}

		UserDeposits(receipts, depositContractAddr)
	})
}

func Fuzz_Nosy_ValidCompressionAlgo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var value CompressionAlgo
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}

		ValidCompressionAlgo(value)
	})
}

// skipping Fuzz_Nosy_checkSpanBatchPrefix__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/log.Logger

func Fuzz_Nosy_convertVToYParity__(f *testing.F) {
	f.Fuzz(func(t *testing.T, v uint64, txType int) {
		convertVToYParity(v, txType)
	})
}

// skipping Fuzz_Nosy_dataAndHashesFromTxs__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/log.Logger

func Fuzz_Nosy_frameSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var frame Frame
		fill_err = tp.Fill(&frame)
		if fill_err != nil {
			return
		}

		frameSize(frame)
	})
}

func Fuzz_Nosy_isEcotoneButNotFirstBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rollupCfg *rollup.Config
		fill_err = tp.Fill(&rollupCfg)
		if fill_err != nil {
			return
		}
		var l2Timestamp uint64
		fill_err = tp.Fill(&l2Timestamp)
		if fill_err != nil {
			return
		}
		if rollupCfg == nil {
			return
		}

		isEcotoneButNotFirstBlock(rollupCfg, l2Timestamp)
	})
}

func Fuzz_Nosy_isInteropButNotFirstBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rollupCfg *rollup.Config
		fill_err = tp.Fill(&rollupCfg)
		if fill_err != nil {
			return
		}
		var l2Timestamp uint64
		fill_err = tp.Fill(&l2Timestamp)
		if fill_err != nil {
			return
		}
		if rollupCfg == nil {
			return
		}

		isInteropButNotFirstBlock(rollupCfg, l2Timestamp)
	})
}

func Fuzz_Nosy_isProtectedV__(f *testing.F) {
	f.Fuzz(func(t *testing.T, v uint64, txType int) {
		isProtectedV(v, txType)
	})
}

// skipping Fuzz_Nosy_isValidBatchTx__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/types.Signer

func Fuzz_Nosy_marshalBinaryWithSignature__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var info *L1BlockInfo
		fill_err = tp.Fill(&info)
		if fill_err != nil {
			return
		}
		var signature []byte
		fill_err = tp.Fill(&signature)
		if fill_err != nil {
			return
		}
		if info == nil {
			return
		}

		marshalBinaryWithSignature(info, signature)
	})
}

func Fuzz_Nosy_marshalDepositVersion0__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var deposit *types.DepositTx
		fill_err = tp.Fill(&deposit)
		if fill_err != nil {
			return
		}
		if deposit == nil {
			return
		}

		marshalDepositVersion0(deposit)
	})
}

func Fuzz_Nosy_pruneFrameQueue__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var frames []Frame
		fill_err = tp.Fill(&frames)
		if fill_err != nil {
			return
		}

		pruneFrameQueue(frames)
	})
}

func Fuzz_Nosy_upgradeToCalldata__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}

		upgradeToCalldata(addr)
	})
}
