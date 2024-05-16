// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ContractsMetaData contains all meta data concerning the Contracts contract.
var ContractsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"StringAdded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_str\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_createdTime\",\"type\":\"uint256\"}],\"name\":\"addString\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_str\",\"type\":\"string\"}],\"name\":\"getString\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f80fd5b50335f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506105838061005b5f395ff3fe608060405234801561000f575f80fd5b5060043610610034575f3560e01c80639c981fcb14610038578063ac5dc56f14610068575b5f80fd5b610052600480360381019061004d91906102e5565b610084565b60405161005f9190610344565b60405180910390f35b610082600480360381019061007d9190610387565b6100ab565b005b5f6001826040516100959190610433565b9081526020016040518091039020549050919050565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610138576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161012f906104c9565b60405180910390fd5b806001836040516101499190610433565b9081526020016040518091039020819055507f1272cb554134672a55be274a7f45d1c3cccd3582735427f5281153024b8c0242828260405161018c92919061051f565b60405180910390a15050565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6101f7826101b1565b810181811067ffffffffffffffff82111715610216576102156101c1565b5b80604052505050565b5f610228610198565b905061023482826101ee565b919050565b5f67ffffffffffffffff821115610253576102526101c1565b5b61025c826101b1565b9050602081019050919050565b828183375f83830152505050565b5f61028961028484610239565b61021f565b9050828152602081018484840111156102a5576102a46101ad565b5b6102b0848285610269565b509392505050565b5f82601f8301126102cc576102cb6101a9565b5b81356102dc848260208601610277565b91505092915050565b5f602082840312156102fa576102f96101a1565b5b5f82013567ffffffffffffffff811115610317576103166101a5565b5b610323848285016102b8565b91505092915050565b5f819050919050565b61033e8161032c565b82525050565b5f6020820190506103575f830184610335565b92915050565b6103668161032c565b8114610370575f80fd5b50565b5f813590506103818161035d565b92915050565b5f806040838503121561039d5761039c6101a1565b5b5f83013567ffffffffffffffff8111156103ba576103b96101a5565b5b6103c6858286016102b8565b92505060206103d785828601610373565b9150509250929050565b5f81519050919050565b5f81905092915050565b8281835e5f83830152505050565b5f61040d826103e1565b61041781856103eb565b93506104278185602086016103f5565b80840191505092915050565b5f61043e8284610403565b915081905092915050565b5f82825260208201905092915050565b7f4f6e6c79206f776e65722063616e2063616c6c20746869732066756e6374696f5f8201527f6e00000000000000000000000000000000000000000000000000000000000000602082015250565b5f6104b3602183610449565b91506104be82610459565b604082019050919050565b5f6020820190508181035f8301526104e0816104a7565b9050919050565b5f6104f1826103e1565b6104fb8185610449565b935061050b8185602086016103f5565b610514816101b1565b840191505092915050565b5f6040820190508181035f83015261053781856104e7565b90506105466020830184610335565b939250505056fea264697066735822122025d968e5f98354c4b232e1d54eb9a162653d9b77362480f40cd42b6f48042ef764736f6c63430008190033",
}

// ContractsABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractsMetaData.ABI instead.
var ContractsABI = ContractsMetaData.ABI

// ContractsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractsMetaData.Bin instead.
var ContractsBin = ContractsMetaData.Bin

// DeployContracts deploys a new Ethereum contract, binding an instance of Contracts to it.
func DeployContracts(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contracts, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// Contracts is an auto generated Go binding around an Ethereum contract.
type Contracts struct {
	ContractsCaller     // Read-only binding to the contract
	ContractsTransactor // Write-only binding to the contract
	ContractsFilterer   // Log filterer for contract events
}

// ContractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractsSession struct {
	Contract     *Contracts        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractsCallerSession struct {
	Contract *ContractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ContractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractsTransactorSession struct {
	Contract     *ContractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ContractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractsRaw struct {
	Contract *Contracts // Generic contract binding to access the raw methods on
}

// ContractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractsCallerRaw struct {
	Contract *ContractsCaller // Generic read-only contract binding to access the raw methods on
}

// ContractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractsTransactorRaw struct {
	Contract *ContractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContracts creates a new instance of Contracts, bound to a specific deployed contract.
func NewContracts(address common.Address, backend bind.ContractBackend) (*Contracts, error) {
	contract, err := bindContracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// NewContractsCaller creates a new read-only instance of Contracts, bound to a specific deployed contract.
func NewContractsCaller(address common.Address, caller bind.ContractCaller) (*ContractsCaller, error) {
	contract, err := bindContracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsCaller{contract: contract}, nil
}

// NewContractsTransactor creates a new write-only instance of Contracts, bound to a specific deployed contract.
func NewContractsTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractsTransactor, error) {
	contract, err := bindContracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsTransactor{contract: contract}, nil
}

// NewContractsFilterer creates a new log filterer instance of Contracts, bound to a specific deployed contract.
func NewContractsFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractsFilterer, error) {
	contract, err := bindContracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractsFilterer{contract: contract}, nil
}

// bindContracts binds a generic wrapper to an already deployed contract.
func bindContracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.ContractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transact(opts, method, params...)
}

