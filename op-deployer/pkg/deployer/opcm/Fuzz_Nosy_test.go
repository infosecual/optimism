package opcm

import (
	"context"
	"testing"

	"github.com/ethereum-optimism/optimism/op-chain-ops/script"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
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

func Fuzz_Nosy_Contract_GenericAddressGetter__(f *testing.F) {
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
		var client *ethclient.Client
		fill_err = tp.Fill(&client)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var functionName string
		fill_err = tp.Fill(&functionName)
		if fill_err != nil {
			return
		}
		if client == nil {
			return
		}

		c := NewContract(addr, client)
		c.GenericAddressGetter(ctx, functionName)
	})
}

func Fuzz_Nosy_Contract_GetAddressByNameViaAddressManager__(f *testing.F) {
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
		var client *ethclient.Client
		fill_err = tp.Fill(&client)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		if client == nil {
			return
		}

		c := NewContract(addr, client)
		c.GetAddressByNameViaAddressManager(ctx, name)
	})
}

func Fuzz_Nosy_Contract_ProtocolVersions__(f *testing.F) {
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
		var client *ethclient.Client
		fill_err = tp.Fill(&client)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if client == nil {
			return
		}

		c := NewContract(addr, client)
		c.ProtocolVersions(ctx)
	})
}

func Fuzz_Nosy_Contract_SuperchainConfig__(f *testing.F) {
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
		var client *ethclient.Client
		fill_err = tp.Fill(&client)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if client == nil {
			return
		}

		c := NewContract(addr, client)
		c.SuperchainConfig(ctx)
	})
}

// skipping Fuzz_Nosy_Contract_callContractMethod__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_Contract_getAddress__(f *testing.F) {
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
		var client *ethclient.Client
		fill_err = tp.Fill(&client)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		if client == nil {
			return
		}

		c := NewContract(addr, client)
		c.getAddress(ctx, name)
	})
}

func Fuzz_Nosy_DeployDelayedWETHInput_InputSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var input *DeployDelayedWETHInput
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		if input == nil {
			return
		}

		input.InputSet()
	})
}

func Fuzz_Nosy_DeployDelayedWETHOutput_CheckOutput__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var host *script.Host
		fill_err = tp.Fill(&host)
		if fill_err != nil {
			return
		}
		var i2 DeployDelayedWETHInput
		fill_err = tp.Fill(&i2)
		if fill_err != nil {
			return
		}
		var i3 common.Address
		fill_err = tp.Fill(&i3)
		if fill_err != nil {
			return
		}
		if host == nil {
			return
		}

		output, err := DeployDelayedWETH(host, i2)
		if err != nil {
			return
		}
		output.CheckOutput(i3)
	})
}

func Fuzz_Nosy_DeployDisputeGameInput_InputSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var input *DeployDisputeGameInput
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		if input == nil {
			return
		}

		input.InputSet()
	})
}

func Fuzz_Nosy_DeployDisputeGameOutput_CheckOutput__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var host *script.Host
		fill_err = tp.Fill(&host)
		if fill_err != nil {
			return
		}
		var i2 DeployDisputeGameInput
		fill_err = tp.Fill(&i2)
		if fill_err != nil {
			return
		}
		var i3 common.Address
		fill_err = tp.Fill(&i3)
		if fill_err != nil {
			return
		}
		if host == nil {
			return
		}

		output, err := DeployDisputeGame(host, i2)
		if err != nil {
			return
		}
		output.CheckOutput(i3)
	})
}

func Fuzz_Nosy_DeployImplementationsInput_InputSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var input *DeployImplementationsInput
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		if input == nil {
			return
		}

		input.InputSet()
	})
}

func Fuzz_Nosy_DeployImplementationsOutput_CheckOutput__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var host *script.Host
		fill_err = tp.Fill(&host)
		if fill_err != nil {
			return
		}
		var i2 DeployImplementationsInput
		fill_err = tp.Fill(&i2)
		if fill_err != nil {
			return
		}
		var i3 common.Address
		fill_err = tp.Fill(&i3)
		if fill_err != nil {
			return
		}
		if host == nil {
			return
		}

		output, err := DeployImplementations(host, i2)
		if err != nil {
			return
		}
		output.CheckOutput(i3)
	})
}

func Fuzz_Nosy_DeployMIPSInput_InputSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var input *DeployMIPSInput
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		if input == nil {
			return
		}

		input.InputSet()
	})
}

func Fuzz_Nosy_DeployMIPSOutput_CheckOutput__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var host *script.Host
		fill_err = tp.Fill(&host)
		if fill_err != nil {
			return
		}
		var i2 DeployMIPSInput
		fill_err = tp.Fill(&i2)
		if fill_err != nil {
			return
		}
		var i3 common.Address
		fill_err = tp.Fill(&i3)
		if fill_err != nil {
			return
		}
		if host == nil {
			return
		}

		output, err := DeployMIPS(host, i2)
		if err != nil {
			return
		}
		output.CheckOutput(i3)
	})
}

func Fuzz_Nosy_DeployOPChainInputV160_InputSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var input *DeployOPChainInputV160
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		if input == nil {
			return
		}

		input.InputSet()
	})
}

func Fuzz_Nosy_DeployOPChainInputV160_StartingAnchorRoots__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var input *DeployOPChainInputV160
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		if input == nil {
			return
		}

		input.StartingAnchorRoots()
	})
}

func Fuzz_Nosy_DeployOPChainOutput_CheckOutput__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var host *script.Host
		fill_err = tp.Fill(&host)
		if fill_err != nil {
			return
		}
		var i2 DeployOPChainInputIsthmus
		fill_err = tp.Fill(&i2)
		if fill_err != nil {
			return
		}
		var i3 common.Address
		fill_err = tp.Fill(&i3)
		if fill_err != nil {
			return
		}
		if host == nil {
			return
		}

		output, err := DeployOPChainIsthmus(host, i2)
		if err != nil {
			return
		}
		output.CheckOutput(i3)
	})
}

func Fuzz_Nosy_DeploySuperchainInput_InputSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dsi *DeploySuperchainInput
		fill_err = tp.Fill(&dsi)
		if fill_err != nil {
			return
		}
		if dsi == nil {
			return
		}

		dsi.InputSet()
	})
}

func Fuzz_Nosy_DeploySuperchainOutput_CheckOutput__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *script.Host
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var i2 DeploySuperchainInput
		fill_err = tp.Fill(&i2)
		if fill_err != nil {
			return
		}
		var i3 common.Address
		fill_err = tp.Fill(&i3)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		output, err := DeploySuperchain(h, i2)
		if err != nil {
			return
		}
		output.CheckOutput(i3)
	})
}
