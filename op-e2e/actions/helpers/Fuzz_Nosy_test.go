package helpers

import (
	"testing"
	"github.com/ethereum-optimism/optimism/op-e2e/actions"
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
	
	func Fuzz_Nosy_BasicUser[B any]_ActCheckReceiptStatusOfLastTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *BasicUser[B]
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var success bool
		fill_err = tp.Fill(&success)
		if fill_err != nil {
			return
		}
	if s == nil {
		return
	}

	s.ActCheckReceiptStatusOfLastTx(success)
	})
}

// skipping Fuzz_Nosy_BasicUser[B any]_ActMakeTx__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

// skipping Fuzz_Nosy_BasicUser[B any]_ActRandomTxData__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

// skipping Fuzz_Nosy_BasicUser[B any]_ActRandomTxToAddr__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

// skipping Fuzz_Nosy_BasicUser[B any]_ActRandomTxValue__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

// skipping Fuzz_Nosy_BasicUser[B any]_ActResetTxOpts__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

func Fuzz_Nosy_BasicUser[B any]_ActSetTxCalldata__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *BasicUser[B]
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var calldata []byte
		fill_err = tp.Fill(&calldata)
		if fill_err != nil {
			return
		}
	if s == nil {
		return
	}

	s.ActSetTxCalldata(calldata)
	})
}

func Fuzz_Nosy_BasicUser[B any]_ActSetTxToAddr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *BasicUser[B]
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var to *common.Address
		fill_err = tp.Fill(&to)
		if fill_err != nil {
			return
		}
	if s == nil|| to == nil {
		return
	}

	s.ActSetTxToAddr(to)
	})
}

func Fuzz_Nosy_BasicUser[B any]_ActSetTxValue__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *BasicUser[B]
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var value *big.Int
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
	if s == nil|| value == nil {
		return
	}

	s.ActSetTxValue(value)
	})
}

// skipping Fuzz_Nosy_BasicUser[B any]_CheckReceipt__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

// skipping Fuzz_Nosy_BasicUser[B any]_LastTxReceipt__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

// skipping Fuzz_Nosy_BasicUser[B any]_MakeTransaction__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

// skipping Fuzz_Nosy_BasicUser[B any]_PendingNonce__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

func Fuzz_Nosy_BasicUser[B any]_SetUserEnv__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *BasicUser[B]
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var env *BasicUserEnv[B]
		fill_err = tp.Fill(&env)
		if fill_err != nil {
			return
		}
	if s == nil|| env == nil {
		return
	}

	s.SetUserEnv(env)
	})
}

func Fuzz_Nosy_BasicUser[B any]_Signer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *BasicUser[B]
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
	if s == nil {
		return
	}

	s.Signer()
	})
}

func Fuzz_Nosy_BasicUser[B any]_TxValue__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *BasicUser[B]
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
	if s == nil {
		return
	}

	s.TxValue()
	})
}

func Fuzz_Nosy_BasicUser[B any]_signerFn__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *BasicUser[B]
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		var tx *types.Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
	if s == nil|| tx == nil {
		return
	}

	s.signerFn(address, tx)
	})
}

func Fuzz_Nosy_CrossLayerUser_ActCheckDepositStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *CrossLayerUser
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var l1Success bool
		fill_err = tp.Fill(&l1Success)
		if fill_err != nil {
			return
		}
		var l2Success bool
		fill_err = tp.Fill(&l2Success)
		if fill_err != nil {
			return
		}
	if s == nil {
		return
	}

	s.ActCheckDepositStatus(l1Success, l2Success)
	})
}

func Fuzz_Nosy_CrossLayerUser_ActCheckStartWithdrawal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *CrossLayerUser
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var success bool
		fill_err = tp.Fill(&success)
		if fill_err != nil {
			return
		}
	if s == nil {
		return
	}

	s.ActCheckStartWithdrawal(success)
	})
}

// skipping Fuzz_Nosy_CrossLayerUser_ActCompleteWithdrawal__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

// skipping Fuzz_Nosy_CrossLayerUser_ActDeposit__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

// skipping Fuzz_Nosy_CrossLayerUser_ActProveWithdrawal__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

// skipping Fuzz_Nosy_CrossLayerUser_ActResolve__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

