package crossdomain

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
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

func Fuzz_Nosy_CrossDomainMessage_Encode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var nonce *big.Int
		fill_err = tp.Fill(&nonce)
		if fill_err != nil {
			return
		}
		var sender common.Address
		fill_err = tp.Fill(&sender)
		if fill_err != nil {
			return
		}
		var t3 common.Address
		fill_err = tp.Fill(&t3)
		if fill_err != nil {
			return
		}
		var value *big.Int
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
		var gasLimit *big.Int
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		var d6 []byte
		fill_err = tp.Fill(&d6)
		if fill_err != nil {
			return
		}
		if nonce == nil || value == nil || gasLimit == nil {
			return
		}

		c := NewCrossDomainMessage(nonce, sender, t3, value, gasLimit, d6)
		c.Encode()
	})
}

func Fuzz_Nosy_CrossDomainMessage_Hash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var nonce *big.Int
		fill_err = tp.Fill(&nonce)
		if fill_err != nil {
			return
		}
		var sender common.Address
		fill_err = tp.Fill(&sender)
		if fill_err != nil {
			return
		}
		var t3 common.Address
		fill_err = tp.Fill(&t3)
		if fill_err != nil {
			return
		}
		var value *big.Int
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
		var gasLimit *big.Int
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		var d6 []byte
		fill_err = tp.Fill(&d6)
		if fill_err != nil {
			return
		}
		if nonce == nil || value == nil || gasLimit == nil {
			return
		}

		c := NewCrossDomainMessage(nonce, sender, t3, value, gasLimit, d6)
		c.Hash()
	})
}

func Fuzz_Nosy_CrossDomainMessage_HashV1__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var nonce *big.Int
		fill_err = tp.Fill(&nonce)
		if fill_err != nil {
			return
		}
		var sender common.Address
		fill_err = tp.Fill(&sender)
		if fill_err != nil {
			return
		}
		var t3 common.Address
		fill_err = tp.Fill(&t3)
		if fill_err != nil {
			return
		}
		var value *big.Int
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
		var gasLimit *big.Int
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		var d6 []byte
		fill_err = tp.Fill(&d6)
		if fill_err != nil {
			return
		}
		if nonce == nil || value == nil || gasLimit == nil {
			return
		}

		c := NewCrossDomainMessage(nonce, sender, t3, value, gasLimit, d6)
		c.HashV1()
	})
}

func Fuzz_Nosy_CrossDomainMessage_Version__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var nonce *big.Int
		fill_err = tp.Fill(&nonce)
		if fill_err != nil {
			return
		}
		var sender common.Address
		fill_err = tp.Fill(&sender)
		if fill_err != nil {
			return
		}
		var t3 common.Address
		fill_err = tp.Fill(&t3)
		if fill_err != nil {
			return
		}
		var value *big.Int
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
		var gasLimit *big.Int
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		var d6 []byte
		fill_err = tp.Fill(&d6)
		if fill_err != nil {
			return
		}
		if nonce == nil || value == nil || gasLimit == nil {
			return
		}

		c := NewCrossDomainMessage(nonce, sender, t3, value, gasLimit, d6)
		c.Version()
	})
}

func Fuzz_Nosy_LegacyReceipt_DecodeRLP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *LegacyReceipt
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var s *rlp.Stream
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if r == nil || s == nil {
			return
		}

		r.DecodeRLP(s)
	})
}

func Fuzz_Nosy_Withdrawal_Decode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var nonce *big.Int
		fill_err = tp.Fill(&nonce)
		if fill_err != nil {
			return
		}
		var sender *common.Address
		fill_err = tp.Fill(&sender)
		if fill_err != nil {
			return
		}
		var t3 *common.Address
		fill_err = tp.Fill(&t3)
		if fill_err != nil {
			return
		}
		var value *big.Int
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
		var gasLimit *big.Int
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		var d6 []byte
		fill_err = tp.Fill(&d6)
		if fill_err != nil {
			return
		}
		var d7 []byte
		fill_err = tp.Fill(&d7)
		if fill_err != nil {
			return
		}
		if nonce == nil || sender == nil || t3 == nil || value == nil || gasLimit == nil {
			return
		}

		w := NewWithdrawal(nonce, sender, t3, value, gasLimit, d6)
		w.Decode(d7)
	})
}

func Fuzz_Nosy_Withdrawal_Encode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var nonce *big.Int
		fill_err = tp.Fill(&nonce)
		if fill_err != nil {
			return
		}
		var sender *common.Address
		fill_err = tp.Fill(&sender)
		if fill_err != nil {
			return
		}
		var t3 *common.Address
		fill_err = tp.Fill(&t3)
		if fill_err != nil {
			return
		}
		var value *big.Int
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
		var gasLimit *big.Int
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		var d6 []byte
		fill_err = tp.Fill(&d6)
		if fill_err != nil {
			return
		}
		if nonce == nil || sender == nil || t3 == nil || value == nil || gasLimit == nil {
			return
		}

		w := NewWithdrawal(nonce, sender, t3, value, gasLimit, d6)
		w.Encode()
	})
}

