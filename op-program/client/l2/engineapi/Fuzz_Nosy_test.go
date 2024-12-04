package engineapi

import (
	"context"
	"testing"

	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum/go-ethereum/beacon/engine"
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

func Fuzz_Nosy_BlockProcessor_AddTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BlockProcessor
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var tx *types.Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if b == nil || tx == nil {
			return
		}

		b.AddTx(tx)
	})
}

func Fuzz_Nosy_BlockProcessor_Assemble__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BlockProcessor
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.Assemble()
	})
}

func Fuzz_Nosy_BlockProcessor_CheckTxWithinGasLimit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BlockProcessor
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var tx *types.Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if b == nil || tx == nil {
			return
		}

		b.CheckTxWithinGasLimit(tx)
	})
}

func Fuzz_Nosy_BlockProcessor_Commit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BlockProcessor
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.Commit()
	})
}

func Fuzz_Nosy_L2EngineAPI_ForcedEmpty__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ea *L2EngineAPI
		fill_err = tp.Fill(&ea)
		if fill_err != nil {
			return
		}
		if ea == nil {
			return
		}

		ea.ForcedEmpty()
	})
}

func Fuzz_Nosy_L2EngineAPI_ForkchoiceUpdatedV1__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ea *L2EngineAPI
		fill_err = tp.Fill(&ea)
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
		if ea == nil || state == nil || attr == nil {
			return
		}

		ea.ForkchoiceUpdatedV1(ctx, state, attr)
	})
}

func Fuzz_Nosy_L2EngineAPI_ForkchoiceUpdatedV2__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ea *L2EngineAPI
		fill_err = tp.Fill(&ea)
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
		if ea == nil || state == nil || attr == nil {
			return
		}

		ea.ForkchoiceUpdatedV2(ctx, state, attr)
	})
}

func Fuzz_Nosy_L2EngineAPI_ForkchoiceUpdatedV3__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ea *L2EngineAPI
		fill_err = tp.Fill(&ea)
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
		if ea == nil || state == nil || attr == nil {
			return
		}

		ea.ForkchoiceUpdatedV3(ctx, state, attr)
	})
}

func Fuzz_Nosy_L2EngineAPI_GetPayloadV1__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ea *L2EngineAPI
		fill_err = tp.Fill(&ea)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var payloadId engine.PayloadID
		fill_err = tp.Fill(&payloadId)
		if fill_err != nil {
			return
		}
		if ea == nil {
			return
		}

		ea.GetPayloadV1(ctx, payloadId)
	})
}

func Fuzz_Nosy_L2EngineAPI_GetPayloadV2__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ea *L2EngineAPI
		fill_err = tp.Fill(&ea)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var payloadId engine.PayloadID
		fill_err = tp.Fill(&payloadId)
		if fill_err != nil {
			return
		}
		if ea == nil {
			return
		}

		ea.GetPayloadV2(ctx, payloadId)
	})
}

func Fuzz_Nosy_L2EngineAPI_GetPayloadV3__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ea *L2EngineAPI
		fill_err = tp.Fill(&ea)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var payloadId engine.PayloadID
		fill_err = tp.Fill(&payloadId)
		if fill_err != nil {
			return
		}
		if ea == nil {
			return
		}

		ea.GetPayloadV3(ctx, payloadId)
	})
}

func Fuzz_Nosy_L2EngineAPI_IncludeTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ea *L2EngineAPI
		fill_err = tp.Fill(&ea)
		if fill_err != nil {
			return
		}
		var tx *types.Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		var from common.Address
		fill_err = tp.Fill(&from)
		if fill_err != nil {
			return
		}
		if ea == nil || tx == nil {
			return
		}

		ea.IncludeTx(tx, from)
	})
}

func Fuzz_Nosy_L2EngineAPI_NewPayloadV1__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ea *L2EngineAPI
		fill_err = tp.Fill(&ea)
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
		if ea == nil || payload == nil {
			return
		}

		ea.NewPayloadV1(ctx, payload)
	})
}

func Fuzz_Nosy_L2EngineAPI_NewPayloadV2__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ea *L2EngineAPI
		fill_err = tp.Fill(&ea)
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
		if ea == nil || payload == nil {
			return
		}

		ea.NewPayloadV2(ctx, payload)
	})
}