// skipping Fuzz_Nosy_CrossLayerUser_ActResolveClaim__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

// skipping Fuzz_Nosy_CrossLayerUser_ActStartWithdrawal__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

func Fuzz_Nosy_CrossLayerUser_Address__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *CrossLayerUser
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
	if s == nil {
		return
	}

	s.Address()
	})
}

// skipping Fuzz_Nosy_CrossLayerUser_CheckDepositTx__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

// skipping Fuzz_Nosy_CrossLayerUser_CompleteWithdrawal__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

// skipping Fuzz_Nosy_CrossLayerUser_ProveWithdrawal__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

// skipping Fuzz_Nosy_CrossLayerUser_Resolve__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

// skipping Fuzz_Nosy_CrossLayerUser_ResolveClaim__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

// skipping Fuzz_Nosy_CrossLayerUser_getDisputeGame__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

// skipping Fuzz_Nosy_CrossLayerUser_getLatestWithdrawalParams__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

func Fuzz_Nosy_GarbageChannelOut_AddBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var cfg *GarbageChannelCfg
		fill_err = tp.Fill(&cfg)
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
	if cfg == nil|| rollupCfg == nil|| block == nil {
		return
	}

	co, err := NewGarbageChannelOut(cfg)
	if err != nil {
		return
	}
	co.AddBlock(rollupCfg, block)
	})
}

func Fuzz_Nosy_GarbageChannelOut_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var cfg *GarbageChannelCfg
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
	if cfg == nil {
		return
	}

	co, err := NewGarbageChannelOut(cfg)
	if err != nil {
		return
	}
	co.Close()
	})
}

func Fuzz_Nosy_GarbageChannelOut_Flush__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var cfg *GarbageChannelCfg
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
	if cfg == nil {
		return
	}

	co, err := NewGarbageChannelOut(cfg)
	if err != nil {
		return
	}
	co.Flush()
	})
}

func Fuzz_Nosy_GarbageChannelOut_ID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var cfg *GarbageChannelCfg
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
	if cfg == nil {
		return
	}

	co, err := NewGarbageChannelOut(cfg)
	if err != nil {
		return
	}
	co.ID()
	})
}

func Fuzz_Nosy_GarbageChannelOut_OutputFrame__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var cfg *GarbageChannelCfg
		fill_err = tp.Fill(&cfg)
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
	if cfg == nil|| w == nil {
		return
	}

	co, err := NewGarbageChannelOut(cfg)
	if err != nil {
		return
	}
	co.OutputFrame(w, maxSize)
	})
}

func Fuzz_Nosy_GarbageChannelOut_ReadyBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var cfg *GarbageChannelCfg
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
	if cfg == nil {
		return
	}

	co, err := NewGarbageChannelOut(cfg)
	if err != nil {
		return
	}
	co.ReadyBytes()
	})
}

func Fuzz_Nosy_GarbageChannelOut_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var cfg *GarbageChannelCfg
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
	if cfg == nil {
		return
	}

	co, err := NewGarbageChannelOut(cfg)
	if err != nil {
		return
	}
	co.Reset()
	})
}

// skipping Fuzz_Nosy_L1Miner_ActEmptyBlock__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

// skipping Fuzz_Nosy_L1Miner_ActL1EndBlock__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

func Fuzz_Nosy_L1Miner_ActL1IncludeTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *L1Miner
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var from common.Address
		fill_err = tp.Fill(&from)
		if fill_err != nil {
			return
		}
	if s == nil {
		return
	}

	s.ActL1IncludeTx(from)
	})
}

func Fuzz_Nosy_L1Miner_ActL1IncludeTxByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *L1Miner
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var txHash common.Hash
		fill_err = tp.Fill(&txHash)
		if fill_err != nil {
			return
		}
	if s == nil {
		return
	}

	s.ActL1IncludeTxByHash(txHash)
	})
}

func Fuzz_Nosy_L1Miner_ActL1SetFeeRecipient__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *L1Miner
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var coinbase common.Address
		fill_err = tp.Fill(&coinbase)
		if fill_err != nil {
			return
		}
	if s == nil {
		return
	}

	s.ActL1SetFeeRecipient(coinbase)
	})
}

