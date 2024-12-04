package testutil

import (
	"testing"

	"github.com/ethereum-optimism/optimism/cannon/mipsevm/multithreaded"
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

func Fuzz_Nosy_ExpectedMTState_ActiveThread__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fromState *multithreaded.State
		fill_err = tp.Fill(&fromState)
		if fill_err != nil {
			return
		}
		if fromState == nil {
			return
		}

		e := NewExpectedMTState(fromState)
		e.ActiveThread()
	})
}

func Fuzz_Nosy_ExpectedMTState_ExpectMemoryWordWrite__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fromState *multithreaded.State
		fill_err = tp.Fill(&fromState)
		if fill_err != nil {
			return
		}
		var addr uint32
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var val uint32
		fill_err = tp.Fill(&val)
		if fill_err != nil {
			return
		}
		if fromState == nil {
			return
		}

		e := NewExpectedMTState(fromState)
		e.ExpectMemoryWordWrite(addr, val)
	})
}

// skipping Fuzz_Nosy_ExpectedMTState_ExpectMemoryWriteUint32__ because parameters include func, chan, or unsupported interface: github.com/stretchr/testify/require.TestingT

func Fuzz_Nosy_ExpectedMTState_ExpectNewThread__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fromState *multithreaded.State
		fill_err = tp.Fill(&fromState)
		if fill_err != nil {
			return
		}
		if fromState == nil {
			return
		}

		e := NewExpectedMTState(fromState)
		e.ExpectNewThread()
	})
}

func Fuzz_Nosy_ExpectedMTState_ExpectPreemption__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fromState *multithreaded.State
		fill_err = tp.Fill(&fromState)
		if fill_err != nil {
			return
		}
		var preState *multithreaded.State
		fill_err = tp.Fill(&preState)
		if fill_err != nil {
			return
		}
		if fromState == nil || preState == nil {
			return
		}

		e := NewExpectedMTState(fromState)
		e.ExpectPreemption(preState)
	})
}

func Fuzz_Nosy_ExpectedMTState_ExpectStep__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fromState *multithreaded.State
		fill_err = tp.Fill(&fromState)
		if fill_err != nil {
			return
		}
		if fromState == nil {
			return
		}

		e := NewExpectedMTState(fromState)
		e.ExpectStep()
	})
}

func Fuzz_Nosy_ExpectedMTState_PrestateActiveThread__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fromState *multithreaded.State
		fill_err = tp.Fill(&fromState)
		if fill_err != nil {
			return
		}
		if fromState == nil {
			return
		}

		e := NewExpectedMTState(fromState)
		e.PrestateActiveThread()
	})
}

func Fuzz_Nosy_ExpectedMTState_Thread__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fromState *multithreaded.State
		fill_err = tp.Fill(&fromState)
		if fill_err != nil {
			return
		}
		var threadId uint32
		fill_err = tp.Fill(&threadId)
		if fill_err != nil {
			return
		}
		if fromState == nil {
			return
		}

		e := NewExpectedMTState(fromState)
		e.Thread(threadId)
	})
}

// skipping Fuzz_Nosy_ExpectedMTState_Validate__ because parameters include func, chan, or unsupported interface: github.com/stretchr/testify/require.TestingT

// skipping Fuzz_Nosy_ExpectedMTState_validateThread__ because parameters include func, chan, or unsupported interface: github.com/stretchr/testify/require.TestingT

func Fuzz_Nosy_StateMutatorMultiThreaded_Randomize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *StateMutatorMultiThreaded
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

func Fuzz_Nosy_StateMutatorMultiThreaded_SetExitCode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *StateMutatorMultiThreaded
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

func Fuzz_Nosy_StateMutatorMultiThreaded_SetExited__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *StateMutatorMultiThreaded
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

func Fuzz_Nosy_StateMutatorMultiThreaded_SetHI__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *StateMutatorMultiThreaded
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

func Fuzz_Nosy_StateMutatorMultiThreaded_SetHeap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *StateMutatorMultiThreaded
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

func Fuzz_Nosy_StateMutatorMultiThreaded_SetLO__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *StateMutatorMultiThreaded
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

func Fuzz_Nosy_StateMutatorMultiThreaded_SetLastHint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *StateMutatorMultiThreaded
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

func Fuzz_Nosy_StateMutatorMultiThreaded_SetNextPC__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *StateMutatorMultiThreaded
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

func Fuzz_Nosy_StateMutatorMultiThreaded_SetPC__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *StateMutatorMultiThreaded
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

func Fuzz_Nosy_StateMutatorMultiThreaded_SetPreimageKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *StateMutatorMultiThreaded
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

func Fuzz_Nosy_StateMutatorMultiThreaded_SetPreimageOffset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *StateMutatorMultiThreaded
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

func Fuzz_Nosy_StateMutatorMultiThreaded_SetStep__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *StateMutatorMultiThreaded
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

func Fuzz_Nosy_ThreadIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var state *multithreaded.State
		fill_err = tp.Fill(&state)
		if fill_err != nil {
			return
		}
		if state == nil {
			return
		}

		i := NewThreadIterator(state)
		i.Next()
	})
}

func Fuzz_Nosy_ThreadIterator_currentThread__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var state *multithreaded.State
		fill_err = tp.Fill(&state)
		if fill_err != nil {
			return
		}
		if state == nil {
			return
		}

		i := NewThreadIterator(state)
		i.currentThread()
	})
}

func Fuzz_Nosy_GetAllThreads__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var state *multithreaded.State
		fill_err = tp.Fill(&state)
		if fill_err != nil {
			return
		}
		if state == nil {
			return
		}

		GetAllThreads(state)
	})
}

func Fuzz_Nosy_GetThreadStacks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var state *multithreaded.State
		fill_err = tp.Fill(&state)
		if fill_err != nil {
			return
		}
		if state == nil {
			return
		}

		GetThreadStacks(state)
	})
}

// skipping Fuzz_Nosy_InitializeSingleThread__ because parameters include func, chan, or unsupported interface: []github.com/ethereum-optimism/optimism/cannon/mipsevm/testutil.StateOption

func Fuzz_Nosy_SetupThreads__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var randomSeed int64
		fill_err = tp.Fill(&randomSeed)
		if fill_err != nil {
			return
		}
		var state *multithreaded.State
		fill_err = tp.Fill(&state)
		if fill_err != nil {
			return
		}
		var traverseRight bool
		fill_err = tp.Fill(&traverseRight)
		if fill_err != nil {
			return
		}
		var activeStackSize int
		fill_err = tp.Fill(&activeStackSize)
		if fill_err != nil {
			return
		}
		var otherStackSize int
		fill_err = tp.Fill(&otherStackSize)
		if fill_err != nil {
			return
		}
		if state == nil {
			return
		}

		SetupThreads(randomSeed, state, traverseRight, activeStackSize, otherStackSize)
	})
}
