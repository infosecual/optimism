package event

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

// skipping Fuzz_Nosy_DeriverMux_OnEvent__ because parameters include func, chan, or unsupported interface: *github.com/ethereum-optimism/optimism/op-node/rollup/event.DeriverMux

// skipping Fuzz_Nosy_GlobalSyncExec_Add__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/event.Executable

func Fuzz_Nosy_GlobalSyncExec_Drain__(f *testing.F) {
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

		gs := NewGlobalSynchronous(ctx)
		gs.Drain()
	})
}

// skipping Fuzz_Nosy_GlobalSyncExec_DrainUntil__ because parameters include func, chan, or unsupported interface: func(ev github.com/ethereum-optimism/optimism/op-node/rollup/event.Event) bool

func Fuzz_Nosy_GlobalSyncExec_Enqueue__(f *testing.F) {
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
		var ev AnnotatedEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}

		gs := NewGlobalSynchronous(ctx)
		gs.Enqueue(ev)
	})
}

func Fuzz_Nosy_GlobalSyncExec_pop__(f *testing.F) {
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

		gs := NewGlobalSynchronous(ctx)
		gs.pop()
	})
}

func Fuzz_Nosy_GlobalSyncExec_processEvent__(f *testing.F) {
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
		var ev AnnotatedEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}

		gs := NewGlobalSynchronous(ctx)
		gs.processEvent(ev)
	})
}

func Fuzz_Nosy_GlobalSyncExec_remove__(f *testing.F) {
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
		var h *globalHandle
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		gs := NewGlobalSynchronous(ctx)
		gs.remove(h)
	})
}

func Fuzz_Nosy_LimiterDrainer_Drain__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *LimiterDrainer
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		l.Drain()
	})
}

// skipping Fuzz_Nosy_LimiterDrainer_DrainUntil__ because parameters include func, chan, or unsupported interface: func(ev github.com/ethereum-optimism/optimism/op-node/rollup/event.Event) bool

// skipping Fuzz_Nosy_LimiterDrainer_Emit__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/event.Event

// skipping Fuzz_Nosy_Limiter[E Emitter]_Emit__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/event.Event

func Fuzz_Nosy_LogTracer_OnDeriveEnd__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var lt *LogTracer
		fill_err = tp.Fill(&lt)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var ev AnnotatedEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}
		var derivContext uint64
		fill_err = tp.Fill(&derivContext)
		if fill_err != nil {
			return
		}
		var startTime time.Time
		fill_err = tp.Fill(&startTime)
		if fill_err != nil {
			return
		}
		var duration time.Duration
		fill_err = tp.Fill(&duration)
		if fill_err != nil {
			return
		}
		var effect bool
		fill_err = tp.Fill(&effect)
		if fill_err != nil {
			return
		}
		if lt == nil {
			return
		}

		lt.OnDeriveEnd(name, ev, derivContext, startTime, duration, effect)
	})
}

func Fuzz_Nosy_LogTracer_OnDeriveStart__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var lt *LogTracer
		fill_err = tp.Fill(&lt)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var ev AnnotatedEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}
		var derivContext uint64
		fill_err = tp.Fill(&derivContext)
		if fill_err != nil {
			return
		}
		var startTime time.Time
		fill_err = tp.Fill(&startTime)
		if fill_err != nil {
			return
		}
		if lt == nil {
			return
		}

		lt.OnDeriveStart(name, ev, derivContext, startTime)
	})
}

func Fuzz_Nosy_LogTracer_OnEmit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var lt *LogTracer
		fill_err = tp.Fill(&lt)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var ev AnnotatedEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}
		var derivContext uint64
		fill_err = tp.Fill(&derivContext)
		if fill_err != nil {
			return
		}
		var emitTime time.Time
		fill_err = tp.Fill(&emitTime)
		if fill_err != nil {
			return
		}
		if lt == nil {
			return
		}

		lt.OnEmit(name, ev, derivContext, emitTime)
	})
}

func Fuzz_Nosy_LogTracer_OnRateLimited__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var lt *LogTracer
		fill_err = tp.Fill(&lt)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var derivContext uint64
		fill_err = tp.Fill(&derivContext)
		if fill_err != nil {
			return
		}
		if lt == nil {
			return
		}

		lt.OnRateLimited(name, derivContext)
	})
}

func Fuzz_Nosy_MetricsTracer_OnDeriveEnd__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var mt *MetricsTracer
		fill_err = tp.Fill(&mt)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var ev AnnotatedEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}
		var derivContext uint64
		fill_err = tp.Fill(&derivContext)
		if fill_err != nil {
			return
		}
		var startTime time.Time
		fill_err = tp.Fill(&startTime)
		if fill_err != nil {
			return
		}
		var duration time.Duration
		fill_err = tp.Fill(&duration)
		if fill_err != nil {
			return
		}
		var effect bool
		fill_err = tp.Fill(&effect)
		if fill_err != nil {
			return
		}
		if mt == nil {
			return
		}

		mt.OnDeriveEnd(name, ev, derivContext, startTime, duration, effect)
	})
}

