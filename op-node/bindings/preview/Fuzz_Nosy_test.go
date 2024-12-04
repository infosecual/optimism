package bindingspreview

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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

func Fuzz_Nosy_OptimismPortal2Caller_CheckWithdrawal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2Caller
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var _withdrawalHash [32]byte
		fill_err = tp.Fill(&_withdrawalHash)
		if fill_err != nil {
			return
		}
		var _proofSubmitter common.Address
		fill_err = tp.Fill(&_proofSubmitter)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil || opts == nil {
			return
		}

		_OptimismPortal2.CheckWithdrawal(opts, _withdrawalHash, _proofSubmitter)
	})
}

func Fuzz_Nosy_OptimismPortal2Caller_DisputeGameBlacklist__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2Caller
		fill_err = tp.Fill(&_OptimismPortal2)
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
		if _OptimismPortal2 == nil || opts == nil {
			return
		}

		_OptimismPortal2.DisputeGameBlacklist(opts, arg0)
	})
}

func Fuzz_Nosy_OptimismPortal2Caller_DisputeGameFactory__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2Caller
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil || opts == nil {
			return
		}

		_OptimismPortal2.DisputeGameFactory(opts)
	})
}

func Fuzz_Nosy_OptimismPortal2Caller_DisputeGameFinalityDelaySeconds__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2Caller
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil || opts == nil {
			return
		}

		_OptimismPortal2.DisputeGameFinalityDelaySeconds(opts)
	})
}

func Fuzz_Nosy_OptimismPortal2Caller_FinalizedWithdrawals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2Caller
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var arg0 [32]byte
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil || opts == nil {
			return
		}

		_OptimismPortal2.FinalizedWithdrawals(opts, arg0)
	})
}

func Fuzz_Nosy_OptimismPortal2Caller_Guardian__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2Caller
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil || opts == nil {
			return
		}

		_OptimismPortal2.Guardian(opts)
	})
}

func Fuzz_Nosy_OptimismPortal2Caller_L2Sender__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2Caller
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil || opts == nil {
			return
		}

		_OptimismPortal2.L2Sender(opts)
	})
}

func Fuzz_Nosy_OptimismPortal2Caller_MinimumGasLimit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2Caller
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var _byteCount uint64
		fill_err = tp.Fill(&_byteCount)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil || opts == nil {
			return
		}

		_OptimismPortal2.MinimumGasLimit(opts, _byteCount)
	})
}

func Fuzz_Nosy_OptimismPortal2Caller_NumProofSubmitters__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2Caller
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var _withdrawalHash [32]byte
		fill_err = tp.Fill(&_withdrawalHash)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil || opts == nil {
			return
		}

		_OptimismPortal2.NumProofSubmitters(opts, _withdrawalHash)
	})
}

func Fuzz_Nosy_OptimismPortal2Caller_Params__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2Caller
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil || opts == nil {
			return
		}

		_OptimismPortal2.Params(opts)
	})
}

func Fuzz_Nosy_OptimismPortal2Caller_Paused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2Caller
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil || opts == nil {
			return
		}

		_OptimismPortal2.Paused(opts)
	})
}

func Fuzz_Nosy_OptimismPortal2Caller_ProofMaturityDelaySeconds__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2Caller
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil || opts == nil {
			return
		}

		_OptimismPortal2.ProofMaturityDelaySeconds(opts)
	})
}

func Fuzz_Nosy_OptimismPortal2Caller_ProofSubmitters__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2Caller
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var arg0 [32]byte
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		var arg1 *big.Int
		fill_err = tp.Fill(&arg1)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil || opts == nil || arg1 == nil {
			return
		}

		_OptimismPortal2.ProofSubmitters(opts, arg0, arg1)
	})
}

func Fuzz_Nosy_OptimismPortal2Caller_ProvenWithdrawals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2Caller
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var arg0 [32]byte
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		var arg1 common.Address
		fill_err = tp.Fill(&arg1)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil || opts == nil {
			return
		}

		_OptimismPortal2.ProvenWithdrawals(opts, arg0, arg1)
	})
}

