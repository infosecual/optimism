package preimages

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum-optimism/optimism/op-challenger/game/keccak/matrix"
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

func Fuzz_Nosy_DirectPreimageUploader_UploadPreimage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *DirectPreimageUploader
		fill_err = tp.Fill(&d)
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
		if d == nil || d4 == nil {
			return
		}

		d.UploadPreimage(ctx, claimIdx, d4)
	})
}

func Fuzz_Nosy_LargePreimageUploader_Squeeze__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *LargePreimageUploader
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var uuid *big.Int
		fill_err = tp.Fill(&uuid)
		if fill_err != nil {
			return
		}
		var stateMatrix *matrix.StateMatrix
		fill_err = tp.Fill(&stateMatrix)
		if fill_err != nil {
			return
		}
		if p == nil || uuid == nil || stateMatrix == nil {
			return
		}

		p.Squeeze(ctx, uuid, stateMatrix)
	})
}

func Fuzz_Nosy_LargePreimageUploader_UploadPreimage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *LargePreimageUploader
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var parent uint64
		fill_err = tp.Fill(&parent)
		if fill_err != nil {
			return
		}
		var d4 *types.PreimageOracleData
		fill_err = tp.Fill(&d4)
		if fill_err != nil {
			return
		}
		if p == nil || d4 == nil {
			return
		}

		p.UploadPreimage(ctx, parent, d4)
	})
}

func Fuzz_Nosy_LargePreimageUploader_addLargePreimageData__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *LargePreimageUploader
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var uuid *big.Int
		fill_err = tp.Fill(&uuid)
		if fill_err != nil {
			return
		}
		var chunks []types.InputData
		fill_err = tp.Fill(&chunks)
		if fill_err != nil {
			return
		}
		if p == nil || uuid == nil {
			return
		}

		p.addLargePreimageData(uuid, chunks)
	})
}

func Fuzz_Nosy_LargePreimageUploader_initLargePreimage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *LargePreimageUploader
		fill_err = tp.Fill(&p)
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
		if p == nil || uuid == nil {
			return
		}

		p.initLargePreimage(uuid, partOffset, claimedSize)
	})
}

func Fuzz_Nosy_LargePreimageUploader_splitCalls__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *LargePreimageUploader
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var d2 *types.PreimageOracleData
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}
		if p == nil || d2 == nil {
			return
		}

		p.splitCalls(d2)
	})
}

func Fuzz_Nosy_SplitPreimageUploader_UploadPreimage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *SplitPreimageUploader
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var parent uint64
		fill_err = tp.Fill(&parent)
		if fill_err != nil {
			return
		}
		var d4 *types.PreimageOracleData
		fill_err = tp.Fill(&d4)
		if fill_err != nil {
			return
		}
		if s == nil || d4 == nil {
			return
		}

		s.UploadPreimage(ctx, parent, d4)
	})
}

// skipping Fuzz_Nosy_PreimageGameContract_UpdateOracleTx__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/preimages.PreimageGameContract

// skipping Fuzz_Nosy_PreimageOracleContract_AddLeaves__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/preimages.PreimageOracleContract

// skipping Fuzz_Nosy_PreimageOracleContract_CallSqueeze__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/preimages.PreimageOracleContract

// skipping Fuzz_Nosy_PreimageOracleContract_ChallengePeriod__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/preimages.PreimageOracleContract

// skipping Fuzz_Nosy_PreimageOracleContract_GetMinBondLPP__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/preimages.PreimageOracleContract

// skipping Fuzz_Nosy_PreimageOracleContract_GetProposalMetadata__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/preimages.PreimageOracleContract

// skipping Fuzz_Nosy_PreimageOracleContract_InitLargePreimage__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/preimages.PreimageOracleContract

// skipping Fuzz_Nosy_PreimageOracleContract_Squeeze__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/preimages.PreimageOracleContract

// skipping Fuzz_Nosy_PreimageUploader_UploadPreimage__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/preimages.PreimageUploader

// skipping Fuzz_Nosy_TxSender_From__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/preimages.TxSender

// skipping Fuzz_Nosy_TxSender_SendAndWaitSimple__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/preimages.TxSender
