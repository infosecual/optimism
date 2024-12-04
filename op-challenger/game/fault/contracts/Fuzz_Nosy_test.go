package contracts

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum-optimism/optimism/op-challenger/game/keccak/merkle"
	"github.com/ethereum-optimism/optimism/op-service/sources/batching"
	"github.com/ethereum-optimism/optimism/op-service/sources/batching/rpcblock"
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

func Fuzz_Nosy_DelayedWETHContract_Addr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *DelayedWETHContract
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.Addr()
	})
}

func Fuzz_Nosy_DelayedWETHContract_GetBalanceAndDelay__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *DelayedWETHContract
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var block rpcblock.Block
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.GetBalanceAndDelay(ctx, block)
	})
}

func Fuzz_Nosy_DelayedWETHContract_GetWithdrawals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *DelayedWETHContract
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var block rpcblock.Block
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		var gameAddr common.Address
		fill_err = tp.Fill(&gameAddr)
		if fill_err != nil {
			return
		}
		var recipients []common.Address
		fill_err = tp.Fill(&recipients)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.GetWithdrawals(ctx, block, gameAddr, recipients...)
	})
}

func Fuzz_Nosy_DisputeGameFactoryContract_CreateTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *DisputeGameFactoryContract
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var traceType uint32
		fill_err = tp.Fill(&traceType)
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
		if f1 == nil {
			return
		}

		f1.CreateTx(ctx, traceType, outputRoot, l2BlockNum)
	})
}

func Fuzz_Nosy_DisputeGameFactoryContract_DecodeDisputeGameCreatedLog__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *DisputeGameFactoryContract
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var rcpt *types.Receipt
		fill_err = tp.Fill(&rcpt)
		if fill_err != nil {
			return
		}
		if f1 == nil || rcpt == nil {
			return
		}

		f1.DecodeDisputeGameCreatedLog(rcpt)
	})
}

func Fuzz_Nosy_DisputeGameFactoryContract_GetAllGames__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *DisputeGameFactoryContract
		fill_err = tp.Fill(&f1)
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
		if f1 == nil {
			return
		}

		f1.GetAllGames(ctx, blockHash)
	})
}

func Fuzz_Nosy_DisputeGameFactoryContract_GetGame__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *DisputeGameFactoryContract
		fill_err = tp.Fill(&f1)
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
		var blockHash common.Hash
		fill_err = tp.Fill(&blockHash)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.GetGame(ctx, idx, blockHash)
	})
}

func Fuzz_Nosy_DisputeGameFactoryContract_GetGameCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *DisputeGameFactoryContract
		fill_err = tp.Fill(&f1)
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
		if f1 == nil {
			return
		}

		f1.GetGameCount(ctx, blockHash)
	})
}

func Fuzz_Nosy_DisputeGameFactoryContract_GetGameFromParameters__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *DisputeGameFactoryContract
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var traceType uint32
		fill_err = tp.Fill(&traceType)
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
		if f1 == nil {
			return
		}

		f1.GetGameFromParameters(ctx, traceType, outputRoot, l2BlockNum)
	})
}

func Fuzz_Nosy_DisputeGameFactoryContract_GetGameImpl__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *DisputeGameFactoryContract
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var gameType types.GameType
		fill_err = tp.Fill(&gameType)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.GetGameImpl(ctx, gameType)
	})
}

func Fuzz_Nosy_DisputeGameFactoryContract_GetGamesAtOrAfter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *DisputeGameFactoryContract
		fill_err = tp.Fill(&f1)
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
		var earliestTimestamp uint64
		fill_err = tp.Fill(&earliestTimestamp)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.GetGamesAtOrAfter(ctx, blockHash, earliestTimestamp)
	})
}

func Fuzz_Nosy_DisputeGameFactoryContract_decodeGame__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *DisputeGameFactoryContract
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var idx uint64
		fill_err = tp.Fill(&idx)
		if fill_err != nil {
			return
		}
		var result *batching.CallResult
		fill_err = tp.Fill(&result)
		if fill_err != nil {
			return
		}
		if f1 == nil || result == nil {
			return
		}

		f1.decodeGame(idx, result)
	})
}

func Fuzz_Nosy_FaultDisputeGameContract0180_AttackTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FaultDisputeGameContract0180
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var parent types.Claim
		fill_err = tp.Fill(&parent)
		if fill_err != nil {
			return
		}
		var pivot common.Hash
		fill_err = tp.Fill(&pivot)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.AttackTx(ctx, parent, pivot)
	})
}

func Fuzz_Nosy_FaultDisputeGameContract0180_ChallengeL2BlockNumberTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FaultDisputeGameContract0180
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var _x2 *types.InvalidL2BlockNumberChallenge
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if f1 == nil || _x2 == nil {
			return
		}

		f1.ChallengeL2BlockNumberTx(_x2)
	})
}

func Fuzz_Nosy_FaultDisputeGameContract0180_DefendTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FaultDisputeGameContract0180
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var parent types.Claim
		fill_err = tp.Fill(&parent)
		if fill_err != nil {
			return
		}
		var pivot common.Hash
		fill_err = tp.Fill(&pivot)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.DefendTx(ctx, parent, pivot)
	})
}

func Fuzz_Nosy_FaultDisputeGameContract0180_GetGameMetadata__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FaultDisputeGameContract0180
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var block rpcblock.Block
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.GetGameMetadata(ctx, block)
	})
}

func Fuzz_Nosy_FaultDisputeGameContract0180_IsL2BlockNumberChallenged__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FaultDisputeGameContract0180
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var _x3 rpcblock.Block
		fill_err = tp.Fill(&_x3)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.IsL2BlockNumberChallenged(_x2, _x3)
	})
}

