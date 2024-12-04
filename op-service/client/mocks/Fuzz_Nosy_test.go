package mocks

import (
	context "context"
	http "net/http"
	url "net/url"
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

func Fuzz_Nosy_HTTP_EXPECT__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *HTTP
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.EXPECT()
	})
}

func Fuzz_Nosy_HTTP_Get__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *HTTP
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var path string
		fill_err = tp.Fill(&path)
		if fill_err != nil {
			return
		}
		var query url.Values
		fill_err = tp.Fill(&query)
		if fill_err != nil {
			return
		}
		var headers http.Header
		fill_err = tp.Fill(&headers)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.Get(ctx, path, query, headers)
	})
}

// skipping Fuzz_Nosy_HTTP_Expecter_Get__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_HTTP_Get_Call_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_HTTP_Get_Call_Run__ because parameters include func, chan, or unsupported interface: func(ctx context.Context, path string, query net/url.Values, headers net/http.Header)

// skipping Fuzz_Nosy_HTTP_Get_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(context.Context, string, net/url.Values, net/http.Header) (*net/http.Response, error)

// skipping Fuzz_Nosy_mockConstructorTestingTNewHTTP_Cleanup__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/client/mocks.mockConstructorTestingTNewHTTP
