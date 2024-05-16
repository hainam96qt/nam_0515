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

// AttendanceContractAttendanceRecord is an auto generated low-level Go binding around an user-defined struct.
type AttendanceContractAttendanceRecord struct {
	EmployeeId   *big.Int
	Index        *big.Int
	CheckInTime  *big.Int
	CheckOutTime *big.Int
	Details      string
}

// ContractsMetaData contains all meta data concerning the Contracts contract.
var ContractsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_entity\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_employeeID\",\"type\":\"uint256\"}],\"name\":\"authorizeEntity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_employeeId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_checkInTime\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_details\",\"type\":\"string\"}],\"name\":\"checkIn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_employeeId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_newCheckOutTime\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_newDetails\",\"type\":\"string\"}],\"name\":\"checkOut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_employeeId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_end\",\"type\":\"uint256\"}],\"name\":\"getAttendanceByDateRange\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"employeeId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"checkInTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"checkOutTime\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"details\",\"type\":\"string\"}],\"internalType\":\"structAttendanceContract.AttendanceRecord[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_employeeId\",\"type\":\"uint256\"}],\"name\":\"getAttendanceByEmployeeId\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"employeeId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"checkInTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"checkOutTime\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"details\",\"type\":\"string\"}],\"internalType\":\"structAttendanceContract.AttendanceRecord[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_entity\",\"type\":\"address\"}],\"name\":\"revokeAuthorization\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_employeeId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_newCheckInTime\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_newDetails\",\"type\":\"string\"}],\"name\":\"updateCheckIn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f80fd5b50620f424060015f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055506117b2806100615f395ff3fe608060405234801561000f575f80fd5b506004361061007a575f3560e01c8063adc6ca2a11610059578063adc6ca2a146100e6578063b48028e314610102578063dd0826c61461011e578063e0da7e1a1461014e5761007a565b806242a39a1461007e5780632b23f0a71461009a5780639f43f46f146100b6575b5f80fd5b61009860048036038101906100939190610d33565b61016a565b005b6100b460048036038101906100af9190610ead565b61022f565b005b6100d060048036038101906100cb9190610f2d565b610461565b6040516100dd919061111a565b60405180910390f35b61010060048036038101906100fb919061113a565b610688565b005b61011c600480360381019061011791906111a6565b610851565b005b610138600480360381019061013391906111d1565b610915565b604051610145919061111a565b60405180910390f35b61016860048036038101906101639190610ead565b610a38565b005b5f60015f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054036101e9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101e090611256565b60405180910390fd5b8060015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055505050565b5f60015f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054036102ae576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102a590611256565b60405180910390fd5b5f808581526020019081526020015f20805490508310610303576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102fa906112be565b60405180910390fd5b620f424060015f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205414158015610390575060015f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20548414155b61045b57815f808681526020019081526020015f2084815481106103b7576103b66112dc565b5b905f5260205f209060050201600301819055505f808581526020019081526020015f2083815481106103ec576103eb6112dc565b5b905f5260205f2090600502016004018160405160200161040d92919061147c565b6040516020818303038152906040525f808681526020019081526020015f20848154811061043e5761043d6112dc565b5b905f5260205f20906005020160040190816104599190611635565b505b50505050565b60605f805f8681526020019081526020015f20805480602002602001604051908101604052809291908181526020015f905b8282101561057a578382905f5260205f2090600502016040518060a00160405290815f82015481526020016001820154815260200160028201548152602001600382015481526020016004820180546104eb90611336565b80601f016020809104026020016040519081016040528092919081815260200182805461051790611336565b80156105625780601f1061053957610100808354040283529160200191610562565b820191905f5260205f20905b81548152906001019060200180831161054557829003601f168201915b50505050508152505081526020019060010190610493565b5050505090505f815167ffffffffffffffff81111561059c5761059b610d89565b5b6040519080825280602002602001820160405280156105d557816020015b6105c2610c6a565b8152602001906001900390816105ba5790505b5090505f5b825181101561067b57858382815181106105f7576105f66112dc565b5b6020026020010151604001511015801561062f5750848382815181106106205761061f6112dc565b5b60200260200101516040015111155b1561066e57828181518110610647576106466112dc565b5b6020026020010151828281518110610662576106616112dc565b5b60200260200101819052505b80806001019150506105da565b5080925050509392505050565b5f60015f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205403610707576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106fe90611256565b60405180910390fd5b620f424060015f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205414158015610794575060015f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20548314155b61084c575f805f8581526020019081526020015f208054905090505f808581526020019081526020015f206040518060a001604052808681526020018381526020018581526020015f815260200184815250908060018154018082558091505060019003905f5260205f2090600502015f909190919091505f820151815f015560208201518160010155604082015181600201556060820151816003015560808201518160040190816108479190611635565b505050505b505050565b5f60015f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054036108d0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108c790611256565b60405180910390fd5b5f60015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f208190555050565b60605f808381526020019081526020015f20805480602002602001604051908101604052809291908181526020015f905b82821015610a2d578382905f5260205f2090600502016040518060a00160405290815f820154815260200160018201548152602001600282015481526020016003820154815260200160048201805461099e90611336565b80601f01602080910402602001604051908101604052809291908181526020018280546109ca90611336565b8015610a155780601f106109ec57610100808354040283529160200191610a15565b820191905f5260205f20905b8154815290600101906020018083116109f857829003601f168201915b50505050508152505081526020019060010190610946565b505050509050919050565b5f60015f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205403610ab7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610aae90611256565b60405180910390fd5b5f808581526020019081526020015f20805490508310610b0c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b03906112be565b60405180910390fd5b620f424060015f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205414158015610b99575060015f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20548414155b610c6457815f808681526020019081526020015f208481548110610bc057610bbf6112dc565b5b905f5260205f209060050201600201819055505f808581526020019081526020015f208381548110610bf557610bf46112dc565b5b905f5260205f20906005020160040181604051602001610c1692919061174e565b6040516020818303038152906040525f808681526020019081526020015f208481548110610c4757610c466112dc565b5b905f5260205f2090600502016004019081610c629190611635565b505b50505050565b6040518060a001604052805f81526020015f81526020015f81526020015f8152602001606081525090565b5f604051905090565b5f80fd5b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610ccf82610ca6565b9050919050565b610cdf81610cc5565b8114610ce9575f80fd5b50565b5f81359050610cfa81610cd6565b92915050565b5f819050919050565b610d1281610d00565b8114610d1c575f80fd5b50565b5f81359050610d2d81610d09565b92915050565b5f8060408385031215610d4957610d48610c9e565b5b5f610d5685828601610cec565b9250506020610d6785828601610d1f565b9150509250929050565b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610dbf82610d79565b810181811067ffffffffffffffff82111715610dde57610ddd610d89565b5b80604052505050565b5f610df0610c95565b9050610dfc8282610db6565b919050565b5f67ffffffffffffffff821115610e1b57610e1a610d89565b5b610e2482610d79565b9050602081019050919050565b828183375f83830152505050565b5f610e51610e4c84610e01565b610de7565b905082815260208101848484011115610e6d57610e6c610d75565b5b610e78848285610e31565b509392505050565b5f82601f830112610e9457610e93610d71565b5b8135610ea4848260208601610e3f565b91505092915050565b5f805f8060808587031215610ec557610ec4610c9e565b5b5f610ed287828801610d1f565b9450506020610ee387828801610d1f565b9350506040610ef487828801610d1f565b925050606085013567ffffffffffffffff811115610f1557610f14610ca2565b5b610f2187828801610e80565b91505092959194509250565b5f805f60608486031215610f4457610f43610c9e565b5b5f610f5186828701610d1f565b9350506020610f6286828701610d1f565b9250506040610f7386828701610d1f565b9150509250925092565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b610faf81610d00565b82525050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f610fe782610fb5565b610ff18185610fbf565b9350611001818560208601610fcf565b61100a81610d79565b840191505092915050565b5f60a083015f83015161102a5f860182610fa6565b50602083015161103d6020860182610fa6565b5060408301516110506040860182610fa6565b5060608301516110636060860182610fa6565b506080830151848203608086015261107b8282610fdd565b9150508091505092915050565b5f6110938383611015565b905092915050565b5f602082019050919050565b5f6110b182610f7d565b6110bb8185610f87565b9350836020820285016110cd85610f97565b805f5b8581101561110857848403895281516110e98582611088565b94506110f48361109b565b925060208a019950506001810190506110d0565b50829750879550505050505092915050565b5f6020820190508181035f83015261113281846110a7565b905092915050565b5f805f6060848603121561115157611150610c9e565b5b5f61115e86828701610d1f565b935050602061116f86828701610d1f565b925050604084013567ffffffffffffffff8111156111905761118f610ca2565b5b61119c86828701610e80565b9150509250925092565b5f602082840312156111bb576111ba610c9e565b5b5f6111c884828501610cec565b91505092915050565b5f602082840312156111e6576111e5610c9e565b5b5f6111f384828501610d1f565b91505092915050565b5f82825260208201905092915050565b7f556e617574686f72697a656420616363657373000000000000000000000000005f82015250565b5f6112406013836111fc565b915061124b8261120c565b602082019050919050565b5f6020820190508181035f83015261126d81611234565b9050919050565b7f496e646578206f7574206f6620626f756e6473000000000000000000000000005f82015250565b5f6112a86013836111fc565b91506112b382611274565b602082019050919050565b5f6020820190508181035f8301526112d58161129c565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061134d57607f821691505b6020821081036113605761135f611309565b5b50919050565b5f81905092915050565b5f819050815f5260205f209050919050565b5f815461138e81611336565b6113988186611366565b9450600182165f81146113b257600181146113c7576113f9565b60ff19831686528115158202860193506113f9565b6113d085611370565b5f5b838110156113f1578154818901526001820191506020810190506113d2565b838801955050505b50505092915050565b7f2e20636865636b6f75743a2000000000000000000000000000000000000000005f82015250565b5f611436600c83611366565b915061144182611402565b600c82019050919050565b5f61145682610fb5565b6114608185611366565b9350611470818560208601610fcf565b80840191505092915050565b5f6114878285611382565b91506114928261142a565b915061149e828461144c565b91508190509392505050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026114f47fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826114b9565b6114fe86836114b9565b95508019841693508086168417925050509392505050565b5f819050919050565b5f61153961153461152f84610d00565b611516565b610d00565b9050919050565b5f819050919050565b6115528361151f565b61156661155e82611540565b8484546114c5565b825550505050565b5f90565b61157a61156e565b611585818484611549565b505050565b5b818110156115a85761159d5f82611572565b60018101905061158b565b5050565b601f8211156115ed576115be81611370565b6115c7846114aa565b810160208510156115d6578190505b6115ea6115e2856114aa565b83018261158a565b50505b505050565b5f82821c905092915050565b5f61160d5f19846008026115f2565b1980831691505092915050565b5f61162583836115fe565b9150826002028217905092915050565b61163e82610fb5565b67ffffffffffffffff81111561165757611656610d89565b5b6116618254611336565b61166c8282856115ac565b5f60209050601f83116001811461169d575f841561168b578287015190505b611695858261161a565b8655506116fc565b601f1984166116ab86611370565b5f5b828110156116d2578489015182556001820191506020850194506020810190506116ad565b868310156116ef57848901516116eb601f8916826115fe565b8355505b6001600288020188555050505b505050505050565b7f2e2075706461746520636865636b696e20726561736f6e203a200000000000005f82015250565b5f611738601a83611366565b915061174382611704565b601a82019050919050565b5f6117598285611382565b91506117648261172c565b9150611770828461144c565b9150819050939250505056fea2646970667358221220f15c2680b35ebf0ea56e600f35ddee2a2f099c3c019aa06e0a3fed5bc9459b5564736f6c63430008190033",
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