func Fuzz_Nosy_FaultDisputeGameContract080_AttackTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FaultDisputeGameContract080
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var parent types.Claim
		fill_err = tp.Fill(&parent)
		if fill_err != nil {
			return
		}
		var pivot common.Hash
		fill_err = tp.Fill(&pivot)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.AttackTx(ctx, parent, pivot)
	})
}

func Fuzz_Nosy_FaultDisputeGameContract080_CallResolveClaim__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FaultDisputeGameContract080
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var claimIdx uint64
		fill_err = tp.Fill(&claimIdx)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.CallResolveClaim(ctx, claimIdx)
	})
}

func Fuzz_Nosy_FaultDisputeGameContract080_ChallengeL2BlockNumberTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FaultDisputeGameContract080
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var _x2 *types.InvalidL2BlockNumberChallenge
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if f1 == nil || _x2 == nil {
			return
		}

		f1.ChallengeL2BlockNumberTx(_x2)
	})
}

func Fuzz_Nosy_FaultDisputeGameContract080_DefendTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FaultDisputeGameContract080
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var parent types.Claim
		fill_err = tp.Fill(&parent)
		if fill_err != nil {
			return
		}
		var pivot common.Hash
		fill_err = tp.Fill(&pivot)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.DefendTx(ctx, parent, pivot)
	})
}

func Fuzz_Nosy_FaultDisputeGameContract080_GetAllClaims__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FaultDisputeGameContract080
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var block rpcblock.Block
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.GetAllClaims(ctx, block)
	})
}

func Fuzz_Nosy_FaultDisputeGameContract080_GetClaim__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FaultDisputeGameContract080
		fill_err = tp.Fill(&f1)
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
		if f1 == nil {
			return
		}

		f1.GetClaim(ctx, idx)
	})
}

func Fuzz_Nosy_FaultDisputeGameContract080_GetGameMetadata__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FaultDisputeGameContract080
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var block rpcblock.Block
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.GetGameMetadata(ctx, block)
	})
}

func Fuzz_Nosy_FaultDisputeGameContract080_GetMaxClockDuration__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FaultDisputeGameContract080
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.GetMaxClockDuration(ctx)
	})
}

func Fuzz_Nosy_FaultDisputeGameContract080_IsL2BlockNumberChallenged__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FaultDisputeGameContract080
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var _x3 rpcblock.Block
		fill_err = tp.Fill(&_x3)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.IsL2BlockNumberChallenged(_x2, _x3)
	})
}

func Fuzz_Nosy_FaultDisputeGameContract080_IsResolved__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FaultDisputeGameContract080
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var block rpcblock.Block
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		var claims []types.Claim
		fill_err = tp.Fill(&claims)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.IsResolved(ctx, block, claims...)
	})
}

func Fuzz_Nosy_FaultDisputeGameContract080_ResolveClaimTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FaultDisputeGameContract080
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var claimIdx uint64
		fill_err = tp.Fill(&claimIdx)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.ResolveClaimTx(claimIdx)
	})
}

func Fuzz_Nosy_FaultDisputeGameContract080_resolveClaimCall__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FaultDisputeGameContract080
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var claimIdx uint64
		fill_err = tp.Fill(&claimIdx)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.resolveClaimCall(claimIdx)
	})
}

func Fuzz_Nosy_FaultDisputeGameContract111_AttackTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FaultDisputeGameContract111
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var parent types.Claim
		fill_err = tp.Fill(&parent)
		if fill_err != nil {
			return
		}
		var pivot common.Hash
		fill_err = tp.Fill(&pivot)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.AttackTx(ctx, parent, pivot)
	})
}

func Fuzz_Nosy_FaultDisputeGameContract111_DefendTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FaultDisputeGameContract111
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var parent types.Claim
		fill_err = tp.Fill(&parent)
		if fill_err != nil {
			return
		}
		var pivot common.Hash
		fill_err = tp.Fill(&pivot)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.DefendTx(ctx, parent, pivot)
	})
}

func Fuzz_Nosy_FaultDisputeGameContractLatest_AttackTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FaultDisputeGameContractLatest
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var parent types.Claim
		fill_err = tp.Fill(&parent)
		if fill_err != nil {
			return
		}
		var pivot common.Hash
		fill_err = tp.Fill(&pivot)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.AttackTx(ctx, parent, pivot)
	})
}

func Fuzz_Nosy_FaultDisputeGameContractLatest_CallResolve__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FaultDisputeGameContractLatest
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.CallResolve(ctx)
	})
}

func Fuzz_Nosy_FaultDisputeGameContractLatest_CallResolveClaim__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FaultDisputeGameContractLatest
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var claimIdx uint64
		fill_err = tp.Fill(&claimIdx)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.CallResolveClaim(ctx, claimIdx)
	})
}

func Fuzz_Nosy_FaultDisputeGameContractLatest_ChallengeL2BlockNumberTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FaultDisputeGameContractLatest
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var challenge *types.InvalidL2BlockNumberChallenge
		fill_err = tp.Fill(&challenge)
		if fill_err != nil {
			return
		}
		if f1 == nil || challenge == nil {
			return
		}

		f1.ChallengeL2BlockNumberTx(challenge)
	})
}

func Fuzz_Nosy_FaultDisputeGameContractLatest_ClaimCreditTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FaultDisputeGameContractLatest
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var recipient common.Address
		fill_err = tp.Fill(&recipient)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.ClaimCreditTx(ctx, recipient)
	})
}

