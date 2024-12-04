package testutil

import (
	"testing"

	"github.com/ethereum-optimism/optimism/cannon/mipsevm"
	"github.com/ethereum-optimism/optimism/cannon/mipsevm/memory"
	"github.com/ethereum-optimism/optimism/op-chain-ops/foundry"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/tracing"
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

// skipping Fuzz_Nosy_EvmValidator_ValidateEVM__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/cannon/mipsevm.FPVM

func Fuzz_Nosy_ExpectedState_ExpectMemoryWriteWord__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *ExpectedState
		fill_err = tp.Fill(&e)
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
		if e == nil {
			return
		}

		e.ExpectMemoryWriteWord(addr, val)
	})
}

func Fuzz_Nosy_ExpectedState_ExpectStep__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *ExpectedState
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.ExpectStep()
	})
}

// skipping Fuzz_Nosy_ExpectedState_Validate__ because parameters include func, chan, or unsupported interface: testing.TB

func Fuzz_Nosy_HintTrackingOracle_GetPreimage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *HintTrackingOracle
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var k [32]byte
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.GetPreimage(k)
	})
}

func Fuzz_Nosy_HintTrackingOracle_Hint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *HintTrackingOracle
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var v []byte
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Hint(v)
	})
}

func Fuzz_Nosy_HintTrackingOracle_Hints__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *HintTrackingOracle
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Hints()
	})
}

// skipping Fuzz_Nosy_MIPSEVM_SetLocalOracle__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/cannon/mipsevm.PreimageOracle

func Fuzz_Nosy_MIPSEVM_SetSourceMapTracer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MIPSEVM
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var t2 *testing.T
		fill_err = tp.Fill(&t2)
		if fill_err != nil {
			return
		}
		var version MipsVersion
		fill_err = tp.Fill(&version)
		if fill_err != nil {
			return
		}
		if m == nil || t2 == nil {
			return
		}

		m.SetSourceMapTracer(t2, version)
	})
}

func Fuzz_Nosy_MIPSEVM_SetTracer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MIPSEVM
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var tracer *tracing.Hooks
		fill_err = tp.Fill(&tracer)
		if fill_err != nil {
			return
		}
		if m == nil || tracer == nil {
			return
		}

		m.SetTracer(tracer)
	})
}

// skipping Fuzz_Nosy_MIPSEVM_Step__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/cannon/mipsevm.HashFn

func Fuzz_Nosy_MIPSEVM_assertPreimageOracleReverts__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MIPSEVM
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var t2 *testing.T
		fill_err = tp.Fill(&t2)
		if fill_err != nil {
			return
		}
		var preimageKey [32]byte
		fill_err = tp.Fill(&preimageKey)
		if fill_err != nil {
			return
		}
		var preimageValue []byte
		fill_err = tp.Fill(&preimageValue)
		if fill_err != nil {
			return
		}
		var preimageOffset uint32
		fill_err = tp.Fill(&preimageOffset)
		if fill_err != nil {
			return
		}
		if m == nil || t2 == nil {
			return
		}

		m.assertPreimageOracleReverts(t2, preimageKey, preimageValue, preimageOffset)
	})
}

func Fuzz_Nosy_MIPSEVM_encodePreimageOracleInput__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MIPSEVM
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var t2 *testing.T
		fill_err = tp.Fill(&t2)
		if fill_err != nil {
			return
		}
		var preimageKey [32]byte
		fill_err = tp.Fill(&preimageKey)
		if fill_err != nil {
			return
		}
		var preimageValue []byte
		fill_err = tp.Fill(&preimageValue)
		if fill_err != nil {
			return
		}
		var preimageOffset uint32
		fill_err = tp.Fill(&preimageOffset)
		if fill_err != nil {
			return
		}
		var localContext mipsevm.LocalContext
		fill_err = tp.Fill(&localContext)
		if fill_err != nil {
			return
		}
		if m == nil || t2 == nil {
			return
		}

		m.encodePreimageOracleInput(t2, preimageKey, preimageValue, preimageOffset, localContext)
	})
}

func Fuzz_Nosy_RandHelper_Fraction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed int64) {
		h := NewRandHelper(seed)
		h.Fraction()
	})
}

