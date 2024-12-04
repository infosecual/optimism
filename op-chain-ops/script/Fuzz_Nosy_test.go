package script

import (
	"testing"
	"github.com/ethereum-optimism/optimism/op-chain-ops"
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
	
	func Fuzz_Nosy_AccessControlledPrecompile_RequiredGas__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *AccessControlledPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.RequiredGas(input)
	})
}

func Fuzz_Nosy_AccessControlledPrecompile_Run__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *AccessControlledPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Run(input)
	})
}

func Fuzz_Nosy_BroadcastType_UnmarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var bt *BroadcastType
		fill_err = tp.Fill(&bt)
		if fill_err != nil {
			return
		}
		var d2 []byte
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}
	if bt == nil {
		return
	}

	bt.UnmarshalText(d2)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_Accesses__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Accesses()
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_ActiveFork__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.ActiveFork()
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_Addr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var privateKey *big.Int
		fill_err = tp.Fill(&privateKey)
		if fill_err != nil {
			return
		}
	if c == nil|| privateKey == nil {
		return
	}

	c.Addr(privateKey)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_AllowCheatcodes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.AllowCheatcodes(addr)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_Breakpoint_f0259e92__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Breakpoint_f0259e92(name)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_Breakpoint_f7d39a8d__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var v bool
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Breakpoint_f7d39a8d(name, v)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_Broadcast_afc98040__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Broadcast_afc98040()
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_Broadcast_e6962cdb__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var who common.Address
		fill_err = tp.Fill(&who)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Broadcast_e6962cdb(who)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_ChainId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var id *big.Int
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
	if c == nil|| id == nil {
		return
	}

	c.ChainId(id)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_ClearMockedCalls__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.ClearMockedCalls()
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_Coinbase__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Coinbase(addr)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_ComputeCreate2Address_890c283b__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var salt [32]byte
		fill_err = tp.Fill(&salt)
		if fill_err != nil {
			return
		}
		var codeHash [32]byte
		fill_err = tp.Fill(&codeHash)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.ComputeCreate2Address_890c283b(salt, codeHash)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_CreateFork_31ba3498__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var urlOrAlias string
		fill_err = tp.Fill(&urlOrAlias)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.CreateFork_31ba3498(urlOrAlias)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_CreateFork_6ba3ba2b__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var urlOrAlias string
		fill_err = tp.Fill(&urlOrAlias)
		if fill_err != nil {
			return
		}
		var block *big.Int
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
	if c == nil|| block == nil {
		return
	}

	c.CreateFork_6ba3ba2b(urlOrAlias, block)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_CreateFork_7ca29682__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var urlOrAlias string
		fill_err = tp.Fill(&urlOrAlias)
		if fill_err != nil {
			return
		}
		var txHash common.Hash
		fill_err = tp.Fill(&txHash)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.CreateFork_7ca29682(urlOrAlias, txHash)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_CreateSelectFork_71ee464d__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var urlOrAlias string
		fill_err = tp.Fill(&urlOrAlias)
		if fill_err != nil {
			return
		}
		var block *big.Int
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
	if c == nil|| block == nil {
		return
	}

	c.CreateSelectFork_71ee464d(urlOrAlias, block)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_CreateSelectFork_84d52b7a__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var urlOrAlias string
		fill_err = tp.Fill(&urlOrAlias)
		if fill_err != nil {
			return
		}
		var txHash common.Hash
		fill_err = tp.Fill(&txHash)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.CreateSelectFork_84d52b7a(urlOrAlias, txHash)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_CreateSelectFork_98680034__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var urlOrAlias string
		fill_err = tp.Fill(&urlOrAlias)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.CreateSelectFork_98680034(urlOrAlias)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_Deal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var who common.Address
		fill_err = tp.Fill(&who)
		if fill_err != nil {
			return
		}
		var newBalance *big.Int
		fill_err = tp.Fill(&newBalance)
		if fill_err != nil {
			return
		}
	if c == nil|| newBalance == nil {
		return
	}

	c.Deal(who, newBalance)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_DeriveKey_6229498b__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var mnemonic string
		fill_err = tp.Fill(&mnemonic)
		if fill_err != nil {
			return
		}
		var index uint32
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.DeriveKey_6229498b(mnemonic, index)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_DeriveKey_6bcb2c1b__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var mnemonic string
		fill_err = tp.Fill(&mnemonic)
		if fill_err != nil {
			return
		}
		var path string
		fill_err = tp.Fill(&path)
		if fill_err != nil {
			return
		}
		var index uint32
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.DeriveKey_6bcb2c1b(mnemonic, path, index)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_Difficulty__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var _x2 *big.Int
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
	if c == nil|| _x2 == nil {
		return
	}

	c.Difficulty(_x2)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_DumpState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var pathToStateJson string
		fill_err = tp.Fill(&pathToStateJson)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.DumpState(pathToStateJson)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_EnvAddress_350d56bf__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var key string
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.EnvAddress_350d56bf(key)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_EnvAddress_ad31b9fa__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var key string
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var delimiter string
		fill_err = tp.Fill(&delimiter)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.EnvAddress_ad31b9fa(key, delimiter)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_EnvBool_7ed1ec7d__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var key string
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.EnvBool_7ed1ec7d(key)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_EnvBool_aaaddeaf__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var key string
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var delimiter string
		fill_err = tp.Fill(&delimiter)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.EnvBool_aaaddeaf(key, delimiter)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_EnvBytes32_5af231c1__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var key string
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var delimiter string
		fill_err = tp.Fill(&delimiter)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.EnvBytes32_5af231c1(key, delimiter)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_EnvBytes32_97949042__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var key string
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.EnvBytes32_97949042(key)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_EnvBytes_4d7baf06__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var key string
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.EnvBytes_4d7baf06(key)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_EnvBytes_ddc2651b__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var key string
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var delimiter string
		fill_err = tp.Fill(&delimiter)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.EnvBytes_ddc2651b(key, delimiter)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_EnvInt_42181150__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var key string
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var delimiter string
		fill_err = tp.Fill(&delimiter)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.EnvInt_42181150(key, delimiter)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_EnvInt_892a0c61__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var key string
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.EnvInt_892a0c61(key)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_EnvOr_2281f367__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var key string
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var delimiter string
		fill_err = tp.Fill(&delimiter)
		if fill_err != nil {
			return
		}
		var defaultValue [][32]byte
		fill_err = tp.Fill(&defaultValue)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.EnvOr_2281f367(key, delimiter, defaultValue)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_EnvOr_4700d74b__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var key string
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var delimiter string
		fill_err = tp.Fill(&delimiter)
		if fill_err != nil {
			return
		}
		var defaultValue []*ABIInt256
		fill_err = tp.Fill(&defaultValue)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.EnvOr_4700d74b(key, delimiter, defaultValue)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_EnvOr_4777f3cf__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var key string
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var defaultValue bool
		fill_err = tp.Fill(&defaultValue)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.EnvOr_4777f3cf(key, defaultValue)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_EnvOr_561fe540__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var key string
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var defaultValue common.Address
		fill_err = tp.Fill(&defaultValue)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.EnvOr_561fe540(key, defaultValue)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_EnvOr_5e97348f__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var key string
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var defaultValue *big.Int
		fill_err = tp.Fill(&defaultValue)
		if fill_err != nil {
			return
		}
	if c == nil|| defaultValue == nil {
		return
	}

	c.EnvOr_5e97348f(key, defaultValue)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_EnvOr_64bc3e64__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var key string
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var delimiter string
		fill_err = tp.Fill(&delimiter)
		if fill_err != nil {
			return
		}
		var defaultValue [][]byte
		fill_err = tp.Fill(&defaultValue)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.EnvOr_64bc3e64(key, delimiter, defaultValue)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_EnvOr_74318528__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var key string
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var delimiter string
		fill_err = tp.Fill(&delimiter)
		if fill_err != nil {
			return
		}
		var defaultValue []*big.Int
		fill_err = tp.Fill(&defaultValue)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.EnvOr_74318528(key, delimiter, defaultValue)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_EnvOr_859216bc__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var key string
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var delimiter string
		fill_err = tp.Fill(&delimiter)
		if fill_err != nil {
			return
		}
		var defaultValue []string
		fill_err = tp.Fill(&defaultValue)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.EnvOr_859216bc(key, delimiter, defaultValue)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_EnvOr_b3e47705__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var key string
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var defaultValue []byte
		fill_err = tp.Fill(&defaultValue)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.EnvOr_b3e47705(key, defaultValue)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_EnvOr_b4a85892__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var key string
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var defaultValue [32]byte
		fill_err = tp.Fill(&defaultValue)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.EnvOr_b4a85892(key, defaultValue)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_EnvOr_bbcb713e__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var key string
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var defaultValue *ABIInt256
		fill_err = tp.Fill(&defaultValue)
		if fill_err != nil {
			return
		}
	if c == nil|| defaultValue == nil {
		return
	}

	c.EnvOr_bbcb713e(key, defaultValue)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_EnvOr_c74e9deb__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var key string
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var delimiter string
		fill_err = tp.Fill(&delimiter)
		if fill_err != nil {
			return
		}
		var defaultValue []common.Address
		fill_err = tp.Fill(&defaultValue)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.EnvOr_c74e9deb(key, delimiter, defaultValue)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_EnvOr_d145736c__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var key string
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var defaultValue string
		fill_err = tp.Fill(&defaultValue)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.EnvOr_d145736c(key, defaultValue)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_EnvOr_eb85e83b__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var key string
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var delimiter string
		fill_err = tp.Fill(&delimiter)
		if fill_err != nil {
			return
		}
		var defaultValue []bool
		fill_err = tp.Fill(&defaultValue)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.EnvOr_eb85e83b(key, delimiter, defaultValue)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_EnvString_14b02bc9__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var key string
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var delimiter string
		fill_err = tp.Fill(&delimiter)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.EnvString_14b02bc9(key, delimiter)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_EnvString_f877cb19__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var key string
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.EnvString_f877cb19(key)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_EnvUint_c1978d1f__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var key string
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.EnvUint_c1978d1f(key)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_EnvUint_f3dec099__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var key string
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var delimiter string
		fill_err = tp.Fill(&delimiter)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.EnvUint_f3dec099(key, delimiter)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_Etch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var who common.Address
		fill_err = tp.Fill(&who)
		if fill_err != nil {
			return
		}
		var code []byte
		fill_err = tp.Fill(&code)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Etch(who, code)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_Fee__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var fee *big.Int
		fill_err = tp.Fill(&fee)
		if fill_err != nil {
			return
		}
	if c == nil|| fee == nil {
		return
	}

	c.Fee(fee)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_Ffi__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var args []string
		fill_err = tp.Fill(&args)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Ffi(args)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_GetBlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.GetBlockNumber()
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_GetBlockTimestamp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.GetBlockTimestamp()
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_GetCode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var input string
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.GetCode(input)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_GetDeployedCode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var input string
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.GetDeployedCode(input)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_GetLabel__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.GetLabel(addr)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_GetNonce__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.GetNonce(addr)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_IsPersistent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.IsPersistent(addr)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_KeyExists__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var jsonData string
		fill_err = tp.Fill(&jsonData)
		if fill_err != nil {
			return
		}
		var key string
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.KeyExists(jsonData, key)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_KeyExistsJson__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var d2 string
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}
		var key string
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.KeyExistsJson(d2, key)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_KeyExistsToml__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var d2 string
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}
		var key string
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.KeyExistsToml(d2, key)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_Label__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var label string
		fill_err = tp.Fill(&label)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Label(addr, label)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_Load__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var account common.Address
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		var slot [32]byte
		fill_err = tp.Fill(&slot)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Load(account, slot)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_LoadAllocs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var pathToAllocsJson string
		fill_err = tp.Fill(&pathToAllocsJson)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.LoadAllocs(pathToAllocsJson)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_MakePersistent_1d9e269e__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var accounts []common.Address
		fill_err = tp.Fill(&accounts)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.MakePersistent_1d9e269e(accounts)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_MakePersistent_4074e0a8__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var account0 common.Address
		fill_err = tp.Fill(&account0)
		if fill_err != nil {
			return
		}
		var account1 common.Address
		fill_err = tp.Fill(&account1)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.MakePersistent_4074e0a8(account0, account1)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_MakePersistent_57e22dde__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var account0 common.Address
		fill_err = tp.Fill(&account0)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.MakePersistent_57e22dde(account0)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_MakePersistent_efb77a75__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var account0 common.Address
		fill_err = tp.Fill(&account0)
		if fill_err != nil {
			return
		}
		var account1 common.Address
		fill_err = tp.Fill(&account1)
		if fill_err != nil {
			return
		}
		var account2 common.Address
		fill_err = tp.Fill(&account2)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.MakePersistent_efb77a75(account0, account1, account2)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_MockCallRevert_d23cd037__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var where common.Address
		fill_err = tp.Fill(&where)
		if fill_err != nil {
			return
		}
		var value *big.Int
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
		var d4 []byte
		fill_err = tp.Fill(&d4)
		if fill_err != nil {
			return
		}
		var retdata []byte
		fill_err = tp.Fill(&retdata)
		if fill_err != nil {
			return
		}
	if c == nil|| value == nil {
		return
	}

	c.MockCallRevert_d23cd037(where, value, d4, retdata)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_MockCallRevert_dbaad147__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var where common.Address
		fill_err = tp.Fill(&where)
		if fill_err != nil {
			return
		}
		var d3 []byte
		fill_err = tp.Fill(&d3)
		if fill_err != nil {
			return
		}
		var retdata []byte
		fill_err = tp.Fill(&retdata)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.MockCallRevert_dbaad147(where, d3, retdata)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_MockCall_81409b91__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var where common.Address
		fill_err = tp.Fill(&where)
		if fill_err != nil {
			return
		}
		var value *big.Int
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
		var d4 []byte
		fill_err = tp.Fill(&d4)
		if fill_err != nil {
			return
		}
		var retdata []byte
		fill_err = tp.Fill(&retdata)
		if fill_err != nil {
			return
		}
	if c == nil|| value == nil {
		return
	}

	c.MockCall_81409b91(where, value, d4, retdata)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_MockCall_b96213e4__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var where common.Address
		fill_err = tp.Fill(&where)
		if fill_err != nil {
			return
		}
		var d3 []byte
		fill_err = tp.Fill(&d3)
		if fill_err != nil {
			return
		}
		var retdata []byte
		fill_err = tp.Fill(&retdata)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.MockCall_b96213e4(where, d3, retdata)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_ParseAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var stringifiedValue string
		fill_err = tp.Fill(&stringifiedValue)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.ParseAddress(stringifiedValue)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_ParseBool__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var stringifiedValue string
		fill_err = tp.Fill(&stringifiedValue)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.ParseBool(stringifiedValue)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_ParseBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var stringifiedValue string
		fill_err = tp.Fill(&stringifiedValue)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.ParseBytes(stringifiedValue)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_ParseBytes32__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var stringifiedValue string
		fill_err = tp.Fill(&stringifiedValue)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.ParseBytes32(stringifiedValue)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_ParseInt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var stringifiedValue string
		fill_err = tp.Fill(&stringifiedValue)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.ParseInt(stringifiedValue)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_ParseJsonKeys__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var d2 string
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}
		var key string
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.ParseJsonKeys(d2, key)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_ParseJson_6a82600a__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var d2 string
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.ParseJson_6a82600a(d2)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_ParseJson_85940ef1__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var d2 string
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}
		var key string
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.ParseJson_85940ef1(d2, key)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_ParseTomlAddress_65e7c844__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var tomlStr string
		fill_err = tp.Fill(&tomlStr)
		if fill_err != nil {
			return
		}
		var key string
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.ParseTomlAddress_65e7c844(tomlStr, key)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_ParseTomlKeys__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var d2 string
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}
		var key string
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.ParseTomlKeys(d2, key)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_ParseToml_37736e08__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var d2 string
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}
		var key string
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.ParseToml_37736e08(d2, key)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_ParseToml_592151f0__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var d2 string
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.ParseToml_592151f0(d2)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_ParseUint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var stringifiedValue string
		fill_err = tp.Fill(&stringifiedValue)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.ParseUint(stringifiedValue)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_PauseGasMetering__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.PauseGasMetering()
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_Prank_47e50cce__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var sender common.Address
		fill_err = tp.Fill(&sender)
		if fill_err != nil {
			return
		}
		var origin common.Address
		fill_err = tp.Fill(&origin)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Prank_47e50cce(sender, origin)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_Prank_ca669fa7__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var sender common.Address
		fill_err = tp.Fill(&sender)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Prank_ca669fa7(sender)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_Prevrandao__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var v [32]byte
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Prevrandao(v)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_ProjectRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.ProjectRoot()
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_Prompt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Prompt()
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_ReadCallers__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.ReadCallers()
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_Record__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Record()
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_RecordLogs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.RecordLogs()
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_RememberKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var privateKey *big.Int
		fill_err = tp.Fill(&privateKey)
		if fill_err != nil {
			return
		}
	if c == nil|| privateKey == nil {
		return
	}

	c.RememberKey(privateKey)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_ResetNonce__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.ResetNonce(addr)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_ResumeGasMetering__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.ResumeGasMetering()
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_RevokePersistent_3ce969e6__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var addrs []common.Address
		fill_err = tp.Fill(&addrs)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.RevokePersistent_3ce969e6(addrs)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_RevokePersistent_997a0222__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.RevokePersistent_997a0222(addr)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_Roll__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var num *big.Int
		fill_err = tp.Fill(&num)
		if fill_err != nil {
			return
		}
	if c == nil|| num == nil {
		return
	}

	c.Roll(num)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_RollFork_0f29772b__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var txHash common.Hash
		fill_err = tp.Fill(&txHash)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.RollFork_0f29772b(txHash)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_RollFork_d74c83a4__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var forkID *big.Int
		fill_err = tp.Fill(&forkID)
		if fill_err != nil {
			return
		}
		var block *big.Int
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
	if c == nil|| forkID == nil|| block == nil {
		return
	}

	c.RollFork_d74c83a4(forkID, block)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_RollFork_d9bbf3a1__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var block *big.Int
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
	if c == nil|| block == nil {
		return
	}

	c.RollFork_d9bbf3a1(block)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_RollFork_f2830f7b__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var forkID *uint256.Int
		fill_err = tp.Fill(&forkID)
		if fill_err != nil {
			return
		}
		var txHash common.Hash
		fill_err = tp.Fill(&txHash)
		if fill_err != nil {
			return
		}
	if c == nil|| forkID == nil {
		return
	}

	c.RollFork_f2830f7b(forkID, txHash)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_SerializeAddress_1e356e1a__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var objectKey string
		fill_err = tp.Fill(&objectKey)
		if fill_err != nil {
			return
		}
		var valueKey string
		fill_err = tp.Fill(&valueKey)
		if fill_err != nil {
			return
		}
		var values []common.Address
		fill_err = tp.Fill(&values)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.SerializeAddress_1e356e1a(objectKey, valueKey, values)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_SerializeAddress_972c6062__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var objectKey string
		fill_err = tp.Fill(&objectKey)
		if fill_err != nil {
			return
		}
		var valueKey string
		fill_err = tp.Fill(&valueKey)
		if fill_err != nil {
			return
		}
		var value common.Address
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.SerializeAddress_972c6062(objectKey, valueKey, value)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_SerializeBool_92925aa1__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var objectKey string
		fill_err = tp.Fill(&objectKey)
		if fill_err != nil {
			return
		}
		var valueKey string
		fill_err = tp.Fill(&valueKey)
		if fill_err != nil {
			return
		}
		var values []bool
		fill_err = tp.Fill(&values)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.SerializeBool_92925aa1(objectKey, valueKey, values)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_SerializeBool_ac22e971__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var objectKey string
		fill_err = tp.Fill(&objectKey)
		if fill_err != nil {
			return
		}
		var valueKey string
		fill_err = tp.Fill(&valueKey)
		if fill_err != nil {
			return
		}
		var value bool
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.SerializeBool_ac22e971(objectKey, valueKey, value)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_SerializeBytes32_201e43e2__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var objectKey string
		fill_err = tp.Fill(&objectKey)
		if fill_err != nil {
			return
		}
		var valueKey string
		fill_err = tp.Fill(&valueKey)
		if fill_err != nil {
			return
		}
		var values []common.Hash
		fill_err = tp.Fill(&values)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.SerializeBytes32_201e43e2(objectKey, valueKey, values)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_SerializeBytes32_2d812b44__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var objectKey string
		fill_err = tp.Fill(&objectKey)
		if fill_err != nil {
			return
		}
		var valueKey string
		fill_err = tp.Fill(&valueKey)
		if fill_err != nil {
			return
		}
		var value common.Hash
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.SerializeBytes32_2d812b44(objectKey, valueKey, value)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_SerializeBytes_9884b232__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var objectKey string
		fill_err = tp.Fill(&objectKey)
		if fill_err != nil {
			return
		}
		var valueKey string
		fill_err = tp.Fill(&valueKey)
		if fill_err != nil {
			return
		}
		var values []hexutil.Bytes
		fill_err = tp.Fill(&values)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.SerializeBytes_9884b232(objectKey, valueKey, values)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_SerializeBytes_f21d52c7__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var objectKey string
		fill_err = tp.Fill(&objectKey)
		if fill_err != nil {
			return
		}
		var valueKey string
		fill_err = tp.Fill(&valueKey)
		if fill_err != nil {
			return
		}
		var value hexutil.Bytes
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.SerializeBytes_f21d52c7(objectKey, valueKey, value)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_SerializeInt_3f33db60__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var objectKey string
		fill_err = tp.Fill(&objectKey)
		if fill_err != nil {
			return
		}
		var valueKey string
		fill_err = tp.Fill(&valueKey)
		if fill_err != nil {
			return
		}
		var value *ABIInt256
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
	if c == nil|| value == nil {
		return
	}

	c.SerializeInt_3f33db60(objectKey, valueKey, value)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_SerializeInt_7676e127__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var objectKey string
		fill_err = tp.Fill(&objectKey)
		if fill_err != nil {
			return
		}
		var valueKey string
		fill_err = tp.Fill(&valueKey)
		if fill_err != nil {
			return
		}
		var values []*ABIInt256
		fill_err = tp.Fill(&values)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.SerializeInt_7676e127(objectKey, valueKey, values)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_SerializeJson__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var objectKey string
		fill_err = tp.Fill(&objectKey)
		if fill_err != nil {
			return
		}
		var value string
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.SerializeJson(objectKey, value)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_SerializeString_561cd6f3__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var objectKey string
		fill_err = tp.Fill(&objectKey)
		if fill_err != nil {
			return
		}
		var valueKey string
		fill_err = tp.Fill(&valueKey)
		if fill_err != nil {
			return
		}
		var values []string
		fill_err = tp.Fill(&values)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.SerializeString_561cd6f3(objectKey, valueKey, values)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_SerializeString_88da6d35__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var objectKey string
		fill_err = tp.Fill(&objectKey)
		if fill_err != nil {
			return
		}
		var valueKey string
		fill_err = tp.Fill(&valueKey)
		if fill_err != nil {
			return
		}
		var value string
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.SerializeString_88da6d35(objectKey, valueKey, value)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_SerializeUint_129e9002__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var objectKey string
		fill_err = tp.Fill(&objectKey)
		if fill_err != nil {
			return
		}
		var valueKey string
		fill_err = tp.Fill(&valueKey)
		if fill_err != nil {
			return
		}
		var value *big.Int
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
	if c == nil|| value == nil {
		return
	}

	c.SerializeUint_129e9002(objectKey, valueKey, value)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_SerializeUint_fee9a469__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var objectKey string
		fill_err = tp.Fill(&objectKey)
		if fill_err != nil {
			return
		}
		var valueKey string
		fill_err = tp.Fill(&valueKey)
		if fill_err != nil {
			return
		}
		var values []*big.Int
		fill_err = tp.Fill(&values)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.SerializeUint_fee9a469(objectKey, valueKey, values)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_SetEnv__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var key string
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var value string
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.SetEnv(key, value)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_SetNonce__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var account common.Address
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		var nonce uint64
		fill_err = tp.Fill(&nonce)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.SetNonce(account, nonce)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_Sign__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Sign()
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_Skip__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Skip()
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_Sleep__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ms *big.Int
		fill_err = tp.Fill(&ms)
		if fill_err != nil {
			return
		}
	if c == nil|| ms == nil {
		return
	}

	c.Sleep(ms)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_StartBroadcast_7fb5297f__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.StartBroadcast_7fb5297f()
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_StartBroadcast_7fec2a8d__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var who common.Address
		fill_err = tp.Fill(&who)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.StartBroadcast_7fec2a8d(who)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_StartPrank_06447d56__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var sender common.Address
		fill_err = tp.Fill(&sender)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.StartPrank_06447d56(sender)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_StartPrank_45b56078__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var sender common.Address
		fill_err = tp.Fill(&sender)
		if fill_err != nil {
			return
		}
		var origin common.Address
		fill_err = tp.Fill(&origin)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.StartPrank_45b56078(sender, origin)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_StartStateDiffRecording__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.StartStateDiffRecording()
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_StopAndReturnStateDiff__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.StopAndReturnStateDiff()
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_StopBroadcast__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.StopBroadcast()
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_StopPrank__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.StopPrank()
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_Store__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var account common.Address
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		var slot [32]byte
		fill_err = tp.Fill(&slot)
		if fill_err != nil {
			return
		}
		var value [32]byte
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Store(account, slot, value)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_ToString_56ca623e__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var v common.Address
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.ToString_56ca623e(v)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_ToString_6900a3ae__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var v *big.Int
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
	if c == nil|| v == nil {
		return
	}

	c.ToString_6900a3ae(v)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_ToString_71aad10d__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var v []byte
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.ToString_71aad10d(v)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_ToString_71dce7da__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var v bool
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.ToString_71dce7da(v)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_ToString_a322c40e__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var v *ABIInt256
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
	if c == nil|| v == nil {
		return
	}

	c.ToString_a322c40e(v)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_ToString_b11a19e8__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var v [32]byte
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.ToString_b11a19e8(v)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_TxGasPrice__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var newGasPrice *big.Int
		fill_err = tp.Fill(&newGasPrice)
		if fill_err != nil {
			return
		}
	if c == nil|| newGasPrice == nil {
		return
	}

	c.TxGasPrice(newGasPrice)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_UnixTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.UnixTime()
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_Warp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var timestamp *big.Int
		fill_err = tp.Fill(&timestamp)
		if fill_err != nil {
			return
		}
	if c == nil|| timestamp == nil {
		return
	}

	c.Warp(timestamp)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_WriteJson_35d6ad46__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var d2 string
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}
		var path string
		fill_err = tp.Fill(&path)
		if fill_err != nil {
			return
		}
		var valueKey string
		fill_err = tp.Fill(&valueKey)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.WriteJson_35d6ad46(d2, path, valueKey)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_WriteJson_e23cd19f__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var d2 string
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}
		var path string
		fill_err = tp.Fill(&path)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.WriteJson_e23cd19f(d2, path)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_WriteToml_51ac6a33__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var d2 string
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}
		var path string
		fill_err = tp.Fill(&path)
		if fill_err != nil {
			return
		}
		var valueKey string
		fill_err = tp.Fill(&valueKey)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.WriteToml_51ac6a33(d2, path, valueKey)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_WriteToml_c0865ba7__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var d2 string
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}
		var path string
		fill_err = tp.Fill(&path)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.WriteToml_c0865ba7(d2, path)
	})
}

