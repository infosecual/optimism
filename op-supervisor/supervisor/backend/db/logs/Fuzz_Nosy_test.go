package logs

import (
	"io"
	"testing"

	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db/entrydb"
	"github.com/ethereum-optimism/optimism/op-supervisor/supervisor/types"
	"github.com/ethereum/go-ethereum/common"
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

func Fuzz_Nosy_DB_AddLog__(f *testing.F) {
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
		var logHash common.Hash
		fill_err = tp.Fill(&logHash)
		if fill_err != nil {
			return
		}
		var parentBlock eth.BlockID
		fill_err = tp.Fill(&parentBlock)
		if fill_err != nil {
			return
		}
		var logIdx uint32
		fill_err = tp.Fill(&logIdx)
		if fill_err != nil {
			return
		}
		var execMsg *types.ExecutingMessage
		fill_err = tp.Fill(&execMsg)
		if fill_err != nil {
			return
		}
		if db == nil || execMsg == nil {
			return
		}

		db.AddLog(logHash, parentBlock, logIdx, execMsg)
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

func Fuzz_Nosy_DB_Contains__(f *testing.F) {
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
		var blockNum uint64
		fill_err = tp.Fill(&blockNum)
		if fill_err != nil {
			return
		}
		var logIdx uint32
		fill_err = tp.Fill(&logIdx)
		if fill_err != nil {
			return
		}
		var logHash common.Hash
		fill_err = tp.Fill(&logHash)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		db.Contains(blockNum, logIdx, logHash)
	})
}

func Fuzz_Nosy_DB_FindSealedBlock__(f *testing.F) {
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
		var number uint64
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		db.FindSealedBlock(number)
	})
}

func Fuzz_Nosy_DB_Get__(f *testing.F) {
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
		var blockNum uint64
		fill_err = tp.Fill(&blockNum)
		if fill_err != nil {
			return
		}
		var logIdx uint32
		fill_err = tp.Fill(&logIdx)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		db.Get(blockNum, logIdx)
	})
}

func Fuzz_Nosy_DB_IteratorStartingAt__(f *testing.F) {
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
		var sealedNum uint64
		fill_err = tp.Fill(&sealedNum)
		if fill_err != nil {
			return
		}
		var logsSince uint32
		fill_err = tp.Fill(&logsSince)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		db.IteratorStartingAt(sealedNum, logsSince)
	})
}

func Fuzz_Nosy_DB_LatestSealedBlockNum__(f *testing.F) {
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

		db.LatestSealedBlockNum()
	})
}

func Fuzz_Nosy_DB_OpenBlock__(f *testing.F) {
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
		var blockNum uint64
		fill_err = tp.Fill(&blockNum)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		db.OpenBlock(blockNum)
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
		var newHeadBlockNum uint64
		fill_err = tp.Fill(&newHeadBlockNum)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		db.Rewind(newHeadBlockNum)
	})
}

func Fuzz_Nosy_DB_SealBlock__(f *testing.F) {
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
		var parentHash common.Hash
		fill_err = tp.Fill(&parentHash)
		if fill_err != nil {
			return
		}
		var block eth.BlockID
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		var timestamp uint64
		fill_err = tp.Fill(&timestamp)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		db.SealBlock(parentHash, block, timestamp)
	})
}

func Fuzz_Nosy_DB_StartingBlock__(f *testing.F) {
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

		db.StartingBlock()
	})
}

func Fuzz_Nosy_DB_debugTip__(f *testing.F) {
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

		db.debugTip()
	})
}

func Fuzz_Nosy_DB_findLogInfo__(f *testing.F) {
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
		var blockNum uint64
		fill_err = tp.Fill(&blockNum)
		if fill_err != nil {
			return
		}
		var logIdx uint32
		fill_err = tp.Fill(&logIdx)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		db.findLogInfo(blockNum, logIdx)
	})
}

func Fuzz_Nosy_DB_flush__(f *testing.F) {
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

		db.flush()
	})
}

func Fuzz_Nosy_DB_init__(f *testing.F) {
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
		var trimToLastSealed bool
		fill_err = tp.Fill(&trimToLastSealed)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		db.init(trimToLastSealed)
	})
}

func Fuzz_Nosy_DB_lastEntryIdx__(f *testing.F) {
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

		db.lastEntryIdx()
	})
}

func Fuzz_Nosy_DB_newIterator__(f *testing.F) {
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
		var index entrydb.EntryIdx
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		db.newIterator(index)
	})
}

func Fuzz_Nosy_DB_newIteratorAt__(f *testing.F) {
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
		var blockNum uint64
		fill_err = tp.Fill(&blockNum)
		if fill_err != nil {
			return
		}
		var logIndex uint32
		fill_err = tp.Fill(&logIndex)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		db.newIteratorAt(blockNum, logIndex)
	})
}

