package metrics

import (
	"math/big"
	"testing"
	"time"

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

// skipping Fuzz_Nosy_ClaimStatuses_ForEachStatus__ because parameters include func, chan, or unsupported interface: func(status github.com/ethereum-optimism/optimism/op-dispute-mon/metrics.ClaimStatus, count int)

func Fuzz_Nosy_ClaimStatuses_RecordClaim__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *ClaimStatuses
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var firstHalf bool
		fill_err = tp.Fill(&firstHalf)
		if fill_err != nil {
			return
		}
		var clockExpired bool
		fill_err = tp.Fill(&clockExpired)
		if fill_err != nil {
			return
		}
		var resolvable bool
		fill_err = tp.Fill(&resolvable)
		if fill_err != nil {
			return
		}
		var resolved bool
		fill_err = tp.Fill(&resolved)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.RecordClaim(firstHalf, clockExpired, resolvable, resolved)
	})
}

func Fuzz_Nosy_Metrics_RecordBondCollateral__(f *testing.F) {
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
		var required *big.Int
		fill_err = tp.Fill(&required)
		if fill_err != nil {
			return
		}
		var available *big.Int
		fill_err = tp.Fill(&available)
		if fill_err != nil {
			return
		}
		if required == nil || available == nil {
			return
		}

		m := NewMetrics()
		m.RecordBondCollateral(addr, required, available)
	})
}

func Fuzz_Nosy_Metrics_RecordClaims__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var statuses *ClaimStatuses
		fill_err = tp.Fill(&statuses)
		if fill_err != nil {
			return
		}
		if statuses == nil {
			return
		}

		m := NewMetrics()
		m.RecordClaims(statuses)
	})
}

func Fuzz_Nosy_Metrics_RecordCredit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var expectation CreditExpectation
		fill_err = tp.Fill(&expectation)
		if fill_err != nil {
			return
		}
		var count int
		fill_err = tp.Fill(&count)
		if fill_err != nil {
			return
		}

		m := NewMetrics()
		m.RecordCredit(expectation, count)
	})
}

func Fuzz_Nosy_Metrics_RecordFailedGames__(f *testing.F) {
	f.Fuzz(func(t *testing.T, count int) {
		m := NewMetrics()
		m.RecordFailedGames(count)
	})
}

func Fuzz_Nosy_Metrics_RecordGameAgreement__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var status GameAgreementStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		var count int
		fill_err = tp.Fill(&count)
		if fill_err != nil {
			return
		}

		m := NewMetrics()
		m.RecordGameAgreement(status, count)
	})
}

func Fuzz_Nosy_Metrics_RecordGameResolutionStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var status ResolutionStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		var count int
		fill_err = tp.Fill(&count)
		if fill_err != nil {
			return
		}

		m := NewMetrics()
		m.RecordGameResolutionStatus(status, count)
	})
}

func Fuzz_Nosy_Metrics_RecordHonestActorClaims__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		var stats *HonestActorData
		fill_err = tp.Fill(&stats)
		if fill_err != nil {
			return
		}
		if stats == nil {
			return
		}

		m := NewMetrics()
		m.RecordHonestActorClaims(address, stats)
	})
}

func Fuzz_Nosy_Metrics_RecordHonestWithdrawableAmounts__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var amounts map[common.Address]*big.Int
		fill_err = tp.Fill(&amounts)
		if fill_err != nil {
			return
		}

		m := NewMetrics()
		m.RecordHonestWithdrawableAmounts(amounts)
	})
}

func Fuzz_Nosy_Metrics_RecordIgnoredGames__(f *testing.F) {
	f.Fuzz(func(t *testing.T, count int) {
		m := NewMetrics()
		m.RecordIgnoredGames(count)
	})
}

func Fuzz_Nosy_Metrics_RecordInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, version string) {
		m := NewMetrics()
		m.RecordInfo(version)
	})
}

func Fuzz_Nosy_Metrics_RecordL2Challenges__(f *testing.F) {
	f.Fuzz(func(t *testing.T, agreement bool, count int) {
		m := NewMetrics()
		m.RecordL2Challenges(agreement, count)
	})
}

func Fuzz_Nosy_Metrics_RecordLatestProposals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, latestValid uint64, latestInvalid uint64) {
		m := NewMetrics()
		m.RecordLatestProposals(latestValid, latestInvalid)
	})
}

