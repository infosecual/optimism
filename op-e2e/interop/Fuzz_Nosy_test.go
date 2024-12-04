package interop

import (
	"crypto/ecdsa"
	"math/big"
	"testing"

	"github.com/ethereum-optimism/optimism/op-chain-ops/devkeys"
	"github.com/ethereum-optimism/optimism/op-chain-ops/interopgen"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils/geth"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils/opnode"
	"github.com/ethereum/go-ethereum/ethclient"
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

func Fuzz_Nosy_interopE2ESystem_AddDependency__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *interopE2ESystem
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var id string
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var dep *big.Int
		fill_err = tp.Fill(&dep)
		if fill_err != nil {
			return
		}
		if s == nil || dep == nil {
			return
		}

		s.AddDependency(id, dep)
	})
}

func Fuzz_Nosy_interopE2ESystem_AddUser__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *interopE2ESystem
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var username string
		fill_err = tp.Fill(&username)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.AddUser(username)
	})
}

func Fuzz_Nosy_interopE2ESystem_Address__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *interopE2ESystem
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var id string
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var username string
		fill_err = tp.Fill(&username)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Address(id, username)
	})
}

func Fuzz_Nosy_interopE2ESystem_Batcher__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *interopE2ESystem
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var id string
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Batcher(id)
	})
}

func Fuzz_Nosy_interopE2ESystem_ChainID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *interopE2ESystem
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var network string
		fill_err = tp.Fill(&network)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.ChainID(network)
	})
}

func Fuzz_Nosy_interopE2ESystem_Contract__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *interopE2ESystem
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var id string
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Contract(id, name)
	})
}

func Fuzz_Nosy_interopE2ESystem_DeployEmitterContract__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *interopE2ESystem
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var id string
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var sender string
		fill_err = tp.Fill(&sender)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.DeployEmitterContract(id, sender)
	})
}

func Fuzz_Nosy_interopE2ESystem_EmitData__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *interopE2ESystem
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var id string
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var sender string
		fill_err = tp.Fill(&sender)
		if fill_err != nil {
			return
		}
		var d4 string
		fill_err = tp.Fill(&d4)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.EmitData(id, sender, d4)
	})
}

// skipping Fuzz_Nosy_interopE2ESystem_ExecuteMessage__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_interopE2ESystem_L1GethClient__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *interopE2ESystem
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.L1GethClient()
	})
}

func Fuzz_Nosy_interopE2ESystem_L2Geth__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *interopE2ESystem
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var id string
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.L2Geth(id)
	})
}

func Fuzz_Nosy_interopE2ESystem_L2GethClient__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *interopE2ESystem
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var id string
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.L2GethClient(id)
	})
}

func Fuzz_Nosy_interopE2ESystem_L2IDs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *interopE2ESystem
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.L2IDs()
	})
}

func Fuzz_Nosy_interopE2ESystem_L2OperatorKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *interopE2ESystem
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var id string
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var role devkeys.ChainOperatorRole
		fill_err = tp.Fill(&role)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.L2OperatorKey(id, role)
	})
}

func Fuzz_Nosy_interopE2ESystem_OpNode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *interopE2ESystem
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var id string
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.OpNode(id)
	})
}

func Fuzz_Nosy_interopE2ESystem_Proposer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *interopE2ESystem
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var id string
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Proposer(id)
	})
}

// skipping Fuzz_Nosy_interopE2ESystem_SendL2Tx__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/system/helpers.TxOptsFn

func Fuzz_Nosy_interopE2ESystem_Supervisor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *interopE2ESystem
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Supervisor()
	})
}

func Fuzz_Nosy_interopE2ESystem_SupervisorClient__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *interopE2ESystem
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.SupervisorClient()
	})
}

func Fuzz_Nosy_interopE2ESystem_UserKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *interopE2ESystem
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var id string
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var username string
		fill_err = tp.Fill(&username)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.UserKey(id, username)
	})
}

func Fuzz_Nosy_interopE2ESystem____(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *interopE2ESystem
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var id string
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var operatorKeys map[devkeys.ChainOperatorRole]ecdsa.PrivateKey
		fill_err = tp.Fill(&operatorKeys)
		if fill_err != nil {
			return
		}
		var opNode *opnode.Opnode
		fill_err = tp.Fill(&opNode)
		if fill_err != nil {
			return
		}
		if s == nil || opNode == nil {
			return
		}

		s._(id, operatorKeys, opNode)
	})
}

func Fuzz_Nosy_interopE2ESystem_addL2GethClient__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sys *interopE2ESystem
		fill_err = tp.Fill(&sys)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var client *ethclient.Client
		fill_err = tp.Fill(&client)
		if fill_err != nil {
			return
		}
		if sys == nil || client == nil {
			return
		}

		sys.addL2GethClient(name, client)
	})
}