// GetString is a free data retrieval call binding the contract method 0x9c981fcb.
//
// Solidity: function getString(string _str) view returns(uint256)
func (_Contracts *ContractsCaller) GetString(opts *bind.CallOpts, _str string) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getString", _str)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetString is a free data retrieval call binding the contract method 0x9c981fcb.
//
// Solidity: function getString(string _str) view returns(uint256)
func (_Contracts *ContractsSession) GetString(_str string) (*big.Int, error) {
	return _Contracts.Contract.GetString(&_Contracts.CallOpts, _str)
}

// GetString is a free data retrieval call binding the contract method 0x9c981fcb.
//
// Solidity: function getString(string _str) view returns(uint256)
func (_Contracts *ContractsCallerSession) GetString(_str string) (*big.Int, error) {
	return _Contracts.Contract.GetString(&_Contracts.CallOpts, _str)
}

// AddString is a paid mutator transaction binding the contract method 0xac5dc56f.
//
// Solidity: function addString(string _str, uint256 _createdTime) returns()
func (_Contracts *ContractsTransactor) AddString(opts *bind.TransactOpts, _str string, _createdTime *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "addString", _str, _createdTime)
}

// AddString is a paid mutator transaction binding the contract method 0xac5dc56f.
//
// Solidity: function addString(string _str, uint256 _createdTime) returns()
func (_Contracts *ContractsSession) AddString(_str string, _createdTime *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.AddString(&_Contracts.TransactOpts, _str, _createdTime)
}

// AddString is a paid mutator transaction binding the contract method 0xac5dc56f.
//
// Solidity: function addString(string _str, uint256 _createdTime) returns()
func (_Contracts *ContractsTransactorSession) AddString(_str string, _createdTime *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.AddString(&_Contracts.TransactOpts, _str, _createdTime)
}

// ContractsStringAddedIterator is returned from FilterStringAdded and is used to iterate over the raw logs and unpacked data for StringAdded events raised by the Contracts contract.
type ContractsStringAddedIterator struct {
	Event *ContractsStringAdded // Event containing the contract specifics and raw log

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
func (it *ContractsStringAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsStringAdded)
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
		it.Event = new(ContractsStringAdded)
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
func (it *ContractsStringAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsStringAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsStringAdded represents a StringAdded event raised by the Contracts contract.
type ContractsStringAdded struct {
	Arg0 string
	Arg1 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterStringAdded is a free log retrieval operation binding the contract event 0x1272cb554134672a55be274a7f45d1c3cccd3582735427f5281153024b8c0242.
//
// Solidity: event StringAdded(string arg0, uint256 arg1)
func (_Contracts *ContractsFilterer) FilterStringAdded(opts *bind.FilterOpts) (*ContractsStringAddedIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "StringAdded")
	if err != nil {
		return nil, err
	}
	return &ContractsStringAddedIterator{contract: _Contracts.contract, event: "StringAdded", logs: logs, sub: sub}, nil
}

// WatchStringAdded is a free log subscription operation binding the contract event 0x1272cb554134672a55be274a7f45d1c3cccd3582735427f5281153024b8c0242.
//
// Solidity: event StringAdded(string arg0, uint256 arg1)
func (_Contracts *ContractsFilterer) WatchStringAdded(opts *bind.WatchOpts, sink chan<- *ContractsStringAdded) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "StringAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsStringAdded)
				if err := _Contracts.contract.UnpackLog(event, "StringAdded", log); err != nil {
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

// ParseStringAdded is a log parse operation binding the contract event 0x1272cb554134672a55be274a7f45d1c3cccd3582735427f5281153024b8c0242.
//
// Solidity: event StringAdded(string arg0, uint256 arg1)
func (_Contracts *ContractsFilterer) ParseStringAdded(log types.Log) (*ContractsStringAdded, error) {
	event := new(ContractsStringAdded)
	if err := _Contracts.contract.UnpackLog(event, "StringAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
