package client

import (
	"testing"

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

func Fuzz_Nosy_BootstrapClient_BootInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var br *BootstrapClient
		fill_err = tp.Fill(&br)
		if fill_err != nil {
			return
		}
		if br == nil {
			return
		}

		br.BootInfo()
	})
}

// skipping Fuzz_Nosy_oracleClient_Get__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-program/client.oracleClient

// skipping Fuzz_Nosy_Main__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/log.Logger
