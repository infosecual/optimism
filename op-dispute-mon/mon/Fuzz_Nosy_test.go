package mon

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum-optimism/optimism/op-dispute-mon/config"
	"github.com/ethereum-optimism/optimism/op-dispute-mon/mon/types"
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

func Fuzz_Nosy_ClaimMonitor_CheckClaims__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *ClaimMonitor
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var games []*types.EnrichedGameData
		fill_err = tp.Fill(&games)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.CheckClaims(games)
	})
}

func Fuzz_Nosy_ClaimMonitor_checkGameClaims__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *ClaimMonitor
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var game *types.EnrichedGameData
		fill_err = tp.Fill(&game)
		if fill_err != nil {
			return
		}
		var claimStatuses *metrics.ClaimStatuses
		fill_err = tp.Fill(&claimStatuses)
		if fill_err != nil {
			return
		}
		var honest map[common.Address]*metrics.HonestActorData
		fill_err = tp.Fill(&honest)
		if fill_err != nil {
			return
		}
		if c == nil || game == nil || claimStatuses == nil {
			return
		}

		c.checkGameClaims(game, claimStatuses, honest)
	})
}

func Fuzz_Nosy_ClaimMonitor_checkUpdateHonestActorStats__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *ClaimMonitor
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var proxy common.Address
		fill_err = tp.Fill(&proxy)
		if fill_err != nil {
			return
		}
		var claim *types.EnrichedClaim
		fill_err = tp.Fill(&claim)
		if fill_err != nil {
			return
		}
		var honest map[common.Address]*metrics.HonestActorData
		fill_err = tp.Fill(&honest)
		if fill_err != nil {
			return
		}
		if c == nil || claim == nil {
			return
		}

		c.checkUpdateHonestActorStats(proxy, claim, honest)
	})
}

func Fuzz_Nosy_Forecast_Forecast__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *Forecast
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var games []*types.EnrichedGameData
		fill_err = tp.Fill(&games)
		if fill_err != nil {
			return
		}
		var ignoredCount int
		fill_err = tp.Fill(&ignoredCount)
		if fill_err != nil {
			return
		}
		var failedCount int
		fill_err = tp.Fill(&failedCount)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.Forecast(games, ignoredCount, failedCount)
	})
}

func Fuzz_Nosy_Forecast_forecastGame__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *Forecast
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var game *types.EnrichedGameData
		fill_err = tp.Fill(&game)
		if fill_err != nil {
			return
		}
		var metrics *forecastBatch
		fill_err = tp.Fill(&metrics)
		if fill_err != nil {
			return
		}
		if f1 == nil || game == nil || metrics == nil {
			return
		}

		f1.forecastGame(game, metrics)
	})
}

func Fuzz_Nosy_Forecast_recordBatch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *Forecast
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var batch forecastBatch
		fill_err = tp.Fill(&batch)
		if fill_err != nil {
			return
		}
		var ignoredCount int
		fill_err = tp.Fill(&ignoredCount)
		if fill_err != nil {
			return
		}
		var failedCount int
		fill_err = tp.Fill(&failedCount)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.recordBatch(batch, ignoredCount, failedCount)
	})
}

func Fuzz_Nosy_L2ChallengesMonitor_CheckL2Challenges__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *L2ChallengesMonitor
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var games []*types.EnrichedGameData
		fill_err = tp.Fill(&games)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.CheckL2Challenges(games)
	})
}

func Fuzz_Nosy_ResolutionMonitor_CheckResolutions__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *ResolutionMonitor
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var games []*types.EnrichedGameData
		fill_err = tp.Fill(&games)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.CheckResolutions(games)
	})
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

func Fuzz_Nosy_Service_initBonds__(f *testing.F) {
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

		s.initBonds()
	})
}

func Fuzz_Nosy_Service_initClaimMonitor__(f *testing.F) {
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

		s.initClaimMonitor(cfg)
	})
}