// GetAttendanceByDateRange is a free data retrieval call binding the contract method 0x9f43f46f.
//
// Solidity: function getAttendanceByDateRange(uint256 _employeeId, uint256 _start, uint256 _end) view returns((uint256,uint256,uint256,uint256,string)[])
func (_Contracts *ContractsCaller) GetAttendanceByDateRange(opts *bind.CallOpts, _employeeId *big.Int, _start *big.Int, _end *big.Int) ([]AttendanceContractAttendanceRecord, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getAttendanceByDateRange", _employeeId, _start, _end)

	if err != nil {
		return *new([]AttendanceContractAttendanceRecord), err
	}

	out0 := *abi.ConvertType(out[0], new([]AttendanceContractAttendanceRecord)).(*[]AttendanceContractAttendanceRecord)

	return out0, err

}

// GetAttendanceByDateRange is a free data retrieval call binding the contract method 0x9f43f46f.
//
// Solidity: function getAttendanceByDateRange(uint256 _employeeId, uint256 _start, uint256 _end) view returns((uint256,uint256,uint256,uint256,string)[])
func (_Contracts *ContractsSession) GetAttendanceByDateRange(_employeeId *big.Int, _start *big.Int, _end *big.Int) ([]AttendanceContractAttendanceRecord, error) {
	return _Contracts.Contract.GetAttendanceByDateRange(&_Contracts.CallOpts, _employeeId, _start, _end)
}

