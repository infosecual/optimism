package metrics

import (
	"testing"
	"time"

	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum/go-ethereum/common"
	"github.com/prometheus/client_golang/prometheus"
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

func Fuzz_Nosy_CacheMetrics_CacheAdd__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *CacheMetrics
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var typeLabel string
		fill_err = tp.Fill(&typeLabel)
		if fill_err != nil {
			return
		}
		var typeCacheSize int
		fill_err = tp.Fill(&typeCacheSize)
		if fill_err != nil {
			return
		}
		var evicted bool
		fill_err = tp.Fill(&evicted)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.CacheAdd(typeLabel, typeCacheSize, evicted)
	})
}

func Fuzz_Nosy_CacheMetrics_CacheGet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *CacheMetrics
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var typeLabel string
		fill_err = tp.Fill(&typeLabel)
		if fill_err != nil {
			return
		}
		var hit bool
		fill_err = tp.Fill(&hit)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.CacheGet(typeLabel, hit)
	})
}

func Fuzz_Nosy_Event_Record__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *Event
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.Record()
	})
}

func Fuzz_Nosy_EventVec_Record__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *EventVec
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var lvs []string
		fill_err = tp.Fill(&lvs)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.Record(lvs...)
	})
}

func Fuzz_Nosy_NoopRPCMetrics_RecordRPCClientRequest__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *NoopRPCMetrics
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var method string
		fill_err = tp.Fill(&method)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.RecordRPCClientRequest(method)
	})
}

// skipping Fuzz_Nosy_NoopRPCMetrics_RecordRPCClientResponse__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_NoopRPCMetrics_RecordRPCServerRequest__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *NoopRPCMetrics
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var method string
		fill_err = tp.Fill(&method)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.RecordRPCServerRequest(method)
	})
}

func Fuzz_Nosy_NoopRefMetrics_RecordL1Ref__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *NoopRefMetrics
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var _x2 string
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var _x3 eth.L1BlockRef
		fill_err = tp.Fill(&_x3)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.RecordL1Ref(_x2, _x3)
	})
}

func Fuzz_Nosy_NoopRefMetrics_RecordL2Ref__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *NoopRefMetrics
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var _x2 string
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var _x3 eth.L2BlockRef
		fill_err = tp.Fill(&_x3)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.RecordL2Ref(_x2, _x3)
	})
}

func Fuzz_Nosy_NoopRefMetrics_RecordRef__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *NoopRefMetrics
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var _x2 string
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var _x3 string
		fill_err = tp.Fill(&_x3)
		if fill_err != nil {
			return
		}
		var _x4 uint64
		fill_err = tp.Fill(&_x4)
		if fill_err != nil {
			return
		}
		var _x5 uint64
		fill_err = tp.Fill(&_x5)
		if fill_err != nil {
			return
		}
		var _x6 common.Hash
		fill_err = tp.Fill(&_x6)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.RecordRef(_x2, _x3, _x4, _x5, _x6)
	})
}

func Fuzz_Nosy_PromHTTPRecorder_RecordHTTPRequest__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *PromHTTPRecorder
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var params *HTTPParams
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		if p == nil || params == nil {
			return
		}

		p.RecordHTTPRequest(params)
	})
}

func Fuzz_Nosy_PromHTTPRecorder_RecordHTTPRequestDuration__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *PromHTTPRecorder
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var params *HTTPParams
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		var dur time.Duration
		fill_err = tp.Fill(&dur)
		if fill_err != nil {
			return
		}
		if p == nil || params == nil {
			return
		}

		p.RecordHTTPRequestDuration(params, dur)
	})
}

func Fuzz_Nosy_PromHTTPRecorder_RecordHTTPResponse__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *PromHTTPRecorder
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var params *HTTPParams
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		if p == nil || params == nil {
			return
		}

		p.RecordHTTPResponse(params)
	})
}

func Fuzz_Nosy_PromHTTPRecorder_RecordHTTPResponseSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *PromHTTPRecorder
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var params *HTTPParams
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		var size int
		fill_err = tp.Fill(&size)
		if fill_err != nil {
			return
		}
		if p == nil || params == nil {
			return
		}

		p.RecordHTTPResponseSize(params, size)
	})
}

func Fuzz_Nosy_PromHTTPRecorder_RecordInflightRequest__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *PromHTTPRecorder
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var params *HTTPParams
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		var quantity int
		fill_err = tp.Fill(&quantity)
		if fill_err != nil {
			return
		}
		if p == nil || params == nil {
			return
		}

		p.RecordInflightRequest(params, quantity)
	})
}