func Fuzz_Nosy_MetricsTracer_OnDeriveStart__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var mt *MetricsTracer
		fill_err = tp.Fill(&mt)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var ev AnnotatedEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}
		var derivContext uint64
		fill_err = tp.Fill(&derivContext)
		if fill_err != nil {
			return
		}
		var startTime time.Time
		fill_err = tp.Fill(&startTime)
		if fill_err != nil {
			return
		}
		if mt == nil {
			return
		}

		mt.OnDeriveStart(name, ev, derivContext, startTime)
	})
}

func Fuzz_Nosy_MetricsTracer_OnEmit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var mt *MetricsTracer
		fill_err = tp.Fill(&mt)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var ev AnnotatedEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}
		var derivContext uint64
		fill_err = tp.Fill(&derivContext)
		if fill_err != nil {
			return
		}
		var emitTime time.Time
		fill_err = tp.Fill(&emitTime)
		if fill_err != nil {
			return
		}
		if mt == nil {
			return
		}

		mt.OnEmit(name, ev, derivContext, emitTime)
	})
}

func Fuzz_Nosy_MetricsTracer_OnRateLimited__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var mt *MetricsTracer
		fill_err = tp.Fill(&mt)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var derivContext uint64
		fill_err = tp.Fill(&derivContext)
		if fill_err != nil {
			return
		}
		if mt == nil {
			return
		}

		mt.OnRateLimited(name, derivContext)
	})
}

func Fuzz_Nosy_SequenceTracer_Output__(f *testing.F) {
	f.Fuzz(func(t *testing.T, showDurations bool) {
		st := NewSequenceTracer()
		st.Output(showDurations)
	})
}

func Fuzz_Nosy_StructTracer_OnDeriveEnd__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var ev AnnotatedEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}
		var derivContext uint64
		fill_err = tp.Fill(&derivContext)
		if fill_err != nil {
			return
		}
		var startTime time.Time
		fill_err = tp.Fill(&startTime)
		if fill_err != nil {
			return
		}
		var duration time.Duration
		fill_err = tp.Fill(&duration)
		if fill_err != nil {
			return
		}
		var effect bool
		fill_err = tp.Fill(&effect)
		if fill_err != nil {
			return
		}

		st := NewStructTracer()
		st.OnDeriveEnd(name, ev, derivContext, startTime, duration, effect)
	})
}

func Fuzz_Nosy_StructTracer_OnDeriveStart__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var ev AnnotatedEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}
		var derivContext uint64
		fill_err = tp.Fill(&derivContext)
		if fill_err != nil {
			return
		}
		var startTime time.Time
		fill_err = tp.Fill(&startTime)
		if fill_err != nil {
			return
		}

		st := NewStructTracer()
		st.OnDeriveStart(name, ev, derivContext, startTime)
	})
}

func Fuzz_Nosy_StructTracer_OnEmit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var ev AnnotatedEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}
		var derivContext uint64
		fill_err = tp.Fill(&derivContext)
		if fill_err != nil {
			return
		}
		var emitTime time.Time
		fill_err = tp.Fill(&emitTime)
		if fill_err != nil {
			return
		}

		st := NewStructTracer()
		st.OnEmit(name, ev, derivContext, emitTime)
	})
}

func Fuzz_Nosy_StructTracer_OnRateLimited__(f *testing.F) {
	f.Fuzz(func(t *testing.T, name string, derivContext uint64) {
		st := NewStructTracer()
		st.OnRateLimited(name, derivContext)
	})
}

// skipping Fuzz_Nosy_Sys_AddTracer__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/event.Tracer

// skipping Fuzz_Nosy_Sys_Register__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/event.Deriver

// skipping Fuzz_Nosy_Sys_RemoveTracer__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/event.Tracer

func Fuzz_Nosy_Sys_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Sys
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Stop()
	})
}

func Fuzz_Nosy_Sys_Unregister__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Sys
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Unregister(name)
	})
}

// skipping Fuzz_Nosy_Sys_emit__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/event.Event

func Fuzz_Nosy_Sys_recordDerivEnd__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Sys
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var ev AnnotatedEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}
		var derivContext uint64
		fill_err = tp.Fill(&derivContext)
		if fill_err != nil {
			return
		}
		var startTime time.Time
		fill_err = tp.Fill(&startTime)
		if fill_err != nil {
			return
		}
		var duration time.Duration
		fill_err = tp.Fill(&duration)
		if fill_err != nil {
			return
		}
		var effect bool
		fill_err = tp.Fill(&effect)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.recordDerivEnd(name, ev, derivContext, startTime, duration, effect)
	})
}

func Fuzz_Nosy_Sys_recordDerivStart__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Sys
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var ev AnnotatedEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}
		var startTime time.Time
		fill_err = tp.Fill(&startTime)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.recordDerivStart(name, ev, startTime)
	})
}

func Fuzz_Nosy_Sys_recordEmit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Sys
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var ev AnnotatedEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}
		var derivContext uint64
		fill_err = tp.Fill(&derivContext)
		if fill_err != nil {
			return
		}
		var emitTime time.Time
		fill_err = tp.Fill(&emitTime)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.recordEmit(name, ev, derivContext, emitTime)
	})
}

