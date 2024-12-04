package sources

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
	
	// skipping Fuzz_Nosy_BasicRPCReceiptsFetcher_FetchReceipts__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/eth.BlockInfo

func Fuzz_Nosy_BasicRPCReceiptsFetcher_deleteBatchCall__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var f1 *BasicRPCReceiptsFetcher
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var blockHash common.Hash
		fill_err = tp.Fill(&blockHash)
		if fill_err != nil {
			return
		}
	if f1 == nil {
		return
	}

	f1.deleteBatchCall(blockHash)
	})
}

func Fuzz_Nosy_BasicRPCReceiptsFetcher_getOrCreateBatchCall__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var f1 *BasicRPCReceiptsFetcher
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var blockHash common.Hash
		fill_err = tp.Fill(&blockHash)
		if fill_err != nil {
			return
		}
		var txHashes []common.Hash
		fill_err = tp.Fill(&txHashes)
		if fill_err != nil {
			return
		}
	if f1 == nil {
		return
	}

	f1.getOrCreateBatchCall(blockHash, txHashes)
	})
}

func Fuzz_Nosy_BeaconHTTPClient_BeaconBlobSideCars__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var cl *BeaconHTTPClient
		fill_err = tp.Fill(&cl)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var fetchAllSidecars bool
		fill_err = tp.Fill(&fetchAllSidecars)
		if fill_err != nil {
			return
		}
		var slot uint64
		fill_err = tp.Fill(&slot)
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

	cl.BeaconBlobSideCars(ctx, fetchAllSidecars, slot, hashes)
	})
}

func Fuzz_Nosy_BeaconHTTPClient_BeaconGenesis__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var cl *BeaconHTTPClient
		fill_err = tp.Fill(&cl)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
	if cl == nil {
		return
	}

	cl.BeaconGenesis(ctx)
	})
}

func Fuzz_Nosy_BeaconHTTPClient_ConfigSpec__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var cl *BeaconHTTPClient
		fill_err = tp.Fill(&cl)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
	if cl == nil {
		return
	}

	cl.ConfigSpec(ctx)
	})
}

func Fuzz_Nosy_BeaconHTTPClient_NodeVersion__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var cl *BeaconHTTPClient
		fill_err = tp.Fill(&cl)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
	if cl == nil {
		return
	}

	cl.NodeVersion(ctx)
	})
}

// skipping Fuzz_Nosy_BeaconHTTPClient_apiReq__ because parameters include func, chan, or unsupported interface: any

// skipping Fuzz_Nosy_CachingReceiptsProvider_FetchReceipts__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/eth.BlockInfo

func Fuzz_Nosy_CachingReceiptsProvider_deleteFetchingLock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var p *CachingReceiptsProvider
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

	p.deleteFetchingLock(blockHash)
	})
}

func Fuzz_Nosy_CachingReceiptsProvider_getOrCreateFetchingLock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var p *CachingReceiptsProvider
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

	p.getOrCreateFetchingLock(blockHash)
	})
}

func Fuzz_Nosy_CachingReceiptsProvider_isInnerNil__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var p *CachingReceiptsProvider
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
	if p == nil {
		return
	}

	p.isInnerNil()
	})
}

func Fuzz_Nosy_ClientPool[T any]_Get__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var p *ClientPool[T]
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
	if p == nil {
		return
	}

	p.Get()
	})
}

func Fuzz_Nosy_ClientPool[T any]_Len__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var p *ClientPool[T]
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
	if p == nil {
		return
	}

	p.Len()
	})
}

func Fuzz_Nosy_ClientPool[T any]_MoveToNext__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var p *ClientPool[T]
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
	if p == nil {
		return
	}

	p.MoveToNext()
	})
}

func Fuzz_Nosy_DebugClient_CodeByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var o *DebugClient
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

	o.CodeByHash(ctx, hash)
	})
}

func Fuzz_Nosy_DebugClient_NodeByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var o *DebugClient
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

	o.NodeByHash(ctx, hash)
	})
}

func Fuzz_Nosy_DebugClient_dbGet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var o *DebugClient
		fill_err = tp.Fill(&o)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
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

	o.dbGet(ctx, key)
	})
}

