package testutil

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
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

func Fuzz_Nosy_StateMutatorSingleThreaded_Randomize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *StateMutatorSingleThreaded
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var randSeed int64
		fill_err = tp.Fill(&randSeed)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Randomize(randSeed)
	})
}

func Fuzz_Nosy_StateMutatorSingleThreaded_SetExitCode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *StateMutatorSingleThreaded
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var val uint8
		fill_err = tp.Fill(&val)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.SetExitCode(val)
	})
}

func Fuzz_Nosy_StateMutatorSingleThreaded_SetExited__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *StateMutatorSingleThreaded
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var val bool
		fill_err = tp.Fill(&val)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.SetExited(val)
	})
}

func Fuzz_Nosy_StateMutatorSingleThreaded_SetHI__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *StateMutatorSingleThreaded
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var val uint32
		fill_err = tp.Fill(&val)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.SetHI(val)
	})
}

func Fuzz_Nosy_StateMutatorSingleThreaded_SetHeap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *StateMutatorSingleThreaded
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var val uint32
		fill_err = tp.Fill(&val)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.SetHeap(val)
	})
}

func Fuzz_Nosy_StateMutatorSingleThreaded_SetLO__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *StateMutatorSingleThreaded
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var val uint32
		fill_err = tp.Fill(&val)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.SetLO(val)
	})
}

func Fuzz_Nosy_StateMutatorSingleThreaded_SetLastHint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *StateMutatorSingleThreaded
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var val hexutil.Bytes
		fill_err = tp.Fill(&val)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.SetLastHint(val)
	})
}

func Fuzz_Nosy_StateMutatorSingleThreaded_SetNextPC__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *StateMutatorSingleThreaded
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var val uint32
		fill_err = tp.Fill(&val)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.SetNextPC(val)
	})
}

func Fuzz_Nosy_StateMutatorSingleThreaded_SetPC__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *StateMutatorSingleThreaded
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var val uint32
		fill_err = tp.Fill(&val)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.SetPC(val)
	})
}

func Fuzz_Nosy_StateMutatorSingleThreaded_SetPreimageKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *StateMutatorSingleThreaded
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var val common.Hash
		fill_err = tp.Fill(&val)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.SetPreimageKey(val)
	})
}

func Fuzz_Nosy_StateMutatorSingleThreaded_SetPreimageOffset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *StateMutatorSingleThreaded
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var val uint32
		fill_err = tp.Fill(&val)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.SetPreimageOffset(val)
	})
}

func Fuzz_Nosy_StateMutatorSingleThreaded_SetStep__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *StateMutatorSingleThreaded
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var val uint64
		fill_err = tp.Fill(&val)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.SetStep(val)
	})
}
