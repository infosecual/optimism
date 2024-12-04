package batcher

import (
	"context"
	"testing"

	"github.com/ethereum-optimism/optimism/op-node/rollup"
	"github.com/ethereum-optimism/optimism/op-node/rollup/derive"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum-optimism/optimism/op-service/txmgr"
	"github.com/ethereum/go-ethereum/core/types"
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

func Fuzz_Nosy_BatchSubmitter_StartBatchSubmitting__(f *testing.F) {
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

		l := NewBatchSubmitter(setup)
		l.StartBatchSubmitting()
	})
}

func Fuzz_Nosy_BatchSubmitter_StopBatchSubmitting__(f *testing.F) {
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

		l := NewBatchSubmitter(setup)
		l.StopBatchSubmitting(ctx)
	})
}

func Fuzz_Nosy_BatchSubmitter_StopBatchSubmittingIfRunning__(f *testing.F) {
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

		l := NewBatchSubmitter(setup)
		l.StopBatchSubmittingIfRunning(ctx)
	})
}

func Fuzz_Nosy_BatchSubmitter_blobTxCandidate__(f *testing.F) {
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
		var d2 txData
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}

		l := NewBatchSubmitter(setup)
		l.blobTxCandidate(d2)
	})
}

func Fuzz_Nosy_BatchSubmitter_calculateL2BlockRangeToStore__(f *testing.F) {
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
		var syncStatus eth.SyncStatus
		fill_err = tp.Fill(&syncStatus)
		if fill_err != nil {
			return
		}

		l := NewBatchSubmitter(setup)
		l.calculateL2BlockRangeToStore(syncStatus)
	})
}

func Fuzz_Nosy_BatchSubmitter_calldataTxCandidate__(f *testing.F) {
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
		var d2 []byte
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}

		l := NewBatchSubmitter(setup)
		l.calldataTxCandidate(d2)
	})
}

// skipping Fuzz_Nosy_BatchSubmitter_cancelBlockingTx__ because parameters include func, chan, or unsupported interface: chan github.com/ethereum-optimism/optimism/op-service/txmgr.TxReceipt[github.com/ethereum-optimism/optimism/op-batcher/batcher.txRef]

// skipping Fuzz_Nosy_BatchSubmitter_checkTxpool__ because parameters include func, chan, or unsupported interface: chan github.com/ethereum-optimism/optimism/op-service/txmgr.TxReceipt[github.com/ethereum-optimism/optimism/op-batcher/batcher.txRef]

func Fuzz_Nosy_BatchSubmitter_clearState__(f *testing.F) {
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

		l := NewBatchSubmitter(setup)
		l.clearState(ctx)
	})
}

func Fuzz_Nosy_BatchSubmitter_getSyncStatus__(f *testing.F) {
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

		l := NewBatchSubmitter(setup)
		l.getSyncStatus(ctx)
	})
}

func Fuzz_Nosy_BatchSubmitter_handleReceipt__(f *testing.F) {
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
		var r txmgr.TxReceipt[txRef]
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}

		l := NewBatchSubmitter(setup)
		l.handleReceipt(r)
	})
}

func Fuzz_Nosy_BatchSubmitter_l1Tip__(f *testing.F) {
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

		l := NewBatchSubmitter(setup)
		l.l1Tip(ctx)
	})
}

func Fuzz_Nosy_BatchSubmitter_loadBlockIntoState__(f *testing.F) {
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
		var blockNumber uint64
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}

		l := NewBatchSubmitter(setup)
		l.loadBlockIntoState(ctx, blockNumber)
	})
}

func Fuzz_Nosy_BatchSubmitter_loadBlocksIntoState__(f *testing.F) {
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
		var syncStatus eth.SyncStatus
		fill_err = tp.Fill(&syncStatus)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		l := NewBatchSubmitter(setup)
		l.loadBlocksIntoState(syncStatus, ctx)
	})
}

// skipping Fuzz_Nosy_BatchSubmitter_mainLoop__ because parameters include func, chan, or unsupported interface: chan github.com/ethereum-optimism/optimism/op-service/txmgr.TxReceipt[github.com/ethereum-optimism/optimism/op-batcher/batcher.txRef]

// skipping Fuzz_Nosy_BatchSubmitter_processReceiptsLoop__ because parameters include func, chan, or unsupported interface: chan github.com/ethereum-optimism/optimism/op-service/txmgr.TxReceipt[github.com/ethereum-optimism/optimism/op-batcher/batcher.txRef]

// skipping Fuzz_Nosy_BatchSubmitter_publishStateToL1__ because parameters include func, chan, or unsupported interface: chan github.com/ethereum-optimism/optimism/op-service/txmgr.TxReceipt[github.com/ethereum-optimism/optimism/op-batcher/batcher.txRef]

// skipping Fuzz_Nosy_BatchSubmitter_publishToAltDAAndL1__ because parameters include func, chan, or unsupported interface: chan github.com/ethereum-optimism/optimism/op-service/txmgr.TxReceipt[github.com/ethereum-optimism/optimism/op-batcher/batcher.txRef]

// skipping Fuzz_Nosy_BatchSubmitter_publishTxToL1__ because parameters include func, chan, or unsupported interface: chan github.com/ethereum-optimism/optimism/op-service/txmgr.TxReceipt[github.com/ethereum-optimism/optimism/op-batcher/batcher.txRef]

