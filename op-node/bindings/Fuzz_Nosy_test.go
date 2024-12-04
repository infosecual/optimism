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

func Fuzz_Nosy_DisputeGameFactoryCaller_FindLatestGames__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DisputeGameFactory *DisputeGameFactoryCaller
		fill_err = tp.Fill(&_DisputeGameFactory)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var _gameType uint32
		fill_err = tp.Fill(&_gameType)
		if fill_err != nil {
			return
		}
		var _start *big.Int
		fill_err = tp.Fill(&_start)
		if fill_err != nil {
			return
		}
		var _n *big.Int
		fill_err = tp.Fill(&_n)
		if fill_err != nil {
			return
		}
		if _DisputeGameFactory == nil || opts == nil || _start == nil || _n == nil {
			return
		}

		_DisputeGameFactory.FindLatestGames(opts, _gameType, _start, _n)
	})
}

func Fuzz_Nosy_DisputeGameFactoryCaller_GameAtIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DisputeGameFactory *DisputeGameFactoryCaller
		fill_err = tp.Fill(&_DisputeGameFactory)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var _index *big.Int
		fill_err = tp.Fill(&_index)
		if fill_err != nil {
			return
		}
		if _DisputeGameFactory == nil || opts == nil || _index == nil {
			return
		}

		_DisputeGameFactory.GameAtIndex(opts, _index)
	})
}

func Fuzz_Nosy_DisputeGameFactoryCaller_GameCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DisputeGameFactory *DisputeGameFactoryCaller
		fill_err = tp.Fill(&_DisputeGameFactory)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _DisputeGameFactory == nil || opts == nil {
			return
		}

		_DisputeGameFactory.GameCount(opts)
	})
}

func Fuzz_Nosy_DisputeGameFactoryCaller_GameImpls__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DisputeGameFactory *DisputeGameFactoryCaller
		fill_err = tp.Fill(&_DisputeGameFactory)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var arg0 uint32
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _DisputeGameFactory == nil || opts == nil {
			return
		}

		_DisputeGameFactory.GameImpls(opts, arg0)
	})
}

func Fuzz_Nosy_DisputeGameFactoryCaller_Games__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DisputeGameFactory *DisputeGameFactoryCaller
		fill_err = tp.Fill(&_DisputeGameFactory)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var _gameType uint32
		fill_err = tp.Fill(&_gameType)
		if fill_err != nil {
			return
		}
		var _rootClaim [32]byte
		fill_err = tp.Fill(&_rootClaim)
		if fill_err != nil {
			return
		}
		var _extraData []byte
		fill_err = tp.Fill(&_extraData)
		if fill_err != nil {
			return
		}
		if _DisputeGameFactory == nil || opts == nil {
			return
		}

		_DisputeGameFactory.Games(opts, _gameType, _rootClaim, _extraData)
	})
}

func Fuzz_Nosy_DisputeGameFactoryCaller_GetGameUUID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DisputeGameFactory *DisputeGameFactoryCaller
		fill_err = tp.Fill(&_DisputeGameFactory)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var _gameType uint32
		fill_err = tp.Fill(&_gameType)
		if fill_err != nil {
			return
		}
		var _rootClaim [32]byte
		fill_err = tp.Fill(&_rootClaim)
		if fill_err != nil {
			return
		}
		var _extraData []byte
		fill_err = tp.Fill(&_extraData)
		if fill_err != nil {
			return
		}
		if _DisputeGameFactory == nil || opts == nil {
			return
		}

		_DisputeGameFactory.GetGameUUID(opts, _gameType, _rootClaim, _extraData)
	})
}

func Fuzz_Nosy_DisputeGameFactoryCaller_InitBonds__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DisputeGameFactory *DisputeGameFactoryCaller
		fill_err = tp.Fill(&_DisputeGameFactory)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var arg0 uint32
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _DisputeGameFactory == nil || opts == nil {
			return
		}

		_DisputeGameFactory.InitBonds(opts, arg0)
	})
}

func Fuzz_Nosy_DisputeGameFactoryCaller_Owner__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DisputeGameFactory *DisputeGameFactoryCaller
		fill_err = tp.Fill(&_DisputeGameFactory)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _DisputeGameFactory == nil || opts == nil {
			return
		}

		_DisputeGameFactory.Owner(opts)
	})
}

func Fuzz_Nosy_DisputeGameFactoryCaller_Version__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DisputeGameFactory *DisputeGameFactoryCaller
		fill_err = tp.Fill(&_DisputeGameFactory)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _DisputeGameFactory == nil || opts == nil {
			return
		}

		_DisputeGameFactory.Version(opts)
	})
}

// skipping Fuzz_Nosy_DisputeGameFactoryCallerRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

func Fuzz_Nosy_DisputeGameFactoryCallerSession_FindLatestGames__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DisputeGameFactory *DisputeGameFactoryCallerSession
		fill_err = tp.Fill(&_DisputeGameFactory)
		if fill_err != nil {
			return
		}
		var _gameType uint32
		fill_err = tp.Fill(&_gameType)
		if fill_err != nil {
			return
		}
		var _start *big.Int
		fill_err = tp.Fill(&_start)
		if fill_err != nil {
			return
		}
		var _n *big.Int
		fill_err = tp.Fill(&_n)
		if fill_err != nil {
			return
		}
		if _DisputeGameFactory == nil || _start == nil || _n == nil {
			return
		}

		_DisputeGameFactory.FindLatestGames(_gameType, _start, _n)
	})
}

func Fuzz_Nosy_DisputeGameFactoryCallerSession_GameAtIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DisputeGameFactory *DisputeGameFactoryCallerSession
		fill_err = tp.Fill(&_DisputeGameFactory)
		if fill_err != nil {
			return
		}
		var _index *big.Int
		fill_err = tp.Fill(&_index)
		if fill_err != nil {
			return
		}
		if _DisputeGameFactory == nil || _index == nil {
			return
		}

		_DisputeGameFactory.GameAtIndex(_index)
	})
}

func Fuzz_Nosy_DisputeGameFactoryCallerSession_GameCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DisputeGameFactory *DisputeGameFactoryCallerSession
		fill_err = tp.Fill(&_DisputeGameFactory)
		if fill_err != nil {
			return
		}
		if _DisputeGameFactory == nil {
			return
		}

		_DisputeGameFactory.GameCount()
	})
}

func Fuzz_Nosy_DisputeGameFactoryCallerSession_GameImpls__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DisputeGameFactory *DisputeGameFactoryCallerSession
		fill_err = tp.Fill(&_DisputeGameFactory)
		if fill_err != nil {
			return
		}
		var arg0 uint32
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _DisputeGameFactory == nil {
			return
		}

		_DisputeGameFactory.GameImpls(arg0)
	})
}

func Fuzz_Nosy_DisputeGameFactoryCallerSession_Games__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DisputeGameFactory *DisputeGameFactoryCallerSession
		fill_err = tp.Fill(&_DisputeGameFactory)
		if fill_err != nil {
			return
		}
		var _gameType uint32
		fill_err = tp.Fill(&_gameType)
		if fill_err != nil {
			return
		}
		var _rootClaim [32]byte
		fill_err = tp.Fill(&_rootClaim)
		if fill_err != nil {
			return
		}
		var _extraData []byte
		fill_err = tp.Fill(&_extraData)
		if fill_err != nil {
			return
		}
		if _DisputeGameFactory == nil {
			return
		}

		_DisputeGameFactory.Games(_gameType, _rootClaim, _extraData)
	})
}

func Fuzz_Nosy_DisputeGameFactoryCallerSession_GetGameUUID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DisputeGameFactory *DisputeGameFactoryCallerSession
		fill_err = tp.Fill(&_DisputeGameFactory)
		if fill_err != nil {
			return
		}
		var _gameType uint32
		fill_err = tp.Fill(&_gameType)
		if fill_err != nil {
			return
		}
		var _rootClaim [32]byte
		fill_err = tp.Fill(&_rootClaim)
		if fill_err != nil {
			return
		}
		var _extraData []byte
		fill_err = tp.Fill(&_extraData)
		if fill_err != nil {
			return
		}
		if _DisputeGameFactory == nil {
			return
		}

		_DisputeGameFactory.GetGameUUID(_gameType, _rootClaim, _extraData)
	})
}

func Fuzz_Nosy_DisputeGameFactoryCallerSession_InitBonds__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DisputeGameFactory *DisputeGameFactoryCallerSession
		fill_err = tp.Fill(&_DisputeGameFactory)
		if fill_err != nil {
			return
		}
		var arg0 uint32
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _DisputeGameFactory == nil {
			return
		}

		_DisputeGameFactory.InitBonds(arg0)
	})
}

func Fuzz_Nosy_DisputeGameFactoryCallerSession_Owner__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DisputeGameFactory *DisputeGameFactoryCallerSession
		fill_err = tp.Fill(&_DisputeGameFactory)
		if fill_err != nil {
			return
		}
		if _DisputeGameFactory == nil {
			return
		}

		_DisputeGameFactory.Owner()
	})
}

func Fuzz_Nosy_DisputeGameFactoryCallerSession_Version__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DisputeGameFactory *DisputeGameFactoryCallerSession
		fill_err = tp.Fill(&_DisputeGameFactory)
		if fill_err != nil {
			return
		}
		if _DisputeGameFactory == nil {
			return
		}

		_DisputeGameFactory.Version()
	})
}

func Fuzz_Nosy_DisputeGameFactoryDisputeGameCreatedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DisputeGameFactoryDisputeGameCreatedIterator
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

func Fuzz_Nosy_DisputeGameFactoryDisputeGameCreatedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DisputeGameFactoryDisputeGameCreatedIterator
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

func Fuzz_Nosy_DisputeGameFactoryDisputeGameCreatedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DisputeGameFactoryDisputeGameCreatedIterator
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

func Fuzz_Nosy_DisputeGameFactoryFilterer_FilterDisputeGameCreated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DisputeGameFactory *DisputeGameFactoryFilterer
		fill_err = tp.Fill(&_DisputeGameFactory)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var disputeProxy []common.Address
		fill_err = tp.Fill(&disputeProxy)
		if fill_err != nil {
			return
		}
		var gameType []uint32
		fill_err = tp.Fill(&gameType)
		if fill_err != nil {
			return
		}
		var rootClaim [][32]byte
		fill_err = tp.Fill(&rootClaim)
		if fill_err != nil {
			return
		}
		if _DisputeGameFactory == nil || opts == nil {
			return
		}

		_DisputeGameFactory.FilterDisputeGameCreated(opts, disputeProxy, gameType, rootClaim)
	})
}

func Fuzz_Nosy_DisputeGameFactoryFilterer_FilterImplementationSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DisputeGameFactory *DisputeGameFactoryFilterer
		fill_err = tp.Fill(&_DisputeGameFactory)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var impl []common.Address
		fill_err = tp.Fill(&impl)
		if fill_err != nil {
			return
		}
		var gameType []uint32
		fill_err = tp.Fill(&gameType)
		if fill_err != nil {
			return
		}
		if _DisputeGameFactory == nil || opts == nil {
			return
		}

		_DisputeGameFactory.FilterImplementationSet(opts, impl, gameType)
	})
}

func Fuzz_Nosy_DisputeGameFactoryFilterer_FilterInitBondUpdated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DisputeGameFactory *DisputeGameFactoryFilterer
		fill_err = tp.Fill(&_DisputeGameFactory)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var gameType []uint32
		fill_err = tp.Fill(&gameType)
		if fill_err != nil {
			return
		}
		var newBond []*big.Int
		fill_err = tp.Fill(&newBond)
		if fill_err != nil {
			return
		}
		if _DisputeGameFactory == nil || opts == nil {
			return
		}

		_DisputeGameFactory.FilterInitBondUpdated(opts, gameType, newBond)
	})
}

func Fuzz_Nosy_DisputeGameFactoryFilterer_FilterInitialized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DisputeGameFactory *DisputeGameFactoryFilterer
		fill_err = tp.Fill(&_DisputeGameFactory)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _DisputeGameFactory == nil || opts == nil {
			return
		}

		_DisputeGameFactory.FilterInitialized(opts)
	})
}

func Fuzz_Nosy_DisputeGameFactoryFilterer_FilterOwnershipTransferred__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DisputeGameFactory *DisputeGameFactoryFilterer
		fill_err = tp.Fill(&_DisputeGameFactory)
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
		if _DisputeGameFactory == nil || opts == nil {
			return
		}

		_DisputeGameFactory.FilterOwnershipTransferred(opts, previousOwner, newOwner)
	})
}

func Fuzz_Nosy_DisputeGameFactoryFilterer_ParseDisputeGameCreated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DisputeGameFactory *DisputeGameFactoryFilterer
		fill_err = tp.Fill(&_DisputeGameFactory)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _DisputeGameFactory == nil {
			return
		}

		_DisputeGameFactory.ParseDisputeGameCreated(log)
	})
}

func Fuzz_Nosy_DisputeGameFactoryFilterer_ParseImplementationSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DisputeGameFactory *DisputeGameFactoryFilterer
		fill_err = tp.Fill(&_DisputeGameFactory)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _DisputeGameFactory == nil {
			return
		}

		_DisputeGameFactory.ParseImplementationSet(log)
	})
}

func Fuzz_Nosy_DisputeGameFactoryFilterer_ParseInitBondUpdated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DisputeGameFactory *DisputeGameFactoryFilterer
		fill_err = tp.Fill(&_DisputeGameFactory)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _DisputeGameFactory == nil {
			return
		}

		_DisputeGameFactory.ParseInitBondUpdated(log)
	})
}

func Fuzz_Nosy_DisputeGameFactoryFilterer_ParseInitialized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DisputeGameFactory *DisputeGameFactoryFilterer
		fill_err = tp.Fill(&_DisputeGameFactory)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _DisputeGameFactory == nil {
			return
		}

		_DisputeGameFactory.ParseInitialized(log)
	})
}

func Fuzz_Nosy_DisputeGameFactoryFilterer_ParseOwnershipTransferred__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DisputeGameFactory *DisputeGameFactoryFilterer
		fill_err = tp.Fill(&_DisputeGameFactory)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _DisputeGameFactory == nil {
			return
		}

		_DisputeGameFactory.ParseOwnershipTransferred(log)
	})
}

// skipping Fuzz_Nosy_DisputeGameFactoryFilterer_WatchDisputeGameCreated__ because parameters include func, chan, or unsupported interface: chan<- *github.com/ethereum-optimism/optimism/op-node/bindings.DisputeGameFactoryDisputeGameCreated

