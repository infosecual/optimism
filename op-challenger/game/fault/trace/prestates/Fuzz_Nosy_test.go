package prestates

import (
	"context"
	"net/url"
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

func Fuzz_Nosy_MultiPrestateProvider_PrestatePath__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MultiPrestateProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.PrestatePath(ctx, hash)
	})
}

func Fuzz_Nosy_MultiPrestateProvider_fetch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MultiPrestateProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var prestateUrl *url.URL
		fill_err = tp.Fill(&prestateUrl)
		if fill_err != nil {
			return
		}
		if m == nil || prestateUrl == nil {
			return
		}

		m.fetch(ctx, prestateUrl)
	})
}

func Fuzz_Nosy_MultiPrestateProvider_fetchPrestate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MultiPrestateProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		var fileType string
		fill_err = tp.Fill(&fileType)
		if fill_err != nil {
			return
		}
		var dest string
		fill_err = tp.Fill(&dest)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.fetchPrestate(ctx, hash, fileType, dest)
	})
}

func Fuzz_Nosy_PrestateProviderCache_GetOrCreate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *PrestateProviderCache
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var prestateHash common.Hash
		fill_err = tp.Fill(&prestateHash)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.GetOrCreate(ctx, prestateHash)
	})
}

func Fuzz_Nosy_SinglePrestateSource_PrestatePath__(f *testing.F) {
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
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var _x3 common.Hash
		fill_err = tp.Fill(&_x3)
		if fill_err != nil {
			return
		}

		s := NewSinglePrestateSource(path)
		s.PrestatePath(_x2, _x3)
	})
}

// skipping Fuzz_Nosy_PrestateSource_PrestatePath__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/trace/prestates.PrestateSource