func Fuzz_Nosy_interopE2ESystem_newBatcherForL2__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *interopE2ESystem
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var id string
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var operatorKeys map[devkeys.ChainOperatorRole]ecdsa.PrivateKey
		fill_err = tp.Fill(&operatorKeys)
		if fill_err != nil {
			return
		}
		var l2Geth *geth.GethInstance
		fill_err = tp.Fill(&l2Geth)
		if fill_err != nil {
			return
		}
		var opNode *opnode.Opnode
		fill_err = tp.Fill(&opNode)
		if fill_err != nil {
			return
		}
		if s == nil || l2Geth == nil || opNode == nil {
			return
		}

		s.newBatcherForL2(id, operatorKeys, l2Geth, opNode)
	})
}

func Fuzz_Nosy_interopE2ESystem_newGethForL2__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *interopE2ESystem
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var id string
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var l2Out *interopgen.L2Output
		fill_err = tp.Fill(&l2Out)
		if fill_err != nil {
			return
		}
		if s == nil || l2Out == nil {
			return
		}

		s.newGethForL2(id, l2Out)
	})
}

func Fuzz_Nosy_interopE2ESystem_newL2__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *interopE2ESystem
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var id string
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var l2Out *interopgen.L2Output
		fill_err = tp.Fill(&l2Out)
		if fill_err != nil {
			return
		}
		if s == nil || l2Out == nil {
			return
		}

		s.newL2(id, l2Out)
	})
}

func Fuzz_Nosy_interopE2ESystem_newNodeForL2__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *interopE2ESystem
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var id string
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var l2Out *interopgen.L2Output
		fill_err = tp.Fill(&l2Out)
		if fill_err != nil {
			return
		}
		var operatorKeys map[devkeys.ChainOperatorRole]ecdsa.PrivateKey
		fill_err = tp.Fill(&operatorKeys)
		if fill_err != nil {
			return
		}
		var l2Geth *geth.GethInstance
		fill_err = tp.Fill(&l2Geth)
		if fill_err != nil {
			return
		}
		if s == nil || l2Out == nil || l2Geth == nil {
			return
		}

		s.newNodeForL2(id, l2Out, operatorKeys, l2Geth)
	})
}

func Fuzz_Nosy_interopE2ESystem_newOperatorKeysForL2__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *interopE2ESystem
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var l2Out *interopgen.L2Output
		fill_err = tp.Fill(&l2Out)
		if fill_err != nil {
			return
		}
		if s == nil || l2Out == nil {
			return
		}

		s.newOperatorKeysForL2(l2Out)
	})
}

func Fuzz_Nosy_interopE2ESystem_prepare__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *interopE2ESystem
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var t2 *testing.T
		fill_err = tp.Fill(&t2)
		if fill_err != nil {
			return
		}
		var w worldResourcePaths
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if s == nil || t2 == nil {
			return
		}

		s.prepare(t2, w)
	})
}

func Fuzz_Nosy_interopE2ESystem_prepareContracts__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *interopE2ESystem
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.prepareContracts()
	})
}

func Fuzz_Nosy_interopE2ESystem_prepareHDWallet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *interopE2ESystem
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.prepareHDWallet()
	})
}

func Fuzz_Nosy_interopE2ESystem_prepareL1__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *interopE2ESystem
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.prepareL1()
	})
}

func Fuzz_Nosy_interopE2ESystem_prepareL2s__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *interopE2ESystem
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.prepareL2s()
	})
}

func Fuzz_Nosy_interopE2ESystem_prepareSupervisor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *interopE2ESystem
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.prepareSupervisor()
	})
}

func Fuzz_Nosy_interopE2ESystem_prepareWorld__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *interopE2ESystem
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var w worldResourcePaths
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.prepareWorld(w)
	})
}

func Fuzz_Nosy_SuperSystem_AddDependency__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var recipe *interopgen.InteropDevRecipe
		fill_err = tp.Fill(&recipe)
		if fill_err != nil {
			return
		}
		var w worldResourcePaths
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var config SuperSystemConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var network string
		fill_err = tp.Fill(&network)
		if fill_err != nil {
			return
		}
		var dep *big.Int
		fill_err = tp.Fill(&dep)
		if fill_err != nil {
			return
		}
		if t1 == nil || recipe == nil || dep == nil {
			return
		}

		_x1 := NewSuperSystem(t1, recipe, w, config)
		_x1.AddDependency(network, dep)
	})
}

