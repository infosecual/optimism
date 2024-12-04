package sequencing

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

func Fuzz_Nosy_L1OriginSelector_CurrentAndNextOrigin__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var los *L1OriginSelector
		fill_err = tp.Fill(&los)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var l2Head eth.L2BlockRef
		fill_err = tp.Fill(&l2Head)
		if fill_err != nil {
			return
		}
		if los == nil {
			return
		}

		los.CurrentAndNextOrigin(ctx, l2Head)
	})
}

func Fuzz_Nosy_L1OriginSelector_FindL1Origin__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var los *L1OriginSelector
		fill_err = tp.Fill(&los)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var l2Head eth.L2BlockRef
		fill_err = tp.Fill(&l2Head)
		if fill_err != nil {
			return
		}
		if los == nil {
			return
		}

		los.FindL1Origin(ctx, l2Head)
	})
}

// skipping Fuzz_Nosy_L1OriginSelector_OnEvent__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/event.Event

func Fuzz_Nosy_L1OriginSelector_fetch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var los *L1OriginSelector
		fill_err = tp.Fill(&los)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var number uint64
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		if los == nil {
			return
		}

		los.fetch(ctx, number)
	})
}

func Fuzz_Nosy_L1OriginSelector_maybeSetNextOrigin__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var los *L1OriginSelector
		fill_err = tp.Fill(&los)
		if fill_err != nil {
			return
		}
		var nextOrigin eth.L1BlockRef
		fill_err = tp.Fill(&nextOrigin)
		if fill_err != nil {
			return
		}
		if los == nil {
			return
		}

		los.maybeSetNextOrigin(nextOrigin)
	})
}

func Fuzz_Nosy_L1OriginSelector_onForkchoiceUpdate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var los *L1OriginSelector
		fill_err = tp.Fill(&los)
		if fill_err != nil {
			return
		}
		var unsafeL2Head eth.L2BlockRef
		fill_err = tp.Fill(&unsafeL2Head)
		if fill_err != nil {
			return
		}
		if los == nil {
			return
		}

		los.onForkchoiceUpdate(unsafeL2Head)
	})
}

func Fuzz_Nosy_L1OriginSelector_reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var los *L1OriginSelector
		fill_err = tp.Fill(&los)
		if fill_err != nil {
			return
		}
		if los == nil {
			return
		}

		los.reset()
	})
}

func Fuzz_Nosy_L1OriginSelector_tryFetchNextOrigin__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var los *L1OriginSelector
		fill_err = tp.Fill(&los)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var currentOrigin eth.L1BlockRef
		fill_err = tp.Fill(&currentOrigin)
		if fill_err != nil {
			return
		}
		var nextOrigin eth.L1BlockRef
		fill_err = tp.Fill(&nextOrigin)
		if fill_err != nil {
			return
		}
		if los == nil {
			return
		}

		los.tryFetchNextOrigin(ctx, currentOrigin, nextOrigin)
	})
}

func Fuzz_Nosy_Sequencer_Active__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *Sequencer
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.Active()
	})
}

// skipping Fuzz_Nosy_Sequencer_AttachEmitter__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/event.Emitter

func Fuzz_Nosy_Sequencer_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *Sequencer
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.Close()
	})
}

func Fuzz_Nosy_Sequencer_ConductorEnabled__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *Sequencer
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

		d.ConductorEnabled(ctx)
	})
}

func Fuzz_Nosy_Sequencer_Init__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *Sequencer
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var active bool
		fill_err = tp.Fill(&active)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.Init(ctx, active)
	})
}

func Fuzz_Nosy_Sequencer_NextAction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *Sequencer
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.NextAction()
	})
}

// skipping Fuzz_Nosy_Sequencer_OnEvent__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/event.Event

func Fuzz_Nosy_Sequencer_OverrideLeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *Sequencer
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

		d.OverrideLeader(ctx)
	})
}

