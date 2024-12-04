package test

import (
	"testing"

	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum/go-ethereum/beacon/engine"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
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

// skipping Fuzz_Nosy_testHelper_Log__ because parameters include func, chan, or unsupported interface: []any

func Fuzz_Nosy_testHelper_addBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *testHelper
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var txs []*types.Transaction
		fill_err = tp.Fill(&txs)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.addBlock(txs...)
	})
}

func Fuzz_Nosy_testHelper_addBlockWithParent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *testHelper
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var head *types.Header
		fill_err = tp.Fill(&head)
		if fill_err != nil {
			return
		}
		var timestamp hexutil.Uint64
		fill_err = tp.Fill(&timestamp)
		if fill_err != nil {
			return
		}
		var txs []*types.Transaction
		fill_err = tp.Fill(&txs)
		if fill_err != nil {
			return
		}
		if h == nil || head == nil {
			return
		}

		h.addBlockWithParent(head, timestamp, txs...)
	})
}

func Fuzz_Nosy_testHelper_callNewPayload__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *testHelper
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var envelope *eth.ExecutionPayloadEnvelope
		fill_err = tp.Fill(&envelope)
		if fill_err != nil {
			return
		}
		if h == nil || envelope == nil {
			return
		}

		h.callNewPayload(envelope)
	})
}

func Fuzz_Nosy_testHelper_finalHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *testHelper
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.finalHash()
	})
}

func Fuzz_Nosy_testHelper_forkChoiceUpdated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *testHelper
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var head common.Hash
		fill_err = tp.Fill(&head)
		if fill_err != nil {
			return
		}
		var safe common.Hash
		fill_err = tp.Fill(&safe)
		if fill_err != nil {
			return
		}
		var finalized common.Hash
		fill_err = tp.Fill(&finalized)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.forkChoiceUpdated(head, safe, finalized)
	})
}

func Fuzz_Nosy_testHelper_getPayload__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *testHelper
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var id *engine.PayloadID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		if h == nil || id == nil {
			return
		}

		h.getPayload(id)
	})
}

func Fuzz_Nosy_testHelper_headHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *testHelper
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.headHash()
	})
}

func Fuzz_Nosy_testHelper_newPayload__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *testHelper
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var envelope *eth.ExecutionPayloadEnvelope
		fill_err = tp.Fill(&envelope)
		if fill_err != nil {
			return
		}
		if h == nil || envelope == nil {
			return
		}

		h.newPayload(envelope)
	})
}

func Fuzz_Nosy_testHelper_safeHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *testHelper
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.safeHash()
	})
}

func Fuzz_Nosy_testHelper_startBlockBuilding__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *testHelper
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var head *types.Header
		fill_err = tp.Fill(&head)
		if fill_err != nil {
			return
		}
		var newBlockTimestamp hexutil.Uint64
		fill_err = tp.Fill(&newBlockTimestamp)
		if fill_err != nil {
			return
		}
		var txs []*types.Transaction
		fill_err = tp.Fill(&txs)
		if fill_err != nil {
			return
		}
		if h == nil || head == nil {
			return
		}

		h.startBlockBuilding(head, newBlockTimestamp, txs...)
	})
}

// skipping Fuzz_Nosy_RunEngineAPITests__ because parameters include func, chan, or unsupported interface: func(t *testing.T) github.com/ethereum-optimism/optimism/op-program/client/l2/engineapi.EngineBackend

func Fuzz_Nosy_updateBlockHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var envelope *eth.ExecutionPayloadEnvelope
		fill_err = tp.Fill(&envelope)
		if fill_err != nil {
			return
		}
		if envelope == nil {
			return
		}

		updateBlockHash(envelope)
	})
}
