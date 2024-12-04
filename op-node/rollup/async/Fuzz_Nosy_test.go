package async

import (
	"context"
	"testing"

	"github.com/ethereum-optimism/optimism/op-service/eth"
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

func Fuzz_Nosy_SimpleAsyncGossiper_Clear__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *SimpleAsyncGossiper
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.Clear()
	})
}

func Fuzz_Nosy_SimpleAsyncGossiper_Get__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *SimpleAsyncGossiper
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.Get()
	})
}

func Fuzz_Nosy_SimpleAsyncGossiper_Gossip__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *SimpleAsyncGossiper
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var payload *eth.ExecutionPayloadEnvelope
		fill_err = tp.Fill(&payload)
		if fill_err != nil {
			return
		}
		if p == nil || payload == nil {
			return
		}

		p.Gossip(payload)
	})
}

func Fuzz_Nosy_SimpleAsyncGossiper_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *SimpleAsyncGossiper
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.Start()
	})
}

func Fuzz_Nosy_SimpleAsyncGossiper_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *SimpleAsyncGossiper
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.Stop()
	})
}

func Fuzz_Nosy_SimpleAsyncGossiper_clearPayload__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *SimpleAsyncGossiper
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.clearPayload()
	})
}

// skipping Fuzz_Nosy_SimpleAsyncGossiper_getPayload__ because parameters include func, chan, or unsupported interface: chan *github.com/ethereum-optimism/optimism/op-service/eth.ExecutionPayloadEnvelope

func Fuzz_Nosy_SimpleAsyncGossiper_gossip__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *SimpleAsyncGossiper
		fill_err = tp.Fill(&p)
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
		if p == nil || payload == nil {
			return
		}

		p.gossip(ctx, payload)
	})
}

// skipping Fuzz_Nosy_AsyncGossiper_Clear__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/async.AsyncGossiper

// skipping Fuzz_Nosy_AsyncGossiper_Get__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/async.AsyncGossiper

// skipping Fuzz_Nosy_AsyncGossiper_Gossip__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/async.AsyncGossiper

// skipping Fuzz_Nosy_AsyncGossiper_Start__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/async.AsyncGossiper

// skipping Fuzz_Nosy_AsyncGossiper_Stop__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/async.AsyncGossiper

// skipping Fuzz_Nosy_Metrics_RecordPublishingError__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/async.Metrics

// skipping Fuzz_Nosy_Network_PublishL2Payload__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/async.Network

func Fuzz_Nosy_NoOpGossiper_Clear__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 NoOpGossiper
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}

		_x1.Clear()
	})
}

func Fuzz_Nosy_NoOpGossiper_Get__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 NoOpGossiper
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}

		_x1.Get()
	})
}

func Fuzz_Nosy_NoOpGossiper_Gossip__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 NoOpGossiper
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var payload *eth.ExecutionPayloadEnvelope
		fill_err = tp.Fill(&payload)
		if fill_err != nil {
			return
		}
		if payload == nil {
			return
		}

		_x1.Gossip(payload)
	})
}

func Fuzz_Nosy_NoOpGossiper_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 NoOpGossiper
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}

		_x1.Start()
	})
}

func Fuzz_Nosy_NoOpGossiper_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 NoOpGossiper
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}

		_x1.Stop()
	})
}
