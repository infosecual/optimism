package reassemble

import (
	"testing"

	"github.com/ethereum-optimism/optimism/op-node/cmd/batch_decoder/fetch"
	"github.com/ethereum-optimism/optimism/op-node/rollup"
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

func Fuzz_Nosy_Channels__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config Config
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var rollupCfg *rollup.Config
		fill_err = tp.Fill(&rollupCfg)
		if fill_err != nil {
			return
		}
		if rollupCfg == nil {
			return
		}

		Channels(config, rollupCfg)
	})
}

func Fuzz_Nosy_LoadFrames__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var directory string
		fill_err = tp.Fill(&directory)
		if fill_err != nil {
			return
		}
		var inbox common.Address
		fill_err = tp.Fill(&inbox)
		if fill_err != nil {
			return
		}

		LoadFrames(directory, inbox)
	})
}

func Fuzz_Nosy_loadTransactions__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dir string
		fill_err = tp.Fill(&dir)
		if fill_err != nil {
			return
		}
		var inbox common.Address
		fill_err = tp.Fill(&inbox)
		if fill_err != nil {
			return
		}

		loadTransactions(dir, inbox)
	})
}

func Fuzz_Nosy_transactionsToFrames__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var txns []fetch.TransactionWithMetadata
		fill_err = tp.Fill(&txns)
		if fill_err != nil {
			return
		}

		transactionsToFrames(txns)
	})
}
