package bindings

import (
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

func Fuzz_Nosy_GasPriceOracleCaller_BaseFee__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _GasPriceOracle *GasPriceOracleCaller
		fill_err = tp.Fill(&_GasPriceOracle)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _GasPriceOracle == nil || opts == nil {
			return
		}

		_GasPriceOracle.BaseFee(opts)
	})
}

func Fuzz_Nosy_GasPriceOracleCaller_BaseFeeScalar__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _GasPriceOracle *GasPriceOracleCaller
		fill_err = tp.Fill(&_GasPriceOracle)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _GasPriceOracle == nil || opts == nil {
			return
		}

		_GasPriceOracle.BaseFeeScalar(opts)
	})
}

func Fuzz_Nosy_GasPriceOracleCaller_BlobBaseFee__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _GasPriceOracle *GasPriceOracleCaller
		fill_err = tp.Fill(&_GasPriceOracle)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _GasPriceOracle == nil || opts == nil {
			return
		}

		_GasPriceOracle.BlobBaseFee(opts)
	})
}

func Fuzz_Nosy_GasPriceOracleCaller_BlobBaseFeeScalar__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _GasPriceOracle *GasPriceOracleCaller
		fill_err = tp.Fill(&_GasPriceOracle)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _GasPriceOracle == nil || opts == nil {
			return
		}

		_GasPriceOracle.BlobBaseFeeScalar(opts)
	})
}

func Fuzz_Nosy_GasPriceOracleCaller_DECIMALS__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _GasPriceOracle *GasPriceOracleCaller
		fill_err = tp.Fill(&_GasPriceOracle)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _GasPriceOracle == nil || opts == nil {
			return
		}

		_GasPriceOracle.DECIMALS(opts)
	})
}

func Fuzz_Nosy_GasPriceOracleCaller_Decimals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _GasPriceOracle *GasPriceOracleCaller
		fill_err = tp.Fill(&_GasPriceOracle)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _GasPriceOracle == nil || opts == nil {
			return
		}

		_GasPriceOracle.Decimals(opts)
	})
}

func Fuzz_Nosy_GasPriceOracleCaller_GasPrice__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _GasPriceOracle *GasPriceOracleCaller
		fill_err = tp.Fill(&_GasPriceOracle)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _GasPriceOracle == nil || opts == nil {
			return
		}

		_GasPriceOracle.GasPrice(opts)
	})
}

func Fuzz_Nosy_GasPriceOracleCaller_GetL1Fee__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _GasPriceOracle *GasPriceOracleCaller
		fill_err = tp.Fill(&_GasPriceOracle)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var _data []byte
		fill_err = tp.Fill(&_data)
		if fill_err != nil {
			return
		}
		if _GasPriceOracle == nil || opts == nil {
			return
		}

		_GasPriceOracle.GetL1Fee(opts, _data)
	})
}

func Fuzz_Nosy_GasPriceOracleCaller_GetL1GasUsed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _GasPriceOracle *GasPriceOracleCaller
		fill_err = tp.Fill(&_GasPriceOracle)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var _data []byte
		fill_err = tp.Fill(&_data)
		if fill_err != nil {
			return
		}
		if _GasPriceOracle == nil || opts == nil {
			return
		}

		_GasPriceOracle.GetL1GasUsed(opts, _data)
	})
}

func Fuzz_Nosy_GasPriceOracleCaller_IsEcotone__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _GasPriceOracle *GasPriceOracleCaller
		fill_err = tp.Fill(&_GasPriceOracle)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _GasPriceOracle == nil || opts == nil {
			return
		}

		_GasPriceOracle.IsEcotone(opts)
	})
}

func Fuzz_Nosy_GasPriceOracleCaller_L1BaseFee__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _GasPriceOracle *GasPriceOracleCaller
		fill_err = tp.Fill(&_GasPriceOracle)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _GasPriceOracle == nil || opts == nil {
			return
		}

		_GasPriceOracle.L1BaseFee(opts)
	})
}

func Fuzz_Nosy_GasPriceOracleCaller_Overhead__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _GasPriceOracle *GasPriceOracleCaller
		fill_err = tp.Fill(&_GasPriceOracle)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _GasPriceOracle == nil || opts == nil {
			return
		}

		_GasPriceOracle.Overhead(opts)
	})
}