func Fuzz_Nosy_OptimismPortal2Caller_RespectedGameType__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2Caller
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil || opts == nil {
			return
		}

		_OptimismPortal2.RespectedGameType(opts)
	})
}

func Fuzz_Nosy_OptimismPortal2Caller_RespectedGameTypeUpdatedAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2Caller
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil || opts == nil {
			return
		}

		_OptimismPortal2.RespectedGameTypeUpdatedAt(opts)
	})
}

func Fuzz_Nosy_OptimismPortal2Caller_SuperchainConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2Caller
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil || opts == nil {
			return
		}

		_OptimismPortal2.SuperchainConfig(opts)
	})
}

func Fuzz_Nosy_OptimismPortal2Caller_SystemConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2Caller
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil || opts == nil {
			return
		}

		_OptimismPortal2.SystemConfig(opts)
	})
}

func Fuzz_Nosy_OptimismPortal2Caller_Version__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2Caller
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil || opts == nil {
			return
		}

		_OptimismPortal2.Version(opts)
	})
}

// skipping Fuzz_Nosy_OptimismPortal2CallerRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

func Fuzz_Nosy_OptimismPortal2CallerSession_CheckWithdrawal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2CallerSession
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		var _withdrawalHash [32]byte
		fill_err = tp.Fill(&_withdrawalHash)
		if fill_err != nil {
			return
		}
		var _proofSubmitter common.Address
		fill_err = tp.Fill(&_proofSubmitter)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil {
			return
		}

		_OptimismPortal2.CheckWithdrawal(_withdrawalHash, _proofSubmitter)
	})
}

func Fuzz_Nosy_OptimismPortal2CallerSession_DisputeGameBlacklist__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2CallerSession
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil {
			return
		}

		_OptimismPortal2.DisputeGameBlacklist(arg0)
	})
}

func Fuzz_Nosy_OptimismPortal2CallerSession_DisputeGameFactory__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2CallerSession
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil {
			return
		}

		_OptimismPortal2.DisputeGameFactory()
	})
}

func Fuzz_Nosy_OptimismPortal2CallerSession_DisputeGameFinalityDelaySeconds__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2CallerSession
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil {
			return
		}

		_OptimismPortal2.DisputeGameFinalityDelaySeconds()
	})
}

func Fuzz_Nosy_OptimismPortal2CallerSession_FinalizedWithdrawals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2CallerSession
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		var arg0 [32]byte
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil {
			return
		}

		_OptimismPortal2.FinalizedWithdrawals(arg0)
	})
}

func Fuzz_Nosy_OptimismPortal2CallerSession_Guardian__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2CallerSession
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil {
			return
		}

		_OptimismPortal2.Guardian()
	})
}

func Fuzz_Nosy_OptimismPortal2CallerSession_L2Sender__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2CallerSession
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil {
			return
		}

		_OptimismPortal2.L2Sender()
	})
}

func Fuzz_Nosy_OptimismPortal2CallerSession_MinimumGasLimit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2CallerSession
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		var _byteCount uint64
		fill_err = tp.Fill(&_byteCount)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil {
			return
		}

		_OptimismPortal2.MinimumGasLimit(_byteCount)
	})
}

func Fuzz_Nosy_OptimismPortal2CallerSession_NumProofSubmitters__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2CallerSession
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		var _withdrawalHash [32]byte
		fill_err = tp.Fill(&_withdrawalHash)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil {
			return
		}

		_OptimismPortal2.NumProofSubmitters(_withdrawalHash)
	})
}

func Fuzz_Nosy_OptimismPortal2CallerSession_Params__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2CallerSession
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil {
			return
		}

		_OptimismPortal2.Params()
	})
}

func Fuzz_Nosy_OptimismPortal2CallerSession_Paused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2CallerSession
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil {
			return
		}

		_OptimismPortal2.Paused()
	})
}

func Fuzz_Nosy_OptimismPortal2CallerSession_ProofMaturityDelaySeconds__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2CallerSession
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil {
			return
		}

		_OptimismPortal2.ProofMaturityDelaySeconds()
	})
}