func Fuzz_Nosy_Sequencer_SetMaxSafeLag__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *Sequencer
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var v uint64
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.SetMaxSafeLag(ctx, v)
	})
}

func Fuzz_Nosy_Sequencer_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *Sequencer
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var head common.Hash
		fill_err = tp.Fill(&head)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.Start(ctx, head)
	})
}

func Fuzz_Nosy_Sequencer_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *Sequencer
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

		d.Stop(ctx)
	})
}

func Fuzz_Nosy_Sequencer_forceStart__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *Sequencer
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.forceStart()
	})
}

func Fuzz_Nosy_Sequencer_handleInvalid__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *Sequencer
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.handleInvalid()
	})
}

func Fuzz_Nosy_Sequencer_onBuildSealed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *Sequencer
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var x engine.BuildSealedEvent
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.onBuildSealed(x)
	})
}

func Fuzz_Nosy_Sequencer_onBuildStarted__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *Sequencer
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var x engine.BuildStartedEvent
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.onBuildStarted(x)
	})
}

func Fuzz_Nosy_Sequencer_onEngineResetConfirmedEvent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *Sequencer
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var _x2 engine.EngineResetConfirmedEvent
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.onEngineResetConfirmedEvent(_x2)
	})
}

func Fuzz_Nosy_Sequencer_onEngineTemporaryError__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *Sequencer
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var x rollup.EngineTemporaryErrorEvent
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.onEngineTemporaryError(x)
	})
}

func Fuzz_Nosy_Sequencer_onForkchoiceUpdate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *Sequencer
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var x engine.ForkchoiceUpdateEvent
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.onForkchoiceUpdate(x)
	})
}

func Fuzz_Nosy_Sequencer_onInvalidPayloadAttributes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *Sequencer
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var x engine.InvalidPayloadAttributesEvent
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.onInvalidPayloadAttributes(x)
	})
}

func Fuzz_Nosy_Sequencer_onPayloadInvalid__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *Sequencer
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var x engine.PayloadInvalidEvent
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.onPayloadInvalid(x)
	})
}

func Fuzz_Nosy_Sequencer_onPayloadSealExpiredError__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *Sequencer
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var x engine.PayloadSealExpiredErrorEvent
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.onPayloadSealExpiredError(x)
	})
}

func Fuzz_Nosy_Sequencer_onPayloadSealInvalid__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *Sequencer
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var x engine.PayloadSealInvalidEvent
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.onPayloadSealInvalid(x)
	})
}

func Fuzz_Nosy_Sequencer_onPayloadSuccess__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *Sequencer
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var x engine.PayloadSuccessEvent
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.onPayloadSuccess(x)
	})
}

func Fuzz_Nosy_Sequencer_onReset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *Sequencer
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var x rollup.ResetEvent
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.onReset(x)
	})
}

func Fuzz_Nosy_Sequencer_onSequencerAction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *Sequencer
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var _x2 SequencerActionEvent
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.onSequencerAction(_x2)
	})
}

func Fuzz_Nosy_Sequencer_setLatestHead__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *Sequencer
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var head eth.L2BlockRef
		fill_err = tp.Fill(&head)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.setLatestHead(head)
	})
}

func Fuzz_Nosy_Sequencer_startBuildingBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *Sequencer
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.startBuildingBlock()
	})
}

// skipping Fuzz_Nosy_AsyncGossiper_Clear__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/sequencing.AsyncGossiper

// skipping Fuzz_Nosy_AsyncGossiper_Get__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/sequencing.AsyncGossiper

// skipping Fuzz_Nosy_AsyncGossiper_Gossip__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/sequencing.AsyncGossiper

// skipping Fuzz_Nosy_AsyncGossiper_Start__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/sequencing.AsyncGossiper

// skipping Fuzz_Nosy_AsyncGossiper_Stop__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/sequencing.AsyncGossiper

