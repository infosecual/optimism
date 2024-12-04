package memory

import (
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

func Fuzz_Nosy_CachedPage_InvalidateFull__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *CachedPage
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.InvalidateFull()
	})
}

func Fuzz_Nosy_CachedPage_MerkleRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *CachedPage
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.MerkleRoot()
	})
}

func Fuzz_Nosy_CachedPage_MerkleizeSubtree__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *CachedPage
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var gindex uint64
		fill_err = tp.Fill(&gindex)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.MerkleizeSubtree(gindex)
	})
}

func Fuzz_Nosy_CachedPage_invalidate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *CachedPage
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var pageAddr uint32
		fill_err = tp.Fill(&pageAddr)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.invalidate(pageAddr)
	})
}

func Fuzz_Nosy_Memory_AllocPage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, pageIndex uint32) {
		m := NewMemory()
		m.AllocPage(pageIndex)
	})
}

func Fuzz_Nosy_Memory_Deserialize__(f *testing.F) {
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

		m := NewMemory()
		m.Deserialize(in)
	})
}

// skipping Fuzz_Nosy_Memory_ForEachPage__ because parameters include func, chan, or unsupported interface: func(pageIndex uint32, page *github.com/ethereum-optimism/optimism/cannon/mipsevm/memory.Page) error

func Fuzz_Nosy_Memory_GetWord__(f *testing.F) {
	f.Fuzz(func(t *testing.T, addr uint32) {
		m := NewMemory()
		m.GetWord(addr)
	})
}

func Fuzz_Nosy_Memory_MerkleProof__(f *testing.F) {
	f.Fuzz(func(t *testing.T, addr uint32) {
		m := NewMemory()
		m.MerkleProof(addr)
	})
}

func Fuzz_Nosy_Memory_MerkleizeSubtree__(f *testing.F) {
	f.Fuzz(func(t *testing.T, gindex uint64) {
		m := NewMemory()
		m.MerkleizeSubtree(gindex)
	})
}

func Fuzz_Nosy_Memory_ReadMemoryRange__(f *testing.F) {
	f.Fuzz(func(t *testing.T, addr uint32, count uint32) {
		m := NewMemory()
		m.ReadMemoryRange(addr, count)
	})
}

func Fuzz_Nosy_Memory_Serialize__(f *testing.F) {
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

		m := NewMemory()
		m.Serialize(out)
	})
}

func Fuzz_Nosy_Memory_SetMemoryRange__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var addr uint32
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var r io.Reader
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}

		m := NewMemory()
		m.SetMemoryRange(addr, r)
	})
}

func Fuzz_Nosy_Memory_SetWord__(f *testing.F) {
	f.Fuzz(func(t *testing.T, addr uint32, v uint32) {
		m := NewMemory()
		m.SetWord(addr, v)
	})
}

func Fuzz_Nosy_Memory_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d1 []byte) {
		m := NewMemory()
		m.UnmarshalJSON(d1)
	})
}

func Fuzz_Nosy_Memory_invalidate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, addr uint32) {
		m := NewMemory()
		m.invalidate(addr)
	})
}

func Fuzz_Nosy_Memory_pageLookup__(f *testing.F) {
	f.Fuzz(func(t *testing.T, pageIndex uint32) {
		m := NewMemory()
		m.pageLookup(pageIndex)
	})
}

func Fuzz_Nosy_Memory_traverseBranch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, parent uint64, addr uint32, depth uint8) {
		m := NewMemory()
		m.traverseBranch(parent, addr, depth)
	})
}

func Fuzz_Nosy_Page_MarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *Page
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.MarshalJSON()
	})
}

func Fuzz_Nosy_Page_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *Page
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var dat []byte
		fill_err = tp.Fill(&dat)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.UnmarshalJSON(dat)
	})
}

func Fuzz_Nosy_Page_UnmarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *Page
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var dat []byte
		fill_err = tp.Fill(&dat)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.UnmarshalText(dat)
	})
}

func Fuzz_Nosy_memReader_Read__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *memReader
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var dest []byte
		fill_err = tp.Fill(&dest)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.Read(dest)
	})
}

func Fuzz_Nosy_HashPair__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var left [32]byte
		fill_err = tp.Fill(&left)
		if fill_err != nil {
			return
		}
		var right [32]byte
		fill_err = tp.Fill(&right)
		if fill_err != nil {
			return
		}

		HashPair(left, right)
	})
}
