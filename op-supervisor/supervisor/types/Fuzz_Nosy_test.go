package types

import (
	"testing"

	"github.com/ethereum-optimism/optimism/op-service/eth"
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

func Fuzz_Nosy_ChainID_ToBig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, i uint64) {
		id := ChainIDFromUInt64(i)
		id.ToBig()
	})
}

func Fuzz_Nosy_ChainID_UnmarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, i uint64, d2 []byte) {
		id := ChainIDFromUInt64(i)
		id.UnmarshalText(d2)
	})
}

func Fuzz_Nosy_ChainIndex_UnmarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ci *ChainIndex
		fill_err = tp.Fill(&ci)
		if fill_err != nil {
			return
		}
		var d2 []byte
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}
		if ci == nil {
			return
		}

		ci.UnmarshalText(d2)
	})
}

func Fuzz_Nosy_ExecutingMessage_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *ExecutingMessage
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.String()
	})
}

func Fuzz_Nosy_Identifier_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id *Identifier
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		if id == nil {
			return
		}

		id.UnmarshalJSON(input)
	})
}

func Fuzz_Nosy_SafetyLevel_AtLeastAsSafe__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var lvl *SafetyLevel
		fill_err = tp.Fill(&lvl)
		if fill_err != nil {
			return
		}
		var min SafetyLevel
		fill_err = tp.Fill(&min)
		if fill_err != nil {
			return
		}
		if lvl == nil {
			return
		}

		lvl.AtLeastAsSafe(min)
	})
}

func Fuzz_Nosy_SafetyLevel_UnmarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var lvl *SafetyLevel
		fill_err = tp.Fill(&lvl)
		if fill_err != nil {
			return
		}
		var text []byte
		fill_err = tp.Fill(&text)
		if fill_err != nil {
			return
		}
		if lvl == nil {
			return
		}

		lvl.UnmarshalText(text)
	})
}

func Fuzz_Nosy_BlockSeal_ForceWithParent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ref eth.L1BlockRef
		fill_err = tp.Fill(&ref)
		if fill_err != nil {
			return
		}
		var parent eth.BlockID
		fill_err = tp.Fill(&parent)
		if fill_err != nil {
			return
		}

		s := BlockSealFromRef(ref)
		s.ForceWithParent(parent)
	})
}

func Fuzz_Nosy_BlockSeal_ID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ref eth.L1BlockRef
		fill_err = tp.Fill(&ref)
		if fill_err != nil {
			return
		}

		s := BlockSealFromRef(ref)
		s.ID()
	})
}

func Fuzz_Nosy_BlockSeal_MustWithParent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ref eth.L1BlockRef
		fill_err = tp.Fill(&ref)
		if fill_err != nil {
			return
		}
		var parent eth.BlockID
		fill_err = tp.Fill(&parent)
		if fill_err != nil {
			return
		}

		s := BlockSealFromRef(ref)
		s.MustWithParent(parent)
	})
}

func Fuzz_Nosy_BlockSeal_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ref eth.L1BlockRef
		fill_err = tp.Fill(&ref)
		if fill_err != nil {
			return
		}

		s := BlockSealFromRef(ref)
		s.String()
	})
}

func Fuzz_Nosy_BlockSeal_WithParent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ref eth.L1BlockRef
		fill_err = tp.Fill(&ref)
		if fill_err != nil {
			return
		}
		var parent eth.BlockID
		fill_err = tp.Fill(&parent)
		if fill_err != nil {
			return
		}

		s := BlockSealFromRef(ref)
		s.WithParent(parent)
	})
}

func Fuzz_Nosy_ChainID_Cmp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var i uint64
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
		var other ChainID
		fill_err = tp.Fill(&other)
		if fill_err != nil {
			return
		}

		id := ChainIDFromUInt64(i)
		id.Cmp(other)
	})
}

func Fuzz_Nosy_ChainID_MarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, i uint64) {
		id := ChainIDFromUInt64(i)
		id.MarshalText()
	})
}

func Fuzz_Nosy_ChainID_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, i uint64) {
		id := ChainIDFromUInt64(i)
		id.String()
	})
}

func Fuzz_Nosy_ChainID_ToUInt32__(f *testing.F) {
	f.Fuzz(func(t *testing.T, i uint64) {
		id := ChainIDFromUInt64(i)
		id.ToUInt32()
	})
}

func Fuzz_Nosy_ChainIndex_MarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ci ChainIndex
		fill_err = tp.Fill(&ci)
		if fill_err != nil {
			return
		}

		ci.MarshalText()
	})
}

func Fuzz_Nosy_ChainIndex_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ci ChainIndex
		fill_err = tp.Fill(&ci)
		if fill_err != nil {
			return
		}

		ci.String()
	})
}

func Fuzz_Nosy_Identifier_MarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id Identifier
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}

		id.MarshalJSON()
	})
}

func Fuzz_Nosy_ReferenceView_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v ReferenceView
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}

		v.String()
	})
}

func Fuzz_Nosy_SafetyLevel_MarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var lvl SafetyLevel
		fill_err = tp.Fill(&lvl)
		if fill_err != nil {
			return
		}

		lvl.MarshalText()
	})
}

func Fuzz_Nosy_SafetyLevel_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var lvl SafetyLevel
		fill_err = tp.Fill(&lvl)
		if fill_err != nil {
			return
		}

		lvl.String()
	})
}

func Fuzz_Nosy_SafetyLevel_Valid__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var lvl SafetyLevel
		fill_err = tp.Fill(&lvl)
		if fill_err != nil {
			return
		}

		lvl.Valid()
	})
}

func Fuzz_Nosy_LogToMessagePayload__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *types.Log
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		LogToMessagePayload(l)
	})
}
