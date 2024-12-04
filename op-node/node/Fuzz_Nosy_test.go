package node

import (
	"context"
	"testing"
	"time"

	"github.com/ethereum-optimism/optimism/op-node/p2p"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/params"
	"github.com/libp2p/go-libp2p/core/peer"
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

func Fuzz_Nosy_ActiveConfigPersistence_SequencerStarted__(f *testing.F) {
	f.Fuzz(func(t *testing.T, file string) {
		p := NewConfigPersistence(file)
		p.SequencerStarted()
	})
}

func Fuzz_Nosy_ActiveConfigPersistence_SequencerState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, file string) {
		p := NewConfigPersistence(file)
		p.SequencerState()
	})
}

func Fuzz_Nosy_ActiveConfigPersistence_SequencerStopped__(f *testing.F) {
	f.Fuzz(func(t *testing.T, file string) {
		p := NewConfigPersistence(file)
		p.SequencerStopped()
	})
}

func Fuzz_Nosy_ActiveConfigPersistence_persist__(f *testing.F) {
	f.Fuzz(func(t *testing.T, file string, sequencerStarted bool) {
		p := NewConfigPersistence(file)
		p.persist(sequencerStarted)
	})
}

func Fuzz_Nosy_ActiveConfigPersistence_read__(f *testing.F) {
	f.Fuzz(func(t *testing.T, file string) {
		p := NewConfigPersistence(file)
		p.read()
	})
}

func Fuzz_Nosy_ConductorClient_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *ConductorClient
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.Close()
	})
}

func Fuzz_Nosy_ConductorClient_CommitUnsafePayload__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *ConductorClient
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var payload *eth.ExecutionPayloadEnvelope
		fill_err = tp.Fill(&payload)
		if fill_err != nil {
			return
		}
		if c == nil || payload == nil {
			return
		}

		c.CommitUnsafePayload(ctx, payload)
	})
}

func Fuzz_Nosy_ConductorClient_Enabled__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *ConductorClient
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.Enabled(ctx)
	})
}

func Fuzz_Nosy_ConductorClient_Leader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *ConductorClient
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.Leader(ctx)
	})
}

func Fuzz_Nosy_ConductorClient_OverrideLeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *ConductorClient
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.OverrideLeader(ctx)
	})
}

func Fuzz_Nosy_ConductorClient_initialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *ConductorClient
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.initialize(ctx)
	})
}

func Fuzz_Nosy_Config_Check__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		cfg.Check()
	})
}

// skipping Fuzz_Nosy_Config_LoadPersisted__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/log.Logger

func Fuzz_Nosy_Config_P2PEnabled__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		cfg.P2PEnabled()
	})
}

func Fuzz_Nosy_L1BeaconEndpointConfig_Check__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *L1BeaconEndpointConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		cfg.Check()
	})
}

// skipping Fuzz_Nosy_L1BeaconEndpointConfig_Setup__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/log.Logger

func Fuzz_Nosy_L1BeaconEndpointConfig_ShouldFetchAllSidecars__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *L1BeaconEndpointConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		cfg.ShouldFetchAllSidecars()
	})
}

func Fuzz_Nosy_L1BeaconEndpointConfig_ShouldIgnoreBeaconCheck__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *L1BeaconEndpointConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		cfg.ShouldIgnoreBeaconCheck()
	})
}

func Fuzz_Nosy_L1EndpointConfig_Check__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *L1EndpointConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		cfg.Check()
	})
}

// skipping Fuzz_Nosy_L1EndpointConfig_Setup__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/log.Logger

func Fuzz_Nosy_L2EndpointConfig_Check__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *L2EndpointConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		cfg.Check()
	})
}

// skipping Fuzz_Nosy_L2EndpointConfig_Setup__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/log.Logger

func Fuzz_Nosy_OpNode_HTTPEndpoint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *OpNode
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.HTTPEndpoint()
	})
}

