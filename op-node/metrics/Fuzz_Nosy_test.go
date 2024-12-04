package metrics

import (
	"context"
	"testing"
	"time"

	"github.com/ethereum-optimism/optimism/op-node/p2p/store"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
	"github.com/libp2p/go-libp2p/core/metrics"
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

func Fuzz_Nosy_Metrics_ClientPayloadByNumberEvent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var procName string
		fill_err = tp.Fill(&procName)
		if fill_err != nil {
			return
		}
		var num uint64
		fill_err = tp.Fill(&num)
		if fill_err != nil {
			return
		}
		var resultCode byte
		fill_err = tp.Fill(&resultCode)
		if fill_err != nil {
			return
		}
		var duration time.Duration
		fill_err = tp.Fill(&duration)
		if fill_err != nil {
			return
		}

		m := NewMetrics(procName)
		m.ClientPayloadByNumberEvent(num, resultCode, duration)
	})
}

func Fuzz_Nosy_Metrics_CountSequencedTxsInBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, procName string, txns int, deposits int) {
		m := NewMetrics(procName)
		m.CountSequencedTxsInBlock(txns, deposits)
	})
}

func Fuzz_Nosy_Metrics_DecPeerCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, procName string) {
		m := NewMetrics(procName)
		m.DecPeerCount()
	})
}

func Fuzz_Nosy_Metrics_DecStreamCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, procName string) {
		m := NewMetrics(procName)
		m.DecStreamCount()
	})
}

func Fuzz_Nosy_Metrics_Document__(f *testing.F) {
	f.Fuzz(func(t *testing.T, procName string) {
		m := NewMetrics(procName)
		m.Document()
	})
}

func Fuzz_Nosy_Metrics_IncPeerCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, procName string) {
		m := NewMetrics(procName)
		m.IncPeerCount()
	})
}

func Fuzz_Nosy_Metrics_IncStreamCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, procName string) {
		m := NewMetrics(procName)
		m.IncStreamCount()
	})
}

func Fuzz_Nosy_Metrics_PayloadsQuarantineSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, procName string, n int) {
		m := NewMetrics(procName)
		m.PayloadsQuarantineSize(n)
	})
}

func Fuzz_Nosy_Metrics_RecordAccept__(f *testing.F) {
	f.Fuzz(func(t *testing.T, procName string, allow bool) {
		m := NewMetrics(procName)
		m.RecordAccept(allow)
	})
}

func Fuzz_Nosy_Metrics_RecordBandwidth__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var procName string
		fill_err = tp.Fill(&procName)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var bwc *metrics.BandwidthCounter
		fill_err = tp.Fill(&bwc)
		if fill_err != nil {
			return
		}
		if bwc == nil {
			return
		}

		m := NewMetrics(procName)
		m.RecordBandwidth(ctx, bwc)
	})
}

func Fuzz_Nosy_Metrics_RecordChannelInputBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, procName string, inputCompressedBytes int) {
		m := NewMetrics(procName)
		m.RecordChannelInputBytes(inputCompressedBytes)
	})
}

func Fuzz_Nosy_Metrics_RecordChannelTimedOut__(f *testing.F) {
	f.Fuzz(func(t *testing.T, procName string) {
		m := NewMetrics(procName)
		m.RecordChannelTimedOut()
	})
}

func Fuzz_Nosy_Metrics_RecordDerivationError__(f *testing.F) {
	f.Fuzz(func(t *testing.T, procName string) {
		m := NewMetrics(procName)
		m.RecordDerivationError()
	})
}

func Fuzz_Nosy_Metrics_RecordDerivedBatches__(f *testing.F) {
	f.Fuzz(func(t *testing.T, procName string, batchType string) {
		m := NewMetrics(procName)
		m.RecordDerivedBatches(batchType)
	})
}

func Fuzz_Nosy_Metrics_RecordDial__(f *testing.F) {
	f.Fuzz(func(t *testing.T, procName string, allow bool) {
		m := NewMetrics(procName)
		m.RecordDial(allow)
	})
}

