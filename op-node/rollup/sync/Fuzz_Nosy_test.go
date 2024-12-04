package sync

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

func Fuzz_Nosy_Mode_Clone__(f *testing.F) {
	f.Fuzz(func(t *testing.T, s string) {
		m, err := StringToMode(s)
		if err != nil {
			return
		}
		m.Clone()
	})
}

func Fuzz_Nosy_Mode_Set__(f *testing.F) {
	f.Fuzz(func(t *testing.T, s string, value string) {
		m, err := StringToMode(s)
		if err != nil {
			return
		}
		m.Set(value)
	})
}

// skipping Fuzz_Nosy_L1Chain_L1BlockRefByHash__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/sync.L1Chain

// skipping Fuzz_Nosy_L1Chain_L1BlockRefByLabel__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/sync.L1Chain

// skipping Fuzz_Nosy_L1Chain_L1BlockRefByNumber__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/sync.L1Chain

// skipping Fuzz_Nosy_L2Chain_L2BlockRefByHash__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/sync.L2Chain

// skipping Fuzz_Nosy_L2Chain_L2BlockRefByLabel__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/sync.L2Chain

func Fuzz_Nosy_Mode_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, s string) {
		m, err := StringToMode(s)
		if err != nil {
			return
		}
		m.String()
	})
}
