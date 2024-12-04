package inspect

import (
	"testing"

	"github.com/ethereum-optimism/optimism/op-chain-ops/script"
	"github.com/ethereum-optimism/optimism/op-deployer/pkg/deployer/state"
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

func Fuzz_Nosy_GenesisAndRollup__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var globalState *state.State
		fill_err = tp.Fill(&globalState)
		if fill_err != nil {
			return
		}
		var chainID common.Hash
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		if globalState == nil {
			return
		}

		GenesisAndRollup(globalState, chainID)
	})
}

func Fuzz_Nosy_ReadSemver__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var host *script.Host
		fill_err = tp.Fill(&host)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if host == nil {
			return
		}

		ReadSemver(host, addr)
	})
}

func Fuzz_Nosy_findSemverBytecode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var host *script.Host
		fill_err = tp.Fill(&host)
		if fill_err != nil {
			return
		}
		var proxyAddr common.Address
		fill_err = tp.Fill(&proxyAddr)
		if fill_err != nil {
			return
		}
		if host == nil {
			return
		}

		findSemverBytecode(host, proxyAddr)
	})
}
