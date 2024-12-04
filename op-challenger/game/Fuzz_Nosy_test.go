package game

import (
	"context"
	"testing"

	"github.com/ethereum-optimism/optimism/op-challenger/config"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum-optimism/optimism/op-service/metrics"
	"github.com/ethereum-optimism/optimism/op-service/oppprof"
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

func Fuzz_Nosy_Service_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Start(ctx)
	})
}

func Fuzz_Nosy_Service_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Stop(ctx)
	})
}

func Fuzz_Nosy_Service_Stopped__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Stopped()
	})
}

func Fuzz_Nosy_Service_initBondClaims__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.initBondClaims()
	})
}

func Fuzz_Nosy_Service_initClaimants__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var cfg *config.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if s == nil || cfg == nil {
			return
		}

		s.initClaimants(cfg)
	})
}

func Fuzz_Nosy_Service_initFactoryContract__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var cfg *config.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if s == nil || cfg == nil {
			return
		}

		s.initFactoryContract(cfg)
	})
}

func Fuzz_Nosy_Service_initFromConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var cfg *config.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if s == nil || cfg == nil {
			return
		}

		s.initFromConfig(ctx, cfg)
	})
}

func Fuzz_Nosy_Service_initL1Client__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var cfg *config.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if s == nil || cfg == nil {
			return
		}

		s.initL1Client(ctx, cfg)
	})
}

func Fuzz_Nosy_Service_initLargePreimages__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.initLargePreimages()
	})
}

func Fuzz_Nosy_Service_initMetricsServer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var cfg *metrics.CLIConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if s == nil || cfg == nil {
			return
		}

		s.initMetricsServer(cfg)
	})
}

func Fuzz_Nosy_Service_initMonitor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var cfg *config.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if s == nil || cfg == nil {
			return
		}

		s.initMonitor(cfg)
	})
}

func Fuzz_Nosy_Service_initPProf__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var cfg *oppprof.CLIConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if s == nil || cfg == nil {
			return
		}

		s.initPProf(cfg)
	})
}

func Fuzz_Nosy_Service_initPollClient__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var cfg *config.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if s == nil || cfg == nil {
			return
		}

		s.initPollClient(ctx, cfg)
	})
}

func Fuzz_Nosy_Service_initRollupClient__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var cfg *config.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if s == nil || cfg == nil {
			return
		}

		s.initRollupClient(ctx, cfg)
	})
}

func Fuzz_Nosy_Service_initScheduler__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var cfg *config.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if s == nil || cfg == nil {
			return
		}

		s.initScheduler(cfg)
	})
}

func Fuzz_Nosy_Service_initTxManager__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var cfg *config.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if s == nil || cfg == nil {
			return
		}

		s.initTxManager(ctx, cfg)
	})
}

func Fuzz_Nosy_Service_registerGameTypes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var cfg *config.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if s == nil || cfg == nil {
			return
		}

		s.registerGameTypes(ctx, cfg)
	})
}

func Fuzz_Nosy_diskManager_DirForGame__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dir string
		fill_err = tp.Fill(&dir)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}

		d := newDiskManager(dir)
		d.DirForGame(addr)
	})
}

func Fuzz_Nosy_diskManager_RemoveAllExcept__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dir string
		fill_err = tp.Fill(&dir)
		if fill_err != nil {
			return
		}
		var keep []common.Address
		fill_err = tp.Fill(&keep)
		if fill_err != nil {
			return
		}

		d := newDiskManager(dir)
		d.RemoveAllExcept(keep)
	})
}

func Fuzz_Nosy_gameMonitor_StartMonitoring__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *gameMonitor
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.StartMonitoring()
	})
}

func Fuzz_Nosy_gameMonitor_StopMonitoring__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *gameMonitor
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.StopMonitoring()
	})
}

func Fuzz_Nosy_gameMonitor_allowedGame__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *gameMonitor
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var g2 common.Address
		fill_err = tp.Fill(&g2)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.allowedGame(g2)
	})
}

func Fuzz_Nosy_gameMonitor_onNewL1Head__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *gameMonitor
		fill_err = tp.Fill(&m)
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
		if m == nil {
			return
		}

		m.onNewL1Head(ctx, sig)
	})
}

func Fuzz_Nosy_gameMonitor_progressGames__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *gameMonitor
		fill_err = tp.Fill(&m)
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
		var blockNumber uint64
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.progressGames(ctx, blockHash, blockNumber)
	})
}

func Fuzz_Nosy_gameMonitor_resubscribeFunction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *gameMonitor
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.resubscribeFunction()
	})
}

// skipping Fuzz_Nosy_headSource_SubscribeNewHead__ because parameters include func, chan, or unsupported interface: chan<- *github.com/ethereum/go-ethereum/core/types.Header

// skipping Fuzz_Nosy_MinimalSubscriber_EthSubscribe__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game.MinimalSubscriber

// skipping Fuzz_Nosy_RWClock_Now__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game.RWClock

// skipping Fuzz_Nosy_RWClock_SetTime__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game.RWClock

// skipping Fuzz_Nosy_claimer_Schedule__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game.claimer

// skipping Fuzz_Nosy_gameScheduler_Schedule__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game.gameScheduler

// skipping Fuzz_Nosy_gameSource_GetGamesAtOrAfter__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game.gameSource

// skipping Fuzz_Nosy_preimageScheduler_Schedule__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game.preimageScheduler