func Fuzz_Nosy_OpNode_OnNewL1Finalized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *OpNode
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var sig eth.L1BlockRef
		fill_err = tp.Fill(&sig)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.OnNewL1Finalized(ctx, sig)
	})
}

func Fuzz_Nosy_OpNode_OnNewL1Head__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *OpNode
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var sig eth.L1BlockRef
		fill_err = tp.Fill(&sig)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.OnNewL1Head(ctx, sig)
	})
}

func Fuzz_Nosy_OpNode_OnNewL1Safe__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *OpNode
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var sig eth.L1BlockRef
		fill_err = tp.Fill(&sig)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.OnNewL1Safe(ctx, sig)
	})
}

func Fuzz_Nosy_OpNode_OnUnsafeL2Payload__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *OpNode
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var from peer.ID
		fill_err = tp.Fill(&from)
		if fill_err != nil {
			return
		}
		var envelope *eth.ExecutionPayloadEnvelope
		fill_err = tp.Fill(&envelope)
		if fill_err != nil {
			return
		}
		if n == nil || envelope == nil {
			return
		}

		n.OnUnsafeL2Payload(ctx, from, envelope)
	})
}

func Fuzz_Nosy_OpNode_P2P__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *OpNode
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.P2P()
	})
}

func Fuzz_Nosy_OpNode_PublishL2Payload__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *OpNode
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var envelope *eth.ExecutionPayloadEnvelope
		fill_err = tp.Fill(&envelope)
		if fill_err != nil {
			return
		}
		if n == nil || envelope == nil {
			return
		}

		n.PublishL2Payload(ctx, envelope)
	})
}

func Fuzz_Nosy_OpNode_RequestL2Range__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *OpNode
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var start eth.L2BlockRef
		fill_err = tp.Fill(&start)
		if fill_err != nil {
			return
		}
		var end eth.L2BlockRef
		fill_err = tp.Fill(&end)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.RequestL2Range(ctx, start, end)
	})
}

func Fuzz_Nosy_OpNode_RuntimeConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *OpNode
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.RuntimeConfig()
	})
}

func Fuzz_Nosy_OpNode_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *OpNode
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.Start(ctx)
	})
}

func Fuzz_Nosy_OpNode_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *OpNode
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.Stop(ctx)
	})
}

func Fuzz_Nosy_OpNode_Stopped__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *OpNode
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.Stopped()
	})
}

func Fuzz_Nosy_OpNode_haltMaybe__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *OpNode
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.haltMaybe()
	})
}

func Fuzz_Nosy_OpNode_handleProtocolVersionsUpdate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *OpNode
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.handleProtocolVersionsUpdate(ctx)
	})
}

func Fuzz_Nosy_OpNode_init__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *OpNode
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var cfg *Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if n == nil || cfg == nil {
			return
		}

		n.init(ctx, cfg)
	})
}

func Fuzz_Nosy_OpNode_initEventSystem__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *OpNode
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.initEventSystem()
	})
}

func Fuzz_Nosy_OpNode_initL1__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *OpNode
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var cfg *Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if n == nil || cfg == nil {
			return
		}

		n.initL1(ctx, cfg)
	})
}

func Fuzz_Nosy_OpNode_initL1BeaconAPI__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *OpNode
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var cfg *Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if n == nil || cfg == nil {
			return
		}

		n.initL1BeaconAPI(ctx, cfg)
	})
}

func Fuzz_Nosy_OpNode_initL2__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *OpNode
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var cfg *Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if n == nil || cfg == nil {
			return
		}

		n.initL2(ctx, cfg)
	})
}

func Fuzz_Nosy_OpNode_initMetricsServer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *OpNode
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var cfg *Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if n == nil || cfg == nil {
			return
		}

		n.initMetricsServer(cfg)
	})
}

func Fuzz_Nosy_OpNode_initP2P__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *OpNode
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var cfg *Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if n == nil || cfg == nil {
			return
		}

		n.initP2P(cfg)
	})
}

