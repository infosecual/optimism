package merkle

import (
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

func Fuzz_Nosy_BinaryMerkleTree_AddLeaf__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}

		m := NewBinaryMerkleTree()
		m.AddLeaf(hash)
	})
}

func Fuzz_Nosy_BinaryMerkleTree_ProofAtIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, index uint64) {
		m := NewBinaryMerkleTree()
		m.ProofAtIndex(index)
	})
}

func Fuzz_Nosy_BinaryMerkleTree_walkDownToLeafCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, subtreeLeafCount uint64) {
		m := NewBinaryMerkleTree()
		m.walkDownToLeafCount(subtreeLeafCount)
	})
}

func Fuzz_Nosy_merkleNode_IsLeftChild__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *merkleNode
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var o *merkleNode
		fill_err = tp.Fill(&o)
		if fill_err != nil {
			return
		}
		if m == nil || o == nil {
			return
		}

		m.IsLeftChild(o)
	})
}

func Fuzz_Nosy_merkleNode_IsRightChild__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *merkleNode
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var o *merkleNode
		fill_err = tp.Fill(&o)
		if fill_err != nil {
			return
		}
		if m == nil || o == nil {
			return
		}

		m.IsRightChild(o)
	})
}