func Fuzz_Nosy_L1Miner_ActL1StartBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *L1Miner
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var timeDelta uint64
		fill_err = tp.Fill(&timeDelta)
		if fill_err != nil {
			return
		}
	if s == nil {
		return
	}

	s.ActL1StartBlock(timeDelta)
	})
}

func Fuzz_Nosy_L1Miner_BlobSource__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *L1Miner
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
	if s == nil {
		return
	}

	s.BlobSource()
	})
}

func Fuzz_Nosy_L1Miner_BlobStore__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *L1Miner
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
	if s == nil {
		return
	}

	s.BlobStore()
	})
}

func Fuzz_Nosy_L1Miner_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *L1Miner
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

// skipping Fuzz_Nosy_L1Miner_IncludeTx__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

// skipping Fuzz_Nosy_L1Replica_ActL1Finalize__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

// skipping Fuzz_Nosy_L1Replica_ActL1FinalizeNext__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

// skipping Fuzz_Nosy_L1Replica_ActL1RPCFail__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

func Fuzz_Nosy_L1Replica_ActL1RewindDepth__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *L1Replica
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var depth uint64
		fill_err = tp.Fill(&depth)
		if fill_err != nil {
			return
		}
	if s == nil {
		return
	}

	s.ActL1RewindDepth(depth)
	})
}

// skipping Fuzz_Nosy_L1Replica_ActL1RewindToParent__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

// skipping Fuzz_Nosy_L1Replica_ActL1Safe__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

// skipping Fuzz_Nosy_L1Replica_ActL1SafeNext__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

// skipping Fuzz_Nosy_L1Replica_ActL1Sync__ because parameters include func, chan, or unsupported interface: func(num uint64) *github.com/ethereum/go-ethereum/core/types.Block

func Fuzz_Nosy_L1Replica_CanonL1Chain__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *L1Replica
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
	if s == nil {
		return
	}

	s.CanonL1Chain()
	})
}

func Fuzz_Nosy_L1Replica_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *L1Replica
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

func Fuzz_Nosy_L1Replica_EthClient__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *L1Replica
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
	if s == nil {
		return
	}

	s.EthClient()
	})
}

func Fuzz_Nosy_L1Replica_FinalizedNum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *L1Replica
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
	if s == nil {
		return
	}

	s.FinalizedNum()
	})
}

func Fuzz_Nosy_L1Replica_HTTPEndpoint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *L1Replica
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
	if s == nil {
		return
	}

	s.HTTPEndpoint()
	})
}

func Fuzz_Nosy_L1Replica_L1Chain__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *L1Replica
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
	if s == nil {
		return
	}

	s.L1Chain()
	})
}

// skipping Fuzz_Nosy_L1Replica_L1Client__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

// skipping Fuzz_Nosy_L1Replica_MockL1RPCErrors__ because parameters include func, chan, or unsupported interface: func() error

func Fuzz_Nosy_L1Replica_RPCClient__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *L1Replica
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
	if s == nil {
		return
	}

	s.RPCClient()
	})
}

func Fuzz_Nosy_L1Replica_SafeNum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *L1Replica
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
	if s == nil {
		return
	}

	s.SafeNum()
	})
}

func Fuzz_Nosy_L1Replica_UnsafeNum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *L1Replica
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
	if s == nil {
		return
	}

	s.UnsafeNum()
	})
}

// skipping Fuzz_Nosy_L2Batcher_ActAddBlockByNumber__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

// skipping Fuzz_Nosy_L2Batcher_ActBufferAll__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

// skipping Fuzz_Nosy_L2Batcher_ActCreateChannel__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

// skipping Fuzz_Nosy_L2Batcher_ActL2BatchBuffer__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

// skipping Fuzz_Nosy_L2Batcher_ActL2BatchSubmit__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

// skipping Fuzz_Nosy_L2Batcher_ActL2BatchSubmitGarbage__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

// skipping Fuzz_Nosy_L2Batcher_ActL2BatchSubmitGarbageRaw__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

// skipping Fuzz_Nosy_L2Batcher_ActL2BatchSubmitMultiBlob__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

// skipping Fuzz_Nosy_L2Batcher_ActL2BatchSubmitRaw__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