func Fuzz_Nosy_OpNode_initP2PSigner__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *OpNode
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var cfg *Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if n == nil || cfg == nil {
			return
		}

		n.initP2PSigner(ctx, cfg)
	})
}

func Fuzz_Nosy_OpNode_initPProf__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *OpNode
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var cfg *Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if n == nil || cfg == nil {
			return
		}

		n.initPProf(cfg)
	})
}

func Fuzz_Nosy_OpNode_initRPCServer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *OpNode
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var cfg *Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if n == nil || cfg == nil {
			return
		}

		n.initRPCServer(cfg)
	})
}

func Fuzz_Nosy_OpNode_initRuntimeConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *OpNode
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var cfg *Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if n == nil || cfg == nil {
			return
		}

		n.initRuntimeConfig(ctx, cfg)
	})
}

func Fuzz_Nosy_OpNode_initTracer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *OpNode
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var cfg *Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if n == nil || cfg == nil {
			return
		}

		n.initTracer(ctx, cfg)
	})
}

// skipping Fuzz_Nosy_OpNode_onEvent__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/event.Event

func Fuzz_Nosy_OpNode_p2pEnabled__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *OpNode
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.p2pEnabled()
	})
}

func Fuzz_Nosy_PreparedL1Endpoint_Check__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *PreparedL1Endpoint
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		cfg.Check()
	})
}

// skipping Fuzz_Nosy_PreparedL1Endpoint_Setup__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/log.Logger

func Fuzz_Nosy_PreparedL2Endpoints_Check__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *PreparedL2Endpoints
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.Check()
	})
}

// skipping Fuzz_Nosy_PreparedL2Endpoints_Setup__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/log.Logger

func Fuzz_Nosy_RPCConfig_HttpEndpoint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *RPCConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		cfg.HttpEndpoint()
	})
}

func Fuzz_Nosy_RuntimeConfig_Load__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *RuntimeConfig
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var l1Ref eth.L1BlockRef
		fill_err = tp.Fill(&l1Ref)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.Load(ctx, l1Ref)
	})
}

func Fuzz_Nosy_RuntimeConfig_P2PSequencerAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *RuntimeConfig
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.P2PSequencerAddress()
	})
}

func Fuzz_Nosy_RuntimeConfig_RecommendedProtocolVersion__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *RuntimeConfig
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.RecommendedProtocolVersion()
	})
}

func Fuzz_Nosy_RuntimeConfig_RequiredProtocolVersion__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *RuntimeConfig
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.RequiredProtocolVersion()
	})
}

func Fuzz_Nosy_SupervisorEndpointConfig_Check__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *SupervisorEndpointConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		cfg.Check()
	})
}

// skipping Fuzz_Nosy_SupervisorEndpointConfig_SupervisorClient__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/log.Logger

func Fuzz_Nosy_adminAPI_ConductorEnabled__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *adminAPI
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.ConductorEnabled(ctx)
	})
}

func Fuzz_Nosy_adminAPI_OverrideLeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *adminAPI
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.OverrideLeader(ctx)
	})
}

func Fuzz_Nosy_adminAPI_PostUnsafePayload__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *adminAPI
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var envelope *eth.ExecutionPayloadEnvelope
		fill_err = tp.Fill(&envelope)
		if fill_err != nil {
			return
		}
		if n == nil || envelope == nil {
			return
		}

		n.PostUnsafePayload(ctx, envelope)
	})
}

func Fuzz_Nosy_adminAPI_ResetDerivationPipeline__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *adminAPI
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.ResetDerivationPipeline(ctx)
	})
}

func Fuzz_Nosy_adminAPI_SequencerActive__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *adminAPI
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.SequencerActive(ctx)
	})
}

func Fuzz_Nosy_adminAPI_StartSequencer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *adminAPI
		fill_err = tp.Fill(&n)
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
		if n == nil {
			return
		}

		n.StartSequencer(ctx, blockHash)
	})
}

