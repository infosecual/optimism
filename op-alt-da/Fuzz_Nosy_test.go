package altda

import (
	"context"
	"testing"
	"time"

	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum/go-ethereum/core/types"
	go_fuzz_utils "github.com/trailofbits/go-fuzz-utils"
	"github.com/urfave/cli/v2"
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

// skipping Fuzz_Nosy_AltDADisabled_AdvanceL1Origin__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-alt-da.L1Fetcher

func Fuzz_Nosy_AltDADisabled_Finalize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *AltDADisabled
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var ref eth.L1BlockRef
		fill_err = tp.Fill(&ref)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.Finalize(ref)
	})
}

// skipping Fuzz_Nosy_AltDADisabled_GetInput__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-alt-da.L1Fetcher

// skipping Fuzz_Nosy_AltDADisabled_OnFinalizedHeadSignal__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-alt-da.HeadSignalFn

func Fuzz_Nosy_AltDADisabled_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *AltDADisabled
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var base eth.L1BlockRef
		fill_err = tp.Fill(&base)
		if fill_err != nil {
			return
		}
		var baseCfg eth.SystemConfig
		fill_err = tp.Fill(&baseCfg)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.Reset(ctx, base, baseCfg)
	})
}

func Fuzz_Nosy_Challenge_key__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *Challenge
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.key()
	})
}

// skipping Fuzz_Nosy_DA_AdvanceChallengeOrigin__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-alt-da.L1Fetcher

// skipping Fuzz_Nosy_DA_AdvanceCommitmentOrigin__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-alt-da.L1Fetcher

// skipping Fuzz_Nosy_DA_AdvanceL1Origin__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-alt-da.L1Fetcher

func Fuzz_Nosy_DA_Finalize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *DA
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var l1Finalized eth.L1BlockRef
		fill_err = tp.Fill(&l1Finalized)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.Finalize(l1Finalized)
	})
}

// skipping Fuzz_Nosy_DA_GetInput__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-alt-da.L1Fetcher

// skipping Fuzz_Nosy_DA_LookAhead__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-alt-da.L1Fetcher

// skipping Fuzz_Nosy_DA_OnFinalizedHeadSignal__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-alt-da.HeadSignalFn

func Fuzz_Nosy_DA_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *DA
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var base eth.L1BlockRef
		fill_err = tp.Fill(&base)
		if fill_err != nil {
			return
		}
		var baseCfg eth.SystemConfig
		fill_err = tp.Fill(&baseCfg)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.Reset(ctx, base, baseCfg)
	})
}

func Fuzz_Nosy_DA_decodeChallengeStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *DA
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var log *types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if d == nil || log == nil {
			return
		}

		d.decodeChallengeStatus(log)
	})
}

// skipping Fuzz_Nosy_DA_fetchChallengeLogs__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-alt-da.L1Fetcher

// skipping Fuzz_Nosy_DA_loadChallengeEvents__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-alt-da.L1Fetcher

// skipping Fuzz_Nosy_DA_updateFinalizedFromL1__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-alt-da.L1Fetcher

func Fuzz_Nosy_DA_updateFinalizedHead__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *DA
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var l1Finalized eth.L1BlockRef
		fill_err = tp.Fill(&l1Finalized)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.updateFinalizedHead(l1Finalized)
	})
}

// skipping Fuzz_Nosy_DAClient_GetInput__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-alt-da.CommitmentData

func Fuzz_Nosy_DAClient_SetInput__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var url string
		fill_err = tp.Fill(&url)
		if fill_err != nil {
			return
		}
		var verify bool
		fill_err = tp.Fill(&verify)
		if fill_err != nil {
			return
		}
		var pc bool
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var img []byte
		fill_err = tp.Fill(&img)
		if fill_err != nil {
			return
		}

		c := NewDAClient(url, verify, pc)
		c.SetInput(ctx, img)
	})
}

func Fuzz_Nosy_DAClient_setInput__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var url string
		fill_err = tp.Fill(&url)
		if fill_err != nil {
			return
		}
		var verify bool
		fill_err = tp.Fill(&verify)
		if fill_err != nil {
			return
		}
		var pc bool
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var img []byte
		fill_err = tp.Fill(&img)
		if fill_err != nil {
			return
		}

		c := NewDAClient(url, verify, pc)
		c.setInput(ctx, img)
	})
}