func Fuzz_Nosy_DB_readSearchCheckpoint__(f *testing.F) {
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
		var entryIdx entrydb.EntryIdx
		fill_err = tp.Fill(&entryIdx)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		db.readSearchCheckpoint(entryIdx)
	})
}

func Fuzz_Nosy_DB_searchCheckpoint__(f *testing.F) {
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
		var sealedBlockNum uint64
		fill_err = tp.Fill(&sealedBlockNum)
		if fill_err != nil {
			return
		}
		var logsSince uint32
		fill_err = tp.Fill(&logsSince)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		db.searchCheckpoint(sealedBlockNum, logsSince)
	})
}

func Fuzz_Nosy_DB_trimToLastSealed__(f *testing.F) {
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

		db.trimToLastSealed()
	})
}

func Fuzz_Nosy_DB_updateEntryCountMetric__(f *testing.F) {
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

		db.updateEntryCountMetric()
	})
}

func Fuzz_Nosy_EntryTypeFlag_Add__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *EntryTypeFlag
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		var v EntryTypeFlag
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.Add(v)
	})
}

func Fuzz_Nosy_EntryTypeFlag_Remove__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *EntryTypeFlag
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		var v EntryTypeFlag
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.Remove(v)
	})
}

func Fuzz_Nosy_iterator_End__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var i *iterator
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
		if i == nil {
			return
		}

		i.End()
	})
}

func Fuzz_Nosy_iterator_ExecMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var i *iterator
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
		if i == nil {
			return
		}

		i.ExecMessage()
	})
}

func Fuzz_Nosy_iterator_InitMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var i *iterator
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
		if i == nil {
			return
		}

		i.InitMessage()
	})
}

func Fuzz_Nosy_iterator_NextBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var i *iterator
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
		if i == nil {
			return
		}

		i.NextBlock()
	})
}

func Fuzz_Nosy_iterator_NextExecMsg__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var i *iterator
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
		if i == nil {
			return
		}

		i.NextExecMsg()
	})
}

func Fuzz_Nosy_iterator_NextIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var i *iterator
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
		if i == nil {
			return
		}

		i.NextIndex()
	})
}

func Fuzz_Nosy_iterator_NextInitMsg__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var i *iterator
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
		if i == nil {
			return
		}

		i.NextInitMsg()
	})
}

func Fuzz_Nosy_iterator_SealedBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var i *iterator
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
		if i == nil {
			return
		}

		i.SealedBlock()
	})
}

func Fuzz_Nosy_iterator_SealedTimestamp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var i *iterator
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
		if i == nil {
			return
		}

		i.SealedTimestamp()
	})
}

// skipping Fuzz_Nosy_iterator_TraverseConditional__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db/logs.traverseConditionalFn

func Fuzz_Nosy_iterator_next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var i *iterator
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
		if i == nil {
			return
		}

		i.next()
	})
}

func Fuzz_Nosy_logContext_ApplyEntry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *logContext
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		var entry Entry
		fill_err = tp.Fill(&entry)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		l.ApplyEntry(entry)
	})
}

func Fuzz_Nosy_logContext_ApplyLog__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *logContext
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		var parentBlock eth.BlockID
		fill_err = tp.Fill(&parentBlock)
		if fill_err != nil {
			return
		}
		var logIdx uint32
		fill_err = tp.Fill(&logIdx)
		if fill_err != nil {
			return
		}
		var logHash common.Hash
		fill_err = tp.Fill(&logHash)
		if fill_err != nil {
			return
		}
		var execMsg *types.ExecutingMessage
		fill_err = tp.Fill(&execMsg)
		if fill_err != nil {
			return
		}
		if l == nil || execMsg == nil {
			return
		}

		l.ApplyLog(parentBlock, logIdx, logHash, execMsg)
	})
}

func Fuzz_Nosy_logContext_ExecMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *logContext
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		l.ExecMessage()
	})
}

func Fuzz_Nosy_logContext_InitMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *logContext
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		l.InitMessage()
	})
}

func Fuzz_Nosy_logContext_NextIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *logContext
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		l.NextIndex()
	})
}

func Fuzz_Nosy_logContext_SealBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *logContext
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		var parent common.Hash
		fill_err = tp.Fill(&parent)
		if fill_err != nil {
			return
		}
		var upd eth.BlockID
		fill_err = tp.Fill(&upd)
		if fill_err != nil {
			return
		}
		var timestamp uint64
		fill_err = tp.Fill(&timestamp)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		l.SealBlock(parent, upd, timestamp)
	})
}

