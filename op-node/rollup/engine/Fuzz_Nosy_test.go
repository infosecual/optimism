package engine

import (
	"context"
	"testing"

	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum/go-ethereum/common/hexutil"
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

// skipping Fuzz_Nosy_EngDeriver_AttachEmitter__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/event.Emitter

// skipping Fuzz_Nosy_EngDeriver_OnEvent__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/event.Event

func Fuzz_Nosy_EngDeriver_emitDepositsOnlyPayloadAttributesRequest__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var eq *EngDeriver
		fill_err = tp.Fill(&eq)
		if fill_err != nil {
			return
		}
		var parent eth.BlockID
		fill_err = tp.Fill(&parent)
		if fill_err != nil {
			return
		}
		var derivedFrom eth.L1BlockRef
		fill_err = tp.Fill(&derivedFrom)
		if fill_err != nil {
			return
		}
		if eq == nil {
			return
		}

		eq.emitDepositsOnlyPayloadAttributesRequest(parent, derivedFrom)
	})
}

func Fuzz_Nosy_EngDeriver_onBuildCancel__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var eq *EngDeriver
		fill_err = tp.Fill(&eq)
		if fill_err != nil {
			return
		}
		var ev BuildCancelEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}
		if eq == nil {
			return
		}

		eq.onBuildCancel(ev)
	})
}

func Fuzz_Nosy_EngDeriver_onBuildInvalid__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var eq *EngDeriver
		fill_err = tp.Fill(&eq)
		if fill_err != nil {
			return
		}
		var ev BuildInvalidEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}
		if eq == nil {
			return
		}

		eq.onBuildInvalid(ev)
	})
}

func Fuzz_Nosy_EngDeriver_onBuildSeal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var eq *EngDeriver
		fill_err = tp.Fill(&eq)
		if fill_err != nil {
			return
		}
		var ev BuildSealEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}
		if eq == nil {
			return
		}

		eq.onBuildSeal(ev)
	})
}

func Fuzz_Nosy_EngDeriver_onBuildSealed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var eq *EngDeriver
		fill_err = tp.Fill(&eq)
		if fill_err != nil {
			return
		}
		var ev BuildSealedEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}
		if eq == nil {
			return
		}

		eq.onBuildSealed(ev)
	})
}

func Fuzz_Nosy_EngDeriver_onBuildStart__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var eq *EngDeriver
		fill_err = tp.Fill(&eq)
		if fill_err != nil {
			return
		}
		var ev BuildStartEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}
		if eq == nil {
			return
		}

		eq.onBuildStart(ev)
	})
}

func Fuzz_Nosy_EngDeriver_onBuildStarted__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var eq *EngDeriver
		fill_err = tp.Fill(&eq)
		if fill_err != nil {
			return
		}
		var ev BuildStartedEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}
		if eq == nil {
			return
		}

		eq.onBuildStarted(ev)
	})
}

func Fuzz_Nosy_EngDeriver_onPayloadInvalid__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var eq *EngDeriver
		fill_err = tp.Fill(&eq)
		if fill_err != nil {
			return
		}
		var ev PayloadInvalidEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}
		if eq == nil {
			return
		}

		eq.onPayloadInvalid(ev)
	})
}

func Fuzz_Nosy_EngDeriver_onPayloadProcess__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var eq *EngDeriver
		fill_err = tp.Fill(&eq)
		if fill_err != nil {
			return
		}
		var ev PayloadProcessEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}
		if eq == nil {
			return
		}

		eq.onPayloadProcess(ev)
	})
}

func Fuzz_Nosy_EngDeriver_onPayloadSuccess__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var eq *EngDeriver
		fill_err = tp.Fill(&eq)
		if fill_err != nil {
			return
		}
		var ev PayloadSuccessEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}
		if eq == nil {
			return
		}

		eq.onPayloadSuccess(ev)
	})
}

func Fuzz_Nosy_EngineController_BackupUnsafeL2Head__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *EngineController
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.BackupUnsafeL2Head()
	})
}

func Fuzz_Nosy_EngineController_CrossUnsafeL2Head__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *EngineController
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.CrossUnsafeL2Head()
	})
}

func Fuzz_Nosy_EngineController_Finalized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *EngineController
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.Finalized()
	})
}

