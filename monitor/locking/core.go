// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package locking

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

// IERC20ABI is the input ABI used to generate the binding from.
const IERC20ABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"x\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IERC20FuncSigs maps the 4-byte function signature to its string representation.
var IERC20FuncSigs = map[string]string{
	"70a08231": "balanceOf(address)",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// IERC20 is an auto generated Go binding around an Ethereum contract.
type IERC20 struct {
	IERC20Caller     // Read-only binding to the contract
	IERC20Transactor // Write-only binding to the contract
	IERC20Filterer   // Log filterer for contract events
}

// IERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20Session struct {
	Contract     *IERC20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CallerSession struct {
	Contract *IERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20TransactorSession struct {
	Contract     *IERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20Raw struct {
	Contract *IERC20 // Generic contract binding to access the raw methods on
}

// IERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CallerRaw struct {
	Contract *IERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20TransactorRaw struct {
	Contract *IERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20 creates a new instance of IERC20, bound to a specific deployed contract.
func NewIERC20(address common.Address, backend bind.ContractBackend) (*IERC20, error) {
	contract, err := bindIERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// NewIERC20Caller creates a new read-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Caller(address common.Address, caller bind.ContractCaller) (*IERC20Caller, error) {
	contract, err := bindIERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Caller{contract: contract}, nil
}

// NewIERC20Transactor creates a new write-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC20Transactor, error) {
	contract, err := bindIERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Transactor{contract: contract}, nil
}

// NewIERC20Filterer creates a new log filterer instance of IERC20, bound to a specific deployed contract.
func NewIERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC20Filterer, error) {
	contract, err := bindIERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20Filterer{contract: contract}, nil
}

// bindIERC20 binds a generic wrapper to an already deployed contract.
func bindIERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.IERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a paid mutator transaction binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address x) returns(uint256)
func (_IERC20 *IERC20Transactor) BalanceOf(opts *bind.TransactOpts, x common.Address) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "balanceOf", x)
}

// BalanceOf is a paid mutator transaction binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address x) returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(x common.Address) (*types.Transaction, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.TransactOpts, x)
}

// BalanceOf is a paid mutator transaction binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address x) returns(uint256)
func (_IERC20 *IERC20TransactorSession) BalanceOf(x common.Address) (*types.Transaction, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.TransactOpts, x)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address target, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, target common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", target, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address target, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Transfer(target common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, target, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address target, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(target common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, target, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address owner, address target, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, owner common.Address, target common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", owner, target, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address owner, address target, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(owner common.Address, target common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, owner, target, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address owner, address target, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(owner common.Address, target common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, owner, target, amount)
}

// LockingABI is the input ABI used to generate the binding from.
const LockingABI = "[{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_baseToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountToken\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LockedBaseToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountTokenP\",\"type\":\"uint256\"}],\"name\":\"SetLockAmount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"UnlockCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"UnlockRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawBaseTokens\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"LockData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"releaseRequest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"amountToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"drain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"isLocked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onTokenTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"timeToWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountToken\",\"type\":\"uint256\"}],\"name\":\"updateLockAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// LockingFuncSigs maps the 4-byte function signature to its string representation.
var LockingFuncSigs = map[string]string{
	"6ecb9d50": "LockData(address)",
	"102ff0b3": "amountToken()",
	"c55dae63": "baseToken()",
	"ce606ee0": "contractOwner()",
	"ece53132": "drain(address)",
	"4a4fbeec": "isLocked(address)",
	"f83d08ba": "lock()",
	"a4c0ed36": "onTokenTransfer(address,uint256,bytes)",
	"49dcc4eb": "timeToWithdraw(address)",
	"a69df4b5": "unlock()",
	"3c37db63": "updateLockAmount(uint256)",
}