func Fuzz_Nosy_Metrics_RecordEmittedEvent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, procName string, eventName string, emitter string) {
		m := NewMetrics(procName)
		m.RecordEmittedEvent(eventName, emitter)
	})
}

func Fuzz_Nosy_Metrics_RecordEventsRateLimited__(f *testing.F) {
	f.Fuzz(func(t *testing.T, procName string) {
		m := NewMetrics(procName)
		m.RecordEventsRateLimited()
	})
}

func Fuzz_Nosy_Metrics_RecordFrame__(f *testing.F) {
	f.Fuzz(func(t *testing.T, procName string) {
		m := NewMetrics(procName)
		m.RecordFrame()
	})
}

func Fuzz_Nosy_Metrics_RecordGossipEvent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, procName string, evType int32) {
		m := NewMetrics(procName)
		m.RecordGossipEvent(evType)
	})
}

func Fuzz_Nosy_Metrics_RecordHeadChannelOpened__(f *testing.F) {
	f.Fuzz(func(t *testing.T, procName string) {
		m := NewMetrics(procName)
		m.RecordHeadChannelOpened()
	})
}

func Fuzz_Nosy_Metrics_RecordIPUnban__(f *testing.F) {
	f.Fuzz(func(t *testing.T, procName string) {
		m := NewMetrics(procName)
		m.RecordIPUnban()
	})
}

func Fuzz_Nosy_Metrics_RecordInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, procName string, version string) {
		m := NewMetrics(procName)
		m.RecordInfo(version)
	})
}

func Fuzz_Nosy_Metrics_RecordL1ReorgDepth__(f *testing.F) {
	f.Fuzz(func(t *testing.T, procName string, d uint64) {
		m := NewMetrics(procName)
		m.RecordL1ReorgDepth(d)
	})
}

func Fuzz_Nosy_Metrics_RecordL1RequestTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var procName string
		fill_err = tp.Fill(&procName)
		if fill_err != nil {
			return
		}
		var method string
		fill_err = tp.Fill(&method)
		if fill_err != nil {
			return
		}
		var duration time.Duration
		fill_err = tp.Fill(&duration)
		if fill_err != nil {
			return
		}

		m := NewMetrics(procName)
		m.RecordL1RequestTime(method, duration)
	})
}

func Fuzz_Nosy_Metrics_RecordPeerUnban__(f *testing.F) {
	f.Fuzz(func(t *testing.T, procName string) {
		m := NewMetrics(procName)
		m.RecordPeerUnban()
	})
}

func Fuzz_Nosy_Metrics_RecordPipelineReset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, procName string) {
		m := NewMetrics(procName)
		m.RecordPipelineReset()
	})
}

func Fuzz_Nosy_Metrics_RecordProcessedEvent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var procName string
		fill_err = tp.Fill(&procName)
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

		m := NewMetrics(procName)
		m.RecordProcessedEvent(eventName, deriver, duration)
	})
}

func Fuzz_Nosy_Metrics_RecordPublishingError__(f *testing.F) {
	f.Fuzz(func(t *testing.T, procName string) {
		m := NewMetrics(procName)
		m.RecordPublishingError()
	})
}

func Fuzz_Nosy_Metrics_RecordReceivedUnsafePayload__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var procName string
		fill_err = tp.Fill(&procName)
		if fill_err != nil {
			return
		}
		var payload *eth.ExecutionPayloadEnvelope
		fill_err = tp.Fill(&payload)
		if fill_err != nil {
			return
		}
		if payload == nil {
			return
		}

		m := NewMetrics(procName)
		m.RecordReceivedUnsafePayload(payload)
	})
}

