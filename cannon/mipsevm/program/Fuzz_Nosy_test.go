package program

import (
	"debug/elf"
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

func Fuzz_Nosy_Metadata_CreateSymbolMatcher__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var elfProgram *elf.File
		fill_err = tp.Fill(&elfProgram)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		if elfProgram == nil {
			return
		}

		m, err := MakeMetadata(elfProgram)
		if err != nil {
			return
		}
		m.CreateSymbolMatcher(name)
	})
}

func Fuzz_Nosy_Metadata_LookupSymbol__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var elfProgram *elf.File
		fill_err = tp.Fill(&elfProgram)
		if fill_err != nil {
			return
		}
		var addr uint32
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if elfProgram == nil {
			return
		}

		m, err := MakeMetadata(elfProgram)
		if err != nil {
			return
		}
		m.LookupSymbol(addr)
	})
}

// skipping Fuzz_Nosy_LoadELF__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/cannon/mipsevm/program.CreateInitialFPVMState[T]

func Fuzz_Nosy_pad__(f *testing.F) {
	f.Fuzz(func(t *testing.T, buf []byte) {
		pad(buf)
	})
}
