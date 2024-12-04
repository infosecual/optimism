package jsonutil

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

// skipping Fuzz_Nosy_LazySortedJsonMap[K comparable, V any]_UnmarshalJSON__ because parameters include func, chan, or unsupported interface: *github.com/ethereum-optimism/optimism/op-service/jsonutil.LazySortedJsonMap[K, V]

// skipping Fuzz_Nosy_jsonDecoder_Decode__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_jsonEncoder_Encode__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_tomlDecoder_Decode__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_tomlEncoder_Encode__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_Decoder_Decode__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_Encoder_Encode__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_LazySortedJsonMap[K comparable, V any]_MarshalJSON__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/jsonutil.LazySortedJsonMap[K, V]

func Fuzz_Nosy_LoadJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, inputPath string) {
		LoadJSON(inputPath)
	})
}

func Fuzz_Nosy_LoadTOML__(f *testing.F) {
	f.Fuzz(func(t *testing.T, inputPath string) {
		LoadTOML(inputPath)
	})
}

// skipping Fuzz_Nosy_MergeJSON__ because parameters include func, chan, or unsupported interface: T

// skipping Fuzz_Nosy_load__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-service/jsonutil.DecoderFactory
