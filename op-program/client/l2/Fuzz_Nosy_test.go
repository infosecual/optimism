package l2

import (
	"context"
	"testing"

	"github.com/ethereum-optimism/optimism/op-program/client/l2/engineapi"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum/go-ethereum/common"
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

func Fuzz_Nosy_CachingOracle_BlockByHash__(f *testing.F) {
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

		o.BlockByHash(blockHash)
	})
}

func Fuzz_Nosy_CachingOracle_CodeByHash__(f *testing.F) {
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
		var codeHash common.Hash
		fill_err = tp.Fill(&codeHash)
		if fill_err != nil {
			return
		}
		if o == nil {
			return
		}

		o.CodeByHash(codeHash)
	})
}

func Fuzz_Nosy_CachingOracle_NodeByHash__(f *testing.F) {
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
		var nodeHash common.Hash
		fill_err = tp.Fill(&nodeHash)
		if fill_err != nil {
			return
		}
		if o == nil {
			return
		}

		o.NodeByHash(nodeHash)
	})
}

func Fuzz_Nosy_CachingOracle_OutputByRoot__(f *testing.F) {
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
		var root common.Hash
		fill_err = tp.Fill(&root)
		if fill_err != nil {
			return
		}
		if o == nil {
			return
		}

		o.OutputByRoot(root)
	})
}

func Fuzz_Nosy_OracleBackedL2Chain_AssembleAndInsertBlockWithoutSetHead__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o *OracleBackedL2Chain
		fill_err = tp.Fill(&o)
		if fill_err != nil {
			return
		}
		var processor *engineapi.BlockProcessor
		fill_err = tp.Fill(&processor)
		if fill_err != nil {
			return
		}
		if o == nil || processor == nil {
			return
		}

		o.AssembleAndInsertBlockWithoutSetHead(processor)
	})
}

func Fuzz_Nosy_OracleBackedL2Chain_Config__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o *OracleBackedL2Chain
		fill_err = tp.Fill(&o)
		if fill_err != nil {
			return
		}
		if o == nil {
			return
		}

		o.Config()
	})
}

func Fuzz_Nosy_OracleBackedL2Chain_CurrentFinalBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o *OracleBackedL2Chain
		fill_err = tp.Fill(&o)
		if fill_err != nil {
			return
		}
		if o == nil {
			return
		}

		o.CurrentFinalBlock()
	})
}

func Fuzz_Nosy_OracleBackedL2Chain_CurrentHeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o *OracleBackedL2Chain
		fill_err = tp.Fill(&o)
		if fill_err != nil {
			return
		}
		if o == nil {
			return
		}

		o.CurrentHeader()
	})
}

func Fuzz_Nosy_OracleBackedL2Chain_CurrentSafeBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o *OracleBackedL2Chain
		fill_err = tp.Fill(&o)
		if fill_err != nil {
			return
		}
		if o == nil {
			return
		}

		o.CurrentSafeBlock()
	})
}

func Fuzz_Nosy_OracleBackedL2Chain_Engine__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o *OracleBackedL2Chain
		fill_err = tp.Fill(&o)
		if fill_err != nil {
			return
		}
		if o == nil {
			return
		}

		o.Engine()
	})
}

func Fuzz_Nosy_OracleBackedL2Chain_GetBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o *OracleBackedL2Chain
		fill_err = tp.Fill(&o)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
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

		o.GetBlock(hash, number)
	})
}

func Fuzz_Nosy_OracleBackedL2Chain_GetBlockByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o *OracleBackedL2Chain
		fill_err = tp.Fill(&o)
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

		o.GetBlockByHash(hash)
	})
}

func Fuzz_Nosy_OracleBackedL2Chain_GetCanonicalHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o *OracleBackedL2Chain
		fill_err = tp.Fill(&o)
		if fill_err != nil {
			return
		}
		var n uint64
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if o == nil {
			return
		}

		o.GetCanonicalHash(n)
	})
}

func Fuzz_Nosy_OracleBackedL2Chain_GetHeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o *OracleBackedL2Chain
		fill_err = tp.Fill(&o)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		var u uint64
		fill_err = tp.Fill(&u)
		if fill_err != nil {
			return
		}
		if o == nil {
			return
		}

		o.GetHeader(hash, u)
	})
}

func Fuzz_Nosy_OracleBackedL2Chain_GetHeaderByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o *OracleBackedL2Chain
		fill_err = tp.Fill(&o)
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

		o.GetHeaderByHash(hash)
	})
}

func Fuzz_Nosy_OracleBackedL2Chain_GetHeaderByNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o *OracleBackedL2Chain
		fill_err = tp.Fill(&o)
		if fill_err != nil {
			return
		}
		var n uint64
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if o == nil {
			return
		}

		o.GetHeaderByNumber(n)
	})
}

func Fuzz_Nosy_OracleBackedL2Chain_GetTd__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o *OracleBackedL2Chain
		fill_err = tp.Fill(&o)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
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

		o.GetTd(hash, number)
	})
}