func Fuzz_Nosy_BatchSubmitter_recordConfirmedTx__(f *testing.F) {
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
		var id txID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var receipt *types.Receipt
		fill_err = tp.Fill(&receipt)
		if fill_err != nil {
			return
		}
		if receipt == nil {
			return
		}

		l := NewBatchSubmitter(setup)
		l.recordConfirmedTx(id, receipt)
	})
}

// skipping Fuzz_Nosy_BatchSubmitter_recordFailedDARequest__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_BatchSubmitter_recordFailedTx__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_BatchSubmitter_recordL1Tip__(f *testing.F) {
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
		var l1tip eth.L1BlockRef
		fill_err = tp.Fill(&l1tip)
		if fill_err != nil {
			return
		}

		l := NewBatchSubmitter(setup)
		l.recordL1Tip(l1tip)
	})
}

func Fuzz_Nosy_BatchSubmitter_safeL1Origin__(f *testing.F) {
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

		l := NewBatchSubmitter(setup)
		l.safeL1Origin(ctx)
	})
}

// skipping Fuzz_Nosy_BatchSubmitter_sendTransaction__ because parameters include func, chan, or unsupported interface: chan github.com/ethereum-optimism/optimism/op-service/txmgr.TxReceipt[github.com/ethereum-optimism/optimism/op-batcher/batcher.txRef]

// skipping Fuzz_Nosy_BatchSubmitter_sendTx__ because parameters include func, chan, or unsupported interface: chan github.com/ethereum-optimism/optimism/op-service/txmgr.TxReceipt[github.com/ethereum-optimism/optimism/op-batcher/batcher.txRef]

func Fuzz_Nosy_BatchSubmitter_setTxPoolState__(f *testing.F) {
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
		var txPoolState TxPoolState
		fill_err = tp.Fill(&txPoolState)
		if fill_err != nil {
			return
		}
		var txPoolBlockedBlob bool
		fill_err = tp.Fill(&txPoolBlockedBlob)
		if fill_err != nil {
			return
		}

		l := NewBatchSubmitter(setup)
		l.setTxPoolState(txPoolState, txPoolBlockedBlob)
	})
}

func Fuzz_Nosy_BatchSubmitter_throttlingLoop__(f *testing.F) {
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

		l := NewBatchSubmitter(setup)
		l.throttlingLoop(ctx)
	})
}

func Fuzz_Nosy_BatchSubmitter_waitForL2Genesis__(f *testing.F) {
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

		l := NewBatchSubmitter(setup)
		l.waitForL2Genesis()
	})
}

func Fuzz_Nosy_BatchSubmitter_waitNodeSync__(f *testing.F) {
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

		l := NewBatchSubmitter(setup)
		l.waitNodeSync()
	})
}

func Fuzz_Nosy_BatchSubmitter_waitNodeSyncAndClearState__(f *testing.F) {
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

		l := NewBatchSubmitter(setup)
		l.waitNodeSyncAndClearState()
	})
}

func Fuzz_Nosy_BatcherService_Kill__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bs *BatcherService
		fill_err = tp.Fill(&bs)
		if fill_err != nil {
			return
		}
		if bs == nil {
			return
		}

		bs.Kill()
	})
}

func Fuzz_Nosy_BatcherService_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bs *BatcherService
		fill_err = tp.Fill(&bs)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if bs == nil {
			return
		}

		bs.Start(_x2)
	})
}

func Fuzz_Nosy_BatcherService_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bs *BatcherService
		fill_err = tp.Fill(&bs)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if bs == nil {
			return
		}

		bs.Stop(ctx)
	})
}

func Fuzz_Nosy_BatcherService_Stopped__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bs *BatcherService
		fill_err = tp.Fill(&bs)
		if fill_err != nil {
			return
		}
		if bs == nil {
			return
		}

		bs.Stopped()
	})
}

func Fuzz_Nosy_BatcherService_TestDriver__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bs *BatcherService
		fill_err = tp.Fill(&bs)
		if fill_err != nil {
			return
		}
		if bs == nil {
			return
		}

		bs.TestDriver()
	})
}

func Fuzz_Nosy_BatcherService_ThrottlingTestDriver__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bs *BatcherService
		fill_err = tp.Fill(&bs)
		if fill_err != nil {
			return
		}
		if bs == nil {
			return
		}

		bs.ThrottlingTestDriver()
	})
}

func Fuzz_Nosy_BatcherService_initAltDA__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bs *BatcherService
		fill_err = tp.Fill(&bs)
		if fill_err != nil {
			return
		}
		var cfg *CLIConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if bs == nil || cfg == nil {
			return
		}

		bs.initAltDA(cfg)
	})
}

func Fuzz_Nosy_BatcherService_initBalanceMonitor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bs *BatcherService
		fill_err = tp.Fill(&bs)
		if fill_err != nil {
			return
		}
		var cfg *CLIConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if bs == nil || cfg == nil {
			return
		}

		bs.initBalanceMonitor(cfg)
	})
}

func Fuzz_Nosy_BatcherService_initChannelConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bs *BatcherService
		fill_err = tp.Fill(&bs)
		if fill_err != nil {
			return
		}
		var cfg *CLIConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if bs == nil || cfg == nil {
			return
		}

		bs.initChannelConfig(cfg)
	})
}

// skipping Fuzz_Nosy_BatcherService_initDriver__ because parameters include func, chan, or unsupported interface: []github.com/ethereum-optimism/optimism/op-batcher/batcher.DriverSetupOption

// skipping Fuzz_Nosy_BatcherService_initFromCLIConfig__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/log.Logger

