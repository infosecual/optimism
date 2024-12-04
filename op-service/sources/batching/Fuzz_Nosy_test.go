package batching

import (
	"testing"
	"github.com/ethereum-optimism/optimism/op-service/sources"
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
	
	// skipping Fuzz_Nosy_BalanceCall_HandleResult__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_BalanceCall_ToBatchElemCreator__(f *testing.F) {
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

	b := NewBalanceCall(addr)
	b.ToBatchElemCreator()
	})
}

func Fuzz_Nosy_BoundContract_Addr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var abi *abi.ABI
		fill_err = tp.Fill(&abi)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
	if abi == nil {
		return
	}

	b := NewBoundContract(abi, addr)
	b.Addr()
	})
}

// skipping Fuzz_Nosy_BoundContract_Call__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_BoundContract_DecodeCall__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var abi *abi.ABI
		fill_err = tp.Fill(&abi)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var d3 []byte
		fill_err = tp.Fill(&d3)
		if fill_err != nil {
			return
		}
	if abi == nil {
		return
	}

	b := NewBoundContract(abi, addr)
	b.DecodeCall(d3)
	})
}

func Fuzz_Nosy_BoundContract_DecodeEvent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var abi *abi.ABI
		fill_err = tp.Fill(&abi)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var log *types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
	if abi == nil|| log == nil {
		return
	}

	b := NewBoundContract(abi, addr)
	b.DecodeEvent(log)
	})
}

func Fuzz_Nosy_CallResult_GetAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CallResult
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var i int
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.GetAddress(i)
	})
}

func Fuzz_Nosy_CallResult_GetBigInt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CallResult
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var i int
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.GetBigInt(i)
	})
}

func Fuzz_Nosy_CallResult_GetBool__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CallResult
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var i int
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.GetBool(i)
	})
}

func Fuzz_Nosy_CallResult_GetBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CallResult
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var i int
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.GetBytes(i)
	})
}

func Fuzz_Nosy_CallResult_GetBytes32__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CallResult
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var i int
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.GetBytes32(i)
	})
}

func Fuzz_Nosy_CallResult_GetBytes32Slice__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CallResult
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var i int
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.GetBytes32Slice(i)
	})
}

func Fuzz_Nosy_CallResult_GetHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CallResult
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var i int
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.GetHash(i)
	})
}

func Fuzz_Nosy_CallResult_GetString__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CallResult
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var i int
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.GetString(i)
	})
}

// skipping Fuzz_Nosy_CallResult_GetStruct__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_CallResult_GetUint32__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CallResult
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var i int
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.GetUint32(i)
	})
}

func Fuzz_Nosy_CallResult_GetUint64__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CallResult
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var i int
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.GetUint64(i)
	})
}

func Fuzz_Nosy_CallResult_GetUint8__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CallResult
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var i int
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.GetUint8(i)
	})
}

func Fuzz_Nosy_ContractCall_CallMethod__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ContractCall
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.CallMethod()
	})
}

func Fuzz_Nosy_ContractCall_CreateResult__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ContractCall
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.CreateResult()
	})
}

// skipping Fuzz_Nosy_ContractCall_HandleResult__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_ContractCall_Pack__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ContractCall
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Pack()
	})
}

func Fuzz_Nosy_ContractCall_ToBatchElemCreator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ContractCall
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.ToBatchElemCreator()
	})
}

func Fuzz_Nosy_ContractCall_ToCallArgs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ContractCall
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.ToCallArgs()
	})
}

func Fuzz_Nosy_ContractCall_ToTxCandidate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ContractCall
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.ToTxCandidate()
	})
}

func Fuzz_Nosy_ContractCall_Unpack__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ContractCall
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var hex hexutil.Bytes
		fill_err = tp.Fill(&hex)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Unpack(hex)
	})
}

func Fuzz_Nosy_IterativeBatchCall[K, V any]_Complete__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var ibc *IterativeBatchCall[K, V]
		fill_err = tp.Fill(&ibc)
		if fill_err != nil {
			return
		}
	if ibc == nil {
		return
	}

	ibc.Complete()
	})
}

func Fuzz_Nosy_IterativeBatchCall[K, V any]_Fetch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var ibc *IterativeBatchCall[K, V]
		fill_err = tp.Fill(&ibc)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
	if ibc == nil {
		return
	}

	ibc.Fetch(ctx)
	})
}

func Fuzz_Nosy_IterativeBatchCall[K, V any]_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var ibc *IterativeBatchCall[K, V]
		fill_err = tp.Fill(&ibc)
		if fill_err != nil {
			return
		}
	if ibc == nil {
		return
	}

	ibc.Reset()
	})
}

func Fuzz_Nosy_IterativeBatchCall[K, V any]_Result__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var ibc *IterativeBatchCall[K, V]
		fill_err = tp.Fill(&ibc)
		if fill_err != nil {
			return
		}
	if ibc == nil {
		return
	}

	ibc.Result()
	})
}

func Fuzz_Nosy_MultiCaller_BatchSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var m *MultiCaller
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
	if m == nil {
		return
	}

	m.BatchSize()
	})
}

// skipping Fuzz_Nosy_MultiCaller_Call__ because parameters include func, chan, or unsupported interface: []github.com/ethereum-optimism/optimism/op-service/sources/batching.Call

// skipping Fuzz_Nosy_MultiCaller_SingleCall__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/sources/batching.Call

// skipping Fuzz_Nosy_Call_HandleResult__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/sources/batching.Call

// skipping Fuzz_Nosy_Call_ToBatchElemCreator__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/sources/batching.Call

// skipping Fuzz_Nosy_EthRpc_BatchCallContext__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/sources/batching.EthRpc

// skipping Fuzz_Nosy_EthRpc_CallContext__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/sources/batching.EthRpc

// skipping Fuzz_Nosy_ReadArray__ because parameters include func, chan, or unsupported interface: func(i *math/big.Int) *github.com/ethereum-optimism/optimism/op-service/sources/batching.ContractCall