func Fuzz_Nosy_OptimismPortal2CallerSession_ProofSubmitters__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2CallerSession
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		var arg0 [32]byte
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		var arg1 *big.Int
		fill_err = tp.Fill(&arg1)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil || arg1 == nil {
			return
		}

		_OptimismPortal2.ProofSubmitters(arg0, arg1)
	})
}

func Fuzz_Nosy_OptimismPortal2CallerSession_ProvenWithdrawals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2CallerSession
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		var arg0 [32]byte
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		var arg1 common.Address
		fill_err = tp.Fill(&arg1)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil {
			return
		}

		_OptimismPortal2.ProvenWithdrawals(arg0, arg1)
	})
}

func Fuzz_Nosy_OptimismPortal2CallerSession_RespectedGameType__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2CallerSession
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil {
			return
		}

		_OptimismPortal2.RespectedGameType()
	})
}

func Fuzz_Nosy_OptimismPortal2CallerSession_RespectedGameTypeUpdatedAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2CallerSession
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil {
			return
		}

		_OptimismPortal2.RespectedGameTypeUpdatedAt()
	})
}

func Fuzz_Nosy_OptimismPortal2CallerSession_SuperchainConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2CallerSession
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil {
			return
		}

		_OptimismPortal2.SuperchainConfig()
	})
}

func Fuzz_Nosy_OptimismPortal2CallerSession_SystemConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2CallerSession
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil {
			return
		}

		_OptimismPortal2.SystemConfig()
	})
}

func Fuzz_Nosy_OptimismPortal2CallerSession_Version__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2CallerSession
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil {
			return
		}

		_OptimismPortal2.Version()
	})
}

func Fuzz_Nosy_OptimismPortal2Filterer_FilterInitialized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2Filterer
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil || opts == nil {
			return
		}

		_OptimismPortal2.FilterInitialized(opts)
	})
}

func Fuzz_Nosy_OptimismPortal2Filterer_FilterTransactionDeposited__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2Filterer
		fill_err = tp.Fill(&_OptimismPortal2)
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
		var version []*big.Int
		fill_err = tp.Fill(&version)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil || opts == nil {
			return
		}

		_OptimismPortal2.FilterTransactionDeposited(opts, from, to, version)
	})
}

func Fuzz_Nosy_OptimismPortal2Filterer_FilterWithdrawalFinalized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2Filterer
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var withdrawalHash [][32]byte
		fill_err = tp.Fill(&withdrawalHash)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil || opts == nil {
			return
		}

		_OptimismPortal2.FilterWithdrawalFinalized(opts, withdrawalHash)
	})
}

func Fuzz_Nosy_OptimismPortal2Filterer_FilterWithdrawalProven__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2Filterer
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var withdrawalHash [][32]byte
		fill_err = tp.Fill(&withdrawalHash)
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
		if _OptimismPortal2 == nil || opts == nil {
			return
		}

		_OptimismPortal2.FilterWithdrawalProven(opts, withdrawalHash, from, to)
	})
}

func Fuzz_Nosy_OptimismPortal2Filterer_ParseInitialized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2Filterer
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil {
			return
		}

		_OptimismPortal2.ParseInitialized(log)
	})
}

func Fuzz_Nosy_OptimismPortal2Filterer_ParseTransactionDeposited__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2Filterer
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil {
			return
		}

		_OptimismPortal2.ParseTransactionDeposited(log)
	})
}

func Fuzz_Nosy_OptimismPortal2Filterer_ParseWithdrawalFinalized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2Filterer
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil {
			return
		}

		_OptimismPortal2.ParseWithdrawalFinalized(log)
	})
}

func Fuzz_Nosy_OptimismPortal2Filterer_ParseWithdrawalProven__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2Filterer
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil {
			return
		}

		_OptimismPortal2.ParseWithdrawalProven(log)
	})
}

