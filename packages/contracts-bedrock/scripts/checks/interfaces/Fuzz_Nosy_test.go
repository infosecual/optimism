package main

import (
	"encoding/json"
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

// skipping Fuzz_Nosy_abiItemLess__ because parameters include func, chan, or unsupported interface: map[string]interface{}

func Fuzz_Nosy_compareABIs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var abi1 json.RawMessage
		fill_err = tp.Fill(&abi1)
		if fill_err != nil {
			return
		}
		var abi2 json.RawMessage
		fill_err = tp.Fill(&abi2)
		if fill_err != nil {
			return
		}

		compareABIs(abi1, abi2)
	})
}

func Fuzz_Nosy_getContractSemver__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var artifact *Artifact
		fill_err = tp.Fill(&artifact)
		if fill_err != nil {
			return
		}
		if artifact == nil {
			return
		}

		getContractSemver(artifact)
	})
}

// skipping Fuzz_Nosy_getString__ because parameters include func, chan, or unsupported interface: map[string]interface{}

func Fuzz_Nosy_glob__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dir string, ext string) {
		glob(dir, ext)
	})
}

func Fuzz_Nosy_isExcluded__(f *testing.F) {
	f.Fuzz(func(t *testing.T, contractName string) {
		isExcluded(contractName)
	})
}

// skipping Fuzz_Nosy_normalizeABIItem__ because parameters include func, chan, or unsupported interface: map[string]interface{}

func Fuzz_Nosy_normalizeInternalType__(f *testing.F) {
	f.Fuzz(func(t *testing.T, internalType string) {
		normalizeInternalType(internalType)
	})
}

// skipping Fuzz_Nosy_writeStderr__ because parameters include func, chan, or unsupported interface: []any
