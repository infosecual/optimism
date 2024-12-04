package driver

import (
	"context"
	"testing"

	"github.com/ethereum-optimism/optimism/op-node/rollup"
	"github.com/ethereum-optimism/optimism/op-node/rollup/engine"
	"github.com/ethereum-optimism/optimism/op-service/eth"
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

func Fuzz_Nosy_Driver_BlockRefWithStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Driver
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var num uint64
		fill_err = tp.Fill(&num)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.BlockRefWithStatus(ctx, num)
	})
}

func Fuzz_Nosy_Driver_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Driver
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Close()
	})
}

func Fuzz_Nosy_Driver_ConductorEnabled__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Driver
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

func Fuzz_Nosy_Driver_OnL1Finalized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Driver
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var finalized eth.L1BlockRef
		fill_err = tp.Fill(&finalized)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.OnL1Finalized(ctx, finalized)
	})
}

func Fuzz_Nosy_Driver_OnL1Head__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Driver
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var unsafe eth.L1BlockRef
		fill_err = tp.Fill(&unsafe)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.OnL1Head(ctx, unsafe)
	})
}

func Fuzz_Nosy_Driver_OnL1Safe__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Driver
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var safe eth.L1BlockRef
		fill_err = tp.Fill(&safe)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.OnL1Safe(ctx, safe)
	})
}

func Fuzz_Nosy_Driver_OnUnsafeL2Payload__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Driver
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var envelope *eth.ExecutionPayloadEnvelope
		fill_err = tp.Fill(&envelope)
		if fill_err != nil {
			return
		}
		if s == nil || envelope == nil {
			return
		}

		s.OnUnsafeL2Payload(ctx, envelope)
	})
}

func Fuzz_Nosy_Driver_OverrideLeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Driver
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

		s.OverrideLeader(ctx)
	})
}

func Fuzz_Nosy_Driver_ResetDerivationPipeline__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Driver
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

		s.ResetDerivationPipeline(ctx)
	})
}

func Fuzz_Nosy_Driver_SequencerActive__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Driver
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

func Fuzz_Nosy_Driver_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Driver
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Start()
	})
}

func Fuzz_Nosy_Driver_StartSequencer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Driver
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var blockHash common.Hash
		fill_err = tp.Fill(&blockHash)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.StartSequencer(ctx, blockHash)
	})
}

func Fuzz_Nosy_Driver_StopSequencer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Driver
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

func Fuzz_Nosy_Driver_SyncStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Driver
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

		s.SyncStatus(ctx)
	})
}

func Fuzz_Nosy_Driver_checkForGapInUnsafeQueue__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Driver
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

		s.checkForGapInUnsafeQueue(ctx)
	})
}

func Fuzz_Nosy_Driver_eventLoop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Driver
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.eventLoop()
	})
}

func Fuzz_Nosy_MeteredL1Fetcher_FetchReceipts__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MeteredL1Fetcher
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var blockHash common.Hash
		fill_err = tp.Fill(&blockHash)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.FetchReceipts(ctx, blockHash)
	})
}

func Fuzz_Nosy_MeteredL1Fetcher_InfoAndTxsByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MeteredL1Fetcher
		fill_err = tp.Fill(&m)
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
		if m == nil {
			return
		}

		m.InfoAndTxsByHash(ctx, hash)
	})
}

func Fuzz_Nosy_MeteredL1Fetcher_InfoByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MeteredL1Fetcher
		fill_err = tp.Fill(&m)
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
		if m == nil {
			return
		}

		m.InfoByHash(ctx, hash)
	})
}

func Fuzz_Nosy_MeteredL1Fetcher_L1BlockRefByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MeteredL1Fetcher
		fill_err = tp.Fill(&m)
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
		if m == nil {
			return
		}

		m.L1BlockRefByHash(ctx, hash)
	})
}

func Fuzz_Nosy_MeteredL1Fetcher_L1BlockRefByLabel__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MeteredL1Fetcher
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var label eth.BlockLabel
		fill_err = tp.Fill(&label)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.L1BlockRefByLabel(ctx, label)
	})
}

func Fuzz_Nosy_MeteredL1Fetcher_L1BlockRefByNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MeteredL1Fetcher
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var num uint64
		fill_err = tp.Fill(&num)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.L1BlockRefByNumber(ctx, num)
	})
}

func Fuzz_Nosy_MeteredL1Fetcher_recordTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MeteredL1Fetcher
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var method string
		fill_err = tp.Fill(&method)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.recordTime(method)
	})
}

// skipping Fuzz_Nosy_StepSchedulingDeriver_AttachEmitter__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/event.Emitter

func Fuzz_Nosy_StepSchedulingDeriver_NextDelayedStep__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *StepSchedulingDeriver
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.NextDelayedStep()
	})
}

func Fuzz_Nosy_StepSchedulingDeriver_NextStep__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *StepSchedulingDeriver
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.NextStep()
	})
}

// skipping Fuzz_Nosy_StepSchedulingDeriver_OnEvent__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/event.Event

// skipping Fuzz_Nosy_SyncDeriver_AttachEmitter__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/event.Emitter

// skipping Fuzz_Nosy_SyncDeriver_OnEvent__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/event.Event

func Fuzz_Nosy_SyncDeriver_SyncStep__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *SyncDeriver
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.SyncStep()
	})
}

func Fuzz_Nosy_SyncDeriver_onEngineConfirmedReset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *SyncDeriver
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var x engine.EngineResetConfirmedEvent
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.onEngineConfirmedReset(x)
	})
}

