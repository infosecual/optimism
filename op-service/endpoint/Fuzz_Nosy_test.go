package endpoint

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

func Fuzz_Nosy_ServerRPC_ClientRPC__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *ServerRPC
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.ClientRPC()
	})
}

func Fuzz_Nosy_ServerRPC_HttpRPC__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *ServerRPC
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.HttpRPC()
	})
}

func Fuzz_Nosy_ServerRPC_RPC__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *ServerRPC
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.RPC()
	})
}

func Fuzz_Nosy_ServerRPC_WsRPC__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *ServerRPC
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.WsRPC()
	})
}

// skipping Fuzz_Nosy_ClientRPC_ClientRPC__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/endpoint.ClientRPC

// skipping Fuzz_Nosy_HttpRPC_HttpRPC__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/endpoint.HttpRPC

func Fuzz_Nosy_HttpURL_HttpRPC__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var url HttpURL
		fill_err = tp.Fill(&url)
		if fill_err != nil {
			return
		}

		url.HttpRPC()
	})
}

func Fuzz_Nosy_HttpURL_RPC__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var url HttpURL
		fill_err = tp.Fill(&url)
		if fill_err != nil {
			return
		}

		url.RPC()
	})
}

// skipping Fuzz_Nosy_RPC_RPC__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/endpoint.RPC

// skipping Fuzz_Nosy_RestHTTP_RestHTTP__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/endpoint.RestHTTP

func Fuzz_Nosy_RestHTTPURL_RestHTTP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var url RestHTTPURL
		fill_err = tp.Fill(&url)
		if fill_err != nil {
			return
		}

		url.RestHTTP()
	})
}

func Fuzz_Nosy_WsOrHttpRPC_HttpRPC__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r WsOrHttpRPC
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}

		r.HttpRPC()
	})
}

func Fuzz_Nosy_WsOrHttpRPC_RPC__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r WsOrHttpRPC
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}

		r.RPC()
	})
}

func Fuzz_Nosy_WsOrHttpRPC_WsRPC__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r WsOrHttpRPC
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}

		r.WsRPC()
	})
}

// skipping Fuzz_Nosy_WsRPC_WsRPC__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/endpoint.WsRPC

func Fuzz_Nosy_WsURL_RPC__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var url WsURL
		fill_err = tp.Fill(&url)
		if fill_err != nil {
			return
		}

		url.RPC()
	})
}

func Fuzz_Nosy_WsURL_WsRPC__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var url WsURL
		fill_err = tp.Fill(&url)
		if fill_err != nil {
			return
		}

		url.WsRPC()
	})
}

// skipping Fuzz_Nosy_SelectRPC__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/endpoint.RPC
