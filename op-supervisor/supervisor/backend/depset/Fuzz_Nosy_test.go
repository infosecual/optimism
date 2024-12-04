package depset

import (
	"context"
	"testing"

	"github.com/ethereum-optimism/optimism/op-supervisor/supervisor/types"
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

func Fuzz_Nosy_JsonDependencySetLoader_LoadDependencySet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var j *JsonDependencySetLoader
		fill_err = tp.Fill(&j)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if j == nil {
			return
		}

		j.LoadDependencySet(ctx)
	})
}

func Fuzz_Nosy_StaticConfigDependencySet_CanExecuteAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dependencies map[types.ChainID]*StaticConfigDependency
		fill_err = tp.Fill(&dependencies)
		if fill_err != nil {
			return
		}
		var chainID types.ChainID
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var execTimestamp uint64
		fill_err = tp.Fill(&execTimestamp)
		if fill_err != nil {
			return
		}

		ds, err := NewStaticConfigDependencySet(dependencies)
		if err != nil {
			return
		}
		ds.CanExecuteAt(chainID, execTimestamp)
	})
}

func Fuzz_Nosy_StaticConfigDependencySet_CanInitiateAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dependencies map[types.ChainID]*StaticConfigDependency
		fill_err = tp.Fill(&dependencies)
		if fill_err != nil {
			return
		}
		var chainID types.ChainID
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var initTimestamp uint64
		fill_err = tp.Fill(&initTimestamp)
		if fill_err != nil {
			return
		}

		ds, err := NewStaticConfigDependencySet(dependencies)
		if err != nil {
			return
		}
		ds.CanInitiateAt(chainID, initTimestamp)
	})
}

func Fuzz_Nosy_StaticConfigDependencySet_ChainIDFromIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dependencies map[types.ChainID]*StaticConfigDependency
		fill_err = tp.Fill(&dependencies)
		if fill_err != nil {
			return
		}
		var index types.ChainIndex
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}

		ds, err := NewStaticConfigDependencySet(dependencies)
		if err != nil {
			return
		}
		ds.ChainIDFromIndex(index)
	})
}

func Fuzz_Nosy_StaticConfigDependencySet_ChainIndexFromID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dependencies map[types.ChainID]*StaticConfigDependency
		fill_err = tp.Fill(&dependencies)
		if fill_err != nil {
			return
		}
		var id types.ChainID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}

		ds, err := NewStaticConfigDependencySet(dependencies)
		if err != nil {
			return
		}
		ds.ChainIndexFromID(id)
	})
}

func Fuzz_Nosy_StaticConfigDependencySet_Chains__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dependencies map[types.ChainID]*StaticConfigDependency
		fill_err = tp.Fill(&dependencies)
		if fill_err != nil {
			return
		}

		ds, err := NewStaticConfigDependencySet(dependencies)
		if err != nil {
			return
		}
		ds.Chains()
	})
}

func Fuzz_Nosy_StaticConfigDependencySet_HasChain__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dependencies map[types.ChainID]*StaticConfigDependency
		fill_err = tp.Fill(&dependencies)
		if fill_err != nil {
			return
		}
		var chainID types.ChainID
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}

		ds, err := NewStaticConfigDependencySet(dependencies)
		if err != nil {
			return
		}
		ds.HasChain(chainID)
	})
}

func Fuzz_Nosy_StaticConfigDependencySet_LoadDependencySet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dependencies map[types.ChainID]*StaticConfigDependency
		fill_err = tp.Fill(&dependencies)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		ds, err := NewStaticConfigDependencySet(dependencies)
		if err != nil {
			return
		}
		ds.LoadDependencySet(ctx)
	})
}

func Fuzz_Nosy_StaticConfigDependencySet_MarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dependencies map[types.ChainID]*StaticConfigDependency
		fill_err = tp.Fill(&dependencies)
		if fill_err != nil {
			return
		}

		ds, err := NewStaticConfigDependencySet(dependencies)
		if err != nil {
			return
		}
		ds.MarshalJSON()
	})
}

func Fuzz_Nosy_StaticConfigDependencySet_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dependencies map[types.ChainID]*StaticConfigDependency
		fill_err = tp.Fill(&dependencies)
		if fill_err != nil {
			return
		}
		var d2 []byte
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}

		ds, err := NewStaticConfigDependencySet(dependencies)
		if err != nil {
			return
		}
		ds.UnmarshalJSON(d2)
	})
}

func Fuzz_Nosy_StaticConfigDependencySet_hydrate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dependencies map[types.ChainID]*StaticConfigDependency
		fill_err = tp.Fill(&dependencies)
		if fill_err != nil {
			return
		}

		ds, err := NewStaticConfigDependencySet(dependencies)
		if err != nil {
			return
		}
		ds.hydrate()
	})
}

// skipping Fuzz_Nosy_DependencySet_CanExecuteAt__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/depset.DependencySet

// skipping Fuzz_Nosy_DependencySet_CanInitiateAt__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/depset.DependencySet

// skipping Fuzz_Nosy_DependencySet_ChainIDFromIndex__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/depset.DependencySet

// skipping Fuzz_Nosy_DependencySet_ChainIndexFromID__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/depset.DependencySet

// skipping Fuzz_Nosy_DependencySet_Chains__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/depset.DependencySet

// skipping Fuzz_Nosy_DependencySet_HasChain__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/depset.DependencySet

// skipping Fuzz_Nosy_DependencySetSource_LoadDependencySet__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/depset.DependencySetSource
