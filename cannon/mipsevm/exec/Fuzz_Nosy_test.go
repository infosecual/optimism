package exec

import (
	"testing"

	"github.com/ethereum-optimism/optimism/cannon/mipsevm"
	"github.com/ethereum-optimism/optimism/cannon/mipsevm/memory"
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

func Fuzz_Nosy_MemoryTrackerImpl_MemProof__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var memory *memory.Memory
		fill_err = tp.Fill(&memory)
		if fill_err != nil {
			return
		}
		if memory == nil {
			return
		}

		m := NewMemoryTracker(memory)
		m.MemProof()
	})
}

func Fuzz_Nosy_MemoryTrackerImpl_MemProof2__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var memory *memory.Memory
		fill_err = tp.Fill(&memory)
		if fill_err != nil {
			return
		}
		if memory == nil {
			return
		}

		m := NewMemoryTracker(memory)
		m.MemProof2()
	})
}

func Fuzz_Nosy_MemoryTrackerImpl_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var memory *memory.Memory
		fill_err = tp.Fill(&memory)
		if fill_err != nil {
			return
		}
		var enableProof bool
		fill_err = tp.Fill(&enableProof)
		if fill_err != nil {
			return
		}
		if memory == nil {
			return
		}

		m := NewMemoryTracker(memory)
		m.Reset(enableProof)
	})
}

func Fuzz_Nosy_MemoryTrackerImpl_TrackMemAccess__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var memory *memory.Memory
		fill_err = tp.Fill(&memory)
		if fill_err != nil {
			return
		}
		var effAddr uint32
		fill_err = tp.Fill(&effAddr)
		if fill_err != nil {
			return
		}
		if memory == nil {
			return
		}

		m := NewMemoryTracker(memory)
		m.TrackMemAccess(effAddr)
	})
}

func Fuzz_Nosy_MemoryTrackerImpl_TrackMemAccess2__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var memory *memory.Memory
		fill_err = tp.Fill(&memory)
		if fill_err != nil {
			return
		}
		var effAddr uint32
		fill_err = tp.Fill(&effAddr)
		if fill_err != nil {
			return
		}
		if memory == nil {
			return
		}

		m := NewMemoryTracker(memory)
		m.TrackMemAccess2(effAddr)
	})
}

func Fuzz_Nosy_NoopMemoryTracker_TrackMemAccess__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *NoopMemoryTracker
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var _x2 uint32
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.TrackMemAccess(_x2)
	})
}

func Fuzz_Nosy_NoopStackTracker_PopStack__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *NoopStackTracker
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.PopStack()
	})
}

func Fuzz_Nosy_NoopStackTracker_PushStack__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *NoopStackTracker
		fill_err = tp.Fill(&n)
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
		if n == nil {
			return
		}

		n.PushStack(caller, t3)
	})
}

func Fuzz_Nosy_NoopStackTracker_Traceback__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *NoopStackTracker
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.Traceback()
	})
}

func Fuzz_Nosy_StackTrackerImpl_PopStack__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *StackTrackerImpl
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.PopStack()
	})
}

func Fuzz_Nosy_StackTrackerImpl_PushStack__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *StackTrackerImpl
		fill_err = tp.Fill(&s)
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
		if s == nil {
			return
		}

		s.PushStack(caller, t3)
	})
}

func Fuzz_Nosy_StackTrackerImpl_Traceback__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *StackTrackerImpl
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Traceback()
	})
}

func Fuzz_Nosy_TrackingPreimageOracleReader_GetPreimage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *TrackingPreimageOracleReader
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var k [32]byte
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.GetPreimage(k)
	})
}

func Fuzz_Nosy_TrackingPreimageOracleReader_Hint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *TrackingPreimageOracleReader
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var v []byte
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.Hint(v)
	})
}

func Fuzz_Nosy_TrackingPreimageOracleReader_LastPreimage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *TrackingPreimageOracleReader
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.LastPreimage()
	})
}

func Fuzz_Nosy_TrackingPreimageOracleReader_NumPreimageRequests__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *TrackingPreimageOracleReader
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.NumPreimageRequests()
	})
}

func Fuzz_Nosy_TrackingPreimageOracleReader_ReadPreimage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *TrackingPreimageOracleReader
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var key [32]byte
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var offset uint32
		fill_err = tp.Fill(&offset)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.ReadPreimage(key, offset)
	})
}

func Fuzz_Nosy_TrackingPreimageOracleReader_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *TrackingPreimageOracleReader
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.Reset()
	})
}

func Fuzz_Nosy_TrackingPreimageOracleReader_TotalPreimageSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *TrackingPreimageOracleReader
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.TotalPreimageSize()
	})
}