// skipping Fuzz_Nosy_DisputeGameFactoryFilterer_WatchImplementationSet__ because parameters include func, chan, or unsupported interface: chan<- *github.com/ethereum-optimism/optimism/op-node/bindings.DisputeGameFactoryImplementationSet

// skipping Fuzz_Nosy_DisputeGameFactoryFilterer_WatchInitBondUpdated__ because parameters include func, chan, or unsupported interface: chan<- *github.com/ethereum-optimism/optimism/op-node/bindings.DisputeGameFactoryInitBondUpdated

// skipping Fuzz_Nosy_DisputeGameFactoryFilterer_WatchInitialized__ because parameters include func, chan, or unsupported interface: chan<- *github.com/ethereum-optimism/optimism/op-node/bindings.DisputeGameFactoryInitialized

// skipping Fuzz_Nosy_DisputeGameFactoryFilterer_WatchOwnershipTransferred__ because parameters include func, chan, or unsupported interface: chan<- *github.com/ethereum-optimism/optimism/op-node/bindings.DisputeGameFactoryOwnershipTransferred

func Fuzz_Nosy_DisputeGameFactoryImplementationSetIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DisputeGameFactoryImplementationSetIterator
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

func Fuzz_Nosy_DisputeGameFactoryImplementationSetIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DisputeGameFactoryImplementationSetIterator
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

func Fuzz_Nosy_DisputeGameFactoryImplementationSetIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DisputeGameFactoryImplementationSetIterator
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

func Fuzz_Nosy_DisputeGameFactoryInitBondUpdatedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DisputeGameFactoryInitBondUpdatedIterator
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

func Fuzz_Nosy_DisputeGameFactoryInitBondUpdatedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DisputeGameFactoryInitBondUpdatedIterator
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

func Fuzz_Nosy_DisputeGameFactoryInitBondUpdatedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DisputeGameFactoryInitBondUpdatedIterator
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

func Fuzz_Nosy_DisputeGameFactoryInitializedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DisputeGameFactoryInitializedIterator
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

func Fuzz_Nosy_DisputeGameFactoryInitializedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DisputeGameFactoryInitializedIterator
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

func Fuzz_Nosy_DisputeGameFactoryInitializedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DisputeGameFactoryInitializedIterator
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

func Fuzz_Nosy_DisputeGameFactoryOwnershipTransferredIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DisputeGameFactoryOwnershipTransferredIterator
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

func Fuzz_Nosy_DisputeGameFactoryOwnershipTransferredIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DisputeGameFactoryOwnershipTransferredIterator
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

func Fuzz_Nosy_DisputeGameFactoryOwnershipTransferredIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *DisputeGameFactoryOwnershipTransferredIterator
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

// skipping Fuzz_Nosy_DisputeGameFactoryRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

// skipping Fuzz_Nosy_DisputeGameFactoryRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_DisputeGameFactoryRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DisputeGameFactory *DisputeGameFactoryRaw
		fill_err = tp.Fill(&_DisputeGameFactory)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _DisputeGameFactory == nil || opts == nil {
			return
		}

		_DisputeGameFactory.Transfer(opts)
	})
}

func Fuzz_Nosy_DisputeGameFactorySession_Create__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DisputeGameFactory *DisputeGameFactorySession
		fill_err = tp.Fill(&_DisputeGameFactory)
		if fill_err != nil {
			return
		}
		var _gameType uint32
		fill_err = tp.Fill(&_gameType)
		if fill_err != nil {
			return
		}
		var _rootClaim [32]byte
		fill_err = tp.Fill(&_rootClaim)
		if fill_err != nil {
			return
		}
		var _extraData []byte
		fill_err = tp.Fill(&_extraData)
		if fill_err != nil {
			return
		}
		if _DisputeGameFactory == nil {
			return
		}

		_DisputeGameFactory.Create(_gameType, _rootClaim, _extraData)
	})
}

func Fuzz_Nosy_DisputeGameFactorySession_FindLatestGames__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DisputeGameFactory *DisputeGameFactorySession
		fill_err = tp.Fill(&_DisputeGameFactory)
		if fill_err != nil {
			return
		}
		var _gameType uint32
		fill_err = tp.Fill(&_gameType)
		if fill_err != nil {
			return
		}
		var _start *big.Int
		fill_err = tp.Fill(&_start)
		if fill_err != nil {
			return
		}
		var _n *big.Int
		fill_err = tp.Fill(&_n)
		if fill_err != nil {
			return
		}
		if _DisputeGameFactory == nil || _start == nil || _n == nil {
			return
		}

		_DisputeGameFactory.FindLatestGames(_gameType, _start, _n)
	})
}

func Fuzz_Nosy_DisputeGameFactorySession_GameAtIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DisputeGameFactory *DisputeGameFactorySession
		fill_err = tp.Fill(&_DisputeGameFactory)
		if fill_err != nil {
			return
		}
		var _index *big.Int
		fill_err = tp.Fill(&_index)
		if fill_err != nil {
			return
		}
		if _DisputeGameFactory == nil || _index == nil {
			return
		}

		_DisputeGameFactory.GameAtIndex(_index)
	})
}

func Fuzz_Nosy_DisputeGameFactorySession_GameCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DisputeGameFactory *DisputeGameFactorySession
		fill_err = tp.Fill(&_DisputeGameFactory)
		if fill_err != nil {
			return
		}
		if _DisputeGameFactory == nil {
			return
		}

		_DisputeGameFactory.GameCount()
	})
}

func Fuzz_Nosy_DisputeGameFactorySession_GameImpls__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DisputeGameFactory *DisputeGameFactorySession
		fill_err = tp.Fill(&_DisputeGameFactory)
		if fill_err != nil {
			return
		}
		var arg0 uint32
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _DisputeGameFactory == nil {
			return
		}

		_DisputeGameFactory.GameImpls(arg0)
	})
}

func Fuzz_Nosy_DisputeGameFactorySession_Games__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DisputeGameFactory *DisputeGameFactorySession
		fill_err = tp.Fill(&_DisputeGameFactory)
		if fill_err != nil {
			return
		}
		var _gameType uint32
		fill_err = tp.Fill(&_gameType)
		if fill_err != nil {
			return
		}
		var _rootClaim [32]byte
		fill_err = tp.Fill(&_rootClaim)
		if fill_err != nil {
			return
		}
		var _extraData []byte
		fill_err = tp.Fill(&_extraData)
		if fill_err != nil {
			return
		}
		if _DisputeGameFactory == nil {
			return
		}

		_DisputeGameFactory.Games(_gameType, _rootClaim, _extraData)
	})
}

func Fuzz_Nosy_DisputeGameFactorySession_GetGameUUID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DisputeGameFactory *DisputeGameFactorySession
		fill_err = tp.Fill(&_DisputeGameFactory)
		if fill_err != nil {
			return
		}
		var _gameType uint32
		fill_err = tp.Fill(&_gameType)
		if fill_err != nil {
			return
		}
		var _rootClaim [32]byte
		fill_err = tp.Fill(&_rootClaim)
		if fill_err != nil {
			return
		}
		var _extraData []byte
		fill_err = tp.Fill(&_extraData)
		if fill_err != nil {
			return
		}
		if _DisputeGameFactory == nil {
			return
		}

		_DisputeGameFactory.GetGameUUID(_gameType, _rootClaim, _extraData)
	})
}

func Fuzz_Nosy_DisputeGameFactorySession_InitBonds__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DisputeGameFactory *DisputeGameFactorySession
		fill_err = tp.Fill(&_DisputeGameFactory)
		if fill_err != nil {
			return
		}
		var arg0 uint32
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _DisputeGameFactory == nil {
			return
		}

		_DisputeGameFactory.InitBonds(arg0)
	})
}

func Fuzz_Nosy_DisputeGameFactorySession_Initialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DisputeGameFactory *DisputeGameFactorySession
		fill_err = tp.Fill(&_DisputeGameFactory)
		if fill_err != nil {
			return
		}
		var _owner common.Address
		fill_err = tp.Fill(&_owner)
		if fill_err != nil {
			return
		}
		if _DisputeGameFactory == nil {
			return
		}

		_DisputeGameFactory.Initialize(_owner)
	})
}

func Fuzz_Nosy_DisputeGameFactorySession_Owner__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DisputeGameFactory *DisputeGameFactorySession
		fill_err = tp.Fill(&_DisputeGameFactory)
		if fill_err != nil {
			return
		}
		if _DisputeGameFactory == nil {
			return
		}

		_DisputeGameFactory.Owner()
	})
}

func Fuzz_Nosy_DisputeGameFactorySession_RenounceOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DisputeGameFactory *DisputeGameFactorySession
		fill_err = tp.Fill(&_DisputeGameFactory)
		if fill_err != nil {
			return
		}
		if _DisputeGameFactory == nil {
			return
		}

		_DisputeGameFactory.RenounceOwnership()
	})
}

func Fuzz_Nosy_DisputeGameFactorySession_SetImplementation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DisputeGameFactory *DisputeGameFactorySession
		fill_err = tp.Fill(&_DisputeGameFactory)
		if fill_err != nil {
			return
		}
		var _gameType uint32
		fill_err = tp.Fill(&_gameType)
		if fill_err != nil {
			return
		}
		var _impl common.Address
		fill_err = tp.Fill(&_impl)
		if fill_err != nil {
			return
		}
		if _DisputeGameFactory == nil {
			return
		}

		_DisputeGameFactory.SetImplementation(_gameType, _impl)
	})
}

func Fuzz_Nosy_DisputeGameFactorySession_SetInitBond__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DisputeGameFactory *DisputeGameFactorySession
		fill_err = tp.Fill(&_DisputeGameFactory)
		if fill_err != nil {
			return
		}
		var _gameType uint32
		fill_err = tp.Fill(&_gameType)
		if fill_err != nil {
			return
		}
		var _initBond *big.Int
		fill_err = tp.Fill(&_initBond)
		if fill_err != nil {
			return
		}
		if _DisputeGameFactory == nil || _initBond == nil {
			return
		}

		_DisputeGameFactory.SetInitBond(_gameType, _initBond)
	})
}

func Fuzz_Nosy_DisputeGameFactorySession_TransferOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DisputeGameFactory *DisputeGameFactorySession
		fill_err = tp.Fill(&_DisputeGameFactory)
		if fill_err != nil {
			return
		}
		var newOwner common.Address
		fill_err = tp.Fill(&newOwner)
		if fill_err != nil {
			return
		}
		if _DisputeGameFactory == nil {
			return
		}

		_DisputeGameFactory.TransferOwnership(newOwner)
	})
}

func Fuzz_Nosy_DisputeGameFactorySession_Version__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DisputeGameFactory *DisputeGameFactorySession
		fill_err = tp.Fill(&_DisputeGameFactory)
		if fill_err != nil {
			return
		}
		if _DisputeGameFactory == nil {
			return
		}

		_DisputeGameFactory.Version()
	})
}

func Fuzz_Nosy_DisputeGameFactoryTransactor_Create__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DisputeGameFactory *DisputeGameFactoryTransactor
		fill_err = tp.Fill(&_DisputeGameFactory)
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
		var _rootClaim [32]byte
		fill_err = tp.Fill(&_rootClaim)
		if fill_err != nil {
			return
		}
		var _extraData []byte
		fill_err = tp.Fill(&_extraData)
		if fill_err != nil {
			return
		}
		if _DisputeGameFactory == nil || opts == nil {
			return
		}

		_DisputeGameFactory.Create(opts, _gameType, _rootClaim, _extraData)
	})
}

func Fuzz_Nosy_DisputeGameFactoryTransactor_Initialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DisputeGameFactory *DisputeGameFactoryTransactor
		fill_err = tp.Fill(&_DisputeGameFactory)
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
		if _DisputeGameFactory == nil || opts == nil {
			return
		}

		_DisputeGameFactory.Initialize(opts, _owner)
	})
}

func Fuzz_Nosy_DisputeGameFactoryTransactor_RenounceOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DisputeGameFactory *DisputeGameFactoryTransactor
		fill_err = tp.Fill(&_DisputeGameFactory)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _DisputeGameFactory == nil || opts == nil {
			return
		}

		_DisputeGameFactory.RenounceOwnership(opts)
	})
}

func Fuzz_Nosy_DisputeGameFactoryTransactor_SetImplementation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DisputeGameFactory *DisputeGameFactoryTransactor
		fill_err = tp.Fill(&_DisputeGameFactory)
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
		var _impl common.Address
		fill_err = tp.Fill(&_impl)
		if fill_err != nil {
			return
		}
		if _DisputeGameFactory == nil || opts == nil {
			return
		}

		_DisputeGameFactory.SetImplementation(opts, _gameType, _impl)
	})
}

func Fuzz_Nosy_DisputeGameFactoryTransactor_SetInitBond__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DisputeGameFactory *DisputeGameFactoryTransactor
		fill_err = tp.Fill(&_DisputeGameFactory)
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
		var _initBond *big.Int
		fill_err = tp.Fill(&_initBond)
		if fill_err != nil {
			return
		}
		if _DisputeGameFactory == nil || opts == nil || _initBond == nil {
			return
		}

		_DisputeGameFactory.SetInitBond(opts, _gameType, _initBond)
	})
}

func Fuzz_Nosy_DisputeGameFactoryTransactor_TransferOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DisputeGameFactory *DisputeGameFactoryTransactor
		fill_err = tp.Fill(&_DisputeGameFactory)
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
		if _DisputeGameFactory == nil || opts == nil {
			return
		}

		_DisputeGameFactory.TransferOwnership(opts, newOwner)
	})
}

// skipping Fuzz_Nosy_DisputeGameFactoryTransactorRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_DisputeGameFactoryTransactorRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DisputeGameFactory *DisputeGameFactoryTransactorRaw
		fill_err = tp.Fill(&_DisputeGameFactory)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _DisputeGameFactory == nil || opts == nil {
			return
		}

		_DisputeGameFactory.Transfer(opts)
	})
}

func Fuzz_Nosy_DisputeGameFactoryTransactorSession_Create__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DisputeGameFactory *DisputeGameFactoryTransactorSession
		fill_err = tp.Fill(&_DisputeGameFactory)
		if fill_err != nil {
			return
		}
		var _gameType uint32
		fill_err = tp.Fill(&_gameType)
		if fill_err != nil {
			return
		}
		var _rootClaim [32]byte
		fill_err = tp.Fill(&_rootClaim)
		if fill_err != nil {
			return
		}
		var _extraData []byte
		fill_err = tp.Fill(&_extraData)
		if fill_err != nil {
			return
		}
		if _DisputeGameFactory == nil {
			return
		}

		_DisputeGameFactory.Create(_gameType, _rootClaim, _extraData)
	})
}