func Fuzz_Nosy_BatcherService_initMetrics__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bs *BatcherService
		fill_err = tp.Fill(&bs)
		if fill_err != nil {
			return
		}
		var cfg *CLIConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if bs == nil || cfg == nil {
			return
		}

		bs.initMetrics(cfg)
	})
}

func Fuzz_Nosy_BatcherService_initMetricsServer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bs *BatcherService
		fill_err = tp.Fill(&bs)
		if fill_err != nil {
			return
		}
		var cfg *CLIConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if bs == nil || cfg == nil {
			return
		}

		bs.initMetricsServer(cfg)
	})
}

func Fuzz_Nosy_BatcherService_initPProf__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bs *BatcherService
		fill_err = tp.Fill(&bs)
		if fill_err != nil {
			return
		}
		var cfg *CLIConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if bs == nil || cfg == nil {
			return
		}

		bs.initPProf(cfg)
	})
}

func Fuzz_Nosy_BatcherService_initRPCClients__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bs *BatcherService
		fill_err = tp.Fill(&bs)
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
		if bs == nil || cfg == nil {
			return
		}

		bs.initRPCClients(ctx, cfg)
	})
}

func Fuzz_Nosy_BatcherService_initRPCServer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bs *BatcherService
		fill_err = tp.Fill(&bs)
		if fill_err != nil {
			return
		}
		var cfg *CLIConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if bs == nil || cfg == nil {
			return
		}

		bs.initRPCServer(cfg)
	})
}

func Fuzz_Nosy_BatcherService_initRollupConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bs *BatcherService
		fill_err = tp.Fill(&bs)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if bs == nil {
			return
		}

		bs.initRollupConfig(ctx)
	})
}

func Fuzz_Nosy_BatcherService_initTxManager__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bs *BatcherService
		fill_err = tp.Fill(&bs)
		if fill_err != nil {
			return
		}
		var cfg *CLIConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if bs == nil || cfg == nil {
			return
		}

		bs.initTxManager(cfg)
	})
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

func Fuzz_Nosy_ChannelBuilder_AddBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg ChannelConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var rollupCfg *rollup.Config
		fill_err = tp.Fill(&rollupCfg)
		if fill_err != nil {
			return
		}
		var latestL1OriginBlockNum uint64
		fill_err = tp.Fill(&latestL1OriginBlockNum)
		if fill_err != nil {
			return
		}
		var block *types.Block
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		if rollupCfg == nil || block == nil {
			return
		}

		c, err := NewChannelBuilder(cfg, rollupCfg, latestL1OriginBlockNum)
		if err != nil {
			return
		}
		c.AddBlock(block)
	})
}

func Fuzz_Nosy_ChannelBuilder_Blocks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg ChannelConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var rollupCfg *rollup.Config
		fill_err = tp.Fill(&rollupCfg)
		if fill_err != nil {
			return
		}
		var latestL1OriginBlockNum uint64
		fill_err = tp.Fill(&latestL1OriginBlockNum)
		if fill_err != nil {
			return
		}
		if rollupCfg == nil {
			return
		}

		c, err := NewChannelBuilder(cfg, rollupCfg, latestL1OriginBlockNum)
		if err != nil {
			return
		}
		c.Blocks()
	})
}

func Fuzz_Nosy_ChannelBuilder_CheckTimeout__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg ChannelConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var rollupCfg *rollup.Config
		fill_err = tp.Fill(&rollupCfg)
		if fill_err != nil {
			return
		}
		var latestL1OriginBlockNum uint64
		fill_err = tp.Fill(&latestL1OriginBlockNum)
		if fill_err != nil {
			return
		}
		var l1BlockNum uint64
		fill_err = tp.Fill(&l1BlockNum)
		if fill_err != nil {
			return
		}
		if rollupCfg == nil {
			return
		}

		c, err := NewChannelBuilder(cfg, rollupCfg, latestL1OriginBlockNum)
		if err != nil {
			return
		}
		c.CheckTimeout(l1BlockNum)
	})
}

func Fuzz_Nosy_ChannelBuilder_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg ChannelConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var rollupCfg *rollup.Config
		fill_err = tp.Fill(&rollupCfg)
		if fill_err != nil {
			return
		}
		var latestL1OriginBlockNum uint64
		fill_err = tp.Fill(&latestL1OriginBlockNum)
		if fill_err != nil {
			return
		}
		if rollupCfg == nil {
			return
		}

		c, err := NewChannelBuilder(cfg, rollupCfg, latestL1OriginBlockNum)
		if err != nil {
			return
		}
		c.Close()
	})
}

func Fuzz_Nosy_ChannelBuilder_FramePublished__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg ChannelConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var rollupCfg *rollup.Config
		fill_err = tp.Fill(&rollupCfg)
		if fill_err != nil {
			return
		}
		var latestL1OriginBlockNum uint64
		fill_err = tp.Fill(&latestL1OriginBlockNum)
		if fill_err != nil {
			return
		}
		var l1BlockNum uint64
		fill_err = tp.Fill(&l1BlockNum)
		if fill_err != nil {
			return
		}
		if rollupCfg == nil {
			return
		}

		c, err := NewChannelBuilder(cfg, rollupCfg, latestL1OriginBlockNum)
		if err != nil {
			return
		}
		c.FramePublished(l1BlockNum)
	})
}

func Fuzz_Nosy_ChannelBuilder_FullErr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg ChannelConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var rollupCfg *rollup.Config
		fill_err = tp.Fill(&rollupCfg)
		if fill_err != nil {
			return
		}
		var latestL1OriginBlockNum uint64
		fill_err = tp.Fill(&latestL1OriginBlockNum)
		if fill_err != nil {
			return
		}
		if rollupCfg == nil {
			return
		}

		c, err := NewChannelBuilder(cfg, rollupCfg, latestL1OriginBlockNum)
		if err != nil {
			return
		}
		c.FullErr()
	})
}

