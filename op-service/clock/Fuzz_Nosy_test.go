package clock

import (
	"context"
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

func Fuzz_Nosy_AdvancingClock_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var advanceEvery time.Duration
		fill_err = tp.Fill(&advanceEvery)
		if fill_err != nil {
			return
		}

		c := NewAdvancingClock(advanceEvery)
		c.Start()
	})
}

func Fuzz_Nosy_AdvancingClock_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var advanceEvery time.Duration
		fill_err = tp.Fill(&advanceEvery)
		if fill_err != nil {
			return
		}

		c := NewAdvancingClock(advanceEvery)
		c.Stop()
	})
}

func Fuzz_Nosy_AdvancingClock_onTick__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var advanceEvery time.Duration
		fill_err = tp.Fill(&advanceEvery)
		if fill_err != nil {
			return
		}
		var now time.Time
		fill_err = tp.Fill(&now)
		if fill_err != nil {
			return
		}

		c := NewAdvancingClock(advanceEvery)
		c.onTick(now)
	})
}

func Fuzz_Nosy_DeterministicClock_AdvanceTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var now time.Time
		fill_err = tp.Fill(&now)
		if fill_err != nil {
			return
		}
		var d time.Duration
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}

		s := NewDeterministicClock(now)
		s.AdvanceTime(d)
	})
}

func Fuzz_Nosy_DeterministicClock_After__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var now time.Time
		fill_err = tp.Fill(&now)
		if fill_err != nil {
			return
		}
		var d time.Duration
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}

		s := NewDeterministicClock(now)
		s.After(d)
	})
}

// skipping Fuzz_Nosy_DeterministicClock_AfterFunc__ because parameters include func, chan, or unsupported interface: func()

func Fuzz_Nosy_DeterministicClock_NewTicker__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var now time.Time
		fill_err = tp.Fill(&now)
		if fill_err != nil {
			return
		}
		var d time.Duration
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}

		s := NewDeterministicClock(now)
		s.NewTicker(d)
	})
}

func Fuzz_Nosy_DeterministicClock_NewTimer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var now time.Time
		fill_err = tp.Fill(&now)
		if fill_err != nil {
			return
		}
		var d time.Duration
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}

		s := NewDeterministicClock(now)
		s.NewTimer(d)
	})
}

func Fuzz_Nosy_DeterministicClock_Now__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var now time.Time
		fill_err = tp.Fill(&now)
		if fill_err != nil {
			return
		}

		s := NewDeterministicClock(now)
		s.Now()
	})
}

func Fuzz_Nosy_DeterministicClock_Since__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var now time.Time
		fill_err = tp.Fill(&now)
		if fill_err != nil {
			return
		}
		var t2 time.Time
		fill_err = tp.Fill(&t2)
		if fill_err != nil {
			return
		}

		s := NewDeterministicClock(now)
		s.Since(t2)
	})
}

func Fuzz_Nosy_DeterministicClock_SleepCtx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var now time.Time
		fill_err = tp.Fill(&now)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var d time.Duration
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}

		s := NewDeterministicClock(now)
		s.SleepCtx(ctx, d)
	})
}

func Fuzz_Nosy_DeterministicClock_WaitForNewPendingTask__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var now time.Time
		fill_err = tp.Fill(&now)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		s := NewDeterministicClock(now)
		s.WaitForNewPendingTask(ctx)
	})
}

func Fuzz_Nosy_DeterministicClock_WaitForNewPendingTaskWithTimeout__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var now time.Time
		fill_err = tp.Fill(&now)
		if fill_err != nil {
			return
		}
		var timeout time.Duration
		fill_err = tp.Fill(&timeout)
		if fill_err != nil {
			return
		}

		s := NewDeterministicClock(now)
		s.WaitForNewPendingTaskWithTimeout(timeout)
	})
}

// skipping Fuzz_Nosy_DeterministicClock_addPending__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/clock.action

func Fuzz_Nosy_LoopFn_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var lf *LoopFn
		fill_err = tp.Fill(&lf)
		if fill_err != nil {
			return
		}
		if lf == nil {
			return
		}

		lf.Close()
	})
}

func Fuzz_Nosy_LoopFn_work__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var lf *LoopFn
		fill_err = tp.Fill(&lf)
		if fill_err != nil {
			return
		}
		if lf == nil {
			return
		}

		lf.work()
	})
}