func Fuzz_Nosy_FaultDisputeGameContractLatest_DefendTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FaultDisputeGameContractLatest
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var parent types.Claim
		fill_err = tp.Fill(&parent)
		if fill_err != nil {
			return
		}
		var pivot common.Hash
		fill_err = tp.Fill(&pivot)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.DefendTx(ctx, parent, pivot)
	})
}

func Fuzz_Nosy_FaultDisputeGameContractLatest_GetAbsolutePrestateHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FaultDisputeGameContractLatest
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.GetAbsolutePrestateHash(ctx)
	})
}

func Fuzz_Nosy_FaultDisputeGameContractLatest_GetAllClaims__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FaultDisputeGameContractLatest
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var block rpcblock.Block
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.GetAllClaims(ctx, block)
	})
}

func Fuzz_Nosy_FaultDisputeGameContractLatest_GetBalanceAndDelay__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FaultDisputeGameContractLatest
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var block rpcblock.Block
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.GetBalanceAndDelay(ctx, block)
	})
}

func Fuzz_Nosy_FaultDisputeGameContractLatest_GetBlockRange__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FaultDisputeGameContractLatest
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.GetBlockRange(ctx)
	})
}

func Fuzz_Nosy_FaultDisputeGameContractLatest_GetClaim__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FaultDisputeGameContractLatest
		fill_err = tp.Fill(&f1)
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
		if f1 == nil {
			return
		}

		f1.GetClaim(ctx, idx)
	})
}

func Fuzz_Nosy_FaultDisputeGameContractLatest_GetClaimCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FaultDisputeGameContractLatest
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.GetClaimCount(ctx)
	})
}

func Fuzz_Nosy_FaultDisputeGameContractLatest_GetCredit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FaultDisputeGameContractLatest
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var recipient common.Address
		fill_err = tp.Fill(&recipient)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.GetCredit(ctx, recipient)
	})
}

func Fuzz_Nosy_FaultDisputeGameContractLatest_GetCredits__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FaultDisputeGameContractLatest
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var block rpcblock.Block
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		var recipients []common.Address
		fill_err = tp.Fill(&recipients)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.GetCredits(ctx, block, recipients...)
	})
}

func Fuzz_Nosy_FaultDisputeGameContractLatest_GetGameMetadata__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FaultDisputeGameContractLatest
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var block rpcblock.Block
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.GetGameMetadata(ctx, block)
	})
}

func Fuzz_Nosy_FaultDisputeGameContractLatest_GetL1Head__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FaultDisputeGameContractLatest
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.GetL1Head(ctx)
	})
}

func Fuzz_Nosy_FaultDisputeGameContractLatest_GetMaxClockDuration__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FaultDisputeGameContractLatest
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.GetMaxClockDuration(ctx)
	})
}

func Fuzz_Nosy_FaultDisputeGameContractLatest_GetMaxGameDepth__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FaultDisputeGameContractLatest
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.GetMaxGameDepth(ctx)
	})
}

func Fuzz_Nosy_FaultDisputeGameContractLatest_GetOracle__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FaultDisputeGameContractLatest
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.GetOracle(ctx)
	})
}

func Fuzz_Nosy_FaultDisputeGameContractLatest_GetRequiredBond__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FaultDisputeGameContractLatest
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var position types.Position
		fill_err = tp.Fill(&position)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.GetRequiredBond(ctx, position)
	})
}

func Fuzz_Nosy_FaultDisputeGameContractLatest_GetRequiredBonds__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FaultDisputeGameContractLatest
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var block rpcblock.Block
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		var positions []*big.Int
		fill_err = tp.Fill(&positions)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.GetRequiredBonds(ctx, block, positions...)
	})
}

func Fuzz_Nosy_FaultDisputeGameContractLatest_GetResolvedAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FaultDisputeGameContractLatest
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var block rpcblock.Block
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.GetResolvedAt(ctx, block)
	})
}

func Fuzz_Nosy_FaultDisputeGameContractLatest_GetSplitDepth__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FaultDisputeGameContractLatest
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.GetSplitDepth(ctx)
	})
}

func Fuzz_Nosy_FaultDisputeGameContractLatest_GetStartingRootHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FaultDisputeGameContractLatest
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.GetStartingRootHash(ctx)
	})
}

func Fuzz_Nosy_FaultDisputeGameContractLatest_GetStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FaultDisputeGameContractLatest
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.GetStatus(ctx)
	})
}

func Fuzz_Nosy_FaultDisputeGameContractLatest_GetWithdrawals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FaultDisputeGameContractLatest
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var block rpcblock.Block
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		var recipients []common.Address
		fill_err = tp.Fill(&recipients)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.GetWithdrawals(ctx, block, recipients...)
	})
}

func Fuzz_Nosy_FaultDisputeGameContractLatest_IsL2BlockNumberChallenged__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FaultDisputeGameContractLatest
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var block rpcblock.Block
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.IsL2BlockNumberChallenged(ctx, block)
	})
}

func Fuzz_Nosy_FaultDisputeGameContractLatest_IsResolved__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FaultDisputeGameContractLatest
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var block rpcblock.Block
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		var claims []types.Claim
		fill_err = tp.Fill(&claims)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.IsResolved(ctx, block, claims...)
	})
}

func Fuzz_Nosy_FaultDisputeGameContractLatest_ResolveClaimTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FaultDisputeGameContractLatest
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var claimIdx uint64
		fill_err = tp.Fill(&claimIdx)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.ResolveClaimTx(claimIdx)
	})
}