func Fuzz_Nosy_ChannelBuilder_HasPendingFrame__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg ChannelConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var rollupCfg *rollup.Config
		fill_err = tp.Fill(&rollupCfg)
		if fill_err != nil {
			return
		}
		var latestL1OriginBlockNum uint64
		fill_err = tp.Fill(&latestL1OriginBlockNum)
		if fill_err != nil {
			return
		}
		if rollupCfg == nil {
			return
		}

		c, err := NewChannelBuilder(cfg, rollupCfg, latestL1OriginBlockNum)
		if err != nil {
			return
		}
		c.HasPendingFrame()
	})
}

func Fuzz_Nosy_ChannelBuilder_ID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg ChannelConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var rollupCfg *rollup.Config
		fill_err = tp.Fill(&rollupCfg)
		if fill_err != nil {
			return
		}
		var latestL1OriginBlockNum uint64
		fill_err = tp.Fill(&latestL1OriginBlockNum)
		if fill_err != nil {
			return
		}
		if rollupCfg == nil {
			return
		}

		c, err := NewChannelBuilder(cfg, rollupCfg, latestL1OriginBlockNum)
		if err != nil {
			return
		}
		c.ID()
	})
}

func Fuzz_Nosy_ChannelBuilder_InputBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg ChannelConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var rollupCfg *rollup.Config
		fill_err = tp.Fill(&rollupCfg)
		if fill_err != nil {
			return
		}
		var latestL1OriginBlockNum uint64
		fill_err = tp.Fill(&latestL1OriginBlockNum)
		if fill_err != nil {
			return
		}
		if rollupCfg == nil {
			return
		}

		c, err := NewChannelBuilder(cfg, rollupCfg, latestL1OriginBlockNum)
		if err != nil {
			return
		}
		c.InputBytes()
	})
}

func Fuzz_Nosy_ChannelBuilder_IsFull__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg ChannelConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var rollupCfg *rollup.Config
		fill_err = tp.Fill(&rollupCfg)
		if fill_err != nil {
			return
		}
		var latestL1OriginBlockNum uint64
		fill_err = tp.Fill(&latestL1OriginBlockNum)
		if fill_err != nil {
			return
		}
		if rollupCfg == nil {
			return
		}

		c, err := NewChannelBuilder(cfg, rollupCfg, latestL1OriginBlockNum)
		if err != nil {
			return
		}
		c.IsFull()
	})
}

func Fuzz_Nosy_ChannelBuilder_LatestL1Origin__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg ChannelConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var rollupCfg *rollup.Config
		fill_err = tp.Fill(&rollupCfg)
		if fill_err != nil {
			return
		}
		var latestL1OriginBlockNum uint64
		fill_err = tp.Fill(&latestL1OriginBlockNum)
		if fill_err != nil {
			return
		}
		if rollupCfg == nil {
			return
		}

		c, err := NewChannelBuilder(cfg, rollupCfg, latestL1OriginBlockNum)
		if err != nil {
			return
		}
		c.LatestL1Origin()
	})
}

func Fuzz_Nosy_ChannelBuilder_LatestL2__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg ChannelConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var rollupCfg *rollup.Config
		fill_err = tp.Fill(&rollupCfg)
		if fill_err != nil {
			return
		}
		var latestL1OriginBlockNum uint64
		fill_err = tp.Fill(&latestL1OriginBlockNum)
		if fill_err != nil {
			return
		}
		if rollupCfg == nil {
			return
		}

		c, err := NewChannelBuilder(cfg, rollupCfg, latestL1OriginBlockNum)
		if err != nil {
			return
		}
		c.LatestL2()
	})
}

func Fuzz_Nosy_ChannelBuilder_NextFrame__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg ChannelConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var rollupCfg *rollup.Config
		fill_err = tp.Fill(&rollupCfg)
		if fill_err != nil {
			return
		}
		var latestL1OriginBlockNum uint64
		fill_err = tp.Fill(&latestL1OriginBlockNum)
		if fill_err != nil {
			return
		}
		if rollupCfg == nil {
			return
		}

		c, err := NewChannelBuilder(cfg, rollupCfg, latestL1OriginBlockNum)
		if err != nil {
			return
		}
		c.NextFrame()
	})
}

func Fuzz_Nosy_ChannelBuilder_OldestL1Origin__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg ChannelConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var rollupCfg *rollup.Config
		fill_err = tp.Fill(&rollupCfg)
		if fill_err != nil {
			return
		}
		var latestL1OriginBlockNum uint64
		fill_err = tp.Fill(&latestL1OriginBlockNum)
		if fill_err != nil {
			return
		}
		if rollupCfg == nil {
			return
		}

		c, err := NewChannelBuilder(cfg, rollupCfg, latestL1OriginBlockNum)
		if err != nil {
			return
		}
		c.OldestL1Origin()
	})
}

func Fuzz_Nosy_ChannelBuilder_OldestL2__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg ChannelConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var rollupCfg *rollup.Config
		fill_err = tp.Fill(&rollupCfg)
		if fill_err != nil {
			return
		}
		var latestL1OriginBlockNum uint64
		fill_err = tp.Fill(&latestL1OriginBlockNum)
		if fill_err != nil {
			return
		}
		if rollupCfg == nil {
			return
		}

		c, err := NewChannelBuilder(cfg, rollupCfg, latestL1OriginBlockNum)
		if err != nil {
			return
		}
		c.OldestL2()
	})
}

