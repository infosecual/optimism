package mocks

import (
	context "context"
	"testing"

	eth "github.com/ethereum-optimism/optimism/op-service/eth"
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

func Fuzz_Nosy_BeaconClient_BeaconBlobSideCars__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *BeaconClient
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var fetchAllSidecars bool
		fill_err = tp.Fill(&fetchAllSidecars)
		if fill_err != nil {
			return
		}
		var slot uint64
		fill_err = tp.Fill(&slot)
		if fill_err != nil {
			return
		}
		var hashes []eth.IndexedBlobHash
		fill_err = tp.Fill(&hashes)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.BeaconBlobSideCars(ctx, fetchAllSidecars, slot, hashes)
	})
}

func Fuzz_Nosy_BeaconClient_BeaconGenesis__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *BeaconClient
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.BeaconGenesis(ctx)
	})
}

func Fuzz_Nosy_BeaconClient_ConfigSpec__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *BeaconClient
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.ConfigSpec(ctx)
	})
}

func Fuzz_Nosy_BeaconClient_EXPECT__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *BeaconClient
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

func Fuzz_Nosy_BeaconClient_NodeVersion__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *BeaconClient
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.NodeVersion(ctx)
	})
}

// skipping Fuzz_Nosy_BeaconClient_BeaconBlobSideCars_Call_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_BeaconClient_BeaconBlobSideCars_Call_Run__ because parameters include func, chan, or unsupported interface: func(ctx context.Context, fetchAllSidecars bool, slot uint64, hashes []github.com/ethereum-optimism/optimism/op-service/eth.IndexedBlobHash)

// skipping Fuzz_Nosy_BeaconClient_BeaconBlobSideCars_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(context.Context, bool, uint64, []github.com/ethereum-optimism/optimism/op-service/eth.IndexedBlobHash) (github.com/ethereum-optimism/optimism/op-service/eth.APIGetBlobSidecarsResponse, error)

// skipping Fuzz_Nosy_BeaconClient_BeaconGenesis_Call_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_BeaconClient_BeaconGenesis_Call_Run__ because parameters include func, chan, or unsupported interface: func(ctx context.Context)

// skipping Fuzz_Nosy_BeaconClient_BeaconGenesis_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(context.Context) (github.com/ethereum-optimism/optimism/op-service/eth.APIGenesisResponse, error)

// skipping Fuzz_Nosy_BeaconClient_ConfigSpec_Call_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_BeaconClient_ConfigSpec_Call_Run__ because parameters include func, chan, or unsupported interface: func(ctx context.Context)

// skipping Fuzz_Nosy_BeaconClient_ConfigSpec_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(context.Context) (github.com/ethereum-optimism/optimism/op-service/eth.APIConfigResponse, error)

// skipping Fuzz_Nosy_BeaconClient_Expecter_BeaconBlobSideCars__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_BeaconClient_Expecter_BeaconGenesis__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_BeaconClient_Expecter_ConfigSpec__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_BeaconClient_Expecter_NodeVersion__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_BeaconClient_NodeVersion_Call_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_BeaconClient_NodeVersion_Call_Run__ because parameters include func, chan, or unsupported interface: func(ctx context.Context)

// skipping Fuzz_Nosy_BeaconClient_NodeVersion_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(context.Context) (string, error)

func Fuzz_Nosy_BlobSideCarsFetcher_BeaconBlobSideCars__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *BlobSideCarsFetcher
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var fetchAllSidecars bool
		fill_err = tp.Fill(&fetchAllSidecars)
		if fill_err != nil {
			return
		}
		var slot uint64
		fill_err = tp.Fill(&slot)
		if fill_err != nil {
			return
		}
		var hashes []eth.IndexedBlobHash
		fill_err = tp.Fill(&hashes)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.BeaconBlobSideCars(ctx, fetchAllSidecars, slot, hashes)
	})
}

func Fuzz_Nosy_BlobSideCarsFetcher_EXPECT__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *BlobSideCarsFetcher
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

// skipping Fuzz_Nosy_BlobSideCarsFetcher_BeaconBlobSideCars_Call_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_BlobSideCarsFetcher_BeaconBlobSideCars_Call_Run__ because parameters include func, chan, or unsupported interface: func(ctx context.Context, fetchAllSidecars bool, slot uint64, hashes []github.com/ethereum-optimism/optimism/op-service/eth.IndexedBlobHash)

// skipping Fuzz_Nosy_BlobSideCarsFetcher_BeaconBlobSideCars_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(context.Context, bool, uint64, []github.com/ethereum-optimism/optimism/op-service/eth.IndexedBlobHash) (github.com/ethereum-optimism/optimism/op-service/eth.APIGetBlobSidecarsResponse, error)

// skipping Fuzz_Nosy_BlobSideCarsFetcher_Expecter_BeaconBlobSideCars__ because parameters include func, chan, or unsupported interface: interface{}
