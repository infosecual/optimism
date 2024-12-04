package transactions

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
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

// skipping Fuzz_Nosy_RequireSendTx__ because parameters include func, chan, or unsupported interface: []github.com/ethereum-optimism/optimism/op-e2e/e2eutils/transactions.SendTxOpt

// skipping Fuzz_Nosy_SendTx__ because parameters include func, chan, or unsupported interface: []github.com/ethereum-optimism/optimism/op-e2e/e2eutils/transactions.SendTxOpt

func Fuzz_Nosy_TransactionsBySender__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var block *types.Block
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		var sender common.Address
		fill_err = tp.Fill(&sender)
		if fill_err != nil {
			return
		}
		if block == nil {
			return
		}

		TransactionsBySender(block, sender)
	})
}
