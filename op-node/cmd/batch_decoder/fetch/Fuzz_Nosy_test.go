package fetch

import (
	"testing"

	"github.com/ethereum-optimism/optimism/op-service/sources"
	"github.com/ethereum/go-ethereum/ethclient"
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

func Fuzz_Nosy_Batches__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var client *ethclient.Client
		fill_err = tp.Fill(&client)
		if fill_err != nil {
			return
		}
		var beacon *sources.L1BeaconClient
		fill_err = tp.Fill(&beacon)
		if fill_err != nil {
			return
		}
		var config Config
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		if client == nil || beacon == nil {
			return
		}

		Batches(client, beacon, config)
	})
}

// skipping Fuzz_Nosy_fetchBatchesPerBlock__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/types.Signer
