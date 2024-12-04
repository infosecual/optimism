package bindings

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

func Fuzz_Nosy_DataAvailabilityChallengeBalanceChangedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DataAvailabilityChallengeBalanceChangedIterator
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

func Fuzz_Nosy_DataAvailabilityChallengeBalanceChangedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DataAvailabilityChallengeBalanceChangedIterator
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

func Fuzz_Nosy_DataAvailabilityChallengeBalanceChangedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DataAvailabilityChallengeBalanceChangedIterator
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

func Fuzz_Nosy_DataAvailabilityChallengeCaller_Balances__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeCaller
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
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
		if _DataAvailabilityChallenge == nil || opts == nil {
			return
		}

		_DataAvailabilityChallenge.Balances(opts, arg0)
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeCaller_BondSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeCaller
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil || opts == nil {
			return
		}

		_DataAvailabilityChallenge.BondSize(opts)
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeCaller_ChallengeWindow__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeCaller
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil || opts == nil {
			return
		}

		_DataAvailabilityChallenge.ChallengeWindow(opts)
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeCaller_FixedResolutionCost__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeCaller
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil || opts == nil {
			return
		}

		_DataAvailabilityChallenge.FixedResolutionCost(opts)
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeCaller_GetChallenge__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeCaller
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var challengedBlockNumber *big.Int
		fill_err = tp.Fill(&challengedBlockNumber)
		if fill_err != nil {
			return
		}
		var challengedCommitment []byte
		fill_err = tp.Fill(&challengedCommitment)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil || opts == nil || challengedBlockNumber == nil {
			return
		}

		_DataAvailabilityChallenge.GetChallenge(opts, challengedBlockNumber, challengedCommitment)
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeCaller_GetChallengeStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeCaller
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var challengedBlockNumber *big.Int
		fill_err = tp.Fill(&challengedBlockNumber)
		if fill_err != nil {
			return
		}
		var challengedCommitment []byte
		fill_err = tp.Fill(&challengedCommitment)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil || opts == nil || challengedBlockNumber == nil {
			return
		}

		_DataAvailabilityChallenge.GetChallengeStatus(opts, challengedBlockNumber, challengedCommitment)
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeCaller_Owner__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeCaller
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil || opts == nil {
			return
		}

		_DataAvailabilityChallenge.Owner(opts)
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeCaller_ResolveWindow__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeCaller
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil || opts == nil {
			return
		}

		_DataAvailabilityChallenge.ResolveWindow(opts)
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeCaller_ResolverRefundPercentage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeCaller
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil || opts == nil {
			return
		}

		_DataAvailabilityChallenge.ResolverRefundPercentage(opts)
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeCaller_ValidateCommitment__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeCaller
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var commitment []byte
		fill_err = tp.Fill(&commitment)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil || opts == nil {
			return
		}

		_DataAvailabilityChallenge.ValidateCommitment(opts, commitment)
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeCaller_VariableResolutionCost__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeCaller
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil || opts == nil {
			return
		}

		_DataAvailabilityChallenge.VariableResolutionCost(opts)
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeCaller_VariableResolutionCostPrecision__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeCaller
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil || opts == nil {
			return
		}

		_DataAvailabilityChallenge.VariableResolutionCostPrecision(opts)
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeCaller_Version__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeCaller
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil || opts == nil {
			return
		}

		_DataAvailabilityChallenge.Version(opts)
	})
}

// skipping Fuzz_Nosy_DataAvailabilityChallengeCallerRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

