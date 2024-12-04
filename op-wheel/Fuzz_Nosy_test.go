package wheel

import (
	"testing"
	"github.com/ethereum-optimism/optimism"
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
	
	func Fuzz_Nosy_TextFlag[T Text]_Clone__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var a *TextFlag[T]
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
	if a == nil {
		return
	}

	a.Clone()
	})
}

func Fuzz_Nosy_TextFlag[T Text]_Get__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var a *TextFlag[T]
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
	if a == nil {
		return
	}

	a.Get()
	})
}

func Fuzz_Nosy_TextFlag[T Text]_Set__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var a *TextFlag[T]
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var value string
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
	if a == nil {
		return
	}

	a.Set(value)
	})
}

func Fuzz_Nosy_TextFlag[T Text]_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var a *TextFlag[T]
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
	if a == nil {
		return
	}

	a.String()
	})
}

func Fuzz_Nosy_prefixEnvVars__(f *testing.F) {
	f.Fuzz(func(t *testing.T, name string) {
	prefixEnvVars(name)
	})
}

// skipping Fuzz_Nosy_withEngineFlags__ because parameters include func, chan, or unsupported interface: []github.com/urfave/cli/v2.Flag

