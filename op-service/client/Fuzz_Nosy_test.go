package client

import (
	"context"
	"math/big"
	"net/http"
	"net/url"
	"testing"

	"github.com/ethereum-optimism/optimism/op-service/metrics"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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

func Fuzz_Nosy_BaseRPCClient_BatchCallContext__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BaseRPCClient
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var batch []rpc.BatchElem
		fill_err = tp.Fill(&batch)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.BatchCallContext(ctx, batch)
	})
}

// skipping Fuzz_Nosy_BaseRPCClient_CallContext__ because parameters include func, chan, or unsupported interface: any

func Fuzz_Nosy_BaseRPCClient_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BaseRPCClient
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.Close()
	})
}

// skipping Fuzz_Nosy_BaseRPCClient_EthSubscribe__ because parameters include func, chan, or unsupported interface: any

func Fuzz_Nosy_BasicHTTPClient_Get__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cl *BasicHTTPClient
		fill_err = tp.Fill(&cl)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var p string
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var query url.Values
		fill_err = tp.Fill(&query)
		if fill_err != nil {
			return
		}
		var headers http.Header
		fill_err = tp.Fill(&headers)
		if fill_err != nil {
			return
		}
		if cl == nil {
			return
		}

		cl.Get(ctx, p, query, headers)
	})
}

func Fuzz_Nosy_InstrumentedClient_BalanceAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var m *metrics.RPCClientMetrics
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
		var blockNumber *big.Int
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		if c == nil || m == nil || blockNumber == nil {
			return
		}

		ic := NewInstrumentedClient(c, m)
		ic.BalanceAt(ctx, account, blockNumber)
	})
}

func Fuzz_Nosy_InstrumentedClient_BlockByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var m *metrics.RPCClientMetrics
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
		if c == nil || m == nil {
			return
		}

		ic := NewInstrumentedClient(c, m)
		ic.BlockByHash(ctx, hash)
	})
}

func Fuzz_Nosy_InstrumentedClient_BlockByNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var m *metrics.RPCClientMetrics
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
		if c == nil || m == nil || number == nil {
			return
		}

		ic := NewInstrumentedClient(c, m)
		ic.BlockByNumber(ctx, number)
	})
}

func Fuzz_Nosy_InstrumentedClient_BlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var m *metrics.RPCClientMetrics
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if c == nil || m == nil {
			return
		}

		ic := NewInstrumentedClient(c, m)
		ic.BlockNumber(ctx)
	})
}

func Fuzz_Nosy_InstrumentedClient_CallContract__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var m *metrics.RPCClientMetrics
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
		if c == nil || m == nil || blockNumber == nil {
			return
		}

		ic := NewInstrumentedClient(c, m)
		ic.CallContract(ctx, msg, blockNumber)
	})
}

func Fuzz_Nosy_InstrumentedClient_CallContractAtHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var m *metrics.RPCClientMetrics
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
		if c == nil || m == nil {
			return
		}

		ic := NewInstrumentedClient(c, m)
		ic.CallContractAtHash(ctx, msg, blockHash)
	})
}

func Fuzz_Nosy_InstrumentedClient_ChainID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var m *metrics.RPCClientMetrics
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if c == nil || m == nil {
			return
		}

		ic := NewInstrumentedClient(c, m)
		ic.ChainID(ctx)
	})
}

func Fuzz_Nosy_InstrumentedClient_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var m *metrics.RPCClientMetrics
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if c == nil || m == nil {
			return
		}

		ic := NewInstrumentedClient(c, m)
		ic.Close()
	})
}

func Fuzz_Nosy_InstrumentedClient_CodeAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var m *metrics.RPCClientMetrics
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
		var blockNumber *big.Int
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		if c == nil || m == nil || blockNumber == nil {
			return
		}

		ic := NewInstrumentedClient(c, m)
		ic.CodeAt(ctx, account, blockNumber)
	})
}

func Fuzz_Nosy_InstrumentedClient_EstimateGas__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var m *metrics.RPCClientMetrics
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
		if c == nil || m == nil {
			return
		}

		ic := NewInstrumentedClient(c, m)
		ic.EstimateGas(ctx, msg)
	})
}

func Fuzz_Nosy_InstrumentedClient_FilterLogs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var m *metrics.RPCClientMetrics
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
		if c == nil || m == nil {
			return
		}

		ic := NewInstrumentedClient(c, m)
		ic.FilterLogs(ctx, q)
	})
}

func Fuzz_Nosy_InstrumentedClient_HeaderByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var m *metrics.RPCClientMetrics
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
		if c == nil || m == nil {
			return
		}

		ic := NewInstrumentedClient(c, m)
		ic.HeaderByHash(ctx, hash)
	})
}

