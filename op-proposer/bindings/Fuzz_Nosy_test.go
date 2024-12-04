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

// skipping Fuzz_Nosy_L2OutputOracleFilterer_WatchInitialized__ because parameters include func, chan, or unsupported interface: chan<- *github.com/ethereum-optimism/optimism/op-proposer/bindings.L2OutputOracleInitialized

// skipping Fuzz_Nosy_L2OutputOracleFilterer_WatchOutputProposed__ because parameters include func, chan, or unsupported interface: chan<- *github.com/ethereum-optimism/optimism/op-proposer/bindings.L2OutputOracleOutputProposed

// skipping Fuzz_Nosy_L2OutputOracleFilterer_WatchOutputsDeleted__ because parameters include func, chan, or unsupported interface: chan<- *github.com/ethereum-optimism/optimism/op-proposer/bindings.L2OutputOracleOutputsDeleted

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

// skipping Fuzz_Nosy_DeployL2OutputOracle__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/abi/bind.ContractBackend
