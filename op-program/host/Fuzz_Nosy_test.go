package host

import (
	"context"
	"testing"

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

func Fuzz_Nosy_L2Client_OutputByRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *L2Client
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var l2OutputRoot common.Hash
		fill_err = tp.Fill(&l2OutputRoot)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.OutputByRoot(ctx, l2OutputRoot)
	})
}

func Fuzz_Nosy_L2Source_CodeByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *L2Source
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		l.CodeByHash(ctx, hash)
	})
}

func Fuzz_Nosy_L2Source_ExecutionWitness__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *L2Source
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var blockNum uint64
		fill_err = tp.Fill(&blockNum)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		l.ExecutionWitness(ctx, blockNum)
	})
}

func Fuzz_Nosy_L2Source_ExperimentalEnabled__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *L2Source
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		l.ExperimentalEnabled()
	})
}

func Fuzz_Nosy_L2Source_GetProof__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *L2Source
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		var storage []common.Hash
		fill_err = tp.Fill(&storage)
		if fill_err != nil {
			return
		}
		var blockTag string
		fill_err = tp.Fill(&blockTag)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		l.GetProof(ctx, address, storage, blockTag)
	})
}

func Fuzz_Nosy_L2Source_InfoAndTxsByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *L2Source
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var blockHash common.Hash
		fill_err = tp.Fill(&blockHash)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		l.InfoAndTxsByHash(ctx, blockHash)
	})
}

func Fuzz_Nosy_L2Source_NodeByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *L2Source
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		l.NodeByHash(ctx, hash)
	})
}

func Fuzz_Nosy_L2Source_OutputByRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *L2Source
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var root common.Hash
		fill_err = tp.Fill(&root)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		l.OutputByRoot(ctx, root)
	})
}

// skipping Fuzz_Nosy_Prefetcher_GetPreimage__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-program/host.Prefetcher

// skipping Fuzz_Nosy_Prefetcher_Hint__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-program/host.Prefetcher

// skipping Fuzz_Nosy_launchOracleServer__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/log.Logger

// skipping Fuzz_Nosy_routeHints__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/log.Logger
