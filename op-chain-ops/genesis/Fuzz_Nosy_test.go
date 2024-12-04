package genesis

import (
	"testing"

	"github.com/ethereum-optimism/optimism/op-node/rollup"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
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

// skipping Fuzz_Nosy_AltDADeployConfig_Check__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/log.Logger

// skipping Fuzz_Nosy_DeployConfig_Check__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/log.Logger

func Fuzz_Nosy_DeployConfig_CheckAddresses__(f *testing.F) {
	f.Fuzz(func(t *testing.T, network string, path string) {
		d, err := NewDeployConfigWithNetwork(network, path)
		if err != nil {
			return
		}
		d.CheckAddresses()
	})
}

func Fuzz_Nosy_DeployConfig_Copy__(f *testing.F) {
	f.Fuzz(func(t *testing.T, network string, path string) {
		d, err := NewDeployConfigWithNetwork(network, path)
		if err != nil {
			return
		}
		d.Copy()
	})
}

func Fuzz_Nosy_DeployConfig_RollupConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var network string
		fill_err = tp.Fill(&network)
		if fill_err != nil {
			return
		}
		var path string
		fill_err = tp.Fill(&path)
		if fill_err != nil {
			return
		}
		var l1StartBlock *types.Header
		fill_err = tp.Fill(&l1StartBlock)
		if fill_err != nil {
			return
		}
		var l2GenesisBlockHash common.Hash
		fill_err = tp.Fill(&l2GenesisBlockHash)
		if fill_err != nil {
			return
		}
		var l2GenesisBlockNumber uint64
		fill_err = tp.Fill(&l2GenesisBlockNumber)
		if fill_err != nil {
			return
		}
		if l1StartBlock == nil {
			return
		}

		d, err := NewDeployConfigWithNetwork(network, path)
		if err != nil {
			return
		}
		d.RollupConfig(l1StartBlock, l2GenesisBlockHash, l2GenesisBlockNumber)
	})
}

func Fuzz_Nosy_DeployConfig_SetDeployments__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var network string
		fill_err = tp.Fill(&network)
		if fill_err != nil {
			return
		}
		var path string
		fill_err = tp.Fill(&path)
		if fill_err != nil {
			return
		}
		var deployments *L1Deployments
		fill_err = tp.Fill(&deployments)
		if fill_err != nil {
			return
		}
		if deployments == nil {
			return
		}

		d, err := NewDeployConfigWithNetwork(network, path)
		if err != nil {
			return
		}
		d.SetDeployments(deployments)
	})
}

// skipping Fuzz_Nosy_EIP1559DeployConfig_Check__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/log.Logger

// skipping Fuzz_Nosy_FaultProofDeployConfig_Check__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/log.Logger

// skipping Fuzz_Nosy_GasPriceOracleDeployConfig_Check__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/log.Logger

func Fuzz_Nosy_GasPriceOracleDeployConfig_FeeScalar__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *GasPriceOracleDeployConfig
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.FeeScalar()
	})
}

// skipping Fuzz_Nosy_GasTokenDeployConfig_Check__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/log.Logger

// skipping Fuzz_Nosy_GovernanceDeployConfig_Check__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/log.Logger

func Fuzz_Nosy_GovernanceDeployConfig_GovernanceEnabled__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *GovernanceDeployConfig
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.GovernanceEnabled()
	})
}

func Fuzz_Nosy_L1DependenciesConfig_CheckAddresses__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *L1DependenciesConfig
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var dependencyContext DependencyContext
		fill_err = tp.Fill(&dependencyContext)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.CheckAddresses(dependencyContext)
	})
}

func Fuzz_Nosy_L1Deployments_Check__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var path string
		fill_err = tp.Fill(&path)
		if fill_err != nil {
			return
		}
		var deployConfig *DeployConfig
		fill_err = tp.Fill(&deployConfig)
		if fill_err != nil {
			return
		}
		if deployConfig == nil {
			return
		}

		d, err := NewL1Deployments(path)
		if err != nil {
			return
		}
		d.Check(deployConfig)
	})
}

func Fuzz_Nosy_L1Deployments_Copy__(f *testing.F) {
	f.Fuzz(func(t *testing.T, path string) {
		d, err := NewL1Deployments(path)
		if err != nil {
			return
		}
		d.Copy()
	})
}