func Fuzz_Nosy_InstrumentedClient_HeaderByNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var m *metrics.RPCClientMetrics
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
		if c == nil || m == nil || number == nil {
			return
		}

		ic := NewInstrumentedClient(c, m)
		ic.HeaderByNumber(ctx, number)
	})
}

func Fuzz_Nosy_InstrumentedClient_NetworkID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var m *metrics.RPCClientMetrics
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if c == nil || m == nil {
			return
		}

		ic := NewInstrumentedClient(c, m)
		ic.NetworkID(ctx)
	})
}

func Fuzz_Nosy_InstrumentedClient_NonceAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var m *metrics.RPCClientMetrics
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
		var blockNumber *big.Int
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		if c == nil || m == nil || blockNumber == nil {
			return
		}

		ic := NewInstrumentedClient(c, m)
		ic.NonceAt(ctx, account, blockNumber)
	})
}

func Fuzz_Nosy_InstrumentedClient_PeerCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var m *metrics.RPCClientMetrics
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if c == nil || m == nil {
			return
		}

		ic := NewInstrumentedClient(c, m)
		ic.PeerCount(ctx)
	})
}

func Fuzz_Nosy_InstrumentedClient_PendingBalanceAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var m *metrics.RPCClientMetrics
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
		if c == nil || m == nil {
			return
		}

		ic := NewInstrumentedClient(c, m)
		ic.PendingBalanceAt(ctx, account)
	})
}

func Fuzz_Nosy_InstrumentedClient_PendingCallContract__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var m *metrics.RPCClientMetrics
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
		if c == nil || m == nil {
			return
		}

		ic := NewInstrumentedClient(c, m)
		ic.PendingCallContract(ctx, msg)
	})
}

func Fuzz_Nosy_InstrumentedClient_PendingCodeAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var m *metrics.RPCClientMetrics
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
		if c == nil || m == nil {
			return
		}

		ic := NewInstrumentedClient(c, m)
		ic.PendingCodeAt(ctx, account)
	})
}

func Fuzz_Nosy_InstrumentedClient_PendingNonceAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var m *metrics.RPCClientMetrics
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
		if c == nil || m == nil {
			return
		}

		ic := NewInstrumentedClient(c, m)
		ic.PendingNonceAt(ctx, account)
	})
}

func Fuzz_Nosy_InstrumentedClient_PendingStorageAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var m *metrics.RPCClientMetrics
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
		if c == nil || m == nil {
			return
		}

		ic := NewInstrumentedClient(c, m)
		ic.PendingStorageAt(ctx, account, key)
	})
}

func Fuzz_Nosy_InstrumentedClient_PendingTransactionCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var m *metrics.RPCClientMetrics
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if c == nil || m == nil {
			return
		}

		ic := NewInstrumentedClient(c, m)
		ic.PendingTransactionCount(ctx)
	})
}

func Fuzz_Nosy_InstrumentedClient_RPC__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var m *metrics.RPCClientMetrics
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if c == nil || m == nil {
			return
		}

		ic := NewInstrumentedClient(c, m)
		ic.RPC()
	})
}

func Fuzz_Nosy_InstrumentedClient_SendTransaction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var m *metrics.RPCClientMetrics
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
		if c == nil || m == nil || tx == nil {
			return
		}

		ic := NewInstrumentedClient(c, m)
		ic.SendTransaction(ctx, tx)
	})
}

func Fuzz_Nosy_InstrumentedClient_StorageAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var m *metrics.RPCClientMetrics
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
		var blockNumber *big.Int
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		if c == nil || m == nil || blockNumber == nil {
			return
		}

		ic := NewInstrumentedClient(c, m)
		ic.StorageAt(ctx, account, key, blockNumber)
	})
}

// skipping Fuzz_Nosy_InstrumentedClient_SubscribeFilterLogs__ because parameters include func, chan, or unsupported interface: chan<- github.com/ethereum/go-ethereum/core/types.Log

// skipping Fuzz_Nosy_InstrumentedClient_SubscribeNewHead__ because parameters include func, chan, or unsupported interface: chan<- *github.com/ethereum/go-ethereum/core/types.Header

func Fuzz_Nosy_InstrumentedClient_SuggestGasPrice__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var m *metrics.RPCClientMetrics
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if c == nil || m == nil {
			return
		}

		ic := NewInstrumentedClient(c, m)
		ic.SuggestGasPrice(ctx)
	})
}

func Fuzz_Nosy_InstrumentedClient_SuggestGasTipCap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var m *metrics.RPCClientMetrics
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if c == nil || m == nil {
			return
		}

		ic := NewInstrumentedClient(c, m)
		ic.SuggestGasTipCap(ctx)
	})
}

func Fuzz_Nosy_InstrumentedClient_SyncProgress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var m *metrics.RPCClientMetrics
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if c == nil || m == nil {
			return
		}

		ic := NewInstrumentedClient(c, m)
		ic.SyncProgress(ctx)
	})
}

