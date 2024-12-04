package main

import (
	"testing"

	"github.com/ethereum-optimism/optimism/cannon/mipsevm/versions"
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

func Fuzz_Nosy_artifact_isValid__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a artifact
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}

		a.isValid()
	})
}

func Fuzz_Nosy_extractTempFile__(f *testing.F) {
	f.Fuzz(func(t *testing.T, name string, d2 []byte) {
		extractTempFile(name, d2)
	})
}

func Fuzz_Nosy_parseFlag__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var args []string
		fill_err = tp.Fill(&args)
		if fill_err != nil {
			return
		}
		var flag string
		fill_err = tp.Fill(&flag)
		if fill_err != nil {
			return
		}

		parseFlag(args, flag)
	})
}

func Fuzz_Nosy_parsePathFlag__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var args []string
		fill_err = tp.Fill(&args)
		if fill_err != nil {
			return
		}
		var flag string
		fill_err = tp.Fill(&flag)
		if fill_err != nil {
			return
		}

		parsePathFlag(args, flag)
	})
}

func Fuzz_Nosy_vmFilename__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ver versions.StateVersion
		fill_err = tp.Fill(&ver)
		if fill_err != nil {
			return
		}

		vmFilename(ver)
	})
}