// GetAttendanceByDateRange is a free data retrieval call binding the contract method 0x9f43f46f.
//
// Solidity: function getAttendanceByDateRange(uint256 _employeeId, uint256 _start, uint256 _end) view returns((uint256,uint256,uint256,uint256,string)[])
func (_Contracts *ContractsCallerSession) GetAttendanceByDateRange(_employeeId *big.Int, _start *big.Int, _end *big.Int) ([]AttendanceContractAttendanceRecord, error) {
	return _Contracts.Contract.GetAttendanceByDateRange(&_Contracts.CallOpts, _employeeId, _start, _end)
}

// GetAttendanceByEmployeeId is a free data retrieval call binding the contract method 0xdd0826c6.
//
// Solidity: function getAttendanceByEmployeeId(uint256 _employeeId) view returns((uint256,uint256,uint256,uint256,string)[])
func (_Contracts *ContractsCaller) GetAttendanceByEmployeeId(opts *bind.CallOpts, _employeeId *big.Int) ([]AttendanceContractAttendanceRecord, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getAttendanceByEmployeeId", _employeeId)

	if err != nil {
		return *new([]AttendanceContractAttendanceRecord), err
	}

	out0 := *abi.ConvertType(out[0], new([]AttendanceContractAttendanceRecord)).(*[]AttendanceContractAttendanceRecord)

	return out0, err

}

