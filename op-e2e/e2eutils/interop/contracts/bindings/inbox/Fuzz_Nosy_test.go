package inbox

import (
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
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

// skipping Fuzz_Nosy_InboxCallerRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

// skipping Fuzz_Nosy_InboxRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

// skipping Fuzz_Nosy_InboxRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_InboxRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Inbox *InboxRaw
		fill_err = tp.Fill(&_Inbox)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _Inbox == nil || opts == nil {
			return
		}

		_Inbox.Transfer(opts)
	})
}

func Fuzz_Nosy_InboxSession_ExecuteMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Inbox *InboxSession
		fill_err = tp.Fill(&_Inbox)
		if fill_err != nil {
			return
		}
		var _id Identifier
		fill_err = tp.Fill(&_id)
		if fill_err != nil {
			return
		}
		var _target common.Address
		fill_err = tp.Fill(&_target)
		if fill_err != nil {
			return
		}
		var _message []byte
		fill_err = tp.Fill(&_message)
		if fill_err != nil {
			return
		}
		if _Inbox == nil {
			return
		}

		_Inbox.ExecuteMessage(_id, _target, _message)
	})
}

func Fuzz_Nosy_InboxSession_ValidateMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Inbox *InboxSession
		fill_err = tp.Fill(&_Inbox)
		if fill_err != nil {
			return
		}
		var _id Identifier
		fill_err = tp.Fill(&_id)
		if fill_err != nil {
			return
		}
		var _msgHash [32]byte
		fill_err = tp.Fill(&_msgHash)
		if fill_err != nil {
			return
		}
		if _Inbox == nil {
			return
		}

		_Inbox.ValidateMessage(_id, _msgHash)
	})
}

func Fuzz_Nosy_InboxTransactor_ExecuteMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Inbox *InboxTransactor
		fill_err = tp.Fill(&_Inbox)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var _id Identifier
		fill_err = tp.Fill(&_id)
		if fill_err != nil {
			return
		}
		var _target common.Address
		fill_err = tp.Fill(&_target)
		if fill_err != nil {
			return
		}
		var _message []byte
		fill_err = tp.Fill(&_message)
		if fill_err != nil {
			return
		}
		if _Inbox == nil || opts == nil {
			return
		}

		_Inbox.ExecuteMessage(opts, _id, _target, _message)
	})
}

func Fuzz_Nosy_InboxTransactor_ValidateMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Inbox *InboxTransactor
		fill_err = tp.Fill(&_Inbox)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var _id Identifier
		fill_err = tp.Fill(&_id)
		if fill_err != nil {
			return
		}
		var _msgHash [32]byte
		fill_err = tp.Fill(&_msgHash)
		if fill_err != nil {
			return
		}
		if _Inbox == nil || opts == nil {
			return
		}

		_Inbox.ValidateMessage(opts, _id, _msgHash)
	})
}

// skipping Fuzz_Nosy_InboxTransactorRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_InboxTransactorRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Inbox *InboxTransactorRaw
		fill_err = tp.Fill(&_Inbox)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _Inbox == nil || opts == nil {
			return
		}

		_Inbox.Transfer(opts)
	})
}

func Fuzz_Nosy_InboxTransactorSession_ExecuteMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Inbox *InboxTransactorSession
		fill_err = tp.Fill(&_Inbox)
		if fill_err != nil {
			return
		}
		var _id Identifier
		fill_err = tp.Fill(&_id)
		if fill_err != nil {
			return
		}
		var _target common.Address
		fill_err = tp.Fill(&_target)
		if fill_err != nil {
			return
		}
		var _message []byte
		fill_err = tp.Fill(&_message)
		if fill_err != nil {
			return
		}
		if _Inbox == nil {
			return
		}

		_Inbox.ExecuteMessage(_id, _target, _message)
	})
}

func Fuzz_Nosy_InboxTransactorSession_ValidateMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _Inbox *InboxTransactorSession
		fill_err = tp.Fill(&_Inbox)
		if fill_err != nil {
			return
		}
		var _id Identifier
		fill_err = tp.Fill(&_id)
		if fill_err != nil {
			return
		}
		var _msgHash [32]byte
		fill_err = tp.Fill(&_msgHash)
		if fill_err != nil {
			return
		}
		if _Inbox == nil {
			return
		}

		_Inbox.ValidateMessage(_id, _msgHash)
	})
}