func Fuzz_Nosy_Metrics_RecordSequencerBuildingDiffTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var procName string
		fill_err = tp.Fill(&procName)
		if fill_err != nil {
			return
		}
		var duration time.Duration
		fill_err = tp.Fill(&duration)
		if fill_err != nil {
			return
		}

		m := NewMetrics(procName)
		m.RecordSequencerBuildingDiffTime(duration)
	})
}

func Fuzz_Nosy_Metrics_RecordSequencerInconsistentL1Origin__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var procName string
		fill_err = tp.Fill(&procName)
		if fill_err != nil {
			return
		}
		var from eth.BlockID
		fill_err = tp.Fill(&from)
		if fill_err != nil {
			return
		}
		var to eth.BlockID
		fill_err = tp.Fill(&to)
		if fill_err != nil {
			return
		}

		m := NewMetrics(procName)
		m.RecordSequencerInconsistentL1Origin(from, to)
	})
}

func Fuzz_Nosy_Metrics_RecordSequencerReset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, procName string) {
		m := NewMetrics(procName)
		m.RecordSequencerReset()
	})
}

func Fuzz_Nosy_Metrics_RecordSequencerSealingTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var procName string
		fill_err = tp.Fill(&procName)
		if fill_err != nil {
			return
		}
		var duration time.Duration
		fill_err = tp.Fill(&duration)
		if fill_err != nil {
			return
		}

		m := NewMetrics(procName)
		m.RecordSequencerSealingTime(duration)
	})
}

func Fuzz_Nosy_Metrics_RecordSequencingError__(f *testing.F) {
	f.Fuzz(func(t *testing.T, procName string) {
		m := NewMetrics(procName)
		m.RecordSequencingError()
	})
}

func Fuzz_Nosy_Metrics_RecordUnsafePayloadsBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var procName string
		fill_err = tp.Fill(&procName)
		if fill_err != nil {
			return
		}
		var length uint64
		fill_err = tp.Fill(&length)
		if fill_err != nil {
			return
		}
		var memSize uint64
		fill_err = tp.Fill(&memSize)
		if fill_err != nil {
			return
		}
		var next eth.BlockID
		fill_err = tp.Fill(&next)
		if fill_err != nil {
			return
		}

		m := NewMetrics(procName)
		m.RecordUnsafePayloadsBuffer(length, memSize, next)
	})
}

func Fuzz_Nosy_Metrics_RecordUp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, procName string) {
		m := NewMetrics(procName)
		m.RecordUp()
	})
}

func Fuzz_Nosy_Metrics_ReportProtocolVersions__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var procName string
		fill_err = tp.Fill(&procName)
		if fill_err != nil {
			return
		}
		var local params.ProtocolVersion
		fill_err = tp.Fill(&local)
		if fill_err != nil {
			return
		}
		var engine params.ProtocolVersion
		fill_err = tp.Fill(&engine)
		if fill_err != nil {
			return
		}
		var recommended params.ProtocolVersion
		fill_err = tp.Fill(&recommended)
		if fill_err != nil {
			return
		}
		var required params.ProtocolVersion
		fill_err = tp.Fill(&required)
		if fill_err != nil {
			return
		}

		m := NewMetrics(procName)
		m.ReportProtocolVersions(local, engine, recommended, required)
	})
}

func Fuzz_Nosy_Metrics_ServerPayloadByNumberEvent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var procName string
		fill_err = tp.Fill(&procName)
		if fill_err != nil {
			return
		}
		var num uint64
		fill_err = tp.Fill(&num)
		if fill_err != nil {
			return
		}
		var resultCode byte
		fill_err = tp.Fill(&resultCode)
		if fill_err != nil {
			return
		}
		var duration time.Duration
		fill_err = tp.Fill(&duration)
		if fill_err != nil {
			return
		}

		m := NewMetrics(procName)
		m.ServerPayloadByNumberEvent(num, resultCode, duration)
	})
}

func Fuzz_Nosy_Metrics_SetDerivationIdle__(f *testing.F) {
	f.Fuzz(func(t *testing.T, procName string, status bool) {
		m := NewMetrics(procName)
		m.SetDerivationIdle(status)
	})
}