func Fuzz_Nosy_EngineAPIClient_EngineVersionProvider__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *EngineAPIClient
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
	if s == nil {
		return
	}

	s.EngineVersionProvider()
	})
}

func Fuzz_Nosy_EngineAPIClient_ForkchoiceUpdate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *EngineAPIClient
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var fc *eth.ForkchoiceState
		fill_err = tp.Fill(&fc)
		if fill_err != nil {
			return
		}
		var attributes *eth.PayloadAttributes
		fill_err = tp.Fill(&attributes)
		if fill_err != nil {
			return
		}
	if s == nil|| fc == nil|| attributes == nil {
		return
	}

	s.ForkchoiceUpdate(ctx, fc, attributes)
	})
}

func Fuzz_Nosy_EngineAPIClient_GetPayload__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *EngineAPIClient
		fill_err = tp.Fill(&s)
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
	if s == nil {
		return
	}

	s.GetPayload(ctx, payloadInfo)
	})
}

func Fuzz_Nosy_EngineAPIClient_NewPayload__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *EngineAPIClient
		fill_err = tp.Fill(&s)
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
	if s == nil|| payload == nil|| parentBeaconBlockRoot == nil {
		return
	}

	s.NewPayload(ctx, payload, parentBeaconBlockRoot)
	})
}

func Fuzz_Nosy_EngineAPIClient_SignalSuperchainV1__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *EngineAPIClient
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var recommended params.ProtocolVersion
		fill_err = tp.Fill(&recommended)
		if fill_err != nil {
			return
		}
		var required params.ProtocolVersion
		fill_err = tp.Fill(&required)
		if fill_err != nil {
			return
		}
	if s == nil {
		return
	}

	s.SignalSuperchainV1(ctx, recommended, required)
	})
}

func Fuzz_Nosy_EthClient_ChainID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *EthClient
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

	s.ChainID(ctx)
	})
}

func Fuzz_Nosy_EthClient_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *EthClient
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
	if s == nil {
		return
	}

	s.Close()
	})
}

func Fuzz_Nosy_EthClient_ExecutionWitness__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *EthClient
		fill_err = tp.Fill(&s)
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
	if s == nil {
		return
	}

	s.ExecutionWitness(ctx, blockNum)
	})
}

func Fuzz_Nosy_EthClient_FetchReceipts__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *EthClient
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

func Fuzz_Nosy_EthClient_GetProof__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *EthClient
		fill_err = tp.Fill(&s)
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
	if s == nil {
		return
	}

	s.GetProof(ctx, address, storage, blockTag)
	})
}

func Fuzz_Nosy_EthClient_GetStorageAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *EthClient
		fill_err = tp.Fill(&s)
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
	if s == nil {
		return
	}

	s.GetStorageAt(ctx, address, storageSlot, blockTag)
	})
}

func Fuzz_Nosy_EthClient_InfoAndTxsByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *EthClient
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

	s.InfoAndTxsByHash(ctx, hash)
	})
}

func Fuzz_Nosy_EthClient_InfoAndTxsByLabel__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *EthClient
		fill_err = tp.Fill(&s)
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
	if s == nil {
		return
	}

	s.InfoAndTxsByLabel(ctx, label)
	})
}

func Fuzz_Nosy_EthClient_InfoAndTxsByNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *EthClient
		fill_err = tp.Fill(&s)
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
	if s == nil {
		return
	}

	s.InfoAndTxsByNumber(ctx, number)
	})
}

func Fuzz_Nosy_EthClient_InfoByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *EthClient
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

	s.InfoByHash(ctx, hash)
	})
}

func Fuzz_Nosy_EthClient_InfoByLabel__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *EthClient
		fill_err = tp.Fill(&s)
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
	if s == nil {
		return
	}

	s.InfoByLabel(ctx, label)
	})
}

func Fuzz_Nosy_EthClient_InfoByNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *EthClient
		fill_err = tp.Fill(&s)
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
	if s == nil {
		return
	}

	s.InfoByNumber(ctx, number)
	})
}