// GetAttendanceByEmployeeId is a free data retrieval call binding the contract method 0xdd0826c6.
//
// Solidity: function getAttendanceByEmployeeId(uint256 _employeeId) view returns((uint256,uint256,uint256,uint256,string)[])
func (_Contracts *ContractsSession) GetAttendanceByEmployeeId(_employeeId *big.Int) ([]AttendanceContractAttendanceRecord, error) {
	return _Contracts.Contract.GetAttendanceByEmployeeId(&_Contracts.CallOpts, _employeeId)
}

// GetAttendanceByEmployeeId is a free data retrieval call binding the contract method 0xdd0826c6.
//
// Solidity: function getAttendanceByEmployeeId(uint256 _employeeId) view returns((uint256,uint256,uint256,uint256,string)[])
func (_Contracts *ContractsCallerSession) GetAttendanceByEmployeeId(_employeeId *big.Int) ([]AttendanceContractAttendanceRecord, error) {
	return _Contracts.Contract.GetAttendanceByEmployeeId(&_Contracts.CallOpts, _employeeId)
}

// AuthorizeEntity is a paid mutator transaction binding the contract method 0x0042a39a.
//
// Solidity: function authorizeEntity(address _entity, uint256 _employeeID) returns()
func (_Contracts *ContractsTransactor) AuthorizeEntity(opts *bind.TransactOpts, _entity common.Address, _employeeID *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "authorizeEntity", _entity, _employeeID)
}

// AuthorizeEntity is a paid mutator transaction binding the contract method 0x0042a39a.
//
// Solidity: function authorizeEntity(address _entity, uint256 _employeeID) returns()
func (_Contracts *ContractsSession) AuthorizeEntity(_entity common.Address, _employeeID *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.AuthorizeEntity(&_Contracts.TransactOpts, _entity, _employeeID)
}

// AuthorizeEntity is a paid mutator transaction binding the contract method 0x0042a39a.
//
// Solidity: function authorizeEntity(address _entity, uint256 _employeeID) returns()
func (_Contracts *ContractsTransactorSession) AuthorizeEntity(_entity common.Address, _employeeID *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.AuthorizeEntity(&_Contracts.TransactOpts, _entity, _employeeID)
}

// CheckIn is a paid mutator transaction binding the contract method 0xadc6ca2a.
//
// Solidity: function checkIn(uint256 _employeeId, uint256 _checkInTime, string _details) returns()
func (_Contracts *ContractsTransactor) CheckIn(opts *bind.TransactOpts, _employeeId *big.Int, _checkInTime *big.Int, _details string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "checkIn", _employeeId, _checkInTime, _details)
}

// CheckIn is a paid mutator transaction binding the contract method 0xadc6ca2a.
//
// Solidity: function checkIn(uint256 _employeeId, uint256 _checkInTime, string _details) returns()
func (_Contracts *ContractsSession) CheckIn(_employeeId *big.Int, _checkInTime *big.Int, _details string) (*types.Transaction, error) {
	return _Contracts.Contract.CheckIn(&_Contracts.TransactOpts, _employeeId, _checkInTime, _details)
}

// CheckIn is a paid mutator transaction binding the contract method 0xadc6ca2a.
//
// Solidity: function checkIn(uint256 _employeeId, uint256 _checkInTime, string _details) returns()
func (_Contracts *ContractsTransactorSession) CheckIn(_employeeId *big.Int, _checkInTime *big.Int, _details string) (*types.Transaction, error) {
	return _Contracts.Contract.CheckIn(&_Contracts.TransactOpts, _employeeId, _checkInTime, _details)
}