// skipping Fuzz_Nosy_L1Deployments_ForEach__ because parameters include func, chan, or unsupported interface: func(name string, addr github.com/ethereum/go-ethereum/common.Address)

func Fuzz_Nosy_L1Deployments_GetName__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var path string
		fill_err = tp.Fill(&path)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}

		d, err := NewL1Deployments(path)
		if err != nil {
			return
		}
		d.GetName(addr)
	})
}

// skipping Fuzz_Nosy_L2CoreDeployConfig_Check__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/log.Logger

// skipping Fuzz_Nosy_L2GenesisBlockDeployConfig_Check__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/log.Logger

// skipping Fuzz_Nosy_L2InitializationConfig_Check__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/log.Logger

// skipping Fuzz_Nosy_L2VaultsDeployConfig_Check__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/log.Logger

// skipping Fuzz_Nosy_LegacyDeployConfig_Check__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/log.Logger

func Fuzz_Nosy_MarshalableRPCBlockNumberOrHash_Hash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MarshalableRPCBlockNumberOrHash
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Hash()
	})
}

func Fuzz_Nosy_MarshalableRPCBlockNumberOrHash_MarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MarshalableRPCBlockNumberOrHash
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalJSON()
	})
}

func Fuzz_Nosy_MarshalableRPCBlockNumberOrHash_Number__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MarshalableRPCBlockNumberOrHash
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Number()
	})
}

func Fuzz_Nosy_MarshalableRPCBlockNumberOrHash_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MarshalableRPCBlockNumberOrHash
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.String()
	})
}

func Fuzz_Nosy_MarshalableRPCBlockNumberOrHash_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MarshalableRPCBlockNumberOrHash
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.UnmarshalJSON(b)
	})
}

// skipping Fuzz_Nosy_OperatorDeployConfig_Check__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/log.Logger

// skipping Fuzz_Nosy_OutputOracleDeployConfig_Check__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/log.Logger

// skipping Fuzz_Nosy_OwnershipDeployConfig_Check__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/log.Logger

// skipping Fuzz_Nosy_SuperchainL1DeployConfig_Check__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/log.Logger

func Fuzz_Nosy_UpgradeScheduleDeployConfig_ActivateForkAtGenesis__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *UpgradeScheduleDeployConfig
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var fork rollup.ForkName
		fill_err = tp.Fill(&fork)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.ActivateForkAtGenesis(fork)
	})
}

func Fuzz_Nosy_UpgradeScheduleDeployConfig_ActivateForkAtOffset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *UpgradeScheduleDeployConfig
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var fork rollup.ForkName
		fill_err = tp.Fill(&fork)
		if fill_err != nil {
			return
		}
		var offset uint64
		fill_err = tp.Fill(&offset)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.ActivateForkAtOffset(fork, offset)
	})
}

func Fuzz_Nosy_UpgradeScheduleDeployConfig_AllocMode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *UpgradeScheduleDeployConfig
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var genesisTime uint64
		fill_err = tp.Fill(&genesisTime)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.AllocMode(genesisTime)
	})
}

func Fuzz_Nosy_UpgradeScheduleDeployConfig_CanyonTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *UpgradeScheduleDeployConfig
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var genesisTime uint64
		fill_err = tp.Fill(&genesisTime)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.CanyonTime(genesisTime)
	})
}

// skipping Fuzz_Nosy_UpgradeScheduleDeployConfig_Check__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/log.Logger

func Fuzz_Nosy_UpgradeScheduleDeployConfig_DeltaTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *UpgradeScheduleDeployConfig
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var genesisTime uint64
		fill_err = tp.Fill(&genesisTime)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.DeltaTime(genesisTime)
	})
}

func Fuzz_Nosy_UpgradeScheduleDeployConfig_EcotoneTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *UpgradeScheduleDeployConfig
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var genesisTime uint64
		fill_err = tp.Fill(&genesisTime)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.EcotoneTime(genesisTime)
	})
}

func Fuzz_Nosy_UpgradeScheduleDeployConfig_FjordTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *UpgradeScheduleDeployConfig
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var genesisTime uint64
		fill_err = tp.Fill(&genesisTime)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.FjordTime(genesisTime)
	})
}

