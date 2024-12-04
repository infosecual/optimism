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

func Fuzz_Nosy_L1CrossDomainMessengerCaller_BaseGas__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerCaller
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var _message []byte
		fill_err = tp.Fill(&_message)
		if fill_err != nil {
			return
		}
		var _minGasLimit uint32
		fill_err = tp.Fill(&_minGasLimit)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil || opts == nil {
			return
		}

		_L1CrossDomainMessenger.BaseGas(opts, _message, _minGasLimit)
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerCaller_FailedMessages__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerCaller
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
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
		if _L1CrossDomainMessenger == nil || opts == nil {
			return
		}

		_L1CrossDomainMessenger.FailedMessages(opts, arg0)
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerCaller_MESSAGEVERSION__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerCaller
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil || opts == nil {
			return
		}

		_L1CrossDomainMessenger.MESSAGEVERSION(opts)
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerCaller_MINGASCALLDATAOVERHEAD__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerCaller
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil || opts == nil {
			return
		}

		_L1CrossDomainMessenger.MINGASCALLDATAOVERHEAD(opts)
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerCaller_MINGASDYNAMICOVERHEADDENOMINATOR__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerCaller
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil || opts == nil {
			return
		}

		_L1CrossDomainMessenger.MINGASDYNAMICOVERHEADDENOMINATOR(opts)
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerCaller_MINGASDYNAMICOVERHEADNUMERATOR__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerCaller
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil || opts == nil {
			return
		}

		_L1CrossDomainMessenger.MINGASDYNAMICOVERHEADNUMERATOR(opts)
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerCaller_MessageNonce__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerCaller
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil || opts == nil {
			return
		}

		_L1CrossDomainMessenger.MessageNonce(opts)
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerCaller_OTHERMESSENGER__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerCaller
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil || opts == nil {
			return
		}

		_L1CrossDomainMessenger.OTHERMESSENGER(opts)
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerCaller_OtherMessenger__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerCaller
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil || opts == nil {
			return
		}

		_L1CrossDomainMessenger.OtherMessenger(opts)
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerCaller_PORTAL__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerCaller
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil || opts == nil {
			return
		}

		_L1CrossDomainMessenger.PORTAL(opts)
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerCaller_Paused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerCaller
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil || opts == nil {
			return
		}

		_L1CrossDomainMessenger.Paused(opts)
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerCaller_Portal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerCaller
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil || opts == nil {
			return
		}

		_L1CrossDomainMessenger.Portal(opts)
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerCaller_RELAYCALLOVERHEAD__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerCaller
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil || opts == nil {
			return
		}

		_L1CrossDomainMessenger.RELAYCALLOVERHEAD(opts)
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerCaller_RELAYCONSTANTOVERHEAD__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerCaller
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil || opts == nil {
			return
		}

		_L1CrossDomainMessenger.RELAYCONSTANTOVERHEAD(opts)
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerCaller_RELAYGASCHECKBUFFER__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerCaller
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil || opts == nil {
			return
		}

		_L1CrossDomainMessenger.RELAYGASCHECKBUFFER(opts)
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerCaller_RELAYRESERVEDGAS__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerCaller
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil || opts == nil {
			return
		}

		_L1CrossDomainMessenger.RELAYRESERVEDGAS(opts)
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerCaller_SuccessfulMessages__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerCaller
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
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
		if _L1CrossDomainMessenger == nil || opts == nil {
			return
		}

		_L1CrossDomainMessenger.SuccessfulMessages(opts, arg0)
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerCaller_SuperchainConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerCaller
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil || opts == nil {
			return
		}

		_L1CrossDomainMessenger.SuperchainConfig(opts)
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerCaller_SystemConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerCaller
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil || opts == nil {
			return
		}

		_L1CrossDomainMessenger.SystemConfig(opts)
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerCaller_Version__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerCaller
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil || opts == nil {
			return
		}

		_L1CrossDomainMessenger.Version(opts)
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerCaller_XDomainMessageSender__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerCaller
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil || opts == nil {
			return
		}

		_L1CrossDomainMessenger.XDomainMessageSender(opts)
	})
}

// skipping Fuzz_Nosy_L1CrossDomainMessengerCallerRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

func Fuzz_Nosy_L1CrossDomainMessengerCallerSession_BaseGas__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerCallerSession
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var _message []byte
		fill_err = tp.Fill(&_message)
		if fill_err != nil {
			return
		}
		var _minGasLimit uint32
		fill_err = tp.Fill(&_minGasLimit)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil {
			return
		}

		_L1CrossDomainMessenger.BaseGas(_message, _minGasLimit)
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerCallerSession_FailedMessages__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerCallerSession
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var arg0 [32]byte
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil {
			return
		}

		_L1CrossDomainMessenger.FailedMessages(arg0)
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerCallerSession_MESSAGEVERSION__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerCallerSession
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil {
			return
		}

		_L1CrossDomainMessenger.MESSAGEVERSION()
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerCallerSession_MINGASCALLDATAOVERHEAD__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerCallerSession
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil {
			return
		}

		_L1CrossDomainMessenger.MINGASCALLDATAOVERHEAD()
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerCallerSession_MINGASDYNAMICOVERHEADDENOMINATOR__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerCallerSession
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil {
			return
		}

		_L1CrossDomainMessenger.MINGASDYNAMICOVERHEADDENOMINATOR()
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerCallerSession_MINGASDYNAMICOVERHEADNUMERATOR__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerCallerSession
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil {
			return
		}

		_L1CrossDomainMessenger.MINGASDYNAMICOVERHEADNUMERATOR()
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerCallerSession_MessageNonce__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerCallerSession
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil {
			return
		}

		_L1CrossDomainMessenger.MessageNonce()
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerCallerSession_OTHERMESSENGER__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerCallerSession
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil {
			return
		}

		_L1CrossDomainMessenger.OTHERMESSENGER()
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerCallerSession_OtherMessenger__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerCallerSession
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil {
			return
		}

		_L1CrossDomainMessenger.OtherMessenger()
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerCallerSession_PORTAL__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerCallerSession
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil {
			return
		}

		_L1CrossDomainMessenger.PORTAL()
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerCallerSession_Paused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerCallerSession
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil {
			return
		}

		_L1CrossDomainMessenger.Paused()
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerCallerSession_Portal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerCallerSession
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil {
			return
		}

		_L1CrossDomainMessenger.Portal()
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerCallerSession_RELAYCALLOVERHEAD__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerCallerSession
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil {
			return
		}

		_L1CrossDomainMessenger.RELAYCALLOVERHEAD()
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerCallerSession_RELAYCONSTANTOVERHEAD__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerCallerSession
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil {
			return
		}

		_L1CrossDomainMessenger.RELAYCONSTANTOVERHEAD()
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerCallerSession_RELAYGASCHECKBUFFER__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerCallerSession
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil {
			return
		}

		_L1CrossDomainMessenger.RELAYGASCHECKBUFFER()
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerCallerSession_RELAYRESERVEDGAS__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerCallerSession
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil {
			return
		}

		_L1CrossDomainMessenger.RELAYRESERVEDGAS()
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerCallerSession_SuccessfulMessages__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerCallerSession
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var arg0 [32]byte
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil {
			return
		}

		_L1CrossDomainMessenger.SuccessfulMessages(arg0)
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerCallerSession_SuperchainConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerCallerSession
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil {
			return
		}

		_L1CrossDomainMessenger.SuperchainConfig()
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerCallerSession_SystemConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerCallerSession
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil {
			return
		}

		_L1CrossDomainMessenger.SystemConfig()
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerCallerSession_Version__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerCallerSession
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil {
			return
		}

		_L1CrossDomainMessenger.Version()
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerCallerSession_XDomainMessageSender__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerCallerSession
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil {
			return
		}

		_L1CrossDomainMessenger.XDomainMessageSender()
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerFailedRelayedMessageIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *L1CrossDomainMessengerFailedRelayedMessageIterator
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

func Fuzz_Nosy_L1CrossDomainMessengerFailedRelayedMessageIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *L1CrossDomainMessengerFailedRelayedMessageIterator
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

func Fuzz_Nosy_L1CrossDomainMessengerFailedRelayedMessageIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *L1CrossDomainMessengerFailedRelayedMessageIterator
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

func Fuzz_Nosy_L1CrossDomainMessengerFilterer_FilterFailedRelayedMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerFilterer
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var msgHash [][32]byte
		fill_err = tp.Fill(&msgHash)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil || opts == nil {
			return
		}

		_L1CrossDomainMessenger.FilterFailedRelayedMessage(opts, msgHash)
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerFilterer_FilterInitialized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerFilterer
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil || opts == nil {
			return
		}

		_L1CrossDomainMessenger.FilterInitialized(opts)
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerFilterer_FilterRelayedMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerFilterer
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var msgHash [][32]byte
		fill_err = tp.Fill(&msgHash)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil || opts == nil {
			return
		}

		_L1CrossDomainMessenger.FilterRelayedMessage(opts, msgHash)
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerFilterer_FilterSentMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerFilterer
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var t3 []common.Address
		fill_err = tp.Fill(&t3)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil || opts == nil {
			return
		}

		_L1CrossDomainMessenger.FilterSentMessage(opts, t3)
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerFilterer_FilterSentMessageExtension1__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerFilterer
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var sender []common.Address
		fill_err = tp.Fill(&sender)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil || opts == nil {
			return
		}

		_L1CrossDomainMessenger.FilterSentMessageExtension1(opts, sender)
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerFilterer_ParseFailedRelayedMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerFilterer
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil {
			return
		}

		_L1CrossDomainMessenger.ParseFailedRelayedMessage(log)
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerFilterer_ParseInitialized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerFilterer
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil {
			return
		}

		_L1CrossDomainMessenger.ParseInitialized(log)
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerFilterer_ParseRelayedMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerFilterer
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil {
			return
		}

		_L1CrossDomainMessenger.ParseRelayedMessage(log)
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerFilterer_ParseSentMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerFilterer
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil {
			return
		}

		_L1CrossDomainMessenger.ParseSentMessage(log)
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerFilterer_ParseSentMessageExtension1__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerFilterer
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil {
			return
		}

		_L1CrossDomainMessenger.ParseSentMessageExtension1(log)
	})
}

// skipping Fuzz_Nosy_L1CrossDomainMessengerFilterer_WatchFailedRelayedMessage__ because parameters include func, chan, or unsupported interface: chan<- *github.com/ethereum-optimism/optimism/op-chain-ops/crossdomain/bindings.L1CrossDomainMessengerFailedRelayedMessage

// skipping Fuzz_Nosy_L1CrossDomainMessengerFilterer_WatchInitialized__ because parameters include func, chan, or unsupported interface: chan<- *github.com/ethereum-optimism/optimism/op-chain-ops/crossdomain/bindings.L1CrossDomainMessengerInitialized

// skipping Fuzz_Nosy_L1CrossDomainMessengerFilterer_WatchRelayedMessage__ because parameters include func, chan, or unsupported interface: chan<- *github.com/ethereum-optimism/optimism/op-chain-ops/crossdomain/bindings.L1CrossDomainMessengerRelayedMessage

// skipping Fuzz_Nosy_L1CrossDomainMessengerFilterer_WatchSentMessage__ because parameters include func, chan, or unsupported interface: chan<- *github.com/ethereum-optimism/optimism/op-chain-ops/crossdomain/bindings.L1CrossDomainMessengerSentMessage

// skipping Fuzz_Nosy_L1CrossDomainMessengerFilterer_WatchSentMessageExtension1__ because parameters include func, chan, or unsupported interface: chan<- *github.com/ethereum-optimism/optimism/op-chain-ops/crossdomain/bindings.L1CrossDomainMessengerSentMessageExtension1

func Fuzz_Nosy_L1CrossDomainMessengerInitializedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *L1CrossDomainMessengerInitializedIterator
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

func Fuzz_Nosy_L1CrossDomainMessengerInitializedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *L1CrossDomainMessengerInitializedIterator
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

func Fuzz_Nosy_L1CrossDomainMessengerInitializedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *L1CrossDomainMessengerInitializedIterator
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

// skipping Fuzz_Nosy_L1CrossDomainMessengerRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

// skipping Fuzz_Nosy_L1CrossDomainMessengerRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_L1CrossDomainMessengerRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerRaw
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil || opts == nil {
			return
		}

		_L1CrossDomainMessenger.Transfer(opts)
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerRelayedMessageIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *L1CrossDomainMessengerRelayedMessageIterator
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

func Fuzz_Nosy_L1CrossDomainMessengerRelayedMessageIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *L1CrossDomainMessengerRelayedMessageIterator
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

func Fuzz_Nosy_L1CrossDomainMessengerRelayedMessageIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *L1CrossDomainMessengerRelayedMessageIterator
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

func Fuzz_Nosy_L1CrossDomainMessengerSentMessageExtension1Iterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *L1CrossDomainMessengerSentMessageExtension1Iterator
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

func Fuzz_Nosy_L1CrossDomainMessengerSentMessageExtension1Iterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *L1CrossDomainMessengerSentMessageExtension1Iterator
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

func Fuzz_Nosy_L1CrossDomainMessengerSentMessageExtension1Iterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *L1CrossDomainMessengerSentMessageExtension1Iterator
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

func Fuzz_Nosy_L1CrossDomainMessengerSentMessageIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *L1CrossDomainMessengerSentMessageIterator
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

