package helpers

import (
	"testing"

	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils"
	"github.com/ethereum/go-ethereum/common/hexutil"
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

func Fuzz_Nosy_ApplyDeltaTimeOffset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dp *e2eutils.DeployParams
		fill_err = tp.Fill(&dp)
		if fill_err != nil {
			return
		}
		var deltaTimeOffset *hexutil.Uint64
		fill_err = tp.Fill(&deltaTimeOffset)
		if fill_err != nil {
			return
		}
		if dp == nil || deltaTimeOffset == nil {
			return
		}

		ApplyDeltaTimeOffset(dp, deltaTimeOffset)
	})
}
