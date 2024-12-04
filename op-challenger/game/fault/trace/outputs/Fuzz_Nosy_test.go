package outputs

import (
	"context"
	"testing"

	"github.com/ethereum-optimism/optimism/op-challenger/game/fault/contracts"
	"github.com/ethereum-optimism/optimism/op-challenger/game/fault/types"
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

func Fuzz_Nosy_OutputPrestateProvider_AbsolutePreStateCommitment__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o *OutputPrestateProvider
		fill_err = tp.Fill(&o)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if o == nil {
			return
		}

		o.AbsolutePreStateCommitment(ctx)
	})
}

func Fuzz_Nosy_OutputPrestateProvider_outputAtBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o *OutputPrestateProvider
		fill_err = tp.Fill(&o)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var block uint64
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		if o == nil {
			return
		}

		o.outputAtBlock(ctx, block)
	})
}

func Fuzz_Nosy_OutputTraceProvider_ClaimedBlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o *OutputTraceProvider
		fill_err = tp.Fill(&o)
		if fill_err != nil {
			return
		}
		var pos types.Position
		fill_err = tp.Fill(&pos)
		if fill_err != nil {
			return
		}
		if o == nil {
			return
		}

		o.ClaimedBlockNumber(pos)
	})
}

func Fuzz_Nosy_OutputTraceProvider_Get__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o *OutputTraceProvider
		fill_err = tp.Fill(&o)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var pos types.Position
		fill_err = tp.Fill(&pos)
		if fill_err != nil {
			return
		}
		if o == nil {
			return
		}

		o.Get(ctx, pos)
	})
}

func Fuzz_Nosy_OutputTraceProvider_GetL2BlockNumberChallenge__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o *OutputTraceProvider
		fill_err = tp.Fill(&o)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if o == nil {
			return
		}

		o.GetL2BlockNumberChallenge(ctx)
	})
}

func Fuzz_Nosy_OutputTraceProvider_GetStepData__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o *OutputTraceProvider
		fill_err = tp.Fill(&o)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var _x3 types.Position
		fill_err = tp.Fill(&_x3)
		if fill_err != nil {
			return
		}
		if o == nil {
			return
		}

		o.GetStepData(_x2, _x3)
	})
}

func Fuzz_Nosy_OutputTraceProvider_HonestBlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o *OutputTraceProvider
		fill_err = tp.Fill(&o)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var pos types.Position
		fill_err = tp.Fill(&pos)
		if fill_err != nil {
			return
		}
		if o == nil {
			return
		}

		o.HonestBlockNumber(ctx, pos)
	})
}

func Fuzz_Nosy_OutputTraceProvider_outputAtBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o *OutputTraceProvider
		fill_err = tp.Fill(&o)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var block uint64
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		if o == nil {
			return
		}

		o.outputAtBlock(ctx, block)
	})
}

func Fuzz_Nosy_ProviderCache_GetOrCreate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *ProviderCache
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var localContext common.Hash
		fill_err = tp.Fill(&localContext)
		if fill_err != nil {
			return
		}
		var depth types.Depth
		fill_err = tp.Fill(&depth)
		if fill_err != nil {
			return
		}
		var agreed contracts.Proposal
		fill_err = tp.Fill(&agreed)
		if fill_err != nil {
			return
		}
		var claimed contracts.Proposal
		fill_err = tp.Fill(&claimed)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.GetOrCreate(ctx, localContext, depth, agreed, claimed)
	})
}

// skipping Fuzz_Nosy_OutputRollupClient_OutputAtBlock__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/trace/outputs.OutputRollupClient

// skipping Fuzz_Nosy_OutputRollupClient_SafeHeadAtL1Block__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/trace/outputs.OutputRollupClient

func Fuzz_Nosy_FetchProposals__(f *testing.F) {
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
		var topProvider *OutputTraceProvider
		fill_err = tp.Fill(&topProvider)
		if fill_err != nil {
			return
		}
		var pre types.Claim
		fill_err = tp.Fill(&pre)
		if fill_err != nil {
			return
		}
		var post types.Claim
		fill_err = tp.Fill(&post)
		if fill_err != nil {
			return
		}
		if topProvider == nil {
			return
		}

		FetchProposals(ctx, topProvider, pre, post)
	})
}

func Fuzz_Nosy_localContextPreimage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pre types.Claim
		fill_err = tp.Fill(&pre)
		if fill_err != nil {
			return
		}
		var post types.Claim
		fill_err = tp.Fill(&post)
		if fill_err != nil {
			return
		}

		localContextPreimage(pre, post)
	})
}