func Fuzz_Nosy_Metrics_SetPeerScores__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var procName string
		fill_err = tp.Fill(&procName)
		if fill_err != nil {
			return
		}
		var allScores []store.PeerScores
		fill_err = tp.Fill(&allScores)
		if fill_err != nil {
			return
		}

		m := NewMetrics(procName)
		m.SetPeerScores(allScores)
	})
}

func Fuzz_Nosy_Metrics_SetSequencerState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, procName string, active bool) {
		m := NewMetrics(procName)
		m.SetSequencerState(active)
	})
}

func Fuzz_Nosy_Metrics_StartServer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, procName string, hostname string, port int) {
		m := NewMetrics(procName)
		m.StartServer(hostname, port)
	})
}

func Fuzz_Nosy_noopMetricer_ClientPayloadByNumberEvent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *noopMetricer
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var num uint64
		fill_err = tp.Fill(&num)
		if fill_err != nil {
			return
		}
		var resultCode byte
		fill_err = tp.Fill(&resultCode)
		if fill_err != nil {
			return
		}
		var duration time.Duration
		fill_err = tp.Fill(&duration)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.ClientPayloadByNumberEvent(num, resultCode, duration)
	})
}

func Fuzz_Nosy_noopMetricer_CountSequencedTxsInBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *noopMetricer
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var txns int
		fill_err = tp.Fill(&txns)
		if fill_err != nil {
			return
		}
		var deposits int
		fill_err = tp.Fill(&deposits)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.CountSequencedTxsInBlock(txns, deposits)
	})
}

func Fuzz_Nosy_noopMetricer_DecPeerCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *noopMetricer
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.DecPeerCount()
	})
}

func Fuzz_Nosy_noopMetricer_DecStreamCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *noopMetricer
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.DecStreamCount()
	})
}

func Fuzz_Nosy_noopMetricer_Document__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *noopMetricer
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.Document()
	})
}

func Fuzz_Nosy_noopMetricer_IncPeerCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *noopMetricer
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.IncPeerCount()
	})
}

func Fuzz_Nosy_noopMetricer_IncStreamCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *noopMetricer
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.IncStreamCount()
	})
}

func Fuzz_Nosy_noopMetricer_PayloadsQuarantineSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *noopMetricer
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var _x2 int
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.PayloadsQuarantineSize(_x2)
	})
}

func Fuzz_Nosy_noopMetricer_RecordAccept__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *noopMetricer
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var allow bool
		fill_err = tp.Fill(&allow)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.RecordAccept(allow)
	})
}

func Fuzz_Nosy_noopMetricer_RecordBandwidth__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *noopMetricer
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var bwc *metrics.BandwidthCounter
		fill_err = tp.Fill(&bwc)
		if fill_err != nil {
			return
		}
		if n == nil || bwc == nil {
			return
		}

		n.RecordBandwidth(ctx, bwc)
	})
}

func Fuzz_Nosy_noopMetricer_RecordChannelInputBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *noopMetricer
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var _x2 int
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.RecordChannelInputBytes(_x2)
	})
}

func Fuzz_Nosy_noopMetricer_RecordChannelTimedOut__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *noopMetricer
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.RecordChannelTimedOut()
	})
}

func Fuzz_Nosy_noopMetricer_RecordDerivationError__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *noopMetricer
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.RecordDerivationError()
	})
}

func Fuzz_Nosy_noopMetricer_RecordDerivedBatches__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *noopMetricer
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var batchType string
		fill_err = tp.Fill(&batchType)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.RecordDerivedBatches(batchType)
	})
}

func Fuzz_Nosy_noopMetricer_RecordDial__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *noopMetricer
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var allow bool
		fill_err = tp.Fill(&allow)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.RecordDial(allow)
	})
}

func Fuzz_Nosy_noopMetricer_RecordEmittedEvent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *noopMetricer
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
		if n == nil {
			return
		}

		n.RecordEmittedEvent(eventName, emitter)
	})
}

