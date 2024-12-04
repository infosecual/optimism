package multithreaded

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

func Fuzz_Nosy_InstrumentedState_assertPostStateChecks__(f *testing.F) {
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

		m.assertPostStateChecks()
	})
}

func Fuzz_Nosy_InstrumentedState_clearLLMemoryReservation__(f *testing.F) {
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

		m.clearLLMemoryReservation()
	})
}

func Fuzz_Nosy_InstrumentedState_doMipsStep__(f *testing.F) {
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

		m.doMipsStep()
	})
}

func Fuzz_Nosy_InstrumentedState_handleMemoryUpdate__(f *testing.F) {
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
		var effMemAddr uint32
		fill_err = tp.Fill(&effMemAddr)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.handleMemoryUpdate(effMemAddr)
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

func Fuzz_Nosy_InstrumentedState_lastThreadRemaining__(f *testing.F) {
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

		m.lastThreadRemaining()
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

func Fuzz_Nosy_InstrumentedState_onWaitComplete__(f *testing.F) {
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
		var thread *ThreadState
		fill_err = tp.Fill(&thread)
		if fill_err != nil {
			return
		}
		var isTimedOut bool
		fill_err = tp.Fill(&isTimedOut)
		if fill_err != nil {
			return
		}
		if m == nil || thread == nil {
			return
		}

		m.onWaitComplete(thread, isTimedOut)
	})
}

func Fuzz_Nosy_InstrumentedState_popThread__(f *testing.F) {
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

		m.popThread()
	})
}

func Fuzz_Nosy_InstrumentedState_preemptThread__(f *testing.F) {
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
		var thread *ThreadState
		fill_err = tp.Fill(&thread)
		if fill_err != nil {
			return
		}
		if m == nil || thread == nil {
			return
		}

		m.preemptThread(thread)
	})
}

func Fuzz_Nosy_InstrumentedState_pushThread__(f *testing.F) {
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
		var thread *ThreadState
		fill_err = tp.Fill(&thread)
		if fill_err != nil {
			return
		}
		if m == nil || thread == nil {
			return
		}

		m.pushThread(thread)
	})
}

func Fuzz_Nosy_NoopThreadedStackTracker_DropThread__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *NoopThreadedStackTracker
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var threadId uint32
		fill_err = tp.Fill(&threadId)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.DropThread(threadId)
	})
}

// skipping Fuzz_Nosy_State_CreateVM__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/log.Logger

func Fuzz_Nosy_State_Deserialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc uint32
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var heapStart uint32
		fill_err = tp.Fill(&heapStart)
		if fill_err != nil {
			return
		}
		var in io.Reader
		fill_err = tp.Fill(&in)
		if fill_err != nil {
			return
		}

		s := CreateInitialState(pc, heapStart)
		s.Deserialize(in)
	})
}

func Fuzz_Nosy_State_EncodeThreadProof__(f *testing.F) {
	f.Fuzz(func(t *testing.T, pc uint32, heapStart uint32) {
		s := CreateInitialState(pc, heapStart)
		s.EncodeThreadProof()
	})
}

func Fuzz_Nosy_State_EncodeWitness__(f *testing.F) {
	f.Fuzz(func(t *testing.T, pc uint32, heapStart uint32) {
		s := CreateInitialState(pc, heapStart)
		s.EncodeWitness()
	})
}

func Fuzz_Nosy_State_GetCpu__(f *testing.F) {
	f.Fuzz(func(t *testing.T, pc uint32, heapStart uint32) {
		s := CreateInitialState(pc, heapStart)
		s.GetCpu()
	})
}

func Fuzz_Nosy_State_GetCurrentThread__(f *testing.F) {
	f.Fuzz(func(t *testing.T, pc uint32, heapStart uint32) {
		s := CreateInitialState(pc, heapStart)
		s.GetCurrentThread()
	})
}

func Fuzz_Nosy_State_GetExitCode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, pc uint32, heapStart uint32) {
		s := CreateInitialState(pc, heapStart)
		s.GetExitCode()
	})
}

func Fuzz_Nosy_State_GetExited__(f *testing.F) {
	f.Fuzz(func(t *testing.T, pc uint32, heapStart uint32) {
		s := CreateInitialState(pc, heapStart)
		s.GetExited()
	})
}

func Fuzz_Nosy_State_GetHeap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, pc uint32, heapStart uint32) {
		s := CreateInitialState(pc, heapStart)
		s.GetHeap()
	})
}

func Fuzz_Nosy_State_GetLastHint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, pc uint32, heapStart uint32) {
		s := CreateInitialState(pc, heapStart)
		s.GetLastHint()
	})
}

func Fuzz_Nosy_State_GetMemory__(f *testing.F) {
	f.Fuzz(func(t *testing.T, pc uint32, heapStart uint32) {
		s := CreateInitialState(pc, heapStart)
		s.GetMemory()
	})
}

