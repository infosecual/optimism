package mipsevm

import (
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

func Fuzz_Nosy_LoggingWriter_Write__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var lw *LoggingWriter
		fill_err = tp.Fill(&lw)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if lw == nil {
			return
		}

		lw.Write(b)
	})
}

func Fuzz_Nosy_StepWitness_HasPreimage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var wit *StepWitness
		fill_err = tp.Fill(&wit)
		if fill_err != nil {
			return
		}
		if wit == nil {
			return
		}

		wit.HasPreimage()
	})
}

// skipping Fuzz_Nosy_FPVM_CheckInfiniteLoop__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/cannon/mipsevm.FPVM

// skipping Fuzz_Nosy_FPVM_GetDebugInfo__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/cannon/mipsevm.FPVM

// skipping Fuzz_Nosy_FPVM_GetState__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/cannon/mipsevm.FPVM

// skipping Fuzz_Nosy_FPVM_InitDebug__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/cannon/mipsevm.FPVM

// skipping Fuzz_Nosy_FPVM_LastPreimage__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/cannon/mipsevm.FPVM

// skipping Fuzz_Nosy_FPVM_LookupSymbol__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/cannon/mipsevm.FPVM

// skipping Fuzz_Nosy_FPVM_Step__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/cannon/mipsevm.FPVM

// skipping Fuzz_Nosy_FPVM_Traceback__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/cannon/mipsevm.FPVM

// skipping Fuzz_Nosy_FPVMState_CreateVM__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/cannon/mipsevm.FPVMState

// skipping Fuzz_Nosy_FPVMState_EncodeWitness__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/cannon/mipsevm.FPVMState

// skipping Fuzz_Nosy_FPVMState_GetCpu__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/cannon/mipsevm.FPVMState

// skipping Fuzz_Nosy_FPVMState_GetExitCode__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/cannon/mipsevm.FPVMState

// skipping Fuzz_Nosy_FPVMState_GetExited__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/cannon/mipsevm.FPVMState

// skipping Fuzz_Nosy_FPVMState_GetHeap__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/cannon/mipsevm.FPVMState

// skipping Fuzz_Nosy_FPVMState_GetLastHint__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/cannon/mipsevm.FPVMState

// skipping Fuzz_Nosy_FPVMState_GetMemory__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/cannon/mipsevm.FPVMState

// skipping Fuzz_Nosy_FPVMState_GetPC__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/cannon/mipsevm.FPVMState

// skipping Fuzz_Nosy_FPVMState_GetPreimageKey__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/cannon/mipsevm.FPVMState

// skipping Fuzz_Nosy_FPVMState_GetPreimageOffset__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/cannon/mipsevm.FPVMState

// skipping Fuzz_Nosy_FPVMState_GetRegistersRef__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/cannon/mipsevm.FPVMState

// skipping Fuzz_Nosy_FPVMState_GetStep__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/cannon/mipsevm.FPVMState

func Fuzz_Nosy_HexU32_MarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v HexU32
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}

		v.MarshalText()
	})
}

func Fuzz_Nosy_HexU32_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v HexU32
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}

		v.String()
	})
}

// skipping Fuzz_Nosy_Metadata_CreateSymbolMatcher__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/cannon/mipsevm.Metadata

// skipping Fuzz_Nosy_Metadata_LookupSymbol__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/cannon/mipsevm.Metadata

// skipping Fuzz_Nosy_PreimageOracle_GetPreimage__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/cannon/mipsevm.PreimageOracle

// skipping Fuzz_Nosy_PreimageOracle_Hint__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/cannon/mipsevm.PreimageOracle

func Fuzz_Nosy_AppendBoolToWitness__(f *testing.F) {
	f.Fuzz(func(t *testing.T, witnessData []byte, boolVal bool) {
		AppendBoolToWitness(witnessData, boolVal)
	})
}

func Fuzz_Nosy_VmStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, exited bool, exitCode uint8) {
		VmStatus(exited, exitCode)
	})
}

func Fuzz_Nosy_logAsText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b string) {
		logAsText(b)
	})
}
