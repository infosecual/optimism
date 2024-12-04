package verify

import (
	"context"
	"testing"

	"github.com/ethereum-optimism/optimism/op-service/sources"
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

func Fuzz_Nosy_Runner_RunBetweenBlocks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l1RpcUrl string
		fill_err = tp.Fill(&l1RpcUrl)
		if fill_err != nil {
			return
		}
		var l1RpcKind string
		fill_err = tp.Fill(&l1RpcKind)
		if fill_err != nil {
			return
		}
		var l1BeaconUrl string
		fill_err = tp.Fill(&l1BeaconUrl)
		if fill_err != nil {
			return
		}
		var l2RpcUrl string
		fill_err = tp.Fill(&l2RpcUrl)
		if fill_err != nil {
			return
		}
		var dataDir string
		fill_err = tp.Fill(&dataDir)
		if fill_err != nil {
			return
		}
		var network string
		fill_err = tp.Fill(&network)
		if fill_err != nil {
			return
		}
		var chainID uint64
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var runInProcess bool
		fill_err = tp.Fill(&runInProcess)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var l1Head common.Hash
		fill_err = tp.Fill(&l1Head)
		if fill_err != nil {
			return
		}
		var startBlockNum uint64
		fill_err = tp.Fill(&startBlockNum)
		if fill_err != nil {
			return
		}
		var endBlockNumber uint64
		fill_err = tp.Fill(&endBlockNumber)
		if fill_err != nil {
			return
		}

		r, err := NewRunner(l1RpcUrl, l1RpcKind, l1BeaconUrl, l2RpcUrl, dataDir, network, chainID, runInProcess)
		if err != nil {
			return
		}
		r.RunBetweenBlocks(ctx, l1Head, startBlockNum, endBlockNumber)
	})
}

func Fuzz_Nosy_Runner_RunToFinalized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l1RpcUrl string
		fill_err = tp.Fill(&l1RpcUrl)
		if fill_err != nil {
			return
		}
		var l1RpcKind string
		fill_err = tp.Fill(&l1RpcKind)
		if fill_err != nil {
			return
		}
		var l1BeaconUrl string
		fill_err = tp.Fill(&l1BeaconUrl)
		if fill_err != nil {
			return
		}
		var l2RpcUrl string
		fill_err = tp.Fill(&l2RpcUrl)
		if fill_err != nil {
			return
		}
		var dataDir string
		fill_err = tp.Fill(&dataDir)
		if fill_err != nil {
			return
		}
		var network string
		fill_err = tp.Fill(&network)
		if fill_err != nil {
			return
		}
		var chainID uint64
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var runInProcess bool
		fill_err = tp.Fill(&runInProcess)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		r, err := NewRunner(l1RpcUrl, l1RpcKind, l1BeaconUrl, l2RpcUrl, dataDir, network, chainID, runInProcess)
		if err != nil {
			return
		}
		r.RunToFinalized(ctx)
	})
}

func Fuzz_Nosy_Runner_createL2Client__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l1RpcUrl string
		fill_err = tp.Fill(&l1RpcUrl)
		if fill_err != nil {
			return
		}
		var l1RpcKind string
		fill_err = tp.Fill(&l1RpcKind)
		if fill_err != nil {
			return
		}
		var l1BeaconUrl string
		fill_err = tp.Fill(&l1BeaconUrl)
		if fill_err != nil {
			return
		}
		var l2RpcUrl string
		fill_err = tp.Fill(&l2RpcUrl)
		if fill_err != nil {
			return
		}
		var dataDir string
		fill_err = tp.Fill(&dataDir)
		if fill_err != nil {
			return
		}
		var network string
		fill_err = tp.Fill(&network)
		if fill_err != nil {
			return
		}
		var chainID uint64
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var runInProcess bool
		fill_err = tp.Fill(&runInProcess)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		r, err := NewRunner(l1RpcUrl, l1RpcKind, l1BeaconUrl, l2RpcUrl, dataDir, network, chainID, runInProcess)
		if err != nil {
			return
		}
		r.createL2Client(ctx)
	})
}

// skipping Fuzz_Nosy_Runner_run__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/eth.BlockInfo

func Fuzz_Nosy_outputAtBlockNum__(f *testing.F) {
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
		var l2Client *sources.L2Client
		fill_err = tp.Fill(&l2Client)
		if fill_err != nil {
			return
		}
		var blockNum uint64
		fill_err = tp.Fill(&blockNum)
		if fill_err != nil {
			return
		}
		if l2Client == nil {
			return
		}

		outputAtBlockNum(ctx, l2Client, blockNum)
	})
}

// skipping Fuzz_Nosy_retryOp__ because parameters include func, chan, or unsupported interface: func() (T, error)