// skipping Fuzz_Nosy_OptimismPortal2Filterer_WatchInitialized__ because parameters include func, chan, or unsupported interface: chan<- *github.com/ethereum-optimism/optimism/op-node/bindings/preview.OptimismPortal2Initialized

// skipping Fuzz_Nosy_OptimismPortal2Filterer_WatchTransactionDeposited__ because parameters include func, chan, or unsupported interface: chan<- *github.com/ethereum-optimism/optimism/op-node/bindings/preview.OptimismPortal2TransactionDeposited

// skipping Fuzz_Nosy_OptimismPortal2Filterer_WatchWithdrawalFinalized__ because parameters include func, chan, or unsupported interface: chan<- *github.com/ethereum-optimism/optimism/op-node/bindings/preview.OptimismPortal2WithdrawalFinalized

// skipping Fuzz_Nosy_OptimismPortal2Filterer_WatchWithdrawalProven__ because parameters include func, chan, or unsupported interface: chan<- *github.com/ethereum-optimism/optimism/op-node/bindings/preview.OptimismPortal2WithdrawalProven

func Fuzz_Nosy_OptimismPortal2InitializedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OptimismPortal2InitializedIterator
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

func Fuzz_Nosy_OptimismPortal2InitializedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OptimismPortal2InitializedIterator
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

func Fuzz_Nosy_OptimismPortal2InitializedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OptimismPortal2InitializedIterator
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

// skipping Fuzz_Nosy_OptimismPortal2Raw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

// skipping Fuzz_Nosy_OptimismPortal2Raw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_OptimismPortal2Raw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2Raw
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil || opts == nil {
			return
		}

		_OptimismPortal2.Transfer(opts)
	})
}

func Fuzz_Nosy_OptimismPortal2Session_BlacklistDisputeGame__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2Session
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		var _disputeGame common.Address
		fill_err = tp.Fill(&_disputeGame)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil {
			return
		}

		_OptimismPortal2.BlacklistDisputeGame(_disputeGame)
	})
}

func Fuzz_Nosy_OptimismPortal2Session_CheckWithdrawal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2Session
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		var _withdrawalHash [32]byte
		fill_err = tp.Fill(&_withdrawalHash)
		if fill_err != nil {
			return
		}
		var _proofSubmitter common.Address
		fill_err = tp.Fill(&_proofSubmitter)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil {
			return
		}

		_OptimismPortal2.CheckWithdrawal(_withdrawalHash, _proofSubmitter)
	})
}

func Fuzz_Nosy_OptimismPortal2Session_DepositTransaction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2Session
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		var _to common.Address
		fill_err = tp.Fill(&_to)
		if fill_err != nil {
			return
		}
		var _value *big.Int
		fill_err = tp.Fill(&_value)
		if fill_err != nil {
			return
		}
		var _gasLimit uint64
		fill_err = tp.Fill(&_gasLimit)
		if fill_err != nil {
			return
		}
		var _isCreation bool
		fill_err = tp.Fill(&_isCreation)
		if fill_err != nil {
			return
		}
		var _data []byte
		fill_err = tp.Fill(&_data)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil || _value == nil {
			return
		}

		_OptimismPortal2.DepositTransaction(_to, _value, _gasLimit, _isCreation, _data)
	})
}

func Fuzz_Nosy_OptimismPortal2Session_DisputeGameBlacklist__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2Session
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil {
			return
		}

		_OptimismPortal2.DisputeGameBlacklist(arg0)
	})
}

func Fuzz_Nosy_OptimismPortal2Session_DisputeGameFactory__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2Session
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil {
			return
		}

		_OptimismPortal2.DisputeGameFactory()
	})
}

func Fuzz_Nosy_OptimismPortal2Session_DisputeGameFinalityDelaySeconds__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2Session
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil {
			return
		}

		_OptimismPortal2.DisputeGameFinalityDelaySeconds()
	})
}

func Fuzz_Nosy_OptimismPortal2Session_DonateETH__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2Session
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil {
			return
		}

		_OptimismPortal2.DonateETH()
	})
}

