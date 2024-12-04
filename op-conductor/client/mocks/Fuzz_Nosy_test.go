package mocks

import (
	context "context"
	"testing"

	eth "github.com/ethereum-optimism/optimism/op-service/eth"
	common "github.com/ethereum/go-ethereum/common"
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

func Fuzz_Nosy_SequencerControl_ConductorEnabled__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *SequencerControl
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

		_m.ConductorEnabled(ctx)
	})
}

func Fuzz_Nosy_SequencerControl_EXPECT__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *SequencerControl
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

func Fuzz_Nosy_SequencerControl_LatestUnsafeBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *SequencerControl
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

		_m.LatestUnsafeBlock(ctx)
	})
}

func Fuzz_Nosy_SequencerControl_PostUnsafePayload__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *SequencerControl
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var payload *eth.ExecutionPayloadEnvelope
		fill_err = tp.Fill(&payload)
		if fill_err != nil {
			return
		}
		if _m == nil || payload == nil {
			return
		}

		_m.PostUnsafePayload(ctx, payload)
	})
}

func Fuzz_Nosy_SequencerControl_SequencerActive__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *SequencerControl
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

		_m.SequencerActive(ctx)
	})
}

func Fuzz_Nosy_SequencerControl_StartSequencer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *SequencerControl
		fill_err = tp.Fill(&_m)
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
		if _m == nil {
			return
		}

		_m.StartSequencer(ctx, hash)
	})
}

func Fuzz_Nosy_SequencerControl_StopSequencer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *SequencerControl
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

		_m.StopSequencer(ctx)
	})
}

// skipping Fuzz_Nosy_SequencerControl_ConductorEnabled_Call_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_SequencerControl_ConductorEnabled_Call_Run__ because parameters include func, chan, or unsupported interface: func(ctx context.Context)

// skipping Fuzz_Nosy_SequencerControl_ConductorEnabled_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(context.Context) (bool, error)

// skipping Fuzz_Nosy_SequencerControl_Expecter_ConductorEnabled__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_SequencerControl_Expecter_LatestUnsafeBlock__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_SequencerControl_Expecter_PostUnsafePayload__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_SequencerControl_Expecter_SequencerActive__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_SequencerControl_Expecter_StartSequencer__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_SequencerControl_Expecter_StopSequencer__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_SequencerControl_LatestUnsafeBlock_Call_Return__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/eth.BlockInfo

// skipping Fuzz_Nosy_SequencerControl_LatestUnsafeBlock_Call_Run__ because parameters include func, chan, or unsupported interface: func(ctx context.Context)

// skipping Fuzz_Nosy_SequencerControl_LatestUnsafeBlock_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(context.Context) (github.com/ethereum-optimism/optimism/op-service/eth.BlockInfo, error)

// skipping Fuzz_Nosy_SequencerControl_PostUnsafePayload_Call_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_SequencerControl_PostUnsafePayload_Call_Run__ because parameters include func, chan, or unsupported interface: func(ctx context.Context, payload *github.com/ethereum-optimism/optimism/op-service/eth.ExecutionPayloadEnvelope)

// skipping Fuzz_Nosy_SequencerControl_PostUnsafePayload_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(context.Context, *github.com/ethereum-optimism/optimism/op-service/eth.ExecutionPayloadEnvelope) error

// skipping Fuzz_Nosy_SequencerControl_SequencerActive_Call_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_SequencerControl_SequencerActive_Call_Run__ because parameters include func, chan, or unsupported interface: func(ctx context.Context)

// skipping Fuzz_Nosy_SequencerControl_SequencerActive_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(context.Context) (bool, error)

// skipping Fuzz_Nosy_SequencerControl_StartSequencer_Call_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_SequencerControl_StartSequencer_Call_Run__ because parameters include func, chan, or unsupported interface: func(ctx context.Context, hash github.com/ethereum/go-ethereum/common.Hash)

// skipping Fuzz_Nosy_SequencerControl_StartSequencer_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(context.Context, github.com/ethereum/go-ethereum/common.Hash) error

// skipping Fuzz_Nosy_SequencerControl_StopSequencer_Call_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_SequencerControl_StopSequencer_Call_Run__ because parameters include func, chan, or unsupported interface: func(ctx context.Context)

// skipping Fuzz_Nosy_SequencerControl_StopSequencer_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(context.Context) (github.com/ethereum/go-ethereum/common.Hash, error)
