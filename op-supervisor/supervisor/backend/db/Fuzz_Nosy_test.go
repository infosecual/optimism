package db

import (
	"testing"

	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum-optimism/optimism/op-supervisor/supervisor/types"
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

// skipping Fuzz_Nosy_ChainsDB_AddCrossDerivedFromDB__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db.CrossDerivedFromStorage

func Fuzz_Nosy_ChainsDB_AddCrossUnsafeTracker__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d1 *ChainsDB
		fill_err = tp.Fill(&d1)
		if fill_err != nil {
			return
		}
		var chainID types.ChainID
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		if d1 == nil {
			return
		}

		d1.AddCrossUnsafeTracker(chainID)
	})
}

// skipping Fuzz_Nosy_ChainsDB_AddLocalDerivedFromDB__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db.LocalDerivedFromStorage

func Fuzz_Nosy_ChainsDB_AddLog__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d1 *ChainsDB
		fill_err = tp.Fill(&d1)
		if fill_err != nil {
			return
		}
		var chain types.ChainID
		fill_err = tp.Fill(&chain)
		if fill_err != nil {
			return
		}
		var logHash common.Hash
		fill_err = tp.Fill(&logHash)
		if fill_err != nil {
			return
		}
		var parentBlock eth.BlockID
		fill_err = tp.Fill(&parentBlock)
		if fill_err != nil {
			return
		}
		var logIdx uint32
		fill_err = tp.Fill(&logIdx)
		if fill_err != nil {
			return
		}
		var execMsg *types.ExecutingMessage
		fill_err = tp.Fill(&execMsg)
		if fill_err != nil {
			return
		}
		if d1 == nil || execMsg == nil {
			return
		}

		d1.AddLog(chain, logHash, parentBlock, logIdx, execMsg)
	})
}

// skipping Fuzz_Nosy_ChainsDB_AddLogDB__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db.LogStorage

func Fuzz_Nosy_ChainsDB_CandidateCrossSafe__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d1 *ChainsDB
		fill_err = tp.Fill(&d1)
		if fill_err != nil {
			return
		}
		var chain types.ChainID
		fill_err = tp.Fill(&chain)
		if fill_err != nil {
			return
		}
		if d1 == nil {
			return
		}

		d1.CandidateCrossSafe(chain)
	})
}

func Fuzz_Nosy_ChainsDB_Check__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d1 *ChainsDB
		fill_err = tp.Fill(&d1)
		if fill_err != nil {
			return
		}
		var chain types.ChainID
		fill_err = tp.Fill(&chain)
		if fill_err != nil {
			return
		}
		var blockNum uint64
		fill_err = tp.Fill(&blockNum)
		if fill_err != nil {
			return
		}
		var logIdx uint32
		fill_err = tp.Fill(&logIdx)
		if fill_err != nil {
			return
		}
		var logHash common.Hash
		fill_err = tp.Fill(&logHash)
		if fill_err != nil {
			return
		}
		if d1 == nil {
			return
		}

		d1.Check(chain, blockNum, logIdx, logHash)
	})
}

func Fuzz_Nosy_ChainsDB_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d1 *ChainsDB
		fill_err = tp.Fill(&d1)
		if fill_err != nil {
			return
		}
		if d1 == nil {
			return
		}

		d1.Close()
	})
}

func Fuzz_Nosy_ChainsDB_CrossDerivedFrom__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d1 *ChainsDB
		fill_err = tp.Fill(&d1)
		if fill_err != nil {
			return
		}
		var chain types.ChainID
		fill_err = tp.Fill(&chain)
		if fill_err != nil {
			return
		}
		var derived eth.BlockID
		fill_err = tp.Fill(&derived)
		if fill_err != nil {
			return
		}
		if d1 == nil {
			return
		}

		d1.CrossDerivedFrom(chain, derived)
	})
}

func Fuzz_Nosy_ChainsDB_CrossDerivedFromBlockRef__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d1 *ChainsDB
		fill_err = tp.Fill(&d1)
		if fill_err != nil {
			return
		}
		var chainID types.ChainID
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var derived eth.BlockID
		fill_err = tp.Fill(&derived)
		if fill_err != nil {
			return
		}
		if d1 == nil {
			return
		}

		d1.CrossDerivedFromBlockRef(chainID, derived)
	})
}

func Fuzz_Nosy_ChainsDB_CrossSafe__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d1 *ChainsDB
		fill_err = tp.Fill(&d1)
		if fill_err != nil {
			return
		}
		var chainID types.ChainID
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		if d1 == nil {
			return
		}

		d1.CrossSafe(chainID)
	})
}

func Fuzz_Nosy_ChainsDB_CrossUnsafe__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d1 *ChainsDB
		fill_err = tp.Fill(&d1)
		if fill_err != nil {
			return
		}
		var chainID types.ChainID
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		if d1 == nil {
			return
		}

		d1.CrossUnsafe(chainID)
	})
}