func Fuzz_Nosy_SuperSystem_AddUser__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var recipe *interopgen.InteropDevRecipe
		fill_err = tp.Fill(&recipe)
		if fill_err != nil {
			return
		}
		var w worldResourcePaths
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var config SuperSystemConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var username string
		fill_err = tp.Fill(&username)
		if fill_err != nil {
			return
		}
		if t1 == nil || recipe == nil {
			return
		}

		_x1 := NewSuperSystem(t1, recipe, w, config)
		_x1.AddUser(username)
	})
}

func Fuzz_Nosy_SuperSystem_Address__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var recipe *interopgen.InteropDevRecipe
		fill_err = tp.Fill(&recipe)
		if fill_err != nil {
			return
		}
		var w worldResourcePaths
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var config SuperSystemConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var network string
		fill_err = tp.Fill(&network)
		if fill_err != nil {
			return
		}
		var username string
		fill_err = tp.Fill(&username)
		if fill_err != nil {
			return
		}
		if t1 == nil || recipe == nil {
			return
		}

		_x1 := NewSuperSystem(t1, recipe, w, config)
		_x1.Address(network, username)
	})
}

func Fuzz_Nosy_SuperSystem_Batcher__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var recipe *interopgen.InteropDevRecipe
		fill_err = tp.Fill(&recipe)
		if fill_err != nil {
			return
		}
		var w worldResourcePaths
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var config SuperSystemConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var network string
		fill_err = tp.Fill(&network)
		if fill_err != nil {
			return
		}
		if t1 == nil || recipe == nil {
			return
		}

		_x1 := NewSuperSystem(t1, recipe, w, config)
		_x1.Batcher(network)
	})
}

func Fuzz_Nosy_SuperSystem_ChainID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var recipe *interopgen.InteropDevRecipe
		fill_err = tp.Fill(&recipe)
		if fill_err != nil {
			return
		}
		var w worldResourcePaths
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var config SuperSystemConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var network string
		fill_err = tp.Fill(&network)
		if fill_err != nil {
			return
		}
		if t1 == nil || recipe == nil {
			return
		}

		_x1 := NewSuperSystem(t1, recipe, w, config)
		_x1.ChainID(network)
	})
}

func Fuzz_Nosy_SuperSystem_Contract__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var recipe *interopgen.InteropDevRecipe
		fill_err = tp.Fill(&recipe)
		if fill_err != nil {
			return
		}
		var w worldResourcePaths
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var config SuperSystemConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var network string
		fill_err = tp.Fill(&network)
		if fill_err != nil {
			return
		}
		var contractName string
		fill_err = tp.Fill(&contractName)
		if fill_err != nil {
			return
		}
		if t1 == nil || recipe == nil {
			return
		}

		_x1 := NewSuperSystem(t1, recipe, w, config)
		_x1.Contract(network, contractName)
	})
}

func Fuzz_Nosy_SuperSystem_DeployEmitterContract__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var recipe *interopgen.InteropDevRecipe
		fill_err = tp.Fill(&recipe)
		if fill_err != nil {
			return
		}
		var w worldResourcePaths
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var config SuperSystemConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var network string
		fill_err = tp.Fill(&network)
		if fill_err != nil {
			return
		}
		var username string
		fill_err = tp.Fill(&username)
		if fill_err != nil {
			return
		}
		if t1 == nil || recipe == nil {
			return
		}

		_x1 := NewSuperSystem(t1, recipe, w, config)
		_x1.DeployEmitterContract(network, username)
	})
}

func Fuzz_Nosy_SuperSystem_EmitData__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var recipe *interopgen.InteropDevRecipe
		fill_err = tp.Fill(&recipe)
		if fill_err != nil {
			return
		}
		var w worldResourcePaths
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var config SuperSystemConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var network string
		fill_err = tp.Fill(&network)
		if fill_err != nil {
			return
		}
		var username string
		fill_err = tp.Fill(&username)
		if fill_err != nil {
			return
		}
		var d7 string
		fill_err = tp.Fill(&d7)
		if fill_err != nil {
			return
		}
		if t1 == nil || recipe == nil {
			return
		}

		_x1 := NewSuperSystem(t1, recipe, w, config)
		_x1.EmitData(network, username, d7)
	})
}

// skipping Fuzz_Nosy_SuperSystem_ExecuteMessage__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_SuperSystem_L2Geth__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var recipe *interopgen.InteropDevRecipe
		fill_err = tp.Fill(&recipe)
		if fill_err != nil {
			return
		}
		var w worldResourcePaths
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var config SuperSystemConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var network string
		fill_err = tp.Fill(&network)
		if fill_err != nil {
			return
		}
		if t1 == nil || recipe == nil {
			return
		}

		_x1 := NewSuperSystem(t1, recipe, w, config)
		_x1.L2Geth(network)
	})
}

