package locks

import (
	"testing"
	"github.com/ethereum-optimism/optimism/op-service"
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
	
	func Fuzz_Nosy_RWMap[K comparable, V any]_Clear__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var m *RWMap[K, V]
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
	if m == nil {
		return
	}

	m.Clear()
	})
}

// skipping Fuzz_Nosy_RWMap[K comparable, V any]_Get__ because parameters include func, chan, or unsupported interface: K

// skipping Fuzz_Nosy_RWMap[K comparable, V any]_Has__ because parameters include func, chan, or unsupported interface: K

// skipping Fuzz_Nosy_RWMap[K comparable, V any]_Range__ because parameters include func, chan, or unsupported interface: func(key K, value V) bool

// skipping Fuzz_Nosy_RWMap[K comparable, V any]_Set__ because parameters include func, chan, or unsupported interface: K

func Fuzz_Nosy_RWValue[E any]_Get__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *RWValue[E]
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Get()
	})
}

// skipping Fuzz_Nosy_RWValue[E any]_Set__ because parameters include func, chan, or unsupported interface: E

