package e2esys

import (
	"testing"
	"time"

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

func Fuzz_Nosy_System_AdvanceTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sys *System
		fill_err = tp.Fill(&sys)
		if fill_err != nil {
			return
		}
		var d time.Duration
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		if sys == nil {
			return
		}

		sys.AdvanceTime(d)
	})
}

func Fuzz_Nosy_System_AllocType__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sys *System
		fill_err = tp.Fill(&sys)
		if fill_err != nil {
			return
		}
		if sys == nil {
			return
		}

		sys.AllocType()
	})
}

func Fuzz_Nosy_System_BatcherHelper__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sys *System
		fill_err = tp.Fill(&sys)
		if fill_err != nil {
			return
		}
		if sys == nil {
			return
		}

		sys.BatcherHelper()
	})
}

func Fuzz_Nosy_System_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sys *System
		fill_err = tp.Fill(&sys)
		if fill_err != nil {
			return
		}
		if sys == nil {
			return
		}

		sys.Close()
	})
}

func Fuzz_Nosy_System_Config__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sys *System
		fill_err = tp.Fill(&sys)
		if fill_err != nil {
			return
		}
		if sys == nil {
			return
		}

		sys.Config()
	})
}

func Fuzz_Nosy_System_L1BeaconEndpoint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sys *System
		fill_err = tp.Fill(&sys)
		if fill_err != nil {
			return
		}
		if sys == nil {
			return
		}

		sys.L1BeaconEndpoint()
	})
}

func Fuzz_Nosy_System_L1BeaconHTTPClient__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sys *System
		fill_err = tp.Fill(&sys)
		if fill_err != nil {
			return
		}
		if sys == nil {
			return
		}

		sys.L1BeaconHTTPClient()
	})
}

func Fuzz_Nosy_System_L1Deployments__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sys *System
		fill_err = tp.Fill(&sys)
		if fill_err != nil {
			return
		}
		if sys == nil {
			return
		}

		sys.L1Deployments()
	})
}

func Fuzz_Nosy_System_L1Slot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sys *System
		fill_err = tp.Fill(&sys)
		if fill_err != nil {
			return
		}
		var l1Timestamp uint64
		fill_err = tp.Fill(&l1Timestamp)
		if fill_err != nil {
			return
		}
		if sys == nil {
			return
		}

		sys.L1Slot(l1Timestamp)
	})
}

func Fuzz_Nosy_System_L2Genesis__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sys *System
		fill_err = tp.Fill(&sys)
		if fill_err != nil {
			return
		}
		if sys == nil {
			return
		}

		sys.L2Genesis()
	})
}

func Fuzz_Nosy_System_NewMockNetPeer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sys *System
		fill_err = tp.Fill(&sys)
		if fill_err != nil {
			return
		}
		if sys == nil {
			return
		}

		sys.NewMockNetPeer()
	})
}

func Fuzz_Nosy_System_NodeClient__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sys *System
		fill_err = tp.Fill(&sys)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		if sys == nil {
			return
		}

		sys.NodeClient(name)
	})
}

func Fuzz_Nosy_System_NodeEndpoint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sys *System
		fill_err = tp.Fill(&sys)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		if sys == nil {
			return
		}

		sys.NodeEndpoint(name)
	})
}

func Fuzz_Nosy_System_RollupCfg__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sys *System
		fill_err = tp.Fill(&sys)
		if fill_err != nil {
			return
		}
		if sys == nil {
			return
		}

		sys.RollupCfg()
	})
}

func Fuzz_Nosy_System_RollupClient__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sys *System
		fill_err = tp.Fill(&sys)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		if sys == nil {
			return
		}

		sys.RollupClient(name)
	})
}

func Fuzz_Nosy_System_RollupEndpoint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sys *System
		fill_err = tp.Fill(&sys)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		if sys == nil {
			return
		}

		sys.RollupEndpoint(name)
	})
}

func Fuzz_Nosy_System_TestAccount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sys *System
		fill_err = tp.Fill(&sys)
		if fill_err != nil {
			return
		}
		var idx int
		fill_err = tp.Fill(&idx)
		if fill_err != nil {
			return
		}
		if sys == nil {
			return
		}

		sys.TestAccount(idx)
	})
}

func Fuzz_Nosy_startOptions_Get__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _opts []StartOption
		fill_err = tp.Fill(&_opts)
		if fill_err != nil {
			return
		}
		var key string
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var role string
		fill_err = tp.Fill(&role)
		if fill_err != nil {
			return
		}

		s, err := parseStartOptions(_opts)
		if err != nil {
			return
		}
		s.Get(key, role)
	})
}

func Fuzz_Nosy_SystemConfig_L1ChainIDBig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg SystemConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}

		cfg.L1ChainIDBig()
	})
}

func Fuzz_Nosy_SystemConfig_L2ChainIDBig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg SystemConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}

		cfg.L2ChainIDBig()
	})
}

func Fuzz_Nosy_SystemConfig_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg SystemConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var t2 *testing.T
		fill_err = tp.Fill(&t2)
		if fill_err != nil {
			return
		}
		var startOpts []StartOption
		fill_err = tp.Fill(&startOpts)
		if fill_err != nil {
			return
		}
		if t2 == nil {
			return
		}

		cfg.Start(t2, startOpts...)
	})
}

// skipping Fuzz_Nosy_ConfigureL1__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/e2eutils/services.EthInstance

// skipping Fuzz_Nosy_ConfigureL2__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/e2eutils/services.EthInstance

// skipping Fuzz_Nosy_writeDefaultJWT__ because parameters include func, chan, or unsupported interface: testing.TB
