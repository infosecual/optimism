package finality

import (
	"testing"

	"github.com/ethereum-optimism/optimism/op-node/rollup"
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

// skipping Fuzz_Nosy_AltDAFinalizer_OnEvent__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/event.Event

// skipping Fuzz_Nosy_Finalizer_AttachEmitter__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/event.Emitter

func Fuzz_Nosy_Finalizer_FinalizedL1__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fi *Finalizer
		fill_err = tp.Fill(&fi)
		if fill_err != nil {
			return
		}
		if fi == nil {
			return
		}

		fi.FinalizedL1()
	})
}

// skipping Fuzz_Nosy_Finalizer_OnEvent__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/event.Event

func Fuzz_Nosy_Finalizer_onDerivationIdle__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fi *Finalizer
		fill_err = tp.Fill(&fi)
		if fill_err != nil {
			return
		}
		var derivedFrom eth.L1BlockRef
		fill_err = tp.Fill(&derivedFrom)
		if fill_err != nil {
			return
		}
		if fi == nil {
			return
		}

		fi.onDerivationIdle(derivedFrom)
	})
}

func Fuzz_Nosy_Finalizer_onDerivedSafeBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fi *Finalizer
		fill_err = tp.Fill(&fi)
		if fill_err != nil {
			return
		}
		var l2Safe eth.L2BlockRef
		fill_err = tp.Fill(&l2Safe)
		if fill_err != nil {
			return
		}
		var derivedFrom eth.L1BlockRef
		fill_err = tp.Fill(&derivedFrom)
		if fill_err != nil {
			return
		}
		if fi == nil {
			return
		}

		fi.onDerivedSafeBlock(l2Safe, derivedFrom)
	})
}

func Fuzz_Nosy_Finalizer_onL1Finalized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fi *Finalizer
		fill_err = tp.Fill(&fi)
		if fill_err != nil {
			return
		}
		var l1Origin eth.L1BlockRef
		fill_err = tp.Fill(&l1Origin)
		if fill_err != nil {
			return
		}
		if fi == nil {
			return
		}

		fi.onL1Finalized(l1Origin)
	})
}

func Fuzz_Nosy_Finalizer_onReset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fi *Finalizer
		fill_err = tp.Fill(&fi)
		if fill_err != nil {
			return
		}
		if fi == nil {
			return
		}

		fi.onReset()
	})
}

func Fuzz_Nosy_Finalizer_tryFinalize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fi *Finalizer
		fill_err = tp.Fill(&fi)
		if fill_err != nil {
			return
		}
		if fi == nil {
			return
		}

		fi.tryFinalize()
	})
}

// skipping Fuzz_Nosy_AltDABackend_Finalize__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/finality.AltDABackend

// skipping Fuzz_Nosy_AltDABackend_OnFinalizedHeadSignal__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/finality.AltDABackend

func Fuzz_Nosy_FinalizeL1Event_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ev FinalizeL1Event
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}

		ev.String()
	})
}

// skipping Fuzz_Nosy_FinalizerEngine_Finalized__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/finality.FinalizerEngine

// skipping Fuzz_Nosy_FinalizerEngine_SetFinalizedHead__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/finality.FinalizerEngine

// skipping Fuzz_Nosy_FinalizerL1Interface_L1BlockRefByNumber__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/finality.FinalizerL1Interface

func Fuzz_Nosy_TryFinalizeEvent_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ev TryFinalizeEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}

		ev.String()
	})
}

func Fuzz_Nosy_calcFinalityLookback__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *rollup.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		calcFinalityLookback(cfg)
	})
}
