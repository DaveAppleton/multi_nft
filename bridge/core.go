// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// BridgeABI is the input ABI used to generate the binding from.
const BridgeABI = "[{\"inputs\":[{\"internalType\":\"contractvUSD\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"cwd\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractvUSD\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"drain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// BridgeFuncSigs maps the 4-byte function signature to its string representation.
var BridgeFuncSigs = map[string]string{
	"e3f7691a": "cwd(address,uint256)",
	"ece53132": "drain(address)",
}

// BridgeBin is the compiled bytecode used for deploying new contracts.
var BridgeBin = "0x608060405234801561001057600080fd5b506102ea806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063e3f7691a1461003b578063ece5313214610069575b600080fd5b6100676004803603604081101561005157600080fd5b506001600160a01b03813516906020013561008f565b005b6100676004803603602081101561007f57600080fd5b50356001600160a01b03166101ad565b604080516370a0823160e01b815230600482015290516001600160a01b038416916370a08231916024808301926020929190829003018186803b1580156100d557600080fd5b505afa1580156100e9573d6000803e3d6000fd5b505050506040513d60208110156100ff57600080fd5b5051811061014b576040805162461bcd60e51b8152602060048201526014602482015273496e73756666696369656e742062616c616e636560601b604482015290519081900360640190fd5b816001600160a01b0316632e1a7d4d826040518263ffffffff1660e01b815260040180828152602001915050600060405180830381600087803b15801561019157600080fd5b505af11580156101a5573d6000803e3d6000fd5b505050505050565b806001600160a01b031663a9059cbb33836001600160a01b03166370a08231306040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b15801561020a57600080fd5b505afa15801561021e573d6000803e3d6000fd5b505050506040513d602081101561023457600080fd5b5051604080516001600160e01b031960e086901b1681526001600160a01b03909316600484015260248301919091525160448083019260209291908290030181600087803b15801561028557600080fd5b505af1158015610299573d6000803e3d6000fd5b505050506040513d60208110156102af57600080fd5b50505056fea264697066735822122008c102267541635acb1e104883c2f3972e1c22d1ecda137957a4d826c1fb8d9c64736f6c63430007050033"

// DeployBridge deploys a new Ethereum contract, binding an instance of Bridge to it.
func DeployBridge(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Bridge, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BridgeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bridge{BridgeCaller: BridgeCaller{contract: contract}, BridgeTransactor: BridgeTransactor{contract: contract}, BridgeFilterer: BridgeFilterer{contract: contract}}, nil
}

// Bridge is an auto generated Go binding around an Ethereum contract.
type Bridge struct {
	BridgeCaller     // Read-only binding to the contract
	BridgeTransactor // Write-only binding to the contract
	BridgeFilterer   // Log filterer for contract events
}

// BridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeSession struct {
	Contract     *Bridge           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeCallerSession struct {
	Contract *BridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeTransactorSession struct {
	Contract     *BridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeRaw struct {
	Contract *Bridge // Generic contract binding to access the raw methods on
}

// BridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeCallerRaw struct {
	Contract *BridgeCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeTransactorRaw struct {
	Contract *BridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridge creates a new instance of Bridge, bound to a specific deployed contract.
func NewBridge(address common.Address, backend bind.ContractBackend) (*Bridge, error) {
	contract, err := bindBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bridge{BridgeCaller: BridgeCaller{contract: contract}, BridgeTransactor: BridgeTransactor{contract: contract}, BridgeFilterer: BridgeFilterer{contract: contract}}, nil
}

// NewBridgeCaller creates a new read-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeCaller(address common.Address, caller bind.ContractCaller) (*BridgeCaller, error) {
	contract, err := bindBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeCaller{contract: contract}, nil
}

// NewBridgeTransactor creates a new write-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeTransactor, error) {
	contract, err := bindBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeTransactor{contract: contract}, nil
}

// NewBridgeFilterer creates a new log filterer instance of Bridge, bound to a specific deployed contract.
func NewBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeFilterer, error) {
	contract, err := bindBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeFilterer{contract: contract}, nil
}

// bindBridge binds a generic wrapper to an already deployed contract.
func bindBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.BridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transact(opts, method, params...)
}

// Cwd is a paid mutator transaction binding the contract method 0xe3f7691a.
//
// Solidity: function cwd(address token, uint256 amount) returns()
func (_Bridge *BridgeTransactor) Cwd(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "cwd", token, amount)
}

// Cwd is a paid mutator transaction binding the contract method 0xe3f7691a.
//
// Solidity: function cwd(address token, uint256 amount) returns()
func (_Bridge *BridgeSession) Cwd(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.Cwd(&_Bridge.TransactOpts, token, amount)
}

// Cwd is a paid mutator transaction binding the contract method 0xe3f7691a.
//
// Solidity: function cwd(address token, uint256 amount) returns()
func (_Bridge *BridgeTransactorSession) Cwd(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.Cwd(&_Bridge.TransactOpts, token, amount)
}

// Drain is a paid mutator transaction binding the contract method 0xece53132.
//
// Solidity: function drain(address token) returns()
func (_Bridge *BridgeTransactor) Drain(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "drain", token)
}

