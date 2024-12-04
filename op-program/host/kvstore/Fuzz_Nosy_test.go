package kvstore

import (
	"testing"

	"github.com/ethereum-optimism/optimism/op-program/host/config"
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

func Fuzz_Nosy_LocalPreimageSource_Get__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config *config.Config
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var key common.Hash
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		if config == nil {
			return
		}

		s := NewLocalPreimageSource(config)
		s.Get(key)
	})
}

func Fuzz_Nosy_MemKV_Get__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k common.Hash
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}

		m := NewMemKV()
		m.Get(k)
	})
}

func Fuzz_Nosy_MemKV_Put__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k common.Hash
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var v []byte
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}

		m := NewMemKV()
		m.Put(k, v)
	})
}

func Fuzz_Nosy_PreimageSourceSplitter_Get__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *PreimageSourceSplitter
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var key [32]byte
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Get(key)
	})
}

func Fuzz_Nosy_directoryKV_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, path string) {
		d := newDirectoryKV(path)
		d.Close()
	})
}

func Fuzz_Nosy_directoryKV_Get__(f *testing.F) {
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
		var k common.Hash
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}

		d := newDirectoryKV(path)
		d.Get(k)
	})
}

func Fuzz_Nosy_directoryKV_Put__(f *testing.F) {
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
		var k common.Hash
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var v []byte
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}

		d := newDirectoryKV(path)
		d.Put(k, v)
	})
}

func Fuzz_Nosy_directoryKV_pathKey__(f *testing.F) {
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
		var k common.Hash
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}

		d := newDirectoryKV(path)
		d.pathKey(k)
	})
}

func Fuzz_Nosy_fileKV_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, path string) {
		d := newFileKV(path)
		d.Close()
	})
}

func Fuzz_Nosy_fileKV_Get__(f *testing.F) {
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
		var k common.Hash
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}

		d := newFileKV(path)
		d.Get(k)
	})
}

func Fuzz_Nosy_fileKV_Put__(f *testing.F) {
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
		var k common.Hash
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var v []byte
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}

		d := newFileKV(path)
		d.Put(k, v)
	})
}

func Fuzz_Nosy_fileKV_pathKey__(f *testing.F) {
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
		var k common.Hash
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}

		d := newFileKV(path)
		d.pathKey(k)
	})
}

func Fuzz_Nosy_pebbleKV_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, path string) {
		d := newPebbleKV(path)
		d.Close()
	})
}

func Fuzz_Nosy_pebbleKV_Get__(f *testing.F) {
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
		var k common.Hash
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}

		d := newPebbleKV(path)
		d.Get(k)
	})
}

func Fuzz_Nosy_pebbleKV_Put__(f *testing.F) {
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
		var k common.Hash
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var v []byte
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}

		d := newPebbleKV(path)
		d.Put(k, v)
	})
}

// skipping Fuzz_Nosy_KV_Close__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-program/host/kvstore.KV

// skipping Fuzz_Nosy_KV_Get__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-program/host/kvstore.KV

// skipping Fuzz_Nosy_KV_Put__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-program/host/kvstore.KV
