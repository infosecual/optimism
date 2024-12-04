package singlethreaded

import (
	"io"
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

func Fuzz_Nosy_InstrumentedState_CheckInfiniteLoop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *InstrumentedState
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.CheckInfiniteLoop()
	})
}

func Fuzz_Nosy_InstrumentedState_GetDebugInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *InstrumentedState
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetDebugInfo()
	})
}

func Fuzz_Nosy_InstrumentedState_GetState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *InstrumentedState
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetState()
	})
}

func Fuzz_Nosy_InstrumentedState_InitDebug__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *InstrumentedState
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.InitDebug()
	})
}

func Fuzz_Nosy_InstrumentedState_LastPreimage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *InstrumentedState
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.LastPreimage()
	})
}

func Fuzz_Nosy_InstrumentedState_LookupSymbol__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *InstrumentedState
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var addr uint32
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.LookupSymbol(addr)
	})
}

func Fuzz_Nosy_InstrumentedState_Step__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *InstrumentedState
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var proof bool
		fill_err = tp.Fill(&proof)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Step(proof)
	})
}

func Fuzz_Nosy_InstrumentedState_Traceback__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *InstrumentedState
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Traceback()
	})
}

func Fuzz_Nosy_InstrumentedState_handleRMWOps__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *InstrumentedState
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var insn uint32
		fill_err = tp.Fill(&insn)
		if fill_err != nil {
			return
		}
		var opcode uint32
		fill_err = tp.Fill(&opcode)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.handleRMWOps(insn, opcode)
	})
}

func Fuzz_Nosy_InstrumentedState_handleSyscall__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *InstrumentedState
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.handleSyscall()
	})
}

func Fuzz_Nosy_InstrumentedState_mipsStep__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *InstrumentedState
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.mipsStep()
	})
}

// skipping Fuzz_Nosy_State_CreateVM__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/log.Logger

func Fuzz_Nosy_State_Deserialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var in io.Reader
		fill_err = tp.Fill(&in)
		if fill_err != nil {
			return
		}

		s := CreateEmptyState()
		s.Deserialize(in)
	})
}

func Fuzz_Nosy_State_Serialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var out io.Writer
		fill_err = tp.Fill(&out)
		if fill_err != nil {
			return
		}

		s := CreateEmptyState()
		s.Serialize(out)
	})
}

func Fuzz_Nosy_State_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d1 []byte) {
		s := CreateEmptyState()
		s.UnmarshalJSON(d1)
	})
}

func Fuzz_Nosy_StateWitness_StateHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sw StateWitness
		fill_err = tp.Fill(&sw)
		if fill_err != nil {
			return
		}

		sw.StateHash()
	})
}