func Fuzz_Nosy_EthClient_PayloadByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *EthClient
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

	s.PayloadByHash(ctx, hash)
	})
}

func Fuzz_Nosy_EthClient_PayloadByLabel__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *EthClient
		fill_err = tp.Fill(&s)
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
	if s == nil {
		return
	}

	s.PayloadByLabel(ctx, label)
	})
}

func Fuzz_Nosy_EthClient_PayloadByNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *EthClient
		fill_err = tp.Fill(&s)
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
	if s == nil {
		return
	}

	s.PayloadByNumber(ctx, number)
	})
}

func Fuzz_Nosy_EthClient_ReadStorageAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *EthClient
		fill_err = tp.Fill(&s)
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
	if s == nil {
		return
	}

	s.ReadStorageAt(ctx, address, storageSlot, blockHash)
	})
}

// skipping Fuzz_Nosy_EthClient_SubscribeNewHead__ because parameters include func, chan, or unsupported interface: chan<- *github.com/ethereum/go-ethereum/core/types.Header

// skipping Fuzz_Nosy_EthClient_blockCall__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/sources.rpcBlockID

// skipping Fuzz_Nosy_EthClient_headerCall__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/sources.rpcBlockID

// skipping Fuzz_Nosy_EthClient_payloadCall__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/sources.rpcBlockID

func Fuzz_Nosy_EthClientConfig_Check__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *EthClientConfig
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Check()
	})
}

func Fuzz_Nosy_L1BeaconClient_GetBlobSidecars__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var cl *L1BeaconClient
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

func Fuzz_Nosy_L1BeaconClient_GetBlobs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var cl *L1BeaconClient
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

func Fuzz_Nosy_L1BeaconClient_GetTimeToSlotFn__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var cl *L1BeaconClient
		fill_err = tp.Fill(&cl)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
	if cl == nil {
		return
	}

	cl.GetTimeToSlotFn(ctx)
	})
}

func Fuzz_Nosy_L1BeaconClient_GetVersion__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var cl *L1BeaconClient
		fill_err = tp.Fill(&cl)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
	if cl == nil {
		return
	}

	cl.GetVersion(ctx)
	})
}

func Fuzz_Nosy_L1BeaconClient_fetchSidecars__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var cl *L1BeaconClient
		fill_err = tp.Fill(&cl)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var slot uint64
		fill_err = tp.Fill(&slot)
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

	cl.fetchSidecars(ctx, slot, hashes)
	})
}

func Fuzz_Nosy_L1Client_L1BlockRefByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *L1Client
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

	s.L1BlockRefByHash(ctx, hash)
	})
}

func Fuzz_Nosy_L1Client_L1BlockRefByLabel__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *L1Client
		fill_err = tp.Fill(&s)
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
	if s == nil {
		return
	}

	s.L1BlockRefByLabel(ctx, label)
	})
}

func Fuzz_Nosy_L1Client_L1BlockRefByNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *L1Client
		fill_err = tp.Fill(&s)
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
	if s == nil {
		return
	}

	s.L1BlockRefByNumber(ctx, num)
	})
}

func Fuzz_Nosy_L2Client_L2BlockRefByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *L2Client
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

	s.L2BlockRefByHash(ctx, hash)
	})
}

func Fuzz_Nosy_L2Client_L2BlockRefByLabel__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *L2Client
		fill_err = tp.Fill(&s)
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
	if s == nil {
		return
	}

	s.L2BlockRefByLabel(ctx, label)
	})
}

func Fuzz_Nosy_L2Client_L2BlockRefByNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *L2Client
		fill_err = tp.Fill(&s)
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
	if s == nil {
		return
	}

	s.L2BlockRefByNumber(ctx, num)
	})
}

func Fuzz_Nosy_L2Client_OutputV0AtBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *L2Client
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

	s.OutputV0AtBlock(ctx, blockHash)
	})
}

func Fuzz_Nosy_L2Client_RollupConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *L2Client
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
	if s == nil {
		return
	}

	s.RollupConfig()
	})
}

func Fuzz_Nosy_L2Client_SystemConfigByL2Hash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *L2Client
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

	s.SystemConfigByL2Hash(ctx, hash)
	})
}