// CheckOut is a paid mutator transaction binding the contract method 0x2b23f0a7.
//
// Solidity: function checkOut(uint256 _employeeId, uint256 _index, uint256 _newCheckOutTime, string _newDetails) returns()
func (_Contracts *ContractsTransactor) CheckOut(opts *bind.TransactOpts, _employeeId *big.Int, _index *big.Int, _newCheckOutTime *big.Int, _newDetails string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "checkOut", _employeeId, _index, _newCheckOutTime, _newDetails)
}

// CheckOut is a paid mutator transaction binding the contract method 0x2b23f0a7.
//
// Solidity: function checkOut(uint256 _employeeId, uint256 _index, uint256 _newCheckOutTime, string _newDetails) returns()
func (_Contracts *ContractsSession) CheckOut(_employeeId *big.Int, _index *big.Int, _newCheckOutTime *big.Int, _newDetails string) (*types.Transaction, error) {
	return _Contracts.Contract.CheckOut(&_Contracts.TransactOpts, _employeeId, _index, _newCheckOutTime, _newDetails)
}

// CheckOut is a paid mutator transaction binding the contract method 0x2b23f0a7.
//
// Solidity: function checkOut(uint256 _employeeId, uint256 _index, uint256 _newCheckOutTime, string _newDetails) returns()
func (_Contracts *ContractsTransactorSession) CheckOut(_employeeId *big.Int, _index *big.Int, _newCheckOutTime *big.Int, _newDetails string) (*types.Transaction, error) {
	return _Contracts.Contract.CheckOut(&_Contracts.TransactOpts, _employeeId, _index, _newCheckOutTime, _newDetails)
}

// RevokeAuthorization is a paid mutator transaction binding the contract method 0xb48028e3.
//
// Solidity: function revokeAuthorization(address _entity) returns()
func (_Contracts *ContractsTransactor) RevokeAuthorization(opts *bind.TransactOpts, _entity common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "revokeAuthorization", _entity)
}

// RevokeAuthorization is a paid mutator transaction binding the contract method 0xb48028e3.
//
// Solidity: function revokeAuthorization(address _entity) returns()
func (_Contracts *ContractsSession) RevokeAuthorization(_entity common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.RevokeAuthorization(&_Contracts.TransactOpts, _entity)
}

// RevokeAuthorization is a paid mutator transaction binding the contract method 0xb48028e3.
//
// Solidity: function revokeAuthorization(address _entity) returns()
func (_Contracts *ContractsTransactorSession) RevokeAuthorization(_entity common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.RevokeAuthorization(&_Contracts.TransactOpts, _entity)
}

// UpdateCheckIn is a paid mutator transaction binding the contract method 0xe0da7e1a.
//
// Solidity: function updateCheckIn(uint256 _employeeId, uint256 _index, uint256 _newCheckInTime, string _newDetails) returns()
func (_Contracts *ContractsTransactor) UpdateCheckIn(opts *bind.TransactOpts, _employeeId *big.Int, _index *big.Int, _newCheckInTime *big.Int, _newDetails string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "updateCheckIn", _employeeId, _index, _newCheckInTime, _newDetails)
}

// UpdateCheckIn is a paid mutator transaction binding the contract method 0xe0da7e1a.
//
// Solidity: function updateCheckIn(uint256 _employeeId, uint256 _index, uint256 _newCheckInTime, string _newDetails) returns()
func (_Contracts *ContractsSession) UpdateCheckIn(_employeeId *big.Int, _index *big.Int, _newCheckInTime *big.Int, _newDetails string) (*types.Transaction, error) {
	return _Contracts.Contract.UpdateCheckIn(&_Contracts.TransactOpts, _employeeId, _index, _newCheckInTime, _newDetails)
}

// UpdateCheckIn is a paid mutator transaction binding the contract method 0xe0da7e1a.
//
// Solidity: function updateCheckIn(uint256 _employeeId, uint256 _index, uint256 _newCheckInTime, string _newDetails) returns()
func (_Contracts *ContractsTransactorSession) UpdateCheckIn(_employeeId *big.Int, _index *big.Int, _newCheckInTime *big.Int, _newDetails string) (*types.Transaction, error) {
	return _Contracts.Contract.UpdateCheckIn(&_Contracts.TransactOpts, _employeeId, _index, _newCheckInTime, _newDetails)
}
