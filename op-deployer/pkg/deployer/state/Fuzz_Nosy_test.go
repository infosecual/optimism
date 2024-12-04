package state

import (
	"testing"
	"github.com/ethereum-optimism/optimism/op-deployer/pkg/deployer"
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
	
	func Fuzz_Nosy_Base64Bytes_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var b *Base64Bytes
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var d2 []byte
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}
	if b == nil {
		return
	}

	b.UnmarshalJSON(d2)
	})
}

func Fuzz_Nosy_ChainIntent_Check__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ChainIntent
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Check()
	})
}

func Fuzz_Nosy_GzipData[T any]_MarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var g *GzipData[T]
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
	if g == nil {
		return
	}

	g.MarshalJSON()
	})
}

func Fuzz_Nosy_GzipData[T any]_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var g *GzipData[T]
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
	if g == nil {
		return
	}

	g.UnmarshalJSON(b)
	})
}

func Fuzz_Nosy_Intent_Chain__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *Intent
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var id common.Hash
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Chain(id)
	})
}

func Fuzz_Nosy_Intent_Check__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *Intent
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Check()
	})
}

func Fuzz_Nosy_Intent_L1ChainIDBig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *Intent
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.L1ChainIDBig()
	})
}

func Fuzz_Nosy_Intent_WriteToFile__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *Intent
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var path string
		fill_err = tp.Fill(&path)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.WriteToFile(path)
	})
}

func Fuzz_Nosy_Intent_checkL1Dev__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *Intent
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.checkL1Dev()
	})
}

func Fuzz_Nosy_Intent_checkL1Prod__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *Intent
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.checkL1Prod()
	})
}

func Fuzz_Nosy_Intent_checkL2Prod__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *Intent
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.checkL2Prod()
	})
}

func Fuzz_Nosy_State_Chain__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *State
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var id common.Hash
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
	if s == nil {
		return
	}

	s.Chain(id)
	})
}

func Fuzz_Nosy_State_WriteToFile__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *State
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var path string
		fill_err = tp.Fill(&path)
		if fill_err != nil {
			return
		}
	if s == nil {
		return
	}

	s.WriteToFile(path)
	})
}

func Fuzz_Nosy_Base64Bytes_MarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var b Base64Bytes
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

	b.MarshalJSON()
	})
}

func Fuzz_Nosy_DeploymentStrategy_Check__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var d DeploymentStrategy
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}

	d.Check()
	})
}

