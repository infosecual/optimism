package status

import (
	"testing"
	"github.com/ethereum-optimism/optimism/op-node/rollup"
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
	
	func Fuzz_Nosy_L1Tracker_L1BlockRefByNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var l *L1Tracker
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var num uint64
		fill_err = tp.Fill(&num)
		if fill_err != nil {
			return
		}
	if l == nil {
		return
	}

	l.L1BlockRefByNumber(ctx, num)
	})
}

// skipping Fuzz_Nosy_L1Tracker_OnEvent__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/event.Event

func Fuzz_Nosy_StatusTracker_L1Head__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var st *StatusTracker
		fill_err = tp.Fill(&st)
		if fill_err != nil {
			return
		}
	if st == nil {
		return
	}

	st.L1Head()
	})
}

// skipping Fuzz_Nosy_StatusTracker_OnEvent__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/event.Event

func Fuzz_Nosy_StatusTracker_SyncStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var st *StatusTracker
		fill_err = tp.Fill(&st)
		if fill_err != nil {
			return
		}
	if st == nil {
		return
	}

	st.SyncStatus()
	})
}

func Fuzz_Nosy_l1HeadBuffer_Get__(f *testing.F) {
	f.Fuzz(func(t *testing.T, size int, num uint64) {
	lhb := newL1HeadBuffer(size)
	lhb.Get(num)
	})
}

func Fuzz_Nosy_l1HeadBuffer_Insert__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var size int
		fill_err = tp.Fill(&size)
		if fill_err != nil {
			return
		}
		var l1Head eth.L1BlockRef
		fill_err = tp.Fill(&l1Head)
		if fill_err != nil {
			return
		}

	lhb := newL1HeadBuffer(size)
	lhb.Insert(l1Head)
	})
}

func Fuzz_Nosy_l1HeadBuffer_get__(f *testing.F) {
	f.Fuzz(func(t *testing.T, size int, num uint64) {
	lhb := newL1HeadBuffer(size)
	lhb.get(num)
	})
}

func Fuzz_Nosy_ringbuffer[T any]_End__(f *testing.F) {
	f.Fuzz(func(t *testing.T, size int) {
	rb := newRingBuffer(size)
	rb.End()
	})
}

func Fuzz_Nosy_ringbuffer[T any]_Get__(f *testing.F) {
	f.Fuzz(func(t *testing.T, size int, idx int) {
	rb := newRingBuffer(size)
	rb.Get(idx)
	})
}

func Fuzz_Nosy_ringbuffer[T any]_Len__(f *testing.F) {
	f.Fuzz(func(t *testing.T, size int) {
	rb := newRingBuffer(size)
	rb.Len()
	})
}

func Fuzz_Nosy_ringbuffer[T any]_Pop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, size int) {
	rb := newRingBuffer(size)
	rb.Pop()
	})
}

// skipping Fuzz_Nosy_ringbuffer[T any]_Push__ because parameters include func, chan, or unsupported interface: T

func Fuzz_Nosy_ringbuffer[T any]_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, size int) {
	rb := newRingBuffer(size)
	rb.Reset()
	})
}

func Fuzz_Nosy_ringbuffer[T any]_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, size int) {
	rb := newRingBuffer(size)
	rb.Start()
	})
}

func Fuzz_Nosy_L1SafeEvent_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var ev L1SafeEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}

	ev.String()
	})
}

func Fuzz_Nosy_L1UnsafeEvent_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var ev L1UnsafeEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}

	ev.String()
	})
}

// skipping Fuzz_Nosy_Metrics_RecordL1Ref__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/status.Metrics

// skipping Fuzz_Nosy_Metrics_RecordL1ReorgDepth__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/status.Metrics

