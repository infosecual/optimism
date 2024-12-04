package batcher

import (
	"context"
	"crypto/ecdsa"
	"testing"

	"github.com/ethereum-optimism/optimism/op-node/rollup"
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

func Fuzz_Nosy_Helper_SendLargeInvalidBatch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var privKey *ecdsa.PrivateKey
		fill_err = tp.Fill(&privKey)
		if fill_err != nil {
			return
		}
		var rollupCfg *rollup.Config
		fill_err = tp.Fill(&rollupCfg)
		if fill_err != nil {
			return
		}
		var l1Client *ethclient.Client
		fill_err = tp.Fill(&l1Client)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if t1 == nil || privKey == nil || rollupCfg == nil || l1Client == nil {
			return
		}

		h := NewHelper(t1, privKey, rollupCfg, l1Client)
		h.SendLargeInvalidBatch(ctx)
	})
}