func Fuzz_Nosy_InstrumentedClient_TransactionByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var m *metrics.RPCClientMetrics
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
		if c == nil || m == nil {
			return
		}

		ic := NewInstrumentedClient(c, m)
		ic.TransactionByHash(ctx, hash)
	})
}

func Fuzz_Nosy_InstrumentedClient_TransactionCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var m *metrics.RPCClientMetrics
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
		if c == nil || m == nil {
			return
		}

		ic := NewInstrumentedClient(c, m)
		ic.TransactionCount(ctx, blockHash)
	})
}

func Fuzz_Nosy_InstrumentedClient_TransactionInBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var m *metrics.RPCClientMetrics
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
		var index uint
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		if c == nil || m == nil {
			return
		}

		ic := NewInstrumentedClient(c, m)
		ic.TransactionInBlock(ctx, blockHash, index)
	})
}

func Fuzz_Nosy_InstrumentedClient_TransactionReceipt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var m *metrics.RPCClientMetrics
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
		if c == nil || m == nil {
			return
		}

		ic := NewInstrumentedClient(c, m)
		ic.TransactionReceipt(ctx, txHash)
	})
}

func Fuzz_Nosy_InstrumentedClient_TransactionSender__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var m *metrics.RPCClientMetrics
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
		if c == nil || m == nil || tx == nil {
			return
		}

		ic := NewInstrumentedClient(c, m)
		ic.TransactionSender(ctx, tx, block, index)
	})
}

func Fuzz_Nosy_InstrumentedRPCClient_BatchCallContext__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ic *InstrumentedRPCClient
		fill_err = tp.Fill(&ic)
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
		if ic == nil {
			return
		}

		ic.BatchCallContext(ctx, b)
	})
}

// skipping Fuzz_Nosy_InstrumentedRPCClient_CallContext__ because parameters include func, chan, or unsupported interface: any

func Fuzz_Nosy_InstrumentedRPCClient_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ic *InstrumentedRPCClient
		fill_err = tp.Fill(&ic)
		if fill_err != nil {
			return
		}
		if ic == nil {
			return
		}

		ic.Close()
	})
}

// skipping Fuzz_Nosy_InstrumentedRPCClient_EthSubscribe__ because parameters include func, chan, or unsupported interface: any

func Fuzz_Nosy_PollingClient_BatchCallContext__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *PollingClient
		fill_err = tp.Fill(&w)
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
		if w == nil {
			return
		}

		w.BatchCallContext(ctx, b)
	})
}

// skipping Fuzz_Nosy_PollingClient_CallContext__ because parameters include func, chan, or unsupported interface: any

func Fuzz_Nosy_PollingClient_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *PollingClient
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		w.Close()
	})
}

// skipping Fuzz_Nosy_PollingClient_EthSubscribe__ because parameters include func, chan, or unsupported interface: any

func Fuzz_Nosy_PollingClient_getLatestHeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *PollingClient
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		w.getLatestHeader()
	})
}

func Fuzz_Nosy_PollingClient_pollHeads__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *PollingClient
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		w.pollHeads()
	})
}

func Fuzz_Nosy_PollingClient_reqPoll__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *PollingClient
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		w.reqPoll()
	})
}

func Fuzz_Nosy_RateLimitingClient_BatchCallContext__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *RateLimitingClient
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var batch []rpc.BatchElem
		fill_err = tp.Fill(&batch)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.BatchCallContext(ctx, batch)
	})
}

// skipping Fuzz_Nosy_RateLimitingClient_CallContext__ because parameters include func, chan, or unsupported interface: any

func Fuzz_Nosy_RateLimitingClient_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *RateLimitingClient
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.Close()
	})
}

// skipping Fuzz_Nosy_RateLimitingClient_EthSubscribe__ because parameters include func, chan, or unsupported interface: any

func Fuzz_Nosy_lazyRPC_BatchCallContext__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var endpoint string
		fill_err = tp.Fill(&endpoint)
		if fill_err != nil {
			return
		}
		var cfg rpcConfig
		fill_err = tp.Fill(&cfg)
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

		l := newLazyRPC(endpoint, cfg)
		l.BatchCallContext(ctx, b)
	})
}

// skipping Fuzz_Nosy_lazyRPC_CallContext__ because parameters include func, chan, or unsupported interface: any

func Fuzz_Nosy_lazyRPC_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var endpoint string
		fill_err = tp.Fill(&endpoint)
		if fill_err != nil {
			return
		}
		var cfg rpcConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}

		l := newLazyRPC(endpoint, cfg)
		l.Close()
	})
}

// skipping Fuzz_Nosy_lazyRPC_EthSubscribe__ because parameters include func, chan, or unsupported interface: any

