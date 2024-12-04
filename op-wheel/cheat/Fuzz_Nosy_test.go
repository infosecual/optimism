package cheat

import (
	"testing"

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

func Fuzz_Nosy_Cheater_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dataDirPath string, readOnly bool) {
		ch, err := OpenGethDB(dataDirPath, readOnly)
		if err != nil {
			return
		}
		ch.Close()
	})
}

// skipping Fuzz_Nosy_Cheater_RunAndClose__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-wheel/cheat.HeadFn

func Fuzz_Nosy_blockBodyKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var number uint64
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}

		blockBodyKey(number, hash)
	})
}

func Fuzz_Nosy_encodeBlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, number uint64) {
		encodeBlockNumber(number)
	})
}
