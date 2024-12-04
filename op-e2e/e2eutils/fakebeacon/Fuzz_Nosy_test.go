package fakebeacon

import (
	"testing"

	"github.com/ethereum/go-ethereum/beacon/engine"
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

func Fuzz_Nosy_FakeBeacon_BeaconAddr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FakeBeacon
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.BeaconAddr()
	})
}

func Fuzz_Nosy_FakeBeacon_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FakeBeacon
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.Close()
	})
}

func Fuzz_Nosy_FakeBeacon_LoadBlobsBundle__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FakeBeacon
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var slot uint64
		fill_err = tp.Fill(&slot)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.LoadBlobsBundle(slot)
	})
}

func Fuzz_Nosy_FakeBeacon_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FakeBeacon
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var addr string
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.Start(addr)
	})
}

func Fuzz_Nosy_FakeBeacon_StoreBlobsBundle__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FakeBeacon
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var slot uint64
		fill_err = tp.Fill(&slot)
		if fill_err != nil {
			return
		}
		var bundle *engine.BlobsBundleV1
		fill_err = tp.Fill(&bundle)
		if fill_err != nil {
			return
		}
		if f1 == nil || bundle == nil {
			return
		}

		f1.StoreBlobsBundle(slot, bundle)
	})
}