// skipping Fuzz_Nosy_L2Batcher_ActL2ChannelClose__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

// skipping Fuzz_Nosy_L2Batcher_ActSubmitAll__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

// skipping Fuzz_Nosy_L2Batcher_ActSubmitAllMultiBlobs__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

// skipping Fuzz_Nosy_L2Batcher_Buffer__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

// skipping Fuzz_Nosy_L2Batcher_ReadNextOutputFrame__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

func Fuzz_Nosy_L2Batcher_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *L2Batcher
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
	if s == nil {
		return
	}

	s.Reset()
	})
}

func Fuzz_Nosy_L2Batcher_SubmittingData__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *L2Batcher
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
	if s == nil {
		return
	}

	s.SubmittingData()
	})
}

func Fuzz_Nosy_L2Engine_ActL2IncludeTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var e *L2Engine
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var from common.Address
		fill_err = tp.Fill(&from)
		if fill_err != nil {
			return
		}
	if e == nil {
		return
	}

	e.ActL2IncludeTx(from)
	})
}

func Fuzz_Nosy_L2Engine_ActL2IncludeTxIgnoreForcedEmpty__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var e *L2Engine
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var from common.Address
		fill_err = tp.Fill(&from)
		if fill_err != nil {
			return
		}
	if e == nil {
		return
	}

	e.ActL2IncludeTxIgnoreForcedEmpty(from)
	})
}

// skipping Fuzz_Nosy_L2Engine_ActL2RPCFail__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

func Fuzz_Nosy_L2Engine_AddPeers__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *L2Engine
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var peers []*enode.Node
		fill_err = tp.Fill(&peers)
		if fill_err != nil {
			return
		}
	if s == nil {
		return
	}

	s.AddPeers(peers...)
	})
}

func Fuzz_Nosy_L2Engine_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var e *L2Engine
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
	if e == nil {
		return
	}

	e.Close()
	})
}

// skipping Fuzz_Nosy_L2Engine_EngineClient__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

func Fuzz_Nosy_L2Engine_Enode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *L2Engine
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
	if s == nil {
		return
	}

	s.Enode()
	})
}

func Fuzz_Nosy_L2Engine_EthClient__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *L2Engine
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
	if s == nil {
		return
	}

	s.EthClient()
	})
}

func Fuzz_Nosy_L2Engine_GethClient__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *L2Engine
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
	if s == nil {
		return
	}

	s.GethClient()
	})
}

func Fuzz_Nosy_L2Engine_HTTPEndpoint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *L2Engine
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
	if s == nil {
		return
	}

	s.HTTPEndpoint()
	})
}

func Fuzz_Nosy_L2Engine_L2Chain__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *L2Engine
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
	if s == nil {
		return
	}

	s.L2Chain()
	})
}

func Fuzz_Nosy_L2Engine_PeerCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *L2Engine
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
	if s == nil {
		return
	}

	s.PeerCount()
	})
}

func Fuzz_Nosy_L2Engine_RPCClient__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var e *L2Engine
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
	if e == nil {
		return
	}

	e.RPCClient()
	})
}

// skipping Fuzz_Nosy_L2Proposer_ActMakeProposalTx__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

// skipping Fuzz_Nosy_L2Proposer_CanPropose__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

func Fuzz_Nosy_L2Proposer_LastProposalTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var p *L2Proposer
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
	if p == nil {
		return
	}

	p.LastProposalTx()
	})
}

// skipping Fuzz_Nosy_L2Proposer_fetchNextOutput__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

// skipping Fuzz_Nosy_L2Proposer_sendTx__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

// skipping Fuzz_Nosy_L2Sequencer_ActBuildL2ToEcotone__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

// skipping Fuzz_Nosy_L2Sequencer_ActBuildL2ToFjord__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

// skipping Fuzz_Nosy_L2Sequencer_ActBuildL2ToGranite__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

// skipping Fuzz_Nosy_L2Sequencer_ActBuildL2ToHolocene__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

// skipping Fuzz_Nosy_L2Sequencer_ActBuildL2ToTime__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

// skipping Fuzz_Nosy_L2Sequencer_ActBuildToL1Head__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

// skipping Fuzz_Nosy_L2Sequencer_ActBuildToL1HeadExcl__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