func Fuzz_Nosy_L2EngineAPI_NewPayloadV3__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ea *L2EngineAPI
		fill_err = tp.Fill(&ea)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var params *eth.ExecutionPayload
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		var versionedHashes []common.Hash
		fill_err = tp.Fill(&versionedHashes)
		if fill_err != nil {
			return
		}
		var beaconRoot *common.Hash
		fill_err = tp.Fill(&beaconRoot)
		if fill_err != nil {
			return
		}
		if ea == nil || params == nil || beaconRoot == nil {
			return
		}

		ea.NewPayloadV3(ctx, params, versionedHashes, beaconRoot)
	})
}

func Fuzz_Nosy_L2EngineAPI_PendingIndices__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ea *L2EngineAPI
		fill_err = tp.Fill(&ea)
		if fill_err != nil {
			return
		}
		var from common.Address
		fill_err = tp.Fill(&from)
		if fill_err != nil {
			return
		}
		if ea == nil {
			return
		}

		ea.PendingIndices(from)
	})
}

func Fuzz_Nosy_L2EngineAPI_RemainingBlockGas__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ea *L2EngineAPI
		fill_err = tp.Fill(&ea)
		if fill_err != nil {
			return
		}
		if ea == nil {
			return
		}

		ea.RemainingBlockGas()
	})
}

func Fuzz_Nosy_L2EngineAPI_config__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ea *L2EngineAPI
		fill_err = tp.Fill(&ea)
		if fill_err != nil {
			return
		}
		if ea == nil {
			return
		}

		ea.config()
	})
}

func Fuzz_Nosy_L2EngineAPI_endBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ea *L2EngineAPI
		fill_err = tp.Fill(&ea)
		if fill_err != nil {
			return
		}
		if ea == nil {
			return
		}

		ea.endBlock()
	})
}

func Fuzz_Nosy_L2EngineAPI_forkchoiceUpdated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ea *L2EngineAPI
		fill_err = tp.Fill(&ea)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
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
		if ea == nil || state == nil || attr == nil {
			return
		}

		ea.forkchoiceUpdated(_x2, state, attr)
	})
}

func Fuzz_Nosy_L2EngineAPI_getPayload__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ea *L2EngineAPI
		fill_err = tp.Fill(&ea)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var payloadId engine.PayloadID
		fill_err = tp.Fill(&payloadId)
		if fill_err != nil {
			return
		}
		if ea == nil {
			return
		}

		ea.getPayload(_x2, payloadId)
	})
}

// skipping Fuzz_Nosy_L2EngineAPI_invalid__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_L2EngineAPI_newPayload__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ea *L2EngineAPI
		fill_err = tp.Fill(&ea)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var payload *eth.ExecutionPayload
		fill_err = tp.Fill(&payload)
		if fill_err != nil {
			return
		}
		var hashes []common.Hash
		fill_err = tp.Fill(&hashes)
		if fill_err != nil {
			return
		}
		var root *common.Hash
		fill_err = tp.Fill(&root)
		if fill_err != nil {
			return
		}
		if ea == nil || payload == nil || root == nil {
			return
		}

		ea.newPayload(_x2, payload, hashes, root)
	})
}

func Fuzz_Nosy_L2EngineAPI_startBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ea *L2EngineAPI
		fill_err = tp.Fill(&ea)
		if fill_err != nil {
			return
		}
		var parent common.Hash
		fill_err = tp.Fill(&parent)
		if fill_err != nil {
			return
		}
		var attrs *eth.PayloadAttributes
		fill_err = tp.Fill(&attrs)
		if fill_err != nil {
			return
		}
		if ea == nil || attrs == nil {
			return
		}

		ea.startBlock(parent, attrs)
	})
}

func Fuzz_Nosy_L2EngineAPI_verifyPayloadAttributes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ea *L2EngineAPI
		fill_err = tp.Fill(&ea)
		if fill_err != nil {
			return
		}
		var attr *eth.PayloadAttributes
		fill_err = tp.Fill(&attr)
		if fill_err != nil {
			return
		}
		if ea == nil || attr == nil {
			return
		}

		ea.verifyPayloadAttributes(attr)
	})
}