func Fuzz_Nosy_noopMetricer_RecordEventsRateLimited__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *noopMetricer
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.RecordEventsRateLimited()
	})
}

func Fuzz_Nosy_noopMetricer_RecordFrame__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *noopMetricer
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.RecordFrame()
	})
}

func Fuzz_Nosy_noopMetricer_RecordGossipEvent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *noopMetricer
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var evType int32
		fill_err = tp.Fill(&evType)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.RecordGossipEvent(evType)
	})
}

func Fuzz_Nosy_noopMetricer_RecordHeadChannelOpened__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *noopMetricer
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.RecordHeadChannelOpened()
	})
}

func Fuzz_Nosy_noopMetricer_RecordIPUnban__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *noopMetricer
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.RecordIPUnban()
	})
}

func Fuzz_Nosy_noopMetricer_RecordInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *noopMetricer
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var version string
		fill_err = tp.Fill(&version)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.RecordInfo(version)
	})
}

func Fuzz_Nosy_noopMetricer_RecordL1Ref__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *noopMetricer
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var ref eth.L1BlockRef
		fill_err = tp.Fill(&ref)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.RecordL1Ref(name, ref)
	})
}

func Fuzz_Nosy_noopMetricer_RecordL1ReorgDepth__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *noopMetricer
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var d uint64
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.RecordL1ReorgDepth(d)
	})
}

func Fuzz_Nosy_noopMetricer_RecordL2Ref__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *noopMetricer
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var ref eth.L2BlockRef
		fill_err = tp.Fill(&ref)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.RecordL2Ref(name, ref)
	})
}

func Fuzz_Nosy_noopMetricer_RecordPeerUnban__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *noopMetricer
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.RecordPeerUnban()
	})
}

func Fuzz_Nosy_noopMetricer_RecordPipelineReset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *noopMetricer
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.RecordPipelineReset()
	})
}

func Fuzz_Nosy_noopMetricer_RecordProcessedEvent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *noopMetricer
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
		if n == nil {
			return
		}

		n.RecordProcessedEvent(eventName, deriver, duration)
	})
}

func Fuzz_Nosy_noopMetricer_RecordPublishingError__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *noopMetricer
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.RecordPublishingError()
	})
}

func Fuzz_Nosy_noopMetricer_RecordReceivedUnsafePayload__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *noopMetricer
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var payload *eth.ExecutionPayloadEnvelope
		fill_err = tp.Fill(&payload)
		if fill_err != nil {
			return
		}
		if n == nil || payload == nil {
			return
		}

		n.RecordReceivedUnsafePayload(payload)
	})
}

func Fuzz_Nosy_noopMetricer_RecordRef__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *noopMetricer
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var layer string
		fill_err = tp.Fill(&layer)
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
		var timestamp uint64
		fill_err = tp.Fill(&timestamp)
		if fill_err != nil {
			return
		}
		var h common.Hash
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.RecordRef(layer, name, num, timestamp, h)
	})
}

func Fuzz_Nosy_noopMetricer_RecordSequencerBuildingDiffTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *noopMetricer
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var duration time.Duration
		fill_err = tp.Fill(&duration)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.RecordSequencerBuildingDiffTime(duration)
	})
}

func Fuzz_Nosy_noopMetricer_RecordSequencerInconsistentL1Origin__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *noopMetricer
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var from eth.BlockID
		fill_err = tp.Fill(&from)
		if fill_err != nil {
			return
		}
		var to eth.BlockID
		fill_err = tp.Fill(&to)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.RecordSequencerInconsistentL1Origin(from, to)
	})
}

func Fuzz_Nosy_noopMetricer_RecordSequencerReset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *noopMetricer
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.RecordSequencerReset()
	})
}

func Fuzz_Nosy_noopMetricer_RecordSequencerSealingTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *noopMetricer
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var duration time.Duration
		fill_err = tp.Fill(&duration)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.RecordSequencerSealingTime(duration)
	})
}