// skipping Fuzz_Nosy_CheatCodesPrecompile_createFork__ because parameters include func, chan, or unsupported interface: []github.com/ethereum-optimism/optimism/op-chain-ops/script.ForkOption

// skipping Fuzz_Nosy_CheatCodesPrecompile_createSelectFork__ because parameters include func, chan, or unsupported interface: []github.com/ethereum-optimism/optimism/op-chain-ops/script.ForkOption

func Fuzz_Nosy_CheatCodesPrecompile_forkURLOption__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var id forking.ForkID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.forkURLOption(id)
	})
}

func Fuzz_Nosy_CheatCodesPrecompile_getArtifact__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *CheatCodesPrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var input string
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.getArtifact(input)
	})
}

// skipping Fuzz_Nosy_CheatCodesPrecompile_rollFork__ because parameters include func, chan, or unsupported interface: []github.com/ethereum-optimism/optimism/op-chain-ops/script.ForkOption

// skipping Fuzz_Nosy_CheatCodesPrecompile_serializeJsonValue__ because parameters include func, chan, or unsupported interface: any

func Fuzz_Nosy_ConsolePrecompile_Log_007150be__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_007150be(p0, p1, p2)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_00dd87b9__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 common.Address
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p1 == nil|| p2 == nil {
		return
	}

	c.Log_00dd87b9(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_013d178b__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 [10]byte
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_013d178b(p0)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_018c84c2__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_018c84c2(p0, p1, p2)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_031c6f73__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 string
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil {
		return
	}

	c.Log_031c6f73(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_04004a2e__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 [11]byte
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_04004a2e(p0)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_0454c079__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 string
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_0454c079(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_078287f5__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 *big.Int
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil|| p3 == nil {
		return
	}

	c.Log_078287f5(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_07831502__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 *big.Int
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p3 == nil {
		return
	}

	c.Log_07831502(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_088ef9d2__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
	if c == nil|| p1 == nil {
		return
	}

	c.Log_088ef9d2(p0, p1, p2)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_091ffaf5__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 bool
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil {
		return
	}

	c.Log_091ffaf5(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_0aa6cfad__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 string
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p2 == nil {
		return
	}

	c.Log_0aa6cfad(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_0b84bc58__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 [25]byte
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_0b84bc58(p0)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_0bb00eab__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 *big.Int
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p2 == nil|| p3 == nil {
		return
	}

	c.Log_0bb00eab(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_0be77f56__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 hexutil.Bytes
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_0be77f56(p0)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_0c66d1be__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 *big.Int
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p3 == nil {
		return
	}

	c.Log_0c66d1be(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_0c9cd9c1__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 *big.Int
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil|| p2 == nil|| p3 == nil {
		return
	}

	c.Log_0c9cd9c1(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_0d26b925__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
	if c == nil|| p2 == nil {
		return
	}

	c.Log_0d26b925(p0, p1, p2)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_0d36fa20__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 common.Address
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_0d36fa20(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_0df12b76__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 bool
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_0df12b76(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_0e378994__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 bool
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_0e378994(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_0ef7e050__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 bool
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p2 == nil {
		return
	}

	c.Log_0ef7e050(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_100f650e__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 *big.Int
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p1 == nil|| p3 == nil {
		return
	}

	c.Log_100f650e(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_1023f7b2__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 common.Address
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p2 == nil {
		return
	}

	c.Log_1023f7b2(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_1078f68d__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_1078f68d(p0, p1, p2)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_1093ee11__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
	if c == nil|| p2 == nil {
		return
	}

	c.Log_1093ee11(p0, p1, p2)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_12d6c788__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 string
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_12d6c788(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_12f21602__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
	if c == nil|| p2 == nil {
		return
	}

	c.Log_12f21602(p0, p1, p2)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_136b05dd__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 common.Address
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p2 == nil {
		return
	}

	c.Log_136b05dd(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_1537dc87__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 *big.Int
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p1 == nil|| p3 == nil {
		return
	}

	c.Log_1537dc87(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_1596a1ce__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 common.Address
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p2 == nil {
		return
	}

	c.Log_1596a1ce(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_159f8927__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 *big.Int
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p3 == nil {
		return
	}

	c.Log_159f8927(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_15c127b5__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 common.Address
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil|| p2 == nil {
		return
	}

	c.Log_15c127b5(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_15cac476__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 bool
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil|| p1 == nil {
		return
	}

	c.Log_15cac476(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_1606a393__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 *big.Int
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p3 == nil {
		return
	}

	c.Log_1606a393(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_1762e32a__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 string
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_1762e32a(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_17fe6185__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
	if c == nil|| p2 == nil {
		return
	}

	c.Log_17fe6185(p0, p1, p2)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_18c9c746__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_18c9c746(p0, p1, p2)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_193fb800__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 *big.Int
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil|| p1 == nil|| p2 == nil|| p3 == nil {
		return
	}

	c.Log_193fb800(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_19fd4956__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 common.Address
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_19fd4956(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_1ad96de6__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 string
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p2 == nil {
		return
	}

	c.Log_1ad96de6(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_1bb3b09a__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 string
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p1 == nil {
		return
	}

	c.Log_1bb3b09a(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_1c41a336__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 common.Address
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_1c41a336(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_1c7ec448__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
	if c == nil|| p1 == nil {
		return
	}

	c.Log_1c7ec448(p0, p1, p2)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_1c9d7eb3__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil {
		return
	}

	c.Log_1c9d7eb3(p0, p1)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_1d14d001__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 common.Address
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_1d14d001(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_1da986ea__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 string
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p1 == nil {
		return
	}

	c.Log_1da986ea(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_1dc8e1b8__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 *big.Int
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p2 == nil|| p3 == nil {
		return
	}

	c.Log_1dc8e1b8(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_1e4b87e5__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 bool
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_1e4b87e5(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_20098014__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil|| p2 == nil {
		return
	}

	c.Log_20098014(p0, p1, p2)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_205871c2__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 common.Address
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_205871c2(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_20718650__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil {
		return
	}

	c.Log_20718650(p0, p1, p2)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_20e3984d__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 common.Address
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p1 == nil|| p2 == nil {
		return
	}

	c.Log_20e3984d(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_212255cc__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_212255cc(p0, p1, p2)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_21ad0683__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 string
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil {
		return
	}

	c.Log_21ad0683(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_21bdaf25__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 string
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_21bdaf25(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_223603bd__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 common.Address
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_223603bd(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_22f6b999__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 *big.Int
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p1 == nil|| p3 == nil {
		return
	}

	c.Log_22f6b999(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_245986f2__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 string
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_245986f2(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_2488b414__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 common.Address
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil {
		return
	}

	c.Log_2488b414(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_24f91465__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 *big.Int
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p3 == nil {
		return
	}

	c.Log_24f91465(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_2555fa46__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_2555fa46(p0, p1, p2)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_26f560a8__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 common.Address
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p1 == nil {
		return
	}

	c.Log_26f560a8(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_27b7cf85__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 [32]byte
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_27b7cf85(p0)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_27d8afd2__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 string
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil|| p1 == nil {
		return
	}

	c.Log_27d8afd2(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_28863fcb__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 *big.Int
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p2 == nil|| p3 == nil {
		return
	}

	c.Log_28863fcb(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_2a110e83__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_2a110e83(p0, p1)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_2ae408d4__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 string
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_2ae408d4(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_2b2b18dc__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 common.Address
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_2b2b18dc(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_2c1754ed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 bool
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_2c1754ed(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_2c1d0746__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 *big.Int
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil|| p3 == nil {
		return
	}

	c.Log_2c1d0746(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_2c2ecbc2__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_2c2ecbc2(p0)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_2cd4134a__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 bool
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_2cd4134a(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_2ced7cef__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_2ced7cef(p0, p1, p2)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_2d5b6cb9__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *ABIInt256
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil {
		return
	}

	c.Log_2d5b6cb9(p0)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_2d834926__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 [3]byte
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_2d834926(p0)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_2d8e33a4__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 string
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_2d8e33a4(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_2dd778e6__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 string
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_2dd778e6(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_319af333__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_319af333(p0, p1)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_32458eed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_32458eed(p0)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_339f673a__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 [17]byte
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_339f673a(p0)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_33e9dd1d__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 common.Address
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_33e9dd1d(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_34f0e636__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 *big.Int
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p1 == nil|| p2 == nil|| p3 == nil {
		return
	}

	c.Log_34f0e636(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_35085f7b__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil {
		return
	}

	c.Log_35085f7b(p0, p1, p2)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_354c36d6__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 bool
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p1 == nil {
		return
	}

	c.Log_354c36d6(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_35a5071f__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 bool
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_35a5071f(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_37103367__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
	if c == nil|| p1 == nil|| p2 == nil {
		return
	}

	c.Log_37103367(p0, p1, p2)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_374bb4b2__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 *big.Int
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p1 == nil|| p2 == nil|| p3 == nil {
		return
	}

	c.Log_374bb4b2(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_37aa7d4c__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil|| p2 == nil {
		return
	}

	c.Log_37aa7d4c(p0, p1, p2)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_386ff5f4__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 *big.Int
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p2 == nil|| p3 == nil {
		return
	}

	c.Log_386ff5f4(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_3971e78c__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 *big.Int
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p3 == nil {
		return
	}

	c.Log_3971e78c(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_399174d3__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
	if c == nil|| p1 == nil {
		return
	}

	c.Log_399174d3(p0, p1)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_3a3757dd__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 [27]byte
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_3a3757dd(p0)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_3b2279b4__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 common.Address
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil|| p2 == nil {
		return
	}

	c.Log_3b2279b4(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_3b2a5ce0__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 bool
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_3b2a5ce0(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_3bf5e537__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 bool
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p1 == nil {
		return
	}

	c.Log_3bf5e537(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_3ca6268e__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *ABIInt256
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
	if c == nil|| p1 == nil {
		return
	}

	c.Log_3ca6268e(p0, p1)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_3e128ca3__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 string
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil {
		return
	}

	c.Log_3e128ca3(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_3e9f866a__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 *big.Int
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p3 == nil {
		return
	}

	c.Log_3e9f866a(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_3f8a701d__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 bool
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_3f8a701d(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_40785869__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 bool
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_40785869(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_41304fac__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_41304fac(p0)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_42d21db7__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 common.Address
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil|| p1 == nil {
		return
	}

	c.Log_42d21db7(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_439c7bef__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 common.Address
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_439c7bef(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_448830a8__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 string
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p2 == nil {
		return
	}

	c.Log_448830a8(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_454d54a5__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 bool
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil {
		return
	}

	c.Log_454d54a5(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_457fe3cf__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 *big.Int
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p3 == nil {
		return
	}

	c.Log_457fe3cf(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_46600be0__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 bool
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_46600be0(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_46826b5d__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 *big.Int
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil|| p3 == nil {
		return
	}

	c.Log_46826b5d(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_475c5c33__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 string
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_475c5c33(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_4766da72__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil|| p1 == nil {
		return
	}

	c.Log_4766da72(p0, p1, p2)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_478d1c62__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 common.Address
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p1 == nil {
		return
	}

	c.Log_478d1c62(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_483d0416__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 string
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_483d0416(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_4a28c017__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 string
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p1 == nil|| p2 == nil {
		return
	}

	c.Log_4a28c017(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_4a66cb34__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 string
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_4a66cb34(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_4b5c4277__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_4b5c4277(p0, p1)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_4b69c3d5__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 [29]byte
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_4b69c3d5(p0)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_4c123d57__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 *big.Int
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p3 == nil {
		return
	}

	c.Log_4c123d57(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_4ceda75a__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil {
		return
	}

	c.Log_4ceda75a(p0, p1, p2)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_4ed57e28__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 [7]byte
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_4ed57e28(p0)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_4f04fdc6__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 *big.Int
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p1 == nil|| p3 == nil {
		return
	}

	c.Log_4f04fdc6(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_4f84252e__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 [8]byte
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_4f84252e(p0)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_50709698__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_50709698(p0, p1, p2)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_50ad461d__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 bool
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_50ad461d(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_515e38b6__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 *big.Int
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p3 == nil {
		return
	}

	c.Log_515e38b6(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_5188e3e9__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 [20]byte
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_5188e3e9(p0)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_51973ec9__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_51973ec9()
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_51f09ff8__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 string
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p2 == nil {
		return
	}

	c.Log_51f09ff8(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_538e06ab__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 common.Address
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_538e06ab(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_54a7a9a0__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 common.Address
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p2 == nil {
		return
	}

	c.Log_54a7a9a0(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_56a5d1b1__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 common.Address
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil|| p1 == nil {
		return
	}

	c.Log_56a5d1b1(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_5821efa1__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
	if c == nil|| p2 == nil {
		return
	}

	c.Log_5821efa1(p0, p1, p2)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_5970e089__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
	if c == nil|| p1 == nil {
		return
	}

	c.Log_5970e089(p0, p1, p2)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_59cfcbe3__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 string
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil|| p1 == nil|| p2 == nil {
		return
	}

	c.Log_59cfcbe3(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_5a477632__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 string
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p2 == nil {
		return
	}

	c.Log_5a477632(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_5a9b5ed5__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil|| p2 == nil {
		return
	}

	c.Log_5a9b5ed5(p0, p1, p2)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_5ab84e1f__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 string
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p1 == nil {
		return
	}

	c.Log_5ab84e1f(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_5abd992a__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 *big.Int
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil|| p3 == nil {
		return
	}

	c.Log_5abd992a(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_5c430d47__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 common.Address
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p1 == nil {
		return
	}

	c.Log_5c430d47(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_5c96b331__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil|| p1 == nil {
		return
	}

	c.Log_5c96b331(p0, p1, p2)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_5ccd4e37__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 bool
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_5ccd4e37(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_5d02c50b__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 string
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_5d02c50b(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_5d08bb05__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 *big.Int
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p3 == nil {
		return
	}

	c.Log_5d08bb05(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_5d1a971a__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 string
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p2 == nil {
		return
	}

	c.Log_5d1a971a(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_5da297eb__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 *big.Int
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil|| p1 == nil|| p3 == nil {
		return
	}

	c.Log_5da297eb(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_5e6b5a33__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 [19]byte
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_5e6b5a33(p0)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_5e84b0ea__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 string
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_5e84b0ea(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_5ea2b7ae__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 common.Address
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p1 == nil {
		return
	}

	c.Log_5ea2b7ae(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_5f15d28c__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 bool
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_5f15d28c(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_5f1d5c9f__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 bool
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_5f1d5c9f(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_5f743a7c__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 bool
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil|| p2 == nil {
		return
	}

	c.Log_5f743a7c(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_5f7b9afb__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
	if c == nil|| p2 == nil {
		return
	}

	c.Log_5f7b9afb(p0, p1, p2)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_6168ed61__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 common.Address
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil {
		return
	}

	c.Log_6168ed61(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_619e4d0e__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 bool
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p2 == nil {
		return
	}

	c.Log_619e4d0e(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_63183678__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 common.Address
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p2 == nil {
		return
	}

	c.Log_63183678(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_63cb41f9__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil {
		return
	}

	c.Log_63cb41f9(p0, p1, p2)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_63fb8bc5__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 common.Address
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p2 == nil {
		return
	}

	c.Log_63fb8bc5(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_643fd0df__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil {
		return
	}

	c.Log_643fd0df(p0, p1)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_64b5bb67__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 *big.Int
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p2 == nil|| p3 == nil {
		return
	}

	c.Log_64b5bb67(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_660375dd__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 common.Address
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_660375dd(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_665bf134__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 common.Address
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_665bf134(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_665c6104__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 [16]byte
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_665c6104(p0)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_66f1bc67__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 bool
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p1 == nil|| p2 == nil {
		return
	}

	c.Log_66f1bc67(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_678209a8__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
	if c == nil|| p1 == nil {
		return
	}

	c.Log_678209a8(p0, p1, p2)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_67dd6ff1__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
	if c == nil|| p2 == nil {
		return
	}

	c.Log_67dd6ff1(p0, p1, p2)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_68c8b8bd__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 string
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil {
		return
	}

	c.Log_68c8b8bd(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_691a8f74__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 bool
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil|| p2 == nil {
		return
	}

	c.Log_691a8f74(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_69276c86__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil {
		return
	}

	c.Log_69276c86(p0, p1)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_69640b59__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 common.Address
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil {
		return
	}

	c.Log_69640b59(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_6a1199e2__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 *big.Int
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p1 == nil|| p3 == nil {
		return
	}

	c.Log_6a1199e2(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_6a9c478b__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 bool
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_6a9c478b(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_6b0e5d53__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 bool
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p2 == nil {
		return
	}

	c.Log_6b0e5d53(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_6cde40b8__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 string
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil|| p1 == nil {
		return
	}

	c.Log_6cde40b8(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_6d1e8751__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 string
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_6d1e8751(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_6d572f44__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 common.Address
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_6d572f44(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_6d7045c1__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 *big.Int
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p3 == nil {
		return
	}

	c.Log_6d7045c1(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_6dd434ca__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 bool
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_6dd434ca(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_6e18a128__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 [1]byte
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_6e18a128(p0)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_6f1a594e__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 bool
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_6f1a594e(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_6f7c603e__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 common.Address
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_6f7c603e(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_7190a529__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 common.Address
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_7190a529(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_71d04af2__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil|| p1 == nil {
		return
	}

	c.Log_71d04af2(p0, p1, p2)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_736efbb6__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 *big.Int
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil|| p3 == nil {
		return
	}

	c.Log_736efbb6(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_742d6ee7__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 string
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p2 == nil {
		return
	}

	c.Log_742d6ee7(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_7464ce23__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 *big.Int
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil|| p3 == nil {
		return
	}

	c.Log_7464ce23(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_759f86bb__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_759f86bb(p0, p1)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_75b605d3__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_75b605d3(p0, p1)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_7626db92__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 bool
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p1 == nil|| p2 == nil {
		return
	}

	c.Log_7626db92(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_79884c2b__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 bool
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_79884c2b(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_7af6ab25__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 bool
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil|| p1 == nil {
		return
	}

	c.Log_7af6ab25(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_7afac959__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil {
		return
	}

	c.Log_7afac959(p0, p1, p2)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_7bc0d848__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
	if c == nil|| p1 == nil {
		return
	}

	c.Log_7bc0d848(p0, p1, p2)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_7be0c3eb__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 *big.Int
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p3 == nil {
		return
	}

	c.Log_7be0c3eb(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_7bf181a1__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 *big.Int
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p2 == nil|| p3 == nil {
		return
	}

	c.Log_7bf181a1(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_7c4632a4__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 common.Address
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p1 == nil {
		return
	}

	c.Log_7c4632a4(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_7cc3c607__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 *big.Int
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p3 == nil {
		return
	}

	c.Log_7cc3c607(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_7d24491d__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 bool
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p1 == nil {
		return
	}

	c.Log_7d24491d(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_7dd4d0e0__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 string
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p2 == nil {
		return
	}

	c.Log_7dd4d0e0(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_7f9bbca2__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 *big.Int
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p1 == nil|| p3 == nil {
		return
	}

	c.Log_7f9bbca2(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_800a1c67__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 string
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_800a1c67(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_80e6a20b__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 *big.Int
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p3 == nil {
		return
	}

	c.Log_80e6a20b(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_82112a42__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 bool
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p1 == nil {
		return
	}

	c.Log_82112a42(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_82c25b74__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 *big.Int
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil|| p2 == nil|| p3 == nil {
		return
	}

	c.Log_82c25b74(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_8309e8a8__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
	if c == nil|| p1 == nil {
		return
	}

	c.Log_8309e8a8(p0, p1)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_850b7ad6__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_850b7ad6(p0, p1, p2)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_853c4849__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_853c4849(p0, p1)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_854b3496__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 string
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p1 == nil|| p2 == nil {
		return
	}

	c.Log_854b3496(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_85775021__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil {
		return
	}

	c.Log_85775021(p0, p1, p2)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_86a06abd__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 [12]byte
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_86a06abd(p0)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_88a8c406__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 string
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p1 == nil {
		return
	}

	c.Log_88a8c406(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_88cb6041__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 common.Address
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil|| p2 == nil {
		return
	}

	c.Log_88cb6041(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_88f6e4b2__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 *big.Int
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil|| p1 == nil|| p3 == nil {
		return
	}

	c.Log_88f6e4b2(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_895af8c5__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 bool
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_895af8c5(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_8af7cf8a__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 bool
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p2 == nil {
		return
	}

	c.Log_8af7cf8a(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_8c329b1a__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 common.Address
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_8c329b1a(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_8c4e5de6__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 *big.Int
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p3 == nil {
		return
	}

	c.Log_8c4e5de6(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_8da6def5__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 common.Address
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p2 == nil {
		return
	}

	c.Log_8da6def5(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_8e3f78a9__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 *big.Int
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p3 == nil {
		return
	}

	c.Log_8e3f78a9(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_8e69fb5d__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 string
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p1 == nil|| p2 == nil {
		return
	}

	c.Log_8e69fb5d(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_8eafb02b__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 *big.Int
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p3 == nil {
		return
	}

	c.Log_8eafb02b(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_8ef3f399__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 *big.Int
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p3 == nil {
		return
	}

	c.Log_8ef3f399(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_8f736d16__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 common.Address
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_8f736d16(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_8feac525__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_8feac525(p0, p1)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_90bd8cd0__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 [9]byte
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_90bd8cd0(p0)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_90c30a56__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 bool
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil {
		return
	}

	c.Log_90c30a56(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_90fb06aa__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 string
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil {
		return
	}

	c.Log_90fb06aa(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_9143dbb1__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 string
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p1 == nil {
		return
	}

	c.Log_9143dbb1(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_91a02e2a__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 bool
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil|| p2 == nil {
		return
	}

	c.Log_91a02e2a(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_91d1112e__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 *big.Int
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p3 == nil {
		return
	}

	c.Log_91d1112e(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_9266f07f__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 [14]byte
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_9266f07f(p0)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_932bbb38__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_932bbb38(p0, p1, p2)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_935e09bf__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 common.Address
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p2 == nil {
		return
	}

	c.Log_935e09bf(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_94250d77__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 *big.Int
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p3 == nil {
		return
	}

	c.Log_94250d77(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_94529e34__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 [13]byte
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_94529e34(p0)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_958c28c6__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 bool
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_958c28c6(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_9591b953__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_9591b953(p0, p1, p2)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_95ed0195__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_95ed0195(p0, p1, p2)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_97d394d8__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 common.Address
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_97d394d8(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_9a816a83__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 common.Address
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil|| p1 == nil {
		return
	}

	c.Log_9a816a83(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_9acd3616__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 common.Address
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p1 == nil {
		return
	}

	c.Log_9acd3616(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_9b4254e2__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 bool
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p2 == nil {
		return
	}

	c.Log_9b4254e2(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_9b6ec042__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil {
		return
	}

	c.Log_9b6ec042(p0, p1, p2)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_9c3adfa1__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 string
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil {
		return
	}

	c.Log_9c3adfa1(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_9c4f99fb__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
	if c == nil|| p2 == nil {
		return
	}

	c.Log_9c4f99fb(p0, p1, p2)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_9cba8fff__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 common.Address
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil {
		return
	}

	c.Log_9cba8fff(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_9d22d5dd__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 string
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_9d22d5dd(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_9f1bc36e__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 common.Address
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_9f1bc36e(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_9ffb2f93__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 string
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p1 == nil {
		return
	}

	c.Log_9ffb2f93(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_a04e2f87__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 common.Address
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_a04e2f87(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_a0a47963__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 string
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_a0a47963(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_a1bcc9b3__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 bool
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p1 == nil {
		return
	}

	c.Log_a1bcc9b3(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_a1ef4cbb__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 common.Address
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil {
		return
	}

	c.Log_a1ef4cbb(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_a1f2e8aa__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
	if c == nil|| p1 == nil {
		return
	}

	c.Log_a1f2e8aa(p0, p1, p2)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_a31bfdcc__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 common.Address
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p1 == nil {
		return
	}

	c.Log_a31bfdcc(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_a5b4fc99__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 string
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil|| p1 == nil {
		return
	}

	c.Log_a5b4fc99(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_a5cada94__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 *big.Int
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p3 == nil {
		return
	}

	c.Log_a5cada94(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_a684808d__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 [5]byte
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_a684808d(p0)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_a6f50b0f__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 bool
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_a6f50b0f(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_a73c1db6__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 string
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_a73c1db6(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_a75c59de__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 *big.Int
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p3 == nil {
		return
	}

	c.Log_a75c59de(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_a7a87853__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 *big.Int
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p1 == nil|| p2 == nil|| p3 == nil {
		return
	}

	c.Log_a7a87853(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_a826caeb__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 string
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_a826caeb(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_aa6540c8__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 string
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_aa6540c8(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_aabc9a31__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 common.Address
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_aabc9a31(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_ab085ae6__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 bool
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil|| p1 == nil {
		return
	}

	c.Log_ab085ae6(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_aba1cf0d__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 [23]byte
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_aba1cf0d(p0)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_abf73a98__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 string
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p1 == nil {
		return
	}

	c.Log_abf73a98(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_ade052c7__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 string
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil {
		return
	}

	c.Log_ade052c7(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_ae2ec581__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 common.Address
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil {
		return
	}

	c.Log_ae2ec581(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_ae84a591__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 [6]byte
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_ae84a591(p0)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_b028c9bd__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 *big.Int
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil|| p3 == nil {
		return
	}

	c.Log_b028c9bd(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_b076847f__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_b076847f(p0, p1, p2)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_b0e0f9b5__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_b0e0f9b5(p0, p1, p2)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_b115611f__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil {
		return
	}

	c.Log_b115611f(p0, p1, p2)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_b3a6b6bd__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 bool
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil {
		return
	}

	c.Log_b3a6b6bd(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_b4c314ff__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 bool
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p1 == nil {
		return
	}

	c.Log_b4c314ff(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_b59dbd60__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 bool
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_b59dbd60(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_b60e72cc__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
	if c == nil|| p1 == nil {
		return
	}

	c.Log_b60e72cc(p0, p1)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_b69bcaf6__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
	if c == nil|| p1 == nil|| p2 == nil {
		return
	}

	c.Log_b69bcaf6(p0, p1, p2)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_b6f577a1__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 bool
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil {
		return
	}

	c.Log_b6f577a1(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_b7b914ca__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 string
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil|| p2 == nil {
		return
	}

	c.Log_b7b914ca(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_b857163a__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 bool
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_b857163a(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_ba535d9c__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 bool
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil {
		return
	}

	c.Log_ba535d9c(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_bc0b61fe__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 string
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_bc0b61fe(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_bcfd9be0__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil {
		return
	}

	c.Log_bcfd9be0(p0, p1, p2)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_be553481__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 *big.Int
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p2 == nil|| p3 == nil {
		return
	}

	c.Log_be553481(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_be984353__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 bool
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p1 == nil|| p2 == nil {
		return
	}

	c.Log_be984353(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_bf01f891__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 *big.Int
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p1 == nil|| p3 == nil {
		return
	}

	c.Log_bf01f891(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_c0a302d8__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 bool
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_c0a302d8(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_c21f64c7__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 *big.Int
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p3 == nil {
		return
	}

	c.Log_c21f64c7(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_c2854d92__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 [31]byte
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_c2854d92(p0)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_c371c7db__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 common.Address
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_c371c7db(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_c3a8a654__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 bool
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p2 == nil {
		return
	}

	c.Log_c3a8a654(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_c3b55635__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_c3b55635(p0, p1)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_c3fc3970__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
	if c == nil|| p1 == nil {
		return
	}

	c.Log_c3fc3970(p0, p1, p2)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_c4643e20__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 bool
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p2 == nil {
		return
	}

	c.Log_c4643e20(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_c4d23d9a__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 [18]byte
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_c4d23d9a(p0)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_c598d185__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 bool
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil|| p1 == nil|| p2 == nil {
		return
	}

	c.Log_c598d185(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_c5ad85f9__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 string
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p1 == nil {
		return
	}

	c.Log_c5ad85f9(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_c67ea9d1__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 *big.Int
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p1 == nil|| p3 == nil {
		return
	}

	c.Log_c67ea9d1(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_c6acc7a8__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 *big.Int
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil|| p2 == nil|| p3 == nil {
		return
	}

	c.Log_c6acc7a8(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_c82aeaee__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 [28]byte
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_c82aeaee(p0)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_c91d5ed4__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_c91d5ed4(p0, p1, p2)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_c95958d6__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
	if c == nil|| p2 == nil {
		return
	}

	c.Log_c95958d6(p0, p1, p2)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_ca47c4eb__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
	if c == nil|| p1 == nil|| p2 == nil {
		return
	}

	c.Log_ca47c4eb(p0, p1, p2)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_ca7733b1__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
	if c == nil|| p1 == nil {
		return
	}

	c.Log_ca7733b1(p0, p1, p2)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_cac43479__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 bool
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_cac43479(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_cc32ab07__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 bool
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil {
		return
	}

	c.Log_cc32ab07(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_ccf790a1__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 common.Address
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p2 == nil {
		return
	}

	c.Log_ccf790a1(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_ceb5f4d7__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 bool
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p1 == nil {
		return
	}

	c.Log_ceb5f4d7(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_cf009880__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 *big.Int
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil|| p3 == nil {
		return
	}

	c.Log_cf009880(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_cf020fb1__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_cf020fb1(p0, p1, p2)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_cf18105c__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 bool
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p1 == nil {
		return
	}

	c.Log_cf18105c(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_cf394485__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 common.Address
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_cf394485(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_d1ed7a3c__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil|| p1 == nil|| p2 == nil {
		return
	}

	c.Log_d1ed7a3c(p0, p1, p2)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_d2763667__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_d2763667(p0, p1, p2)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_d2d423cd__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 string
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil {
		return
	}

	c.Log_d2d423cd(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_d583c602__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 common.Address
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil {
		return
	}

	c.Log_d583c602(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_d5fae89c__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 [22]byte
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_d5fae89c(p0)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_d6019f1c__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 bool
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p2 == nil {
		return
	}

	c.Log_d6019f1c(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_d6aefad2__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 *big.Int
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p3 == nil {
		return
	}

	c.Log_d6aefad2(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_d812a167__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 string
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_d812a167(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_da9574e0__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 [15]byte
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_da9574e0(p0)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_daf0d4aa__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_daf0d4aa(p0, p1)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_dbb4c247__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_dbb4c247(p0, p1, p2)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_dc5e935b__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 bool
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_dc5e935b(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_ddb06521__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 string
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil|| p2 == nil {
		return
	}

	c.Log_ddb06521(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_dddb9561__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 string
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil {
		return
	}

	c.Log_dddb9561(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_de03e774__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 string
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil|| p2 == nil {
		return
	}

	c.Log_de03e774(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_de68f20a__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 string
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_de68f20a(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_de9a9270__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_de9a9270(p0, p1, p2)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_dfc4a2e8__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 string
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_dfc4a2e8(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_e05f48d1__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 [4]byte
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_e05f48d1(p0)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_e0625b29__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 common.Address
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_e0625b29(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_e0e95b98__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 common.Address
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p1 == nil {
		return
	}

	c.Log_e0e95b98(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_e0e9ad4f__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_e0e9ad4f(p0, p1, p2)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_e21de278__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 common.Address
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p1 == nil|| p2 == nil {
		return
	}

	c.Log_e21de278(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_e298f47d__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_e298f47d(p0, p1, p2)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_e2bfd60b__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 bool
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_e2bfd60b(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_e351140f__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 bool
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil {
		return
	}

	c.Log_e351140f(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_e3a9ca2f__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 *big.Int
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p3 == nil {
		return
	}

	c.Log_e3a9ca2f(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_e41b6f6f__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 *big.Int
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p1 == nil|| p3 == nil {
		return
	}

	c.Log_e41b6f6f(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_e5e70b2b__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 bool
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p1 == nil {
		return
	}

	c.Log_e5e70b2b(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_e8d3018d__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 *big.Int
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil|| p3 == nil {
		return
	}

	c.Log_e8d3018d(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_e8defba9__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
	if c == nil|| p1 == nil {
		return
	}

	c.Log_e8defba9(p0, p1, p2)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_e9b62296__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 [2]byte
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_e9b62296(p0)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_e9da3560__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 [21]byte
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_e9da3560(p0)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_eb1bff80__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 string
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_eb1bff80(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_eb7f6fd2__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 *big.Int
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil|| p1 == nil|| p3 == nil {
		return
	}

	c.Log_eb7f6fd2(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_eb830c92__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_eb830c92(p0, p1, p2)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_eb928d7f__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 bool
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil {
		return
	}

	c.Log_eb928d7f(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_ed8f28f6__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 common.Address
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_ed8f28f6(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_ee12c4ed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 [30]byte
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_ee12c4ed(p0)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_ef1cefe7__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 *big.Int
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p3 == nil {
		return
	}

	c.Log_ef1cefe7(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_ef529018__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 common.Address
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil {
		return
	}

	c.Log_ef529018(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_ef72c513__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 common.Address
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil {
		return
	}

	c.Log_ef72c513(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_f08744e8__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_f08744e8(p0, p1, p2)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_f11699ed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_f11699ed(p0, p1, p2)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_f1b35b34__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 [24]byte
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_f1b35b34(p0)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_f2a66286__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 bool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_f2a66286(p0, p1, p2)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_f45d7d2c__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 *big.Int
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p2 == nil|| p3 == nil {
		return
	}

	c.Log_f45d7d2c(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_f4880ea4__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 common.Address
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_f4880ea4(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_f5bc2249__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 string
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p1 == nil {
		return
	}

	c.Log_f5bc2249(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_f666715a__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil|| p1 == nil {
		return
	}

	c.Log_f666715a(p0, p1)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_f7e36245__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 string
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_f7e36245(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_f808da20__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 string
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_f808da20(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_f82c50f1__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil {
		return
	}

	c.Log_f82c50f1(p0)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_f8b149f1__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 [26]byte
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_f8b149f1(p0)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_f8f51b1e__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 *big.Int
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p2 == nil|| p3 == nil {
		return
	}

	c.Log_f8f51b1e(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_f9ad2b89__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 bool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 common.Address
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_f9ad2b89(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_fa8185af__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 *big.Int
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 common.Address
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p0 == nil|| p1 == nil|| p2 == nil {
		return
	}

	c.Log_fa8185af(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_fb772265__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 string
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_fb772265(p0, p1, p2)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_fc4845f0__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 bool
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p2 == nil {
		return
	}

	c.Log_fc4845f0(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_fcec75e0__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 string
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 common.Address
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
	if c == nil {
		return
	}

	c.Log_fcec75e0(p0, p1, p2)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_fdb4f990__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 common.Address
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 common.Address
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *big.Int
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 string
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p2 == nil {
		return
	}

	c.Log_fdb4f990(p0, p1, p2, p3)
	})
}

func Fuzz_Nosy_ConsolePrecompile_Log_fedd1fff__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var c *ConsolePrecompile
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p0 bool
		fill_err = tp.Fill(&p0)
		if fill_err != nil {
			return
		}
		var p1 *big.Int
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 string
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 common.Address
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
	if c == nil|| p1 == nil {
		return
	}

	c.Log_fedd1fff(p0, p1, p2, p3)
	})
}

// skipping Fuzz_Nosy_ConsolePrecompile_log__ because parameters include func, chan, or unsupported interface: []any

func Fuzz_Nosy_Host_AllowCheatcodes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var h *Host
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
	if h == nil {
		return
	}

	h.AllowCheatcodes(addr)
	})
}

func Fuzz_Nosy_Host_AllowedCheatcodes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var h *Host
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
	if h == nil {
		return
	}

	h.AllowedCheatcodes(addr)
	})
}

func Fuzz_Nosy_Host_Artifacts__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var h *Host
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
	if h == nil {
		return
	}

	h.Artifacts()
	})
}

func Fuzz_Nosy_Host_Call__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var h *Host
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var from common.Address
		fill_err = tp.Fill(&from)
		if fill_err != nil {
			return
		}
		var to common.Address
		fill_err = tp.Fill(&to)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		var gas uint64
		fill_err = tp.Fill(&gas)
		if fill_err != nil {
			return
		}
		var value *uint256.Int
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
	if h == nil|| value == nil {
		return
	}

	h.Call(from, to, input, gas, value)
	})
}

func Fuzz_Nosy_Host_CallerMode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var h *Host
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
	if h == nil {
		return
	}

	h.CallerMode()
	})
}

func Fuzz_Nosy_Host_ChainID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var h *Host
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
	if h == nil {
		return
	}

	h.ChainID()
	})
}

func Fuzz_Nosy_Host_Create__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var h *Host
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var from common.Address
		fill_err = tp.Fill(&from)
		if fill_err != nil {
			return
		}
		var initCode []byte
		fill_err = tp.Fill(&initCode)
		if fill_err != nil {
			return
		}
	if h == nil {
		return
	}

	h.Create(from, initCode)
	})
}

// skipping Fuzz_Nosy_Host_CreateSelectFork__ because parameters include func, chan, or unsupported interface: []github.com/ethereum-optimism/optimism/op-chain-ops/script.ForkOption

func Fuzz_Nosy_Host_CurrentCall__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var h *Host
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
	if h == nil {
		return
	}

	h.CurrentCall()
	})
}

func Fuzz_Nosy_Host_EnableCheats__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var h *Host
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
	if h == nil {
		return
	}

	h.EnableCheats()
	})
}

func Fuzz_Nosy_Host_EnforceMaxCodeSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var h *Host
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var v bool
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
	if h == nil {
		return
	}

	h.EnforceMaxCodeSize(v)
	})
}

func Fuzz_Nosy_Host_GetCode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var h *Host
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
	if h == nil {
		return
	}

	h.GetCode(addr)
	})
}

func Fuzz_Nosy_Host_GetEnvVar__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var h *Host
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var key string
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
	if h == nil {
		return
	}

	h.GetEnvVar(key)
	})
}

func Fuzz_Nosy_Host_GetNonce__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var h *Host
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
	if h == nil {
		return
	}

	h.GetNonce(addr)
	})
}

func Fuzz_Nosy_Host_HasPrecompileOverride__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var h *Host
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
	if h == nil {
		return
	}

	h.HasPrecompileOverride(addr)
	})
}

func Fuzz_Nosy_Host_ImportAccount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var h *Host
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var account types.Account
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
	if h == nil {
		return
	}

	h.ImportAccount(addr, account)
	})
}

func Fuzz_Nosy_Host_ImportState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var h *Host
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var allocs *foundry.ForgeAllocs
		fill_err = tp.Fill(&allocs)
		if fill_err != nil {
			return
		}
	if h == nil|| allocs == nil {
		return
	}

	h.ImportState(allocs)
	})
}

func Fuzz_Nosy_Host_Label__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var h *Host
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var label string
		fill_err = tp.Fill(&label)
		if fill_err != nil {
			return
		}
	if h == nil {
		return
	}

	h.Label(addr, label)
	})
}

func Fuzz_Nosy_Host_LoadContract__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var h *Host
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var artifactName string
		fill_err = tp.Fill(&artifactName)
		if fill_err != nil {
			return
		}
		var contractName string
		fill_err = tp.Fill(&contractName)
		if fill_err != nil {
			return
		}
	if h == nil {
		return
	}

	h.LoadContract(artifactName, contractName)
	})
}

func Fuzz_Nosy_Host_LogCallStack__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var h *Host
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
	if h == nil {
		return
	}

	h.LogCallStack()
	})
}

func Fuzz_Nosy_Host_MsgSender__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var h *Host
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
	if h == nil {
		return
	}

	h.MsgSender()
	})
}

func Fuzz_Nosy_Host_NewScriptAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var h *Host
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
	if h == nil {
		return
	}

	h.NewScriptAddress()
	})
}

func Fuzz_Nosy_Host_Prank__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var h *Host
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var msgSender *common.Address
		fill_err = tp.Fill(&msgSender)
		if fill_err != nil {
			return
		}
		var txOrigin *common.Address
		fill_err = tp.Fill(&txOrigin)
		if fill_err != nil {
			return
		}
		var repeat bool
		fill_err = tp.Fill(&repeat)
		if fill_err != nil {
			return
		}
		var broadcast bool
		fill_err = tp.Fill(&broadcast)
		if fill_err != nil {
			return
		}
	if h == nil|| msgSender == nil|| txOrigin == nil {
		return
	}

	h.Prank(msgSender, txOrigin, repeat, broadcast)
	})
}

func Fuzz_Nosy_Host_RememberArtifact__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var h *Host
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var artifact *foundry.Artifact
		fill_err = tp.Fill(&artifact)
		if fill_err != nil {
			return
		}
		var contract string
		fill_err = tp.Fill(&contract)
		if fill_err != nil {
			return
		}
	if h == nil|| artifact == nil {
		return
	}

	h.RememberArtifact(addr, artifact, contract)
	})
}

func Fuzz_Nosy_Host_RememberOnLabel__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var h *Host
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var label string
		fill_err = tp.Fill(&label)
		if fill_err != nil {
			return
		}
		var srcFile string
		fill_err = tp.Fill(&srcFile)
		if fill_err != nil {
			return
		}
		var contract string
		fill_err = tp.Fill(&contract)
		if fill_err != nil {
			return
		}
	if h == nil {
		return
	}

	h.RememberOnLabel(label, srcFile, contract)
	})
}

func Fuzz_Nosy_Host_ScriptBackendFn__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var h *Host
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var to common.Address
		fill_err = tp.Fill(&to)
		if fill_err != nil {
			return
		}
	if h == nil {
		return
	}

	h.ScriptBackendFn(to)
	})
}

func Fuzz_Nosy_Host_SelfAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var h *Host
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
	if h == nil {
		return
	}

	h.SelfAddress()
	})
}

func Fuzz_Nosy_Host_SetEnvVar__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var h *Host
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var key string
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var value string
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
	if h == nil {
		return
	}

	h.SetEnvVar(key, value)
	})
}

func Fuzz_Nosy_Host_SetNonce__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var h *Host
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var nonce uint64
		fill_err = tp.Fill(&nonce)
		if fill_err != nil {
			return
		}
	if h == nil {
		return
	}

	h.SetNonce(addr, nonce)
	})
}

// skipping Fuzz_Nosy_Host_SetPrecompile__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/vm.PrecompiledContract

func Fuzz_Nosy_Host_SetTxOrigin__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var h *Host
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
	if h == nil {
		return
	}

	h.SetTxOrigin(addr)
	})
}

func Fuzz_Nosy_Host_StateDump__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var h *Host
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
	if h == nil {
		return
	}

	h.StateDump()
	})
}

func Fuzz_Nosy_Host_StopPrank__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var h *Host
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var broadcast bool
		fill_err = tp.Fill(&broadcast)
		if fill_err != nil {
			return
		}
	if h == nil {
		return
	}

	h.StopPrank(broadcast)
	})
}

func Fuzz_Nosy_Host_TxOrigin__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var h *Host
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
	if h == nil {
		return
	}

	h.TxOrigin()
	})
}

func Fuzz_Nosy_Host_Wipe__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var h *Host
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
	if h == nil {
		return
	}

	h.Wipe(addr)
	})
}

// skipping Fuzz_Nosy_Host_getPrecompile__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/vm.PrecompiledContract

// skipping Fuzz_Nosy_Host_handleCaller__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/vm.ContractRef

func Fuzz_Nosy_Host_onEnter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var h *Host
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var depth int
		fill_err = tp.Fill(&depth)
		if fill_err != nil {
			return
		}
		var typ byte
		fill_err = tp.Fill(&typ)
		if fill_err != nil {
			return
		}
		var from common.Address
		fill_err = tp.Fill(&from)
		if fill_err != nil {
			return
		}
		var to common.Address
		fill_err = tp.Fill(&to)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		var gas uint64
		fill_err = tp.Fill(&gas)
		if fill_err != nil {
			return
		}
		var value *big.Int
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
	if h == nil|| value == nil {
		return
	}

	h.onEnter(depth, typ, from, to, input, gas, value)
	})
}

// skipping Fuzz_Nosy_Host_onExit__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_Host_onFault__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/tracing.OpContext

// skipping Fuzz_Nosy_Host_onFork__ because parameters include func, chan, or unsupported interface: []github.com/ethereum-optimism/optimism/op-chain-ops/script.ForkOption

func Fuzz_Nosy_Host_onLog__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var h *Host
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var ev *types.Log
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}
	if h == nil|| ev == nil {
		return
	}

	h.onLog(ev)
	})
}

// skipping Fuzz_Nosy_Host_onOpcode__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/tracing.OpContext

func Fuzz_Nosy_Host_onStorageChange__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var h *Host
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var slot common.Hash
		fill_err = tp.Fill(&slot)
		if fill_err != nil {
			return
		}
		var prev common.Hash
		fill_err = tp.Fill(&prev)
		if fill_err != nil {
			return
		}
		var new common.Hash
		fill_err = tp.Fill(&new)
		if fill_err != nil {
			return
		}
	if h == nil {
		return
	}

	h.onStorageChange(addr, slot, prev, new)
	})
}

func Fuzz_Nosy_Host_prelude__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var h *Host
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var from common.Address
		fill_err = tp.Fill(&from)
		if fill_err != nil {
			return
		}
		var to *common.Address
		fill_err = tp.Fill(&to)
		if fill_err != nil {
			return
		}
	if h == nil|| to == nil {
		return
	}

	h.prelude(from, to)
	})
}

func Fuzz_Nosy_Host_unwindCallstack__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var h *Host
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var depth int
		fill_err = tp.Fill(&depth)
		if fill_err != nil {
			return
		}
	if h == nil {
		return
	}

	h.unwindCallstack(depth)
	})
}

func Fuzz_Nosy_Precompile[E any]_RequiredGas__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var p *Precompile[E]
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
	if p == nil {
		return
	}

	p.RequiredGas(input)
	})
}

func Fuzz_Nosy_Precompile[E any]_Run__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var p *Precompile[E]
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
	if p == nil {
		return
	}

	p.Run(input)
	})
}

func Fuzz_Nosy_Precompile[E any]_setupFieldSetter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var p *Precompile[E]
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
	if p == nil {
		return
	}

	p.setupFieldSetter()
	})
}

func Fuzz_Nosy_Precompile[E any]_setupFields__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var p *Precompile[E]
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var val *reflect.Value
		fill_err = tp.Fill(&val)
		if fill_err != nil {
			return
		}
	if p == nil|| val == nil {
		return
	}

	p.setupFields(val)
	})
}

func Fuzz_Nosy_Precompile[E any]_setupMethod__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var p *Precompile[E]
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var selfVal *reflect.Value
		fill_err = tp.Fill(&selfVal)
		if fill_err != nil {
			return
		}
		var methodDef *reflect.Method
		fill_err = tp.Fill(&methodDef)
		if fill_err != nil {
			return
		}
	if p == nil|| selfVal == nil|| methodDef == nil {
		return
	}

	p.setupMethod(selfVal, methodDef)
	})
}

func Fuzz_Nosy_Precompile[E any]_setupMethods__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var p *Precompile[E]
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var val *reflect.Value
		fill_err = tp.Fill(&val)
		if fill_err != nil {
			return
		}
	if p == nil|| val == nil {
		return
	}

	p.setupMethods(val)
	})
}

func Fuzz_Nosy_Precompile[E any]_setupStructField__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var p *Precompile[E]
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var fieldDef *reflect.StructField
		fill_err = tp.Fill(&fieldDef)
		if fill_err != nil {
			return
		}
		var fieldVal *reflect.Value
		fill_err = tp.Fill(&fieldVal)
		if fill_err != nil {
			return
		}
	if p == nil|| fieldDef == nil|| fieldVal == nil {
		return
	}

	p.setupStructField(fieldDef, fieldVal)
	})
}

func Fuzz_Nosy_prankRef_Address__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var p *prankRef
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
	if p == nil {
		return
	}

	p.Address()
	})
}

func Fuzz_Nosy_prankRef_Value__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var p *prankRef
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
	if p == nil {
		return
	}

	p.Value()
	})
}

func Fuzz_Nosy_Broadcast_ID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var parent *CallFrame
		fill_err = tp.Fill(&parent)
		if fill_err != nil {
			return
		}
		var current *CallFrame
		fill_err = tp.Fill(&current)
		if fill_err != nil {
			return
		}
	if parent == nil|| current == nil {
		return
	}

	b := NewBroadcast(parent, current)
	b.ID()
	})
}

func Fuzz_Nosy_BroadcastType_MarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var bt BroadcastType
		fill_err = tp.Fill(&bt)
		if fill_err != nil {
			return
		}

	bt.MarshalText()
	})
}

func Fuzz_Nosy_BroadcastType_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var bt BroadcastType
		fill_err = tp.Fill(&bt)
		if fill_err != nil {
			return
		}

	bt.String()
	})
}

func Fuzz_Nosy_CallerMode_Big__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var cm CallerMode
		fill_err = tp.Fill(&cm)
		if fill_err != nil {
			return
		}

	cm.Big()
	})
}

// skipping Fuzz_Nosy_MakeBindings__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-chain-ops/script.CallBackendFn

func Fuzz_Nosy_SplitJSONPathKeys__(f *testing.F) {
	f.Fuzz(func(t *testing.T, path string) {
	SplitJSONPathKeys(path)
	})
}

func Fuzz_Nosy_WithFieldSetter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var p *Precompile[E]
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
	if p == nil {
		return
	}

	WithFieldSetter(p)
	})
}

func Fuzz_Nosy_WithFieldsOnly__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var p *Precompile[E]
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
	if p == nil {
		return
	}

	WithFieldsOnly(p)
	})
}

// skipping Fuzz_Nosy_WithPrecompileAtAddress__ because parameters include func, chan, or unsupported interface: E

func Fuzz_Nosy_WithScript__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var h *Host
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var contract string
		fill_err = tp.Fill(&contract)
		if fill_err != nil {
			return
		}
	if h == nil {
		return
	}

	WithScript(h, name, contract)
	})
}

func Fuzz_Nosy_b32__(f *testing.F) {
	f.Fuzz(func(t *testing.T, v uint64) {
	b32(v)
	})
}

func Fuzz_Nosy_bytes4__(f *testing.F) {
	f.Fuzz(func(t *testing.T, sig string) {
	bytes4(sig)
	})
}

func Fuzz_Nosy_checkABI__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var abiData *abi.ABI
		fill_err = tp.Fill(&abiData)
		if fill_err != nil {
			return
		}
		var methodSignature string
		fill_err = tp.Fill(&methodSignature)
		if fill_err != nil {
			return
		}
	if abiData == nil {
		return
	}

	checkABI(abiData, methodSignature)
	})
}

// skipping Fuzz_Nosy_consoleFormat__ because parameters include func, chan, or unsupported interface: []any

// skipping Fuzz_Nosy_convertType__ because parameters include func, chan, or unsupported interface: any

// skipping Fuzz_Nosy_encodeRevert__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_envList__ because parameters include func, chan, or unsupported interface: func(key string) (string, bool)

// skipping Fuzz_Nosy_envOrList__ because parameters include func, chan, or unsupported interface: func(key string) (string, bool)

// skipping Fuzz_Nosy_envOrSingular__ because parameters include func, chan, or unsupported interface: func(key string) (string, bool)

// skipping Fuzz_Nosy_envSingular__ because parameters include func, chan, or unsupported interface: func(key string) (string, bool)

func Fuzz_Nosy_formatBigInt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var x *big.Int
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		var precision int
		fill_err = tp.Fill(&precision)
		if fill_err != nil {
			return
		}
	if x == nil {
		return
	}

	formatBigInt(x, precision)
	})
}

func Fuzz_Nosy_formatBigIntFixedPrecision__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var x *big.Int
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		var precision uint
		fill_err = tp.Fill(&precision)
		if fill_err != nil {
			return
		}
	if x == nil {
		return
	}

	formatBigIntFixedPrecision(x, precision)
	})
}

// skipping Fuzz_Nosy_formatValue__ because parameters include func, chan, or unsupported interface: any

// skipping Fuzz_Nosy_goTypeToSolidityType__ because parameters include func, chan, or unsupported interface: reflect.Type

// skipping Fuzz_Nosy_hasTrailingError__ because parameters include func, chan, or unsupported interface: func(i int) reflect.Type

func Fuzz_Nosy_leftPad32__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d1 []byte) {
	leftPad32(d1)
	})
}

// skipping Fuzz_Nosy_lookupKeys__ because parameters include func, chan, or unsupported interface: any

// skipping Fuzz_Nosy_makeArgAllocators__ because parameters include func, chan, or unsupported interface: func(i int) reflect.Type

func Fuzz_Nosy_makeArgTypes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var args abi.Arguments
		fill_err = tp.Fill(&args)
		if fill_err != nil {
			return
		}

	makeArgTypes(args)
	})
}

// skipping Fuzz_Nosy_makeFn__ because parameters include func, chan, or unsupported interface: []func() any

func Fuzz_Nosy_rightPad32__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d1 []byte) {
	rightPad32(d1)
	})
}

func Fuzz_Nosy_takePath__(f *testing.F) {
	f.Fuzz(func(t *testing.T, query string) {
	takePath(query)
	})
}

