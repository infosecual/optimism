package emit

import (
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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

// skipping Fuzz_Nosy_EmitCallerRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

func Fuzz_Nosy_EmitDataEmittedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *EmitDataEmittedIterator
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

func Fuzz_Nosy_EmitDataEmittedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *EmitDataEmittedIterator
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

func Fuzz_Nosy_EmitDataEmittedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *EmitDataEmittedIterator
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

func Fuzz_Nosy_EmitFilterer_FilterDataEmitted__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Emit *EmitFilterer
		fill_err = tp.Fill(&_Emit)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var _data [][]byte
		fill_err = tp.Fill(&_data)
		if fill_err != nil {
			return
		}
		if _Emit == nil || opts == nil {
			return
		}

		_Emit.FilterDataEmitted(opts, _data)
	})
}

func Fuzz_Nosy_EmitFilterer_ParseDataEmitted__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Emit *EmitFilterer
		fill_err = tp.Fill(&_Emit)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _Emit == nil {
			return
		}

		_Emit.ParseDataEmitted(log)
	})
}

// skipping Fuzz_Nosy_EmitFilterer_WatchDataEmitted__ because parameters include func, chan, or unsupported interface: chan<- *github.com/ethereum-optimism/optimism/op-e2e/e2eutils/interop/contracts/bindings/emit.EmitDataEmitted

// skipping Fuzz_Nosy_EmitRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

// skipping Fuzz_Nosy_EmitRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_EmitRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Emit *EmitRaw
		fill_err = tp.Fill(&_Emit)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _Emit == nil || opts == nil {
			return
		}

		_Emit.Transfer(opts)
	})
}

func Fuzz_Nosy_EmitSession_EmitData__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Emit *EmitSession
		fill_err = tp.Fill(&_Emit)
		if fill_err != nil {
			return
		}
		var _data []byte
		fill_err = tp.Fill(&_data)
		if fill_err != nil {
			return
		}
		if _Emit == nil {
			return
		}

		_Emit.EmitData(_data)
	})
}

func Fuzz_Nosy_EmitTransactor_EmitData__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Emit *EmitTransactor
		fill_err = tp.Fill(&_Emit)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var _data []byte
		fill_err = tp.Fill(&_data)
		if fill_err != nil {
			return
		}
		if _Emit == nil || opts == nil {
			return
		}

		_Emit.EmitData(opts, _data)
	})
}

// skipping Fuzz_Nosy_EmitTransactorRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_EmitTransactorRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Emit *EmitTransactorRaw
		fill_err = tp.Fill(&_Emit)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _Emit == nil || opts == nil {
			return
		}

		_Emit.Transfer(opts)
	})
}

func Fuzz_Nosy_EmitTransactorSession_EmitData__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Emit *EmitTransactorSession
		fill_err = tp.Fill(&_Emit)
		if fill_err != nil {
			return
		}
		var _data []byte
		fill_err = tp.Fill(&_data)
		if fill_err != nil {
			return
		}
		if _Emit == nil {
			return
		}

		_Emit.EmitData(_data)
	})
}

// skipping Fuzz_Nosy_DeployEmit__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/abi/bind.ContractBackend