func Fuzz_Nosy_lazyRPC_dial__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var endpoint string
		fill_err = tp.Fill(&endpoint)
		if fill_err != nil {
			return
		}
		var cfg rpcConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		l := newLazyRPC(endpoint, cfg)
		l.dial(ctx)
	})
}

func Fuzz_Nosy_BasicHTTPClientOption_Apply__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h http.Header
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var c *BasicHTTPClient
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		_x1 := WithHeader(h)
		_x1.Apply(c)
	})
}

// skipping Fuzz_Nosy_BasicHTTPClientOptionFn_Apply__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/client.BasicHTTPClientOptionFn

// skipping Fuzz_Nosy_Client_BalanceAt__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/client.Client

// skipping Fuzz_Nosy_Client_BlockByHash__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/client.Client

// skipping Fuzz_Nosy_Client_BlockByNumber__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/client.Client

// skipping Fuzz_Nosy_Client_BlockNumber__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/client.Client

// skipping Fuzz_Nosy_Client_CallContract__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/client.Client

// skipping Fuzz_Nosy_Client_CallContractAtHash__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/client.Client

// skipping Fuzz_Nosy_Client_ChainID__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/client.Client

// skipping Fuzz_Nosy_Client_Close__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/client.Client

// skipping Fuzz_Nosy_Client_CodeAt__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/client.Client

// skipping Fuzz_Nosy_Client_EstimateGas__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/client.Client

// skipping Fuzz_Nosy_Client_FilterLogs__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/client.Client

// skipping Fuzz_Nosy_Client_HeaderByHash__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/client.Client

// skipping Fuzz_Nosy_Client_HeaderByNumber__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/client.Client

// skipping Fuzz_Nosy_Client_NetworkID__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/client.Client

// skipping Fuzz_Nosy_Client_NonceAt__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/client.Client

// skipping Fuzz_Nosy_Client_PeerCount__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/client.Client

// skipping Fuzz_Nosy_Client_PendingBalanceAt__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/client.Client

// skipping Fuzz_Nosy_Client_PendingCallContract__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/client.Client

// skipping Fuzz_Nosy_Client_PendingCodeAt__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/client.Client

// skipping Fuzz_Nosy_Client_PendingNonceAt__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/client.Client

// skipping Fuzz_Nosy_Client_PendingStorageAt__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/client.Client

// skipping Fuzz_Nosy_Client_PendingTransactionCount__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/client.Client

// skipping Fuzz_Nosy_Client_RPC__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/client.Client

// skipping Fuzz_Nosy_Client_SendTransaction__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/client.Client

// skipping Fuzz_Nosy_Client_StorageAt__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/client.Client

// skipping Fuzz_Nosy_Client_SubscribeFilterLogs__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/client.Client

// skipping Fuzz_Nosy_Client_SubscribeNewHead__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/client.Client

// skipping Fuzz_Nosy_Client_SuggestGasPrice__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/client.Client

// skipping Fuzz_Nosy_Client_SuggestGasTipCap__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/client.Client

// skipping Fuzz_Nosy_Client_SyncProgress__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/client.Client

// skipping Fuzz_Nosy_Client_TransactionByHash__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/client.Client

// skipping Fuzz_Nosy_Client_TransactionCount__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/client.Client

// skipping Fuzz_Nosy_Client_TransactionInBlock__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/client.Client

// skipping Fuzz_Nosy_Client_TransactionReceipt__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/client.Client

// skipping Fuzz_Nosy_Client_TransactionSender__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/client.Client

// skipping Fuzz_Nosy_HTTP_Get__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/client.HTTP

func Fuzz_Nosy_RPC_BatchCallContext__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var cfg rpcConfig
		fill_err = tp.Fill(&cfg)
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
		if c == nil {
			return
		}

		_x1 := wrapClient(c, cfg)
		_x1.BatchCallContext(ctx, b)
	})
}

// skipping Fuzz_Nosy_RPC_CallContext__ because parameters include func, chan, or unsupported interface: any

func Fuzz_Nosy_RPC_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var cfg rpcConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		_x1 := wrapClient(c, cfg)
		_x1.Close()
	})
}

// skipping Fuzz_Nosy_RPC_EthSubscribe__ because parameters include func, chan, or unsupported interface: any

func Fuzz_Nosy_IsURLAvailable__(f *testing.F) {
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
		var address string
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}

		IsURLAvailable(ctx, address)
	})
}

func Fuzz_Nosy_addHTTPHeaders__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header http.Header
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		var hs []http.Header
		fill_err = tp.Fill(&hs)
		if fill_err != nil {
			return
		}

		addHTTPHeaders(header, hs...)
	})
}

// skipping Fuzz_Nosy_instrument2__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/metrics.RPCClientMetricer