func Fuzz_Nosy_bn256PairingOracle_RequiredGas__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *bn256PairingOracle
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.RequiredGas(input)
	})
}

func Fuzz_Nosy_bn256PairingOracle_Run__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *bn256PairingOracle
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.Run(input)
	})
}

func Fuzz_Nosy_bn256PairingOracleGranite_Run__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *bn256PairingOracleGranite
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.Run(input)
	})
}

func Fuzz_Nosy_ecrecoverOracle_RequiredGas__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *ecrecoverOracle
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.RequiredGas(input)
	})
}

func Fuzz_Nosy_ecrecoverOracle_Run__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *ecrecoverOracle
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.Run(input)
	})
}

func Fuzz_Nosy_kzgPointEvaluationOracle_RequiredGas__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *kzgPointEvaluationOracle
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.RequiredGas(input)
	})
}

func Fuzz_Nosy_kzgPointEvaluationOracle_Run__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *kzgPointEvaluationOracle
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.Run(input)
	})
}

// skipping Fuzz_Nosy_BlockDataProvider_Config__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-program/client/l2/engineapi.BlockDataProvider

// skipping Fuzz_Nosy_BlockDataProvider_Engine__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-program/client/l2/engineapi.BlockDataProvider

// skipping Fuzz_Nosy_BlockDataProvider_GetHeader__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-program/client/l2/engineapi.BlockDataProvider

// skipping Fuzz_Nosy_BlockDataProvider_GetVMConfig__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-program/client/l2/engineapi.BlockDataProvider

// skipping Fuzz_Nosy_BlockDataProvider_StateAt__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-program/client/l2/engineapi.BlockDataProvider

// skipping Fuzz_Nosy_CachingEngineBackend_AssembleAndInsertBlockWithoutSetHead__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-program/client/l2/engineapi.CachingEngineBackend

// skipping Fuzz_Nosy_EngineBackend_Config__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-program/client/l2/engineapi.EngineBackend

// skipping Fuzz_Nosy_EngineBackend_CurrentFinalBlock__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-program/client/l2/engineapi.EngineBackend

// skipping Fuzz_Nosy_EngineBackend_CurrentSafeBlock__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-program/client/l2/engineapi.EngineBackend

// skipping Fuzz_Nosy_EngineBackend_Engine__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-program/client/l2/engineapi.EngineBackend

// skipping Fuzz_Nosy_EngineBackend_GetBlock__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-program/client/l2/engineapi.EngineBackend

// skipping Fuzz_Nosy_EngineBackend_GetBlockByHash__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-program/client/l2/engineapi.EngineBackend

// skipping Fuzz_Nosy_EngineBackend_GetCanonicalHash__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-program/client/l2/engineapi.EngineBackend

// skipping Fuzz_Nosy_EngineBackend_GetVMConfig__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-program/client/l2/engineapi.EngineBackend

// skipping Fuzz_Nosy_EngineBackend_HasBlockAndState__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-program/client/l2/engineapi.EngineBackend

// skipping Fuzz_Nosy_EngineBackend_InsertBlockWithoutSetHead__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-program/client/l2/engineapi.EngineBackend

// skipping Fuzz_Nosy_EngineBackend_SetCanonical__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-program/client/l2/engineapi.EngineBackend

// skipping Fuzz_Nosy_EngineBackend_SetFinalized__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-program/client/l2/engineapi.EngineBackend

// skipping Fuzz_Nosy_EngineBackend_SetSafe__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-program/client/l2/engineapi.EngineBackend

// skipping Fuzz_Nosy_EngineBackend_StateAt__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-program/client/l2/engineapi.EngineBackend

// skipping Fuzz_Nosy_PrecompileOracle_Precompile__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-program/client/l2/engineapi.PrecompileOracle

func Fuzz_Nosy_allZero__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		allZero(b)
	})
}

func Fuzz_Nosy_toGethWithdrawals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var payload *eth.ExecutionPayload
		fill_err = tp.Fill(&payload)
		if fill_err != nil {
			return
		}
		if payload == nil {
			return
		}

		toGethWithdrawals(payload)
	})
}
