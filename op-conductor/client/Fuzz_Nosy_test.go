package client

import (
	"context"
	"testing"

	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum-optimism/optimism/op-service/sources"
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

func Fuzz_Nosy_sequencerController_ConductorEnabled__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *sequencerController
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.ConductorEnabled(ctx)
	})
}

func Fuzz_Nosy_sequencerController_LatestUnsafeBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *sequencerController
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.LatestUnsafeBlock(ctx)
	})
}

func Fuzz_Nosy_sequencerController_PostUnsafePayload__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *sequencerController
		fill_err = tp.Fill(&s)
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
		if s == nil || payload == nil {
			return
		}

		s.PostUnsafePayload(ctx, payload)
	})
}

func Fuzz_Nosy_sequencerController_SequencerActive__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *sequencerController
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.SequencerActive(ctx)
	})
}

func Fuzz_Nosy_sequencerController_StartSequencer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *sequencerController
		fill_err = tp.Fill(&s)
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
		if s == nil {
			return
		}

		s.StartSequencer(ctx, hash)
	})
}

func Fuzz_Nosy_sequencerController_StopSequencer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *sequencerController
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.StopSequencer(ctx)
	})
}

func Fuzz_Nosy_SequencerControl_ConductorEnabled__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var exec *sources.EthClient
		fill_err = tp.Fill(&exec)
		if fill_err != nil {
			return
		}
		var node *sources.RollupClient
		fill_err = tp.Fill(&node)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if exec == nil || node == nil {
			return
		}

		_x1 := NewSequencerControl(exec, node)
		_x1.ConductorEnabled(ctx)
	})
}

func Fuzz_Nosy_SequencerControl_LatestUnsafeBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var exec *sources.EthClient
		fill_err = tp.Fill(&exec)
		if fill_err != nil {
			return
		}
		var node *sources.RollupClient
		fill_err = tp.Fill(&node)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if exec == nil || node == nil {
			return
		}

		_x1 := NewSequencerControl(exec, node)
		_x1.LatestUnsafeBlock(ctx)
	})
}

func Fuzz_Nosy_SequencerControl_PostUnsafePayload__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var exec *sources.EthClient
		fill_err = tp.Fill(&exec)
		if fill_err != nil {
			return
		}
		var node *sources.RollupClient
		fill_err = tp.Fill(&node)
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
		if exec == nil || node == nil || payload == nil {
			return
		}

		_x1 := NewSequencerControl(exec, node)
		_x1.PostUnsafePayload(ctx, payload)
	})
}

func Fuzz_Nosy_SequencerControl_SequencerActive__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var exec *sources.EthClient
		fill_err = tp.Fill(&exec)
		if fill_err != nil {
			return
		}
		var node *sources.RollupClient
		fill_err = tp.Fill(&node)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if exec == nil || node == nil {
			return
		}

		_x1 := NewSequencerControl(exec, node)
		_x1.SequencerActive(ctx)
	})
}

func Fuzz_Nosy_SequencerControl_StartSequencer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var exec *sources.EthClient
		fill_err = tp.Fill(&exec)
		if fill_err != nil {
			return
		}
		var node *sources.RollupClient
		fill_err = tp.Fill(&node)
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
		if exec == nil || node == nil {
			return
		}

		_x1 := NewSequencerControl(exec, node)
		_x1.StartSequencer(ctx, hash)
	})
}

func Fuzz_Nosy_SequencerControl_StopSequencer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var exec *sources.EthClient
		fill_err = tp.Fill(&exec)
		if fill_err != nil {
			return
		}
		var node *sources.RollupClient
		fill_err = tp.Fill(&node)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if exec == nil || node == nil {
			return
		}

		_x1 := NewSequencerControl(exec, node)
		_x1.StopSequencer(ctx)
	})
}