func Fuzz_Nosy_DisputeGameFactoryTransactorSession_Initialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DisputeGameFactory *DisputeGameFactoryTransactorSession
		fill_err = tp.Fill(&_DisputeGameFactory)
		if fill_err != nil {
			return
		}
		var _owner common.Address
		fill_err = tp.Fill(&_owner)
		if fill_err != nil {
			return
		}
		if _DisputeGameFactory == nil {
			return
		}

		_DisputeGameFactory.Initialize(_owner)
	})
}

func Fuzz_Nosy_DisputeGameFactoryTransactorSession_RenounceOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DisputeGameFactory *DisputeGameFactoryTransactorSession
		fill_err = tp.Fill(&_DisputeGameFactory)
		if fill_err != nil {
			return
		}
		if _DisputeGameFactory == nil {
			return
		}

		_DisputeGameFactory.RenounceOwnership()
	})
}

func Fuzz_Nosy_DisputeGameFactoryTransactorSession_SetImplementation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DisputeGameFactory *DisputeGameFactoryTransactorSession
		fill_err = tp.Fill(&_DisputeGameFactory)
		if fill_err != nil {
			return
		}
		var _gameType uint32
		fill_err = tp.Fill(&_gameType)
		if fill_err != nil {
			return
		}
		var _impl common.Address
		fill_err = tp.Fill(&_impl)
		if fill_err != nil {
			return
		}
		if _DisputeGameFactory == nil {
			return
		}

		_DisputeGameFactory.SetImplementation(_gameType, _impl)
	})
}

func Fuzz_Nosy_DisputeGameFactoryTransactorSession_SetInitBond__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DisputeGameFactory *DisputeGameFactoryTransactorSession
		fill_err = tp.Fill(&_DisputeGameFactory)
		if fill_err != nil {
			return
		}
		var _gameType uint32
		fill_err = tp.Fill(&_gameType)
		if fill_err != nil {
			return
		}
		var _initBond *big.Int
		fill_err = tp.Fill(&_initBond)
		if fill_err != nil {
			return
		}
		if _DisputeGameFactory == nil || _initBond == nil {
			return
		}

		_DisputeGameFactory.SetInitBond(_gameType, _initBond)
	})
}

func Fuzz_Nosy_DisputeGameFactoryTransactorSession_TransferOwnership__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _DisputeGameFactory *DisputeGameFactoryTransactorSession
		fill_err = tp.Fill(&_DisputeGameFactory)
		if fill_err != nil {
			return
		}
		var newOwner common.Address
		fill_err = tp.Fill(&newOwner)
		if fill_err != nil {
			return
		}
		if _DisputeGameFactory == nil {
			return
		}

		_DisputeGameFactory.TransferOwnership(newOwner)
	})
}

func Fuzz_Nosy_L1BlockCaller_BaseFeeScalar__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1Block *L1BlockCaller
		fill_err = tp.Fill(&_L1Block)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L1Block == nil || opts == nil {
			return
		}

		_L1Block.BaseFeeScalar(opts)
	})
}

func Fuzz_Nosy_L1BlockCaller_Basefee__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1Block *L1BlockCaller
		fill_err = tp.Fill(&_L1Block)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L1Block == nil || opts == nil {
			return
		}

		_L1Block.Basefee(opts)
	})
}

func Fuzz_Nosy_L1BlockCaller_BatcherHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1Block *L1BlockCaller
		fill_err = tp.Fill(&_L1Block)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L1Block == nil || opts == nil {
			return
		}

		_L1Block.BatcherHash(opts)
	})
}

func Fuzz_Nosy_L1BlockCaller_BlobBaseFee__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1Block *L1BlockCaller
		fill_err = tp.Fill(&_L1Block)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L1Block == nil || opts == nil {
			return
		}

		_L1Block.BlobBaseFee(opts)
	})
}

func Fuzz_Nosy_L1BlockCaller_BlobBaseFeeScalar__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1Block *L1BlockCaller
		fill_err = tp.Fill(&_L1Block)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L1Block == nil || opts == nil {
			return
		}

		_L1Block.BlobBaseFeeScalar(opts)
	})
}

func Fuzz_Nosy_L1BlockCaller_DEPOSITORACCOUNT__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1Block *L1BlockCaller
		fill_err = tp.Fill(&_L1Block)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L1Block == nil || opts == nil {
			return
		}

		_L1Block.DEPOSITORACCOUNT(opts)
	})
}

func Fuzz_Nosy_L1BlockCaller_Hash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1Block *L1BlockCaller
		fill_err = tp.Fill(&_L1Block)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L1Block == nil || opts == nil {
			return
		}

		_L1Block.Hash(opts)
	})
}

func Fuzz_Nosy_L1BlockCaller_L1FeeOverhead__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1Block *L1BlockCaller
		fill_err = tp.Fill(&_L1Block)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L1Block == nil || opts == nil {
			return
		}

		_L1Block.L1FeeOverhead(opts)
	})
}

func Fuzz_Nosy_L1BlockCaller_L1FeeScalar__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1Block *L1BlockCaller
		fill_err = tp.Fill(&_L1Block)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L1Block == nil || opts == nil {
			return
		}

		_L1Block.L1FeeScalar(opts)
	})
}

func Fuzz_Nosy_L1BlockCaller_Number__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1Block *L1BlockCaller
		fill_err = tp.Fill(&_L1Block)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L1Block == nil || opts == nil {
			return
		}

		_L1Block.Number(opts)
	})
}

func Fuzz_Nosy_L1BlockCaller_SequenceNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1Block *L1BlockCaller
		fill_err = tp.Fill(&_L1Block)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L1Block == nil || opts == nil {
			return
		}

		_L1Block.SequenceNumber(opts)
	})
}

func Fuzz_Nosy_L1BlockCaller_Timestamp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1Block *L1BlockCaller
		fill_err = tp.Fill(&_L1Block)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L1Block == nil || opts == nil {
			return
		}

		_L1Block.Timestamp(opts)
	})
}

func Fuzz_Nosy_L1BlockCaller_Version__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1Block *L1BlockCaller
		fill_err = tp.Fill(&_L1Block)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L1Block == nil || opts == nil {
			return
		}

		_L1Block.Version(opts)
	})
}

// skipping Fuzz_Nosy_L1BlockCallerRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

func Fuzz_Nosy_L1BlockCallerSession_BaseFeeScalar__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1Block *L1BlockCallerSession
		fill_err = tp.Fill(&_L1Block)
		if fill_err != nil {
			return
		}
		if _L1Block == nil {
			return
		}

		_L1Block.BaseFeeScalar()
	})
}

func Fuzz_Nosy_L1BlockCallerSession_Basefee__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1Block *L1BlockCallerSession
		fill_err = tp.Fill(&_L1Block)
		if fill_err != nil {
			return
		}
		if _L1Block == nil {
			return
		}

		_L1Block.Basefee()
	})
}

func Fuzz_Nosy_L1BlockCallerSession_BatcherHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1Block *L1BlockCallerSession
		fill_err = tp.Fill(&_L1Block)
		if fill_err != nil {
			return
		}
		if _L1Block == nil {
			return
		}

		_L1Block.BatcherHash()
	})
}

func Fuzz_Nosy_L1BlockCallerSession_BlobBaseFee__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1Block *L1BlockCallerSession
		fill_err = tp.Fill(&_L1Block)
		if fill_err != nil {
			return
		}
		if _L1Block == nil {
			return
		}

		_L1Block.BlobBaseFee()
	})
}

func Fuzz_Nosy_L1BlockCallerSession_BlobBaseFeeScalar__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1Block *L1BlockCallerSession
		fill_err = tp.Fill(&_L1Block)
		if fill_err != nil {
			return
		}
		if _L1Block == nil {
			return
		}

		_L1Block.BlobBaseFeeScalar()
	})
}

func Fuzz_Nosy_L1BlockCallerSession_DEPOSITORACCOUNT__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1Block *L1BlockCallerSession
		fill_err = tp.Fill(&_L1Block)
		if fill_err != nil {
			return
		}
		if _L1Block == nil {
			return
		}

		_L1Block.DEPOSITORACCOUNT()
	})
}

func Fuzz_Nosy_L1BlockCallerSession_Hash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1Block *L1BlockCallerSession
		fill_err = tp.Fill(&_L1Block)
		if fill_err != nil {
			return
		}
		if _L1Block == nil {
			return
		}

		_L1Block.Hash()
	})
}

func Fuzz_Nosy_L1BlockCallerSession_L1FeeOverhead__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1Block *L1BlockCallerSession
		fill_err = tp.Fill(&_L1Block)
		if fill_err != nil {
			return
		}
		if _L1Block == nil {
			return
		}

		_L1Block.L1FeeOverhead()
	})
}

func Fuzz_Nosy_L1BlockCallerSession_L1FeeScalar__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1Block *L1BlockCallerSession
		fill_err = tp.Fill(&_L1Block)
		if fill_err != nil {
			return
		}
		if _L1Block == nil {
			return
		}

		_L1Block.L1FeeScalar()
	})
}

func Fuzz_Nosy_L1BlockCallerSession_Number__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1Block *L1BlockCallerSession
		fill_err = tp.Fill(&_L1Block)
		if fill_err != nil {
			return
		}
		if _L1Block == nil {
			return
		}

		_L1Block.Number()
	})
}

func Fuzz_Nosy_L1BlockCallerSession_SequenceNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1Block *L1BlockCallerSession
		fill_err = tp.Fill(&_L1Block)
		if fill_err != nil {
			return
		}
		if _L1Block == nil {
			return
		}

		_L1Block.SequenceNumber()
	})
}

func Fuzz_Nosy_L1BlockCallerSession_Timestamp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1Block *L1BlockCallerSession
		fill_err = tp.Fill(&_L1Block)
		if fill_err != nil {
			return
		}
		if _L1Block == nil {
			return
		}

		_L1Block.Timestamp()
	})
}

func Fuzz_Nosy_L1BlockCallerSession_Version__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1Block *L1BlockCallerSession
		fill_err = tp.Fill(&_L1Block)
		if fill_err != nil {
			return
		}
		if _L1Block == nil {
			return
		}

		_L1Block.Version()
	})
}

// skipping Fuzz_Nosy_L1BlockRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

// skipping Fuzz_Nosy_L1BlockRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_L1BlockRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1Block *L1BlockRaw
		fill_err = tp.Fill(&_L1Block)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L1Block == nil || opts == nil {
			return
		}

		_L1Block.Transfer(opts)
	})
}

func Fuzz_Nosy_L1BlockSession_BaseFeeScalar__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1Block *L1BlockSession
		fill_err = tp.Fill(&_L1Block)
		if fill_err != nil {
			return
		}
		if _L1Block == nil {
			return
		}

		_L1Block.BaseFeeScalar()
	})
}

func Fuzz_Nosy_L1BlockSession_Basefee__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1Block *L1BlockSession
		fill_err = tp.Fill(&_L1Block)
		if fill_err != nil {
			return
		}
		if _L1Block == nil {
			return
		}

		_L1Block.Basefee()
	})
}

func Fuzz_Nosy_L1BlockSession_BatcherHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1Block *L1BlockSession
		fill_err = tp.Fill(&_L1Block)
		if fill_err != nil {
			return
		}
		if _L1Block == nil {
			return
		}

		_L1Block.BatcherHash()
	})
}

func Fuzz_Nosy_L1BlockSession_BlobBaseFee__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1Block *L1BlockSession
		fill_err = tp.Fill(&_L1Block)
		if fill_err != nil {
			return
		}
		if _L1Block == nil {
			return
		}

		_L1Block.BlobBaseFee()
	})
}

func Fuzz_Nosy_L1BlockSession_BlobBaseFeeScalar__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1Block *L1BlockSession
		fill_err = tp.Fill(&_L1Block)
		if fill_err != nil {
			return
		}
		if _L1Block == nil {
			return
		}

		_L1Block.BlobBaseFeeScalar()
	})
}

func Fuzz_Nosy_L1BlockSession_DEPOSITORACCOUNT__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1Block *L1BlockSession
		fill_err = tp.Fill(&_L1Block)
		if fill_err != nil {
			return
		}
		if _L1Block == nil {
			return
		}

		_L1Block.DEPOSITORACCOUNT()
	})
}

func Fuzz_Nosy_L1BlockSession_Hash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1Block *L1BlockSession
		fill_err = tp.Fill(&_L1Block)
		if fill_err != nil {
			return
		}
		if _L1Block == nil {
			return
		}

		_L1Block.Hash()
	})
}

func Fuzz_Nosy_L1BlockSession_L1FeeOverhead__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1Block *L1BlockSession
		fill_err = tp.Fill(&_L1Block)
		if fill_err != nil {
			return
		}
		if _L1Block == nil {
			return
		}

		_L1Block.L1FeeOverhead()
	})
}

func Fuzz_Nosy_L1BlockSession_L1FeeScalar__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1Block *L1BlockSession
		fill_err = tp.Fill(&_L1Block)
		if fill_err != nil {
			return
		}
		if _L1Block == nil {
			return
		}

		_L1Block.L1FeeScalar()
	})
}

func Fuzz_Nosy_L1BlockSession_Number__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1Block *L1BlockSession
		fill_err = tp.Fill(&_L1Block)
		if fill_err != nil {
			return
		}
		if _L1Block == nil {
			return
		}

		_L1Block.Number()
	})
}

func Fuzz_Nosy_L1BlockSession_SequenceNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1Block *L1BlockSession
		fill_err = tp.Fill(&_L1Block)
		if fill_err != nil {
			return
		}
		if _L1Block == nil {
			return
		}

		_L1Block.SequenceNumber()
	})
}

