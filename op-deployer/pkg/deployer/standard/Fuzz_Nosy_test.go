package standard

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

func Fuzz_Nosy_ChainNameFor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, chainID uint64) {
		ChainNameFor(chainID)
	})
}

func Fuzz_Nosy_CommitForDeployTag__(f *testing.F) {
	f.Fuzz(func(t *testing.T, tag string) {
		CommitForDeployTag(tag)
	})
}

func Fuzz_Nosy_L1VersionsDataFor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, chainID uint64) {
		L1VersionsDataFor(chainID)
	})
}

func Fuzz_Nosy_standardArtifactsURL__(f *testing.F) {
	f.Fuzz(func(t *testing.T, checksum string) {
		standardArtifactsURL(checksum)
	})
}
