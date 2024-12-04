package flags

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

func Fuzz_Nosy_DataAvailabilityType_Clone__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var kind *DataAvailabilityType
		fill_err = tp.Fill(&kind)
		if fill_err != nil {
			return
		}
		if kind == nil {
			return
		}

		kind.Clone()
	})
}

func Fuzz_Nosy_DataAvailabilityType_Set__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var kind *DataAvailabilityType
		fill_err = tp.Fill(&kind)
		if fill_err != nil {
			return
		}
		var value string
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
		if kind == nil {
			return
		}

		kind.Set(value)
	})
}

func Fuzz_Nosy_DataAvailabilityType_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var kind DataAvailabilityType
		fill_err = tp.Fill(&kind)
		if fill_err != nil {
			return
		}

		kind.String()
	})
}

func Fuzz_Nosy_ValidDataAvailabilityType__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var value DataAvailabilityType
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}

		ValidDataAvailabilityType(value)
	})
}

func Fuzz_Nosy_prefixEnvVars__(f *testing.F) {
	f.Fuzz(func(t *testing.T, name string) {
		prefixEnvVars(name)
	})
}