func Fuzz_Nosy_L1BlockSession_SetL1BlockValues__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1Block *L1BlockSession
		fill_err = tp.Fill(&_L1Block)
		if fill_err != nil {
			return
		}
		var _number uint64
		fill_err = tp.Fill(&_number)
		if fill_err != nil {
			return
		}
		var _timestamp uint64
		fill_err = tp.Fill(&_timestamp)
		if fill_err != nil {
			return
		}
		var _basefee *big.Int
		fill_err = tp.Fill(&_basefee)
		if fill_err != nil {
			return
		}
		var _hash [32]byte
		fill_err = tp.Fill(&_hash)
		if fill_err != nil {
			return
		}
		var _sequenceNumber uint64
		fill_err = tp.Fill(&_sequenceNumber)
		if fill_err != nil {
			return
		}
		var _batcherHash [32]byte
		fill_err = tp.Fill(&_batcherHash)
		if fill_err != nil {
			return
		}
		var _l1FeeOverhead *big.Int
		fill_err = tp.Fill(&_l1FeeOverhead)
		if fill_err != nil {
			return
		}
		var _l1FeeScalar *big.Int
		fill_err = tp.Fill(&_l1FeeScalar)
		if fill_err != nil {
			return
		}
		if _L1Block == nil || _basefee == nil || _l1FeeOverhead == nil || _l1FeeScalar == nil {
			return
		}

		_L1Block.SetL1BlockValues(_number, _timestamp, _basefee, _hash, _sequenceNumber, _batcherHash, _l1FeeOverhead, _l1FeeScalar)
	})
}

func Fuzz_Nosy_L1BlockSession_SetL1BlockValuesEcotone__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1Block *L1BlockSession
		fill_err = tp.Fill(&_L1Block)
		if fill_err != nil {
			return
		}
		if _L1Block == nil {
			return
		}

		_L1Block.SetL1BlockValuesEcotone()
	})
}

func Fuzz_Nosy_L1BlockSession_Timestamp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1Block *L1BlockSession
		fill_err = tp.Fill(&_L1Block)
		if fill_err != nil {
			return
		}
		if _L1Block == nil {
			return
		}

		_L1Block.Timestamp()
	})
}

func Fuzz_Nosy_L1BlockSession_Version__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1Block *L1BlockSession
		fill_err = tp.Fill(&_L1Block)
		if fill_err != nil {
			return
		}
		if _L1Block == nil {
			return
		}

		_L1Block.Version()
	})
}

func Fuzz_Nosy_L1BlockTransactor_SetL1BlockValues__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1Block *L1BlockTransactor
		fill_err = tp.Fill(&_L1Block)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var _number uint64
		fill_err = tp.Fill(&_number)
		if fill_err != nil {
			return
		}
		var _timestamp uint64
		fill_err = tp.Fill(&_timestamp)
		if fill_err != nil {
			return
		}
		var _basefee *big.Int
		fill_err = tp.Fill(&_basefee)
		if fill_err != nil {
			return
		}
		var _hash [32]byte
		fill_err = tp.Fill(&_hash)
		if fill_err != nil {
			return
		}
		var _sequenceNumber uint64
		fill_err = tp.Fill(&_sequenceNumber)
		if fill_err != nil {
			return
		}
		var _batcherHash [32]byte
		fill_err = tp.Fill(&_batcherHash)
		if fill_err != nil {
			return
		}
		var _l1FeeOverhead *big.Int
		fill_err = tp.Fill(&_l1FeeOverhead)
		if fill_err != nil {
			return
		}
		var _l1FeeScalar *big.Int
		fill_err = tp.Fill(&_l1FeeScalar)
		if fill_err != nil {
			return
		}
		if _L1Block == nil || opts == nil || _basefee == nil || _l1FeeOverhead == nil || _l1FeeScalar == nil {
			return
		}

		_L1Block.SetL1BlockValues(opts, _number, _timestamp, _basefee, _hash, _sequenceNumber, _batcherHash, _l1FeeOverhead, _l1FeeScalar)
	})
}

func Fuzz_Nosy_L1BlockTransactor_SetL1BlockValuesEcotone__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1Block *L1BlockTransactor
		fill_err = tp.Fill(&_L1Block)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L1Block == nil || opts == nil {
			return
		}

		_L1Block.SetL1BlockValuesEcotone(opts)
	})
}

// skipping Fuzz_Nosy_L1BlockTransactorRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_L1BlockTransactorRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1Block *L1BlockTransactorRaw
		fill_err = tp.Fill(&_L1Block)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L1Block == nil || opts == nil {
			return
		}

		_L1Block.Transfer(opts)
	})
}

func Fuzz_Nosy_L1BlockTransactorSession_SetL1BlockValues__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1Block *L1BlockTransactorSession
		fill_err = tp.Fill(&_L1Block)
		if fill_err != nil {
			return
		}
		var _number uint64
		fill_err = tp.Fill(&_number)
		if fill_err != nil {
			return
		}
		var _timestamp uint64
		fill_err = tp.Fill(&_timestamp)
		if fill_err != nil {
			return
		}
		var _basefee *big.Int
		fill_err = tp.Fill(&_basefee)
		if fill_err != nil {
			return
		}
		var _hash [32]byte
		fill_err = tp.Fill(&_hash)
		if fill_err != nil {
			return
		}
		var _sequenceNumber uint64
		fill_err = tp.Fill(&_sequenceNumber)
		if fill_err != nil {
			return
		}
		var _batcherHash [32]byte
		fill_err = tp.Fill(&_batcherHash)
		if fill_err != nil {
			return
		}
		var _l1FeeOverhead *big.Int
		fill_err = tp.Fill(&_l1FeeOverhead)
		if fill_err != nil {
			return
		}
		var _l1FeeScalar *big.Int
		fill_err = tp.Fill(&_l1FeeScalar)
		if fill_err != nil {
			return
		}
		if _L1Block == nil || _basefee == nil || _l1FeeOverhead == nil || _l1FeeScalar == nil {
			return
		}

		_L1Block.SetL1BlockValues(_number, _timestamp, _basefee, _hash, _sequenceNumber, _batcherHash, _l1FeeOverhead, _l1FeeScalar)
	})
}

func Fuzz_Nosy_L1BlockTransactorSession_SetL1BlockValuesEcotone__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1Block *L1BlockTransactorSession
		fill_err = tp.Fill(&_L1Block)
		if fill_err != nil {
			return
		}
		if _L1Block == nil {
			return
		}

		_L1Block.SetL1BlockValuesEcotone()
	})
}

func Fuzz_Nosy_L2OutputOracleCaller_CHALLENGER__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleCaller
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil || opts == nil {
			return
		}

		_L2OutputOracle.CHALLENGER(opts)
	})
}

func Fuzz_Nosy_L2OutputOracleCaller_Challenger__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleCaller
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil || opts == nil {
			return
		}

		_L2OutputOracle.Challenger(opts)
	})
}

func Fuzz_Nosy_L2OutputOracleCaller_ComputeL2Timestamp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleCaller
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var _l2BlockNumber *big.Int
		fill_err = tp.Fill(&_l2BlockNumber)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil || opts == nil || _l2BlockNumber == nil {
			return
		}

		_L2OutputOracle.ComputeL2Timestamp(opts, _l2BlockNumber)
	})
}

func Fuzz_Nosy_L2OutputOracleCaller_FINALIZATIONPERIODSECONDS__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleCaller
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil || opts == nil {
			return
		}

		_L2OutputOracle.FINALIZATIONPERIODSECONDS(opts)
	})
}

func Fuzz_Nosy_L2OutputOracleCaller_FinalizationPeriodSeconds__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleCaller
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil || opts == nil {
			return
		}

		_L2OutputOracle.FinalizationPeriodSeconds(opts)
	})
}

func Fuzz_Nosy_L2OutputOracleCaller_GetL2Output__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleCaller
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var _l2OutputIndex *big.Int
		fill_err = tp.Fill(&_l2OutputIndex)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil || opts == nil || _l2OutputIndex == nil {
			return
		}

		_L2OutputOracle.GetL2Output(opts, _l2OutputIndex)
	})
}

func Fuzz_Nosy_L2OutputOracleCaller_GetL2OutputAfter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleCaller
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var _l2BlockNumber *big.Int
		fill_err = tp.Fill(&_l2BlockNumber)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil || opts == nil || _l2BlockNumber == nil {
			return
		}

		_L2OutputOracle.GetL2OutputAfter(opts, _l2BlockNumber)
	})
}

func Fuzz_Nosy_L2OutputOracleCaller_GetL2OutputIndexAfter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleCaller
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var _l2BlockNumber *big.Int
		fill_err = tp.Fill(&_l2BlockNumber)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil || opts == nil || _l2BlockNumber == nil {
			return
		}

		_L2OutputOracle.GetL2OutputIndexAfter(opts, _l2BlockNumber)
	})
}

func Fuzz_Nosy_L2OutputOracleCaller_L2BLOCKTIME__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleCaller
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil || opts == nil {
			return
		}

		_L2OutputOracle.L2BLOCKTIME(opts)
	})
}

func Fuzz_Nosy_L2OutputOracleCaller_L2BlockTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleCaller
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil || opts == nil {
			return
		}

		_L2OutputOracle.L2BlockTime(opts)
	})
}

func Fuzz_Nosy_L2OutputOracleCaller_LatestBlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleCaller
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil || opts == nil {
			return
		}

		_L2OutputOracle.LatestBlockNumber(opts)
	})
}

func Fuzz_Nosy_L2OutputOracleCaller_LatestOutputIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleCaller
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil || opts == nil {
			return
		}

		_L2OutputOracle.LatestOutputIndex(opts)
	})
}

func Fuzz_Nosy_L2OutputOracleCaller_NextBlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleCaller
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil || opts == nil {
			return
		}

		_L2OutputOracle.NextBlockNumber(opts)
	})
}

func Fuzz_Nosy_L2OutputOracleCaller_NextOutputIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleCaller
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil || opts == nil {
			return
		}

		_L2OutputOracle.NextOutputIndex(opts)
	})
}

func Fuzz_Nosy_L2OutputOracleCaller_PROPOSER__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleCaller
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil || opts == nil {
			return
		}

		_L2OutputOracle.PROPOSER(opts)
	})
}

func Fuzz_Nosy_L2OutputOracleCaller_Proposer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleCaller
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil || opts == nil {
			return
		}

		_L2OutputOracle.Proposer(opts)
	})
}

func Fuzz_Nosy_L2OutputOracleCaller_SUBMISSIONINTERVAL__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleCaller
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil || opts == nil {
			return
		}

		_L2OutputOracle.SUBMISSIONINTERVAL(opts)
	})
}

func Fuzz_Nosy_L2OutputOracleCaller_StartingBlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleCaller
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil || opts == nil {
			return
		}

		_L2OutputOracle.StartingBlockNumber(opts)
	})
}

func Fuzz_Nosy_L2OutputOracleCaller_StartingTimestamp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleCaller
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil || opts == nil {
			return
		}

		_L2OutputOracle.StartingTimestamp(opts)
	})
}

func Fuzz_Nosy_L2OutputOracleCaller_SubmissionInterval__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleCaller
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil || opts == nil {
			return
		}

		_L2OutputOracle.SubmissionInterval(opts)
	})
}

func Fuzz_Nosy_L2OutputOracleCaller_Version__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleCaller
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil || opts == nil {
			return
		}

		_L2OutputOracle.Version(opts)
	})
}

// skipping Fuzz_Nosy_L2OutputOracleCallerRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

func Fuzz_Nosy_L2OutputOracleCallerSession_CHALLENGER__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleCallerSession
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil {
			return
		}

		_L2OutputOracle.CHALLENGER()
	})
}

func Fuzz_Nosy_L2OutputOracleCallerSession_Challenger__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleCallerSession
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil {
			return
		}

		_L2OutputOracle.Challenger()
	})
}

func Fuzz_Nosy_L2OutputOracleCallerSession_ComputeL2Timestamp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleCallerSession
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		var _l2BlockNumber *big.Int
		fill_err = tp.Fill(&_l2BlockNumber)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil || _l2BlockNumber == nil {
			return
		}

		_L2OutputOracle.ComputeL2Timestamp(_l2BlockNumber)
	})
}

func Fuzz_Nosy_L2OutputOracleCallerSession_FINALIZATIONPERIODSECONDS__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleCallerSession
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil {
			return
		}

		_L2OutputOracle.FINALIZATIONPERIODSECONDS()
	})
}

func Fuzz_Nosy_L2OutputOracleCallerSession_FinalizationPeriodSeconds__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleCallerSession
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil {
			return
		}

		_L2OutputOracle.FinalizationPeriodSeconds()
	})
}

func Fuzz_Nosy_L2OutputOracleCallerSession_GetL2Output__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleCallerSession
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		var _l2OutputIndex *big.Int
		fill_err = tp.Fill(&_l2OutputIndex)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil || _l2OutputIndex == nil {
			return
		}

		_L2OutputOracle.GetL2Output(_l2OutputIndex)
	})
}

func Fuzz_Nosy_L2OutputOracleCallerSession_GetL2OutputAfter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleCallerSession
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		var _l2BlockNumber *big.Int
		fill_err = tp.Fill(&_l2BlockNumber)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil || _l2BlockNumber == nil {
			return
		}

		_L2OutputOracle.GetL2OutputAfter(_l2BlockNumber)
	})
}

func Fuzz_Nosy_L2OutputOracleCallerSession_GetL2OutputIndexAfter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleCallerSession
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		var _l2BlockNumber *big.Int
		fill_err = tp.Fill(&_l2BlockNumber)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil || _l2BlockNumber == nil {
			return
		}

		_L2OutputOracle.GetL2OutputIndexAfter(_l2BlockNumber)
	})
}

func Fuzz_Nosy_L2OutputOracleCallerSession_L2BLOCKTIME__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleCallerSession
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil {
			return
		}

		_L2OutputOracle.L2BLOCKTIME()
	})
}

func Fuzz_Nosy_L2OutputOracleCallerSession_L2BlockTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleCallerSession
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil {
			return
		}

		_L2OutputOracle.L2BlockTime()
	})
}

func Fuzz_Nosy_L2OutputOracleCallerSession_LatestBlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleCallerSession
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil {
			return
		}

		_L2OutputOracle.LatestBlockNumber()
	})
}

func Fuzz_Nosy_L2OutputOracleCallerSession_LatestOutputIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleCallerSession
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil {
			return
		}

		_L2OutputOracle.LatestOutputIndex()
	})
}

func Fuzz_Nosy_L2OutputOracleCallerSession_NextBlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleCallerSession
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil {
			return
		}

		_L2OutputOracle.NextBlockNumber()
	})
}

func Fuzz_Nosy_L2OutputOracleCallerSession_NextOutputIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleCallerSession
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil {
			return
		}

		_L2OutputOracle.NextOutputIndex()
	})
}

func Fuzz_Nosy_L2OutputOracleCallerSession_PROPOSER__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleCallerSession
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil {
			return
		}

		_L2OutputOracle.PROPOSER()
	})
}