func Fuzz_Nosy_PromNodeRecorder_RecordInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *PromNodeRecorder
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var version string
		fill_err = tp.Fill(&version)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.RecordInfo(version)
	})
}

func Fuzz_Nosy_PromNodeRecorder_RecordUp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *PromNodeRecorder
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.RecordUp()
	})
}

func Fuzz_Nosy_RPCClientMetrics_RecordRPCClientRequest__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *RPCClientMetrics
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var method string
		fill_err = tp.Fill(&method)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.RecordRPCClientRequest(method)
	})
}

// skipping Fuzz_Nosy_RPCClientMetrics_RecordRPCClientResponse__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_RPCServerMetrics_RecordRPCServerRequest__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *RPCServerMetrics
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var method string
		fill_err = tp.Fill(&method)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.RecordRPCServerRequest(method)
	})
}

func Fuzz_Nosy_RefMetrics_RecordL1Ref__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *RefMetrics
		fill_err = tp.Fill(&m)
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
		if m == nil {
			return
		}

		m.RecordL1Ref(name, ref)
	})
}

func Fuzz_Nosy_RefMetrics_RecordL2Ref__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *RefMetrics
		fill_err = tp.Fill(&m)
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
		if m == nil {
			return
		}

		m.RecordL2Ref(name, ref)
	})
}

func Fuzz_Nosy_RefMetrics_RecordRef__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *RefMetrics
		fill_err = tp.Fill(&m)
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
		if m == nil {
			return
		}

		m.RecordRef(layer, name, num, timestamp, h)
	})
}

func Fuzz_Nosy_documentor_Document__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *documentor
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.Document()
	})
}

func Fuzz_Nosy_documentor_NewCounter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *documentor
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var opts prometheus.CounterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.NewCounter(opts)
	})
}

func Fuzz_Nosy_documentor_NewCounterVec__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *documentor
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var opts prometheus.CounterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var labelNames []string
		fill_err = tp.Fill(&labelNames)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.NewCounterVec(opts, labelNames)
	})
}

func Fuzz_Nosy_documentor_NewGauge__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *documentor
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var opts prometheus.GaugeOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.NewGauge(opts)
	})
}

// skipping Fuzz_Nosy_documentor_NewGaugeFunc__ because parameters include func, chan, or unsupported interface: func() float64

func Fuzz_Nosy_documentor_NewGaugeVec__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *documentor
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var opts prometheus.GaugeOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var labelNames []string
		fill_err = tp.Fill(&labelNames)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.NewGaugeVec(opts, labelNames)
	})
}

func Fuzz_Nosy_documentor_NewHistogram__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *documentor
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var opts prometheus.HistogramOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.NewHistogram(opts)
	})
}

func Fuzz_Nosy_documentor_NewHistogramVec__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *documentor
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var opts prometheus.HistogramOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var labelNames []string
		fill_err = tp.Fill(&labelNames)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.NewHistogramVec(opts, labelNames)
	})
}

func Fuzz_Nosy_documentor_NewSummary__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *documentor
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var opts prometheus.SummaryOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.NewSummary(opts)
	})
}

func Fuzz_Nosy_documentor_NewSummaryVec__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *documentor
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var opts prometheus.SummaryOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var labelNames []string
		fill_err = tp.Fill(&labelNames)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.NewSummaryVec(opts, labelNames)
	})
}

func Fuzz_Nosy_noopHTTPRecorder_RecordHTTPRequest__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *noopHTTPRecorder
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var _x2 *HTTPParams
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if n == nil || _x2 == nil {
			return
		}

		n.RecordHTTPRequest(_x2)
	})
}

func Fuzz_Nosy_noopHTTPRecorder_RecordHTTPRequestDuration__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *noopHTTPRecorder
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var _x2 *HTTPParams
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var _x3 time.Duration
		fill_err = tp.Fill(&_x3)
		if fill_err != nil {
			return
		}
		if n == nil || _x2 == nil {
			return
		}

		n.RecordHTTPRequestDuration(_x2, _x3)
	})
}

func Fuzz_Nosy_noopHTTPRecorder_RecordHTTPResponse__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *noopHTTPRecorder
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var _x2 *HTTPParams
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if n == nil || _x2 == nil {
			return
		}

		n.RecordHTTPResponse(_x2)
	})
}