// skipping Fuzz_Nosy_MemTracker_TrackMemAccess__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/cannon/mipsevm/exec.MemTracker

// skipping Fuzz_Nosy_PreimageReader_ReadPreimage__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/cannon/mipsevm/exec.PreimageReader

// skipping Fuzz_Nosy_StackTracker_PopStack__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/cannon/mipsevm/exec.StackTracker

// skipping Fuzz_Nosy_StackTracker_PushStack__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/cannon/mipsevm/exec.StackTracker

// skipping Fuzz_Nosy_TraceableStackTracker_Traceback__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/cannon/mipsevm/exec.TraceableStackTracker

// skipping Fuzz_Nosy_ExecMipsCoreStepLogic__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/cannon/mipsevm/exec.MemTracker

func Fuzz_Nosy_ExecuteMipsInstruction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, insn uint32, opcode uint32, fun uint32, rs uint32, rt uint32, mem uint32) {
		ExecuteMipsInstruction(insn, opcode, fun, rs, rt, mem)
	})
}

func Fuzz_Nosy_GetInstructionDetails__(f *testing.F) {
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
		var memory *memory.Memory
		fill_err = tp.Fill(&memory)
		if fill_err != nil {
			return
		}
		if memory == nil {
			return
		}

		GetInstructionDetails(pc, memory)
	})
}

func Fuzz_Nosy_GetSyscallArgs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var registers *[32]uint32
		fill_err = tp.Fill(&registers)
		if fill_err != nil {
			return
		}
		if registers == nil {
			return
		}

		GetSyscallArgs(registers)
	})
}

func Fuzz_Nosy_HandleSysFcntl__(f *testing.F) {
	f.Fuzz(func(t *testing.T, a0 uint32, a1 uint32) {
		HandleSysFcntl(a0, a1)
	})
}

func Fuzz_Nosy_HandleSysMmap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, a0 uint32, a1 uint32, heap uint32) {
		HandleSysMmap(a0, a1, heap)
	})
}

// skipping Fuzz_Nosy_HandleSysRead__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/cannon/mipsevm/exec.PreimageReader

// skipping Fuzz_Nosy_HandleSysWrite__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/cannon/mipsevm.PreimageOracle

func Fuzz_Nosy_HandleSyscallUpdates__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cpu *mipsevm.CpuScalars
		fill_err = tp.Fill(&cpu)
		if fill_err != nil {
			return
		}
		var registers *[32]uint32
		fill_err = tp.Fill(&registers)
		if fill_err != nil {
			return
		}
		var v0 uint32
		fill_err = tp.Fill(&v0)
		if fill_err != nil {
			return
		}
		var v1 uint32
		fill_err = tp.Fill(&v1)
		if fill_err != nil {
			return
		}
		if cpu == nil || registers == nil {
			return
		}

		HandleSyscallUpdates(cpu, registers, v0, v1)
	})
}

// skipping Fuzz_Nosy_LoadSubWord__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/cannon/mipsevm/exec.MemTracker

func Fuzz_Nosy_SelectSubWord__(f *testing.F) {
	f.Fuzz(func(t *testing.T, vaddr uint32, memWord uint32, byteLength uint32, signExtend bool) {
		SelectSubWord(vaddr, memWord, byteLength, signExtend)
	})
}

func Fuzz_Nosy_SignExtend__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dat uint32, idx uint32) {
		SignExtend(dat, idx)
	})
}

func Fuzz_Nosy_SignExtendImmediate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, insn uint32) {
		SignExtendImmediate(insn)
	})
}

// skipping Fuzz_Nosy_StoreSubWord__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/cannon/mipsevm/exec.MemTracker

func Fuzz_Nosy_UpdateSubWord__(f *testing.F) {
	f.Fuzz(func(t *testing.T, vaddr uint32, memWord uint32, byteLength uint32, value uint32) {
		UpdateSubWord(vaddr, memWord, byteLength, value)
	})
}

func Fuzz_Nosy_abs64__(f *testing.F) {
	f.Fuzz(func(t *testing.T, x int64) {
		abs64(x)
	})
}

func Fuzz_Nosy_assertMips64__(f *testing.F) {
	f.Fuzz(func(t *testing.T, insn uint32) {
		assertMips64(insn)
	})
}

func Fuzz_Nosy_assertMips64Fun__(f *testing.F) {
	f.Fuzz(func(t *testing.T, fun uint32) {
		assertMips64Fun(fun)
	})
}

func Fuzz_Nosy_calculateSubWordMaskAndOffset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, vaddr uint32, byteLength uint32) {
		calculateSubWordMaskAndOffset(vaddr, byteLength)
	})
}