func Fuzz_Nosy_RandHelper_Intn__(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed int64, n int) {
		h := NewRandHelper(seed)
		h.Intn(n)
	})
}

func Fuzz_Nosy_RandHelper_RandHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed int64) {
		h := NewRandHelper(seed)
		h.RandHash()
	})
}

func Fuzz_Nosy_RandHelper_RandHint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed int64) {
		h := NewRandHelper(seed)
		h.RandHint()
	})
}

func Fuzz_Nosy_RandHelper_RandPC__(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed int64) {
		h := NewRandHelper(seed)
		h.RandPC()
	})
}

func Fuzz_Nosy_RandHelper_RandRegisters__(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed int64) {
		h := NewRandHelper(seed)
		h.RandRegisters()
	})
}

func Fuzz_Nosy_RandHelper_RandStep__(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed int64) {
		h := NewRandHelper(seed)
		h.RandStep()
	})
}

// skipping Fuzz_Nosy_RandHelper_RandomBytes__ because parameters include func, chan, or unsupported interface: github.com/stretchr/testify/require.TestingT

func Fuzz_Nosy_RandHelper_Uint32__(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed int64) {
		h := NewRandHelper(seed)
		h.Uint32()
	})
}

func Fuzz_Nosy_RandHelper_Word__(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed int64) {
		h := NewRandHelper(seed)
		h.Word()
	})
}

func Fuzz_Nosy_TestOracle_GetPreimage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var preimageData []byte
		fill_err = tp.Fill(&preimageData)
		if fill_err != nil {
			return
		}
		var k [32]byte
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1 := StaticOracle(t1, preimageData)
		t1.GetPreimage(k)
	})
}

func Fuzz_Nosy_TestOracle_Hint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var preimageData []byte
		fill_err = tp.Fill(&preimageData)
		if fill_err != nil {
			return
		}
		var v []byte
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1 := StaticOracle(t1, preimageData)
		t1.Hint(v)
	})
}

func Fuzz_Nosy_testChain_Engine__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *testChain
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.Engine()
	})
}

func Fuzz_Nosy_testChain_GetHeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *testChain
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var h common.Hash
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var n uint64
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.GetHeader(h, n)
	})
}

// skipping Fuzz_Nosy_StateMutator_Randomize__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/cannon/mipsevm/testutil.StateMutator

// skipping Fuzz_Nosy_StateMutator_SetExitCode__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/cannon/mipsevm/testutil.StateMutator

// skipping Fuzz_Nosy_StateMutator_SetExited__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/cannon/mipsevm/testutil.StateMutator

// skipping Fuzz_Nosy_StateMutator_SetHI__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/cannon/mipsevm/testutil.StateMutator

// skipping Fuzz_Nosy_StateMutator_SetHeap__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/cannon/mipsevm/testutil.StateMutator

// skipping Fuzz_Nosy_StateMutator_SetLO__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/cannon/mipsevm/testutil.StateMutator

// skipping Fuzz_Nosy_StateMutator_SetLastHint__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/cannon/mipsevm/testutil.StateMutator

// skipping Fuzz_Nosy_StateMutator_SetNextPC__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/cannon/mipsevm/testutil.StateMutator

// skipping Fuzz_Nosy_StateMutator_SetPC__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/cannon/mipsevm/testutil.StateMutator

// skipping Fuzz_Nosy_StateMutator_SetPreimageKey__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/cannon/mipsevm/testutil.StateMutator

// skipping Fuzz_Nosy_StateMutator_SetPreimageOffset__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/cannon/mipsevm/testutil.StateMutator

// skipping Fuzz_Nosy_StateMutator_SetStep__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/cannon/mipsevm/testutil.StateMutator

func Fuzz_Nosy_AddHintLengthPrefix__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d1 []byte) {
		AddHintLengthPrefix(d1)
	})
}

func Fuzz_Nosy_AddPreimageLengthPrefix__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d1 []byte) {
		AddPreimageLengthPrefix(d1)
	})
}