func Fuzz_Nosy_noopHTTPRecorder_RecordHTTPResponseSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *noopHTTPRecorder
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var _x2 *HTTPParams
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var _x3 int
		fill_err = tp.Fill(&_x3)
		if fill_err != nil {
			return
		}
		if n == nil || _x2 == nil {
			return
		}

		n.RecordHTTPResponseSize(_x2, _x3)
	})
}

func Fuzz_Nosy_noopHTTPRecorder_RecordInflightRequest__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *noopHTTPRecorder
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var _x2 *HTTPParams
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var _x3 int
		fill_err = tp.Fill(&_x3)
		if fill_err != nil {
			return
		}
		if n == nil || _x2 == nil {
			return
		}

		n.RecordInflightRequest(_x2, _x3)
	})
}

func Fuzz_Nosy_noopNodeRecorder_RecordInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *noopNodeRecorder
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var _x2 string
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.RecordInfo(_x2)
	})
}

func Fuzz_Nosy_noopNodeRecorder_RecordUp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *noopNodeRecorder
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

func Fuzz_Nosy_CLIConfig_Check__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx *cli.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if ctx == nil {
			return
		}

		m := ReadCLIConfig(ctx)
		m.Check()
	})
}

func Fuzz_Nosy_Factory_Document__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var registry *prometheus.Registry
		fill_err = tp.Fill(&registry)
		if fill_err != nil {
			return
		}
		if registry == nil {
			return
		}

		_x1 := With(registry)
		_x1.Document()
	})
}

func Fuzz_Nosy_Factory_NewCounter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var registry *prometheus.Registry
		fill_err = tp.Fill(&registry)
		if fill_err != nil {
			return
		}
		var opts prometheus.CounterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if registry == nil {
			return
		}

		_x1 := With(registry)
		_x1.NewCounter(opts)
	})
}

func Fuzz_Nosy_Factory_NewCounterVec__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var registry *prometheus.Registry
		fill_err = tp.Fill(&registry)
		if fill_err != nil {
			return
		}
		var opts prometheus.CounterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var labelNames []string
		fill_err = tp.Fill(&labelNames)
		if fill_err != nil {
			return
		}
		if registry == nil {
			return
		}

		_x1 := With(registry)
		_x1.NewCounterVec(opts, labelNames)
	})
}

func Fuzz_Nosy_Factory_NewGauge__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var registry *prometheus.Registry
		fill_err = tp.Fill(&registry)
		if fill_err != nil {
			return
		}
		var opts prometheus.GaugeOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if registry == nil {
			return
		}

		_x1 := With(registry)
		_x1.NewGauge(opts)
	})
}

// skipping Fuzz_Nosy_Factory_NewGaugeFunc__ because parameters include func, chan, or unsupported interface: func() float64

func Fuzz_Nosy_Factory_NewGaugeVec__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var registry *prometheus.Registry
		fill_err = tp.Fill(&registry)
		if fill_err != nil {
			return
		}
		var opts prometheus.GaugeOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var labelNames []string
		fill_err = tp.Fill(&labelNames)
		if fill_err != nil {
			return
		}
		if registry == nil {
			return
		}

		_x1 := With(registry)
		_x1.NewGaugeVec(opts, labelNames)
	})
}

func Fuzz_Nosy_Factory_NewHistogram__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var registry *prometheus.Registry
		fill_err = tp.Fill(&registry)
		if fill_err != nil {
			return
		}
		var opts prometheus.HistogramOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if registry == nil {
			return
		}

		_x1 := With(registry)
		_x1.NewHistogram(opts)
	})
}

func Fuzz_Nosy_Factory_NewHistogramVec__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var registry *prometheus.Registry
		fill_err = tp.Fill(&registry)
		if fill_err != nil {
			return
		}
		var opts prometheus.HistogramOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var labelNames []string
		fill_err = tp.Fill(&labelNames)
		if fill_err != nil {
			return
		}
		if registry == nil {
			return
		}

		_x1 := With(registry)
		_x1.NewHistogramVec(opts, labelNames)
	})
}

func Fuzz_Nosy_Factory_NewSummary__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var registry *prometheus.Registry
		fill_err = tp.Fill(&registry)
		if fill_err != nil {
			return
		}
		var opts prometheus.SummaryOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if registry == nil {
			return
		}

		_x1 := With(registry)
		_x1.NewSummary(opts)
	})
}

