package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
)

type StateDB struct {
	Bytecode []byte
}

func (s *StateDB) AddAddressToAccessList(addr common.Address)                {}
func (s *StateDB) AddBalance(addr common.Address, amount *big.Int)           {}
func (s *StateDB) AddLog(log *types.Log)                                     {}
func (s *StateDB) AddPreimage(hash common.Hash, preimage []byte)             {}
func (s *StateDB) AddRefund(gas uint64)                                      {}
func (s *StateDB) AddSlotToAccessList(addr common.Address, slot common.Hash) {}
func (s *StateDB) AddressInAccessList(addr common.Address) bool              { return true }
func (s *StateDB) CreateAccount(addr common.Address)                         {}
func (s *StateDB) Empty(addr common.Address) bool                            { return false }
func (s *StateDB) Exist(addr common.Address) bool                            { return true }
func (b *StateDB) ForEachStorage(addr common.Address, cb func(key, value common.Hash) bool) error {
	return nil
}
func (s *StateDB) GetBalance(addr common.Address) *big.Int { return common.Big0 }
func (s *StateDB) GetCode(addr common.Address) []byte {
	fmt.Println("GetCode", addr)
	return s.Bytecode
}
func (s *StateDB) GetCodeHash(addr common.Address) common.Hash { return common.Hash{} }
func (s *StateDB) GetCodeSize(addr common.Address) int         { return 100 }
func (s *StateDB) GetCommittedState(addr common.Address, hash common.Hash) common.Hash {
	return common.Hash{}
}
func (s *StateDB) GetNonce(addr common.Address) uint64 { return 0 }
func (s *StateDB) GetRefund() uint64                   { return 0 }
func (s *StateDB) GetState(addr common.Address, hash common.Hash) common.Hash {
	fmt.Println("GetState", addr, hash)
	return common.Hash{}
}
func (s *StateDB) HasSuicided(addr common.Address) bool { return false }
func (s *StateDB) PrepareAccessList(sender common.Address, dst *common.Address, precompiles []common.Address, list types.AccessList) {
}
func (s *StateDB) RevertToSnapshot(revid int)                           {}
func (s *StateDB) SetCode(addr common.Address, code []byte)             {}
func (s *StateDB) SetNonce(addr common.Address, nonce uint64)           {}
func (s *StateDB) SetState(addr common.Address, key, value common.Hash) {}
func (s *StateDB) SlotInAccessList(addr common.Address, slot common.Hash) (addressPresent bool, slotPresent bool) {
	return true, true
}
func (s *StateDB) Snapshot() int                                   { return 0 }
func (s *StateDB) SubBalance(addr common.Address, amount *big.Int) {}
func (s *StateDB) SubRefund(gas uint64)                            {}
func (s *StateDB) Suicide(addr common.Address) bool                { return true }

// **** stub Tracer ****
type Tracer struct{}

// CaptureStart implements the Tracer interface to initialize the tracing operation.
func (jst *Tracer) CaptureStart(env *vm.EVM, from common.Address, to common.Address, create bool, input []byte, gas uint64, value *big.Int) {
}

// CaptureState implements the Tracer interface to trace a single step of VM execution.
func (jst *Tracer) CaptureState(env *vm.EVM, pc uint64, op vm.OpCode, gas, cost uint64, scope *vm.ScopeContext, rData []byte, depth int, err error) {
	//fmt.Println(pc, op, gas)
}

// CaptureFault implements the Tracer interface to trace an execution fault
func (jst *Tracer) CaptureFault(env *vm.EVM, pc uint64, op vm.OpCode, gas, cost uint64, scope *vm.ScopeContext, depth int, err error) {
}

// CaptureEnd is called after the call finishes to finalize the tracing.
func (jst *Tracer) CaptureEnd(output []byte, gasUsed uint64, t time.Duration, err error) {
}