func Fuzz_Nosy_DataAvailabilityChallengeCallerSession_Balances__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeCallerSession
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil {
			return
		}

		_DataAvailabilityChallenge.Balances(arg0)
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeCallerSession_BondSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeCallerSession
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil {
			return
		}

		_DataAvailabilityChallenge.BondSize()
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeCallerSession_ChallengeWindow__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeCallerSession
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil {
			return
		}

		_DataAvailabilityChallenge.ChallengeWindow()
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeCallerSession_FixedResolutionCost__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeCallerSession
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil {
			return
		}

		_DataAvailabilityChallenge.FixedResolutionCost()
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeCallerSession_GetChallenge__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeCallerSession
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		var challengedBlockNumber *big.Int
		fill_err = tp.Fill(&challengedBlockNumber)
		if fill_err != nil {
			return
		}
		var challengedCommitment []byte
		fill_err = tp.Fill(&challengedCommitment)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil || challengedBlockNumber == nil {
			return
		}

		_DataAvailabilityChallenge.GetChallenge(challengedBlockNumber, challengedCommitment)
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeCallerSession_GetChallengeStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeCallerSession
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		var challengedBlockNumber *big.Int
		fill_err = tp.Fill(&challengedBlockNumber)
		if fill_err != nil {
			return
		}
		var challengedCommitment []byte
		fill_err = tp.Fill(&challengedCommitment)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil || challengedBlockNumber == nil {
			return
		}

		_DataAvailabilityChallenge.GetChallengeStatus(challengedBlockNumber, challengedCommitment)
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeCallerSession_Owner__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeCallerSession
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil {
			return
		}

		_DataAvailabilityChallenge.Owner()
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeCallerSession_ResolveWindow__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeCallerSession
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil {
			return
		}

		_DataAvailabilityChallenge.ResolveWindow()
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeCallerSession_ResolverRefundPercentage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeCallerSession
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil {
			return
		}

		_DataAvailabilityChallenge.ResolverRefundPercentage()
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeCallerSession_ValidateCommitment__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeCallerSession
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		var commitment []byte
		fill_err = tp.Fill(&commitment)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil {
			return
		}

		_DataAvailabilityChallenge.ValidateCommitment(commitment)
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeCallerSession_VariableResolutionCost__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeCallerSession
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil {
			return
		}

		_DataAvailabilityChallenge.VariableResolutionCost()
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeCallerSession_VariableResolutionCostPrecision__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeCallerSession
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil {
			return
		}

		_DataAvailabilityChallenge.VariableResolutionCostPrecision()
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeCallerSession_Version__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeCallerSession
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil {
			return
		}

		_DataAvailabilityChallenge.Version()
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeChallengeStatusChangedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DataAvailabilityChallengeChallengeStatusChangedIterator
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

func Fuzz_Nosy_DataAvailabilityChallengeChallengeStatusChangedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DataAvailabilityChallengeChallengeStatusChangedIterator
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

func Fuzz_Nosy_DataAvailabilityChallengeChallengeStatusChangedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DataAvailabilityChallengeChallengeStatusChangedIterator
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

func Fuzz_Nosy_DataAvailabilityChallengeFilterer_FilterBalanceChanged__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeFilterer
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil || opts == nil {
			return
		}

		_DataAvailabilityChallenge.FilterBalanceChanged(opts)
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeFilterer_FilterChallengeStatusChanged__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeFilterer
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var challengedBlockNumber []*big.Int
		fill_err = tp.Fill(&challengedBlockNumber)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil || opts == nil {
			return
		}

		_DataAvailabilityChallenge.FilterChallengeStatusChanged(opts, challengedBlockNumber)
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeFilterer_FilterInitialized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeFilterer
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil || opts == nil {
			return
		}

		_DataAvailabilityChallenge.FilterInitialized(opts)
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeFilterer_FilterOwnershipTransferred__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeFilterer
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var previousOwner []common.Address
		fill_err = tp.Fill(&previousOwner)
		if fill_err != nil {
			return
		}
		var newOwner []common.Address
		fill_err = tp.Fill(&newOwner)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil || opts == nil {
			return
		}

		_DataAvailabilityChallenge.FilterOwnershipTransferred(opts, previousOwner, newOwner)
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeFilterer_FilterRequiredBondSizeChanged__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeFilterer
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil || opts == nil {
			return
		}

		_DataAvailabilityChallenge.FilterRequiredBondSizeChanged(opts)
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeFilterer_FilterResolverRefundPercentageChanged__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeFilterer
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil || opts == nil {
			return
		}

		_DataAvailabilityChallenge.FilterResolverRefundPercentageChanged(opts)
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeFilterer_ParseBalanceChanged__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeFilterer
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil {
			return
		}

		_DataAvailabilityChallenge.ParseBalanceChanged(log)
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeFilterer_ParseChallengeStatusChanged__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeFilterer
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil {
			return
		}

		_DataAvailabilityChallenge.ParseChallengeStatusChanged(log)
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeFilterer_ParseInitialized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeFilterer
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil {
			return
		}

		_DataAvailabilityChallenge.ParseInitialized(log)
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeFilterer_ParseOwnershipTransferred__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeFilterer
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil {
			return
		}

		_DataAvailabilityChallenge.ParseOwnershipTransferred(log)
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeFilterer_ParseRequiredBondSizeChanged__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeFilterer
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil {
			return
		}

		_DataAvailabilityChallenge.ParseRequiredBondSizeChanged(log)
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeFilterer_ParseResolverRefundPercentageChanged__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeFilterer
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil {
			return
		}

		_DataAvailabilityChallenge.ParseResolverRefundPercentageChanged(log)
	})
}