func Fuzz_Nosy_ChannelBuilder_OutputBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg ChannelConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var rollupCfg *rollup.Config
		fill_err = tp.Fill(&rollupCfg)
		if fill_err != nil {
			return
		}
		var latestL1OriginBlockNum uint64
		fill_err = tp.Fill(&latestL1OriginBlockNum)
		if fill_err != nil {
			return
		}
		if rollupCfg == nil {
			return
		}

		c, err := NewChannelBuilder(cfg, rollupCfg, latestL1OriginBlockNum)
		if err != nil {
			return
		}
		c.OutputBytes()
	})
}

func Fuzz_Nosy_ChannelBuilder_OutputFrames__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg ChannelConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var rollupCfg *rollup.Config
		fill_err = tp.Fill(&rollupCfg)
		if fill_err != nil {
			return
		}
		var latestL1OriginBlockNum uint64
		fill_err = tp.Fill(&latestL1OriginBlockNum)
		if fill_err != nil {
			return
		}
		if rollupCfg == nil {
			return
		}

		c, err := NewChannelBuilder(cfg, rollupCfg, latestL1OriginBlockNum)
		if err != nil {
			return
		}
		c.OutputFrames()
	})
}

func Fuzz_Nosy_ChannelBuilder_PendingFrames__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg ChannelConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var rollupCfg *rollup.Config
		fill_err = tp.Fill(&rollupCfg)
		if fill_err != nil {
			return
		}
		var latestL1OriginBlockNum uint64
		fill_err = tp.Fill(&latestL1OriginBlockNum)
		if fill_err != nil {
			return
		}
		if rollupCfg == nil {
			return
		}

		c, err := NewChannelBuilder(cfg, rollupCfg, latestL1OriginBlockNum)
		if err != nil {
			return
		}
		c.PendingFrames()
	})
}

func Fuzz_Nosy_ChannelBuilder_ReadyBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg ChannelConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var rollupCfg *rollup.Config
		fill_err = tp.Fill(&rollupCfg)
		if fill_err != nil {
			return
		}
		var latestL1OriginBlockNum uint64
		fill_err = tp.Fill(&latestL1OriginBlockNum)
		if fill_err != nil {
			return
		}
		if rollupCfg == nil {
			return
		}

		c, err := NewChannelBuilder(cfg, rollupCfg, latestL1OriginBlockNum)
		if err != nil {
			return
		}
		c.ReadyBytes()
	})
}

func Fuzz_Nosy_ChannelBuilder_RewindFrameCursor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg ChannelConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var rollupCfg *rollup.Config
		fill_err = tp.Fill(&rollupCfg)
		if fill_err != nil {
			return
		}
		var latestL1OriginBlockNum uint64
		fill_err = tp.Fill(&latestL1OriginBlockNum)
		if fill_err != nil {
			return
		}
		var frame frameData
		fill_err = tp.Fill(&frame)
		if fill_err != nil {
			return
		}
		if rollupCfg == nil {
			return
		}

		c, err := NewChannelBuilder(cfg, rollupCfg, latestL1OriginBlockNum)
		if err != nil {
			return
		}
		c.RewindFrameCursor(frame)
	})
}

func Fuzz_Nosy_ChannelBuilder_TimedOut__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg ChannelConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var rollupCfg *rollup.Config
		fill_err = tp.Fill(&rollupCfg)
		if fill_err != nil {
			return
		}
		var latestL1OriginBlockNum uint64
		fill_err = tp.Fill(&latestL1OriginBlockNum)
		if fill_err != nil {
			return
		}
		var blockNum uint64
		fill_err = tp.Fill(&blockNum)
		if fill_err != nil {
			return
		}
		if rollupCfg == nil {
			return
		}

		c, err := NewChannelBuilder(cfg, rollupCfg, latestL1OriginBlockNum)
		if err != nil {
			return
		}
		c.TimedOut(blockNum)
	})
}

func Fuzz_Nosy_ChannelBuilder_Timeout__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg ChannelConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var rollupCfg *rollup.Config
		fill_err = tp.Fill(&rollupCfg)
		if fill_err != nil {
			return
		}
		var latestL1OriginBlockNum uint64
		fill_err = tp.Fill(&latestL1OriginBlockNum)
		if fill_err != nil {
			return
		}
		if rollupCfg == nil {
			return
		}

		c, err := NewChannelBuilder(cfg, rollupCfg, latestL1OriginBlockNum)
		if err != nil {
			return
		}
		c.Timeout()
	})
}

func Fuzz_Nosy_ChannelBuilder_TotalFrames__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg ChannelConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var rollupCfg *rollup.Config
		fill_err = tp.Fill(&rollupCfg)
		if fill_err != nil {
			return
		}
		var latestL1OriginBlockNum uint64
		fill_err = tp.Fill(&latestL1OriginBlockNum)
		if fill_err != nil {
			return
		}
		if rollupCfg == nil {
			return
		}

		c, err := NewChannelBuilder(cfg, rollupCfg, latestL1OriginBlockNum)
		if err != nil {
			return
		}
		c.TotalFrames()
	})
}

