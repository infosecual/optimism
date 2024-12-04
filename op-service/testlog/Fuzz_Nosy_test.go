package testlog

import (
	"context"
	"log/slog"
	"testing"

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

// skipping Fuzz_Nosy_CapturedAttributes_Attrs__ because parameters include func, chan, or unsupported interface: func(log/slog.Attr) bool

func Fuzz_Nosy_CapturedRecord_AttrValue__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *CapturedRecord
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.AttrValue(name)
	})
}

// skipping Fuzz_Nosy_CapturedRecord_Attrs__ because parameters include func, chan, or unsupported interface: func(log/slog.Attr) bool

func Fuzz_Nosy_CapturingHandler_Clear__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *CapturingHandler
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.Clear()
	})
}

func Fuzz_Nosy_CapturingHandler_Enabled__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *CapturingHandler
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var _x3 slog.Level
		fill_err = tp.Fill(&_x3)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.Enabled(_x2, _x3)
	})
}

// skipping Fuzz_Nosy_CapturingHandler_FindLog__ because parameters include func, chan, or unsupported interface: []github.com/ethereum-optimism/optimism/op-service/testlog.LogFilter

// skipping Fuzz_Nosy_CapturingHandler_FindLogs__ because parameters include func, chan, or unsupported interface: []github.com/ethereum-optimism/optimism/op-service/testlog.LogFilter

func Fuzz_Nosy_CapturingHandler_Handle__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *CapturingHandler
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var r slog.Record
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.Handle(ctx, r)
	})
}

func Fuzz_Nosy_CapturingHandler_WithAttrs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *CapturingHandler
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var attrs []slog.Attr
		fill_err = tp.Fill(&attrs)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.WithAttrs(attrs)
	})
}

func Fuzz_Nosy_CapturingHandler_WithGroup__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *CapturingHandler
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.WithGroup(name)
	})
}

// skipping Fuzz_Nosy_logger_Crit__ because parameters include func, chan, or unsupported interface: []any

// skipping Fuzz_Nosy_logger_Debug__ because parameters include func, chan, or unsupported interface: []any

func Fuzz_Nosy_logger_Enabled__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *logger
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var level slog.Level
		fill_err = tp.Fill(&level)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		l.Enabled(ctx, level)
	})
}

// skipping Fuzz_Nosy_logger_Error__ because parameters include func, chan, or unsupported interface: []any

func Fuzz_Nosy_logger_Handler__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *logger
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		l.Handler()
	})
}

// skipping Fuzz_Nosy_logger_Info__ because parameters include func, chan, or unsupported interface: []any

// skipping Fuzz_Nosy_logger_Log__ because parameters include func, chan, or unsupported interface: []any

// skipping Fuzz_Nosy_logger_New__ because parameters include func, chan, or unsupported interface: []any

// skipping Fuzz_Nosy_logger_Trace__ because parameters include func, chan, or unsupported interface: []any

// skipping Fuzz_Nosy_logger_Warn__ because parameters include func, chan, or unsupported interface: []any

// skipping Fuzz_Nosy_logger_With__ because parameters include func, chan, or unsupported interface: []any

// skipping Fuzz_Nosy_logger_Write__ because parameters include func, chan, or unsupported interface: []any

func Fuzz_Nosy_logger_flush__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *logger
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		l.flush()
	})
}

// skipping Fuzz_Nosy_Testing_Helper__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/testlog.Testing

// skipping Fuzz_Nosy_Testing_Logf__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/testlog.Testing

// skipping Fuzz_Nosy_CaptureLogger__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/testlog.Testing

func Fuzz_Nosy_estimateInfoLen__(f *testing.F) {
	f.Fuzz(func(t *testing.T, frameSkip int) {
		estimateInfoLen(frameSkip)
	})
}