func Fuzz_Nosy_L2OutputOracleCallerSession_Proposer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleCallerSession
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil {
			return
		}

		_L2OutputOracle.Proposer()
	})
}

func Fuzz_Nosy_L2OutputOracleCallerSession_SUBMISSIONINTERVAL__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleCallerSession
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil {
			return
		}

		_L2OutputOracle.SUBMISSIONINTERVAL()
	})
}

func Fuzz_Nosy_L2OutputOracleCallerSession_StartingBlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleCallerSession
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil {
			return
		}

		_L2OutputOracle.StartingBlockNumber()
	})
}

func Fuzz_Nosy_L2OutputOracleCallerSession_StartingTimestamp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleCallerSession
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil {
			return
		}

		_L2OutputOracle.StartingTimestamp()
	})
}

func Fuzz_Nosy_L2OutputOracleCallerSession_SubmissionInterval__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleCallerSession
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil {
			return
		}

		_L2OutputOracle.SubmissionInterval()
	})
}

func Fuzz_Nosy_L2OutputOracleCallerSession_Version__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleCallerSession
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil {
			return
		}

		_L2OutputOracle.Version()
	})
}

func Fuzz_Nosy_L2OutputOracleFilterer_FilterInitialized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleFilterer
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil || opts == nil {
			return
		}

		_L2OutputOracle.FilterInitialized(opts)
	})
}

func Fuzz_Nosy_L2OutputOracleFilterer_FilterOutputProposed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleFilterer
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var outputRoot [][32]byte
		fill_err = tp.Fill(&outputRoot)
		if fill_err != nil {
			return
		}
		var l2OutputIndex []*big.Int
		fill_err = tp.Fill(&l2OutputIndex)
		if fill_err != nil {
			return
		}
		var l2BlockNumber []*big.Int
		fill_err = tp.Fill(&l2BlockNumber)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil || opts == nil {
			return
		}

		_L2OutputOracle.FilterOutputProposed(opts, outputRoot, l2OutputIndex, l2BlockNumber)
	})
}

func Fuzz_Nosy_L2OutputOracleFilterer_FilterOutputsDeleted__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleFilterer
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var prevNextOutputIndex []*big.Int
		fill_err = tp.Fill(&prevNextOutputIndex)
		if fill_err != nil {
			return
		}
		var newNextOutputIndex []*big.Int
		fill_err = tp.Fill(&newNextOutputIndex)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil || opts == nil {
			return
		}

		_L2OutputOracle.FilterOutputsDeleted(opts, prevNextOutputIndex, newNextOutputIndex)
	})
}

func Fuzz_Nosy_L2OutputOracleFilterer_ParseInitialized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleFilterer
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil {
			return
		}

		_L2OutputOracle.ParseInitialized(log)
	})
}

func Fuzz_Nosy_L2OutputOracleFilterer_ParseOutputProposed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleFilterer
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil {
			return
		}

		_L2OutputOracle.ParseOutputProposed(log)
	})
}

func Fuzz_Nosy_L2OutputOracleFilterer_ParseOutputsDeleted__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleFilterer
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil {
			return
		}

		_L2OutputOracle.ParseOutputsDeleted(log)
	})
}

// skipping Fuzz_Nosy_L2OutputOracleFilterer_WatchInitialized__ because parameters include func, chan, or unsupported interface: chan<- *github.com/ethereum-optimism/optimism/op-node/bindings.L2OutputOracleInitialized

// skipping Fuzz_Nosy_L2OutputOracleFilterer_WatchOutputProposed__ because parameters include func, chan, or unsupported interface: chan<- *github.com/ethereum-optimism/optimism/op-node/bindings.L2OutputOracleOutputProposed

// skipping Fuzz_Nosy_L2OutputOracleFilterer_WatchOutputsDeleted__ because parameters include func, chan, or unsupported interface: chan<- *github.com/ethereum-optimism/optimism/op-node/bindings.L2OutputOracleOutputsDeleted

func Fuzz_Nosy_L2OutputOracleInitializedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *L2OutputOracleInitializedIterator
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

func Fuzz_Nosy_L2OutputOracleInitializedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *L2OutputOracleInitializedIterator
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

func Fuzz_Nosy_L2OutputOracleInitializedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *L2OutputOracleInitializedIterator
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

func Fuzz_Nosy_L2OutputOracleOutputProposedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *L2OutputOracleOutputProposedIterator
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

func Fuzz_Nosy_L2OutputOracleOutputProposedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *L2OutputOracleOutputProposedIterator
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

func Fuzz_Nosy_L2OutputOracleOutputProposedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *L2OutputOracleOutputProposedIterator
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

func Fuzz_Nosy_L2OutputOracleOutputsDeletedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *L2OutputOracleOutputsDeletedIterator
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

func Fuzz_Nosy_L2OutputOracleOutputsDeletedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *L2OutputOracleOutputsDeletedIterator
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

func Fuzz_Nosy_L2OutputOracleOutputsDeletedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *L2OutputOracleOutputsDeletedIterator
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

// skipping Fuzz_Nosy_L2OutputOracleRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

// skipping Fuzz_Nosy_L2OutputOracleRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_L2OutputOracleRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleRaw
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil || opts == nil {
			return
		}

		_L2OutputOracle.Transfer(opts)
	})
}

func Fuzz_Nosy_L2OutputOracleSession_CHALLENGER__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleSession
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil {
			return
		}

		_L2OutputOracle.CHALLENGER()
	})
}

func Fuzz_Nosy_L2OutputOracleSession_Challenger__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleSession
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil {
			return
		}

		_L2OutputOracle.Challenger()
	})
}

func Fuzz_Nosy_L2OutputOracleSession_ComputeL2Timestamp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleSession
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		var _l2BlockNumber *big.Int
		fill_err = tp.Fill(&_l2BlockNumber)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil || _l2BlockNumber == nil {
			return
		}

		_L2OutputOracle.ComputeL2Timestamp(_l2BlockNumber)
	})
}

func Fuzz_Nosy_L2OutputOracleSession_DeleteL2Outputs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleSession
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		var _l2OutputIndex *big.Int
		fill_err = tp.Fill(&_l2OutputIndex)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil || _l2OutputIndex == nil {
			return
		}

		_L2OutputOracle.DeleteL2Outputs(_l2OutputIndex)
	})
}

func Fuzz_Nosy_L2OutputOracleSession_FINALIZATIONPERIODSECONDS__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleSession
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil {
			return
		}

		_L2OutputOracle.FINALIZATIONPERIODSECONDS()
	})
}

func Fuzz_Nosy_L2OutputOracleSession_FinalizationPeriodSeconds__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleSession
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil {
			return
		}

		_L2OutputOracle.FinalizationPeriodSeconds()
	})
}

func Fuzz_Nosy_L2OutputOracleSession_GetL2Output__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleSession
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		var _l2OutputIndex *big.Int
		fill_err = tp.Fill(&_l2OutputIndex)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil || _l2OutputIndex == nil {
			return
		}

		_L2OutputOracle.GetL2Output(_l2OutputIndex)
	})
}

func Fuzz_Nosy_L2OutputOracleSession_GetL2OutputAfter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleSession
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		var _l2BlockNumber *big.Int
		fill_err = tp.Fill(&_l2BlockNumber)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil || _l2BlockNumber == nil {
			return
		}

		_L2OutputOracle.GetL2OutputAfter(_l2BlockNumber)
	})
}

func Fuzz_Nosy_L2OutputOracleSession_GetL2OutputIndexAfter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleSession
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		var _l2BlockNumber *big.Int
		fill_err = tp.Fill(&_l2BlockNumber)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil || _l2BlockNumber == nil {
			return
		}

		_L2OutputOracle.GetL2OutputIndexAfter(_l2BlockNumber)
	})
}

func Fuzz_Nosy_L2OutputOracleSession_Initialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleSession
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		var _submissionInterval *big.Int
		fill_err = tp.Fill(&_submissionInterval)
		if fill_err != nil {
			return
		}
		var _l2BlockTime *big.Int
		fill_err = tp.Fill(&_l2BlockTime)
		if fill_err != nil {
			return
		}
		var _startingBlockNumber *big.Int
		fill_err = tp.Fill(&_startingBlockNumber)
		if fill_err != nil {
			return
		}
		var _startingTimestamp *big.Int
		fill_err = tp.Fill(&_startingTimestamp)
		if fill_err != nil {
			return
		}
		var _proposer common.Address
		fill_err = tp.Fill(&_proposer)
		if fill_err != nil {
			return
		}
		var _challenger common.Address
		fill_err = tp.Fill(&_challenger)
		if fill_err != nil {
			return
		}
		var _finalizationPeriodSeconds *big.Int
		fill_err = tp.Fill(&_finalizationPeriodSeconds)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil || _submissionInterval == nil || _l2BlockTime == nil || _startingBlockNumber == nil || _startingTimestamp == nil || _finalizationPeriodSeconds == nil {
			return
		}

		_L2OutputOracle.Initialize(_submissionInterval, _l2BlockTime, _startingBlockNumber, _startingTimestamp, _proposer, _challenger, _finalizationPeriodSeconds)
	})
}

func Fuzz_Nosy_L2OutputOracleSession_L2BLOCKTIME__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleSession
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil {
			return
		}

		_L2OutputOracle.L2BLOCKTIME()
	})
}

func Fuzz_Nosy_L2OutputOracleSession_L2BlockTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleSession
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil {
			return
		}

		_L2OutputOracle.L2BlockTime()
	})
}

func Fuzz_Nosy_L2OutputOracleSession_LatestBlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleSession
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil {
			return
		}

		_L2OutputOracle.LatestBlockNumber()
	})
}

func Fuzz_Nosy_L2OutputOracleSession_LatestOutputIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleSession
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil {
			return
		}

		_L2OutputOracle.LatestOutputIndex()
	})
}

func Fuzz_Nosy_L2OutputOracleSession_NextBlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleSession
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil {
			return
		}

		_L2OutputOracle.NextBlockNumber()
	})
}

func Fuzz_Nosy_L2OutputOracleSession_NextOutputIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleSession
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil {
			return
		}

		_L2OutputOracle.NextOutputIndex()
	})
}

func Fuzz_Nosy_L2OutputOracleSession_PROPOSER__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleSession
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil {
			return
		}

		_L2OutputOracle.PROPOSER()
	})
}

func Fuzz_Nosy_L2OutputOracleSession_ProposeL2Output__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleSession
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		var _outputRoot [32]byte
		fill_err = tp.Fill(&_outputRoot)
		if fill_err != nil {
			return
		}
		var _l2BlockNumber *big.Int
		fill_err = tp.Fill(&_l2BlockNumber)
		if fill_err != nil {
			return
		}
		var _l1BlockHash [32]byte
		fill_err = tp.Fill(&_l1BlockHash)
		if fill_err != nil {
			return
		}
		var _l1BlockNumber *big.Int
		fill_err = tp.Fill(&_l1BlockNumber)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil || _l2BlockNumber == nil || _l1BlockNumber == nil {
			return
		}

		_L2OutputOracle.ProposeL2Output(_outputRoot, _l2BlockNumber, _l1BlockHash, _l1BlockNumber)
	})
}

func Fuzz_Nosy_L2OutputOracleSession_Proposer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleSession
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil {
			return
		}

		_L2OutputOracle.Proposer()
	})
}

func Fuzz_Nosy_L2OutputOracleSession_SUBMISSIONINTERVAL__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleSession
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil {
			return
		}

		_L2OutputOracle.SUBMISSIONINTERVAL()
	})
}

func Fuzz_Nosy_L2OutputOracleSession_StartingBlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleSession
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil {
			return
		}

		_L2OutputOracle.StartingBlockNumber()
	})
}

func Fuzz_Nosy_L2OutputOracleSession_StartingTimestamp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleSession
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil {
			return
		}

		_L2OutputOracle.StartingTimestamp()
	})
}

func Fuzz_Nosy_L2OutputOracleSession_SubmissionInterval__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleSession
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil {
			return
		}

		_L2OutputOracle.SubmissionInterval()
	})
}

func Fuzz_Nosy_L2OutputOracleSession_Version__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleSession
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil {
			return
		}

		_L2OutputOracle.Version()
	})
}

func Fuzz_Nosy_L2OutputOracleTransactor_DeleteL2Outputs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleTransactor
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var _l2OutputIndex *big.Int
		fill_err = tp.Fill(&_l2OutputIndex)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil || opts == nil || _l2OutputIndex == nil {
			return
		}

		_L2OutputOracle.DeleteL2Outputs(opts, _l2OutputIndex)
	})
}

func Fuzz_Nosy_L2OutputOracleTransactor_Initialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleTransactor
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var _submissionInterval *big.Int
		fill_err = tp.Fill(&_submissionInterval)
		if fill_err != nil {
			return
		}
		var _l2BlockTime *big.Int
		fill_err = tp.Fill(&_l2BlockTime)
		if fill_err != nil {
			return
		}
		var _startingBlockNumber *big.Int
		fill_err = tp.Fill(&_startingBlockNumber)
		if fill_err != nil {
			return
		}
		var _startingTimestamp *big.Int
		fill_err = tp.Fill(&_startingTimestamp)
		if fill_err != nil {
			return
		}
		var _proposer common.Address
		fill_err = tp.Fill(&_proposer)
		if fill_err != nil {
			return
		}
		var _challenger common.Address
		fill_err = tp.Fill(&_challenger)
		if fill_err != nil {
			return
		}
		var _finalizationPeriodSeconds *big.Int
		fill_err = tp.Fill(&_finalizationPeriodSeconds)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil || opts == nil || _submissionInterval == nil || _l2BlockTime == nil || _startingBlockNumber == nil || _startingTimestamp == nil || _finalizationPeriodSeconds == nil {
			return
		}

		_L2OutputOracle.Initialize(opts, _submissionInterval, _l2BlockTime, _startingBlockNumber, _startingTimestamp, _proposer, _challenger, _finalizationPeriodSeconds)
	})
}

func Fuzz_Nosy_L2OutputOracleTransactor_ProposeL2Output__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleTransactor
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var _outputRoot [32]byte
		fill_err = tp.Fill(&_outputRoot)
		if fill_err != nil {
			return
		}
		var _l2BlockNumber *big.Int
		fill_err = tp.Fill(&_l2BlockNumber)
		if fill_err != nil {
			return
		}
		var _l1BlockHash [32]byte
		fill_err = tp.Fill(&_l1BlockHash)
		if fill_err != nil {
			return
		}
		var _l1BlockNumber *big.Int
		fill_err = tp.Fill(&_l1BlockNumber)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil || opts == nil || _l2BlockNumber == nil || _l1BlockNumber == nil {
			return
		}

		_L2OutputOracle.ProposeL2Output(opts, _outputRoot, _l2BlockNumber, _l1BlockHash, _l1BlockNumber)
	})
}