func Fuzz_Nosy_L1CrossDomainMessengerSentMessageIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *L1CrossDomainMessengerSentMessageIterator
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

func Fuzz_Nosy_L1CrossDomainMessengerSentMessageIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *L1CrossDomainMessengerSentMessageIterator
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

func Fuzz_Nosy_L1CrossDomainMessengerSession_BaseGas__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerSession
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var _message []byte
		fill_err = tp.Fill(&_message)
		if fill_err != nil {
			return
		}
		var _minGasLimit uint32
		fill_err = tp.Fill(&_minGasLimit)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil {
			return
		}

		_L1CrossDomainMessenger.BaseGas(_message, _minGasLimit)
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerSession_FailedMessages__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerSession
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var arg0 [32]byte
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil {
			return
		}

		_L1CrossDomainMessenger.FailedMessages(arg0)
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerSession_Initialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerSession
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var _superchainConfig common.Address
		fill_err = tp.Fill(&_superchainConfig)
		if fill_err != nil {
			return
		}
		var _portal common.Address
		fill_err = tp.Fill(&_portal)
		if fill_err != nil {
			return
		}
		var _systemConfig common.Address
		fill_err = tp.Fill(&_systemConfig)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil {
			return
		}

		_L1CrossDomainMessenger.Initialize(_superchainConfig, _portal, _systemConfig)
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerSession_MESSAGEVERSION__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerSession
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil {
			return
		}

		_L1CrossDomainMessenger.MESSAGEVERSION()
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerSession_MINGASCALLDATAOVERHEAD__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerSession
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil {
			return
		}

		_L1CrossDomainMessenger.MINGASCALLDATAOVERHEAD()
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerSession_MINGASDYNAMICOVERHEADDENOMINATOR__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerSession
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil {
			return
		}

		_L1CrossDomainMessenger.MINGASDYNAMICOVERHEADDENOMINATOR()
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerSession_MINGASDYNAMICOVERHEADNUMERATOR__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerSession
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil {
			return
		}

		_L1CrossDomainMessenger.MINGASDYNAMICOVERHEADNUMERATOR()
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerSession_MessageNonce__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerSession
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil {
			return
		}

		_L1CrossDomainMessenger.MessageNonce()
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerSession_OTHERMESSENGER__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerSession
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil {
			return
		}

		_L1CrossDomainMessenger.OTHERMESSENGER()
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerSession_OtherMessenger__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerSession
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil {
			return
		}

		_L1CrossDomainMessenger.OtherMessenger()
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerSession_PORTAL__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerSession
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil {
			return
		}

		_L1CrossDomainMessenger.PORTAL()
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerSession_Paused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerSession
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil {
			return
		}

		_L1CrossDomainMessenger.Paused()
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerSession_Portal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerSession
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil {
			return
		}

		_L1CrossDomainMessenger.Portal()
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerSession_RELAYCALLOVERHEAD__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerSession
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil {
			return
		}

		_L1CrossDomainMessenger.RELAYCALLOVERHEAD()
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerSession_RELAYCONSTANTOVERHEAD__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerSession
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil {
			return
		}

		_L1CrossDomainMessenger.RELAYCONSTANTOVERHEAD()
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerSession_RELAYGASCHECKBUFFER__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerSession
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil {
			return
		}

		_L1CrossDomainMessenger.RELAYGASCHECKBUFFER()
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerSession_RELAYRESERVEDGAS__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerSession
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil {
			return
		}

		_L1CrossDomainMessenger.RELAYRESERVEDGAS()
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerSession_RelayMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerSession
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var _nonce *big.Int
		fill_err = tp.Fill(&_nonce)
		if fill_err != nil {
			return
		}
		var _sender common.Address
		fill_err = tp.Fill(&_sender)
		if fill_err != nil {
			return
		}
		var _target common.Address
		fill_err = tp.Fill(&_target)
		if fill_err != nil {
			return
		}
		var _value *big.Int
		fill_err = tp.Fill(&_value)
		if fill_err != nil {
			return
		}
		var _minGasLimit *big.Int
		fill_err = tp.Fill(&_minGasLimit)
		if fill_err != nil {
			return
		}
		var _message []byte
		fill_err = tp.Fill(&_message)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil || _nonce == nil || _value == nil || _minGasLimit == nil {
			return
		}

		_L1CrossDomainMessenger.RelayMessage(_nonce, _sender, _target, _value, _minGasLimit, _message)
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerSession_SendMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerSession
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
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
		var _minGasLimit uint32
		fill_err = tp.Fill(&_minGasLimit)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil {
			return
		}

		_L1CrossDomainMessenger.SendMessage(_target, _message, _minGasLimit)
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerSession_SuccessfulMessages__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerSession
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var arg0 [32]byte
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil {
			return
		}

		_L1CrossDomainMessenger.SuccessfulMessages(arg0)
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerSession_SuperchainConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerSession
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil {
			return
		}

		_L1CrossDomainMessenger.SuperchainConfig()
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerSession_SystemConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerSession
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil {
			return
		}

		_L1CrossDomainMessenger.SystemConfig()
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerSession_Version__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerSession
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil {
			return
		}

		_L1CrossDomainMessenger.Version()
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerSession_XDomainMessageSender__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerSession
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil {
			return
		}

		_L1CrossDomainMessenger.XDomainMessageSender()
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerTransactor_Initialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerTransactor
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var _superchainConfig common.Address
		fill_err = tp.Fill(&_superchainConfig)
		if fill_err != nil {
			return
		}
		var _portal common.Address
		fill_err = tp.Fill(&_portal)
		if fill_err != nil {
			return
		}
		var _systemConfig common.Address
		fill_err = tp.Fill(&_systemConfig)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil || opts == nil {
			return
		}

		_L1CrossDomainMessenger.Initialize(opts, _superchainConfig, _portal, _systemConfig)
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerTransactor_RelayMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerTransactor
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var _nonce *big.Int
		fill_err = tp.Fill(&_nonce)
		if fill_err != nil {
			return
		}
		var _sender common.Address
		fill_err = tp.Fill(&_sender)
		if fill_err != nil {
			return
		}
		var _target common.Address
		fill_err = tp.Fill(&_target)
		if fill_err != nil {
			return
		}
		var _value *big.Int
		fill_err = tp.Fill(&_value)
		if fill_err != nil {
			return
		}
		var _minGasLimit *big.Int
		fill_err = tp.Fill(&_minGasLimit)
		if fill_err != nil {
			return
		}
		var _message []byte
		fill_err = tp.Fill(&_message)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil || opts == nil || _nonce == nil || _value == nil || _minGasLimit == nil {
			return
		}

		_L1CrossDomainMessenger.RelayMessage(opts, _nonce, _sender, _target, _value, _minGasLimit, _message)
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerTransactor_SendMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerTransactor
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
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
		var _message []byte
		fill_err = tp.Fill(&_message)
		if fill_err != nil {
			return
		}
		var _minGasLimit uint32
		fill_err = tp.Fill(&_minGasLimit)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil || opts == nil {
			return
		}

		_L1CrossDomainMessenger.SendMessage(opts, _target, _message, _minGasLimit)
	})
}

// skipping Fuzz_Nosy_L1CrossDomainMessengerTransactorRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_L1CrossDomainMessengerTransactorRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerTransactorRaw
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil || opts == nil {
			return
		}

		_L1CrossDomainMessenger.Transfer(opts)
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerTransactorSession_Initialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerTransactorSession
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var _superchainConfig common.Address
		fill_err = tp.Fill(&_superchainConfig)
		if fill_err != nil {
			return
		}
		var _portal common.Address
		fill_err = tp.Fill(&_portal)
		if fill_err != nil {
			return
		}
		var _systemConfig common.Address
		fill_err = tp.Fill(&_systemConfig)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil {
			return
		}

		_L1CrossDomainMessenger.Initialize(_superchainConfig, _portal, _systemConfig)
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerTransactorSession_RelayMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerTransactorSession
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var _nonce *big.Int
		fill_err = tp.Fill(&_nonce)
		if fill_err != nil {
			return
		}
		var _sender common.Address
		fill_err = tp.Fill(&_sender)
		if fill_err != nil {
			return
		}
		var _target common.Address
		fill_err = tp.Fill(&_target)
		if fill_err != nil {
			return
		}
		var _value *big.Int
		fill_err = tp.Fill(&_value)
		if fill_err != nil {
			return
		}
		var _minGasLimit *big.Int
		fill_err = tp.Fill(&_minGasLimit)
		if fill_err != nil {
			return
		}
		var _message []byte
		fill_err = tp.Fill(&_message)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil || _nonce == nil || _value == nil || _minGasLimit == nil {
			return
		}

		_L1CrossDomainMessenger.RelayMessage(_nonce, _sender, _target, _value, _minGasLimit, _message)
	})
}

func Fuzz_Nosy_L1CrossDomainMessengerTransactorSession_SendMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1CrossDomainMessenger *L1CrossDomainMessengerTransactorSession
		fill_err = tp.Fill(&_L1CrossDomainMessenger)
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
		var _minGasLimit uint32
		fill_err = tp.Fill(&_minGasLimit)
		if fill_err != nil {
			return
		}
		if _L1CrossDomainMessenger == nil {
			return
		}

		_L1CrossDomainMessenger.SendMessage(_target, _message, _minGasLimit)
	})
}

func Fuzz_Nosy_L1StandardBridgeCaller_Deposits__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeCaller
		fill_err = tp.Fill(&_L1StandardBridge)
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
		if _L1StandardBridge == nil || opts == nil {
			return
		}

		_L1StandardBridge.Deposits(opts, arg0, arg1)
	})
}

func Fuzz_Nosy_L1StandardBridgeCaller_L2TokenBridge__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeCaller
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil || opts == nil {
			return
		}

		_L1StandardBridge.L2TokenBridge(opts)
	})
}

func Fuzz_Nosy_L1StandardBridgeCaller_MESSENGER__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeCaller
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil || opts == nil {
			return
		}

		_L1StandardBridge.MESSENGER(opts)
	})
}

func Fuzz_Nosy_L1StandardBridgeCaller_Messenger__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeCaller
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil || opts == nil {
			return
		}

		_L1StandardBridge.Messenger(opts)
	})
}

func Fuzz_Nosy_L1StandardBridgeCaller_OTHERBRIDGE__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeCaller
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil || opts == nil {
			return
		}

		_L1StandardBridge.OTHERBRIDGE(opts)
	})
}

func Fuzz_Nosy_L1StandardBridgeCaller_OtherBridge__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeCaller
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil || opts == nil {
			return
		}

		_L1StandardBridge.OtherBridge(opts)
	})
}

func Fuzz_Nosy_L1StandardBridgeCaller_Paused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeCaller
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil || opts == nil {
			return
		}

		_L1StandardBridge.Paused(opts)
	})
}

func Fuzz_Nosy_L1StandardBridgeCaller_SuperchainConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeCaller
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil || opts == nil {
			return
		}

		_L1StandardBridge.SuperchainConfig(opts)
	})
}

func Fuzz_Nosy_L1StandardBridgeCaller_SystemConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeCaller
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil || opts == nil {
			return
		}

		_L1StandardBridge.SystemConfig(opts)
	})
}

func Fuzz_Nosy_L1StandardBridgeCaller_Version__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeCaller
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil || opts == nil {
			return
		}

		_L1StandardBridge.Version(opts)
	})
}

// skipping Fuzz_Nosy_L1StandardBridgeCallerRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

func Fuzz_Nosy_L1StandardBridgeCallerSession_Deposits__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeCallerSession
		fill_err = tp.Fill(&_L1StandardBridge)
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
		if _L1StandardBridge == nil {
			return
		}

		_L1StandardBridge.Deposits(arg0, arg1)
	})
}

func Fuzz_Nosy_L1StandardBridgeCallerSession_L2TokenBridge__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeCallerSession
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil {
			return
		}

		_L1StandardBridge.L2TokenBridge()
	})
}

func Fuzz_Nosy_L1StandardBridgeCallerSession_MESSENGER__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeCallerSession
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil {
			return
		}

		_L1StandardBridge.MESSENGER()
	})
}

func Fuzz_Nosy_L1StandardBridgeCallerSession_Messenger__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeCallerSession
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil {
			return
		}

		_L1StandardBridge.Messenger()
	})
}

func Fuzz_Nosy_L1StandardBridgeCallerSession_OTHERBRIDGE__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeCallerSession
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil {
			return
		}

		_L1StandardBridge.OTHERBRIDGE()
	})
}

func Fuzz_Nosy_L1StandardBridgeCallerSession_OtherBridge__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeCallerSession
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil {
			return
		}

		_L1StandardBridge.OtherBridge()
	})
}

func Fuzz_Nosy_L1StandardBridgeCallerSession_Paused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeCallerSession
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil {
			return
		}

		_L1StandardBridge.Paused()
	})
}

func Fuzz_Nosy_L1StandardBridgeCallerSession_SuperchainConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeCallerSession
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil {
			return
		}

		_L1StandardBridge.SuperchainConfig()
	})
}

func Fuzz_Nosy_L1StandardBridgeCallerSession_SystemConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeCallerSession
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil {
			return
		}

		_L1StandardBridge.SystemConfig()
	})
}

func Fuzz_Nosy_L1StandardBridgeCallerSession_Version__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeCallerSession
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil {
			return
		}

		_L1StandardBridge.Version()
	})
}