// skipping Fuzz_Nosy_DAClient_setInputWithCommit__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-alt-da.CommitmentData

func Fuzz_Nosy_DAErrFaker_ActGetPreImageFail__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *DAErrFaker
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.ActGetPreImageFail()
	})
}

func Fuzz_Nosy_DAErrFaker_ActSetPreImageFail__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *DAErrFaker
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.ActSetPreImageFail()
	})
}

// skipping Fuzz_Nosy_DAErrFaker_GetInput__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-alt-da.CommitmentData

func Fuzz_Nosy_DAErrFaker_SetInput__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *DAErrFaker
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var d3 []byte
		fill_err = tp.Fill(&d3)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.SetInput(ctx, d3)
	})
}

// skipping Fuzz_Nosy_DAServer_HandleGet__ because parameters include func, chan, or unsupported interface: net/http.ResponseWriter

// skipping Fuzz_Nosy_DAServer_HandlePut__ because parameters include func, chan, or unsupported interface: net/http.ResponseWriter

func Fuzz_Nosy_DAServer_HttpEndpoint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *DAServer
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.HttpEndpoint()
	})
}

func Fuzz_Nosy_DAServer_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *DAServer
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.Start()
	})
}

func Fuzz_Nosy_DAServer_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *DAServer
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.Stop()
	})
}

// skipping Fuzz_Nosy_FakeDAServer_HandleGet__ because parameters include func, chan, or unsupported interface: net/http.ResponseWriter

// skipping Fuzz_Nosy_FakeDAServer_HandlePut__ because parameters include func, chan, or unsupported interface: net/http.ResponseWriter

func Fuzz_Nosy_FakeDAServer_SetGetRequestLatency__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *FakeDAServer
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var latency time.Duration
		fill_err = tp.Fill(&latency)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.SetGetRequestLatency(latency)
	})
}

func Fuzz_Nosy_FakeDAServer_SetPutRequestLatency__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *FakeDAServer
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var latency time.Duration
		fill_err = tp.Fill(&latency)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.SetPutRequestLatency(latency)
	})
}

func Fuzz_Nosy_FakeDAServer_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *FakeDAServer
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

func Fuzz_Nosy_MemStore_Get__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var key []byte
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}

		s := NewMemStore()
		s.Get(ctx, key)
	})
}

func Fuzz_Nosy_MemStore_Put__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var key []byte
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var value []byte
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}

		s := NewMemStore()
		s.Put(ctx, key, value)
	})
}

func Fuzz_Nosy_Metrics_RecordActiveChallenge__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *Metrics
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var commBlock uint64
		fill_err = tp.Fill(&commBlock)
		if fill_err != nil {
			return
		}
		var startBlock uint64
		fill_err = tp.Fill(&startBlock)
		if fill_err != nil {
			return
		}
		var hash []byte
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.RecordActiveChallenge(commBlock, startBlock, hash)
	})
}

func Fuzz_Nosy_Metrics_RecordChallenge__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *Metrics
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var status string
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.RecordChallenge(status)
	})
}

func Fuzz_Nosy_Metrics_RecordChallengesHead__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *Metrics
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var num uint64
		fill_err = tp.Fill(&num)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.RecordChallengesHead(name, num)
	})
}

func Fuzz_Nosy_Metrics_RecordExpiredChallenge__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *Metrics
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var hash []byte
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.RecordExpiredChallenge(hash)
	})
}

func Fuzz_Nosy_Metrics_RecordResolvedChallenge__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *Metrics
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var hash []byte
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.RecordResolvedChallenge(hash)
	})
}

func Fuzz_Nosy_Metrics_RecordStorageError__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *Metrics
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.RecordStorageError()
	})
}

func Fuzz_Nosy_MockDAClient_DeleteData__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *MockDAClient
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var key []byte
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.DeleteData(key)
	})
}

// skipping Fuzz_Nosy_MockDAClient_GetInput__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-alt-da.CommitmentData

func Fuzz_Nosy_MockDAClient_SetInput__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *MockDAClient
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var d3 []byte
		fill_err = tp.Fill(&d3)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.SetInput(ctx, d3)
	})
}

