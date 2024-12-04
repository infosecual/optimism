package ioutil

import (
	"io"
	"io/fs"
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

func Fuzz_Nosy_AtomicWriter_Abort__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var path string
		fill_err = tp.Fill(&path)
		if fill_err != nil {
			return
		}
		var perm fs.FileMode
		fill_err = tp.Fill(&perm)
		if fill_err != nil {
			return
		}

		a, err := NewAtomicWriterCompressed(path, perm)
		if err != nil {
			return
		}
		a.Abort()
	})
}

func Fuzz_Nosy_AtomicWriter_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var path string
		fill_err = tp.Fill(&path)
		if fill_err != nil {
			return
		}
		var perm fs.FileMode
		fill_err = tp.Fill(&perm)
		if fill_err != nil {
			return
		}

		a, err := NewAtomicWriterCompressed(path, perm)
		if err != nil {
			return
		}
		a.Close()
	})
}

func Fuzz_Nosy_AtomicWriter_Write__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var path string
		fill_err = tp.Fill(&path)
		if fill_err != nil {
			return
		}
		var perm fs.FileMode
		fill_err = tp.Fill(&perm)
		if fill_err != nil {
			return
		}
		var p []byte
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}

		a, err := NewAtomicWriterCompressed(path, perm)
		if err != nil {
			return
		}
		a.Write(p)
	})
}

func Fuzz_Nosy_WrappedReadCloser_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r io.ReadCloser
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var c io.Closer
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}

		g := NewWrappedReadCloser(r, c)
		g.Close()
	})
}

func Fuzz_Nosy_WrappedWriteCloser_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *WrappedWriteCloser
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.Close()
	})
}

func Fuzz_Nosy_noopCloser_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *noopCloser
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.Close()
	})
}

func Fuzz_Nosy_IsGzip__(f *testing.F) {
	f.Fuzz(func(t *testing.T, path string) {
		IsGzip(path)
	})
}
