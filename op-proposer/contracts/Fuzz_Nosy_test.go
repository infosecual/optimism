package contracts

import (
	"context"
	"testing"
	"time"

	"github.com/ethereum-optimism/optimism/op-service/sources/batching"
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

func Fuzz_Nosy_DisputeGameFactory_HasProposedSince__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var caller *batching.MultiCaller
		fill_err = tp.Fill(&caller)
		if fill_err != nil {
			return
		}
		var networkTimeout time.Duration
		fill_err = tp.Fill(&networkTimeout)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var proposer common.Address
		fill_err = tp.Fill(&proposer)
		if fill_err != nil {
			return
		}
		var cutoff time.Time
		fill_err = tp.Fill(&cutoff)
		if fill_err != nil {
			return
		}
		var gameType uint32
		fill_err = tp.Fill(&gameType)
		if fill_err != nil {
			return
		}
		if caller == nil {
			return
		}

		f1 := NewDisputeGameFactory(addr, caller, networkTimeout)
		f1.HasProposedSince(ctx, proposer, cutoff, gameType)
	})
}

func Fuzz_Nosy_DisputeGameFactory_ProposalTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var caller *batching.MultiCaller
		fill_err = tp.Fill(&caller)
		if fill_err != nil {
			return
		}
		var networkTimeout time.Duration
		fill_err = tp.Fill(&networkTimeout)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var gameType uint32
		fill_err = tp.Fill(&gameType)
		if fill_err != nil {
			return
		}
		var outputRoot common.Hash
		fill_err = tp.Fill(&outputRoot)
		if fill_err != nil {
			return
		}
		var l2BlockNum uint64
		fill_err = tp.Fill(&l2BlockNum)
		if fill_err != nil {
			return
		}
		if caller == nil {
			return
		}

		f1 := NewDisputeGameFactory(addr, caller, networkTimeout)
		f1.ProposalTx(ctx, gameType, outputRoot, l2BlockNum)
	})
}

func Fuzz_Nosy_DisputeGameFactory_Version__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var caller *batching.MultiCaller
		fill_err = tp.Fill(&caller)
		if fill_err != nil {
			return
		}
		var networkTimeout time.Duration
		fill_err = tp.Fill(&networkTimeout)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if caller == nil {
			return
		}

		f1 := NewDisputeGameFactory(addr, caller, networkTimeout)
		f1.Version(ctx)
	})
}

func Fuzz_Nosy_DisputeGameFactory_gameAtIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var caller *batching.MultiCaller
		fill_err = tp.Fill(&caller)
		if fill_err != nil {
			return
		}
		var networkTimeout time.Duration
		fill_err = tp.Fill(&networkTimeout)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var idx uint64
		fill_err = tp.Fill(&idx)
		if fill_err != nil {
			return
		}
		if caller == nil {
			return
		}

		f1 := NewDisputeGameFactory(addr, caller, networkTimeout)
		f1.gameAtIndex(ctx, idx)
	})
}

func Fuzz_Nosy_DisputeGameFactory_gameCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var caller *batching.MultiCaller
		fill_err = tp.Fill(&caller)
		if fill_err != nil {
			return
		}
		var networkTimeout time.Duration
		fill_err = tp.Fill(&networkTimeout)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if caller == nil {
			return
		}

		f1 := NewDisputeGameFactory(addr, caller, networkTimeout)
		f1.gameCount(ctx)
	})
}