func Fuzz_Nosy_SimpleClock_Set__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v time.Time
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}

		c := NewSimpleClock()
		c.Set(v)
	})
}

func Fuzz_Nosy_SimpleClock_SetTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, u uint64) {
		c := NewSimpleClock()
		c.SetTime(u)
	})
}

func Fuzz_Nosy_SystemTicker_Ch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *SystemTicker
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Ch()
	})
}

func Fuzz_Nosy_SystemTimer_Ch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *SystemTimer
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Ch()
	})
}

func Fuzz_Nosy_ticker_Ch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *ticker
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Ch()
	})
}

func Fuzz_Nosy_ticker_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *ticker
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var d time.Duration
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Reset(d)
	})
}

func Fuzz_Nosy_ticker_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *ticker
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Stop()
	})
}

func Fuzz_Nosy_ticker_fire__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *ticker
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var now time.Time
		fill_err = tp.Fill(&now)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.fire(now)
	})
}

func Fuzz_Nosy_ticker_isDue__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *ticker
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var now time.Time
		fill_err = tp.Fill(&now)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.isDue(now)
	})
}

func Fuzz_Nosy_timer_Ch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *timer
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Ch()
	})
}

func Fuzz_Nosy_timer_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *timer
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Stop()
	})
}

func Fuzz_Nosy_timer_fire__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *timer
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var now time.Time
		fill_err = tp.Fill(&now)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.fire(now)
	})
}

func Fuzz_Nosy_timer_isDue__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *timer
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var now time.Time
		fill_err = tp.Fill(&now)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.isDue(now)
	})
}

// skipping Fuzz_Nosy_Clock_After__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/clock.Clock

// skipping Fuzz_Nosy_Clock_AfterFunc__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/clock.Clock

// skipping Fuzz_Nosy_Clock_NewTicker__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/clock.Clock

// skipping Fuzz_Nosy_Clock_NewTimer__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/clock.Clock

// skipping Fuzz_Nosy_Clock_Now__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/clock.Clock

// skipping Fuzz_Nosy_Clock_Since__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/clock.Clock

// skipping Fuzz_Nosy_Clock_SleepCtx__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/clock.Clock

// skipping Fuzz_Nosy_RWClock_Now__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/clock.RWClock

// skipping Fuzz_Nosy_Ticker_Ch__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/clock.Ticker

// skipping Fuzz_Nosy_Ticker_Reset__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/clock.Ticker

// skipping Fuzz_Nosy_Ticker_Stop__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/clock.Ticker

// skipping Fuzz_Nosy_Timer_Ch__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/clock.Timer

// skipping Fuzz_Nosy_Timer_Stop__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/clock.Timer

// skipping Fuzz_Nosy_action_fire__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/clock.action

// skipping Fuzz_Nosy_action_isDue__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/clock.action

func Fuzz_Nosy_systemClock_After__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s systemClock
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var d time.Duration
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}

		s.After(d)
	})
}

// skipping Fuzz_Nosy_systemClock_AfterFunc__ because parameters include func, chan, or unsupported interface: func()

func Fuzz_Nosy_systemClock_NewTicker__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s systemClock
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var d time.Duration
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}

		s.NewTicker(d)
	})
}

func Fuzz_Nosy_systemClock_NewTimer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s systemClock
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var d time.Duration
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}

		s.NewTimer(d)
	})
}

func Fuzz_Nosy_systemClock_Now__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s systemClock
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}

		s.Now()
	})
}

func Fuzz_Nosy_systemClock_Since__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s systemClock
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var t2 time.Time
		fill_err = tp.Fill(&t2)
		if fill_err != nil {
			return
		}

		s.Since(t2)
	})
}

func Fuzz_Nosy_systemClock_SleepCtx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s systemClock
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var d time.Duration
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}

		s.SleepCtx(ctx, d)
	})
}

func Fuzz_Nosy_task_fire__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 task
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var now time.Time
		fill_err = tp.Fill(&now)
		if fill_err != nil {
			return
		}

		t1.fire(now)
	})
}

func Fuzz_Nosy_task_isDue__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 task
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var now time.Time
		fill_err = tp.Fill(&now)
		if fill_err != nil {
			return
		}

		t1.isDue(now)
	})
}

// skipping Fuzz_Nosy_MinCheckedTimestamp__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/clock.RWClock