// skipping Fuzz_Nosy_DataAvailabilityChallengeFilterer_WatchBalanceChanged__ because parameters include func, chan, or unsupported interface: chan<- *github.com/ethereum-optimism/optimism/op-alt-da/bindings.DataAvailabilityChallengeBalanceChanged

// skipping Fuzz_Nosy_DataAvailabilityChallengeFilterer_WatchChallengeStatusChanged__ because parameters include func, chan, or unsupported interface: chan<- *github.com/ethereum-optimism/optimism/op-alt-da/bindings.DataAvailabilityChallengeChallengeStatusChanged

// skipping Fuzz_Nosy_DataAvailabilityChallengeFilterer_WatchInitialized__ because parameters include func, chan, or unsupported interface: chan<- *github.com/ethereum-optimism/optimism/op-alt-da/bindings.DataAvailabilityChallengeInitialized

// skipping Fuzz_Nosy_DataAvailabilityChallengeFilterer_WatchOwnershipTransferred__ because parameters include func, chan, or unsupported interface: chan<- *github.com/ethereum-optimism/optimism/op-alt-da/bindings.DataAvailabilityChallengeOwnershipTransferred

// skipping Fuzz_Nosy_DataAvailabilityChallengeFilterer_WatchRequiredBondSizeChanged__ because parameters include func, chan, or unsupported interface: chan<- *github.com/ethereum-optimism/optimism/op-alt-da/bindings.DataAvailabilityChallengeRequiredBondSizeChanged

// skipping Fuzz_Nosy_DataAvailabilityChallengeFilterer_WatchResolverRefundPercentageChanged__ because parameters include func, chan, or unsupported interface: chan<- *github.com/ethereum-optimism/optimism/op-alt-da/bindings.DataAvailabilityChallengeResolverRefundPercentageChanged

func Fuzz_Nosy_DataAvailabilityChallengeInitializedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DataAvailabilityChallengeInitializedIterator
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

func Fuzz_Nosy_DataAvailabilityChallengeInitializedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DataAvailabilityChallengeInitializedIterator
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

func Fuzz_Nosy_DataAvailabilityChallengeInitializedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DataAvailabilityChallengeInitializedIterator
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

func Fuzz_Nosy_DataAvailabilityChallengeOwnershipTransferredIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DataAvailabilityChallengeOwnershipTransferredIterator
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

func Fuzz_Nosy_DataAvailabilityChallengeOwnershipTransferredIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DataAvailabilityChallengeOwnershipTransferredIterator
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

func Fuzz_Nosy_DataAvailabilityChallengeOwnershipTransferredIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DataAvailabilityChallengeOwnershipTransferredIterator
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

// skipping Fuzz_Nosy_DataAvailabilityChallengeRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