func Fuzz_Nosy_DisabledSequencer_Active__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ds DisabledSequencer
		fill_err = tp.Fill(&ds)
		if fill_err != nil {
			return
		}

		ds.Active()
	})
}

func Fuzz_Nosy_DisabledSequencer_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ds DisabledSequencer
		fill_err = tp.Fill(&ds)
		if fill_err != nil {
			return
		}

		ds.Close()
	})
}

func Fuzz_Nosy_DisabledSequencer_ConductorEnabled__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ds DisabledSequencer
		fill_err = tp.Fill(&ds)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		ds.ConductorEnabled(ctx)
	})
}

func Fuzz_Nosy_DisabledSequencer_Init__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ds DisabledSequencer
		fill_err = tp.Fill(&ds)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var active bool
		fill_err = tp.Fill(&active)
		if fill_err != nil {
			return
		}

		ds.Init(ctx, active)
	})
}

func Fuzz_Nosy_DisabledSequencer_NextAction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ds DisabledSequencer
		fill_err = tp.Fill(&ds)
		if fill_err != nil {
			return
		}

		ds.NextAction()
	})
}

// skipping Fuzz_Nosy_DisabledSequencer_OnEvent__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/event.Event

func Fuzz_Nosy_DisabledSequencer_OverrideLeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ds DisabledSequencer
		fill_err = tp.Fill(&ds)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		ds.OverrideLeader(ctx)
	})
}

func Fuzz_Nosy_DisabledSequencer_SetMaxSafeLag__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ds DisabledSequencer
		fill_err = tp.Fill(&ds)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var v uint64
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}

		ds.SetMaxSafeLag(ctx, v)
	})
}

func Fuzz_Nosy_DisabledSequencer_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ds DisabledSequencer
		fill_err = tp.Fill(&ds)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var head common.Hash
		fill_err = tp.Fill(&head)
		if fill_err != nil {
			return
		}

		ds.Start(ctx, head)
	})
}

func Fuzz_Nosy_DisabledSequencer_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ds DisabledSequencer
		fill_err = tp.Fill(&ds)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		ds.Stop(ctx)
	})
}

// skipping Fuzz_Nosy_L1OriginSelectorIface_FindL1Origin__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/sequencing.L1OriginSelectorIface

// skipping Fuzz_Nosy_Metrics_RecordSequencerInconsistentL1Origin__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/sequencing.Metrics

// skipping Fuzz_Nosy_Metrics_RecordSequencerReset__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/sequencing.Metrics

// skipping Fuzz_Nosy_Metrics_RecordSequencingError__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/sequencing.Metrics

// skipping Fuzz_Nosy_Metrics_SetSequencerState__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/sequencing.Metrics

func Fuzz_Nosy_SequencerActionEvent_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ev SequencerActionEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}

		ev.String()
	})
}

// skipping Fuzz_Nosy_SequencerIface_Active__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/sequencing.SequencerIface

// skipping Fuzz_Nosy_SequencerIface_Close__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/sequencing.SequencerIface

// skipping Fuzz_Nosy_SequencerIface_ConductorEnabled__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/sequencing.SequencerIface

// skipping Fuzz_Nosy_SequencerIface_Init__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/sequencing.SequencerIface

// skipping Fuzz_Nosy_SequencerIface_NextAction__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/sequencing.SequencerIface

// skipping Fuzz_Nosy_SequencerIface_OverrideLeader__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/sequencing.SequencerIface

// skipping Fuzz_Nosy_SequencerIface_SetMaxSafeLag__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/sequencing.SequencerIface

// skipping Fuzz_Nosy_SequencerIface_Start__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/sequencing.SequencerIface

// skipping Fuzz_Nosy_SequencerIface_Stop__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/sequencing.SequencerIface

// skipping Fuzz_Nosy_SequencerStateListener_SequencerStarted__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/sequencing.SequencerStateListener

// skipping Fuzz_Nosy_SequencerStateListener_SequencerStopped__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/sequencing.SequencerStateListener
