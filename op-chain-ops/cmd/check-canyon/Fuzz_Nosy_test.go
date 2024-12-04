package main

import (
	"bytes"
	"testing"

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

// skipping Fuzz_Nosy_L2Client_BlockByNumber__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-chain-ops/cmd/check-canyon.L2Client

// skipping Fuzz_Nosy_L2Client_CodeAt__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-chain-ops/cmd/check-canyon.L2Client

// skipping Fuzz_Nosy_L2Client_FetchReceipts__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-chain-ops/cmd/check-canyon.L2Client

// skipping Fuzz_Nosy_L2Client_InfoByNumber__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-chain-ops/cmd/check-canyon.L2Client

func Fuzz_Nosy_rawReceipts_EncodeIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s rawReceipts
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var i int
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
		var w *bytes.Buffer
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		s.EncodeIndex(i, w)
	})
}

func Fuzz_Nosy_rawReceipts_Len__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s rawReceipts
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}

		s.Len()
	})
}

// skipping Fuzz_Nosy_CheckActivation__ because parameters include func, chan, or unsupported interface: func(github.com/ethereum-optimism/optimism/op-chain-ops/cmd/check-canyon.Args, bool) error

// skipping Fuzz_Nosy_CheckInactivation__ because parameters include func, chan, or unsupported interface: func(github.com/ethereum-optimism/optimism/op-chain-ops/cmd/check-canyon.Args, bool) error

func Fuzz_Nosy_ManuallyEncodeReceipts__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var receipts types.Receipts
		fill_err = tp.Fill(&receipts)
		if fill_err != nil {
			return
		}
		var canyonActive bool
		fill_err = tp.Fill(&canyonActive)
		if fill_err != nil {
			return
		}

		ManuallyEncodeReceipts(receipts, canyonActive)
	})
}
