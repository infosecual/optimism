package fromda

import (
	"io"
	"testing"

	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db/entrydb"
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

func Fuzz_Nosy_DB_AddDerived__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *DB
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var derivedFrom eth.L1BlockRef
		fill_err = tp.Fill(&derivedFrom)
		if fill_err != nil {
			return
		}
		var derived eth.L1BlockRef
		fill_err = tp.Fill(&derived)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		db.AddDerived(derivedFrom, derived)
	})
}

func Fuzz_Nosy_DB_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *DB
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		db.Close()
	})
}

func Fuzz_Nosy_DB_DerivedFrom__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *DB
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var derived eth.BlockID
		fill_err = tp.Fill(&derived)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		db.DerivedFrom(derived)
	})
}

func Fuzz_Nosy_DB_First__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *DB
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		db.First()
	})
}

func Fuzz_Nosy_DB_FirstAfter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *DB
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var derivedFrom eth.BlockID
		fill_err = tp.Fill(&derivedFrom)
		if fill_err != nil {
			return
		}
		var derived eth.BlockID
		fill_err = tp.Fill(&derived)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		db.FirstAfter(derivedFrom, derived)
	})
}

func Fuzz_Nosy_DB_LastDerivedAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *DB
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var derivedFrom eth.BlockID
		fill_err = tp.Fill(&derivedFrom)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		db.LastDerivedAt(derivedFrom)
	})
}

func Fuzz_Nosy_DB_Latest__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *DB
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		db.Latest()
	})
}

func Fuzz_Nosy_DB_NextDerived__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *DB
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var derived eth.BlockID
		fill_err = tp.Fill(&derived)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		db.NextDerived(derived)
	})
}

func Fuzz_Nosy_DB_NextDerivedFrom__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *DB
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var derivedFrom eth.BlockID
		fill_err = tp.Fill(&derivedFrom)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		db.NextDerivedFrom(derivedFrom)
	})
}

func Fuzz_Nosy_DB_PreviousDerived__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *DB
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var derived eth.BlockID
		fill_err = tp.Fill(&derived)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		db.PreviousDerived(derived)
	})
}

func Fuzz_Nosy_DB_PreviousDerivedFrom__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *DB
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var derivedFrom eth.BlockID
		fill_err = tp.Fill(&derivedFrom)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		db.PreviousDerivedFrom(derivedFrom)
	})
}

func Fuzz_Nosy_DB_Rewind__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *DB
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var derivedFrom uint64
		fill_err = tp.Fill(&derivedFrom)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		db.Rewind(derivedFrom)
	})
}

// skipping Fuzz_Nosy_DB_find__ because parameters include func, chan, or unsupported interface: func(link github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db/fromda.LinkEntry) int

func Fuzz_Nosy_DB_firstDerivedAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *DB
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var derivedFrom uint64
		fill_err = tp.Fill(&derivedFrom)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		db.firstDerivedAt(derivedFrom)
	})
}

func Fuzz_Nosy_DB_firstDerivedFrom__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *DB
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var derived uint64
		fill_err = tp.Fill(&derived)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		db.firstDerivedFrom(derived)
	})
}

func Fuzz_Nosy_DB_lastDerivedAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *DB
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var derivedFrom uint64
		fill_err = tp.Fill(&derivedFrom)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		db.lastDerivedAt(derivedFrom)
	})
}

func Fuzz_Nosy_DB_lastDerivedFrom__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *DB
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var derived uint64
		fill_err = tp.Fill(&derived)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		db.lastDerivedFrom(derived)
	})
}

func Fuzz_Nosy_DB_latest__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *DB
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		db.latest()
	})
}

func Fuzz_Nosy_DB_lookup__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *DB
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var derivedFrom uint64
		fill_err = tp.Fill(&derivedFrom)
		if fill_err != nil {
			return
		}
		var derived uint64
		fill_err = tp.Fill(&derived)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		db.lookup(derivedFrom, derived)
	})
}

func Fuzz_Nosy_DB_readAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *DB
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var i entrydb.EntryIdx
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		db.readAt(i)
	})
}

func Fuzz_Nosy_LinkEntry_decode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *LinkEntry
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var e Entry
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.decode(e)
	})
}

func Fuzz_Nosy_LinkEntry_encode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *LinkEntry
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.encode()
	})
}

func Fuzz_Nosy_delegate_RecordDBDerivedEntryCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *delegate
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var count int64
		fill_err = tp.Fill(&count)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.RecordDBDerivedEntryCount(count)
	})
}

// skipping Fuzz_Nosy_ChainMetrics_RecordDBEntryCount__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db/fromda.ChainMetrics

func Fuzz_Nosy_Entry_Type__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e Entry
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}

		e.Type()
	})
}

func Fuzz_Nosy_EntryBinary_Append__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 EntryBinary
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var dest []byte
		fill_err = tp.Fill(&dest)
		if fill_err != nil {
			return
		}
		var e *Entry
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		_x1.Append(dest, e)
	})
}

func Fuzz_Nosy_EntryBinary_EntrySize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 EntryBinary
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}

		_x1.EntrySize()
	})
}

func Fuzz_Nosy_EntryBinary_ReadAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 EntryBinary
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var dest *Entry
		fill_err = tp.Fill(&dest)
		if fill_err != nil {
			return
		}
		var r io.ReaderAt
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var at int64
		fill_err = tp.Fill(&at)
		if fill_err != nil {
			return
		}
		if dest == nil {
			return
		}

		_x1.ReadAt(dest, r, at)
	})
}

// skipping Fuzz_Nosy_EntryStore_Append__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db/fromda.EntryStore

// skipping Fuzz_Nosy_EntryStore_Close__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db/fromda.EntryStore

// skipping Fuzz_Nosy_EntryStore_LastEntryIdx__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db/fromda.EntryStore

// skipping Fuzz_Nosy_EntryStore_Read__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db/fromda.EntryStore

// skipping Fuzz_Nosy_EntryStore_Size__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db/fromda.EntryStore

// skipping Fuzz_Nosy_EntryStore_Truncate__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db/fromda.EntryStore

func Fuzz_Nosy_EntryType_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s EntryType
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}

		s.String()
	})
}

func Fuzz_Nosy_LinkEntry_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d LinkEntry
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}

		d.String()
	})
}

// skipping Fuzz_Nosy_Metrics_RecordDBDerivedEntryCount__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db/fromda.Metrics
