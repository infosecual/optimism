package test

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum-optimism/optimism/op-service/txmgr"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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

func Fuzz_Nosy_AbiBasedRpc_AddContract__(f *testing.F) {
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
		var t2 common.Address
		fill_err = tp.Fill(&t2)
		if fill_err != nil {
			return
		}
		var c3 *abi.ABI
		fill_err = tp.Fill(&c3)
		if fill_err != nil {
			return
		}
		var t4 common.Address
		fill_err = tp.Fill(&t4)
		if fill_err != nil {
			return
		}
		var c5 *abi.ABI
		fill_err = tp.Fill(&c5)
		if fill_err != nil {
			return
		}
		if t1 == nil || c3 == nil || c5 == nil {
			return
		}

		l := NewAbiBasedRpc(t1, t2, c3)
		l.AddContract(t4, c5)
	})
}

// skipping Fuzz_Nosy_AbiBasedRpc_SetError__ because parameters include func, chan, or unsupported interface: []interface{}

// skipping Fuzz_Nosy_AbiBasedRpc_SetResponse__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_AbiBasedRpc_VerifyTxCandidate__(f *testing.F) {
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
		var to common.Address
		fill_err = tp.Fill(&to)
		if fill_err != nil {
			return
		}
		var contractAbi *abi.ABI
		fill_err = tp.Fill(&contractAbi)
		if fill_err != nil {
			return
		}
		var candidate txmgr.TxCandidate
		fill_err = tp.Fill(&candidate)
		if fill_err != nil {
			return
		}
		if t1 == nil || contractAbi == nil {
			return
		}

		l := NewAbiBasedRpc(t1, to, contractAbi)
		l.VerifyTxCandidate(candidate)
	})
}

func Fuzz_Nosy_AbiBasedRpc_abi__(f *testing.F) {
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
		var t2 common.Address
		fill_err = tp.Fill(&t2)
		if fill_err != nil {
			return
		}
		var contractAbi *abi.ABI
		fill_err = tp.Fill(&contractAbi)
		if fill_err != nil {
			return
		}
		var t4 common.Address
		fill_err = tp.Fill(&t4)
		if fill_err != nil {
			return
		}
		if t1 == nil || contractAbi == nil {
			return
		}

		l := NewAbiBasedRpc(t1, t2, contractAbi)
		l.abi(t4)
	})
}

func Fuzz_Nosy_ERC20ApprovalIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ERC20ApprovalIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_ERC20ApprovalIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ERC20ApprovalIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_ERC20ApprovalIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ERC20ApprovalIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_ERC20Caller_Allowance__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ERC20 *ERC20Caller
		fill_err = tp.Fill(&_ERC20)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		var arg1 common.Address
		fill_err = tp.Fill(&arg1)
		if fill_err != nil {
			return
		}
		if _ERC20 == nil || opts == nil {
			return
		}

		_ERC20.Allowance(opts, arg0, arg1)
	})
}

func Fuzz_Nosy_ERC20Caller_BalanceOf__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ERC20 *ERC20Caller
		fill_err = tp.Fill(&_ERC20)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _ERC20 == nil || opts == nil {
			return
		}

		_ERC20.BalanceOf(opts, arg0)
	})
}

func Fuzz_Nosy_ERC20Caller_DOMAINSEPARATOR__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ERC20 *ERC20Caller
		fill_err = tp.Fill(&_ERC20)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ERC20 == nil || opts == nil {
			return
		}

		_ERC20.DOMAINSEPARATOR(opts)
	})
}

func Fuzz_Nosy_ERC20Caller_Decimals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ERC20 *ERC20Caller
		fill_err = tp.Fill(&_ERC20)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ERC20 == nil || opts == nil {
			return
		}

		_ERC20.Decimals(opts)
	})
}

func Fuzz_Nosy_ERC20Caller_Name__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ERC20 *ERC20Caller
		fill_err = tp.Fill(&_ERC20)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ERC20 == nil || opts == nil {
			return
		}

		_ERC20.Name(opts)
	})
}