func Fuzz_Nosy_FaultDisputeGameContractLatest_ResolveTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FaultDisputeGameContractLatest
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.ResolveTx()
	})
}

func Fuzz_Nosy_FaultDisputeGameContractLatest_StepTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FaultDisputeGameContractLatest
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var claimIdx uint64
		fill_err = tp.Fill(&claimIdx)
		if fill_err != nil {
			return
		}
		var isAttack bool
		fill_err = tp.Fill(&isAttack)
		if fill_err != nil {
			return
		}
		var stateData []byte
		fill_err = tp.Fill(&stateData)
		if fill_err != nil {
			return
		}
		var proof []byte
		fill_err = tp.Fill(&proof)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.StepTx(claimIdx, isAttack, stateData, proof)
	})
}

func Fuzz_Nosy_FaultDisputeGameContractLatest_UpdateOracleTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FaultDisputeGameContractLatest
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var claimIdx uint64
		fill_err = tp.Fill(&claimIdx)
		if fill_err != nil {
			return
		}
		var d4 *types.PreimageOracleData
		fill_err = tp.Fill(&d4)
		if fill_err != nil {
			return
		}
		if f1 == nil || d4 == nil {
			return
		}

		f1.UpdateOracleTx(ctx, claimIdx, d4)
	})
}

func Fuzz_Nosy_FaultDisputeGameContractLatest_Vm__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FaultDisputeGameContractLatest
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.Vm(ctx)
	})
}

func Fuzz_Nosy_FaultDisputeGameContractLatest_addGlobalDataTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FaultDisputeGameContractLatest
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var d3 *types.PreimageOracleData
		fill_err = tp.Fill(&d3)
		if fill_err != nil {
			return
		}
		if f1 == nil || d3 == nil {
			return
		}

		f1.addGlobalDataTx(ctx, d3)
	})
}

func Fuzz_Nosy_FaultDisputeGameContractLatest_addLocalDataTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FaultDisputeGameContractLatest
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var claimIdx uint64
		fill_err = tp.Fill(&claimIdx)
		if fill_err != nil {
			return
		}
		var d3 *types.PreimageOracleData
		fill_err = tp.Fill(&d3)
		if fill_err != nil {
			return
		}
		if f1 == nil || d3 == nil {
			return
		}

		f1.addLocalDataTx(claimIdx, d3)
	})
}

func Fuzz_Nosy_FaultDisputeGameContractLatest_decodeClaim__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FaultDisputeGameContractLatest
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var result *batching.CallResult
		fill_err = tp.Fill(&result)
		if fill_err != nil {
			return
		}
		var contractIndex int
		fill_err = tp.Fill(&contractIndex)
		if fill_err != nil {
			return
		}
		if f1 == nil || result == nil {
			return
		}

		f1.decodeClaim(result, contractIndex)
	})
}

func Fuzz_Nosy_FaultDisputeGameContractLatest_getDelayedWETH__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FaultDisputeGameContractLatest
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var block rpcblock.Block
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.getDelayedWETH(ctx, block)
	})
}

func Fuzz_Nosy_FaultDisputeGameContractLatest_resolveCall__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FaultDisputeGameContractLatest
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.resolveCall()
	})
}

func Fuzz_Nosy_FaultDisputeGameContractLatest_resolveClaimCall__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FaultDisputeGameContractLatest
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var claimIdx uint64
		fill_err = tp.Fill(&claimIdx)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.resolveClaimCall(claimIdx)
	})
}

func Fuzz_Nosy_FaultDisputeGameContractLatest_txWithBond__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FaultDisputeGameContractLatest
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var position types.Position
		fill_err = tp.Fill(&position)
		if fill_err != nil {
			return
		}
		var call *batching.ContractCall
		fill_err = tp.Fill(&call)
		if fill_err != nil {
			return
		}
		if f1 == nil || call == nil {
			return
		}

		f1.txWithBond(ctx, position, call)
	})
}

func Fuzz_Nosy_PreimageOracleContract100_AddGlobalDataTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *PreimageOracleContract100
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var d2 *types.PreimageOracleData
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}
		if c == nil || d2 == nil {
			return
		}

		c.AddGlobalDataTx(d2)
	})
}

func Fuzz_Nosy_PreimageOracleContractLatest_AddGlobalDataTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *PreimageOracleContractLatest
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var d2 *types.PreimageOracleData
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}
		if c == nil || d2 == nil {
			return
		}

		c.AddGlobalDataTx(d2)
	})
}

func Fuzz_Nosy_PreimageOracleContractLatest_AddLeaves__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *PreimageOracleContractLatest
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var uuid *big.Int
		fill_err = tp.Fill(&uuid)
		if fill_err != nil {
			return
		}
		var startingBlockIndex *big.Int
		fill_err = tp.Fill(&startingBlockIndex)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		var commitments []common.Hash
		fill_err = tp.Fill(&commitments)
		if fill_err != nil {
			return
		}
		var finalize bool
		fill_err = tp.Fill(&finalize)
		if fill_err != nil {
			return
		}
		if c == nil || uuid == nil || startingBlockIndex == nil {
			return
		}

		c.AddLeaves(uuid, startingBlockIndex, input, commitments, finalize)
	})
}

func Fuzz_Nosy_PreimageOracleContractLatest_Addr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *PreimageOracleContractLatest
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.Addr()
	})
}

