package attributes

import (
	"testing"

	"github.com/ethereum-optimism/optimism/op-node/rollup/derive"
	"github.com/ethereum-optimism/optimism/op-node/rollup/engine"
	"github.com/ethereum-optimism/optimism/op-service/eth"
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

// skipping Fuzz_Nosy_AttributesHandler_AttachEmitter__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/event.Emitter

// skipping Fuzz_Nosy_AttributesHandler_OnEvent__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/event.Event

func Fuzz_Nosy_AttributesHandler_consolidateNextSafeAttributes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var eq *AttributesHandler
		fill_err = tp.Fill(&eq)
		if fill_err != nil {
			return
		}
		var a2 *derive.AttributesWithParent
		fill_err = tp.Fill(&a2)
		if fill_err != nil {
			return
		}
		var onto eth.L2BlockRef
		fill_err = tp.Fill(&onto)
		if fill_err != nil {
			return
		}
		if eq == nil || a2 == nil {
			return
		}

		eq.consolidateNextSafeAttributes(a2, onto)
	})
}

func Fuzz_Nosy_AttributesHandler_onPendingSafeUpdate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var eq *AttributesHandler
		fill_err = tp.Fill(&eq)
		if fill_err != nil {
			return
		}
		var x engine.PendingSafeUpdateEvent
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if eq == nil {
			return
		}

		eq.onPendingSafeUpdate(x)
	})
}

// skipping Fuzz_Nosy_L2_PayloadByNumber__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/attributes.L2

// skipping Fuzz_Nosy_getMissingTxnHashes__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/log.Logger

// skipping Fuzz_Nosy_logL1InfoTxns__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/log.Logger