func Fuzz_Nosy_GasPriceOracleCaller_Scalar__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _GasPriceOracle *GasPriceOracleCaller
		fill_err = tp.Fill(&_GasPriceOracle)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _GasPriceOracle == nil || opts == nil {
			return
		}

		_GasPriceOracle.Scalar(opts)
	})
}

func Fuzz_Nosy_GasPriceOracleCaller_Version__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _GasPriceOracle *GasPriceOracleCaller
		fill_err = tp.Fill(&_GasPriceOracle)
		if fill_err != nil {
			return
		}
		var opts *bind.CallOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _GasPriceOracle == nil || opts == nil {
			return
		}

		_GasPriceOracle.Version(opts)
	})
}

// skipping Fuzz_Nosy_GasPriceOracleCallerRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

func Fuzz_Nosy_GasPriceOracleCallerSession_BaseFee__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _GasPriceOracle *GasPriceOracleCallerSession
		fill_err = tp.Fill(&_GasPriceOracle)
		if fill_err != nil {
			return
		}
		if _GasPriceOracle == nil {
			return
		}

		_GasPriceOracle.BaseFee()
	})
}

func Fuzz_Nosy_GasPriceOracleCallerSession_BaseFeeScalar__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _GasPriceOracle *GasPriceOracleCallerSession
		fill_err = tp.Fill(&_GasPriceOracle)
		if fill_err != nil {
			return
		}
		if _GasPriceOracle == nil {
			return
		}

		_GasPriceOracle.BaseFeeScalar()
	})
}

func Fuzz_Nosy_GasPriceOracleCallerSession_BlobBaseFee__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _GasPriceOracle *GasPriceOracleCallerSession
		fill_err = tp.Fill(&_GasPriceOracle)
		if fill_err != nil {
			return
		}
		if _GasPriceOracle == nil {
			return
		}

		_GasPriceOracle.BlobBaseFee()
	})
}

func Fuzz_Nosy_GasPriceOracleCallerSession_BlobBaseFeeScalar__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _GasPriceOracle *GasPriceOracleCallerSession
		fill_err = tp.Fill(&_GasPriceOracle)
		if fill_err != nil {
			return
		}
		if _GasPriceOracle == nil {
			return
		}

		_GasPriceOracle.BlobBaseFeeScalar()
	})
}

func Fuzz_Nosy_GasPriceOracleCallerSession_DECIMALS__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _GasPriceOracle *GasPriceOracleCallerSession
		fill_err = tp.Fill(&_GasPriceOracle)
		if fill_err != nil {
			return
		}
		if _GasPriceOracle == nil {
			return
		}

		_GasPriceOracle.DECIMALS()
	})
}

func Fuzz_Nosy_GasPriceOracleCallerSession_Decimals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _GasPriceOracle *GasPriceOracleCallerSession
		fill_err = tp.Fill(&_GasPriceOracle)
		if fill_err != nil {
			return
		}
		if _GasPriceOracle == nil {
			return
		}

		_GasPriceOracle.Decimals()
	})
}

func Fuzz_Nosy_GasPriceOracleCallerSession_GasPrice__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _GasPriceOracle *GasPriceOracleCallerSession
		fill_err = tp.Fill(&_GasPriceOracle)
		if fill_err != nil {
			return
		}
		if _GasPriceOracle == nil {
			return
		}

		_GasPriceOracle.GasPrice()
	})
}

func Fuzz_Nosy_GasPriceOracleCallerSession_GetL1Fee__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _GasPriceOracle *GasPriceOracleCallerSession
		fill_err = tp.Fill(&_GasPriceOracle)
		if fill_err != nil {
			return
		}
		var _data []byte
		fill_err = tp.Fill(&_data)
		if fill_err != nil {
			return
		}
		if _GasPriceOracle == nil {
			return
		}

		_GasPriceOracle.GetL1Fee(_data)
	})
}

func Fuzz_Nosy_GasPriceOracleCallerSession_GetL1GasUsed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _GasPriceOracle *GasPriceOracleCallerSession
		fill_err = tp.Fill(&_GasPriceOracle)
		if fill_err != nil {
			return
		}
		var _data []byte
		fill_err = tp.Fill(&_data)
		if fill_err != nil {
			return
		}
		if _GasPriceOracle == nil {
			return
		}

		_GasPriceOracle.GetL1GasUsed(_data)
	})
}

