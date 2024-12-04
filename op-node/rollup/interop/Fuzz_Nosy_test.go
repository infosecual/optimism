package interop

import (
	"testing"

	"github.com/ethereum-optimism/optimism/op-node/rollup/engine"
	"github.com/ethereum-optimism/optimism/op-node/rollup/finality"
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

// skipping Fuzz_Nosy_InteropDeriver_AttachEmitter__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/event.Emitter

// skipping Fuzz_Nosy_InteropDeriver_OnEvent__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/event.Event

func Fuzz_Nosy_InteropDeriver_onCrossSafeUpdateEvent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *InteropDeriver
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var x engine.CrossSafeUpdateEvent
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.onCrossSafeUpdateEvent(x)
	})
}

func Fuzz_Nosy_InteropDeriver_onCrossUnsafe__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *InteropDeriver
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var x engine.CrossUnsafeUpdateEvent
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.onCrossUnsafe(x)
	})
}

func Fuzz_Nosy_InteropDeriver_onFinalizedL1__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *InteropDeriver
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var x finality.FinalizeL1Event
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.onFinalizedL1(x)
	})
}

func Fuzz_Nosy_InteropDeriver_onFinalizedUpdate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *InteropDeriver
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var x engine.FinalizedUpdateEvent
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.onFinalizedUpdate(x)
	})
}

func Fuzz_Nosy_InteropDeriver_onInteropPendingSafeChangedEvent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *InteropDeriver
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var x engine.InteropPendingSafeChangedEvent
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.onInteropPendingSafeChangedEvent(x)
	})
}

func Fuzz_Nosy_InteropDeriver_onLocalUnsafeUpdate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *InteropDeriver
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var x engine.UnsafeUpdateEvent
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.onLocalUnsafeUpdate(x)
	})
}

// skipping Fuzz_Nosy_InteropBackend_CrossDerivedFrom__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/interop.InteropBackend

// skipping Fuzz_Nosy_InteropBackend_Finalized__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/interop.InteropBackend

// skipping Fuzz_Nosy_InteropBackend_SafeView__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/interop.InteropBackend

// skipping Fuzz_Nosy_InteropBackend_UnsafeView__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/interop.InteropBackend

// skipping Fuzz_Nosy_InteropBackend_UpdateFinalizedL1__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/interop.InteropBackend

// skipping Fuzz_Nosy_InteropBackend_UpdateLocalSafe__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/interop.InteropBackend

// skipping Fuzz_Nosy_InteropBackend_UpdateLocalUnsafe__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/interop.InteropBackend

// skipping Fuzz_Nosy_L2Source_L2BlockRefByHash__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/interop.L2Source

// skipping Fuzz_Nosy_L2Source_L2BlockRefByNumber__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/interop.L2Source