func Fuzz_Nosy_State_GetPC__(f *testing.F) {
	f.Fuzz(func(t *testing.T, pc uint32, heapStart uint32) {
		s := CreateInitialState(pc, heapStart)
		s.GetPC()
	})
}

func Fuzz_Nosy_State_GetPreimageKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, pc uint32, heapStart uint32) {
		s := CreateInitialState(pc, heapStart)
		s.GetPreimageKey()
	})
}

func Fuzz_Nosy_State_GetPreimageOffset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, pc uint32, heapStart uint32) {
		s := CreateInitialState(pc, heapStart)
		s.GetPreimageOffset()
	})
}

func Fuzz_Nosy_State_GetRegistersRef__(f *testing.F) {
	f.Fuzz(func(t *testing.T, pc uint32, heapStart uint32) {
		s := CreateInitialState(pc, heapStart)
		s.GetRegistersRef()
	})
}

func Fuzz_Nosy_State_GetStep__(f *testing.F) {
	f.Fuzz(func(t *testing.T, pc uint32, heapStart uint32) {
		s := CreateInitialState(pc, heapStart)
		s.GetStep()
	})
}

func Fuzz_Nosy_State_Serialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc uint32
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var heapStart uint32
		fill_err = tp.Fill(&heapStart)
		if fill_err != nil {
			return
		}
		var out io.Writer
		fill_err = tp.Fill(&out)
		if fill_err != nil {
			return
		}

		s := CreateInitialState(pc, heapStart)
		s.Serialize(out)
	})
}

func Fuzz_Nosy_State_ThreadCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, pc uint32, heapStart uint32) {
		s := CreateInitialState(pc, heapStart)
		s.ThreadCount()
	})
}

func Fuzz_Nosy_State_VMStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, pc uint32, heapStart uint32) {
		s := CreateInitialState(pc, heapStart)
		s.VMStatus()
	})
}

func Fuzz_Nosy_State_calculateThreadStackRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc uint32
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var heapStart uint32
		fill_err = tp.Fill(&heapStart)
		if fill_err != nil {
			return
		}
		var stack []*ThreadState
		fill_err = tp.Fill(&stack)
		if fill_err != nil {
			return
		}

		s := CreateInitialState(pc, heapStart)
		s.calculateThreadStackRoot(stack)
	})
}

func Fuzz_Nosy_State_getActiveThreadStack__(f *testing.F) {
	f.Fuzz(func(t *testing.T, pc uint32, heapStart uint32) {
		s := CreateInitialState(pc, heapStart)
		s.getActiveThreadStack()
	})
}

func Fuzz_Nosy_State_getCpuRef__(f *testing.F) {
	f.Fuzz(func(t *testing.T, pc uint32, heapStart uint32) {
		s := CreateInitialState(pc, heapStart)
		s.getCpuRef()
	})
}

func Fuzz_Nosy_State_getLeftThreadStackRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, pc uint32, heapStart uint32) {
		s := CreateInitialState(pc, heapStart)
		s.getLeftThreadStackRoot()
	})
}

func Fuzz_Nosy_State_getRightThreadStackRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, pc uint32, heapStart uint32) {
		s := CreateInitialState(pc, heapStart)
		s.getRightThreadStackRoot()
	})
}

func Fuzz_Nosy_ThreadState_Deserialize__(f *testing.F) {
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

		t1 := CreateEmptyThread()
		t1.Deserialize(in)
	})
}

func Fuzz_Nosy_ThreadState_Serialize__(f *testing.F) {
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

		t1 := CreateEmptyThread()
		t1.Serialize(out)
	})
}

func Fuzz_Nosy_ThreadedStackTrackerImpl_DropThread__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *ThreadedStackTrackerImpl
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var threadId uint32
		fill_err = tp.Fill(&threadId)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.DropThread(threadId)
	})
}

func Fuzz_Nosy_ThreadedStackTrackerImpl_PopStack__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *ThreadedStackTrackerImpl
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.PopStack()
	})
}

func Fuzz_Nosy_ThreadedStackTrackerImpl_PushStack__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *ThreadedStackTrackerImpl
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var caller uint32
		fill_err = tp.Fill(&caller)
		if fill_err != nil {
			return
		}
		var t3 uint32
		fill_err = tp.Fill(&t3)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.PushStack(caller, t3)
	})
}

func Fuzz_Nosy_ThreadedStackTrackerImpl_Traceback__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *ThreadedStackTrackerImpl
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Traceback()
	})
}

func Fuzz_Nosy_ThreadedStackTrackerImpl_getCurrentTracker__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *ThreadedStackTrackerImpl
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.getCurrentTracker()
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

// skipping Fuzz_Nosy_ThreadedStackTracker_DropThread__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/cannon/mipsevm/multithreaded.ThreadedStackTracker
