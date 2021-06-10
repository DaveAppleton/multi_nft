// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pSale

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
const IERC20ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IERC20FuncSigs maps the 4-byte function signature to its string representation.
var IERC20FuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"18160ddd": "totalSupply()",
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

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Session) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// IERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20 contract.
type IERC20ApprovalIterator struct {
	Event *IERC20Approval // Event containing the contract specifics and raw log

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
func (it *IERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Approval)
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
		it.Event = new(IERC20Approval)
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
func (it *IERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Approval represents a Approval event raised by the IERC20 contract.
type IERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20ApprovalIterator{contract: _IERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Approval)
				if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) ParseApproval(log types.Log) (*IERC20Approval, error) {
	event := new(IERC20Approval)
	if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20 contract.
type IERC20TransferIterator struct {
	Event *IERC20Transfer // Event containing the contract specifics and raw log

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
func (it *IERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Transfer)
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
		it.Event = new(IERC20Transfer)
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
func (it *IERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Transfer represents a Transfer event raised by the IERC20 contract.
type IERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20TransferIterator{contract: _IERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Transfer)
				if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) ParseTransfer(log types.Log) (*IERC20Transfer, error) {
	event := new(IERC20Transfer)
	if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IMintyTokenABI is the input ABI used to generate the binding from.
const IMintyTokenABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"creator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"artist\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"ipfsHash\",\"type\":\"string\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IMintyTokenFuncSigs maps the 4-byte function signature to its string representation.
var IMintyTokenFuncSigs = map[string]string{
	"510b5158": "creator(uint256)",
	"081812fc": "getApproved(uint256)",
	"e985e9c5": "isApprovedForAll(address,address)",
	"b85cbc79": "mint(address,address,uint256,string)",
	"6352211e": "ownerOf(uint256)",
	"b88d4fde": "safeTransferFrom(address,address,uint256,bytes)",
	"00923f9e": "tokenExists(uint256)",
}

// IMintyToken is an auto generated Go binding around an Ethereum contract.
type IMintyToken struct {
	IMintyTokenCaller     // Read-only binding to the contract
	IMintyTokenTransactor // Write-only binding to the contract
	IMintyTokenFilterer   // Log filterer for contract events
}

// IMintyTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type IMintyTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMintyTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IMintyTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMintyTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IMintyTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMintyTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IMintyTokenSession struct {
	Contract     *IMintyToken      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IMintyTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IMintyTokenCallerSession struct {
	Contract *IMintyTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// IMintyTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IMintyTokenTransactorSession struct {
	Contract     *IMintyTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IMintyTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type IMintyTokenRaw struct {
	Contract *IMintyToken // Generic contract binding to access the raw methods on
}

// IMintyTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IMintyTokenCallerRaw struct {
	Contract *IMintyTokenCaller // Generic read-only contract binding to access the raw methods on
}

// IMintyTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IMintyTokenTransactorRaw struct {
	Contract *IMintyTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIMintyToken creates a new instance of IMintyToken, bound to a specific deployed contract.
func NewIMintyToken(address common.Address, backend bind.ContractBackend) (*IMintyToken, error) {
	contract, err := bindIMintyToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IMintyToken{IMintyTokenCaller: IMintyTokenCaller{contract: contract}, IMintyTokenTransactor: IMintyTokenTransactor{contract: contract}, IMintyTokenFilterer: IMintyTokenFilterer{contract: contract}}, nil
}

// NewIMintyTokenCaller creates a new read-only instance of IMintyToken, bound to a specific deployed contract.
func NewIMintyTokenCaller(address common.Address, caller bind.ContractCaller) (*IMintyTokenCaller, error) {
	contract, err := bindIMintyToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IMintyTokenCaller{contract: contract}, nil
}

// NewIMintyTokenTransactor creates a new write-only instance of IMintyToken, bound to a specific deployed contract.
func NewIMintyTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*IMintyTokenTransactor, error) {
	contract, err := bindIMintyToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IMintyTokenTransactor{contract: contract}, nil
}

// NewIMintyTokenFilterer creates a new log filterer instance of IMintyToken, bound to a specific deployed contract.
func NewIMintyTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*IMintyTokenFilterer, error) {
	contract, err := bindIMintyToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IMintyTokenFilterer{contract: contract}, nil
}

// bindIMintyToken binds a generic wrapper to an already deployed contract.
func bindIMintyToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IMintyTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMintyToken *IMintyTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMintyToken.Contract.IMintyTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMintyToken *IMintyTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMintyToken.Contract.IMintyTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMintyToken *IMintyTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMintyToken.Contract.IMintyTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMintyToken *IMintyTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMintyToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMintyToken *IMintyTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMintyToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMintyToken *IMintyTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMintyToken.Contract.contract.Transact(opts, method, params...)
}

// Creator is a free data retrieval call binding the contract method 0x510b5158.
//
// Solidity: function creator(uint256 tokenId) view returns(address)
func (_IMintyToken *IMintyTokenCaller) Creator(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IMintyToken.contract.Call(opts, &out, "creator", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Creator is a free data retrieval call binding the contract method 0x510b5158.
//
// Solidity: function creator(uint256 tokenId) view returns(address)
func (_IMintyToken *IMintyTokenSession) Creator(tokenId *big.Int) (common.Address, error) {
	return _IMintyToken.Contract.Creator(&_IMintyToken.CallOpts, tokenId)
}

// Creator is a free data retrieval call binding the contract method 0x510b5158.
//
// Solidity: function creator(uint256 tokenId) view returns(address)
func (_IMintyToken *IMintyTokenCallerSession) Creator(tokenId *big.Int) (common.Address, error) {
	return _IMintyToken.Contract.Creator(&_IMintyToken.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_IMintyToken *IMintyTokenCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IMintyToken.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_IMintyToken *IMintyTokenSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _IMintyToken.Contract.GetApproved(&_IMintyToken.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_IMintyToken *IMintyTokenCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _IMintyToken.Contract.GetApproved(&_IMintyToken.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_IMintyToken *IMintyTokenCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _IMintyToken.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_IMintyToken *IMintyTokenSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _IMintyToken.Contract.IsApprovedForAll(&_IMintyToken.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_IMintyToken *IMintyTokenCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _IMintyToken.Contract.IsApprovedForAll(&_IMintyToken.CallOpts, owner, operator)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_IMintyToken *IMintyTokenCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IMintyToken.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_IMintyToken *IMintyTokenSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _IMintyToken.Contract.OwnerOf(&_IMintyToken.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_IMintyToken *IMintyTokenCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _IMintyToken.Contract.OwnerOf(&_IMintyToken.CallOpts, tokenId)
}

// TokenExists is a free data retrieval call binding the contract method 0x00923f9e.
//
// Solidity: function tokenExists(uint256 tokenId) view returns(bool)
func (_IMintyToken *IMintyTokenCaller) TokenExists(opts *bind.CallOpts, tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _IMintyToken.contract.Call(opts, &out, "tokenExists", tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TokenExists is a free data retrieval call binding the contract method 0x00923f9e.
//
// Solidity: function tokenExists(uint256 tokenId) view returns(bool)
func (_IMintyToken *IMintyTokenSession) TokenExists(tokenId *big.Int) (bool, error) {
	return _IMintyToken.Contract.TokenExists(&_IMintyToken.CallOpts, tokenId)
}

// TokenExists is a free data retrieval call binding the contract method 0x00923f9e.
//
// Solidity: function tokenExists(uint256 tokenId) view returns(bool)
func (_IMintyToken *IMintyTokenCallerSession) TokenExists(tokenId *big.Int) (bool, error) {
	return _IMintyToken.Contract.TokenExists(&_IMintyToken.CallOpts, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0xb85cbc79.
//
// Solidity: function mint(address buyer, address artist, uint256 tokenId, string ipfsHash) returns()
func (_IMintyToken *IMintyTokenTransactor) Mint(opts *bind.TransactOpts, buyer common.Address, artist common.Address, tokenId *big.Int, ipfsHash string) (*types.Transaction, error) {
	return _IMintyToken.contract.Transact(opts, "mint", buyer, artist, tokenId, ipfsHash)
}

// Mint is a paid mutator transaction binding the contract method 0xb85cbc79.
//
// Solidity: function mint(address buyer, address artist, uint256 tokenId, string ipfsHash) returns()
func (_IMintyToken *IMintyTokenSession) Mint(buyer common.Address, artist common.Address, tokenId *big.Int, ipfsHash string) (*types.Transaction, error) {
	return _IMintyToken.Contract.Mint(&_IMintyToken.TransactOpts, buyer, artist, tokenId, ipfsHash)
}

// Mint is a paid mutator transaction binding the contract method 0xb85cbc79.
//
// Solidity: function mint(address buyer, address artist, uint256 tokenId, string ipfsHash) returns()
func (_IMintyToken *IMintyTokenTransactorSession) Mint(buyer common.Address, artist common.Address, tokenId *big.Int, ipfsHash string) (*types.Transaction, error) {
	return _IMintyToken.Contract.Mint(&_IMintyToken.TransactOpts, buyer, artist, tokenId, ipfsHash)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IMintyToken *IMintyTokenTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IMintyToken.contract.Transact(opts, "safeTransferFrom", from, to, tokenId, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IMintyToken *IMintyTokenSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IMintyToken.Contract.SafeTransferFrom(&_IMintyToken.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IMintyToken *IMintyTokenTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IMintyToken.Contract.SafeTransferFrom(&_IMintyToken.TransactOpts, from, to, tokenId, data)
}

// PMintysaleABI is the input ABI used to generate the binding from.
const PMintysaleABI = "[{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_weth\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_ownerPerMille\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_creatorPerMille\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_divisor\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previous_bid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"this_bid\",\"type\":\"uint256\"}],\"name\":\"BidIncreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"}],\"name\":\"BidReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"hash\",\"type\":\"string\"}],\"name\":\"NewOffer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"OfferAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"Payment\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"ResaleOffer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"SaleResubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"SaleRetracted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ownerShare\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"creatorShare\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"divisor\",\"type\":\"uint256\"}],\"name\":\"SharesUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"}],\"name\":\"acceptBid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"acceptOffer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"available\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"bids\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creatorPerMille\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"divisor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isMinter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"items\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"itemHash\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"available\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"minted\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"topup\",\"type\":\"uint256\"}],\"name\":\"makeBid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minty\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"ipfsString\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"offerNew\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"offerResale\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"ipfsString\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"offerSpecial\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ownerPerMille\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"reSubmitOffer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"retractOffer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"setMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIMintyToken\",\"name\":\"m\",\"type\":\"address\"}],\"name\":\"setMintyUnique\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIMintyToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_ownerPerMille\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_creatorPerMille\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_divisor\",\"type\":\"uint256\"}],\"name\":\"updateShares\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"weth\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"withdrawBid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// PMintysaleFuncSigs maps the 4-byte function signature to its string representation.
var PMintysaleFuncSigs = map[string]string{
	"02068664": "acceptBid(uint256,address)",
	"c815729d": "acceptOffer(uint256)",
	"96e494e8": "available(uint256)",
	"3f1ffcec": "bids(uint256,address)",
	"ccb3eae8": "creatorPerMille()",
	"1f2dc5ef": "divisor()",
	"aa271e1a": "isMinter(address)",
	"bfb231d2": "items(uint256)",
	"adfcdc11": "makeBid(uint256,uint256)",
	"cbfcab1d": "minty()",
	"9499ac54": "nextToken()",
	"e35263a7": "offerNew(uint256,string,uint256)",
	"f21d6afe": "offerResale(uint256,uint256)",
	"4aad4e03": "offerSpecial(uint256,address,string,uint256)",
	"8da5cb5b": "owner()",
	"20540adc": "ownerPerMille()",
	"07c6061c": "reSubmitOffer(uint256,uint256)",
	"c4c6c87d": "retractOffer(uint256)",
	"cf456ae7": "setMinter(address,bool)",
	"b04baaa0": "setMintyUnique(address)",
	"fc0c546a": "token()",
	"cb61b42d": "updateShares(uint256,uint256,uint256)",
	"3fc8cef3": "weth()",
	"0eaaf4c8": "withdrawBid(uint256)",
}

// PMintysaleBin is the compiled bytecode used for deploying new contracts.
var PMintysaleBin = "0x6080604052600080546001600160a01b031916331790553480156200002357600080fd5b50604051620034d6380380620034d6833981810160405260a08110156200004957600080fd5b508051602082015160408301516060840151608090940151929391929091906103e882840114620000ac5760405162461bcd60e51b8152600401808060200182810382526036815260200180620034a06036913960400191505060405180910390fd5b6103e881101562000104576040805162461bcd60e51b815260206004820152601960248201527f64697669736f72206973206c657373207468616e203130303000000000000000604482015290519081900360640190fd5b600280546001600160a01b038088166001600160a01b031992831617909255600a805492871692909116919091179055600783905560088290556009819055604080518481526020810184905280820183905290517ff81b1f9d8b5d36a5a67d13006a8b091a9b63e2a38b331d315d99c6daefd2ddcc9181900360600190a1505050505061330880620001986000396000f3fe608060405234801561001057600080fd5b50600436106101585760003560e01c8063adfcdc11116100c3578063cbfcab1d1161007c578063cbfcab1d146104cb578063ccb3eae8146104d3578063cf456ae7146104db578063e35263a714610509578063f21d6afe146105b8578063fc0c546a146105db57610158565b8063adfcdc1114610361578063b04baaa014610384578063bfb231d2146103aa578063c4c6c87d14610468578063c815729d14610485578063cb61b42d146104a257610158565b80633fc8cef3116101155780633fc8cef3146102195780634aad4e031461023d5780638da5cb5b146102fa5780639499ac541461030257806396e494e81461030a578063aa271e1a1461033b57610158565b8063020686641461015d57806307c6061c1461018b5780630eaaf4c8146101ae5780631f2dc5ef146101cb57806320540adc146101e55780633f1ffcec146101ed575b600080fd5b6101896004803603604081101561017357600080fd5b50803590602001356001600160a01b03166105e3565b005b610189600480360360408110156101a157600080fd5b5080359060200135610c2c565b610189600480360360208110156101c457600080fd5b5035610f3e565b6101d36110e6565b60408051918252519081900360200190f35b6101d36110ec565b6101d36004803603604081101561020357600080fd5b50803590602001356001600160a01b03166110f2565b61022161110f565b604080516001600160a01b039092168252519081900360200190f35b6101896004803603608081101561025357600080fd5b8135916001600160a01b036020820135169181019060608101604082013564010000000081111561028357600080fd5b82018360208201111561029557600080fd5b803590602001918460018302840111640100000000831117156102b757600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550509135925061111e915050565b61022161139c565b6101d36113ab565b6103276004803603602081101561032057600080fd5b50356113b1565b604080519115158252519081900360200190f35b6103276004803603602081101561035157600080fd5b50356001600160a01b0316611687565b6101896004803603604081101561037757600080fd5b508035906020013561169c565b6101896004803603602081101561039a57600080fd5b50356001600160a01b03166119af565b6103c7600480360360208110156103c057600080fd5b5035611a1f565b60405180866001600160a01b031681526020018060200185815260200184151581526020018315158152602001828103825286818151815260200191508051906020019080838360005b83811015610429578181015183820152602001610411565b50505050905090810190601f1680156104565780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390f35b6101896004803603602081101561047e57600080fd5b5035611aec565b6101896004803603602081101561049b57600080fd5b5035611dee565b610189600480360360608110156104b857600080fd5b5080359060208101359060400135611ed2565b61022161200f565b6101d361201e565b610189600480360360408110156104f157600080fd5b506001600160a01b0381351690602001351515612024565b6101896004803603606081101561051f57600080fd5b8135919081019060408101602082013564010000000081111561054157600080fd5b82018360208201111561055357600080fd5b8035906020019184600183028401116401000000008311171561057557600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550509135925061209d915050565b610189600480360360408110156105ce57600080fd5b5080359060200135612407565b6102216127d3565b60065460609060ff1615610635576040805162461bcd60e51b81526020600482015260146024820152734e6f207265656e7472616e637920706c6561736560601b604482015290519081900360640190fd5b6006805460ff1916600117905561064a61311e565b600084815260046020908152604091829020825160a08101845281546001600160a01b03168152600180830180548651600261010094831615949094026000190190911692909204601f81018690048602830186019096528582529194929385810193919291908301828280156107025780601f106106d757610100808354040283529160200191610702565b820191906000526020600020905b8154815290600101906020018083116106e557829003601f168201915b5050509183525050600282015460208083019190915260039092015460ff8082161515604080850191909152610100909204161515606090920191909152825160015482516331a9108f60e11b8152600481018a9052925194955090936000936001600160a01b0390921692636352211e926024808301939192829003018186803b15801561079057600080fd5b505afa1580156107a4573d6000803e3d6000fd5b505050506040513d60208110156107ba57600080fd5b5051606084015190915061080a576040805162461bcd60e51b81526020600482015260126024820152714974656d206e6f7420617661696c61626c6560701b604482015290519081900360640190fd5b816001600160a01b0316816001600160a01b03161461086c576040805162461bcd60e51b815260206004820152601960248201527824ba32b6903737ba1037bbb732b210313c9037b33332b932b960391b604482015290519081900360640190fd5b336001600160a01b038316146108c1576040805162461bcd60e51b8152602060048201526015602482015274139bdd081e5bdd5c881a5d195b481d1bc81cd95b1b605a1b604482015290519081900360640190fd5b6000868152600b602090815260408083206001600160a01b0389168452909152812080549190556080840151156109df57600154604051635c46a7ef60e11b81526001600160a01b03858116600483019081528982166024840152604483018b9052608060648401908152895160848501528951929094169363b88d4fde9388938c938e938d939192909160a40190602085019080838360005b8381101561097357818101518382015260200161095b565b50505050905090810190601f1680156109a05780820380516001836020036101000a031916815260200191505b5095505050505050600060405180830381600087803b1580156109c257600080fd5b505af11580156109d6573d6000803e3d6000fd5b50505050610aee565b600160009054906101000a90046001600160a01b03166001600160a01b031663b85cbc798786600001518a88602001516040518563ffffffff1660e01b815260040180856001600160a01b03168152602001846001600160a01b0316815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610a7f578181015183820152602001610a67565b50505050905090810190601f168015610aac5780820380516001836020036101000a031916815260200191505b5095505050505050600060405180830381600087803b158015610ace57600080fd5b505af1158015610ae2573d6000803e3d6000fd5b50506001608087015250505b600060608501819052878152600460209081526040909120855181546001600160a01b0319166001600160a01b039091161781558186015180518793610b3b92600185019291019061315a565b5060408281015160028301556060808401516003909301805460809095015115156101000261ff001994151560ff19909616959095179390931693909317909155600a54865182516001600160a01b03928316815290821660208201529086168183015290517f8703df07b43faff82a32cfdce460e22eb4b029f946f61df422e57d5ef66f3962929181900390910190a18351610bd99084836127e2565b6006805460ff19169055604080513381526020810189905280820183905290517f2e91b12fed89e8218510d4f7c38c6db746acd0984f09b436b69bd7db5e93da8b9181900360600190a150505050505050565b610c3461311e565b600083815260046020908152604091829020825160a08101845281546001600160a01b03168152600180830180548651600261010094831615949094026000190190911692909204601f8101869004860283018601909652858252919492938581019391929190830182828015610cec5780601f10610cc157610100808354040283529160200191610cec565b820191906000526020600020905b815481529060010190602001808311610ccf57829003601f168201915b5050509183525050600282015460208083019190915260039092015460ff80821615156040808501919091526101009092041615156060909201919091528251600154825162491fcf60e11b815260048101899052925194955090936001600160a01b039091169262923f9e9260248082019391829003018186803b158015610d7457600080fd5b505afa158015610d88573d6000803e3d6000fd5b505050506040513d6020811015610d9e57600080fd5b505115610e2057600154604080516331a9108f60e11b81526004810187905290516001600160a01b0390921691636352211e91602480820192602092909190829003018186803b158015610df157600080fd5b505afa158015610e05573d6000803e3d6000fd5b505050506040513d6020811015610e1b57600080fd5b505190505b6001600160a01b0381163314610e6c576040805162461bcd60e51b815260206004820152600c60248201526b155b985d5d1a1bdc9a5cd95960a21b604482015290519081900360640190fd5b6001606083018190526040808401859052600086815260046020908152919020845181546001600160a01b0319166001600160a01b0390911617815581850151805186949293610ec093850192019061315a565b50604082810151600283015560608301516003909201805460809094015115156101000261ff001993151560ff19909516949094179290921692909217905580518581526020810185905281517fe36efdb123434f0d9dcc119c468c16b0d6de5fa64252b13e049b2c87dafc2a1e929181900390910190a150505050565b60065460ff1615610f8d576040805162461bcd60e51b81526020600482015260146024820152734e6f207265656e7472616e637920706c6561736560601b604482015290519081900360640190fd5b6006805460ff191660011790556000818152600b6020908152604080832033845290915290205480610ffc576040805162461bcd60e51b81526020600482015260136024820152726e6f7468696e6720746f20776974686472617760681b604482015290519081900360640190fd5b6000828152600b6020908152604080832033808552908352818420849055600254825163a9059cbb60e01b815260048101929092526024820186905291516001600160a01b039092169363a9059cbb9360448084019491939192918390030190829087803b15801561106d57600080fd5b505af1158015611081573d6000803e3d6000fd5b505050506040513d602081101561109757600080fd5b50516110d8576040805162461bcd60e51b81526020600482015260156024820152600080516020613255833981519152604482015290519081900360640190fd5b50506006805460ff19169055565b60095481565b60075481565b600b60209081526000928352604080842090915290825290205481565b6002546001600160a01b031681565b6000546001600160a01b0316331461116c576040805162461bcd60e51b815260206004820152600c60248201526b1d5b985d5d1a1bdc9a5cd95960a21b604482015290519081900360640190fd5b6001546040805162491fcf60e11b81526004810187905290516001600160a01b039092169162923f9e91602480820192602092909190829003018186803b1580156111b657600080fd5b505afa1580156111ca573d6000803e3d6000fd5b505050506040513d60208110156111e057600080fd5b505115611227576040805162461bcd60e51b815260206004820152601060248201526f125b9d985b1a59081d1bdad95b88125160821b604482015290519081900360640190fd5b6040805160a0810182526001600160a01b03858116825260208083018681528385018690526001606085018190526000608086018190528a81526004845295909520845181546001600160a01b03191694169390931783555180519394929361129793850192919091019061315a565b5060408281015160028301556060808401516003909301805460809586015115156101000261ff001995151560ff1990921691909117949094169390931790925580518781526001600160a01b03871660208083019190915291810185905291820183815285519383019390935284517f3697349224444578adf7f32ac71b635cf39a0924633e5700169ed207bf91f7b49388938893879389939160a08401919085019080838360005b83811015611359578181015183820152602001611341565b50505050905090810190601f1680156113865780820380516001836020036101000a031916815260200191505b509550505050505060405180910390a150505050565b6000546001600160a01b031681565b60035481565b60006113bb61311e565b600083815260046020908152604091829020825160a08101845281546001600160a01b03168152600180830180548651600261010094831615949094026000190190911692909204601f81018690048602830186019096528582529194929385810193919291908301828280156114735780601f1061144857610100808354040283529160200191611473565b820191906000526020600020905b81548152906001019060200180831161145657829003601f168201915b50505091835250506002820154602082015260039091015460ff808216151560408401526101009091041615156060918201528101519091506114ba576000915050611682565b80608001516114cd576001915050611682565b6000838152600460208181526040928390205460015484516331a9108f60e11b815293840188905293516001600160a01b039182169490911692636352211e926024808301939192829003018186803b15801561152957600080fd5b505afa15801561153d573d6000803e3d6000fd5b505050506040513d602081101561155357600080fd5b50516001600160a01b03161461156d576000915050611682565b60015481516040805163e985e9c560e01b81526001600160a01b0392831660048201523060248201529051919092169163e985e9c5916044808301926020929190829003018186803b1580156115c257600080fd5b505afa1580156115d6573d6000803e3d6000fd5b505050506040513d60208110156115ec57600080fd5b5051156115fd576001915050611682565b6001546040805163020604bf60e21b815260048101869052905130926001600160a01b03169163081812fc916024808301926020929190829003018186803b15801561164857600080fd5b505afa15801561165c573d6000803e3d6000fd5b505050506040513d602081101561167257600080fd5b50516001600160a01b0316149150505b919050565b60056020526000908152604090205460ff1681565b600254604080516323b872dd60e01b81523360048201523060248201526044810184905290516001600160a01b03909216916323b872dd916064808201926020929091908290030181600087803b1580156116f657600080fd5b505af115801561170a573d6000803e3d6000fd5b505050506040513d602081101561172057600080fd5b5051611761576040805162461bcd60e51b81526020600482015260156024820152600080516020613255833981519152604482015290519081900360640190fd5b61176961311e565b600083815260046020908152604091829020825160a08101845281546001600160a01b03168152600180830180548651600261010094831615949094026000190190911692909204601f81018690048602830186019096528582529194929385810193919291908301828280156118215780601f106117f657610100808354040283529160200191611821565b820191906000526020600020905b81548152906001019060200180831161180457829003601f168201915b50505091835250506002820154602082015260039091015460ff808216151560408401526101009091041615156060918201528101519091506118a0576040805162461bcd60e51b81526020600482015260126024820152714974656d206e6f7420617661696c61626c6560701b604482015290519081900360640190fd5b6000838152600b60209081526040808320338452909152908190205490820151908301908111156118f6576000848152600b602090815260408083203384529091528120556118ef8482612b38565b50506119ab565b6000848152600b6020908152604080832033845290915290208190558281141561195f57604080513381526020810186905280820183905290517f7f06de72678b8f755d2b9cce42fb004f364fad13cafec8e3b43e8808d9db828c9181900360600190a16119a8565b6040805133815260208101869052848303818301526060810185905290517f34762811225e202a6ffaa6152f171af086267130932de1c3951c1e36ff89b28a9181900360800190a15b50505b5050565b6000546001600160a01b031633146119fd576040805162461bcd60e51b815260206004820152600c60248201526b1d5b985d5d1a1bdc9a5cd95960a21b604482015290519081900360640190fd5b600180546001600160a01b0319166001600160a01b0392909216919091179055565b6004602090815260009182526040918290208054600180830180548651600261010094831615949094026000190190911692909204601f81018690048602830186019096528582526001600160a01b03909216949293909290830182828015611ac95780601f10611a9e57610100808354040283529160200191611ac9565b820191906000526020600020905b815481529060010190602001808311611aac57829003601f168201915b50505050600283015460039093015491929160ff80821692506101009091041685565b611af461311e565b600082815260046020908152604091829020825160a08101845281546001600160a01b03168152600180830180548651600261010094831615949094026000190190911692909204601f8101869004860283018601909652858252919492938581019391929190830182828015611bac5780601f10611b8157610100808354040283529160200191611bac565b820191906000526020600020905b815481529060010190602001808311611b8f57829003601f168201915b5050509183525050600282015460208083019190915260039092015460ff80821615156040808501919091526101009092041615156060909201919091528251600154825162491fcf60e11b815260048101889052925194955090936001600160a01b039091169262923f9e9260248082019391829003018186803b158015611c3457600080fd5b505afa158015611c48573d6000803e3d6000fd5b505050506040513d6020811015611c5e57600080fd5b505115611ce057600154604080516331a9108f60e11b81526004810186905290516001600160a01b0390921691636352211e91602480820192602092909190829003018186803b158015611cb157600080fd5b505afa158015611cc5573d6000803e3d6000fd5b505050506040513d6020811015611cdb57600080fd5b505190505b6001600160a01b0381163314611d2c576040805162461bcd60e51b815260206004820152600c60248201526b155b985d5d1a1bdc9a5cd95960a21b604482015290519081900360640190fd5b600060608301819052838152600460209081526040909120835181546001600160a01b0319166001600160a01b039091161781558184015180518593611d7992600185019291019061315a565b50604082810151600283015560608301516003909201805460809094015115156101000261ff001993151560ff199095169490941792909216929092179055805184815290517f567b7c6ccef395faa0faf04e2da702ec06c19c0cff01cde647c3d285257a0bb09181900360200190a1505050565b6000818152600460208181526040808420600290810154905482516323b872dd60e01b8152339581019590955230602486015260448501829052915190946001600160a01b03909216936323b872dd936064808301949193928390030190829087803b158015611e5d57600080fd5b505af1158015611e71573d6000803e3d6000fd5b505050506040513d6020811015611e8757600080fd5b5051611ec8576040805162461bcd60e51b81526020600482015260156024820152600080516020613255833981519152604482015290519081900360640190fd5b6119ab8282612b38565b6000546001600160a01b03163314611f20576040805162461bcd60e51b815260206004820152600c60248201526b1d5b985d5d1a1bdc9a5cd95960a21b604482015290519081900360640190fd5b8282016103e814611f625760405162461bcd60e51b81526004018080602001828103825260368152602001806131fc6036913960400191505060405180910390fd5b6103e86009541015611fbb576040805162461bcd60e51b815260206004820152601960248201527f64697669736f72206973206c657373207468616e203130303000000000000000604482015290519081900360640190fd5b600783905560088290556009819055604080518481526020810184905280820183905290517ff81b1f9d8b5d36a5a67d13006a8b091a9b63e2a38b331d315d99c6daefd2ddcc9181900360600190a1505050565b600a546001600160a01b031681565b60085481565b6000546001600160a01b03163314612072576040805162461bcd60e51b815260206004820152600c60248201526b1d5b985d5d1a1bdc9a5cd95960a21b604482015290519081900360640190fd5b6001600160a01b03919091166000908152600560205260409020805460ff1916911515919091179055565b6001546040805162491fcf60e11b81526004810186905290516001600160a01b039092169162923f9e91602480820192602092909190829003018186803b1580156120e757600080fd5b505afa1580156120fb573d6000803e3d6000fd5b505050506040513d602081101561211157600080fd5b505115612158576040805162461bcd60e51b815260206004820152601060248201526f125b9d985b1a59081d1bdad95b88125160821b604482015290519081900360640190fd5b3360009081526005602052604090205460ff166121a65760405162461bcd60e51b81526004018080602001828103825260278152602001806132756027913960400191505060405180910390fd5b6121ae61311e565b600084815260046020908152604091829020825160a08101845281546001600160a01b03168152600180830180548651600261010094831615949094026000190190911692909204601f81018690048602830186019096528582529194929385810193919291908301828280156122665780601f1061223b57610100808354040283529160200191612266565b820191906000526020600020905b81548152906001019060200180831161224957829003601f168201915b50505091835250506002820154602082015260039091015460ff8082161515604084015261010090910416151560609091015280519091506001600160a01b0316156122e35760405162461bcd60e51b81526004018080602001828103825260238152602001806132326023913960400191505060405180910390fd5b6040805160a08101825233815260208082018681528284018690526001606084018190526000608085018190528981526004845294909420835181546001600160a01b0319166001600160a01b0390911617815590518051939491936123519392850192919091019061315a565b5060408281015160028301556060808401516003909301805460809586015115156101000261ff001995151560ff19909216919091179490941693909317909255805187815233602080830182905292820187905292810184815287519482019490945286517f3697349224444578adf7f32ac71b635cf39a0924633e5700169ed207bf91f7b49489949388938a9390929160a08401919085019080838360008315611359578181015183820152602001611341565b6001546040805162491fcf60e11b81526004810185905290516001600160a01b039092169162923f9e91602480820192602092909190829003018186803b15801561245157600080fd5b505afa158015612465573d6000803e3d6000fd5b505050506040513d602081101561247b57600080fd5b50516124c6576040805162461bcd60e51b8152602060048201526015602482015274546f6b656e20646f6573206e6f742065786973742160581b604482015290519081900360640190fd5b600154604080516331a9108f60e11b81526004810185905290516001600160a01b0390921691636352211e91602480820192602092909190829003018186803b15801561251257600080fd5b505afa158015612526573d6000803e3d6000fd5b505050506040513d602081101561253c57600080fd5b50516001600160a01b0316331461259a576040805162461bcd60e51b815260206004820152601960248201527f596f7520646f206e6f74206f776e207468697320746f6b656e00000000000000604482015290519081900360640190fd5b6001546040805163e985e9c560e01b815233600482015230602482015290516001600160a01b039092169163e985e9c591604480820192602092909190829003018186803b1580156125eb57600080fd5b505afa1580156125ff573d6000803e3d6000fd5b505050506040513d602081101561261557600080fd5b50518061269f57506001546040805163020604bf60e21b815260048101859052905130926001600160a01b03169163081812fc916024808301926020929190829003018186803b15801561266857600080fd5b505afa15801561267c573d6000803e3d6000fd5b505050506040513d602081101561269257600080fd5b50516001600160a01b0316145b6126da5760405162461bcd60e51b815260040180806020018281038252603781526020018061329c6037913960400191505060405180910390fd5b6040805160a08101825233815281516020808201845260008083528184019283528385018690526001606085018190526080850181905287825260048352949020835181546001600160a01b0319166001600160a01b03909116178155915180519394929361275093850192919091019061315a565b5060408281015160028301556060808401516003909301805460809095015115156101000261ff001994151560ff19909616959095179390931693909317909155805184815233602082015280820184905290517febe4318722b948bede41b4decd121e3e2ab187f11ad89499eb7a586de7ffcd90929181900390910190a15050565b6001546001600160a01b031681565b60006009546008548302816127f357fe5b049050600060095460075484028161280757fe5b04905081810183036001600160a01b0386811690861614156128ea576002546040805163a9059cbb60e01b81526001600160a01b03898116600483015285870160248301529151919092169163a9059cbb9160448083019260209291908290030181600087803b15801561287a57600080fd5b505af115801561288e573d6000803e3d6000fd5b505050506040513d60208110156128a457600080fd5b50516128e5576040805162461bcd60e51b81526020600482015260156024820152600080516020613255833981519152604482015290519081900360640190fd5b612a6c565b6002546040805163a9059cbb60e01b81526001600160a01b038981166004830152602482018790529151919092169163a9059cbb9160448083019260209291908290030181600087803b15801561294057600080fd5b505af1158015612954573d6000803e3d6000fd5b505050506040513d602081101561296a57600080fd5b50516129ab576040805162461bcd60e51b81526020600482015260156024820152600080516020613255833981519152604482015290519081900360640190fd5b6002546040805163a9059cbb60e01b81526001600160a01b038881166004830152602482018690529151919092169163a9059cbb9160448083019260209291908290030181600087803b158015612a0157600080fd5b505af1158015612a15573d6000803e3d6000fd5b505050506040513d6020811015612a2b57600080fd5b5051612a6c576040805162461bcd60e51b81526020600482015260156024820152600080516020613255833981519152604482015290519081900360640190fd5b600254600a546040805163a9059cbb60e01b81526001600160a01b039283166004820152602481018590529051919092169163a9059cbb9160448083019260209291908290030181600087803b158015612ac557600080fd5b505af1158015612ad9573d6000803e3d6000fd5b505050506040513d6020811015612aef57600080fd5b5051612b30576040805162461bcd60e51b81526020600482015260156024820152600080516020613255833981519152604482015290519081900360640190fd5b505050505050565b60065460ff1615612b87576040805162461bcd60e51b81526020600482015260146024820152734e6f207265656e7472616e637920706c6561736560601b604482015290519081900360640190fd5b6006805460ff191660011790556060612b9e61311e565b600084815260046020908152604091829020825160a08101845281546001600160a01b03168152600180830180548651600261010094831615949094026000190190911692909204601f8101869004860283018601909652858252919492938581019391929190830182828015612c565780601f10612c2b57610100808354040283529160200191612c56565b820191906000526020600020905b815481529060010190602001808311612c3957829003601f168201915b50505091835250506002820154602082015260039091015460ff8082161515604084015261010090910416151560609182015281519082015191925090612cd9576040805162461bcd60e51b81526020600482015260126024820152714974656d206e6f7420617661696c61626c6560701b604482015290519081900360640190fd5b8160400151841015612d22576040805162461bcd60e51b815260206004820152600d60248201526c141c9a58d9481b9bdd081b595d609a1b604482015290519081900360640190fd5b816080015115612eee57600154604080516331a9108f60e11b81526004810188905290516000926001600160a01b031691636352211e916024808301926020929190829003018186803b158015612d7857600080fd5b505afa158015612d8c573d6000803e3d6000fd5b505050506040513d6020811015612da257600080fd5b505190506001600160a01b0380821690831614612e02576040805162461bcd60e51b815260206004820152601960248201527824ba32b6903737ba1037bbb732b210313c9037b33332b932b960391b604482015290519081900360640190fd5b600154604051635c46a7ef60e11b81526001600160a01b03848116600483019081523360248401819052604484018b9052608060648501908152895160848601528951939095169463b88d4fde94889492938d938c93919260a40190602085019080838360005b83811015612e81578181015183820152602001612e69565b50505050905090810190601f168015612eae5780820380516001836020036101000a031916815260200191505b5095505050505050600060405180830381600087803b158015612ed057600080fd5b505af1158015612ee4573d6000803e3d6000fd5b5050505050612fe2565b600154825160208085015160405163b85cbc7960e01b815233600482018181526001600160a01b038681166024850152604484018d9052608060648501908152855160848601528551919098169763b85cbc7997939693958e959094919260a49091019185019080838360005b83811015612f73578181015183820152602001612f5b565b50505050905090810190601f168015612fa05780820380516001836020036101000a031916815260200191505b5095505050505050600060405180830381600087803b158015612fc257600080fd5b505af1158015612fd6573d6000803e3d6000fd5b50506001608085015250505b600060608301819052858152600460209081526040909120835181546001600160a01b0319166001600160a01b03909116178155818401518051859361302f92600185019291019061315a565b5060408281015160028301556060808401516003909301805460809095015115156101000261ff001994151560ff19909616959095179390931693909317909155600a54845182516001600160a01b03928316815290821660208201529084168183015290517f8703df07b43faff82a32cfdce460e22eb4b029f946f61df422e57d5ef66f3962929181900390910190a181516130cd9082866127e2565b6006805460ff19169055604080513381526020810187905280820186905290517f2e91b12fed89e8218510d4f7c38c6db746acd0984f09b436b69bd7db5e93da8b9181900360600190a15050505050565b6040518060a0016040528060006001600160a01b0316815260200160608152602001600081526020016000151581526020016000151581525090565b828054600181600116156101000203166002900490600052602060002090601f01602090048101928261319057600085556131d6565b82601f106131a957805160ff19168380011785556131d6565b828001600101855582156131d6579182015b828111156131d65782518255916020019190600101906131bb565b506131e29291506131e6565b5090565b5b808211156131e257600081556001016131e756fe73756d285f63726561746f725065724d696c6c65202b205f6f776e65725065724d696c6c6529206d75737420657175616c2031303030417474656d707420746f206d6f6469667920616e206578697374696e67206f6666657243616e6e6f74207472616e736665722066756e64730000000000000000000000596f7520617265206e6f7420616c6c6f77656420746f206d696e7420746f6b656e732068657265596f752068617665206e6f7420617070726f766564207468697320636f6e747261637420746f2073656c6c20796f757220746f6b656e73a2646970667358221220c4668b3bfb6621cff5fcb0f004496ec01c23132f4c50c66dcc82f7ab1128cf1e64736f6c6343000705003373756d285f63726561746f725065724d696c6c65202b205f6f776e65725065724d696c6c6529206d75737420657175616c2031303030"

// DeployPMintysale deploys a new Ethereum contract, binding an instance of PMintysale to it.
func DeployPMintysale(auth *bind.TransactOpts, backend bind.ContractBackend, _weth common.Address, wallet common.Address, _ownerPerMille *big.Int, _creatorPerMille *big.Int, _divisor *big.Int) (common.Address, *types.Transaction, *PMintysale, error) {
	parsed, err := abi.JSON(strings.NewReader(PMintysaleABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PMintysaleBin), backend, _weth, wallet, _ownerPerMille, _creatorPerMille, _divisor)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PMintysale{PMintysaleCaller: PMintysaleCaller{contract: contract}, PMintysaleTransactor: PMintysaleTransactor{contract: contract}, PMintysaleFilterer: PMintysaleFilterer{contract: contract}}, nil
}

// PMintysale is an auto generated Go binding around an Ethereum contract.
type PMintysale struct {
	PMintysaleCaller     // Read-only binding to the contract
	PMintysaleTransactor // Write-only binding to the contract
	PMintysaleFilterer   // Log filterer for contract events
}

// PMintysaleCaller is an auto generated read-only Go binding around an Ethereum contract.
type PMintysaleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PMintysaleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PMintysaleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PMintysaleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PMintysaleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PMintysaleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PMintysaleSession struct {
	Contract     *PMintysale       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PMintysaleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PMintysaleCallerSession struct {
	Contract *PMintysaleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// PMintysaleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PMintysaleTransactorSession struct {
	Contract     *PMintysaleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// PMintysaleRaw is an auto generated low-level Go binding around an Ethereum contract.
type PMintysaleRaw struct {
	Contract *PMintysale // Generic contract binding to access the raw methods on
}

// PMintysaleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PMintysaleCallerRaw struct {
	Contract *PMintysaleCaller // Generic read-only contract binding to access the raw methods on
}

// PMintysaleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PMintysaleTransactorRaw struct {
	Contract *PMintysaleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPMintysale creates a new instance of PMintysale, bound to a specific deployed contract.
func NewPMintysale(address common.Address, backend bind.ContractBackend) (*PMintysale, error) {
	contract, err := bindPMintysale(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PMintysale{PMintysaleCaller: PMintysaleCaller{contract: contract}, PMintysaleTransactor: PMintysaleTransactor{contract: contract}, PMintysaleFilterer: PMintysaleFilterer{contract: contract}}, nil
}

// NewPMintysaleCaller creates a new read-only instance of PMintysale, bound to a specific deployed contract.
func NewPMintysaleCaller(address common.Address, caller bind.ContractCaller) (*PMintysaleCaller, error) {
	contract, err := bindPMintysale(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PMintysaleCaller{contract: contract}, nil
}

// NewPMintysaleTransactor creates a new write-only instance of PMintysale, bound to a specific deployed contract.
func NewPMintysaleTransactor(address common.Address, transactor bind.ContractTransactor) (*PMintysaleTransactor, error) {
	contract, err := bindPMintysale(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PMintysaleTransactor{contract: contract}, nil
}

// NewPMintysaleFilterer creates a new log filterer instance of PMintysale, bound to a specific deployed contract.
func NewPMintysaleFilterer(address common.Address, filterer bind.ContractFilterer) (*PMintysaleFilterer, error) {
	contract, err := bindPMintysale(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PMintysaleFilterer{contract: contract}, nil
}

// bindPMintysale binds a generic wrapper to an already deployed contract.
func bindPMintysale(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PMintysaleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PMintysale *PMintysaleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PMintysale.Contract.PMintysaleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PMintysale *PMintysaleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PMintysale.Contract.PMintysaleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PMintysale *PMintysaleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PMintysale.Contract.PMintysaleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PMintysale *PMintysaleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PMintysale.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PMintysale *PMintysaleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PMintysale.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PMintysale *PMintysaleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PMintysale.Contract.contract.Transact(opts, method, params...)
}

// Available is a free data retrieval call binding the contract method 0x96e494e8.
//
// Solidity: function available(uint256 tokenId) view returns(bool)
func (_PMintysale *PMintysaleCaller) Available(opts *bind.CallOpts, tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _PMintysale.contract.Call(opts, &out, "available", tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Available is a free data retrieval call binding the contract method 0x96e494e8.
//
// Solidity: function available(uint256 tokenId) view returns(bool)
func (_PMintysale *PMintysaleSession) Available(tokenId *big.Int) (bool, error) {
	return _PMintysale.Contract.Available(&_PMintysale.CallOpts, tokenId)
}

// Available is a free data retrieval call binding the contract method 0x96e494e8.
//
// Solidity: function available(uint256 tokenId) view returns(bool)
func (_PMintysale *PMintysaleCallerSession) Available(tokenId *big.Int) (bool, error) {
	return _PMintysale.Contract.Available(&_PMintysale.CallOpts, tokenId)
}

// Bids is a free data retrieval call binding the contract method 0x3f1ffcec.
//
// Solidity: function bids(uint256 , address ) view returns(uint256)
func (_PMintysale *PMintysaleCaller) Bids(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PMintysale.contract.Call(opts, &out, "bids", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Bids is a free data retrieval call binding the contract method 0x3f1ffcec.
//
// Solidity: function bids(uint256 , address ) view returns(uint256)
func (_PMintysale *PMintysaleSession) Bids(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _PMintysale.Contract.Bids(&_PMintysale.CallOpts, arg0, arg1)
}

// Bids is a free data retrieval call binding the contract method 0x3f1ffcec.
//
// Solidity: function bids(uint256 , address ) view returns(uint256)
func (_PMintysale *PMintysaleCallerSession) Bids(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _PMintysale.Contract.Bids(&_PMintysale.CallOpts, arg0, arg1)
}

// CreatorPerMille is a free data retrieval call binding the contract method 0xccb3eae8.
//
// Solidity: function creatorPerMille() view returns(uint256)
func (_PMintysale *PMintysaleCaller) CreatorPerMille(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PMintysale.contract.Call(opts, &out, "creatorPerMille")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CreatorPerMille is a free data retrieval call binding the contract method 0xccb3eae8.
//
// Solidity: function creatorPerMille() view returns(uint256)
func (_PMintysale *PMintysaleSession) CreatorPerMille() (*big.Int, error) {
	return _PMintysale.Contract.CreatorPerMille(&_PMintysale.CallOpts)
}

// CreatorPerMille is a free data retrieval call binding the contract method 0xccb3eae8.
//
// Solidity: function creatorPerMille() view returns(uint256)
func (_PMintysale *PMintysaleCallerSession) CreatorPerMille() (*big.Int, error) {
	return _PMintysale.Contract.CreatorPerMille(&_PMintysale.CallOpts)
}

// Divisor is a free data retrieval call binding the contract method 0x1f2dc5ef.
//
// Solidity: function divisor() view returns(uint256)
func (_PMintysale *PMintysaleCaller) Divisor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PMintysale.contract.Call(opts, &out, "divisor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Divisor is a free data retrieval call binding the contract method 0x1f2dc5ef.
//
// Solidity: function divisor() view returns(uint256)
func (_PMintysale *PMintysaleSession) Divisor() (*big.Int, error) {
	return _PMintysale.Contract.Divisor(&_PMintysale.CallOpts)
}

// Divisor is a free data retrieval call binding the contract method 0x1f2dc5ef.
//
// Solidity: function divisor() view returns(uint256)
func (_PMintysale *PMintysaleCallerSession) Divisor() (*big.Int, error) {
	return _PMintysale.Contract.Divisor(&_PMintysale.CallOpts)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address ) view returns(bool)
func (_PMintysale *PMintysaleCaller) IsMinter(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _PMintysale.contract.Call(opts, &out, "isMinter", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address ) view returns(bool)
func (_PMintysale *PMintysaleSession) IsMinter(arg0 common.Address) (bool, error) {
	return _PMintysale.Contract.IsMinter(&_PMintysale.CallOpts, arg0)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address ) view returns(bool)
func (_PMintysale *PMintysaleCallerSession) IsMinter(arg0 common.Address) (bool, error) {
	return _PMintysale.Contract.IsMinter(&_PMintysale.CallOpts, arg0)
}

// Items is a free data retrieval call binding the contract method 0xbfb231d2.
//
// Solidity: function items(uint256 ) view returns(address creator, string itemHash, uint256 price, bool available, bool minted)
func (_PMintysale *PMintysaleCaller) Items(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Creator   common.Address
	ItemHash  string
	Price     *big.Int
	Available bool
	Minted    bool
}, error) {
	var out []interface{}
	err := _PMintysale.contract.Call(opts, &out, "items", arg0)

	outstruct := new(struct {
		Creator   common.Address
		ItemHash  string
		Price     *big.Int
		Available bool
		Minted    bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Creator = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.ItemHash = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Price = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Available = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.Minted = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// Items is a free data retrieval call binding the contract method 0xbfb231d2.
//
// Solidity: function items(uint256 ) view returns(address creator, string itemHash, uint256 price, bool available, bool minted)
func (_PMintysale *PMintysaleSession) Items(arg0 *big.Int) (struct {
	Creator   common.Address
	ItemHash  string
	Price     *big.Int
	Available bool
	Minted    bool
}, error) {
	return _PMintysale.Contract.Items(&_PMintysale.CallOpts, arg0)
}

// Items is a free data retrieval call binding the contract method 0xbfb231d2.
//
// Solidity: function items(uint256 ) view returns(address creator, string itemHash, uint256 price, bool available, bool minted)
func (_PMintysale *PMintysaleCallerSession) Items(arg0 *big.Int) (struct {
	Creator   common.Address
	ItemHash  string
	Price     *big.Int
	Available bool
	Minted    bool
}, error) {
	return _PMintysale.Contract.Items(&_PMintysale.CallOpts, arg0)
}

// Minty is a free data retrieval call binding the contract method 0xcbfcab1d.
//
// Solidity: function minty() view returns(address)
func (_PMintysale *PMintysaleCaller) Minty(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PMintysale.contract.Call(opts, &out, "minty")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Minty is a free data retrieval call binding the contract method 0xcbfcab1d.
//
// Solidity: function minty() view returns(address)
func (_PMintysale *PMintysaleSession) Minty() (common.Address, error) {
	return _PMintysale.Contract.Minty(&_PMintysale.CallOpts)
}

// Minty is a free data retrieval call binding the contract method 0xcbfcab1d.
//
// Solidity: function minty() view returns(address)
func (_PMintysale *PMintysaleCallerSession) Minty() (common.Address, error) {
	return _PMintysale.Contract.Minty(&_PMintysale.CallOpts)
}

// NextToken is a free data retrieval call binding the contract method 0x9499ac54.
//
// Solidity: function nextToken() view returns(uint256)
func (_PMintysale *PMintysaleCaller) NextToken(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PMintysale.contract.Call(opts, &out, "nextToken")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextToken is a free data retrieval call binding the contract method 0x9499ac54.
//
// Solidity: function nextToken() view returns(uint256)
func (_PMintysale *PMintysaleSession) NextToken() (*big.Int, error) {
	return _PMintysale.Contract.NextToken(&_PMintysale.CallOpts)
}

// NextToken is a free data retrieval call binding the contract method 0x9499ac54.
//
// Solidity: function nextToken() view returns(uint256)
func (_PMintysale *PMintysaleCallerSession) NextToken() (*big.Int, error) {
	return _PMintysale.Contract.NextToken(&_PMintysale.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PMintysale *PMintysaleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PMintysale.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PMintysale *PMintysaleSession) Owner() (common.Address, error) {
	return _PMintysale.Contract.Owner(&_PMintysale.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PMintysale *PMintysaleCallerSession) Owner() (common.Address, error) {
	return _PMintysale.Contract.Owner(&_PMintysale.CallOpts)
}

// OwnerPerMille is a free data retrieval call binding the contract method 0x20540adc.
//
// Solidity: function ownerPerMille() view returns(uint256)
func (_PMintysale *PMintysaleCaller) OwnerPerMille(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PMintysale.contract.Call(opts, &out, "ownerPerMille")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OwnerPerMille is a free data retrieval call binding the contract method 0x20540adc.
//
// Solidity: function ownerPerMille() view returns(uint256)
func (_PMintysale *PMintysaleSession) OwnerPerMille() (*big.Int, error) {
	return _PMintysale.Contract.OwnerPerMille(&_PMintysale.CallOpts)
}

// OwnerPerMille is a free data retrieval call binding the contract method 0x20540adc.
//
// Solidity: function ownerPerMille() view returns(uint256)
func (_PMintysale *PMintysaleCallerSession) OwnerPerMille() (*big.Int, error) {
	return _PMintysale.Contract.OwnerPerMille(&_PMintysale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_PMintysale *PMintysaleCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PMintysale.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_PMintysale *PMintysaleSession) Token() (common.Address, error) {
	return _PMintysale.Contract.Token(&_PMintysale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_PMintysale *PMintysaleCallerSession) Token() (common.Address, error) {
	return _PMintysale.Contract.Token(&_PMintysale.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_PMintysale *PMintysaleCaller) Weth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PMintysale.contract.Call(opts, &out, "weth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_PMintysale *PMintysaleSession) Weth() (common.Address, error) {
	return _PMintysale.Contract.Weth(&_PMintysale.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_PMintysale *PMintysaleCallerSession) Weth() (common.Address, error) {
	return _PMintysale.Contract.Weth(&_PMintysale.CallOpts)
}

// AcceptBid is a paid mutator transaction binding the contract method 0x02068664.
//
// Solidity: function acceptBid(uint256 tokenId, address bidder) returns()
func (_PMintysale *PMintysaleTransactor) AcceptBid(opts *bind.TransactOpts, tokenId *big.Int, bidder common.Address) (*types.Transaction, error) {
	return _PMintysale.contract.Transact(opts, "acceptBid", tokenId, bidder)
}

// AcceptBid is a paid mutator transaction binding the contract method 0x02068664.
//
// Solidity: function acceptBid(uint256 tokenId, address bidder) returns()
func (_PMintysale *PMintysaleSession) AcceptBid(tokenId *big.Int, bidder common.Address) (*types.Transaction, error) {
	return _PMintysale.Contract.AcceptBid(&_PMintysale.TransactOpts, tokenId, bidder)
}

// AcceptBid is a paid mutator transaction binding the contract method 0x02068664.
//
// Solidity: function acceptBid(uint256 tokenId, address bidder) returns()
func (_PMintysale *PMintysaleTransactorSession) AcceptBid(tokenId *big.Int, bidder common.Address) (*types.Transaction, error) {
	return _PMintysale.Contract.AcceptBid(&_PMintysale.TransactOpts, tokenId, bidder)
}

// AcceptOffer is a paid mutator transaction binding the contract method 0xc815729d.
//
// Solidity: function acceptOffer(uint256 tokenId) returns()
func (_PMintysale *PMintysaleTransactor) AcceptOffer(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _PMintysale.contract.Transact(opts, "acceptOffer", tokenId)
}

// AcceptOffer is a paid mutator transaction binding the contract method 0xc815729d.
//
// Solidity: function acceptOffer(uint256 tokenId) returns()
func (_PMintysale *PMintysaleSession) AcceptOffer(tokenId *big.Int) (*types.Transaction, error) {
	return _PMintysale.Contract.AcceptOffer(&_PMintysale.TransactOpts, tokenId)
}

// AcceptOffer is a paid mutator transaction binding the contract method 0xc815729d.
//
// Solidity: function acceptOffer(uint256 tokenId) returns()
func (_PMintysale *PMintysaleTransactorSession) AcceptOffer(tokenId *big.Int) (*types.Transaction, error) {
	return _PMintysale.Contract.AcceptOffer(&_PMintysale.TransactOpts, tokenId)
}

// MakeBid is a paid mutator transaction binding the contract method 0xadfcdc11.
//
// Solidity: function makeBid(uint256 tokenId, uint256 topup) returns()
func (_PMintysale *PMintysaleTransactor) MakeBid(opts *bind.TransactOpts, tokenId *big.Int, topup *big.Int) (*types.Transaction, error) {
	return _PMintysale.contract.Transact(opts, "makeBid", tokenId, topup)
}

// MakeBid is a paid mutator transaction binding the contract method 0xadfcdc11.
//
// Solidity: function makeBid(uint256 tokenId, uint256 topup) returns()
func (_PMintysale *PMintysaleSession) MakeBid(tokenId *big.Int, topup *big.Int) (*types.Transaction, error) {
	return _PMintysale.Contract.MakeBid(&_PMintysale.TransactOpts, tokenId, topup)
}

// MakeBid is a paid mutator transaction binding the contract method 0xadfcdc11.
//
// Solidity: function makeBid(uint256 tokenId, uint256 topup) returns()
func (_PMintysale *PMintysaleTransactorSession) MakeBid(tokenId *big.Int, topup *big.Int) (*types.Transaction, error) {
	return _PMintysale.Contract.MakeBid(&_PMintysale.TransactOpts, tokenId, topup)
}

// OfferNew is a paid mutator transaction binding the contract method 0xe35263a7.
//
// Solidity: function offerNew(uint256 tokenId, string ipfsString, uint256 price) returns()
func (_PMintysale *PMintysaleTransactor) OfferNew(opts *bind.TransactOpts, tokenId *big.Int, ipfsString string, price *big.Int) (*types.Transaction, error) {
	return _PMintysale.contract.Transact(opts, "offerNew", tokenId, ipfsString, price)
}

// OfferNew is a paid mutator transaction binding the contract method 0xe35263a7.
//
// Solidity: function offerNew(uint256 tokenId, string ipfsString, uint256 price) returns()
func (_PMintysale *PMintysaleSession) OfferNew(tokenId *big.Int, ipfsString string, price *big.Int) (*types.Transaction, error) {
	return _PMintysale.Contract.OfferNew(&_PMintysale.TransactOpts, tokenId, ipfsString, price)
}

// OfferNew is a paid mutator transaction binding the contract method 0xe35263a7.
//
// Solidity: function offerNew(uint256 tokenId, string ipfsString, uint256 price) returns()
func (_PMintysale *PMintysaleTransactorSession) OfferNew(tokenId *big.Int, ipfsString string, price *big.Int) (*types.Transaction, error) {
	return _PMintysale.Contract.OfferNew(&_PMintysale.TransactOpts, tokenId, ipfsString, price)
}

// OfferResale is a paid mutator transaction binding the contract method 0xf21d6afe.
//
// Solidity: function offerResale(uint256 tokenId, uint256 price) returns()
func (_PMintysale *PMintysaleTransactor) OfferResale(opts *bind.TransactOpts, tokenId *big.Int, price *big.Int) (*types.Transaction, error) {
	return _PMintysale.contract.Transact(opts, "offerResale", tokenId, price)
}

// OfferResale is a paid mutator transaction binding the contract method 0xf21d6afe.
//
// Solidity: function offerResale(uint256 tokenId, uint256 price) returns()
func (_PMintysale *PMintysaleSession) OfferResale(tokenId *big.Int, price *big.Int) (*types.Transaction, error) {
	return _PMintysale.Contract.OfferResale(&_PMintysale.TransactOpts, tokenId, price)
}

// OfferResale is a paid mutator transaction binding the contract method 0xf21d6afe.
//
// Solidity: function offerResale(uint256 tokenId, uint256 price) returns()
func (_PMintysale *PMintysaleTransactorSession) OfferResale(tokenId *big.Int, price *big.Int) (*types.Transaction, error) {
	return _PMintysale.Contract.OfferResale(&_PMintysale.TransactOpts, tokenId, price)
}

// OfferSpecial is a paid mutator transaction binding the contract method 0x4aad4e03.
//
// Solidity: function offerSpecial(uint256 tokenId, address creator, string ipfsString, uint256 price) returns()
func (_PMintysale *PMintysaleTransactor) OfferSpecial(opts *bind.TransactOpts, tokenId *big.Int, creator common.Address, ipfsString string, price *big.Int) (*types.Transaction, error) {
	return _PMintysale.contract.Transact(opts, "offerSpecial", tokenId, creator, ipfsString, price)
}

// OfferSpecial is a paid mutator transaction binding the contract method 0x4aad4e03.
//
// Solidity: function offerSpecial(uint256 tokenId, address creator, string ipfsString, uint256 price) returns()
func (_PMintysale *PMintysaleSession) OfferSpecial(tokenId *big.Int, creator common.Address, ipfsString string, price *big.Int) (*types.Transaction, error) {
	return _PMintysale.Contract.OfferSpecial(&_PMintysale.TransactOpts, tokenId, creator, ipfsString, price)
}

// OfferSpecial is a paid mutator transaction binding the contract method 0x4aad4e03.
//
// Solidity: function offerSpecial(uint256 tokenId, address creator, string ipfsString, uint256 price) returns()
func (_PMintysale *PMintysaleTransactorSession) OfferSpecial(tokenId *big.Int, creator common.Address, ipfsString string, price *big.Int) (*types.Transaction, error) {
	return _PMintysale.Contract.OfferSpecial(&_PMintysale.TransactOpts, tokenId, creator, ipfsString, price)
}

// ReSubmitOffer is a paid mutator transaction binding the contract method 0x07c6061c.
//
// Solidity: function reSubmitOffer(uint256 tokenId, uint256 price) returns()
func (_PMintysale *PMintysaleTransactor) ReSubmitOffer(opts *bind.TransactOpts, tokenId *big.Int, price *big.Int) (*types.Transaction, error) {
	return _PMintysale.contract.Transact(opts, "reSubmitOffer", tokenId, price)
}

// ReSubmitOffer is a paid mutator transaction binding the contract method 0x07c6061c.
//
// Solidity: function reSubmitOffer(uint256 tokenId, uint256 price) returns()
func (_PMintysale *PMintysaleSession) ReSubmitOffer(tokenId *big.Int, price *big.Int) (*types.Transaction, error) {
	return _PMintysale.Contract.ReSubmitOffer(&_PMintysale.TransactOpts, tokenId, price)
}

// ReSubmitOffer is a paid mutator transaction binding the contract method 0x07c6061c.
//
// Solidity: function reSubmitOffer(uint256 tokenId, uint256 price) returns()
func (_PMintysale *PMintysaleTransactorSession) ReSubmitOffer(tokenId *big.Int, price *big.Int) (*types.Transaction, error) {
	return _PMintysale.Contract.ReSubmitOffer(&_PMintysale.TransactOpts, tokenId, price)
}

// RetractOffer is a paid mutator transaction binding the contract method 0xc4c6c87d.
//
// Solidity: function retractOffer(uint256 tokenId) returns()
func (_PMintysale *PMintysaleTransactor) RetractOffer(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _PMintysale.contract.Transact(opts, "retractOffer", tokenId)
}

// RetractOffer is a paid mutator transaction binding the contract method 0xc4c6c87d.
//
// Solidity: function retractOffer(uint256 tokenId) returns()
func (_PMintysale *PMintysaleSession) RetractOffer(tokenId *big.Int) (*types.Transaction, error) {
	return _PMintysale.Contract.RetractOffer(&_PMintysale.TransactOpts, tokenId)
}

// RetractOffer is a paid mutator transaction binding the contract method 0xc4c6c87d.
//
// Solidity: function retractOffer(uint256 tokenId) returns()
func (_PMintysale *PMintysaleTransactorSession) RetractOffer(tokenId *big.Int) (*types.Transaction, error) {
	return _PMintysale.Contract.RetractOffer(&_PMintysale.TransactOpts, tokenId)
}

// SetMinter is a paid mutator transaction binding the contract method 0xcf456ae7.
//
// Solidity: function setMinter(address minter, bool status) returns()
func (_PMintysale *PMintysaleTransactor) SetMinter(opts *bind.TransactOpts, minter common.Address, status bool) (*types.Transaction, error) {
	return _PMintysale.contract.Transact(opts, "setMinter", minter, status)
}

// SetMinter is a paid mutator transaction binding the contract method 0xcf456ae7.
//
// Solidity: function setMinter(address minter, bool status) returns()
func (_PMintysale *PMintysaleSession) SetMinter(minter common.Address, status bool) (*types.Transaction, error) {
	return _PMintysale.Contract.SetMinter(&_PMintysale.TransactOpts, minter, status)
}

// SetMinter is a paid mutator transaction binding the contract method 0xcf456ae7.
//
// Solidity: function setMinter(address minter, bool status) returns()
func (_PMintysale *PMintysaleTransactorSession) SetMinter(minter common.Address, status bool) (*types.Transaction, error) {
	return _PMintysale.Contract.SetMinter(&_PMintysale.TransactOpts, minter, status)
}

// SetMintyUnique is a paid mutator transaction binding the contract method 0xb04baaa0.
//
// Solidity: function setMintyUnique(address m) returns()
func (_PMintysale *PMintysaleTransactor) SetMintyUnique(opts *bind.TransactOpts, m common.Address) (*types.Transaction, error) {
	return _PMintysale.contract.Transact(opts, "setMintyUnique", m)
}

// SetMintyUnique is a paid mutator transaction binding the contract method 0xb04baaa0.
//
// Solidity: function setMintyUnique(address m) returns()
func (_PMintysale *PMintysaleSession) SetMintyUnique(m common.Address) (*types.Transaction, error) {
	return _PMintysale.Contract.SetMintyUnique(&_PMintysale.TransactOpts, m)
}

// SetMintyUnique is a paid mutator transaction binding the contract method 0xb04baaa0.
//
// Solidity: function setMintyUnique(address m) returns()
func (_PMintysale *PMintysaleTransactorSession) SetMintyUnique(m common.Address) (*types.Transaction, error) {
	return _PMintysale.Contract.SetMintyUnique(&_PMintysale.TransactOpts, m)
}

// UpdateShares is a paid mutator transaction binding the contract method 0xcb61b42d.
//
// Solidity: function updateShares(uint256 _ownerPerMille, uint256 _creatorPerMille, uint256 _divisor) returns()
func (_PMintysale *PMintysaleTransactor) UpdateShares(opts *bind.TransactOpts, _ownerPerMille *big.Int, _creatorPerMille *big.Int, _divisor *big.Int) (*types.Transaction, error) {
	return _PMintysale.contract.Transact(opts, "updateShares", _ownerPerMille, _creatorPerMille, _divisor)
}

// UpdateShares is a paid mutator transaction binding the contract method 0xcb61b42d.
//
// Solidity: function updateShares(uint256 _ownerPerMille, uint256 _creatorPerMille, uint256 _divisor) returns()
func (_PMintysale *PMintysaleSession) UpdateShares(_ownerPerMille *big.Int, _creatorPerMille *big.Int, _divisor *big.Int) (*types.Transaction, error) {
	return _PMintysale.Contract.UpdateShares(&_PMintysale.TransactOpts, _ownerPerMille, _creatorPerMille, _divisor)
}

// UpdateShares is a paid mutator transaction binding the contract method 0xcb61b42d.
//
// Solidity: function updateShares(uint256 _ownerPerMille, uint256 _creatorPerMille, uint256 _divisor) returns()
func (_PMintysale *PMintysaleTransactorSession) UpdateShares(_ownerPerMille *big.Int, _creatorPerMille *big.Int, _divisor *big.Int) (*types.Transaction, error) {
	return _PMintysale.Contract.UpdateShares(&_PMintysale.TransactOpts, _ownerPerMille, _creatorPerMille, _divisor)
}

// WithdrawBid is a paid mutator transaction binding the contract method 0x0eaaf4c8.
//
// Solidity: function withdrawBid(uint256 tokenId) returns()
func (_PMintysale *PMintysaleTransactor) WithdrawBid(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _PMintysale.contract.Transact(opts, "withdrawBid", tokenId)
}

// WithdrawBid is a paid mutator transaction binding the contract method 0x0eaaf4c8.
//
// Solidity: function withdrawBid(uint256 tokenId) returns()
func (_PMintysale *PMintysaleSession) WithdrawBid(tokenId *big.Int) (*types.Transaction, error) {
	return _PMintysale.Contract.WithdrawBid(&_PMintysale.TransactOpts, tokenId)
}

// WithdrawBid is a paid mutator transaction binding the contract method 0x0eaaf4c8.
//
// Solidity: function withdrawBid(uint256 tokenId) returns()
func (_PMintysale *PMintysaleTransactorSession) WithdrawBid(tokenId *big.Int) (*types.Transaction, error) {
	return _PMintysale.Contract.WithdrawBid(&_PMintysale.TransactOpts, tokenId)
}

// PMintysaleBidIncreasedIterator is returned from FilterBidIncreased and is used to iterate over the raw logs and unpacked data for BidIncreased events raised by the PMintysale contract.
type PMintysaleBidIncreasedIterator struct {
	Event *PMintysaleBidIncreased // Event containing the contract specifics and raw log

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
func (it *PMintysaleBidIncreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PMintysaleBidIncreased)
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
		it.Event = new(PMintysaleBidIncreased)
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
func (it *PMintysaleBidIncreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PMintysaleBidIncreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PMintysaleBidIncreased represents a BidIncreased event raised by the PMintysale contract.
type PMintysaleBidIncreased struct {
	Bidder      common.Address
	TokenId     *big.Int
	PreviousBid *big.Int
	ThisBid     *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBidIncreased is a free log retrieval operation binding the contract event 0x34762811225e202a6ffaa6152f171af086267130932de1c3951c1e36ff89b28a.
//
// Solidity: event BidIncreased(address bidder, uint256 tokenId, uint256 previous_bid, uint256 this_bid)
func (_PMintysale *PMintysaleFilterer) FilterBidIncreased(opts *bind.FilterOpts) (*PMintysaleBidIncreasedIterator, error) {

	logs, sub, err := _PMintysale.contract.FilterLogs(opts, "BidIncreased")
	if err != nil {
		return nil, err
	}
	return &PMintysaleBidIncreasedIterator{contract: _PMintysale.contract, event: "BidIncreased", logs: logs, sub: sub}, nil
}

// WatchBidIncreased is a free log subscription operation binding the contract event 0x34762811225e202a6ffaa6152f171af086267130932de1c3951c1e36ff89b28a.
//
// Solidity: event BidIncreased(address bidder, uint256 tokenId, uint256 previous_bid, uint256 this_bid)
func (_PMintysale *PMintysaleFilterer) WatchBidIncreased(opts *bind.WatchOpts, sink chan<- *PMintysaleBidIncreased) (event.Subscription, error) {

	logs, sub, err := _PMintysale.contract.WatchLogs(opts, "BidIncreased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PMintysaleBidIncreased)
				if err := _PMintysale.contract.UnpackLog(event, "BidIncreased", log); err != nil {
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

// ParseBidIncreased is a log parse operation binding the contract event 0x34762811225e202a6ffaa6152f171af086267130932de1c3951c1e36ff89b28a.
//
// Solidity: event BidIncreased(address bidder, uint256 tokenId, uint256 previous_bid, uint256 this_bid)
func (_PMintysale *PMintysaleFilterer) ParseBidIncreased(log types.Log) (*PMintysaleBidIncreased, error) {
	event := new(PMintysaleBidIncreased)
	if err := _PMintysale.contract.UnpackLog(event, "BidIncreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PMintysaleBidReceivedIterator is returned from FilterBidReceived and is used to iterate over the raw logs and unpacked data for BidReceived events raised by the PMintysale contract.
type PMintysaleBidReceivedIterator struct {
	Event *PMintysaleBidReceived // Event containing the contract specifics and raw log

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
func (it *PMintysaleBidReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PMintysaleBidReceived)
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
		it.Event = new(PMintysaleBidReceived)
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
func (it *PMintysaleBidReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PMintysaleBidReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PMintysaleBidReceived represents a BidReceived event raised by the PMintysale contract.
type PMintysaleBidReceived struct {
	Bidder  common.Address
	TokenId *big.Int
	Bid     *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBidReceived is a free log retrieval operation binding the contract event 0x7f06de72678b8f755d2b9cce42fb004f364fad13cafec8e3b43e8808d9db828c.
//
// Solidity: event BidReceived(address bidder, uint256 tokenId, uint256 bid)
func (_PMintysale *PMintysaleFilterer) FilterBidReceived(opts *bind.FilterOpts) (*PMintysaleBidReceivedIterator, error) {

	logs, sub, err := _PMintysale.contract.FilterLogs(opts, "BidReceived")
	if err != nil {
		return nil, err
	}
	return &PMintysaleBidReceivedIterator{contract: _PMintysale.contract, event: "BidReceived", logs: logs, sub: sub}, nil
}

// WatchBidReceived is a free log subscription operation binding the contract event 0x7f06de72678b8f755d2b9cce42fb004f364fad13cafec8e3b43e8808d9db828c.
//
// Solidity: event BidReceived(address bidder, uint256 tokenId, uint256 bid)
func (_PMintysale *PMintysaleFilterer) WatchBidReceived(opts *bind.WatchOpts, sink chan<- *PMintysaleBidReceived) (event.Subscription, error) {

	logs, sub, err := _PMintysale.contract.WatchLogs(opts, "BidReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PMintysaleBidReceived)
				if err := _PMintysale.contract.UnpackLog(event, "BidReceived", log); err != nil {
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

// ParseBidReceived is a log parse operation binding the contract event 0x7f06de72678b8f755d2b9cce42fb004f364fad13cafec8e3b43e8808d9db828c.
//
// Solidity: event BidReceived(address bidder, uint256 tokenId, uint256 bid)
func (_PMintysale *PMintysaleFilterer) ParseBidReceived(log types.Log) (*PMintysaleBidReceived, error) {
	event := new(PMintysaleBidReceived)
	if err := _PMintysale.contract.UnpackLog(event, "BidReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PMintysaleNewOfferIterator is returned from FilterNewOffer and is used to iterate over the raw logs and unpacked data for NewOffer events raised by the PMintysale contract.
type PMintysaleNewOfferIterator struct {
	Event *PMintysaleNewOffer // Event containing the contract specifics and raw log

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
func (it *PMintysaleNewOfferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PMintysaleNewOffer)
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
		it.Event = new(PMintysaleNewOffer)
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
func (it *PMintysaleNewOfferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PMintysaleNewOfferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PMintysaleNewOffer represents a NewOffer event raised by the PMintysale contract.
type PMintysaleNewOffer struct {
	TokenId *big.Int
	Owner   common.Address
	Price   *big.Int
	Hash    string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewOffer is a free log retrieval operation binding the contract event 0x3697349224444578adf7f32ac71b635cf39a0924633e5700169ed207bf91f7b4.
//
// Solidity: event NewOffer(uint256 tokenId, address owner, uint256 price, string hash)
func (_PMintysale *PMintysaleFilterer) FilterNewOffer(opts *bind.FilterOpts) (*PMintysaleNewOfferIterator, error) {

	logs, sub, err := _PMintysale.contract.FilterLogs(opts, "NewOffer")
	if err != nil {
		return nil, err
	}
	return &PMintysaleNewOfferIterator{contract: _PMintysale.contract, event: "NewOffer", logs: logs, sub: sub}, nil
}

// WatchNewOffer is a free log subscription operation binding the contract event 0x3697349224444578adf7f32ac71b635cf39a0924633e5700169ed207bf91f7b4.
//
// Solidity: event NewOffer(uint256 tokenId, address owner, uint256 price, string hash)
func (_PMintysale *PMintysaleFilterer) WatchNewOffer(opts *bind.WatchOpts, sink chan<- *PMintysaleNewOffer) (event.Subscription, error) {

	logs, sub, err := _PMintysale.contract.WatchLogs(opts, "NewOffer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PMintysaleNewOffer)
				if err := _PMintysale.contract.UnpackLog(event, "NewOffer", log); err != nil {
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

// ParseNewOffer is a log parse operation binding the contract event 0x3697349224444578adf7f32ac71b635cf39a0924633e5700169ed207bf91f7b4.
//
// Solidity: event NewOffer(uint256 tokenId, address owner, uint256 price, string hash)
func (_PMintysale *PMintysaleFilterer) ParseNewOffer(log types.Log) (*PMintysaleNewOffer, error) {
	event := new(PMintysaleNewOffer)
	if err := _PMintysale.contract.UnpackLog(event, "NewOffer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PMintysaleOfferAcceptedIterator is returned from FilterOfferAccepted and is used to iterate over the raw logs and unpacked data for OfferAccepted events raised by the PMintysale contract.
type PMintysaleOfferAcceptedIterator struct {
	Event *PMintysaleOfferAccepted // Event containing the contract specifics and raw log

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
func (it *PMintysaleOfferAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PMintysaleOfferAccepted)
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
		it.Event = new(PMintysaleOfferAccepted)
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
func (it *PMintysaleOfferAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PMintysaleOfferAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PMintysaleOfferAccepted represents a OfferAccepted event raised by the PMintysale contract.
type PMintysaleOfferAccepted struct {
	Buyer   common.Address
	TokenId *big.Int
	Price   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOfferAccepted is a free log retrieval operation binding the contract event 0x2e91b12fed89e8218510d4f7c38c6db746acd0984f09b436b69bd7db5e93da8b.
//
// Solidity: event OfferAccepted(address buyer, uint256 tokenId, uint256 price)
func (_PMintysale *PMintysaleFilterer) FilterOfferAccepted(opts *bind.FilterOpts) (*PMintysaleOfferAcceptedIterator, error) {

	logs, sub, err := _PMintysale.contract.FilterLogs(opts, "OfferAccepted")
	if err != nil {
		return nil, err
	}
	return &PMintysaleOfferAcceptedIterator{contract: _PMintysale.contract, event: "OfferAccepted", logs: logs, sub: sub}, nil
}

// WatchOfferAccepted is a free log subscription operation binding the contract event 0x2e91b12fed89e8218510d4f7c38c6db746acd0984f09b436b69bd7db5e93da8b.
//
// Solidity: event OfferAccepted(address buyer, uint256 tokenId, uint256 price)
func (_PMintysale *PMintysaleFilterer) WatchOfferAccepted(opts *bind.WatchOpts, sink chan<- *PMintysaleOfferAccepted) (event.Subscription, error) {

	logs, sub, err := _PMintysale.contract.WatchLogs(opts, "OfferAccepted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PMintysaleOfferAccepted)
				if err := _PMintysale.contract.UnpackLog(event, "OfferAccepted", log); err != nil {
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

// ParseOfferAccepted is a log parse operation binding the contract event 0x2e91b12fed89e8218510d4f7c38c6db746acd0984f09b436b69bd7db5e93da8b.
//
// Solidity: event OfferAccepted(address buyer, uint256 tokenId, uint256 price)
func (_PMintysale *PMintysaleFilterer) ParseOfferAccepted(log types.Log) (*PMintysaleOfferAccepted, error) {
	event := new(PMintysaleOfferAccepted)
	if err := _PMintysale.contract.UnpackLog(event, "OfferAccepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PMintysalePaymentIterator is returned from FilterPayment and is used to iterate over the raw logs and unpacked data for Payment events raised by the PMintysale contract.
type PMintysalePaymentIterator struct {
	Event *PMintysalePayment // Event containing the contract specifics and raw log

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
func (it *PMintysalePaymentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PMintysalePayment)
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
		it.Event = new(PMintysalePayment)
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
func (it *PMintysalePaymentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PMintysalePaymentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PMintysalePayment represents a Payment event raised by the PMintysale contract.
type PMintysalePayment struct {
	Wallet  common.Address
	Creator common.Address
	Owner   common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPayment is a free log retrieval operation binding the contract event 0x8703df07b43faff82a32cfdce460e22eb4b029f946f61df422e57d5ef66f3962.
//
// Solidity: event Payment(address wallet, address creator, address _owner)
func (_PMintysale *PMintysaleFilterer) FilterPayment(opts *bind.FilterOpts) (*PMintysalePaymentIterator, error) {

	logs, sub, err := _PMintysale.contract.FilterLogs(opts, "Payment")
	if err != nil {
		return nil, err
	}
	return &PMintysalePaymentIterator{contract: _PMintysale.contract, event: "Payment", logs: logs, sub: sub}, nil
}

// WatchPayment is a free log subscription operation binding the contract event 0x8703df07b43faff82a32cfdce460e22eb4b029f946f61df422e57d5ef66f3962.
//
// Solidity: event Payment(address wallet, address creator, address _owner)
func (_PMintysale *PMintysaleFilterer) WatchPayment(opts *bind.WatchOpts, sink chan<- *PMintysalePayment) (event.Subscription, error) {

	logs, sub, err := _PMintysale.contract.WatchLogs(opts, "Payment")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PMintysalePayment)
				if err := _PMintysale.contract.UnpackLog(event, "Payment", log); err != nil {
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

// ParsePayment is a log parse operation binding the contract event 0x8703df07b43faff82a32cfdce460e22eb4b029f946f61df422e57d5ef66f3962.
//
// Solidity: event Payment(address wallet, address creator, address _owner)
func (_PMintysale *PMintysaleFilterer) ParsePayment(log types.Log) (*PMintysalePayment, error) {
	event := new(PMintysalePayment)
	if err := _PMintysale.contract.UnpackLog(event, "Payment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PMintysaleResaleOfferIterator is returned from FilterResaleOffer and is used to iterate over the raw logs and unpacked data for ResaleOffer events raised by the PMintysale contract.
type PMintysaleResaleOfferIterator struct {
	Event *PMintysaleResaleOffer // Event containing the contract specifics and raw log

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
func (it *PMintysaleResaleOfferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PMintysaleResaleOffer)
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
		it.Event = new(PMintysaleResaleOffer)
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
func (it *PMintysaleResaleOfferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PMintysaleResaleOfferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PMintysaleResaleOffer represents a ResaleOffer event raised by the PMintysale contract.
type PMintysaleResaleOffer struct {
	TokenId *big.Int
	Owner   common.Address
	Price   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterResaleOffer is a free log retrieval operation binding the contract event 0xebe4318722b948bede41b4decd121e3e2ab187f11ad89499eb7a586de7ffcd90.
//
// Solidity: event ResaleOffer(uint256 tokenId, address owner, uint256 price)
func (_PMintysale *PMintysaleFilterer) FilterResaleOffer(opts *bind.FilterOpts) (*PMintysaleResaleOfferIterator, error) {

	logs, sub, err := _PMintysale.contract.FilterLogs(opts, "ResaleOffer")
	if err != nil {
		return nil, err
	}
	return &PMintysaleResaleOfferIterator{contract: _PMintysale.contract, event: "ResaleOffer", logs: logs, sub: sub}, nil
}

// WatchResaleOffer is a free log subscription operation binding the contract event 0xebe4318722b948bede41b4decd121e3e2ab187f11ad89499eb7a586de7ffcd90.
//
// Solidity: event ResaleOffer(uint256 tokenId, address owner, uint256 price)
func (_PMintysale *PMintysaleFilterer) WatchResaleOffer(opts *bind.WatchOpts, sink chan<- *PMintysaleResaleOffer) (event.Subscription, error) {

	logs, sub, err := _PMintysale.contract.WatchLogs(opts, "ResaleOffer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PMintysaleResaleOffer)
				if err := _PMintysale.contract.UnpackLog(event, "ResaleOffer", log); err != nil {
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

// ParseResaleOffer is a log parse operation binding the contract event 0xebe4318722b948bede41b4decd121e3e2ab187f11ad89499eb7a586de7ffcd90.
//
// Solidity: event ResaleOffer(uint256 tokenId, address owner, uint256 price)
func (_PMintysale *PMintysaleFilterer) ParseResaleOffer(log types.Log) (*PMintysaleResaleOffer, error) {
	event := new(PMintysaleResaleOffer)
	if err := _PMintysale.contract.UnpackLog(event, "ResaleOffer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PMintysaleSaleResubmittedIterator is returned from FilterSaleResubmitted and is used to iterate over the raw logs and unpacked data for SaleResubmitted events raised by the PMintysale contract.
type PMintysaleSaleResubmittedIterator struct {
	Event *PMintysaleSaleResubmitted // Event containing the contract specifics and raw log

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
func (it *PMintysaleSaleResubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PMintysaleSaleResubmitted)
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
		it.Event = new(PMintysaleSaleResubmitted)
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
func (it *PMintysaleSaleResubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PMintysaleSaleResubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PMintysaleSaleResubmitted represents a SaleResubmitted event raised by the PMintysale contract.
type PMintysaleSaleResubmitted struct {
	TokenId *big.Int
	Price   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSaleResubmitted is a free log retrieval operation binding the contract event 0xe36efdb123434f0d9dcc119c468c16b0d6de5fa64252b13e049b2c87dafc2a1e.
//
// Solidity: event SaleResubmitted(uint256 tokenId, uint256 price)
func (_PMintysale *PMintysaleFilterer) FilterSaleResubmitted(opts *bind.FilterOpts) (*PMintysaleSaleResubmittedIterator, error) {

	logs, sub, err := _PMintysale.contract.FilterLogs(opts, "SaleResubmitted")
	if err != nil {
		return nil, err
	}
	return &PMintysaleSaleResubmittedIterator{contract: _PMintysale.contract, event: "SaleResubmitted", logs: logs, sub: sub}, nil
}

// WatchSaleResubmitted is a free log subscription operation binding the contract event 0xe36efdb123434f0d9dcc119c468c16b0d6de5fa64252b13e049b2c87dafc2a1e.
//
// Solidity: event SaleResubmitted(uint256 tokenId, uint256 price)
func (_PMintysale *PMintysaleFilterer) WatchSaleResubmitted(opts *bind.WatchOpts, sink chan<- *PMintysaleSaleResubmitted) (event.Subscription, error) {

	logs, sub, err := _PMintysale.contract.WatchLogs(opts, "SaleResubmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PMintysaleSaleResubmitted)
				if err := _PMintysale.contract.UnpackLog(event, "SaleResubmitted", log); err != nil {
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

// ParseSaleResubmitted is a log parse operation binding the contract event 0xe36efdb123434f0d9dcc119c468c16b0d6de5fa64252b13e049b2c87dafc2a1e.
//
// Solidity: event SaleResubmitted(uint256 tokenId, uint256 price)
func (_PMintysale *PMintysaleFilterer) ParseSaleResubmitted(log types.Log) (*PMintysaleSaleResubmitted, error) {
	event := new(PMintysaleSaleResubmitted)
	if err := _PMintysale.contract.UnpackLog(event, "SaleResubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PMintysaleSaleRetractedIterator is returned from FilterSaleRetracted and is used to iterate over the raw logs and unpacked data for SaleRetracted events raised by the PMintysale contract.
type PMintysaleSaleRetractedIterator struct {
	Event *PMintysaleSaleRetracted // Event containing the contract specifics and raw log

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
func (it *PMintysaleSaleRetractedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PMintysaleSaleRetracted)
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
		it.Event = new(PMintysaleSaleRetracted)
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
func (it *PMintysaleSaleRetractedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PMintysaleSaleRetractedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PMintysaleSaleRetracted represents a SaleRetracted event raised by the PMintysale contract.
type PMintysaleSaleRetracted struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSaleRetracted is a free log retrieval operation binding the contract event 0x567b7c6ccef395faa0faf04e2da702ec06c19c0cff01cde647c3d285257a0bb0.
//
// Solidity: event SaleRetracted(uint256 tokenId)
func (_PMintysale *PMintysaleFilterer) FilterSaleRetracted(opts *bind.FilterOpts) (*PMintysaleSaleRetractedIterator, error) {

	logs, sub, err := _PMintysale.contract.FilterLogs(opts, "SaleRetracted")
	if err != nil {
		return nil, err
	}
	return &PMintysaleSaleRetractedIterator{contract: _PMintysale.contract, event: "SaleRetracted", logs: logs, sub: sub}, nil
}

// WatchSaleRetracted is a free log subscription operation binding the contract event 0x567b7c6ccef395faa0faf04e2da702ec06c19c0cff01cde647c3d285257a0bb0.
//
// Solidity: event SaleRetracted(uint256 tokenId)
func (_PMintysale *PMintysaleFilterer) WatchSaleRetracted(opts *bind.WatchOpts, sink chan<- *PMintysaleSaleRetracted) (event.Subscription, error) {

	logs, sub, err := _PMintysale.contract.WatchLogs(opts, "SaleRetracted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PMintysaleSaleRetracted)
				if err := _PMintysale.contract.UnpackLog(event, "SaleRetracted", log); err != nil {
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

// ParseSaleRetracted is a log parse operation binding the contract event 0x567b7c6ccef395faa0faf04e2da702ec06c19c0cff01cde647c3d285257a0bb0.
//
// Solidity: event SaleRetracted(uint256 tokenId)
func (_PMintysale *PMintysaleFilterer) ParseSaleRetracted(log types.Log) (*PMintysaleSaleRetracted, error) {
	event := new(PMintysaleSaleRetracted)
	if err := _PMintysale.contract.UnpackLog(event, "SaleRetracted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PMintysaleSharesUpdatedIterator is returned from FilterSharesUpdated and is used to iterate over the raw logs and unpacked data for SharesUpdated events raised by the PMintysale contract.
type PMintysaleSharesUpdatedIterator struct {
	Event *PMintysaleSharesUpdated // Event containing the contract specifics and raw log

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
func (it *PMintysaleSharesUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PMintysaleSharesUpdated)
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
		it.Event = new(PMintysaleSharesUpdated)
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
func (it *PMintysaleSharesUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PMintysaleSharesUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PMintysaleSharesUpdated represents a SharesUpdated event raised by the PMintysale contract.
type PMintysaleSharesUpdated struct {
	OwnerShare   *big.Int
	CreatorShare *big.Int
	Divisor      *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSharesUpdated is a free log retrieval operation binding the contract event 0xf81b1f9d8b5d36a5a67d13006a8b091a9b63e2a38b331d315d99c6daefd2ddcc.
//
// Solidity: event SharesUpdated(uint256 ownerShare, uint256 creatorShare, uint256 divisor)
func (_PMintysale *PMintysaleFilterer) FilterSharesUpdated(opts *bind.FilterOpts) (*PMintysaleSharesUpdatedIterator, error) {

	logs, sub, err := _PMintysale.contract.FilterLogs(opts, "SharesUpdated")
	if err != nil {
		return nil, err
	}
	return &PMintysaleSharesUpdatedIterator{contract: _PMintysale.contract, event: "SharesUpdated", logs: logs, sub: sub}, nil
}

// WatchSharesUpdated is a free log subscription operation binding the contract event 0xf81b1f9d8b5d36a5a67d13006a8b091a9b63e2a38b331d315d99c6daefd2ddcc.
//
// Solidity: event SharesUpdated(uint256 ownerShare, uint256 creatorShare, uint256 divisor)
func (_PMintysale *PMintysaleFilterer) WatchSharesUpdated(opts *bind.WatchOpts, sink chan<- *PMintysaleSharesUpdated) (event.Subscription, error) {

	logs, sub, err := _PMintysale.contract.WatchLogs(opts, "SharesUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PMintysaleSharesUpdated)
				if err := _PMintysale.contract.UnpackLog(event, "SharesUpdated", log); err != nil {
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

// ParseSharesUpdated is a log parse operation binding the contract event 0xf81b1f9d8b5d36a5a67d13006a8b091a9b63e2a38b331d315d99c6daefd2ddcc.
//
// Solidity: event SharesUpdated(uint256 ownerShare, uint256 creatorShare, uint256 divisor)
func (_PMintysale *PMintysaleFilterer) ParseSharesUpdated(log types.Log) (*PMintysaleSharesUpdated, error) {
	event := new(PMintysaleSharesUpdated)
	if err := _PMintysale.contract.UnpackLog(event, "SharesUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
