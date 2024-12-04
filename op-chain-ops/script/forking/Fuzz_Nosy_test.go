package forking

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/tracing"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	"github.com/holiman/uint256"
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

func Fuzz_Nosy_AccountDiff_Copy__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *AccountDiff
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.Copy()
	})
}

func Fuzz_Nosy_CachedSource_Balance__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *CachedSource
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

		c.Balance(addr)
	})
}

func Fuzz_Nosy_CachedSource_Code__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *CachedSource
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

		c.Code(addr)
	})
}

func Fuzz_Nosy_CachedSource_Nonce__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *CachedSource
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

		c.Nonce(addr)
	})
}

func Fuzz_Nosy_CachedSource_StateRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *CachedSource
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.StateRoot()
	})
}

func Fuzz_Nosy_CachedSource_StorageAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *CachedSource
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var key common.Hash
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.StorageAt(addr, key)
	})
}

func Fuzz_Nosy_CachedSource_URLOrAlias__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *CachedSource
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.URLOrAlias()
	})
}

func Fuzz_Nosy_ForkDB_ContractCode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *ForkDB
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var codeHash common.Hash
		fill_err = tp.Fill(&codeHash)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.ContractCode(addr, codeHash)
	})
}

func Fuzz_Nosy_ForkDB_ContractCodeSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *ForkDB
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var codeHash common.Hash
		fill_err = tp.Fill(&codeHash)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.ContractCodeSize(addr, codeHash)
	})
}

// skipping Fuzz_Nosy_ForkDB_CopyTrie__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/state.Trie

func Fuzz_Nosy_ForkDB_DiskDB__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *ForkDB
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.DiskDB()
	})
}

// skipping Fuzz_Nosy_ForkDB_OpenStorageTrie__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/state.Trie

func Fuzz_Nosy_ForkDB_OpenTrie__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *ForkDB
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var root common.Hash
		fill_err = tp.Fill(&root)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.OpenTrie(root)
	})
}

func Fuzz_Nosy_ForkDB_PointCache__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *ForkDB
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.PointCache()
	})
}

func Fuzz_Nosy_ForkDB_Reader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *ForkDB
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var root common.Hash
		fill_err = tp.Fill(&root)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.Reader(root)
	})
}

func Fuzz_Nosy_ForkDB_Snapshot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *ForkDB
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.Snapshot()
	})
}

func Fuzz_Nosy_ForkDB_TrieDB__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *ForkDB
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.TrieDB()
	})
}

func Fuzz_Nosy_ForkID_U256__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *big.Int
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		id := ForkIDFromBig(b)
		id.U256()
	})
}

func Fuzz_Nosy_ForkableState_ActiveFork__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fst *ForkableState
		fill_err = tp.Fill(&fst)
		if fill_err != nil {
			return
		}
		if fst == nil {
			return
		}

		fst.ActiveFork()
	})
}

func Fuzz_Nosy_ForkableState_AddAddressToAccessList__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fst *ForkableState
		fill_err = tp.Fill(&fst)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if fst == nil {
			return
		}

		fst.AddAddressToAccessList(addr)
	})
}

func Fuzz_Nosy_ForkableState_AddBalance__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fst *ForkableState
		fill_err = tp.Fill(&fst)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		var u *uint256.Int
		fill_err = tp.Fill(&u)
		if fill_err != nil {
			return
		}
		var reason tracing.BalanceChangeReason
		fill_err = tp.Fill(&reason)
		if fill_err != nil {
			return
		}
		if fst == nil || u == nil {
			return
		}

		fst.AddBalance(address, u, reason)
	})
}

func Fuzz_Nosy_ForkableState_AddLog__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fst *ForkableState
		fill_err = tp.Fill(&fst)
		if fill_err != nil {
			return
		}
		var log *types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if fst == nil || log == nil {
			return
		}

		fst.AddLog(log)
	})
}

