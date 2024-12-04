package config

import (
	"testing"

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

func Fuzz_Nosy_Config_Check__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var gameFactoryAddress common.Address
		fill_err = tp.Fill(&gameFactoryAddress)
		if fill_err != nil {
			return
		}
		var l1EthRpc string
		fill_err = tp.Fill(&l1EthRpc)
		if fill_err != nil {
			return
		}
		var l1BeaconApi string
		fill_err = tp.Fill(&l1BeaconApi)
		if fill_err != nil {
			return
		}
		var l2RollupRpc string
		fill_err = tp.Fill(&l2RollupRpc)
		if fill_err != nil {
			return
		}
		var l2EthRpc string
		fill_err = tp.Fill(&l2EthRpc)
		if fill_err != nil {
			return
		}
		var datadir string
		fill_err = tp.Fill(&datadir)
		if fill_err != nil {
			return
		}
		var supportedTraceTypes []types.TraceType
		fill_err = tp.Fill(&supportedTraceTypes)
		if fill_err != nil {
			return
		}

		c := NewConfig(gameFactoryAddress, l1EthRpc, l1BeaconApi, l2RollupRpc, l2EthRpc, datadir, supportedTraceTypes...)
		c.Check()
	})
}

func Fuzz_Nosy_Config_TraceTypeEnabled__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var gameFactoryAddress common.Address
		fill_err = tp.Fill(&gameFactoryAddress)
		if fill_err != nil {
			return
		}
		var l1EthRpc string
		fill_err = tp.Fill(&l1EthRpc)
		if fill_err != nil {
			return
		}
		var l1BeaconApi string
		fill_err = tp.Fill(&l1BeaconApi)
		if fill_err != nil {
			return
		}
		var l2RollupRpc string
		fill_err = tp.Fill(&l2RollupRpc)
		if fill_err != nil {
			return
		}
		var l2EthRpc string
		fill_err = tp.Fill(&l2EthRpc)
		if fill_err != nil {
			return
		}
		var datadir string
		fill_err = tp.Fill(&datadir)
		if fill_err != nil {
			return
		}
		var supportedTraceTypes []types.TraceType
		fill_err = tp.Fill(&supportedTraceTypes)
		if fill_err != nil {
			return
		}
		var t8 types.TraceType
		fill_err = tp.Fill(&t8)
		if fill_err != nil {
			return
		}

		c := NewConfig(gameFactoryAddress, l1EthRpc, l1BeaconApi, l2RollupRpc, l2EthRpc, datadir, supportedTraceTypes...)
		c.TraceTypeEnabled(t8)
	})
}