func Fuzz_Nosy_ERC20Caller_Nonces__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ERC20 *ERC20Caller
		fill_err = tp.Fill(&_ERC20)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _ERC20 == nil || opts == nil {
			return
		}

		_ERC20.Nonces(opts, arg0)
	})
}

func Fuzz_Nosy_ERC20Caller_Symbol__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ERC20 *ERC20Caller
		fill_err = tp.Fill(&_ERC20)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ERC20 == nil || opts == nil {
			return
		}

		_ERC20.Symbol(opts)
	})
}

func Fuzz_Nosy_ERC20Caller_TotalSupply__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ERC20 *ERC20Caller
		fill_err = tp.Fill(&_ERC20)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ERC20 == nil || opts == nil {
			return
		}

		_ERC20.TotalSupply(opts)
	})
}

// skipping Fuzz_Nosy_ERC20CallerRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

func Fuzz_Nosy_ERC20CallerSession_Allowance__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ERC20 *ERC20CallerSession
		fill_err = tp.Fill(&_ERC20)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		var arg1 common.Address
		fill_err = tp.Fill(&arg1)
		if fill_err != nil {
			return
		}
		if _ERC20 == nil {
			return
		}

		_ERC20.Allowance(arg0, arg1)
	})
}

func Fuzz_Nosy_ERC20CallerSession_BalanceOf__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ERC20 *ERC20CallerSession
		fill_err = tp.Fill(&_ERC20)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _ERC20 == nil {
			return
		}

		_ERC20.BalanceOf(arg0)
	})
}

func Fuzz_Nosy_ERC20CallerSession_DOMAINSEPARATOR__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ERC20 *ERC20CallerSession
		fill_err = tp.Fill(&_ERC20)
		if fill_err != nil {
			return
		}
		if _ERC20 == nil {
			return
		}

		_ERC20.DOMAINSEPARATOR()
	})
}

func Fuzz_Nosy_ERC20CallerSession_Decimals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ERC20 *ERC20CallerSession
		fill_err = tp.Fill(&_ERC20)
		if fill_err != nil {
			return
		}
		if _ERC20 == nil {
			return
		}

		_ERC20.Decimals()
	})
}

func Fuzz_Nosy_ERC20CallerSession_Name__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ERC20 *ERC20CallerSession
		fill_err = tp.Fill(&_ERC20)
		if fill_err != nil {
			return
		}
		if _ERC20 == nil {
			return
		}

		_ERC20.Name()
	})
}

func Fuzz_Nosy_ERC20CallerSession_Nonces__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ERC20 *ERC20CallerSession
		fill_err = tp.Fill(&_ERC20)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _ERC20 == nil {
			return
		}

		_ERC20.Nonces(arg0)
	})
}

func Fuzz_Nosy_ERC20CallerSession_Symbol__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ERC20 *ERC20CallerSession
		fill_err = tp.Fill(&_ERC20)
		if fill_err != nil {
			return
		}
		if _ERC20 == nil {
			return
		}

		_ERC20.Symbol()
	})
}

func Fuzz_Nosy_ERC20CallerSession_TotalSupply__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ERC20 *ERC20CallerSession
		fill_err = tp.Fill(&_ERC20)
		if fill_err != nil {
			return
		}
		if _ERC20 == nil {
			return
		}

		_ERC20.TotalSupply()
	})
}

func Fuzz_Nosy_ERC20Filterer_FilterApproval__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ERC20 *ERC20Filterer
		fill_err = tp.Fill(&_ERC20)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var owner []common.Address
		fill_err = tp.Fill(&owner)
		if fill_err != nil {
			return
		}
		var spender []common.Address
		fill_err = tp.Fill(&spender)
		if fill_err != nil {
			return
		}
		if _ERC20 == nil || opts == nil {
			return
		}

		_ERC20.FilterApproval(opts, owner, spender)
	})
}