func Fuzz_Nosy_L1StandardBridgeERC20BridgeFinalizedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *L1StandardBridgeERC20BridgeFinalizedIterator
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

func Fuzz_Nosy_L1StandardBridgeERC20BridgeFinalizedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *L1StandardBridgeERC20BridgeFinalizedIterator
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

func Fuzz_Nosy_L1StandardBridgeERC20BridgeFinalizedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *L1StandardBridgeERC20BridgeFinalizedIterator
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

func Fuzz_Nosy_L1StandardBridgeERC20BridgeInitiatedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *L1StandardBridgeERC20BridgeInitiatedIterator
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

func Fuzz_Nosy_L1StandardBridgeERC20BridgeInitiatedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *L1StandardBridgeERC20BridgeInitiatedIterator
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

func Fuzz_Nosy_L1StandardBridgeERC20BridgeInitiatedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *L1StandardBridgeERC20BridgeInitiatedIterator
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

func Fuzz_Nosy_L1StandardBridgeERC20DepositInitiatedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *L1StandardBridgeERC20DepositInitiatedIterator
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

func Fuzz_Nosy_L1StandardBridgeERC20DepositInitiatedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *L1StandardBridgeERC20DepositInitiatedIterator
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

func Fuzz_Nosy_L1StandardBridgeERC20DepositInitiatedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *L1StandardBridgeERC20DepositInitiatedIterator
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

func Fuzz_Nosy_L1StandardBridgeERC20WithdrawalFinalizedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *L1StandardBridgeERC20WithdrawalFinalizedIterator
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

func Fuzz_Nosy_L1StandardBridgeERC20WithdrawalFinalizedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *L1StandardBridgeERC20WithdrawalFinalizedIterator
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

func Fuzz_Nosy_L1StandardBridgeERC20WithdrawalFinalizedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *L1StandardBridgeERC20WithdrawalFinalizedIterator
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

func Fuzz_Nosy_L1StandardBridgeETHBridgeFinalizedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *L1StandardBridgeETHBridgeFinalizedIterator
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

func Fuzz_Nosy_L1StandardBridgeETHBridgeFinalizedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *L1StandardBridgeETHBridgeFinalizedIterator
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

func Fuzz_Nosy_L1StandardBridgeETHBridgeFinalizedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *L1StandardBridgeETHBridgeFinalizedIterator
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

func Fuzz_Nosy_L1StandardBridgeETHBridgeInitiatedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *L1StandardBridgeETHBridgeInitiatedIterator
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

func Fuzz_Nosy_L1StandardBridgeETHBridgeInitiatedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *L1StandardBridgeETHBridgeInitiatedIterator
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

func Fuzz_Nosy_L1StandardBridgeETHBridgeInitiatedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *L1StandardBridgeETHBridgeInitiatedIterator
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

func Fuzz_Nosy_L1StandardBridgeETHDepositInitiatedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *L1StandardBridgeETHDepositInitiatedIterator
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

func Fuzz_Nosy_L1StandardBridgeETHDepositInitiatedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *L1StandardBridgeETHDepositInitiatedIterator
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

func Fuzz_Nosy_L1StandardBridgeETHDepositInitiatedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *L1StandardBridgeETHDepositInitiatedIterator
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

func Fuzz_Nosy_L1StandardBridgeETHWithdrawalFinalizedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *L1StandardBridgeETHWithdrawalFinalizedIterator
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

func Fuzz_Nosy_L1StandardBridgeETHWithdrawalFinalizedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *L1StandardBridgeETHWithdrawalFinalizedIterator
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

func Fuzz_Nosy_L1StandardBridgeETHWithdrawalFinalizedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *L1StandardBridgeETHWithdrawalFinalizedIterator
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

func Fuzz_Nosy_L1StandardBridgeFilterer_FilterERC20BridgeFinalized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeFilterer
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var localToken []common.Address
		fill_err = tp.Fill(&localToken)
		if fill_err != nil {
			return
		}
		var remoteToken []common.Address
		fill_err = tp.Fill(&remoteToken)
		if fill_err != nil {
			return
		}
		var from []common.Address
		fill_err = tp.Fill(&from)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil || opts == nil {
			return
		}

		_L1StandardBridge.FilterERC20BridgeFinalized(opts, localToken, remoteToken, from)
	})
}

func Fuzz_Nosy_L1StandardBridgeFilterer_FilterERC20BridgeInitiated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeFilterer
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var localToken []common.Address
		fill_err = tp.Fill(&localToken)
		if fill_err != nil {
			return
		}
		var remoteToken []common.Address
		fill_err = tp.Fill(&remoteToken)
		if fill_err != nil {
			return
		}
		var from []common.Address
		fill_err = tp.Fill(&from)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil || opts == nil {
			return
		}

		_L1StandardBridge.FilterERC20BridgeInitiated(opts, localToken, remoteToken, from)
	})
}

func Fuzz_Nosy_L1StandardBridgeFilterer_FilterERC20DepositInitiated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeFilterer
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var l1Token []common.Address
		fill_err = tp.Fill(&l1Token)
		if fill_err != nil {
			return
		}
		var l2Token []common.Address
		fill_err = tp.Fill(&l2Token)
		if fill_err != nil {
			return
		}
		var from []common.Address
		fill_err = tp.Fill(&from)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil || opts == nil {
			return
		}

		_L1StandardBridge.FilterERC20DepositInitiated(opts, l1Token, l2Token, from)
	})
}

func Fuzz_Nosy_L1StandardBridgeFilterer_FilterERC20WithdrawalFinalized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeFilterer
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var l1Token []common.Address
		fill_err = tp.Fill(&l1Token)
		if fill_err != nil {
			return
		}
		var l2Token []common.Address
		fill_err = tp.Fill(&l2Token)
		if fill_err != nil {
			return
		}
		var from []common.Address
		fill_err = tp.Fill(&from)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil || opts == nil {
			return
		}

		_L1StandardBridge.FilterERC20WithdrawalFinalized(opts, l1Token, l2Token, from)
	})
}

func Fuzz_Nosy_L1StandardBridgeFilterer_FilterETHBridgeFinalized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeFilterer
		fill_err = tp.Fill(&_L1StandardBridge)
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
		if _L1StandardBridge == nil || opts == nil {
			return
		}

		_L1StandardBridge.FilterETHBridgeFinalized(opts, from, to)
	})
}

func Fuzz_Nosy_L1StandardBridgeFilterer_FilterETHBridgeInitiated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeFilterer
		fill_err = tp.Fill(&_L1StandardBridge)
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
		if _L1StandardBridge == nil || opts == nil {
			return
		}

		_L1StandardBridge.FilterETHBridgeInitiated(opts, from, to)
	})
}

func Fuzz_Nosy_L1StandardBridgeFilterer_FilterETHDepositInitiated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeFilterer
		fill_err = tp.Fill(&_L1StandardBridge)
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
		if _L1StandardBridge == nil || opts == nil {
			return
		}

		_L1StandardBridge.FilterETHDepositInitiated(opts, from, to)
	})
}

func Fuzz_Nosy_L1StandardBridgeFilterer_FilterETHWithdrawalFinalized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeFilterer
		fill_err = tp.Fill(&_L1StandardBridge)
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
		if _L1StandardBridge == nil || opts == nil {
			return
		}

		_L1StandardBridge.FilterETHWithdrawalFinalized(opts, from, to)
	})
}

func Fuzz_Nosy_L1StandardBridgeFilterer_FilterInitialized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeFilterer
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil || opts == nil {
			return
		}

		_L1StandardBridge.FilterInitialized(opts)
	})
}

func Fuzz_Nosy_L1StandardBridgeFilterer_ParseERC20BridgeFinalized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeFilterer
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil {
			return
		}

		_L1StandardBridge.ParseERC20BridgeFinalized(log)
	})
}

func Fuzz_Nosy_L1StandardBridgeFilterer_ParseERC20BridgeInitiated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeFilterer
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil {
			return
		}

		_L1StandardBridge.ParseERC20BridgeInitiated(log)
	})
}

func Fuzz_Nosy_L1StandardBridgeFilterer_ParseERC20DepositInitiated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeFilterer
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil {
			return
		}

		_L1StandardBridge.ParseERC20DepositInitiated(log)
	})
}

func Fuzz_Nosy_L1StandardBridgeFilterer_ParseERC20WithdrawalFinalized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeFilterer
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil {
			return
		}

		_L1StandardBridge.ParseERC20WithdrawalFinalized(log)
	})
}

func Fuzz_Nosy_L1StandardBridgeFilterer_ParseETHBridgeFinalized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeFilterer
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil {
			return
		}

		_L1StandardBridge.ParseETHBridgeFinalized(log)
	})
}

func Fuzz_Nosy_L1StandardBridgeFilterer_ParseETHBridgeInitiated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeFilterer
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil {
			return
		}

		_L1StandardBridge.ParseETHBridgeInitiated(log)
	})
}

func Fuzz_Nosy_L1StandardBridgeFilterer_ParseETHDepositInitiated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeFilterer
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil {
			return
		}

		_L1StandardBridge.ParseETHDepositInitiated(log)
	})
}

func Fuzz_Nosy_L1StandardBridgeFilterer_ParseETHWithdrawalFinalized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeFilterer
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil {
			return
		}

		_L1StandardBridge.ParseETHWithdrawalFinalized(log)
	})
}

func Fuzz_Nosy_L1StandardBridgeFilterer_ParseInitialized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeFilterer
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil {
			return
		}

		_L1StandardBridge.ParseInitialized(log)
	})
}

// skipping Fuzz_Nosy_L1StandardBridgeFilterer_WatchERC20BridgeFinalized__ because parameters include func, chan, or unsupported interface: chan<- *github.com/ethereum-optimism/optimism/op-chain-ops/crossdomain/bindings.L1StandardBridgeERC20BridgeFinalized

// skipping Fuzz_Nosy_L1StandardBridgeFilterer_WatchERC20BridgeInitiated__ because parameters include func, chan, or unsupported interface: chan<- *github.com/ethereum-optimism/optimism/op-chain-ops/crossdomain/bindings.L1StandardBridgeERC20BridgeInitiated

// skipping Fuzz_Nosy_L1StandardBridgeFilterer_WatchERC20DepositInitiated__ because parameters include func, chan, or unsupported interface: chan<- *github.com/ethereum-optimism/optimism/op-chain-ops/crossdomain/bindings.L1StandardBridgeERC20DepositInitiated

// skipping Fuzz_Nosy_L1StandardBridgeFilterer_WatchERC20WithdrawalFinalized__ because parameters include func, chan, or unsupported interface: chan<- *github.com/ethereum-optimism/optimism/op-chain-ops/crossdomain/bindings.L1StandardBridgeERC20WithdrawalFinalized

// skipping Fuzz_Nosy_L1StandardBridgeFilterer_WatchETHBridgeFinalized__ because parameters include func, chan, or unsupported interface: chan<- *github.com/ethereum-optimism/optimism/op-chain-ops/crossdomain/bindings.L1StandardBridgeETHBridgeFinalized

// skipping Fuzz_Nosy_L1StandardBridgeFilterer_WatchETHBridgeInitiated__ because parameters include func, chan, or unsupported interface: chan<- *github.com/ethereum-optimism/optimism/op-chain-ops/crossdomain/bindings.L1StandardBridgeETHBridgeInitiated

// skipping Fuzz_Nosy_L1StandardBridgeFilterer_WatchETHDepositInitiated__ because parameters include func, chan, or unsupported interface: chan<- *github.com/ethereum-optimism/optimism/op-chain-ops/crossdomain/bindings.L1StandardBridgeETHDepositInitiated

// skipping Fuzz_Nosy_L1StandardBridgeFilterer_WatchETHWithdrawalFinalized__ because parameters include func, chan, or unsupported interface: chan<- *github.com/ethereum-optimism/optimism/op-chain-ops/crossdomain/bindings.L1StandardBridgeETHWithdrawalFinalized

// skipping Fuzz_Nosy_L1StandardBridgeFilterer_WatchInitialized__ because parameters include func, chan, or unsupported interface: chan<- *github.com/ethereum-optimism/optimism/op-chain-ops/crossdomain/bindings.L1StandardBridgeInitialized

func Fuzz_Nosy_L1StandardBridgeInitializedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *L1StandardBridgeInitializedIterator
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

func Fuzz_Nosy_L1StandardBridgeInitializedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *L1StandardBridgeInitializedIterator
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

func Fuzz_Nosy_L1StandardBridgeInitializedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *L1StandardBridgeInitializedIterator
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

// skipping Fuzz_Nosy_L1StandardBridgeRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

// skipping Fuzz_Nosy_L1StandardBridgeRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_L1StandardBridgeRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeRaw
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil || opts == nil {
			return
		}

		_L1StandardBridge.Transfer(opts)
	})
}

func Fuzz_Nosy_L1StandardBridgeSession_BridgeERC20__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeSession
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		var _localToken common.Address
		fill_err = tp.Fill(&_localToken)
		if fill_err != nil {
			return
		}
		var _remoteToken common.Address
		fill_err = tp.Fill(&_remoteToken)
		if fill_err != nil {
			return
		}
		var _amount *big.Int
		fill_err = tp.Fill(&_amount)
		if fill_err != nil {
			return
		}
		var _minGasLimit uint32
		fill_err = tp.Fill(&_minGasLimit)
		if fill_err != nil {
			return
		}
		var _extraData []byte
		fill_err = tp.Fill(&_extraData)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil || _amount == nil {
			return
		}

		_L1StandardBridge.BridgeERC20(_localToken, _remoteToken, _amount, _minGasLimit, _extraData)
	})
}