// skipping Fuzz_Nosy_L2OutputOracleTransactorRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_L2OutputOracleTransactorRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleTransactorRaw
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil || opts == nil {
			return
		}

		_L2OutputOracle.Transfer(opts)
	})
}

func Fuzz_Nosy_L2OutputOracleTransactorSession_DeleteL2Outputs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleTransactorSession
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		var _l2OutputIndex *big.Int
		fill_err = tp.Fill(&_l2OutputIndex)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil || _l2OutputIndex == nil {
			return
		}

		_L2OutputOracle.DeleteL2Outputs(_l2OutputIndex)
	})
}

func Fuzz_Nosy_L2OutputOracleTransactorSession_Initialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleTransactorSession
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		var _submissionInterval *big.Int
		fill_err = tp.Fill(&_submissionInterval)
		if fill_err != nil {
			return
		}
		var _l2BlockTime *big.Int
		fill_err = tp.Fill(&_l2BlockTime)
		if fill_err != nil {
			return
		}
		var _startingBlockNumber *big.Int
		fill_err = tp.Fill(&_startingBlockNumber)
		if fill_err != nil {
			return
		}
		var _startingTimestamp *big.Int
		fill_err = tp.Fill(&_startingTimestamp)
		if fill_err != nil {
			return
		}
		var _proposer common.Address
		fill_err = tp.Fill(&_proposer)
		if fill_err != nil {
			return
		}
		var _challenger common.Address
		fill_err = tp.Fill(&_challenger)
		if fill_err != nil {
			return
		}
		var _finalizationPeriodSeconds *big.Int
		fill_err = tp.Fill(&_finalizationPeriodSeconds)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil || _submissionInterval == nil || _l2BlockTime == nil || _startingBlockNumber == nil || _startingTimestamp == nil || _finalizationPeriodSeconds == nil {
			return
		}

		_L2OutputOracle.Initialize(_submissionInterval, _l2BlockTime, _startingBlockNumber, _startingTimestamp, _proposer, _challenger, _finalizationPeriodSeconds)
	})
}

func Fuzz_Nosy_L2OutputOracleTransactorSession_ProposeL2Output__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2OutputOracle *L2OutputOracleTransactorSession
		fill_err = tp.Fill(&_L2OutputOracle)
		if fill_err != nil {
			return
		}
		var _outputRoot [32]byte
		fill_err = tp.Fill(&_outputRoot)
		if fill_err != nil {
			return
		}
		var _l2BlockNumber *big.Int
		fill_err = tp.Fill(&_l2BlockNumber)
		if fill_err != nil {
			return
		}
		var _l1BlockHash [32]byte
		fill_err = tp.Fill(&_l1BlockHash)
		if fill_err != nil {
			return
		}
		var _l1BlockNumber *big.Int
		fill_err = tp.Fill(&_l1BlockNumber)
		if fill_err != nil {
			return
		}
		if _L2OutputOracle == nil || _l2BlockNumber == nil || _l1BlockNumber == nil {
			return
		}

		_L2OutputOracle.ProposeL2Output(_outputRoot, _l2BlockNumber, _l1BlockHash, _l1BlockNumber)
	})
}

func Fuzz_Nosy_L2ToL1MessagePasserCaller_MESSAGEVERSION__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2ToL1MessagePasser *L2ToL1MessagePasserCaller
		fill_err = tp.Fill(&_L2ToL1MessagePasser)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L2ToL1MessagePasser == nil || opts == nil {
			return
		}

		_L2ToL1MessagePasser.MESSAGEVERSION(opts)
	})
}

func Fuzz_Nosy_L2ToL1MessagePasserCaller_MessageNonce__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2ToL1MessagePasser *L2ToL1MessagePasserCaller
		fill_err = tp.Fill(&_L2ToL1MessagePasser)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L2ToL1MessagePasser == nil || opts == nil {
			return
		}

		_L2ToL1MessagePasser.MessageNonce(opts)
	})
}

func Fuzz_Nosy_L2ToL1MessagePasserCaller_SentMessages__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2ToL1MessagePasser *L2ToL1MessagePasserCaller
		fill_err = tp.Fill(&_L2ToL1MessagePasser)
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
		if _L2ToL1MessagePasser == nil || opts == nil {
			return
		}

		_L2ToL1MessagePasser.SentMessages(opts, arg0)
	})
}

func Fuzz_Nosy_L2ToL1MessagePasserCaller_Version__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2ToL1MessagePasser *L2ToL1MessagePasserCaller
		fill_err = tp.Fill(&_L2ToL1MessagePasser)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L2ToL1MessagePasser == nil || opts == nil {
			return
		}

		_L2ToL1MessagePasser.Version(opts)
	})
}

// skipping Fuzz_Nosy_L2ToL1MessagePasserCallerRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

func Fuzz_Nosy_L2ToL1MessagePasserCallerSession_MESSAGEVERSION__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2ToL1MessagePasser *L2ToL1MessagePasserCallerSession
		fill_err = tp.Fill(&_L2ToL1MessagePasser)
		if fill_err != nil {
			return
		}
		if _L2ToL1MessagePasser == nil {
			return
		}

		_L2ToL1MessagePasser.MESSAGEVERSION()
	})
}

func Fuzz_Nosy_L2ToL1MessagePasserCallerSession_MessageNonce__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2ToL1MessagePasser *L2ToL1MessagePasserCallerSession
		fill_err = tp.Fill(&_L2ToL1MessagePasser)
		if fill_err != nil {
			return
		}
		if _L2ToL1MessagePasser == nil {
			return
		}

		_L2ToL1MessagePasser.MessageNonce()
	})
}

func Fuzz_Nosy_L2ToL1MessagePasserCallerSession_SentMessages__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2ToL1MessagePasser *L2ToL1MessagePasserCallerSession
		fill_err = tp.Fill(&_L2ToL1MessagePasser)
		if fill_err != nil {
			return
		}
		var arg0 [32]byte
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _L2ToL1MessagePasser == nil {
			return
		}

		_L2ToL1MessagePasser.SentMessages(arg0)
	})
}

func Fuzz_Nosy_L2ToL1MessagePasserCallerSession_Version__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2ToL1MessagePasser *L2ToL1MessagePasserCallerSession
		fill_err = tp.Fill(&_L2ToL1MessagePasser)
		if fill_err != nil {
			return
		}
		if _L2ToL1MessagePasser == nil {
			return
		}

		_L2ToL1MessagePasser.Version()
	})
}

func Fuzz_Nosy_L2ToL1MessagePasserFilterer_FilterMessagePassed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2ToL1MessagePasser *L2ToL1MessagePasserFilterer
		fill_err = tp.Fill(&_L2ToL1MessagePasser)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var nonce []*big.Int
		fill_err = tp.Fill(&nonce)
		if fill_err != nil {
			return
		}
		var sender []common.Address
		fill_err = tp.Fill(&sender)
		if fill_err != nil {
			return
		}
		var t5 []common.Address
		fill_err = tp.Fill(&t5)
		if fill_err != nil {
			return
		}
		if _L2ToL1MessagePasser == nil || opts == nil {
			return
		}

		_L2ToL1MessagePasser.FilterMessagePassed(opts, nonce, sender, t5)
	})
}

func Fuzz_Nosy_L2ToL1MessagePasserFilterer_FilterWithdrawerBalanceBurnt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2ToL1MessagePasser *L2ToL1MessagePasserFilterer
		fill_err = tp.Fill(&_L2ToL1MessagePasser)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var amount []*big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if _L2ToL1MessagePasser == nil || opts == nil {
			return
		}

		_L2ToL1MessagePasser.FilterWithdrawerBalanceBurnt(opts, amount)
	})
}

func Fuzz_Nosy_L2ToL1MessagePasserFilterer_ParseMessagePassed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2ToL1MessagePasser *L2ToL1MessagePasserFilterer
		fill_err = tp.Fill(&_L2ToL1MessagePasser)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _L2ToL1MessagePasser == nil {
			return
		}

		_L2ToL1MessagePasser.ParseMessagePassed(log)
	})
}

func Fuzz_Nosy_L2ToL1MessagePasserFilterer_ParseWithdrawerBalanceBurnt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2ToL1MessagePasser *L2ToL1MessagePasserFilterer
		fill_err = tp.Fill(&_L2ToL1MessagePasser)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _L2ToL1MessagePasser == nil {
			return
		}

		_L2ToL1MessagePasser.ParseWithdrawerBalanceBurnt(log)
	})
}

// skipping Fuzz_Nosy_L2ToL1MessagePasserFilterer_WatchMessagePassed__ because parameters include func, chan, or unsupported interface: chan<- *github.com/ethereum-optimism/optimism/op-node/bindings.L2ToL1MessagePasserMessagePassed

// skipping Fuzz_Nosy_L2ToL1MessagePasserFilterer_WatchWithdrawerBalanceBurnt__ because parameters include func, chan, or unsupported interface: chan<- *github.com/ethereum-optimism/optimism/op-node/bindings.L2ToL1MessagePasserWithdrawerBalanceBurnt

func Fuzz_Nosy_L2ToL1MessagePasserMessagePassedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *L2ToL1MessagePasserMessagePassedIterator
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

func Fuzz_Nosy_L2ToL1MessagePasserMessagePassedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *L2ToL1MessagePasserMessagePassedIterator
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

func Fuzz_Nosy_L2ToL1MessagePasserMessagePassedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *L2ToL1MessagePasserMessagePassedIterator
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

// skipping Fuzz_Nosy_L2ToL1MessagePasserRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

// skipping Fuzz_Nosy_L2ToL1MessagePasserRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_L2ToL1MessagePasserRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2ToL1MessagePasser *L2ToL1MessagePasserRaw
		fill_err = tp.Fill(&_L2ToL1MessagePasser)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L2ToL1MessagePasser == nil || opts == nil {
			return
		}

		_L2ToL1MessagePasser.Transfer(opts)
	})
}

func Fuzz_Nosy_L2ToL1MessagePasserSession_Burn__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2ToL1MessagePasser *L2ToL1MessagePasserSession
		fill_err = tp.Fill(&_L2ToL1MessagePasser)
		if fill_err != nil {
			return
		}
		if _L2ToL1MessagePasser == nil {
			return
		}

		_L2ToL1MessagePasser.Burn()
	})
}

func Fuzz_Nosy_L2ToL1MessagePasserSession_InitiateWithdrawal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2ToL1MessagePasser *L2ToL1MessagePasserSession
		fill_err = tp.Fill(&_L2ToL1MessagePasser)
		if fill_err != nil {
			return
		}
		var _target common.Address
		fill_err = tp.Fill(&_target)
		if fill_err != nil {
			return
		}
		var _gasLimit *big.Int
		fill_err = tp.Fill(&_gasLimit)
		if fill_err != nil {
			return
		}
		var _data []byte
		fill_err = tp.Fill(&_data)
		if fill_err != nil {
			return
		}
		if _L2ToL1MessagePasser == nil || _gasLimit == nil {
			return
		}

		_L2ToL1MessagePasser.InitiateWithdrawal(_target, _gasLimit, _data)
	})
}

func Fuzz_Nosy_L2ToL1MessagePasserSession_MESSAGEVERSION__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2ToL1MessagePasser *L2ToL1MessagePasserSession
		fill_err = tp.Fill(&_L2ToL1MessagePasser)
		if fill_err != nil {
			return
		}
		if _L2ToL1MessagePasser == nil {
			return
		}

		_L2ToL1MessagePasser.MESSAGEVERSION()
	})
}

func Fuzz_Nosy_L2ToL1MessagePasserSession_MessageNonce__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2ToL1MessagePasser *L2ToL1MessagePasserSession
		fill_err = tp.Fill(&_L2ToL1MessagePasser)
		if fill_err != nil {
			return
		}
		if _L2ToL1MessagePasser == nil {
			return
		}

		_L2ToL1MessagePasser.MessageNonce()
	})
}

func Fuzz_Nosy_L2ToL1MessagePasserSession_Receive__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2ToL1MessagePasser *L2ToL1MessagePasserSession
		fill_err = tp.Fill(&_L2ToL1MessagePasser)
		if fill_err != nil {
			return
		}
		if _L2ToL1MessagePasser == nil {
			return
		}

		_L2ToL1MessagePasser.Receive()
	})
}

func Fuzz_Nosy_L2ToL1MessagePasserSession_SentMessages__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2ToL1MessagePasser *L2ToL1MessagePasserSession
		fill_err = tp.Fill(&_L2ToL1MessagePasser)
		if fill_err != nil {
			return
		}
		var arg0 [32]byte
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _L2ToL1MessagePasser == nil {
			return
		}

		_L2ToL1MessagePasser.SentMessages(arg0)
	})
}

func Fuzz_Nosy_L2ToL1MessagePasserSession_Version__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2ToL1MessagePasser *L2ToL1MessagePasserSession
		fill_err = tp.Fill(&_L2ToL1MessagePasser)
		if fill_err != nil {
			return
		}
		if _L2ToL1MessagePasser == nil {
			return
		}

		_L2ToL1MessagePasser.Version()
	})
}

func Fuzz_Nosy_L2ToL1MessagePasserTransactor_Burn__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2ToL1MessagePasser *L2ToL1MessagePasserTransactor
		fill_err = tp.Fill(&_L2ToL1MessagePasser)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L2ToL1MessagePasser == nil || opts == nil {
			return
		}

		_L2ToL1MessagePasser.Burn(opts)
	})
}

func Fuzz_Nosy_L2ToL1MessagePasserTransactor_InitiateWithdrawal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2ToL1MessagePasser *L2ToL1MessagePasserTransactor
		fill_err = tp.Fill(&_L2ToL1MessagePasser)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var _target common.Address
		fill_err = tp.Fill(&_target)
		if fill_err != nil {
			return
		}
		var _gasLimit *big.Int
		fill_err = tp.Fill(&_gasLimit)
		if fill_err != nil {
			return
		}
		var _data []byte
		fill_err = tp.Fill(&_data)
		if fill_err != nil {
			return
		}
		if _L2ToL1MessagePasser == nil || opts == nil || _gasLimit == nil {
			return
		}

		_L2ToL1MessagePasser.InitiateWithdrawal(opts, _target, _gasLimit, _data)
	})
}