func Fuzz_Nosy_OptimismPortal2Session_FinalizeWithdrawalTransaction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2Session
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		var _tx TypesWithdrawalTransaction
		fill_err = tp.Fill(&_tx)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil {
			return
		}

		_OptimismPortal2.FinalizeWithdrawalTransaction(_tx)
	})
}

func Fuzz_Nosy_OptimismPortal2Session_FinalizeWithdrawalTransactionExternalProof__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2Session
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		var _tx TypesWithdrawalTransaction
		fill_err = tp.Fill(&_tx)
		if fill_err != nil {
			return
		}
		var _proofSubmitter common.Address
		fill_err = tp.Fill(&_proofSubmitter)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil {
			return
		}

		_OptimismPortal2.FinalizeWithdrawalTransactionExternalProof(_tx, _proofSubmitter)
	})
}

func Fuzz_Nosy_OptimismPortal2Session_FinalizedWithdrawals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2Session
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		var arg0 [32]byte
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil {
			return
		}

		_OptimismPortal2.FinalizedWithdrawals(arg0)
	})
}

func Fuzz_Nosy_OptimismPortal2Session_Guardian__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2Session
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil {
			return
		}

		_OptimismPortal2.Guardian()
	})
}

func Fuzz_Nosy_OptimismPortal2Session_Initialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2Session
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		var _disputeGameFactory common.Address
		fill_err = tp.Fill(&_disputeGameFactory)
		if fill_err != nil {
			return
		}
		var _systemConfig common.Address
		fill_err = tp.Fill(&_systemConfig)
		if fill_err != nil {
			return
		}
		var _superchainConfig common.Address
		fill_err = tp.Fill(&_superchainConfig)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil {
			return
		}

		_OptimismPortal2.Initialize(_disputeGameFactory, _systemConfig, _superchainConfig)
	})
}

func Fuzz_Nosy_OptimismPortal2Session_L2Sender__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2Session
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil {
			return
		}

		_OptimismPortal2.L2Sender()
	})
}

func Fuzz_Nosy_OptimismPortal2Session_MinimumGasLimit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2Session
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		var _byteCount uint64
		fill_err = tp.Fill(&_byteCount)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil {
			return
		}

		_OptimismPortal2.MinimumGasLimit(_byteCount)
	})
}

func Fuzz_Nosy_OptimismPortal2Session_NumProofSubmitters__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2Session
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		var _withdrawalHash [32]byte
		fill_err = tp.Fill(&_withdrawalHash)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil {
			return
		}

		_OptimismPortal2.NumProofSubmitters(_withdrawalHash)
	})
}

func Fuzz_Nosy_OptimismPortal2Session_Params__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2Session
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil {
			return
		}

		_OptimismPortal2.Params()
	})
}

func Fuzz_Nosy_OptimismPortal2Session_Paused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2Session
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil {
			return
		}

		_OptimismPortal2.Paused()
	})
}

func Fuzz_Nosy_OptimismPortal2Session_ProofMaturityDelaySeconds__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2Session
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil {
			return
		}

		_OptimismPortal2.ProofMaturityDelaySeconds()
	})
}

func Fuzz_Nosy_OptimismPortal2Session_ProofSubmitters__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2Session
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		var arg0 [32]byte
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		var arg1 *big.Int
		fill_err = tp.Fill(&arg1)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil || arg1 == nil {
			return
		}

		_OptimismPortal2.ProofSubmitters(arg0, arg1)
	})
}

func Fuzz_Nosy_OptimismPortal2Session_ProveWithdrawalTransaction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2Session
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		var _tx TypesWithdrawalTransaction
		fill_err = tp.Fill(&_tx)
		if fill_err != nil {
			return
		}
		var _disputeGameIndex *big.Int
		fill_err = tp.Fill(&_disputeGameIndex)
		if fill_err != nil {
			return
		}
		var _outputRootProof TypesOutputRootProof
		fill_err = tp.Fill(&_outputRootProof)
		if fill_err != nil {
			return
		}
		var _withdrawalProof [][]byte
		fill_err = tp.Fill(&_withdrawalProof)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil || _disputeGameIndex == nil {
			return
		}

		_OptimismPortal2.ProveWithdrawalTransaction(_tx, _disputeGameIndex, _outputRootProof, _withdrawalProof)
	})
}

