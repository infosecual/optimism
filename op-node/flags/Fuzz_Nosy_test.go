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

func Fuzz_Nosy_P2PFlags__(f *testing.F) {
	f.Fuzz(func(t *testing.T, envPrefix string) {
		P2PFlags(envPrefix)
	})
}

func Fuzz_Nosy_deprecatedP2PFlags__(f *testing.F) {
	f.Fuzz(func(t *testing.T, envPrefix string) {
		deprecatedP2PFlags(envPrefix)
	})
}

func Fuzz_Nosy_p2pEnv__(f *testing.F) {
	f.Fuzz(func(t *testing.T, envprefix string, v string) {
		p2pEnv(envprefix, v)
	})
}

func Fuzz_Nosy_prefixEnvVars__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var names []string
		fill_err = tp.Fill(&names)
		if fill_err != nil {
			return
		}

		prefixEnvVars(names...)
	})
}
