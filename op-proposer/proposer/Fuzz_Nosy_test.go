package proposer

import (
	"context"
	"testing"

	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum/go-ethereum/accounts/abi"
	go_fuzz_utils "github.com/trailofbits/go-fuzz-utils"
	"github.com/urfave/cli/v2"
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

func Fuzz_Nosy_CLIConfig_Check__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx *cli.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if ctx == nil {
			return
		}

		c := NewConfig(ctx)
		c.Check()
	})
}

func Fuzz_Nosy_L2OutputSubmitter_FetchCurrentBlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var setup DriverSetup
		fill_err = tp.Fill(&setup)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		l, err := NewL2OutputSubmitter(setup)
		if err != nil {
			return
		}
		l.FetchCurrentBlockNumber(ctx)
	})
}

func Fuzz_Nosy_L2OutputSubmitter_FetchDGFOutput__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var setup DriverSetup
		fill_err = tp.Fill(&setup)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		l, err := NewL2OutputSubmitter(setup)
		if err != nil {
			return
		}
		l.FetchDGFOutput(ctx)
	})
}

func Fuzz_Nosy_L2OutputSubmitter_FetchL2OOOutput__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var setup DriverSetup
		fill_err = tp.Fill(&setup)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		l, err := NewL2OutputSubmitter(setup)
		if err != nil {
			return
		}
		l.FetchL2OOOutput(ctx)
	})
}

func Fuzz_Nosy_L2OutputSubmitter_FetchOutput__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var setup DriverSetup
		fill_err = tp.Fill(&setup)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var block uint64
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}

		l, err := NewL2OutputSubmitter(setup)
		if err != nil {
			return
		}
		l.FetchOutput(ctx, block)
	})
}

func Fuzz_Nosy_L2OutputSubmitter_ProposeL2OutputDGFTxCandidate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var setup DriverSetup
		fill_err = tp.Fill(&setup)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var output *eth.OutputResponse
		fill_err = tp.Fill(&output)
		if fill_err != nil {
			return
		}
		if output == nil {
			return
		}

		l, err := NewL2OutputSubmitter(setup)
		if err != nil {
			return
		}
		l.ProposeL2OutputDGFTxCandidate(ctx, output)
	})
}

func Fuzz_Nosy_L2OutputSubmitter_ProposeL2OutputTxData__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var setup DriverSetup
		fill_err = tp.Fill(&setup)
		if fill_err != nil {
			return
		}
		var output *eth.OutputResponse
		fill_err = tp.Fill(&output)
		if fill_err != nil {
			return
		}
		if output == nil {
			return
		}

		l, err := NewL2OutputSubmitter(setup)
		if err != nil {
			return
		}
		l.ProposeL2OutputTxData(output)
	})
}

func Fuzz_Nosy_L2OutputSubmitter_StartL2OutputSubmitting__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var setup DriverSetup
		fill_err = tp.Fill(&setup)
		if fill_err != nil {
			return
		}

		l, err := NewL2OutputSubmitter(setup)
		if err != nil {
			return
		}
		l.StartL2OutputSubmitting()
	})
}

func Fuzz_Nosy_L2OutputSubmitter_StopL2OutputSubmitting__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var setup DriverSetup
		fill_err = tp.Fill(&setup)
		if fill_err != nil {
			return
		}

		l, err := NewL2OutputSubmitter(setup)
		if err != nil {
			return
		}
		l.StopL2OutputSubmitting()
	})
}

func Fuzz_Nosy_L2OutputSubmitter_StopL2OutputSubmittingIfRunning__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var setup DriverSetup
		fill_err = tp.Fill(&setup)
		if fill_err != nil {
			return
		}

		l, err := NewL2OutputSubmitter(setup)
		if err != nil {
			return
		}
		l.StopL2OutputSubmittingIfRunning()
	})
}

func Fuzz_Nosy_L2OutputSubmitter_loop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var setup DriverSetup
		fill_err = tp.Fill(&setup)
		if fill_err != nil {
			return
		}

		l, err := NewL2OutputSubmitter(setup)
		if err != nil {
			return
		}
		l.loop()
	})
}

func Fuzz_Nosy_L2OutputSubmitter_proposeOutput__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var setup DriverSetup
		fill_err = tp.Fill(&setup)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var output *eth.OutputResponse
		fill_err = tp.Fill(&output)
		if fill_err != nil {
			return
		}
		if output == nil {
			return
		}

		l, err := NewL2OutputSubmitter(setup)
		if err != nil {
			return
		}
		l.proposeOutput(ctx, output)
	})
}

func Fuzz_Nosy_L2OutputSubmitter_sendTransaction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var setup DriverSetup
		fill_err = tp.Fill(&setup)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var output *eth.OutputResponse
		fill_err = tp.Fill(&output)
		if fill_err != nil {
			return
		}
		if output == nil {
			return
		}

		l, err := NewL2OutputSubmitter(setup)
		if err != nil {
			return
		}
		l.sendTransaction(ctx, output)
	})
}

func Fuzz_Nosy_L2OutputSubmitter_waitForL1Head__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var setup DriverSetup
		fill_err = tp.Fill(&setup)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var blockNum uint64
		fill_err = tp.Fill(&blockNum)
		if fill_err != nil {
			return
		}

		l, err := NewL2OutputSubmitter(setup)
		if err != nil {
			return
		}
		l.waitForL1Head(ctx, blockNum)
	})
}

func Fuzz_Nosy_L2OutputSubmitter_waitNodeSync__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var setup DriverSetup
		fill_err = tp.Fill(&setup)
		if fill_err != nil {
			return
		}

		l, err := NewL2OutputSubmitter(setup)
		if err != nil {
			return
		}
		l.waitNodeSync()
	})
}

