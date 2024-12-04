package compressor

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

func Fuzz_Nosy_NonCompressor_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *NonCompressor
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Close()
	})
}

func Fuzz_Nosy_NonCompressor_Flush__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *NonCompressor
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Flush()
	})
}

func Fuzz_Nosy_NonCompressor_FullErr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *NonCompressor
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.FullErr()
	})
}

func Fuzz_Nosy_NonCompressor_Len__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *NonCompressor
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Len()
	})
}

func Fuzz_Nosy_NonCompressor_Read__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *NonCompressor
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var p []byte
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Read(p)
	})
}

func Fuzz_Nosy_NonCompressor_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *NonCompressor
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Reset()
	})
}

func Fuzz_Nosy_NonCompressor_Write__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *NonCompressor
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var p []byte
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Write(p)
	})
}

func Fuzz_Nosy_RatioCompressor_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *RatioCompressor
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Close()
	})
}

func Fuzz_Nosy_RatioCompressor_Flush__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *RatioCompressor
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Flush()
	})
}

func Fuzz_Nosy_RatioCompressor_FullErr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *RatioCompressor
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.FullErr()
	})
}

func Fuzz_Nosy_RatioCompressor_InputThreshold__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *RatioCompressor
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.InputThreshold()
	})
}

func Fuzz_Nosy_RatioCompressor_Len__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *RatioCompressor
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Len()
	})
}

func Fuzz_Nosy_RatioCompressor_Read__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *RatioCompressor
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var p []byte
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Read(p)
	})
}

func Fuzz_Nosy_RatioCompressor_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *RatioCompressor
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Reset()
	})
}

func Fuzz_Nosy_RatioCompressor_Write__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *RatioCompressor
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var p []byte
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Write(p)
	})
}

func Fuzz_Nosy_RatioCompressor_inputTargetReached__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *RatioCompressor
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.inputTargetReached()
	})
}

func Fuzz_Nosy_ShadowCompressor_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *ShadowCompressor
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Close()
	})
}

func Fuzz_Nosy_ShadowCompressor_Flush__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *ShadowCompressor
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Flush()
	})
}

func Fuzz_Nosy_ShadowCompressor_FullErr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *ShadowCompressor
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.FullErr()
	})
}

func Fuzz_Nosy_ShadowCompressor_Len__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *ShadowCompressor
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Len()
	})
}

func Fuzz_Nosy_ShadowCompressor_Read__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *ShadowCompressor
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var p []byte
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Read(p)
	})
}

func Fuzz_Nosy_ShadowCompressor_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *ShadowCompressor
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Reset()
	})
}

func Fuzz_Nosy_ShadowCompressor_Write__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *ShadowCompressor
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var p []byte
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Write(p)
	})
}

func Fuzz_Nosy_Config_NewCompressor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c Config
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}

		c.NewCompressor()
	})
}
