package main

import (
	"testing"

	go_fuzz_utils "github.com/trailofbits/go-fuzz-utils"
	"github.com/urfave/cli/v2"
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

func Fuzz_Nosy_AddrFromFlag__(f *testing.F) {
	f.Fuzz(func(t *testing.T, flagName string) {
		AddrFromFlag(flagName)
	})
}

// skipping Fuzz_Nosy_NewContractWithTxMgr__ because parameters include func, chan, or unsupported interface: func(ctx *github.com/urfave/cli/v2.Context) (github.com/ethereum/go-ethereum/common.Address, error)

func Fuzz_Nosy_newClientsFromCLI__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx *cli.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if ctx == nil {
			return
		}

		newClientsFromCLI(ctx)
	})
}

// skipping Fuzz_Nosy_newContractFromCLI__ because parameters include func, chan, or unsupported interface: func(ctx *github.com/urfave/cli/v2.Context) (github.com/ethereum/go-ethereum/common.Address, error)

func Fuzz_Nosy_parseRunArgs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var args []string
		fill_err = tp.Fill(&args)
		if fill_err != nil {
			return
		}

		parseRunArgs(args)
	})
}