func Fuzz_Nosy_Service_initExtractor__(f *testing.F) {
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

		s.initExtractor(cfg)
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

func Fuzz_Nosy_Service_initForecast__(f *testing.F) {
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

		s.initForecast(cfg)
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

func Fuzz_Nosy_Service_initGameCallerCreator__(f *testing.F) {
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

		s.initGameCallerCreator()
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

		s.initMonitor(ctx, cfg)
	})
}

func Fuzz_Nosy_Service_initOutputRollupClient__(f *testing.F) {
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

		s.initOutputRollupClient(ctx, cfg)
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

func Fuzz_Nosy_Service_initResolutionMonitor__(f *testing.F) {
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

		s.initResolutionMonitor()
	})
}

func Fuzz_Nosy_Service_initWithdrawalMonitor__(f *testing.F) {
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

		s.initWithdrawalMonitor()
	})
}

func Fuzz_Nosy_UpdateTimeMonitor_CheckUpdateTimes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *UpdateTimeMonitor
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var games []*types.EnrichedGameData
		fill_err = tp.Fill(&games)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.CheckUpdateTimes(games)
	})
}

func Fuzz_Nosy_WithdrawalMonitor_CheckWithdrawals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *WithdrawalMonitor
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var games []*types.EnrichedGameData
		fill_err = tp.Fill(&games)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		w.CheckWithdrawals(games)
	})
}

func Fuzz_Nosy_WithdrawalMonitor_validateGameWithdrawals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *WithdrawalMonitor
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var game *types.EnrichedGameData
		fill_err = tp.Fill(&game)
		if fill_err != nil {
			return
		}
		var now time.Time
		fill_err = tp.Fill(&now)
		if fill_err != nil {
			return
		}
		var honestWithdrawableAmounts map[common.Address]*big.Int
		fill_err = tp.Fill(&honestWithdrawableAmounts)
		if fill_err != nil {
			return
		}
		if w == nil || game == nil {
			return
		}

		w.validateGameWithdrawals(game, now, honestWithdrawableAmounts)
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

func Fuzz_Nosy_gameMonitor_loop__(f *testing.F) {
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

		m.loop()
	})
}

func Fuzz_Nosy_gameMonitor_monitorGames__(f *testing.F) {
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

		m.monitorGames()
	})
}

// skipping Fuzz_Nosy_ClaimMetrics_RecordClaims__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-dispute-mon/mon.ClaimMetrics

// skipping Fuzz_Nosy_ClaimMetrics_RecordHonestActorClaims__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-dispute-mon/mon.ClaimMetrics

// skipping Fuzz_Nosy_ForecastMetrics_RecordFailedGames__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-dispute-mon/mon.ForecastMetrics

// skipping Fuzz_Nosy_ForecastMetrics_RecordGameAgreement__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-dispute-mon/mon.ForecastMetrics

// skipping Fuzz_Nosy_ForecastMetrics_RecordIgnoredGames__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-dispute-mon/mon.ForecastMetrics

// skipping Fuzz_Nosy_ForecastMetrics_RecordLatestProposals__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-dispute-mon/mon.ForecastMetrics

// skipping Fuzz_Nosy_ForecastMetrics_RecordLatestValidProposalL2Block__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-dispute-mon/mon.ForecastMetrics

// skipping Fuzz_Nosy_L2ChallengesMetrics_RecordL2Challenges__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-dispute-mon/mon.L2ChallengesMetrics

// skipping Fuzz_Nosy_MonitorMetrics_RecordMonitorDuration__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-dispute-mon/mon.MonitorMetrics

// skipping Fuzz_Nosy_RClock_Now__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-dispute-mon/mon.RClock

// skipping Fuzz_Nosy_ResolutionMetrics_RecordGameResolutionStatus__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-dispute-mon/mon.ResolutionMetrics

// skipping Fuzz_Nosy_UpdateTimeMetrics_RecordOldestGameUpdateTime__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-dispute-mon/mon.UpdateTimeMetrics

// skipping Fuzz_Nosy_WithdrawalMetrics_RecordHonestWithdrawableAmounts__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-dispute-mon/mon.WithdrawalMetrics

// skipping Fuzz_Nosy_WithdrawalMetrics_RecordWithdrawalRequests__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-dispute-mon/mon.WithdrawalMetrics