func Fuzz_Nosy_EngineController_InsertUnsafePayload__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *EngineController
		fill_err = tp.Fill(&e)
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
		var ref eth.L2BlockRef
		fill_err = tp.Fill(&ref)
		if fill_err != nil {
			return
		}
		if e == nil || envelope == nil {
			return
		}

		e.InsertUnsafePayload(ctx, envelope, ref)
	})
}

func Fuzz_Nosy_EngineController_IsEngineSyncing__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *EngineController
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.IsEngineSyncing()
	})
}

func Fuzz_Nosy_EngineController_LocalSafeL2Head__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *EngineController
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.LocalSafeL2Head()
	})
}

func Fuzz_Nosy_EngineController_PendingSafeL2Head__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *EngineController
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.PendingSafeL2Head()
	})
}

func Fuzz_Nosy_EngineController_SafeL2Head__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *EngineController
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.SafeL2Head()
	})
}

func Fuzz_Nosy_EngineController_SetBackupUnsafeL2Head__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *EngineController
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var r eth.L2BlockRef
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var triggerReorg bool
		fill_err = tp.Fill(&triggerReorg)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.SetBackupUnsafeL2Head(r, triggerReorg)
	})
}

func Fuzz_Nosy_EngineController_SetCrossUnsafeHead__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *EngineController
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var r eth.L2BlockRef
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.SetCrossUnsafeHead(r)
	})
}

func Fuzz_Nosy_EngineController_SetFinalizedHead__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *EngineController
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var r eth.L2BlockRef
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.SetFinalizedHead(r)
	})
}

func Fuzz_Nosy_EngineController_SetLocalSafeHead__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *EngineController
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var r eth.L2BlockRef
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.SetLocalSafeHead(r)
	})
}

func Fuzz_Nosy_EngineController_SetPendingSafeL2Head__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *EngineController
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var r eth.L2BlockRef
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.SetPendingSafeL2Head(r)
	})
}

func Fuzz_Nosy_EngineController_SetSafeHead__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *EngineController
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var r eth.L2BlockRef
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.SetSafeHead(r)
	})
}

func Fuzz_Nosy_EngineController_SetUnsafeHead__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *EngineController
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var r eth.L2BlockRef
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.SetUnsafeHead(r)
	})
}

func Fuzz_Nosy_EngineController_TryBackupUnsafeReorg__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *EngineController
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.TryBackupUnsafeReorg(ctx)
	})
}

func Fuzz_Nosy_EngineController_TryUpdateEngine__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *EngineController
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.TryUpdateEngine(ctx)
	})
}

func Fuzz_Nosy_EngineController_UnsafeL2Head__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *EngineController
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.UnsafeL2Head()
	})
}

func Fuzz_Nosy_EngineController_checkForkchoiceUpdatedStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *EngineController
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var status eth.ExecutePayloadStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.checkForkchoiceUpdatedStatus(status)
	})
}

func Fuzz_Nosy_EngineController_checkNewPayloadStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *EngineController
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var status eth.ExecutePayloadStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.checkNewPayloadStatus(status)
	})
}

func Fuzz_Nosy_EngineController_logSyncProgressMaybe__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *EngineController
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.logSyncProgressMaybe()
	})
}

func Fuzz_Nosy_EngineController_shouldTryBackupUnsafeReorg__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *EngineController
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.shouldTryBackupUnsafeReorg()
	})
}

// skipping Fuzz_Nosy_EngineResetDeriver_AttachEmitter__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/event.Emitter

// skipping Fuzz_Nosy_EngineResetDeriver_OnEvent__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/event.Event

func Fuzz_Nosy_Kind_Clone__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var kind *Kind
		fill_err = tp.Fill(&kind)
		if fill_err != nil {
			return
		}
		if kind == nil {
			return
		}

		kind.Clone()
	})
}

func Fuzz_Nosy_Kind_Set__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var kind *Kind
		fill_err = tp.Fill(&kind)
		if fill_err != nil {
			return
		}
		var value string
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
		if kind == nil {
			return
		}

		kind.Set(value)
	})
}

func Fuzz_Nosy_BuildCancelEvent_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ev BuildCancelEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}

		ev.String()
	})
}

func Fuzz_Nosy_BuildInvalidEvent_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ev BuildInvalidEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}

		ev.String()
	})
}

