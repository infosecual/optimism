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

func Fuzz_Nosy_getAllContractsSources__(f *testing.F) {
	f.Fuzz(func(t *testing.T, srcDir string) {
		getAllContractsSources(srcDir)
	})
}

// skipping Fuzz_Nosy_jsonEqual__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_parseArtifactName__(f *testing.F) {
	f.Fuzz(func(t *testing.T, artifactVersionFile string) {
		parseArtifactName(artifactVersionFile)
	})
}