func Fuzz_Nosy_ForkableState_AddPreimage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fst *ForkableState
		fill_err = tp.Fill(&fst)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		var img []byte
		fill_err = tp.Fill(&img)
		if fill_err != nil {
			return
		}
		if fst == nil {
			return
		}

		fst.AddPreimage(hash, img)
	})
}

func Fuzz_Nosy_ForkableState_AddRefund__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fst *ForkableState
		fill_err = tp.Fill(&fst)
		if fill_err != nil {
			return
		}
		var u uint64
		fill_err = tp.Fill(&u)
		if fill_err != nil {
			return
		}
		if fst == nil {
			return
		}

		fst.AddRefund(u)
	})
}

func Fuzz_Nosy_ForkableState_AddSlotToAccessList__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fst *ForkableState
		fill_err = tp.Fill(&fst)
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
		if fst == nil {
			return
		}

		fst.AddSlotToAccessList(addr, slot)
	})
}

func Fuzz_Nosy_ForkableState_AddressInAccessList__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fst *ForkableState
		fill_err = tp.Fill(&fst)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if fst == nil {
			return
		}

		fst.AddressInAccessList(addr)
	})
}

func Fuzz_Nosy_ForkableState_CreateAccount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fst *ForkableState
		fill_err = tp.Fill(&fst)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		if fst == nil {
			return
		}

		fst.CreateAccount(address)
	})
}

func Fuzz_Nosy_ForkableState_CreateContract__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fst *ForkableState
		fill_err = tp.Fill(&fst)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		if fst == nil {
			return
		}

		fst.CreateContract(address)
	})
}

// skipping Fuzz_Nosy_ForkableState_CreateFork__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-chain-ops/script/forking.ForkSource

// skipping Fuzz_Nosy_ForkableState_CreateSelectFork__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-chain-ops/script/forking.ForkSource

func Fuzz_Nosy_ForkableState_Empty__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fst *ForkableState
		fill_err = tp.Fill(&fst)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		if fst == nil {
			return
		}

		fst.Empty(address)
	})
}

func Fuzz_Nosy_ForkableState_Exist__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fst *ForkableState
		fill_err = tp.Fill(&fst)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		if fst == nil {
			return
		}

		fst.Exist(address)
	})
}

func Fuzz_Nosy_ForkableState_ExportDiff__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fst *ForkableState
		fill_err = tp.Fill(&fst)
		if fill_err != nil {
			return
		}
		var id ForkID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		if fst == nil {
			return
		}

		fst.ExportDiff(id)
	})
}

func Fuzz_Nosy_ForkableState_Finalise__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fst *ForkableState
		fill_err = tp.Fill(&fst)
		if fill_err != nil {
			return
		}
		var deleteEmptyObjects bool
		fill_err = tp.Fill(&deleteEmptyObjects)
		if fill_err != nil {
			return
		}
		if fst == nil {
			return
		}

		fst.Finalise(deleteEmptyObjects)
	})
}

func Fuzz_Nosy_ForkableState_ForkURLOrAlias__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fst *ForkableState
		fill_err = tp.Fill(&fst)
		if fill_err != nil {
			return
		}
		var id ForkID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		if fst == nil {
			return
		}

		fst.ForkURLOrAlias(id)
	})
}

func Fuzz_Nosy_ForkableState_GetBalance__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fst *ForkableState
		fill_err = tp.Fill(&fst)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		if fst == nil {
			return
		}

		fst.GetBalance(address)
	})
}

func Fuzz_Nosy_ForkableState_GetCode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fst *ForkableState
		fill_err = tp.Fill(&fst)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		if fst == nil {
			return
		}

		fst.GetCode(address)
	})
}

func Fuzz_Nosy_ForkableState_GetCodeHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fst *ForkableState
		fill_err = tp.Fill(&fst)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		if fst == nil {
			return
		}

		fst.GetCodeHash(address)
	})
}

