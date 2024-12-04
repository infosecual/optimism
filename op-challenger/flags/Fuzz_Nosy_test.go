package flags

import (
	"testing"

	"github.com/ethereum-optimism/optimism/op-challenger/game/fault/types"
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

func Fuzz_Nosy_VMFlag_DefaultName__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *VMFlag
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.DefaultName()
	})
}

func Fuzz_Nosy_VMFlag_Flags__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *VMFlag
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.Flags()
	})
}

func Fuzz_Nosy_VMFlag_IsSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *VMFlag
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var ctx *cli.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var vm types.TraceType
		fill_err = tp.Fill(&vm)
		if fill_err != nil {
			return
		}
		if f1 == nil || ctx == nil {
			return
		}

		f1.IsSet(ctx, vm)
	})
}

func Fuzz_Nosy_VMFlag_SourceFlagName__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *VMFlag
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var ctx *cli.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var vm types.TraceType
		fill_err = tp.Fill(&vm)
		if fill_err != nil {
			return
		}
		if f1 == nil || ctx == nil {
			return
		}

		f1.SourceFlagName(ctx, vm)
	})
}

func Fuzz_Nosy_VMFlag_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *VMFlag
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var ctx *cli.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var vm types.TraceType
		fill_err = tp.Fill(&vm)
		if fill_err != nil {
			return
		}
		if f1 == nil || ctx == nil {
			return
		}

		f1.String(ctx, vm)
	})
}

func Fuzz_Nosy_VMFlag_flagName__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *VMFlag
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var vm types.TraceType
		fill_err = tp.Fill(&vm)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.flagName(vm)
	})
}

// skipping Fuzz_Nosy_getL2Rpc__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/log.Logger

func Fuzz_Nosy_parseTraceTypes__(f *testing.F) {
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

		parseTraceTypes(ctx)
	})
}

func Fuzz_Nosy_prefixEnvVars__(f *testing.F) {
	f.Fuzz(func(t *testing.T, name string) {
		prefixEnvVars(name)
	})
}