func Fuzz_Nosy_Sys_recordRateLimited__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Sys
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var derivContext uint64
		fill_err = tp.Fill(&derivContext)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.recordRateLimited(name, derivContext)
	})
}

func Fuzz_Nosy_Sys_unregister__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Sys
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.unregister(name)
	})
}

func Fuzz_Nosy_globalHandle_leave__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var gh *globalHandle
		fill_err = tp.Fill(&gh)
		if fill_err != nil {
			return
		}
		if gh == nil {
			return
		}

		gh.leave()
	})
}

func Fuzz_Nosy_globalHandle_onEvent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var gh *globalHandle
		fill_err = tp.Fill(&gh)
		if fill_err != nil {
			return
		}
		var ev AnnotatedEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}
		if gh == nil {
			return
		}

		gh.onEvent(ev)
	})
}

// skipping Fuzz_Nosy_systemActor_Emit__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/event.Event

func Fuzz_Nosy_systemActor_RunEvent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *systemActor
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var ev AnnotatedEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.RunEvent(ev)
	})
}

// skipping Fuzz_Nosy_AttachEmitter_AttachEmitter__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/event.AttachEmitter

func Fuzz_Nosy_CriticalErrorEvent_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ev CriticalErrorEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}

		ev.String()
	})
}

// skipping Fuzz_Nosy_DebugDeriver_OnEvent__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/event.Event

// skipping Fuzz_Nosy_Deriver_OnEvent__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/event.Deriver

// skipping Fuzz_Nosy_DeriverFunc_OnEvent__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/event.DeriverFunc

// skipping Fuzz_Nosy_Drainer_Drain__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/event.Drainer

// skipping Fuzz_Nosy_Drainer_DrainUntil__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/event.Drainer

// skipping Fuzz_Nosy_Emitter_Emit__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/event.Emitter

// skipping Fuzz_Nosy_EmitterFunc_Emit__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/event.EmitterFunc

// skipping Fuzz_Nosy_Event_String__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/event.Event

// skipping Fuzz_Nosy_Executable_RunEvent__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/event.Executable

// skipping Fuzz_Nosy_ExecutableFunc_RunEvent__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/event.ExecutableFunc

// skipping Fuzz_Nosy_Executor_Add__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/event.Executor

// skipping Fuzz_Nosy_Executor_Enqueue__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/event.Executor

// skipping Fuzz_Nosy_Metrics_RecordEmittedEvent__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/event.Metrics

// skipping Fuzz_Nosy_Metrics_RecordEventsRateLimited__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/event.Metrics

// skipping Fuzz_Nosy_Metrics_RecordProcessedEvent__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/event.Metrics

// skipping Fuzz_Nosy_NoopDeriver_OnEvent__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/event.Event

// skipping Fuzz_Nosy_NoopEmitter_Emit__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/event.Event

func Fuzz_Nosy_NoopMetrics_RecordEmittedEvent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n NoopMetrics
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var eventName string
		fill_err = tp.Fill(&eventName)
		if fill_err != nil {
			return
		}
		var emitter string
		fill_err = tp.Fill(&emitter)
		if fill_err != nil {
			return
		}

		n.RecordEmittedEvent(eventName, emitter)
	})
}

func Fuzz_Nosy_NoopMetrics_RecordEventsRateLimited__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n NoopMetrics
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}

		n.RecordEventsRateLimited()
	})
}

func Fuzz_Nosy_NoopMetrics_RecordProcessedEvent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n NoopMetrics
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var eventName string
		fill_err = tp.Fill(&eventName)
		if fill_err != nil {
			return
		}
		var deriver string
		fill_err = tp.Fill(&deriver)
		if fill_err != nil {
			return
		}
		var duration time.Duration
		fill_err = tp.Fill(&duration)
		if fill_err != nil {
			return
		}

		n.RecordProcessedEvent(eventName, deriver, duration)
	})
}

// skipping Fuzz_Nosy_Registry_Register__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/event.Registry

// skipping Fuzz_Nosy_Registry_Unregister__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/event.Registry

// skipping Fuzz_Nosy_System_AddTracer__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/event.System

// skipping Fuzz_Nosy_System_RemoveTracer__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/event.System

// skipping Fuzz_Nosy_System_Stop__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/event.System

// skipping Fuzz_Nosy_Tracer_OnDeriveEnd__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/event.Tracer

// skipping Fuzz_Nosy_Tracer_OnDeriveStart__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/event.Tracer

// skipping Fuzz_Nosy_Tracer_OnEmit__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/event.Tracer

// skipping Fuzz_Nosy_Tracer_OnRateLimited__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/event.Tracer

// skipping Fuzz_Nosy_Any__ because parameters include func, chan, or unsupported interface: []func(ev github.com/ethereum-optimism/optimism/op-node/rollup/event.Event) bool

// skipping Fuzz_Nosy_Is__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/event.Event