func Fuzz_Nosy_ProposerService_Driver__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ps *ProposerService
		fill_err = tp.Fill(&ps)
		if fill_err != nil {
			return
		}
		if ps == nil {
			return
		}

		ps.Driver()
	})
}

func Fuzz_Nosy_ProposerService_Kill__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ps *ProposerService
		fill_err = tp.Fill(&ps)
		if fill_err != nil {
			return
		}
		if ps == nil {
			return
		}

		ps.Kill()
	})
}

func Fuzz_Nosy_ProposerService_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ps *ProposerService
		fill_err = tp.Fill(&ps)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if ps == nil {
			return
		}

		ps.Start(_x2)
	})
}

func Fuzz_Nosy_ProposerService_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ps *ProposerService
		fill_err = tp.Fill(&ps)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if ps == nil {
			return
		}

		ps.Stop(ctx)
	})
}

func Fuzz_Nosy_ProposerService_Stopped__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ps *ProposerService
		fill_err = tp.Fill(&ps)
		if fill_err != nil {
			return
		}
		if ps == nil {
			return
		}

		ps.Stopped()
	})
}

func Fuzz_Nosy_ProposerService_initBalanceMonitor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ps *ProposerService
		fill_err = tp.Fill(&ps)
		if fill_err != nil {
			return
		}
		var cfg *CLIConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if ps == nil || cfg == nil {
			return
		}

		ps.initBalanceMonitor(cfg)
	})
}

func Fuzz_Nosy_ProposerService_initDGF__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ps *ProposerService
		fill_err = tp.Fill(&ps)
		if fill_err != nil {
			return
		}
		var cfg *CLIConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if ps == nil || cfg == nil {
			return
		}

		ps.initDGF(cfg)
	})
}

func Fuzz_Nosy_ProposerService_initDriver__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ps *ProposerService
		fill_err = tp.Fill(&ps)
		if fill_err != nil {
			return
		}
		if ps == nil {
			return
		}

		ps.initDriver()
	})
}

// skipping Fuzz_Nosy_ProposerService_initFromCLIConfig__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/log.Logger

func Fuzz_Nosy_ProposerService_initL2ooAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ps *ProposerService
		fill_err = tp.Fill(&ps)
		if fill_err != nil {
			return
		}
		var cfg *CLIConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if ps == nil || cfg == nil {
			return
		}

		ps.initL2ooAddress(cfg)
	})
}

func Fuzz_Nosy_ProposerService_initMetrics__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ps *ProposerService
		fill_err = tp.Fill(&ps)
		if fill_err != nil {
			return
		}
		var cfg *CLIConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if ps == nil || cfg == nil {
			return
		}

		ps.initMetrics(cfg)
	})
}

func Fuzz_Nosy_ProposerService_initMetricsServer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ps *ProposerService
		fill_err = tp.Fill(&ps)
		if fill_err != nil {
			return
		}
		var cfg *CLIConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if ps == nil || cfg == nil {
			return
		}

		ps.initMetricsServer(cfg)
	})
}

func Fuzz_Nosy_ProposerService_initPProf__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ps *ProposerService
		fill_err = tp.Fill(&ps)
		if fill_err != nil {
			return
		}
		var cfg *CLIConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if ps == nil || cfg == nil {
			return
		}

		ps.initPProf(cfg)
	})
}

func Fuzz_Nosy_ProposerService_initRPCClients__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ps *ProposerService
		fill_err = tp.Fill(&ps)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var cfg *CLIConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if ps == nil || cfg == nil {
			return
		}

		ps.initRPCClients(ctx, cfg)
	})
}

func Fuzz_Nosy_ProposerService_initRPCServer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ps *ProposerService
		fill_err = tp.Fill(&ps)
		if fill_err != nil {
			return
		}
		var cfg *CLIConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if ps == nil || cfg == nil {
			return
		}

		ps.initRPCServer(cfg)
	})
}

func Fuzz_Nosy_ProposerService_initTxManager__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ps *ProposerService
		fill_err = tp.Fill(&ps)
		if fill_err != nil {
			return
		}
		var cfg *CLIConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if ps == nil || cfg == nil {
			return
		}

		ps.initTxManager(cfg)
	})
}

// skipping Fuzz_Nosy_DGFContract_HasProposedSince__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-proposer/proposer.DGFContract

// skipping Fuzz_Nosy_DGFContract_ProposalTx__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-proposer/proposer.DGFContract

// skipping Fuzz_Nosy_DGFContract_Version__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-proposer/proposer.DGFContract

// skipping Fuzz_Nosy_L1Client_CallContract__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-proposer/proposer.L1Client

// skipping Fuzz_Nosy_L1Client_CodeAt__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-proposer/proposer.L1Client

// skipping Fuzz_Nosy_L1Client_HeaderByNumber__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-proposer/proposer.L1Client

// skipping Fuzz_Nosy_L2OOContract_NextBlockNumber__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-proposer/proposer.L2OOContract

// skipping Fuzz_Nosy_L2OOContract_Version__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-proposer/proposer.L2OOContract

// skipping Fuzz_Nosy_RollupClient_OutputAtBlock__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-proposer/proposer.RollupClient

// skipping Fuzz_Nosy_RollupClient_SyncStatus__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-proposer/proposer.RollupClient

func Fuzz_Nosy_proposeL2OutputTxData__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var abi *abi.ABI
		fill_err = tp.Fill(&abi)
		if fill_err != nil {
			return
		}
		var output *eth.OutputResponse
		fill_err = tp.Fill(&output)
		if fill_err != nil {
			return
		}
		if abi == nil || output == nil {
			return
		}

		proposeL2OutputTxData(abi, output)
	})
}
