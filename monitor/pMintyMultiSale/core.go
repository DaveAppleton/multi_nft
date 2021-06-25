// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pMintyMultiSale

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

// PoolEntry is an auto generated low-level Go binding around an user-defined struct.
type PoolEntry struct {
	Beneficiary common.Address
	Share       *big.Int
}

// IERC20ABI is the input ABI used to generate the binding from.
const IERC20ABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IERC20FuncSigs maps the 4-byte function signature to its string representation.
var IERC20FuncSigs = map[string]string{
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

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address owner, address receiver, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, owner common.Address, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", owner, receiver, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address owner, address receiver, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(owner common.Address, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, owner, receiver, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address owner, address receiver, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(owner common.Address, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, owner, receiver, amount)
}

// IMintyMultiTokenABI is the input ABI used to generate the binding from.
const IMintyMultiTokenABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"creator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"saleNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getRoyalties\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"share\",\"type\":\"uint256\"}],\"internalType\":\"structPoolEntry[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"ipfsHash\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"quantities\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"hashes\",\"type\":\"string[]\"},{\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"}],\"name\":\"mintBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"minted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"royaltyPerMille\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"}],\"name\":\"validateBuyer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IMintyMultiTokenFuncSigs maps the 4-byte function signature to its string representation.
var IMintyMultiTokenFuncSigs = map[string]string{
	"00fdd58e": "balanceOf(address,uint256)",
	"510b5158": "creator(uint256)",
	"e03650b9": "getRoyalties(uint256,uint256)",
	"e985e9c5": "isApprovedForAll(address,address)",
	"cf2ed32e": "mint(uint256,uint256,string,uint256)",
	"2a637d5c": "mintBatch(uint256[],uint256[],string[],uint256)",
	"7dc0bf3f": "minted(uint256)",
	"8da5cb5b": "owner()",
	"e8f3359f": "royaltyPerMille()",
	"f242432a": "safeTransferFrom(address,address,uint256,uint256,bytes)",
	"0e89341c": "uri(uint256)",
	"5773bd07": "validateBuyer(address)",
}

// IMintyMultiToken is an auto generated Go binding around an Ethereum contract.
type IMintyMultiToken struct {
	IMintyMultiTokenCaller     // Read-only binding to the contract
	IMintyMultiTokenTransactor // Write-only binding to the contract
	IMintyMultiTokenFilterer   // Log filterer for contract events
}

// IMintyMultiTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type IMintyMultiTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMintyMultiTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IMintyMultiTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMintyMultiTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IMintyMultiTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMintyMultiTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IMintyMultiTokenSession struct {
	Contract     *IMintyMultiToken // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IMintyMultiTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IMintyMultiTokenCallerSession struct {
	Contract *IMintyMultiTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IMintyMultiTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IMintyMultiTokenTransactorSession struct {
	Contract     *IMintyMultiTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IMintyMultiTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type IMintyMultiTokenRaw struct {
	Contract *IMintyMultiToken // Generic contract binding to access the raw methods on
}

// IMintyMultiTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IMintyMultiTokenCallerRaw struct {
	Contract *IMintyMultiTokenCaller // Generic read-only contract binding to access the raw methods on
}

// IMintyMultiTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IMintyMultiTokenTransactorRaw struct {
	Contract *IMintyMultiTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIMintyMultiToken creates a new instance of IMintyMultiToken, bound to a specific deployed contract.
func NewIMintyMultiToken(address common.Address, backend bind.ContractBackend) (*IMintyMultiToken, error) {
	contract, err := bindIMintyMultiToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IMintyMultiToken{IMintyMultiTokenCaller: IMintyMultiTokenCaller{contract: contract}, IMintyMultiTokenTransactor: IMintyMultiTokenTransactor{contract: contract}, IMintyMultiTokenFilterer: IMintyMultiTokenFilterer{contract: contract}}, nil
}

// NewIMintyMultiTokenCaller creates a new read-only instance of IMintyMultiToken, bound to a specific deployed contract.
func NewIMintyMultiTokenCaller(address common.Address, caller bind.ContractCaller) (*IMintyMultiTokenCaller, error) {
	contract, err := bindIMintyMultiToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IMintyMultiTokenCaller{contract: contract}, nil
}

// NewIMintyMultiTokenTransactor creates a new write-only instance of IMintyMultiToken, bound to a specific deployed contract.
func NewIMintyMultiTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*IMintyMultiTokenTransactor, error) {
	contract, err := bindIMintyMultiToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IMintyMultiTokenTransactor{contract: contract}, nil
}

// NewIMintyMultiTokenFilterer creates a new log filterer instance of IMintyMultiToken, bound to a specific deployed contract.
func NewIMintyMultiTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*IMintyMultiTokenFilterer, error) {
	contract, err := bindIMintyMultiToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IMintyMultiTokenFilterer{contract: contract}, nil
}

