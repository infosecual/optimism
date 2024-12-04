package arch

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

// skipping Fuzz_Nosy_ByteOrder_AppendWord__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/cannon/mipsevm/arch.ByteOrder

// skipping Fuzz_Nosy_ByteOrder_PutWord__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/cannon/mipsevm/arch.ByteOrder

// skipping Fuzz_Nosy_ByteOrder_Word__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/cannon/mipsevm/arch.ByteOrder

func Fuzz_Nosy_byteOrder32_AppendWord__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bo byteOrder32
		fill_err = tp.Fill(&bo)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var v uint32
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}

		bo.AppendWord(b, v)
	})
}

func Fuzz_Nosy_byteOrder32_PutWord__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bo byteOrder32
		fill_err = tp.Fill(&bo)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var v uint32
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}

		bo.PutWord(b, v)
	})
}

func Fuzz_Nosy_byteOrder32_Word__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bo byteOrder32
		fill_err = tp.Fill(&bo)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		bo.Word(b)
	})
}