func Fuzz_Nosy_OptimismPortal2Session_ProvenWithdrawals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2Session
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		var arg0 [32]byte
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		var arg1 common.Address
		fill_err = tp.Fill(&arg1)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil {
			return
		}

		_OptimismPortal2.ProvenWithdrawals(arg0, arg1)
	})
}

func Fuzz_Nosy_OptimismPortal2Session_Receive__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2Session
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil {
			return
		}

		_OptimismPortal2.Receive()
	})
}

func Fuzz_Nosy_OptimismPortal2Session_RespectedGameType__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2Session
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil {
			return
		}

		_OptimismPortal2.RespectedGameType()
	})
}

func Fuzz_Nosy_OptimismPortal2Session_RespectedGameTypeUpdatedAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2Session
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil {
			return
		}

		_OptimismPortal2.RespectedGameTypeUpdatedAt()
	})
}

func Fuzz_Nosy_OptimismPortal2Session_SetRespectedGameType__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2Session
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		var _gameType uint32
		fill_err = tp.Fill(&_gameType)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil {
			return
		}

		_OptimismPortal2.SetRespectedGameType(_gameType)
	})
}

func Fuzz_Nosy_OptimismPortal2Session_SuperchainConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2Session
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil {
			return
		}

		_OptimismPortal2.SuperchainConfig()
	})
}

func Fuzz_Nosy_OptimismPortal2Session_SystemConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2Session
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil {
			return
		}

		_OptimismPortal2.SystemConfig()
	})
}

func Fuzz_Nosy_OptimismPortal2Session_Version__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2Session
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil {
			return
		}

		_OptimismPortal2.Version()
	})
}

func Fuzz_Nosy_OptimismPortal2TransactionDepositedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OptimismPortal2TransactionDepositedIterator
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

func Fuzz_Nosy_OptimismPortal2TransactionDepositedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OptimismPortal2TransactionDepositedIterator
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

func Fuzz_Nosy_OptimismPortal2TransactionDepositedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OptimismPortal2TransactionDepositedIterator
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

func Fuzz_Nosy_OptimismPortal2Transactor_BlacklistDisputeGame__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2Transactor
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var _disputeGame common.Address
		fill_err = tp.Fill(&_disputeGame)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil || opts == nil {
			return
		}

		_OptimismPortal2.BlacklistDisputeGame(opts, _disputeGame)
	})
}

func Fuzz_Nosy_OptimismPortal2Transactor_DepositTransaction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2Transactor
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var _to common.Address
		fill_err = tp.Fill(&_to)
		if fill_err != nil {
			return
		}
		var _value *big.Int
		fill_err = tp.Fill(&_value)
		if fill_err != nil {
			return
		}
		var _gasLimit uint64
		fill_err = tp.Fill(&_gasLimit)
		if fill_err != nil {
			return
		}
		var _isCreation bool
		fill_err = tp.Fill(&_isCreation)
		if fill_err != nil {
			return
		}
		var _data []byte
		fill_err = tp.Fill(&_data)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil || opts == nil || _value == nil {
			return
		}

		_OptimismPortal2.DepositTransaction(opts, _to, _value, _gasLimit, _isCreation, _data)
	})
}

func Fuzz_Nosy_OptimismPortal2Transactor_DonateETH__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2Transactor
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil || opts == nil {
			return
		}

		_OptimismPortal2.DonateETH(opts)
	})
}

func Fuzz_Nosy_OptimismPortal2Transactor_FinalizeWithdrawalTransaction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2Transactor
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var _tx TypesWithdrawalTransaction
		fill_err = tp.Fill(&_tx)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil || opts == nil {
			return
		}

		_OptimismPortal2.FinalizeWithdrawalTransaction(opts, _tx)
	})
}

