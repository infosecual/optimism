package geth

import (
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
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

func Fuzz_Nosy_GethInstance_AuthRPC__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var gi *GethInstance
		fill_err = tp.Fill(&gi)
		if fill_err != nil {
			return
		}
		if gi == nil {
			return
		}

		gi.AuthRPC()
	})
}

func Fuzz_Nosy_GethInstance_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var gi *GethInstance
		fill_err = tp.Fill(&gi)
		if fill_err != nil {
			return
		}
		if gi == nil {
			return
		}

		gi.Close()
	})
}

func Fuzz_Nosy_GethInstance_UserRPC__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var gi *GethInstance
		fill_err = tp.Fill(&gi)
		if fill_err != nil {
			return
		}
		if gi == nil {
			return
		}

		gi.UserRPC()
	})
}

func Fuzz_Nosy_fakePoS_FakeBeaconBlockRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *fakePoS
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var time uint64
		fill_err = tp.Fill(&time)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.FakeBeaconBlockRoot(time)
	})
}

func Fuzz_Nosy_fakePoS_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *fakePoS
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.Start()
	})
}

func Fuzz_Nosy_fakePoS_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *fakePoS
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.Stop()
	})
}

// skipping Fuzz_Nosy_Beacon_StoreBlobsBundle__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/e2eutils/geth.Beacon

func Fuzz_Nosy_ConnectP2P__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var node1 *ethclient.Client
		fill_err = tp.Fill(&node1)
		if fill_err != nil {
			return
		}
		var node2 *ethclient.Client
		fill_err = tp.Fill(&node2)
		if fill_err != nil {
			return
		}
		if t1 == nil || node1 == nil || node2 == nil {
			return
		}

		ConnectP2P(t1, node1, node2)
	})
}