func Fuzz_Nosy_RPCBlock_ExecutionPayloadEnvelope__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var block *RPCBlock
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		var trustCache bool
		fill_err = tp.Fill(&trustCache)
		if fill_err != nil {
			return
		}
	if block == nil {
		return
	}

	block.ExecutionPayloadEnvelope(trustCache)
	})
}

func Fuzz_Nosy_RPCBlock_Info__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var block *RPCBlock
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		var trustCache bool
		fill_err = tp.Fill(&trustCache)
		if fill_err != nil {
			return
		}
		var mustBePostMerge bool
		fill_err = tp.Fill(&mustBePostMerge)
		if fill_err != nil {
			return
		}
	if block == nil {
		return
	}

	block.Info(trustCache, mustBePostMerge)
	})
}

func Fuzz_Nosy_RPCBlock_verify__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var block *RPCBlock
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
	if block == nil {
		return
	}

	block.verify()
	})
}

func Fuzz_Nosy_RPCHeader_BlockID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var hdr *RPCHeader
		fill_err = tp.Fill(&hdr)
		if fill_err != nil {
			return
		}
	if hdr == nil {
		return
	}

	hdr.BlockID()
	})
}

func Fuzz_Nosy_RPCHeader_Info__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var hdr *RPCHeader
		fill_err = tp.Fill(&hdr)
		if fill_err != nil {
			return
		}
		var trustCache bool
		fill_err = tp.Fill(&trustCache)
		if fill_err != nil {
			return
		}
		var mustBePostMerge bool
		fill_err = tp.Fill(&mustBePostMerge)
		if fill_err != nil {
			return
		}
	if hdr == nil {
		return
	}

	hdr.Info(trustCache, mustBePostMerge)
	})
}

func Fuzz_Nosy_RPCHeader_checkPostMerge__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var hdr *RPCHeader
		fill_err = tp.Fill(&hdr)
		if fill_err != nil {
			return
		}
	if hdr == nil {
		return
	}

	hdr.checkPostMerge()
	})
}

func Fuzz_Nosy_RPCHeader_computeBlockHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var hdr *RPCHeader
		fill_err = tp.Fill(&hdr)
		if fill_err != nil {
			return
		}
	if hdr == nil {
		return
	}

	hdr.computeBlockHash()
	})
}

func Fuzz_Nosy_RPCHeader_createGethHeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var hdr *RPCHeader
		fill_err = tp.Fill(&hdr)
		if fill_err != nil {
			return
		}
	if hdr == nil {
		return
	}

	hdr.createGethHeader()
	})
}

func Fuzz_Nosy_RPCProviderKind_Clone__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var kind *RPCProviderKind
		fill_err = tp.Fill(&kind)
		if fill_err != nil {
			return
		}
	if kind == nil {
		return
	}

	kind.Clone()
	})
}

func Fuzz_Nosy_RPCProviderKind_Set__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var kind *RPCProviderKind
		fill_err = tp.Fill(&kind)
		if fill_err != nil {
			return
		}
		var value string
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
	if kind == nil {
		return
	}

	kind.Set(value)
	})
}

// skipping Fuzz_Nosy_RPCReceiptsFetcher_FetchReceipts__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/eth.BlockInfo

// skipping Fuzz_Nosy_RPCReceiptsFetcher_OnReceiptsMethodErr__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_RPCReceiptsFetcher_PickReceiptsMethod__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var f1 *RPCReceiptsFetcher
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var txCount int
		fill_err = tp.Fill(&txCount)
		if fill_err != nil {
			return
		}
	if f1 == nil {
		return
	}

	f1.PickReceiptsMethod(txCount)
	})
}

func Fuzz_Nosy_RollupClient_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var r *RollupClient
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
	if r == nil {
		return
	}

	r.Close()
	})
}

func Fuzz_Nosy_RollupClient_ConductorEnabled__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var r *RollupClient
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
	if r == nil {
		return
	}

	r.ConductorEnabled(ctx)
	})
}