func Fuzz_Nosy_NoopMetrics_RecordActiveChallenge__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *NoopMetrics
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var commBlock uint64
		fill_err = tp.Fill(&commBlock)
		if fill_err != nil {
			return
		}
		var startBlock uint64
		fill_err = tp.Fill(&startBlock)
		if fill_err != nil {
			return
		}
		var hash []byte
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.RecordActiveChallenge(commBlock, startBlock, hash)
	})
}

func Fuzz_Nosy_NoopMetrics_RecordChallengesHead__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *NoopMetrics
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var num uint64
		fill_err = tp.Fill(&num)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.RecordChallengesHead(name, num)
	})
}

func Fuzz_Nosy_NoopMetrics_RecordExpiredChallenge__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *NoopMetrics
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var hash []byte
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.RecordExpiredChallenge(hash)
	})
}

func Fuzz_Nosy_NoopMetrics_RecordResolvedChallenge__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *NoopMetrics
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var hash []byte
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.RecordResolvedChallenge(hash)
	})
}

func Fuzz_Nosy_NoopMetrics_RecordStorageError__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *NoopMetrics
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.RecordStorageError()
	})
}

func Fuzz_Nosy_State_ClearCommitments__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *State
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.ClearCommitments()
	})
}

// skipping Fuzz_Nosy_State_CreateChallenge__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-alt-da.CommitmentData

func Fuzz_Nosy_State_ExpireChallenges__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *State
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var origin eth.BlockID
		fill_err = tp.Fill(&origin)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.ExpireChallenges(origin)
	})
}

func Fuzz_Nosy_State_ExpireCommitments__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *State
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var origin eth.BlockID
		fill_err = tp.Fill(&origin)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.ExpireCommitments(origin)
	})
}

// skipping Fuzz_Nosy_State_GetChallenge__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-alt-da.CommitmentData

// skipping Fuzz_Nosy_State_GetChallengeStatus__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-alt-da.CommitmentData

func Fuzz_Nosy_State_NoCommitments__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *State
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.NoCommitments()
	})
}

func Fuzz_Nosy_State_Prune__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *State
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var origin eth.BlockID
		fill_err = tp.Fill(&origin)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Prune(origin)
	})
}

func Fuzz_Nosy_State_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *State
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Reset()
	})
}

// skipping Fuzz_Nosy_State_ResolveChallenge__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-alt-da.CommitmentData

// skipping Fuzz_Nosy_State_TrackCommitment__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-alt-da.CommitmentData

func Fuzz_Nosy_State_pruneChallenges__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *State
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var origin eth.BlockID
		fill_err = tp.Fill(&origin)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.pruneChallenges(origin)
	})
}

func Fuzz_Nosy_State_pruneCommitments__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *State
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var origin eth.BlockID
		fill_err = tp.Fill(&origin)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.pruneCommitments(origin)
	})
}

func Fuzz_Nosy_CLIConfig_Check__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *cli.Context
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c1 := ReadCLIConfig(c)
		c1.Check()
	})
}

func Fuzz_Nosy_CLIConfig_NewDAClient__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *cli.Context
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c1 := ReadCLIConfig(c)
		c1.NewDAClient()
	})
}

func Fuzz_Nosy_CommitmentData_CommitmentType__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 CommitmentType
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}

		_x1 := NewCommitmentData(t1, input)
		_x1.CommitmentType()
	})
}

func Fuzz_Nosy_CommitmentData_Encode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 CommitmentType
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}

		_x1 := NewCommitmentData(t1, input)
		_x1.Encode()
	})
}

func Fuzz_Nosy_CommitmentData_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 CommitmentType
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}

		_x1 := NewCommitmentData(t1, input)
		_x1.String()
	})
}

func Fuzz_Nosy_CommitmentData_TxData__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 CommitmentType
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}

		_x1 := NewCommitmentData(t1, input)
		_x1.TxData()
	})
}

func Fuzz_Nosy_CommitmentData_Verify__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 CommitmentType
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var i2 []byte
		fill_err = tp.Fill(&i2)
		if fill_err != nil {
			return
		}
		var i3 []byte
		fill_err = tp.Fill(&i3)
		if fill_err != nil {
			return
		}

		_x1 := NewCommitmentData(t1, i2)
		_x1.Verify(i3)
	})
}

