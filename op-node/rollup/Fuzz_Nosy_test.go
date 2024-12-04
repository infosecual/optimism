package rollup

import (
	"io"
	"testing"

	"github.com/ethereum-optimism/optimism/op-service/eth"
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

func Fuzz_Nosy_ChainSpec_ChannelTimeout__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config *Config
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var t2 uint64
		fill_err = tp.Fill(&t2)
		if fill_err != nil {
			return
		}
		if config == nil {
			return
		}

		s := NewChainSpec(config)
		s.ChannelTimeout(t2)
	})
}

// skipping Fuzz_Nosy_ChainSpec_CheckForkActivation__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/log.Logger

func Fuzz_Nosy_ChainSpec_IsCanyon__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config *Config
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var t2 uint64
		fill_err = tp.Fill(&t2)
		if fill_err != nil {
			return
		}
		if config == nil {
			return
		}

		s := NewChainSpec(config)
		s.IsCanyon(t2)
	})
}

func Fuzz_Nosy_ChainSpec_IsFeatMaxSequencerDriftConstant__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config *Config
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var t2 uint64
		fill_err = tp.Fill(&t2)
		if fill_err != nil {
			return
		}
		if config == nil {
			return
		}

		s := NewChainSpec(config)
		s.IsFeatMaxSequencerDriftConstant(t2)
	})
}

func Fuzz_Nosy_ChainSpec_IsHolocene__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config *Config
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var t2 uint64
		fill_err = tp.Fill(&t2)
		if fill_err != nil {
			return
		}
		if config == nil {
			return
		}

		s := NewChainSpec(config)
		s.IsHolocene(t2)
	})
}

func Fuzz_Nosy_ChainSpec_L2ChainID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config *Config
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		if config == nil {
			return
		}

		s := NewChainSpec(config)
		s.L2ChainID()
	})
}

func Fuzz_Nosy_ChainSpec_L2GenesisTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config *Config
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		if config == nil {
			return
		}

		s := NewChainSpec(config)
		s.L2GenesisTime()
	})
}

func Fuzz_Nosy_ChainSpec_MaxChannelBankSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config *Config
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var t2 uint64
		fill_err = tp.Fill(&t2)
		if fill_err != nil {
			return
		}
		if config == nil {
			return
		}

		s := NewChainSpec(config)
		s.MaxChannelBankSize(t2)
	})
}

func Fuzz_Nosy_ChainSpec_MaxRLPBytesPerChannel__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config *Config
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var t2 uint64
		fill_err = tp.Fill(&t2)
		if fill_err != nil {
			return
		}
		if config == nil {
			return
		}

		s := NewChainSpec(config)
		s.MaxRLPBytesPerChannel(t2)
	})
}

func Fuzz_Nosy_ChainSpec_MaxSequencerDrift__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config *Config
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var t2 uint64
		fill_err = tp.Fill(&t2)
		if fill_err != nil {
			return
		}
		if config == nil {
			return
		}

		s := NewChainSpec(config)
		s.MaxSequencerDrift(t2)
	})
}

func Fuzz_Nosy_Config_ActivateAtGenesis__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainID uint64
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var hardfork ForkName
		fill_err = tp.Fill(&hardfork)
		if fill_err != nil {
			return
		}

		c, err := LoadOPStackRollupConfig(chainID)
		if err != nil {
			return
		}
		c.ActivateAtGenesis(hardfork)
	})
}

func Fuzz_Nosy_Config_AltDAEnabled__(f *testing.F) {
	f.Fuzz(func(t *testing.T, chainID uint64) {
		c, err := LoadOPStackRollupConfig(chainID)
		if err != nil {
			return
		}
		c.AltDAEnabled()
	})
}

func Fuzz_Nosy_Config_Check__(f *testing.F) {
	f.Fuzz(func(t *testing.T, chainID uint64) {
		cfg, err := LoadOPStackRollupConfig(chainID)
		if err != nil {
			return
		}
		cfg.Check()
	})
}

// skipping Fuzz_Nosy_Config_CheckL1ChainID__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup.L1Client

// skipping Fuzz_Nosy_Config_CheckL1GenesisBlockHash__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup.L1Client

// skipping Fuzz_Nosy_Config_CheckL2ChainID__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup.L2Client

// skipping Fuzz_Nosy_Config_CheckL2GenesisBlockHash__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup.L2Client

func Fuzz_Nosy_Config_Description__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainID uint64
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var l2Chains map[string]string
		fill_err = tp.Fill(&l2Chains)
		if fill_err != nil {
			return
		}

		c, err := LoadOPStackRollupConfig(chainID)
		if err != nil {
			return
		}
		c.Description(l2Chains)
	})
}

