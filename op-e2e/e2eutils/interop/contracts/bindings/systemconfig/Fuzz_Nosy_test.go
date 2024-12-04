package systemconfig

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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

// skipping Fuzz_Nosy_SystemconfigCallerRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

// skipping Fuzz_Nosy_SystemconfigRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

// skipping Fuzz_Nosy_SystemconfigRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_SystemconfigRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Systemconfig *SystemconfigRaw
		fill_err = tp.Fill(&_Systemconfig)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _Systemconfig == nil || opts == nil {
			return
		}

		_Systemconfig.Transfer(opts)
	})
}

func Fuzz_Nosy_SystemconfigSession_AddDependency__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Systemconfig *SystemconfigSession
		fill_err = tp.Fill(&_Systemconfig)
		if fill_err != nil {
			return
		}
		var _chainId *big.Int
		fill_err = tp.Fill(&_chainId)
		if fill_err != nil {
			return
		}
		if _Systemconfig == nil || _chainId == nil {
			return
		}

		_Systemconfig.AddDependency(_chainId)
	})
}

func Fuzz_Nosy_SystemconfigTransactor_AddDependency__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Systemconfig *SystemconfigTransactor
		fill_err = tp.Fill(&_Systemconfig)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var _chainId *big.Int
		fill_err = tp.Fill(&_chainId)
		if fill_err != nil {
			return
		}
		if _Systemconfig == nil || opts == nil || _chainId == nil {
			return
		}

		_Systemconfig.AddDependency(opts, _chainId)
	})
}

// skipping Fuzz_Nosy_SystemconfigTransactorRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_SystemconfigTransactorRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Systemconfig *SystemconfigTransactorRaw
		fill_err = tp.Fill(&_Systemconfig)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _Systemconfig == nil || opts == nil {
			return
		}

		_Systemconfig.Transfer(opts)
	})
}

func Fuzz_Nosy_SystemconfigTransactorSession_AddDependency__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Systemconfig *SystemconfigTransactorSession
		fill_err = tp.Fill(&_Systemconfig)
		if fill_err != nil {
			return
		}
		var _chainId *big.Int
		fill_err = tp.Fill(&_chainId)
		if fill_err != nil {
			return
		}
		if _Systemconfig == nil || _chainId == nil {
			return
		}

		_Systemconfig.AddDependency(_chainId)
	})
}
