package sender

import (
	"testing"

	"github.com/ethereum-optimism/optimism/op-service/txmgr"
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

func Fuzz_Nosy_TxSender_From__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *TxSender
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.From()
	})
}

func Fuzz_Nosy_TxSender_SendAndWaitDetailed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *TxSender
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var txPurpose string
		fill_err = tp.Fill(&txPurpose)
		if fill_err != nil {
			return
		}
		var txs []txmgr.TxCandidate
		fill_err = tp.Fill(&txs)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.SendAndWaitDetailed(txPurpose, txs...)
	})
}

func Fuzz_Nosy_TxSender_SendAndWaitSimple__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *TxSender
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var txPurpose string
		fill_err = tp.Fill(&txPurpose)
		if fill_err != nil {
			return
		}
		var txs []txmgr.TxCandidate
		fill_err = tp.Fill(&txs)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.SendAndWaitSimple(txPurpose, txs...)
	})
}