func Fuzz_Nosy_Metrics_RecordLatestValidProposalL2Block__(f *testing.F) {
	f.Fuzz(func(t *testing.T, latestValid uint64) {
		m := NewMetrics()
		m.RecordLatestValidProposalL2Block(latestValid)
	})
}

func Fuzz_Nosy_Metrics_RecordMonitorDuration__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dur time.Duration
		fill_err = tp.Fill(&dur)
		if fill_err != nil {
			return
		}

		m := NewMetrics()
		m.RecordMonitorDuration(dur)
	})
}

func Fuzz_Nosy_Metrics_RecordOldestGameUpdateTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 time.Time
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}

		m := NewMetrics()
		m.RecordOldestGameUpdateTime(t1)
	})
}

func Fuzz_Nosy_Metrics_RecordOutputFetchTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, timestamp float64) {
		m := NewMetrics()
		m.RecordOutputFetchTime(timestamp)
	})
}

func Fuzz_Nosy_Metrics_RecordWithdrawalRequests__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var delayedWeth common.Address
		fill_err = tp.Fill(&delayedWeth)
		if fill_err != nil {
			return
		}
		var matches bool
		fill_err = tp.Fill(&matches)
		if fill_err != nil {
			return
		}
		var count int
		fill_err = tp.Fill(&count)
		if fill_err != nil {
			return
		}

		m := NewMetrics()
		m.RecordWithdrawalRequests(delayedWeth, matches, count)
	})
}

func Fuzz_Nosy_Metrics_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, host string, port int) {
		m := NewMetrics()
		m.Start(host, port)
	})
}

// skipping Fuzz_Nosy_Metrics_StartBalanceMetrics__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/log.Logger

func Fuzz_Nosy_NoopMetricsImpl_CacheAdd__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *NoopMetricsImpl
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var _x2 string
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var _x3 int
		fill_err = tp.Fill(&_x3)
		if fill_err != nil {
			return
		}
		var _x4 bool
		fill_err = tp.Fill(&_x4)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.CacheAdd(_x2, _x3, _x4)
	})
}

func Fuzz_Nosy_NoopMetricsImpl_CacheGet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *NoopMetricsImpl
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var _x2 string
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var _x3 bool
		fill_err = tp.Fill(&_x3)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.CacheGet(_x2, _x3)
	})
}

func Fuzz_Nosy_NoopMetricsImpl_RecordBondCollateral__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *NoopMetricsImpl
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var _x2 common.Address
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var _x3 *big.Int
		fill_err = tp.Fill(&_x3)
		if fill_err != nil {
			return
		}
		var _x4 *big.Int
		fill_err = tp.Fill(&_x4)
		if fill_err != nil {
			return
		}
		if _x1 == nil || _x3 == nil || _x4 == nil {
			return
		}

		_x1.RecordBondCollateral(_x2, _x3, _x4)
	})
}

func Fuzz_Nosy_NoopMetricsImpl_RecordClaims__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *NoopMetricsImpl
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var _x2 *ClaimStatuses
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if _x1 == nil || _x2 == nil {
			return
		}

		_x1.RecordClaims(_x2)
	})
}

func Fuzz_Nosy_NoopMetricsImpl_RecordCredit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *NoopMetricsImpl
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var _x2 CreditExpectation
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var _x3 int
		fill_err = tp.Fill(&_x3)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.RecordCredit(_x2, _x3)
	})
}

func Fuzz_Nosy_NoopMetricsImpl_RecordFailedGames__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *NoopMetricsImpl
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var _x2 int
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.RecordFailedGames(_x2)
	})
}

func Fuzz_Nosy_NoopMetricsImpl_RecordGameAgreement__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *NoopMetricsImpl
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var _x2 GameAgreementStatus
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var _x3 int
		fill_err = tp.Fill(&_x3)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.RecordGameAgreement(_x2, _x3)
	})
}

func Fuzz_Nosy_NoopMetricsImpl_RecordGameResolutionStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *NoopMetricsImpl
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var _x2 ResolutionStatus
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var _x3 int
		fill_err = tp.Fill(&_x3)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.RecordGameResolutionStatus(_x2, _x3)
	})
}

func Fuzz_Nosy_NoopMetricsImpl_RecordHonestActorClaims__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *NoopMetricsImpl
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var _x2 common.Address
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var _x3 *HonestActorData
		fill_err = tp.Fill(&_x3)
		if fill_err != nil {
			return
		}
		if _x1 == nil || _x3 == nil {
			return
		}

		_x1.RecordHonestActorClaims(_x2, _x3)
	})
}