func Fuzz_Nosy_Config_ForkchoiceUpdatedVersion__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainID uint64
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var attr *eth.PayloadAttributes
		fill_err = tp.Fill(&attr)
		if fill_err != nil {
			return
		}
		if attr == nil {
			return
		}

		c, err := LoadOPStackRollupConfig(chainID)
		if err != nil {
			return
		}
		c.ForkchoiceUpdatedVersion(attr)
	})
}

func Fuzz_Nosy_Config_GetOPAltDAConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, chainID uint64) {
		c, err := LoadOPStackRollupConfig(chainID)
		if err != nil {
			return
		}
		c.GetOPAltDAConfig()
	})
}

func Fuzz_Nosy_Config_GetPayloadVersion__(f *testing.F) {
	f.Fuzz(func(t *testing.T, chainID uint64, timestamp uint64) {
		c, err := LoadOPStackRollupConfig(chainID)
		if err != nil {
			return
		}
		c.GetPayloadVersion(timestamp)
	})
}

func Fuzz_Nosy_Config_IsActivationBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, chainID uint64, oldTime uint64, newTime uint64) {
		c, err := LoadOPStackRollupConfig(chainID)
		if err != nil {
			return
		}
		c.IsActivationBlock(oldTime, newTime)
	})
}

func Fuzz_Nosy_Config_IsCanyon__(f *testing.F) {
	f.Fuzz(func(t *testing.T, chainID uint64, timestamp uint64) {
		c, err := LoadOPStackRollupConfig(chainID)
		if err != nil {
			return
		}
		c.IsCanyon(timestamp)
	})
}

func Fuzz_Nosy_Config_IsCanyonActivationBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, chainID uint64, l2BlockTime uint64) {
		c, err := LoadOPStackRollupConfig(chainID)
		if err != nil {
			return
		}
		c.IsCanyonActivationBlock(l2BlockTime)
	})
}

func Fuzz_Nosy_Config_IsDelta__(f *testing.F) {
	f.Fuzz(func(t *testing.T, chainID uint64, timestamp uint64) {
		c, err := LoadOPStackRollupConfig(chainID)
		if err != nil {
			return
		}
		c.IsDelta(timestamp)
	})
}

func Fuzz_Nosy_Config_IsDeltaActivationBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, chainID uint64, l2BlockTime uint64) {
		c, err := LoadOPStackRollupConfig(chainID)
		if err != nil {
			return
		}
		c.IsDeltaActivationBlock(l2BlockTime)
	})
}

func Fuzz_Nosy_Config_IsEcotone__(f *testing.F) {
	f.Fuzz(func(t *testing.T, chainID uint64, timestamp uint64) {
		c, err := LoadOPStackRollupConfig(chainID)
		if err != nil {
			return
		}
		c.IsEcotone(timestamp)
	})
}

func Fuzz_Nosy_Config_IsEcotoneActivationBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, chainID uint64, l2BlockTime uint64) {
		c, err := LoadOPStackRollupConfig(chainID)
		if err != nil {
			return
		}
		c.IsEcotoneActivationBlock(l2BlockTime)
	})
}

func Fuzz_Nosy_Config_IsFjord__(f *testing.F) {
	f.Fuzz(func(t *testing.T, chainID uint64, timestamp uint64) {
		c, err := LoadOPStackRollupConfig(chainID)
		if err != nil {
			return
		}
		c.IsFjord(timestamp)
	})
}

func Fuzz_Nosy_Config_IsFjordActivationBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, chainID uint64, l2BlockTime uint64) {
		c, err := LoadOPStackRollupConfig(chainID)
		if err != nil {
			return
		}
		c.IsFjordActivationBlock(l2BlockTime)
	})
}

func Fuzz_Nosy_Config_IsGranite__(f *testing.F) {
	f.Fuzz(func(t *testing.T, chainID uint64, timestamp uint64) {
		c, err := LoadOPStackRollupConfig(chainID)
		if err != nil {
			return
		}
		c.IsGranite(timestamp)
	})
}

func Fuzz_Nosy_Config_IsGraniteActivationBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, chainID uint64, l2BlockTime uint64) {
		c, err := LoadOPStackRollupConfig(chainID)
		if err != nil {
			return
		}
		c.IsGraniteActivationBlock(l2BlockTime)
	})
}

func Fuzz_Nosy_Config_IsHolocene__(f *testing.F) {
	f.Fuzz(func(t *testing.T, chainID uint64, timestamp uint64) {
		c, err := LoadOPStackRollupConfig(chainID)
		if err != nil {
			return
		}
		c.IsHolocene(timestamp)
	})
}

func Fuzz_Nosy_Config_IsHoloceneActivationBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, chainID uint64, l2BlockTime uint64) {
		c, err := LoadOPStackRollupConfig(chainID)
		if err != nil {
			return
		}
		c.IsHoloceneActivationBlock(l2BlockTime)
	})
}