func Fuzz_Nosy_ChainsDB_DependencySet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d1 *ChainsDB
		fill_err = tp.Fill(&d1)
		if fill_err != nil {
			return
		}
		if d1 == nil {
			return
		}

		d1.DependencySet()
	})
}

func Fuzz_Nosy_ChainsDB_Finalized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d1 *ChainsDB
		fill_err = tp.Fill(&d1)
		if fill_err != nil {
			return
		}
		var chainID types.ChainID
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		if d1 == nil {
			return
		}

		d1.Finalized(chainID)
	})
}

func Fuzz_Nosy_ChainsDB_FindSealedBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d1 *ChainsDB
		fill_err = tp.Fill(&d1)
		if fill_err != nil {
			return
		}
		var chain types.ChainID
		fill_err = tp.Fill(&chain)
		if fill_err != nil {
			return
		}
		var number uint64
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		if d1 == nil {
			return
		}

		d1.FindSealedBlock(chain, number)
	})
}

func Fuzz_Nosy_ChainsDB_IsCrossUnsafe__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d1 *ChainsDB
		fill_err = tp.Fill(&d1)
		if fill_err != nil {
			return
		}
		var chainID types.ChainID
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var block eth.BlockID
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		if d1 == nil {
			return
		}

		d1.IsCrossUnsafe(chainID, block)
	})
}

func Fuzz_Nosy_ChainsDB_IsLocalUnsafe__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d1 *ChainsDB
		fill_err = tp.Fill(&d1)
		if fill_err != nil {
			return
		}
		var chainID types.ChainID
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var block eth.BlockID
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		if d1 == nil {
			return
		}

		d1.IsLocalUnsafe(chainID, block)
	})
}

func Fuzz_Nosy_ChainsDB_IteratorStartingAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d1 *ChainsDB
		fill_err = tp.Fill(&d1)
		if fill_err != nil {
			return
		}
		var chain types.ChainID
		fill_err = tp.Fill(&chain)
		if fill_err != nil {
			return
		}
		var sealedNum uint64
		fill_err = tp.Fill(&sealedNum)
		if fill_err != nil {
			return
		}
		var logIndex uint32
		fill_err = tp.Fill(&logIndex)
		if fill_err != nil {
			return
		}
		if d1 == nil {
			return
		}

		d1.IteratorStartingAt(chain, sealedNum, logIndex)
	})
}

func Fuzz_Nosy_ChainsDB_LastDerivedFrom__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d1 *ChainsDB
		fill_err = tp.Fill(&d1)
		if fill_err != nil {
			return
		}
		var chainID types.ChainID
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var derivedFrom eth.BlockID
		fill_err = tp.Fill(&derivedFrom)
		if fill_err != nil {
			return
		}
		if d1 == nil {
			return
		}

		d1.LastDerivedFrom(chainID, derivedFrom)
	})
}

func Fuzz_Nosy_ChainsDB_LatestBlockNum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d1 *ChainsDB
		fill_err = tp.Fill(&d1)
		if fill_err != nil {
			return
		}
		var chain types.ChainID
		fill_err = tp.Fill(&chain)
		if fill_err != nil {
			return
		}
		if d1 == nil {
			return
		}

		d1.LatestBlockNum(chain)
	})
}

func Fuzz_Nosy_ChainsDB_LocalDerivedFrom__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d1 *ChainsDB
		fill_err = tp.Fill(&d1)
		if fill_err != nil {
			return
		}
		var chain types.ChainID
		fill_err = tp.Fill(&chain)
		if fill_err != nil {
			return
		}
		var derived eth.BlockID
		fill_err = tp.Fill(&derived)
		if fill_err != nil {
			return
		}
		if d1 == nil {
			return
		}

		d1.LocalDerivedFrom(chain, derived)
	})
}

func Fuzz_Nosy_ChainsDB_LocalSafe__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d1 *ChainsDB
		fill_err = tp.Fill(&d1)
		if fill_err != nil {
			return
		}
		var chainID types.ChainID
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		if d1 == nil {
			return
		}

		d1.LocalSafe(chainID)
	})
}

func Fuzz_Nosy_ChainsDB_LocalUnsafe__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d1 *ChainsDB
		fill_err = tp.Fill(&d1)
		if fill_err != nil {
			return
		}
		var chainID types.ChainID
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		if d1 == nil {
			return
		}

		d1.LocalUnsafe(chainID)
	})
}