func Fuzz_Nosy_Withdrawal_Hash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var nonce *big.Int
		fill_err = tp.Fill(&nonce)
		if fill_err != nil {
			return
		}
		var sender *common.Address
		fill_err = tp.Fill(&sender)
		if fill_err != nil {
			return
		}
		var t3 *common.Address
		fill_err = tp.Fill(&t3)
		if fill_err != nil {
			return
		}
		var value *big.Int
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
		var gasLimit *big.Int
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		var d6 []byte
		fill_err = tp.Fill(&d6)
		if fill_err != nil {
			return
		}
		if nonce == nil || sender == nil || t3 == nil || value == nil || gasLimit == nil {
			return
		}

		w := NewWithdrawal(nonce, sender, t3, value, gasLimit, d6)
		w.Hash()
	})
}

func Fuzz_Nosy_Withdrawal_StorageSlot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var nonce *big.Int
		fill_err = tp.Fill(&nonce)
		if fill_err != nil {
			return
		}
		var sender *common.Address
		fill_err = tp.Fill(&sender)
		if fill_err != nil {
			return
		}
		var t3 *common.Address
		fill_err = tp.Fill(&t3)
		if fill_err != nil {
			return
		}
		var value *big.Int
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
		var gasLimit *big.Int
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		var d6 []byte
		fill_err = tp.Fill(&d6)
		if fill_err != nil {
			return
		}
		if nonce == nil || sender == nil || t3 == nil || value == nil || gasLimit == nil {
			return
		}

		w := NewWithdrawal(nonce, sender, t3, value, gasLimit, d6)
		w.StorageSlot()
	})
}

func Fuzz_Nosy_Withdrawal_WithdrawalTransaction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var nonce *big.Int
		fill_err = tp.Fill(&nonce)
		if fill_err != nil {
			return
		}
		var sender *common.Address
		fill_err = tp.Fill(&sender)
		if fill_err != nil {
			return
		}
		var t3 *common.Address
		fill_err = tp.Fill(&t3)
		if fill_err != nil {
			return
		}
		var value *big.Int
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
		var gasLimit *big.Int
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		var d6 []byte
		fill_err = tp.Fill(&d6)
		if fill_err != nil {
			return
		}
		if nonce == nil || sender == nil || t3 == nil || value == nil || gasLimit == nil {
			return
		}

		w := NewWithdrawal(nonce, sender, t3, value, gasLimit, d6)
		w.WithdrawalTransaction()
	})
}

// skipping Fuzz_Nosy_WithdrawalMessage_Decode__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-chain-ops/crossdomain.WithdrawalMessage

// skipping Fuzz_Nosy_WithdrawalMessage_Encode__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-chain-ops/crossdomain.WithdrawalMessage

// skipping Fuzz_Nosy_WithdrawalMessage_Hash__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-chain-ops/crossdomain.WithdrawalMessage

// skipping Fuzz_Nosy_WithdrawalMessage_StorageSlot__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-chain-ops/crossdomain.WithdrawalMessage

func Fuzz_Nosy_DecodeVersionedNonce__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var versioned *big.Int
		fill_err = tp.Fill(&versioned)
		if fill_err != nil {
			return
		}
		if versioned == nil {
			return
		}

		DecodeVersionedNonce(versioned)
	})
}

func Fuzz_Nosy_EncodeCrossDomainMessageV0__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 common.Address
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var sender common.Address
		fill_err = tp.Fill(&sender)
		if fill_err != nil {
			return
		}
		var message []byte
		fill_err = tp.Fill(&message)
		if fill_err != nil {
			return
		}
		var nonce *big.Int
		fill_err = tp.Fill(&nonce)
		if fill_err != nil {
			return
		}
		if nonce == nil {
			return
		}

		EncodeCrossDomainMessageV0(t1, sender, message, nonce)
	})
}

func Fuzz_Nosy_EncodeCrossDomainMessageV1__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var nonce *big.Int
		fill_err = tp.Fill(&nonce)
		if fill_err != nil {
			return
		}
		var sender common.Address
		fill_err = tp.Fill(&sender)
		if fill_err != nil {
			return
		}
		var t3 common.Address
		fill_err = tp.Fill(&t3)
		if fill_err != nil {
			return
		}
		var value *big.Int
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
		var gasLimit *big.Int
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		var d6 []byte
		fill_err = tp.Fill(&d6)
		if fill_err != nil {
			return
		}
		if nonce == nil || value == nil || gasLimit == nil {
			return
		}

		EncodeCrossDomainMessageV1(nonce, sender, t3, value, gasLimit, d6)
	})
}

// skipping Fuzz_Nosy_ReadLegacyReceipts__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.Reader
