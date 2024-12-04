package testutil

import (
	"debug/elf"
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

func Fuzz_Nosy_TrackableReaderAt_ReadAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *TrackableReaderAt
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var p []byte
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var offset int64
		fill_err = tp.Fill(&offset)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.ReadAt(p, offset)
	})
}

// skipping Fuzz_Nosy_MockFPVMState_CreateVM__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/log.Logger

func Fuzz_Nosy_MockFPVMState_EncodeWitness__(f *testing.F) {
	f.Fuzz(func(t *testing.T, pc uint32, heapStart uint32) {
		m := MockCreateInitState(pc, heapStart)
		m.EncodeWitness()
	})
}

func Fuzz_Nosy_MockFPVMState_GetCpu__(f *testing.F) {
	f.Fuzz(func(t *testing.T, pc uint32, heapStart uint32) {
		m := MockCreateInitState(pc, heapStart)
		m.GetCpu()
	})
}

func Fuzz_Nosy_MockFPVMState_GetExitCode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, pc uint32, heapStart uint32) {
		m := MockCreateInitState(pc, heapStart)
		m.GetExitCode()
	})
}

func Fuzz_Nosy_MockFPVMState_GetExited__(f *testing.F) {
	f.Fuzz(func(t *testing.T, pc uint32, heapStart uint32) {
		m := MockCreateInitState(pc, heapStart)
		m.GetExited()
	})
}

func Fuzz_Nosy_MockFPVMState_GetHeap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, pc uint32, heapStart uint32) {
		m := MockCreateInitState(pc, heapStart)
		m.GetHeap()
	})
}

func Fuzz_Nosy_MockFPVMState_GetLastHint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, pc uint32, heapStart uint32) {
		m := MockCreateInitState(pc, heapStart)
		m.GetLastHint()
	})
}

func Fuzz_Nosy_MockFPVMState_GetMemory__(f *testing.F) {
	f.Fuzz(func(t *testing.T, pc uint32, heapStart uint32) {
		m := MockCreateInitState(pc, heapStart)
		m.GetMemory()
	})
}

func Fuzz_Nosy_MockFPVMState_GetPC__(f *testing.F) {
	f.Fuzz(func(t *testing.T, pc uint32, heapStart uint32) {
		m := MockCreateInitState(pc, heapStart)
		m.GetPC()
	})
}

func Fuzz_Nosy_MockFPVMState_GetPreimageKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, pc uint32, heapStart uint32) {
		m := MockCreateInitState(pc, heapStart)
		m.GetPreimageKey()
	})
}

func Fuzz_Nosy_MockFPVMState_GetPreimageOffset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, pc uint32, heapStart uint32) {
		m := MockCreateInitState(pc, heapStart)
		m.GetPreimageOffset()
	})
}

func Fuzz_Nosy_MockFPVMState_GetRegistersRef__(f *testing.F) {
	f.Fuzz(func(t *testing.T, pc uint32, heapStart uint32) {
		m := MockCreateInitState(pc, heapStart)
		m.GetRegistersRef()
	})
}

func Fuzz_Nosy_MockFPVMState_GetStep__(f *testing.F) {
	f.Fuzz(func(t *testing.T, pc uint32, heapStart uint32) {
		m := MockCreateInitState(pc, heapStart)
		m.GetStep()
	})
}

func Fuzz_Nosy_MockFPVMState_Serialize__(f *testing.F) {
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

		m := MockCreateInitState(pc, heapStart)
		m.Serialize(out)
	})
}

func Fuzz_Nosy_MockProgWithReader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var progType elf.ProgType
		fill_err = tp.Fill(&progType)
		if fill_err != nil {
			return
		}
		var filesz uint64
		fill_err = tp.Fill(&filesz)
		if fill_err != nil {
			return
		}
		var memsz uint64
		fill_err = tp.Fill(&memsz)
		if fill_err != nil {
			return
		}
		var vaddr uint64
		fill_err = tp.Fill(&vaddr)
		if fill_err != nil {
			return
		}
		var d5 []byte
		fill_err = tp.Fill(&d5)
		if fill_err != nil {
			return
		}

		MockProgWithReader(progType, filesz, memsz, vaddr, d5)
	})
}