func Fuzz_Nosy_PreimageOracleContractLatest_CallSqueeze__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *PreimageOracleContractLatest
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var claimant common.Address
		fill_err = tp.Fill(&claimant)
		if fill_err != nil {
			return
		}
		var uuid *big.Int
		fill_err = tp.Fill(&uuid)
		if fill_err != nil {
			return
		}
		var prestateMatrix types.StateSnapshot
		fill_err = tp.Fill(&prestateMatrix)
		if fill_err != nil {
			return
		}
		var preState types.Leaf
		fill_err = tp.Fill(&preState)
		if fill_err != nil {
			return
		}
		var preStateProof merkle.Proof
		fill_err = tp.Fill(&preStateProof)
		if fill_err != nil {
			return
		}
		var postState types.Leaf
		fill_err = tp.Fill(&postState)
		if fill_err != nil {
			return
		}
		var postStateProof merkle.Proof
		fill_err = tp.Fill(&postStateProof)
		if fill_err != nil {
			return
		}
		if c == nil || uuid == nil {
			return
		}

		c.CallSqueeze(ctx, claimant, uuid, prestateMatrix, preState, preStateProof, postState, postStateProof)
	})
}

func Fuzz_Nosy_PreimageOracleContractLatest_ChallengePeriod__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *PreimageOracleContractLatest
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.ChallengePeriod(ctx)
	})
}

func Fuzz_Nosy_PreimageOracleContractLatest_ChallengeTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *PreimageOracleContractLatest
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ident types.LargePreimageIdent
		fill_err = tp.Fill(&ident)
		if fill_err != nil {
			return
		}
		var challenge types.Challenge
		fill_err = tp.Fill(&challenge)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.ChallengeTx(ident, challenge)
	})
}

func Fuzz_Nosy_PreimageOracleContractLatest_DecodeInputData__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *PreimageOracleContractLatest
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var d2 []byte
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.DecodeInputData(d2)
	})
}

func Fuzz_Nosy_PreimageOracleContractLatest_GetActivePreimages__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *PreimageOracleContractLatest
		fill_err = tp.Fill(&c)
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
		if c == nil {
			return
		}

		c.GetActivePreimages(ctx, blockHash)
	})
}

func Fuzz_Nosy_PreimageOracleContractLatest_GetInputDataBlocks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *PreimageOracleContractLatest
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var block rpcblock.Block
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		var ident types.LargePreimageIdent
		fill_err = tp.Fill(&ident)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.GetInputDataBlocks(ctx, block, ident)
	})
}

func Fuzz_Nosy_PreimageOracleContractLatest_GetMinBondLPP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *PreimageOracleContractLatest
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.GetMinBondLPP(ctx)
	})
}

func Fuzz_Nosy_PreimageOracleContractLatest_GetProposalMetadata__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *PreimageOracleContractLatest
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var block rpcblock.Block
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		var idents []types.LargePreimageIdent
		fill_err = tp.Fill(&idents)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.GetProposalMetadata(ctx, block, idents...)
	})
}

func Fuzz_Nosy_PreimageOracleContractLatest_GetProposalTreeRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *PreimageOracleContractLatest
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var block rpcblock.Block
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		var ident types.LargePreimageIdent
		fill_err = tp.Fill(&ident)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.GetProposalTreeRoot(ctx, block, ident)
	})
}

func Fuzz_Nosy_PreimageOracleContractLatest_GlobalDataExists__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *PreimageOracleContractLatest
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var d3 *types.PreimageOracleData
		fill_err = tp.Fill(&d3)
		if fill_err != nil {
			return
		}
		if c == nil || d3 == nil {
			return
		}

		c.GlobalDataExists(ctx, d3)
	})
}

func Fuzz_Nosy_PreimageOracleContractLatest_InitLargePreimage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *PreimageOracleContractLatest
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var uuid *big.Int
		fill_err = tp.Fill(&uuid)
		if fill_err != nil {
			return
		}
		var partOffset uint32
		fill_err = tp.Fill(&partOffset)
		if fill_err != nil {
			return
		}
		var claimedSize uint32
		fill_err = tp.Fill(&claimedSize)
		if fill_err != nil {
			return
		}
		if c == nil || uuid == nil {
			return
		}

		c.InitLargePreimage(uuid, partOffset, claimedSize)
	})
}

func Fuzz_Nosy_PreimageOracleContractLatest_MinLargePreimageSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *PreimageOracleContractLatest
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.MinLargePreimageSize(ctx)
	})
}

func Fuzz_Nosy_PreimageOracleContractLatest_Squeeze__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *PreimageOracleContractLatest
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var claimant common.Address
		fill_err = tp.Fill(&claimant)
		if fill_err != nil {
			return
		}
		var uuid *big.Int
		fill_err = tp.Fill(&uuid)
		if fill_err != nil {
			return
		}
		var prestateMatrix types.StateSnapshot
		fill_err = tp.Fill(&prestateMatrix)
		if fill_err != nil {
			return
		}
		var preState types.Leaf
		fill_err = tp.Fill(&preState)
		if fill_err != nil {
			return
		}
		var preStateProof merkle.Proof
		fill_err = tp.Fill(&preStateProof)
		if fill_err != nil {
			return
		}
		var postState types.Leaf
		fill_err = tp.Fill(&postState)
		if fill_err != nil {
			return
		}
		var postStateProof merkle.Proof
		fill_err = tp.Fill(&postStateProof)
		if fill_err != nil {
			return
		}
		if c == nil || uuid == nil {
			return
		}

		c.Squeeze(claimant, uuid, prestateMatrix, preState, preStateProof, postState, postStateProof)
	})
}

func Fuzz_Nosy_PreimageOracleContractLatest_decodePreimageIdent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *PreimageOracleContractLatest
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var result *batching.CallResult
		fill_err = tp.Fill(&result)
		if fill_err != nil {
			return
		}
		if c == nil || result == nil {
			return
		}

		c.decodePreimageIdent(result)
	})
}