// skipping Fuzz_Nosy_DAStorage_GetInput__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-alt-da.DAStorage

// skipping Fuzz_Nosy_DAStorage_SetInput__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-alt-da.DAStorage

func Fuzz_Nosy_GenericCommitment_CommitmentType__(f *testing.F) {
	f.Fuzz(func(t *testing.T, commitment []byte) {
		c, err := DecodeGenericCommitment(commitment)
		if err != nil {
			return
		}
		c.CommitmentType()
	})
}

func Fuzz_Nosy_GenericCommitment_Encode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, commitment []byte) {
		c, err := DecodeGenericCommitment(commitment)
		if err != nil {
			return
		}
		c.Encode()
	})
}

func Fuzz_Nosy_GenericCommitment_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, commitment []byte) {
		c, err := DecodeGenericCommitment(commitment)
		if err != nil {
			return
		}
		c.String()
	})
}

func Fuzz_Nosy_GenericCommitment_TxData__(f *testing.F) {
	f.Fuzz(func(t *testing.T, commitment []byte) {
		c, err := DecodeGenericCommitment(commitment)
		if err != nil {
			return
		}
		c.TxData()
	})
}

func Fuzz_Nosy_GenericCommitment_Verify__(f *testing.F) {
	f.Fuzz(func(t *testing.T, commitment []byte, input []byte) {
		c, err := DecodeGenericCommitment(commitment)
		if err != nil {
			return
		}
		c.Verify(input)
	})
}

// skipping Fuzz_Nosy_KVStore_Get__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-alt-da.KVStore

// skipping Fuzz_Nosy_KVStore_Put__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-alt-da.KVStore

func Fuzz_Nosy_Keccak256Commitment_CommitmentType__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input []byte) {
		c := NewKeccak256Commitment(input)
		c.CommitmentType()
	})
}

func Fuzz_Nosy_Keccak256Commitment_Encode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input []byte) {
		c := NewKeccak256Commitment(input)
		c.Encode()
	})
}

func Fuzz_Nosy_Keccak256Commitment_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input []byte) {
		c := NewKeccak256Commitment(input)
		c.String()
	})
}

func Fuzz_Nosy_Keccak256Commitment_TxData__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input []byte) {
		c := NewKeccak256Commitment(input)
		c.TxData()
	})
}

func Fuzz_Nosy_Keccak256Commitment_Verify__(f *testing.F) {
	f.Fuzz(func(t *testing.T, i1 []byte, i2 []byte) {
		c := NewKeccak256Commitment(i1)
		c.Verify(i2)
	})
}

// skipping Fuzz_Nosy_L1Fetcher_FetchReceipts__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-alt-da.L1Fetcher

// skipping Fuzz_Nosy_L1Fetcher_InfoAndTxsByHash__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-alt-da.L1Fetcher

// skipping Fuzz_Nosy_L1Fetcher_L1BlockRefByNumber__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-alt-da.L1Fetcher

// skipping Fuzz_Nosy_Metricer_RecordActiveChallenge__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-alt-da.Metricer

// skipping Fuzz_Nosy_Metricer_RecordChallengesHead__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-alt-da.Metricer

// skipping Fuzz_Nosy_Metricer_RecordExpiredChallenge__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-alt-da.Metricer

// skipping Fuzz_Nosy_Metricer_RecordResolvedChallenge__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-alt-da.Metricer

// skipping Fuzz_Nosy_Metricer_RecordStorageError__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-alt-da.Metricer

func Fuzz_Nosy_CLIFlags__(f *testing.F) {
	f.Fuzz(func(t *testing.T, envPrefix string, category string) {
		CLIFlags(envPrefix, category)
	})
}

func Fuzz_Nosy_DecodeResolvedInput__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d1 []byte) {
		DecodeResolvedInput(d1)
	})
}

func Fuzz_Nosy_altDAEnvs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, envprefix string, v string) {
		altDAEnvs(envprefix, v)
	})
}

func Fuzz_Nosy_altDAFlags__(f *testing.F) {
	f.Fuzz(func(t *testing.T, v string) {
		altDAFlags(v)
	})
}

// skipping Fuzz_Nosy_challengeKey__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-alt-da.CommitmentData