func Fuzz_Nosy_L1StandardBridgeSession_BridgeERC20To__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeSession
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		var _localToken common.Address
		fill_err = tp.Fill(&_localToken)
		if fill_err != nil {
			return
		}
		var _remoteToken common.Address
		fill_err = tp.Fill(&_remoteToken)
		if fill_err != nil {
			return
		}
		var _to common.Address
		fill_err = tp.Fill(&_to)
		if fill_err != nil {
			return
		}
		var _amount *big.Int
		fill_err = tp.Fill(&_amount)
		if fill_err != nil {
			return
		}
		var _minGasLimit uint32
		fill_err = tp.Fill(&_minGasLimit)
		if fill_err != nil {
			return
		}
		var _extraData []byte
		fill_err = tp.Fill(&_extraData)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil || _amount == nil {
			return
		}

		_L1StandardBridge.BridgeERC20To(_localToken, _remoteToken, _to, _amount, _minGasLimit, _extraData)
	})
}

func Fuzz_Nosy_L1StandardBridgeSession_BridgeETH__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeSession
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		var _minGasLimit uint32
		fill_err = tp.Fill(&_minGasLimit)
		if fill_err != nil {
			return
		}
		var _extraData []byte
		fill_err = tp.Fill(&_extraData)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil {
			return
		}

		_L1StandardBridge.BridgeETH(_minGasLimit, _extraData)
	})
}

func Fuzz_Nosy_L1StandardBridgeSession_BridgeETHTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeSession
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		var _to common.Address
		fill_err = tp.Fill(&_to)
		if fill_err != nil {
			return
		}
		var _minGasLimit uint32
		fill_err = tp.Fill(&_minGasLimit)
		if fill_err != nil {
			return
		}
		var _extraData []byte
		fill_err = tp.Fill(&_extraData)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil {
			return
		}

		_L1StandardBridge.BridgeETHTo(_to, _minGasLimit, _extraData)
	})
}

func Fuzz_Nosy_L1StandardBridgeSession_DepositERC20__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeSession
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		var _l1Token common.Address
		fill_err = tp.Fill(&_l1Token)
		if fill_err != nil {
			return
		}
		var _l2Token common.Address
		fill_err = tp.Fill(&_l2Token)
		if fill_err != nil {
			return
		}
		var _amount *big.Int
		fill_err = tp.Fill(&_amount)
		if fill_err != nil {
			return
		}
		var _minGasLimit uint32
		fill_err = tp.Fill(&_minGasLimit)
		if fill_err != nil {
			return
		}
		var _extraData []byte
		fill_err = tp.Fill(&_extraData)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil || _amount == nil {
			return
		}

		_L1StandardBridge.DepositERC20(_l1Token, _l2Token, _amount, _minGasLimit, _extraData)
	})
}

func Fuzz_Nosy_L1StandardBridgeSession_DepositERC20To__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeSession
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		var _l1Token common.Address
		fill_err = tp.Fill(&_l1Token)
		if fill_err != nil {
			return
		}
		var _l2Token common.Address
		fill_err = tp.Fill(&_l2Token)
		if fill_err != nil {
			return
		}
		var _to common.Address
		fill_err = tp.Fill(&_to)
		if fill_err != nil {
			return
		}
		var _amount *big.Int
		fill_err = tp.Fill(&_amount)
		if fill_err != nil {
			return
		}
		var _minGasLimit uint32
		fill_err = tp.Fill(&_minGasLimit)
		if fill_err != nil {
			return
		}
		var _extraData []byte
		fill_err = tp.Fill(&_extraData)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil || _amount == nil {
			return
		}

		_L1StandardBridge.DepositERC20To(_l1Token, _l2Token, _to, _amount, _minGasLimit, _extraData)
	})
}

func Fuzz_Nosy_L1StandardBridgeSession_DepositETH__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeSession
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		var _minGasLimit uint32
		fill_err = tp.Fill(&_minGasLimit)
		if fill_err != nil {
			return
		}
		var _extraData []byte
		fill_err = tp.Fill(&_extraData)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil {
			return
		}

		_L1StandardBridge.DepositETH(_minGasLimit, _extraData)
	})
}

func Fuzz_Nosy_L1StandardBridgeSession_DepositETHTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeSession
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		var _to common.Address
		fill_err = tp.Fill(&_to)
		if fill_err != nil {
			return
		}
		var _minGasLimit uint32
		fill_err = tp.Fill(&_minGasLimit)
		if fill_err != nil {
			return
		}
		var _extraData []byte
		fill_err = tp.Fill(&_extraData)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil {
			return
		}

		_L1StandardBridge.DepositETHTo(_to, _minGasLimit, _extraData)
	})
}

func Fuzz_Nosy_L1StandardBridgeSession_Deposits__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeSession
		fill_err = tp.Fill(&_L1StandardBridge)
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
		if _L1StandardBridge == nil {
			return
		}

		_L1StandardBridge.Deposits(arg0, arg1)
	})
}

func Fuzz_Nosy_L1StandardBridgeSession_FinalizeBridgeERC20__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeSession
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		var _localToken common.Address
		fill_err = tp.Fill(&_localToken)
		if fill_err != nil {
			return
		}
		var _remoteToken common.Address
		fill_err = tp.Fill(&_remoteToken)
		if fill_err != nil {
			return
		}
		var _from common.Address
		fill_err = tp.Fill(&_from)
		if fill_err != nil {
			return
		}
		var _to common.Address
		fill_err = tp.Fill(&_to)
		if fill_err != nil {
			return
		}
		var _amount *big.Int
		fill_err = tp.Fill(&_amount)
		if fill_err != nil {
			return
		}
		var _extraData []byte
		fill_err = tp.Fill(&_extraData)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil || _amount == nil {
			return
		}

		_L1StandardBridge.FinalizeBridgeERC20(_localToken, _remoteToken, _from, _to, _amount, _extraData)
	})
}

func Fuzz_Nosy_L1StandardBridgeSession_FinalizeBridgeETH__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeSession
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		var _from common.Address
		fill_err = tp.Fill(&_from)
		if fill_err != nil {
			return
		}
		var _to common.Address
		fill_err = tp.Fill(&_to)
		if fill_err != nil {
			return
		}
		var _amount *big.Int
		fill_err = tp.Fill(&_amount)
		if fill_err != nil {
			return
		}
		var _extraData []byte
		fill_err = tp.Fill(&_extraData)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil || _amount == nil {
			return
		}

		_L1StandardBridge.FinalizeBridgeETH(_from, _to, _amount, _extraData)
	})
}

func Fuzz_Nosy_L1StandardBridgeSession_FinalizeERC20Withdrawal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeSession
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		var _l1Token common.Address
		fill_err = tp.Fill(&_l1Token)
		if fill_err != nil {
			return
		}
		var _l2Token common.Address
		fill_err = tp.Fill(&_l2Token)
		if fill_err != nil {
			return
		}
		var _from common.Address
		fill_err = tp.Fill(&_from)
		if fill_err != nil {
			return
		}
		var _to common.Address
		fill_err = tp.Fill(&_to)
		if fill_err != nil {
			return
		}
		var _amount *big.Int
		fill_err = tp.Fill(&_amount)
		if fill_err != nil {
			return
		}
		var _extraData []byte
		fill_err = tp.Fill(&_extraData)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil || _amount == nil {
			return
		}

		_L1StandardBridge.FinalizeERC20Withdrawal(_l1Token, _l2Token, _from, _to, _amount, _extraData)
	})
}

func Fuzz_Nosy_L1StandardBridgeSession_FinalizeETHWithdrawal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeSession
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		var _from common.Address
		fill_err = tp.Fill(&_from)
		if fill_err != nil {
			return
		}
		var _to common.Address
		fill_err = tp.Fill(&_to)
		if fill_err != nil {
			return
		}
		var _amount *big.Int
		fill_err = tp.Fill(&_amount)
		if fill_err != nil {
			return
		}
		var _extraData []byte
		fill_err = tp.Fill(&_extraData)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil || _amount == nil {
			return
		}

		_L1StandardBridge.FinalizeETHWithdrawal(_from, _to, _amount, _extraData)
	})
}

func Fuzz_Nosy_L1StandardBridgeSession_Initialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeSession
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		var _messenger common.Address
		fill_err = tp.Fill(&_messenger)
		if fill_err != nil {
			return
		}
		var _superchainConfig common.Address
		fill_err = tp.Fill(&_superchainConfig)
		if fill_err != nil {
			return
		}
		var _systemConfig common.Address
		fill_err = tp.Fill(&_systemConfig)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil {
			return
		}

		_L1StandardBridge.Initialize(_messenger, _superchainConfig, _systemConfig)
	})
}

func Fuzz_Nosy_L1StandardBridgeSession_L2TokenBridge__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeSession
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil {
			return
		}

		_L1StandardBridge.L2TokenBridge()
	})
}

func Fuzz_Nosy_L1StandardBridgeSession_MESSENGER__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeSession
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil {
			return
		}

		_L1StandardBridge.MESSENGER()
	})
}

func Fuzz_Nosy_L1StandardBridgeSession_Messenger__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeSession
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil {
			return
		}

		_L1StandardBridge.Messenger()
	})
}

func Fuzz_Nosy_L1StandardBridgeSession_OTHERBRIDGE__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeSession
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil {
			return
		}

		_L1StandardBridge.OTHERBRIDGE()
	})
}

func Fuzz_Nosy_L1StandardBridgeSession_OtherBridge__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeSession
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil {
			return
		}

		_L1StandardBridge.OtherBridge()
	})
}

func Fuzz_Nosy_L1StandardBridgeSession_Paused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeSession
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil {
			return
		}

		_L1StandardBridge.Paused()
	})
}

func Fuzz_Nosy_L1StandardBridgeSession_Receive__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeSession
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil {
			return
		}

		_L1StandardBridge.Receive()
	})
}

func Fuzz_Nosy_L1StandardBridgeSession_SuperchainConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeSession
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil {
			return
		}

		_L1StandardBridge.SuperchainConfig()
	})
}

func Fuzz_Nosy_L1StandardBridgeSession_SystemConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeSession
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil {
			return
		}

		_L1StandardBridge.SystemConfig()
	})
}

func Fuzz_Nosy_L1StandardBridgeSession_Version__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeSession
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil {
			return
		}

		_L1StandardBridge.Version()
	})
}

func Fuzz_Nosy_L1StandardBridgeTransactor_BridgeERC20__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeTransactor
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var _localToken common.Address
		fill_err = tp.Fill(&_localToken)
		if fill_err != nil {
			return
		}
		var _remoteToken common.Address
		fill_err = tp.Fill(&_remoteToken)
		if fill_err != nil {
			return
		}
		var _amount *big.Int
		fill_err = tp.Fill(&_amount)
		if fill_err != nil {
			return
		}
		var _minGasLimit uint32
		fill_err = tp.Fill(&_minGasLimit)
		if fill_err != nil {
			return
		}
		var _extraData []byte
		fill_err = tp.Fill(&_extraData)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil || opts == nil || _amount == nil {
			return
		}

		_L1StandardBridge.BridgeERC20(opts, _localToken, _remoteToken, _amount, _minGasLimit, _extraData)
	})
}

func Fuzz_Nosy_L1StandardBridgeTransactor_BridgeERC20To__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeTransactor
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var _localToken common.Address
		fill_err = tp.Fill(&_localToken)
		if fill_err != nil {
			return
		}
		var _remoteToken common.Address
		fill_err = tp.Fill(&_remoteToken)
		if fill_err != nil {
			return
		}
		var _to common.Address
		fill_err = tp.Fill(&_to)
		if fill_err != nil {
			return
		}
		var _amount *big.Int
		fill_err = tp.Fill(&_amount)
		if fill_err != nil {
			return
		}
		var _minGasLimit uint32
		fill_err = tp.Fill(&_minGasLimit)
		if fill_err != nil {
			return
		}
		var _extraData []byte
		fill_err = tp.Fill(&_extraData)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil || opts == nil || _amount == nil {
			return
		}

		_L1StandardBridge.BridgeERC20To(opts, _localToken, _remoteToken, _to, _amount, _minGasLimit, _extraData)
	})
}

func Fuzz_Nosy_L1StandardBridgeTransactor_BridgeETH__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeTransactor
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var _minGasLimit uint32
		fill_err = tp.Fill(&_minGasLimit)
		if fill_err != nil {
			return
		}
		var _extraData []byte
		fill_err = tp.Fill(&_extraData)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil || opts == nil {
			return
		}

		_L1StandardBridge.BridgeETH(opts, _minGasLimit, _extraData)
	})
}

func Fuzz_Nosy_L1StandardBridgeTransactor_BridgeETHTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeTransactor
		fill_err = tp.Fill(&_L1StandardBridge)
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
		var _minGasLimit uint32
		fill_err = tp.Fill(&_minGasLimit)
		if fill_err != nil {
			return
		}
		var _extraData []byte
		fill_err = tp.Fill(&_extraData)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil || opts == nil {
			return
		}

		_L1StandardBridge.BridgeETHTo(opts, _to, _minGasLimit, _extraData)
	})
}

