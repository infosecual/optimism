package clsync

import (
	"testing"

	"github.com/ethereum-optimism/optimism/op-node/rollup/engine"
	"github.com/ethereum-optimism/optimism/op-service/eth"
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

// skipping Fuzz_Nosy_CLSync_AttachEmitter__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/event.Emitter

func Fuzz_Nosy_CLSync_LowestQueuedUnsafeBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var eq *CLSync
		fill_err = tp.Fill(&eq)
		if fill_err != nil {
			return
		}
		if eq == nil {
			return
		}

		eq.LowestQueuedUnsafeBlock()
	})
}

// skipping Fuzz_Nosy_CLSync_OnEvent__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/event.Event

func Fuzz_Nosy_CLSync_fromQueue__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var eq *CLSync
		fill_err = tp.Fill(&eq)
		if fill_err != nil {
			return
		}
		var x engine.ForkchoiceUpdateEvent
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if eq == nil {
			return
		}

		eq.fromQueue(x)
	})
}

func Fuzz_Nosy_CLSync_onForkchoiceUpdate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var eq *CLSync
		fill_err = tp.Fill(&eq)
		if fill_err != nil {
			return
		}
		var x engine.ForkchoiceUpdateEvent
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if eq == nil {
			return
		}

		eq.onForkchoiceUpdate(x)
	})
}

func Fuzz_Nosy_CLSync_onInvalidPayload__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var eq *CLSync
		fill_err = tp.Fill(&eq)
		if fill_err != nil {
			return
		}
		var x engine.PayloadInvalidEvent
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if eq == nil {
			return
		}

		eq.onInvalidPayload(x)
	})
}

func Fuzz_Nosy_CLSync_onUnsafePayload__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var eq *CLSync
		fill_err = tp.Fill(&eq)
		if fill_err != nil {
			return
		}
		var x ReceivedUnsafePayloadEvent
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if eq == nil {
			return
		}

		eq.onUnsafePayload(x)
	})
}

func Fuzz_Nosy_PayloadsQueue_Len__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var upq *PayloadsQueue
		fill_err = tp.Fill(&upq)
		if fill_err != nil {
			return
		}
		if upq == nil {
			return
		}

		upq.Len()
	})
}

func Fuzz_Nosy_PayloadsQueue_MemSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var upq *PayloadsQueue
		fill_err = tp.Fill(&upq)
		if fill_err != nil {
			return
		}
		if upq == nil {
			return
		}

		upq.MemSize()
	})
}

func Fuzz_Nosy_PayloadsQueue_Peek__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var upq *PayloadsQueue
		fill_err = tp.Fill(&upq)
		if fill_err != nil {
			return
		}
		if upq == nil {
			return
		}

		upq.Peek()
	})
}

func Fuzz_Nosy_PayloadsQueue_Pop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var upq *PayloadsQueue
		fill_err = tp.Fill(&upq)
		if fill_err != nil {
			return
		}
		if upq == nil {
			return
		}

		upq.Pop()
	})
}

func Fuzz_Nosy_PayloadsQueue_Push__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var upq *PayloadsQueue
		fill_err = tp.Fill(&upq)
		if fill_err != nil {
			return
		}
		var e *eth.ExecutionPayloadEnvelope
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if upq == nil || e == nil {
			return
		}

		upq.Push(e)
	})
}

func Fuzz_Nosy_payloadsByNumber_Pop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pq *payloadsByNumber
		fill_err = tp.Fill(&pq)
		if fill_err != nil {
			return
		}
		if pq == nil {
			return
		}

		pq.Pop()
	})
}

// skipping Fuzz_Nosy_payloadsByNumber_Push__ because parameters include func, chan, or unsupported interface: any

// skipping Fuzz_Nosy_Metrics_RecordUnsafePayloadsBuffer__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/rollup/clsync.Metrics

func Fuzz_Nosy_ReceivedUnsafePayloadEvent_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ev ReceivedUnsafePayloadEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}

		ev.String()
	})
}

func Fuzz_Nosy_payloadsByNumber_Len__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pq payloadsByNumber
		fill_err = tp.Fill(&pq)
		if fill_err != nil {
			return
		}

		pq.Len()
	})
}

func Fuzz_Nosy_payloadsByNumber_Less__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pq payloadsByNumber
		fill_err = tp.Fill(&pq)
		if fill_err != nil {
			return
		}
		var i int
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
		var j int
		fill_err = tp.Fill(&j)
		if fill_err != nil {
			return
		}

		pq.Less(i, j)
	})
}

func Fuzz_Nosy_payloadsByNumber_Swap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pq payloadsByNumber
		fill_err = tp.Fill(&pq)
		if fill_err != nil {
			return
		}
		var i int
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
		var j int
		fill_err = tp.Fill(&j)
		if fill_err != nil {
			return
		}

		pq.Swap(i, j)
	})
}

func Fuzz_Nosy_payloadMemSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *eth.ExecutionPayloadEnvelope
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		payloadMemSize(p)
	})
}
