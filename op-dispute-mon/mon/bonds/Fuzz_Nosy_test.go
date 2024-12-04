package bonds

import (
	"testing"

	"github.com/ethereum-optimism/optimism/op-dispute-mon/mon/types"
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

func Fuzz_Nosy_Bonds_CheckBonds__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *Bonds
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var games []*types.EnrichedGameData
		fill_err = tp.Fill(&games)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.CheckBonds(games)
	})
}

func Fuzz_Nosy_Bonds_checkCredits__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *Bonds
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var games []*types.EnrichedGameData
		fill_err = tp.Fill(&games)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.checkCredits(games)
	})
}

// skipping Fuzz_Nosy_BondMetrics_RecordBondCollateral__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-dispute-mon/mon/bonds.BondMetrics

// skipping Fuzz_Nosy_BondMetrics_RecordCredit__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-dispute-mon/mon/bonds.BondMetrics

// skipping Fuzz_Nosy_RClock_Now__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-dispute-mon/mon/bonds.RClock

func Fuzz_Nosy_CalculateRequiredCollateral__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var games []*types.EnrichedGameData
		fill_err = tp.Fill(&games)
		if fill_err != nil {
			return
		}

		CalculateRequiredCollateral(games)
	})
}