type jsoncontract struct {
	Bytecode         string `json:"bytecode"`
	DeployedBytecode string `json:"deployedBytecode"`
}

//var ram []byte
//var regs [4096]byte

var pcCount int = 0
var debug int = 0
var ram map[uint64](uint32)

func opStaticCall(pc *uint64, interpreter *vm.EVMInterpreter, scope *vm.ScopeContext) ([]byte, error) {
	// Pop gas. The actual gas is in interpreter.evm.callGasTemp.
	stack := scope.Stack

	temp := stack.Pop()
	returnGas := temp.Uint64()
	_, inOffset, inSize, retOffset, retSize := stack.Pop(), stack.Pop(), stack.Pop(), stack.Pop(), stack.Pop()
	//fmt.Println(temp, addr, inOffset, inSize, retOffset, retSize)

	temp.SetOne()
	stack.Push(&temp)

	// Get arguments from the memory.
	args := scope.Memory.GetPtr(int64(inOffset.Uint64()), int64(inSize.Uint64()))
	if args[0] == 98 {
		// read
		addr := common.BytesToHash(args[4:]).Big().Uint64()
		nret := ram[addr]

		//scope.Memory.GetPtr(int64(inOffset.Uint64()), int64(inSize.Uint64()))

		ret := common.BigToHash(big.NewInt(int64(nret))).Bytes()
		if debug >= 2 {
			fmt.Println("HOOKED READ!   ", fmt.Sprintf("%x = %x", addr, nret))
		}
		if addr == 0xc0000080 && debug >= 1 {
			fmt.Printf("%7d: PC %x\n", pcCount, nret)
			pcCount += 1
		}
		scope.Memory.Set(retOffset.Uint64(), retSize.Uint64(), ret)
	} else if args[0] == 184 {
		addr := common.BytesToHash(args[0x24:0x44]).Big().Uint64()
		dat := common.BytesToHash(args[0x44:0x64]).Big().Uint64()
		if debug >= 2 {
			fmt.Println("HOOKED WRITE!  ", fmt.Sprintf("%x = %x", addr, dat))
		}
		ram[addr] = uint32(dat)

		// pass through stateRoot
		scope.Memory.Set(retOffset.Uint64(), retSize.Uint64(), args[0x4:0x24])
	}

	// ram access is free here
	scope.Contract.Gas = returnGas
	// what is the return value here?
	return common.Hash{}.Bytes(), nil
}

func runMinigeth(fn string, interpreter *vm.EVMInterpreter, bytecode []byte) {
	ram = make(map[uint64](uint32))
	dat, _ := ioutil.ReadFile(fn)
	for i := 0; i < len(dat); i += 4 {
		ram[uint64(i)] = uint32(dat[i])<<24 |
			uint32(dat[i+1])<<16 |
			uint32(dat[i+2])<<8 |
			uint32(dat[i+3])<<0
	}

	steps := 100000
	gas := 10000 * uint64(steps)

	// 0xdb7df598
	from := common.Address{}
	to := common.HexToAddress("0x1337")
	input := []byte{0xdb, 0x7d, 0xf5, 0x98} // Steps(bytes32, uint256)
	input = append(input, common.BigToHash(common.Big0).Bytes()...)
	input = append(input, common.BigToHash(big.NewInt(int64(steps))).Bytes()...)

	debug = 1
	for i := 0; i < 1; i++ {
		contract := vm.NewContract(vm.AccountRef(from), vm.AccountRef(to), common.Big0, gas)
		contract.SetCallCode(&to, crypto.Keccak256Hash(bytecode), bytecode)
		_, err := interpreter.Run(contract, input, false)
		fmt.Printf("step:%d pc:0x%x v0:%d ", i, ram[0xc0000080], ram[0xc0000008])
		fmt.Println(err)

		ram[0xc0000008] = 0
		ram[0xc0000000+7*4] = 0
	}
}

