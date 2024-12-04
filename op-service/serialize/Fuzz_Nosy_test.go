package serialize

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

func Fuzz_Nosy_BinaryReader_ReadBool__(f *testing.F) {
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
		var t2 *bool
		fill_err = tp.Fill(&t2)
		if fill_err != nil {
			return
		}
		if t2 == nil {
			return
		}

		r := NewBinaryReader(in)
		r.ReadBool(t2)
	})
}

func Fuzz_Nosy_BinaryReader_ReadBytes__(f *testing.F) {
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
		var t2 *[]byte
		fill_err = tp.Fill(&t2)
		if fill_err != nil {
			return
		}
		if t2 == nil {
			return
		}

		r := NewBinaryReader(in)
		r.ReadBytes(t2)
	})
}

func Fuzz_Nosy_BinaryReader_ReadHash__(f *testing.F) {
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
		var t2 *common.Hash
		fill_err = tp.Fill(&t2)
		if fill_err != nil {
			return
		}
		if t2 == nil {
			return
		}

		r := NewBinaryReader(in)
		r.ReadHash(t2)
	})
}

// skipping Fuzz_Nosy_BinaryReader_ReadUInt__ because parameters include func, chan, or unsupported interface: any

func Fuzz_Nosy_BinaryWriter_WriteBool__(f *testing.F) {
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
		var v bool
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}

		w := NewBinaryWriter(out)
		w.WriteBool(v)
	})
}

func Fuzz_Nosy_BinaryWriter_WriteBytes__(f *testing.F) {
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
		var v []byte
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}

		w := NewBinaryWriter(out)
		w.WriteBytes(v)
	})
}

func Fuzz_Nosy_BinaryWriter_WriteHash__(f *testing.F) {
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
		var v common.Hash
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}

		w := NewBinaryWriter(out)
		w.WriteHash(v)
	})
}

// skipping Fuzz_Nosy_BinaryWriter_WriteUInt__ because parameters include func, chan, or unsupported interface: any

// skipping Fuzz_Nosy_Deserializable_Deserialize__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/serialize.Deserializable

// skipping Fuzz_Nosy_Serializable_Serialize__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/serialize.Serializable

func Fuzz_Nosy_IsBinaryFile__(f *testing.F) {
	f.Fuzz(func(t *testing.T, path string) {
		IsBinaryFile(path)
	})
}

func Fuzz_Nosy_LoadSerializedBinary__(f *testing.F) {
	f.Fuzz(func(t *testing.T, inputPath string) {
		LoadSerializedBinary(inputPath)
	})
}