// skipping Fuzz_Nosy_L2Sequencer_ActBuildToL1HeadExclUnsafe__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

// skipping Fuzz_Nosy_L2Sequencer_ActBuildToL1HeadUnsafe__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

// skipping Fuzz_Nosy_L2Sequencer_ActL2EmptyBlock__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

// skipping Fuzz_Nosy_L2Sequencer_ActL2EndBlock__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

// skipping Fuzz_Nosy_L2Sequencer_ActL2ForceAdvanceL1Origin__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

// skipping Fuzz_Nosy_L2Sequencer_ActL2KeepL1Origin__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

// skipping Fuzz_Nosy_L2Sequencer_ActL2StartBlock__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

// skipping Fuzz_Nosy_L2Verifier_ActInteropBackendCheck__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

// skipping Fuzz_Nosy_L2Verifier_ActL1FinalizedSignal__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

// skipping Fuzz_Nosy_L2Verifier_ActL1HeadSignal__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

// skipping Fuzz_Nosy_L2Verifier_ActL1SafeSignal__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

// skipping Fuzz_Nosy_L2Verifier_ActL2EventsUntil__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

// skipping Fuzz_Nosy_L2Verifier_ActL2EventsUntilPending__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

func Fuzz_Nosy_L2Verifier_ActL2InsertUnsafePayload__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *L2Verifier
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var payload *eth.ExecutionPayloadEnvelope
		fill_err = tp.Fill(&payload)
		if fill_err != nil {
			return
		}
	if s == nil|| payload == nil {
		return
	}

	s.ActL2InsertUnsafePayload(payload)
	})
}

// skipping Fuzz_Nosy_L2Verifier_ActL2PipelineFull__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

func Fuzz_Nosy_L2Verifier_ActL2UnsafeGossipReceive__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *L2Verifier
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var payload *eth.ExecutionPayloadEnvelope
		fill_err = tp.Fill(&payload)
		if fill_err != nil {
			return
		}
	if s == nil|| payload == nil {
		return
	}

	s.ActL2UnsafeGossipReceive(payload)
	})
}

// skipping Fuzz_Nosy_L2Verifier_ActRPCFail__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

func Fuzz_Nosy_L2Verifier_DerivationMetricsTracer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *L2Verifier
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
	if s == nil {
		return
	}

	s.DerivationMetricsTracer()
	})
}

func Fuzz_Nosy_L2Verifier_L2BackupUnsafe__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *L2Verifier
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
	if s == nil {
		return
	}

	s.L2BackupUnsafe()
	})
}

func Fuzz_Nosy_L2Verifier_L2Finalized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *L2Verifier
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
	if s == nil {
		return
	}

	s.L2Finalized()
	})
}

func Fuzz_Nosy_L2Verifier_L2PendingSafe__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *L2Verifier
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
	if s == nil {
		return
	}

	s.L2PendingSafe()
	})
}

func Fuzz_Nosy_L2Verifier_L2Safe__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *L2Verifier
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
	if s == nil {
		return
	}

	s.L2Safe()
	})
}

func Fuzz_Nosy_L2Verifier_L2Unsafe__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *L2Verifier
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
	if s == nil {
		return
	}

	s.L2Unsafe()
	})
}

// skipping Fuzz_Nosy_L2Verifier_OnEvent__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/event.Event

func Fuzz_Nosy_L2Verifier_RPCClient__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *L2Verifier
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
	if s == nil {
		return
	}

	s.RPCClient()
	})
}

func Fuzz_Nosy_L2Verifier_RollupClient__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *L2Verifier
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
	if s == nil {
		return
	}

	s.RollupClient()
	})
}

func Fuzz_Nosy_L2Verifier_SyncStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *L2Verifier
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
	if s == nil {
		return
	}

	s.SyncStatus()
	})
}

func Fuzz_Nosy_MockL1OriginSelector_FindL1Origin__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var m *MockL1OriginSelector
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var l2Head eth.L2BlockRef
		fill_err = tp.Fill(&l2Head)
		if fill_err != nil {
			return
		}
	if m == nil {
		return
	}

	m.FindL1Origin(ctx, l2Head)
	})
}