// LockingBin is the compiled bytecode used for deploying new contracts.
var LockingBin = "0x608060405234801561001057600080fd5b50604051610e00380380610e008339818101604052604081101561003357600080fd5b50805160209091015160008054336001600160a01b031991821617909155600180549091166001600160a01b03841617905561006e81610075565b505061010f565b6000546001600160a01b031633146100d4576040805162461bcd60e51b815260206004820152601360248201527f556e617574686f72697365642061636365737300000000000000000000000000604482015290519081900360640190fd5b60028190556040805182815290517f97e0c5afc28fab93c1626222b791d272e9d38ca5e480e54ce8cbb7aeb615f4a59181900360200190a150565b610ce28061011e6000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c8063a4c0ed3611610071578063a4c0ed361461018d578063a69df4b514610248578063c55dae6314610250578063ce606ee014610274578063ece531321461027c578063f83d08ba146102a2576100a9565b8063102ff0b3146100ae5780633c37db63146100c857806349dcc4eb146100e75780634a4fbeec1461010d5780636ecb9d5014610147575b600080fd5b6100b66102aa565b60408051918252519081900360200190f35b6100e5600480360360208110156100de57600080fd5b50356102b0565b005b6100b6600480360360208110156100fd57600080fd5b50356001600160a01b0316610340565b6101336004803603602081101561012357600080fd5b50356001600160a01b0316610420565b604080519115158252519081900360200190f35b61016d6004803603602081101561015d57600080fd5b50356001600160a01b0316610467565b604080519384526020840192909252151582820152519081900360600190f35b6100e5600480360360608110156101a357600080fd5b6001600160a01b03823516916020810135918101906060810160408201356401000000008111156101d357600080fd5b8201836020820111156101e557600080fd5b8035906020019184600183028401116401000000008311171561020757600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061048b945050505050565b6100e561053c565b6102586107a0565b604080516001600160a01b039092168252519081900360200190f35b6102586107af565b6100e56004803603602081101561029257600080fd5b50356001600160a01b03166107be565b6100e56109c1565b60025481565b6000546001600160a01b03163314610305576040805162461bcd60e51b8152602060048201526013602482015272556e617574686f72697365642061636365737360681b604482015290519081900360640190fd5b60028190556040805182815290517f97e0c5afc28fab93c1626222b791d272e9d38ca5e480e54ce8cbb7aeb615f4a59181900360200190a150565b6001600160a01b03811660009081526003602052604081206002015460ff166103b0576040805162461bcd60e51b815260206004820152601960248201527f4f776e657220686173206e6f2066756e6473206c6f636b656400000000000000604482015290519081900360640190fd5b6001600160a01b038216600090815260036020526040902054806104055760405162461bcd60e51b8152600401808060200182810382526025815260200180610c886025913960400191505060405180910390fd5b42811061041657600091505061041b565b420390505b919050565b6001600160a01b03811660009081526003602052604081206002015460ff16801561046157506001600160a01b038216600090815260036020526040902054155b92915050565b60036020526000908152604090208054600182015460029092015490919060ff1683565b6001546001600160a01b031633146104e0576040805162461bcd60e51b8152602060048201526013602482015272556e617574686f726973656420736f7572636560681b604482015290519081900360640190fd5b600254821461052d576040805162461bcd60e51b8152602060048201526014602482015273125b98dbdc9c9958dd081d985b1d59481cd95b9d60621b604482015290519081900360640190fd5b6105378383610af4565b505050565b610544610c1b565b50336000908152600360209081526040918290208251606081018452815481526001820154928101929092526002015460ff161515918101829052906105d1576040805162461bcd60e51b815260206004820152601d60248201527f596f7520646f206e6f742068617665206c6f636b656420746f6b656e73000000604482015290519081900360640190fd5b602081015161061b576040805162461bcd60e51b81526020600482015260116024820152704e6f7468696e6720746f20756e6c6f636b60781b604482015290519081900360640190fd5b8051801561074f57428111156106625760405162461bcd60e51b8152600401808060200182810382526027815260200180610c3f6027913960400191505060405180910390fd5b6001546020808401516040805163a9059cbb60e01b81523360048201526024810192909252516001600160a01b039093169263a9059cbb926044808401939192918290030181600087803b1580156106b957600080fd5b505af11580156106cd573d6000803e3d6000fd5b505050506040513d60208110156106e357600080fd5b505033600081815260036020908152604080832083815560018101939093556002909201805460ff191690558481015182519384529083015280517ffe3b7acaad2fd3b2a0c954f99523de4322c0792bcfa0051a3e284e05e8530e4c9281900390910190a1505061079e565b336000818152600360209081526040918290206206978042019055815192835290517f9626a556bdcacf287e1acddc78733a299483d4511693e5211ffaacad062137bc9281900390910190a150505b565b6001546001600160a01b031681565b6000546001600160a01b031681565b6000546001600160a01b0316331461080c576040805162461bcd60e51b815260206004820152600c60248201526b155b985d5d1a1bdc9a5cd95960a21b604482015290519081900360640190fd5b6001600160a01b03811661085a57600080546040516001600160a01b03909116914780156108fc02929091818181858888f19350505050158015610854573d6000803e3d6000fd5b506109be565b6001546001600160a01b03828116911614156108a75760405162461bcd60e51b8152600401808060200182810382526022815260200180610c666022913960400191505060405180910390fd5b806001600160a01b031663a9059cbb60008054906101000a90046001600160a01b0316836001600160a01b03166370a08231306040518263ffffffff1660e01b815260040180826001600160a01b03168152602001915050602060405180830381600087803b15801561091957600080fd5b505af115801561092d573d6000803e3d6000fd5b505050506040513d602081101561094357600080fd5b5051604080516001600160e01b031960e086901b1681526001600160a01b03909316600484015260248301919091525160448083019260209291908290030181600087803b15801561099457600080fd5b505af11580156109a8573d6000803e3d6000fd5b505050506040513d602081101561053757600080fd5b50565b3360009081526003602052604090205415610a225733600081815260036020908152604080832092909255815192835290517ffa044b7b93a40365dc68049797c2eb06918523d694e5d56e406cac3eb35578e59281900390910190a161079e565b600254600154604080516323b872dd60e01b81523360048201523060248201526044810184905290516001600160a01b03909216916323b872dd916064808201926020929091908290030181600087803b158015610a7f57600080fd5b505af1158015610a93573d6000803e3d6000fd5b505050506040513d6020811015610aa957600080fd5b5051610aee576040805162461bcd60e51b815260206004820152600f60248201526e1d1c985b9cd9995c8819985a5b1959608a1b604482015290519081900360640190fd5b6109be33825b610afc610c1b565b506001600160a01b0382166000908152600360209081526040918290208251606081018452815481526001820154928101929092526002015460ff161580159282019290925290610b94576040805162461bcd60e51b815260206004820152601e60248201527f596f7520616c72656164792068617665206c6f636b656420746f6b656e730000604482015290519081900360640190fd5b6001604082810182815260208085018681526001600160a01b0388166000818152600384528590208751815591519582019590955591516002909201805460ff1916921515929092179091558151928352820184905280517fb00e71ae128757f10a89bed65c351adc01e3a28456245a1e31f738f364bf24919281900390910190a1505050565b60405180606001604052806000815260200160008152602001600015158152509056fe72656c6561736520616c72656164792072657175657374656420627574206e6f74207265616479596f752063616e6e6f7420776974686472617720746865206261736520746f6b656e4f776e657220686173206e6f74207265717565737465642066756e64732072656c65617365a264697066735822122078d633542696be01d3bcd576a59ca8e6992030e577a55c0d46e4d576a8c8b80464736f6c63430007050033"