func Fuzz_Nosy_SuperSystem_L2GethClient__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var recipe *interopgen.InteropDevRecipe
		fill_err = tp.Fill(&recipe)
		if fill_err != nil {
			return
		}
		var w worldResourcePaths
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var config SuperSystemConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var network string
		fill_err = tp.Fill(&network)
		if fill_err != nil {
			return
		}
		if t1 == nil || recipe == nil {
			return
		}

		_x1 := NewSuperSystem(t1, recipe, w, config)
		_x1.L2GethClient(network)
	})
}

func Fuzz_Nosy_SuperSystem_L2IDs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var recipe *interopgen.InteropDevRecipe
		fill_err = tp.Fill(&recipe)
		if fill_err != nil {
			return
		}
		var w worldResourcePaths
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var config SuperSystemConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		if t1 == nil || recipe == nil {
			return
		}

		_x1 := NewSuperSystem(t1, recipe, w, config)
		_x1.L2IDs()
	})
}

func Fuzz_Nosy_SuperSystem_L2OperatorKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var recipe *interopgen.InteropDevRecipe
		fill_err = tp.Fill(&recipe)
		if fill_err != nil {
			return
		}
		var w worldResourcePaths
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var config SuperSystemConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var network string
		fill_err = tp.Fill(&network)
		if fill_err != nil {
			return
		}
		var role devkeys.ChainOperatorRole
		fill_err = tp.Fill(&role)
		if fill_err != nil {
			return
		}
		if t1 == nil || recipe == nil {
			return
		}

		_x1 := NewSuperSystem(t1, recipe, w, config)
		_x1.L2OperatorKey(network, role)
	})
}

func Fuzz_Nosy_SuperSystem_OpNode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var recipe *interopgen.InteropDevRecipe
		fill_err = tp.Fill(&recipe)
		if fill_err != nil {
			return
		}
		var w worldResourcePaths
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var config SuperSystemConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var network string
		fill_err = tp.Fill(&network)
		if fill_err != nil {
			return
		}
		if t1 == nil || recipe == nil {
			return
		}

		_x1 := NewSuperSystem(t1, recipe, w, config)
		_x1.OpNode(network)
	})
}

func Fuzz_Nosy_SuperSystem_Proposer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var recipe *interopgen.InteropDevRecipe
		fill_err = tp.Fill(&recipe)
		if fill_err != nil {
			return
		}
		var w worldResourcePaths
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var config SuperSystemConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var network string
		fill_err = tp.Fill(&network)
		if fill_err != nil {
			return
		}
		if t1 == nil || recipe == nil {
			return
		}

		_x1 := NewSuperSystem(t1, recipe, w, config)
		_x1.Proposer(network)
	})
}

// skipping Fuzz_Nosy_SuperSystem_SendL2Tx__ because parameters include func, chan, or unsupported interface: github.com/ethereum-optimism/optimism/op-e2e/system/helpers.TxOptsFn

func Fuzz_Nosy_SuperSystem_Supervisor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var recipe *interopgen.InteropDevRecipe
		fill_err = tp.Fill(&recipe)
		if fill_err != nil {
			return
		}
		var w worldResourcePaths
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var config SuperSystemConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		if t1 == nil || recipe == nil {
			return
		}

		_x1 := NewSuperSystem(t1, recipe, w, config)
		_x1.Supervisor()
	})
}

func Fuzz_Nosy_SuperSystem_SupervisorClient__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var recipe *interopgen.InteropDevRecipe
		fill_err = tp.Fill(&recipe)
		if fill_err != nil {
			return
		}
		var w worldResourcePaths
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var config SuperSystemConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		if t1 == nil || recipe == nil {
			return
		}

		_x1 := NewSuperSystem(t1, recipe, w, config)
		_x1.SupervisorClient()
	})
}

func Fuzz_Nosy_SuperSystem_UserKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var recipe *interopgen.InteropDevRecipe
		fill_err = tp.Fill(&recipe)
		if fill_err != nil {
			return
		}
		var w worldResourcePaths
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var config SuperSystemConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var id string
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var username string
		fill_err = tp.Fill(&username)
		if fill_err != nil {
			return
		}
		if t1 == nil || recipe == nil {
			return
		}

		_x1 := NewSuperSystem(t1, recipe, w, config)
		_x1.UserKey(id, username)
	})
}

// skipping Fuzz_Nosy_mustDial__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/log.Logger

// skipping Fuzz_Nosy_writeDefaultJWT__ because parameters include func, chan, or unsupported interface: testing.TB