func Fuzz_Nosy_L1StandardBridgeTransactor_DepositERC20__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeTransactor
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var _l1Token common.Address
		fill_err = tp.Fill(&_l1Token)
		if fill_err != nil {
			return
		}
		var _l2Token common.Address
		fill_err = tp.Fill(&_l2Token)
		if fill_err != nil {
			return
		}
		var _amount *big.Int
		fill_err = tp.Fill(&_amount)
		if fill_err != nil {
			return
		}
		var _minGasLimit uint32
		fill_err = tp.Fill(&_minGasLimit)
		if fill_err != nil {
			return
		}
		var _extraData []byte
		fill_err = tp.Fill(&_extraData)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil || opts == nil || _amount == nil {
			return
		}

		_L1StandardBridge.DepositERC20(opts, _l1Token, _l2Token, _amount, _minGasLimit, _extraData)
	})
}

func Fuzz_Nosy_L1StandardBridgeTransactor_DepositERC20To__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeTransactor
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var _l1Token common.Address
		fill_err = tp.Fill(&_l1Token)
		if fill_err != nil {
			return
		}
		var _l2Token common.Address
		fill_err = tp.Fill(&_l2Token)
		if fill_err != nil {
			return
		}
		var _to common.Address
		fill_err = tp.Fill(&_to)
		if fill_err != nil {
			return
		}
		var _amount *big.Int
		fill_err = tp.Fill(&_amount)
		if fill_err != nil {
			return
		}
		var _minGasLimit uint32
		fill_err = tp.Fill(&_minGasLimit)
		if fill_err != nil {
			return
		}
		var _extraData []byte
		fill_err = tp.Fill(&_extraData)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil || opts == nil || _amount == nil {
			return
		}

		_L1StandardBridge.DepositERC20To(opts, _l1Token, _l2Token, _to, _amount, _minGasLimit, _extraData)
	})
}

func Fuzz_Nosy_L1StandardBridgeTransactor_DepositETH__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeTransactor
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var _minGasLimit uint32
		fill_err = tp.Fill(&_minGasLimit)
		if fill_err != nil {
			return
		}
		var _extraData []byte
		fill_err = tp.Fill(&_extraData)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil || opts == nil {
			return
		}

		_L1StandardBridge.DepositETH(opts, _minGasLimit, _extraData)
	})
}

func Fuzz_Nosy_L1StandardBridgeTransactor_DepositETHTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeTransactor
		fill_err = tp.Fill(&_L1StandardBridge)
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
		var _minGasLimit uint32
		fill_err = tp.Fill(&_minGasLimit)
		if fill_err != nil {
			return
		}
		var _extraData []byte
		fill_err = tp.Fill(&_extraData)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil || opts == nil {
			return
		}

		_L1StandardBridge.DepositETHTo(opts, _to, _minGasLimit, _extraData)
	})
}

func Fuzz_Nosy_L1StandardBridgeTransactor_FinalizeBridgeERC20__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeTransactor
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var _localToken common.Address
		fill_err = tp.Fill(&_localToken)
		if fill_err != nil {
			return
		}
		var _remoteToken common.Address
		fill_err = tp.Fill(&_remoteToken)
		if fill_err != nil {
			return
		}
		var _from common.Address
		fill_err = tp.Fill(&_from)
		if fill_err != nil {
			return
		}
		var _to common.Address
		fill_err = tp.Fill(&_to)
		if fill_err != nil {
			return
		}
		var _amount *big.Int
		fill_err = tp.Fill(&_amount)
		if fill_err != nil {
			return
		}
		var _extraData []byte
		fill_err = tp.Fill(&_extraData)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil || opts == nil || _amount == nil {
			return
		}

		_L1StandardBridge.FinalizeBridgeERC20(opts, _localToken, _remoteToken, _from, _to, _amount, _extraData)
	})
}

func Fuzz_Nosy_L1StandardBridgeTransactor_FinalizeBridgeETH__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeTransactor
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var _from common.Address
		fill_err = tp.Fill(&_from)
		if fill_err != nil {
			return
		}
		var _to common.Address
		fill_err = tp.Fill(&_to)
		if fill_err != nil {
			return
		}
		var _amount *big.Int
		fill_err = tp.Fill(&_amount)
		if fill_err != nil {
			return
		}
		var _extraData []byte
		fill_err = tp.Fill(&_extraData)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil || opts == nil || _amount == nil {
			return
		}

		_L1StandardBridge.FinalizeBridgeETH(opts, _from, _to, _amount, _extraData)
	})
}

func Fuzz_Nosy_L1StandardBridgeTransactor_FinalizeERC20Withdrawal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeTransactor
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var _l1Token common.Address
		fill_err = tp.Fill(&_l1Token)
		if fill_err != nil {
			return
		}
		var _l2Token common.Address
		fill_err = tp.Fill(&_l2Token)
		if fill_err != nil {
			return
		}
		var _from common.Address
		fill_err = tp.Fill(&_from)
		if fill_err != nil {
			return
		}
		var _to common.Address
		fill_err = tp.Fill(&_to)
		if fill_err != nil {
			return
		}
		var _amount *big.Int
		fill_err = tp.Fill(&_amount)
		if fill_err != nil {
			return
		}
		var _extraData []byte
		fill_err = tp.Fill(&_extraData)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil || opts == nil || _amount == nil {
			return
		}

		_L1StandardBridge.FinalizeERC20Withdrawal(opts, _l1Token, _l2Token, _from, _to, _amount, _extraData)
	})
}

func Fuzz_Nosy_L1StandardBridgeTransactor_FinalizeETHWithdrawal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeTransactor
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var _from common.Address
		fill_err = tp.Fill(&_from)
		if fill_err != nil {
			return
		}
		var _to common.Address
		fill_err = tp.Fill(&_to)
		if fill_err != nil {
			return
		}
		var _amount *big.Int
		fill_err = tp.Fill(&_amount)
		if fill_err != nil {
			return
		}
		var _extraData []byte
		fill_err = tp.Fill(&_extraData)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil || opts == nil || _amount == nil {
			return
		}

		_L1StandardBridge.FinalizeETHWithdrawal(opts, _from, _to, _amount, _extraData)
	})
}

func Fuzz_Nosy_L1StandardBridgeTransactor_Initialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeTransactor
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var _messenger common.Address
		fill_err = tp.Fill(&_messenger)
		if fill_err != nil {
			return
		}
		var _superchainConfig common.Address
		fill_err = tp.Fill(&_superchainConfig)
		if fill_err != nil {
			return
		}
		var _systemConfig common.Address
		fill_err = tp.Fill(&_systemConfig)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil || opts == nil {
			return
		}

		_L1StandardBridge.Initialize(opts, _messenger, _superchainConfig, _systemConfig)
	})
}

func Fuzz_Nosy_L1StandardBridgeTransactor_Receive__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeTransactor
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil || opts == nil {
			return
		}

		_L1StandardBridge.Receive(opts)
	})
}

// skipping Fuzz_Nosy_L1StandardBridgeTransactorRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_L1StandardBridgeTransactorRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeTransactorRaw
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil || opts == nil {
			return
		}

		_L1StandardBridge.Transfer(opts)
	})
}

func Fuzz_Nosy_L1StandardBridgeTransactorSession_BridgeERC20__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeTransactorSession
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		var _localToken common.Address
		fill_err = tp.Fill(&_localToken)
		if fill_err != nil {
			return
		}
		var _remoteToken common.Address
		fill_err = tp.Fill(&_remoteToken)
		if fill_err != nil {
			return
		}
		var _amount *big.Int
		fill_err = tp.Fill(&_amount)
		if fill_err != nil {
			return
		}
		var _minGasLimit uint32
		fill_err = tp.Fill(&_minGasLimit)
		if fill_err != nil {
			return
		}
		var _extraData []byte
		fill_err = tp.Fill(&_extraData)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil || _amount == nil {
			return
		}

		_L1StandardBridge.BridgeERC20(_localToken, _remoteToken, _amount, _minGasLimit, _extraData)
	})
}

func Fuzz_Nosy_L1StandardBridgeTransactorSession_BridgeERC20To__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeTransactorSession
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		var _localToken common.Address
		fill_err = tp.Fill(&_localToken)
		if fill_err != nil {
			return
		}
		var _remoteToken common.Address
		fill_err = tp.Fill(&_remoteToken)
		if fill_err != nil {
			return
		}
		var _to common.Address
		fill_err = tp.Fill(&_to)
		if fill_err != nil {
			return
		}
		var _amount *big.Int
		fill_err = tp.Fill(&_amount)
		if fill_err != nil {
			return
		}
		var _minGasLimit uint32
		fill_err = tp.Fill(&_minGasLimit)
		if fill_err != nil {
			return
		}
		var _extraData []byte
		fill_err = tp.Fill(&_extraData)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil || _amount == nil {
			return
		}

		_L1StandardBridge.BridgeERC20To(_localToken, _remoteToken, _to, _amount, _minGasLimit, _extraData)
	})
}

func Fuzz_Nosy_L1StandardBridgeTransactorSession_BridgeETH__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeTransactorSession
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		var _minGasLimit uint32
		fill_err = tp.Fill(&_minGasLimit)
		if fill_err != nil {
			return
		}
		var _extraData []byte
		fill_err = tp.Fill(&_extraData)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil {
			return
		}

		_L1StandardBridge.BridgeETH(_minGasLimit, _extraData)
	})
}

func Fuzz_Nosy_L1StandardBridgeTransactorSession_BridgeETHTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeTransactorSession
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		var _to common.Address
		fill_err = tp.Fill(&_to)
		if fill_err != nil {
			return
		}
		var _minGasLimit uint32
		fill_err = tp.Fill(&_minGasLimit)
		if fill_err != nil {
			return
		}
		var _extraData []byte
		fill_err = tp.Fill(&_extraData)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil {
			return
		}

		_L1StandardBridge.BridgeETHTo(_to, _minGasLimit, _extraData)
	})
}

func Fuzz_Nosy_L1StandardBridgeTransactorSession_DepositERC20__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeTransactorSession
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		var _l1Token common.Address
		fill_err = tp.Fill(&_l1Token)
		if fill_err != nil {
			return
		}
		var _l2Token common.Address
		fill_err = tp.Fill(&_l2Token)
		if fill_err != nil {
			return
		}
		var _amount *big.Int
		fill_err = tp.Fill(&_amount)
		if fill_err != nil {
			return
		}
		var _minGasLimit uint32
		fill_err = tp.Fill(&_minGasLimit)
		if fill_err != nil {
			return
		}
		var _extraData []byte
		fill_err = tp.Fill(&_extraData)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil || _amount == nil {
			return
		}

		_L1StandardBridge.DepositERC20(_l1Token, _l2Token, _amount, _minGasLimit, _extraData)
	})
}

func Fuzz_Nosy_L1StandardBridgeTransactorSession_DepositERC20To__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeTransactorSession
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		var _l1Token common.Address
		fill_err = tp.Fill(&_l1Token)
		if fill_err != nil {
			return
		}
		var _l2Token common.Address
		fill_err = tp.Fill(&_l2Token)
		if fill_err != nil {
			return
		}
		var _to common.Address
		fill_err = tp.Fill(&_to)
		if fill_err != nil {
			return
		}
		var _amount *big.Int
		fill_err = tp.Fill(&_amount)
		if fill_err != nil {
			return
		}
		var _minGasLimit uint32
		fill_err = tp.Fill(&_minGasLimit)
		if fill_err != nil {
			return
		}
		var _extraData []byte
		fill_err = tp.Fill(&_extraData)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil || _amount == nil {
			return
		}

		_L1StandardBridge.DepositERC20To(_l1Token, _l2Token, _to, _amount, _minGasLimit, _extraData)
	})
}

func Fuzz_Nosy_L1StandardBridgeTransactorSession_DepositETH__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeTransactorSession
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		var _minGasLimit uint32
		fill_err = tp.Fill(&_minGasLimit)
		if fill_err != nil {
			return
		}
		var _extraData []byte
		fill_err = tp.Fill(&_extraData)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil {
			return
		}

		_L1StandardBridge.DepositETH(_minGasLimit, _extraData)
	})
}

func Fuzz_Nosy_L1StandardBridgeTransactorSession_DepositETHTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeTransactorSession
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		var _to common.Address
		fill_err = tp.Fill(&_to)
		if fill_err != nil {
			return
		}
		var _minGasLimit uint32
		fill_err = tp.Fill(&_minGasLimit)
		if fill_err != nil {
			return
		}
		var _extraData []byte
		fill_err = tp.Fill(&_extraData)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil {
			return
		}

		_L1StandardBridge.DepositETHTo(_to, _minGasLimit, _extraData)
	})
}

func Fuzz_Nosy_L1StandardBridgeTransactorSession_FinalizeBridgeERC20__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeTransactorSession
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		var _localToken common.Address
		fill_err = tp.Fill(&_localToken)
		if fill_err != nil {
			return
		}
		var _remoteToken common.Address
		fill_err = tp.Fill(&_remoteToken)
		if fill_err != nil {
			return
		}
		var _from common.Address
		fill_err = tp.Fill(&_from)
		if fill_err != nil {
			return
		}
		var _to common.Address
		fill_err = tp.Fill(&_to)
		if fill_err != nil {
			return
		}
		var _amount *big.Int
		fill_err = tp.Fill(&_amount)
		if fill_err != nil {
			return
		}
		var _extraData []byte
		fill_err = tp.Fill(&_extraData)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil || _amount == nil {
			return
		}

		_L1StandardBridge.FinalizeBridgeERC20(_localToken, _remoteToken, _from, _to, _amount, _extraData)
	})
}