// skipping Fuzz_Nosy_DataAvailabilityChallengeRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_DataAvailabilityChallengeRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeRaw
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil || opts == nil {
			return
		}

		_DataAvailabilityChallenge.Transfer(opts)
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeRequiredBondSizeChangedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DataAvailabilityChallengeRequiredBondSizeChangedIterator
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

func Fuzz_Nosy_DataAvailabilityChallengeRequiredBondSizeChangedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DataAvailabilityChallengeRequiredBondSizeChangedIterator
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

func Fuzz_Nosy_DataAvailabilityChallengeRequiredBondSizeChangedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DataAvailabilityChallengeRequiredBondSizeChangedIterator
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

func Fuzz_Nosy_DataAvailabilityChallengeResolverRefundPercentageChangedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DataAvailabilityChallengeResolverRefundPercentageChangedIterator
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

func Fuzz_Nosy_DataAvailabilityChallengeResolverRefundPercentageChangedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DataAvailabilityChallengeResolverRefundPercentageChangedIterator
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

func Fuzz_Nosy_DataAvailabilityChallengeResolverRefundPercentageChangedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DataAvailabilityChallengeResolverRefundPercentageChangedIterator
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

func Fuzz_Nosy_DataAvailabilityChallengeSession_Balances__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeSession
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		var arg0 common.Address
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil {
			return
		}

		_DataAvailabilityChallenge.Balances(arg0)
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeSession_BondSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeSession
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil {
			return
		}

		_DataAvailabilityChallenge.BondSize()
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeSession_Challenge__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeSession
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		var challengedBlockNumber *big.Int
		fill_err = tp.Fill(&challengedBlockNumber)
		if fill_err != nil {
			return
		}
		var challengedCommitment []byte
		fill_err = tp.Fill(&challengedCommitment)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil || challengedBlockNumber == nil {
			return
		}

		_DataAvailabilityChallenge.Challenge(challengedBlockNumber, challengedCommitment)
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeSession_ChallengeWindow__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeSession
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil {
			return
		}

		_DataAvailabilityChallenge.ChallengeWindow()
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeSession_Deposit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeSession
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil {
			return
		}

		_DataAvailabilityChallenge.Deposit()
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeSession_FixedResolutionCost__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeSession
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil {
			return
		}

		_DataAvailabilityChallenge.FixedResolutionCost()
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeSession_GetChallenge__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeSession
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		var challengedBlockNumber *big.Int
		fill_err = tp.Fill(&challengedBlockNumber)
		if fill_err != nil {
			return
		}
		var challengedCommitment []byte
		fill_err = tp.Fill(&challengedCommitment)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil || challengedBlockNumber == nil {
			return
		}

		_DataAvailabilityChallenge.GetChallenge(challengedBlockNumber, challengedCommitment)
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeSession_GetChallengeStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeSession
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		var challengedBlockNumber *big.Int
		fill_err = tp.Fill(&challengedBlockNumber)
		if fill_err != nil {
			return
		}
		var challengedCommitment []byte
		fill_err = tp.Fill(&challengedCommitment)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil || challengedBlockNumber == nil {
			return
		}

		_DataAvailabilityChallenge.GetChallengeStatus(challengedBlockNumber, challengedCommitment)
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeSession_Initialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeSession
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		var _owner common.Address
		fill_err = tp.Fill(&_owner)
		if fill_err != nil {
			return
		}
		var _challengeWindow *big.Int
		fill_err = tp.Fill(&_challengeWindow)
		if fill_err != nil {
			return
		}
		var _resolveWindow *big.Int
		fill_err = tp.Fill(&_resolveWindow)
		if fill_err != nil {
			return
		}
		var _bondSize *big.Int
		fill_err = tp.Fill(&_bondSize)
		if fill_err != nil {
			return
		}
		var _resolverRefundPercentage *big.Int
		fill_err = tp.Fill(&_resolverRefundPercentage)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil || _challengeWindow == nil || _resolveWindow == nil || _bondSize == nil || _resolverRefundPercentage == nil {
			return
		}

		_DataAvailabilityChallenge.Initialize(_owner, _challengeWindow, _resolveWindow, _bondSize, _resolverRefundPercentage)
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeSession_Owner__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeSession
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil {
			return
		}

		_DataAvailabilityChallenge.Owner()
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeSession_Receive__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeSession
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil {
			return
		}

		_DataAvailabilityChallenge.Receive()
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeSession_RenounceOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeSession
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil {
			return
		}

		_DataAvailabilityChallenge.RenounceOwnership()
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeSession_Resolve__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeSession
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		var challengedBlockNumber *big.Int
		fill_err = tp.Fill(&challengedBlockNumber)
		if fill_err != nil {
			return
		}
		var challengedCommitment []byte
		fill_err = tp.Fill(&challengedCommitment)
		if fill_err != nil {
			return
		}
		var resolveData []byte
		fill_err = tp.Fill(&resolveData)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil || challengedBlockNumber == nil {
			return
		}

		_DataAvailabilityChallenge.Resolve(challengedBlockNumber, challengedCommitment, resolveData)
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeSession_ResolveWindow__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeSession
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil {
			return
		}

		_DataAvailabilityChallenge.ResolveWindow()
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeSession_ResolverRefundPercentage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeSession
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil {
			return
		}

		_DataAvailabilityChallenge.ResolverRefundPercentage()
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeSession_SetBondSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeSession
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		var _bondSize *big.Int
		fill_err = tp.Fill(&_bondSize)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil || _bondSize == nil {
			return
		}

		_DataAvailabilityChallenge.SetBondSize(_bondSize)
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeSession_SetResolverRefundPercentage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeSession
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		var _resolverRefundPercentage *big.Int
		fill_err = tp.Fill(&_resolverRefundPercentage)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil || _resolverRefundPercentage == nil {
			return
		}

		_DataAvailabilityChallenge.SetResolverRefundPercentage(_resolverRefundPercentage)
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeSession_TransferOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeSession
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		var newOwner common.Address
		fill_err = tp.Fill(&newOwner)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil {
			return
		}

		_DataAvailabilityChallenge.TransferOwnership(newOwner)
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeSession_UnlockBond__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeSession
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		var challengedBlockNumber *big.Int
		fill_err = tp.Fill(&challengedBlockNumber)
		if fill_err != nil {
			return
		}
		var challengedCommitment []byte
		fill_err = tp.Fill(&challengedCommitment)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil || challengedBlockNumber == nil {
			return
		}

		_DataAvailabilityChallenge.UnlockBond(challengedBlockNumber, challengedCommitment)
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeSession_ValidateCommitment__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeSession
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		var commitment []byte
		fill_err = tp.Fill(&commitment)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil {
			return
		}

		_DataAvailabilityChallenge.ValidateCommitment(commitment)
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeSession_VariableResolutionCost__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeSession
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil {
			return
		}

		_DataAvailabilityChallenge.VariableResolutionCost()
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeSession_VariableResolutionCostPrecision__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeSession
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil {
			return
		}

		_DataAvailabilityChallenge.VariableResolutionCostPrecision()
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeSession_Version__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeSession
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil {
			return
		}

		_DataAvailabilityChallenge.Version()
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeSession_Withdraw__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeSession
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil {
			return
		}

		_DataAvailabilityChallenge.Withdraw()
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeTransactor_Challenge__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeTransactor
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var challengedBlockNumber *big.Int
		fill_err = tp.Fill(&challengedBlockNumber)
		if fill_err != nil {
			return
		}
		var challengedCommitment []byte
		fill_err = tp.Fill(&challengedCommitment)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil || opts == nil || challengedBlockNumber == nil {
			return
		}

		_DataAvailabilityChallenge.Challenge(opts, challengedBlockNumber, challengedCommitment)
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeTransactor_Deposit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeTransactor
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil || opts == nil {
			return
		}

		_DataAvailabilityChallenge.Deposit(opts)
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeTransactor_Initialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeTransactor
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var _owner common.Address
		fill_err = tp.Fill(&_owner)
		if fill_err != nil {
			return
		}
		var _challengeWindow *big.Int
		fill_err = tp.Fill(&_challengeWindow)
		if fill_err != nil {
			return
		}
		var _resolveWindow *big.Int
		fill_err = tp.Fill(&_resolveWindow)
		if fill_err != nil {
			return
		}
		var _bondSize *big.Int
		fill_err = tp.Fill(&_bondSize)
		if fill_err != nil {
			return
		}
		var _resolverRefundPercentage *big.Int
		fill_err = tp.Fill(&_resolverRefundPercentage)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil || opts == nil || _challengeWindow == nil || _resolveWindow == nil || _bondSize == nil || _resolverRefundPercentage == nil {
			return
		}

		_DataAvailabilityChallenge.Initialize(opts, _owner, _challengeWindow, _resolveWindow, _bondSize, _resolverRefundPercentage)
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeTransactor_Receive__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeTransactor
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil || opts == nil {
			return
		}

		_DataAvailabilityChallenge.Receive(opts)
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeTransactor_RenounceOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeTransactor
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil || opts == nil {
			return
		}

		_DataAvailabilityChallenge.RenounceOwnership(opts)
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeTransactor_Resolve__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeTransactor
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var challengedBlockNumber *big.Int
		fill_err = tp.Fill(&challengedBlockNumber)
		if fill_err != nil {
			return
		}
		var challengedCommitment []byte
		fill_err = tp.Fill(&challengedCommitment)
		if fill_err != nil {
			return
		}
		var resolveData []byte
		fill_err = tp.Fill(&resolveData)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil || opts == nil || challengedBlockNumber == nil {
			return
		}

		_DataAvailabilityChallenge.Resolve(opts, challengedBlockNumber, challengedCommitment, resolveData)
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeTransactor_SetBondSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeTransactor
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var _bondSize *big.Int
		fill_err = tp.Fill(&_bondSize)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil || opts == nil || _bondSize == nil {
			return
		}

		_DataAvailabilityChallenge.SetBondSize(opts, _bondSize)
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeTransactor_SetResolverRefundPercentage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeTransactor
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var _resolverRefundPercentage *big.Int
		fill_err = tp.Fill(&_resolverRefundPercentage)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil || opts == nil || _resolverRefundPercentage == nil {
			return
		}

		_DataAvailabilityChallenge.SetResolverRefundPercentage(opts, _resolverRefundPercentage)
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeTransactor_TransferOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeTransactor
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var newOwner common.Address
		fill_err = tp.Fill(&newOwner)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil || opts == nil {
			return
		}

		_DataAvailabilityChallenge.TransferOwnership(opts, newOwner)
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeTransactor_UnlockBond__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeTransactor
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var challengedBlockNumber *big.Int
		fill_err = tp.Fill(&challengedBlockNumber)
		if fill_err != nil {
			return
		}
		var challengedCommitment []byte
		fill_err = tp.Fill(&challengedCommitment)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil || opts == nil || challengedBlockNumber == nil {
			return
		}

		_DataAvailabilityChallenge.UnlockBond(opts, challengedBlockNumber, challengedCommitment)
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeTransactor_Withdraw__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeTransactor
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil || opts == nil {
			return
		}

		_DataAvailabilityChallenge.Withdraw(opts)
	})
}