func Fuzz_Nosy_OracleBackedL2Chain_GetVMConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o *OracleBackedL2Chain
		fill_err = tp.Fill(&o)
		if fill_err != nil {
			return
		}
		if o == nil {
			return
		}

		o.GetVMConfig()
	})
}

func Fuzz_Nosy_OracleBackedL2Chain_HasBlockAndState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o *OracleBackedL2Chain
		fill_err = tp.Fill(&o)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
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

		o.HasBlockAndState(hash, number)
	})
}

func Fuzz_Nosy_OracleBackedL2Chain_InsertBlockWithoutSetHead__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o *OracleBackedL2Chain
		fill_err = tp.Fill(&o)
		if fill_err != nil {
			return
		}
		var block *types.Block
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		var makeWitness bool
		fill_err = tp.Fill(&makeWitness)
		if fill_err != nil {
			return
		}
		if o == nil || block == nil {
			return
		}

		o.InsertBlockWithoutSetHead(block, makeWitness)
	})
}

func Fuzz_Nosy_OracleBackedL2Chain_SetCanonical__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o *OracleBackedL2Chain
		fill_err = tp.Fill(&o)
		if fill_err != nil {
			return
		}
		var head *types.Block
		fill_err = tp.Fill(&head)
		if fill_err != nil {
			return
		}
		if o == nil || head == nil {
			return
		}

		o.SetCanonical(head)
	})
}

func Fuzz_Nosy_OracleBackedL2Chain_SetFinalized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o *OracleBackedL2Chain
		fill_err = tp.Fill(&o)
		if fill_err != nil {
			return
		}
		var header *types.Header
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		if o == nil || header == nil {
			return
		}

		o.SetFinalized(header)
	})
}

func Fuzz_Nosy_OracleBackedL2Chain_SetSafe__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o *OracleBackedL2Chain
		fill_err = tp.Fill(&o)
		if fill_err != nil {
			return
		}
		var header *types.Header
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		if o == nil || header == nil {
			return
		}

		o.SetSafe(header)
	})
}

func Fuzz_Nosy_OracleBackedL2Chain_StateAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o *OracleBackedL2Chain
		fill_err = tp.Fill(&o)
		if fill_err != nil {
			return
		}
		var root common.Hash
		fill_err = tp.Fill(&root)
		if fill_err != nil {
			return
		}
		if o == nil {
			return
		}

		o.StateAt(root)
	})
}

func Fuzz_Nosy_OracleEngine_ForkchoiceUpdate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o *OracleEngine
		fill_err = tp.Fill(&o)
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
		if o == nil || state == nil || attr == nil {
			return
		}

		o.ForkchoiceUpdate(ctx, state, attr)
	})
}

func Fuzz_Nosy_OracleEngine_GetPayload__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o *OracleEngine
		fill_err = tp.Fill(&o)
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
		if o == nil {
			return
		}

		o.GetPayload(ctx, payloadInfo)
	})
}

func Fuzz_Nosy_OracleEngine_L2BlockRefByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o *OracleEngine
		fill_err = tp.Fill(&o)
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
		if o == nil {
			return
		}

		o.L2BlockRefByHash(ctx, l2Hash)
	})
}

func Fuzz_Nosy_OracleEngine_L2BlockRefByLabel__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o *OracleEngine
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

		o.L2BlockRefByLabel(ctx, label)
	})
}

func Fuzz_Nosy_OracleEngine_L2BlockRefByNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o *OracleEngine
		fill_err = tp.Fill(&o)
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
		if o == nil {
			return
		}

		o.L2BlockRefByNumber(ctx, n)
	})
}

func Fuzz_Nosy_OracleEngine_L2OutputRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o *OracleEngine
		fill_err = tp.Fill(&o)
		if fill_err != nil {
			return
		}
		var l2ClaimBlockNum uint64
		fill_err = tp.Fill(&l2ClaimBlockNum)
		if fill_err != nil {
			return
		}
		if o == nil {
			return
		}

		o.L2OutputRoot(l2ClaimBlockNum)
	})
}

func Fuzz_Nosy_OracleEngine_NewPayload__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o *OracleEngine
		fill_err = tp.Fill(&o)
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
		if o == nil || payload == nil || parentBeaconBlockRoot == nil {
			return
		}

		o.NewPayload(ctx, payload, parentBeaconBlockRoot)
	})
}

func Fuzz_Nosy_OracleEngine_PayloadByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o *OracleEngine
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

		o.PayloadByHash(ctx, hash)
	})
}

func Fuzz_Nosy_OracleEngine_PayloadByNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o *OracleEngine
		fill_err = tp.Fill(&o)
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
		if o == nil {
			return
		}

		o.PayloadByNumber(ctx, n)
	})
}

func Fuzz_Nosy_OracleEngine_SystemConfigByL2Hash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o *OracleEngine
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

		o.SystemConfigByL2Hash(ctx, hash)
	})
}

func Fuzz_Nosy_OracleKeyValueStore_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o *OracleKeyValueStore
		fill_err = tp.Fill(&o)
		if fill_err != nil {
			return
		}
		if o == nil {
			return
		}

		o.Close()
	})
}