func Fuzz_Nosy_ForkableState_GetCodeSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fst *ForkableState
		fill_err = tp.Fill(&fst)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		if fst == nil {
			return
		}

		fst.GetCodeSize(address)
	})
}

func Fuzz_Nosy_ForkableState_GetCommittedState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fst *ForkableState
		fill_err = tp.Fill(&fst)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if fst == nil {
			return
		}

		fst.GetCommittedState(address, hash)
	})
}

func Fuzz_Nosy_ForkableState_GetNonce__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fst *ForkableState
		fill_err = tp.Fill(&fst)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		if fst == nil {
			return
		}

		fst.GetNonce(address)
	})
}

func Fuzz_Nosy_ForkableState_GetRefund__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fst *ForkableState
		fill_err = tp.Fill(&fst)
		if fill_err != nil {
			return
		}
		if fst == nil {
			return
		}

		fst.GetRefund()
	})
}

func Fuzz_Nosy_ForkableState_GetState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fst *ForkableState
		fill_err = tp.Fill(&fst)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		var k common.Hash
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		if fst == nil {
			return
		}

		fst.GetState(address, k)
	})
}

func Fuzz_Nosy_ForkableState_GetStorageRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fst *ForkableState
		fill_err = tp.Fill(&fst)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if fst == nil {
			return
		}

		fst.GetStorageRoot(addr)
	})
}

func Fuzz_Nosy_ForkableState_GetTransientState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fst *ForkableState
		fill_err = tp.Fill(&fst)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var key common.Hash
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		if fst == nil {
			return
		}

		fst.GetTransientState(addr, key)
	})
}

func Fuzz_Nosy_ForkableState_HasSelfDestructed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fst *ForkableState
		fill_err = tp.Fill(&fst)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		if fst == nil {
			return
		}

		fst.HasSelfDestructed(address)
	})
}

func Fuzz_Nosy_ForkableState_IsPersistent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fst *ForkableState
		fill_err = tp.Fill(&fst)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if fst == nil {
			return
		}

		fst.IsPersistent(addr)
	})
}

func Fuzz_Nosy_ForkableState_MakeExcluded__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fst *ForkableState
		fill_err = tp.Fill(&fst)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if fst == nil {
			return
		}

		fst.MakeExcluded(addr)
	})
}

func Fuzz_Nosy_ForkableState_MakePersistent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fst *ForkableState
		fill_err = tp.Fill(&fst)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if fst == nil {
			return
		}

		fst.MakePersistent(addr)
	})
}

func Fuzz_Nosy_ForkableState_PointCache__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fst *ForkableState
		fill_err = tp.Fill(&fst)
		if fill_err != nil {
			return
		}
		if fst == nil {
			return
		}

		fst.PointCache()
	})
}

func Fuzz_Nosy_ForkableState_Prepare__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fst *ForkableState
		fill_err = tp.Fill(&fst)
		if fill_err != nil {
			return
		}
		var rules params.Rules
		fill_err = tp.Fill(&rules)
		if fill_err != nil {
			return
		}
		var sender common.Address
		fill_err = tp.Fill(&sender)
		if fill_err != nil {
			return
		}
		var coinbase common.Address
		fill_err = tp.Fill(&coinbase)
		if fill_err != nil {
			return
		}
		var dest *common.Address
		fill_err = tp.Fill(&dest)
		if fill_err != nil {
			return
		}
		var precompiles []common.Address
		fill_err = tp.Fill(&precompiles)
		if fill_err != nil {
			return
		}
		var txAccesses types.AccessList
		fill_err = tp.Fill(&txAccesses)
		if fill_err != nil {
			return
		}
		if fst == nil || dest == nil {
			return
		}

		fst.Prepare(rules, sender, coinbase, dest, precompiles, txAccesses)
	})
}

// skipping Fuzz_Nosy_ForkableState_ResetFork__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-chain-ops/script/forking.ForkSource

