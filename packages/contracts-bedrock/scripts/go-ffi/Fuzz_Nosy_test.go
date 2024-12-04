package main

import (
	"math/big"
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

func Fuzz_Nosy_proofList_Delete__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *proofList
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var key []byte
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.Delete(key)
	})
}

func Fuzz_Nosy_proofList_Put__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *proofList
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var key []byte
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var value []byte
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.Put(key, value)
	})
}

func Fuzz_Nosy_trieTestCase_AbiEncode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, selectEmptyKey bool) {
		t1 := genTrieTestCase(selectEmptyKey)
		t1.AbiEncode()
	})
}

func Fuzz_Nosy_VerifyMerkleProof__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var root common.Hash
		fill_err = tp.Fill(&root)
		if fill_err != nil {
			return
		}
		var leaf common.Hash
		fill_err = tp.Fill(&leaf)
		if fill_err != nil {
			return
		}
		var index uint64
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		var proof [16]common.Hash
		fill_err = tp.Fill(&proof)
		if fill_err != nil {
			return
		}

		VerifyMerkleProof(root, leaf, index, proof)
	})
}

// skipping Fuzz_Nosy_checkErr__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_checkOk__(f *testing.F) {
	f.Fuzz(func(t *testing.T, ok bool) {
		checkOk(ok)
	})
}

func Fuzz_Nosy_encodeCrossDomainMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var nonce *big.Int
		fill_err = tp.Fill(&nonce)
		if fill_err != nil {
			return
		}
		var sender common.Address
		fill_err = tp.Fill(&sender)
		if fill_err != nil {
			return
		}
		var t3 common.Address
		fill_err = tp.Fill(&t3)
		if fill_err != nil {
			return
		}
		var value *big.Int
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
		var gasLimit *big.Int
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		var d6 []byte
		fill_err = tp.Fill(&d6)
		if fill_err != nil {
			return
		}
		if nonce == nil || value == nil || gasLimit == nil {
			return
		}

		encodeCrossDomainMessage(nonce, sender, t3, value, gasLimit, d6)
	})
}

func Fuzz_Nosy_randRange__(f *testing.F) {
	f.Fuzz(func(t *testing.T, min int64, max int64) {
		randRange(min, max)
	})
}
