package registry

import (
	"testing"

	"github.com/ethereum-optimism/optimism/op-challenger/game/types"
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

func Fuzz_Nosy_GameTypeRegistry_CreateBondContract__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var game types.GameMetadata
		fill_err = tp.Fill(&game)
		if fill_err != nil {
			return
		}

		r := NewGameTypeRegistry()
		r.CreateBondContract(game)
	})
}

func Fuzz_Nosy_GameTypeRegistry_CreatePlayer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var game types.GameMetadata
		fill_err = tp.Fill(&game)
		if fill_err != nil {
			return
		}
		var dir string
		fill_err = tp.Fill(&dir)
		if fill_err != nil {
			return
		}

		r := NewGameTypeRegistry()
		r.CreatePlayer(game, dir)
	})
}

// skipping Fuzz_Nosy_GameTypeRegistry_RegisterBondContract__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/fault/claims.BondContractCreator

// skipping Fuzz_Nosy_GameTypeRegistry_RegisterGameType__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/scheduler.PlayerCreator

// skipping Fuzz_Nosy_OracleRegistry_RegisterOracle__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/keccak/types.LargePreimageOracle
