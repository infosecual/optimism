package keccak

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

func Fuzz_Nosy_LargePreimageScheduler_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *LargePreimageScheduler
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

func Fuzz_Nosy_LargePreimageScheduler_Schedule__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *LargePreimageScheduler
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var blockHash common.Hash
		fill_err = tp.Fill(&blockHash)
		if fill_err != nil {
			return
		}
		var _x3 uint64
		fill_err = tp.Fill(&_x3)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Schedule(blockHash, _x3)
	})
}

func Fuzz_Nosy_LargePreimageScheduler_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *LargePreimageScheduler
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

func Fuzz_Nosy_LargePreimageScheduler_run__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *LargePreimageScheduler
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

// skipping Fuzz_Nosy_LargePreimageScheduler_verifyOraclePreimages__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/keccak/types.LargePreimageOracle

func Fuzz_Nosy_LargePreimageScheduler_verifyPreimages__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *LargePreimageScheduler
		fill_err = tp.Fill(&s)
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
		if s == nil {
			return
		}

		s.verifyPreimages(ctx, blockHash)
	})
}

// skipping Fuzz_Nosy_PreimageChallenger_Challenge__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/keccak.Oracle

// skipping Fuzz_Nosy_PreimageVerifier_CreateChallenge__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/keccak.VerifierPreimageOracle

// skipping Fuzz_Nosy_ChallengeMetrics_RecordPreimageChallengeFailed__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/keccak.ChallengeMetrics

// skipping Fuzz_Nosy_ChallengeMetrics_RecordPreimageChallenged__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/keccak.ChallengeMetrics

// skipping Fuzz_Nosy_Challenger_Challenge__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/keccak.Challenger

// skipping Fuzz_Nosy_Fetcher_FetchInputs__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/keccak.Fetcher

// skipping Fuzz_Nosy_Metrics_RecordLargePreimageCount__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/keccak.Metrics

// skipping Fuzz_Nosy_Oracle_ChallengeTx__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/keccak.Oracle

// skipping Fuzz_Nosy_OracleSource_Oracles__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/keccak.OracleSource

// skipping Fuzz_Nosy_Sender_SendAndWaitSimple__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/keccak.Sender

// skipping Fuzz_Nosy_Verifier_CreateChallenge__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/keccak.Verifier

// skipping Fuzz_Nosy_VerifierPreimageOracle_GetProposalTreeRoot__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/keccak.VerifierPreimageOracle