func Fuzz_Nosy_defaultTesting_Ctx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var st *defaultTesting
		fill_err = tp.Fill(&st)
		if fill_err != nil {
			return
		}
	if st == nil {
		return
	}

	st.Ctx()
	})
}

// skipping Fuzz_Nosy_defaultTesting_InvalidAction__ because parameters include func, chan, or unsupported interface: []any

func Fuzz_Nosy_defaultTesting_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var st *defaultTesting
		fill_err = tp.Fill(&st)
		if fill_err != nil {
			return
		}
		var actionCtx context.Context
		fill_err = tp.Fill(&actionCtx)
		if fill_err != nil {
			return
		}
	if st == nil {
		return
	}

	st.Reset(actionCtx)
	})
}

func Fuzz_Nosy_defaultTesting_State__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var st *defaultTesting
		fill_err = tp.Fill(&st)
		if fill_err != nil {
			return
		}
	if st == nil {
		return
	}

	st.State()
	})
}

func Fuzz_Nosy_engineApiBackend_Database__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var e *engineApiBackend
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
	if e == nil {
		return
	}

	e.Database()
	})
}

func Fuzz_Nosy_engineApiBackend_Genesis__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var e *engineApiBackend
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
	if e == nil {
		return
	}

	e.Genesis()
	})
}

func Fuzz_Nosy_l2VerifierBackend_BlockRefWithStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *l2VerifierBackend
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

	s.BlockRefWithStatus(ctx, num)
	})
}

func Fuzz_Nosy_l2VerifierBackend_ConductorEnabled__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *l2VerifierBackend
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

	s.ConductorEnabled(ctx)
	})
}

func Fuzz_Nosy_l2VerifierBackend_OnUnsafeL2Payload__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *l2VerifierBackend
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var envelope *eth.ExecutionPayloadEnvelope
		fill_err = tp.Fill(&envelope)
		if fill_err != nil {
			return
		}
	if s == nil|| envelope == nil {
		return
	}

	s.OnUnsafeL2Payload(ctx, envelope)
	})
}

func Fuzz_Nosy_l2VerifierBackend_OverrideLeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *l2VerifierBackend
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

	s.OverrideLeader(ctx)
	})
}

func Fuzz_Nosy_l2VerifierBackend_ResetDerivationPipeline__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *l2VerifierBackend
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

	s.ResetDerivationPipeline(ctx)
	})
}

func Fuzz_Nosy_l2VerifierBackend_SequencerActive__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *l2VerifierBackend
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

	s.SequencerActive(ctx)
	})
}

func Fuzz_Nosy_l2VerifierBackend_StartSequencer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *l2VerifierBackend
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

	s.StartSequencer(ctx, blockHash)
	})
}

func Fuzz_Nosy_l2VerifierBackend_StopSequencer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *l2VerifierBackend
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

	s.StopSequencer(ctx)
	})
}

func Fuzz_Nosy_l2VerifierBackend_SyncStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *l2VerifierBackend
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

	s.SyncStatus(ctx)
	})
}

// skipping Fuzz_Nosy_AltDAInputSetter_SetInput__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.AltDAInputSetter

// skipping Fuzz_Nosy_BlocksAPI_BlockByNumber__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.BlocksAPI

// skipping Fuzz_Nosy_ChannelOutIface_AddBlock__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.ChannelOutIface

// skipping Fuzz_Nosy_ChannelOutIface_Close__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.ChannelOutIface

// skipping Fuzz_Nosy_ChannelOutIface_Flush__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.ChannelOutIface

// skipping Fuzz_Nosy_ChannelOutIface_ID__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.ChannelOutIface

// skipping Fuzz_Nosy_ChannelOutIface_OutputFrame__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.ChannelOutIface

// skipping Fuzz_Nosy_ChannelOutIface_ReadyBytes__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.ChannelOutIface

// skipping Fuzz_Nosy_ChannelOutIface_Reset__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.ChannelOutIface

// skipping Fuzz_Nosy_Env_ActBatchSubmitAllAndMine__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

// skipping Fuzz_Nosy_Env_L1UserEnv__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

// skipping Fuzz_Nosy_Env_L2UserEnv__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