func Fuzz_Nosy_GasPriceOracleCallerSession_IsEcotone__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _GasPriceOracle *GasPriceOracleCallerSession
		fill_err = tp.Fill(&_GasPriceOracle)
		if fill_err != nil {
			return
		}
		if _GasPriceOracle == nil {
			return
		}

		_GasPriceOracle.IsEcotone()
	})
}

func Fuzz_Nosy_GasPriceOracleCallerSession_L1BaseFee__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _GasPriceOracle *GasPriceOracleCallerSession
		fill_err = tp.Fill(&_GasPriceOracle)
		if fill_err != nil {
			return
		}
		if _GasPriceOracle == nil {
			return
		}

		_GasPriceOracle.L1BaseFee()
	})
}

func Fuzz_Nosy_GasPriceOracleCallerSession_Overhead__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _GasPriceOracle *GasPriceOracleCallerSession
		fill_err = tp.Fill(&_GasPriceOracle)
		if fill_err != nil {
			return
		}
		if _GasPriceOracle == nil {
			return
		}

		_GasPriceOracle.Overhead()
	})
}

func Fuzz_Nosy_GasPriceOracleCallerSession_Scalar__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _GasPriceOracle *GasPriceOracleCallerSession
		fill_err = tp.Fill(&_GasPriceOracle)
		if fill_err != nil {
			return
		}
		if _GasPriceOracle == nil {
			return
		}

		_GasPriceOracle.Scalar()
	})
}

func Fuzz_Nosy_GasPriceOracleCallerSession_Version__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _GasPriceOracle *GasPriceOracleCallerSession
		fill_err = tp.Fill(&_GasPriceOracle)
		if fill_err != nil {
			return
		}
		if _GasPriceOracle == nil {
			return
		}

		_GasPriceOracle.Version()
	})
}

// skipping Fuzz_Nosy_GasPriceOracleRaw_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

// skipping Fuzz_Nosy_GasPriceOracleRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_GasPriceOracleRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _GasPriceOracle *GasPriceOracleRaw
		fill_err = tp.Fill(&_GasPriceOracle)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _GasPriceOracle == nil || opts == nil {
			return
		}

		_GasPriceOracle.Transfer(opts)
	})
}

func Fuzz_Nosy_GasPriceOracleSession_BaseFee__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _GasPriceOracle *GasPriceOracleSession
		fill_err = tp.Fill(&_GasPriceOracle)
		if fill_err != nil {
			return
		}
		if _GasPriceOracle == nil {
			return
		}

		_GasPriceOracle.BaseFee()
	})
}

func Fuzz_Nosy_GasPriceOracleSession_BaseFeeScalar__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _GasPriceOracle *GasPriceOracleSession
		fill_err = tp.Fill(&_GasPriceOracle)
		if fill_err != nil {
			return
		}
		if _GasPriceOracle == nil {
			return
		}

		_GasPriceOracle.BaseFeeScalar()
	})
}

func Fuzz_Nosy_GasPriceOracleSession_BlobBaseFee__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _GasPriceOracle *GasPriceOracleSession
		fill_err = tp.Fill(&_GasPriceOracle)
		if fill_err != nil {
			return
		}
		if _GasPriceOracle == nil {
			return
		}

		_GasPriceOracle.BlobBaseFee()
	})
}

func Fuzz_Nosy_GasPriceOracleSession_BlobBaseFeeScalar__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _GasPriceOracle *GasPriceOracleSession
		fill_err = tp.Fill(&_GasPriceOracle)
		if fill_err != nil {
			return
		}
		if _GasPriceOracle == nil {
			return
		}

		_GasPriceOracle.BlobBaseFeeScalar()
	})
}

func Fuzz_Nosy_GasPriceOracleSession_DECIMALS__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _GasPriceOracle *GasPriceOracleSession
		fill_err = tp.Fill(&_GasPriceOracle)
		if fill_err != nil {
			return
		}
		if _GasPriceOracle == nil {
			return
		}

		_GasPriceOracle.DECIMALS()
	})
}

func Fuzz_Nosy_GasPriceOracleSession_Decimals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _GasPriceOracle *GasPriceOracleSession
		fill_err = tp.Fill(&_GasPriceOracle)
		if fill_err != nil {
			return
		}
		if _GasPriceOracle == nil {
			return
		}

		_GasPriceOracle.Decimals()
	})
}

func Fuzz_Nosy_GasPriceOracleSession_GasPrice__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _GasPriceOracle *GasPriceOracleSession
		fill_err = tp.Fill(&_GasPriceOracle)
		if fill_err != nil {
			return
		}
		if _GasPriceOracle == nil {
			return
		}

		_GasPriceOracle.GasPrice()
	})
}