func Fuzz_Nosy_ForkableState_RevertToSnapshot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fst *ForkableState
		fill_err = tp.Fill(&fst)
		if fill_err != nil {
			return
		}
		var i int
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
		if fst == nil {
			return
		}

		fst.RevertToSnapshot(i)
	})
}

func Fuzz_Nosy_ForkableState_RevokeExcluded__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fst *ForkableState
		fill_err = tp.Fill(&fst)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if fst == nil {
			return
		}

		fst.RevokeExcluded(addr)
	})
}

func Fuzz_Nosy_ForkableState_RevokePersistent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fst *ForkableState
		fill_err = tp.Fill(&fst)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if fst == nil {
			return
		}

		fst.RevokePersistent(addr)
	})
}

func Fuzz_Nosy_ForkableState_SelectFork__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fst *ForkableState
		fill_err = tp.Fill(&fst)
		if fill_err != nil {
			return
		}
		var id ForkID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		if fst == nil {
			return
		}

		fst.SelectFork(id)
	})
}

func Fuzz_Nosy_ForkableState_SelfDestruct__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fst *ForkableState
		fill_err = tp.Fill(&fst)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		if fst == nil {
			return
		}

		fst.SelfDestruct(address)
	})
}

func Fuzz_Nosy_ForkableState_Selfdestruct6780__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fst *ForkableState
		fill_err = tp.Fill(&fst)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		if fst == nil {
			return
		}

		fst.Selfdestruct6780(address)
	})
}

func Fuzz_Nosy_ForkableState_SetBalance__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fst *ForkableState
		fill_err = tp.Fill(&fst)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var amount *uint256.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var reason tracing.BalanceChangeReason
		fill_err = tp.Fill(&reason)
		if fill_err != nil {
			return
		}
		if fst == nil || amount == nil {
			return
		}

		fst.SetBalance(addr, amount, reason)
	})
}

func Fuzz_Nosy_ForkableState_SetCode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fst *ForkableState
		fill_err = tp.Fill(&fst)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		var bytes []byte
		fill_err = tp.Fill(&bytes)
		if fill_err != nil {
			return
		}
		if fst == nil {
			return
		}

		fst.SetCode(address, bytes)
	})
}

func Fuzz_Nosy_ForkableState_SetNonce__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fst *ForkableState
		fill_err = tp.Fill(&fst)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		var u uint64
		fill_err = tp.Fill(&u)
		if fill_err != nil {
			return
		}
		if fst == nil {
			return
		}

		fst.SetNonce(address, u)
	})
}

func Fuzz_Nosy_ForkableState_SetState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fst *ForkableState
		fill_err = tp.Fill(&fst)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		var k common.Hash
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var v common.Hash
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		if fst == nil {
			return
		}

		fst.SetState(address, k, v)
	})
}

func Fuzz_Nosy_ForkableState_SetTransientState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fst *ForkableState
		fill_err = tp.Fill(&fst)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var key common.Hash
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var value common.Hash
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
		if fst == nil {
			return
		}

		fst.SetTransientState(addr, key, value)
	})
}

func Fuzz_Nosy_ForkableState_SlotInAccessList__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fst *ForkableState
		fill_err = tp.Fill(&fst)
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
		if fst == nil {
			return
		}

		fst.SlotInAccessList(addr, slot)
	})
}

func Fuzz_Nosy_ForkableState_Snapshot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fst *ForkableState
		fill_err = tp.Fill(&fst)
		if fill_err != nil {
			return
		}
		if fst == nil {
			return
		}

		fst.Snapshot()
	})
}

func Fuzz_Nosy_ForkableState_SubBalance__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fst *ForkableState
		fill_err = tp.Fill(&fst)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		var u *uint256.Int
		fill_err = tp.Fill(&u)
		if fill_err != nil {
			return
		}
		var reason tracing.BalanceChangeReason
		fill_err = tp.Fill(&reason)
		if fill_err != nil {
			return
		}
		if fst == nil || u == nil {
			return
		}

		fst.SubBalance(address, u, reason)
	})
}

