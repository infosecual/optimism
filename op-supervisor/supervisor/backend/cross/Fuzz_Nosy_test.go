package cross

import (
	"testing"

	"github.com/ethereum-optimism/optimism/op-service/eth"
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

func Fuzz_Nosy_Worker_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Worker
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

func Fuzz_Nosy_Worker_OnNewData__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Worker
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.OnNewData()
	})
}

func Fuzz_Nosy_Worker_ProcessWork__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Worker
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.ProcessWork()
	})
}

func Fuzz_Nosy_Worker_StartBackground__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Worker
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.StartBackground()
	})
}

func Fuzz_Nosy_Worker_worker__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Worker
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.worker()
	})
}

func Fuzz_Nosy_graph_addEdge__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *graph
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		var from node
		fill_err = tp.Fill(&from)
		if fill_err != nil {
			return
		}
		var to node
		fill_err = tp.Fill(&to)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.addEdge(from, to)
	})
}

// skipping Fuzz_Nosy_CrossSafeDeps_CandidateCrossSafe__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/cross.CrossSafeDeps

// skipping Fuzz_Nosy_CrossSafeDeps_CrossSafe__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/cross.CrossSafeDeps

// skipping Fuzz_Nosy_CrossSafeDeps_NextDerivedFrom__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/cross.CrossSafeDeps

// skipping Fuzz_Nosy_CrossSafeDeps_OpenBlock__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/cross.CrossSafeDeps

// skipping Fuzz_Nosy_CrossSafeDeps_PreviousDerived__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/cross.CrossSafeDeps

// skipping Fuzz_Nosy_CrossSafeDeps_UpdateCrossSafe__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/cross.CrossSafeDeps

// skipping Fuzz_Nosy_CrossUnsafeDeps_CrossUnsafe__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/cross.CrossUnsafeDeps

// skipping Fuzz_Nosy_CrossUnsafeDeps_OpenBlock__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/cross.CrossUnsafeDeps

// skipping Fuzz_Nosy_CrossUnsafeDeps_UpdateCrossUnsafe__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/cross.CrossUnsafeDeps

// skipping Fuzz_Nosy_CycleCheckDeps_OpenBlock__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/cross.CycleCheckDeps

// skipping Fuzz_Nosy_SafeFrontierCheckDeps_CandidateCrossSafe__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/cross.SafeFrontierCheckDeps

// skipping Fuzz_Nosy_SafeFrontierCheckDeps_CrossDerivedFrom__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/cross.SafeFrontierCheckDeps

// skipping Fuzz_Nosy_SafeFrontierCheckDeps_DependencySet__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/cross.SafeFrontierCheckDeps

// skipping Fuzz_Nosy_SafeStartDeps_Check__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/cross.SafeStartDeps

// skipping Fuzz_Nosy_SafeStartDeps_CrossDerivedFrom__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/cross.SafeStartDeps

// skipping Fuzz_Nosy_SafeStartDeps_DependencySet__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/cross.SafeStartDeps

// skipping Fuzz_Nosy_UnsafeFrontierCheckDeps_DependencySet__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/cross.UnsafeFrontierCheckDeps

// skipping Fuzz_Nosy_UnsafeFrontierCheckDeps_IsCrossUnsafe__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/cross.UnsafeFrontierCheckDeps

// skipping Fuzz_Nosy_UnsafeFrontierCheckDeps_IsLocalUnsafe__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/cross.UnsafeFrontierCheckDeps

// skipping Fuzz_Nosy_UnsafeFrontierCheckDeps_ParentBlock__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/cross.UnsafeFrontierCheckDeps

// skipping Fuzz_Nosy_UnsafeStartDeps_Check__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/cross.UnsafeStartDeps

// skipping Fuzz_Nosy_UnsafeStartDeps_DependencySet__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/cross.UnsafeStartDeps

// skipping Fuzz_Nosy_UnsafeStartDeps_IsCrossUnsafe__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/cross.UnsafeStartDeps

// skipping Fuzz_Nosy_CrossSafeHazards__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/cross.SafeStartDeps

// skipping Fuzz_Nosy_CrossUnsafeHazards__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/cross.UnsafeStartDeps

func Fuzz_Nosy_GenerateMermaidDiagram__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *graph
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		GenerateMermaidDiagram(g)
	})
}

func Fuzz_Nosy_blockSealMatchesRef__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var seal types.BlockSeal
		fill_err = tp.Fill(&seal)
		if fill_err != nil {
			return
		}
		var ref eth.L1BlockRef
		fill_err = tp.Fill(&ref)
		if fill_err != nil {
			return
		}

		blockSealMatchesRef(seal, ref)
	})
}

// skipping Fuzz_Nosy_gatherLogs__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/cross.CycleCheckDeps

func Fuzz_Nosy_sliceOfExecMsgs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var execMsgs map[uint32]*types.ExecutingMessage
		fill_err = tp.Fill(&execMsgs)
		if fill_err != nil {
			return
		}

		sliceOfExecMsgs(execMsgs)
	})
}