// skipping Fuzz_Nosy_DataAvailabilityChallengeTransactorRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_DataAvailabilityChallengeTransactorRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeTransactorRaw
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil || opts == nil {
			return
		}

		_DataAvailabilityChallenge.Transfer(opts)
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeTransactorSession_Challenge__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeTransactorSession
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		var challengedBlockNumber *big.Int
		fill_err = tp.Fill(&challengedBlockNumber)
		if fill_err != nil {
			return
		}
		var challengedCommitment []byte
		fill_err = tp.Fill(&challengedCommitment)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil || challengedBlockNumber == nil {
			return
		}

		_DataAvailabilityChallenge.Challenge(challengedBlockNumber, challengedCommitment)
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeTransactorSession_Deposit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeTransactorSession
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil {
			return
		}

		_DataAvailabilityChallenge.Deposit()
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeTransactorSession_Initialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeTransactorSession
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		var _owner common.Address
		fill_err = tp.Fill(&_owner)
		if fill_err != nil {
			return
		}
		var _challengeWindow *big.Int
		fill_err = tp.Fill(&_challengeWindow)
		if fill_err != nil {
			return
		}
		var _resolveWindow *big.Int
		fill_err = tp.Fill(&_resolveWindow)
		if fill_err != nil {
			return
		}
		var _bondSize *big.Int
		fill_err = tp.Fill(&_bondSize)
		if fill_err != nil {
			return
		}
		var _resolverRefundPercentage *big.Int
		fill_err = tp.Fill(&_resolverRefundPercentage)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil || _challengeWindow == nil || _resolveWindow == nil || _bondSize == nil || _resolverRefundPercentage == nil {
			return
		}

		_DataAvailabilityChallenge.Initialize(_owner, _challengeWindow, _resolveWindow, _bondSize, _resolverRefundPercentage)
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeTransactorSession_Receive__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeTransactorSession
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil {
			return
		}

		_DataAvailabilityChallenge.Receive()
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeTransactorSession_RenounceOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeTransactorSession
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil {
			return
		}

		_DataAvailabilityChallenge.RenounceOwnership()
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeTransactorSession_Resolve__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeTransactorSession
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		var challengedBlockNumber *big.Int
		fill_err = tp.Fill(&challengedBlockNumber)
		if fill_err != nil {
			return
		}
		var challengedCommitment []byte
		fill_err = tp.Fill(&challengedCommitment)
		if fill_err != nil {
			return
		}
		var resolveData []byte
		fill_err = tp.Fill(&resolveData)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil || challengedBlockNumber == nil {
			return
		}

		_DataAvailabilityChallenge.Resolve(challengedBlockNumber, challengedCommitment, resolveData)
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeTransactorSession_SetBondSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeTransactorSession
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		var _bondSize *big.Int
		fill_err = tp.Fill(&_bondSize)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil || _bondSize == nil {
			return
		}

		_DataAvailabilityChallenge.SetBondSize(_bondSize)
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeTransactorSession_SetResolverRefundPercentage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeTransactorSession
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		var _resolverRefundPercentage *big.Int
		fill_err = tp.Fill(&_resolverRefundPercentage)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil || _resolverRefundPercentage == nil {
			return
		}

		_DataAvailabilityChallenge.SetResolverRefundPercentage(_resolverRefundPercentage)
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeTransactorSession_TransferOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeTransactorSession
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		var newOwner common.Address
		fill_err = tp.Fill(&newOwner)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil {
			return
		}

		_DataAvailabilityChallenge.TransferOwnership(newOwner)
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeTransactorSession_UnlockBond__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeTransactorSession
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		var challengedBlockNumber *big.Int
		fill_err = tp.Fill(&challengedBlockNumber)
		if fill_err != nil {
			return
		}
		var challengedCommitment []byte
		fill_err = tp.Fill(&challengedCommitment)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil || challengedBlockNumber == nil {
			return
		}

		_DataAvailabilityChallenge.UnlockBond(challengedBlockNumber, challengedCommitment)
	})
}

func Fuzz_Nosy_DataAvailabilityChallengeTransactorSession_Withdraw__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DataAvailabilityChallenge *DataAvailabilityChallengeTransactorSession
		fill_err = tp.Fill(&_DataAvailabilityChallenge)
		if fill_err != nil {
			return
		}
		if _DataAvailabilityChallenge == nil {
			return
		}

		_DataAvailabilityChallenge.Withdraw()
	})
}

// skipping Fuzz_Nosy_DeployDataAvailabilityChallenge__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/abi/bind.ContractBackend