func Fuzz_Nosy_VMContract_Addr__(f *testing.F) {
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
		if caller == nil {
			return
		}

		c := NewVMContract(addr, caller)
		c.Addr()
	})
}

func Fuzz_Nosy_VMContract_Oracle__(f *testing.F) {
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
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if caller == nil {
			return
		}

		c := NewVMContract(addr, caller)
		c.Oracle(ctx)
	})
}

// skipping Fuzz_Nosy_VersionedBuilder[C any]_AddVersion__ because parameters include func, chan, or unsupported interface: func() (C, error)

// skipping Fuzz_Nosy_VersionedBuilder[C any]_Build__ because parameters include func, chan, or unsupported interface: func() (C, error)

func Fuzz_Nosy_metadata_blocksProcessed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *metadata
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.blocksProcessed()
	})
}

func Fuzz_Nosy_metadata_bytesProcessed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *metadata
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.bytesProcessed()
	})
}

func Fuzz_Nosy_metadata_claimedSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *metadata
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.claimedSize()
	})
}

func Fuzz_Nosy_metadata_countered__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *metadata
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.countered()
	})
}

func Fuzz_Nosy_metadata_partOffset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *metadata
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.partOffset()
	})
}

func Fuzz_Nosy_metadata_setBlocksProcessed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *metadata
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var value uint32
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.setBlocksProcessed(value)
	})
}

func Fuzz_Nosy_metadata_setBytesProcessed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *metadata
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var value uint32
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.setBytesProcessed(value)
	})
}

func Fuzz_Nosy_metadata_setClaimedSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *metadata
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var value uint32
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.setClaimedSize(value)
	})
}

func Fuzz_Nosy_metadata_setCountered__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *metadata
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var value bool
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.setCountered(value)
	})
}

func Fuzz_Nosy_metadata_setPartOffset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *metadata
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var value uint32
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.setPartOffset(value)
	})
}

func Fuzz_Nosy_metadata_setTimestamp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *metadata
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var timestamp uint64
		fill_err = tp.Fill(&timestamp)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.setTimestamp(timestamp)
	})
}

func Fuzz_Nosy_metadata_timestamp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *metadata
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.timestamp()
	})
}

// skipping Fuzz_Nosy_FaultDisputeGameContract_AttackTx__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/contracts.FaultDisputeGameContract

// skipping Fuzz_Nosy_FaultDisputeGameContract_CallResolve__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/contracts.FaultDisputeGameContract

// skipping Fuzz_Nosy_FaultDisputeGameContract_CallResolveClaim__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/contracts.FaultDisputeGameContract

// skipping Fuzz_Nosy_FaultDisputeGameContract_ChallengeL2BlockNumberTx__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/contracts.FaultDisputeGameContract

// skipping Fuzz_Nosy_FaultDisputeGameContract_ClaimCreditTx__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/contracts.FaultDisputeGameContract

// skipping Fuzz_Nosy_FaultDisputeGameContract_DefendTx__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/contracts.FaultDisputeGameContract

// skipping Fuzz_Nosy_FaultDisputeGameContract_GetAbsolutePrestateHash__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/contracts.FaultDisputeGameContract

// skipping Fuzz_Nosy_FaultDisputeGameContract_GetAllClaims__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/contracts.FaultDisputeGameContract

// skipping Fuzz_Nosy_FaultDisputeGameContract_GetBalanceAndDelay__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/contracts.FaultDisputeGameContract

// skipping Fuzz_Nosy_FaultDisputeGameContract_GetBlockRange__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/contracts.FaultDisputeGameContract

// skipping Fuzz_Nosy_FaultDisputeGameContract_GetClaim__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/contracts.FaultDisputeGameContract

// skipping Fuzz_Nosy_FaultDisputeGameContract_GetClaimCount__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/contracts.FaultDisputeGameContract

// skipping Fuzz_Nosy_FaultDisputeGameContract_GetCredit__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/contracts.FaultDisputeGameContract

// skipping Fuzz_Nosy_FaultDisputeGameContract_GetCredits__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/contracts.FaultDisputeGameContract

// skipping Fuzz_Nosy_FaultDisputeGameContract_GetGameMetadata__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/contracts.FaultDisputeGameContract

// skipping Fuzz_Nosy_FaultDisputeGameContract_GetL1Head__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/contracts.FaultDisputeGameContract

// skipping Fuzz_Nosy_FaultDisputeGameContract_GetMaxClockDuration__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/contracts.FaultDisputeGameContract

// skipping Fuzz_Nosy_FaultDisputeGameContract_GetMaxGameDepth__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/contracts.FaultDisputeGameContract

// skipping Fuzz_Nosy_FaultDisputeGameContract_GetOracle__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/contracts.FaultDisputeGameContract

// skipping Fuzz_Nosy_FaultDisputeGameContract_GetRequiredBond__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/contracts.FaultDisputeGameContract

// skipping Fuzz_Nosy_FaultDisputeGameContract_GetRequiredBonds__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/contracts.FaultDisputeGameContract

// skipping Fuzz_Nosy_FaultDisputeGameContract_GetResolvedAt__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/contracts.FaultDisputeGameContract

// skipping Fuzz_Nosy_FaultDisputeGameContract_GetSplitDepth__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/contracts.FaultDisputeGameContract

// skipping Fuzz_Nosy_FaultDisputeGameContract_GetStartingRootHash__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/contracts.FaultDisputeGameContract

