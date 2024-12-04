package srcmap

import (
	"io"
	"testing"

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

func Fuzz_Nosy_SourceID_UnmarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id *SourceID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var d2 []byte
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}
		if id == nil {
			return
		}

		id.UnmarshalText(d2)
	})
}

func Fuzz_Nosy_SourceMap_FormattedInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *SourceMap
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var pc uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.FormattedInfo(pc)
	})
}

func Fuzz_Nosy_SourceMap_Info__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *SourceMap
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var pc uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Info(pc)
	})
}

func Fuzz_Nosy_SourceMapTracer_Hooks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var srcMaps map[common.Address]*SourceMap
		fill_err = tp.Fill(&srcMaps)
		if fill_err != nil {
			return
		}
		var out io.Writer
		fill_err = tp.Fill(&out)
		if fill_err != nil {
			return
		}

		s := NewSourceMapTracer(srcMaps, out)
		s.Hooks()
	})
}

// skipping Fuzz_Nosy_SourceMapTracer_OnFault__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/tracing.OpContext

// skipping Fuzz_Nosy_SourceMapTracer_OnOpCode__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/tracing.OpContext

func Fuzz_Nosy_SourceMapTracer_info__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var srcMaps map[common.Address]*SourceMap
		fill_err = tp.Fill(&srcMaps)
		if fill_err != nil {
			return
		}
		var out io.Writer
		fill_err = tp.Fill(&out)
		if fill_err != nil {
			return
		}
		var codeAddr common.Address
		fill_err = tp.Fill(&codeAddr)
		if fill_err != nil {
			return
		}
		var pc uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}

		s := NewSourceMapTracer(srcMaps, out)
		s.info(codeAddr, pc)
	})
}

// skipping Fuzz_Nosy_loadLineColData__ because parameters include func, chan, or unsupported interface: io/fs.FS