// bindIMintyMultiToken binds a generic wrapper to an already deployed contract.
func bindIMintyMultiToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IMintyMultiTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMintyMultiToken *IMintyMultiTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMintyMultiToken.Contract.IMintyMultiTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMintyMultiToken *IMintyMultiTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMintyMultiToken.Contract.IMintyMultiTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMintyMultiToken *IMintyMultiTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMintyMultiToken.Contract.IMintyMultiTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMintyMultiToken *IMintyMultiTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMintyMultiToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMintyMultiToken *IMintyMultiTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMintyMultiToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMintyMultiToken *IMintyMultiTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMintyMultiToken.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_IMintyMultiToken *IMintyMultiTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IMintyMultiToken.contract.Call(opts, &out, "balanceOf", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_IMintyMultiToken *IMintyMultiTokenSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _IMintyMultiToken.Contract.BalanceOf(&_IMintyMultiToken.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_IMintyMultiToken *IMintyMultiTokenCallerSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _IMintyMultiToken.Contract.BalanceOf(&_IMintyMultiToken.CallOpts, account, id)
}

// Creator is a free data retrieval call binding the contract method 0x510b5158.
//
// Solidity: function creator(uint256 tokenId) view returns(address)
func (_IMintyMultiToken *IMintyMultiTokenCaller) Creator(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IMintyMultiToken.contract.Call(opts, &out, "creator", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Creator is a free data retrieval call binding the contract method 0x510b5158.
//
// Solidity: function creator(uint256 tokenId) view returns(address)
func (_IMintyMultiToken *IMintyMultiTokenSession) Creator(tokenId *big.Int) (common.Address, error) {
	return _IMintyMultiToken.Contract.Creator(&_IMintyMultiToken.CallOpts, tokenId)
}

// Creator is a free data retrieval call binding the contract method 0x510b5158.
//
// Solidity: function creator(uint256 tokenId) view returns(address)
func (_IMintyMultiToken *IMintyMultiTokenCallerSession) Creator(tokenId *big.Int) (common.Address, error) {
	return _IMintyMultiToken.Contract.Creator(&_IMintyMultiToken.CallOpts, tokenId)
}

// GetRoyalties is a free data retrieval call binding the contract method 0xe03650b9.
//
// Solidity: function getRoyalties(uint256 saleNumber, uint256 tokenId) view returns((address,uint256)[])
func (_IMintyMultiToken *IMintyMultiTokenCaller) GetRoyalties(opts *bind.CallOpts, saleNumber *big.Int, tokenId *big.Int) ([]PoolEntry, error) {
	var out []interface{}
	err := _IMintyMultiToken.contract.Call(opts, &out, "getRoyalties", saleNumber, tokenId)

	if err != nil {
		return *new([]PoolEntry), err
	}

	out0 := *abi.ConvertType(out[0], new([]PoolEntry)).(*[]PoolEntry)

	return out0, err

}

// GetRoyalties is a free data retrieval call binding the contract method 0xe03650b9.
//
// Solidity: function getRoyalties(uint256 saleNumber, uint256 tokenId) view returns((address,uint256)[])
func (_IMintyMultiToken *IMintyMultiTokenSession) GetRoyalties(saleNumber *big.Int, tokenId *big.Int) ([]PoolEntry, error) {
	return _IMintyMultiToken.Contract.GetRoyalties(&_IMintyMultiToken.CallOpts, saleNumber, tokenId)
}

// GetRoyalties is a free data retrieval call binding the contract method 0xe03650b9.
//
// Solidity: function getRoyalties(uint256 saleNumber, uint256 tokenId) view returns((address,uint256)[])
func (_IMintyMultiToken *IMintyMultiTokenCallerSession) GetRoyalties(saleNumber *big.Int, tokenId *big.Int) ([]PoolEntry, error) {
	return _IMintyMultiToken.Contract.GetRoyalties(&_IMintyMultiToken.CallOpts, saleNumber, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_IMintyMultiToken *IMintyMultiTokenCaller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _IMintyMultiToken.contract.Call(opts, &out, "isApprovedForAll", account, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_IMintyMultiToken *IMintyMultiTokenSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _IMintyMultiToken.Contract.IsApprovedForAll(&_IMintyMultiToken.CallOpts, account, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_IMintyMultiToken *IMintyMultiTokenCallerSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _IMintyMultiToken.Contract.IsApprovedForAll(&_IMintyMultiToken.CallOpts, account, operator)
}

// Minted is a free data retrieval call binding the contract method 0x7dc0bf3f.
//
// Solidity: function minted(uint256 id) view returns(bool)
func (_IMintyMultiToken *IMintyMultiTokenCaller) Minted(opts *bind.CallOpts, id *big.Int) (bool, error) {
	var out []interface{}
	err := _IMintyMultiToken.contract.Call(opts, &out, "minted", id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Minted is a free data retrieval call binding the contract method 0x7dc0bf3f.
//
// Solidity: function minted(uint256 id) view returns(bool)
func (_IMintyMultiToken *IMintyMultiTokenSession) Minted(id *big.Int) (bool, error) {
	return _IMintyMultiToken.Contract.Minted(&_IMintyMultiToken.CallOpts, id)
}

// Minted is a free data retrieval call binding the contract method 0x7dc0bf3f.
//
// Solidity: function minted(uint256 id) view returns(bool)
func (_IMintyMultiToken *IMintyMultiTokenCallerSession) Minted(id *big.Int) (bool, error) {
	return _IMintyMultiToken.Contract.Minted(&_IMintyMultiToken.CallOpts, id)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IMintyMultiToken *IMintyMultiTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IMintyMultiToken.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IMintyMultiToken *IMintyMultiTokenSession) Owner() (common.Address, error) {
	return _IMintyMultiToken.Contract.Owner(&_IMintyMultiToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IMintyMultiToken *IMintyMultiTokenCallerSession) Owner() (common.Address, error) {
	return _IMintyMultiToken.Contract.Owner(&_IMintyMultiToken.CallOpts)
}

// RoyaltyPerMille is a free data retrieval call binding the contract method 0xe8f3359f.
//
// Solidity: function royaltyPerMille() view returns(uint256)
func (_IMintyMultiToken *IMintyMultiTokenCaller) RoyaltyPerMille(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IMintyMultiToken.contract.Call(opts, &out, "royaltyPerMille")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RoyaltyPerMille is a free data retrieval call binding the contract method 0xe8f3359f.
//
// Solidity: function royaltyPerMille() view returns(uint256)
func (_IMintyMultiToken *IMintyMultiTokenSession) RoyaltyPerMille() (*big.Int, error) {
	return _IMintyMultiToken.Contract.RoyaltyPerMille(&_IMintyMultiToken.CallOpts)
}

// RoyaltyPerMille is a free data retrieval call binding the contract method 0xe8f3359f.
//
// Solidity: function royaltyPerMille() view returns(uint256)
func (_IMintyMultiToken *IMintyMultiTokenCallerSession) RoyaltyPerMille() (*big.Int, error) {
	return _IMintyMultiToken.Contract.RoyaltyPerMille(&_IMintyMultiToken.CallOpts)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 id) view returns(string)
func (_IMintyMultiToken *IMintyMultiTokenCaller) Uri(opts *bind.CallOpts, id *big.Int) (string, error) {
	var out []interface{}
	err := _IMintyMultiToken.contract.Call(opts, &out, "uri", id)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 id) view returns(string)
func (_IMintyMultiToken *IMintyMultiTokenSession) Uri(id *big.Int) (string, error) {
	return _IMintyMultiToken.Contract.Uri(&_IMintyMultiToken.CallOpts, id)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 id) view returns(string)
func (_IMintyMultiToken *IMintyMultiTokenCallerSession) Uri(id *big.Int) (string, error) {
	return _IMintyMultiToken.Contract.Uri(&_IMintyMultiToken.CallOpts, id)
}

// Mint is a paid mutator transaction binding the contract method 0xcf2ed32e.
//
// Solidity: function mint(uint256 tokenId, uint256 quantity, string ipfsHash, uint256 poolId) returns()
func (_IMintyMultiToken *IMintyMultiTokenTransactor) Mint(opts *bind.TransactOpts, tokenId *big.Int, quantity *big.Int, ipfsHash string, poolId *big.Int) (*types.Transaction, error) {
	return _IMintyMultiToken.contract.Transact(opts, "mint", tokenId, quantity, ipfsHash, poolId)
}

// Mint is a paid mutator transaction binding the contract method 0xcf2ed32e.
//
// Solidity: function mint(uint256 tokenId, uint256 quantity, string ipfsHash, uint256 poolId) returns()
func (_IMintyMultiToken *IMintyMultiTokenSession) Mint(tokenId *big.Int, quantity *big.Int, ipfsHash string, poolId *big.Int) (*types.Transaction, error) {
	return _IMintyMultiToken.Contract.Mint(&_IMintyMultiToken.TransactOpts, tokenId, quantity, ipfsHash, poolId)
}

// Mint is a paid mutator transaction binding the contract method 0xcf2ed32e.
//
// Solidity: function mint(uint256 tokenId, uint256 quantity, string ipfsHash, uint256 poolId) returns()
func (_IMintyMultiToken *IMintyMultiTokenTransactorSession) Mint(tokenId *big.Int, quantity *big.Int, ipfsHash string, poolId *big.Int) (*types.Transaction, error) {
	return _IMintyMultiToken.Contract.Mint(&_IMintyMultiToken.TransactOpts, tokenId, quantity, ipfsHash, poolId)
}

// MintBatch is a paid mutator transaction binding the contract method 0x2a637d5c.
//
// Solidity: function mintBatch(uint256[] tokenIds, uint256[] quantities, string[] hashes, uint256 poolId) returns()
func (_IMintyMultiToken *IMintyMultiTokenTransactor) MintBatch(opts *bind.TransactOpts, tokenIds []*big.Int, quantities []*big.Int, hashes []string, poolId *big.Int) (*types.Transaction, error) {
	return _IMintyMultiToken.contract.Transact(opts, "mintBatch", tokenIds, quantities, hashes, poolId)
}

// MintBatch is a paid mutator transaction binding the contract method 0x2a637d5c.
//
// Solidity: function mintBatch(uint256[] tokenIds, uint256[] quantities, string[] hashes, uint256 poolId) returns()
func (_IMintyMultiToken *IMintyMultiTokenSession) MintBatch(tokenIds []*big.Int, quantities []*big.Int, hashes []string, poolId *big.Int) (*types.Transaction, error) {
	return _IMintyMultiToken.Contract.MintBatch(&_IMintyMultiToken.TransactOpts, tokenIds, quantities, hashes, poolId)
}

// MintBatch is a paid mutator transaction binding the contract method 0x2a637d5c.
//
// Solidity: function mintBatch(uint256[] tokenIds, uint256[] quantities, string[] hashes, uint256 poolId) returns()
func (_IMintyMultiToken *IMintyMultiTokenTransactorSession) MintBatch(tokenIds []*big.Int, quantities []*big.Int, hashes []string, poolId *big.Int) (*types.Transaction, error) {
	return _IMintyMultiToken.Contract.MintBatch(&_IMintyMultiToken.TransactOpts, tokenIds, quantities, hashes, poolId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, uint256 amount, bytes data) returns()
func (_IMintyMultiToken *IMintyMultiTokenTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IMintyMultiToken.contract.Transact(opts, "safeTransferFrom", from, to, tokenId, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, uint256 amount, bytes data) returns()
func (_IMintyMultiToken *IMintyMultiTokenSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IMintyMultiToken.Contract.SafeTransferFrom(&_IMintyMultiToken.TransactOpts, from, to, tokenId, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, uint256 amount, bytes data) returns()
func (_IMintyMultiToken *IMintyMultiTokenTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IMintyMultiToken.Contract.SafeTransferFrom(&_IMintyMultiToken.TransactOpts, from, to, tokenId, amount, data)
}

// ValidateBuyer is a paid mutator transaction binding the contract method 0x5773bd07.
//
// Solidity: function validateBuyer(address buyer) returns()
func (_IMintyMultiToken *IMintyMultiTokenTransactor) ValidateBuyer(opts *bind.TransactOpts, buyer common.Address) (*types.Transaction, error) {
	return _IMintyMultiToken.contract.Transact(opts, "validateBuyer", buyer)
}

// ValidateBuyer is a paid mutator transaction binding the contract method 0x5773bd07.
//
// Solidity: function validateBuyer(address buyer) returns()
func (_IMintyMultiToken *IMintyMultiTokenSession) ValidateBuyer(buyer common.Address) (*types.Transaction, error) {
	return _IMintyMultiToken.Contract.ValidateBuyer(&_IMintyMultiToken.TransactOpts, buyer)
}

// ValidateBuyer is a paid mutator transaction binding the contract method 0x5773bd07.
//
// Solidity: function validateBuyer(address buyer) returns()
func (_IMintyMultiToken *IMintyMultiTokenTransactorSession) ValidateBuyer(buyer common.Address) (*types.Transaction, error) {
	return _IMintyMultiToken.Contract.ValidateBuyer(&_IMintyMultiToken.TransactOpts, buyer)
}

// PMintyMultiSaleABI is the input ABI used to generate the binding from.
const PMintyMultiSaleABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"_weth\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_divisor\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIMintyMultiToken\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previous_bid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"this_bid\",\"type\":\"uint256\"}],\"name\":\"BidIncreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIMintyMultiToken\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"}],\"name\":\"BidReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIMintyMultiToken\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"BidWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"divisor\",\"type\":\"uint256\"}],\"name\":\"FeeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIMintyMultiToken\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"hash\",\"type\":\"string\"}],\"name\":\"NewOffer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIMintyMultiToken\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pos\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"OfferAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIMintyMultiToken\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"OfferWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Payment\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIMintyMultiToken\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"position\",\"type\":\"uint256\"}],\"name\":\"ResaleOffer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIMintyMultiToken\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pos\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"SaleRepriced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIMintyMultiToken\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pos\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"SaleRetracted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousWallet\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newWallet\",\"type\":\"address\"}],\"name\":\"WalletTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"contractIMintyMultiToken\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pos\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"name\":\"acceptOffer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIMintyMultiToken\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"offerId\",\"type\":\"uint256\"}],\"name\":\"available\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newWallet\",\"type\":\"address\"}],\"name\":\"changeWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"divisor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIMintyMultiToken\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"items\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"itemHash\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"unitPrice\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minty\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIMintyMultiToken\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"multiTokenOwners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIMintyMultiToken\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"numberOfOffers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIMintyMultiToken\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"ipfsString\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"}],\"name\":\"offerNew\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIMintyMultiToken\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"ipfsStrings\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"quantities\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"}],\"name\":\"offerNewBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIMintyMultiToken\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"offerResale\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIMintyMultiToken\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"offerId\",\"type\":\"uint256\"}],\"name\":\"price\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIMintyMultiToken\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pos\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"reSubmitOffer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIMintyMultiToken\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pos\",\"type\":\"uint256\"}],\"name\":\"retractRemainingOffer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_divisor\",\"type\":\"uint256\"}],\"name\":\"updateShares\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"weth\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// PMintyMultiSaleFuncSigs maps the 4-byte function signature to its string representation.
var PMintyMultiSaleFuncSigs = map[string]string{
	"6dd043d7": "acceptOffer(address,uint256,uint256,uint256)",
	"eac54d98": "available(address,uint256,uint256)",
	"98b9a2dc": "changeWallet(address)",
	"1f2dc5ef": "divisor()",
	"6d873c5d": "items(address,uint256,uint256)",
	"cbfcab1d": "minty()",
	"35b4511a": "multiTokenOwners(address)",
	"a3f7a5f8": "numberOfOffers(address,uint256)",
	"8bc844c7": "offerNew(address,uint256,string,uint256,uint256,uint256)",
	"5af9c479": "offerNewBatch(address,uint256[],string[],uint256[],uint256[],uint256)",
	"7d5c7217": "offerResale(address,uint256,uint256,uint256)",
	"8da5cb5b": "owner()",
	"3a06daa2": "price(address,uint256,uint256)",
	"4e900c02": "reSubmitOffer(address,uint256,uint256,uint256)",
	"ecf13ef4": "retractRemainingOffer(address,uint256,uint256)",
	"f2fde38b": "transferOwnership(address)",
	"b6608467": "updateShares(uint256)",
	"3fc8cef3": "weth()",
}

// PMintyMultiSaleBin is the compiled bytecode used for deploying new contracts.
var PMintyMultiSaleBin = "0x6080604052600280546001600160a01b031916331790553480156200002357600080fd5b5060405162002fcb38038062002fcb8339810160408190526200004691620000ec565b6103e8811015620000745760405162461bcd60e51b81526004016200006b9062000133565b60405180910390fd5b600580546001600160a01b038086166001600160a01b031992831617909255600380549285169290911691909117905560048190556040517f8c4d35e54a3f2ef1134138fd8ea3daee6a3c89e10d2665996babdf70261e2c7690620000db9083906200016a565b60405180910390a15050506200018c565b60008060006060848603121562000101578283fd5b83516200010e8162000173565b6020850151909350620001218162000173565b80925050604084015190509250925092565b60208082526017908201527f44697669736f72206d757374206265203e3d2031303030000000000000000000604082015260600190565b90815260200190565b6001600160a01b03811681146200018957600080fd5b50565b612e2f806200019c6000396000f3fe608060405234801561001057600080fd5b50600436106101165760003560e01c80638bc844c7116100a2578063b660846711610071578063b660846714610226578063cbfcab1d14610239578063eac54d9814610241578063ecf13ef414610254578063f2fde38b1461026757610116565b80638bc844c7146101e55780638da5cb5b146101f857806398b9a2dc14610200578063a3f7a5f81461021357610116565b80634e900c02116100e95780634e900c02146101745780635af9c479146101895780636d873c5d1461019c5780636dd043d7146101bf5780637d5c7217146101d257610116565b80631f2dc5ef1461011b57806335b4511a146101395780633a06daa2146101595780633fc8cef31461016c575b600080fd5b61012361027a565b6040516101309190612d17565b60405180910390f35b61014c61014736600461235a565b610280565b6040516101309190612759565b6101236101673660046125e2565b61029b565b61014c61048d565b610187610182366004612616565b61049c565b005b610187610197366004612483565b6106fd565b6101af6101aa3660046125e2565b610b75565b604051610130949392919061285e565b6101876101cd366004612616565b610c63565b6101876101e0366004612616565b61109f565b6101876101f336600461256f565b61133d565b61014c611695565b61018761020e36600461235a565b6116a4565b610123610221366004612544565b611750565b6101876102343660046126c3565b611779565b61014c611805565b61012361024f3660046125e2565b611814565b6101876102623660046125e2565b611a8a565b61018761027536600461235a565b611ce6565b60045481565b6001602052600090815260409020546001600160a01b031681565b6001600160a01b03831660009081526020818152604080832085845290915281205482106102e45760405162461bcd60e51b81526004016102db90612a01565b60405180910390fd5b6102ec612164565b6001600160a01b038516600090815260208181526040808320878452909152902080548490811061031957fe5b60009182526020918290206040805160808101825260049390930290910180546001600160a01b03168352600180820154848601526002808301805485516101009482161594909402600019011691909104601f81018790048702830187018552808352949592949386019391929091908301828280156103db5780601f106103b0576101008083540402835291602001916103db565b820191906000526020600020905b8154815290600101906020018083116103be57829003601f168201915b505050918352505060039190910154602090910152805160405163e985e9c560e01b81529192506001600160a01b0387169163e985e9c59161042191309060040161276d565b60206040518083038186803b15801561043957600080fd5b505afa15801561044d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104719190612463565b61047f576000915050610486565b6060015190505b9392505050565b6003546001600160a01b031681565b6001600160a01b03841660009081526020818152604080832086845290915290205482106104dc5760405162461bcd60e51b81526004016102db90612be1565b6104e4612164565b6001600160a01b038516600090815260208181526040808320878452909152902080548490811061051157fe5b60009182526020918290206040805160808101825260049390930290910180546001600160a01b03168352600180820154848601526002808301805485516101009482161594909402600019011691909104601f81018790048702830187018552808352949592949386019391929091908301828280156105d35780601f106105a8576101008083540402835291602001916105d3565b820191906000526020600020905b8154815290600101906020018083116105b657829003601f168201915b50505050508152602001600382015481525050905080600001516001600160a01b0316336001600160a01b03161461061d5760405162461bcd60e51b81526004016102db90612a5a565b606081018290526001600160a01b038516600090815260208181526040808320878452909152902080548291908590811061065457fe5b600091825260209182902083516004929092020180546001600160a01b0319166001600160a01b039092169190911781558282015160018201556040830151805191926106a992600285019290910190612195565b50606082015181600301559050507f0365ed613ce46864c2596157898bfc904a905c87c70c265de1473f51eedb1b7785858585336040516106ee9594939291906129d1565b60405180910390a15050505050565b60405163e985e9c560e01b81526001600160a01b0387169063e985e9c59061072b903390309060040161276d565b60206040518083038186803b15801561074357600080fd5b505afa158015610757573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061077b9190612463565b6107975760405162461bcd60e51b81526004016102db90612c77565b336001600160a01b0316866001600160a01b0316638da5cb5b6040518163ffffffff1660e01b815260040160206040518083038186803b1580156107da57600080fd5b505afa1580156107ee573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108129190612376565b6001600160a01b0316146108385760405162461bcd60e51b81526004016102db90612b84565b60005b8551811015610b0857866001600160a01b0316637dc0bf3f87838151811061085f57fe5b60200260200101516040518263ffffffff1660e01b81526004016108839190612d17565b60206040518083038186803b15801561089b57600080fd5b505afa1580156108af573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108d39190612463565b156108f05760405162461bcd60e51b81526004016102db90612b0c565b600080886001600160a01b03166001600160a01b03168152602001908152602001600020600087838151811061092257fe5b602002602001015181526020019081526020016000208054905060001461095b5760405162461bcd60e51b81526004016102db90612ab9565b600080886001600160a01b03166001600160a01b03168152602001908152602001600020600087838151811061098d57fe5b602002602001015181526020019081526020016000206040518060800160405280336001600160a01b031681526020018684815181106109c957fe5b602002602001015181526020018784815181106109e257fe5b602002602001015181526020018584815181106109fb57fe5b602090810291909101810151909152825460018082018555600094855293829020835160049092020180546001600160a01b0319166001600160a01b03909216919091178155828201519381019390935560408201518051929392610a669260028501920190612195565b506060820151816003015550507f10575dae8255319b2807b08ada1038743fd34f2a06f1fba49e9407e7aa91f8f687878381518110610aa157fe5b602002602001015133878581518110610ab657fe5b6020026020010151878681518110610aca57fe5b60200260200101518a8781518110610ade57fe5b6020026020010151604051610af896959493929190612923565b60405180910390a160010161083b565b50604051630a98df5760e21b81526001600160a01b03871690632a637d5c90610b3b908890879089908790600401612896565b600060405180830381600087803b158015610b5557600080fd5b505af1158015610b69573d6000803e3d6000fd5b50505050505050505050565b60006020528260005260406000206020528160005260406000208181548110610b9d57600080fd5b6000918252602091829020600491909102018054600180830154600280850180546040805161010096831615969096026000190190911692909204601f81018890048802850188019092528184526001600160a01b03909416985090965092945092909190830182828015610c535780601f10610c2857610100808354040283529160200191610c53565b820191906000526020600020905b815481529060010190602001808311610c3657829003601f168201915b5050505050908060030154905084565b600354600160a01b900460ff1615610c8d5760405162461bcd60e51b81526004016102db90612a2c565b81610cf157604051635773bd0760e01b81526001600160a01b03851690635773bd0790610cbe903390600401612759565b600060405180830381600087803b158015610cd857600080fd5b505af1158015610cec573d6000803e3d6000fd5b505050505b6003805460ff60a01b1916600160a01b1790556060610d0e612164565b6001600160a01b0386166000908152602081815260408083208884529091529020805485908110610d3b57fe5b60009182526020918290206040805160808101825260049390930290910180546001600160a01b03168352600180820154848601526002808301805485516101009482161594909402600019011691909104601f8101879004870283018701855280835294959294938601939192909190830182828015610dfd5780601f10610dd257610100808354040283529160200191610dfd565b820191906000526020600020905b815481529060010190602001808311610de057829003601f168201915b5050505050815260200160038201548152505090506000816000015190508382602001511015610e3f5760405162461bcd60e51b81526004016102db90612a82565b6000610e4f836060015186611d92565b905084886001600160a01b031662fdd58e848a6040518363ffffffff1660e01b8152600401610e7f9291906127e1565b60206040518083038186803b158015610e9757600080fd5b505afa158015610eab573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ecf91906126db565b1015610eed5760405162461bcd60e51b81526004016102db90612b43565b604051637921219560e11b81526001600160a01b0389169063f242432a90610f2190859033908c908b908b906004016127fa565b600060405180830381600087803b158015610f3b57600080fd5b505af1158015610f4f573d6000803e3d6000fd5b5050506020808501805188900390526001600160a01b038a16600090815280825260408082208b835290925220805485925088908110610f8b57fe5b600091825260209182902083516004929092020180546001600160a01b0319166001600160a01b03909216919091178155828201516001820155604083015180519192610fe092600285019290910190612195565b506060919091015160039091015560055483516040517f0913aa4718e76655dbdbde41862a88aba8459d8721849484c97f53af3b5028e692611031926001600160a01b039091169186908690612834565b60405180910390a16110468888888585611db7565b6003805460ff60a01b191690556040517faf47c9ab14ff5ce679b4cb416be29684283e7b30fcfe405bfdebcc135f4b1e6f9061108d9033908b908b908b908b9088906127ab565b60405180910390a15050505050505050565b604051627eeac760e11b81526001600160a01b0385169062fdd58e906110cb90339087906004016127e1565b60206040518083038186803b1580156110e357600080fd5b505afa1580156110f7573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061111b91906126db565b821115801561112a5750600082115b6111465760405162461bcd60e51b81526004016102db90612cd4565b60405163e985e9c560e01b81526001600160a01b0385169063e985e9c590611174903390309060040161276d565b60206040518083038186803b15801561118c57600080fd5b505afa1580156111a0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111c49190612463565b6111e05760405162461bcd60e51b81526004016102db90612c77565b6001600160a01b0384166000818152602081815260408083208784528252918290208054835160808101855233815292830187905283516303a24d0760e21b81529094919383019190630e89341c9061123d908a90600401612d17565b60006040518083038186803b15801561125557600080fd5b505afa158015611269573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526112919190810190612650565b81526020908101859052825460018082018555600094855293829020835160049092020180546001600160a01b0319166001600160a01b039092169190911781558282015193810193909355604082015180519293926112f79260028501920190612195565b506060820151816003015550507f0df75d277ebbb180db8152f8c392e8bf6ba2720b35f9f98456eca3557b91182e8585853386866040516106ee96959493929190612997565b60405163e985e9c560e01b81526001600160a01b0387169063e985e9c59061136b903390309060040161276d565b60206040518083038186803b15801561138357600080fd5b505afa158015611397573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113bb9190612463565b6113d75760405162461bcd60e51b81526004016102db90612c77565b604051637dc0bf3f60e01b81526001600160a01b03871690637dc0bf3f90611403908890600401612d17565b60206040518083038186803b15801561141b57600080fd5b505afa15801561142f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114539190612463565b156114705760405162461bcd60e51b81526004016102db90612b0c565b336001600160a01b0316866001600160a01b0316638da5cb5b6040518163ffffffff1660e01b815260040160206040518083038186803b1580156114b357600080fd5b505afa1580156114c7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114eb9190612376565b6001600160a01b0316146115115760405162461bcd60e51b81526004016102db90612b84565b6001600160a01b038616600090815260208181526040808320888452909152902054156115505760405162461bcd60e51b81526004016102db90612ab9565b604051636797699760e11b81526001600160a01b0387169063cf2ed32e90611582908890879089908790600401612d2e565b600060405180830381600087803b15801561159c57600080fd5b505af11580156115b0573d6000803e3d6000fd5b505050506001600160a01b03868116600090815260208181526040808320898452825280832081516080810183523381528084018981529281018a81526060820189905282546001808201855593875295859020825160049097020180546001600160a01b031916969097169590951786559151908501559151805192939261163f9260028501920190612195565b506060820151816003015550507f10575dae8255319b2807b08ada1038743fd34f2a06f1fba49e9407e7aa91f8f686863386868960405161168596959493929190612923565b60405180910390a1505050505050565b6002546001600160a01b031681565b6002546001600160a01b031633146116ce5760405162461bcd60e51b81526004016102db90612ae6565b6001600160a01b0381166116f45760405162461bcd60e51b81526004016102db90612baa565b6005546040516001600160a01b038084169216907ff404d79dbd5d19bc2bf99a74f0e5affb6502b3d6012140f1d908ce3fcd0716d890600090a3600580546001600160a01b0319166001600160a01b0392909216919091179055565b6001600160a01b0382166000908152602081815260408083208484529091529020545b92915050565b6002546001600160a01b031633146117a35760405162461bcd60e51b81526004016102db90612ae6565b6103e88110156117c55760405162461bcd60e51b81526004016102db90612c11565b60048190556040517f8c4d35e54a3f2ef1134138fd8ea3daee6a3c89e10d2665996babdf70261e2c76906117fa908390612d17565b60405180910390a150565b6005546001600160a01b031681565b6001600160a01b03831660009081526020818152604080832085845290915281205482106118545760405162461bcd60e51b81526004016102db90612a01565b61185c612164565b6001600160a01b038516600090815260208181526040808320878452909152902080548490811061188957fe5b60009182526020918290206040805160808101825260049390930290910180546001600160a01b03168352600180820154848601526002808301805485516101009482161594909402600019011691909104601f810187900487028301870185528083529495929493860193919290919083018282801561194b5780601f106119205761010080835404028352916020019161194b565b820191906000526020600020905b81548152906001019060200180831161192e57829003601f168201915b505050918352505060039190910154602090910152805160405163e985e9c560e01b81529192506001600160a01b0387169163e985e9c59161199191309060040161276d565b60206040518083038186803b1580156119a957600080fd5b505afa1580156119bd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119e19190612463565b6119ef576000915050610486565b60208101518151604051627eeac760e11b81526000916001600160a01b0389169162fdd58e91611a23918a906004016127e1565b60206040518083038186803b158015611a3b57600080fd5b505afa158015611a4f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a7391906126db565b9050611a7f828261214d565b979650505050505050565b6001600160a01b0383166000908152602081815260408083208584529091529020548110611aca5760405162461bcd60e51b81526004016102db90612be1565b611ad2612164565b6001600160a01b0384166000908152602081815260408083208684529091529020805483908110611aff57fe5b60009182526020918290206040805160808101825260049390930290910180546001600160a01b03168352600180820154848601526002808301805485516101009482161594909402600019011691909104601f8101879004870283018701855280835294959294938601939192909190830182828015611bc15780601f10611b9657610100808354040283529160200191611bc1565b820191906000526020600020905b815481529060010190602001808311611ba457829003601f168201915b50505050508152602001600382015481525050905080600001516001600160a01b0316336001600160a01b031614611c0b5760405162461bcd60e51b81526004016102db90612a5a565b600060208083018290526001600160a01b038616825281815260408083208684529091529020805482919084908110611c4057fe5b600091825260209182902083516004929092020180546001600160a01b0319166001600160a01b03909216919091178155828201516001820155604083015180519192611c9592600285019290910190612195565b50606082015181600301559050507f74ee16c30acdc0986deebeda1b36d871a7ebf1dc7e2fd327dc7628617c37726d84848433604051611cd8949392919061296c565b60405180910390a150505050565b6002546001600160a01b03163314611d105760405162461bcd60e51b81526004016102db90612ae6565b6001600160a01b038116611d365760405162461bcd60e51b81526004016102db90612baa565b6002546040516001600160a01b038084169216907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3600280546001600160a01b0319166001600160a01b0392909216919091179055565b600082611da157506000611773565b5081810281838281611daf57fe5b041461177357fe5b6000856001600160a01b031663e8f3359f6040518163ffffffff1660e01b815260040160206040518083038186803b158015611df257600080fd5b505afa158015611e06573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e2a91906126db565b90506000600454611e3b8484611d92565b81611e4257fe5b0490506000600454611e5885856103e803611d92565b81611e5f57fe5b6003546040516323b872dd60e01b81529290910492506001600160a01b0316906323b872dd90611e9790339089908690600401612787565b602060405180830381600087803b158015611eb157600080fd5b505af1158015611ec5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611ee99190612463565b611f055760405162461bcd60e51b81526004016102db90612c48565b60405163e03650b960e01b815281906060906001600160a01b038b169063e03650b990611f38908b908d90600401612d20565b60006040518083038186803b158015611f5057600080fd5b505afa158015611f64573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611f8c9190810190612392565b905060005b81518110156120965760006103e8611fc087858581518110611faf57fe5b602002602001015160200151611d92565b81611fc757fe5b60035485519290910492506001600160a01b0316906323b872dd903390869086908110611ff057fe5b602002602001015160000151846040518463ffffffff1660e01b815260040161201b93929190612787565b602060405180830381600087803b15801561203557600080fd5b505af1158015612049573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061206d9190612463565b6120895760405162461bcd60e51b81526004016102db90612c48565b9290920191600101611f91565b506003546005546040516323b872dd60e01b8152848903926001600160a01b03908116926323b872dd926120d292339216908690600401612787565b602060405180830381600087803b1580156120ec57600080fd5b505af1158015612100573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906121249190612463565b6121405760405162461bcd60e51b81526004016102db90612c48565b5050505050505050505050565b60008183101561215e575081611773565b50919050565b604051806080016040528060006001600160a01b031681526020016000815260200160608152602001600081525090565b828054600181600116156101000203166002900490600052602060002090601f0160209004810192826121cb5760008555612211565b82601f106121e457805160ff1916838001178555612211565b82800160010185558215612211579182015b828111156122115782518255916020019190600101906121f6565b5061221d929150612221565b5090565b5b8082111561221d5760008155600101612222565b600082601f830112612246578081fd5b813561225961225482612d71565b612d4d565b818152915060208083019084810160005b8481101561229357612281888484358a010161230c565b8452928201929082019060010161226a565b505050505092915050565b600082601f8301126122ae578081fd5b81356122bc61225482612d71565b8181529150602080830190848101818402860182018710156122dd57600080fd5b60005b84811015612293578135845292820192908201906001016122e0565b803561230781612de1565b919050565b600082601f83011261231c578081fd5b813561232a61225482612d8f565b915080825283602082850101111561234157600080fd5b8060208401602084013760009082016020015292915050565b60006020828403121561236b578081fd5b813561048681612de1565b600060208284031215612387578081fd5b815161048681612de1565b600060208083850312156123a4578182fd5b825167ffffffffffffffff808211156123bb578384fd5b818501915085601f8301126123ce578384fd5b81516123dc61225482612d71565b818152848101908486016040808502870188018b10156123fa578889fd5b8896505b848710156124545780828c031215612414578889fd5b8051818101818110888211171561242757fe5b8252825161243481612de1565b8152828901518982015284526001969096019592870192908101906123fe565b50909998505050505050505050565b600060208284031215612474578081fd5b81518015158114610486578182fd5b60008060008060008060c0878903121561249b578182fd5b6124a4876122fc565b9550602087013567ffffffffffffffff808211156124c0578384fd5b6124cc8a838b0161229e565b965060408901359150808211156124e1578384fd5b6124ed8a838b01612236565b95506060890135915080821115612502578384fd5b61250e8a838b0161229e565b94506080890135915080821115612523578384fd5b5061253089828a0161229e565b92505060a087013590509295509295509295565b60008060408385031215612556578182fd5b823561256181612de1565b946020939093013593505050565b60008060008060008060c08789031215612587578182fd5b863561259281612de1565b955060208701359450604087013567ffffffffffffffff8111156125b4578283fd5b6125c089828a0161230c565b945050606087013592506080870135915060a087013590509295509295509295565b6000806000606084860312156125f6578081fd5b833561260181612de1565b95602085013595506040909401359392505050565b6000806000806080858703121561262b578182fd5b843561263681612de1565b966020860135965060408601359560600135945092505050565b600060208284031215612661578081fd5b815167ffffffffffffffff811115612677578182fd5b8201601f81018413612687578182fd5b805161269561225482612d8f565b8181528560208385010111156126a9578384fd5b6126ba826020830160208601612db1565b95945050505050565b6000602082840312156126d4578081fd5b5035919050565b6000602082840312156126ec578081fd5b5051919050565b6000815180845260208085019450808401835b8381101561272257815187529582019590820190600101612706565b509495945050505050565b60008151808452612745816020860160208601612db1565b601f01601f19169290920160200192915050565b6001600160a01b0391909116815260200190565b6001600160a01b0392831681529116602082015260400190565b6001600160a01b039384168152919092166020820152604081019190915260600190565b6001600160a01b03968716815294909516602085015260408401929092526060830152608082015260a081019190915260c00190565b6001600160a01b03929092168252602082015260400190565b6001600160a01b03868116825285166020820152604081018490526060810183905260a060808201819052600090611a7f9083018461272d565b6001600160a01b039485168152928416602084015292166040820152606081019190915260800190565b600060018060a01b038616825284602083015260806040830152612885608083018561272d565b905082606083015295945050505050565b6000608082526128a960808301876126f3565b6020838203818501526128bc82886126f3565b848103604086015286518082529092508183019082810284018301838901865b8381101561290a57601f198784030185526128f883835161272d565b948601949250908501906001016128dc565b5050809550505050505082606083015295945050505050565b600060018060a01b03808916835287602084015280871660408401525084606083015283608083015260c060a083015261296060c083018461272d565b98975050505050505050565b6001600160a01b03948516815260208101939093526040830191909152909116606082015260800190565b6001600160a01b0396871681526020810195909552604085019390935293166060830152608082019290925260a081019190915260c00190565b6001600160a01b039586168152602081019490945260408401929092526060830152909116608082015260a00190565b60208082526011908201527013d999995c9251081b9bdd081d985b1a59607a1b604082015260600190565b6020808252601490820152734e6f207265656e7472616e637920706c6561736560601b604082015260600190565b6020808252600e908201526d3737ba103cb7bab91037b33332b960911b604082015260600190565b6020808252601a908201527f6e6f7420656e6f756768206974656d7320617661696c61626c65000000000000604082015260600190565b602080825260139082015272556e61626c6520746f206f66666572206e657760681b604082015260600190565b6020808252600c908201526b1d5b985d5d1a1bdc9a5cd95960a21b604082015260600190565b60208082526017908201527f546f6b656e20494420616c7265616479206d696e746564000000000000000000604082015260600190565b60208082526021908201527f6e6f7420656e6f756768206974656d73206f776e6564206279206f66666572656040820152603960f91b606082015260800190565b6020808252600c908201526b155b985d5d1a1bdc9a5cd95960a21b604082015260600190565b6020808252601a908201527f446f206e6f742073657420746f2061646472657373207a65726f000000000000604082015260600190565b60208082526016908201527534b73b30b634b21037b33332b9103837b9b4ba34b7b760511b604082015260600190565b60208082526017908201527f44697669736f72206d757374206265203e3d2031303030000000000000000000604082015260600190565b60208082526015908201527463616e6e6f74207472616e736665722066756e647360581b604082015260600190565b60208082526037908201527f596f752068617665206e6f7420617070726f766564207468697320636f6e747260408201527f61637420746f2073656c6c20796f757220746f6b656e73000000000000000000606082015260800190565b60208082526023908201527f596f7520646f206e6f74206f776e20656e6f756768206f66207468697320746f60408201526235b2b760e91b606082015260800190565b90815260200190565b918252602082015260400190565b600085825284602083015260806040830152612885608083018561272d565b60405181810167ffffffffffffffff81118282101715612d6957fe5b604052919050565b600067ffffffffffffffff821115612d8557fe5b5060209081020190565b600067ffffffffffffffff821115612da357fe5b50601f01601f191660200190565b60005b83811015612dcc578181015183820152602001612db4565b83811115612ddb576000848401525b50505050565b6001600160a01b0381168114612df657600080fd5b5056fea264697066735822122068fd956f85227073b9c90c8e2b33df4945e56b897b65e0cfc0f3a674ad432c3064736f6c63430007050033"

// DeployPMintyMultiSale deploys a new Ethereum contract, binding an instance of PMintyMultiSale to it.
func DeployPMintyMultiSale(auth *bind.TransactOpts, backend bind.ContractBackend, wallet common.Address, _weth common.Address, _divisor *big.Int) (common.Address, *types.Transaction, *PMintyMultiSale, error) {
	parsed, err := abi.JSON(strings.NewReader(PMintyMultiSaleABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PMintyMultiSaleBin), backend, wallet, _weth, _divisor)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PMintyMultiSale{PMintyMultiSaleCaller: PMintyMultiSaleCaller{contract: contract}, PMintyMultiSaleTransactor: PMintyMultiSaleTransactor{contract: contract}, PMintyMultiSaleFilterer: PMintyMultiSaleFilterer{contract: contract}}, nil
}

// PMintyMultiSale is an auto generated Go binding around an Ethereum contract.
type PMintyMultiSale struct {
	PMintyMultiSaleCaller     // Read-only binding to the contract
	PMintyMultiSaleTransactor // Write-only binding to the contract
	PMintyMultiSaleFilterer   // Log filterer for contract events
}

// PMintyMultiSaleCaller is an auto generated read-only Go binding around an Ethereum contract.
type PMintyMultiSaleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PMintyMultiSaleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PMintyMultiSaleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PMintyMultiSaleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PMintyMultiSaleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PMintyMultiSaleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PMintyMultiSaleSession struct {
	Contract     *PMintyMultiSale  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PMintyMultiSaleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PMintyMultiSaleCallerSession struct {
	Contract *PMintyMultiSaleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// PMintyMultiSaleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PMintyMultiSaleTransactorSession struct {
	Contract     *PMintyMultiSaleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// PMintyMultiSaleRaw is an auto generated low-level Go binding around an Ethereum contract.
type PMintyMultiSaleRaw struct {
	Contract *PMintyMultiSale // Generic contract binding to access the raw methods on
}

// PMintyMultiSaleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PMintyMultiSaleCallerRaw struct {
	Contract *PMintyMultiSaleCaller // Generic read-only contract binding to access the raw methods on
}

// PMintyMultiSaleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PMintyMultiSaleTransactorRaw struct {
	Contract *PMintyMultiSaleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPMintyMultiSale creates a new instance of PMintyMultiSale, bound to a specific deployed contract.
func NewPMintyMultiSale(address common.Address, backend bind.ContractBackend) (*PMintyMultiSale, error) {
	contract, err := bindPMintyMultiSale(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PMintyMultiSale{PMintyMultiSaleCaller: PMintyMultiSaleCaller{contract: contract}, PMintyMultiSaleTransactor: PMintyMultiSaleTransactor{contract: contract}, PMintyMultiSaleFilterer: PMintyMultiSaleFilterer{contract: contract}}, nil
}

// NewPMintyMultiSaleCaller creates a new read-only instance of PMintyMultiSale, bound to a specific deployed contract.
func NewPMintyMultiSaleCaller(address common.Address, caller bind.ContractCaller) (*PMintyMultiSaleCaller, error) {
	contract, err := bindPMintyMultiSale(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PMintyMultiSaleCaller{contract: contract}, nil
}

// NewPMintyMultiSaleTransactor creates a new write-only instance of PMintyMultiSale, bound to a specific deployed contract.
func NewPMintyMultiSaleTransactor(address common.Address, transactor bind.ContractTransactor) (*PMintyMultiSaleTransactor, error) {
	contract, err := bindPMintyMultiSale(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PMintyMultiSaleTransactor{contract: contract}, nil
}

// NewPMintyMultiSaleFilterer creates a new log filterer instance of PMintyMultiSale, bound to a specific deployed contract.
func NewPMintyMultiSaleFilterer(address common.Address, filterer bind.ContractFilterer) (*PMintyMultiSaleFilterer, error) {
	contract, err := bindPMintyMultiSale(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PMintyMultiSaleFilterer{contract: contract}, nil
}

// bindPMintyMultiSale binds a generic wrapper to an already deployed contract.
func bindPMintyMultiSale(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PMintyMultiSaleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PMintyMultiSale *PMintyMultiSaleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PMintyMultiSale.Contract.PMintyMultiSaleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PMintyMultiSale *PMintyMultiSaleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PMintyMultiSale.Contract.PMintyMultiSaleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PMintyMultiSale *PMintyMultiSaleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PMintyMultiSale.Contract.PMintyMultiSaleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PMintyMultiSale *PMintyMultiSaleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PMintyMultiSale.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PMintyMultiSale *PMintyMultiSaleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PMintyMultiSale.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PMintyMultiSale *PMintyMultiSaleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PMintyMultiSale.Contract.contract.Transact(opts, method, params...)
}

// Available is a free data retrieval call binding the contract method 0xeac54d98.
//
// Solidity: function available(address token, uint256 tokenId, uint256 offerId) view returns(uint256)
func (_PMintyMultiSale *PMintyMultiSaleCaller) Available(opts *bind.CallOpts, token common.Address, tokenId *big.Int, offerId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PMintyMultiSale.contract.Call(opts, &out, "available", token, tokenId, offerId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Available is a free data retrieval call binding the contract method 0xeac54d98.
//
// Solidity: function available(address token, uint256 tokenId, uint256 offerId) view returns(uint256)
func (_PMintyMultiSale *PMintyMultiSaleSession) Available(token common.Address, tokenId *big.Int, offerId *big.Int) (*big.Int, error) {
	return _PMintyMultiSale.Contract.Available(&_PMintyMultiSale.CallOpts, token, tokenId, offerId)
}

// Available is a free data retrieval call binding the contract method 0xeac54d98.
//
// Solidity: function available(address token, uint256 tokenId, uint256 offerId) view returns(uint256)
func (_PMintyMultiSale *PMintyMultiSaleCallerSession) Available(token common.Address, tokenId *big.Int, offerId *big.Int) (*big.Int, error) {
	return _PMintyMultiSale.Contract.Available(&_PMintyMultiSale.CallOpts, token, tokenId, offerId)
}

// Divisor is a free data retrieval call binding the contract method 0x1f2dc5ef.
//
// Solidity: function divisor() view returns(uint256)
func (_PMintyMultiSale *PMintyMultiSaleCaller) Divisor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PMintyMultiSale.contract.Call(opts, &out, "divisor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Divisor is a free data retrieval call binding the contract method 0x1f2dc5ef.
//
// Solidity: function divisor() view returns(uint256)
func (_PMintyMultiSale *PMintyMultiSaleSession) Divisor() (*big.Int, error) {
	return _PMintyMultiSale.Contract.Divisor(&_PMintyMultiSale.CallOpts)
}

// Divisor is a free data retrieval call binding the contract method 0x1f2dc5ef.
//
// Solidity: function divisor() view returns(uint256)
func (_PMintyMultiSale *PMintyMultiSaleCallerSession) Divisor() (*big.Int, error) {
	return _PMintyMultiSale.Contract.Divisor(&_PMintyMultiSale.CallOpts)
}

// Items is a free data retrieval call binding the contract method 0x6d873c5d.
//
// Solidity: function items(address , uint256 , uint256 ) view returns(address creator, uint256 quantity, string itemHash, uint256 unitPrice)
func (_PMintyMultiSale *PMintyMultiSaleCaller) Items(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int, arg2 *big.Int) (struct {
	Creator   common.Address
	Quantity  *big.Int
	ItemHash  string
	UnitPrice *big.Int
}, error) {
	var out []interface{}
	err := _PMintyMultiSale.contract.Call(opts, &out, "items", arg0, arg1, arg2)

	outstruct := new(struct {
		Creator   common.Address
		Quantity  *big.Int
		ItemHash  string
		UnitPrice *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Creator = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Quantity = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ItemHash = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.UnitPrice = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Items is a free data retrieval call binding the contract method 0x6d873c5d.
//
// Solidity: function items(address , uint256 , uint256 ) view returns(address creator, uint256 quantity, string itemHash, uint256 unitPrice)
func (_PMintyMultiSale *PMintyMultiSaleSession) Items(arg0 common.Address, arg1 *big.Int, arg2 *big.Int) (struct {
	Creator   common.Address
	Quantity  *big.Int
	ItemHash  string
	UnitPrice *big.Int
}, error) {
	return _PMintyMultiSale.Contract.Items(&_PMintyMultiSale.CallOpts, arg0, arg1, arg2)
}

// Items is a free data retrieval call binding the contract method 0x6d873c5d.
//
// Solidity: function items(address , uint256 , uint256 ) view returns(address creator, uint256 quantity, string itemHash, uint256 unitPrice)
func (_PMintyMultiSale *PMintyMultiSaleCallerSession) Items(arg0 common.Address, arg1 *big.Int, arg2 *big.Int) (struct {
	Creator   common.Address
	Quantity  *big.Int
	ItemHash  string
	UnitPrice *big.Int
}, error) {
	return _PMintyMultiSale.Contract.Items(&_PMintyMultiSale.CallOpts, arg0, arg1, arg2)
}

// Minty is a free data retrieval call binding the contract method 0xcbfcab1d.
//
// Solidity: function minty() view returns(address)
func (_PMintyMultiSale *PMintyMultiSaleCaller) Minty(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PMintyMultiSale.contract.Call(opts, &out, "minty")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Minty is a free data retrieval call binding the contract method 0xcbfcab1d.
//
// Solidity: function minty() view returns(address)
func (_PMintyMultiSale *PMintyMultiSaleSession) Minty() (common.Address, error) {
	return _PMintyMultiSale.Contract.Minty(&_PMintyMultiSale.CallOpts)
}

// Minty is a free data retrieval call binding the contract method 0xcbfcab1d.
//
// Solidity: function minty() view returns(address)
func (_PMintyMultiSale *PMintyMultiSaleCallerSession) Minty() (common.Address, error) {
	return _PMintyMultiSale.Contract.Minty(&_PMintyMultiSale.CallOpts)
}

// MultiTokenOwners is a free data retrieval call binding the contract method 0x35b4511a.
//
// Solidity: function multiTokenOwners(address ) view returns(address)
func (_PMintyMultiSale *PMintyMultiSaleCaller) MultiTokenOwners(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _PMintyMultiSale.contract.Call(opts, &out, "multiTokenOwners", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MultiTokenOwners is a free data retrieval call binding the contract method 0x35b4511a.
//
// Solidity: function multiTokenOwners(address ) view returns(address)
func (_PMintyMultiSale *PMintyMultiSaleSession) MultiTokenOwners(arg0 common.Address) (common.Address, error) {
	return _PMintyMultiSale.Contract.MultiTokenOwners(&_PMintyMultiSale.CallOpts, arg0)
}

// MultiTokenOwners is a free data retrieval call binding the contract method 0x35b4511a.
//
// Solidity: function multiTokenOwners(address ) view returns(address)
func (_PMintyMultiSale *PMintyMultiSaleCallerSession) MultiTokenOwners(arg0 common.Address) (common.Address, error) {
	return _PMintyMultiSale.Contract.MultiTokenOwners(&_PMintyMultiSale.CallOpts, arg0)
}

// NumberOfOffers is a free data retrieval call binding the contract method 0xa3f7a5f8.
//
// Solidity: function numberOfOffers(address token, uint256 tokenId) view returns(uint256)
func (_PMintyMultiSale *PMintyMultiSaleCaller) NumberOfOffers(opts *bind.CallOpts, token common.Address, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PMintyMultiSale.contract.Call(opts, &out, "numberOfOffers", token, tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumberOfOffers is a free data retrieval call binding the contract method 0xa3f7a5f8.
//
// Solidity: function numberOfOffers(address token, uint256 tokenId) view returns(uint256)
func (_PMintyMultiSale *PMintyMultiSaleSession) NumberOfOffers(token common.Address, tokenId *big.Int) (*big.Int, error) {
	return _PMintyMultiSale.Contract.NumberOfOffers(&_PMintyMultiSale.CallOpts, token, tokenId)
}

// NumberOfOffers is a free data retrieval call binding the contract method 0xa3f7a5f8.
//
// Solidity: function numberOfOffers(address token, uint256 tokenId) view returns(uint256)
func (_PMintyMultiSale *PMintyMultiSaleCallerSession) NumberOfOffers(token common.Address, tokenId *big.Int) (*big.Int, error) {
	return _PMintyMultiSale.Contract.NumberOfOffers(&_PMintyMultiSale.CallOpts, token, tokenId)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PMintyMultiSale *PMintyMultiSaleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PMintyMultiSale.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PMintyMultiSale *PMintyMultiSaleSession) Owner() (common.Address, error) {
	return _PMintyMultiSale.Contract.Owner(&_PMintyMultiSale.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PMintyMultiSale *PMintyMultiSaleCallerSession) Owner() (common.Address, error) {
	return _PMintyMultiSale.Contract.Owner(&_PMintyMultiSale.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0x3a06daa2.
//
// Solidity: function price(address token, uint256 tokenId, uint256 offerId) view returns(uint256)
func (_PMintyMultiSale *PMintyMultiSaleCaller) Price(opts *bind.CallOpts, token common.Address, tokenId *big.Int, offerId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PMintyMultiSale.contract.Call(opts, &out, "price", token, tokenId, offerId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price is a free data retrieval call binding the contract method 0x3a06daa2.
//
// Solidity: function price(address token, uint256 tokenId, uint256 offerId) view returns(uint256)
func (_PMintyMultiSale *PMintyMultiSaleSession) Price(token common.Address, tokenId *big.Int, offerId *big.Int) (*big.Int, error) {
	return _PMintyMultiSale.Contract.Price(&_PMintyMultiSale.CallOpts, token, tokenId, offerId)
}

// Price is a free data retrieval call binding the contract method 0x3a06daa2.
//
// Solidity: function price(address token, uint256 tokenId, uint256 offerId) view returns(uint256)
func (_PMintyMultiSale *PMintyMultiSaleCallerSession) Price(token common.Address, tokenId *big.Int, offerId *big.Int) (*big.Int, error) {
	return _PMintyMultiSale.Contract.Price(&_PMintyMultiSale.CallOpts, token, tokenId, offerId)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_PMintyMultiSale *PMintyMultiSaleCaller) Weth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PMintyMultiSale.contract.Call(opts, &out, "weth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_PMintyMultiSale *PMintyMultiSaleSession) Weth() (common.Address, error) {
	return _PMintyMultiSale.Contract.Weth(&_PMintyMultiSale.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_PMintyMultiSale *PMintyMultiSaleCallerSession) Weth() (common.Address, error) {
	return _PMintyMultiSale.Contract.Weth(&_PMintyMultiSale.CallOpts)
}

// AcceptOffer is a paid mutator transaction binding the contract method 0x6dd043d7.
//
// Solidity: function acceptOffer(address token, uint256 tokenId, uint256 pos, uint256 quantity) returns()
func (_PMintyMultiSale *PMintyMultiSaleTransactor) AcceptOffer(opts *bind.TransactOpts, token common.Address, tokenId *big.Int, pos *big.Int, quantity *big.Int) (*types.Transaction, error) {
	return _PMintyMultiSale.contract.Transact(opts, "acceptOffer", token, tokenId, pos, quantity)
}

// AcceptOffer is a paid mutator transaction binding the contract method 0x6dd043d7.
//
// Solidity: function acceptOffer(address token, uint256 tokenId, uint256 pos, uint256 quantity) returns()
func (_PMintyMultiSale *PMintyMultiSaleSession) AcceptOffer(token common.Address, tokenId *big.Int, pos *big.Int, quantity *big.Int) (*types.Transaction, error) {
	return _PMintyMultiSale.Contract.AcceptOffer(&_PMintyMultiSale.TransactOpts, token, tokenId, pos, quantity)
}

// AcceptOffer is a paid mutator transaction binding the contract method 0x6dd043d7.
//
// Solidity: function acceptOffer(address token, uint256 tokenId, uint256 pos, uint256 quantity) returns()
func (_PMintyMultiSale *PMintyMultiSaleTransactorSession) AcceptOffer(token common.Address, tokenId *big.Int, pos *big.Int, quantity *big.Int) (*types.Transaction, error) {
	return _PMintyMultiSale.Contract.AcceptOffer(&_PMintyMultiSale.TransactOpts, token, tokenId, pos, quantity)
}

// ChangeWallet is a paid mutator transaction binding the contract method 0x98b9a2dc.
//
// Solidity: function changeWallet(address newWallet) returns()
func (_PMintyMultiSale *PMintyMultiSaleTransactor) ChangeWallet(opts *bind.TransactOpts, newWallet common.Address) (*types.Transaction, error) {
	return _PMintyMultiSale.contract.Transact(opts, "changeWallet", newWallet)
}

// ChangeWallet is a paid mutator transaction binding the contract method 0x98b9a2dc.
//
// Solidity: function changeWallet(address newWallet) returns()
func (_PMintyMultiSale *PMintyMultiSaleSession) ChangeWallet(newWallet common.Address) (*types.Transaction, error) {
	return _PMintyMultiSale.Contract.ChangeWallet(&_PMintyMultiSale.TransactOpts, newWallet)
}

// ChangeWallet is a paid mutator transaction binding the contract method 0x98b9a2dc.
//
// Solidity: function changeWallet(address newWallet) returns()
func (_PMintyMultiSale *PMintyMultiSaleTransactorSession) ChangeWallet(newWallet common.Address) (*types.Transaction, error) {
	return _PMintyMultiSale.Contract.ChangeWallet(&_PMintyMultiSale.TransactOpts, newWallet)
}

// OfferNew is a paid mutator transaction binding the contract method 0x8bc844c7.
//
// Solidity: function offerNew(address token, uint256 tokenId, string ipfsString, uint256 quantity, uint256 price, uint256 poolId) returns()
func (_PMintyMultiSale *PMintyMultiSaleTransactor) OfferNew(opts *bind.TransactOpts, token common.Address, tokenId *big.Int, ipfsString string, quantity *big.Int, price *big.Int, poolId *big.Int) (*types.Transaction, error) {
	return _PMintyMultiSale.contract.Transact(opts, "offerNew", token, tokenId, ipfsString, quantity, price, poolId)
}

// OfferNew is a paid mutator transaction binding the contract method 0x8bc844c7.
//
// Solidity: function offerNew(address token, uint256 tokenId, string ipfsString, uint256 quantity, uint256 price, uint256 poolId) returns()
func (_PMintyMultiSale *PMintyMultiSaleSession) OfferNew(token common.Address, tokenId *big.Int, ipfsString string, quantity *big.Int, price *big.Int, poolId *big.Int) (*types.Transaction, error) {
	return _PMintyMultiSale.Contract.OfferNew(&_PMintyMultiSale.TransactOpts, token, tokenId, ipfsString, quantity, price, poolId)
}

// OfferNew is a paid mutator transaction binding the contract method 0x8bc844c7.
//
// Solidity: function offerNew(address token, uint256 tokenId, string ipfsString, uint256 quantity, uint256 price, uint256 poolId) returns()
func (_PMintyMultiSale *PMintyMultiSaleTransactorSession) OfferNew(token common.Address, tokenId *big.Int, ipfsString string, quantity *big.Int, price *big.Int, poolId *big.Int) (*types.Transaction, error) {
	return _PMintyMultiSale.Contract.OfferNew(&_PMintyMultiSale.TransactOpts, token, tokenId, ipfsString, quantity, price, poolId)
}

// OfferNewBatch is a paid mutator transaction binding the contract method 0x5af9c479.
//
// Solidity: function offerNewBatch(address token, uint256[] tokenIds, string[] ipfsStrings, uint256[] quantities, uint256[] prices, uint256 poolId) returns()
func (_PMintyMultiSale *PMintyMultiSaleTransactor) OfferNewBatch(opts *bind.TransactOpts, token common.Address, tokenIds []*big.Int, ipfsStrings []string, quantities []*big.Int, prices []*big.Int, poolId *big.Int) (*types.Transaction, error) {
	return _PMintyMultiSale.contract.Transact(opts, "offerNewBatch", token, tokenIds, ipfsStrings, quantities, prices, poolId)
}

// OfferNewBatch is a paid mutator transaction binding the contract method 0x5af9c479.
//
// Solidity: function offerNewBatch(address token, uint256[] tokenIds, string[] ipfsStrings, uint256[] quantities, uint256[] prices, uint256 poolId) returns()
func (_PMintyMultiSale *PMintyMultiSaleSession) OfferNewBatch(token common.Address, tokenIds []*big.Int, ipfsStrings []string, quantities []*big.Int, prices []*big.Int, poolId *big.Int) (*types.Transaction, error) {
	return _PMintyMultiSale.Contract.OfferNewBatch(&_PMintyMultiSale.TransactOpts, token, tokenIds, ipfsStrings, quantities, prices, poolId)
}

// OfferNewBatch is a paid mutator transaction binding the contract method 0x5af9c479.
//
// Solidity: function offerNewBatch(address token, uint256[] tokenIds, string[] ipfsStrings, uint256[] quantities, uint256[] prices, uint256 poolId) returns()
func (_PMintyMultiSale *PMintyMultiSaleTransactorSession) OfferNewBatch(token common.Address, tokenIds []*big.Int, ipfsStrings []string, quantities []*big.Int, prices []*big.Int, poolId *big.Int) (*types.Transaction, error) {
	return _PMintyMultiSale.Contract.OfferNewBatch(&_PMintyMultiSale.TransactOpts, token, tokenIds, ipfsStrings, quantities, prices, poolId)
}

// OfferResale is a paid mutator transaction binding the contract method 0x7d5c7217.
//
// Solidity: function offerResale(address token, uint256 tokenId, uint256 quantity, uint256 price) returns()
func (_PMintyMultiSale *PMintyMultiSaleTransactor) OfferResale(opts *bind.TransactOpts, token common.Address, tokenId *big.Int, quantity *big.Int, price *big.Int) (*types.Transaction, error) {
	return _PMintyMultiSale.contract.Transact(opts, "offerResale", token, tokenId, quantity, price)
}

// OfferResale is a paid mutator transaction binding the contract method 0x7d5c7217.
//
// Solidity: function offerResale(address token, uint256 tokenId, uint256 quantity, uint256 price) returns()
func (_PMintyMultiSale *PMintyMultiSaleSession) OfferResale(token common.Address, tokenId *big.Int, quantity *big.Int, price *big.Int) (*types.Transaction, error) {
	return _PMintyMultiSale.Contract.OfferResale(&_PMintyMultiSale.TransactOpts, token, tokenId, quantity, price)
}

// OfferResale is a paid mutator transaction binding the contract method 0x7d5c7217.
//
// Solidity: function offerResale(address token, uint256 tokenId, uint256 quantity, uint256 price) returns()
func (_PMintyMultiSale *PMintyMultiSaleTransactorSession) OfferResale(token common.Address, tokenId *big.Int, quantity *big.Int, price *big.Int) (*types.Transaction, error) {
	return _PMintyMultiSale.Contract.OfferResale(&_PMintyMultiSale.TransactOpts, token, tokenId, quantity, price)
}

// ReSubmitOffer is a paid mutator transaction binding the contract method 0x4e900c02.
//
// Solidity: function reSubmitOffer(address token, uint256 tokenId, uint256 pos, uint256 price) returns()
func (_PMintyMultiSale *PMintyMultiSaleTransactor) ReSubmitOffer(opts *bind.TransactOpts, token common.Address, tokenId *big.Int, pos *big.Int, price *big.Int) (*types.Transaction, error) {
	return _PMintyMultiSale.contract.Transact(opts, "reSubmitOffer", token, tokenId, pos, price)
}

// ReSubmitOffer is a paid mutator transaction binding the contract method 0x4e900c02.
//
// Solidity: function reSubmitOffer(address token, uint256 tokenId, uint256 pos, uint256 price) returns()
func (_PMintyMultiSale *PMintyMultiSaleSession) ReSubmitOffer(token common.Address, tokenId *big.Int, pos *big.Int, price *big.Int) (*types.Transaction, error) {
	return _PMintyMultiSale.Contract.ReSubmitOffer(&_PMintyMultiSale.TransactOpts, token, tokenId, pos, price)
}

// ReSubmitOffer is a paid mutator transaction binding the contract method 0x4e900c02.
//
// Solidity: function reSubmitOffer(address token, uint256 tokenId, uint256 pos, uint256 price) returns()
func (_PMintyMultiSale *PMintyMultiSaleTransactorSession) ReSubmitOffer(token common.Address, tokenId *big.Int, pos *big.Int, price *big.Int) (*types.Transaction, error) {
	return _PMintyMultiSale.Contract.ReSubmitOffer(&_PMintyMultiSale.TransactOpts, token, tokenId, pos, price)
}

// RetractRemainingOffer is a paid mutator transaction binding the contract method 0xecf13ef4.
//
// Solidity: function retractRemainingOffer(address token, uint256 tokenId, uint256 pos) returns()
func (_PMintyMultiSale *PMintyMultiSaleTransactor) RetractRemainingOffer(opts *bind.TransactOpts, token common.Address, tokenId *big.Int, pos *big.Int) (*types.Transaction, error) {
	return _PMintyMultiSale.contract.Transact(opts, "retractRemainingOffer", token, tokenId, pos)
}

// RetractRemainingOffer is a paid mutator transaction binding the contract method 0xecf13ef4.
//
// Solidity: function retractRemainingOffer(address token, uint256 tokenId, uint256 pos) returns()
func (_PMintyMultiSale *PMintyMultiSaleSession) RetractRemainingOffer(token common.Address, tokenId *big.Int, pos *big.Int) (*types.Transaction, error) {
	return _PMintyMultiSale.Contract.RetractRemainingOffer(&_PMintyMultiSale.TransactOpts, token, tokenId, pos)
}

// RetractRemainingOffer is a paid mutator transaction binding the contract method 0xecf13ef4.
//
// Solidity: function retractRemainingOffer(address token, uint256 tokenId, uint256 pos) returns()
func (_PMintyMultiSale *PMintyMultiSaleTransactorSession) RetractRemainingOffer(token common.Address, tokenId *big.Int, pos *big.Int) (*types.Transaction, error) {
	return _PMintyMultiSale.Contract.RetractRemainingOffer(&_PMintyMultiSale.TransactOpts, token, tokenId, pos)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PMintyMultiSale *PMintyMultiSaleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PMintyMultiSale.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PMintyMultiSale *PMintyMultiSaleSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PMintyMultiSale.Contract.TransferOwnership(&_PMintyMultiSale.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PMintyMultiSale *PMintyMultiSaleTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PMintyMultiSale.Contract.TransferOwnership(&_PMintyMultiSale.TransactOpts, newOwner)
}

// UpdateShares is a paid mutator transaction binding the contract method 0xb6608467.
//
// Solidity: function updateShares(uint256 _divisor) returns()
func (_PMintyMultiSale *PMintyMultiSaleTransactor) UpdateShares(opts *bind.TransactOpts, _divisor *big.Int) (*types.Transaction, error) {
	return _PMintyMultiSale.contract.Transact(opts, "updateShares", _divisor)
}

// UpdateShares is a paid mutator transaction binding the contract method 0xb6608467.
//
// Solidity: function updateShares(uint256 _divisor) returns()
func (_PMintyMultiSale *PMintyMultiSaleSession) UpdateShares(_divisor *big.Int) (*types.Transaction, error) {
	return _PMintyMultiSale.Contract.UpdateShares(&_PMintyMultiSale.TransactOpts, _divisor)
}

// UpdateShares is a paid mutator transaction binding the contract method 0xb6608467.
//
// Solidity: function updateShares(uint256 _divisor) returns()
func (_PMintyMultiSale *PMintyMultiSaleTransactorSession) UpdateShares(_divisor *big.Int) (*types.Transaction, error) {
	return _PMintyMultiSale.Contract.UpdateShares(&_PMintyMultiSale.TransactOpts, _divisor)
}

// PMintyMultiSaleBidIncreasedIterator is returned from FilterBidIncreased and is used to iterate over the raw logs and unpacked data for BidIncreased events raised by the PMintyMultiSale contract.
type PMintyMultiSaleBidIncreasedIterator struct {
	Event *PMintyMultiSaleBidIncreased // Event containing the contract specifics and raw log

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
func (it *PMintyMultiSaleBidIncreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PMintyMultiSaleBidIncreased)
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
		it.Event = new(PMintyMultiSaleBidIncreased)
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
func (it *PMintyMultiSaleBidIncreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PMintyMultiSaleBidIncreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PMintyMultiSaleBidIncreased represents a BidIncreased event raised by the PMintyMultiSale contract.
type PMintyMultiSaleBidIncreased struct {
	Bidder      common.Address
	Token       common.Address
	TokenId     *big.Int
	PreviousBid *big.Int
	ThisBid     *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBidIncreased is a free log retrieval operation binding the contract event 0x239c9fe244a9bb0983901f35e4fc2fa1e99aed4a9896e48469975487fb877432.
//
// Solidity: event BidIncreased(address bidder, address token, uint256 tokenId, uint256 previous_bid, uint256 this_bid)
func (_PMintyMultiSale *PMintyMultiSaleFilterer) FilterBidIncreased(opts *bind.FilterOpts) (*PMintyMultiSaleBidIncreasedIterator, error) {

	logs, sub, err := _PMintyMultiSale.contract.FilterLogs(opts, "BidIncreased")
	if err != nil {
		return nil, err
	}
	return &PMintyMultiSaleBidIncreasedIterator{contract: _PMintyMultiSale.contract, event: "BidIncreased", logs: logs, sub: sub}, nil
}

// WatchBidIncreased is a free log subscription operation binding the contract event 0x239c9fe244a9bb0983901f35e4fc2fa1e99aed4a9896e48469975487fb877432.
//
// Solidity: event BidIncreased(address bidder, address token, uint256 tokenId, uint256 previous_bid, uint256 this_bid)
func (_PMintyMultiSale *PMintyMultiSaleFilterer) WatchBidIncreased(opts *bind.WatchOpts, sink chan<- *PMintyMultiSaleBidIncreased) (event.Subscription, error) {

	logs, sub, err := _PMintyMultiSale.contract.WatchLogs(opts, "BidIncreased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PMintyMultiSaleBidIncreased)
				if err := _PMintyMultiSale.contract.UnpackLog(event, "BidIncreased", log); err != nil {
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

// ParseBidIncreased is a log parse operation binding the contract event 0x239c9fe244a9bb0983901f35e4fc2fa1e99aed4a9896e48469975487fb877432.
//
// Solidity: event BidIncreased(address bidder, address token, uint256 tokenId, uint256 previous_bid, uint256 this_bid)
func (_PMintyMultiSale *PMintyMultiSaleFilterer) ParseBidIncreased(log types.Log) (*PMintyMultiSaleBidIncreased, error) {
	event := new(PMintyMultiSaleBidIncreased)
	if err := _PMintyMultiSale.contract.UnpackLog(event, "BidIncreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PMintyMultiSaleBidReceivedIterator is returned from FilterBidReceived and is used to iterate over the raw logs and unpacked data for BidReceived events raised by the PMintyMultiSale contract.
type PMintyMultiSaleBidReceivedIterator struct {
	Event *PMintyMultiSaleBidReceived // Event containing the contract specifics and raw log

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
func (it *PMintyMultiSaleBidReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PMintyMultiSaleBidReceived)
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
		it.Event = new(PMintyMultiSaleBidReceived)
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
func (it *PMintyMultiSaleBidReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PMintyMultiSaleBidReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PMintyMultiSaleBidReceived represents a BidReceived event raised by the PMintyMultiSale contract.
type PMintyMultiSaleBidReceived struct {
	Bidder  common.Address
	Token   common.Address
	TokenId *big.Int
	Bid     *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBidReceived is a free log retrieval operation binding the contract event 0x3c24064e704c839f701e9424574677be08efac54f5a23728c4da106999fcf503.
//
// Solidity: event BidReceived(address bidder, address token, uint256 tokenId, uint256 bid)
func (_PMintyMultiSale *PMintyMultiSaleFilterer) FilterBidReceived(opts *bind.FilterOpts) (*PMintyMultiSaleBidReceivedIterator, error) {

	logs, sub, err := _PMintyMultiSale.contract.FilterLogs(opts, "BidReceived")
	if err != nil {
		return nil, err
	}
	return &PMintyMultiSaleBidReceivedIterator{contract: _PMintyMultiSale.contract, event: "BidReceived", logs: logs, sub: sub}, nil
}

// WatchBidReceived is a free log subscription operation binding the contract event 0x3c24064e704c839f701e9424574677be08efac54f5a23728c4da106999fcf503.
//
// Solidity: event BidReceived(address bidder, address token, uint256 tokenId, uint256 bid)
func (_PMintyMultiSale *PMintyMultiSaleFilterer) WatchBidReceived(opts *bind.WatchOpts, sink chan<- *PMintyMultiSaleBidReceived) (event.Subscription, error) {

	logs, sub, err := _PMintyMultiSale.contract.WatchLogs(opts, "BidReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PMintyMultiSaleBidReceived)
				if err := _PMintyMultiSale.contract.UnpackLog(event, "BidReceived", log); err != nil {
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

// ParseBidReceived is a log parse operation binding the contract event 0x3c24064e704c839f701e9424574677be08efac54f5a23728c4da106999fcf503.
//
// Solidity: event BidReceived(address bidder, address token, uint256 tokenId, uint256 bid)
func (_PMintyMultiSale *PMintyMultiSaleFilterer) ParseBidReceived(log types.Log) (*PMintyMultiSaleBidReceived, error) {
	event := new(PMintyMultiSaleBidReceived)
	if err := _PMintyMultiSale.contract.UnpackLog(event, "BidReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PMintyMultiSaleBidWithdrawnIterator is returned from FilterBidWithdrawn and is used to iterate over the raw logs and unpacked data for BidWithdrawn events raised by the PMintyMultiSale contract.
type PMintyMultiSaleBidWithdrawnIterator struct {
	Event *PMintyMultiSaleBidWithdrawn // Event containing the contract specifics and raw log

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
func (it *PMintyMultiSaleBidWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PMintyMultiSaleBidWithdrawn)
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
		it.Event = new(PMintyMultiSaleBidWithdrawn)
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
func (it *PMintyMultiSaleBidWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PMintyMultiSaleBidWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PMintyMultiSaleBidWithdrawn represents a BidWithdrawn event raised by the PMintyMultiSale contract.
type PMintyMultiSaleBidWithdrawn struct {
	Token   common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBidWithdrawn is a free log retrieval operation binding the contract event 0x03f0427e8bcfdf5f69217150cf160ebe2dac5fa607336fd7643bfd61a9019080.
//
// Solidity: event BidWithdrawn(address token, uint256 tokenId)
func (_PMintyMultiSale *PMintyMultiSaleFilterer) FilterBidWithdrawn(opts *bind.FilterOpts) (*PMintyMultiSaleBidWithdrawnIterator, error) {

	logs, sub, err := _PMintyMultiSale.contract.FilterLogs(opts, "BidWithdrawn")
	if err != nil {
		return nil, err
	}
	return &PMintyMultiSaleBidWithdrawnIterator{contract: _PMintyMultiSale.contract, event: "BidWithdrawn", logs: logs, sub: sub}, nil
}

// WatchBidWithdrawn is a free log subscription operation binding the contract event 0x03f0427e8bcfdf5f69217150cf160ebe2dac5fa607336fd7643bfd61a9019080.
//
// Solidity: event BidWithdrawn(address token, uint256 tokenId)
func (_PMintyMultiSale *PMintyMultiSaleFilterer) WatchBidWithdrawn(opts *bind.WatchOpts, sink chan<- *PMintyMultiSaleBidWithdrawn) (event.Subscription, error) {

	logs, sub, err := _PMintyMultiSale.contract.WatchLogs(opts, "BidWithdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PMintyMultiSaleBidWithdrawn)
				if err := _PMintyMultiSale.contract.UnpackLog(event, "BidWithdrawn", log); err != nil {
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

// ParseBidWithdrawn is a log parse operation binding the contract event 0x03f0427e8bcfdf5f69217150cf160ebe2dac5fa607336fd7643bfd61a9019080.
//
// Solidity: event BidWithdrawn(address token, uint256 tokenId)
func (_PMintyMultiSale *PMintyMultiSaleFilterer) ParseBidWithdrawn(log types.Log) (*PMintyMultiSaleBidWithdrawn, error) {
	event := new(PMintyMultiSaleBidWithdrawn)
	if err := _PMintyMultiSale.contract.UnpackLog(event, "BidWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PMintyMultiSaleFeeUpdatedIterator is returned from FilterFeeUpdated and is used to iterate over the raw logs and unpacked data for FeeUpdated events raised by the PMintyMultiSale contract.
type PMintyMultiSaleFeeUpdatedIterator struct {
	Event *PMintyMultiSaleFeeUpdated // Event containing the contract specifics and raw log

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
func (it *PMintyMultiSaleFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PMintyMultiSaleFeeUpdated)
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
		it.Event = new(PMintyMultiSaleFeeUpdated)
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
func (it *PMintyMultiSaleFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PMintyMultiSaleFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PMintyMultiSaleFeeUpdated represents a FeeUpdated event raised by the PMintyMultiSale contract.
type PMintyMultiSaleFeeUpdated struct {
	Divisor *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFeeUpdated is a free log retrieval operation binding the contract event 0x8c4d35e54a3f2ef1134138fd8ea3daee6a3c89e10d2665996babdf70261e2c76.
//
// Solidity: event FeeUpdated(uint256 divisor)
func (_PMintyMultiSale *PMintyMultiSaleFilterer) FilterFeeUpdated(opts *bind.FilterOpts) (*PMintyMultiSaleFeeUpdatedIterator, error) {

	logs, sub, err := _PMintyMultiSale.contract.FilterLogs(opts, "FeeUpdated")
	if err != nil {
		return nil, err
	}
	return &PMintyMultiSaleFeeUpdatedIterator{contract: _PMintyMultiSale.contract, event: "FeeUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeUpdated is a free log subscription operation binding the contract event 0x8c4d35e54a3f2ef1134138fd8ea3daee6a3c89e10d2665996babdf70261e2c76.
//
// Solidity: event FeeUpdated(uint256 divisor)
func (_PMintyMultiSale *PMintyMultiSaleFilterer) WatchFeeUpdated(opts *bind.WatchOpts, sink chan<- *PMintyMultiSaleFeeUpdated) (event.Subscription, error) {

	logs, sub, err := _PMintyMultiSale.contract.WatchLogs(opts, "FeeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PMintyMultiSaleFeeUpdated)
				if err := _PMintyMultiSale.contract.UnpackLog(event, "FeeUpdated", log); err != nil {
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

// ParseFeeUpdated is a log parse operation binding the contract event 0x8c4d35e54a3f2ef1134138fd8ea3daee6a3c89e10d2665996babdf70261e2c76.
//
// Solidity: event FeeUpdated(uint256 divisor)
func (_PMintyMultiSale *PMintyMultiSaleFilterer) ParseFeeUpdated(log types.Log) (*PMintyMultiSaleFeeUpdated, error) {
	event := new(PMintyMultiSaleFeeUpdated)
	if err := _PMintyMultiSale.contract.UnpackLog(event, "FeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PMintyMultiSaleNewOfferIterator is returned from FilterNewOffer and is used to iterate over the raw logs and unpacked data for NewOffer events raised by the PMintyMultiSale contract.
type PMintyMultiSaleNewOfferIterator struct {
	Event *PMintyMultiSaleNewOffer // Event containing the contract specifics and raw log

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
func (it *PMintyMultiSaleNewOfferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PMintyMultiSaleNewOffer)
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
		it.Event = new(PMintyMultiSaleNewOffer)
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
func (it *PMintyMultiSaleNewOfferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PMintyMultiSaleNewOfferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PMintyMultiSaleNewOffer represents a NewOffer event raised by the PMintyMultiSale contract.
type PMintyMultiSaleNewOffer struct {
	Token    common.Address
	TokenId  *big.Int
	Owner    common.Address
	Quantity *big.Int
	Price    *big.Int
	Hash     string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewOffer is a free log retrieval operation binding the contract event 0x10575dae8255319b2807b08ada1038743fd34f2a06f1fba49e9407e7aa91f8f6.
//
// Solidity: event NewOffer(address token, uint256 tokenId, address owner, uint256 quantity, uint256 price, string hash)
func (_PMintyMultiSale *PMintyMultiSaleFilterer) FilterNewOffer(opts *bind.FilterOpts) (*PMintyMultiSaleNewOfferIterator, error) {

	logs, sub, err := _PMintyMultiSale.contract.FilterLogs(opts, "NewOffer")
	if err != nil {
		return nil, err
	}
	return &PMintyMultiSaleNewOfferIterator{contract: _PMintyMultiSale.contract, event: "NewOffer", logs: logs, sub: sub}, nil
}

// WatchNewOffer is a free log subscription operation binding the contract event 0x10575dae8255319b2807b08ada1038743fd34f2a06f1fba49e9407e7aa91f8f6.
//
// Solidity: event NewOffer(address token, uint256 tokenId, address owner, uint256 quantity, uint256 price, string hash)
func (_PMintyMultiSale *PMintyMultiSaleFilterer) WatchNewOffer(opts *bind.WatchOpts, sink chan<- *PMintyMultiSaleNewOffer) (event.Subscription, error) {

	logs, sub, err := _PMintyMultiSale.contract.WatchLogs(opts, "NewOffer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PMintyMultiSaleNewOffer)
				if err := _PMintyMultiSale.contract.UnpackLog(event, "NewOffer", log); err != nil {
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

// ParseNewOffer is a log parse operation binding the contract event 0x10575dae8255319b2807b08ada1038743fd34f2a06f1fba49e9407e7aa91f8f6.
//
// Solidity: event NewOffer(address token, uint256 tokenId, address owner, uint256 quantity, uint256 price, string hash)
func (_PMintyMultiSale *PMintyMultiSaleFilterer) ParseNewOffer(log types.Log) (*PMintyMultiSaleNewOffer, error) {
	event := new(PMintyMultiSaleNewOffer)
	if err := _PMintyMultiSale.contract.UnpackLog(event, "NewOffer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PMintyMultiSaleOfferAcceptedIterator is returned from FilterOfferAccepted and is used to iterate over the raw logs and unpacked data for OfferAccepted events raised by the PMintyMultiSale contract.
type PMintyMultiSaleOfferAcceptedIterator struct {
	Event *PMintyMultiSaleOfferAccepted // Event containing the contract specifics and raw log

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
func (it *PMintyMultiSaleOfferAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PMintyMultiSaleOfferAccepted)
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
		it.Event = new(PMintyMultiSaleOfferAccepted)
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
func (it *PMintyMultiSaleOfferAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PMintyMultiSaleOfferAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PMintyMultiSaleOfferAccepted represents a OfferAccepted event raised by the PMintyMultiSale contract.
type PMintyMultiSaleOfferAccepted struct {
	Buyer    common.Address
	Token    common.Address
	TokenId  *big.Int
	Pos      *big.Int
	Quantity *big.Int
	Price    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOfferAccepted is a free log retrieval operation binding the contract event 0xaf47c9ab14ff5ce679b4cb416be29684283e7b30fcfe405bfdebcc135f4b1e6f.
//
// Solidity: event OfferAccepted(address buyer, address token, uint256 tokenId, uint256 pos, uint256 quantity, uint256 price)
func (_PMintyMultiSale *PMintyMultiSaleFilterer) FilterOfferAccepted(opts *bind.FilterOpts) (*PMintyMultiSaleOfferAcceptedIterator, error) {

	logs, sub, err := _PMintyMultiSale.contract.FilterLogs(opts, "OfferAccepted")
	if err != nil {
		return nil, err
	}
	return &PMintyMultiSaleOfferAcceptedIterator{contract: _PMintyMultiSale.contract, event: "OfferAccepted", logs: logs, sub: sub}, nil
}

// WatchOfferAccepted is a free log subscription operation binding the contract event 0xaf47c9ab14ff5ce679b4cb416be29684283e7b30fcfe405bfdebcc135f4b1e6f.
//
// Solidity: event OfferAccepted(address buyer, address token, uint256 tokenId, uint256 pos, uint256 quantity, uint256 price)
func (_PMintyMultiSale *PMintyMultiSaleFilterer) WatchOfferAccepted(opts *bind.WatchOpts, sink chan<- *PMintyMultiSaleOfferAccepted) (event.Subscription, error) {

	logs, sub, err := _PMintyMultiSale.contract.WatchLogs(opts, "OfferAccepted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PMintyMultiSaleOfferAccepted)
				if err := _PMintyMultiSale.contract.UnpackLog(event, "OfferAccepted", log); err != nil {
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

// ParseOfferAccepted is a log parse operation binding the contract event 0xaf47c9ab14ff5ce679b4cb416be29684283e7b30fcfe405bfdebcc135f4b1e6f.
//
// Solidity: event OfferAccepted(address buyer, address token, uint256 tokenId, uint256 pos, uint256 quantity, uint256 price)
func (_PMintyMultiSale *PMintyMultiSaleFilterer) ParseOfferAccepted(log types.Log) (*PMintyMultiSaleOfferAccepted, error) {
	event := new(PMintyMultiSaleOfferAccepted)
	if err := _PMintyMultiSale.contract.UnpackLog(event, "OfferAccepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PMintyMultiSaleOfferWithdrawnIterator is returned from FilterOfferWithdrawn and is used to iterate over the raw logs and unpacked data for OfferWithdrawn events raised by the PMintyMultiSale contract.
type PMintyMultiSaleOfferWithdrawnIterator struct {
	Event *PMintyMultiSaleOfferWithdrawn // Event containing the contract specifics and raw log

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
func (it *PMintyMultiSaleOfferWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PMintyMultiSaleOfferWithdrawn)
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
		it.Event = new(PMintyMultiSaleOfferWithdrawn)
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
func (it *PMintyMultiSaleOfferWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PMintyMultiSaleOfferWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PMintyMultiSaleOfferWithdrawn represents a OfferWithdrawn event raised by the PMintyMultiSale contract.
type PMintyMultiSaleOfferWithdrawn struct {
	Token   common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOfferWithdrawn is a free log retrieval operation binding the contract event 0xe79a2b74049aac2661ef1f29587692b91005c0faa808d1e678f7e1dccc8011e8.
//
// Solidity: event OfferWithdrawn(address token, uint256 tokenId)
func (_PMintyMultiSale *PMintyMultiSaleFilterer) FilterOfferWithdrawn(opts *bind.FilterOpts) (*PMintyMultiSaleOfferWithdrawnIterator, error) {

	logs, sub, err := _PMintyMultiSale.contract.FilterLogs(opts, "OfferWithdrawn")
	if err != nil {
		return nil, err
	}
	return &PMintyMultiSaleOfferWithdrawnIterator{contract: _PMintyMultiSale.contract, event: "OfferWithdrawn", logs: logs, sub: sub}, nil
}

// WatchOfferWithdrawn is a free log subscription operation binding the contract event 0xe79a2b74049aac2661ef1f29587692b91005c0faa808d1e678f7e1dccc8011e8.
//
// Solidity: event OfferWithdrawn(address token, uint256 tokenId)
func (_PMintyMultiSale *PMintyMultiSaleFilterer) WatchOfferWithdrawn(opts *bind.WatchOpts, sink chan<- *PMintyMultiSaleOfferWithdrawn) (event.Subscription, error) {

	logs, sub, err := _PMintyMultiSale.contract.WatchLogs(opts, "OfferWithdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PMintyMultiSaleOfferWithdrawn)
				if err := _PMintyMultiSale.contract.UnpackLog(event, "OfferWithdrawn", log); err != nil {
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

// ParseOfferWithdrawn is a log parse operation binding the contract event 0xe79a2b74049aac2661ef1f29587692b91005c0faa808d1e678f7e1dccc8011e8.
//
// Solidity: event OfferWithdrawn(address token, uint256 tokenId)
func (_PMintyMultiSale *PMintyMultiSaleFilterer) ParseOfferWithdrawn(log types.Log) (*PMintyMultiSaleOfferWithdrawn, error) {
	event := new(PMintyMultiSaleOfferWithdrawn)
	if err := _PMintyMultiSale.contract.UnpackLog(event, "OfferWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PMintyMultiSaleOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PMintyMultiSale contract.
type PMintyMultiSaleOwnershipTransferredIterator struct {
	Event *PMintyMultiSaleOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PMintyMultiSaleOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PMintyMultiSaleOwnershipTransferred)
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
		it.Event = new(PMintyMultiSaleOwnershipTransferred)
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
func (it *PMintyMultiSaleOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PMintyMultiSaleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PMintyMultiSaleOwnershipTransferred represents a OwnershipTransferred event raised by the PMintyMultiSale contract.
type PMintyMultiSaleOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PMintyMultiSale *PMintyMultiSaleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PMintyMultiSaleOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PMintyMultiSale.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PMintyMultiSaleOwnershipTransferredIterator{contract: _PMintyMultiSale.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PMintyMultiSale *PMintyMultiSaleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PMintyMultiSaleOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PMintyMultiSale.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PMintyMultiSaleOwnershipTransferred)
				if err := _PMintyMultiSale.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PMintyMultiSale *PMintyMultiSaleFilterer) ParseOwnershipTransferred(log types.Log) (*PMintyMultiSaleOwnershipTransferred, error) {
	event := new(PMintyMultiSaleOwnershipTransferred)
	if err := _PMintyMultiSale.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PMintyMultiSalePaymentIterator is returned from FilterPayment and is used to iterate over the raw logs and unpacked data for Payment events raised by the PMintyMultiSale contract.
type PMintyMultiSalePaymentIterator struct {
	Event *PMintyMultiSalePayment // Event containing the contract specifics and raw log

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
func (it *PMintyMultiSalePaymentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PMintyMultiSalePayment)
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
		it.Event = new(PMintyMultiSalePayment)
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
func (it *PMintyMultiSalePaymentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PMintyMultiSalePaymentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PMintyMultiSalePayment represents a Payment event raised by the PMintyMultiSale contract.
type PMintyMultiSalePayment struct {
	Wallet  common.Address
	Creator common.Address
	Owner   common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPayment is a free log retrieval operation binding the contract event 0x0913aa4718e76655dbdbde41862a88aba8459d8721849484c97f53af3b5028e6.
//
// Solidity: event Payment(address wallet, address creator, address _owner, uint256 amount)
func (_PMintyMultiSale *PMintyMultiSaleFilterer) FilterPayment(opts *bind.FilterOpts) (*PMintyMultiSalePaymentIterator, error) {

	logs, sub, err := _PMintyMultiSale.contract.FilterLogs(opts, "Payment")
	if err != nil {
		return nil, err
	}
	return &PMintyMultiSalePaymentIterator{contract: _PMintyMultiSale.contract, event: "Payment", logs: logs, sub: sub}, nil
}

// WatchPayment is a free log subscription operation binding the contract event 0x0913aa4718e76655dbdbde41862a88aba8459d8721849484c97f53af3b5028e6.
//
// Solidity: event Payment(address wallet, address creator, address _owner, uint256 amount)
func (_PMintyMultiSale *PMintyMultiSaleFilterer) WatchPayment(opts *bind.WatchOpts, sink chan<- *PMintyMultiSalePayment) (event.Subscription, error) {

	logs, sub, err := _PMintyMultiSale.contract.WatchLogs(opts, "Payment")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PMintyMultiSalePayment)
				if err := _PMintyMultiSale.contract.UnpackLog(event, "Payment", log); err != nil {
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

// ParsePayment is a log parse operation binding the contract event 0x0913aa4718e76655dbdbde41862a88aba8459d8721849484c97f53af3b5028e6.
//
// Solidity: event Payment(address wallet, address creator, address _owner, uint256 amount)
func (_PMintyMultiSale *PMintyMultiSaleFilterer) ParsePayment(log types.Log) (*PMintyMultiSalePayment, error) {
	event := new(PMintyMultiSalePayment)
	if err := _PMintyMultiSale.contract.UnpackLog(event, "Payment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PMintyMultiSaleResaleOfferIterator is returned from FilterResaleOffer and is used to iterate over the raw logs and unpacked data for ResaleOffer events raised by the PMintyMultiSale contract.
type PMintyMultiSaleResaleOfferIterator struct {
	Event *PMintyMultiSaleResaleOffer // Event containing the contract specifics and raw log

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
func (it *PMintyMultiSaleResaleOfferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PMintyMultiSaleResaleOffer)
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
		it.Event = new(PMintyMultiSaleResaleOffer)
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
func (it *PMintyMultiSaleResaleOfferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PMintyMultiSaleResaleOfferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PMintyMultiSaleResaleOffer represents a ResaleOffer event raised by the PMintyMultiSale contract.
type PMintyMultiSaleResaleOffer struct {
	Token    common.Address
	TokenId  *big.Int
	Quantity *big.Int
	Owner    common.Address
	Price    *big.Int
	Position *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterResaleOffer is a free log retrieval operation binding the contract event 0x0df75d277ebbb180db8152f8c392e8bf6ba2720b35f9f98456eca3557b91182e.
//
// Solidity: event ResaleOffer(address token, uint256 tokenId, uint256 quantity, address owner, uint256 price, uint256 position)
func (_PMintyMultiSale *PMintyMultiSaleFilterer) FilterResaleOffer(opts *bind.FilterOpts) (*PMintyMultiSaleResaleOfferIterator, error) {

	logs, sub, err := _PMintyMultiSale.contract.FilterLogs(opts, "ResaleOffer")
	if err != nil {
		return nil, err
	}
	return &PMintyMultiSaleResaleOfferIterator{contract: _PMintyMultiSale.contract, event: "ResaleOffer", logs: logs, sub: sub}, nil
}

// WatchResaleOffer is a free log subscription operation binding the contract event 0x0df75d277ebbb180db8152f8c392e8bf6ba2720b35f9f98456eca3557b91182e.
//
// Solidity: event ResaleOffer(address token, uint256 tokenId, uint256 quantity, address owner, uint256 price, uint256 position)
func (_PMintyMultiSale *PMintyMultiSaleFilterer) WatchResaleOffer(opts *bind.WatchOpts, sink chan<- *PMintyMultiSaleResaleOffer) (event.Subscription, error) {

	logs, sub, err := _PMintyMultiSale.contract.WatchLogs(opts, "ResaleOffer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PMintyMultiSaleResaleOffer)
				if err := _PMintyMultiSale.contract.UnpackLog(event, "ResaleOffer", log); err != nil {
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

// ParseResaleOffer is a log parse operation binding the contract event 0x0df75d277ebbb180db8152f8c392e8bf6ba2720b35f9f98456eca3557b91182e.
//
// Solidity: event ResaleOffer(address token, uint256 tokenId, uint256 quantity, address owner, uint256 price, uint256 position)
func (_PMintyMultiSale *PMintyMultiSaleFilterer) ParseResaleOffer(log types.Log) (*PMintyMultiSaleResaleOffer, error) {
	event := new(PMintyMultiSaleResaleOffer)
	if err := _PMintyMultiSale.contract.UnpackLog(event, "ResaleOffer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PMintyMultiSaleSaleRepricedIterator is returned from FilterSaleRepriced and is used to iterate over the raw logs and unpacked data for SaleRepriced events raised by the PMintyMultiSale contract.
type PMintyMultiSaleSaleRepricedIterator struct {
	Event *PMintyMultiSaleSaleRepriced // Event containing the contract specifics and raw log

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
func (it *PMintyMultiSaleSaleRepricedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PMintyMultiSaleSaleRepriced)
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
		it.Event = new(PMintyMultiSaleSaleRepriced)
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
func (it *PMintyMultiSaleSaleRepricedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PMintyMultiSaleSaleRepricedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PMintyMultiSaleSaleRepriced represents a SaleRepriced event raised by the PMintyMultiSale contract.
type PMintyMultiSaleSaleRepriced struct {
	Token   common.Address
	TokenId *big.Int
	Pos     *big.Int
	Price   *big.Int
	Owner   common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSaleRepriced is a free log retrieval operation binding the contract event 0x0365ed613ce46864c2596157898bfc904a905c87c70c265de1473f51eedb1b77.
//
// Solidity: event SaleRepriced(address token, uint256 tokenId, uint256 pos, uint256 price, address owner)
func (_PMintyMultiSale *PMintyMultiSaleFilterer) FilterSaleRepriced(opts *bind.FilterOpts) (*PMintyMultiSaleSaleRepricedIterator, error) {

	logs, sub, err := _PMintyMultiSale.contract.FilterLogs(opts, "SaleRepriced")
	if err != nil {
		return nil, err
	}
	return &PMintyMultiSaleSaleRepricedIterator{contract: _PMintyMultiSale.contract, event: "SaleRepriced", logs: logs, sub: sub}, nil
}

// WatchSaleRepriced is a free log subscription operation binding the contract event 0x0365ed613ce46864c2596157898bfc904a905c87c70c265de1473f51eedb1b77.
//
// Solidity: event SaleRepriced(address token, uint256 tokenId, uint256 pos, uint256 price, address owner)
func (_PMintyMultiSale *PMintyMultiSaleFilterer) WatchSaleRepriced(opts *bind.WatchOpts, sink chan<- *PMintyMultiSaleSaleRepriced) (event.Subscription, error) {

	logs, sub, err := _PMintyMultiSale.contract.WatchLogs(opts, "SaleRepriced")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PMintyMultiSaleSaleRepriced)
				if err := _PMintyMultiSale.contract.UnpackLog(event, "SaleRepriced", log); err != nil {
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

// ParseSaleRepriced is a log parse operation binding the contract event 0x0365ed613ce46864c2596157898bfc904a905c87c70c265de1473f51eedb1b77.
//
// Solidity: event SaleRepriced(address token, uint256 tokenId, uint256 pos, uint256 price, address owner)
func (_PMintyMultiSale *PMintyMultiSaleFilterer) ParseSaleRepriced(log types.Log) (*PMintyMultiSaleSaleRepriced, error) {
	event := new(PMintyMultiSaleSaleRepriced)
	if err := _PMintyMultiSale.contract.UnpackLog(event, "SaleRepriced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PMintyMultiSaleSaleRetractedIterator is returned from FilterSaleRetracted and is used to iterate over the raw logs and unpacked data for SaleRetracted events raised by the PMintyMultiSale contract.
type PMintyMultiSaleSaleRetractedIterator struct {
	Event *PMintyMultiSaleSaleRetracted // Event containing the contract specifics and raw log

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
func (it *PMintyMultiSaleSaleRetractedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PMintyMultiSaleSaleRetracted)
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
		it.Event = new(PMintyMultiSaleSaleRetracted)
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
func (it *PMintyMultiSaleSaleRetractedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PMintyMultiSaleSaleRetractedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PMintyMultiSaleSaleRetracted represents a SaleRetracted event raised by the PMintyMultiSale contract.
type PMintyMultiSaleSaleRetracted struct {
	Token   common.Address
	TokenId *big.Int
	Pos     *big.Int
	Owner   common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSaleRetracted is a free log retrieval operation binding the contract event 0x74ee16c30acdc0986deebeda1b36d871a7ebf1dc7e2fd327dc7628617c37726d.
//
// Solidity: event SaleRetracted(address token, uint256 tokenId, uint256 pos, address owner)
func (_PMintyMultiSale *PMintyMultiSaleFilterer) FilterSaleRetracted(opts *bind.FilterOpts) (*PMintyMultiSaleSaleRetractedIterator, error) {

	logs, sub, err := _PMintyMultiSale.contract.FilterLogs(opts, "SaleRetracted")
	if err != nil {
		return nil, err
	}
	return &PMintyMultiSaleSaleRetractedIterator{contract: _PMintyMultiSale.contract, event: "SaleRetracted", logs: logs, sub: sub}, nil
}

// WatchSaleRetracted is a free log subscription operation binding the contract event 0x74ee16c30acdc0986deebeda1b36d871a7ebf1dc7e2fd327dc7628617c37726d.
//
// Solidity: event SaleRetracted(address token, uint256 tokenId, uint256 pos, address owner)
func (_PMintyMultiSale *PMintyMultiSaleFilterer) WatchSaleRetracted(opts *bind.WatchOpts, sink chan<- *PMintyMultiSaleSaleRetracted) (event.Subscription, error) {

	logs, sub, err := _PMintyMultiSale.contract.WatchLogs(opts, "SaleRetracted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PMintyMultiSaleSaleRetracted)
				if err := _PMintyMultiSale.contract.UnpackLog(event, "SaleRetracted", log); err != nil {
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

// ParseSaleRetracted is a log parse operation binding the contract event 0x74ee16c30acdc0986deebeda1b36d871a7ebf1dc7e2fd327dc7628617c37726d.
//
// Solidity: event SaleRetracted(address token, uint256 tokenId, uint256 pos, address owner)
func (_PMintyMultiSale *PMintyMultiSaleFilterer) ParseSaleRetracted(log types.Log) (*PMintyMultiSaleSaleRetracted, error) {
	event := new(PMintyMultiSaleSaleRetracted)
	if err := _PMintyMultiSale.contract.UnpackLog(event, "SaleRetracted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PMintyMultiSaleWalletTransferredIterator is returned from FilterWalletTransferred and is used to iterate over the raw logs and unpacked data for WalletTransferred events raised by the PMintyMultiSale contract.
type PMintyMultiSaleWalletTransferredIterator struct {
	Event *PMintyMultiSaleWalletTransferred // Event containing the contract specifics and raw log

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
func (it *PMintyMultiSaleWalletTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PMintyMultiSaleWalletTransferred)
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
		it.Event = new(PMintyMultiSaleWalletTransferred)
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
func (it *PMintyMultiSaleWalletTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PMintyMultiSaleWalletTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PMintyMultiSaleWalletTransferred represents a WalletTransferred event raised by the PMintyMultiSale contract.
type PMintyMultiSaleWalletTransferred struct {
	PreviousWallet common.Address
	NewWallet      common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWalletTransferred is a free log retrieval operation binding the contract event 0xf404d79dbd5d19bc2bf99a74f0e5affb6502b3d6012140f1d908ce3fcd0716d8.
//
// Solidity: event WalletTransferred(address indexed previousWallet, address indexed newWallet)
func (_PMintyMultiSale *PMintyMultiSaleFilterer) FilterWalletTransferred(opts *bind.FilterOpts, previousWallet []common.Address, newWallet []common.Address) (*PMintyMultiSaleWalletTransferredIterator, error) {

	var previousWalletRule []interface{}
	for _, previousWalletItem := range previousWallet {
		previousWalletRule = append(previousWalletRule, previousWalletItem)
	}
	var newWalletRule []interface{}
	for _, newWalletItem := range newWallet {
		newWalletRule = append(newWalletRule, newWalletItem)
	}

	logs, sub, err := _PMintyMultiSale.contract.FilterLogs(opts, "WalletTransferred", previousWalletRule, newWalletRule)
	if err != nil {
		return nil, err
	}
	return &PMintyMultiSaleWalletTransferredIterator{contract: _PMintyMultiSale.contract, event: "WalletTransferred", logs: logs, sub: sub}, nil
}

// WatchWalletTransferred is a free log subscription operation binding the contract event 0xf404d79dbd5d19bc2bf99a74f0e5affb6502b3d6012140f1d908ce3fcd0716d8.
//
// Solidity: event WalletTransferred(address indexed previousWallet, address indexed newWallet)
func (_PMintyMultiSale *PMintyMultiSaleFilterer) WatchWalletTransferred(opts *bind.WatchOpts, sink chan<- *PMintyMultiSaleWalletTransferred, previousWallet []common.Address, newWallet []common.Address) (event.Subscription, error) {

	var previousWalletRule []interface{}
	for _, previousWalletItem := range previousWallet {
		previousWalletRule = append(previousWalletRule, previousWalletItem)
	}
	var newWalletRule []interface{}
	for _, newWalletItem := range newWallet {
		newWalletRule = append(newWalletRule, newWalletItem)
	}

	logs, sub, err := _PMintyMultiSale.contract.WatchLogs(opts, "WalletTransferred", previousWalletRule, newWalletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PMintyMultiSaleWalletTransferred)
				if err := _PMintyMultiSale.contract.UnpackLog(event, "WalletTransferred", log); err != nil {
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

// ParseWalletTransferred is a log parse operation binding the contract event 0xf404d79dbd5d19bc2bf99a74f0e5affb6502b3d6012140f1d908ce3fcd0716d8.
//
// Solidity: event WalletTransferred(address indexed previousWallet, address indexed newWallet)
func (_PMintyMultiSale *PMintyMultiSaleFilterer) ParseWalletTransferred(log types.Log) (*PMintyMultiSaleWalletTransferred, error) {
	event := new(PMintyMultiSaleWalletTransferred)
	if err := _PMintyMultiSale.contract.UnpackLog(event, "WalletTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