func Fuzz_Nosy_adminAPI_StopSequencer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *adminAPI
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.StopSequencer(ctx)
	})
}

func Fuzz_Nosy_nodeAPI_OutputAtBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *nodeAPI
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var number hexutil.Uint64
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.OutputAtBlock(ctx, number)
	})
}

func Fuzz_Nosy_nodeAPI_RollupConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *nodeAPI
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.RollupConfig(_x2)
	})
}

func Fuzz_Nosy_nodeAPI_SafeHeadAtL1Block__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *nodeAPI
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var number hexutil.Uint64
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.SafeHeadAtL1Block(ctx, number)
	})
}

func Fuzz_Nosy_nodeAPI_SyncStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *nodeAPI
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.SyncStatus(ctx)
	})
}

func Fuzz_Nosy_nodeAPI_Version__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *nodeAPI
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.Version(ctx)
	})
}

func Fuzz_Nosy_rpcServer_Addr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *rpcServer
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.Addr()
	})
}

func Fuzz_Nosy_rpcServer_EnableAdminAPI__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *rpcServer
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var api *adminAPI
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		if s == nil || api == nil {
			return
		}

		s.EnableAdminAPI(api)
	})
}

func Fuzz_Nosy_rpcServer_EnableP2P__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *rpcServer
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var backend *p2p.APIBackend
		fill_err = tp.Fill(&backend)
		if fill_err != nil {
			return
		}
		if s == nil || backend == nil {
			return
		}

		s.EnableP2P(backend)
	})
}

func Fuzz_Nosy_rpcServer_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *rpcServer
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Start()
	})
}

func Fuzz_Nosy_rpcServer_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *rpcServer
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.Stop(ctx)
	})
}

// skipping Fuzz_Nosy_ConfigPersistence_SequencerStarted__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/node.ConfigPersistence

// skipping Fuzz_Nosy_ConfigPersistence_SequencerState__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/node.ConfigPersistence

// skipping Fuzz_Nosy_ConfigPersistence_SequencerStopped__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/node.ConfigPersistence

func Fuzz_Nosy_DisabledConfigPersistence_SequencerStarted__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d DisabledConfigPersistence
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}

		d.SequencerStarted()
	})
}

func Fuzz_Nosy_DisabledConfigPersistence_SequencerState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d DisabledConfigPersistence
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}

		d.SequencerState()
	})
}

func Fuzz_Nosy_DisabledConfigPersistence_SequencerStopped__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d DisabledConfigPersistence
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}

		d.SequencerStopped()
	})
}

// skipping Fuzz_Nosy_L1BeaconEndpointSetup_Check__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/node.L1BeaconEndpointSetup

// skipping Fuzz_Nosy_L1BeaconEndpointSetup_Setup__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/node.L1BeaconEndpointSetup

// skipping Fuzz_Nosy_L1BeaconEndpointSetup_ShouldFetchAllSidecars__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/node.L1BeaconEndpointSetup

// skipping Fuzz_Nosy_L1BeaconEndpointSetup_ShouldIgnoreBeaconCheck__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/node.L1BeaconEndpointSetup

// skipping Fuzz_Nosy_L1EndpointSetup_Check__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/node.L1EndpointSetup

// skipping Fuzz_Nosy_L1EndpointSetup_Setup__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/node.L1EndpointSetup

// skipping Fuzz_Nosy_L2EndpointSetup_Check__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/node.L2EndpointSetup

// skipping Fuzz_Nosy_L2EndpointSetup_Setup__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/node.L2EndpointSetup

func Fuzz_Nosy_MetricsConfig_Check__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m MetricsConfig
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}

		m.Check()
	})
}

// skipping Fuzz_Nosy_ReadonlyRuntimeConfig_P2PSequencerAddress__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/node.ReadonlyRuntimeConfig

// skipping Fuzz_Nosy_ReadonlyRuntimeConfig_RecommendedProtocolVersion__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/node.ReadonlyRuntimeConfig

