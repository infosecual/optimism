package withdrawals

import (
	"testing"

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

func Fuzz_Nosy_proofDB_Get__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var proof []string
		fill_err = tp.Fill(&proof)
		if fill_err != nil {
			return
		}
		var key []byte
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}

		p := GenerateProofDB(proof)
		p.Get(key)
	})
}

func Fuzz_Nosy_proofDB_Has__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var proof []string
		fill_err = tp.Fill(&proof)
		if fill_err != nil {
			return
		}
		var key []byte
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}

		p := GenerateProofDB(proof)
		p.Has(key)
	})
}

// skipping Fuzz_Nosy_BlockClient_BlockByNumber__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/withdrawals.BlockClient

// skipping Fuzz_Nosy_ProofClient_GetProof__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/withdrawals.ProofClient

// skipping Fuzz_Nosy_ReceiptClient_TransactionReceipt__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/withdrawals.ReceiptClient
