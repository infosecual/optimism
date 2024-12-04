package helpers

import (
	"crypto/ecdsa"
	"testing"

	"github.com/ethereum-optimism/optimism/op-e2e/system/e2esys"
	"github.com/ethereum-optimism/optimism/op-node/withdrawals"
	"github.com/ethereum/go-ethereum/core/types"
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

func Fuzz_Nosy_TxOpts_VerifyOnClients__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var clients []*ethclient.Client
		fill_err = tp.Fill(&clients)
		if fill_err != nil {
			return
		}

		o := defaultTxOpts()
		o.VerifyOnClients(clients...)
	})
}

func Fuzz_Nosy_WithdrawalTxOpts_VerifyOnClients__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var clients []*ethclient.Client
		fill_err = tp.Fill(&clients)
		if fill_err != nil {
			return
		}

		o := defaultWithdrawalTxOpts()
		o.VerifyOnClients(clients...)
	})
}

// skipping Fuzz_Nosy_ClientProvider_NodeClient__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/system/helpers.ClientProvider

func Fuzz_Nosy_FinalizeWithdrawal__(f *testing.F) {
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
		var cfg e2esys.SystemConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var l1Client *ethclient.Client
		fill_err = tp.Fill(&l1Client)
		if fill_err != nil {
			return
		}
		var privKey *ecdsa.PrivateKey
		fill_err = tp.Fill(&privKey)
		if fill_err != nil {
			return
		}
		var withdrawalProofReceipt *types.Receipt
		fill_err = tp.Fill(&withdrawalProofReceipt)
		if fill_err != nil {
			return
		}
		var params withdrawals.ProvenWithdrawalParameters
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		if t1 == nil || l1Client == nil || privKey == nil || withdrawalProofReceipt == nil {
			return
		}

		FinalizeWithdrawal(t1, cfg, l1Client, privKey, withdrawalProofReceipt, params)
	})
}

// skipping Fuzz_Nosy_ProveAndFinalizeWithdrawal__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/system/helpers.ClientProvider

// skipping Fuzz_Nosy_ProveWithdrawal__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/system/helpers.ClientProvider

// skipping Fuzz_Nosy_SendWithdrawal__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/system/helpers.WithdrawalTxOptsFn
