package vm

import (
	"context"
	"testing"

	"github.com/ethereum-optimism/optimism/op-challenger/game/fault/trace/utils"
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

func Fuzz_Nosy_Executor_DoGenerateProof__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *Executor
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var dir string
		fill_err = tp.Fill(&dir)
		if fill_err != nil {
			return
		}
		var begin uint64
		fill_err = tp.Fill(&begin)
		if fill_err != nil {
			return
		}
		var end uint64
		fill_err = tp.Fill(&end)
		if fill_err != nil {
			return
		}
		var extraVmArgs []string
		fill_err = tp.Fill(&extraVmArgs)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.DoGenerateProof(ctx, dir, begin, end, extraVmArgs...)
	})
}

func Fuzz_Nosy_Executor_GenerateProof__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *Executor
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var dir string
		fill_err = tp.Fill(&dir)
		if fill_err != nil {
			return
		}
		var i uint64
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.GenerateProof(ctx, dir, i)
	})
}

func Fuzz_Nosy_KonaExecutor_OracleCommand__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var dataDir string
		fill_err = tp.Fill(&dataDir)
		if fill_err != nil {
			return
		}
		var inputs utils.LocalGameInputs
		fill_err = tp.Fill(&inputs)
		if fill_err != nil {
			return
		}

		s := NewKonaExecutor()
		s.OracleCommand(cfg, dataDir, inputs)
	})
}

func Fuzz_Nosy_OpProgramServerExecutor_OracleCommand__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *OpProgramServerExecutor
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var cfg Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var dataDir string
		fill_err = tp.Fill(&dataDir)
		if fill_err != nil {
			return
		}
		var inputs utils.LocalGameInputs
		fill_err = tp.Fill(&inputs)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.OracleCommand(cfg, dataDir, inputs)
	})
}

func Fuzz_Nosy_PrestateProvider_AbsolutePreStateCommitment__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *PrestateProvider
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.AbsolutePreStateCommitment(ctx)
	})
}

func Fuzz_Nosy_PrestateProvider_PrestatePath__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *PrestateProvider
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.PrestatePath()
	})
}

// skipping Fuzz_Nosy_Metricer_RecordExecutionTime__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/trace/vm.Metricer

// skipping Fuzz_Nosy_Metricer_RecordMemoryUsed__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/trace/vm.Metricer

// skipping Fuzz_Nosy_OracleServerExecutor_OracleCommand__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/trace/vm.OracleServerExecutor

// skipping Fuzz_Nosy_StateConverter_ConvertStateToProof__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/trace/vm.StateConverter

func Fuzz_Nosy_FinalStatePath__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dir string, binarySnapshots bool) {
		FinalStatePath(dir, binarySnapshots)
	})
}

// skipping Fuzz_Nosy_FindStartingSnapshot__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/log.Logger

func Fuzz_Nosy_PreimageDir__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dir string) {
		PreimageDir(dir)
	})
}