func Fuzz_Nosy_ChainsDB_NextDerivedFrom__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d1 *ChainsDB
		fill_err = tp.Fill(&d1)
		if fill_err != nil {
			return
		}
		var chain types.ChainID
		fill_err = tp.Fill(&chain)
		if fill_err != nil {
			return
		}
		var derivedFrom eth.BlockID
		fill_err = tp.Fill(&derivedFrom)
		if fill_err != nil {
			return
		}
		if d1 == nil {
			return
		}

		d1.NextDerivedFrom(chain, derivedFrom)
	})
}

func Fuzz_Nosy_ChainsDB_OpenBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d1 *ChainsDB
		fill_err = tp.Fill(&d1)
		if fill_err != nil {
			return
		}
		var chainID types.ChainID
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var blockNum uint64
		fill_err = tp.Fill(&blockNum)
		if fill_err != nil {
			return
		}
		if d1 == nil {
			return
		}

		d1.OpenBlock(chainID, blockNum)
	})
}

func Fuzz_Nosy_ChainsDB_ParentBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d1 *ChainsDB
		fill_err = tp.Fill(&d1)
		if fill_err != nil {
			return
		}
		var chainID types.ChainID
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var parentOf eth.BlockID
		fill_err = tp.Fill(&parentOf)
		if fill_err != nil {
			return
		}
		if d1 == nil {
			return
		}

		d1.ParentBlock(chainID, parentOf)
	})
}

func Fuzz_Nosy_ChainsDB_PreviousDerived__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d1 *ChainsDB
		fill_err = tp.Fill(&d1)
		if fill_err != nil {
			return
		}
		var chain types.ChainID
		fill_err = tp.Fill(&chain)
		if fill_err != nil {
			return
		}
		var derived eth.BlockID
		fill_err = tp.Fill(&derived)
		if fill_err != nil {
			return
		}
		if d1 == nil {
			return
		}

		d1.PreviousDerived(chain, derived)
	})
}

func Fuzz_Nosy_ChainsDB_PreviousDerivedFrom__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d1 *ChainsDB
		fill_err = tp.Fill(&d1)
		if fill_err != nil {
			return
		}
		var chain types.ChainID
		fill_err = tp.Fill(&chain)
		if fill_err != nil {
			return
		}
		var derivedFrom eth.BlockID
		fill_err = tp.Fill(&derivedFrom)
		if fill_err != nil {
			return
		}
		if d1 == nil {
			return
		}

		d1.PreviousDerivedFrom(chain, derivedFrom)
	})
}

func Fuzz_Nosy_ChainsDB_ResumeFromLastSealedBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d1 *ChainsDB
		fill_err = tp.Fill(&d1)
		if fill_err != nil {
			return
		}
		if d1 == nil {
			return
		}

		d1.ResumeFromLastSealedBlock()
	})
}

func Fuzz_Nosy_ChainsDB_Rewind__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d1 *ChainsDB
		fill_err = tp.Fill(&d1)
		if fill_err != nil {
			return
		}
		var chain types.ChainID
		fill_err = tp.Fill(&chain)
		if fill_err != nil {
			return
		}
		var headBlockNum uint64
		fill_err = tp.Fill(&headBlockNum)
		if fill_err != nil {
			return
		}
		if d1 == nil {
			return
		}

		d1.Rewind(chain, headBlockNum)
	})
}

func Fuzz_Nosy_ChainsDB_Safest__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d1 *ChainsDB
		fill_err = tp.Fill(&d1)
		if fill_err != nil {
			return
		}
		var chainID types.ChainID
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var blockNum uint64
		fill_err = tp.Fill(&blockNum)
		if fill_err != nil {
			return
		}
		var index uint32
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		if d1 == nil {
			return
		}

		d1.Safest(chainID, blockNum, index)
	})
}

func Fuzz_Nosy_ChainsDB_SealBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d1 *ChainsDB
		fill_err = tp.Fill(&d1)
		if fill_err != nil {
			return
		}
		var chain types.ChainID
		fill_err = tp.Fill(&chain)
		if fill_err != nil {
			return
		}
		var block eth.L1BlockRef
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		if d1 == nil {
			return
		}

		d1.SealBlock(chain, block)
	})
}

func Fuzz_Nosy_ChainsDB_UpdateCrossSafe__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d1 *ChainsDB
		fill_err = tp.Fill(&d1)
		if fill_err != nil {
			return
		}
		var chain types.ChainID
		fill_err = tp.Fill(&chain)
		if fill_err != nil {
			return
		}
		var l1View eth.L1BlockRef
		fill_err = tp.Fill(&l1View)
		if fill_err != nil {
			return
		}
		var lastCrossDerived eth.L1BlockRef
		fill_err = tp.Fill(&lastCrossDerived)
		if fill_err != nil {
			return
		}
		if d1 == nil {
			return
		}

		d1.UpdateCrossSafe(chain, l1View, lastCrossDerived)
	})
}