func Fuzz_Nosy_OptimismPortal2Transactor_FinalizeWithdrawalTransactionExternalProof__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2Transactor
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var _tx TypesWithdrawalTransaction
		fill_err = tp.Fill(&_tx)
		if fill_err != nil {
			return
		}
		var _proofSubmitter common.Address
		fill_err = tp.Fill(&_proofSubmitter)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil || opts == nil {
			return
		}

		_OptimismPortal2.FinalizeWithdrawalTransactionExternalProof(opts, _tx, _proofSubmitter)
	})
}

func Fuzz_Nosy_OptimismPortal2Transactor_Initialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2Transactor
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var _disputeGameFactory common.Address
		fill_err = tp.Fill(&_disputeGameFactory)
		if fill_err != nil {
			return
		}
		var _systemConfig common.Address
		fill_err = tp.Fill(&_systemConfig)
		if fill_err != nil {
			return
		}
		var _superchainConfig common.Address
		fill_err = tp.Fill(&_superchainConfig)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil || opts == nil {
			return
		}

		_OptimismPortal2.Initialize(opts, _disputeGameFactory, _systemConfig, _superchainConfig)
	})
}

func Fuzz_Nosy_OptimismPortal2Transactor_ProveWithdrawalTransaction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2Transactor
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var _tx TypesWithdrawalTransaction
		fill_err = tp.Fill(&_tx)
		if fill_err != nil {
			return
		}
		var _disputeGameIndex *big.Int
		fill_err = tp.Fill(&_disputeGameIndex)
		if fill_err != nil {
			return
		}
		var _outputRootProof TypesOutputRootProof
		fill_err = tp.Fill(&_outputRootProof)
		if fill_err != nil {
			return
		}
		var _withdrawalProof [][]byte
		fill_err = tp.Fill(&_withdrawalProof)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil || opts == nil || _disputeGameIndex == nil {
			return
		}

		_OptimismPortal2.ProveWithdrawalTransaction(opts, _tx, _disputeGameIndex, _outputRootProof, _withdrawalProof)
	})
}

func Fuzz_Nosy_OptimismPortal2Transactor_Receive__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2Transactor
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil || opts == nil {
			return
		}

		_OptimismPortal2.Receive(opts)
	})
}

func Fuzz_Nosy_OptimismPortal2Transactor_SetRespectedGameType__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2Transactor
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var _gameType uint32
		fill_err = tp.Fill(&_gameType)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil || opts == nil {
			return
		}

		_OptimismPortal2.SetRespectedGameType(opts, _gameType)
	})
}

// skipping Fuzz_Nosy_OptimismPortal2TransactorRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_OptimismPortal2TransactorRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2TransactorRaw
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil || opts == nil {
			return
		}

		_OptimismPortal2.Transfer(opts)
	})
}

func Fuzz_Nosy_OptimismPortal2TransactorSession_BlacklistDisputeGame__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2TransactorSession
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		var _disputeGame common.Address
		fill_err = tp.Fill(&_disputeGame)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil {
			return
		}

		_OptimismPortal2.BlacklistDisputeGame(_disputeGame)
	})
}

func Fuzz_Nosy_OptimismPortal2TransactorSession_DepositTransaction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2TransactorSession
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		var _to common.Address
		fill_err = tp.Fill(&_to)
		if fill_err != nil {
			return
		}
		var _value *big.Int
		fill_err = tp.Fill(&_value)
		if fill_err != nil {
			return
		}
		var _gasLimit uint64
		fill_err = tp.Fill(&_gasLimit)
		if fill_err != nil {
			return
		}
		var _isCreation bool
		fill_err = tp.Fill(&_isCreation)
		if fill_err != nil {
			return
		}
		var _data []byte
		fill_err = tp.Fill(&_data)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil || _value == nil {
			return
		}

		_OptimismPortal2.DepositTransaction(_to, _value, _gasLimit, _isCreation, _data)
	})
}

func Fuzz_Nosy_OptimismPortal2TransactorSession_DonateETH__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2TransactorSession
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil {
			return
		}

		_OptimismPortal2.DonateETH()
	})
}