func Fuzz_Nosy_RollupClient_OutputAtBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var r *RollupClient
		fill_err = tp.Fill(&r)
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
	if r == nil {
		return
	}

	r.OutputAtBlock(ctx, blockNum)
	})
}

func Fuzz_Nosy_RollupClient_OverrideLeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var r *RollupClient
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
	if r == nil {
		return
	}

	r.OverrideLeader(ctx)
	})
}

func Fuzz_Nosy_RollupClient_PostUnsafePayload__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var r *RollupClient
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var payload *eth.ExecutionPayloadEnvelope
		fill_err = tp.Fill(&payload)
		if fill_err != nil {
			return
		}
	if r == nil|| payload == nil {
		return
	}

	r.PostUnsafePayload(ctx, payload)
	})
}

func Fuzz_Nosy_RollupClient_RollupConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var r *RollupClient
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
	if r == nil {
		return
	}

	r.RollupConfig(ctx)
	})
}

func Fuzz_Nosy_RollupClient_SafeHeadAtL1Block__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var r *RollupClient
		fill_err = tp.Fill(&r)
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
	if r == nil {
		return
	}

	r.SafeHeadAtL1Block(ctx, blockNum)
	})
}

func Fuzz_Nosy_RollupClient_SequencerActive__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var r *RollupClient
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
	if r == nil {
		return
	}

	r.SequencerActive(ctx)
	})
}

func Fuzz_Nosy_RollupClient_SetLogLevel__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var r *RollupClient
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var lvl slog.Level
		fill_err = tp.Fill(&lvl)
		if fill_err != nil {
			return
		}
	if r == nil {
		return
	}

	r.SetLogLevel(ctx, lvl)
	})
}

func Fuzz_Nosy_RollupClient_StartSequencer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var r *RollupClient
		fill_err = tp.Fill(&r)
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
	if r == nil {
		return
	}

	r.StartSequencer(ctx, unsafeHead)
	})
}

func Fuzz_Nosy_RollupClient_StopSequencer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var r *RollupClient
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
	if r == nil {
		return
	}

	r.StopSequencer(ctx)
	})
}

func Fuzz_Nosy_RollupClient_SyncStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var r *RollupClient
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
	if r == nil {
		return
	}

	r.SyncStatus(ctx)
	})
}

func Fuzz_Nosy_RollupClient_Version__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var r *RollupClient
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
	if r == nil {
		return
	}

	r.Version(ctx)
	})
}

func Fuzz_Nosy_SupervisorClient_AddL2RPC__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var cl *SupervisorClient
		fill_err = tp.Fill(&cl)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var rpc string
		fill_err = tp.Fill(&rpc)
		if fill_err != nil {
			return
		}
	if cl == nil {
		return
	}

	cl.AddL2RPC(ctx, rpc)
	})
}

func Fuzz_Nosy_SupervisorClient_CheckMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var cl *SupervisorClient
		fill_err = tp.Fill(&cl)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var identifier types.Identifier
		fill_err = tp.Fill(&identifier)
		if fill_err != nil {
			return
		}
		var logHash common.Hash
		fill_err = tp.Fill(&logHash)
		if fill_err != nil {
			return
		}
	if cl == nil {
		return
	}

	cl.CheckMessage(ctx, identifier, logHash)
	})
}

func Fuzz_Nosy_SupervisorClient_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var cl *SupervisorClient
		fill_err = tp.Fill(&cl)
		if fill_err != nil {
			return
		}
	if cl == nil {
		return
	}

	cl.Close()
	})
}

func Fuzz_Nosy_SupervisorClient_CrossDerivedFrom__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var cl *SupervisorClient
		fill_err = tp.Fill(&cl)
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
	if cl == nil {
		return
	}

	cl.CrossDerivedFrom(ctx, chainID, derived)
	})
}

func Fuzz_Nosy_SupervisorClient_Finalized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var cl *SupervisorClient
		fill_err = tp.Fill(&cl)
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
	if cl == nil {
		return
	}

	cl.Finalized(ctx, chainID)
	})
}

func Fuzz_Nosy_SupervisorClient_SafeView__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var cl *SupervisorClient
		fill_err = tp.Fill(&cl)
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
	if cl == nil {
		return
	}

	cl.SafeView(ctx, chainID, safe)
	})
}

