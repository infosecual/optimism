package config

import (
	"testing"

	"github.com/ethereum-optimism/optimism/op-node/rollup"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
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

func Fuzz_Nosy_Config_Check__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rollupCfg *rollup.Config
		fill_err = tp.Fill(&rollupCfg)
		if fill_err != nil {
			return
		}
		var l2Genesis *params.ChainConfig
		fill_err = tp.Fill(&l2Genesis)
		if fill_err != nil {
			return
		}
		var l1Head common.Hash
		fill_err = tp.Fill(&l1Head)
		if fill_err != nil {
			return
		}
		var l2Head common.Hash
		fill_err = tp.Fill(&l2Head)
		if fill_err != nil {
			return
		}
		var l2OutputRoot common.Hash
		fill_err = tp.Fill(&l2OutputRoot)
		if fill_err != nil {
			return
		}
		var l2Claim common.Hash
		fill_err = tp.Fill(&l2Claim)
		if fill_err != nil {
			return
		}
		var l2ClaimBlockNum uint64
		fill_err = tp.Fill(&l2ClaimBlockNum)
		if fill_err != nil {
			return
		}
		if rollupCfg == nil || l2Genesis == nil {
			return
		}

		c := NewConfig(rollupCfg, l2Genesis, l1Head, l2Head, l2OutputRoot, l2Claim, l2ClaimBlockNum)
		c.Check()
	})
}

func Fuzz_Nosy_Config_FetchingEnabled__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rollupCfg *rollup.Config
		fill_err = tp.Fill(&rollupCfg)
		if fill_err != nil {
			return
		}
		var l2Genesis *params.ChainConfig
		fill_err = tp.Fill(&l2Genesis)
		if fill_err != nil {
			return
		}
		var l1Head common.Hash
		fill_err = tp.Fill(&l1Head)
		if fill_err != nil {
			return
		}
		var l2Head common.Hash
		fill_err = tp.Fill(&l2Head)
		if fill_err != nil {
			return
		}
		var l2OutputRoot common.Hash
		fill_err = tp.Fill(&l2OutputRoot)
		if fill_err != nil {
			return
		}
		var l2Claim common.Hash
		fill_err = tp.Fill(&l2Claim)
		if fill_err != nil {
			return
		}
		var l2ClaimBlockNum uint64
		fill_err = tp.Fill(&l2ClaimBlockNum)
		if fill_err != nil {
			return
		}
		if rollupCfg == nil || l2Genesis == nil {
			return
		}

		c := NewConfig(rollupCfg, l2Genesis, l1Head, l2Head, l2OutputRoot, l2Claim, l2ClaimBlockNum)
		c.FetchingEnabled()
	})
}