func Fuzz_Nosy_ChannelBuilder_closeAndOutputAllFrames__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg ChannelConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var rollupCfg *rollup.Config
		fill_err = tp.Fill(&rollupCfg)
		if fill_err != nil {
			return
		}
		var latestL1OriginBlockNum uint64
		fill_err = tp.Fill(&latestL1OriginBlockNum)
		if fill_err != nil {
			return
		}
		if rollupCfg == nil {
			return
		}

		c, err := NewChannelBuilder(cfg, rollupCfg, latestL1OriginBlockNum)
		if err != nil {
			return
		}
		c.closeAndOutputAllFrames()
	})
}

func Fuzz_Nosy_ChannelBuilder_outputFrame__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg ChannelConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var rollupCfg *rollup.Config
		fill_err = tp.Fill(&rollupCfg)
		if fill_err != nil {
			return
		}
		var latestL1OriginBlockNum uint64
		fill_err = tp.Fill(&latestL1OriginBlockNum)
		if fill_err != nil {
			return
		}
		if rollupCfg == nil {
			return
		}

		c, err := NewChannelBuilder(cfg, rollupCfg, latestL1OriginBlockNum)
		if err != nil {
			return
		}
		c.outputFrame()
	})
}

func Fuzz_Nosy_ChannelBuilder_outputReadyFrames__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg ChannelConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var rollupCfg *rollup.Config
		fill_err = tp.Fill(&rollupCfg)
		if fill_err != nil {
			return
		}
		var latestL1OriginBlockNum uint64
		fill_err = tp.Fill(&latestL1OriginBlockNum)
		if fill_err != nil {
			return
		}
		if rollupCfg == nil {
			return
		}

		c, err := NewChannelBuilder(cfg, rollupCfg, latestL1OriginBlockNum)
		if err != nil {
			return
		}
		c.outputReadyFrames()
	})
}

// skipping Fuzz_Nosy_ChannelBuilder_setFullErr__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_ChannelBuilder_updateDurationTimeout__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg ChannelConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var rollupCfg *rollup.Config
		fill_err = tp.Fill(&rollupCfg)
		if fill_err != nil {
			return
		}
		var latestL1OriginBlockNum uint64
		fill_err = tp.Fill(&latestL1OriginBlockNum)
		if fill_err != nil {
			return
		}
		var l1BlockNum uint64
		fill_err = tp.Fill(&l1BlockNum)
		if fill_err != nil {
			return
		}
		if rollupCfg == nil {
			return
		}

		c, err := NewChannelBuilder(cfg, rollupCfg, latestL1OriginBlockNum)
		if err != nil {
			return
		}
		c.updateDurationTimeout(l1BlockNum)
	})
}

func Fuzz_Nosy_ChannelBuilder_updateSwTimeout__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg ChannelConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var rollupCfg *rollup.Config
		fill_err = tp.Fill(&rollupCfg)
		if fill_err != nil {
			return
		}
		var latestL1OriginBlockNum uint64
		fill_err = tp.Fill(&latestL1OriginBlockNum)
		if fill_err != nil {
			return
		}
		var l1InfoNumber uint64
		fill_err = tp.Fill(&l1InfoNumber)
		if fill_err != nil {
			return
		}
		if rollupCfg == nil {
			return
		}

		c, err := NewChannelBuilder(cfg, rollupCfg, latestL1OriginBlockNum)
		if err != nil {
			return
		}
		c.updateSwTimeout(l1InfoNumber)
	})
}

// skipping Fuzz_Nosy_ChannelBuilder_updateTimeout__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_ChannelConfig_Check__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cc *ChannelConfig
		fill_err = tp.Fill(&cc)
		if fill_err != nil {
			return
		}
		if cc == nil {
			return
		}

		cc.Check()
	})
}

func Fuzz_Nosy_ChannelConfig_InitCompressorConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cc *ChannelConfig
		fill_err = tp.Fill(&cc)
		if fill_err != nil {
			return
		}
		var approxComprRatio float64
		fill_err = tp.Fill(&approxComprRatio)
		if fill_err != nil {
			return
		}
		var comprKind string
		fill_err = tp.Fill(&comprKind)
		if fill_err != nil {
			return
		}
		var compressionAlgo derive.CompressionAlgo
		fill_err = tp.Fill(&compressionAlgo)
		if fill_err != nil {
			return
		}
		if cc == nil {
			return
		}

		cc.InitCompressorConfig(approxComprRatio, comprKind, compressionAlgo)
	})
}

func Fuzz_Nosy_ChannelConfig_InitNoneCompressor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cc *ChannelConfig
		fill_err = tp.Fill(&cc)
		if fill_err != nil {
			return
		}
		if cc == nil {
			return
		}

		cc.InitNoneCompressor()
	})
}

func Fuzz_Nosy_ChannelConfig_InitRatioCompressor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cc *ChannelConfig
		fill_err = tp.Fill(&cc)
		if fill_err != nil {
			return
		}
		var approxComprRatio float64
		fill_err = tp.Fill(&approxComprRatio)
		if fill_err != nil {
			return
		}
		var compressionAlgo derive.CompressionAlgo
		fill_err = tp.Fill(&compressionAlgo)
		if fill_err != nil {
			return
		}
		if cc == nil {
			return
		}

		cc.InitRatioCompressor(approxComprRatio, compressionAlgo)
	})
}

func Fuzz_Nosy_ChannelConfig_InitShadowCompressor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cc *ChannelConfig
		fill_err = tp.Fill(&cc)
		if fill_err != nil {
			return
		}
		var compressionAlgo derive.CompressionAlgo
		fill_err = tp.Fill(&compressionAlgo)
		if fill_err != nil {
			return
		}
		if cc == nil {
			return
		}

		cc.InitShadowCompressor(compressionAlgo)
	})
}

