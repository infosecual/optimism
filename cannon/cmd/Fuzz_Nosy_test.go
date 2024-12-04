package cmd

import (
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

func Fuzz_Nosy_ProcessPreimageOracle_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *ProcessPreimageOracle
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.Close()
	})
}

func Fuzz_Nosy_ProcessPreimageOracle_GetPreimage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *ProcessPreimageOracle
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var k [32]byte
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.GetPreimage(k)
	})
}

func Fuzz_Nosy_ProcessPreimageOracle_Hint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *ProcessPreimageOracle
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var v []byte
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.Hint(v)
	})
}

func Fuzz_Nosy_ProcessPreimageOracle_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *ProcessPreimageOracle
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.Start()
	})
}

func Fuzz_Nosy_ProcessPreimageOracle_wait__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *ProcessPreimageOracle
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.wait()
	})
}

func Fuzz_Nosy_StepMatcherFlag_Clone__(f *testing.F) {
	f.Fuzz(func(t *testing.T, pattern string) {
		m := MustStepMatcherFlag(pattern)
		m.Clone()
	})
}

func Fuzz_Nosy_StepMatcherFlag_Matcher__(f *testing.F) {
	f.Fuzz(func(t *testing.T, pattern string) {
		m := MustStepMatcherFlag(pattern)
		m.Matcher()
	})
}

func Fuzz_Nosy_StepMatcherFlag_Set__(f *testing.F) {
	f.Fuzz(func(t *testing.T, pattern string, value string) {
		m := MustStepMatcherFlag(pattern)
		m.Set(value)
	})
}

func Fuzz_Nosy_StepMatcherFlag_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, pattern string) {
		m := MustStepMatcherFlag(pattern)
		m.String()
	})
}

func Fuzz_Nosy_leveler_Level__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *leveler
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		l.Level()
	})
}

// skipping Fuzz_Nosy_VMState_GetStep__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/cannon/cmd.VMState

func Fuzz_Nosy_rawHint_Hint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rh rawHint
		fill_err = tp.Fill(&rh)
		if fill_err != nil {
			return
		}

		rh.Hint()
	})
}

func Fuzz_Nosy_rawKey_PreimageKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rk rawKey
		fill_err = tp.Fill(&rk)
		if fill_err != nil {
			return
		}

		rk.PreimageKey()
	})
}