func Fuzz_Nosy_logContext_SealedBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *logContext
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		l.SealedBlock()
	})
}

func Fuzz_Nosy_logContext_SealedTimestamp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *logContext
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		l.SealedTimestamp()
	})
}

// skipping Fuzz_Nosy_logContext_appendEntry__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db/logs.EntryObj

func Fuzz_Nosy_logContext_forceBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *logContext
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		var upd eth.BlockID
		fill_err = tp.Fill(&upd)
		if fill_err != nil {
			return
		}
		var timestamp uint64
		fill_err = tp.Fill(&timestamp)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		l.forceBlock(upd, timestamp)
	})
}

func Fuzz_Nosy_logContext_hasCompleteBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *logContext
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		l.hasCompleteBlock()
	})
}

func Fuzz_Nosy_logContext_hasIncompleteLog__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *logContext
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		l.hasIncompleteLog()
	})
}

func Fuzz_Nosy_logContext_hasReadableLog__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *logContext
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		l.hasReadableLog()
	})
}

func Fuzz_Nosy_logContext_infer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *logContext
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		l.infer()
	})
}

func Fuzz_Nosy_logContext_inferFull__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *logContext
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		l.inferFull()
	})
}

func Fuzz_Nosy_logContext_processEntry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *logContext
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		var entry Entry
		fill_err = tp.Fill(&entry)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		l.processEntry(entry)
	})
}

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

// skipping Fuzz_Nosy_EntryObj_encode__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db/logs.EntryObj

func Fuzz_Nosy_EntryType_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x EntryType
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.String()
	})
}

func Fuzz_Nosy_EntryTypeFlag_Any__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x EntryTypeFlag
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		var v EntryTypeFlag
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}

		x.Any(v)
	})
}

func Fuzz_Nosy_EntryTypeFlag_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x EntryTypeFlag
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.String()
	})
}

// skipping Fuzz_Nosy_Iterator_End__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db/logs.Iterator

// skipping Fuzz_Nosy_Iterator_NextBlock__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db/logs.Iterator

// skipping Fuzz_Nosy_Iterator_NextExecMsg__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db/logs.Iterator

// skipping Fuzz_Nosy_Iterator_NextInitMsg__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db/logs.Iterator

// skipping Fuzz_Nosy_Iterator_TraverseConditional__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db/logs.Iterator

// skipping Fuzz_Nosy_IteratorState_ExecMessage__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db/logs.IteratorState

// skipping Fuzz_Nosy_IteratorState_InitMessage__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db/logs.IteratorState

// skipping Fuzz_Nosy_IteratorState_NextIndex__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db/logs.IteratorState

// skipping Fuzz_Nosy_IteratorState_SealedBlock__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db/logs.IteratorState

// skipping Fuzz_Nosy_IteratorState_SealedTimestamp__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db/logs.IteratorState

// skipping Fuzz_Nosy_Metrics_RecordDBEntryCount__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db/logs.Metrics

// skipping Fuzz_Nosy_Metrics_RecordDBSearchEntriesRead__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/db/logs.Metrics

func Fuzz_Nosy_canonicalHash_encode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d1 Entry
		fill_err = tp.Fill(&d1)
		if fill_err != nil {
			return
		}

		c, err := newCanonicalHashFromEntry(d1)
		if err != nil {
			return
		}
		c.encode()
	})
}

func Fuzz_Nosy_executingCheck_encode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d1 Entry
		fill_err = tp.Fill(&d1)
		if fill_err != nil {
			return
		}

		e, err := newExecutingCheckFromEntry(d1)
		if err != nil {
			return
		}
		e.encode()
	})
}

func Fuzz_Nosy_executingLink_encode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg types.ExecutingMessage
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}

		e, err := newExecutingLink(msg)
		if err != nil {
			return
		}
		e.encode()
	})
}

func Fuzz_Nosy_initiatingEvent_encode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var logHash common.Hash
		fill_err = tp.Fill(&logHash)
		if fill_err != nil {
			return
		}
		var hasExecMsg bool
		fill_err = tp.Fill(&hasExecMsg)
		if fill_err != nil {
			return
		}

		i := newInitiatingEvent(logHash, hasExecMsg)
		i.encode()
	})
}

func Fuzz_Nosy_paddingEntry_encode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e paddingEntry
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}

		e.encode()
	})
}

func Fuzz_Nosy_searchCheckpoint_encode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d1 Entry
		fill_err = tp.Fill(&d1)
		if fill_err != nil {
			return
		}

		s, err := newSearchCheckpointFromEntry(d1)
		if err != nil {
			return
		}
		s.encode()
	})
}