func Fuzz_Nosy_GarbageKind_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var gk GarbageKind
		fill_err = tp.Fill(&gk)
		if fill_err != nil {
			return
		}

	gk.String()
	})
}

// skipping Fuzz_Nosy_L1TxAPI_HeaderByNumber__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.L1TxAPI

// skipping Fuzz_Nosy_L1TxAPI_PendingNonceAt__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.L1TxAPI

// skipping Fuzz_Nosy_L1TxAPI_SendTransaction__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.L1TxAPI

// skipping Fuzz_Nosy_L2API_GetProof__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.L2API

// skipping Fuzz_Nosy_L2API_InfoByHash__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.L2API

// skipping Fuzz_Nosy_L2API_L2BlockRefByNumber__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.L2API

// skipping Fuzz_Nosy_L2API_OutputV0AtBlock__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.L2API

// skipping Fuzz_Nosy_L2BlockRefs_L2BlockRefByHash__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.L2BlockRefs

// skipping Fuzz_Nosy_StatefulTesting_Reset__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.StatefulTesting

// skipping Fuzz_Nosy_StatefulTesting_State__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.StatefulTesting

// skipping Fuzz_Nosy_SyncStatusAPI_SyncStatus__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.SyncStatusAPI

// skipping Fuzz_Nosy_Testing_Ctx__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

// skipping Fuzz_Nosy_Testing_InvalidAction__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

// skipping Fuzz_Nosy_Writer_Close__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Writer

// skipping Fuzz_Nosy_Writer_Flush__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Writer

// skipping Fuzz_Nosy_Writer_Reset__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Writer

// skipping Fuzz_Nosy_Writer_Write__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Writer

func Fuzz_Nosy_fakeTxMgr_API__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var f1 fakeTxMgr
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}

	f1.API()
	})
}

func Fuzz_Nosy_fakeTxMgr_BlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var f1 fakeTxMgr
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}

	f1.BlockNumber(_x2)
	})
}

func Fuzz_Nosy_fakeTxMgr_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var f1 fakeTxMgr
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}

	f1.Close()
	})
}

func Fuzz_Nosy_fakeTxMgr_From__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var f1 fakeTxMgr
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}

	f1.From()
	})
}

func Fuzz_Nosy_fakeTxMgr_IsClosed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var f1 fakeTxMgr
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}

	f1.IsClosed()
	})
}

func Fuzz_Nosy_fakeTxMgr_Send__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var f1 fakeTxMgr
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var _x3 txmgr.TxCandidate
		fill_err = tp.Fill(&_x3)
		if fill_err != nil {
			return
		}

	f1.Send(_x2, _x3)
	})
}

// skipping Fuzz_Nosy_fakeTxMgr_SendAsync__ because parameters include func, chan, or unsupported interface: chan github.com/ethereum-optimism/optimism/op-service/txmgr.SendResponse

func Fuzz_Nosy_fakeTxMgr_SuggestGasPriceCaps__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var f1 fakeTxMgr
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}

	f1.SuggestGasPriceCaps(_x2)
	})
}

// skipping Fuzz_Nosy_BlockLogger__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/e2eutils.TestingBase

// skipping Fuzz_Nosy_SetupReorgTest__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

// skipping Fuzz_Nosy_SetupReorgTestActors__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

// skipping Fuzz_Nosy_SetupSequencerTest__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

// skipping Fuzz_Nosy_SetupVerifier__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

// skipping Fuzz_Nosy_SetupVerifierOnlyTest__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

func Fuzz_Nosy_blockToBatch__(f *testing.F) {
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
	if rollupCfg == nil|| block == nil {
		return
	}

	blockToBatch(rollupCfg, block)
	})
}

func Fuzz_Nosy_estimateGasPending__(f *testing.F) {
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
		var ec *ethclient.Client
		fill_err = tp.Fill(&ec)
		if fill_err != nil {
			return
		}
		var msg ethereum.CallMsg
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
	if ec == nil {
		return
	}

	estimateGasPending(ctx, ec, msg)
	})
}

// skipping Fuzz_Nosy_newBackend__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/e2eutils.TestingBase

func Fuzz_Nosy_toCallArg__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var msg ethereum.CallMsg
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}

	toCallArg(msg)
	})
}

