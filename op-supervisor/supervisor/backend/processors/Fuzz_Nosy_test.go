package processors

import (
	"context"
	"testing"

	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum/go-ethereum/core/types"
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

func Fuzz_Nosy_ChainProcessor_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *ChainProcessor
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

func Fuzz_Nosy_ChainProcessor_OnNewHead__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *ChainProcessor
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var head eth.L1BlockRef
		fill_err = tp.Fill(&head)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.OnNewHead(head)
	})
}

func Fuzz_Nosy_ChainProcessor_ProcessToHead__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *ChainProcessor
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.ProcessToHead()
	})
}

// skipping Fuzz_Nosy_ChainProcessor_SetSource__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/processors.Source

func Fuzz_Nosy_ChainProcessor_StartBackground__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *ChainProcessor
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.StartBackground()
	})
}

func Fuzz_Nosy_ChainProcessor_nextNum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *ChainProcessor
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.nextNum()
	})
}

func Fuzz_Nosy_ChainProcessor_update__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *ChainProcessor
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var nextNum uint64
		fill_err = tp.Fill(&nextNum)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.update(nextNum)
	})
}

func Fuzz_Nosy_ChainProcessor_work__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *ChainProcessor
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.work()
	})
}

func Fuzz_Nosy_ChainProcessor_worker__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *ChainProcessor
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.worker()
	})
}

func Fuzz_Nosy_logProcessor_ProcessLogs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *logProcessor
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var block eth.L1BlockRef
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		var rcpts types.Receipts
		fill_err = tp.Fill(&rcpts)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.ProcessLogs(_x2, block, rcpts)
	})
}

// skipping Fuzz_Nosy_BlockProcessorFn_ProcessBlock__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/processors.BlockProcessorFn

// skipping Fuzz_Nosy_ChainsDBClientForLogProcessor_AddLog__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/processors.ChainsDBClientForLogProcessor

// skipping Fuzz_Nosy_ChainsDBClientForLogProcessor_SealBlock__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/processors.ChainsDBClientForLogProcessor

// skipping Fuzz_Nosy_DatabaseRewinder_LatestBlockNum__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/processors.DatabaseRewinder

// skipping Fuzz_Nosy_DatabaseRewinder_Rewind__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/processors.DatabaseRewinder

// skipping Fuzz_Nosy_LogProcessor_ProcessLogs__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/processors.LogProcessor

// skipping Fuzz_Nosy_LogStorage_AddLog__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/processors.LogStorage

// skipping Fuzz_Nosy_LogStorage_SealBlock__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/processors.LogStorage

// skipping Fuzz_Nosy_Source_FetchReceipts__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/processors.Source

// skipping Fuzz_Nosy_Source_L1BlockRefByNumber__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-supervisor/supervisor/backend/processors.Source
