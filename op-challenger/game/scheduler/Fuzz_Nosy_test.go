package scheduler

import (
	"context"
	"testing"

	"github.com/ethereum-optimism/optimism/op-challenger/game/types"
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

func Fuzz_Nosy_Scheduler_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Scheduler
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Close()
	})
}

func Fuzz_Nosy_Scheduler_Schedule__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Scheduler
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var games []types.GameMetadata
		fill_err = tp.Fill(&games)
		if fill_err != nil {
			return
		}
		var blockNumber uint64
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Schedule(games, blockNumber)
	})
}

func Fuzz_Nosy_Scheduler_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Scheduler
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

func Fuzz_Nosy_Scheduler_ThreadActive__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Scheduler
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.ThreadActive()
	})
}

func Fuzz_Nosy_Scheduler_ThreadIdle__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Scheduler
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.ThreadIdle()
	})
}

func Fuzz_Nosy_Scheduler_loop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Scheduler
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

		s.loop(ctx)
	})
}

func Fuzz_Nosy_coordinator_createJob__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *coordinator
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var game types.GameMetadata
		fill_err = tp.Fill(&game)
		if fill_err != nil {
			return
		}
		var blockNumber uint64
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.createJob(ctx, game, blockNumber)
	})
}

func Fuzz_Nosy_coordinator_deleteResolvedGameFiles__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *coordinator
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.deleteResolvedGameFiles()
	})
}

func Fuzz_Nosy_coordinator_enqueueJob__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *coordinator
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var j job
		fill_err = tp.Fill(&j)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.enqueueJob(ctx, j)
	})
}

func Fuzz_Nosy_coordinator_processResult__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *coordinator
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var j job
		fill_err = tp.Fill(&j)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.processResult(j)
	})
}

func Fuzz_Nosy_coordinator_schedule__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *coordinator
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var games []types.GameMetadata
		fill_err = tp.Fill(&games)
		if fill_err != nil {
			return
		}
		var blockNumber uint64
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.schedule(ctx, games, blockNumber)
	})
}

// skipping Fuzz_Nosy_CoordinatorMetricer_RecordActedL1Block__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/scheduler.CoordinatorMetricer

// skipping Fuzz_Nosy_CoordinatorMetricer_RecordGameUpdateCompleted__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/scheduler.CoordinatorMetricer

// skipping Fuzz_Nosy_CoordinatorMetricer_RecordGameUpdateScheduled__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/scheduler.CoordinatorMetricer

// skipping Fuzz_Nosy_CoordinatorMetricer_RecordGamesStatus__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/scheduler.CoordinatorMetricer

// skipping Fuzz_Nosy_DiskManager_DirForGame__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/scheduler.DiskManager

// skipping Fuzz_Nosy_DiskManager_RemoveAllExcept__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/scheduler.DiskManager

// skipping Fuzz_Nosy_GamePlayer_ProgressGame__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/scheduler.GamePlayer

// skipping Fuzz_Nosy_GamePlayer_Status__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/scheduler.GamePlayer

// skipping Fuzz_Nosy_GamePlayer_ValidatePrestate__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/scheduler.GamePlayer

// skipping Fuzz_Nosy_SchedulerMetricer_DecActiveExecutors__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/scheduler.SchedulerMetricer

// skipping Fuzz_Nosy_SchedulerMetricer_DecIdleExecutors__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/scheduler.SchedulerMetricer

// skipping Fuzz_Nosy_SchedulerMetricer_IncActiveExecutors__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/scheduler.SchedulerMetricer

// skipping Fuzz_Nosy_SchedulerMetricer_IncIdleExecutors__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/scheduler.SchedulerMetricer

// skipping Fuzz_Nosy_SchedulerMetricer_RecordActedL1Block__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/scheduler.SchedulerMetricer

// skipping Fuzz_Nosy_SchedulerMetricer_RecordGameUpdateCompleted__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/scheduler.SchedulerMetricer

// skipping Fuzz_Nosy_SchedulerMetricer_RecordGameUpdateScheduled__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/scheduler.SchedulerMetricer

// skipping Fuzz_Nosy_SchedulerMetricer_RecordGamesStatus__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-challenger/game/scheduler.SchedulerMetricer

// skipping Fuzz_Nosy_progressGames__ because parameters include func, chan, or unsupported interface: <-chan github.com/ethereum-optimism/optimism/op-challenger/game/scheduler.job