func Fuzz_Nosy_ChainsDB_UpdateCrossUnsafe__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d1 *ChainsDB
		fill_err = tp.Fill(&d1)
		if fill_err != nil {
			return
		}
		var chain types.ChainID
		fill_err = tp.Fill(&chain)
		if fill_err != nil {
			return
		}
		var crossUnsafe types.BlockSeal
		fill_err = tp.Fill(&crossUnsafe)
		if fill_err != nil {
			return
		}
		if d1 == nil {
			return
		}

		d1.UpdateCrossUnsafe(chain, crossUnsafe)
	})
}

func Fuzz_Nosy_ChainsDB_UpdateFinalizedL1__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d1 *ChainsDB
		fill_err = tp.Fill(&d1)
		if fill_err != nil {
			return
		}
		var finalized eth.L1BlockRef
		fill_err = tp.Fill(&finalized)
		if fill_err != nil {
			return
		}
		if d1 == nil {
			return
		}

		d1.UpdateFinalizedL1(finalized)
	})
}

func Fuzz_Nosy_ChainsDB_UpdateLocalSafe__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d1 *ChainsDB
		fill_err = tp.Fill(&d1)
		if fill_err != nil {
			return
		}
		var chain types.ChainID
		fill_err = tp.Fill(&chain)
		if fill_err != nil {
			return
		}
		var derivedFrom eth.L1BlockRef
		fill_err = tp.Fill(&derivedFrom)
		if fill_err != nil {
			return
		}
		var lastDerived eth.L1BlockRef
		fill_err = tp.Fill(&lastDerived)
		if fill_err != nil {
			return
		}
		if d1 == nil {
			return
		}

		d1.UpdateLocalSafe(chain, derivedFrom, lastDerived)
	})
}

// skipping Fuzz_Nosy_LocalDerivedFromStorage_AddDerived__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db.LocalDerivedFromStorage

// skipping Fuzz_Nosy_LocalDerivedFromStorage_DerivedFrom__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db.LocalDerivedFromStorage

// skipping Fuzz_Nosy_LocalDerivedFromStorage_First__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db.LocalDerivedFromStorage

// skipping Fuzz_Nosy_LocalDerivedFromStorage_FirstAfter__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db.LocalDerivedFromStorage

// skipping Fuzz_Nosy_LocalDerivedFromStorage_LastDerivedAt__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db.LocalDerivedFromStorage

// skipping Fuzz_Nosy_LocalDerivedFromStorage_Latest__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db.LocalDerivedFromStorage

// skipping Fuzz_Nosy_LocalDerivedFromStorage_NextDerived__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db.LocalDerivedFromStorage

// skipping Fuzz_Nosy_LocalDerivedFromStorage_NextDerivedFrom__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db.LocalDerivedFromStorage

// skipping Fuzz_Nosy_LocalDerivedFromStorage_PreviousDerived__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db.LocalDerivedFromStorage

// skipping Fuzz_Nosy_LocalDerivedFromStorage_PreviousDerivedFrom__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db.LocalDerivedFromStorage

// skipping Fuzz_Nosy_LogStorage_AddLog__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db.LogStorage

// skipping Fuzz_Nosy_LogStorage_Contains__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db.LogStorage

// skipping Fuzz_Nosy_LogStorage_FindSealedBlock__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db.LogStorage

// skipping Fuzz_Nosy_LogStorage_IteratorStartingAt__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db.LogStorage

// skipping Fuzz_Nosy_LogStorage_LatestSealedBlockNum__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db.LogStorage

// skipping Fuzz_Nosy_LogStorage_OpenBlock__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db.LogStorage

// skipping Fuzz_Nosy_LogStorage_Rewind__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db.LogStorage

// skipping Fuzz_Nosy_LogStorage_SealBlock__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db.LogStorage

func Fuzz_Nosy_prepChainDir__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainID types.ChainID
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var datadir string
		fill_err = tp.Fill(&datadir)
		if fill_err != nil {
			return
		}

		prepChainDir(chainID, datadir)
	})
}

func Fuzz_Nosy_prepCrossDerivedFromDBPath__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainID types.ChainID
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var datadir string
		fill_err = tp.Fill(&datadir)
		if fill_err != nil {
			return
		}

		prepCrossDerivedFromDBPath(chainID, datadir)
	})
}

func Fuzz_Nosy_prepLocalDerivedFromDBPath__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainID types.ChainID
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var datadir string
		fill_err = tp.Fill(&datadir)
		if fill_err != nil {
			return
		}

		prepLocalDerivedFromDBPath(chainID, datadir)
	})
}

func Fuzz_Nosy_prepLogDBPath__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainID types.ChainID
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var datadir string
		fill_err = tp.Fill(&datadir)
		if fill_err != nil {
			return
		}

		prepLogDBPath(chainID, datadir)
	})
}
