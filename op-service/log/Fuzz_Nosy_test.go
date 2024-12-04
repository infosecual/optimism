package log

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

func Fuzz_Nosy_DynamicLogHandler_Enabled__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *DynamicLogHandler
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var lvl slog.Level
		fill_err = tp.Fill(&lvl)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.Enabled(ctx, lvl)
	})
}

func Fuzz_Nosy_DynamicLogHandler_Handle__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *DynamicLogHandler
		fill_err = tp.Fill(&d)
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
		if d == nil {
			return
		}

		d.Handle(ctx, r)
	})
}

func Fuzz_Nosy_DynamicLogHandler_SetLogLevel__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *DynamicLogHandler
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var lvl slog.Level
		fill_err = tp.Fill(&lvl)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.SetLogLevel(lvl)
	})
}

func Fuzz_Nosy_DynamicLogHandler_WithAttrs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *DynamicLogHandler
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var attrs []slog.Attr
		fill_err = tp.Fill(&attrs)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.WithAttrs(attrs)
	})
}

func Fuzz_Nosy_DynamicLogHandler_WithGroup__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *DynamicLogHandler
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.WithGroup(name)
	})
}

func Fuzz_Nosy_FormatFlagValue_Clone__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fmtType FormatType
		fill_err = tp.Fill(&fmtType)
		if fill_err != nil {
			return
		}

		fv := NewFormatFlagValue(fmtType)
		fv.Clone()
	})
}

func Fuzz_Nosy_FormatFlagValue_Set__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fmtType FormatType
		fill_err = tp.Fill(&fmtType)
		if fill_err != nil {
			return
		}
		var value string
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}

		fv := NewFormatFlagValue(fmtType)
		fv.Set(value)
	})
}

func Fuzz_Nosy_LevelFlagValue_Clone__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var lvl slog.Level
		fill_err = tp.Fill(&lvl)
		if fill_err != nil {
			return
		}

		fv := NewLevelFlagValue(lvl)
		fv.Clone()
	})
}

func Fuzz_Nosy_LevelFlagValue_Set__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var lvl slog.Level
		fill_err = tp.Fill(&lvl)
		if fill_err != nil {
			return
		}
		var value string
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}

		fv := NewLevelFlagValue(lvl)
		fv.Set(value)
	})
}

func Fuzz_Nosy_Writer_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *Writer
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		w.Close()
	})
}

func Fuzz_Nosy_Writer_Write__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *Writer
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		w.Write(b)
	})
}

func Fuzz_Nosy_FormatFlagValue_FormatType__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fmtType FormatType
		fill_err = tp.Fill(&fmtType)
		if fill_err != nil {
			return
		}

		fv := NewFormatFlagValue(fmtType)
		fv.FormatType()
	})
}

func Fuzz_Nosy_FormatFlagValue_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fmtType FormatType
		fill_err = tp.Fill(&fmtType)
		if fill_err != nil {
			return
		}

		fv := NewFormatFlagValue(fmtType)
		fv.String()
	})
}

func Fuzz_Nosy_FormatType_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ft FormatType
		fill_err = tp.Fill(&ft)
		if fill_err != nil {
			return
		}

		ft.String()
	})
}

func Fuzz_Nosy_LevelFlagValue_Level__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var lvl slog.Level
		fill_err = tp.Fill(&lvl)
		if fill_err != nil {
			return
		}

		fv := NewLevelFlagValue(lvl)
		fv.Level()
	})
}

func Fuzz_Nosy_LevelFlagValue_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var lvl slog.Level
		fill_err = tp.Fill(&lvl)
		if fill_err != nil {
			return
		}

		fv := NewLevelFlagValue(lvl)
		fv.String()
	})
}

// skipping Fuzz_Nosy_LvlSetter_SetLogLevel__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/log.LvlSetter

func Fuzz_Nosy_CLIFlags__(f *testing.F) {
	f.Fuzz(func(t *testing.T, envPrefix string) {
		CLIFlags(envPrefix)
	})
}

func Fuzz_Nosy_CLIFlagsWithCategory__(f *testing.F) {
	f.Fuzz(func(t *testing.T, envPrefix string, category string) {
		CLIFlagsWithCategory(envPrefix, category)
	})
}

func Fuzz_Nosy_FormatHandler__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ft FormatType
		fill_err = tp.Fill(&ft)
		if fill_err != nil {
			return
		}
		var color bool
		fill_err = tp.Fill(&color)
		if fill_err != nil {
			return
		}

		FormatHandler(ft, color)
	})
}

// skipping Fuzz_Nosy_SetGlobalLogHandler__ because parameters include func, chan, or unsupported interface: log/slog.Handler
