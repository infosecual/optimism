package main

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

func Fuzz_Nosy_checkUnusedImports__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var imports []string
		fill_err = tp.Fill(&imports)
		if fill_err != nil {
			return
		}
		var content string
		fill_err = tp.Fill(&content)
		if fill_err != nil {
			return
		}

		checkUnusedImports(imports, content)
	})
}

func Fuzz_Nosy_findImports__(f *testing.F) {
	f.Fuzz(func(t *testing.T, content string) {
		findImports(content)
	})
}

func Fuzz_Nosy_isImportUsed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, imp string, content string) {
		isImportUsed(imp, content)
	})
}

// skipping Fuzz_Nosy_processFile__ because parameters include func, chan, or unsupported interface: func(string, ...any)

// skipping Fuzz_Nosy_writeStderr__ because parameters include func, chan, or unsupported interface: []any
