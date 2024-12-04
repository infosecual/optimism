package safedb

import (
	"context"
	"testing"

	"github.com/ethereum-optimism/optimism/op-service/eth"
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

func Fuzz_Nosy_DisabledDB_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *DisabledDB
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.Close()
	})
}

func Fuzz_Nosy_DisabledDB_Enabled__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *DisabledDB
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.Enabled()
	})
}

func Fuzz_Nosy_DisabledDB_SafeHeadAtL1__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *DisabledDB
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var _x3 uint64
		fill_err = tp.Fill(&_x3)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.SafeHeadAtL1(_x2, _x3)
	})
}

func Fuzz_Nosy_DisabledDB_SafeHeadReset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *DisabledDB
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var _x2 eth.L2BlockRef
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.SafeHeadReset(_x2)
	})
}

func Fuzz_Nosy_DisabledDB_SafeHeadUpdated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *DisabledDB
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var _x2 eth.L2BlockRef
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var _x3 eth.BlockID
		fill_err = tp.Fill(&_x3)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.SafeHeadUpdated(_x2, _x3)
	})
}

func Fuzz_Nosy_SafeDB_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *SafeDB
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.Close()
	})
}

func Fuzz_Nosy_SafeDB_Enabled__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *SafeDB
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.Enabled()
	})
}

func Fuzz_Nosy_SafeDB_SafeHeadAtL1__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *SafeDB
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var l1BlockNum uint64
		fill_err = tp.Fill(&l1BlockNum)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.SafeHeadAtL1(ctx, l1BlockNum)
	})
}

func Fuzz_Nosy_SafeDB_SafeHeadReset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *SafeDB
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var safeHead eth.L2BlockRef
		fill_err = tp.Fill(&safeHead)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.SafeHeadReset(safeHead)
	})
}

func Fuzz_Nosy_SafeDB_SafeHeadUpdated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *SafeDB
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var safeHead eth.L2BlockRef
		fill_err = tp.Fill(&safeHead)
		if fill_err != nil {
			return
		}
		var l1Head eth.BlockID
		fill_err = tp.Fill(&l1Head)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.SafeHeadUpdated(safeHead, l1Head)
	})
}

func Fuzz_Nosy_uint64Key_IterRange__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c uint64Key
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}

		c.IterRange()
	})
}

func Fuzz_Nosy_uint64Key_Max__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c uint64Key
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}

		c.Max()
	})
}

func Fuzz_Nosy_uint64Key_Of__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c uint64Key
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var num uint64
		fill_err = tp.Fill(&num)
		if fill_err != nil {
			return
		}

		c.Of(num)
	})
}

func Fuzz_Nosy_decodeSafeByL1BlockNum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, key []byte, val []byte) {
		decodeSafeByL1BlockNum(key, val)
	})
}

func Fuzz_Nosy_safeByL1BlockNumValue__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l1 eth.BlockID
		fill_err = tp.Fill(&l1)
		if fill_err != nil {
			return
		}
		var l2 eth.BlockID
		fill_err = tp.Fill(&l2)
		if fill_err != nil {
			return
		}

		safeByL1BlockNumValue(l1, l2)
	})
}
