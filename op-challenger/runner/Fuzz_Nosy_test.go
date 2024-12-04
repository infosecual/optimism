package runner

import (
	"context"
	"testing"
	"time"

	"github.com/ethereum-optimism/optimism/op-challenger/game/fault/types"
	"github.com/ethereum-optimism/optimism/op-service/metrics"
	"github.com/ethereum-optimism/optimism/op-service/sources"
	"github.com/ethereum-optimism/optimism/op-service/sources/batching"
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

func Fuzz_Nosy_Metrics_RecordFailure__(f *testing.F) {
	f.Fuzz(func(t *testing.T, vmType string) {
		m := NewMetrics()
		m.RecordFailure(vmType)
	})
}

func Fuzz_Nosy_Metrics_RecordInvalid__(f *testing.F) {
	f.Fuzz(func(t *testing.T, vmType string) {
		m := NewMetrics()
		m.RecordInvalid(vmType)
	})
}

func Fuzz_Nosy_Metrics_RecordSuccess__(f *testing.F) {
	f.Fuzz(func(t *testing.T, vmType string) {
		m := NewMetrics()
		m.RecordSuccess(vmType)
	})
}

func Fuzz_Nosy_Metrics_RecordVmExecutionTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var vmType string
		fill_err = tp.Fill(&vmType)
		if fill_err != nil {
			return
		}
		var dur time.Duration
		fill_err = tp.Fill(&dur)
		if fill_err != nil {
			return
		}

		m := NewMetrics()
		m.RecordVmExecutionTime(vmType, dur)
	})
}

func Fuzz_Nosy_Metrics_RecordVmMemoryUsed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, vmType string, memoryUsed uint64) {
		m := NewMetrics()
		m.RecordVmMemoryUsed(vmType, memoryUsed)
	})
}

func Fuzz_Nosy_Runner_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *Runner
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.Start(ctx)
	})
}

func Fuzz_Nosy_Runner_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *Runner
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.Stop(ctx)
	})
}

func Fuzz_Nosy_Runner_Stopped__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *Runner
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.Stopped()
	})
}

func Fuzz_Nosy_Runner_createGameInputs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *Runner
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var client *sources.RollupClient
		fill_err = tp.Fill(&client)
		if fill_err != nil {
			return
		}
		if r == nil || client == nil {
			return
		}

		r.createGameInputs(ctx, client)
	})
}

func Fuzz_Nosy_Runner_findL2BlockNumberToDispute__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *Runner
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var client *sources.RollupClient
		fill_err = tp.Fill(&client)
		if fill_err != nil {
			return
		}
		var l1HeadNum uint64
		fill_err = tp.Fill(&l1HeadNum)
		if fill_err != nil {
			return
		}
		var l2BlockNum uint64
		fill_err = tp.Fill(&l2BlockNum)
		if fill_err != nil {
			return
		}
		if r == nil || client == nil {
			return
		}

		r.findL2BlockNumberToDispute(ctx, client, l1HeadNum, l2BlockNum)
	})
}

func Fuzz_Nosy_Runner_getPrestateHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *Runner
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var traceType types.TraceType
		fill_err = tp.Fill(&traceType)
		if fill_err != nil {
			return
		}
		var caller *batching.MultiCaller
		fill_err = tp.Fill(&caller)
		if fill_err != nil {
			return
		}
		if r == nil || caller == nil {
			return
		}

		r.getPrestateHash(ctx, traceType, caller)
	})
}

func Fuzz_Nosy_Runner_initMetricsServer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *Runner
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var cfg *metrics.CLIConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if r == nil || cfg == nil {
			return
		}

		r.initMetricsServer(cfg)
	})
}

func Fuzz_Nosy_Runner_loop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *Runner
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var runConfig RunConfig
		fill_err = tp.Fill(&runConfig)
		if fill_err != nil {
			return
		}
		var client *sources.RollupClient
		fill_err = tp.Fill(&client)
		if fill_err != nil {
			return
		}
		var caller *batching.MultiCaller
		fill_err = tp.Fill(&caller)
		if fill_err != nil {
			return
		}
		if r == nil || client == nil || caller == nil {
			return
		}

		r.loop(ctx, runConfig, client, caller)
	})
}

func Fuzz_Nosy_Runner_prepDatadir__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *Runner
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.prepDatadir(name)
	})
}

func Fuzz_Nosy_Runner_runAndRecordOnce__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *Runner
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var runConfig RunConfig
		fill_err = tp.Fill(&runConfig)
		if fill_err != nil {
			return
		}
		var client *sources.RollupClient
		fill_err = tp.Fill(&client)
		if fill_err != nil {
			return
		}
		var caller *batching.MultiCaller
		fill_err = tp.Fill(&caller)
		if fill_err != nil {
			return
		}
		if r == nil || client == nil || caller == nil {
			return
		}

		r.runAndRecordOnce(ctx, runConfig, client, caller)
	})
}

// skipping Fuzz_Nosy_Runner_runOnce__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/log.Logger

// skipping Fuzz_Nosy_Metricer_RecordFailure__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/runner.Metricer

// skipping Fuzz_Nosy_Metricer_RecordInvalid__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/runner.Metricer

// skipping Fuzz_Nosy_Metricer_RecordSuccess__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/runner.Metricer

// skipping Fuzz_Nosy_Metricer_RecordVmExecutionTime__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/runner.Metricer

// skipping Fuzz_Nosy_Metricer_RecordVmMemoryUsed__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/runner.Metricer

// skipping Fuzz_Nosy_getPrestate__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/trace/vm.StateConverter