func Fuzz_Nosy_ERC20Filterer_FilterTransfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ERC20 *ERC20Filterer
		fill_err = tp.Fill(&_ERC20)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var from []common.Address
		fill_err = tp.Fill(&from)
		if fill_err != nil {
			return
		}
		var to []common.Address
		fill_err = tp.Fill(&to)
		if fill_err != nil {
			return
		}
		if _ERC20 == nil || opts == nil {
			return
		}

		_ERC20.FilterTransfer(opts, from, to)
	})
}

func Fuzz_Nosy_ERC20Filterer_ParseApproval__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ERC20 *ERC20Filterer
		fill_err = tp.Fill(&_ERC20)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ERC20 == nil {
			return
		}

		_ERC20.ParseApproval(log)
	})
}

func Fuzz_Nosy_ERC20Filterer_ParseTransfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ERC20 *ERC20Filterer
		fill_err = tp.Fill(&_ERC20)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _ERC20 == nil {
			return
		}

		_ERC20.ParseTransfer(log)
	})
}

// skipping Fuzz_Nosy_ERC20Filterer_WatchApproval__ because parameters include func, chan, or unsupported interface: chan<- *github.com/ethereum-optimism/optimism/op-service/sources/batching/test.ERC20Approval

// skipping Fuzz_Nosy_ERC20Filterer_WatchTransfer__ because parameters include func, chan, or unsupported interface: chan<- *github.com/ethereum-optimism/optimism/op-service/sources/batching/test.ERC20Transfer

// skipping Fuzz_Nosy_ERC20Raw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

// skipping Fuzz_Nosy_ERC20Raw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_ERC20Raw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ERC20 *ERC20Raw
		fill_err = tp.Fill(&_ERC20)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ERC20 == nil || opts == nil {
			return
		}

		_ERC20.Transfer(opts)
	})
}

func Fuzz_Nosy_ERC20Session_Allowance__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ERC20 *ERC20Session
		fill_err = tp.Fill(&_ERC20)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		var arg1 common.Address
		fill_err = tp.Fill(&arg1)
		if fill_err != nil {
			return
		}
		if _ERC20 == nil {
			return
		}

		_ERC20.Allowance(arg0, arg1)
	})
}

func Fuzz_Nosy_ERC20Session_Approve__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ERC20 *ERC20Session
		fill_err = tp.Fill(&_ERC20)
		if fill_err != nil {
			return
		}
		var spender common.Address
		fill_err = tp.Fill(&spender)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if _ERC20 == nil || amount == nil {
			return
		}

		_ERC20.Approve(spender, amount)
	})
}

func Fuzz_Nosy_ERC20Session_BalanceOf__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ERC20 *ERC20Session
		fill_err = tp.Fill(&_ERC20)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _ERC20 == nil {
			return
		}

		_ERC20.BalanceOf(arg0)
	})
}

func Fuzz_Nosy_ERC20Session_DOMAINSEPARATOR__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ERC20 *ERC20Session
		fill_err = tp.Fill(&_ERC20)
		if fill_err != nil {
			return
		}
		if _ERC20 == nil {
			return
		}

		_ERC20.DOMAINSEPARATOR()
	})
}

func Fuzz_Nosy_ERC20Session_Decimals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ERC20 *ERC20Session
		fill_err = tp.Fill(&_ERC20)
		if fill_err != nil {
			return
		}
		if _ERC20 == nil {
			return
		}

		_ERC20.Decimals()
	})
}

func Fuzz_Nosy_ERC20Session_Name__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ERC20 *ERC20Session
		fill_err = tp.Fill(&_ERC20)
		if fill_err != nil {
			return
		}
		if _ERC20 == nil {
			return
		}

		_ERC20.Name()
	})
}

func Fuzz_Nosy_ERC20Session_Nonces__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ERC20 *ERC20Session
		fill_err = tp.Fill(&_ERC20)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _ERC20 == nil {
			return
		}

		_ERC20.Nonces(arg0)
	})
}