// DeployLocking deploys a new Ethereum contract, binding an instance of Locking to it.
func DeployLocking(auth *bind.TransactOpts, backend bind.ContractBackend, _baseToken common.Address, _amountToken *big.Int) (common.Address, *types.Transaction, *Locking, error) {
	parsed, err := abi.JSON(strings.NewReader(LockingABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(LockingBin), backend, _baseToken, _amountToken)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Locking{LockingCaller: LockingCaller{contract: contract}, LockingTransactor: LockingTransactor{contract: contract}, LockingFilterer: LockingFilterer{contract: contract}}, nil
}

// Locking is an auto generated Go binding around an Ethereum contract.
type Locking struct {
	LockingCaller     // Read-only binding to the contract
	LockingTransactor // Write-only binding to the contract
	LockingFilterer   // Log filterer for contract events
}

// LockingCaller is an auto generated read-only Go binding around an Ethereum contract.
type LockingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LockingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LockingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LockingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LockingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LockingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LockingSession struct {
	Contract     *Locking          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LockingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LockingCallerSession struct {
	Contract *LockingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// LockingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LockingTransactorSession struct {
	Contract     *LockingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// LockingRaw is an auto generated low-level Go binding around an Ethereum contract.
type LockingRaw struct {
	Contract *Locking // Generic contract binding to access the raw methods on
}

// LockingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LockingCallerRaw struct {
	Contract *LockingCaller // Generic read-only contract binding to access the raw methods on
}

// LockingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LockingTransactorRaw struct {
	Contract *LockingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLocking creates a new instance of Locking, bound to a specific deployed contract.
func NewLocking(address common.Address, backend bind.ContractBackend) (*Locking, error) {
	contract, err := bindLocking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Locking{LockingCaller: LockingCaller{contract: contract}, LockingTransactor: LockingTransactor{contract: contract}, LockingFilterer: LockingFilterer{contract: contract}}, nil
}

// NewLockingCaller creates a new read-only instance of Locking, bound to a specific deployed contract.
func NewLockingCaller(address common.Address, caller bind.ContractCaller) (*LockingCaller, error) {
	contract, err := bindLocking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LockingCaller{contract: contract}, nil
}

// NewLockingTransactor creates a new write-only instance of Locking, bound to a specific deployed contract.
func NewLockingTransactor(address common.Address, transactor bind.ContractTransactor) (*LockingTransactor, error) {
	contract, err := bindLocking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LockingTransactor{contract: contract}, nil
}

// NewLockingFilterer creates a new log filterer instance of Locking, bound to a specific deployed contract.
func NewLockingFilterer(address common.Address, filterer bind.ContractFilterer) (*LockingFilterer, error) {
	contract, err := bindLocking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LockingFilterer{contract: contract}, nil
}

// bindLocking binds a generic wrapper to an already deployed contract.
func bindLocking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LockingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Locking *LockingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Locking.Contract.LockingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Locking *LockingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Locking.Contract.LockingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Locking *LockingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Locking.Contract.LockingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Locking *LockingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Locking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Locking *LockingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Locking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Locking *LockingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Locking.Contract.contract.Transact(opts, method, params...)
}

// LockData is a free data retrieval call binding the contract method 0x6ecb9d50.
//
// Solidity: function LockData(address ) view returns(uint256 releaseRequest, uint256 amount, bool locked)
func (_Locking *LockingCaller) LockData(opts *bind.CallOpts, arg0 common.Address) (struct {
	ReleaseRequest *big.Int
	Amount         *big.Int
	Locked         bool
}, error) {
	var out []interface{}
	err := _Locking.contract.Call(opts, &out, "LockData", arg0)

	outstruct := new(struct {
		ReleaseRequest *big.Int
		Amount         *big.Int
		Locked         bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ReleaseRequest = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Locked = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// LockData is a free data retrieval call binding the contract method 0x6ecb9d50.
//
// Solidity: function LockData(address ) view returns(uint256 releaseRequest, uint256 amount, bool locked)
func (_Locking *LockingSession) LockData(arg0 common.Address) (struct {
	ReleaseRequest *big.Int
	Amount         *big.Int
	Locked         bool
}, error) {
	return _Locking.Contract.LockData(&_Locking.CallOpts, arg0)
}

// LockData is a free data retrieval call binding the contract method 0x6ecb9d50.
//
// Solidity: function LockData(address ) view returns(uint256 releaseRequest, uint256 amount, bool locked)
func (_Locking *LockingCallerSession) LockData(arg0 common.Address) (struct {
	ReleaseRequest *big.Int
	Amount         *big.Int
	Locked         bool
}, error) {
	return _Locking.Contract.LockData(&_Locking.CallOpts, arg0)
}

// AmountToken is a free data retrieval call binding the contract method 0x102ff0b3.
//
// Solidity: function amountToken() view returns(uint256)
func (_Locking *LockingCaller) AmountToken(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Locking.contract.Call(opts, &out, "amountToken")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AmountToken is a free data retrieval call binding the contract method 0x102ff0b3.
//
// Solidity: function amountToken() view returns(uint256)
func (_Locking *LockingSession) AmountToken() (*big.Int, error) {
	return _Locking.Contract.AmountToken(&_Locking.CallOpts)
}

// AmountToken is a free data retrieval call binding the contract method 0x102ff0b3.
//
// Solidity: function amountToken() view returns(uint256)
func (_Locking *LockingCallerSession) AmountToken() (*big.Int, error) {
	return _Locking.Contract.AmountToken(&_Locking.CallOpts)
}

// BaseToken is a free data retrieval call binding the contract method 0xc55dae63.
//
// Solidity: function baseToken() view returns(address)
func (_Locking *LockingCaller) BaseToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Locking.contract.Call(opts, &out, "baseToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BaseToken is a free data retrieval call binding the contract method 0xc55dae63.
//
// Solidity: function baseToken() view returns(address)
func (_Locking *LockingSession) BaseToken() (common.Address, error) {
	return _Locking.Contract.BaseToken(&_Locking.CallOpts)
}

// BaseToken is a free data retrieval call binding the contract method 0xc55dae63.
//
// Solidity: function baseToken() view returns(address)
func (_Locking *LockingCallerSession) BaseToken() (common.Address, error) {
	return _Locking.Contract.BaseToken(&_Locking.CallOpts)
}

// ContractOwner is a free data retrieval call binding the contract method 0xce606ee0.
//
// Solidity: function contractOwner() view returns(address)
func (_Locking *LockingCaller) ContractOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Locking.contract.Call(opts, &out, "contractOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ContractOwner is a free data retrieval call binding the contract method 0xce606ee0.
//
// Solidity: function contractOwner() view returns(address)
func (_Locking *LockingSession) ContractOwner() (common.Address, error) {
	return _Locking.Contract.ContractOwner(&_Locking.CallOpts)
}

// ContractOwner is a free data retrieval call binding the contract method 0xce606ee0.
//
// Solidity: function contractOwner() view returns(address)
func (_Locking *LockingCallerSession) ContractOwner() (common.Address, error) {
	return _Locking.Contract.ContractOwner(&_Locking.CallOpts)
}

// IsLocked is a free data retrieval call binding the contract method 0x4a4fbeec.
//
// Solidity: function isLocked(address owner) view returns(bool)
func (_Locking *LockingCaller) IsLocked(opts *bind.CallOpts, owner common.Address) (bool, error) {
	var out []interface{}
	err := _Locking.contract.Call(opts, &out, "isLocked", owner)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLocked is a free data retrieval call binding the contract method 0x4a4fbeec.
//
// Solidity: function isLocked(address owner) view returns(bool)
func (_Locking *LockingSession) IsLocked(owner common.Address) (bool, error) {
	return _Locking.Contract.IsLocked(&_Locking.CallOpts, owner)
}

// IsLocked is a free data retrieval call binding the contract method 0x4a4fbeec.
//
// Solidity: function isLocked(address owner) view returns(bool)
func (_Locking *LockingCallerSession) IsLocked(owner common.Address) (bool, error) {
	return _Locking.Contract.IsLocked(&_Locking.CallOpts, owner)
}

// TimeToWithdraw is a free data retrieval call binding the contract method 0x49dcc4eb.
//
// Solidity: function timeToWithdraw(address owner) view returns(uint256)
func (_Locking *LockingCaller) TimeToWithdraw(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Locking.contract.Call(opts, &out, "timeToWithdraw", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TimeToWithdraw is a free data retrieval call binding the contract method 0x49dcc4eb.
//
// Solidity: function timeToWithdraw(address owner) view returns(uint256)
func (_Locking *LockingSession) TimeToWithdraw(owner common.Address) (*big.Int, error) {
	return _Locking.Contract.TimeToWithdraw(&_Locking.CallOpts, owner)
}

// TimeToWithdraw is a free data retrieval call binding the contract method 0x49dcc4eb.
//
// Solidity: function timeToWithdraw(address owner) view returns(uint256)
func (_Locking *LockingCallerSession) TimeToWithdraw(owner common.Address) (*big.Int, error) {
	return _Locking.Contract.TimeToWithdraw(&_Locking.CallOpts, owner)
}

// Drain is a paid mutator transaction binding the contract method 0xece53132.
//
// Solidity: function drain(address token) returns()
func (_Locking *LockingTransactor) Drain(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Locking.contract.Transact(opts, "drain", token)
}

// Drain is a paid mutator transaction binding the contract method 0xece53132.
//
// Solidity: function drain(address token) returns()
func (_Locking *LockingSession) Drain(token common.Address) (*types.Transaction, error) {
	return _Locking.Contract.Drain(&_Locking.TransactOpts, token)
}

// Drain is a paid mutator transaction binding the contract method 0xece53132.
//
// Solidity: function drain(address token) returns()
func (_Locking *LockingTransactorSession) Drain(token common.Address) (*types.Transaction, error) {
	return _Locking.Contract.Drain(&_Locking.TransactOpts, token)
}

// Lock is a paid mutator transaction binding the contract method 0xf83d08ba.
//
// Solidity: function lock() returns()
func (_Locking *LockingTransactor) Lock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Locking.contract.Transact(opts, "lock")
}

// Lock is a paid mutator transaction binding the contract method 0xf83d08ba.
//
// Solidity: function lock() returns()
func (_Locking *LockingSession) Lock() (*types.Transaction, error) {
	return _Locking.Contract.Lock(&_Locking.TransactOpts)
}

// Lock is a paid mutator transaction binding the contract method 0xf83d08ba.
//
// Solidity: function lock() returns()
func (_Locking *LockingTransactorSession) Lock() (*types.Transaction, error) {
	return _Locking.Contract.Lock(&_Locking.TransactOpts)
}

// OnTokenTransfer is a paid mutator transaction binding the contract method 0xa4c0ed36.
//
// Solidity: function onTokenTransfer(address _sender, uint256 _value, bytes ) returns()
func (_Locking *LockingTransactor) OnTokenTransfer(opts *bind.TransactOpts, _sender common.Address, _value *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _Locking.contract.Transact(opts, "onTokenTransfer", _sender, _value, arg2)
}

// OnTokenTransfer is a paid mutator transaction binding the contract method 0xa4c0ed36.
//
// Solidity: function onTokenTransfer(address _sender, uint256 _value, bytes ) returns()
func (_Locking *LockingSession) OnTokenTransfer(_sender common.Address, _value *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _Locking.Contract.OnTokenTransfer(&_Locking.TransactOpts, _sender, _value, arg2)
}

// OnTokenTransfer is a paid mutator transaction binding the contract method 0xa4c0ed36.
//
// Solidity: function onTokenTransfer(address _sender, uint256 _value, bytes ) returns()
func (_Locking *LockingTransactorSession) OnTokenTransfer(_sender common.Address, _value *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _Locking.Contract.OnTokenTransfer(&_Locking.TransactOpts, _sender, _value, arg2)
}

// Unlock is a paid mutator transaction binding the contract method 0xa69df4b5.
//
// Solidity: function unlock() returns()
func (_Locking *LockingTransactor) Unlock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Locking.contract.Transact(opts, "unlock")
}

// Unlock is a paid mutator transaction binding the contract method 0xa69df4b5.
//
// Solidity: function unlock() returns()
func (_Locking *LockingSession) Unlock() (*types.Transaction, error) {
	return _Locking.Contract.Unlock(&_Locking.TransactOpts)
}

// Unlock is a paid mutator transaction binding the contract method 0xa69df4b5.
//
// Solidity: function unlock() returns()
func (_Locking *LockingTransactorSession) Unlock() (*types.Transaction, error) {
	return _Locking.Contract.Unlock(&_Locking.TransactOpts)
}

// UpdateLockAmount is a paid mutator transaction binding the contract method 0x3c37db63.
//
// Solidity: function updateLockAmount(uint256 _amountToken) returns()
func (_Locking *LockingTransactor) UpdateLockAmount(opts *bind.TransactOpts, _amountToken *big.Int) (*types.Transaction, error) {
	return _Locking.contract.Transact(opts, "updateLockAmount", _amountToken)
}

// UpdateLockAmount is a paid mutator transaction binding the contract method 0x3c37db63.
//
// Solidity: function updateLockAmount(uint256 _amountToken) returns()
func (_Locking *LockingSession) UpdateLockAmount(_amountToken *big.Int) (*types.Transaction, error) {
	return _Locking.Contract.UpdateLockAmount(&_Locking.TransactOpts, _amountToken)
}

// UpdateLockAmount is a paid mutator transaction binding the contract method 0x3c37db63.
//
// Solidity: function updateLockAmount(uint256 _amountToken) returns()
func (_Locking *LockingTransactorSession) UpdateLockAmount(_amountToken *big.Int) (*types.Transaction, error) {
	return _Locking.Contract.UpdateLockAmount(&_Locking.TransactOpts, _amountToken)
}

// LockingLockedBaseTokenIterator is returned from FilterLockedBaseToken and is used to iterate over the raw logs and unpacked data for LockedBaseToken events raised by the Locking contract.
type LockingLockedBaseTokenIterator struct {
	Event *LockingLockedBaseToken // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LockingLockedBaseTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockingLockedBaseToken)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LockingLockedBaseToken)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LockingLockedBaseTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockingLockedBaseTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockingLockedBaseToken represents a LockedBaseToken event raised by the Locking contract.
type LockingLockedBaseToken struct {
	Owner  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLockedBaseToken is a free log retrieval operation binding the contract event 0xb00e71ae128757f10a89bed65c351adc01e3a28456245a1e31f738f364bf2491.
//
// Solidity: event LockedBaseToken(address owner, uint256 amount)
func (_Locking *LockingFilterer) FilterLockedBaseToken(opts *bind.FilterOpts) (*LockingLockedBaseTokenIterator, error) {

	logs, sub, err := _Locking.contract.FilterLogs(opts, "LockedBaseToken")
	if err != nil {
		return nil, err
	}
	return &LockingLockedBaseTokenIterator{contract: _Locking.contract, event: "LockedBaseToken", logs: logs, sub: sub}, nil
}

// WatchLockedBaseToken is a free log subscription operation binding the contract event 0xb00e71ae128757f10a89bed65c351adc01e3a28456245a1e31f738f364bf2491.
//
// Solidity: event LockedBaseToken(address owner, uint256 amount)
func (_Locking *LockingFilterer) WatchLockedBaseToken(opts *bind.WatchOpts, sink chan<- *LockingLockedBaseToken) (event.Subscription, error) {

	logs, sub, err := _Locking.contract.WatchLogs(opts, "LockedBaseToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockingLockedBaseToken)
				if err := _Locking.contract.UnpackLog(event, "LockedBaseToken", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLockedBaseToken is a log parse operation binding the contract event 0xb00e71ae128757f10a89bed65c351adc01e3a28456245a1e31f738f364bf2491.
//
// Solidity: event LockedBaseToken(address owner, uint256 amount)
func (_Locking *LockingFilterer) ParseLockedBaseToken(log types.Log) (*LockingLockedBaseToken, error) {
	event := new(LockingLockedBaseToken)
	if err := _Locking.contract.UnpackLog(event, "LockedBaseToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LockingSetLockAmountIterator is returned from FilterSetLockAmount and is used to iterate over the raw logs and unpacked data for SetLockAmount events raised by the Locking contract.
type LockingSetLockAmountIterator struct {
	Event *LockingSetLockAmount // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LockingSetLockAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockingSetLockAmount)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LockingSetLockAmount)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LockingSetLockAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockingSetLockAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockingSetLockAmount represents a SetLockAmount event raised by the Locking contract.
type LockingSetLockAmount struct {
	AmountTokenP *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSetLockAmount is a free log retrieval operation binding the contract event 0x97e0c5afc28fab93c1626222b791d272e9d38ca5e480e54ce8cbb7aeb615f4a5.
//
// Solidity: event SetLockAmount(uint256 amountTokenP)
func (_Locking *LockingFilterer) FilterSetLockAmount(opts *bind.FilterOpts) (*LockingSetLockAmountIterator, error) {

	logs, sub, err := _Locking.contract.FilterLogs(opts, "SetLockAmount")
	if err != nil {
		return nil, err
	}
	return &LockingSetLockAmountIterator{contract: _Locking.contract, event: "SetLockAmount", logs: logs, sub: sub}, nil
}

// WatchSetLockAmount is a free log subscription operation binding the contract event 0x97e0c5afc28fab93c1626222b791d272e9d38ca5e480e54ce8cbb7aeb615f4a5.
//
// Solidity: event SetLockAmount(uint256 amountTokenP)
func (_Locking *LockingFilterer) WatchSetLockAmount(opts *bind.WatchOpts, sink chan<- *LockingSetLockAmount) (event.Subscription, error) {

	logs, sub, err := _Locking.contract.WatchLogs(opts, "SetLockAmount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockingSetLockAmount)
				if err := _Locking.contract.UnpackLog(event, "SetLockAmount", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetLockAmount is a log parse operation binding the contract event 0x97e0c5afc28fab93c1626222b791d272e9d38ca5e480e54ce8cbb7aeb615f4a5.
//
// Solidity: event SetLockAmount(uint256 amountTokenP)
func (_Locking *LockingFilterer) ParseSetLockAmount(log types.Log) (*LockingSetLockAmount, error) {
	event := new(LockingSetLockAmount)
	if err := _Locking.contract.UnpackLog(event, "SetLockAmount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LockingUnlockCancelledIterator is returned from FilterUnlockCancelled and is used to iterate over the raw logs and unpacked data for UnlockCancelled events raised by the Locking contract.
type LockingUnlockCancelledIterator struct {
	Event *LockingUnlockCancelled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LockingUnlockCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockingUnlockCancelled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LockingUnlockCancelled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LockingUnlockCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockingUnlockCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockingUnlockCancelled represents a UnlockCancelled event raised by the Locking contract.
type LockingUnlockCancelled struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUnlockCancelled is a free log retrieval operation binding the contract event 0xfa044b7b93a40365dc68049797c2eb06918523d694e5d56e406cac3eb35578e5.
//
// Solidity: event UnlockCancelled(address owner)
func (_Locking *LockingFilterer) FilterUnlockCancelled(opts *bind.FilterOpts) (*LockingUnlockCancelledIterator, error) {

	logs, sub, err := _Locking.contract.FilterLogs(opts, "UnlockCancelled")
	if err != nil {
		return nil, err
	}
	return &LockingUnlockCancelledIterator{contract: _Locking.contract, event: "UnlockCancelled", logs: logs, sub: sub}, nil
}

// WatchUnlockCancelled is a free log subscription operation binding the contract event 0xfa044b7b93a40365dc68049797c2eb06918523d694e5d56e406cac3eb35578e5.
//
// Solidity: event UnlockCancelled(address owner)
func (_Locking *LockingFilterer) WatchUnlockCancelled(opts *bind.WatchOpts, sink chan<- *LockingUnlockCancelled) (event.Subscription, error) {

	logs, sub, err := _Locking.contract.WatchLogs(opts, "UnlockCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockingUnlockCancelled)
				if err := _Locking.contract.UnpackLog(event, "UnlockCancelled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnlockCancelled is a log parse operation binding the contract event 0xfa044b7b93a40365dc68049797c2eb06918523d694e5d56e406cac3eb35578e5.
//
// Solidity: event UnlockCancelled(address owner)
func (_Locking *LockingFilterer) ParseUnlockCancelled(log types.Log) (*LockingUnlockCancelled, error) {
	event := new(LockingUnlockCancelled)
	if err := _Locking.contract.UnpackLog(event, "UnlockCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LockingUnlockRequestedIterator is returned from FilterUnlockRequested and is used to iterate over the raw logs and unpacked data for UnlockRequested events raised by the Locking contract.
type LockingUnlockRequestedIterator struct {
	Event *LockingUnlockRequested // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LockingUnlockRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockingUnlockRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LockingUnlockRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LockingUnlockRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockingUnlockRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockingUnlockRequested represents a UnlockRequested event raised by the Locking contract.
type LockingUnlockRequested struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUnlockRequested is a free log retrieval operation binding the contract event 0x9626a556bdcacf287e1acddc78733a299483d4511693e5211ffaacad062137bc.
//
// Solidity: event UnlockRequested(address owner)
func (_Locking *LockingFilterer) FilterUnlockRequested(opts *bind.FilterOpts) (*LockingUnlockRequestedIterator, error) {

	logs, sub, err := _Locking.contract.FilterLogs(opts, "UnlockRequested")
	if err != nil {
		return nil, err
	}
	return &LockingUnlockRequestedIterator{contract: _Locking.contract, event: "UnlockRequested", logs: logs, sub: sub}, nil
}

// WatchUnlockRequested is a free log subscription operation binding the contract event 0x9626a556bdcacf287e1acddc78733a299483d4511693e5211ffaacad062137bc.
//
// Solidity: event UnlockRequested(address owner)
func (_Locking *LockingFilterer) WatchUnlockRequested(opts *bind.WatchOpts, sink chan<- *LockingUnlockRequested) (event.Subscription, error) {

	logs, sub, err := _Locking.contract.WatchLogs(opts, "UnlockRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockingUnlockRequested)
				if err := _Locking.contract.UnpackLog(event, "UnlockRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnlockRequested is a log parse operation binding the contract event 0x9626a556bdcacf287e1acddc78733a299483d4511693e5211ffaacad062137bc.
//
// Solidity: event UnlockRequested(address owner)
func (_Locking *LockingFilterer) ParseUnlockRequested(log types.Log) (*LockingUnlockRequested, error) {
	event := new(LockingUnlockRequested)
	if err := _Locking.contract.UnpackLog(event, "UnlockRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LockingWithdrawBaseTokensIterator is returned from FilterWithdrawBaseTokens and is used to iterate over the raw logs and unpacked data for WithdrawBaseTokens events raised by the Locking contract.
type LockingWithdrawBaseTokensIterator struct {
	Event *LockingWithdrawBaseTokens // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LockingWithdrawBaseTokensIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockingWithdrawBaseTokens)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LockingWithdrawBaseTokens)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LockingWithdrawBaseTokensIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockingWithdrawBaseTokensIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockingWithdrawBaseTokens represents a WithdrawBaseTokens event raised by the Locking contract.
type LockingWithdrawBaseTokens struct {
	Owner  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawBaseTokens is a free log retrieval operation binding the contract event 0xfe3b7acaad2fd3b2a0c954f99523de4322c0792bcfa0051a3e284e05e8530e4c.
//
// Solidity: event WithdrawBaseTokens(address owner, uint256 amount)
func (_Locking *LockingFilterer) FilterWithdrawBaseTokens(opts *bind.FilterOpts) (*LockingWithdrawBaseTokensIterator, error) {

	logs, sub, err := _Locking.contract.FilterLogs(opts, "WithdrawBaseTokens")
	if err != nil {
		return nil, err
	}
	return &LockingWithdrawBaseTokensIterator{contract: _Locking.contract, event: "WithdrawBaseTokens", logs: logs, sub: sub}, nil
}

// WatchWithdrawBaseTokens is a free log subscription operation binding the contract event 0xfe3b7acaad2fd3b2a0c954f99523de4322c0792bcfa0051a3e284e05e8530e4c.
//
// Solidity: event WithdrawBaseTokens(address owner, uint256 amount)
func (_Locking *LockingFilterer) WatchWithdrawBaseTokens(opts *bind.WatchOpts, sink chan<- *LockingWithdrawBaseTokens) (event.Subscription, error) {

	logs, sub, err := _Locking.contract.WatchLogs(opts, "WithdrawBaseTokens")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockingWithdrawBaseTokens)
				if err := _Locking.contract.UnpackLog(event, "WithdrawBaseTokens", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawBaseTokens is a log parse operation binding the contract event 0xfe3b7acaad2fd3b2a0c954f99523de4322c0792bcfa0051a3e284e05e8530e4c.
//
// Solidity: event WithdrawBaseTokens(address owner, uint256 amount)
func (_Locking *LockingFilterer) ParseWithdrawBaseTokens(log types.Log) (*LockingWithdrawBaseTokens, error) {
	event := new(LockingWithdrawBaseTokens)
	if err := _Locking.contract.UnpackLog(event, "WithdrawBaseTokens", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