// skipping Fuzz_Nosy_ReadonlyRuntimeConfig_RequiredProtocolVersion__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/node.ReadonlyRuntimeConfig

// skipping Fuzz_Nosy_RuntimeCfgL1Source_ReadStorageAt__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/node.RuntimeCfgL1Source

// skipping Fuzz_Nosy_SafeDBReader_SafeHeadAtL1__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/node.SafeDBReader

// skipping Fuzz_Nosy_SupervisorEndpointSetup_Check__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/node.SupervisorEndpointSetup

// skipping Fuzz_Nosy_SupervisorEndpointSetup_SupervisorClient__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/node.SupervisorEndpointSetup

// skipping Fuzz_Nosy_Tracer_OnNewL1Head__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/node.Tracer

// skipping Fuzz_Nosy_Tracer_OnPublishL2Payload__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/node.Tracer

// skipping Fuzz_Nosy_Tracer_OnUnsafeL2Payload__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/node.Tracer

// skipping Fuzz_Nosy_driverClient_BlockRefWithStatus__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/node.driverClient

// skipping Fuzz_Nosy_driverClient_ConductorEnabled__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/node.driverClient

// skipping Fuzz_Nosy_driverClient_OnUnsafeL2Payload__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/node.driverClient

// skipping Fuzz_Nosy_driverClient_OverrideLeader__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/node.driverClient

// skipping Fuzz_Nosy_driverClient_ResetDerivationPipeline__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/node.driverClient

// skipping Fuzz_Nosy_driverClient_SequencerActive__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/node.driverClient

// skipping Fuzz_Nosy_driverClient_StartSequencer__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/node.driverClient

// skipping Fuzz_Nosy_driverClient_StopSequencer__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/node.driverClient

// skipping Fuzz_Nosy_driverClient_SyncStatus__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/node.driverClient

// skipping Fuzz_Nosy_l2EthClient_GetProof__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/node.l2EthClient

// skipping Fuzz_Nosy_l2EthClient_InfoByHash__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/node.l2EthClient

// skipping Fuzz_Nosy_l2EthClient_OutputV0AtBlock__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/node.l2EthClient

func Fuzz_Nosy_noOpTracer_OnNewL1Head__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n noOpTracer
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var sig eth.L1BlockRef
		fill_err = tp.Fill(&sig)
		if fill_err != nil {
			return
		}

		n.OnNewL1Head(ctx, sig)
	})
}

func Fuzz_Nosy_noOpTracer_OnPublishL2Payload__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n noOpTracer
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var payload *eth.ExecutionPayloadEnvelope
		fill_err = tp.Fill(&payload)
		if fill_err != nil {
			return
		}
		if payload == nil {
			return
		}

		n.OnPublishL2Payload(ctx, payload)
	})
}

func Fuzz_Nosy_noOpTracer_OnUnsafeL2Payload__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n noOpTracer
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var from peer.ID
		fill_err = tp.Fill(&from)
		if fill_err != nil {
			return
		}
		var payload *eth.ExecutionPayloadEnvelope
		fill_err = tp.Fill(&payload)
		if fill_err != nil {
			return
		}
		if payload == nil {
			return
		}

		n.OnUnsafeL2Payload(ctx, from, payload)
	})
}

func Fuzz_Nosy_haltMaybe__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var haltOption string
		fill_err = tp.Fill(&haltOption)
		if fill_err != nil {
			return
		}
		var reqCmp params.ProtocolVersionComparison
		fill_err = tp.Fill(&reqCmp)
		if fill_err != nil {
			return
		}

		haltMaybe(haltOption, reqCmp)
	})
}

func Fuzz_Nosy_unixTimeStale__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var timestamp uint64
		fill_err = tp.Fill(&timestamp)
		if fill_err != nil {
			return
		}
		var duration time.Duration
		fill_err = tp.Fill(&duration)
		if fill_err != nil {
			return
		}

		unixTimeStale(timestamp, duration)
	})
}