func Fuzz_Nosy_ForkableState_SubRefund__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fst *ForkableState
		fill_err = tp.Fill(&fst)
		if fill_err != nil {
			return
		}
		var u uint64
		fill_err = tp.Fill(&u)
		if fill_err != nil {
			return
		}
		if fst == nil {
			return
		}

		fst.SubRefund(u)
	})
}

// skipping Fuzz_Nosy_ForkableState_SubstituteBaseState__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-chain-ops/script/forking.VMStateDB

func Fuzz_Nosy_ForkableState_Witness__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fst *ForkableState
		fill_err = tp.Fill(&fst)
		if fill_err != nil {
			return
		}
		if fst == nil {
			return
		}

		fst.Witness()
	})
}

func Fuzz_Nosy_ForkableState_stateFor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fst *ForkableState
		fill_err = tp.Fill(&fst)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if fst == nil {
			return
		}

		fst.stateFor(addr)
	})
}

func Fuzz_Nosy_ForkedAccountsTrie_ClearDiff__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *ForkedAccountsTrie
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.ClearDiff()
	})
}

func Fuzz_Nosy_ForkedAccountsTrie_Commit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *ForkedAccountsTrie
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var collectLeaf bool
		fill_err = tp.Fill(&collectLeaf)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.Commit(collectLeaf)
	})
}

func Fuzz_Nosy_ForkedAccountsTrie_ContractCode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *ForkedAccountsTrie
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var codeHash common.Hash
		fill_err = tp.Fill(&codeHash)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.ContractCode(addr, codeHash)
	})
}

func Fuzz_Nosy_ForkedAccountsTrie_ContractCodeSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *ForkedAccountsTrie
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var codeHash common.Hash
		fill_err = tp.Fill(&codeHash)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.ContractCodeSize(addr, codeHash)
	})
}

func Fuzz_Nosy_ForkedAccountsTrie_Copy__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *ForkedAccountsTrie
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.Copy()
	})
}

func Fuzz_Nosy_ForkedAccountsTrie_DeleteAccount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *ForkedAccountsTrie
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.DeleteAccount(address)
	})
}

func Fuzz_Nosy_ForkedAccountsTrie_DeleteStorage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *ForkedAccountsTrie
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var key []byte
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.DeleteStorage(addr, key)
	})
}

func Fuzz_Nosy_ForkedAccountsTrie_ExportDiff__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *ForkedAccountsTrie
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.ExportDiff()
	})
}

func Fuzz_Nosy_ForkedAccountsTrie_GetAccount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *ForkedAccountsTrie
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.GetAccount(address)
	})
}

func Fuzz_Nosy_ForkedAccountsTrie_GetKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *ForkedAccountsTrie
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var bytes []byte
		fill_err = tp.Fill(&bytes)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.GetKey(bytes)
	})
}

func Fuzz_Nosy_ForkedAccountsTrie_GetStorage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *ForkedAccountsTrie
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var key []byte
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.GetStorage(addr, key)
	})
}

func Fuzz_Nosy_ForkedAccountsTrie_HasDiff__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *ForkedAccountsTrie
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.HasDiff()
	})
}

func Fuzz_Nosy_ForkedAccountsTrie_Hash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *ForkedAccountsTrie
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.Hash()
	})
}

func Fuzz_Nosy_ForkedAccountsTrie_IsVerkle__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *ForkedAccountsTrie
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.IsVerkle()
	})
}

func Fuzz_Nosy_ForkedAccountsTrie_NodeIterator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *ForkedAccountsTrie
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var startKey []byte
		fill_err = tp.Fill(&startKey)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.NodeIterator(startKey)
	})
}

// skipping Fuzz_Nosy_ForkedAccountsTrie_Prove__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