func Fuzz_Nosy_L2ToL1MessagePasserTransactor_Receive__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2ToL1MessagePasser *L2ToL1MessagePasserTransactor
		fill_err = tp.Fill(&_L2ToL1MessagePasser)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L2ToL1MessagePasser == nil || opts == nil {
			return
		}

		_L2ToL1MessagePasser.Receive(opts)
	})
}

// skipping Fuzz_Nosy_L2ToL1MessagePasserTransactorRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_L2ToL1MessagePasserTransactorRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2ToL1MessagePasser *L2ToL1MessagePasserTransactorRaw
		fill_err = tp.Fill(&_L2ToL1MessagePasser)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L2ToL1MessagePasser == nil || opts == nil {
			return
		}

		_L2ToL1MessagePasser.Transfer(opts)
	})
}

func Fuzz_Nosy_L2ToL1MessagePasserTransactorSession_Burn__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2ToL1MessagePasser *L2ToL1MessagePasserTransactorSession
		fill_err = tp.Fill(&_L2ToL1MessagePasser)
		if fill_err != nil {
			return
		}
		if _L2ToL1MessagePasser == nil {
			return
		}

		_L2ToL1MessagePasser.Burn()
	})
}

func Fuzz_Nosy_L2ToL1MessagePasserTransactorSession_InitiateWithdrawal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2ToL1MessagePasser *L2ToL1MessagePasserTransactorSession
		fill_err = tp.Fill(&_L2ToL1MessagePasser)
		if fill_err != nil {
			return
		}
		var _target common.Address
		fill_err = tp.Fill(&_target)
		if fill_err != nil {
			return
		}
		var _gasLimit *big.Int
		fill_err = tp.Fill(&_gasLimit)
		if fill_err != nil {
			return
		}
		var _data []byte
		fill_err = tp.Fill(&_data)
		if fill_err != nil {
			return
		}
		if _L2ToL1MessagePasser == nil || _gasLimit == nil {
			return
		}

		_L2ToL1MessagePasser.InitiateWithdrawal(_target, _gasLimit, _data)
	})
}

func Fuzz_Nosy_L2ToL1MessagePasserTransactorSession_Receive__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2ToL1MessagePasser *L2ToL1MessagePasserTransactorSession
		fill_err = tp.Fill(&_L2ToL1MessagePasser)
		if fill_err != nil {
			return
		}
		if _L2ToL1MessagePasser == nil {
			return
		}

		_L2ToL1MessagePasser.Receive()
	})
}

func Fuzz_Nosy_L2ToL1MessagePasserWithdrawerBalanceBurntIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *L2ToL1MessagePasserWithdrawerBalanceBurntIterator
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

func Fuzz_Nosy_L2ToL1MessagePasserWithdrawerBalanceBurntIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *L2ToL1MessagePasserWithdrawerBalanceBurntIterator
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

func Fuzz_Nosy_L2ToL1MessagePasserWithdrawerBalanceBurntIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *L2ToL1MessagePasserWithdrawerBalanceBurntIterator
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

func Fuzz_Nosy_OptimismPortalCaller_FinalizedWithdrawals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal *OptimismPortalCaller
		fill_err = tp.Fill(&_OptimismPortal)
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
		if _OptimismPortal == nil || opts == nil {
			return
		}

		_OptimismPortal.FinalizedWithdrawals(opts, arg0)
	})
}

func Fuzz_Nosy_OptimismPortalCaller_Guardian__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal *OptimismPortalCaller
		fill_err = tp.Fill(&_OptimismPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OptimismPortal == nil || opts == nil {
			return
		}

		_OptimismPortal.Guardian(opts)
	})
}

func Fuzz_Nosy_OptimismPortalCaller_IsOutputFinalized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal *OptimismPortalCaller
		fill_err = tp.Fill(&_OptimismPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var _l2OutputIndex *big.Int
		fill_err = tp.Fill(&_l2OutputIndex)
		if fill_err != nil {
			return
		}
		if _OptimismPortal == nil || opts == nil || _l2OutputIndex == nil {
			return
		}

		_OptimismPortal.IsOutputFinalized(opts, _l2OutputIndex)
	})
}

func Fuzz_Nosy_OptimismPortalCaller_L2Oracle__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal *OptimismPortalCaller
		fill_err = tp.Fill(&_OptimismPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OptimismPortal == nil || opts == nil {
			return
		}

		_OptimismPortal.L2Oracle(opts)
	})
}

func Fuzz_Nosy_OptimismPortalCaller_L2Sender__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal *OptimismPortalCaller
		fill_err = tp.Fill(&_OptimismPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OptimismPortal == nil || opts == nil {
			return
		}

		_OptimismPortal.L2Sender(opts)
	})
}

func Fuzz_Nosy_OptimismPortalCaller_MinimumGasLimit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal *OptimismPortalCaller
		fill_err = tp.Fill(&_OptimismPortal)
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
		if _OptimismPortal == nil || opts == nil {
			return
		}

		_OptimismPortal.MinimumGasLimit(opts, _byteCount)
	})
}

func Fuzz_Nosy_OptimismPortalCaller_Params__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal *OptimismPortalCaller
		fill_err = tp.Fill(&_OptimismPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OptimismPortal == nil || opts == nil {
			return
		}

		_OptimismPortal.Params(opts)
	})
}

func Fuzz_Nosy_OptimismPortalCaller_Paused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal *OptimismPortalCaller
		fill_err = tp.Fill(&_OptimismPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OptimismPortal == nil || opts == nil {
			return
		}

		_OptimismPortal.Paused(opts)
	})
}

func Fuzz_Nosy_OptimismPortalCaller_ProvenWithdrawals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal *OptimismPortalCaller
		fill_err = tp.Fill(&_OptimismPortal)
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
		if _OptimismPortal == nil || opts == nil {
			return
		}

		_OptimismPortal.ProvenWithdrawals(opts, arg0)
	})
}

func Fuzz_Nosy_OptimismPortalCaller_SuperchainConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal *OptimismPortalCaller
		fill_err = tp.Fill(&_OptimismPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OptimismPortal == nil || opts == nil {
			return
		}

		_OptimismPortal.SuperchainConfig(opts)
	})
}

func Fuzz_Nosy_OptimismPortalCaller_SystemConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal *OptimismPortalCaller
		fill_err = tp.Fill(&_OptimismPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OptimismPortal == nil || opts == nil {
			return
		}

		_OptimismPortal.SystemConfig(opts)
	})
}

func Fuzz_Nosy_OptimismPortalCaller_Version__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal *OptimismPortalCaller
		fill_err = tp.Fill(&_OptimismPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OptimismPortal == nil || opts == nil {
			return
		}

		_OptimismPortal.Version(opts)
	})
}

// skipping Fuzz_Nosy_OptimismPortalCallerRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

func Fuzz_Nosy_OptimismPortalCallerSession_FinalizedWithdrawals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal *OptimismPortalCallerSession
		fill_err = tp.Fill(&_OptimismPortal)
		if fill_err != nil {
			return
		}
		var arg0 [32]byte
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _OptimismPortal == nil {
			return
		}

		_OptimismPortal.FinalizedWithdrawals(arg0)
	})
}

func Fuzz_Nosy_OptimismPortalCallerSession_Guardian__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal *OptimismPortalCallerSession
		fill_err = tp.Fill(&_OptimismPortal)
		if fill_err != nil {
			return
		}
		if _OptimismPortal == nil {
			return
		}

		_OptimismPortal.Guardian()
	})
}

func Fuzz_Nosy_OptimismPortalCallerSession_IsOutputFinalized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal *OptimismPortalCallerSession
		fill_err = tp.Fill(&_OptimismPortal)
		if fill_err != nil {
			return
		}
		var _l2OutputIndex *big.Int
		fill_err = tp.Fill(&_l2OutputIndex)
		if fill_err != nil {
			return
		}
		if _OptimismPortal == nil || _l2OutputIndex == nil {
			return
		}

		_OptimismPortal.IsOutputFinalized(_l2OutputIndex)
	})
}

func Fuzz_Nosy_OptimismPortalCallerSession_L2Oracle__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal *OptimismPortalCallerSession
		fill_err = tp.Fill(&_OptimismPortal)
		if fill_err != nil {
			return
		}
		if _OptimismPortal == nil {
			return
		}

		_OptimismPortal.L2Oracle()
	})
}

func Fuzz_Nosy_OptimismPortalCallerSession_L2Sender__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal *OptimismPortalCallerSession
		fill_err = tp.Fill(&_OptimismPortal)
		if fill_err != nil {
			return
		}
		if _OptimismPortal == nil {
			return
		}

		_OptimismPortal.L2Sender()
	})
}

func Fuzz_Nosy_OptimismPortalCallerSession_MinimumGasLimit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal *OptimismPortalCallerSession
		fill_err = tp.Fill(&_OptimismPortal)
		if fill_err != nil {
			return
		}
		var _byteCount uint64
		fill_err = tp.Fill(&_byteCount)
		if fill_err != nil {
			return
		}
		if _OptimismPortal == nil {
			return
		}

		_OptimismPortal.MinimumGasLimit(_byteCount)
	})
}

func Fuzz_Nosy_OptimismPortalCallerSession_Params__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal *OptimismPortalCallerSession
		fill_err = tp.Fill(&_OptimismPortal)
		if fill_err != nil {
			return
		}
		if _OptimismPortal == nil {
			return
		}

		_OptimismPortal.Params()
	})
}

func Fuzz_Nosy_OptimismPortalCallerSession_Paused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal *OptimismPortalCallerSession
		fill_err = tp.Fill(&_OptimismPortal)
		if fill_err != nil {
			return
		}
		if _OptimismPortal == nil {
			return
		}

		_OptimismPortal.Paused()
	})
}

func Fuzz_Nosy_OptimismPortalCallerSession_ProvenWithdrawals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal *OptimismPortalCallerSession
		fill_err = tp.Fill(&_OptimismPortal)
		if fill_err != nil {
			return
		}
		var arg0 [32]byte
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _OptimismPortal == nil {
			return
		}

		_OptimismPortal.ProvenWithdrawals(arg0)
	})
}

func Fuzz_Nosy_OptimismPortalCallerSession_SuperchainConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal *OptimismPortalCallerSession
		fill_err = tp.Fill(&_OptimismPortal)
		if fill_err != nil {
			return
		}
		if _OptimismPortal == nil {
			return
		}

		_OptimismPortal.SuperchainConfig()
	})
}

func Fuzz_Nosy_OptimismPortalCallerSession_SystemConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal *OptimismPortalCallerSession
		fill_err = tp.Fill(&_OptimismPortal)
		if fill_err != nil {
			return
		}
		if _OptimismPortal == nil {
			return
		}

		_OptimismPortal.SystemConfig()
	})
}

func Fuzz_Nosy_OptimismPortalCallerSession_Version__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal *OptimismPortalCallerSession
		fill_err = tp.Fill(&_OptimismPortal)
		if fill_err != nil {
			return
		}
		if _OptimismPortal == nil {
			return
		}

		_OptimismPortal.Version()
	})
}

func Fuzz_Nosy_OptimismPortalFilterer_FilterInitialized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal *OptimismPortalFilterer
		fill_err = tp.Fill(&_OptimismPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OptimismPortal == nil || opts == nil {
			return
		}

		_OptimismPortal.FilterInitialized(opts)
	})
}

func Fuzz_Nosy_OptimismPortalFilterer_FilterTransactionDeposited__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal *OptimismPortalFilterer
		fill_err = tp.Fill(&_OptimismPortal)
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
		if _OptimismPortal == nil || opts == nil {
			return
		}

		_OptimismPortal.FilterTransactionDeposited(opts, from, to, version)
	})
}

func Fuzz_Nosy_OptimismPortalFilterer_FilterWithdrawalFinalized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal *OptimismPortalFilterer
		fill_err = tp.Fill(&_OptimismPortal)
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
		if _OptimismPortal == nil || opts == nil {
			return
		}

		_OptimismPortal.FilterWithdrawalFinalized(opts, withdrawalHash)
	})
}

func Fuzz_Nosy_OptimismPortalFilterer_FilterWithdrawalProven__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal *OptimismPortalFilterer
		fill_err = tp.Fill(&_OptimismPortal)
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
		if _OptimismPortal == nil || opts == nil {
			return
		}

		_OptimismPortal.FilterWithdrawalProven(opts, withdrawalHash, from, to)
	})
}

func Fuzz_Nosy_OptimismPortalFilterer_ParseInitialized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal *OptimismPortalFilterer
		fill_err = tp.Fill(&_OptimismPortal)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _OptimismPortal == nil {
			return
		}

		_OptimismPortal.ParseInitialized(log)
	})
}

func Fuzz_Nosy_OptimismPortalFilterer_ParseTransactionDeposited__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal *OptimismPortalFilterer
		fill_err = tp.Fill(&_OptimismPortal)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _OptimismPortal == nil {
			return
		}

		_OptimismPortal.ParseTransactionDeposited(log)
	})
}

func Fuzz_Nosy_OptimismPortalFilterer_ParseWithdrawalFinalized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal *OptimismPortalFilterer
		fill_err = tp.Fill(&_OptimismPortal)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _OptimismPortal == nil {
			return
		}

		_OptimismPortal.ParseWithdrawalFinalized(log)
	})
}

func Fuzz_Nosy_OptimismPortalFilterer_ParseWithdrawalProven__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal *OptimismPortalFilterer
		fill_err = tp.Fill(&_OptimismPortal)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _OptimismPortal == nil {
			return
		}

		_OptimismPortal.ParseWithdrawalProven(log)
	})
}

// skipping Fuzz_Nosy_OptimismPortalFilterer_WatchInitialized__ because parameters include func, chan, or unsupported interface: chan<- *github.com/ethereum-optimism/optimism/op-node/bindings.OptimismPortalInitialized

// skipping Fuzz_Nosy_OptimismPortalFilterer_WatchTransactionDeposited__ because parameters include func, chan, or unsupported interface: chan<- *github.com/ethereum-optimism/optimism/op-node/bindings.OptimismPortalTransactionDeposited

// skipping Fuzz_Nosy_OptimismPortalFilterer_WatchWithdrawalFinalized__ because parameters include func, chan, or unsupported interface: chan<- *github.com/ethereum-optimism/optimism/op-node/bindings.OptimismPortalWithdrawalFinalized

// skipping Fuzz_Nosy_OptimismPortalFilterer_WatchWithdrawalProven__ because parameters include func, chan, or unsupported interface: chan<- *github.com/ethereum-optimism/optimism/op-node/bindings.OptimismPortalWithdrawalProven