func Fuzz_Nosy_L1StandardBridgeTransactorSession_FinalizeBridgeETH__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeTransactorSession
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		var _from common.Address
		fill_err = tp.Fill(&_from)
		if fill_err != nil {
			return
		}
		var _to common.Address
		fill_err = tp.Fill(&_to)
		if fill_err != nil {
			return
		}
		var _amount *big.Int
		fill_err = tp.Fill(&_amount)
		if fill_err != nil {
			return
		}
		var _extraData []byte
		fill_err = tp.Fill(&_extraData)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil || _amount == nil {
			return
		}

		_L1StandardBridge.FinalizeBridgeETH(_from, _to, _amount, _extraData)
	})
}

func Fuzz_Nosy_L1StandardBridgeTransactorSession_FinalizeERC20Withdrawal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeTransactorSession
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		var _l1Token common.Address
		fill_err = tp.Fill(&_l1Token)
		if fill_err != nil {
			return
		}
		var _l2Token common.Address
		fill_err = tp.Fill(&_l2Token)
		if fill_err != nil {
			return
		}
		var _from common.Address
		fill_err = tp.Fill(&_from)
		if fill_err != nil {
			return
		}
		var _to common.Address
		fill_err = tp.Fill(&_to)
		if fill_err != nil {
			return
		}
		var _amount *big.Int
		fill_err = tp.Fill(&_amount)
		if fill_err != nil {
			return
		}
		var _extraData []byte
		fill_err = tp.Fill(&_extraData)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil || _amount == nil {
			return
		}

		_L1StandardBridge.FinalizeERC20Withdrawal(_l1Token, _l2Token, _from, _to, _amount, _extraData)
	})
}

func Fuzz_Nosy_L1StandardBridgeTransactorSession_FinalizeETHWithdrawal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeTransactorSession
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		var _from common.Address
		fill_err = tp.Fill(&_from)
		if fill_err != nil {
			return
		}
		var _to common.Address
		fill_err = tp.Fill(&_to)
		if fill_err != nil {
			return
		}
		var _amount *big.Int
		fill_err = tp.Fill(&_amount)
		if fill_err != nil {
			return
		}
		var _extraData []byte
		fill_err = tp.Fill(&_extraData)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil || _amount == nil {
			return
		}

		_L1StandardBridge.FinalizeETHWithdrawal(_from, _to, _amount, _extraData)
	})
}

func Fuzz_Nosy_L1StandardBridgeTransactorSession_Initialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeTransactorSession
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		var _messenger common.Address
		fill_err = tp.Fill(&_messenger)
		if fill_err != nil {
			return
		}
		var _superchainConfig common.Address
		fill_err = tp.Fill(&_superchainConfig)
		if fill_err != nil {
			return
		}
		var _systemConfig common.Address
		fill_err = tp.Fill(&_systemConfig)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil {
			return
		}

		_L1StandardBridge.Initialize(_messenger, _superchainConfig, _systemConfig)
	})
}

func Fuzz_Nosy_L1StandardBridgeTransactorSession_Receive__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L1StandardBridge *L1StandardBridgeTransactorSession
		fill_err = tp.Fill(&_L1StandardBridge)
		if fill_err != nil {
			return
		}
		if _L1StandardBridge == nil {
			return
		}

		_L1StandardBridge.Receive()
	})
}

func Fuzz_Nosy_L2CrossDomainMessengerCaller_BaseGas__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2CrossDomainMessenger *L2CrossDomainMessengerCaller
		fill_err = tp.Fill(&_L2CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var _message []byte
		fill_err = tp.Fill(&_message)
		if fill_err != nil {
			return
		}
		var _minGasLimit uint32
		fill_err = tp.Fill(&_minGasLimit)
		if fill_err != nil {
			return
		}
		if _L2CrossDomainMessenger == nil || opts == nil {
			return
		}

		_L2CrossDomainMessenger.BaseGas(opts, _message, _minGasLimit)
	})
}

func Fuzz_Nosy_L2CrossDomainMessengerCaller_FailedMessages__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2CrossDomainMessenger *L2CrossDomainMessengerCaller
		fill_err = tp.Fill(&_L2CrossDomainMessenger)
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
		if _L2CrossDomainMessenger == nil || opts == nil {
			return
		}

		_L2CrossDomainMessenger.FailedMessages(opts, arg0)
	})
}

func Fuzz_Nosy_L2CrossDomainMessengerCaller_L1CrossDomainMessenger__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2CrossDomainMessenger *L2CrossDomainMessengerCaller
		fill_err = tp.Fill(&_L2CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L2CrossDomainMessenger == nil || opts == nil {
			return
		}

		_L2CrossDomainMessenger.L1CrossDomainMessenger(opts)
	})
}

func Fuzz_Nosy_L2CrossDomainMessengerCaller_MESSAGEVERSION__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2CrossDomainMessenger *L2CrossDomainMessengerCaller
		fill_err = tp.Fill(&_L2CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L2CrossDomainMessenger == nil || opts == nil {
			return
		}

		_L2CrossDomainMessenger.MESSAGEVERSION(opts)
	})
}

func Fuzz_Nosy_L2CrossDomainMessengerCaller_MINGASCALLDATAOVERHEAD__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2CrossDomainMessenger *L2CrossDomainMessengerCaller
		fill_err = tp.Fill(&_L2CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L2CrossDomainMessenger == nil || opts == nil {
			return
		}

		_L2CrossDomainMessenger.MINGASCALLDATAOVERHEAD(opts)
	})
}

func Fuzz_Nosy_L2CrossDomainMessengerCaller_MINGASDYNAMICOVERHEADDENOMINATOR__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2CrossDomainMessenger *L2CrossDomainMessengerCaller
		fill_err = tp.Fill(&_L2CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L2CrossDomainMessenger == nil || opts == nil {
			return
		}

		_L2CrossDomainMessenger.MINGASDYNAMICOVERHEADDENOMINATOR(opts)
	})
}

func Fuzz_Nosy_L2CrossDomainMessengerCaller_MINGASDYNAMICOVERHEADNUMERATOR__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2CrossDomainMessenger *L2CrossDomainMessengerCaller
		fill_err = tp.Fill(&_L2CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L2CrossDomainMessenger == nil || opts == nil {
			return
		}

		_L2CrossDomainMessenger.MINGASDYNAMICOVERHEADNUMERATOR(opts)
	})
}

func Fuzz_Nosy_L2CrossDomainMessengerCaller_MessageNonce__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2CrossDomainMessenger *L2CrossDomainMessengerCaller
		fill_err = tp.Fill(&_L2CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L2CrossDomainMessenger == nil || opts == nil {
			return
		}

		_L2CrossDomainMessenger.MessageNonce(opts)
	})
}

func Fuzz_Nosy_L2CrossDomainMessengerCaller_OTHERMESSENGER__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2CrossDomainMessenger *L2CrossDomainMessengerCaller
		fill_err = tp.Fill(&_L2CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L2CrossDomainMessenger == nil || opts == nil {
			return
		}

		_L2CrossDomainMessenger.OTHERMESSENGER(opts)
	})
}

func Fuzz_Nosy_L2CrossDomainMessengerCaller_OtherMessenger__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2CrossDomainMessenger *L2CrossDomainMessengerCaller
		fill_err = tp.Fill(&_L2CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L2CrossDomainMessenger == nil || opts == nil {
			return
		}

		_L2CrossDomainMessenger.OtherMessenger(opts)
	})
}

func Fuzz_Nosy_L2CrossDomainMessengerCaller_Paused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2CrossDomainMessenger *L2CrossDomainMessengerCaller
		fill_err = tp.Fill(&_L2CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L2CrossDomainMessenger == nil || opts == nil {
			return
		}

		_L2CrossDomainMessenger.Paused(opts)
	})
}

func Fuzz_Nosy_L2CrossDomainMessengerCaller_RELAYCALLOVERHEAD__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2CrossDomainMessenger *L2CrossDomainMessengerCaller
		fill_err = tp.Fill(&_L2CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L2CrossDomainMessenger == nil || opts == nil {
			return
		}

		_L2CrossDomainMessenger.RELAYCALLOVERHEAD(opts)
	})
}

func Fuzz_Nosy_L2CrossDomainMessengerCaller_RELAYCONSTANTOVERHEAD__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2CrossDomainMessenger *L2CrossDomainMessengerCaller
		fill_err = tp.Fill(&_L2CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L2CrossDomainMessenger == nil || opts == nil {
			return
		}

		_L2CrossDomainMessenger.RELAYCONSTANTOVERHEAD(opts)
	})
}

func Fuzz_Nosy_L2CrossDomainMessengerCaller_RELAYGASCHECKBUFFER__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2CrossDomainMessenger *L2CrossDomainMessengerCaller
		fill_err = tp.Fill(&_L2CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L2CrossDomainMessenger == nil || opts == nil {
			return
		}

		_L2CrossDomainMessenger.RELAYGASCHECKBUFFER(opts)
	})
}

func Fuzz_Nosy_L2CrossDomainMessengerCaller_RELAYRESERVEDGAS__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2CrossDomainMessenger *L2CrossDomainMessengerCaller
		fill_err = tp.Fill(&_L2CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L2CrossDomainMessenger == nil || opts == nil {
			return
		}

		_L2CrossDomainMessenger.RELAYRESERVEDGAS(opts)
	})
}

func Fuzz_Nosy_L2CrossDomainMessengerCaller_SuccessfulMessages__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2CrossDomainMessenger *L2CrossDomainMessengerCaller
		fill_err = tp.Fill(&_L2CrossDomainMessenger)
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
		if _L2CrossDomainMessenger == nil || opts == nil {
			return
		}

		_L2CrossDomainMessenger.SuccessfulMessages(opts, arg0)
	})
}

func Fuzz_Nosy_L2CrossDomainMessengerCaller_Version__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2CrossDomainMessenger *L2CrossDomainMessengerCaller
		fill_err = tp.Fill(&_L2CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L2CrossDomainMessenger == nil || opts == nil {
			return
		}

		_L2CrossDomainMessenger.Version(opts)
	})
}

func Fuzz_Nosy_L2CrossDomainMessengerCaller_XDomainMessageSender__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2CrossDomainMessenger *L2CrossDomainMessengerCaller
		fill_err = tp.Fill(&_L2CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L2CrossDomainMessenger == nil || opts == nil {
			return
		}

		_L2CrossDomainMessenger.XDomainMessageSender(opts)
	})
}

// skipping Fuzz_Nosy_L2CrossDomainMessengerCallerRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

func Fuzz_Nosy_L2CrossDomainMessengerCallerSession_BaseGas__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2CrossDomainMessenger *L2CrossDomainMessengerCallerSession
		fill_err = tp.Fill(&_L2CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var _message []byte
		fill_err = tp.Fill(&_message)
		if fill_err != nil {
			return
		}
		var _minGasLimit uint32
		fill_err = tp.Fill(&_minGasLimit)
		if fill_err != nil {
			return
		}
		if _L2CrossDomainMessenger == nil {
			return
		}

		_L2CrossDomainMessenger.BaseGas(_message, _minGasLimit)
	})
}

func Fuzz_Nosy_L2CrossDomainMessengerCallerSession_FailedMessages__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2CrossDomainMessenger *L2CrossDomainMessengerCallerSession
		fill_err = tp.Fill(&_L2CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var arg0 [32]byte
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _L2CrossDomainMessenger == nil {
			return
		}

		_L2CrossDomainMessenger.FailedMessages(arg0)
	})
}

func Fuzz_Nosy_L2CrossDomainMessengerCallerSession_L1CrossDomainMessenger__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2CrossDomainMessenger *L2CrossDomainMessengerCallerSession
		fill_err = tp.Fill(&_L2CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		if _L2CrossDomainMessenger == nil {
			return
		}

		_L2CrossDomainMessenger.L1CrossDomainMessenger()
	})
}

func Fuzz_Nosy_L2CrossDomainMessengerCallerSession_MESSAGEVERSION__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2CrossDomainMessenger *L2CrossDomainMessengerCallerSession
		fill_err = tp.Fill(&_L2CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		if _L2CrossDomainMessenger == nil {
			return
		}

		_L2CrossDomainMessenger.MESSAGEVERSION()
	})
}

func Fuzz_Nosy_L2CrossDomainMessengerCallerSession_MINGASCALLDATAOVERHEAD__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2CrossDomainMessenger *L2CrossDomainMessengerCallerSession
		fill_err = tp.Fill(&_L2CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		if _L2CrossDomainMessenger == nil {
			return
		}

		_L2CrossDomainMessenger.MINGASCALLDATAOVERHEAD()
	})
}

func Fuzz_Nosy_L2CrossDomainMessengerCallerSession_MINGASDYNAMICOVERHEADDENOMINATOR__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2CrossDomainMessenger *L2CrossDomainMessengerCallerSession
		fill_err = tp.Fill(&_L2CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		if _L2CrossDomainMessenger == nil {
			return
		}

		_L2CrossDomainMessenger.MINGASDYNAMICOVERHEADDENOMINATOR()
	})
}

