package health

import (
	"context"
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

func Fuzz_Nosy_SequencerHealthMonitor_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hm *SequencerHealthMonitor
		fill_err = tp.Fill(&hm)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if hm == nil {
			return
		}

		hm.Start(ctx)
	})
}

func Fuzz_Nosy_SequencerHealthMonitor_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hm *SequencerHealthMonitor
		fill_err = tp.Fill(&hm)
		if fill_err != nil {
			return
		}
		if hm == nil {
			return
		}

		hm.Stop()
	})
}

func Fuzz_Nosy_SequencerHealthMonitor_Subscribe__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hm *SequencerHealthMonitor
		fill_err = tp.Fill(&hm)
		if fill_err != nil {
			return
		}
		if hm == nil {
			return
		}

		hm.Subscribe()
	})
}

func Fuzz_Nosy_SequencerHealthMonitor_healthCheck__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hm *SequencerHealthMonitor
		fill_err = tp.Fill(&hm)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if hm == nil {
			return
		}

		hm.healthCheck(ctx)
	})
}

func Fuzz_Nosy_SequencerHealthMonitor_loop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hm *SequencerHealthMonitor
		fill_err = tp.Fill(&hm)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if hm == nil {
			return
		}

		hm.loop(ctx)
	})
}

// skipping Fuzz_Nosy_HealthMonitor_Start__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-conductor/health.HealthMonitor

// skipping Fuzz_Nosy_HealthMonitor_Stop__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-conductor/health.HealthMonitor

// skipping Fuzz_Nosy_HealthMonitor_Subscribe__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-conductor/health.HealthMonitor

func Fuzz_Nosy_calculateTimeDiff__(f *testing.F) {
	f.Fuzz(func(t *testing.T, now uint64, then uint64) {
		calculateTimeDiff(now, then)
	})
}
