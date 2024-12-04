package matrix

import (
	"io"
	"testing"

	"github.com/ethereum/go-ethereum/common"
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

func Fuzz_Nosy_StateMatrix_AbsorbUpTo__(f *testing.F) {
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
		var maxLen int
		fill_err = tp.Fill(&maxLen)
		if fill_err != nil {
			return
		}

		d := NewStateMatrix()
		d.AbsorbUpTo(in, maxLen)
	})
}

func Fuzz_Nosy_StateMatrix_absorbLeafInput__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d1 []byte, final bool) {
		d := NewStateMatrix()
		d.absorbLeafInput(d1, final)
	})
}

// skipping Fuzz_Nosy_StateMatrix_absorbNextLeafInput__ because parameters include func, chan, or unsupported interface: func() github.com/ethereum/go-ethereum/common.Hash

func Fuzz_Nosy_StateMatrix_newLeafWithPadding__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		var index uint64
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		var commitment common.Hash
		fill_err = tp.Fill(&commitment)
		if fill_err != nil {
			return
		}
		var final bool
		fill_err = tp.Fill(&final)
		if fill_err != nil {
			return
		}

		d := NewStateMatrix()
		d.newLeafWithPadding(input, index, commitment, final)
	})
}

func Fuzz_Nosy_state_Read__(f *testing.F) {
	f.Fuzz(func(t *testing.T, out []byte) {
		d := newLegacyKeccak256()
		d.Read(out)
	})
}

func Fuzz_Nosy_state_Sum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, in []byte) {
		d := newLegacyKeccak256()
		d.Sum(in)
	})
}

func Fuzz_Nosy_state_Write__(f *testing.F) {
	f.Fuzz(func(t *testing.T, p []byte) {
		d := newLegacyKeccak256()
		d.Write(p)
	})
}

func Fuzz_Nosy_state_padAndPermute__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dsbyte byte) {
		d := newLegacyKeccak256()
		d.padAndPermute(dsbyte)
	})
}

func Fuzz_Nosy_storageBuf_asBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *storageBuf
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.asBytes()
	})
}

func Fuzz_Nosy_copyOutGeneric__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *state
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		copyOutGeneric(d, b)
	})
}

func Fuzz_Nosy_keccakF1600__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a *[25]uint64
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if a == nil {
			return
		}

		keccakF1600(a)
	})
}

func Fuzz_Nosy_pad__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		var paddedInput *[136]byte
		fill_err = tp.Fill(&paddedInput)
		if fill_err != nil {
			return
		}
		var dsbyte byte
		fill_err = tp.Fill(&dsbyte)
		if fill_err != nil {
			return
		}
		if paddedInput == nil {
			return
		}

		pad(input, paddedInput, dsbyte)
	})
}

func Fuzz_Nosy_xorInGeneric__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *state
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var buf []byte
		fill_err = tp.Fill(&buf)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		xorInGeneric(d, buf)
	})
}