func Fuzz_Nosy_ForkedAccountsTrie_UpdateAccount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *ForkedAccountsTrie
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		var account *types.StateAccount
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		var codeLen int
		fill_err = tp.Fill(&codeLen)
		if fill_err != nil {
			return
		}
		if f1 == nil || account == nil {
			return
		}

		f1.UpdateAccount(address, account, codeLen)
	})
}

func Fuzz_Nosy_ForkedAccountsTrie_UpdateContractCode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *ForkedAccountsTrie
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var codeHash common.Hash
		fill_err = tp.Fill(&codeHash)
		if fill_err != nil {
			return
		}
		var code []byte
		fill_err = tp.Fill(&code)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.UpdateContractCode(addr, codeHash, code)
	})
}

func Fuzz_Nosy_ForkedAccountsTrie_UpdateStorage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *ForkedAccountsTrie
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
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
		if f1 == nil {
			return
		}

		f1.UpdateStorage(addr, key, value)
	})
}

func Fuzz_Nosy_ForkedAccountsTrie_Witness__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *ForkedAccountsTrie
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.Witness()
	})
}

func Fuzz_Nosy_RPCSource_Balance__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *RPCSource
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.Balance(addr)
	})
}

func Fuzz_Nosy_RPCSource_BlockHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *RPCSource
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.BlockHash()
	})
}

func Fuzz_Nosy_RPCSource_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *RPCSource
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.Close()
	})
}

func Fuzz_Nosy_RPCSource_Code__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *RPCSource
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.Code(addr)
	})
}

func Fuzz_Nosy_RPCSource_Nonce__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *RPCSource
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.Nonce(addr)
	})
}

func Fuzz_Nosy_RPCSource_StateRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *RPCSource
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.StateRoot()
	})
}

func Fuzz_Nosy_RPCSource_StorageAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *RPCSource
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var key common.Hash
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.StorageAt(addr, key)
	})
}

func Fuzz_Nosy_RPCSource_URLOrAlias__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *RPCSource
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.URLOrAlias()
	})
}

// skipping Fuzz_Nosy_RPCSource_init__ because parameters include func, chan, or unsupported interface: any

func Fuzz_Nosy_forkStateEntry_DB__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fe *forkStateEntry
		fill_err = tp.Fill(&fe)
		if fill_err != nil {
			return
		}
		if fe == nil {
			return
		}

		fe.DB()
	})
}

func Fuzz_Nosy_forkStateReader_Account__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *forkStateReader
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.Account(addr)
	})
}

func Fuzz_Nosy_forkStateReader_Copy__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *forkStateReader
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.Copy()
	})
}

func Fuzz_Nosy_forkStateReader_Storage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *forkStateReader
		fill_err = tp.Fill(&f1)
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
		if f1 == nil {
			return
		}

		f1.Storage(addr, slot)
	})
}

func Fuzz_Nosy_ForkID_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *big.Int
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		id := ForkIDFromBig(b)
		id.String()
	})
}

// skipping Fuzz_Nosy_ForkSource_Balance__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-chain-ops/script/forking.ForkSource

// skipping Fuzz_Nosy_ForkSource_Code__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-chain-ops/script/forking.ForkSource

// skipping Fuzz_Nosy_ForkSource_Nonce__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-chain-ops/script/forking.ForkSource

// skipping Fuzz_Nosy_ForkSource_StateRoot__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-chain-ops/script/forking.ForkSource

// skipping Fuzz_Nosy_ForkSource_StorageAt__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-chain-ops/script/forking.ForkSource

// skipping Fuzz_Nosy_ForkSource_URLOrAlias__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-chain-ops/script/forking.ForkSource

// skipping Fuzz_Nosy_RPCClient_CallContext__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-chain-ops/script/forking.RPCClient

// skipping Fuzz_Nosy_VMStateDB_Finalise__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-chain-ops/script/forking.VMStateDB

// skipping Fuzz_Nosy_VMStateDB_SetBalance__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-chain-ops/script/forking.VMStateDB
