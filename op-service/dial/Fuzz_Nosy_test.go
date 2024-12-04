package dial

import (
	"context"
	"testing"

	"github.com/ethereum-optimism/optimism/op-service/sources"
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

func Fuzz_Nosy_ActiveL2EndpointProvider_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *ActiveL2EndpointProvider
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.Close()
	})
}

func Fuzz_Nosy_ActiveL2EndpointProvider_EthClient__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *ActiveL2EndpointProvider
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.EthClient(ctx)
	})
}

func Fuzz_Nosy_ActiveL2RollupProvider_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *ActiveL2RollupProvider
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.Close()
	})
}

func Fuzz_Nosy_ActiveL2RollupProvider_RollupClient__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *ActiveL2RollupProvider
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.RollupClient(ctx)
	})
}

func Fuzz_Nosy_ActiveL2RollupProvider_checkCurrentSequencer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *ActiveL2RollupProvider
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.checkCurrentSequencer(ctx)
	})
}

func Fuzz_Nosy_ActiveL2RollupProvider_dialSequencer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *ActiveL2RollupProvider
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var idx int
		fill_err = tp.Fill(&idx)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.dialSequencer(ctx, idx)
	})
}

func Fuzz_Nosy_ActiveL2RollupProvider_ensureActiveEndpoint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *ActiveL2RollupProvider
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.ensureActiveEndpoint(ctx)
	})
}

func Fuzz_Nosy_ActiveL2RollupProvider_findActiveEndpoints__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *ActiveL2RollupProvider
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.findActiveEndpoints(ctx)
	})
}

func Fuzz_Nosy_ActiveL2RollupProvider_numEndpoints__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *ActiveL2RollupProvider
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.numEndpoints()
	})
}

func Fuzz_Nosy_ActiveL2RollupProvider_shouldCheck__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *ActiveL2RollupProvider
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.shouldCheck()
	})
}

func Fuzz_Nosy_StaticL2EndpointProvider_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *StaticL2EndpointProvider
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.Close()
	})
}

func Fuzz_Nosy_StaticL2EndpointProvider_EthClient__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *StaticL2EndpointProvider
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.EthClient(_x2)
	})
}

func Fuzz_Nosy_StaticL2RollupProvider_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rollupCl *sources.RollupClient
		fill_err = tp.Fill(&rollupCl)
		if fill_err != nil {
			return
		}
		if rollupCl == nil {
			return
		}

		p, err := NewStaticL2RollupProviderFromExistingRollup(rollupCl)
		if err != nil {
			return
		}
		p.Close()
	})
}

func Fuzz_Nosy_StaticL2RollupProvider_RollupClient__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rollupCl *sources.RollupClient
		fill_err = tp.Fill(&rollupCl)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if rollupCl == nil {
			return
		}

		p, err := NewStaticL2RollupProviderFromExistingRollup(rollupCl)
		if err != nil {
			return
		}
		p.RollupClient(_x2)
	})
}

// skipping Fuzz_Nosy_EthClientInterface_BlockByNumber__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/dial.EthClientInterface

// skipping Fuzz_Nosy_EthClientInterface_Client__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/dial.EthClientInterface

// skipping Fuzz_Nosy_EthClientInterface_Close__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/dial.EthClientInterface

// skipping Fuzz_Nosy_L2EndpointProvider_EthClient__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/dial.L2EndpointProvider

// skipping Fuzz_Nosy_RollupClientInterface_Close__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/dial.RollupClientInterface

// skipping Fuzz_Nosy_RollupClientInterface_OutputAtBlock__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/dial.RollupClientInterface

// skipping Fuzz_Nosy_RollupClientInterface_RollupConfig__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/dial.RollupClientInterface

// skipping Fuzz_Nosy_RollupClientInterface_SequencerActive__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/dial.RollupClientInterface

// skipping Fuzz_Nosy_RollupClientInterface_StartSequencer__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/dial.RollupClientInterface

// skipping Fuzz_Nosy_RollupProvider_Close__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/dial.RollupProvider

// skipping Fuzz_Nosy_RollupProvider_RollupClient__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/dial.RollupProvider

// skipping Fuzz_Nosy_SyncStatusProvider_SyncStatus__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/dial.SyncStatusProvider
