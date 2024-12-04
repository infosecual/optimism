package versions

import (
	"io"
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

func Fuzz_Nosy_VersionedState_Deserialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var path string
		fill_err = tp.Fill(&path)
		if fill_err != nil {
			return
		}
		var in io.Reader
		fill_err = tp.Fill(&in)
		if fill_err != nil {
			return
		}

		s, err := LoadStateFromFile(path)
		if err != nil {
			return
		}
		s.Deserialize(in)
	})
}

func Fuzz_Nosy_VersionedState_MarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, path string) {
		s, err := LoadStateFromFile(path)
		if err != nil {
			return
		}
		s.MarshalJSON()
	})
}

func Fuzz_Nosy_VersionedState_Serialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var path string
		fill_err = tp.Fill(&path)
		if fill_err != nil {
			return
		}
		var w io.Writer
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}

		s, err := LoadStateFromFile(path)
		if err != nil {
			return
		}
		s.Serialize(w)
	})
}

func Fuzz_Nosy_StateVersion_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, path string) {
		s, err := DetectVersion(path)
		if err != nil {
			return
		}
		s.String()
	})
}