func Fuzz_Nosy_L2CrossDomainMessengerCallerSession_MINGASDYNAMICOVERHEADNUMERATOR__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2CrossDomainMessenger *L2CrossDomainMessengerCallerSession
		fill_err = tp.Fill(&_L2CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		if _L2CrossDomainMessenger == nil {
			return
		}

		_L2CrossDomainMessenger.MINGASDYNAMICOVERHEADNUMERATOR()
	})
}

func Fuzz_Nosy_L2CrossDomainMessengerCallerSession_MessageNonce__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2CrossDomainMessenger *L2CrossDomainMessengerCallerSession
		fill_err = tp.Fill(&_L2CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		if _L2CrossDomainMessenger == nil {
			return
		}

		_L2CrossDomainMessenger.MessageNonce()
	})
}

func Fuzz_Nosy_L2CrossDomainMessengerCallerSession_OTHERMESSENGER__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2CrossDomainMessenger *L2CrossDomainMessengerCallerSession
		fill_err = tp.Fill(&_L2CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		if _L2CrossDomainMessenger == nil {
			return
		}

		_L2CrossDomainMessenger.OTHERMESSENGER()
	})
}

func Fuzz_Nosy_L2CrossDomainMessengerCallerSession_OtherMessenger__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2CrossDomainMessenger *L2CrossDomainMessengerCallerSession
		fill_err = tp.Fill(&_L2CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		if _L2CrossDomainMessenger == nil {
			return
		}

		_L2CrossDomainMessenger.OtherMessenger()
	})
}

func Fuzz_Nosy_L2CrossDomainMessengerCallerSession_Paused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2CrossDomainMessenger *L2CrossDomainMessengerCallerSession
		fill_err = tp.Fill(&_L2CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		if _L2CrossDomainMessenger == nil {
			return
		}

		_L2CrossDomainMessenger.Paused()
	})
}

func Fuzz_Nosy_L2CrossDomainMessengerCallerSession_RELAYCALLOVERHEAD__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2CrossDomainMessenger *L2CrossDomainMessengerCallerSession
		fill_err = tp.Fill(&_L2CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		if _L2CrossDomainMessenger == nil {
			return
		}

		_L2CrossDomainMessenger.RELAYCALLOVERHEAD()
	})
}

func Fuzz_Nosy_L2CrossDomainMessengerCallerSession_RELAYCONSTANTOVERHEAD__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2CrossDomainMessenger *L2CrossDomainMessengerCallerSession
		fill_err = tp.Fill(&_L2CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		if _L2CrossDomainMessenger == nil {
			return
		}

		_L2CrossDomainMessenger.RELAYCONSTANTOVERHEAD()
	})
}

func Fuzz_Nosy_L2CrossDomainMessengerCallerSession_RELAYGASCHECKBUFFER__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2CrossDomainMessenger *L2CrossDomainMessengerCallerSession
		fill_err = tp.Fill(&_L2CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		if _L2CrossDomainMessenger == nil {
			return
		}

		_L2CrossDomainMessenger.RELAYGASCHECKBUFFER()
	})
}

func Fuzz_Nosy_L2CrossDomainMessengerCallerSession_RELAYRESERVEDGAS__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2CrossDomainMessenger *L2CrossDomainMessengerCallerSession
		fill_err = tp.Fill(&_L2CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		if _L2CrossDomainMessenger == nil {
			return
		}

		_L2CrossDomainMessenger.RELAYRESERVEDGAS()
	})
}

func Fuzz_Nosy_L2CrossDomainMessengerCallerSession_SuccessfulMessages__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2CrossDomainMessenger *L2CrossDomainMessengerCallerSession
		fill_err = tp.Fill(&_L2CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var arg0 [32]byte
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _L2CrossDomainMessenger == nil {
			return
		}

		_L2CrossDomainMessenger.SuccessfulMessages(arg0)
	})
}

func Fuzz_Nosy_L2CrossDomainMessengerCallerSession_Version__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2CrossDomainMessenger *L2CrossDomainMessengerCallerSession
		fill_err = tp.Fill(&_L2CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		if _L2CrossDomainMessenger == nil {
			return
		}

		_L2CrossDomainMessenger.Version()
	})
}

func Fuzz_Nosy_L2CrossDomainMessengerCallerSession_XDomainMessageSender__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2CrossDomainMessenger *L2CrossDomainMessengerCallerSession
		fill_err = tp.Fill(&_L2CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		if _L2CrossDomainMessenger == nil {
			return
		}

		_L2CrossDomainMessenger.XDomainMessageSender()
	})
}

func Fuzz_Nosy_L2CrossDomainMessengerFailedRelayedMessageIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *L2CrossDomainMessengerFailedRelayedMessageIterator
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

func Fuzz_Nosy_L2CrossDomainMessengerFailedRelayedMessageIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *L2CrossDomainMessengerFailedRelayedMessageIterator
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

func Fuzz_Nosy_L2CrossDomainMessengerFailedRelayedMessageIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *L2CrossDomainMessengerFailedRelayedMessageIterator
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

func Fuzz_Nosy_L2CrossDomainMessengerFilterer_FilterFailedRelayedMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2CrossDomainMessenger *L2CrossDomainMessengerFilterer
		fill_err = tp.Fill(&_L2CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var msgHash [][32]byte
		fill_err = tp.Fill(&msgHash)
		if fill_err != nil {
			return
		}
		if _L2CrossDomainMessenger == nil || opts == nil {
			return
		}

		_L2CrossDomainMessenger.FilterFailedRelayedMessage(opts, msgHash)
	})
}

func Fuzz_Nosy_L2CrossDomainMessengerFilterer_FilterInitialized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2CrossDomainMessenger *L2CrossDomainMessengerFilterer
		fill_err = tp.Fill(&_L2CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L2CrossDomainMessenger == nil || opts == nil {
			return
		}

		_L2CrossDomainMessenger.FilterInitialized(opts)
	})
}

func Fuzz_Nosy_L2CrossDomainMessengerFilterer_FilterRelayedMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2CrossDomainMessenger *L2CrossDomainMessengerFilterer
		fill_err = tp.Fill(&_L2CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var msgHash [][32]byte
		fill_err = tp.Fill(&msgHash)
		if fill_err != nil {
			return
		}
		if _L2CrossDomainMessenger == nil || opts == nil {
			return
		}

		_L2CrossDomainMessenger.FilterRelayedMessage(opts, msgHash)
	})
}

func Fuzz_Nosy_L2CrossDomainMessengerFilterer_FilterSentMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2CrossDomainMessenger *L2CrossDomainMessengerFilterer
		fill_err = tp.Fill(&_L2CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var t3 []common.Address
		fill_err = tp.Fill(&t3)
		if fill_err != nil {
			return
		}
		if _L2CrossDomainMessenger == nil || opts == nil {
			return
		}

		_L2CrossDomainMessenger.FilterSentMessage(opts, t3)
	})
}

func Fuzz_Nosy_L2CrossDomainMessengerFilterer_FilterSentMessageExtension1__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2CrossDomainMessenger *L2CrossDomainMessengerFilterer
		fill_err = tp.Fill(&_L2CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var opts *bind.FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var sender []common.Address
		fill_err = tp.Fill(&sender)
		if fill_err != nil {
			return
		}
		if _L2CrossDomainMessenger == nil || opts == nil {
			return
		}

		_L2CrossDomainMessenger.FilterSentMessageExtension1(opts, sender)
	})
}

func Fuzz_Nosy_L2CrossDomainMessengerFilterer_ParseFailedRelayedMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2CrossDomainMessenger *L2CrossDomainMessengerFilterer
		fill_err = tp.Fill(&_L2CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _L2CrossDomainMessenger == nil {
			return
		}

		_L2CrossDomainMessenger.ParseFailedRelayedMessage(log)
	})
}

func Fuzz_Nosy_L2CrossDomainMessengerFilterer_ParseInitialized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2CrossDomainMessenger *L2CrossDomainMessengerFilterer
		fill_err = tp.Fill(&_L2CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _L2CrossDomainMessenger == nil {
			return
		}

		_L2CrossDomainMessenger.ParseInitialized(log)
	})
}

func Fuzz_Nosy_L2CrossDomainMessengerFilterer_ParseRelayedMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2CrossDomainMessenger *L2CrossDomainMessengerFilterer
		fill_err = tp.Fill(&_L2CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _L2CrossDomainMessenger == nil {
			return
		}

		_L2CrossDomainMessenger.ParseRelayedMessage(log)
	})
}

func Fuzz_Nosy_L2CrossDomainMessengerFilterer_ParseSentMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2CrossDomainMessenger *L2CrossDomainMessengerFilterer
		fill_err = tp.Fill(&_L2CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _L2CrossDomainMessenger == nil {
			return
		}

		_L2CrossDomainMessenger.ParseSentMessage(log)
	})
}

func Fuzz_Nosy_L2CrossDomainMessengerFilterer_ParseSentMessageExtension1__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2CrossDomainMessenger *L2CrossDomainMessengerFilterer
		fill_err = tp.Fill(&_L2CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var log types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if _L2CrossDomainMessenger == nil {
			return
		}

		_L2CrossDomainMessenger.ParseSentMessageExtension1(log)
	})
}

// skipping Fuzz_Nosy_L2CrossDomainMessengerFilterer_WatchFailedRelayedMessage__ because parameters include func, chan, or unsupported interface: chan<- *github.com/ethereum-optimism/optimism/op-chain-ops/crossdomain/bindings.L2CrossDomainMessengerFailedRelayedMessage

// skipping Fuzz_Nosy_L2CrossDomainMessengerFilterer_WatchInitialized__ because parameters include func, chan, or unsupported interface: chan<- *github.com/ethereum-optimism/optimism/op-chain-ops/crossdomain/bindings.L2CrossDomainMessengerInitialized

// skipping Fuzz_Nosy_L2CrossDomainMessengerFilterer_WatchRelayedMessage__ because parameters include func, chan, or unsupported interface: chan<- *github.com/ethereum-optimism/optimism/op-chain-ops/crossdomain/bindings.L2CrossDomainMessengerRelayedMessage

// skipping Fuzz_Nosy_L2CrossDomainMessengerFilterer_WatchSentMessage__ because parameters include func, chan, or unsupported interface: chan<- *github.com/ethereum-optimism/optimism/op-chain-ops/crossdomain/bindings.L2CrossDomainMessengerSentMessage

// skipping Fuzz_Nosy_L2CrossDomainMessengerFilterer_WatchSentMessageExtension1__ because parameters include func, chan, or unsupported interface: chan<- *github.com/ethereum-optimism/optimism/op-chain-ops/crossdomain/bindings.L2CrossDomainMessengerSentMessageExtension1

func Fuzz_Nosy_L2CrossDomainMessengerInitializedIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *L2CrossDomainMessengerInitializedIterator
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

func Fuzz_Nosy_L2CrossDomainMessengerInitializedIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *L2CrossDomainMessengerInitializedIterator
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

func Fuzz_Nosy_L2CrossDomainMessengerInitializedIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *L2CrossDomainMessengerInitializedIterator
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

// skipping Fuzz_Nosy_L2CrossDomainMessengerRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

// skipping Fuzz_Nosy_L2CrossDomainMessengerRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_L2CrossDomainMessengerRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2CrossDomainMessenger *L2CrossDomainMessengerRaw
		fill_err = tp.Fill(&_L2CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L2CrossDomainMessenger == nil || opts == nil {
			return
		}

		_L2CrossDomainMessenger.Transfer(opts)
	})
}

func Fuzz_Nosy_L2CrossDomainMessengerRelayedMessageIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *L2CrossDomainMessengerRelayedMessageIterator
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

func Fuzz_Nosy_L2CrossDomainMessengerRelayedMessageIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *L2CrossDomainMessengerRelayedMessageIterator
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

func Fuzz_Nosy_L2CrossDomainMessengerRelayedMessageIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *L2CrossDomainMessengerRelayedMessageIterator
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

func Fuzz_Nosy_L2CrossDomainMessengerSentMessageExtension1Iterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *L2CrossDomainMessengerSentMessageExtension1Iterator
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

func Fuzz_Nosy_L2CrossDomainMessengerSentMessageExtension1Iterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *L2CrossDomainMessengerSentMessageExtension1Iterator
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

func Fuzz_Nosy_L2CrossDomainMessengerSentMessageExtension1Iterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *L2CrossDomainMessengerSentMessageExtension1Iterator
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

func Fuzz_Nosy_L2CrossDomainMessengerSentMessageIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *L2CrossDomainMessengerSentMessageIterator
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

func Fuzz_Nosy_L2CrossDomainMessengerSentMessageIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *L2CrossDomainMessengerSentMessageIterator
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

func Fuzz_Nosy_L2CrossDomainMessengerSentMessageIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *L2CrossDomainMessengerSentMessageIterator
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

func Fuzz_Nosy_L2CrossDomainMessengerSession_BaseGas__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2CrossDomainMessenger *L2CrossDomainMessengerSession
		fill_err = tp.Fill(&_L2CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var _message []byte
		fill_err = tp.Fill(&_message)
		if fill_err != nil {
			return
		}
		var _minGasLimit uint32
		fill_err = tp.Fill(&_minGasLimit)
		if fill_err != nil {
			return
		}
		if _L2CrossDomainMessenger == nil {
			return
		}

		_L2CrossDomainMessenger.BaseGas(_message, _minGasLimit)
	})
}