func Fuzz_Nosy_UpgradeScheduleDeployConfig_ForkTimeOffset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *UpgradeScheduleDeployConfig
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var fork rollup.ForkName
		fill_err = tp.Fill(&fork)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.ForkTimeOffset(fork)
	})
}

func Fuzz_Nosy_UpgradeScheduleDeployConfig_GraniteTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *UpgradeScheduleDeployConfig
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var genesisTime uint64
		fill_err = tp.Fill(&genesisTime)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.GraniteTime(genesisTime)
	})
}

func Fuzz_Nosy_UpgradeScheduleDeployConfig_HoloceneTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *UpgradeScheduleDeployConfig
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var genesisTime uint64
		fill_err = tp.Fill(&genesisTime)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.HoloceneTime(genesisTime)
	})
}

func Fuzz_Nosy_UpgradeScheduleDeployConfig_InteropTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *UpgradeScheduleDeployConfig
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var genesisTime uint64
		fill_err = tp.Fill(&genesisTime)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.InteropTime(genesisTime)
	})
}

func Fuzz_Nosy_UpgradeScheduleDeployConfig_RegolithTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *UpgradeScheduleDeployConfig
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var genesisTime uint64
		fill_err = tp.Fill(&genesisTime)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.RegolithTime(genesisTime)
	})
}

func Fuzz_Nosy_UpgradeScheduleDeployConfig_SetForkTimeOffset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *UpgradeScheduleDeployConfig
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var fork rollup.ForkName
		fill_err = tp.Fill(&fork)
		if fill_err != nil {
			return
		}
		var offset *uint64
		fill_err = tp.Fill(&offset)
		if fill_err != nil {
			return
		}
		if d == nil || offset == nil {
			return
		}

		d.SetForkTimeOffset(fork, offset)
	})
}

func Fuzz_Nosy_UpgradeScheduleDeployConfig_forks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *UpgradeScheduleDeployConfig
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.forks()
	})
}

func Fuzz_Nosy_WithdrawalNetwork_ToUint8__(f *testing.F) {
	f.Fuzz(func(t *testing.T, i uint8) {
		w := FromUint8(i)
		w.ToUint8()
	})
}

func Fuzz_Nosy_WithdrawalNetwork_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, i uint8, b []byte) {
		w := FromUint8(i)
		w.UnmarshalJSON(b)
	})
}

func Fuzz_Nosy_WithdrawalNetwork_Valid__(f *testing.F) {
	f.Fuzz(func(t *testing.T, i uint8) {
		w := FromUint8(i)
		w.Valid()
	})
}

// skipping Fuzz_Nosy_ConfigChecker_Check__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-chain-ops/genesis.ConfigChecker

func Fuzz_Nosy_WithdrawalNetwork_MarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, i uint8) {
		w := FromUint8(i)
		w.MarshalJSON()
	})
}

func Fuzz_Nosy_WithdrawalNetwork_ToABI__(f *testing.F) {
	f.Fuzz(func(t *testing.T, i uint8) {
		w := FromUint8(i)
		w.ToABI()
	})
}

func Fuzz_Nosy_FundDevAccounts__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var gen *core.Genesis
		fill_err = tp.Fill(&gen)
		if fill_err != nil {
			return
		}
		if gen == nil {
			return
		}

		FundDevAccounts(gen)
	})
}

func Fuzz_Nosy_HasAnyDevAccounts__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var allocs types.GenesisAlloc
		fill_err = tp.Fill(&allocs)
		if fill_err != nil {
			return
		}

		HasAnyDevAccounts(allocs)
	})
}

func Fuzz_Nosy_SetPrecompileBalances__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var gen *core.Genesis
		fill_err = tp.Fill(&gen)
		if fill_err != nil {
			return
		}
		if gen == nil {
			return
		}

		SetPrecompileBalances(gen)
	})
}

func Fuzz_Nosy_offsetToUpgradeTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var offset *hexutil.Uint64
		fill_err = tp.Fill(&offset)
		if fill_err != nil {
			return
		}
		var genesisTime uint64
		fill_err = tp.Fill(&genesisTime)
		if fill_err != nil {
			return
		}
		if offset == nil {
			return
		}

		offsetToUpgradeTime(offset, genesisTime)
	})
}

func Fuzz_Nosy_u64ptr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, n uint64) {
		u64ptr(n)
	})
}
