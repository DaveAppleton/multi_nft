// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pMintyMultiToken

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

// AddressABI is the input ABI used to generate the binding from.
const AddressABI = "[]"

// AddressBin is the compiled bytecode used for deploying new contracts.
var AddressBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220d979ab20612eb26d97c44c787908db0031cac3fb7fddb3630439bc4d2671033564736f6c63430007050033"

// DeployAddress deploys a new Ethereum contract, binding an instance of Address to it.
func DeployAddress(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Address, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AddressBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// Address is an auto generated Go binding around an Ethereum contract.
type Address struct {
	AddressCaller     // Read-only binding to the contract
	AddressTransactor // Write-only binding to the contract
	AddressFilterer   // Log filterer for contract events
}

// AddressCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddressCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddressSession struct {
	Contract     *Address          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AddressCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddressCallerSession struct {
	Contract *AddressCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// AddressTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddressTransactorSession struct {
	Contract     *AddressTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AddressRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddressRaw struct {
	Contract *Address // Generic contract binding to access the raw methods on
}

// AddressCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddressCallerRaw struct {
	Contract *AddressCaller // Generic read-only contract binding to access the raw methods on
}

// AddressTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddressTransactorRaw struct {
	Contract *AddressTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddress creates a new instance of Address, bound to a specific deployed contract.
func NewAddress(address common.Address, backend bind.ContractBackend) (*Address, error) {
	contract, err := bindAddress(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// NewAddressCaller creates a new read-only instance of Address, bound to a specific deployed contract.
func NewAddressCaller(address common.Address, caller bind.ContractCaller) (*AddressCaller, error) {
	contract, err := bindAddress(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressCaller{contract: contract}, nil
}

// NewAddressTransactor creates a new write-only instance of Address, bound to a specific deployed contract.
func NewAddressTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressTransactor, error) {
	contract, err := bindAddress(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressTransactor{contract: contract}, nil
}

// NewAddressFilterer creates a new log filterer instance of Address, bound to a specific deployed contract.
func NewAddressFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressFilterer, error) {
	contract, err := bindAddress(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressFilterer{contract: contract}, nil
}

// bindAddress binds a generic wrapper to an already deployed contract.
func bindAddress(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Address.Contract.AddressCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Address.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.contract.Transact(opts, method, params...)
}

// ContextABI is the input ABI used to generate the binding from.
const ContextABI = "[]"

// Context is an auto generated Go binding around an Ethereum contract.
type Context struct {
	ContextCaller     // Read-only binding to the contract
	ContextTransactor // Write-only binding to the contract
	ContextFilterer   // Log filterer for contract events
}

// ContextCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContextCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContextTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContextFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContextSession struct {
	Contract     *Context          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContextCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContextCallerSession struct {
	Contract *ContextCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ContextTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContextTransactorSession struct {
	Contract     *ContextTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ContextRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContextRaw struct {
	Contract *Context // Generic contract binding to access the raw methods on
}

// ContextCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContextCallerRaw struct {
	Contract *ContextCaller // Generic read-only contract binding to access the raw methods on
}

// ContextTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContextTransactorRaw struct {
	Contract *ContextTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContext creates a new instance of Context, bound to a specific deployed contract.
func NewContext(address common.Address, backend bind.ContractBackend) (*Context, error) {
	contract, err := bindContext(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Context{ContextCaller: ContextCaller{contract: contract}, ContextTransactor: ContextTransactor{contract: contract}, ContextFilterer: ContextFilterer{contract: contract}}, nil
}

// NewContextCaller creates a new read-only instance of Context, bound to a specific deployed contract.
func NewContextCaller(address common.Address, caller bind.ContractCaller) (*ContextCaller, error) {
	contract, err := bindContext(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContextCaller{contract: contract}, nil
}

// NewContextTransactor creates a new write-only instance of Context, bound to a specific deployed contract.
func NewContextTransactor(address common.Address, transactor bind.ContractTransactor) (*ContextTransactor, error) {
	contract, err := bindContext(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContextTransactor{contract: contract}, nil
}

// NewContextFilterer creates a new log filterer instance of Context, bound to a specific deployed contract.
func NewContextFilterer(address common.Address, filterer bind.ContractFilterer) (*ContextFilterer, error) {
	contract, err := bindContext(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContextFilterer{contract: contract}, nil
}

// bindContext binds a generic wrapper to an already deployed contract.
func bindContext(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContextABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Context.Contract.ContextCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Context.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.contract.Transact(opts, method, params...)
}

// ERC1155ABI is the input ABI used to generate the binding from.
const ERC1155ABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uri_\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ERC1155FuncSigs maps the 4-byte function signature to its string representation.
var ERC1155FuncSigs = map[string]string{
	"00fdd58e": "balanceOf(address,uint256)",
	"4e1273f4": "balanceOfBatch(address[],uint256[])",
	"e985e9c5": "isApprovedForAll(address,address)",
	"2eb2c2d6": "safeBatchTransferFrom(address,address,uint256[],uint256[],bytes)",
	"f242432a": "safeTransferFrom(address,address,uint256,uint256,bytes)",
	"a22cb465": "setApprovalForAll(address,bool)",
	"01ffc9a7": "supportsInterface(bytes4)",
	"0e89341c": "uri(uint256)",
}

// ERC1155Bin is the compiled bytecode used for deploying new contracts.
var ERC1155Bin = "0x60806040523480156200001157600080fd5b50604051620016963803806200169683398101604081905262000034916200019c565b620000466301ffc9a760e01b6200007c565b6200005181620000d7565b62000063636cdb3d1360e11b6200007c565b620000756303a24d0760e21b6200007c565b5062000292565b6001600160e01b03198082161415620000b25760405162461bcd60e51b8152600401620000a9906200025b565b60405180910390fd5b6001600160e01b0319166000908152602081905260409020805460ff19166001179055565b8051620000ec906003906020840190620000f0565b5050565b828054600181600116156101000203166002900490600052602060002090601f01602090048101928262000128576000855562000173565b82601f106200014357805160ff191683800117855562000173565b8280016001018555821562000173579182015b828111156200017357825182559160200191906001019062000156565b506200018192915062000185565b5090565b5b8082111562000181576000815560010162000186565b60006020808385031215620001af578182fd5b82516001600160401b0380821115620001c6578384fd5b818501915085601f830112620001da578384fd5b815181811115620001e757fe5b604051601f8201601f19168101850183811182821017156200020557fe5b60405281815283820185018810156200021c578586fd5b8592505b818310156200023f578383018501518184018601529184019162000220565b818311156200025057858583830101525b979650505050505050565b6020808252601c908201527f4552433136353a20696e76616c696420696e7465726661636520696400000000604082015260600190565b6113f480620002a26000396000f3fe608060405234801561001057600080fd5b50600436106100875760003560e01c80634e1273f41161005b5780634e1273f41461010a578063a22cb4651461012a578063e985e9c51461013d578063f242432a1461015057610087565b8062fdd58e1461008c57806301ffc9a7146100b55780630e89341c146100d55780632eb2c2d6146100f5575b600080fd5b61009f61009a366004610cde565b610163565b6040516100ac9190611277565b60405180910390f35b6100c86100c3366004610dc8565b6101bc565b6040516100ac9190610f81565b6100e86100e3366004610e00565b6101df565b6040516100ac9190610f8c565b610108610103366004610b9b565b610277565b005b61011d610118366004610d07565b610491565b6040516100ac9190610f40565b610108610138366004610ca4565b61055e565b6100c861014b366004610b69565b61062c565b61010861015e366004610c41565b61065a565b60006001600160a01b0383166101945760405162461bcd60e51b815260040161018b9061103b565b60405180910390fd5b5060009081526001602090815260408083206001600160a01b03949094168352929052205490565b6001600160e01b0319811660009081526020819052604090205460ff165b919050565b60038054604080516020601f600260001961010060018816150201909516949094049384018190048102820181019092528281526060939092909183018282801561026b5780601f106102405761010080835404028352916020019161026b565b820191906000526020600020905b81548152906001019060200180831161024e57829003601f168201915b50505050509050919050565b81518351146102985760405162461bcd60e51b815260040161018b9061122f565b6001600160a01b0384166102be5760405162461bcd60e51b815260040161018b90611106565b6102c66107ee565b6001600160a01b0316856001600160a01b031614806102ec57506102ec8561014b6107ee565b6103085760405162461bcd60e51b815260040161018b9061114b565b60006103126107ee565b9050610322818787878787610489565b60005b845181101561042357600085828151811061033c57fe5b60200260200101519050600085838151811061035457fe5b602002602001015190506103c1816040518060600160405280602a8152602001611395602a91396001600086815260200190815260200160002060008d6001600160a01b03166001600160a01b03168152602001908152602001600020546107f39092919063ffffffff16565b60008381526001602090815260408083206001600160a01b038e811685529252808320939093558a16815220546103f8908261081f565b60009283526001602081815260408086206001600160a01b038d168752909152909320555001610325565b50846001600160a01b0316866001600160a01b0316826001600160a01b03167f4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb8787604051610473929190610f53565b60405180910390a461048981878787878761084b565b505050505050565b606081518351146104b45760405162461bcd60e51b815260040161018b906111e6565b6060835167ffffffffffffffff811180156104ce57600080fd5b506040519080825280602002602001820160405280156104f8578160200160208202803683370190505b50905060005b84518110156105565761053785828151811061051657fe5b602002602001015185838151811061052a57fe5b6020026020010151610163565b82828151811061054357fe5b60209081029190910101526001016104fe565b509392505050565b816001600160a01b03166105706107ee565b6001600160a01b031614156105975760405162461bcd60e51b815260040161018b9061119d565b80600260006105a46107ee565b6001600160a01b03908116825260208083019390935260409182016000908120918716808252919093529120805460ff1916921515929092179091556105e86107ee565b6001600160a01b03167f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31836040516106209190610f81565b60405180910390a35050565b6001600160a01b03918216600090815260026020908152604080832093909416825291909152205460ff1690565b6001600160a01b0384166106805760405162461bcd60e51b815260040161018b90611106565b6106886107ee565b6001600160a01b0316856001600160a01b031614806106ae57506106ae8561014b6107ee565b6106ca5760405162461bcd60e51b815260040161018b906110bd565b60006106d46107ee565b90506106f48187876106e588610962565b6106ee88610962565b87610489565b61073b836040518060600160405280602a8152602001611395602a913960008781526001602090815260408083206001600160a01b038d16845290915290205491906107f3565b60008581526001602090815260408083206001600160a01b038b81168552925280832093909355871681522054610772908461081f565b60008581526001602090815260408083206001600160a01b03808b168086529190935292819020939093559151909188811691908416907fc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62906107d89089908990611280565b60405180910390a46104898187878787876109a6565b335b90565b600081848411156108175760405162461bcd60e51b815260040161018b9190610f8c565b505050900390565b6000828201838110156108445760405162461bcd60e51b815260040161018b90611086565b9392505050565b61085d846001600160a01b0316610a77565b156104895760405163bc197c8160e01b81526001600160a01b0385169063bc197c81906108969089908990889088908890600401610e9d565b602060405180830381600087803b1580156108b057600080fd5b505af19250505080156108e0575060408051601f3d908101601f191682019092526108dd91810190610de4565b60015b610929576108ec6112d6565b806108f75750610911565b8060405162461bcd60e51b815260040161018b9190610f8c565b60405162461bcd60e51b815260040161018b90610f9f565b6001600160e01b0319811663bc197c8160e01b146109595760405162461bcd60e51b815260040161018b90610ff3565b50505050505050565b60408051600180825281830190925260609182919060208083019080368337019050509050828160008151811061099557fe5b602090810291909101015292915050565b6109b8846001600160a01b0316610a77565b156104895760405163f23a6e6160e01b81526001600160a01b0385169063f23a6e61906109f19089908990889088908890600401610efb565b602060405180830381600087803b158015610a0b57600080fd5b505af1925050508015610a3b575060408051601f3d908101601f19168201909252610a3891810190610de4565b60015b610a47576108ec6112d6565b6001600160e01b0319811663f23a6e6160e01b146109595760405162461bcd60e51b815260040161018b90610ff3565b3b151590565b80356001600160a01b03811681146101da57600080fd5b600082601f830112610aa4578081fd5b8135610ab7610ab2826112b2565b61128e565b818152915060208083019084810181840286018201871015610ad857600080fd5b60005b84811015610af757813584529282019290820190600101610adb565b505050505092915050565b600082601f830112610b12578081fd5b813567ffffffffffffffff811115610b2657fe5b610b39601f8201601f191660200161128e565b9150808252836020828501011115610b5057600080fd5b8060208401602084013760009082016020015292915050565b60008060408385031215610b7b578182fd5b610b8483610a7d565b9150610b9260208401610a7d565b90509250929050565b600080600080600060a08688031215610bb2578081fd5b610bbb86610a7d565b9450610bc960208701610a7d565b9350604086013567ffffffffffffffff80821115610be5578283fd5b610bf189838a01610a94565b94506060880135915080821115610c06578283fd5b610c1289838a01610a94565b93506080880135915080821115610c27578283fd5b50610c3488828901610b02565b9150509295509295909350565b600080600080600060a08688031215610c58578081fd5b610c6186610a7d565b9450610c6f60208701610a7d565b93506040860135925060608601359150608086013567ffffffffffffffff811115610c98578182fd5b610c3488828901610b02565b60008060408385031215610cb6578182fd5b610cbf83610a7d565b915060208301358015158114610cd3578182fd5b809150509250929050565b60008060408385031215610cf0578182fd5b610cf983610a7d565b946020939093013593505050565b60008060408385031215610d19578182fd5b823567ffffffffffffffff80821115610d30578384fd5b818501915085601f830112610d43578384fd5b8135610d51610ab2826112b2565b80828252602080830192508086018a828387028901011115610d71578889fd5b8896505b84871015610d9a57610d8681610a7d565b845260019690960195928101928101610d75565b509096508701359350505080821115610db1578283fd5b50610dbe85828601610a94565b9150509250929050565b600060208284031215610dd9578081fd5b81356108448161137b565b600060208284031215610df5578081fd5b81516108448161137b565b600060208284031215610e11578081fd5b5035919050565b6000815180845260208085019450808401835b83811015610e4757815187529582019590820190600101610e2b565b509495945050505050565b60008151808452815b81811015610e7757602081850181015186830182015201610e5b565b81811115610e885782602083870101525b50601f01601f19169290920160200192915050565b6001600160a01b0386811682528516602082015260a060408201819052600090610ec990830186610e18565b8281036060840152610edb8186610e18565b90508281036080840152610eef8185610e52565b98975050505050505050565b6001600160a01b03868116825285166020820152604081018490526060810183905260a060808201819052600090610f3590830184610e52565b979650505050505050565b6000602082526108446020830184610e18565b600060408252610f666040830185610e18565b8281036020840152610f788185610e18565b95945050505050565b901515815260200190565b6000602082526108446020830184610e52565b60208082526034908201527f455243313135353a207472616e7366657220746f206e6f6e20455243313135356040820152732932b1b2b4bb32b91034b6b83632b6b2b73a32b960611b606082015260800190565b60208082526028908201527f455243313135353a204552433131353552656365697665722072656a656374656040820152676420746f6b656e7360c01b606082015260800190565b6020808252602b908201527f455243313135353a2062616c616e636520717565727920666f7220746865207a60408201526a65726f206164647265737360a81b606082015260800190565b6020808252601b908201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604082015260600190565b60208082526029908201527f455243313135353a2063616c6c6572206973206e6f74206f776e6572206e6f7260408201526808185c1c1c9bdd995960ba1b606082015260800190565b60208082526025908201527f455243313135353a207472616e7366657220746f20746865207a65726f206164604082015264647265737360d81b606082015260800190565b60208082526032908201527f455243313135353a207472616e736665722063616c6c6572206973206e6f74206040820152711bdddb995c881b9bdc88185c1c1c9bdd995960721b606082015260800190565b60208082526029908201527f455243313135353a2073657474696e6720617070726f76616c20737461747573604082015268103337b91039b2b63360b91b606082015260800190565b60208082526029908201527f455243313135353a206163636f756e747320616e6420696473206c656e677468604082015268040dad2e6dac2e8c6d60bb1b606082015260800190565b60208082526028908201527f455243313135353a2069647320616e6420616d6f756e7473206c656e677468206040820152670dad2e6dac2e8c6d60c31b606082015260800190565b90815260200190565b918252602082015260400190565b60405181810167ffffffffffffffff811182821017156112aa57fe5b604052919050565b600067ffffffffffffffff8211156112c657fe5b5060209081020190565b60e01c90565b600060443d10156112e6576107f0565b600481823e6308c379a06112fa82516112d0565b14611304576107f0565b6040513d600319016004823e80513d67ffffffffffffffff816024840111818411171561133457505050506107f0565b8284019250825191508082111561134e57505050506107f0565b503d83016020828401011115611366575050506107f0565b601f01601f1916810160200160405291505090565b6001600160e01b03198116811461139157600080fd5b5056fe455243313135353a20696e73756666696369656e742062616c616e636520666f72207472616e73666572a264697066735822122005dd4eb17a7b113a1e85fedc928ebe4b08f3d2a60801451430c92af2f544590764736f6c63430007050033"

// DeployERC1155 deploys a new Ethereum contract, binding an instance of ERC1155 to it.
func DeployERC1155(auth *bind.TransactOpts, backend bind.ContractBackend, uri_ string) (common.Address, *types.Transaction, *ERC1155, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC1155ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC1155Bin), backend, uri_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC1155{ERC1155Caller: ERC1155Caller{contract: contract}, ERC1155Transactor: ERC1155Transactor{contract: contract}, ERC1155Filterer: ERC1155Filterer{contract: contract}}, nil
}

// ERC1155 is an auto generated Go binding around an Ethereum contract.
type ERC1155 struct {
	ERC1155Caller     // Read-only binding to the contract
	ERC1155Transactor // Write-only binding to the contract
	ERC1155Filterer   // Log filterer for contract events
}

// ERC1155Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC1155Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1155Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC1155Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1155Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC1155Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1155Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC1155Session struct {
	Contract     *ERC1155          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC1155CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC1155CallerSession struct {
	Contract *ERC1155Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ERC1155TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC1155TransactorSession struct {
	Contract     *ERC1155Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ERC1155Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC1155Raw struct {
	Contract *ERC1155 // Generic contract binding to access the raw methods on
}

// ERC1155CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC1155CallerRaw struct {
	Contract *ERC1155Caller // Generic read-only contract binding to access the raw methods on
}

// ERC1155TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC1155TransactorRaw struct {
	Contract *ERC1155Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC1155 creates a new instance of ERC1155, bound to a specific deployed contract.
func NewERC1155(address common.Address, backend bind.ContractBackend) (*ERC1155, error) {
	contract, err := bindERC1155(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC1155{ERC1155Caller: ERC1155Caller{contract: contract}, ERC1155Transactor: ERC1155Transactor{contract: contract}, ERC1155Filterer: ERC1155Filterer{contract: contract}}, nil
}

// NewERC1155Caller creates a new read-only instance of ERC1155, bound to a specific deployed contract.
func NewERC1155Caller(address common.Address, caller bind.ContractCaller) (*ERC1155Caller, error) {
	contract, err := bindERC1155(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC1155Caller{contract: contract}, nil
}

// NewERC1155Transactor creates a new write-only instance of ERC1155, bound to a specific deployed contract.
func NewERC1155Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC1155Transactor, error) {
	contract, err := bindERC1155(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC1155Transactor{contract: contract}, nil
}

// NewERC1155Filterer creates a new log filterer instance of ERC1155, bound to a specific deployed contract.
func NewERC1155Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC1155Filterer, error) {
	contract, err := bindERC1155(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC1155Filterer{contract: contract}, nil
}

// bindERC1155 binds a generic wrapper to an already deployed contract.
func bindERC1155(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC1155ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC1155 *ERC1155Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC1155.Contract.ERC1155Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC1155 *ERC1155Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC1155.Contract.ERC1155Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC1155 *ERC1155Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC1155.Contract.ERC1155Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC1155 *ERC1155CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC1155.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC1155 *ERC1155TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC1155.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC1155 *ERC1155TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC1155.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_ERC1155 *ERC1155Caller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ERC1155.contract.Call(opts, &out, "balanceOf", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_ERC1155 *ERC1155Session) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _ERC1155.Contract.BalanceOf(&_ERC1155.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_ERC1155 *ERC1155CallerSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _ERC1155.Contract.BalanceOf(&_ERC1155.CallOpts, account, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_ERC1155 *ERC1155Caller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _ERC1155.contract.Call(opts, &out, "balanceOfBatch", accounts, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_ERC1155 *ERC1155Session) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _ERC1155.Contract.BalanceOfBatch(&_ERC1155.CallOpts, accounts, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_ERC1155 *ERC1155CallerSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _ERC1155.Contract.BalanceOfBatch(&_ERC1155.CallOpts, accounts, ids)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_ERC1155 *ERC1155Caller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _ERC1155.contract.Call(opts, &out, "isApprovedForAll", account, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_ERC1155 *ERC1155Session) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _ERC1155.Contract.IsApprovedForAll(&_ERC1155.CallOpts, account, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_ERC1155 *ERC1155CallerSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _ERC1155.Contract.IsApprovedForAll(&_ERC1155.CallOpts, account, operator)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC1155 *ERC1155Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ERC1155.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC1155 *ERC1155Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC1155.Contract.SupportsInterface(&_ERC1155.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC1155 *ERC1155CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC1155.Contract.SupportsInterface(&_ERC1155.CallOpts, interfaceId)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_ERC1155 *ERC1155Caller) Uri(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _ERC1155.contract.Call(opts, &out, "uri", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_ERC1155 *ERC1155Session) Uri(arg0 *big.Int) (string, error) {
	return _ERC1155.Contract.Uri(&_ERC1155.CallOpts, arg0)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_ERC1155 *ERC1155CallerSession) Uri(arg0 *big.Int) (string, error) {
	return _ERC1155.Contract.Uri(&_ERC1155.CallOpts, arg0)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_ERC1155 *ERC1155Transactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _ERC1155.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_ERC1155 *ERC1155Session) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _ERC1155.Contract.SafeBatchTransferFrom(&_ERC1155.TransactOpts, from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_ERC1155 *ERC1155TransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _ERC1155.Contract.SafeBatchTransferFrom(&_ERC1155.TransactOpts, from, to, ids, amounts, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_ERC1155 *ERC1155Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC1155.contract.Transact(opts, "safeTransferFrom", from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_ERC1155 *ERC1155Session) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC1155.Contract.SafeTransferFrom(&_ERC1155.TransactOpts, from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_ERC1155 *ERC1155TransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC1155.Contract.SafeTransferFrom(&_ERC1155.TransactOpts, from, to, id, amount, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ERC1155 *ERC1155Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _ERC1155.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ERC1155 *ERC1155Session) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ERC1155.Contract.SetApprovalForAll(&_ERC1155.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ERC1155 *ERC1155TransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ERC1155.Contract.SetApprovalForAll(&_ERC1155.TransactOpts, operator, approved)
}

// ERC1155ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ERC1155 contract.
type ERC1155ApprovalForAllIterator struct {
	Event *ERC1155ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ERC1155ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155ApprovalForAll)
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
		it.Event = new(ERC1155ApprovalForAll)
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
func (it *ERC1155ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155ApprovalForAll represents a ApprovalForAll event raised by the ERC1155 contract.
type ERC1155ApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_ERC1155 *ERC1155Filterer) FilterApprovalForAll(opts *bind.FilterOpts, account []common.Address, operator []common.Address) (*ERC1155ApprovalForAllIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ERC1155.contract.FilterLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155ApprovalForAllIterator{contract: _ERC1155.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_ERC1155 *ERC1155Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ERC1155ApprovalForAll, account []common.Address, operator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ERC1155.contract.WatchLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155ApprovalForAll)
				if err := _ERC1155.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_ERC1155 *ERC1155Filterer) ParseApprovalForAll(log types.Log) (*ERC1155ApprovalForAll, error) {
	event := new(ERC1155ApprovalForAll)
	if err := _ERC1155.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155TransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the ERC1155 contract.
type ERC1155TransferBatchIterator struct {
	Event *ERC1155TransferBatch // Event containing the contract specifics and raw log

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
func (it *ERC1155TransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155TransferBatch)
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
		it.Event = new(ERC1155TransferBatch)
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
func (it *ERC1155TransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155TransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155TransferBatch represents a TransferBatch event raised by the ERC1155 contract.
type ERC1155TransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_ERC1155 *ERC1155Filterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*ERC1155TransferBatchIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC1155.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155TransferBatchIterator{contract: _ERC1155.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_ERC1155 *ERC1155Filterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *ERC1155TransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC1155.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155TransferBatch)
				if err := _ERC1155.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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

// ParseTransferBatch is a log parse operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_ERC1155 *ERC1155Filterer) ParseTransferBatch(log types.Log) (*ERC1155TransferBatch, error) {
	event := new(ERC1155TransferBatch)
	if err := _ERC1155.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155TransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the ERC1155 contract.
type ERC1155TransferSingleIterator struct {
	Event *ERC1155TransferSingle // Event containing the contract specifics and raw log

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
func (it *ERC1155TransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155TransferSingle)
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
		it.Event = new(ERC1155TransferSingle)
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
func (it *ERC1155TransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155TransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155TransferSingle represents a TransferSingle event raised by the ERC1155 contract.
type ERC1155TransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_ERC1155 *ERC1155Filterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*ERC1155TransferSingleIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC1155.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155TransferSingleIterator{contract: _ERC1155.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_ERC1155 *ERC1155Filterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *ERC1155TransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC1155.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155TransferSingle)
				if err := _ERC1155.contract.UnpackLog(event, "TransferSingle", log); err != nil {
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

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_ERC1155 *ERC1155Filterer) ParseTransferSingle(log types.Log) (*ERC1155TransferSingle, error) {
	event := new(ERC1155TransferSingle)
	if err := _ERC1155.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155URIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the ERC1155 contract.
type ERC1155URIIterator struct {
	Event *ERC1155URI // Event containing the contract specifics and raw log

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
func (it *ERC1155URIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155URI)
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
		it.Event = new(ERC1155URI)
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
func (it *ERC1155URIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155URIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155URI represents a URI event raised by the ERC1155 contract.
type ERC1155URI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_ERC1155 *ERC1155Filterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*ERC1155URIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ERC1155.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155URIIterator{contract: _ERC1155.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_ERC1155 *ERC1155Filterer) WatchURI(opts *bind.WatchOpts, sink chan<- *ERC1155URI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ERC1155.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155URI)
				if err := _ERC1155.contract.UnpackLog(event, "URI", log); err != nil {
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

// ParseURI is a log parse operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_ERC1155 *ERC1155Filterer) ParseURI(log types.Log) (*ERC1155URI, error) {
	event := new(ERC1155URI)
	if err := _ERC1155.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC165ABI is the input ABI used to generate the binding from.
const ERC165ABI = "[{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ERC165FuncSigs maps the 4-byte function signature to its string representation.
var ERC165FuncSigs = map[string]string{
	"01ffc9a7": "supportsInterface(bytes4)",
}

// ERC165 is an auto generated Go binding around an Ethereum contract.
type ERC165 struct {
	ERC165Caller     // Read-only binding to the contract
	ERC165Transactor // Write-only binding to the contract
	ERC165Filterer   // Log filterer for contract events
}

// ERC165Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC165Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC165Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC165Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC165Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC165Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC165Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC165Session struct {
	Contract     *ERC165           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC165CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC165CallerSession struct {
	Contract *ERC165Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC165TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC165TransactorSession struct {
	Contract     *ERC165Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC165Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC165Raw struct {
	Contract *ERC165 // Generic contract binding to access the raw methods on
}

// ERC165CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC165CallerRaw struct {
	Contract *ERC165Caller // Generic read-only contract binding to access the raw methods on
}

// ERC165TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC165TransactorRaw struct {
	Contract *ERC165Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC165 creates a new instance of ERC165, bound to a specific deployed contract.
func NewERC165(address common.Address, backend bind.ContractBackend) (*ERC165, error) {
	contract, err := bindERC165(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC165{ERC165Caller: ERC165Caller{contract: contract}, ERC165Transactor: ERC165Transactor{contract: contract}, ERC165Filterer: ERC165Filterer{contract: contract}}, nil
}

// NewERC165Caller creates a new read-only instance of ERC165, bound to a specific deployed contract.
func NewERC165Caller(address common.Address, caller bind.ContractCaller) (*ERC165Caller, error) {
	contract, err := bindERC165(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC165Caller{contract: contract}, nil
}

// NewERC165Transactor creates a new write-only instance of ERC165, bound to a specific deployed contract.
func NewERC165Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC165Transactor, error) {
	contract, err := bindERC165(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC165Transactor{contract: contract}, nil
}

// NewERC165Filterer creates a new log filterer instance of ERC165, bound to a specific deployed contract.
func NewERC165Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC165Filterer, error) {
	contract, err := bindERC165(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC165Filterer{contract: contract}, nil
}

// bindERC165 binds a generic wrapper to an already deployed contract.
func bindERC165(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC165ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC165 *ERC165Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC165.Contract.ERC165Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC165 *ERC165Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC165.Contract.ERC165Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC165 *ERC165Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC165.Contract.ERC165Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC165 *ERC165CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC165.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC165 *ERC165TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC165.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC165 *ERC165TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC165.Contract.contract.Transact(opts, method, params...)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC165 *ERC165Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ERC165.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC165 *ERC165Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC165.Contract.SupportsInterface(&_ERC165.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC165 *ERC165CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC165.Contract.SupportsInterface(&_ERC165.CallOpts, interfaceId)
}

// IERC1155ABI is the input ABI used to generate the binding from.
const IERC1155ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IERC1155FuncSigs maps the 4-byte function signature to its string representation.
var IERC1155FuncSigs = map[string]string{
	"00fdd58e": "balanceOf(address,uint256)",
	"4e1273f4": "balanceOfBatch(address[],uint256[])",
	"e985e9c5": "isApprovedForAll(address,address)",
	"2eb2c2d6": "safeBatchTransferFrom(address,address,uint256[],uint256[],bytes)",
	"f242432a": "safeTransferFrom(address,address,uint256,uint256,bytes)",
	"a22cb465": "setApprovalForAll(address,bool)",
	"01ffc9a7": "supportsInterface(bytes4)",
}

// IERC1155 is an auto generated Go binding around an Ethereum contract.
type IERC1155 struct {
	IERC1155Caller     // Read-only binding to the contract
	IERC1155Transactor // Write-only binding to the contract
	IERC1155Filterer   // Log filterer for contract events
}

// IERC1155Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC1155Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC1155Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC1155Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC1155Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC1155Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC1155Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC1155Session struct {
	Contract     *IERC1155         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC1155CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC1155CallerSession struct {
	Contract *IERC1155Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IERC1155TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC1155TransactorSession struct {
	Contract     *IERC1155Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IERC1155Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC1155Raw struct {
	Contract *IERC1155 // Generic contract binding to access the raw methods on
}

// IERC1155CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC1155CallerRaw struct {
	Contract *IERC1155Caller // Generic read-only contract binding to access the raw methods on
}

// IERC1155TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC1155TransactorRaw struct {
	Contract *IERC1155Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC1155 creates a new instance of IERC1155, bound to a specific deployed contract.
func NewIERC1155(address common.Address, backend bind.ContractBackend) (*IERC1155, error) {
	contract, err := bindIERC1155(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC1155{IERC1155Caller: IERC1155Caller{contract: contract}, IERC1155Transactor: IERC1155Transactor{contract: contract}, IERC1155Filterer: IERC1155Filterer{contract: contract}}, nil
}

// NewIERC1155Caller creates a new read-only instance of IERC1155, bound to a specific deployed contract.
func NewIERC1155Caller(address common.Address, caller bind.ContractCaller) (*IERC1155Caller, error) {
	contract, err := bindIERC1155(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC1155Caller{contract: contract}, nil
}

// NewIERC1155Transactor creates a new write-only instance of IERC1155, bound to a specific deployed contract.
func NewIERC1155Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC1155Transactor, error) {
	contract, err := bindIERC1155(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC1155Transactor{contract: contract}, nil
}

// NewIERC1155Filterer creates a new log filterer instance of IERC1155, bound to a specific deployed contract.
func NewIERC1155Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC1155Filterer, error) {
	contract, err := bindIERC1155(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC1155Filterer{contract: contract}, nil
}

// bindIERC1155 binds a generic wrapper to an already deployed contract.
func bindIERC1155(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC1155ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC1155 *IERC1155Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC1155.Contract.IERC1155Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC1155 *IERC1155Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC1155.Contract.IERC1155Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC1155 *IERC1155Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC1155.Contract.IERC1155Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC1155 *IERC1155CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC1155.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC1155 *IERC1155TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC1155.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC1155 *IERC1155TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC1155.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_IERC1155 *IERC1155Caller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IERC1155.contract.Call(opts, &out, "balanceOf", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_IERC1155 *IERC1155Session) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _IERC1155.Contract.BalanceOf(&_IERC1155.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_IERC1155 *IERC1155CallerSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _IERC1155.Contract.BalanceOf(&_IERC1155.CallOpts, account, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_IERC1155 *IERC1155Caller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _IERC1155.contract.Call(opts, &out, "balanceOfBatch", accounts, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_IERC1155 *IERC1155Session) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _IERC1155.Contract.BalanceOfBatch(&_IERC1155.CallOpts, accounts, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_IERC1155 *IERC1155CallerSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _IERC1155.Contract.BalanceOfBatch(&_IERC1155.CallOpts, accounts, ids)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_IERC1155 *IERC1155Caller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _IERC1155.contract.Call(opts, &out, "isApprovedForAll", account, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_IERC1155 *IERC1155Session) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _IERC1155.Contract.IsApprovedForAll(&_IERC1155.CallOpts, account, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_IERC1155 *IERC1155CallerSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _IERC1155.Contract.IsApprovedForAll(&_IERC1155.CallOpts, account, operator)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC1155 *IERC1155Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _IERC1155.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC1155 *IERC1155Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC1155.Contract.SupportsInterface(&_IERC1155.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC1155 *IERC1155CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC1155.Contract.SupportsInterface(&_IERC1155.CallOpts, interfaceId)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_IERC1155 *IERC1155Transactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _IERC1155.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_IERC1155 *IERC1155Session) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _IERC1155.Contract.SafeBatchTransferFrom(&_IERC1155.TransactOpts, from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_IERC1155 *IERC1155TransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _IERC1155.Contract.SafeBatchTransferFrom(&_IERC1155.TransactOpts, from, to, ids, amounts, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_IERC1155 *IERC1155Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC1155.contract.Transact(opts, "safeTransferFrom", from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_IERC1155 *IERC1155Session) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC1155.Contract.SafeTransferFrom(&_IERC1155.TransactOpts, from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_IERC1155 *IERC1155TransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC1155.Contract.SafeTransferFrom(&_IERC1155.TransactOpts, from, to, id, amount, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_IERC1155 *IERC1155Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _IERC1155.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_IERC1155 *IERC1155Session) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _IERC1155.Contract.SetApprovalForAll(&_IERC1155.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_IERC1155 *IERC1155TransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _IERC1155.Contract.SetApprovalForAll(&_IERC1155.TransactOpts, operator, approved)
}

// IERC1155ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the IERC1155 contract.
type IERC1155ApprovalForAllIterator struct {
	Event *IERC1155ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *IERC1155ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC1155ApprovalForAll)
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
		it.Event = new(IERC1155ApprovalForAll)
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
func (it *IERC1155ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC1155ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC1155ApprovalForAll represents a ApprovalForAll event raised by the IERC1155 contract.
type IERC1155ApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_IERC1155 *IERC1155Filterer) FilterApprovalForAll(opts *bind.FilterOpts, account []common.Address, operator []common.Address) (*IERC1155ApprovalForAllIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IERC1155.contract.FilterLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &IERC1155ApprovalForAllIterator{contract: _IERC1155.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_IERC1155 *IERC1155Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *IERC1155ApprovalForAll, account []common.Address, operator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IERC1155.contract.WatchLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC1155ApprovalForAll)
				if err := _IERC1155.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_IERC1155 *IERC1155Filterer) ParseApprovalForAll(log types.Log) (*IERC1155ApprovalForAll, error) {
	event := new(IERC1155ApprovalForAll)
	if err := _IERC1155.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC1155TransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the IERC1155 contract.
type IERC1155TransferBatchIterator struct {
	Event *IERC1155TransferBatch // Event containing the contract specifics and raw log

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
func (it *IERC1155TransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC1155TransferBatch)
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
		it.Event = new(IERC1155TransferBatch)
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
func (it *IERC1155TransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC1155TransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC1155TransferBatch represents a TransferBatch event raised by the IERC1155 contract.
type IERC1155TransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_IERC1155 *IERC1155Filterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*IERC1155TransferBatchIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC1155.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC1155TransferBatchIterator{contract: _IERC1155.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_IERC1155 *IERC1155Filterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *IERC1155TransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC1155.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC1155TransferBatch)
				if err := _IERC1155.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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

// ParseTransferBatch is a log parse operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_IERC1155 *IERC1155Filterer) ParseTransferBatch(log types.Log) (*IERC1155TransferBatch, error) {
	event := new(IERC1155TransferBatch)
	if err := _IERC1155.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC1155TransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the IERC1155 contract.
type IERC1155TransferSingleIterator struct {
	Event *IERC1155TransferSingle // Event containing the contract specifics and raw log

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
func (it *IERC1155TransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC1155TransferSingle)
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
		it.Event = new(IERC1155TransferSingle)
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
func (it *IERC1155TransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC1155TransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC1155TransferSingle represents a TransferSingle event raised by the IERC1155 contract.
type IERC1155TransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_IERC1155 *IERC1155Filterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*IERC1155TransferSingleIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC1155.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC1155TransferSingleIterator{contract: _IERC1155.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_IERC1155 *IERC1155Filterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *IERC1155TransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC1155.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC1155TransferSingle)
				if err := _IERC1155.contract.UnpackLog(event, "TransferSingle", log); err != nil {
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

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_IERC1155 *IERC1155Filterer) ParseTransferSingle(log types.Log) (*IERC1155TransferSingle, error) {
	event := new(IERC1155TransferSingle)
	if err := _IERC1155.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC1155URIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the IERC1155 contract.
type IERC1155URIIterator struct {
	Event *IERC1155URI // Event containing the contract specifics and raw log

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
func (it *IERC1155URIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC1155URI)
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
		it.Event = new(IERC1155URI)
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
func (it *IERC1155URIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC1155URIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC1155URI represents a URI event raised by the IERC1155 contract.
type IERC1155URI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_IERC1155 *IERC1155Filterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*IERC1155URIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IERC1155.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &IERC1155URIIterator{contract: _IERC1155.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_IERC1155 *IERC1155Filterer) WatchURI(opts *bind.WatchOpts, sink chan<- *IERC1155URI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IERC1155.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC1155URI)
				if err := _IERC1155.contract.UnpackLog(event, "URI", log); err != nil {
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

// ParseURI is a log parse operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_IERC1155 *IERC1155Filterer) ParseURI(log types.Log) (*IERC1155URI, error) {
	event := new(IERC1155URI)
	if err := _IERC1155.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC1155MetadataURIABI is the input ABI used to generate the binding from.
const IERC1155MetadataURIABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IERC1155MetadataURIFuncSigs maps the 4-byte function signature to its string representation.
var IERC1155MetadataURIFuncSigs = map[string]string{
	"00fdd58e": "balanceOf(address,uint256)",
	"4e1273f4": "balanceOfBatch(address[],uint256[])",
	"e985e9c5": "isApprovedForAll(address,address)",
	"2eb2c2d6": "safeBatchTransferFrom(address,address,uint256[],uint256[],bytes)",
	"f242432a": "safeTransferFrom(address,address,uint256,uint256,bytes)",
	"a22cb465": "setApprovalForAll(address,bool)",
	"01ffc9a7": "supportsInterface(bytes4)",
	"0e89341c": "uri(uint256)",
}

// IERC1155MetadataURI is an auto generated Go binding around an Ethereum contract.
type IERC1155MetadataURI struct {
	IERC1155MetadataURICaller     // Read-only binding to the contract
	IERC1155MetadataURITransactor // Write-only binding to the contract
	IERC1155MetadataURIFilterer   // Log filterer for contract events
}

// IERC1155MetadataURICaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC1155MetadataURICaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC1155MetadataURITransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC1155MetadataURITransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC1155MetadataURIFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC1155MetadataURIFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC1155MetadataURISession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC1155MetadataURISession struct {
	Contract     *IERC1155MetadataURI // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IERC1155MetadataURICallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC1155MetadataURICallerSession struct {
	Contract *IERC1155MetadataURICaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// IERC1155MetadataURITransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC1155MetadataURITransactorSession struct {
	Contract     *IERC1155MetadataURITransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// IERC1155MetadataURIRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC1155MetadataURIRaw struct {
	Contract *IERC1155MetadataURI // Generic contract binding to access the raw methods on
}

// IERC1155MetadataURICallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC1155MetadataURICallerRaw struct {
	Contract *IERC1155MetadataURICaller // Generic read-only contract binding to access the raw methods on
}

// IERC1155MetadataURITransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC1155MetadataURITransactorRaw struct {
	Contract *IERC1155MetadataURITransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC1155MetadataURI creates a new instance of IERC1155MetadataURI, bound to a specific deployed contract.
func NewIERC1155MetadataURI(address common.Address, backend bind.ContractBackend) (*IERC1155MetadataURI, error) {
	contract, err := bindIERC1155MetadataURI(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC1155MetadataURI{IERC1155MetadataURICaller: IERC1155MetadataURICaller{contract: contract}, IERC1155MetadataURITransactor: IERC1155MetadataURITransactor{contract: contract}, IERC1155MetadataURIFilterer: IERC1155MetadataURIFilterer{contract: contract}}, nil
}

// NewIERC1155MetadataURICaller creates a new read-only instance of IERC1155MetadataURI, bound to a specific deployed contract.
func NewIERC1155MetadataURICaller(address common.Address, caller bind.ContractCaller) (*IERC1155MetadataURICaller, error) {
	contract, err := bindIERC1155MetadataURI(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC1155MetadataURICaller{contract: contract}, nil
}

// NewIERC1155MetadataURITransactor creates a new write-only instance of IERC1155MetadataURI, bound to a specific deployed contract.
func NewIERC1155MetadataURITransactor(address common.Address, transactor bind.ContractTransactor) (*IERC1155MetadataURITransactor, error) {
	contract, err := bindIERC1155MetadataURI(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC1155MetadataURITransactor{contract: contract}, nil
}

// NewIERC1155MetadataURIFilterer creates a new log filterer instance of IERC1155MetadataURI, bound to a specific deployed contract.
func NewIERC1155MetadataURIFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC1155MetadataURIFilterer, error) {
	contract, err := bindIERC1155MetadataURI(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC1155MetadataURIFilterer{contract: contract}, nil
}

// bindIERC1155MetadataURI binds a generic wrapper to an already deployed contract.
func bindIERC1155MetadataURI(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC1155MetadataURIABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC1155MetadataURI *IERC1155MetadataURIRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC1155MetadataURI.Contract.IERC1155MetadataURICaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC1155MetadataURI *IERC1155MetadataURIRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC1155MetadataURI.Contract.IERC1155MetadataURITransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC1155MetadataURI *IERC1155MetadataURIRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC1155MetadataURI.Contract.IERC1155MetadataURITransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC1155MetadataURI *IERC1155MetadataURICallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC1155MetadataURI.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC1155MetadataURI *IERC1155MetadataURITransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC1155MetadataURI.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC1155MetadataURI *IERC1155MetadataURITransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC1155MetadataURI.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_IERC1155MetadataURI *IERC1155MetadataURICaller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IERC1155MetadataURI.contract.Call(opts, &out, "balanceOf", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_IERC1155MetadataURI *IERC1155MetadataURISession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _IERC1155MetadataURI.Contract.BalanceOf(&_IERC1155MetadataURI.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_IERC1155MetadataURI *IERC1155MetadataURICallerSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _IERC1155MetadataURI.Contract.BalanceOf(&_IERC1155MetadataURI.CallOpts, account, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_IERC1155MetadataURI *IERC1155MetadataURICaller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _IERC1155MetadataURI.contract.Call(opts, &out, "balanceOfBatch", accounts, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_IERC1155MetadataURI *IERC1155MetadataURISession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _IERC1155MetadataURI.Contract.BalanceOfBatch(&_IERC1155MetadataURI.CallOpts, accounts, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_IERC1155MetadataURI *IERC1155MetadataURICallerSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _IERC1155MetadataURI.Contract.BalanceOfBatch(&_IERC1155MetadataURI.CallOpts, accounts, ids)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_IERC1155MetadataURI *IERC1155MetadataURICaller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _IERC1155MetadataURI.contract.Call(opts, &out, "isApprovedForAll", account, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_IERC1155MetadataURI *IERC1155MetadataURISession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _IERC1155MetadataURI.Contract.IsApprovedForAll(&_IERC1155MetadataURI.CallOpts, account, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_IERC1155MetadataURI *IERC1155MetadataURICallerSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _IERC1155MetadataURI.Contract.IsApprovedForAll(&_IERC1155MetadataURI.CallOpts, account, operator)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC1155MetadataURI *IERC1155MetadataURICaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _IERC1155MetadataURI.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC1155MetadataURI *IERC1155MetadataURISession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC1155MetadataURI.Contract.SupportsInterface(&_IERC1155MetadataURI.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC1155MetadataURI *IERC1155MetadataURICallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC1155MetadataURI.Contract.SupportsInterface(&_IERC1155MetadataURI.CallOpts, interfaceId)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 id) view returns(string)
func (_IERC1155MetadataURI *IERC1155MetadataURICaller) Uri(opts *bind.CallOpts, id *big.Int) (string, error) {
	var out []interface{}
	err := _IERC1155MetadataURI.contract.Call(opts, &out, "uri", id)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 id) view returns(string)
func (_IERC1155MetadataURI *IERC1155MetadataURISession) Uri(id *big.Int) (string, error) {
	return _IERC1155MetadataURI.Contract.Uri(&_IERC1155MetadataURI.CallOpts, id)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 id) view returns(string)
func (_IERC1155MetadataURI *IERC1155MetadataURICallerSession) Uri(id *big.Int) (string, error) {
	return _IERC1155MetadataURI.Contract.Uri(&_IERC1155MetadataURI.CallOpts, id)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_IERC1155MetadataURI *IERC1155MetadataURITransactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _IERC1155MetadataURI.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_IERC1155MetadataURI *IERC1155MetadataURISession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _IERC1155MetadataURI.Contract.SafeBatchTransferFrom(&_IERC1155MetadataURI.TransactOpts, from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_IERC1155MetadataURI *IERC1155MetadataURITransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _IERC1155MetadataURI.Contract.SafeBatchTransferFrom(&_IERC1155MetadataURI.TransactOpts, from, to, ids, amounts, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_IERC1155MetadataURI *IERC1155MetadataURITransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC1155MetadataURI.contract.Transact(opts, "safeTransferFrom", from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_IERC1155MetadataURI *IERC1155MetadataURISession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC1155MetadataURI.Contract.SafeTransferFrom(&_IERC1155MetadataURI.TransactOpts, from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_IERC1155MetadataURI *IERC1155MetadataURITransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC1155MetadataURI.Contract.SafeTransferFrom(&_IERC1155MetadataURI.TransactOpts, from, to, id, amount, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_IERC1155MetadataURI *IERC1155MetadataURITransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _IERC1155MetadataURI.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_IERC1155MetadataURI *IERC1155MetadataURISession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _IERC1155MetadataURI.Contract.SetApprovalForAll(&_IERC1155MetadataURI.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_IERC1155MetadataURI *IERC1155MetadataURITransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _IERC1155MetadataURI.Contract.SetApprovalForAll(&_IERC1155MetadataURI.TransactOpts, operator, approved)
}

// IERC1155MetadataURIApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the IERC1155MetadataURI contract.
type IERC1155MetadataURIApprovalForAllIterator struct {
	Event *IERC1155MetadataURIApprovalForAll // Event containing the contract specifics and raw log

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
func (it *IERC1155MetadataURIApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC1155MetadataURIApprovalForAll)
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
		it.Event = new(IERC1155MetadataURIApprovalForAll)
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
func (it *IERC1155MetadataURIApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC1155MetadataURIApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC1155MetadataURIApprovalForAll represents a ApprovalForAll event raised by the IERC1155MetadataURI contract.
type IERC1155MetadataURIApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_IERC1155MetadataURI *IERC1155MetadataURIFilterer) FilterApprovalForAll(opts *bind.FilterOpts, account []common.Address, operator []common.Address) (*IERC1155MetadataURIApprovalForAllIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IERC1155MetadataURI.contract.FilterLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &IERC1155MetadataURIApprovalForAllIterator{contract: _IERC1155MetadataURI.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_IERC1155MetadataURI *IERC1155MetadataURIFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *IERC1155MetadataURIApprovalForAll, account []common.Address, operator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IERC1155MetadataURI.contract.WatchLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC1155MetadataURIApprovalForAll)
				if err := _IERC1155MetadataURI.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_IERC1155MetadataURI *IERC1155MetadataURIFilterer) ParseApprovalForAll(log types.Log) (*IERC1155MetadataURIApprovalForAll, error) {
	event := new(IERC1155MetadataURIApprovalForAll)
	if err := _IERC1155MetadataURI.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC1155MetadataURITransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the IERC1155MetadataURI contract.
type IERC1155MetadataURITransferBatchIterator struct {
	Event *IERC1155MetadataURITransferBatch // Event containing the contract specifics and raw log

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
func (it *IERC1155MetadataURITransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC1155MetadataURITransferBatch)
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
		it.Event = new(IERC1155MetadataURITransferBatch)
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
func (it *IERC1155MetadataURITransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC1155MetadataURITransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC1155MetadataURITransferBatch represents a TransferBatch event raised by the IERC1155MetadataURI contract.
type IERC1155MetadataURITransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_IERC1155MetadataURI *IERC1155MetadataURIFilterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*IERC1155MetadataURITransferBatchIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC1155MetadataURI.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC1155MetadataURITransferBatchIterator{contract: _IERC1155MetadataURI.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_IERC1155MetadataURI *IERC1155MetadataURIFilterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *IERC1155MetadataURITransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC1155MetadataURI.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC1155MetadataURITransferBatch)
				if err := _IERC1155MetadataURI.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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

// ParseTransferBatch is a log parse operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_IERC1155MetadataURI *IERC1155MetadataURIFilterer) ParseTransferBatch(log types.Log) (*IERC1155MetadataURITransferBatch, error) {
	event := new(IERC1155MetadataURITransferBatch)
	if err := _IERC1155MetadataURI.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC1155MetadataURITransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the IERC1155MetadataURI contract.
type IERC1155MetadataURITransferSingleIterator struct {
	Event *IERC1155MetadataURITransferSingle // Event containing the contract specifics and raw log

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
func (it *IERC1155MetadataURITransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC1155MetadataURITransferSingle)
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
		it.Event = new(IERC1155MetadataURITransferSingle)
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
func (it *IERC1155MetadataURITransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC1155MetadataURITransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC1155MetadataURITransferSingle represents a TransferSingle event raised by the IERC1155MetadataURI contract.
type IERC1155MetadataURITransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_IERC1155MetadataURI *IERC1155MetadataURIFilterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*IERC1155MetadataURITransferSingleIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC1155MetadataURI.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC1155MetadataURITransferSingleIterator{contract: _IERC1155MetadataURI.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_IERC1155MetadataURI *IERC1155MetadataURIFilterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *IERC1155MetadataURITransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC1155MetadataURI.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC1155MetadataURITransferSingle)
				if err := _IERC1155MetadataURI.contract.UnpackLog(event, "TransferSingle", log); err != nil {
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

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_IERC1155MetadataURI *IERC1155MetadataURIFilterer) ParseTransferSingle(log types.Log) (*IERC1155MetadataURITransferSingle, error) {
	event := new(IERC1155MetadataURITransferSingle)
	if err := _IERC1155MetadataURI.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC1155MetadataURIURIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the IERC1155MetadataURI contract.
type IERC1155MetadataURIURIIterator struct {
	Event *IERC1155MetadataURIURI // Event containing the contract specifics and raw log

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
func (it *IERC1155MetadataURIURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC1155MetadataURIURI)
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
		it.Event = new(IERC1155MetadataURIURI)
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
func (it *IERC1155MetadataURIURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC1155MetadataURIURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC1155MetadataURIURI represents a URI event raised by the IERC1155MetadataURI contract.
type IERC1155MetadataURIURI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_IERC1155MetadataURI *IERC1155MetadataURIFilterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*IERC1155MetadataURIURIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IERC1155MetadataURI.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &IERC1155MetadataURIURIIterator{contract: _IERC1155MetadataURI.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_IERC1155MetadataURI *IERC1155MetadataURIFilterer) WatchURI(opts *bind.WatchOpts, sink chan<- *IERC1155MetadataURIURI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IERC1155MetadataURI.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC1155MetadataURIURI)
				if err := _IERC1155MetadataURI.contract.UnpackLog(event, "URI", log); err != nil {
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

// ParseURI is a log parse operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_IERC1155MetadataURI *IERC1155MetadataURIFilterer) ParseURI(log types.Log) (*IERC1155MetadataURIURI, error) {
	event := new(IERC1155MetadataURIURI)
	if err := _IERC1155MetadataURI.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC1155ReceiverABI is the input ABI used to generate the binding from.
const IERC1155ReceiverABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IERC1155ReceiverFuncSigs maps the 4-byte function signature to its string representation.
var IERC1155ReceiverFuncSigs = map[string]string{
	"bc197c81": "onERC1155BatchReceived(address,address,uint256[],uint256[],bytes)",
	"f23a6e61": "onERC1155Received(address,address,uint256,uint256,bytes)",
	"01ffc9a7": "supportsInterface(bytes4)",
}

// IERC1155Receiver is an auto generated Go binding around an Ethereum contract.
type IERC1155Receiver struct {
	IERC1155ReceiverCaller     // Read-only binding to the contract
	IERC1155ReceiverTransactor // Write-only binding to the contract
	IERC1155ReceiverFilterer   // Log filterer for contract events
}

// IERC1155ReceiverCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC1155ReceiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC1155ReceiverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC1155ReceiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC1155ReceiverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC1155ReceiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC1155ReceiverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC1155ReceiverSession struct {
	Contract     *IERC1155Receiver // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC1155ReceiverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC1155ReceiverCallerSession struct {
	Contract *IERC1155ReceiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IERC1155ReceiverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC1155ReceiverTransactorSession struct {
	Contract     *IERC1155ReceiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IERC1155ReceiverRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC1155ReceiverRaw struct {
	Contract *IERC1155Receiver // Generic contract binding to access the raw methods on
}

// IERC1155ReceiverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC1155ReceiverCallerRaw struct {
	Contract *IERC1155ReceiverCaller // Generic read-only contract binding to access the raw methods on
}

// IERC1155ReceiverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC1155ReceiverTransactorRaw struct {
	Contract *IERC1155ReceiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC1155Receiver creates a new instance of IERC1155Receiver, bound to a specific deployed contract.
func NewIERC1155Receiver(address common.Address, backend bind.ContractBackend) (*IERC1155Receiver, error) {
	contract, err := bindIERC1155Receiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC1155Receiver{IERC1155ReceiverCaller: IERC1155ReceiverCaller{contract: contract}, IERC1155ReceiverTransactor: IERC1155ReceiverTransactor{contract: contract}, IERC1155ReceiverFilterer: IERC1155ReceiverFilterer{contract: contract}}, nil
}

// NewIERC1155ReceiverCaller creates a new read-only instance of IERC1155Receiver, bound to a specific deployed contract.
func NewIERC1155ReceiverCaller(address common.Address, caller bind.ContractCaller) (*IERC1155ReceiverCaller, error) {
	contract, err := bindIERC1155Receiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC1155ReceiverCaller{contract: contract}, nil
}

// NewIERC1155ReceiverTransactor creates a new write-only instance of IERC1155Receiver, bound to a specific deployed contract.
func NewIERC1155ReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC1155ReceiverTransactor, error) {
	contract, err := bindIERC1155Receiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC1155ReceiverTransactor{contract: contract}, nil
}

// NewIERC1155ReceiverFilterer creates a new log filterer instance of IERC1155Receiver, bound to a specific deployed contract.
func NewIERC1155ReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC1155ReceiverFilterer, error) {
	contract, err := bindIERC1155Receiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC1155ReceiverFilterer{contract: contract}, nil
}

// bindIERC1155Receiver binds a generic wrapper to an already deployed contract.
func bindIERC1155Receiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC1155ReceiverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC1155Receiver *IERC1155ReceiverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC1155Receiver.Contract.IERC1155ReceiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC1155Receiver *IERC1155ReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC1155Receiver.Contract.IERC1155ReceiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC1155Receiver *IERC1155ReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC1155Receiver.Contract.IERC1155ReceiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC1155Receiver *IERC1155ReceiverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC1155Receiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC1155Receiver *IERC1155ReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC1155Receiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC1155Receiver *IERC1155ReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC1155Receiver.Contract.contract.Transact(opts, method, params...)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC1155Receiver *IERC1155ReceiverCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _IERC1155Receiver.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC1155Receiver *IERC1155ReceiverSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC1155Receiver.Contract.SupportsInterface(&_IERC1155Receiver.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC1155Receiver *IERC1155ReceiverCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC1155Receiver.Contract.SupportsInterface(&_IERC1155Receiver.CallOpts, interfaceId)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address operator, address from, uint256[] ids, uint256[] values, bytes data) returns(bytes4)
func (_IERC1155Receiver *IERC1155ReceiverTransactor) OnERC1155BatchReceived(opts *bind.TransactOpts, operator common.Address, from common.Address, ids []*big.Int, values []*big.Int, data []byte) (*types.Transaction, error) {
	return _IERC1155Receiver.contract.Transact(opts, "onERC1155BatchReceived", operator, from, ids, values, data)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address operator, address from, uint256[] ids, uint256[] values, bytes data) returns(bytes4)
func (_IERC1155Receiver *IERC1155ReceiverSession) OnERC1155BatchReceived(operator common.Address, from common.Address, ids []*big.Int, values []*big.Int, data []byte) (*types.Transaction, error) {
	return _IERC1155Receiver.Contract.OnERC1155BatchReceived(&_IERC1155Receiver.TransactOpts, operator, from, ids, values, data)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address operator, address from, uint256[] ids, uint256[] values, bytes data) returns(bytes4)
func (_IERC1155Receiver *IERC1155ReceiverTransactorSession) OnERC1155BatchReceived(operator common.Address, from common.Address, ids []*big.Int, values []*big.Int, data []byte) (*types.Transaction, error) {
	return _IERC1155Receiver.Contract.OnERC1155BatchReceived(&_IERC1155Receiver.TransactOpts, operator, from, ids, values, data)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address operator, address from, uint256 id, uint256 value, bytes data) returns(bytes4)
func (_IERC1155Receiver *IERC1155ReceiverTransactor) OnERC1155Received(opts *bind.TransactOpts, operator common.Address, from common.Address, id *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC1155Receiver.contract.Transact(opts, "onERC1155Received", operator, from, id, value, data)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address operator, address from, uint256 id, uint256 value, bytes data) returns(bytes4)
func (_IERC1155Receiver *IERC1155ReceiverSession) OnERC1155Received(operator common.Address, from common.Address, id *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC1155Receiver.Contract.OnERC1155Received(&_IERC1155Receiver.TransactOpts, operator, from, id, value, data)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address operator, address from, uint256 id, uint256 value, bytes data) returns(bytes4)
func (_IERC1155Receiver *IERC1155ReceiverTransactorSession) OnERC1155Received(operator common.Address, from common.Address, id *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC1155Receiver.Contract.OnERC1155Received(&_IERC1155Receiver.TransactOpts, operator, from, id, value, data)
}

// IERC165ABI is the input ABI used to generate the binding from.
const IERC165ABI = "[{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IERC165FuncSigs maps the 4-byte function signature to its string representation.
var IERC165FuncSigs = map[string]string{
	"01ffc9a7": "supportsInterface(bytes4)",
}

// IERC165 is an auto generated Go binding around an Ethereum contract.
type IERC165 struct {
	IERC165Caller     // Read-only binding to the contract
	IERC165Transactor // Write-only binding to the contract
	IERC165Filterer   // Log filterer for contract events
}

// IERC165Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC165Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC165Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC165Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC165Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC165Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC165Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC165Session struct {
	Contract     *IERC165          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC165CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC165CallerSession struct {
	Contract *IERC165Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IERC165TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC165TransactorSession struct {
	Contract     *IERC165Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IERC165Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC165Raw struct {
	Contract *IERC165 // Generic contract binding to access the raw methods on
}

// IERC165CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC165CallerRaw struct {
	Contract *IERC165Caller // Generic read-only contract binding to access the raw methods on
}

// IERC165TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC165TransactorRaw struct {
	Contract *IERC165Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC165 creates a new instance of IERC165, bound to a specific deployed contract.
func NewIERC165(address common.Address, backend bind.ContractBackend) (*IERC165, error) {
	contract, err := bindIERC165(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC165{IERC165Caller: IERC165Caller{contract: contract}, IERC165Transactor: IERC165Transactor{contract: contract}, IERC165Filterer: IERC165Filterer{contract: contract}}, nil
}

// NewIERC165Caller creates a new read-only instance of IERC165, bound to a specific deployed contract.
func NewIERC165Caller(address common.Address, caller bind.ContractCaller) (*IERC165Caller, error) {
	contract, err := bindIERC165(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC165Caller{contract: contract}, nil
}

// NewIERC165Transactor creates a new write-only instance of IERC165, bound to a specific deployed contract.
func NewIERC165Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC165Transactor, error) {
	contract, err := bindIERC165(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC165Transactor{contract: contract}, nil
}

// NewIERC165Filterer creates a new log filterer instance of IERC165, bound to a specific deployed contract.
func NewIERC165Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC165Filterer, error) {
	contract, err := bindIERC165(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC165Filterer{contract: contract}, nil
}

// bindIERC165 binds a generic wrapper to an already deployed contract.
func bindIERC165(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC165ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC165 *IERC165Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC165.Contract.IERC165Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC165 *IERC165Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC165.Contract.IERC165Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC165 *IERC165Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC165.Contract.IERC165Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC165 *IERC165CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC165.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC165 *IERC165TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC165.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC165 *IERC165TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC165.Contract.contract.Transact(opts, method, params...)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC165 *IERC165Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _IERC165.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC165 *IERC165Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC165.Contract.SupportsInterface(&_IERC165.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC165 *IERC165CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC165.Contract.SupportsInterface(&_IERC165.CallOpts, interfaceId)
}

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

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
var SafeMathBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212208295b5a03db604cabb970ed55bb389ed4ab8dcfa95a01c6ca5299409a42b9cf564736f6c63430007050033"

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
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
var LockingBin = "0x608060405234801561001057600080fd5b50604051610ef1380380610ef183398101604081905261002f916100db565b60008054336001600160a01b031991821617909155600180549091166001600160a01b03841617905561006181610068565b5050610153565b6000546001600160a01b0316331461009b5760405162461bcd60e51b815260040161009290610113565b60405180910390fd5b60028190556040517f97e0c5afc28fab93c1626222b791d272e9d38ca5e480e54ce8cbb7aeb615f4a5906100d090839061014a565b60405180910390a150565b600080604083850312156100ed578182fd5b82516001600160a01b0381168114610103578283fd5b6020939093015192949293505050565b60208082526013908201527f556e617574686f72697365642061636365737300000000000000000000000000604082015260600190565b90815260200190565b610d8f806101626000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c8063a4c0ed3611610071578063a4c0ed3614610136578063a69df4b514610149578063c55dae6314610151578063ce606ee014610166578063ece531321461016e578063f83d08ba14610181576100a9565b8063102ff0b3146100ae5780633c37db63146100cc57806349dcc4eb146100e15780634a4fbeec146100f45780636ecb9d5014610114575b600080fd5b6100b6610189565b6040516100c39190610d23565b60405180910390f35b6100df6100da366004610a22565b61018f565b005b6100b66100ef36600461092c565b610202565b61010761010236600461092c565b61028e565b6040516100c39190610aa3565b61012761012236600461092c565b6102d5565b6040516100c393929190610d2c565b6100df61014436600461094f565b6102f9565b6100df610353565b610159610542565b6040516100c39190610a52565b610159610551565b6100df61017c36600461092c565b610560565b6100df610707565b60025481565b6000546001600160a01b031633146101c25760405162461bcd60e51b81526004016101b990610c24565b60405180910390fd5b60028190556040517f97e0c5afc28fab93c1626222b791d272e9d38ca5e480e54ce8cbb7aeb615f4a5906101f7908390610d23565b60405180910390a150565b6001600160a01b03811660009081526003602052604081206002015460ff1661023d5760405162461bcd60e51b81526004016101b990610cec565b6001600160a01b038216600090815260036020526040902054806102735760405162461bcd60e51b81526004016101b990610ca7565b428110610284576000915050610289565b420390505b919050565b6001600160a01b03811660009081526003602052604081206002015460ff1680156102cf57506001600160a01b038216600090815260036020526040902054155b92915050565b60036020526000908152604090208054600182015460029092015490919060ff1683565b6001546001600160a01b031633146103235760405162461bcd60e51b81526004016101b990610c7a565b60025482146103445760405162461bcd60e51b81526004016101b990610b2c565b61034e8383610814565b505050565b61035b610909565b50336000908152600360209081526040918290208251606081018452815481526001820154928101929092526002015460ff161515918101829052906103b35760405162461bcd60e51b81526004016101b990610b80565b60208101516103d45760405162461bcd60e51b81526004016101b990610bb7565b805180156104f057428111156103fc5760405162461bcd60e51b81526004016101b990610ae5565b600154602083015160405163a9059cbb60e01b81526001600160a01b039092169163a9059cbb9161043291339190600401610a8a565b602060405180830381600087803b15801561044c57600080fd5b505af1158015610460573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104849190610a02565b5033600081815260036020908152604080832083815560018101939093556002909201805460ff1916905584015190517ffe3b7acaad2fd3b2a0c954f99523de4322c0792bcfa0051a3e284e05e8530e4c926104e1929091610a8a565b60405180910390a15050610540565b33600081815260036020526040908190206206978042019055517f9626a556bdcacf287e1acddc78733a299483d4511693e5211ffaacad062137bc9161053591610a52565b60405180910390a150505b565b6001546001600160a01b031681565b6000546001600160a01b031681565b6000546001600160a01b0316331461058a5760405162461bcd60e51b81526004016101b990610b5a565b6001600160a01b0381166105d857600080546040516001600160a01b03909116914780156108fc02929091818181858888f193505050501580156105d2573d6000803e3d6000fd5b50610704565b6001546001600160a01b03828116911614156106065760405162461bcd60e51b81526004016101b990610be2565b6000546040516370a0823160e01b81526001600160a01b038084169263a9059cbb9291169083906370a0823190610641903090600401610a52565b602060405180830381600087803b15801561065b57600080fd5b505af115801561066f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106939190610a3a565b6040518363ffffffff1660e01b81526004016106b0929190610a8a565b602060405180830381600087803b1580156106ca57600080fd5b505af11580156106de573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107029190610a02565b505b50565b3360009081526003602052604090205415610769573360008181526003602052604080822091909155517ffa044b7b93a40365dc68049797c2eb06918523d694e5d56e406cac3eb35578e59161075c91610a52565b60405180910390a1610540565b6002546001546040516323b872dd60e01b81526001600160a01b03909116906323b872dd906107a090339030908690600401610a66565b602060405180830381600087803b1580156107ba57600080fd5b505af11580156107ce573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107f29190610a02565b61080e5760405162461bcd60e51b81526004016101b990610c51565b61070433825b61081c610909565b506001600160a01b0382166000908152600360209081526040918290208251606081018452815481526001820154928101929092526002015460ff16158015928201929092529061087f5760405162461bcd60e51b81526004016101b990610aae565b6001604082810182815260208085018681526001600160a01b038816600090815260039092529083902085518155905193810193909355516002909201805460ff191692151592909217909155517fb00e71ae128757f10a89bed65c351adc01e3a28456245a1e31f738f364bf2491906108fc9085908590610a8a565b60405180910390a1505050565b604051806060016040528060008152602001600081526020016000151581525090565b60006020828403121561093d578081fd5b813561094881610d44565b9392505050565b600080600060608486031215610963578182fd5b833561096e81610d44565b92506020848101359250604085013567ffffffffffffffff80821115610992578384fd5b818701915087601f8301126109a5578384fd5b8135818111156109b157fe5b604051601f8201601f19168101850183811182821017156109ce57fe5b60405281815283820185018a10156109e4578586fd5b81858501868301378585838301015280955050505050509250925092565b600060208284031215610a13578081fd5b81518015158114610948578182fd5b600060208284031215610a33578081fd5b5035919050565b600060208284031215610a4b578081fd5b5051919050565b6001600160a01b0391909116815260200190565b6001600160a01b039384168152919092166020820152604081019190915260600190565b6001600160a01b03929092168252602082015260400190565b901515815260200190565b6020808252601e908201527f596f7520616c72656164792068617665206c6f636b656420746f6b656e730000604082015260600190565b60208082526027908201527f72656c6561736520616c72656164792072657175657374656420627574206e6f6040820152667420726561647960c81b606082015260800190565b602080825260149082015273125b98dbdc9c9958dd081d985b1d59481cd95b9d60621b604082015260600190565b6020808252600c908201526b155b985d5d1a1bdc9a5cd95960a21b604082015260600190565b6020808252601d908201527f596f7520646f206e6f742068617665206c6f636b656420746f6b656e73000000604082015260600190565b6020808252601190820152704e6f7468696e6720746f20756e6c6f636b60781b604082015260600190565b60208082526022908201527f596f752063616e6e6f7420776974686472617720746865206261736520746f6b60408201526132b760f11b606082015260800190565b602080825260139082015272556e617574686f72697365642061636365737360681b604082015260600190565b6020808252600f908201526e1d1c985b9cd9995c8819985a5b1959608a1b604082015260600190565b602080825260139082015272556e617574686f726973656420736f7572636560681b604082015260600190565b60208082526025908201527f4f776e657220686173206e6f74207265717565737465642066756e64732072656040820152646c6561736560d81b606082015260800190565b60208082526019908201527f4f776e657220686173206e6f2066756e6473206c6f636b656400000000000000604082015260600190565b90815260200190565b92835260208301919091521515604082015260600190565b6001600160a01b038116811461070457600080fdfea2646970667358221220bd7b201066a604c85d2ddd22d5a7de43f1e474023c73503247b073bfc511cbb264736f6c63430007050033"

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

// PMintyMultiTokenABI is the input ABI used to generate the binding from.
const PMintyMultiTokenABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"saleContract\",\"type\":\"address\"},{\"internalType\":\"contractlocking[]\",\"name\":\"_locs\",\"type\":\"address[]\"},{\"internalType\":\"string\",\"name\":\"_lockError\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_royaltyPerMille\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"share\",\"type\":\"uint256\"}],\"internalType\":\"structPoolEntry[]\",\"name\":\"_initialEntries\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"share\",\"type\":\"uint256\"}],\"internalType\":\"structPoolEntry[]\",\"name\":\"_resaleEntries\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"OperatorSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"saleNumber\",\"type\":\"uint256\"}],\"name\":\"getRoyalties\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"share\",\"type\":\"uint256\"}],\"internalType\":\"structPoolEntry[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"initialPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"share\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockError\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"locs\",\"outputs\":[{\"internalType\":\"contractlocking\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"hash\",\"type\":\"string\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"quantities\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"hashes\",\"type\":\"string[]\"}],\"name\":\"mintBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"minted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"resalePool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"share\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"royaltyPerMille\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setAuth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_base\",\"type\":\"string\"}],\"name\":\"setBase\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_uri\",\"type\":\"string\"}],\"name\":\"setContractURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"}],\"name\":\"validateBuyer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// PMintyMultiTokenFuncSigs maps the 4-byte function signature to its string representation.
var PMintyMultiTokenFuncSigs = map[string]string{
	"00fdd58e": "balanceOf(address,uint256)",
	"4e1273f4": "balanceOfBatch(address[],uint256[])",
	"e8a3d485": "contractURI()",
	"bb3bafd6": "getRoyalties(uint256)",
	"f637e019": "initialPool(uint256)",
	"e985e9c5": "isApprovedForAll(address,address)",
	"88e485c7": "lockError()",
	"79f3ca18": "locs(uint256)",
	"c6e64e53": "mint(uint256,uint256,string)",
	"c51bb6dc": "mintBatch(uint256[],uint256[],string[])",
	"7dc0bf3f": "minted(uint256)",
	"8da5cb5b": "owner()",
	"43bc5902": "resalePool(uint256)",
	"e8f3359f": "royaltyPerMille()",
	"2eb2c2d6": "safeBatchTransferFrom(address,address,uint256[],uint256[],bytes)",
	"f242432a": "safeTransferFrom(address,address,uint256,uint256,bytes)",
	"a22cb465": "setApprovalForAll(address,bool)",
	"0b44a218": "setAuth(address,bool)",
	"664aa26b": "setBase(string)",
	"938e3d7b": "setContractURI(string)",
	"01ffc9a7": "supportsInterface(bytes4)",
	"0e89341c": "uri(uint256)",
	"5773bd07": "validateBuyer(address)",
}

// PMintyMultiTokenBin is the compiled bytecode used for deploying new contracts.
var PMintyMultiTokenBin = "0x60806040523480156200001157600080fd5b5060405162002c6438038062002c648339810160408190526200003491620006ad565b604080516020810190915260008152620000556301ffc9a760e01b62000351565b6200006081620003a3565b62000072636cdb3d1360e11b62000351565b620000846303a24d0760e21b62000351565b50600780546001600160a01b0319166001600160a01b0389811691909117909155861660009081526008602052604090819020805460ff1916600190811790915590517f1a594081ae893ab78e67d9b9e843547318164322d32c65369d78a96172d9dc8f91620000f79189919062000791565b60405180910390a184516200011490600a906020880190620003bc565b5083516200012a90600b90602087019062000426565b5060405180606001604052806022815260200162002c426022913980516200015b9160069160209091019062000426565b50600c8390556000805b8351811015620002285762000179620004a9565b8482815181106200018657fe5b602090810291909101810151600d80546001808201835560009290925282517fd7b6990105719101dabeb77144f2a3385c8033acd3af97e9423a695e81ad1eb5600290920291820180546001600160a01b0319166001600160a01b0390921691909117905591909201517fd7b6990105719101dabeb77144f2a3385c8033acd3af97e9423a695e81ad1eb6909101819055939093019291909101905062000165565b50806103e814620002565760405162461bcd60e51b81526004016200024d906200081a565b60405180910390fd5b6000805b83518110156200031d576200026e620004a9565b8482815181106200027b57fe5b602090810291909101810151600e80546001808201835560009290925282517fbb7b4a454dc3493923482f07822329ed19e8244eff582cc204f8554c3620c3fd600290920291820180546001600160a01b0319166001600160a01b0390921691909117905591909201517fbb7b4a454dc3493923482f07822329ed19e8244eff582cc204f8554c3620c3fe90910181905593909301929190910190506200025a565b50806103e814620003425760405162461bcd60e51b81526004016200024d90620007e3565b505050505050505050620008ac565b6001600160e01b031980821614156200037e5760405162461bcd60e51b81526004016200024d90620007ac565b6001600160e01b0319166000908152602081905260409020805460ff19166001179055565b8051620003b890600390602084019062000426565b5050565b82805482825590600052602060002090810192821562000414579160200282015b828111156200041457825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190620003dd565b5062000422929150620004c0565b5090565b828054600181600116156101000203166002900490600052602060002090601f0160209004810192826200045e576000855562000414565b82601f106200047957805160ff191683800117855562000414565b8280016001018555821562000414579182015b82811115620004145782518255916020019190600101906200048c565b604080518082019091526000808252602082015290565b5b80821115620004225760008155600101620004c1565b8051620004e48162000893565b919050565b600082601f830112620004fa578081fd5b8151620005116200050b8262000875565b62000851565b8181529150602080830190848101818402860182018710156200053357600080fd5b60005b848110156200055f5781516200054c8162000893565b8452928201929082019060010162000536565b505050505092915050565b600082601f8301126200057b578081fd5b81516200058c6200050b8262000875565b8181529150602080830190848101604080850287018301881015620005b057600080fd5b6000805b86811015620006135782848b031215620005cc578182fd5b82518084016001600160401b0381118282101715620005e757fe5b84528451620005f68162000893565b8152848601518682015286529484019492820192600101620005b4565b5050505050505092915050565b600082601f83011262000631578081fd5b81516001600160401b038111156200064557fe5b60206200065b601f8301601f1916820162000851565b925081835284818386010111156200067257600080fd5b60005b828110156200069257848101820151848201830152810162000675565b82811115620006a45760008284860101525b50505092915050565b600080600080600080600060e0888a031215620006c8578283fd5b620006d388620004d7565b9650620006e360208901620004d7565b60408901519096506001600160401b038082111562000700578485fd5b6200070e8b838c01620004e9565b965060608a015191508082111562000724578485fd5b620007328b838c0162000620565b955060808a0151945060a08a01519150808211156200074f578384fd5b6200075d8b838c016200056a565b935060c08a015191508082111562000773578283fd5b50620007828a828b016200056a565b91505092959891949750929550565b6001600160a01b039290921682521515602082015260400190565b6020808252601c908201527f4552433136353a20696e76616c696420696e7465726661636520696400000000604082015260600190565b6020808252601b908201527f526573616c652073686172657320646f206e6f74206164642075700000000000604082015260600190565b6020808252601c908201527f496e697469616c2073686172657320646f206e6f742061646420757000000000604082015260600190565b6040518181016001600160401b03811182821017156200086d57fe5b604052919050565b60006001600160401b038211156200088957fe5b5060209081020190565b6001600160a01b0381168114620008a957600080fd5b50565b61238680620008bc6000396000f3fe608060405234801561001057600080fd5b506004361061014c5760003560e01c806388e485c7116100c3578063c6e64e531161007c578063c6e64e53146102e5578063e8a3d485146102f8578063e8f3359f14610300578063e985e9c514610308578063f242432a1461031b578063f637e0191461032e5761014c565b806388e485c71461027c5780638da5cb5b14610284578063938e3d7b1461028c578063a22cb4651461029f578063bb3bafd6146102b2578063c51bb6dc146102d25761014c565b806343bc59021161011557806343bc5902146101e25780634e1273f4146102035780635773bd0714610223578063664aa26b1461023657806379f3ca18146102495780637dc0bf3f146102695761014c565b8062fdd58e1461015157806301ffc9a71461017a5780630b44a2181461019a5780630e89341c146101af5780632eb2c2d6146101cf575b600080fd5b61016461015f36600461188c565b610341565b60405161017191906121c6565b60405180910390f35b61018d610188366004611a65565b61039a565b6040516101719190611da9565b6101ad6101a8366004611856565b6103bd565b005b6101c26101bd366004611ad8565b610463565b6040516101719190611db4565b6101ad6101dd36600461174d565b61052d565b6101f56101f0366004611ad8565b610747565b604051610171929190611cf7565b6102166102113660046118b5565b61077f565b6040516101719190611d68565b6101ad610231366004611701565b61084c565b6101ad610244366004611a9d565b610933565b61025c610257366004611ad8565b61098a565b6040516101719190611c25565b61018d610277366004611ad8565b6109b4565b6101c26109c9565b61025c610a57565b6101ad61029a366004611a9d565b610a66565b6101ad6102ad366004611856565b610abd565b6102c56102c0366004611ad8565b610b8b565b6040516101719190611d10565b6101ad6102e0366004611976565b610c7c565b6101ad6102f3366004611af0565b610d9f565b6101c2610e39565b610164610e94565b61018d61031636600461171b565b610e9a565b6101ad6103293660046117f3565b610ec8565b6101f561033c366004611ad8565b61105c565b60006001600160a01b0383166103725760405162461bcd60e51b815260040161036990611eec565b60405180910390fd5b5060009081526001602090815260408083206001600160a01b03949094168352929052205490565b6001600160e01b0319811660009081526020819052604090205460ff165b919050565b3360009081526008602052604090205460ff16806103e557506007546001600160a01b031633145b6104015760405162461bcd60e51b815260040161036990611f6e565b6001600160a01b03821660009081526008602052604090819020805460ff1916831515179055517f1a594081ae893ab78e67d9b9e843547318164322d32c65369d78a96172d9dc8f906104579084908490611cdc565b60405180910390a15050565b6000818152600560209081526040918290208054835160026001831615610100026000190190921691909104601f810184900484028201840190945283815260609384939192918301828280156104fb5780601f106104d0576101008083540402835291602001916104fb565b820191906000526020600020905b8154815290600101906020018083116104de57829003601f168201915b50505050509050600681604051602001610516929190611ba4565b604051602081830303815290604052915050919050565b815183511461054e5760405162461bcd60e51b81526004016103699061213d565b6001600160a01b0384166105745760405162461bcd60e51b815260040161036990611fdd565b61057c61106c565b6001600160a01b0316856001600160a01b031614806105a257506105a28561031661106c565b6105be5760405162461bcd60e51b815260040161036990612022565b60006105c861106c565b90506105d881878787878761073f565b60005b84518110156106d95760008582815181106105f257fe5b60200260200101519050600085838151811061060a57fe5b60200260200101519050610677816040518060600160405280602a8152602001612327602a91396001600086815260200190815260200160002060008d6001600160a01b03166001600160a01b03168152602001908152602001600020546110719092919063ffffffff16565b60008381526001602090815260408083206001600160a01b038e811685529252808320939093558a16815220546106ae908261109d565b60009283526001602081815260408086206001600160a01b038d1687529091529093205550016105db565b50846001600160a01b0316866001600160a01b0316826001600160a01b03167f4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb8787604051610729929190611d7b565b60405180910390a461073f8187878787876110c9565b505050505050565b600e818154811061075757600080fd5b6000918252602090912060029091020180546001909101546001600160a01b03909116915082565b606081518351146107a25760405162461bcd60e51b8152600401610369906120f4565b6060835167ffffffffffffffff811180156107bc57600080fd5b506040519080825280602002602001820160405280156107e6578160200160208202803683370190505b50905060005b84518110156108445761082585828151811061080457fe5b602002602001015185838151811061081857fe5b6020026020010151610341565b82828151811061083157fe5b60209081029190910101526001016107ec565b509392505050565b60005b600a54811015610909576000600a828154811061086857fe5b600091825260209091200154604051631293efbb60e21b81526001600160a01b0390911690634a4fbeec906108a1908690600401611c25565b60206040518083038186803b1580156108b957600080fd5b505afa1580156108cd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108f19190611a49565b90508015610900575050610930565b5060010161084f565b50600a54600b901561092e5760405162461bcd60e51b81526004016103699190611dc7565b505b50565b3360009081526008602052604090205460ff168061095b57506007546001600160a01b031633145b6109775760405162461bcd60e51b815260040161036990611f6e565b805161092e906006906020840190611574565b600a818154811061099a57600080fd5b6000918252602090912001546001600160a01b0316905081565b60046020526000908152604090205460ff1681565b600b805460408051602060026001851615610100026000190190941693909304601f81018490048402820184019092528181529291830182828015610a4f5780601f10610a2457610100808354040283529160200191610a4f565b820191906000526020600020905b815481529060010190602001808311610a3257829003601f168201915b505050505081565b6007546001600160a01b031681565b3360009081526008602052604090205460ff1680610a8e57506007546001600160a01b031633145b610aaa5760405162461bcd60e51b815260040161036990611f6e565b805161092e906009906020840190611574565b816001600160a01b0316610acf61106c565b6001600160a01b03161415610af65760405162461bcd60e51b8152600401610369906120ab565b8060026000610b0361106c565b6001600160a01b03908116825260208083019390935260409182016000908120918716808252919093529120805460ff191692151592909217909155610b4761106c565b6001600160a01b03167f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c3183604051610b7f9190611da9565b60405180910390a35050565b606081610c0757600d805480602002602001604051908101604052809291908181526020016000905b82821015610bfc576000848152602090819020604080518082019091526002850290910180546001600160a01b03168252600190810154828401529083529092019101610bb4565b5050505090506103b8565b600e805480602002602001604051908101604052809291908181526020016000905b82821015610c71576000848152602090819020604080518082019091526002850290910180546001600160a01b03168252600190810154828401529083529092019101610c29565b505050509050919050565b3360009081526008602052604090205460ff1680610ca457506007546001600160a01b031633145b610cc05760405162461bcd60e51b815260040161036990611f6e565b60608151845114610ce35760405162461bcd60e51b815260040161036990612074565b60005b8451811015610d8057828181518110610cfb57fe5b602002602001015160056000878481518110610d1357fe5b602002602001015181526020019081526020016000209080519060200190610d3c929190611574565b50600160046000878481518110610d4f57fe5b6020908102919091018101518252810191909152604001600020805460ff1916911515919091179055600101610ce6565b50600754610d99906001600160a01b03168585846111e0565b50505050565b3360009081526008602052604090205460ff1680610dc757506007546001600160a01b031633145b610de35760405162461bcd60e51b815260040161036990611f6e565b60008381526005602090815260409091208251606092610e07929190850190611574565b506000848152600460205260409020805460ff19166001179055600754610d99906001600160a01b0316858584611375565b6009805460408051602060026001851615610100026000190190941693909304601f81018490048402820184019092528181529291830182828015610a4f5780601f10610a2457610100808354040283529160200191610a4f565b600c5481565b6001600160a01b03918216600090815260026020908152604080832093909416825291909152205460ff1690565b6001600160a01b038416610eee5760405162461bcd60e51b815260040161036990611fdd565b610ef661106c565b6001600160a01b0316856001600160a01b03161480610f1c5750610f1c8561031661106c565b610f385760405162461bcd60e51b815260040161036990611f94565b6000610f4261106c565b9050610f62818787610f5388611459565b610f5c88611459565b8761073f565b610fa9836040518060600160405280602a8152602001612327602a913960008781526001602090815260408083206001600160a01b038d1684529091529020549190611071565b60008581526001602090815260408083206001600160a01b038b81168552925280832093909355871681522054610fe0908461109d565b60008581526001602090815260408083206001600160a01b03808b168086529190935292819020939093559151909188811691908416907fc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f629061104690899089906121cf565b60405180910390a461073f81878787878761149d565b600d818154811061075757600080fd5b335b90565b600081848411156110955760405162461bcd60e51b81526004016103699190611db4565b505050900390565b6000828201838110156110c25760405162461bcd60e51b815260040161036990611f37565b9392505050565b6110db846001600160a01b031661156e565b1561073f5760405163bc197c8160e01b81526001600160a01b0385169063bc197c81906111149089908990889088908890600401611c39565b602060405180830381600087803b15801561112e57600080fd5b505af192505050801561115e575060408051601f3d908101601f1916820190925261115b91810190611a81565b60015b6111a75761116a61225d565b80611175575061118f565b8060405162461bcd60e51b81526004016103699190611db4565b60405162461bcd60e51b815260040161036990611e50565b6001600160e01b0319811663bc197c8160e01b146111d75760405162461bcd60e51b815260040161036990611ea4565b50505050505050565b6001600160a01b0384166112065760405162461bcd60e51b815260040161036990612185565b81518351146112275760405162461bcd60e51b81526004016103699061213d565b600061123161106c565b90506112428160008787878761073f565b60005b8451811015611306576112bd6001600087848151811061126157fe5b602002602001015181526020019081526020016000206000886001600160a01b03166001600160a01b03168152602001908152602001600020548583815181106112a757fe5b602002602001015161109d90919063ffffffff16565b600160008784815181106112cd57fe5b602090810291909101810151825281810192909252604090810160009081206001600160a01b038b168252909252902055600101611245565b50846001600160a01b031660006001600160a01b0316826001600160a01b03167f4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb8787604051611357929190611d7b565b60405180910390a461136e816000878787876110c9565b5050505050565b6001600160a01b03841661139b5760405162461bcd60e51b815260040161036990612185565b60006113a561106c565b90506113b781600087610f5388611459565b60008481526001602090815260408083206001600160a01b03891684529091529020546113e4908461109d565b60008581526001602090815260408083206001600160a01b03808b16808652919093528184209490945551908416907fc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f629061144290899089906121cf565b60405180910390a461136e8160008787878761149d565b60408051600180825281830190925260609182919060208083019080368337019050509050828160008151811061148c57fe5b602090810291909101015292915050565b6114af846001600160a01b031661156e565b1561073f5760405163f23a6e6160e01b81526001600160a01b0385169063f23a6e61906114e89089908990889088908890600401611c97565b602060405180830381600087803b15801561150257600080fd5b505af1925050508015611532575060408051601f3d908101601f1916820190925261152f91810190611a81565b60015b61153e5761116a61225d565b6001600160e01b0319811663f23a6e6160e01b146111d75760405162461bcd60e51b815260040161036990611ea4565b3b151590565b828054600181600116156101000203166002900490600052602060002090601f0160209004810192826115aa57600085556115f0565b82601f106115c357805160ff19168380011785556115f0565b828001600101855582156115f0579182015b828111156115f05782518255916020019190600101906115d5565b506115fc929150611600565b5090565b5b808211156115fc5760008155600101611601565b80356001600160a01b03811681146103b857600080fd5b600082601f83011261163c578081fd5b813561164f61164a82612201565b6121dd565b81815291506020808301908481018184028601820187101561167057600080fd5b60005b8481101561168f57813584529282019290820190600101611673565b505050505092915050565b600082601f8301126116aa578081fd5b813567ffffffffffffffff8111156116be57fe5b6116d1601f8201601f19166020016121dd565b91508082528360208285010111156116e857600080fd5b8060208401602084013760009082016020015292915050565b600060208284031215611712578081fd5b6110c282611615565b6000806040838503121561172d578081fd5b61173683611615565b915061174460208401611615565b90509250929050565b600080600080600060a08688031215611764578081fd5b61176d86611615565b945061177b60208701611615565b9350604086013567ffffffffffffffff80821115611797578283fd5b6117a389838a0161162c565b945060608801359150808211156117b8578283fd5b6117c489838a0161162c565b935060808801359150808211156117d9578283fd5b506117e68882890161169a565b9150509295509295909350565b600080600080600060a0868803121561180a578081fd5b61181386611615565b945061182160208701611615565b93506040860135925060608601359150608086013567ffffffffffffffff81111561184a578182fd5b6117e68882890161169a565b60008060408385031215611868578182fd5b61187183611615565b9150602083013561188181612302565b809150509250929050565b6000806040838503121561189e578182fd5b6118a783611615565b946020939093013593505050565b600080604083850312156118c7578182fd5b823567ffffffffffffffff808211156118de578384fd5b818501915085601f8301126118f1578384fd5b81356118ff61164a82612201565b80828252602080830192508086018a82838702890101111561191f578889fd5b8896505b848710156119485761193481611615565b845260019690960195928101928101611923565b50909650870135935050508082111561195f578283fd5b5061196c8582860161162c565b9150509250929050565b60008060006060848603121561198a578283fd5b833567ffffffffffffffff808211156119a1578485fd5b6119ad8783880161162c565b94506020915081860135818111156119c3578485fd5b6119cf8882890161162c565b9450506040860135818111156119e3578384fd5b86019050601f810187136119f5578283fd5b8035611a0361164a82612201565b81815283810190838501865b84811015611a3857611a268c88843589010161169a565b84529286019290860190600101611a0f565b505080955050505050509250925092565b600060208284031215611a5a578081fd5b81516110c281612302565b600060208284031215611a76578081fd5b81356110c281612310565b600060208284031215611a92578081fd5b81516110c281612310565b600060208284031215611aae578081fd5b813567ffffffffffffffff811115611ac4578182fd5b611ad08482850161169a565b949350505050565b600060208284031215611ae9578081fd5b5035919050565b600080600060608486031215611b04578081fd5b8335925060208401359150604084013567ffffffffffffffff811115611b28578182fd5b611b348682870161169a565b9150509250925092565b6000815180845260208085019450808401835b83811015611b6d57815187529582019590820190600101611b51565b509495945050505050565b60008151808452611b9081602086016020860161222b565b601f01601f19169290920160200192915050565b6000808454600180821660008114611bc35760018114611bda57611c09565b60ff198316865260028304607f1686019350611c09565b600283048886526020808720875b83811015611c015781548a820152908501908201611be8565b505050860193505b5050508351611c1c81836020880161222b565b01949350505050565b6001600160a01b0391909116815260200190565b6001600160a01b0386811682528516602082015260a060408201819052600090611c6590830186611b3e565b8281036060840152611c778186611b3e565b90508281036080840152611c8b8185611b78565b98975050505050505050565b6001600160a01b03868116825285166020820152604081018490526060810183905260a060808201819052600090611cd190830184611b78565b979650505050505050565b6001600160a01b039290921682521515602082015260400190565b6001600160a01b03929092168252602082015260400190565b602080825282518282018190526000919060409081850190868401855b82811015611d5b57815180516001600160a01b03168552860151868501529284019290850190600101611d2d565b5091979650505050505050565b6000602082526110c26020830184611b3e565b600060408252611d8e6040830185611b3e565b8281036020840152611da08185611b3e565b95945050505050565b901515815260200190565b6000602082526110c26020830184611b78565b60006020808301818452828554600180821660008114611dee5760018114611e0c57611d5b565b60028304607f16855260ff1983166040890152606088019350611d5b565b60028304808652611e1c8a61221f565b885b82811015611e3a5781548b820160400152908401908801611e1e565b8a01604001955050505091979650505050505050565b60208082526034908201527f455243313135353a207472616e7366657220746f206e6f6e20455243313135356040820152732932b1b2b4bb32b91034b6b83632b6b2b73a32b960611b606082015260800190565b60208082526028908201527f455243313135353a204552433131353552656365697665722072656a656374656040820152676420746f6b656e7360c01b606082015260800190565b6020808252602b908201527f455243313135353a2062616c616e636520717565727920666f7220746865207a60408201526a65726f206164647265737360a81b606082015260800190565b6020808252601b908201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604082015260600190565b6020808252600c908201526b1d5b985d5d1a1bdc9a5cd95960a21b604082015260600190565b60208082526029908201527f455243313135353a2063616c6c6572206973206e6f74206f776e6572206e6f7260408201526808185c1c1c9bdd995960ba1b606082015260800190565b60208082526025908201527f455243313135353a207472616e7366657220746f20746865207a65726f206164604082015264647265737360d81b606082015260800190565b60208082526032908201527f455243313135353a207472616e736665722063616c6c6572206973206e6f74206040820152711bdddb995c881b9bdc88185c1c1c9bdd995960721b606082015260800190565b6020808252601a908201527f6172726179206c656e6774687320646f206e6f74206d61746368000000000000604082015260600190565b60208082526029908201527f455243313135353a2073657474696e6720617070726f76616c20737461747573604082015268103337b91039b2b63360b91b606082015260800190565b60208082526029908201527f455243313135353a206163636f756e747320616e6420696473206c656e677468604082015268040dad2e6dac2e8c6d60bb1b606082015260800190565b60208082526028908201527f455243313135353a2069647320616e6420616d6f756e7473206c656e677468206040820152670dad2e6dac2e8c6d60c31b606082015260800190565b60208082526021908201527f455243313135353a206d696e7420746f20746865207a65726f206164647265736040820152607360f81b606082015260800190565b90815260200190565b918252602082015260400190565b60405181810167ffffffffffffffff811182821017156121f957fe5b604052919050565b600067ffffffffffffffff82111561221557fe5b5060209081020190565b60009081526020902090565b60005b8381101561224657818101518382015260200161222e565b83811115610d995750506000910152565b60e01c90565b600060443d101561226d5761106e565b600481823e6308c379a06122818251612257565b1461228b5761106e565b6040513d600319016004823e80513d67ffffffffffffffff81602484011181841117156122bb575050505061106e565b828401925082519150808211156122d5575050505061106e565b503d830160208284010111156122ed5750505061106e565b601f01601f1916810160200160405291505090565b801515811461093057600080fd5b6001600160e01b03198116811461093057600080fdfe455243313135353a20696e73756666696369656e742062616c616e636520666f72207472616e73666572a264697066735822122080dd1428b6aff1d542fb705990002b36ab431e9d62edd84ad18e5a6695e63ead64736f6c6343000705003368747470733a2f2f6d696e74792e6d7970696e6174612e636c6f75642f697066732f"

// DeployPMintyMultiToken deploys a new Ethereum contract, binding an instance of PMintyMultiToken to it.
func DeployPMintyMultiToken(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address, saleContract common.Address, _locs []common.Address, _lockError string, _royaltyPerMille *big.Int, _initialEntries []PoolEntry, _resaleEntries []PoolEntry) (common.Address, *types.Transaction, *PMintyMultiToken, error) {
	parsed, err := abi.JSON(strings.NewReader(PMintyMultiTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PMintyMultiTokenBin), backend, _owner, saleContract, _locs, _lockError, _royaltyPerMille, _initialEntries, _resaleEntries)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PMintyMultiToken{PMintyMultiTokenCaller: PMintyMultiTokenCaller{contract: contract}, PMintyMultiTokenTransactor: PMintyMultiTokenTransactor{contract: contract}, PMintyMultiTokenFilterer: PMintyMultiTokenFilterer{contract: contract}}, nil
}

// PMintyMultiToken is an auto generated Go binding around an Ethereum contract.
type PMintyMultiToken struct {
	PMintyMultiTokenCaller     // Read-only binding to the contract
	PMintyMultiTokenTransactor // Write-only binding to the contract
	PMintyMultiTokenFilterer   // Log filterer for contract events
}

// PMintyMultiTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type PMintyMultiTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PMintyMultiTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PMintyMultiTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PMintyMultiTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PMintyMultiTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PMintyMultiTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PMintyMultiTokenSession struct {
	Contract     *PMintyMultiToken // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PMintyMultiTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PMintyMultiTokenCallerSession struct {
	Contract *PMintyMultiTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// PMintyMultiTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PMintyMultiTokenTransactorSession struct {
	Contract     *PMintyMultiTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// PMintyMultiTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type PMintyMultiTokenRaw struct {
	Contract *PMintyMultiToken // Generic contract binding to access the raw methods on
}

// PMintyMultiTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PMintyMultiTokenCallerRaw struct {
	Contract *PMintyMultiTokenCaller // Generic read-only contract binding to access the raw methods on
}

// PMintyMultiTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PMintyMultiTokenTransactorRaw struct {
	Contract *PMintyMultiTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPMintyMultiToken creates a new instance of PMintyMultiToken, bound to a specific deployed contract.
func NewPMintyMultiToken(address common.Address, backend bind.ContractBackend) (*PMintyMultiToken, error) {
	contract, err := bindPMintyMultiToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PMintyMultiToken{PMintyMultiTokenCaller: PMintyMultiTokenCaller{contract: contract}, PMintyMultiTokenTransactor: PMintyMultiTokenTransactor{contract: contract}, PMintyMultiTokenFilterer: PMintyMultiTokenFilterer{contract: contract}}, nil
}

// NewPMintyMultiTokenCaller creates a new read-only instance of PMintyMultiToken, bound to a specific deployed contract.
func NewPMintyMultiTokenCaller(address common.Address, caller bind.ContractCaller) (*PMintyMultiTokenCaller, error) {
	contract, err := bindPMintyMultiToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PMintyMultiTokenCaller{contract: contract}, nil
}

// NewPMintyMultiTokenTransactor creates a new write-only instance of PMintyMultiToken, bound to a specific deployed contract.
func NewPMintyMultiTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*PMintyMultiTokenTransactor, error) {
	contract, err := bindPMintyMultiToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PMintyMultiTokenTransactor{contract: contract}, nil
}

// NewPMintyMultiTokenFilterer creates a new log filterer instance of PMintyMultiToken, bound to a specific deployed contract.
func NewPMintyMultiTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*PMintyMultiTokenFilterer, error) {
	contract, err := bindPMintyMultiToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PMintyMultiTokenFilterer{contract: contract}, nil
}

// bindPMintyMultiToken binds a generic wrapper to an already deployed contract.
func bindPMintyMultiToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PMintyMultiTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PMintyMultiToken *PMintyMultiTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PMintyMultiToken.Contract.PMintyMultiTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PMintyMultiToken *PMintyMultiTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PMintyMultiToken.Contract.PMintyMultiTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PMintyMultiToken *PMintyMultiTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PMintyMultiToken.Contract.PMintyMultiTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PMintyMultiToken *PMintyMultiTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PMintyMultiToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PMintyMultiToken *PMintyMultiTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PMintyMultiToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PMintyMultiToken *PMintyMultiTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PMintyMultiToken.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_PMintyMultiToken *PMintyMultiTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PMintyMultiToken.contract.Call(opts, &out, "balanceOf", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_PMintyMultiToken *PMintyMultiTokenSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _PMintyMultiToken.Contract.BalanceOf(&_PMintyMultiToken.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_PMintyMultiToken *PMintyMultiTokenCallerSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _PMintyMultiToken.Contract.BalanceOf(&_PMintyMultiToken.CallOpts, account, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_PMintyMultiToken *PMintyMultiTokenCaller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _PMintyMultiToken.contract.Call(opts, &out, "balanceOfBatch", accounts, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_PMintyMultiToken *PMintyMultiTokenSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _PMintyMultiToken.Contract.BalanceOfBatch(&_PMintyMultiToken.CallOpts, accounts, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_PMintyMultiToken *PMintyMultiTokenCallerSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _PMintyMultiToken.Contract.BalanceOfBatch(&_PMintyMultiToken.CallOpts, accounts, ids)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_PMintyMultiToken *PMintyMultiTokenCaller) ContractURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PMintyMultiToken.contract.Call(opts, &out, "contractURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_PMintyMultiToken *PMintyMultiTokenSession) ContractURI() (string, error) {
	return _PMintyMultiToken.Contract.ContractURI(&_PMintyMultiToken.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_PMintyMultiToken *PMintyMultiTokenCallerSession) ContractURI() (string, error) {
	return _PMintyMultiToken.Contract.ContractURI(&_PMintyMultiToken.CallOpts)
}

// GetRoyalties is a free data retrieval call binding the contract method 0xbb3bafd6.
//
// Solidity: function getRoyalties(uint256 saleNumber) view returns((address,uint256)[])
func (_PMintyMultiToken *PMintyMultiTokenCaller) GetRoyalties(opts *bind.CallOpts, saleNumber *big.Int) ([]PoolEntry, error) {
	var out []interface{}
	err := _PMintyMultiToken.contract.Call(opts, &out, "getRoyalties", saleNumber)

	if err != nil {
		return *new([]PoolEntry), err
	}

	out0 := *abi.ConvertType(out[0], new([]PoolEntry)).(*[]PoolEntry)

	return out0, err

}

// GetRoyalties is a free data retrieval call binding the contract method 0xbb3bafd6.
//
// Solidity: function getRoyalties(uint256 saleNumber) view returns((address,uint256)[])
func (_PMintyMultiToken *PMintyMultiTokenSession) GetRoyalties(saleNumber *big.Int) ([]PoolEntry, error) {
	return _PMintyMultiToken.Contract.GetRoyalties(&_PMintyMultiToken.CallOpts, saleNumber)
}

// GetRoyalties is a free data retrieval call binding the contract method 0xbb3bafd6.
//
// Solidity: function getRoyalties(uint256 saleNumber) view returns((address,uint256)[])
func (_PMintyMultiToken *PMintyMultiTokenCallerSession) GetRoyalties(saleNumber *big.Int) ([]PoolEntry, error) {
	return _PMintyMultiToken.Contract.GetRoyalties(&_PMintyMultiToken.CallOpts, saleNumber)
}

// InitialPool is a free data retrieval call binding the contract method 0xf637e019.
//
// Solidity: function initialPool(uint256 ) view returns(address beneficiary, uint256 share)
func (_PMintyMultiToken *PMintyMultiTokenCaller) InitialPool(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Beneficiary common.Address
	Share       *big.Int
}, error) {
	var out []interface{}
	err := _PMintyMultiToken.contract.Call(opts, &out, "initialPool", arg0)

	outstruct := new(struct {
		Beneficiary common.Address
		Share       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Beneficiary = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Share = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// InitialPool is a free data retrieval call binding the contract method 0xf637e019.
//
// Solidity: function initialPool(uint256 ) view returns(address beneficiary, uint256 share)
func (_PMintyMultiToken *PMintyMultiTokenSession) InitialPool(arg0 *big.Int) (struct {
	Beneficiary common.Address
	Share       *big.Int
}, error) {
	return _PMintyMultiToken.Contract.InitialPool(&_PMintyMultiToken.CallOpts, arg0)
}

// InitialPool is a free data retrieval call binding the contract method 0xf637e019.
//
// Solidity: function initialPool(uint256 ) view returns(address beneficiary, uint256 share)
func (_PMintyMultiToken *PMintyMultiTokenCallerSession) InitialPool(arg0 *big.Int) (struct {
	Beneficiary common.Address
	Share       *big.Int
}, error) {
	return _PMintyMultiToken.Contract.InitialPool(&_PMintyMultiToken.CallOpts, arg0)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_PMintyMultiToken *PMintyMultiTokenCaller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _PMintyMultiToken.contract.Call(opts, &out, "isApprovedForAll", account, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_PMintyMultiToken *PMintyMultiTokenSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _PMintyMultiToken.Contract.IsApprovedForAll(&_PMintyMultiToken.CallOpts, account, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_PMintyMultiToken *PMintyMultiTokenCallerSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _PMintyMultiToken.Contract.IsApprovedForAll(&_PMintyMultiToken.CallOpts, account, operator)
}

// LockError is a free data retrieval call binding the contract method 0x88e485c7.
//
// Solidity: function lockError() view returns(string)
func (_PMintyMultiToken *PMintyMultiTokenCaller) LockError(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PMintyMultiToken.contract.Call(opts, &out, "lockError")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// LockError is a free data retrieval call binding the contract method 0x88e485c7.
//
// Solidity: function lockError() view returns(string)
func (_PMintyMultiToken *PMintyMultiTokenSession) LockError() (string, error) {
	return _PMintyMultiToken.Contract.LockError(&_PMintyMultiToken.CallOpts)
}

// LockError is a free data retrieval call binding the contract method 0x88e485c7.
//
// Solidity: function lockError() view returns(string)
func (_PMintyMultiToken *PMintyMultiTokenCallerSession) LockError() (string, error) {
	return _PMintyMultiToken.Contract.LockError(&_PMintyMultiToken.CallOpts)
}

// Locs is a free data retrieval call binding the contract method 0x79f3ca18.
//
// Solidity: function locs(uint256 ) view returns(address)
func (_PMintyMultiToken *PMintyMultiTokenCaller) Locs(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _PMintyMultiToken.contract.Call(opts, &out, "locs", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Locs is a free data retrieval call binding the contract method 0x79f3ca18.
//
// Solidity: function locs(uint256 ) view returns(address)
func (_PMintyMultiToken *PMintyMultiTokenSession) Locs(arg0 *big.Int) (common.Address, error) {
	return _PMintyMultiToken.Contract.Locs(&_PMintyMultiToken.CallOpts, arg0)
}

// Locs is a free data retrieval call binding the contract method 0x79f3ca18.
//
// Solidity: function locs(uint256 ) view returns(address)
func (_PMintyMultiToken *PMintyMultiTokenCallerSession) Locs(arg0 *big.Int) (common.Address, error) {
	return _PMintyMultiToken.Contract.Locs(&_PMintyMultiToken.CallOpts, arg0)
}

// Minted is a free data retrieval call binding the contract method 0x7dc0bf3f.
//
// Solidity: function minted(uint256 ) view returns(bool)
func (_PMintyMultiToken *PMintyMultiTokenCaller) Minted(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _PMintyMultiToken.contract.Call(opts, &out, "minted", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Minted is a free data retrieval call binding the contract method 0x7dc0bf3f.
//
// Solidity: function minted(uint256 ) view returns(bool)
func (_PMintyMultiToken *PMintyMultiTokenSession) Minted(arg0 *big.Int) (bool, error) {
	return _PMintyMultiToken.Contract.Minted(&_PMintyMultiToken.CallOpts, arg0)
}

// Minted is a free data retrieval call binding the contract method 0x7dc0bf3f.
//
// Solidity: function minted(uint256 ) view returns(bool)
func (_PMintyMultiToken *PMintyMultiTokenCallerSession) Minted(arg0 *big.Int) (bool, error) {
	return _PMintyMultiToken.Contract.Minted(&_PMintyMultiToken.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PMintyMultiToken *PMintyMultiTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PMintyMultiToken.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PMintyMultiToken *PMintyMultiTokenSession) Owner() (common.Address, error) {
	return _PMintyMultiToken.Contract.Owner(&_PMintyMultiToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PMintyMultiToken *PMintyMultiTokenCallerSession) Owner() (common.Address, error) {
	return _PMintyMultiToken.Contract.Owner(&_PMintyMultiToken.CallOpts)
}

// ResalePool is a free data retrieval call binding the contract method 0x43bc5902.
//
// Solidity: function resalePool(uint256 ) view returns(address beneficiary, uint256 share)
func (_PMintyMultiToken *PMintyMultiTokenCaller) ResalePool(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Beneficiary common.Address
	Share       *big.Int
}, error) {
	var out []interface{}
	err := _PMintyMultiToken.contract.Call(opts, &out, "resalePool", arg0)

	outstruct := new(struct {
		Beneficiary common.Address
		Share       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Beneficiary = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Share = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ResalePool is a free data retrieval call binding the contract method 0x43bc5902.
//
// Solidity: function resalePool(uint256 ) view returns(address beneficiary, uint256 share)
func (_PMintyMultiToken *PMintyMultiTokenSession) ResalePool(arg0 *big.Int) (struct {
	Beneficiary common.Address
	Share       *big.Int
}, error) {
	return _PMintyMultiToken.Contract.ResalePool(&_PMintyMultiToken.CallOpts, arg0)
}

// ResalePool is a free data retrieval call binding the contract method 0x43bc5902.
//
// Solidity: function resalePool(uint256 ) view returns(address beneficiary, uint256 share)
func (_PMintyMultiToken *PMintyMultiTokenCallerSession) ResalePool(arg0 *big.Int) (struct {
	Beneficiary common.Address
	Share       *big.Int
}, error) {
	return _PMintyMultiToken.Contract.ResalePool(&_PMintyMultiToken.CallOpts, arg0)
}

// RoyaltyPerMille is a free data retrieval call binding the contract method 0xe8f3359f.
//
// Solidity: function royaltyPerMille() view returns(uint256)
func (_PMintyMultiToken *PMintyMultiTokenCaller) RoyaltyPerMille(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PMintyMultiToken.contract.Call(opts, &out, "royaltyPerMille")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RoyaltyPerMille is a free data retrieval call binding the contract method 0xe8f3359f.
//
// Solidity: function royaltyPerMille() view returns(uint256)
func (_PMintyMultiToken *PMintyMultiTokenSession) RoyaltyPerMille() (*big.Int, error) {
	return _PMintyMultiToken.Contract.RoyaltyPerMille(&_PMintyMultiToken.CallOpts)
}

// RoyaltyPerMille is a free data retrieval call binding the contract method 0xe8f3359f.
//
// Solidity: function royaltyPerMille() view returns(uint256)
func (_PMintyMultiToken *PMintyMultiTokenCallerSession) RoyaltyPerMille() (*big.Int, error) {
	return _PMintyMultiToken.Contract.RoyaltyPerMille(&_PMintyMultiToken.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PMintyMultiToken *PMintyMultiTokenCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _PMintyMultiToken.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PMintyMultiToken *PMintyMultiTokenSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _PMintyMultiToken.Contract.SupportsInterface(&_PMintyMultiToken.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PMintyMultiToken *PMintyMultiTokenCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _PMintyMultiToken.Contract.SupportsInterface(&_PMintyMultiToken.CallOpts, interfaceId)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 tokenId) view returns(string)
func (_PMintyMultiToken *PMintyMultiTokenCaller) Uri(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _PMintyMultiToken.contract.Call(opts, &out, "uri", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 tokenId) view returns(string)
func (_PMintyMultiToken *PMintyMultiTokenSession) Uri(tokenId *big.Int) (string, error) {
	return _PMintyMultiToken.Contract.Uri(&_PMintyMultiToken.CallOpts, tokenId)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 tokenId) view returns(string)
func (_PMintyMultiToken *PMintyMultiTokenCallerSession) Uri(tokenId *big.Int) (string, error) {
	return _PMintyMultiToken.Contract.Uri(&_PMintyMultiToken.CallOpts, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0xc6e64e53.
//
// Solidity: function mint(uint256 tokenId, uint256 quantity, string hash) returns()
func (_PMintyMultiToken *PMintyMultiTokenTransactor) Mint(opts *bind.TransactOpts, tokenId *big.Int, quantity *big.Int, hash string) (*types.Transaction, error) {
	return _PMintyMultiToken.contract.Transact(opts, "mint", tokenId, quantity, hash)
}

// Mint is a paid mutator transaction binding the contract method 0xc6e64e53.
//
// Solidity: function mint(uint256 tokenId, uint256 quantity, string hash) returns()
func (_PMintyMultiToken *PMintyMultiTokenSession) Mint(tokenId *big.Int, quantity *big.Int, hash string) (*types.Transaction, error) {
	return _PMintyMultiToken.Contract.Mint(&_PMintyMultiToken.TransactOpts, tokenId, quantity, hash)
}

// Mint is a paid mutator transaction binding the contract method 0xc6e64e53.
//
// Solidity: function mint(uint256 tokenId, uint256 quantity, string hash) returns()
func (_PMintyMultiToken *PMintyMultiTokenTransactorSession) Mint(tokenId *big.Int, quantity *big.Int, hash string) (*types.Transaction, error) {
	return _PMintyMultiToken.Contract.Mint(&_PMintyMultiToken.TransactOpts, tokenId, quantity, hash)
}

// MintBatch is a paid mutator transaction binding the contract method 0xc51bb6dc.
//
// Solidity: function mintBatch(uint256[] tokenIds, uint256[] quantities, string[] hashes) returns()
func (_PMintyMultiToken *PMintyMultiTokenTransactor) MintBatch(opts *bind.TransactOpts, tokenIds []*big.Int, quantities []*big.Int, hashes []string) (*types.Transaction, error) {
	return _PMintyMultiToken.contract.Transact(opts, "mintBatch", tokenIds, quantities, hashes)
}

// MintBatch is a paid mutator transaction binding the contract method 0xc51bb6dc.
//
// Solidity: function mintBatch(uint256[] tokenIds, uint256[] quantities, string[] hashes) returns()
func (_PMintyMultiToken *PMintyMultiTokenSession) MintBatch(tokenIds []*big.Int, quantities []*big.Int, hashes []string) (*types.Transaction, error) {
	return _PMintyMultiToken.Contract.MintBatch(&_PMintyMultiToken.TransactOpts, tokenIds, quantities, hashes)
}

// MintBatch is a paid mutator transaction binding the contract method 0xc51bb6dc.
//
// Solidity: function mintBatch(uint256[] tokenIds, uint256[] quantities, string[] hashes) returns()
func (_PMintyMultiToken *PMintyMultiTokenTransactorSession) MintBatch(tokenIds []*big.Int, quantities []*big.Int, hashes []string) (*types.Transaction, error) {
	return _PMintyMultiToken.Contract.MintBatch(&_PMintyMultiToken.TransactOpts, tokenIds, quantities, hashes)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_PMintyMultiToken *PMintyMultiTokenTransactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _PMintyMultiToken.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_PMintyMultiToken *PMintyMultiTokenSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _PMintyMultiToken.Contract.SafeBatchTransferFrom(&_PMintyMultiToken.TransactOpts, from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_PMintyMultiToken *PMintyMultiTokenTransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _PMintyMultiToken.Contract.SafeBatchTransferFrom(&_PMintyMultiToken.TransactOpts, from, to, ids, amounts, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_PMintyMultiToken *PMintyMultiTokenTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _PMintyMultiToken.contract.Transact(opts, "safeTransferFrom", from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_PMintyMultiToken *PMintyMultiTokenSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _PMintyMultiToken.Contract.SafeTransferFrom(&_PMintyMultiToken.TransactOpts, from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_PMintyMultiToken *PMintyMultiTokenTransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _PMintyMultiToken.Contract.SafeTransferFrom(&_PMintyMultiToken.TransactOpts, from, to, id, amount, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_PMintyMultiToken *PMintyMultiTokenTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _PMintyMultiToken.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_PMintyMultiToken *PMintyMultiTokenSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _PMintyMultiToken.Contract.SetApprovalForAll(&_PMintyMultiToken.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_PMintyMultiToken *PMintyMultiTokenTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _PMintyMultiToken.Contract.SetApprovalForAll(&_PMintyMultiToken.TransactOpts, operator, approved)
}

// SetAuth is a paid mutator transaction binding the contract method 0x0b44a218.
//
// Solidity: function setAuth(address operator, bool enabled) returns()
func (_PMintyMultiToken *PMintyMultiTokenTransactor) SetAuth(opts *bind.TransactOpts, operator common.Address, enabled bool) (*types.Transaction, error) {
	return _PMintyMultiToken.contract.Transact(opts, "setAuth", operator, enabled)
}

// SetAuth is a paid mutator transaction binding the contract method 0x0b44a218.
//
// Solidity: function setAuth(address operator, bool enabled) returns()
func (_PMintyMultiToken *PMintyMultiTokenSession) SetAuth(operator common.Address, enabled bool) (*types.Transaction, error) {
	return _PMintyMultiToken.Contract.SetAuth(&_PMintyMultiToken.TransactOpts, operator, enabled)
}

// SetAuth is a paid mutator transaction binding the contract method 0x0b44a218.
//
// Solidity: function setAuth(address operator, bool enabled) returns()
func (_PMintyMultiToken *PMintyMultiTokenTransactorSession) SetAuth(operator common.Address, enabled bool) (*types.Transaction, error) {
	return _PMintyMultiToken.Contract.SetAuth(&_PMintyMultiToken.TransactOpts, operator, enabled)
}

// SetBase is a paid mutator transaction binding the contract method 0x664aa26b.
//
// Solidity: function setBase(string _base) returns()
func (_PMintyMultiToken *PMintyMultiTokenTransactor) SetBase(opts *bind.TransactOpts, _base string) (*types.Transaction, error) {
	return _PMintyMultiToken.contract.Transact(opts, "setBase", _base)
}

// SetBase is a paid mutator transaction binding the contract method 0x664aa26b.
//
// Solidity: function setBase(string _base) returns()
func (_PMintyMultiToken *PMintyMultiTokenSession) SetBase(_base string) (*types.Transaction, error) {
	return _PMintyMultiToken.Contract.SetBase(&_PMintyMultiToken.TransactOpts, _base)
}

// SetBase is a paid mutator transaction binding the contract method 0x664aa26b.
//
// Solidity: function setBase(string _base) returns()
func (_PMintyMultiToken *PMintyMultiTokenTransactorSession) SetBase(_base string) (*types.Transaction, error) {
	return _PMintyMultiToken.Contract.SetBase(&_PMintyMultiToken.TransactOpts, _base)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string _uri) returns()
func (_PMintyMultiToken *PMintyMultiTokenTransactor) SetContractURI(opts *bind.TransactOpts, _uri string) (*types.Transaction, error) {
	return _PMintyMultiToken.contract.Transact(opts, "setContractURI", _uri)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string _uri) returns()
func (_PMintyMultiToken *PMintyMultiTokenSession) SetContractURI(_uri string) (*types.Transaction, error) {
	return _PMintyMultiToken.Contract.SetContractURI(&_PMintyMultiToken.TransactOpts, _uri)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string _uri) returns()
func (_PMintyMultiToken *PMintyMultiTokenTransactorSession) SetContractURI(_uri string) (*types.Transaction, error) {
	return _PMintyMultiToken.Contract.SetContractURI(&_PMintyMultiToken.TransactOpts, _uri)
}

// ValidateBuyer is a paid mutator transaction binding the contract method 0x5773bd07.
//
// Solidity: function validateBuyer(address buyer) returns()
func (_PMintyMultiToken *PMintyMultiTokenTransactor) ValidateBuyer(opts *bind.TransactOpts, buyer common.Address) (*types.Transaction, error) {
	return _PMintyMultiToken.contract.Transact(opts, "validateBuyer", buyer)
}

// ValidateBuyer is a paid mutator transaction binding the contract method 0x5773bd07.
//
// Solidity: function validateBuyer(address buyer) returns()
func (_PMintyMultiToken *PMintyMultiTokenSession) ValidateBuyer(buyer common.Address) (*types.Transaction, error) {
	return _PMintyMultiToken.Contract.ValidateBuyer(&_PMintyMultiToken.TransactOpts, buyer)
}

// ValidateBuyer is a paid mutator transaction binding the contract method 0x5773bd07.
//
// Solidity: function validateBuyer(address buyer) returns()
func (_PMintyMultiToken *PMintyMultiTokenTransactorSession) ValidateBuyer(buyer common.Address) (*types.Transaction, error) {
	return _PMintyMultiToken.Contract.ValidateBuyer(&_PMintyMultiToken.TransactOpts, buyer)
}

// PMintyMultiTokenApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the PMintyMultiToken contract.
type PMintyMultiTokenApprovalForAllIterator struct {
	Event *PMintyMultiTokenApprovalForAll // Event containing the contract specifics and raw log

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
func (it *PMintyMultiTokenApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PMintyMultiTokenApprovalForAll)
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
		it.Event = new(PMintyMultiTokenApprovalForAll)
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
func (it *PMintyMultiTokenApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PMintyMultiTokenApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PMintyMultiTokenApprovalForAll represents a ApprovalForAll event raised by the PMintyMultiToken contract.
type PMintyMultiTokenApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_PMintyMultiToken *PMintyMultiTokenFilterer) FilterApprovalForAll(opts *bind.FilterOpts, account []common.Address, operator []common.Address) (*PMintyMultiTokenApprovalForAllIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _PMintyMultiToken.contract.FilterLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &PMintyMultiTokenApprovalForAllIterator{contract: _PMintyMultiToken.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_PMintyMultiToken *PMintyMultiTokenFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *PMintyMultiTokenApprovalForAll, account []common.Address, operator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _PMintyMultiToken.contract.WatchLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PMintyMultiTokenApprovalForAll)
				if err := _PMintyMultiToken.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_PMintyMultiToken *PMintyMultiTokenFilterer) ParseApprovalForAll(log types.Log) (*PMintyMultiTokenApprovalForAll, error) {
	event := new(PMintyMultiTokenApprovalForAll)
	if err := _PMintyMultiToken.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PMintyMultiTokenOperatorSetIterator is returned from FilterOperatorSet and is used to iterate over the raw logs and unpacked data for OperatorSet events raised by the PMintyMultiToken contract.
type PMintyMultiTokenOperatorSetIterator struct {
	Event *PMintyMultiTokenOperatorSet // Event containing the contract specifics and raw log

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
func (it *PMintyMultiTokenOperatorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PMintyMultiTokenOperatorSet)
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
		it.Event = new(PMintyMultiTokenOperatorSet)
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
func (it *PMintyMultiTokenOperatorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PMintyMultiTokenOperatorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PMintyMultiTokenOperatorSet represents a OperatorSet event raised by the PMintyMultiToken contract.
type PMintyMultiTokenOperatorSet struct {
	Operator common.Address
	Enabled  bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorSet is a free log retrieval operation binding the contract event 0x1a594081ae893ab78e67d9b9e843547318164322d32c65369d78a96172d9dc8f.
//
// Solidity: event OperatorSet(address operator, bool enabled)
func (_PMintyMultiToken *PMintyMultiTokenFilterer) FilterOperatorSet(opts *bind.FilterOpts) (*PMintyMultiTokenOperatorSetIterator, error) {

	logs, sub, err := _PMintyMultiToken.contract.FilterLogs(opts, "OperatorSet")
	if err != nil {
		return nil, err
	}
	return &PMintyMultiTokenOperatorSetIterator{contract: _PMintyMultiToken.contract, event: "OperatorSet", logs: logs, sub: sub}, nil
}

// WatchOperatorSet is a free log subscription operation binding the contract event 0x1a594081ae893ab78e67d9b9e843547318164322d32c65369d78a96172d9dc8f.
//
// Solidity: event OperatorSet(address operator, bool enabled)
func (_PMintyMultiToken *PMintyMultiTokenFilterer) WatchOperatorSet(opts *bind.WatchOpts, sink chan<- *PMintyMultiTokenOperatorSet) (event.Subscription, error) {

	logs, sub, err := _PMintyMultiToken.contract.WatchLogs(opts, "OperatorSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PMintyMultiTokenOperatorSet)
				if err := _PMintyMultiToken.contract.UnpackLog(event, "OperatorSet", log); err != nil {
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

// ParseOperatorSet is a log parse operation binding the contract event 0x1a594081ae893ab78e67d9b9e843547318164322d32c65369d78a96172d9dc8f.
//
// Solidity: event OperatorSet(address operator, bool enabled)
func (_PMintyMultiToken *PMintyMultiTokenFilterer) ParseOperatorSet(log types.Log) (*PMintyMultiTokenOperatorSet, error) {
	event := new(PMintyMultiTokenOperatorSet)
	if err := _PMintyMultiToken.contract.UnpackLog(event, "OperatorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PMintyMultiTokenTransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the PMintyMultiToken contract.
type PMintyMultiTokenTransferBatchIterator struct {
	Event *PMintyMultiTokenTransferBatch // Event containing the contract specifics and raw log

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
func (it *PMintyMultiTokenTransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PMintyMultiTokenTransferBatch)
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
		it.Event = new(PMintyMultiTokenTransferBatch)
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
func (it *PMintyMultiTokenTransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PMintyMultiTokenTransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PMintyMultiTokenTransferBatch represents a TransferBatch event raised by the PMintyMultiToken contract.
type PMintyMultiTokenTransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_PMintyMultiToken *PMintyMultiTokenFilterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*PMintyMultiTokenTransferBatchIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PMintyMultiToken.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PMintyMultiTokenTransferBatchIterator{contract: _PMintyMultiToken.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_PMintyMultiToken *PMintyMultiTokenFilterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *PMintyMultiTokenTransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PMintyMultiToken.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PMintyMultiTokenTransferBatch)
				if err := _PMintyMultiToken.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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

// ParseTransferBatch is a log parse operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_PMintyMultiToken *PMintyMultiTokenFilterer) ParseTransferBatch(log types.Log) (*PMintyMultiTokenTransferBatch, error) {
	event := new(PMintyMultiTokenTransferBatch)
	if err := _PMintyMultiToken.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PMintyMultiTokenTransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the PMintyMultiToken contract.
type PMintyMultiTokenTransferSingleIterator struct {
	Event *PMintyMultiTokenTransferSingle // Event containing the contract specifics and raw log

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
func (it *PMintyMultiTokenTransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PMintyMultiTokenTransferSingle)
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
		it.Event = new(PMintyMultiTokenTransferSingle)
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
func (it *PMintyMultiTokenTransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PMintyMultiTokenTransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PMintyMultiTokenTransferSingle represents a TransferSingle event raised by the PMintyMultiToken contract.
type PMintyMultiTokenTransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_PMintyMultiToken *PMintyMultiTokenFilterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*PMintyMultiTokenTransferSingleIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PMintyMultiToken.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PMintyMultiTokenTransferSingleIterator{contract: _PMintyMultiToken.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_PMintyMultiToken *PMintyMultiTokenFilterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *PMintyMultiTokenTransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PMintyMultiToken.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PMintyMultiTokenTransferSingle)
				if err := _PMintyMultiToken.contract.UnpackLog(event, "TransferSingle", log); err != nil {
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

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_PMintyMultiToken *PMintyMultiTokenFilterer) ParseTransferSingle(log types.Log) (*PMintyMultiTokenTransferSingle, error) {
	event := new(PMintyMultiTokenTransferSingle)
	if err := _PMintyMultiToken.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PMintyMultiTokenURIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the PMintyMultiToken contract.
type PMintyMultiTokenURIIterator struct {
	Event *PMintyMultiTokenURI // Event containing the contract specifics and raw log

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
func (it *PMintyMultiTokenURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PMintyMultiTokenURI)
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
		it.Event = new(PMintyMultiTokenURI)
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
func (it *PMintyMultiTokenURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PMintyMultiTokenURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PMintyMultiTokenURI represents a URI event raised by the PMintyMultiToken contract.
type PMintyMultiTokenURI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_PMintyMultiToken *PMintyMultiTokenFilterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*PMintyMultiTokenURIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _PMintyMultiToken.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &PMintyMultiTokenURIIterator{contract: _PMintyMultiToken.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_PMintyMultiToken *PMintyMultiTokenFilterer) WatchURI(opts *bind.WatchOpts, sink chan<- *PMintyMultiTokenURI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _PMintyMultiToken.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PMintyMultiTokenURI)
				if err := _PMintyMultiToken.contract.UnpackLog(event, "URI", log); err != nil {
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

// ParseURI is a log parse operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_PMintyMultiToken *PMintyMultiTokenFilterer) ParseURI(log types.Log) (*PMintyMultiTokenURI, error) {
	event := new(PMintyMultiTokenURI)
	if err := _PMintyMultiToken.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
