package wait

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
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

func Fuzz_Nosy_ReceiptStatusError_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rse *ReceiptStatusError
		fill_err = tp.Fill(&rse)
		if fill_err != nil {
			return
		}
		if rse == nil {
			return
		}

		rse.Error()
	})
}

// skipping Fuzz_Nosy_BlockCaller_BlockByNumber__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/e2eutils/wait.BlockCaller

// skipping Fuzz_Nosy_BlockCaller_BlockNumber__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/e2eutils/wait.BlockCaller

// skipping Fuzz_Nosy_AndGet__ because parameters include func, chan, or unsupported interface: func() (T, error)

func Fuzz_Nosy_ForGamePublished__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var client *ethclient.Client
		fill_err = tp.Fill(&client)
		if fill_err != nil {
			return
		}
		var optimismPortalAddr common.Address
		fill_err = tp.Fill(&optimismPortalAddr)
		if fill_err != nil {
			return
		}
		var disputeGameFactoryAddr common.Address
		fill_err = tp.Fill(&disputeGameFactoryAddr)
		if fill_err != nil {
			return
		}
		var l2BlockNumber *big.Int
		fill_err = tp.Fill(&l2BlockNumber)
		if fill_err != nil {
			return
		}
		if client == nil || l2BlockNumber == nil {
			return
		}

		ForGamePublished(ctx, client, optimismPortalAddr, disputeGameFactoryAddr, l2BlockNumber)
	})
}

func Fuzz_Nosy_ForOutputRootPublished__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var client *ethclient.Client
		fill_err = tp.Fill(&client)
		if fill_err != nil {
			return
		}
		var l2OutputOracleAddr common.Address
		fill_err = tp.Fill(&l2OutputOracleAddr)
		if fill_err != nil {
			return
		}
		var l2BlockNumber *big.Int
		fill_err = tp.Fill(&l2BlockNumber)
		if fill_err != nil {
			return
		}
		if client == nil || l2BlockNumber == nil {
			return
		}

		ForOutputRootPublished(ctx, client, l2OutputOracleAddr, l2BlockNumber)
	})
}