func runTest(fn string, steps int, interpreter *vm.EVMInterpreter, bytecode []byte, gas uint64) {
	ram = make(map[uint64](uint32))
	ram[0xC000007C] = 0xDEAD0000
	//fmt.Println("starting", fn)
	dat, _ := ioutil.ReadFile(fn)
	for i := 0; i < len(dat); i += 4 {
		ram[uint64(i)] = uint32(dat[i])<<24 |
			uint32(dat[i+1])<<16 |
			uint32(dat[i+2])<<8 |
			uint32(dat[i+3])<<0
	}

	// 0xdb7df598
	from := common.Address{}
	to := common.HexToAddress("0x1337")
	input := []byte{0xdb, 0x7d, 0xf5, 0x98} // Steps(bytes32, uint256)
	input = append(input, common.BigToHash(common.Big0).Bytes()...)
	input = append(input, common.BigToHash(big.NewInt(int64(steps))).Bytes()...)
	contract := vm.NewContract(vm.AccountRef(from), vm.AccountRef(to), common.Big0, gas)
	//fmt.Println(bytecodehash, bytecode)
	contract.SetCallCode(&to, crypto.Keccak256Hash(bytecode), bytecode)

	start := time.Now()
	_, err := interpreter.Run(contract, input, false)
	elapsed := time.Now().Sub(start)

	fmt.Println(err, contract.Gas, elapsed,
		ram[0xbffffff4], ram[0xbffffff8], fmt.Sprintf("%x", ram[0xc0000080]), fn)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Println("hello")

	/*var parent types.Header
	database := state.NewDatabase(parent)
	statedb, _ := state.New(parent.Root, database, nil)*/

	var jj jsoncontract
	mipsjson, _ := ioutil.ReadFile("../artifacts/contracts/MIPS.sol/MIPS.json")
	json.NewDecoder(bytes.NewReader(mipsjson)).Decode(&jj)
	bytecode := common.Hex2Bytes(jj.DeployedBytecode[2:])
	//fmt.Println(bytecode, jj.Bytecode)
	statedb := &StateDB{Bytecode: bytecode}

	var header types.Header
	header.Number = big.NewInt(13284469)
	header.Difficulty = common.Big0
	bc := core.NewBlockChain(&header)
	author := common.Address{}
	blockContext := core.NewEVMBlockContext(&header, bc, &author)
	txContext := vm.TxContext{}
	config := vm.Config{}
	config.Debug = true
	tracer := Tracer{}
	config.Tracer = &tracer
	evm := vm.NewEVM(blockContext, txContext, statedb, params.MainnetChainConfig, config)
	//fmt.Println(evm)

	/*ret, gas, err := evm.Call(vm.AccountRef(from), to, []byte{}, 20000000, common.Big0)
	fmt.Println(ret, gas, err)*/

	interpreter := vm.NewEVMInterpreter(evm, config)
	interpreter.GetCfg().JumpTable[vm.STATICCALL].SetExecute(opStaticCall)

	/*input := []byte{0x69, 0x37, 0x33, 0x72} // Step(bytes32)
	input = append(input, common.Hash{}.Bytes()...)*/

	// 1.26s for 100000 steps
	//steps := 20
	// 19.100079097s for 1_000_000 new steps
	//steps := 1000000
	//debug = true

	if len(os.Args) > 1 {
		if os.Args[1] == "/tmp/minigeth.bin" {
			/*debug = 1
			steps := 1000000
			runTest(os.Args[1], steps, interpreter, bytecode, uint64(steps)*10000)*/
			runMinigeth(os.Args[1], interpreter, bytecode)
		} else {
			debug = 2
			runTest(os.Args[1], 20, interpreter, bytecode, 1000000)
		}
	} else {
		files, err := ioutil.ReadDir("test/bin")
		if err != nil {
			log.Fatal(err)
		}
		for _, f := range files {
			runTest("test/bin/"+f.Name(), 100, interpreter, bytecode, 1000000)
		}
	}

}