func Fuzz_Nosy_ERC20Session_Permit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ERC20 *ERC20Session
		fill_err = tp.Fill(&_ERC20)
		if fill_err != nil {
			return
		}
		var owner common.Address
		fill_err = tp.Fill(&owner)
		if fill_err != nil {
			return
		}
		var spender common.Address
		fill_err = tp.Fill(&spender)
		if fill_err != nil {
			return
		}
		var value *big.Int
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
		var deadline *big.Int
		fill_err = tp.Fill(&deadline)
		if fill_err != nil {
			return
		}
		var v uint8
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		var r [32]byte
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var s [32]byte
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if _ERC20 == nil || value == nil || deadline == nil {
			return
		}

		_ERC20.Permit(owner, spender, value, deadline, v, r, s)
	})
}

func Fuzz_Nosy_ERC20Session_Symbol__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ERC20 *ERC20Session
		fill_err = tp.Fill(&_ERC20)
		if fill_err != nil {
			return
		}
		if _ERC20 == nil {
			return
		}

		_ERC20.Symbol()
	})
}

func Fuzz_Nosy_ERC20Session_TotalSupply__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ERC20 *ERC20Session
		fill_err = tp.Fill(&_ERC20)
		if fill_err != nil {
			return
		}
		if _ERC20 == nil {
			return
		}

		_ERC20.TotalSupply()
	})
}

func Fuzz_Nosy_ERC20Session_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ERC20 *ERC20Session
		fill_err = tp.Fill(&_ERC20)
		if fill_err != nil {
			return
		}
		var to common.Address
		fill_err = tp.Fill(&to)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if _ERC20 == nil || amount == nil {
			return
		}

		_ERC20.Transfer(to, amount)
	})
}

func Fuzz_Nosy_ERC20Session_TransferFrom__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ERC20 *ERC20Session
		fill_err = tp.Fill(&_ERC20)
		if fill_err != nil {
			return
		}
		var from common.Address
		fill_err = tp.Fill(&from)
		if fill_err != nil {
			return
		}
		var to common.Address
		fill_err = tp.Fill(&to)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if _ERC20 == nil || amount == nil {
			return
		}

		_ERC20.TransferFrom(from, to, amount)
	})
}

func Fuzz_Nosy_ERC20Transactor_Approve__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ERC20 *ERC20Transactor
		fill_err = tp.Fill(&_ERC20)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var spender common.Address
		fill_err = tp.Fill(&spender)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if _ERC20 == nil || opts == nil || amount == nil {
			return
		}

		_ERC20.Approve(opts, spender, amount)
	})
}

func Fuzz_Nosy_ERC20Transactor_Permit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ERC20 *ERC20Transactor
		fill_err = tp.Fill(&_ERC20)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var owner common.Address
		fill_err = tp.Fill(&owner)
		if fill_err != nil {
			return
		}
		var spender common.Address
		fill_err = tp.Fill(&spender)
		if fill_err != nil {
			return
		}
		var value *big.Int
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
		var deadline *big.Int
		fill_err = tp.Fill(&deadline)
		if fill_err != nil {
			return
		}
		var v uint8
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		var r [32]byte
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var s [32]byte
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if _ERC20 == nil || opts == nil || value == nil || deadline == nil {
			return
		}

		_ERC20.Permit(opts, owner, spender, value, deadline, v, r, s)
	})
}

func Fuzz_Nosy_ERC20Transactor_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ERC20 *ERC20Transactor
		fill_err = tp.Fill(&_ERC20)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var to common.Address
		fill_err = tp.Fill(&to)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if _ERC20 == nil || opts == nil || amount == nil {
			return
		}

		_ERC20.Transfer(opts, to, amount)
	})
}

func Fuzz_Nosy_ERC20Transactor_TransferFrom__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ERC20 *ERC20Transactor
		fill_err = tp.Fill(&_ERC20)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var from common.Address
		fill_err = tp.Fill(&from)
		if fill_err != nil {
			return
		}
		var to common.Address
		fill_err = tp.Fill(&to)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if _ERC20 == nil || opts == nil || amount == nil {
			return
		}

		_ERC20.TransferFrom(opts, from, to, amount)
	})
}

