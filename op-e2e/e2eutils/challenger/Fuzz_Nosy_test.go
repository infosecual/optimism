package challenger

import (
	"context"
	"testing"

	"github.com/ethereum-optimism/optimism/op-node/rollup"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
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

func Fuzz_Nosy_CapturingMetrics_RecordActedL1Block__(f *testing.F) {
	f.Fuzz(func(t *testing.T, block uint64) {
		c := NewCapturingMetrics()
		c.RecordActedL1Block(block)
	})
}

func Fuzz_Nosy_Helper_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *Helper
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.Close()
	})
}

// skipping Fuzz_Nosy_Helper_VerifyGameDataExists__ because parameters include func, chan, or unsupported interface: []github.com/ethereum-optimism/optimism/op-e2e/e2eutils/challenger.GameAddr

// skipping Fuzz_Nosy_Helper_WaitForGameDataDeletion__ because parameters include func, chan, or unsupported interface: []github.com/ethereum-optimism/optimism/op-e2e/e2eutils/challenger.GameAddr

func Fuzz_Nosy_Helper_WaitForHighestActedL1Block__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *Helper
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var head uint64
		fill_err = tp.Fill(&head)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.WaitForHighestActedL1Block(ctx, head)
	})
}

func Fuzz_Nosy_Helper_WaitL1HeadActedOn__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *Helper
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var client *ethclient.Client
		fill_err = tp.Fill(&client)
		if fill_err != nil {
			return
		}
		if h == nil || client == nil {
			return
		}

		h.WaitL1HeadActedOn(ctx, client)
	})
}

func Fuzz_Nosy_Helper_gameDataDir__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *Helper
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.gameDataDir(addr)
	})
}

// skipping Fuzz_Nosy_EndpointProvider_L1BeaconEndpoint__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/e2eutils/challenger.EndpointProvider

// skipping Fuzz_Nosy_EndpointProvider_NodeEndpoint__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/e2eutils/challenger.EndpointProvider

// skipping Fuzz_Nosy_EndpointProvider_RollupEndpoint__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/e2eutils/challenger.EndpointProvider

// skipping Fuzz_Nosy_GameAddr_Addr__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/e2eutils/challenger.GameAddr

// skipping Fuzz_Nosy_System_AllocType__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/e2eutils/challenger.System

// skipping Fuzz_Nosy_System_L2Genesis__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/e2eutils/challenger.System

// skipping Fuzz_Nosy_System_RollupCfg__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/e2eutils/challenger.System

func Fuzz_Nosy_FindMonorepoRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		FindMonorepoRoot(t1)
	})
}

func Fuzz_Nosy_applyCannonConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *config.Config
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var t2 *testing.T
		fill_err = tp.Fill(&t2)
		if fill_err != nil {
			return
		}
		var rollupCfg *rollup.Config
		fill_err = tp.Fill(&rollupCfg)
		if fill_err != nil {
			return
		}
		var l2Genesis *core.Genesis
		fill_err = tp.Fill(&l2Genesis)
		if fill_err != nil {
			return
		}
		var allocType config.AllocType
		fill_err = tp.Fill(&allocType)
		if fill_err != nil {
			return
		}
		if c == nil || t2 == nil || rollupCfg == nil || l2Genesis == nil {
			return
		}

		applyCannonConfig(c, t2, rollupCfg, l2Genesis, allocType)
	})
}
