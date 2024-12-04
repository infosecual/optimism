package extract

import (
	"context"
	"testing"

	"github.com/ethereum-optimism/optimism/op-challenger/game/types"
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

// skipping Fuzz_Nosy_AgreementEnricher_Enrich__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-dispute-mon/mon/extract.GameCaller

// skipping Fuzz_Nosy_BalanceEnricher_Enrich__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-dispute-mon/mon/extract.GameCaller

// skipping Fuzz_Nosy_BondEnricher_Enrich__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-dispute-mon/mon/extract.GameCaller

// skipping Fuzz_Nosy_ClaimEnricher_Enrich__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-dispute-mon/mon/extract.GameCaller

func Fuzz_Nosy_Extractor_Extract__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *Extractor
		fill_err = tp.Fill(&e)
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
		var minTimestamp uint64
		fill_err = tp.Fill(&minTimestamp)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.Extract(ctx, blockHash, minTimestamp)
	})
}

// skipping Fuzz_Nosy_Extractor_applyEnrichers__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-dispute-mon/mon/extract.GameCaller

func Fuzz_Nosy_Extractor_enrichGame__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *Extractor
		fill_err = tp.Fill(&e)
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
		var game types.GameMetadata
		fill_err = tp.Fill(&game)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.enrichGame(ctx, blockHash, game)
	})
}

func Fuzz_Nosy_Extractor_enrichGames__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *Extractor
		fill_err = tp.Fill(&e)
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
		var games []types.GameMetadata
		fill_err = tp.Fill(&games)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.enrichGames(ctx, blockHash, games)
	})
}

func Fuzz_Nosy_GameCallerCreator_CreateContract__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *GameCallerCreator
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var game types.GameMetadata
		fill_err = tp.Fill(&game)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.CreateContract(ctx, game)
	})
}

// skipping Fuzz_Nosy_L1HeadBlockNumEnricher_Enrich__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-dispute-mon/mon/extract.GameCaller

// skipping Fuzz_Nosy_RecipientEnricher_Enrich__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-dispute-mon/mon/extract.GameCaller

// skipping Fuzz_Nosy_WithdrawalsEnricher_Enrich__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-dispute-mon/mon/extract.GameCaller

// skipping Fuzz_Nosy_BalanceCaller_GetBalanceAndDelay__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-dispute-mon/mon/extract.BalanceCaller

// skipping Fuzz_Nosy_BlockFetcher_HeaderByHash__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-dispute-mon/mon/extract.BlockFetcher

// skipping Fuzz_Nosy_BondCaller_GetCredits__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-dispute-mon/mon/extract.BondCaller

// skipping Fuzz_Nosy_ClaimCaller_IsResolved__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-dispute-mon/mon/extract.ClaimCaller

// skipping Fuzz_Nosy_Enricher_Enrich__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-dispute-mon/mon/extract.Enricher

// skipping Fuzz_Nosy_GameCaller_GetAllClaims__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-dispute-mon/mon/extract.GameCaller

// skipping Fuzz_Nosy_GameCaller_GetGameMetadata__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-dispute-mon/mon/extract.GameCaller

// skipping Fuzz_Nosy_GameCaller_GetWithdrawals__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-dispute-mon/mon/extract.GameCaller

// skipping Fuzz_Nosy_OutputMetrics_RecordOutputFetchTime__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-dispute-mon/mon/extract.OutputMetrics

// skipping Fuzz_Nosy_OutputRollupClient_OutputAtBlock__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-dispute-mon/mon/extract.OutputRollupClient

// skipping Fuzz_Nosy_OutputRollupClient_SafeHeadAtL1Block__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-dispute-mon/mon/extract.OutputRollupClient
