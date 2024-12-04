package devkeys

import (
	"math/big"
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

func Fuzz_Nosy_ChainOperatorRole_MarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var role *ChainOperatorRole
		fill_err = tp.Fill(&role)
		if fill_err != nil {
			return
		}
		if role == nil {
			return
		}

		role.MarshalText()
	})
}

func Fuzz_Nosy_ChainOperatorRole_UnmarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var role *ChainOperatorRole
		fill_err = tp.Fill(&role)
		if fill_err != nil {
			return
		}
		var d2 []byte
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}
		if role == nil {
			return
		}

		role.UnmarshalText(d2)
	})
}

// skipping Fuzz_Nosy_MnemonicDevKeys_Address__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-chain-ops/devkeys.Key

// skipping Fuzz_Nosy_MnemonicDevKeys_Secret__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-chain-ops/devkeys.Key

func Fuzz_Nosy_SuperchainOperatorRole_MarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var role *SuperchainOperatorRole
		fill_err = tp.Fill(&role)
		if fill_err != nil {
			return
		}
		if role == nil {
			return
		}

		role.MarshalText()
	})
}

func Fuzz_Nosy_SuperchainOperatorRole_UnmarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var role *SuperchainOperatorRole
		fill_err = tp.Fill(&role)
		if fill_err != nil {
			return
		}
		var d2 []byte
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}
		if role == nil {
			return
		}

		role.UnmarshalText(d2)
	})
}

// skipping Fuzz_Nosy_Addresses_Address__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-chain-ops/devkeys.Addresses

func Fuzz_Nosy_ChainOperatorKey_HDPath__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k ChainOperatorKey
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}

		k.HDPath()
	})
}

func Fuzz_Nosy_ChainOperatorKey_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k ChainOperatorKey
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}

		k.String()
	})
}

func Fuzz_Nosy_ChainOperatorRole_Key__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var role ChainOperatorRole
		fill_err = tp.Fill(&role)
		if fill_err != nil {
			return
		}
		var chainID *big.Int
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		if chainID == nil {
			return
		}

		role.Key(chainID)
	})
}

func Fuzz_Nosy_ChainOperatorRole_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var role ChainOperatorRole
		fill_err = tp.Fill(&role)
		if fill_err != nil {
			return
		}

		role.String()
	})
}

func Fuzz_Nosy_ChainUserKey_HDPath__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k ChainUserKey
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}

		k.HDPath()
	})
}

func Fuzz_Nosy_ChainUserKey_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k ChainUserKey
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}

		k.String()
	})
}

// skipping Fuzz_Nosy_Key_HDPath__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-chain-ops/devkeys.Key

// skipping Fuzz_Nosy_Key_String__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-chain-ops/devkeys.Key

// skipping Fuzz_Nosy_Role_Key__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-chain-ops/devkeys.Role

// skipping Fuzz_Nosy_Secrets_Secret__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-chain-ops/devkeys.Secrets

func Fuzz_Nosy_SuperchainOperatorKey_HDPath__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k SuperchainOperatorKey
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}

		k.HDPath()
	})
}

func Fuzz_Nosy_SuperchainOperatorKey_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k SuperchainOperatorKey
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}

		k.String()
	})
}

func Fuzz_Nosy_SuperchainOperatorRole_Key__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var role SuperchainOperatorRole
		fill_err = tp.Fill(&role)
		if fill_err != nil {
			return
		}
		var chainID *big.Int
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		if chainID == nil {
			return
		}

		role.Key(chainID)
	})
}

func Fuzz_Nosy_SuperchainOperatorRole_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var role SuperchainOperatorRole
		fill_err = tp.Fill(&role)
		if fill_err != nil {
			return
		}

		role.String()
	})
}

func Fuzz_Nosy_UserKey_HDPath__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k UserKey
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}

		k.HDPath()
	})
}

func Fuzz_Nosy_UserKey_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k UserKey
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}

		k.String()
	})
}

func Fuzz_Nosy_ChainOperatorKeys__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainID *big.Int
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		if chainID == nil {
			return
		}

		ChainOperatorKeys(chainID)
	})
}

func Fuzz_Nosy_ChainUserKeys__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainID *big.Int
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		if chainID == nil {
			return
		}

		ChainUserKeys(chainID)
	})
}

func Fuzz_Nosy_SuperchainOperatorKeys__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainID *big.Int
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		if chainID == nil {
			return
		}

		SuperchainOperatorKeys(chainID)
	})
}