func Fuzz_Nosy_L2CrossDomainMessengerSession_FailedMessages__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2CrossDomainMessenger *L2CrossDomainMessengerSession
		fill_err = tp.Fill(&_L2CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var arg0 [32]byte
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _L2CrossDomainMessenger == nil {
			return
		}

		_L2CrossDomainMessenger.FailedMessages(arg0)
	})
}

func Fuzz_Nosy_L2CrossDomainMessengerSession_Initialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2CrossDomainMessenger *L2CrossDomainMessengerSession
		fill_err = tp.Fill(&_L2CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var _l1CrossDomainMessenger common.Address
		fill_err = tp.Fill(&_l1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		if _L2CrossDomainMessenger == nil {
			return
		}

		_L2CrossDomainMessenger.Initialize(_l1CrossDomainMessenger)
	})
}

func Fuzz_Nosy_L2CrossDomainMessengerSession_L1CrossDomainMessenger__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2CrossDomainMessenger *L2CrossDomainMessengerSession
		fill_err = tp.Fill(&_L2CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		if _L2CrossDomainMessenger == nil {
			return
		}

		_L2CrossDomainMessenger.L1CrossDomainMessenger()
	})
}

func Fuzz_Nosy_L2CrossDomainMessengerSession_MESSAGEVERSION__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2CrossDomainMessenger *L2CrossDomainMessengerSession
		fill_err = tp.Fill(&_L2CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		if _L2CrossDomainMessenger == nil {
			return
		}

		_L2CrossDomainMessenger.MESSAGEVERSION()
	})
}

func Fuzz_Nosy_L2CrossDomainMessengerSession_MINGASCALLDATAOVERHEAD__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2CrossDomainMessenger *L2CrossDomainMessengerSession
		fill_err = tp.Fill(&_L2CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		if _L2CrossDomainMessenger == nil {
			return
		}

		_L2CrossDomainMessenger.MINGASCALLDATAOVERHEAD()
	})
}

func Fuzz_Nosy_L2CrossDomainMessengerSession_MINGASDYNAMICOVERHEADDENOMINATOR__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2CrossDomainMessenger *L2CrossDomainMessengerSession
		fill_err = tp.Fill(&_L2CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		if _L2CrossDomainMessenger == nil {
			return
		}

		_L2CrossDomainMessenger.MINGASDYNAMICOVERHEADDENOMINATOR()
	})
}

func Fuzz_Nosy_L2CrossDomainMessengerSession_MINGASDYNAMICOVERHEADNUMERATOR__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2CrossDomainMessenger *L2CrossDomainMessengerSession
		fill_err = tp.Fill(&_L2CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		if _L2CrossDomainMessenger == nil {
			return
		}

		_L2CrossDomainMessenger.MINGASDYNAMICOVERHEADNUMERATOR()
	})
}

func Fuzz_Nosy_L2CrossDomainMessengerSession_MessageNonce__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2CrossDomainMessenger *L2CrossDomainMessengerSession
		fill_err = tp.Fill(&_L2CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		if _L2CrossDomainMessenger == nil {
			return
		}

		_L2CrossDomainMessenger.MessageNonce()
	})
}

func Fuzz_Nosy_L2CrossDomainMessengerSession_OTHERMESSENGER__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2CrossDomainMessenger *L2CrossDomainMessengerSession
		fill_err = tp.Fill(&_L2CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		if _L2CrossDomainMessenger == nil {
			return
		}

		_L2CrossDomainMessenger.OTHERMESSENGER()
	})
}

func Fuzz_Nosy_L2CrossDomainMessengerSession_OtherMessenger__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2CrossDomainMessenger *L2CrossDomainMessengerSession
		fill_err = tp.Fill(&_L2CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		if _L2CrossDomainMessenger == nil {
			return
		}

		_L2CrossDomainMessenger.OtherMessenger()
	})
}

func Fuzz_Nosy_L2CrossDomainMessengerSession_Paused__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2CrossDomainMessenger *L2CrossDomainMessengerSession
		fill_err = tp.Fill(&_L2CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		if _L2CrossDomainMessenger == nil {
			return
		}

		_L2CrossDomainMessenger.Paused()
	})
}

func Fuzz_Nosy_L2CrossDomainMessengerSession_RELAYCALLOVERHEAD__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2CrossDomainMessenger *L2CrossDomainMessengerSession
		fill_err = tp.Fill(&_L2CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		if _L2CrossDomainMessenger == nil {
			return
		}

		_L2CrossDomainMessenger.RELAYCALLOVERHEAD()
	})
}

func Fuzz_Nosy_L2CrossDomainMessengerSession_RELAYCONSTANTOVERHEAD__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2CrossDomainMessenger *L2CrossDomainMessengerSession
		fill_err = tp.Fill(&_L2CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		if _L2CrossDomainMessenger == nil {
			return
		}

		_L2CrossDomainMessenger.RELAYCONSTANTOVERHEAD()
	})
}

func Fuzz_Nosy_L2CrossDomainMessengerSession_RELAYGASCHECKBUFFER__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2CrossDomainMessenger *L2CrossDomainMessengerSession
		fill_err = tp.Fill(&_L2CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		if _L2CrossDomainMessenger == nil {
			return
		}

		_L2CrossDomainMessenger.RELAYGASCHECKBUFFER()
	})
}

func Fuzz_Nosy_L2CrossDomainMessengerSession_RELAYRESERVEDGAS__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2CrossDomainMessenger *L2CrossDomainMessengerSession
		fill_err = tp.Fill(&_L2CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		if _L2CrossDomainMessenger == nil {
			return
		}

		_L2CrossDomainMessenger.RELAYRESERVEDGAS()
	})
}

func Fuzz_Nosy_L2CrossDomainMessengerSession_RelayMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2CrossDomainMessenger *L2CrossDomainMessengerSession
		fill_err = tp.Fill(&_L2CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var _nonce *big.Int
		fill_err = tp.Fill(&_nonce)
		if fill_err != nil {
			return
		}
		var _sender common.Address
		fill_err = tp.Fill(&_sender)
		if fill_err != nil {
			return
		}
		var _target common.Address
		fill_err = tp.Fill(&_target)
		if fill_err != nil {
			return
		}
		var _value *big.Int
		fill_err = tp.Fill(&_value)
		if fill_err != nil {
			return
		}
		var _minGasLimit *big.Int
		fill_err = tp.Fill(&_minGasLimit)
		if fill_err != nil {
			return
		}
		var _message []byte
		fill_err = tp.Fill(&_message)
		if fill_err != nil {
			return
		}
		if _L2CrossDomainMessenger == nil || _nonce == nil || _value == nil || _minGasLimit == nil {
			return
		}

		_L2CrossDomainMessenger.RelayMessage(_nonce, _sender, _target, _value, _minGasLimit, _message)
	})
}

func Fuzz_Nosy_L2CrossDomainMessengerSession_SendMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2CrossDomainMessenger *L2CrossDomainMessengerSession
		fill_err = tp.Fill(&_L2CrossDomainMessenger)
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
		var _minGasLimit uint32
		fill_err = tp.Fill(&_minGasLimit)
		if fill_err != nil {
			return
		}
		if _L2CrossDomainMessenger == nil {
			return
		}

		_L2CrossDomainMessenger.SendMessage(_target, _message, _minGasLimit)
	})
}

func Fuzz_Nosy_L2CrossDomainMessengerSession_SuccessfulMessages__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2CrossDomainMessenger *L2CrossDomainMessengerSession
		fill_err = tp.Fill(&_L2CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var arg0 [32]byte
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		if _L2CrossDomainMessenger == nil {
			return
		}

		_L2CrossDomainMessenger.SuccessfulMessages(arg0)
	})
}

func Fuzz_Nosy_L2CrossDomainMessengerSession_Version__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2CrossDomainMessenger *L2CrossDomainMessengerSession
		fill_err = tp.Fill(&_L2CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		if _L2CrossDomainMessenger == nil {
			return
		}

		_L2CrossDomainMessenger.Version()
	})
}

func Fuzz_Nosy_L2CrossDomainMessengerSession_XDomainMessageSender__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2CrossDomainMessenger *L2CrossDomainMessengerSession
		fill_err = tp.Fill(&_L2CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		if _L2CrossDomainMessenger == nil {
			return
		}

		_L2CrossDomainMessenger.XDomainMessageSender()
	})
}

func Fuzz_Nosy_L2CrossDomainMessengerTransactor_Initialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2CrossDomainMessenger *L2CrossDomainMessengerTransactor
		fill_err = tp.Fill(&_L2CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var _l1CrossDomainMessenger common.Address
		fill_err = tp.Fill(&_l1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		if _L2CrossDomainMessenger == nil || opts == nil {
			return
		}

		_L2CrossDomainMessenger.Initialize(opts, _l1CrossDomainMessenger)
	})
}

func Fuzz_Nosy_L2CrossDomainMessengerTransactor_RelayMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2CrossDomainMessenger *L2CrossDomainMessengerTransactor
		fill_err = tp.Fill(&_L2CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var _nonce *big.Int
		fill_err = tp.Fill(&_nonce)
		if fill_err != nil {
			return
		}
		var _sender common.Address
		fill_err = tp.Fill(&_sender)
		if fill_err != nil {
			return
		}
		var _target common.Address
		fill_err = tp.Fill(&_target)
		if fill_err != nil {
			return
		}
		var _value *big.Int
		fill_err = tp.Fill(&_value)
		if fill_err != nil {
			return
		}
		var _minGasLimit *big.Int
		fill_err = tp.Fill(&_minGasLimit)
		if fill_err != nil {
			return
		}
		var _message []byte
		fill_err = tp.Fill(&_message)
		if fill_err != nil {
			return
		}
		if _L2CrossDomainMessenger == nil || opts == nil || _nonce == nil || _value == nil || _minGasLimit == nil {
			return
		}

		_L2CrossDomainMessenger.RelayMessage(opts, _nonce, _sender, _target, _value, _minGasLimit, _message)
	})
}

func Fuzz_Nosy_L2CrossDomainMessengerTransactor_SendMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2CrossDomainMessenger *L2CrossDomainMessengerTransactor
		fill_err = tp.Fill(&_L2CrossDomainMessenger)
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
		var _message []byte
		fill_err = tp.Fill(&_message)
		if fill_err != nil {
			return
		}
		var _minGasLimit uint32
		fill_err = tp.Fill(&_minGasLimit)
		if fill_err != nil {
			return
		}
		if _L2CrossDomainMessenger == nil || opts == nil {
			return
		}

		_L2CrossDomainMessenger.SendMessage(opts, _target, _message, _minGasLimit)
	})
}

// skipping Fuzz_Nosy_L2CrossDomainMessengerTransactorRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_L2CrossDomainMessengerTransactorRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2CrossDomainMessenger *L2CrossDomainMessengerTransactorRaw
		fill_err = tp.Fill(&_L2CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _L2CrossDomainMessenger == nil || opts == nil {
			return
		}

		_L2CrossDomainMessenger.Transfer(opts)
	})
}

func Fuzz_Nosy_L2CrossDomainMessengerTransactorSession_Initialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2CrossDomainMessenger *L2CrossDomainMessengerTransactorSession
		fill_err = tp.Fill(&_L2CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var _l1CrossDomainMessenger common.Address
		fill_err = tp.Fill(&_l1CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		if _L2CrossDomainMessenger == nil {
			return
		}

		_L2CrossDomainMessenger.Initialize(_l1CrossDomainMessenger)
	})
}

func Fuzz_Nosy_L2CrossDomainMessengerTransactorSession_RelayMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2CrossDomainMessenger *L2CrossDomainMessengerTransactorSession
		fill_err = tp.Fill(&_L2CrossDomainMessenger)
		if fill_err != nil {
			return
		}
		var _nonce *big.Int
		fill_err = tp.Fill(&_nonce)
		if fill_err != nil {
			return
		}
		var _sender common.Address
		fill_err = tp.Fill(&_sender)
		if fill_err != nil {
			return
		}
		var _target common.Address
		fill_err = tp.Fill(&_target)
		if fill_err != nil {
			return
		}
		var _value *big.Int
		fill_err = tp.Fill(&_value)
		if fill_err != nil {
			return
		}
		var _minGasLimit *big.Int
		fill_err = tp.Fill(&_minGasLimit)
		if fill_err != nil {
			return
		}
		var _message []byte
		fill_err = tp.Fill(&_message)
		if fill_err != nil {
			return
		}
		if _L2CrossDomainMessenger == nil || _nonce == nil || _value == nil || _minGasLimit == nil {
			return
		}

		_L2CrossDomainMessenger.RelayMessage(_nonce, _sender, _target, _value, _minGasLimit, _message)
	})
}

func Fuzz_Nosy_L2CrossDomainMessengerTransactorSession_SendMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _L2CrossDomainMessenger *L2CrossDomainMessengerTransactorSession
		fill_err = tp.Fill(&_L2CrossDomainMessenger)
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
		var _minGasLimit uint32
		fill_err = tp.Fill(&_minGasLimit)
		if fill_err != nil {
			return
		}
		if _L2CrossDomainMessenger == nil {
			return
		}

		_L2CrossDomainMessenger.SendMessage(_target, _message, _minGasLimit)
	})
}

// skipping Fuzz_Nosy_DeployL1CrossDomainMessenger__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/abi/bind.ContractBackend

// skipping Fuzz_Nosy_DeployL1StandardBridge__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/abi/bind.ContractBackend

// skipping Fuzz_Nosy_DeployL2CrossDomainMessenger__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/abi/bind.ContractBackend