func Fuzz_Nosy_NoopMetricsImpl_RecordHonestWithdrawableAmounts__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *NoopMetricsImpl
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var _x2 map[common.Address]*big.Int
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.RecordHonestWithdrawableAmounts(_x2)
	})
}

func Fuzz_Nosy_NoopMetricsImpl_RecordIgnoredGames__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *NoopMetricsImpl
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var _x2 int
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.RecordIgnoredGames(_x2)
	})
}

func Fuzz_Nosy_NoopMetricsImpl_RecordInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *NoopMetricsImpl
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var _x2 string
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.RecordInfo(_x2)
	})
}

func Fuzz_Nosy_NoopMetricsImpl_RecordL2Challenges__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *NoopMetricsImpl
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var _x2 bool
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var _x3 int
		fill_err = tp.Fill(&_x3)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.RecordL2Challenges(_x2, _x3)
	})
}

func Fuzz_Nosy_NoopMetricsImpl_RecordLatestProposals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *NoopMetricsImpl
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var _x2 uint64
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var _x3 uint64
		fill_err = tp.Fill(&_x3)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.RecordLatestProposals(_x2, _x3)
	})
}

func Fuzz_Nosy_NoopMetricsImpl_RecordLatestValidProposalL2Block__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *NoopMetricsImpl
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var _x2 uint64
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.RecordLatestValidProposalL2Block(_x2)
	})
}

func Fuzz_Nosy_NoopMetricsImpl_RecordMonitorDuration__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *NoopMetricsImpl
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var _x2 time.Duration
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.RecordMonitorDuration(_x2)
	})
}

func Fuzz_Nosy_NoopMetricsImpl_RecordOldestGameUpdateTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *NoopMetricsImpl
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var _x2 time.Time
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.RecordOldestGameUpdateTime(_x2)
	})
}

func Fuzz_Nosy_NoopMetricsImpl_RecordOutputFetchTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *NoopMetricsImpl
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var _x2 float64
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.RecordOutputFetchTime(_x2)
	})
}

func Fuzz_Nosy_NoopMetricsImpl_RecordUp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *NoopMetricsImpl
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.RecordUp()
	})
}

func Fuzz_Nosy_NoopMetricsImpl_RecordWithdrawalRequests__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *NoopMetricsImpl
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var _x2 common.Address
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var _x3 bool
		fill_err = tp.Fill(&_x3)
		if fill_err != nil {
			return
		}
		var _x4 int
		fill_err = tp.Fill(&_x4)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.RecordWithdrawalRequests(_x2, _x3, _x4)
	})
}

func Fuzz_Nosy_ClaimStatus_AsLabels__(f *testing.F) {
	f.Fuzz(func(t *testing.T, firstHalf bool, clockExpired bool, resolvable bool, resolved bool) {
		s := NewClaimStatus(firstHalf, clockExpired, resolvable, resolved)
		s.AsLabels()
	})
}

func Fuzz_Nosy_ClaimStatus_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, firstHalf bool, clockExpired bool, resolvable bool, resolved bool) {
		s := NewClaimStatus(firstHalf, clockExpired, resolvable, resolved)
		s.String()
	})
}

// skipping Fuzz_Nosy_Metricer_RecordBondCollateral__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-dispute-mon/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordClaims__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-dispute-mon/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordCredit__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-dispute-mon/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordFailedGames__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-dispute-mon/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordGameAgreement__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-dispute-mon/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordGameResolutionStatus__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-dispute-mon/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordHonestActorClaims__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-dispute-mon/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordHonestWithdrawableAmounts__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-dispute-mon/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordIgnoredGames__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-dispute-mon/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordInfo__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-dispute-mon/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordL2Challenges__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-dispute-mon/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordLatestProposals__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-dispute-mon/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordLatestValidProposalL2Block__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-dispute-mon/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordMonitorDuration__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-dispute-mon/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordOldestGameUpdateTime__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-dispute-mon/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordOutputFetchTime__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-dispute-mon/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordUp__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-dispute-mon/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordWithdrawalRequests__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-dispute-mon/metrics.Metricer

func Fuzz_Nosy_labelValuesFor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var status GameAgreementStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}

		labelValuesFor(status)
	})
}

func Fuzz_Nosy_weiToEther__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var wei *big.Int
		fill_err = tp.Fill(&wei)
		if fill_err != nil {
			return
		}
		if wei == nil {
			return
		}

		weiToEther(wei)
	})
}
