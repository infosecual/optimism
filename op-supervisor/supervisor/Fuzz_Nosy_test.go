package supervisor

import (
	"context"
	"testing"

	"github.com/ethereum-optimism/optimism/op-supervisor/config"
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

func Fuzz_Nosy_SupervisorService_RPC__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var su *SupervisorService
		fill_err = tp.Fill(&su)
		if fill_err != nil {
			return
		}
		if su == nil {
			return
		}

		su.RPC()
	})
}

func Fuzz_Nosy_SupervisorService_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var su *SupervisorService
		fill_err = tp.Fill(&su)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if su == nil {
			return
		}

		su.Start(ctx)
	})
}

func Fuzz_Nosy_SupervisorService_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var su *SupervisorService
		fill_err = tp.Fill(&su)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if su == nil {
			return
		}

		su.Stop(ctx)
	})
}

func Fuzz_Nosy_SupervisorService_Stopped__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var su *SupervisorService
		fill_err = tp.Fill(&su)
		if fill_err != nil {
			return
		}
		if su == nil {
			return
		}

		su.Stopped()
	})
}

func Fuzz_Nosy_SupervisorService_initBackend__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var su *SupervisorService
		fill_err = tp.Fill(&su)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var cfg *config.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if su == nil || cfg == nil {
			return
		}

		su.initBackend(ctx, cfg)
	})
}

func Fuzz_Nosy_SupervisorService_initFromCLIConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var su *SupervisorService
		fill_err = tp.Fill(&su)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var cfg *config.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if su == nil || cfg == nil {
			return
		}

		su.initFromCLIConfig(ctx, cfg)
	})
}

func Fuzz_Nosy_SupervisorService_initMetrics__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var su *SupervisorService
		fill_err = tp.Fill(&su)
		if fill_err != nil {
			return
		}
		var cfg *config.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if su == nil || cfg == nil {
			return
		}

		su.initMetrics(cfg)
	})
}

func Fuzz_Nosy_SupervisorService_initMetricsServer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var su *SupervisorService
		fill_err = tp.Fill(&su)
		if fill_err != nil {
			return
		}
		var cfg *config.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if su == nil || cfg == nil {
			return
		}

		su.initMetricsServer(cfg)
	})
}

func Fuzz_Nosy_SupervisorService_initPProf__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var su *SupervisorService
		fill_err = tp.Fill(&su)
		if fill_err != nil {
			return
		}
		var cfg *config.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if su == nil || cfg == nil {
			return
		}

		su.initPProf(cfg)
	})
}

func Fuzz_Nosy_SupervisorService_initRPCServer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var su *SupervisorService
		fill_err = tp.Fill(&su)
		if fill_err != nil {
			return
		}
		var cfg *config.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if su == nil || cfg == nil {
			return
		}

		su.initRPCServer(cfg)
	})
}