func Fuzz_Nosy_OptimismPortal2TransactorSession_FinalizeWithdrawalTransaction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2TransactorSession
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		var _tx TypesWithdrawalTransaction
		fill_err = tp.Fill(&_tx)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil {
			return
		}

		_OptimismPortal2.FinalizeWithdrawalTransaction(_tx)
	})
}

func Fuzz_Nosy_OptimismPortal2TransactorSession_FinalizeWithdrawalTransactionExternalProof__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2TransactorSession
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		var _tx TypesWithdrawalTransaction
		fill_err = tp.Fill(&_tx)
		if fill_err != nil {
			return
		}
		var _proofSubmitter common.Address
		fill_err = tp.Fill(&_proofSubmitter)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil {
			return
		}

		_OptimismPortal2.FinalizeWithdrawalTransactionExternalProof(_tx, _proofSubmitter)
	})
}

func Fuzz_Nosy_OptimismPortal2TransactorSession_Initialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2TransactorSession
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		var _disputeGameFactory common.Address
		fill_err = tp.Fill(&_disputeGameFactory)
		if fill_err != nil {
			return
		}
		var _systemConfig common.Address
		fill_err = tp.Fill(&_systemConfig)
		if fill_err != nil {
			return
		}
		var _superchainConfig common.Address
		fill_err = tp.Fill(&_superchainConfig)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil {
			return
		}

		_OptimismPortal2.Initialize(_disputeGameFactory, _systemConfig, _superchainConfig)
	})
}

func Fuzz_Nosy_OptimismPortal2TransactorSession_ProveWithdrawalTransaction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2TransactorSession
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		var _tx TypesWithdrawalTransaction
		fill_err = tp.Fill(&_tx)
		if fill_err != nil {
			return
		}
		var _disputeGameIndex *big.Int
		fill_err = tp.Fill(&_disputeGameIndex)
		if fill_err != nil {
			return
		}
		var _outputRootProof TypesOutputRootProof
		fill_err = tp.Fill(&_outputRootProof)
		if fill_err != nil {
			return
		}
		var _withdrawalProof [][]byte
		fill_err = tp.Fill(&_withdrawalProof)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil || _disputeGameIndex == nil {
			return
		}

		_OptimismPortal2.ProveWithdrawalTransaction(_tx, _disputeGameIndex, _outputRootProof, _withdrawalProof)
	})
}

func Fuzz_Nosy_OptimismPortal2TransactorSession_Receive__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2TransactorSession
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil {
			return
		}

		_OptimismPortal2.Receive()
	})
}

func Fuzz_Nosy_OptimismPortal2TransactorSession_SetRespectedGameType__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal2 *OptimismPortal2TransactorSession
		fill_err = tp.Fill(&_OptimismPortal2)
		if fill_err != nil {
			return
		}
		var _gameType uint32
		fill_err = tp.Fill(&_gameType)
		if fill_err != nil {
			return
		}
		if _OptimismPortal2 == nil {
			return
		}

		_OptimismPortal2.SetRespectedGameType(_gameType)
	})
}

func Fuzz_Nosy_OptimismPortal2WithdrawalFinalizedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OptimismPortal2WithdrawalFinalizedIterator
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

func Fuzz_Nosy_OptimismPortal2WithdrawalFinalizedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OptimismPortal2WithdrawalFinalizedIterator
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

func Fuzz_Nosy_OptimismPortal2WithdrawalFinalizedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OptimismPortal2WithdrawalFinalizedIterator
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

func Fuzz_Nosy_OptimismPortal2WithdrawalProvenIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OptimismPortal2WithdrawalProvenIterator
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

func Fuzz_Nosy_OptimismPortal2WithdrawalProvenIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OptimismPortal2WithdrawalProvenIterator
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

func Fuzz_Nosy_OptimismPortal2WithdrawalProvenIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OptimismPortal2WithdrawalProvenIterator
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

// skipping Fuzz_Nosy_DeployOptimismPortal2__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/abi/bind.ContractBackend