func Fuzz_Nosy_GasPriceOracleSession_GetL1Fee__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _GasPriceOracle *GasPriceOracleSession
		fill_err = tp.Fill(&_GasPriceOracle)
		if fill_err != nil {
			return
		}
		var _data []byte
		fill_err = tp.Fill(&_data)
		if fill_err != nil {
			return
		}
		if _GasPriceOracle == nil {
			return
		}

		_GasPriceOracle.GetL1Fee(_data)
	})
}

func Fuzz_Nosy_GasPriceOracleSession_GetL1GasUsed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _GasPriceOracle *GasPriceOracleSession
		fill_err = tp.Fill(&_GasPriceOracle)
		if fill_err != nil {
			return
		}
		var _data []byte
		fill_err = tp.Fill(&_data)
		if fill_err != nil {
			return
		}
		if _GasPriceOracle == nil {
			return
		}

		_GasPriceOracle.GetL1GasUsed(_data)
	})
}

func Fuzz_Nosy_GasPriceOracleSession_IsEcotone__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _GasPriceOracle *GasPriceOracleSession
		fill_err = tp.Fill(&_GasPriceOracle)
		if fill_err != nil {
			return
		}
		if _GasPriceOracle == nil {
			return
		}

		_GasPriceOracle.IsEcotone()
	})
}

func Fuzz_Nosy_GasPriceOracleSession_L1BaseFee__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _GasPriceOracle *GasPriceOracleSession
		fill_err = tp.Fill(&_GasPriceOracle)
		if fill_err != nil {
			return
		}
		if _GasPriceOracle == nil {
			return
		}

		_GasPriceOracle.L1BaseFee()
	})
}

func Fuzz_Nosy_GasPriceOracleSession_Overhead__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _GasPriceOracle *GasPriceOracleSession
		fill_err = tp.Fill(&_GasPriceOracle)
		if fill_err != nil {
			return
		}
		if _GasPriceOracle == nil {
			return
		}

		_GasPriceOracle.Overhead()
	})
}

func Fuzz_Nosy_GasPriceOracleSession_Scalar__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _GasPriceOracle *GasPriceOracleSession
		fill_err = tp.Fill(&_GasPriceOracle)
		if fill_err != nil {
			return
		}
		if _GasPriceOracle == nil {
			return
		}

		_GasPriceOracle.Scalar()
	})
}

func Fuzz_Nosy_GasPriceOracleSession_SetEcotone__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _GasPriceOracle *GasPriceOracleSession
		fill_err = tp.Fill(&_GasPriceOracle)
		if fill_err != nil {
			return
		}
		if _GasPriceOracle == nil {
			return
		}

		_GasPriceOracle.SetEcotone()
	})
}

func Fuzz_Nosy_GasPriceOracleSession_Version__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _GasPriceOracle *GasPriceOracleSession
		fill_err = tp.Fill(&_GasPriceOracle)
		if fill_err != nil {
			return
		}
		if _GasPriceOracle == nil {
			return
		}

		_GasPriceOracle.Version()
	})
}

func Fuzz_Nosy_GasPriceOracleTransactor_SetEcotone__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _GasPriceOracle *GasPriceOracleTransactor
		fill_err = tp.Fill(&_GasPriceOracle)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _GasPriceOracle == nil || opts == nil {
			return
		}

		_GasPriceOracle.SetEcotone(opts)
	})
}

// skipping Fuzz_Nosy_GasPriceOracleTransactorRaw_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_GasPriceOracleTransactorRaw_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _GasPriceOracle *GasPriceOracleTransactorRaw
		fill_err = tp.Fill(&_GasPriceOracle)
		if fill_err != nil {
			return
		}
		var opts *bind.TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if _GasPriceOracle == nil || opts == nil {
			return
		}

		_GasPriceOracle.Transfer(opts)
	})
}

func Fuzz_Nosy_GasPriceOracleTransactorSession_SetEcotone__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _GasPriceOracle *GasPriceOracleTransactorSession
		fill_err = tp.Fill(&_GasPriceOracle)
		if fill_err != nil {
			return
		}
		if _GasPriceOracle == nil {
			return
		}

		_GasPriceOracle.SetEcotone()
	})
}

// skipping Fuzz_Nosy_DeployGasPriceOracle__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/abi/bind.ContractBackend