func Fuzz_Nosy_SupervisorClient_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var cl *SupervisorClient
		fill_err = tp.Fill(&cl)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
	if cl == nil {
		return
	}

	cl.Start(ctx)
	})
}

func Fuzz_Nosy_SupervisorClient_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var cl *SupervisorClient
		fill_err = tp.Fill(&cl)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
	if cl == nil {
		return
	}

	cl.Stop(ctx)
	})
}

func Fuzz_Nosy_SupervisorClient_UnsafeView__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var cl *SupervisorClient
		fill_err = tp.Fill(&cl)
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
	if cl == nil {
		return
	}

	cl.UnsafeView(ctx, chainID, unsafe)
	})
}

func Fuzz_Nosy_SupervisorClient_UpdateFinalizedL1__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var cl *SupervisorClient
		fill_err = tp.Fill(&cl)
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
		var finalizedL1 eth.L1BlockRef
		fill_err = tp.Fill(&finalizedL1)
		if fill_err != nil {
			return
		}
	if cl == nil {
		return
	}

	cl.UpdateFinalizedL1(ctx, chainID, finalizedL1)
	})
}

func Fuzz_Nosy_SupervisorClient_UpdateLocalSafe__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var cl *SupervisorClient
		fill_err = tp.Fill(&cl)
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
	if cl == nil {
		return
	}

	cl.UpdateLocalSafe(ctx, chainID, derivedFrom, lastDerived)
	})
}

func Fuzz_Nosy_SupervisorClient_UpdateLocalUnsafe__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var cl *SupervisorClient
		fill_err = tp.Fill(&cl)
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
	if cl == nil {
		return
	}

	cl.UpdateLocalUnsafe(ctx, chainID, head)
	})
}

func Fuzz_Nosy_limitClient_BatchCallContext__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var lc *limitClient
		fill_err = tp.Fill(&lc)
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
	if lc == nil {
		return
	}

	lc.BatchCallContext(ctx, b)
	})
}

// skipping Fuzz_Nosy_limitClient_CallContext__ because parameters include func, chan, or unsupported interface: any

func Fuzz_Nosy_limitClient_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var lc *limitClient
		fill_err = tp.Fill(&lc)
		if fill_err != nil {
			return
		}
	if lc == nil {
		return
	}

	lc.Close()
	})
}

// skipping Fuzz_Nosy_limitClient_EthSubscribe__ because parameters include func, chan, or unsupported interface: any

func Fuzz_Nosy_limitClient_joinWaitGroup__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var lc *limitClient
		fill_err = tp.Fill(&lc)
		if fill_err != nil {
			return
		}
	if lc == nil {
		return
	}

	lc.joinWaitGroup()
	})
}

// skipping Fuzz_Nosy_BeaconClient_BeaconBlobSideCars__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/sources.BeaconClient

// skipping Fuzz_Nosy_BeaconClient_BeaconGenesis__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/sources.BeaconClient

// skipping Fuzz_Nosy_BeaconClient_ConfigSpec__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/sources.BeaconClient

// skipping Fuzz_Nosy_BeaconClient_NodeVersion__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/sources.BeaconClient

// skipping Fuzz_Nosy_BlobSideCarsFetcher_BeaconBlobSideCars__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/sources.BlobSideCarsFetcher

// skipping Fuzz_Nosy_EngineVersionProvider_ForkchoiceUpdatedVersion__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/sources.EngineVersionProvider

// skipping Fuzz_Nosy_EngineVersionProvider_GetPayloadVersion__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/sources.EngineVersionProvider

// skipping Fuzz_Nosy_EngineVersionProvider_NewPayloadVersion__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/sources.EngineVersionProvider

func Fuzz_Nosy_RPCProviderKind_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var kind RPCProviderKind
		fill_err = tp.Fill(&kind)
		if fill_err != nil {
			return
		}

	kind.String()
	})
}