func Fuzz_Nosy_BuildSealEvent_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ev BuildSealEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}

		ev.String()
	})
}

func Fuzz_Nosy_BuildSealedEvent_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ev BuildSealedEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}

		ev.String()
	})
}

func Fuzz_Nosy_BuildStartEvent_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ev BuildStartEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}

		ev.String()
	})
}

func Fuzz_Nosy_BuildStartedEvent_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ev BuildStartedEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}

		ev.String()
	})
}

func Fuzz_Nosy_CrossSafeUpdateEvent_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ev CrossSafeUpdateEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}

		ev.String()
	})
}

func Fuzz_Nosy_CrossUnsafeUpdateEvent_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ev CrossUnsafeUpdateEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}

		ev.String()
	})
}

func Fuzz_Nosy_CrossUpdateRequestEvent_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ev CrossUpdateRequestEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}

		ev.String()
	})
}

func Fuzz_Nosy_EngineResetConfirmedEvent_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ev EngineResetConfirmedEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}

		ev.String()
	})
}

// skipping Fuzz_Nosy_EngineState_Finalized__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/engine.EngineState

// skipping Fuzz_Nosy_EngineState_SafeL2Head__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/engine.EngineState

// skipping Fuzz_Nosy_EngineState_UnsafeL2Head__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/engine.EngineState

// skipping Fuzz_Nosy_ExecEngine_ForkchoiceUpdate__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/engine.ExecEngine

// skipping Fuzz_Nosy_ExecEngine_GetPayload__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/engine.ExecEngine

// skipping Fuzz_Nosy_ExecEngine_L2BlockRefByLabel__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/engine.ExecEngine

// skipping Fuzz_Nosy_ExecEngine_NewPayload__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/engine.ExecEngine

func Fuzz_Nosy_FinalizedUpdateEvent_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ev FinalizedUpdateEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}

		ev.String()
	})
}

func Fuzz_Nosy_ForceEngineResetEvent_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ev ForceEngineResetEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}

		ev.String()
	})
}

func Fuzz_Nosy_ForkchoiceRequestEvent_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ev ForkchoiceRequestEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}

		ev.String()
	})
}

func Fuzz_Nosy_ForkchoiceUpdateEvent_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ev ForkchoiceUpdateEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}

		ev.String()
	})
}

func Fuzz_Nosy_InteropPendingSafeChangedEvent_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ev InteropPendingSafeChangedEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}

		ev.String()
	})
}

func Fuzz_Nosy_InvalidPayloadAttributesEvent_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ev InvalidPayloadAttributesEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}

		ev.String()
	})
}

func Fuzz_Nosy_Kind_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var kind Kind
		fill_err = tp.Fill(&kind)
		if fill_err != nil {
			return
		}

		kind.String()
	})
}

func Fuzz_Nosy_Kind_SupportsPostFinalizationELSync__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var kind Kind
		fill_err = tp.Fill(&kind)
		if fill_err != nil {
			return
		}

		kind.SupportsPostFinalizationELSync()
	})
}

// skipping Fuzz_Nosy_LocalEngineState_BackupUnsafeL2Head__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/engine.LocalEngineState

// skipping Fuzz_Nosy_LocalEngineState_PendingSafeL2Head__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/engine.LocalEngineState

func Fuzz_Nosy_LocalSafeUpdateEvent_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ev LocalSafeUpdateEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}

		ev.String()
	})
}

// skipping Fuzz_Nosy_Metrics_CountSequencedTxsInBlock__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/engine.Metrics

// skipping Fuzz_Nosy_Metrics_RecordSequencerBuildingDiffTime__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/engine.Metrics

// skipping Fuzz_Nosy_Metrics_RecordSequencerSealingTime__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/engine.Metrics

func Fuzz_Nosy_PayloadInvalidEvent_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ev PayloadInvalidEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}

		ev.String()
	})
}

func Fuzz_Nosy_PayloadProcessEvent_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ev PayloadProcessEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}

		ev.String()
	})
}

func Fuzz_Nosy_PayloadSealExpiredErrorEvent_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ev PayloadSealExpiredErrorEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}

		ev.String()
	})
}

func Fuzz_Nosy_PayloadSealInvalidEvent_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ev PayloadSealInvalidEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}

		ev.String()
	})
}