func Fuzz_Nosy_Factory_NewSummaryVec__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var registry *prometheus.Registry
		fill_err = tp.Fill(&registry)
		if fill_err != nil {
			return
		}
		var opts prometheus.SummaryOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var labelNames []string
		fill_err = tp.Fill(&labelNames)
		if fill_err != nil {
			return
		}
		if registry == nil {
			return
		}

		_x1 := With(registry)
		_x1.NewSummaryVec(opts, labelNames)
	})
}

func Fuzz_Nosy_HTTPRecorder_RecordHTTPRequest__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *prometheus.Registry
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var ns string
		fill_err = tp.Fill(&ns)
		if fill_err != nil {
			return
		}
		var params *HTTPParams
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		if r == nil || params == nil {
			return
		}

		_x1 := NewPromHTTPRecorder(r, ns)
		_x1.RecordHTTPRequest(params)
	})
}

func Fuzz_Nosy_HTTPRecorder_RecordHTTPRequestDuration__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *prometheus.Registry
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var ns string
		fill_err = tp.Fill(&ns)
		if fill_err != nil {
			return
		}
		var params *HTTPParams
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		var dur time.Duration
		fill_err = tp.Fill(&dur)
		if fill_err != nil {
			return
		}
		if r == nil || params == nil {
			return
		}

		_x1 := NewPromHTTPRecorder(r, ns)
		_x1.RecordHTTPRequestDuration(params, dur)
	})
}

func Fuzz_Nosy_HTTPRecorder_RecordHTTPResponse__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *prometheus.Registry
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var ns string
		fill_err = tp.Fill(&ns)
		if fill_err != nil {
			return
		}
		var params *HTTPParams
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		if r == nil || params == nil {
			return
		}

		_x1 := NewPromHTTPRecorder(r, ns)
		_x1.RecordHTTPResponse(params)
	})
}

func Fuzz_Nosy_HTTPRecorder_RecordHTTPResponseSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *prometheus.Registry
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var ns string
		fill_err = tp.Fill(&ns)
		if fill_err != nil {
			return
		}
		var params *HTTPParams
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		var size int
		fill_err = tp.Fill(&size)
		if fill_err != nil {
			return
		}
		if r == nil || params == nil {
			return
		}

		_x1 := NewPromHTTPRecorder(r, ns)
		_x1.RecordHTTPResponseSize(params, size)
	})
}

func Fuzz_Nosy_HTTPRecorder_RecordInflightRequest__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *prometheus.Registry
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var ns string
		fill_err = tp.Fill(&ns)
		if fill_err != nil {
			return
		}
		var params *HTTPParams
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		var quantity int
		fill_err = tp.Fill(&quantity)
		if fill_err != nil {
			return
		}
		if r == nil || params == nil {
			return
		}

		_x1 := NewPromHTTPRecorder(r, ns)
		_x1.RecordInflightRequest(params, quantity)
	})
}

func Fuzz_Nosy_NodeRecorder_RecordInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *prometheus.Registry
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var ns string
		fill_err = tp.Fill(&ns)
		if fill_err != nil {
			return
		}
		var version string
		fill_err = tp.Fill(&version)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		_x1 := NewPromNodeRecorder(r, ns)
		_x1.RecordInfo(version)
	})
}

func Fuzz_Nosy_NodeRecorder_RecordUp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *prometheus.Registry
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var ns string
		fill_err = tp.Fill(&ns)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		_x1 := NewPromNodeRecorder(r, ns)
		_x1.RecordUp()
	})
}

// skipping Fuzz_Nosy_RPCClientMetricer_RecordRPCClientRequest__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/metrics.RPCClientMetricer

// skipping Fuzz_Nosy_RPCClientMetricer_RecordRPCClientResponse__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/metrics.RPCClientMetricer

// skipping Fuzz_Nosy_RPCServerMetricer_RecordRPCServerRequest__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/metrics.RPCServerMetricer

// skipping Fuzz_Nosy_RefMetricer_RecordL1Ref__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/metrics.RefMetricer

// skipping Fuzz_Nosy_RefMetricer_RecordL2Ref__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/metrics.RefMetricer

// skipping Fuzz_Nosy_RefMetricer_RecordRef__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/metrics.RefMetricer

// skipping Fuzz_Nosy_RegistryMetricer_Registry__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/metrics.RegistryMetricer

func Fuzz_Nosy_CLIFlags__(f *testing.F) {
	f.Fuzz(func(t *testing.T, envPrefix string) {
		CLIFlags(envPrefix)
	})
}

func Fuzz_Nosy_fullName__(f *testing.F) {
	f.Fuzz(func(t *testing.T, ns string, subsystem string, name string) {
		fullName(ns, subsystem, name)
	})
}