func Fuzz_Nosy_OracleKeyValueStore_Compact__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o *OracleKeyValueStore
		fill_err = tp.Fill(&o)
		if fill_err != nil {
			return
		}
		var start []byte
		fill_err = tp.Fill(&start)
		if fill_err != nil {
			return
		}
		var limit []byte
		fill_err = tp.Fill(&limit)
		if fill_err != nil {
			return
		}
		if o == nil {
			return
		}

		o.Compact(start, limit)
	})
}

func Fuzz_Nosy_OracleKeyValueStore_Delete__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o *OracleKeyValueStore
		fill_err = tp.Fill(&o)
		if fill_err != nil {
			return
		}
		var key []byte
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		if o == nil {
			return
		}

		o.Delete(key)
	})
}

func Fuzz_Nosy_OracleKeyValueStore_Get__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o *OracleKeyValueStore
		fill_err = tp.Fill(&o)
		if fill_err != nil {
			return
		}
		var key []byte
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		if o == nil {
			return
		}

		o.Get(key)
	})
}

func Fuzz_Nosy_OracleKeyValueStore_Has__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o *OracleKeyValueStore
		fill_err = tp.Fill(&o)
		if fill_err != nil {
			return
		}
		var key []byte
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		if o == nil {
			return
		}

		o.Has(key)
	})
}

func Fuzz_Nosy_OracleKeyValueStore_NewBatch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o *OracleKeyValueStore
		fill_err = tp.Fill(&o)
		if fill_err != nil {
			return
		}
		if o == nil {
			return
		}

		o.NewBatch()
	})
}

func Fuzz_Nosy_OracleKeyValueStore_NewBatchWithSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o *OracleKeyValueStore
		fill_err = tp.Fill(&o)
		if fill_err != nil {
			return
		}
		var size int
		fill_err = tp.Fill(&size)
		if fill_err != nil {
			return
		}
		if o == nil {
			return
		}

		o.NewBatchWithSize(size)
	})
}

func Fuzz_Nosy_OracleKeyValueStore_NewIterator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o *OracleKeyValueStore
		fill_err = tp.Fill(&o)
		if fill_err != nil {
			return
		}
		var prefix []byte
		fill_err = tp.Fill(&prefix)
		if fill_err != nil {
			return
		}
		var start []byte
		fill_err = tp.Fill(&start)
		if fill_err != nil {
			return
		}
		if o == nil {
			return
		}

		o.NewIterator(prefix, start)
	})
}

func Fuzz_Nosy_OracleKeyValueStore_Put__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o *OracleKeyValueStore
		fill_err = tp.Fill(&o)
		if fill_err != nil {
			return
		}
		var key []byte
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var value []byte
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
		if o == nil {
			return
		}

		o.Put(key, value)
	})
}

func Fuzz_Nosy_OracleKeyValueStore_Stat__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o *OracleKeyValueStore
		fill_err = tp.Fill(&o)
		if fill_err != nil {
			return
		}
		if o == nil {
			return
		}

		o.Stat()
	})
}

func Fuzz_Nosy_PreimageOracle_BlockByHash__(f *testing.F) {
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

		p.BlockByHash(blockHash)
	})
}

func Fuzz_Nosy_PreimageOracle_CodeByHash__(f *testing.F) {
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
		var codeHash common.Hash
		fill_err = tp.Fill(&codeHash)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.CodeByHash(codeHash)
	})
}

func Fuzz_Nosy_PreimageOracle_LoadTransactions__(f *testing.F) {
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
		var txHash common.Hash
		fill_err = tp.Fill(&txHash)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.LoadTransactions(blockHash, txHash)
	})
}

func Fuzz_Nosy_PreimageOracle_NodeByHash__(f *testing.F) {
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
		var nodeHash common.Hash
		fill_err = tp.Fill(&nodeHash)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.NodeByHash(nodeHash)
	})
}

func Fuzz_Nosy_PreimageOracle_OutputByRoot__(f *testing.F) {
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
		var l2OutputRoot common.Hash
		fill_err = tp.Fill(&l2OutputRoot)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.OutputByRoot(l2OutputRoot)
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

func Fuzz_Nosy_CodeHint_Hint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l CodeHint
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}

		l.Hint()
	})
}

func Fuzz_Nosy_L2OutputHint_Hint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l L2OutputHint
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}

		l.Hint()
	})
}

// skipping Fuzz_Nosy_Oracle_BlockByHash__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-program/client/l2.Oracle

// skipping Fuzz_Nosy_Oracle_OutputByRoot__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-program/client/l2.Oracle

func Fuzz_Nosy_StateNodeHint_Hint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l StateNodeHint
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}

		l.Hint()
	})
}

// skipping Fuzz_Nosy_StateOracle_CodeByHash__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-program/client/l2.StateOracle

// skipping Fuzz_Nosy_StateOracle_NodeByHash__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-program/client/l2.StateOracle

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
