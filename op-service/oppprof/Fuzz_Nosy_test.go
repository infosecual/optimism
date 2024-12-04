package oppprof

import (
	"context"
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

func Fuzz_Nosy_Service_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var listenEnabled bool
		fill_err = tp.Fill(&listenEnabled)
		if fill_err != nil {
			return
		}
		var listenAddr string
		fill_err = tp.Fill(&listenAddr)
		if fill_err != nil {
			return
		}
		var listenPort int
		fill_err = tp.Fill(&listenPort)
		if fill_err != nil {
			return
		}
		var profType profileType
		fill_err = tp.Fill(&profType)
		if fill_err != nil {
			return
		}
		var profileDir string
		fill_err = tp.Fill(&profileDir)
		if fill_err != nil {
			return
		}
		var profileFilename string
		fill_err = tp.Fill(&profileFilename)
		if fill_err != nil {
			return
		}

		s := New(listenEnabled, listenAddr, listenPort, profType, profileDir, profileFilename)
		s.Start()
	})
}

func Fuzz_Nosy_Service_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var listenEnabled bool
		fill_err = tp.Fill(&listenEnabled)
		if fill_err != nil {
			return
		}
		var listenAddr string
		fill_err = tp.Fill(&listenAddr)
		if fill_err != nil {
			return
		}
		var listenPort int
		fill_err = tp.Fill(&listenPort)
		if fill_err != nil {
			return
		}
		var profType profileType
		fill_err = tp.Fill(&profType)
		if fill_err != nil {
			return
		}
		var profileDir string
		fill_err = tp.Fill(&profileDir)
		if fill_err != nil {
			return
		}
		var profileFilename string
		fill_err = tp.Fill(&profileFilename)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		s := New(listenEnabled, listenAddr, listenPort, profType, profileDir, profileFilename)
		s.Stop(ctx)
	})
}

func Fuzz_Nosy_Service_buildTargetFilePath__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var listenEnabled bool
		fill_err = tp.Fill(&listenEnabled)
		if fill_err != nil {
			return
		}
		var listenAddr string
		fill_err = tp.Fill(&listenAddr)
		if fill_err != nil {
			return
		}
		var listenPort int
		fill_err = tp.Fill(&listenPort)
		if fill_err != nil {
			return
		}
		var profType profileType
		fill_err = tp.Fill(&profType)
		if fill_err != nil {
			return
		}
		var profileDir string
		fill_err = tp.Fill(&profileDir)
		if fill_err != nil {
			return
		}
		var profileFilename string
		fill_err = tp.Fill(&profileFilename)
		if fill_err != nil {
			return
		}

		s := New(listenEnabled, listenAddr, listenPort, profType, profileDir, profileFilename)
		s.buildTargetFilePath()
	})
}

func Fuzz_Nosy_Service_startCPUProfile__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var listenEnabled bool
		fill_err = tp.Fill(&listenEnabled)
		if fill_err != nil {
			return
		}
		var listenAddr string
		fill_err = tp.Fill(&listenAddr)
		if fill_err != nil {
			return
		}
		var listenPort int
		fill_err = tp.Fill(&listenPort)
		if fill_err != nil {
			return
		}
		var profType profileType
		fill_err = tp.Fill(&profType)
		if fill_err != nil {
			return
		}
		var profileDir string
		fill_err = tp.Fill(&profileDir)
		if fill_err != nil {
			return
		}
		var profileFilename string
		fill_err = tp.Fill(&profileFilename)
		if fill_err != nil {
			return
		}

		s := New(listenEnabled, listenAddr, listenPort, profType, profileDir, profileFilename)
		s.startCPUProfile()
	})
}

func Fuzz_Nosy_Service_startServer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var listenEnabled bool
		fill_err = tp.Fill(&listenEnabled)
		if fill_err != nil {
			return
		}
		var listenAddr string
		fill_err = tp.Fill(&listenAddr)
		if fill_err != nil {
			return
		}
		var listenPort int
		fill_err = tp.Fill(&listenPort)
		if fill_err != nil {
			return
		}
		var profType profileType
		fill_err = tp.Fill(&profType)
		if fill_err != nil {
			return
		}
		var profileDir string
		fill_err = tp.Fill(&profileDir)
		if fill_err != nil {
			return
		}
		var profileFilename string
		fill_err = tp.Fill(&profileFilename)
		if fill_err != nil {
			return
		}

		s := New(listenEnabled, listenAddr, listenPort, profType, profileDir, profileFilename)
		s.startServer()
	})
}

func Fuzz_Nosy_profileType_Clone__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *profileType
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Clone()
	})
}

func Fuzz_Nosy_profileType_Set__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *profileType
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var value string
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Set(value)
	})
}

func Fuzz_Nosy_profileType_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 profileType
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}

		t1.String()
	})
}

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

func Fuzz_Nosy_validProfileType__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var value profileType
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}

		validProfileType(value)
	})
}