func Fuzz_Nosy_AlignPC__(f *testing.F) {
	f.Fuzz(func(t *testing.T, pc uint32) {
		AlignPC(pc)
	})
}

// skipping Fuzz_Nosy_AssertEVMReverts__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/cannon/mipsevm.FPVMState

// skipping Fuzz_Nosy_AssertPreimageOracleReverts__ because parameters include func, chan, or unsupported interface: []github.com/ethereum-optimism/optimism/cannon/mipsevm/testutil.evmOption

func Fuzz_Nosy_BoundStep__(f *testing.F) {
	f.Fuzz(func(t *testing.T, step uint64) {
		BoundStep(step)
	})
}

// skipping Fuzz_Nosy_Cannon32OnlyTest__ because parameters include func, chan, or unsupported interface: testing.TB

func Fuzz_Nosy_ClaimTestOracle__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		ClaimTestOracle(t1)
	})
}

func Fuzz_Nosy_EncodeStepInput__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var wit *mipsevm.StepWitness
		fill_err = tp.Fill(&wit)
		if fill_err != nil {
			return
		}
		var localContext mipsevm.LocalContext
		fill_err = tp.Fill(&localContext)
		if fill_err != nil {
			return
		}
		var mips *foundry.Artifact
		fill_err = tp.Fill(&mips)
		if fill_err != nil {
			return
		}
		if t1 == nil || wit == nil || mips == nil {
			return
		}

		EncodeStepInput(t1, wit, localContext, mips)
	})
}

func Fuzz_Nosy_GetInstruction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var mem *memory.Memory
		fill_err = tp.Fill(&mem)
		if fill_err != nil {
			return
		}
		var pc uint32
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		if mem == nil {
			return
		}

		GetInstruction(mem, pc)
	})
}

// skipping Fuzz_Nosy_LoadELFProgram__ because parameters include func, chan, or unsupported interface: github.com/stretchr/testify/require.TestingT

func Fuzz_Nosy_LogStepFailureAtCleanup__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var mipsEvm *MIPSEVM
		fill_err = tp.Fill(&mipsEvm)
		if fill_err != nil {
			return
		}
		if t1 == nil || mipsEvm == nil {
			return
		}

		LogStepFailureAtCleanup(t1, mipsEvm)
	})
}

func Fuzz_Nosy_NewEVMEnv__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var contracts *ContractMetadata
		fill_err = tp.Fill(&contracts)
		if fill_err != nil {
			return
		}
		if contracts == nil {
			return
		}

		NewEVMEnv(contracts)
	})
}

func Fuzz_Nosy_ProgramPath__(f *testing.F) {
	f.Fuzz(func(t *testing.T, programName string) {
		ProgramPath(programName)
	})
}

// skipping Fuzz_Nosy_RunVMTest_Claim__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/cannon/mipsevm/program.CreateInitialFPVMState[T]

// skipping Fuzz_Nosy_RunVMTest_Hello__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/cannon/mipsevm/program.CreateInitialFPVMState[T]

// skipping Fuzz_Nosy_RunVMTests_OpenMips__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/cannon/mipsevm/testutil.StateFactory[T]

// skipping Fuzz_Nosy_SetMemoryUint64__ because parameters include func, chan, or unsupported interface: github.com/stretchr/testify/require.TestingT

func Fuzz_Nosy_StoreInstruction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var mem *memory.Memory
		fill_err = tp.Fill(&mem)
		if fill_err != nil {
			return
		}
		var pc uint32
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var insn uint32
		fill_err = tp.Fill(&insn)
		if fill_err != nil {
			return
		}
		if mem == nil {
			return
		}

		StoreInstruction(mem, pc, insn)
	})
}

func Fuzz_Nosy_ToSignedInteger__(f *testing.F) {
	f.Fuzz(func(t *testing.T, x uint32) {
		ToSignedInteger(x)
	})
}

func Fuzz_Nosy_Uint32ToBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, val uint32) {
		Uint32ToBytes(val)
	})
}

func Fuzz_Nosy_Uint64ToBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, val uint64) {
		Uint64ToBytes(val)
	})
}

// skipping Fuzz_Nosy_ValidateEVM__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/cannon/mipsevm.FPVM
