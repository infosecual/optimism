package config

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

func Fuzz_Nosy_AllocType_Check__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a AllocType
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}

		a.Check()
	})
}

func Fuzz_Nosy_AllocType_UsesProofs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a AllocType
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}

		a.UsesProofs()
	})
}

func Fuzz_Nosy_initAllocType__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var root string
		fill_err = tp.Fill(&root)
		if fill_err != nil {
			return
		}
		var allocType AllocType
		fill_err = tp.Fill(&allocType)
		if fill_err != nil {
			return
		}

		initAllocType(root, allocType)
	})
}
