package preimage

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

func Fuzz_Nosy_FilePoller_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *FilePoller
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.Close()
	})
}

func Fuzz_Nosy_FilePoller_Read__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FilePoller
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.Read(b)
	})
}

func Fuzz_Nosy_FilePoller_Write__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FilePoller
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.Write(b)
	})
}

// skipping Fuzz_Nosy_HintReader_NextHint__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-preimage.HintHandler

// skipping Fuzz_Nosy_HintWriter_Hint__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-preimage.Hint

// skipping Fuzz_Nosy_OracleClient_Get__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-preimage.Key

// skipping Fuzz_Nosy_OracleServer_NextPreimageRequest__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-preimage.PreimageGetter

func Fuzz_Nosy_ReadWritePair_Read__(f *testing.F) {
	f.Fuzz(func(t *testing.T, p []byte) {
		rw := ClientHinterChannel()
		rw.Read(p)
	})
}

func Fuzz_Nosy_ReadWritePair_Write__(f *testing.F) {
	f.Fuzz(func(t *testing.T, p []byte) {
		rw := ClientHinterChannel()
		rw.Write(p)
	})
}

func Fuzz_Nosy_BlobKey_PreimageKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k BlobKey
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}

		k.PreimageKey()
	})
}

func Fuzz_Nosy_BlobKey_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k BlobKey
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}

		k.String()
	})
}

func Fuzz_Nosy_BlobKey_TerminalString__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k BlobKey
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}

		k.TerminalString()
	})
}

// skipping Fuzz_Nosy_FileChannel_Reader__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-preimage.FileChannel

// skipping Fuzz_Nosy_FileChannel_Writer__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-preimage.FileChannel

// skipping Fuzz_Nosy_Hint_Hint__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-preimage.Hint

// skipping Fuzz_Nosy_Hinter_Hint__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-preimage.Hinter

// skipping Fuzz_Nosy_HinterFn_Hint__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-preimage.HinterFn

func Fuzz_Nosy_Keccak256Key_PreimageKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keccak256Key
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}

		k.PreimageKey()
	})
}

func Fuzz_Nosy_Keccak256Key_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keccak256Key
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}

		k.String()
	})
}

func Fuzz_Nosy_Keccak256Key_TerminalString__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keccak256Key
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}

		k.TerminalString()
	})
}

// skipping Fuzz_Nosy_Key_PreimageKey__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-preimage.Key

func Fuzz_Nosy_LocalIndexKey_PreimageKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k LocalIndexKey
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}

		k.PreimageKey()
	})
}

// skipping Fuzz_Nosy_Oracle_Get__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-preimage.Oracle

// skipping Fuzz_Nosy_OracleFn_Get__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-preimage.OracleFn

func Fuzz_Nosy_PrecompileKey_PreimageKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k PrecompileKey
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}

		k.PreimageKey()
	})
}

func Fuzz_Nosy_PrecompileKey_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k PrecompileKey
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}

		k.String()
	})
}

func Fuzz_Nosy_PrecompileKey_TerminalString__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k PrecompileKey
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}

		k.TerminalString()
	})
}

func Fuzz_Nosy_Sha256Key_PreimageKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Sha256Key
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}

		k.PreimageKey()
	})
}

func Fuzz_Nosy_Sha256Key_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Sha256Key
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}

		k.String()
	})
}

func Fuzz_Nosy_Sha256Key_TerminalString__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Sha256Key
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}

		k.TerminalString()
	})
}

func Fuzz_Nosy_Keccak256__(f *testing.F) {
	f.Fuzz(func(t *testing.T, v []byte) {
		Keccak256(v)
	})
}