func Fuzz_Nosy_ReceiptsFetchingMethod_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var kind RPCProviderKind
		fill_err = tp.Fill(&kind)
		if fill_err != nil {
			return
		}
		var available ReceiptsFetchingMethod
		fill_err = tp.Fill(&available)
		if fill_err != nil {
			return
		}
		var txCount uint64
		fill_err = tp.Fill(&txCount)
		if fill_err != nil {
			return
		}

	r := PickBestReceiptsFetchingMethod(kind, available, txCount)
	r.String()
	})
}

// skipping Fuzz_Nosy_ReceiptsProvider_FetchReceipts__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/sources.ReceiptsProvider

func Fuzz_Nosy_hashID_Arg__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var h hashID
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}

	h.Arg()
	})
}

func Fuzz_Nosy_hashID_CheckID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var h hashID
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var id eth.BlockID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}

	h.CheckID(id)
	})
}

func Fuzz_Nosy_headerInfo_BaseFee__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var h headerInfo
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}

	h.BaseFee()
	})
}

func Fuzz_Nosy_headerInfo_BlobBaseFee__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var h headerInfo
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}

	h.BlobBaseFee()
	})
}

func Fuzz_Nosy_headerInfo_Coinbase__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var h headerInfo
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}

	h.Coinbase()
	})
}

func Fuzz_Nosy_headerInfo_GasLimit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var h headerInfo
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}

	h.GasLimit()
	})
}

func Fuzz_Nosy_headerInfo_GasUsed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var h headerInfo
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}

	h.GasUsed()
	})
}

func Fuzz_Nosy_headerInfo_Hash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var h headerInfo
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}

	h.Hash()
	})
}

func Fuzz_Nosy_headerInfo_HeaderRLP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var h headerInfo
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}

	h.HeaderRLP()
	})
}

func Fuzz_Nosy_headerInfo_MixDigest__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var h headerInfo
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}

	h.MixDigest()
	})
}

func Fuzz_Nosy_headerInfo_NumberU64__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var h headerInfo
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}

	h.NumberU64()
	})
}

func Fuzz_Nosy_headerInfo_ParentBeaconRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var h headerInfo
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}

	h.ParentBeaconRoot()
	})
}

func Fuzz_Nosy_headerInfo_ParentHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var h headerInfo
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}

	h.ParentHash()
	})
}

func Fuzz_Nosy_headerInfo_ReceiptHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var h headerInfo
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}

	h.ReceiptHash()
	})
}

func Fuzz_Nosy_headerInfo_Root__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var h headerInfo
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}

	h.Root()
	})
}

func Fuzz_Nosy_headerInfo_Time__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var h headerInfo
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}

	h.Time()
	})
}

func Fuzz_Nosy_numberID_Arg__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var n numberID
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}

	n.Arg()
	})
}

func Fuzz_Nosy_numberID_CheckID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var n numberID
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var id eth.BlockID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}

	n.CheckID(id)
	})
}

// skipping Fuzz_Nosy_rpcBlockID_Arg__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/sources.rpcBlockID

// skipping Fuzz_Nosy_rpcBlockID_CheckID__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/sources.rpcBlockID

// skipping Fuzz_Nosy_rpcClient_BatchCallContext__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/sources.rpcClient

// skipping Fuzz_Nosy_rpcClient_CallContext__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/sources.rpcClient

func Fuzz_Nosy_ValidRPCProviderKind__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var value RPCProviderKind
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}

	ValidRPCProviderKind(value)
	})
}

func Fuzz_Nosy_blobsFromSidecars__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var blobSidecars []*eth.BlobSidecar
		fill_err = tp.Fill(&blobSidecars)
		if fill_err != nil {
			return
		}
		var hashes []eth.IndexedBlobHash
		fill_err = tp.Fill(&hashes)
		if fill_err != nil {
			return
		}

	blobsFromSidecars(blobSidecars, hashes)
	})
}

func Fuzz_Nosy_makeReceiptRequest__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var txHash common.Hash
		fill_err = tp.Fill(&txHash)
		if fill_err != nil {
			return
		}

	makeReceiptRequest(txHash)
	})
}

// skipping Fuzz_Nosy_unusableMethod__ because parameters include func, chan, or unsupported interface: error