// skipping Fuzz_Nosy_ERC20TransactorRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_ERC20TransactorRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ERC20 *ERC20TransactorRaw
		fill_err = tp.Fill(&_ERC20)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _ERC20 == nil || opts == nil {
			return
		}

		_ERC20.Transfer(opts)
	})
}

func Fuzz_Nosy_ERC20TransactorSession_Approve__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ERC20 *ERC20TransactorSession
		fill_err = tp.Fill(&_ERC20)
		if fill_err != nil {
			return
		}
		var spender common.Address
		fill_err = tp.Fill(&spender)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if _ERC20 == nil || amount == nil {
			return
		}

		_ERC20.Approve(spender, amount)
	})
}

func Fuzz_Nosy_ERC20TransactorSession_Permit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ERC20 *ERC20TransactorSession
		fill_err = tp.Fill(&_ERC20)
		if fill_err != nil {
			return
		}
		var owner common.Address
		fill_err = tp.Fill(&owner)
		if fill_err != nil {
			return
		}
		var spender common.Address
		fill_err = tp.Fill(&spender)
		if fill_err != nil {
			return
		}
		var value *big.Int
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
		var deadline *big.Int
		fill_err = tp.Fill(&deadline)
		if fill_err != nil {
			return
		}
		var v uint8
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		var r [32]byte
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var s [32]byte
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if _ERC20 == nil || value == nil || deadline == nil {
			return
		}

		_ERC20.Permit(owner, spender, value, deadline, v, r, s)
	})
}

func Fuzz_Nosy_ERC20TransactorSession_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ERC20 *ERC20TransactorSession
		fill_err = tp.Fill(&_ERC20)
		if fill_err != nil {
			return
		}
		var to common.Address
		fill_err = tp.Fill(&to)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if _ERC20 == nil || amount == nil {
			return
		}

		_ERC20.Transfer(to, amount)
	})
}

func Fuzz_Nosy_ERC20TransactorSession_TransferFrom__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _ERC20 *ERC20TransactorSession
		fill_err = tp.Fill(&_ERC20)
		if fill_err != nil {
			return
		}
		var from common.Address
		fill_err = tp.Fill(&from)
		if fill_err != nil {
			return
		}
		var to common.Address
		fill_err = tp.Fill(&to)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if _ERC20 == nil || amount == nil {
			return
		}

		_ERC20.TransferFrom(from, to, amount)
	})
}

func Fuzz_Nosy_ERC20TransferIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ERC20TransferIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_ERC20TransferIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ERC20TransferIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_ERC20TransferIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *ERC20TransferIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

// skipping Fuzz_Nosy_GenericExpectedCall_Execute__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_GenericExpectedCall_Matches__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_GenericExpectedCall_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *GenericExpectedCall
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.String()
	})
}

// skipping Fuzz_Nosy_RpcStub_AddExpectedCall__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/sources/batching/test.ExpectedRpcCall

func Fuzz_Nosy_RpcStub_BatchCallContext__(f *testing.F) {
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
		if t1 == nil {
			return
		}

		r := NewRpcStub(t1)
		r.BatchCallContext(ctx, b)
	})
}

// skipping Fuzz_Nosy_RpcStub_CallContext__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_RpcStub_ClearResponses__(f *testing.F) {
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
		if t1 == nil {
			return
		}

		r := NewRpcStub(t1)
		r.ClearResponses()
	})
}

// skipping Fuzz_Nosy_RpcStub_findExpectedCall__ because parameters include func, chan, or unsupported interface: []interface{}

// skipping Fuzz_Nosy_expectedCall_Execute__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_expectedCall_Matches__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_expectedCall_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *expectedCall
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.String()
	})
}

// skipping Fuzz_Nosy_ExpectedRpcCall_Execute__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_ExpectedRpcCall_Matches__ because parameters include func, chan, or unsupported interface: []interface{}
