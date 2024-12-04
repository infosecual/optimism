package mpt

import (
	"bytes"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
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

func Fuzz_Nosy_DB_Ancient__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *DB
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var kind string
		fill_err = tp.Fill(&kind)
		if fill_err != nil {
			return
		}
		var number uint64
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.Ancient(kind, number)
	})
}

func Fuzz_Nosy_DB_AncientDatadir__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *DB
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.AncientDatadir()
	})
}

func Fuzz_Nosy_DB_AncientRange__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *DB
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var kind string
		fill_err = tp.Fill(&kind)
		if fill_err != nil {
			return
		}
		var start uint64
		fill_err = tp.Fill(&start)
		if fill_err != nil {
			return
		}
		var count uint64
		fill_err = tp.Fill(&count)
		if fill_err != nil {
			return
		}
		var maxBytes uint64
		fill_err = tp.Fill(&maxBytes)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.AncientRange(kind, start, count, maxBytes)
	})
}

func Fuzz_Nosy_DB_AncientSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *DB
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var kind string
		fill_err = tp.Fill(&kind)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.AncientSize(kind)
	})
}

func Fuzz_Nosy_DB_Ancients__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *DB
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.Ancients()
	})
}

func Fuzz_Nosy_DB_Get__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *DB
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var key []byte
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.Get(key)
	})
}

func Fuzz_Nosy_DB_Has__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *DB
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var key []byte
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.Has(key)
	})
}

func Fuzz_Nosy_DB_HasAncient__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *DB
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var kind string
		fill_err = tp.Fill(&kind)
		if fill_err != nil {
			return
		}
		var number uint64
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.HasAncient(kind, number)
	})
}

// skipping Fuzz_Nosy_DB_MigrateTable__ because parameters include func, chan, or unsupported interface: func([]byte) ([]byte, error)

// skipping Fuzz_Nosy_DB_ModifyAncients__ because parameters include func, chan, or unsupported interface: func(github.com/ethereum/go-ethereum/ethdb.AncientWriteOp) error

func Fuzz_Nosy_DB_Put__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *DB
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var key []byte
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var value []byte
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.Put(key, value)
	})
}

// skipping Fuzz_Nosy_DB_ReadAncients__ because parameters include func, chan, or unsupported interface: func(github.com/ethereum/go-ethereum/ethdb.AncientReaderOp) error

func Fuzz_Nosy_DB_Sync__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *DB
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.Sync()
	})
}

func Fuzz_Nosy_DB_Tail__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *DB
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.Tail()
	})
}

func Fuzz_Nosy_DB_TruncateHead__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *DB
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var n uint64
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.TruncateHead(n)
	})
}

func Fuzz_Nosy_DB_TruncateTail__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *DB
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var n uint64
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.TruncateTail(n)
	})
}

func Fuzz_Nosy_DB_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p DB
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}

		p.Close()
	})
}

func Fuzz_Nosy_DB_Compact__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p DB
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var start []byte
		fill_err = tp.Fill(&start)
		if fill_err != nil {
			return
		}
		var limit []byte
		fill_err = tp.Fill(&limit)
		if fill_err != nil {
			return
		}

		p.Compact(start, limit)
	})
}

func Fuzz_Nosy_DB_Delete__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p DB
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var key []byte
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}

		p.Delete(key)
	})
}

func Fuzz_Nosy_DB_NewBatch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p DB
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}

		p.NewBatch()
	})
}

func Fuzz_Nosy_DB_NewBatchWithSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p DB
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var size int
		fill_err = tp.Fill(&size)
		if fill_err != nil {
			return
		}

		p.NewBatchWithSize(size)
	})
}

func Fuzz_Nosy_DB_NewIterator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p DB
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var prefix []byte
		fill_err = tp.Fill(&prefix)
		if fill_err != nil {
			return
		}
		var start []byte
		fill_err = tp.Fill(&start)
		if fill_err != nil {
			return
		}

		p.NewIterator(prefix, start)
	})
}

func Fuzz_Nosy_DB_Stat__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p DB
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}

		p.Stat()
	})
}

func Fuzz_Nosy_noResetHasher_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n noResetHasher
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}

		n.Reset()
	})
}

func Fuzz_Nosy_rawList_EncodeIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r rawList
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var i int
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
		var buf *bytes.Buffer
		fill_err = tp.Fill(&buf)
		if fill_err != nil {
			return
		}
		if buf == nil {
			return
		}

		r.EncodeIndex(i, buf)
	})
}

func Fuzz_Nosy_rawList_Len__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r rawList
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}

		r.Len()
	})
}

// skipping Fuzz_Nosy_ReadTrie__ because parameters include func, chan, or unsupported interface: func(key github.com/ethereum/go-ethereum/common.Hash) []byte

func Fuzz_Nosy_WriteTrie__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var values []hexutil.Bytes
		fill_err = tp.Fill(&values)
		if fill_err != nil {
			return
		}

		WriteTrie(values)
	})
}