// skipping Fuzz_Nosy_FaultDisputeGameContract_GetStatus__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/contracts.FaultDisputeGameContract

// skipping Fuzz_Nosy_FaultDisputeGameContract_GetWithdrawals__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/contracts.FaultDisputeGameContract

// skipping Fuzz_Nosy_FaultDisputeGameContract_IsL2BlockNumberChallenged__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/contracts.FaultDisputeGameContract

// skipping Fuzz_Nosy_FaultDisputeGameContract_IsResolved__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/contracts.FaultDisputeGameContract

// skipping Fuzz_Nosy_FaultDisputeGameContract_ResolveClaimTx__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/contracts.FaultDisputeGameContract

// skipping Fuzz_Nosy_FaultDisputeGameContract_ResolveTx__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/contracts.FaultDisputeGameContract

// skipping Fuzz_Nosy_FaultDisputeGameContract_StepTx__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/contracts.FaultDisputeGameContract

// skipping Fuzz_Nosy_FaultDisputeGameContract_UpdateOracleTx__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/contracts.FaultDisputeGameContract

// skipping Fuzz_Nosy_FaultDisputeGameContract_Vm__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/contracts.FaultDisputeGameContract

func Fuzz_Nosy_PreimageOracleContract_AddGlobalDataTx__(f *testing.F) {
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
		var d4 *types.PreimageOracleData
		fill_err = tp.Fill(&d4)
		if fill_err != nil {
			return
		}
		if caller == nil || d4 == nil {
			return
		}

		_x1, err := NewPreimageOracleContract(ctx, addr, caller)
		if err != nil {
			return
		}
		_x1.AddGlobalDataTx(d4)
	})
}

func Fuzz_Nosy_PreimageOracleContract_AddLeaves__(f *testing.F) {
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
		var uuid *big.Int
		fill_err = tp.Fill(&uuid)
		if fill_err != nil {
			return
		}
		var startingBlockIndex *big.Int
		fill_err = tp.Fill(&startingBlockIndex)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		var commitments []common.Hash
		fill_err = tp.Fill(&commitments)
		if fill_err != nil {
			return
		}
		var finalize bool
		fill_err = tp.Fill(&finalize)
		if fill_err != nil {
			return
		}
		if caller == nil || uuid == nil || startingBlockIndex == nil {
			return
		}

		_x1, err := NewPreimageOracleContract(ctx, addr, caller)
		if err != nil {
			return
		}
		_x1.AddLeaves(uuid, startingBlockIndex, input, commitments, finalize)
	})
}

func Fuzz_Nosy_PreimageOracleContract_Addr__(f *testing.F) {
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
		if caller == nil {
			return
		}

		_x1, err := NewPreimageOracleContract(ctx, addr, caller)
		if err != nil {
			return
		}
		_x1.Addr()
	})
}

func Fuzz_Nosy_PreimageOracleContract_CallSqueeze__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
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
		var c4 context.Context
		fill_err = tp.Fill(&c4)
		if fill_err != nil {
			return
		}
		var claimant common.Address
		fill_err = tp.Fill(&claimant)
		if fill_err != nil {
			return
		}
		var uuid *big.Int
		fill_err = tp.Fill(&uuid)
		if fill_err != nil {
			return
		}
		var prestateMatrix types.StateSnapshot
		fill_err = tp.Fill(&prestateMatrix)
		if fill_err != nil {
			return
		}
		var preState types.Leaf
		fill_err = tp.Fill(&preState)
		if fill_err != nil {
			return
		}
		var preStateProof merkle.Proof
		fill_err = tp.Fill(&preStateProof)
		if fill_err != nil {
			return
		}
		var postState types.Leaf
		fill_err = tp.Fill(&postState)
		if fill_err != nil {
			return
		}
		var postStateProof merkle.Proof
		fill_err = tp.Fill(&postStateProof)
		if fill_err != nil {
			return
		}
		if caller == nil || uuid == nil {
			return
		}

		_x1, err := NewPreimageOracleContract(c1, addr, caller)
		if err != nil {
			return
		}
		_x1.CallSqueeze(c4, claimant, uuid, prestateMatrix, preState, preStateProof, postState, postStateProof)
	})
}

func Fuzz_Nosy_PreimageOracleContract_ChallengePeriod__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
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
		var c4 context.Context
		fill_err = tp.Fill(&c4)
		if fill_err != nil {
			return
		}
		if caller == nil {
			return
		}

		_x1, err := NewPreimageOracleContract(c1, addr, caller)
		if err != nil {
			return
		}
		_x1.ChallengePeriod(c4)
	})
}

func Fuzz_Nosy_PreimageOracleContract_ChallengeTx__(f *testing.F) {
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
		var ident types.LargePreimageIdent
		fill_err = tp.Fill(&ident)
		if fill_err != nil {
			return
		}
		var challenge types.Challenge
		fill_err = tp.Fill(&challenge)
		if fill_err != nil {
			return
		}
		if caller == nil {
			return
		}

		_x1, err := NewPreimageOracleContract(ctx, addr, caller)
		if err != nil {
			return
		}
		_x1.ChallengeTx(ident, challenge)
	})
}

func Fuzz_Nosy_PreimageOracleContract_DecodeInputData__(f *testing.F) {
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
		var d4 []byte
		fill_err = tp.Fill(&d4)
		if fill_err != nil {
			return
		}
		if caller == nil {
			return
		}

		_x1, err := NewPreimageOracleContract(ctx, addr, caller)
		if err != nil {
			return
		}
		_x1.DecodeInputData(d4)
	})
}

