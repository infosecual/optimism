package interop

import (
	"testing"

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

func Fuzz_Nosy_InteropSetup_CreateActors__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var is *InteropSetup
		fill_err = tp.Fill(&is)
		if fill_err != nil {
			return
		}
		if is == nil {
			return
		}

		is.CreateActors()
	})
}

// skipping Fuzz_Nosy_SupervisorActor_SyncCrossSafe__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

// skipping Fuzz_Nosy_SupervisorActor_SyncCrossUnsafe__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing

// skipping Fuzz_Nosy_SupervisorActor_SyncEvents__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/actions/helpers.Testing