func Fuzz_Nosy_OptimismPortalInitializedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OptimismPortalInitializedIterator
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

func Fuzz_Nosy_OptimismPortalInitializedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OptimismPortalInitializedIterator
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

func Fuzz_Nosy_OptimismPortalInitializedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OptimismPortalInitializedIterator
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

// skipping Fuzz_Nosy_OptimismPortalRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

// skipping Fuzz_Nosy_OptimismPortalRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_OptimismPortalRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal *OptimismPortalRaw
		fill_err = tp.Fill(&_OptimismPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OptimismPortal == nil || opts == nil {
			return
		}

		_OptimismPortal.Transfer(opts)
	})
}

func Fuzz_Nosy_OptimismPortalSession_DepositTransaction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal *OptimismPortalSession
		fill_err = tp.Fill(&_OptimismPortal)
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
		if _OptimismPortal == nil || _value == nil {
			return
		}

		_OptimismPortal.DepositTransaction(_to, _value, _gasLimit, _isCreation, _data)
	})
}

func Fuzz_Nosy_OptimismPortalSession_DonateETH__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal *OptimismPortalSession
		fill_err = tp.Fill(&_OptimismPortal)
		if fill_err != nil {
			return
		}
		if _OptimismPortal == nil {
			return
		}

		_OptimismPortal.DonateETH()
	})
}

func Fuzz_Nosy_OptimismPortalSession_FinalizeWithdrawalTransaction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal *OptimismPortalSession
		fill_err = tp.Fill(&_OptimismPortal)
		if fill_err != nil {
			return
		}
		var _tx TypesWithdrawalTransaction
		fill_err = tp.Fill(&_tx)
		if fill_err != nil {
			return
		}
		if _OptimismPortal == nil {
			return
		}

		_OptimismPortal.FinalizeWithdrawalTransaction(_tx)
	})
}

func Fuzz_Nosy_OptimismPortalSession_FinalizedWithdrawals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal *OptimismPortalSession
		fill_err = tp.Fill(&_OptimismPortal)
		if fill_err != nil {
			return
		}
		var arg0 [32]byte
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _OptimismPortal == nil {
			return
		}

		_OptimismPortal.FinalizedWithdrawals(arg0)
	})
}

func Fuzz_Nosy_OptimismPortalSession_Guardian__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal *OptimismPortalSession
		fill_err = tp.Fill(&_OptimismPortal)
		if fill_err != nil {
			return
		}
		if _OptimismPortal == nil {
			return
		}

		_OptimismPortal.Guardian()
	})
}

func Fuzz_Nosy_OptimismPortalSession_Initialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal *OptimismPortalSession
		fill_err = tp.Fill(&_OptimismPortal)
		if fill_err != nil {
			return
		}
		var _l2Oracle common.Address
		fill_err = tp.Fill(&_l2Oracle)
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
		if _OptimismPortal == nil {
			return
		}

		_OptimismPortal.Initialize(_l2Oracle, _systemConfig, _superchainConfig)
	})
}

func Fuzz_Nosy_OptimismPortalSession_IsOutputFinalized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal *OptimismPortalSession
		fill_err = tp.Fill(&_OptimismPortal)
		if fill_err != nil {
			return
		}
		var _l2OutputIndex *big.Int
		fill_err = tp.Fill(&_l2OutputIndex)
		if fill_err != nil {
			return
		}
		if _OptimismPortal == nil || _l2OutputIndex == nil {
			return
		}

		_OptimismPortal.IsOutputFinalized(_l2OutputIndex)
	})
}

func Fuzz_Nosy_OptimismPortalSession_L2Oracle__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal *OptimismPortalSession
		fill_err = tp.Fill(&_OptimismPortal)
		if fill_err != nil {
			return
		}
		if _OptimismPortal == nil {
			return
		}

		_OptimismPortal.L2Oracle()
	})
}

func Fuzz_Nosy_OptimismPortalSession_L2Sender__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal *OptimismPortalSession
		fill_err = tp.Fill(&_OptimismPortal)
		if fill_err != nil {
			return
		}
		if _OptimismPortal == nil {
			return
		}

		_OptimismPortal.L2Sender()
	})
}

func Fuzz_Nosy_OptimismPortalSession_MinimumGasLimit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal *OptimismPortalSession
		fill_err = tp.Fill(&_OptimismPortal)
		if fill_err != nil {
			return
		}
		var _byteCount uint64
		fill_err = tp.Fill(&_byteCount)
		if fill_err != nil {
			return
		}
		if _OptimismPortal == nil {
			return
		}

		_OptimismPortal.MinimumGasLimit(_byteCount)
	})
}

func Fuzz_Nosy_OptimismPortalSession_Params__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal *OptimismPortalSession
		fill_err = tp.Fill(&_OptimismPortal)
		if fill_err != nil {
			return
		}
		if _OptimismPortal == nil {
			return
		}

		_OptimismPortal.Params()
	})
}

func Fuzz_Nosy_OptimismPortalSession_Paused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal *OptimismPortalSession
		fill_err = tp.Fill(&_OptimismPortal)
		if fill_err != nil {
			return
		}
		if _OptimismPortal == nil {
			return
		}

		_OptimismPortal.Paused()
	})
}

func Fuzz_Nosy_OptimismPortalSession_ProveWithdrawalTransaction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal *OptimismPortalSession
		fill_err = tp.Fill(&_OptimismPortal)
		if fill_err != nil {
			return
		}
		var _tx TypesWithdrawalTransaction
		fill_err = tp.Fill(&_tx)
		if fill_err != nil {
			return
		}
		var _l2OutputIndex *big.Int
		fill_err = tp.Fill(&_l2OutputIndex)
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
		if _OptimismPortal == nil || _l2OutputIndex == nil {
			return
		}

		_OptimismPortal.ProveWithdrawalTransaction(_tx, _l2OutputIndex, _outputRootProof, _withdrawalProof)
	})
}

func Fuzz_Nosy_OptimismPortalSession_ProvenWithdrawals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal *OptimismPortalSession
		fill_err = tp.Fill(&_OptimismPortal)
		if fill_err != nil {
			return
		}
		var arg0 [32]byte
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _OptimismPortal == nil {
			return
		}

		_OptimismPortal.ProvenWithdrawals(arg0)
	})
}

func Fuzz_Nosy_OptimismPortalSession_Receive__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal *OptimismPortalSession
		fill_err = tp.Fill(&_OptimismPortal)
		if fill_err != nil {
			return
		}
		if _OptimismPortal == nil {
			return
		}

		_OptimismPortal.Receive()
	})
}

func Fuzz_Nosy_OptimismPortalSession_SuperchainConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal *OptimismPortalSession
		fill_err = tp.Fill(&_OptimismPortal)
		if fill_err != nil {
			return
		}
		if _OptimismPortal == nil {
			return
		}

		_OptimismPortal.SuperchainConfig()
	})
}

func Fuzz_Nosy_OptimismPortalSession_SystemConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal *OptimismPortalSession
		fill_err = tp.Fill(&_OptimismPortal)
		if fill_err != nil {
			return
		}
		if _OptimismPortal == nil {
			return
		}

		_OptimismPortal.SystemConfig()
	})
}

func Fuzz_Nosy_OptimismPortalSession_Version__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal *OptimismPortalSession
		fill_err = tp.Fill(&_OptimismPortal)
		if fill_err != nil {
			return
		}
		if _OptimismPortal == nil {
			return
		}

		_OptimismPortal.Version()
	})
}

func Fuzz_Nosy_OptimismPortalTransactionDepositedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OptimismPortalTransactionDepositedIterator
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

func Fuzz_Nosy_OptimismPortalTransactionDepositedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OptimismPortalTransactionDepositedIterator
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

func Fuzz_Nosy_OptimismPortalTransactionDepositedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OptimismPortalTransactionDepositedIterator
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

func Fuzz_Nosy_OptimismPortalTransactor_DepositTransaction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal *OptimismPortalTransactor
		fill_err = tp.Fill(&_OptimismPortal)
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
		if _OptimismPortal == nil || opts == nil || _value == nil {
			return
		}

		_OptimismPortal.DepositTransaction(opts, _to, _value, _gasLimit, _isCreation, _data)
	})
}

func Fuzz_Nosy_OptimismPortalTransactor_DonateETH__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal *OptimismPortalTransactor
		fill_err = tp.Fill(&_OptimismPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OptimismPortal == nil || opts == nil {
			return
		}

		_OptimismPortal.DonateETH(opts)
	})
}

func Fuzz_Nosy_OptimismPortalTransactor_FinalizeWithdrawalTransaction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal *OptimismPortalTransactor
		fill_err = tp.Fill(&_OptimismPortal)
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
		if _OptimismPortal == nil || opts == nil {
			return
		}

		_OptimismPortal.FinalizeWithdrawalTransaction(opts, _tx)
	})
}

func Fuzz_Nosy_OptimismPortalTransactor_Initialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal *OptimismPortalTransactor
		fill_err = tp.Fill(&_OptimismPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var _l2Oracle common.Address
		fill_err = tp.Fill(&_l2Oracle)
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
		if _OptimismPortal == nil || opts == nil {
			return
		}

		_OptimismPortal.Initialize(opts, _l2Oracle, _systemConfig, _superchainConfig)
	})
}

func Fuzz_Nosy_OptimismPortalTransactor_ProveWithdrawalTransaction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal *OptimismPortalTransactor
		fill_err = tp.Fill(&_OptimismPortal)
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
		var _l2OutputIndex *big.Int
		fill_err = tp.Fill(&_l2OutputIndex)
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
		if _OptimismPortal == nil || opts == nil || _l2OutputIndex == nil {
			return
		}

		_OptimismPortal.ProveWithdrawalTransaction(opts, _tx, _l2OutputIndex, _outputRootProof, _withdrawalProof)
	})
}

func Fuzz_Nosy_OptimismPortalTransactor_Receive__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal *OptimismPortalTransactor
		fill_err = tp.Fill(&_OptimismPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OptimismPortal == nil || opts == nil {
			return
		}

		_OptimismPortal.Receive(opts)
	})
}

// skipping Fuzz_Nosy_OptimismPortalTransactorRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_OptimismPortalTransactorRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal *OptimismPortalTransactorRaw
		fill_err = tp.Fill(&_OptimismPortal)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _OptimismPortal == nil || opts == nil {
			return
		}

		_OptimismPortal.Transfer(opts)
	})
}

func Fuzz_Nosy_OptimismPortalTransactorSession_DepositTransaction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal *OptimismPortalTransactorSession
		fill_err = tp.Fill(&_OptimismPortal)
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
		if _OptimismPortal == nil || _value == nil {
			return
		}

		_OptimismPortal.DepositTransaction(_to, _value, _gasLimit, _isCreation, _data)
	})
}

func Fuzz_Nosy_OptimismPortalTransactorSession_DonateETH__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal *OptimismPortalTransactorSession
		fill_err = tp.Fill(&_OptimismPortal)
		if fill_err != nil {
			return
		}
		if _OptimismPortal == nil {
			return
		}

		_OptimismPortal.DonateETH()
	})
}

func Fuzz_Nosy_OptimismPortalTransactorSession_FinalizeWithdrawalTransaction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal *OptimismPortalTransactorSession
		fill_err = tp.Fill(&_OptimismPortal)
		if fill_err != nil {
			return
		}
		var _tx TypesWithdrawalTransaction
		fill_err = tp.Fill(&_tx)
		if fill_err != nil {
			return
		}
		if _OptimismPortal == nil {
			return
		}

		_OptimismPortal.FinalizeWithdrawalTransaction(_tx)
	})
}

func Fuzz_Nosy_OptimismPortalTransactorSession_Initialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal *OptimismPortalTransactorSession
		fill_err = tp.Fill(&_OptimismPortal)
		if fill_err != nil {
			return
		}
		var _l2Oracle common.Address
		fill_err = tp.Fill(&_l2Oracle)
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
		if _OptimismPortal == nil {
			return
		}

		_OptimismPortal.Initialize(_l2Oracle, _systemConfig, _superchainConfig)
	})
}

func Fuzz_Nosy_OptimismPortalTransactorSession_ProveWithdrawalTransaction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal *OptimismPortalTransactorSession
		fill_err = tp.Fill(&_OptimismPortal)
		if fill_err != nil {
			return
		}
		var _tx TypesWithdrawalTransaction
		fill_err = tp.Fill(&_tx)
		if fill_err != nil {
			return
		}
		var _l2OutputIndex *big.Int
		fill_err = tp.Fill(&_l2OutputIndex)
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
		if _OptimismPortal == nil || _l2OutputIndex == nil {
			return
		}

		_OptimismPortal.ProveWithdrawalTransaction(_tx, _l2OutputIndex, _outputRootProof, _withdrawalProof)
	})
}

func Fuzz_Nosy_OptimismPortalTransactorSession_Receive__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _OptimismPortal *OptimismPortalTransactorSession
		fill_err = tp.Fill(&_OptimismPortal)
		if fill_err != nil {
			return
		}
		if _OptimismPortal == nil {
			return
		}

		_OptimismPortal.Receive()
	})
}

func Fuzz_Nosy_OptimismPortalWithdrawalFinalizedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OptimismPortalWithdrawalFinalizedIterator
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

func Fuzz_Nosy_OptimismPortalWithdrawalFinalizedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OptimismPortalWithdrawalFinalizedIterator
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

func Fuzz_Nosy_OptimismPortalWithdrawalFinalizedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OptimismPortalWithdrawalFinalizedIterator
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

func Fuzz_Nosy_OptimismPortalWithdrawalProvenIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OptimismPortalWithdrawalProvenIterator
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

func Fuzz_Nosy_OptimismPortalWithdrawalProvenIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OptimismPortalWithdrawalProvenIterator
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

func Fuzz_Nosy_OptimismPortalWithdrawalProvenIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *OptimismPortalWithdrawalProvenIterator
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

// skipping Fuzz_Nosy_DeployDisputeGameFactory__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/abi/bind.ContractBackend

// skipping Fuzz_Nosy_DeployL1Block__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/abi/bind.ContractBackend

// skipping Fuzz_Nosy_DeployL2OutputOracle__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/abi/bind.ContractBackend

// skipping Fuzz_Nosy_DeployL2ToL1MessagePasser__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/abi/bind.ContractBackend

// skipping Fuzz_Nosy_DeployOptimismPortal__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/abi/bind.ContractBackend