func Fuzz_Nosy_PayloadSuccessEvent_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ev PayloadSuccessEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}

		ev.String()
	})
}

func Fuzz_Nosy_PendingSafeRequestEvent_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ev PendingSafeRequestEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}

		ev.String()
	})
}

func Fuzz_Nosy_PendingSafeUpdateEvent_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ev PendingSafeUpdateEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}

		ev.String()
	})
}

func Fuzz_Nosy_ProcessAttributesEvent_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ev ProcessAttributesEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}

		ev.String()
	})
}

func Fuzz_Nosy_ProcessUnsafePayloadEvent_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ev ProcessUnsafePayloadEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}

		ev.String()
	})
}

func Fuzz_Nosy_PromoteCrossUnsafeEvent_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ev PromoteCrossUnsafeEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}

		ev.String()
	})
}

func Fuzz_Nosy_PromoteFinalizedEvent_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ev PromoteFinalizedEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}

		ev.String()
	})
}

func Fuzz_Nosy_PromoteLocalSafeEvent_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ev PromoteLocalSafeEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}

		ev.String()
	})
}

func Fuzz_Nosy_PromotePendingSafeEvent_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ev PromotePendingSafeEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}

		ev.String()
	})
}

func Fuzz_Nosy_PromoteSafeEvent_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ev PromoteSafeEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}

		ev.String()
	})
}

func Fuzz_Nosy_PromoteUnsafeEvent_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ev PromoteUnsafeEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}

		ev.String()
	})
}

func Fuzz_Nosy_RequestCrossSafeEvent_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ev RequestCrossSafeEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}

		ev.String()
	})
}

func Fuzz_Nosy_RequestCrossUnsafeEvent_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ev RequestCrossUnsafeEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}

		ev.String()
	})
}

func Fuzz_Nosy_RequestFinalizedUpdateEvent_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ev RequestFinalizedUpdateEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}

		ev.String()
	})
}

// skipping Fuzz_Nosy_ResetEngineControl_SetBackupUnsafeL2Head__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/engine.ResetEngineControl

// skipping Fuzz_Nosy_ResetEngineControl_SetCrossUnsafeHead__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/engine.ResetEngineControl

// skipping Fuzz_Nosy_ResetEngineControl_SetFinalizedHead__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/engine.ResetEngineControl

// skipping Fuzz_Nosy_ResetEngineControl_SetLocalSafeHead__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/engine.ResetEngineControl

// skipping Fuzz_Nosy_ResetEngineControl_SetPendingSafeL2Head__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/engine.ResetEngineControl

// skipping Fuzz_Nosy_ResetEngineControl_SetSafeHead__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/engine.ResetEngineControl

// skipping Fuzz_Nosy_ResetEngineControl_SetUnsafeHead__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/engine.ResetEngineControl

func Fuzz_Nosy_ResetEngineRequestEvent_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ev ResetEngineRequestEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}

		ev.String()
	})
}

func Fuzz_Nosy_SafeDerivedEvent_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ev SafeDerivedEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}

		ev.String()
	})
}

func Fuzz_Nosy_TryBackupUnsafeReorgEvent_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ev TryBackupUnsafeReorgEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}

		ev.String()
	})
}

func Fuzz_Nosy_TryUpdateEngineEvent_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ev TryUpdateEngineEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}

		ev.String()
	})
}

func Fuzz_Nosy_UnsafeUpdateEvent_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ev UnsafeUpdateEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}

		ev.String()
	})
}

// skipping Fuzz_Nosy_ForceEngineReset__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/engine.ResetEngineControl

func Fuzz_Nosy_ValidEngineKind__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var value Kind
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}

		ValidEngineKind(value)
	})
}

func Fuzz_Nosy_isDepositTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var opaqueTx hexutil.Bytes
		fill_err = tp.Fill(&opaqueTx)
		if fill_err != nil {
			return
		}

		isDepositTx(opaqueTx)
	})
}

func Fuzz_Nosy_lastDeposit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var txns []hexutil.Bytes
		fill_err = tp.Fill(&txns)
		if fill_err != nil {
			return
		}

		lastDeposit(txns)
	})
}

// skipping Fuzz_Nosy_startPayload__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/engine.ExecEngine