func Fuzz_Nosy_ChannelConfig_MaxFramesPerTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cc *ChannelConfig
		fill_err = tp.Fill(&cc)
		if fill_err != nil {
			return
		}
		if cc == nil {
			return
		}

		cc.MaxFramesPerTx()
	})
}

func Fuzz_Nosy_ChannelConfig_ReinitCompressorConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cc *ChannelConfig
		fill_err = tp.Fill(&cc)
		if fill_err != nil {
			return
		}
		if cc == nil {
			return
		}

		cc.ReinitCompressorConfig()
	})
}

func Fuzz_Nosy_ChannelFullError_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *ChannelFullError
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.Error()
	})
}

func Fuzz_Nosy_ChannelFullError_Unwrap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *ChannelFullError
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.Unwrap()
	})
}

func Fuzz_Nosy_DynamicEthChannelConfig_ChannelConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dec *DynamicEthChannelConfig
		fill_err = tp.Fill(&dec)
		if fill_err != nil {
			return
		}
		if dec == nil {
			return
		}

		dec.ChannelConfig()
	})
}

func Fuzz_Nosy_TestBatchSubmitter_JamTxPool__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *TestBatchSubmitter
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		l.JamTxPool(ctx)
	})
}

func Fuzz_Nosy_TestBatchSubmitter_WaitOnJammingTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *TestBatchSubmitter
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		l.WaitOnJammingTx(ctx)
	})
}

func Fuzz_Nosy_channel_AddBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *channel
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var block *types.Block
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		if c == nil || block == nil {
			return
		}

		c.AddBlock(block)
	})
}

func Fuzz_Nosy_channel_CheckTimeout__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *channel
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var l1BlockNum uint64
		fill_err = tp.Fill(&l1BlockNum)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.CheckTimeout(l1BlockNum)
	})
}

func Fuzz_Nosy_channel_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *channel
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.Close()
	})
}

func Fuzz_Nosy_channel_FullErr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *channel
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.FullErr()
	})
}

func Fuzz_Nosy_channel_HasTxData__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *channel
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.HasTxData()
	})
}

func Fuzz_Nosy_channel_ID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *channel
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.ID()
	})
}

func Fuzz_Nosy_channel_InputBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *channel
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.InputBytes()
	})
}

func Fuzz_Nosy_channel_IsFull__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *channel
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.IsFull()
	})
}

func Fuzz_Nosy_channel_LatestL1Origin__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *channel
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.LatestL1Origin()
	})
}

func Fuzz_Nosy_channel_LatestL2__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *channel
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.LatestL2()
	})
}

func Fuzz_Nosy_channel_NextTxData__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *channel
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.NextTxData()
	})
}

func Fuzz_Nosy_channel_NoneSubmitted__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *channel
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.NoneSubmitted()
	})
}

func Fuzz_Nosy_channel_OldestL1Origin__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *channel
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.OldestL1Origin()
	})
}

func Fuzz_Nosy_channel_OldestL2__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *channel
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.OldestL2()
	})
}

func Fuzz_Nosy_channel_OutputBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *channel
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.OutputBytes()
	})
}

func Fuzz_Nosy_channel_OutputFrames__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *channel
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.OutputFrames()
	})
}

func Fuzz_Nosy_channel_PendingFrames__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *channel
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.PendingFrames()
	})
}

func Fuzz_Nosy_channel_ReadyBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *channel
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.ReadyBytes()
	})
}

func Fuzz_Nosy_channel_Timeout__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *channel
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.Timeout()
	})
}

func Fuzz_Nosy_channel_TotalFrames__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *channel
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.TotalFrames()
	})
}

func Fuzz_Nosy_channel_TxConfirmed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *channel
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var id string
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var inclusionBlock eth.BlockID
		fill_err = tp.Fill(&inclusionBlock)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.TxConfirmed(id, inclusionBlock)
	})
}

func Fuzz_Nosy_channel_TxFailed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *channel
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var id string
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.TxFailed(id)
	})
}

func Fuzz_Nosy_channel_isFullySubmitted__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *channel
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.isFullySubmitted()
	})
}

func Fuzz_Nosy_channel_isTimedOut__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *channel
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.isTimedOut()
	})
}

func Fuzz_Nosy_channelManager_AddL2Block__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *channelManager
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var block *types.Block
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		if s == nil || block == nil {
			return
		}

		s.AddL2Block(block)
	})
}

func Fuzz_Nosy_channelManager_CheckExpectedProgress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *channelManager
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var syncStatus eth.SyncStatus
		fill_err = tp.Fill(&syncStatus)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.CheckExpectedProgress(syncStatus)
	})
}

func Fuzz_Nosy_channelManager_Clear__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *channelManager
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var l1OriginLastSubmittedChannel eth.BlockID
		fill_err = tp.Fill(&l1OriginLastSubmittedChannel)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Clear(l1OriginLastSubmittedChannel)
	})
}

func Fuzz_Nosy_channelManager_PendingDABytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *channelManager
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.PendingDABytes()
	})
}

// skipping Fuzz_Nosy_channelManager_SetChannelOutFactory__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-batcher/batcher.ChannelOutFactory

func Fuzz_Nosy_channelManager_TxConfirmed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *channelManager
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var _id txID
		fill_err = tp.Fill(&_id)
		if fill_err != nil {
			return
		}
		var inclusionBlock eth.BlockID
		fill_err = tp.Fill(&inclusionBlock)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.TxConfirmed(_id, inclusionBlock)
	})
}