func Fuzz_Nosy_SyncDeriver_onResetEvent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *SyncDeriver
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var x rollup.ResetEvent
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.onResetEvent(x)
	})
}

func Fuzz_Nosy_SyncDeriver_onSafeDerivedBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *SyncDeriver
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var x engine.SafeDerivedEvent
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.onSafeDerivedBlock(x)
	})
}

// skipping Fuzz_Nosy_AltDAIface_Finalize__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/driver.AltDAIface

// skipping Fuzz_Nosy_AltDAIface_OnFinalizedHeadSignal__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/driver.AltDAIface

// skipping Fuzz_Nosy_AltSync_RequestL2Range__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/driver.AltSync

// skipping Fuzz_Nosy_AttributesHandler_HasAttributes__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/driver.AttributesHandler

// skipping Fuzz_Nosy_AttributesHandler_Proceed__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/driver.AttributesHandler

// skipping Fuzz_Nosy_AttributesHandler_SetAttributes__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/driver.AttributesHandler

// skipping Fuzz_Nosy_CLSync_LowestQueuedUnsafeBlock__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/driver.CLSync

// skipping Fuzz_Nosy_DerivationPipeline_ConfirmEngineReset__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/driver.DerivationPipeline

// skipping Fuzz_Nosy_DerivationPipeline_DerivationReady__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/driver.DerivationPipeline

// skipping Fuzz_Nosy_DerivationPipeline_Origin__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/driver.DerivationPipeline

// skipping Fuzz_Nosy_DerivationPipeline_Reset__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/driver.DerivationPipeline

// skipping Fuzz_Nosy_DerivationPipeline_Step__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/driver.DerivationPipeline

// skipping Fuzz_Nosy_Drain_Drain__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/driver.Drain

// skipping Fuzz_Nosy_EngineController_InsertUnsafePayload__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/driver.EngineController

// skipping Fuzz_Nosy_EngineController_IsEngineSyncing__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/driver.EngineController

// skipping Fuzz_Nosy_EngineController_TryBackupUnsafeReorg__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/driver.EngineController

// skipping Fuzz_Nosy_EngineController_TryUpdateEngine__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/driver.EngineController

// skipping Fuzz_Nosy_Finalizer_FinalizedL1__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/driver.Finalizer

// skipping Fuzz_Nosy_L1Chain_L1BlockRefByLabel__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/driver.L1Chain

// skipping Fuzz_Nosy_L1FetcherMetrics_RecordL1RequestTime__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/driver.L1FetcherMetrics

// skipping Fuzz_Nosy_L2Chain_L2BlockRefByHash__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/driver.L2Chain

// skipping Fuzz_Nosy_L2Chain_L2BlockRefByLabel__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/driver.L2Chain

// skipping Fuzz_Nosy_L2Chain_L2BlockRefByNumber__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/driver.L2Chain

// skipping Fuzz_Nosy_Metrics_RecordChannelInputBytes__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/driver.Metrics

// skipping Fuzz_Nosy_Metrics_RecordChannelTimedOut__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/driver.Metrics

// skipping Fuzz_Nosy_Metrics_RecordDerivationError__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/driver.Metrics

// skipping Fuzz_Nosy_Metrics_RecordDerivedBatches__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/driver.Metrics

// skipping Fuzz_Nosy_Metrics_RecordFrame__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/driver.Metrics

// skipping Fuzz_Nosy_Metrics_RecordHeadChannelOpened__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/driver.Metrics

// skipping Fuzz_Nosy_Metrics_RecordL1Ref__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/driver.Metrics

// skipping Fuzz_Nosy_Metrics_RecordL1ReorgDepth__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/driver.Metrics

// skipping Fuzz_Nosy_Metrics_RecordL2Ref__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/driver.Metrics

// skipping Fuzz_Nosy_Metrics_RecordPipelineReset__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/driver.Metrics

// skipping Fuzz_Nosy_Metrics_RecordPublishingError__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/driver.Metrics

// skipping Fuzz_Nosy_Metrics_RecordReceivedUnsafePayload__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/driver.Metrics

// skipping Fuzz_Nosy_Metrics_RecordUnsafePayloadsBuffer__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/driver.Metrics

// skipping Fuzz_Nosy_Metrics_SetDerivationIdle__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/driver.Metrics

// skipping Fuzz_Nosy_Metrics_SetSequencerState__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/driver.Metrics

// skipping Fuzz_Nosy_Network_PublishL2Payload__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/driver.Network

func Fuzz_Nosy_ResetStepBackoffEvent_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ev ResetStepBackoffEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}

		ev.String()
	})
}

// skipping Fuzz_Nosy_SequencerStateListener_SequencerStarted__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/driver.SequencerStateListener

// skipping Fuzz_Nosy_SequencerStateListener_SequencerStopped__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/driver.SequencerStateListener

func Fuzz_Nosy_StepAttemptEvent_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ev StepAttemptEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}

		ev.String()
	})
}

func Fuzz_Nosy_StepDelayedReqEvent_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ev StepDelayedReqEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}

		ev.String()
	})
}

func Fuzz_Nosy_StepEvent_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ev StepEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}

		ev.String()
	})
}

func Fuzz_Nosy_StepReqEvent_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ev StepReqEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}

		ev.String()
	})
}

// skipping Fuzz_Nosy_SyncStatusTracker_L1Head__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/driver.SyncStatusTracker

// skipping Fuzz_Nosy_SyncStatusTracker_SyncStatus__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/driver.SyncStatusTracker
