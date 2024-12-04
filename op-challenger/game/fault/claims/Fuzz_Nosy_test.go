package claims

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

func Fuzz_Nosy_BondClaimScheduler_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *BondClaimScheduler
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

func Fuzz_Nosy_BondClaimScheduler_Schedule__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *BondClaimScheduler
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var blockNumber uint64
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		var games []types.GameMetadata
		fill_err = tp.Fill(&games)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Schedule(blockNumber, games)
	})
}

func Fuzz_Nosy_BondClaimScheduler_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *BondClaimScheduler
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Start(ctx)
	})
}

func Fuzz_Nosy_BondClaimScheduler_run__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *BondClaimScheduler
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.run(ctx)
	})
}

func Fuzz_Nosy_Claimer_ClaimBonds__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *Claimer
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var games []types.GameMetadata
		fill_err = tp.Fill(&games)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.ClaimBonds(ctx, games)
	})
}

func Fuzz_Nosy_Claimer_claimBond__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *Claimer
		fill_err = tp.Fill(&c)
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
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.claimBond(ctx, game, addr)
	})
}

// skipping Fuzz_Nosy_BondClaimMetrics_RecordBondClaimed__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/claims.BondClaimMetrics

// skipping Fuzz_Nosy_BondClaimSchedulerMetrics_RecordBondClaimFailed__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/claims.BondClaimSchedulerMetrics

// skipping Fuzz_Nosy_BondClaimer_ClaimBonds__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/claims.BondClaimer

// skipping Fuzz_Nosy_BondContract_ClaimCreditTx__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/claims.BondContract

// skipping Fuzz_Nosy_BondContract_GetCredit__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/claims.BondContract

// skipping Fuzz_Nosy_TxSender_SendAndWaitSimple__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/claims.TxSender