// Drain is a paid mutator transaction binding the contract method 0xece53132.
//
// Solidity: function drain(address token) returns()
func (_Bridge *BridgeSession) Drain(token common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.Drain(&_Bridge.TransactOpts, token)
}

// Drain is a paid mutator transaction binding the contract method 0xece53132.
//
// Solidity: function drain(address token) returns()
func (_Bridge *BridgeTransactorSession) Drain(token common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.Drain(&_Bridge.TransactOpts, token)
}

// VUSDABI is the input ABI used to generate the binding from.
const VUSDABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dest\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// VUSDFuncSigs maps the 4-byte function signature to its string representation.
var VUSDFuncSigs = map[string]string{
	"70a08231": "balanceOf(address)",
	"a9059cbb": "transfer(address,uint256)",
	"2e1a7d4d": "withdraw(uint256)",
}

// VUSD is an auto generated Go binding around an Ethereum contract.
type VUSD struct {
	VUSDCaller     // Read-only binding to the contract
	VUSDTransactor // Write-only binding to the contract
	VUSDFilterer   // Log filterer for contract events
}

// VUSDCaller is an auto generated read-only Go binding around an Ethereum contract.
type VUSDCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VUSDTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VUSDTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VUSDFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VUSDFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VUSDSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VUSDSession struct {
	Contract     *VUSD             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VUSDCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VUSDCallerSession struct {
	Contract *VUSDCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VUSDTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VUSDTransactorSession struct {
	Contract     *VUSDTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VUSDRaw is an auto generated low-level Go binding around an Ethereum contract.
type VUSDRaw struct {
	Contract *VUSD // Generic contract binding to access the raw methods on
}

// VUSDCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VUSDCallerRaw struct {
	Contract *VUSDCaller // Generic read-only contract binding to access the raw methods on
}

// VUSDTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VUSDTransactorRaw struct {
	Contract *VUSDTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVUSD creates a new instance of VUSD, bound to a specific deployed contract.
func NewVUSD(address common.Address, backend bind.ContractBackend) (*VUSD, error) {
	contract, err := bindVUSD(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VUSD{VUSDCaller: VUSDCaller{contract: contract}, VUSDTransactor: VUSDTransactor{contract: contract}, VUSDFilterer: VUSDFilterer{contract: contract}}, nil
}

// NewVUSDCaller creates a new read-only instance of VUSD, bound to a specific deployed contract.
func NewVUSDCaller(address common.Address, caller bind.ContractCaller) (*VUSDCaller, error) {
	contract, err := bindVUSD(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VUSDCaller{contract: contract}, nil
}

// NewVUSDTransactor creates a new write-only instance of VUSD, bound to a specific deployed contract.
func NewVUSDTransactor(address common.Address, transactor bind.ContractTransactor) (*VUSDTransactor, error) {
	contract, err := bindVUSD(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VUSDTransactor{contract: contract}, nil
}

// NewVUSDFilterer creates a new log filterer instance of VUSD, bound to a specific deployed contract.
func NewVUSDFilterer(address common.Address, filterer bind.ContractFilterer) (*VUSDFilterer, error) {
	contract, err := bindVUSD(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VUSDFilterer{contract: contract}, nil
}

// bindVUSD binds a generic wrapper to an already deployed contract.
func bindVUSD(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VUSDABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VUSD *VUSDRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VUSD.Contract.VUSDCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VUSD *VUSDRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VUSD.Contract.VUSDTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VUSD *VUSDRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VUSD.Contract.VUSDTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VUSD *VUSDCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VUSD.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VUSD *VUSDTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VUSD.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VUSD *VUSDTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VUSD.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address dest) view returns(uint256)
func (_VUSD *VUSDCaller) BalanceOf(opts *bind.CallOpts, dest common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VUSD.contract.Call(opts, &out, "balanceOf", dest)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address dest) view returns(uint256)
func (_VUSD *VUSDSession) BalanceOf(dest common.Address) (*big.Int, error) {
	return _VUSD.Contract.BalanceOf(&_VUSD.CallOpts, dest)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address dest) view returns(uint256)
func (_VUSD *VUSDCallerSession) BalanceOf(dest common.Address) (*big.Int, error) {
	return _VUSD.Contract.BalanceOf(&_VUSD.CallOpts, dest)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address , uint256 ) returns(bool)
func (_VUSD *VUSDTransactor) Transfer(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _VUSD.contract.Transact(opts, "transfer", arg0, arg1)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address , uint256 ) returns(bool)
func (_VUSD *VUSDSession) Transfer(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _VUSD.Contract.Transfer(&_VUSD.TransactOpts, arg0, arg1)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address , uint256 ) returns(bool)
func (_VUSD *VUSDTransactorSession) Transfer(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _VUSD.Contract.Transfer(&_VUSD.TransactOpts, arg0, arg1)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_VUSD *VUSDTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _VUSD.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_VUSD *VUSDSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _VUSD.Contract.Withdraw(&_VUSD.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_VUSD *VUSDTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _VUSD.Contract.Withdraw(&_VUSD.TransactOpts, amount)
}