func Fuzz_Nosy_Config_IsInterop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, chainID uint64, timestamp uint64) {
		c, err := LoadOPStackRollupConfig(chainID)
		if err != nil {
			return
		}
		c.IsInterop(timestamp)
	})
}

func Fuzz_Nosy_Config_IsInteropActivationBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, chainID uint64, l2BlockTime uint64) {
		c, err := LoadOPStackRollupConfig(chainID)
		if err != nil {
			return
		}
		c.IsInteropActivationBlock(l2BlockTime)
	})
}

func Fuzz_Nosy_Config_IsRegolith__(f *testing.F) {
	f.Fuzz(func(t *testing.T, chainID uint64, timestamp uint64) {
		c, err := LoadOPStackRollupConfig(chainID)
		if err != nil {
			return
		}
		c.IsRegolith(timestamp)
	})
}

func Fuzz_Nosy_Config_IsRegolithActivationBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, chainID uint64, l2BlockTime uint64) {
		c, err := LoadOPStackRollupConfig(chainID)
		if err != nil {
			return
		}
		c.IsRegolithActivationBlock(l2BlockTime)
	})
}

func Fuzz_Nosy_Config_L1Signer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, chainID uint64) {
		c, err := LoadOPStackRollupConfig(chainID)
		if err != nil {
			return
		}
		c.L1Signer()
	})
}

// skipping Fuzz_Nosy_Config_LogDescription__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/log.Logger

func Fuzz_Nosy_Config_NewPayloadVersion__(f *testing.F) {
	f.Fuzz(func(t *testing.T, chainID uint64, timestamp uint64) {
		c, err := LoadOPStackRollupConfig(chainID)
		if err != nil {
			return
		}
		c.NewPayloadVersion(timestamp)
	})
}

func Fuzz_Nosy_Config_ParseRollupConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainID uint64
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var in io.Reader
		fill_err = tp.Fill(&in)
		if fill_err != nil {
			return
		}

		c, err := LoadOPStackRollupConfig(chainID)
		if err != nil {
			return
		}
		c.ParseRollupConfig(in)
	})
}

func Fuzz_Nosy_Config_SyncLookback__(f *testing.F) {
	f.Fuzz(func(t *testing.T, chainID uint64) {
		c, err := LoadOPStackRollupConfig(chainID)
		if err != nil {
			return
		}
		c.SyncLookback()
	})
}

func Fuzz_Nosy_Config_TargetBlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, chainID uint64, timestamp uint64) {
		cfg, err := LoadOPStackRollupConfig(chainID)
		if err != nil {
			return
		}
		cfg.TargetBlockNumber(timestamp)
	})
}

func Fuzz_Nosy_Config_TimestampForBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, chainID uint64, blockNumber uint64) {
		cfg, err := LoadOPStackRollupConfig(chainID)
		if err != nil {
			return
		}
		cfg.TimestampForBlock(blockNumber)
	})
}

// skipping Fuzz_Nosy_Config_ValidateL1Config__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup.L1Client

// skipping Fuzz_Nosy_Config_ValidateL2Config__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup.L2Client

func Fuzz_Nosy_EngineTemporaryErrorEvent_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ev EngineTemporaryErrorEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}

		ev.String()
	})
}

// skipping Fuzz_Nosy_L1Client_ChainID__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup.L1Client

// skipping Fuzz_Nosy_L1Client_L1BlockRefByNumber__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup.L1Client

func Fuzz_Nosy_L1TemporaryErrorEvent_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ev L1TemporaryErrorEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}

		ev.String()
	})
}

// skipping Fuzz_Nosy_L2Client_ChainID__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup.L2Client

// skipping Fuzz_Nosy_L2Client_L2BlockRefByNumber__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup.L2Client

func Fuzz_Nosy_ResetEvent_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ev ResetEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}

		ev.String()
	})
}

// skipping Fuzz_Nosy_SafeHeadListener_Enabled__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup.SafeHeadListener

// skipping Fuzz_Nosy_SafeHeadListener_SafeHeadReset__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup.SafeHeadListener

// skipping Fuzz_Nosy_SafeHeadListener_SafeHeadUpdated__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup.SafeHeadListener

func Fuzz_Nosy_ForksFrom__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fork ForkName
		fill_err = tp.Fill(&fork)
		if fill_err != nil {
			return
		}

		ForksFrom(fork)
	})
}

func Fuzz_Nosy_IsValidFork__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fork ForkName
		fill_err = tp.Fill(&fork)
		if fill_err != nil {
			return
		}

		IsValidFork(fork)
	})
}

func Fuzz_Nosy_fmtForkTimeOrUnset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v *uint64
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		if v == nil {
			return
		}

		fmtForkTimeOrUnset(v)
	})
}

func Fuzz_Nosy_fmtTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, v uint64) {
		fmtTime(v)
	})
}
