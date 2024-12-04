package e2eutils

import (
	"context"
	"testing"

	"github.com/ethereum-optimism/optimism/op-chain-ops/genesis"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum/go-ethereum/core/types"
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

func Fuzz_Nosy_Addresses_All__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a *Addresses
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if a == nil {
			return
		}

		a.All()
	})
}

func Fuzz_Nosy_BlobsStore_GetAllSidecars__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var l1Timestamp uint64
		fill_err = tp.Fill(&l1Timestamp)
		if fill_err != nil {
			return
		}

		store := NewBlobStore()
		store.GetAllSidecars(ctx, l1Timestamp)
	})
}

func Fuzz_Nosy_BlobsStore_GetBlobSidecars__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var ref eth.L1BlockRef
		fill_err = tp.Fill(&ref)
		if fill_err != nil {
			return
		}
		var hashes []eth.IndexedBlobHash
		fill_err = tp.Fill(&hashes)
		if fill_err != nil {
			return
		}

		store := NewBlobStore()
		store.GetBlobSidecars(ctx, ref, hashes)
	})
}

func Fuzz_Nosy_BlobsStore_GetBlobs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var ref eth.L1BlockRef
		fill_err = tp.Fill(&ref)
		if fill_err != nil {
			return
		}
		var hashes []eth.IndexedBlobHash
		fill_err = tp.Fill(&hashes)
		if fill_err != nil {
			return
		}

		store := NewBlobStore()
		store.GetBlobs(ctx, ref, hashes)
	})
}

func Fuzz_Nosy_BlobsStore_StoreBlob__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var blockTime uint64
		fill_err = tp.Fill(&blockTime)
		if fill_err != nil {
			return
		}
		var indexedHash eth.IndexedBlobHash
		fill_err = tp.Fill(&indexedHash)
		if fill_err != nil {
			return
		}
		var blob *eth.Blob
		fill_err = tp.Fill(&blob)
		if fill_err != nil {
			return
		}
		if blob == nil {
			return
		}

		store := NewBlobStore()
		store.StoreBlob(blockTime, indexedHash, blob)
	})
}

func Fuzz_Nosy_MnemonicConfig_Secrets__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MnemonicConfig
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Secrets()
	})
}

func Fuzz_Nosy_Secrets_AccountAtIdx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Secrets
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var idx int
		fill_err = tp.Fill(&idx)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.AccountAtIdx(idx)
	})
}

func Fuzz_Nosy_Secrets_Addresses__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Secrets
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Addresses()
	})
}

// skipping Fuzz_Nosy_TestingBase_Cleanup__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/e2eutils.TestingBase

// skipping Fuzz_Nosy_TestingBase_Error__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/e2eutils.TestingBase

// skipping Fuzz_Nosy_TestingBase_Errorf__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/e2eutils.TestingBase

// skipping Fuzz_Nosy_TestingBase_Fail__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/e2eutils.TestingBase

// skipping Fuzz_Nosy_TestingBase_FailNow__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/e2eutils.TestingBase

// skipping Fuzz_Nosy_TestingBase_Failed__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/e2eutils.TestingBase

// skipping Fuzz_Nosy_TestingBase_Fatal__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/e2eutils.TestingBase

// skipping Fuzz_Nosy_TestingBase_Fatalf__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/e2eutils.TestingBase

// skipping Fuzz_Nosy_TestingBase_Helper__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/e2eutils.TestingBase

// skipping Fuzz_Nosy_TestingBase_Log__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/e2eutils.TestingBase

// skipping Fuzz_Nosy_TestingBase_Logf__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/e2eutils.TestingBase

// skipping Fuzz_Nosy_TestingBase_Name__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/e2eutils.TestingBase

// skipping Fuzz_Nosy_TestingBase_Parallel__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/e2eutils.TestingBase

// skipping Fuzz_Nosy_TestingBase_Setenv__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/e2eutils.TestingBase

// skipping Fuzz_Nosy_TestingBase_Skip__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/e2eutils.TestingBase

// skipping Fuzz_Nosy_TestingBase_SkipNow__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/e2eutils.TestingBase

// skipping Fuzz_Nosy_TestingBase_Skipf__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/e2eutils.TestingBase

// skipping Fuzz_Nosy_TestingBase_Skipped__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/e2eutils.TestingBase

// skipping Fuzz_Nosy_TestingBase_TempDir__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/e2eutils.TestingBase

func Fuzz_Nosy_ApplyDeployConfigForks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var deployConfig *genesis.DeployConfig
		fill_err = tp.Fill(&deployConfig)
		if fill_err != nil {
			return
		}
		if deployConfig == nil {
			return
		}

		ApplyDeployConfigForks(deployConfig)
	})
}

func Fuzz_Nosy_CollectAddresses__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sd *SetupData
		fill_err = tp.Fill(&sd)
		if fill_err != nil {
			return
		}
		var dp *DeployParams
		fill_err = tp.Fill(&dp)
		if fill_err != nil {
			return
		}
		if sd == nil || dp == nil {
			return
		}

		CollectAddresses(sd, dp)
	})
}

// skipping Fuzz_Nosy_WriteDefaultJWT__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/e2eutils.TestingBase

func Fuzz_Nosy_collectAllocAddrs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var alloc types.GenesisAlloc
		fill_err = tp.Fill(&alloc)
		if fill_err != nil {
			return
		}

		collectAllocAddrs(alloc)
	})
}