func Fuzz_Nosy_PreimageOracleContract_GetActivePreimages__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
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
		var c4 context.Context
		fill_err = tp.Fill(&c4)
		if fill_err != nil {
			return
		}
		var blockHash common.Hash
		fill_err = tp.Fill(&blockHash)
		if fill_err != nil {
			return
		}
		if caller == nil {
			return
		}

		_x1, err := NewPreimageOracleContract(c1, addr, caller)
		if err != nil {
			return
		}
		_x1.GetActivePreimages(c4, blockHash)
	})
}

func Fuzz_Nosy_PreimageOracleContract_GetInputDataBlocks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
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
		var c4 context.Context
		fill_err = tp.Fill(&c4)
		if fill_err != nil {
			return
		}
		var block rpcblock.Block
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		var ident types.LargePreimageIdent
		fill_err = tp.Fill(&ident)
		if fill_err != nil {
			return
		}
		if caller == nil {
			return
		}

		_x1, err := NewPreimageOracleContract(c1, addr, caller)
		if err != nil {
			return
		}
		_x1.GetInputDataBlocks(c4, block, ident)
	})
}

func Fuzz_Nosy_PreimageOracleContract_GetMinBondLPP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
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
		var c4 context.Context
		fill_err = tp.Fill(&c4)
		if fill_err != nil {
			return
		}
		if caller == nil {
			return
		}

		_x1, err := NewPreimageOracleContract(c1, addr, caller)
		if err != nil {
			return
		}
		_x1.GetMinBondLPP(c4)
	})
}

func Fuzz_Nosy_PreimageOracleContract_GetProposalMetadata__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
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
		var c4 context.Context
		fill_err = tp.Fill(&c4)
		if fill_err != nil {
			return
		}
		var block rpcblock.Block
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		var idents []types.LargePreimageIdent
		fill_err = tp.Fill(&idents)
		if fill_err != nil {
			return
		}
		if caller == nil {
			return
		}

		_x1, err := NewPreimageOracleContract(c1, addr, caller)
		if err != nil {
			return
		}
		_x1.GetProposalMetadata(c4, block, idents...)
	})
}

func Fuzz_Nosy_PreimageOracleContract_GetProposalTreeRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
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
		var c4 context.Context
		fill_err = tp.Fill(&c4)
		if fill_err != nil {
			return
		}
		var block rpcblock.Block
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		var ident types.LargePreimageIdent
		fill_err = tp.Fill(&ident)
		if fill_err != nil {
			return
		}
		if caller == nil {
			return
		}

		_x1, err := NewPreimageOracleContract(c1, addr, caller)
		if err != nil {
			return
		}
		_x1.GetProposalTreeRoot(c4, block, ident)
	})
}

func Fuzz_Nosy_PreimageOracleContract_GlobalDataExists__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
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
		var c4 context.Context
		fill_err = tp.Fill(&c4)
		if fill_err != nil {
			return
		}
		var d5 *types.PreimageOracleData
		fill_err = tp.Fill(&d5)
		if fill_err != nil {
			return
		}
		if caller == nil || d5 == nil {
			return
		}

		_x1, err := NewPreimageOracleContract(c1, addr, caller)
		if err != nil {
			return
		}
		_x1.GlobalDataExists(c4, d5)
	})
}

func Fuzz_Nosy_PreimageOracleContract_InitLargePreimage__(f *testing.F) {
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
		var uuid *big.Int
		fill_err = tp.Fill(&uuid)
		if fill_err != nil {
			return
		}
		var partOffset uint32
		fill_err = tp.Fill(&partOffset)
		if fill_err != nil {
			return
		}
		var claimedSize uint32
		fill_err = tp.Fill(&claimedSize)
		if fill_err != nil {
			return
		}
		if caller == nil || uuid == nil {
			return
		}

		_x1, err := NewPreimageOracleContract(ctx, addr, caller)
		if err != nil {
			return
		}
		_x1.InitLargePreimage(uuid, partOffset, claimedSize)
	})
}

func Fuzz_Nosy_PreimageOracleContract_MinLargePreimageSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
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
		var c4 context.Context
		fill_err = tp.Fill(&c4)
		if fill_err != nil {
			return
		}
		if caller == nil {
			return
		}

		_x1, err := NewPreimageOracleContract(c1, addr, caller)
		if err != nil {
			return
		}
		_x1.MinLargePreimageSize(c4)
	})
}

func Fuzz_Nosy_PreimageOracleContract_Squeeze__(f *testing.F) {
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
		var claimant common.Address
		fill_err = tp.Fill(&claimant)
		if fill_err != nil {
			return
		}
		var uuid *big.Int
		fill_err = tp.Fill(&uuid)
		if fill_err != nil {
			return
		}
		var prestateMatrix types.StateSnapshot
		fill_err = tp.Fill(&prestateMatrix)
		if fill_err != nil {
			return
		}
		var preState types.Leaf
		fill_err = tp.Fill(&preState)
		if fill_err != nil {
			return
		}
		var preStateProof merkle.Proof
		fill_err = tp.Fill(&preStateProof)
		if fill_err != nil {
			return
		}
		var postState types.Leaf
		fill_err = tp.Fill(&postState)
		if fill_err != nil {
			return
		}
		var postStateProof merkle.Proof
		fill_err = tp.Fill(&postStateProof)
		if fill_err != nil {
			return
		}
		if caller == nil || uuid == nil {
			return
		}

		_x1, err := NewPreimageOracleContract(ctx, addr, caller)
		if err != nil {
			return
		}
		_x1.Squeeze(claimant, uuid, prestateMatrix, preState, preStateProof, postState, postStateProof)
	})
}