func Fuzz_Nosy_noopMetricer_RecordSequencingError__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *noopMetricer
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.RecordSequencingError()
	})
}

func Fuzz_Nosy_noopMetricer_RecordUnsafePayloadsBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *noopMetricer
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var length uint64
		fill_err = tp.Fill(&length)
		if fill_err != nil {
			return
		}
		var memSize uint64
		fill_err = tp.Fill(&memSize)
		if fill_err != nil {
			return
		}
		var next eth.BlockID
		fill_err = tp.Fill(&next)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.RecordUnsafePayloadsBuffer(length, memSize, next)
	})
}

func Fuzz_Nosy_noopMetricer_RecordUp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *noopMetricer
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.RecordUp()
	})
}

func Fuzz_Nosy_noopMetricer_ReportProtocolVersions__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *noopMetricer
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var local params.ProtocolVersion
		fill_err = tp.Fill(&local)
		if fill_err != nil {
			return
		}
		var engine params.ProtocolVersion
		fill_err = tp.Fill(&engine)
		if fill_err != nil {
			return
		}
		var recommended params.ProtocolVersion
		fill_err = tp.Fill(&recommended)
		if fill_err != nil {
			return
		}
		var required params.ProtocolVersion
		fill_err = tp.Fill(&required)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.ReportProtocolVersions(local, engine, recommended, required)
	})
}

func Fuzz_Nosy_noopMetricer_ServerPayloadByNumberEvent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *noopMetricer
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var num uint64
		fill_err = tp.Fill(&num)
		if fill_err != nil {
			return
		}
		var resultCode byte
		fill_err = tp.Fill(&resultCode)
		if fill_err != nil {
			return
		}
		var duration time.Duration
		fill_err = tp.Fill(&duration)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.ServerPayloadByNumberEvent(num, resultCode, duration)
	})
}

func Fuzz_Nosy_noopMetricer_SetDerivationIdle__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *noopMetricer
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var status bool
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.SetDerivationIdle(status)
	})
}

func Fuzz_Nosy_noopMetricer_SetPeerScores__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *noopMetricer
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var allScores []store.PeerScores
		fill_err = tp.Fill(&allScores)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.SetPeerScores(allScores)
	})
}

func Fuzz_Nosy_noopMetricer_SetSequencerState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *noopMetricer
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var active bool
		fill_err = tp.Fill(&active)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.SetSequencerState(active)
	})
}

// skipping Fuzz_Nosy_Metricer_ClientPayloadByNumberEvent__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_CountSequencedTxsInBlock__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_DecPeerCount__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_DecStreamCount__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_Document__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_IncPeerCount__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_IncStreamCount__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_PayloadsQuarantineSize__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordAccept__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordBandwidth__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordChannelInputBytes__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordChannelTimedOut__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordDerivationError__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordDerivedBatches__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordDial__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordEmittedEvent__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordEventsRateLimited__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordFrame__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordGossipEvent__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordHeadChannelOpened__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordIPUnban__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordInfo__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordL1Ref__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordL1ReorgDepth__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordL2Ref__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordPeerUnban__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordPipelineReset__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordProcessedEvent__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordPublishingError__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordRPCClientRequest__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordRPCClientResponse__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordRPCServerRequest__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordReceivedUnsafePayload__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordRef__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordSequencerBuildingDiffTime__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordSequencerInconsistentL1Origin__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordSequencerReset__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordSequencerSealingTime__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordSequencingError__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordUnsafePayloadsBuffer__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_RecordUp__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_ReportProtocolVersions__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_ServerPayloadByNumberEvent__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_SetDerivationIdle__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_SetPeerScores__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/metrics.Metricer

// skipping Fuzz_Nosy_Metricer_SetSequencerState__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-node/metrics.Metricer