func Fuzz_Nosy_channelManager_TxData__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *channelManager
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var l1Head eth.BlockID
		fill_err = tp.Fill(&l1Head)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.TxData(l1Head)
	})
}

func Fuzz_Nosy_channelManager_TxFailed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *channelManager
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var _id txID
		fill_err = tp.Fill(&_id)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.TxFailed(_id)
	})
}

func Fuzz_Nosy_channelManager_ensureChannelWithSpace__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *channelManager
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var l1Head eth.BlockID
		fill_err = tp.Fill(&l1Head)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.ensureChannelWithSpace(l1Head)
	})
}

func Fuzz_Nosy_channelManager_getReadyChannel__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *channelManager
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var l1Head eth.BlockID
		fill_err = tp.Fill(&l1Head)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.getReadyChannel(l1Head)
	})
}

func Fuzz_Nosy_channelManager_handleChannelInvalidated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *channelManager
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var c *channel
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if s == nil || c == nil {
			return
		}

		s.handleChannelInvalidated(c)
	})
}

func Fuzz_Nosy_channelManager_nextTxData__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *channelManager
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var channel *channel
		fill_err = tp.Fill(&channel)
		if fill_err != nil {
			return
		}
		if s == nil || channel == nil {
			return
		}

		s.nextTxData(channel)
	})
}

func Fuzz_Nosy_channelManager_outputFrames__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *channelManager
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.outputFrames()
	})
}

func Fuzz_Nosy_channelManager_pendingBlocks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *channelManager
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.pendingBlocks()
	})
}

func Fuzz_Nosy_channelManager_processBlocks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *channelManager
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.processBlocks()
	})
}

func Fuzz_Nosy_channelManager_pruneChannels__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *channelManager
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var newSafeHead eth.L2BlockRef
		fill_err = tp.Fill(&newSafeHead)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.pruneChannels(newSafeHead)
	})
}

func Fuzz_Nosy_channelManager_pruneSafeBlocks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *channelManager
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var newSafeHead eth.L2BlockRef
		fill_err = tp.Fill(&newSafeHead)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.pruneSafeBlocks(newSafeHead)
	})
}

func Fuzz_Nosy_channelManager_registerL1Block__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *channelManager
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var l1Head eth.BlockID
		fill_err = tp.Fill(&l1Head)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.registerL1Block(l1Head)
	})
}

func Fuzz_Nosy_channelManager_rewindToBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *channelManager
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var block eth.BlockID
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.rewindToBlock(block)
	})
}

func Fuzz_Nosy_txData_Blobs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var frame frameData
		fill_err = tp.Fill(&frame)
		if fill_err != nil {
			return
		}

		td := singleFrameTxData(frame)
		td.Blobs()
	})
}

func Fuzz_Nosy_txData_CallData__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var frame frameData
		fill_err = tp.Fill(&frame)
		if fill_err != nil {
			return
		}

		td := singleFrameTxData(frame)
		td.CallData()
	})
}

func Fuzz_Nosy_txData_Frames__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var frame frameData
		fill_err = tp.Fill(&frame)
		if fill_err != nil {
			return
		}

		td := singleFrameTxData(frame)
		td.Frames()
	})
}

func Fuzz_Nosy_txData_ID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var frame frameData
		fill_err = tp.Fill(&frame)
		if fill_err != nil {
			return
		}

		td := singleFrameTxData(frame)
		td.ID()
	})
}

func Fuzz_Nosy_txData_Len__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var frame frameData
		fill_err = tp.Fill(&frame)
		if fill_err != nil {
			return
		}

		td := singleFrameTxData(frame)
		td.Len()
	})
}

func Fuzz_Nosy_ChannelConfig_ChannelConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cc ChannelConfig
		fill_err = tp.Fill(&cc)
		if fill_err != nil {
			return
		}

		cc.ChannelConfig()
	})
}

// skipping Fuzz_Nosy_ChannelConfigProvider_ChannelConfig__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-batcher/batcher.ChannelConfigProvider

// skipping Fuzz_Nosy_GasPricer_SuggestGasPriceCaps__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-batcher/batcher.GasPricer

// skipping Fuzz_Nosy_L1Client_HeaderByNumber__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-batcher/batcher.L1Client

// skipping Fuzz_Nosy_L1Client_NonceAt__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-batcher/batcher.L1Client

// skipping Fuzz_Nosy_L2Client_BlockByNumber__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-batcher/batcher.L2Client

// skipping Fuzz_Nosy_RollupClient_SyncStatus__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-batcher/batcher.RollupClient

func Fuzz_Nosy_txID_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id txID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}

		id.String()
	})
}

func Fuzz_Nosy_txID_TerminalString__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id txID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}

		id.TerminalString()
	})
}

// skipping Fuzz_Nosy_txID_string__ because parameters include func, chan, or unsupported interface: func(id github.com/ethereum-optimism/optimism/op-node/rollup/derive.ChannelID) string

func Fuzz_Nosy_txRef_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r txRef
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}

		r.String()
	})
}

func Fuzz_Nosy_txRef_TerminalString__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r txRef
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}

		r.TerminalString()
	})
}

// skipping Fuzz_Nosy_txRef_string__ because parameters include func, chan, or unsupported interface: func(github.com/ethereum-optimism/optimism/op-batcher/batcher.txID) string

func Fuzz_Nosy_MaxDataSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, numFrames int, maxFrameSize uint64) {
		MaxDataSize(numFrames, maxFrameSize)
	})
}

// skipping Fuzz_Nosy_logFields__ because parameters include func, chan, or unsupported interface: []any
