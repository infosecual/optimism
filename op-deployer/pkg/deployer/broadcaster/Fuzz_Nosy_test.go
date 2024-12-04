package broadcaster

import (
	"context"
	"testing"

	"github.com/ethereum-optimism/optimism/op-chain-ops/script"
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

func Fuzz_Nosy_KeyedBroadcaster_Broadcast__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg KeyedBroadcasterOpts
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		t1, err := NewKeyedBroadcaster(cfg)
		if err != nil {
			return
		}
		t1.Broadcast(ctx)
	})
}

func Fuzz_Nosy_KeyedBroadcaster_Hook__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg KeyedBroadcasterOpts
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var bcast script.Broadcast
		fill_err = tp.Fill(&bcast)
		if fill_err != nil {
			return
		}

		t1, err := NewKeyedBroadcaster(cfg)
		if err != nil {
			return
		}
		t1.Hook(bcast)
	})
}

func Fuzz_Nosy_KeyedBroadcaster_broadcast__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg KeyedBroadcasterOpts
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var bcast script.Broadcast
		fill_err = tp.Fill(&bcast)
		if fill_err != nil {
			return
		}
		var blockGasLimit uint64
		fill_err = tp.Fill(&blockGasLimit)
		if fill_err != nil {
			return
		}

		t1, err := NewKeyedBroadcaster(cfg)
		if err != nil {
			return
		}
		t1.broadcast(ctx, bcast, blockGasLimit)
	})
}

func Fuzz_Nosy_discardBroadcaster_Broadcast__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *discardBroadcaster
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.Broadcast(ctx)
	})
}

func Fuzz_Nosy_discardBroadcaster_Hook__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *discardBroadcaster
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var bcast script.Broadcast
		fill_err = tp.Fill(&bcast)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.Hook(bcast)
	})
}

func Fuzz_Nosy_Broadcaster_Broadcast__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		_x1 := NoopBroadcaster()
		_x1.Broadcast(ctx)
	})
}

func Fuzz_Nosy_Broadcaster_Hook__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bcast script.Broadcast
		fill_err = tp.Fill(&bcast)
		if fill_err != nil {
			return
		}

		_x1 := NoopBroadcaster()
		_x1.Hook(bcast)
	})
}

// skipping Fuzz_Nosy_DeployerGasPriceEstimator__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/txmgr.ETHBackend

func Fuzz_Nosy_padGasLimit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d1 []byte, gasUsed uint64, creation bool, blockGasLimit uint64) {
		padGasLimit(d1, gasUsed, creation, blockGasLimit)
	})
}
