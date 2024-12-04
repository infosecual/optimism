package entrydb

import (
	"testing"
	"github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db"
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
	
	// skipping Fuzz_Nosy_EntryDB[T EntryType, E Entry[T], B Binary[T, E]]_Append__ because parameters include func, chan, or unsupported interface: []E

func Fuzz_Nosy_EntryDB[T EntryType, E Entry[T], B Binary[T, E]]_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var e *EntryDB[T, E, B]
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
	if e == nil {
		return
	}

	e.Close()
	})
}

func Fuzz_Nosy_EntryDB[T EntryType, E Entry[T], B Binary[T, E]]_LastEntryIdx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var e *EntryDB[T, E, B]
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
	if e == nil {
		return
	}

	e.LastEntryIdx()
	})
}

func Fuzz_Nosy_EntryDB[T EntryType, E Entry[T], B Binary[T, E]]_Read__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var e *EntryDB[T, E, B]
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var idx EntryIdx
		fill_err = tp.Fill(&idx)
		if fill_err != nil {
			return
		}
	if e == nil {
		return
	}

	e.Read(idx)
	})
}

func Fuzz_Nosy_EntryDB[T EntryType, E Entry[T], B Binary[T, E]]_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var e *EntryDB[T, E, B]
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
	if e == nil {
		return
	}

	e.Size()
	})
}

func Fuzz_Nosy_EntryDB[T EntryType, E Entry[T], B Binary[T, E]]_Truncate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var e *EntryDB[T, E, B]
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var idx EntryIdx
		fill_err = tp.Fill(&idx)
		if fill_err != nil {
			return
		}
	if e == nil {
		return
	}

	e.Truncate(idx)
	})
}

func Fuzz_Nosy_EntryDB[T EntryType, E Entry[T], B Binary[T, E]]_recover__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var e *EntryDB[T, E, B]
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
	if e == nil {
		return
	}

	e.recover()
	})
}

// skipping Fuzz_Nosy_MemEntryStore[T EntryType, E Entry[T]]_Append__ because parameters include func, chan, or unsupported interface: []E

func Fuzz_Nosy_MemEntryStore[T EntryType, E Entry[T]]_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *MemEntryStore[T, E]
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
	if s == nil {
		return
	}

	s.Close()
	})
}

func Fuzz_Nosy_MemEntryStore[T EntryType, E Entry[T]]_LastEntryIdx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *MemEntryStore[T, E]
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
	if s == nil {
		return
	}

	s.LastEntryIdx()
	})
}

func Fuzz_Nosy_MemEntryStore[T EntryType, E Entry[T]]_Read__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *MemEntryStore[T, E]
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var idx EntryIdx
		fill_err = tp.Fill(&idx)
		if fill_err != nil {
			return
		}
	if s == nil {
		return
	}

	s.Read(idx)
	})
}

func Fuzz_Nosy_MemEntryStore[T EntryType, E Entry[T]]_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *MemEntryStore[T, E]
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
	if s == nil {
		return
	}

	s.Size()
	})
}

func Fuzz_Nosy_MemEntryStore[T EntryType, E Entry[T]]_Truncate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *MemEntryStore[T, E]
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var idx EntryIdx
		fill_err = tp.Fill(&idx)
		if fill_err != nil {
			return
		}
	if s == nil {
		return
	}

	s.Truncate(idx)
	})
}

// skipping Fuzz_Nosy_Binary[T EntryType, E Entry[T]]_Append__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db/entrydb.Binary[T github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db/entrydb.EntryType, E github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db/entrydb.Entry[T]]

// skipping Fuzz_Nosy_Binary[T EntryType, E Entry[T]]_EntrySize__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db/entrydb.Binary[T github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db/entrydb.EntryType, E github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db/entrydb.Entry[T]]

// skipping Fuzz_Nosy_Binary[T EntryType, E Entry[T]]_ReadAt__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db/entrydb.Binary[T github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db/entrydb.EntryType, E github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db/entrydb.Entry[T]]

// skipping Fuzz_Nosy_EntryStore[T EntryType, E Entry[T]]_Append__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db/entrydb.EntryStore[T github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db/entrydb.EntryType, E github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db/entrydb.Entry[T]]

// skipping Fuzz_Nosy_EntryStore[T EntryType, E Entry[T]]_Close__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db/entrydb.EntryStore[T github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db/entrydb.EntryType, E github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db/entrydb.Entry[T]]

// skipping Fuzz_Nosy_EntryStore[T EntryType, E Entry[T]]_LastEntryIdx__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db/entrydb.EntryStore[T github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db/entrydb.EntryType, E github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db/entrydb.Entry[T]]

// skipping Fuzz_Nosy_EntryStore[T EntryType, E Entry[T]]_Read__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db/entrydb.EntryStore[T github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db/entrydb.EntryType, E github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db/entrydb.Entry[T]]

// skipping Fuzz_Nosy_EntryStore[T EntryType, E Entry[T]]_Size__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db/entrydb.EntryStore[T github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db/entrydb.EntryType, E github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db/entrydb.Entry[T]]

// skipping Fuzz_Nosy_EntryStore[T EntryType, E Entry[T]]_Truncate__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db/entrydb.EntryStore[T github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db/entrydb.EntryType, E github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db/entrydb.Entry[T]]

// skipping Fuzz_Nosy_EntryType_String__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db/entrydb.EntryType

// skipping Fuzz_Nosy_Entry[T EntryType]_Type__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db/entrydb.Entry[T github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db/entrydb.EntryType]

// skipping Fuzz_Nosy_dataAccess_Truncate__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db/entrydb.dataAccess

