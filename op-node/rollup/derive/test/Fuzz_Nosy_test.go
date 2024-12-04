package test

import (
	"math/rand"
	"testing"
	"time"

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

func Fuzz_Nosy_RandomL2Block__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rng *rand.Rand
		fill_err = tp.Fill(&rng)
		if fill_err != nil {
			return
		}
		var txCount int
		fill_err = tp.Fill(&txCount)
		if fill_err != nil {
			return
		}
		var t3 time.Time
		fill_err = tp.Fill(&t3)
		if fill_err != nil {
			return
		}
		if rng == nil {
			return
		}

		RandomL2Block(rng, txCount, t3)
	})
}
