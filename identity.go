// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package identity

import (
	"math/big"
	"strings"

	cpchain "bitbucket.org/cpchain/chain"
	"bitbucket.org/cpchain/chain/accounts/abi"
	"bitbucket.org/cpchain/chain/accounts/abi/bind"
	"bitbucket.org/cpchain/chain/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"
)

// IdentityABI is the input ABI used to generate the binding from.
const IdentityABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"enabled\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"enableContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"disableContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"size\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"remove\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"get\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"content\",\"type\":\"string\"}],\"name\":\"register\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"who\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"identity\",\"type\":\"string\"}],\"name\":\"NewIdentity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"who\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"identity\",\"type\":\"string\"}],\"name\":\"UpdateIdentity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"who\",\"type\":\"address\"}],\"name\":\"RemoveIdentity\",\"type\":\"event\"}]"

// IdentityBin is the compiled bytecode used for deploying new contracts.
const IdentityBin = `0x60806040526000805460a060020a60ff0219167401000000000000000000000000000000000000000017815560055534801561003a57600080fd5b5060008054600160a060020a031916331790556108ba8061005c6000396000f3006080604052600436106100825763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663238dafe08114610087578063367edd32146100b0578063894ba833146100c7578063949d225d146100dc578063a7f4377914610103578063c2bc2efc1461010b578063f2c298be146101a1575b600080fd5b34801561009357600080fd5b5061009c6101ed565b604080519115158252519081900360200190f35b3480156100bc57600080fd5b506100c561020e565b005b3480156100d357600080fd5b506100c561025c565b3480156100e857600080fd5b506100f1610293565b60408051918252519081900360200190f35b6100c561029a565b34801561011757600080fd5b5061012c600160a060020a0360043516610330565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561016657818101518382015260200161014e565b50505050905090810190601f1680156101935780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6040805160206004803580820135601f81018490048402850184019095528484526100c59436949293602493928401919081908401838280828437509497506103f79650505050505050565b60005474010000000000000000000000000000000000000000900460ff1681565b600054600160a060020a0316331461022557600080fd5b6000805474ff0000000000000000000000000000000000000000191674010000000000000000000000000000000000000000179055565b600054600160a060020a0316331461027357600080fd5b6000805474ff000000000000000000000000000000000000000019169055565b6005545b90565b60005474010000000000000000000000000000000000000000900460ff1615156102c357600080fd5b6102d460013363ffffffff6105d516565b15156102df57600080fd5b6102f060013363ffffffff6105f816565b50600580546000190190556040805133815290517fbf872b0739f3bc6313ee30ed542a610dfcb6141e0e7f49ccc2ea8eeaa1735e8a9181900360200190a1565b606061034360018363ffffffff6105d516565b151561034e57600080fd5b600160a060020a03821660009081526004602090815260409182902080548351601f6002600019610100600186161502019093169290920491820184900484028101840190945280845290918301828280156103eb5780601f106103c0576101008083540402835291602001916103eb565b820191906000526020600020905b8154815290600101906020018083116103ce57829003601f168201915b50505050509050919050565b60005474010000000000000000000000000000000000000000900460ff16151561042057600080fd5b3360009081526004602090815260409091208251610440928401906107cd565b5061045260013363ffffffff6105d516565b151561051e5761046960013363ffffffff61073e16565b506005805460010190556040805133808252602080830184815285519484019490945284517fb43d84571ed93e552eba32d34b9bdb8da1e8cda0af378c8c5b413705c00f5a58949293869390929091606084019185019080838360005b838110156104de5781810151838201526020016104c6565b50505050905090810190601f16801561050b5780820380516001836020036101000a031916815260200191505b50935050505060405180910390a16105d2565b7f4d9aa4761c45cac26a1fd345f602649a95c61b251013baf7d61d1ae105bd9a9533826040518083600160a060020a0316600160a060020a0316815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561059657818101518382015260200161057e565b50505050905090810190601f1680156105c35780820380516001836020036101000a031916815260200191505b50935050505060405180910390a15b50565b600160a060020a03811660009081526020839052604090205460ff165b92915050565b600160a060020a03811660009081526020839052604081205481908190819060ff1615156106295760009350610735565b600160a060020a038516600090815260208781526040808320805460ff1916905560028901805460018b019093529220549094509250600019840184811061066d57fe5b600091825260209091200154600287018054600160a060020a03909216925082918490811061069857fe5b6000918252602080832091909101805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039485161790559183168152600188019091526040902082905560028601805460001985019081106106f457fe5b6000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff191690556002860180549061072f90600019830161084b565b50600193505b50505092915050565b600160a060020a03811660009081526020839052604081205460ff1615610767575060006105f2565b50600160a060020a0316600081815260208381526040808320805460ff19166001908117909155600286018054968201845291842086905585810182559083529120909201805473ffffffffffffffffffffffffffffffffffffffff1916909117905590565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061080e57805160ff191683800117855561083b565b8280016001018555821561083b579182015b8281111561083b578251825591602001919060010190610820565b50610847929150610874565b5090565b81548183558181111561086f5760008381526020902061086f918101908301610874565b505050565b61029791905b80821115610847576000815560010161087a5600a165627a7a7230582049bf35f614d3bc6a117d33a925404092d4636aea3ce501a5b1b6a789ab3ec2940029`

// DeployIdentity deploys a new cpchain contract, binding an instance of Identity to it.
func DeployIdentity(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Identity, error) {
	parsed, err := abi.JSON(strings.NewReader(IdentityABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IdentityBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Identity{IdentityCaller: IdentityCaller{contract: contract}, IdentityTransactor: IdentityTransactor{contract: contract}, IdentityFilterer: IdentityFilterer{contract: contract}}, nil
}

// Identity is an auto generated Go binding around an cpchain contract.
type Identity struct {
	IdentityCaller     // Read-only binding to the contract
	IdentityTransactor // Write-only binding to the contract
	IdentityFilterer   // Log filterer for contract events
}

// IdentityCaller is an auto generated read-only Go binding around an cpchain contract.
type IdentityCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityTransactor is an auto generated write-only Go binding around an cpchain contract.
type IdentityTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityFilterer is an auto generated log filtering Go binding around an cpchain contract events.
type IdentityFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentitySession is an auto generated Go binding around an cpchain contract,
// with pre-set call and transact options.
type IdentitySession struct {
	Contract     *Identity         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IdentityCallerSession is an auto generated read-only Go binding around an cpchain contract,
// with pre-set call options.
type IdentityCallerSession struct {
	Contract *IdentityCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IdentityTransactorSession is an auto generated write-only Go binding around an cpchain contract,
// with pre-set transact options.
type IdentityTransactorSession struct {
	Contract     *IdentityTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IdentityRaw is an auto generated low-level Go binding around an cpchain contract.
type IdentityRaw struct {
	Contract *Identity // Generic contract binding to access the raw methods on
}

// IdentityCallerRaw is an auto generated low-level read-only Go binding around an cpchain contract.
type IdentityCallerRaw struct {
	Contract *IdentityCaller // Generic read-only contract binding to access the raw methods on
}

// IdentityTransactorRaw is an auto generated low-level write-only Go binding around an cpchain contract.
type IdentityTransactorRaw struct {
	Contract *IdentityTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIdentity creates a new instance of Identity, bound to a specific deployed contract.
func NewIdentity(address common.Address, backend bind.ContractBackend) (*Identity, error) {
	contract, err := bindIdentity(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Identity{IdentityCaller: IdentityCaller{contract: contract}, IdentityTransactor: IdentityTransactor{contract: contract}, IdentityFilterer: IdentityFilterer{contract: contract}}, nil
}

// NewIdentityCaller creates a new read-only instance of Identity, bound to a specific deployed contract.
func NewIdentityCaller(address common.Address, caller bind.ContractCaller) (*IdentityCaller, error) {
	contract, err := bindIdentity(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IdentityCaller{contract: contract}, nil
}

// NewIdentityTransactor creates a new write-only instance of Identity, bound to a specific deployed contract.
func NewIdentityTransactor(address common.Address, transactor bind.ContractTransactor) (*IdentityTransactor, error) {
	contract, err := bindIdentity(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IdentityTransactor{contract: contract}, nil
}

// NewIdentityFilterer creates a new log filterer instance of Identity, bound to a specific deployed contract.
func NewIdentityFilterer(address common.Address, filterer bind.ContractFilterer) (*IdentityFilterer, error) {
	contract, err := bindIdentity(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IdentityFilterer{contract: contract}, nil
}

// bindIdentity binds a generic wrapper to an already deployed contract.
func bindIdentity(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IdentityABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Identity *IdentityRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Identity.Contract.IdentityCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Identity *IdentityRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Identity.Contract.IdentityTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Identity *IdentityRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Identity.Contract.IdentityTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Identity *IdentityCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Identity.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Identity *IdentityTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Identity.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Identity *IdentityTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Identity.Contract.contract.Transact(opts, method, params...)
}

// Enabled is a free data retrieval call binding the contract method 0x238dafe0.
//
// Solidity: function enabled() constant returns(bool)
func (_Identity *IdentityCaller) Enabled(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Identity.contract.Call(opts, out, "enabled")
	return *ret0, err
}

// Enabled is a free data retrieval call binding the contract method 0x238dafe0.
//
// Solidity: function enabled() constant returns(bool)
func (_Identity *IdentitySession) Enabled() (bool, error) {
	return _Identity.Contract.Enabled(&_Identity.CallOpts)
}

// Enabled is a free data retrieval call binding the contract method 0x238dafe0.
//
// Solidity: function enabled() constant returns(bool)
func (_Identity *IdentityCallerSession) Enabled() (bool, error) {
	return _Identity.Contract.Enabled(&_Identity.CallOpts)
}

// Get is a free data retrieval call binding the contract method 0xc2bc2efc.
//
// Solidity: function get(addr address) constant returns(string)
func (_Identity *IdentityCaller) Get(opts *bind.CallOpts, addr common.Address) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Identity.contract.Call(opts, out, "get", addr)
	return *ret0, err
}

// Get is a free data retrieval call binding the contract method 0xc2bc2efc.
//
// Solidity: function get(addr address) constant returns(string)
func (_Identity *IdentitySession) Get(addr common.Address) (string, error) {
	return _Identity.Contract.Get(&_Identity.CallOpts, addr)
}

// Get is a free data retrieval call binding the contract method 0xc2bc2efc.
//
// Solidity: function get(addr address) constant returns(string)
func (_Identity *IdentityCallerSession) Get(addr common.Address) (string, error) {
	return _Identity.Contract.Get(&_Identity.CallOpts, addr)
}

// Size is a free data retrieval call binding the contract method 0x949d225d.
//
// Solidity: function size() constant returns(uint256)
func (_Identity *IdentityCaller) Size(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Identity.contract.Call(opts, out, "size")
	return *ret0, err
}

// Size is a free data retrieval call binding the contract method 0x949d225d.
//
// Solidity: function size() constant returns(uint256)
func (_Identity *IdentitySession) Size() (*big.Int, error) {
	return _Identity.Contract.Size(&_Identity.CallOpts)
}

// Size is a free data retrieval call binding the contract method 0x949d225d.
//
// Solidity: function size() constant returns(uint256)
func (_Identity *IdentityCallerSession) Size() (*big.Int, error) {
	return _Identity.Contract.Size(&_Identity.CallOpts)
}

// DisableContract is a paid mutator transaction binding the contract method 0x894ba833.
//
// Solidity: function disableContract() returns()
func (_Identity *IdentityTransactor) DisableContract(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Identity.contract.Transact(opts, "disableContract")
}

// DisableContract is a paid mutator transaction binding the contract method 0x894ba833.
//
// Solidity: function disableContract() returns()
func (_Identity *IdentitySession) DisableContract() (*types.Transaction, error) {
	return _Identity.Contract.DisableContract(&_Identity.TransactOpts)
}

// DisableContract is a paid mutator transaction binding the contract method 0x894ba833.
//
// Solidity: function disableContract() returns()
func (_Identity *IdentityTransactorSession) DisableContract() (*types.Transaction, error) {
	return _Identity.Contract.DisableContract(&_Identity.TransactOpts)
}

// EnableContract is a paid mutator transaction binding the contract method 0x367edd32.
//
// Solidity: function enableContract() returns()
func (_Identity *IdentityTransactor) EnableContract(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Identity.contract.Transact(opts, "enableContract")
}

// EnableContract is a paid mutator transaction binding the contract method 0x367edd32.
//
// Solidity: function enableContract() returns()
func (_Identity *IdentitySession) EnableContract() (*types.Transaction, error) {
	return _Identity.Contract.EnableContract(&_Identity.TransactOpts)
}

// EnableContract is a paid mutator transaction binding the contract method 0x367edd32.
//
// Solidity: function enableContract() returns()
func (_Identity *IdentityTransactorSession) EnableContract() (*types.Transaction, error) {
	return _Identity.Contract.EnableContract(&_Identity.TransactOpts)
}

// Register is a paid mutator transaction binding the contract method 0xf2c298be.
//
// Solidity: function register(content string) returns()
func (_Identity *IdentityTransactor) Register(opts *bind.TransactOpts, content string) (*types.Transaction, error) {
	return _Identity.contract.Transact(opts, "register", content)
}

// Register is a paid mutator transaction binding the contract method 0xf2c298be.
//
// Solidity: function register(content string) returns()
func (_Identity *IdentitySession) Register(content string) (*types.Transaction, error) {
	return _Identity.Contract.Register(&_Identity.TransactOpts, content)
}

// Register is a paid mutator transaction binding the contract method 0xf2c298be.
//
// Solidity: function register(content string) returns()
func (_Identity *IdentityTransactorSession) Register(content string) (*types.Transaction, error) {
	return _Identity.Contract.Register(&_Identity.TransactOpts, content)
}

// Remove is a paid mutator transaction binding the contract method 0xa7f43779.
//
// Solidity: function remove() returns()
func (_Identity *IdentityTransactor) Remove(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Identity.contract.Transact(opts, "remove")
}

// Remove is a paid mutator transaction binding the contract method 0xa7f43779.
//
// Solidity: function remove() returns()
func (_Identity *IdentitySession) Remove() (*types.Transaction, error) {
	return _Identity.Contract.Remove(&_Identity.TransactOpts)
}

// Remove is a paid mutator transaction binding the contract method 0xa7f43779.
//
// Solidity: function remove() returns()
func (_Identity *IdentityTransactorSession) Remove() (*types.Transaction, error) {
	return _Identity.Contract.Remove(&_Identity.TransactOpts)
}

// IdentityNewIdentityIterator is returned from FilterNewIdentity and is used to iterate over the raw logs and unpacked data for NewIdentity events raised by the Identity contract.
type IdentityNewIdentityIterator struct {
	Event *IdentityNewIdentity // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log       // Log channel receiving the found contract events
	sub  cpchain.Subscription // Subscription for errors, completion and termination
	done bool                 // Whether the subscription completed delivering logs
	fail error                // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IdentityNewIdentityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityNewIdentity)
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
		it.Event = new(IdentityNewIdentity)
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
func (it *IdentityNewIdentityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityNewIdentityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityNewIdentity represents a NewIdentity event raised by the Identity contract.
type IdentityNewIdentity struct {
	Who      common.Address
	Identity string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewIdentity is a free log retrieval operation binding the contract event 0xb43d84571ed93e552eba32d34b9bdb8da1e8cda0af378c8c5b413705c00f5a58.
//
// Solidity: e NewIdentity(who address, identity string)
func (_Identity *IdentityFilterer) FilterNewIdentity(opts *bind.FilterOpts) (*IdentityNewIdentityIterator, error) {

	logs, sub, err := _Identity.contract.FilterLogs(opts, "NewIdentity")
	if err != nil {
		return nil, err
	}
	return &IdentityNewIdentityIterator{contract: _Identity.contract, event: "NewIdentity", logs: logs, sub: sub}, nil
}

// WatchNewIdentity is a free log subscription operation binding the contract event 0xb43d84571ed93e552eba32d34b9bdb8da1e8cda0af378c8c5b413705c00f5a58.
//
// Solidity: e NewIdentity(who address, identity string)
func (_Identity *IdentityFilterer) WatchNewIdentity(opts *bind.WatchOpts, sink chan<- *IdentityNewIdentity) (event.Subscription, error) {

	logs, sub, err := _Identity.contract.WatchLogs(opts, "NewIdentity")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityNewIdentity)
				if err := _Identity.contract.UnpackLog(event, "NewIdentity", log); err != nil {
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

// IdentityRemoveIdentityIterator is returned from FilterRemoveIdentity and is used to iterate over the raw logs and unpacked data for RemoveIdentity events raised by the Identity contract.
type IdentityRemoveIdentityIterator struct {
	Event *IdentityRemoveIdentity // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log       // Log channel receiving the found contract events
	sub  cpchain.Subscription // Subscription for errors, completion and termination
	done bool                 // Whether the subscription completed delivering logs
	fail error                // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IdentityRemoveIdentityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityRemoveIdentity)
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
		it.Event = new(IdentityRemoveIdentity)
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
func (it *IdentityRemoveIdentityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityRemoveIdentityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityRemoveIdentity represents a RemoveIdentity event raised by the Identity contract.
type IdentityRemoveIdentity struct {
	Who common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRemoveIdentity is a free log retrieval operation binding the contract event 0xbf872b0739f3bc6313ee30ed542a610dfcb6141e0e7f49ccc2ea8eeaa1735e8a.
//
// Solidity: e RemoveIdentity(who address)
func (_Identity *IdentityFilterer) FilterRemoveIdentity(opts *bind.FilterOpts) (*IdentityRemoveIdentityIterator, error) {

	logs, sub, err := _Identity.contract.FilterLogs(opts, "RemoveIdentity")
	if err != nil {
		return nil, err
	}
	return &IdentityRemoveIdentityIterator{contract: _Identity.contract, event: "RemoveIdentity", logs: logs, sub: sub}, nil
}

// WatchRemoveIdentity is a free log subscription operation binding the contract event 0xbf872b0739f3bc6313ee30ed542a610dfcb6141e0e7f49ccc2ea8eeaa1735e8a.
//
// Solidity: e RemoveIdentity(who address)
func (_Identity *IdentityFilterer) WatchRemoveIdentity(opts *bind.WatchOpts, sink chan<- *IdentityRemoveIdentity) (event.Subscription, error) {

	logs, sub, err := _Identity.contract.WatchLogs(opts, "RemoveIdentity")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityRemoveIdentity)
				if err := _Identity.contract.UnpackLog(event, "RemoveIdentity", log); err != nil {
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

// IdentityUpdateIdentityIterator is returned from FilterUpdateIdentity and is used to iterate over the raw logs and unpacked data for UpdateIdentity events raised by the Identity contract.
type IdentityUpdateIdentityIterator struct {
	Event *IdentityUpdateIdentity // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log       // Log channel receiving the found contract events
	sub  cpchain.Subscription // Subscription for errors, completion and termination
	done bool                 // Whether the subscription completed delivering logs
	fail error                // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IdentityUpdateIdentityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityUpdateIdentity)
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
		it.Event = new(IdentityUpdateIdentity)
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
func (it *IdentityUpdateIdentityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityUpdateIdentityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityUpdateIdentity represents a UpdateIdentity event raised by the Identity contract.
type IdentityUpdateIdentity struct {
	Who      common.Address
	Identity string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdateIdentity is a free log retrieval operation binding the contract event 0x4d9aa4761c45cac26a1fd345f602649a95c61b251013baf7d61d1ae105bd9a95.
//
// Solidity: e UpdateIdentity(who address, identity string)
func (_Identity *IdentityFilterer) FilterUpdateIdentity(opts *bind.FilterOpts) (*IdentityUpdateIdentityIterator, error) {

	logs, sub, err := _Identity.contract.FilterLogs(opts, "UpdateIdentity")
	if err != nil {
		return nil, err
	}
	return &IdentityUpdateIdentityIterator{contract: _Identity.contract, event: "UpdateIdentity", logs: logs, sub: sub}, nil
}

// WatchUpdateIdentity is a free log subscription operation binding the contract event 0x4d9aa4761c45cac26a1fd345f602649a95c61b251013baf7d61d1ae105bd9a95.
//
// Solidity: e UpdateIdentity(who address, identity string)
func (_Identity *IdentityFilterer) WatchUpdateIdentity(opts *bind.WatchOpts, sink chan<- *IdentityUpdateIdentity) (event.Subscription, error) {

	logs, sub, err := _Identity.contract.WatchLogs(opts, "UpdateIdentity")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityUpdateIdentity)
				if err := _Identity.contract.UnpackLog(event, "UpdateIdentity", log); err != nil {
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

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
const SafeMathBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a72305820bfcebf208a6b47298db9e91472ad2e4449856125bbdc54864629ed3202a046d10029`

// DeploySafeMath deploys a new cpchain contract, binding an instance of SafeMath to it.
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

// SafeMath is an auto generated Go binding around an cpchain contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an cpchain contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an cpchain contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an cpchain contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an cpchain contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an cpchain contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an cpchain contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an cpchain contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an cpchain contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an cpchain contract.
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
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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

// SetABI is the input ABI used to generate the binding from.
const SetABI = "[]"

// SetBin is the compiled bytecode used for deploying new contracts.
const SetBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a72305820dc4e241056ed9c009ded35c024ea35a74033445b1ecfc6bb46770a1f78670f7d0029`

// DeploySet deploys a new cpchain contract, binding an instance of Set to it.
func DeploySet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Set, error) {
	parsed, err := abi.JSON(strings.NewReader(SetABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SetBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Set{SetCaller: SetCaller{contract: contract}, SetTransactor: SetTransactor{contract: contract}, SetFilterer: SetFilterer{contract: contract}}, nil
}

// Set is an auto generated Go binding around an cpchain contract.
type Set struct {
	SetCaller     // Read-only binding to the contract
	SetTransactor // Write-only binding to the contract
	SetFilterer   // Log filterer for contract events
}

// SetCaller is an auto generated read-only Go binding around an cpchain contract.
type SetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SetTransactor is an auto generated write-only Go binding around an cpchain contract.
type SetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SetFilterer is an auto generated log filtering Go binding around an cpchain contract events.
type SetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SetSession is an auto generated Go binding around an cpchain contract,
// with pre-set call and transact options.
type SetSession struct {
	Contract     *Set              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SetCallerSession is an auto generated read-only Go binding around an cpchain contract,
// with pre-set call options.
type SetCallerSession struct {
	Contract *SetCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SetTransactorSession is an auto generated write-only Go binding around an cpchain contract,
// with pre-set transact options.
type SetTransactorSession struct {
	Contract     *SetTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SetRaw is an auto generated low-level Go binding around an cpchain contract.
type SetRaw struct {
	Contract *Set // Generic contract binding to access the raw methods on
}

// SetCallerRaw is an auto generated low-level read-only Go binding around an cpchain contract.
type SetCallerRaw struct {
	Contract *SetCaller // Generic read-only contract binding to access the raw methods on
}

// SetTransactorRaw is an auto generated low-level write-only Go binding around an cpchain contract.
type SetTransactorRaw struct {
	Contract *SetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSet creates a new instance of Set, bound to a specific deployed contract.
func NewSet(address common.Address, backend bind.ContractBackend) (*Set, error) {
	contract, err := bindSet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Set{SetCaller: SetCaller{contract: contract}, SetTransactor: SetTransactor{contract: contract}, SetFilterer: SetFilterer{contract: contract}}, nil
}

// NewSetCaller creates a new read-only instance of Set, bound to a specific deployed contract.
func NewSetCaller(address common.Address, caller bind.ContractCaller) (*SetCaller, error) {
	contract, err := bindSet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SetCaller{contract: contract}, nil
}

// NewSetTransactor creates a new write-only instance of Set, bound to a specific deployed contract.
func NewSetTransactor(address common.Address, transactor bind.ContractTransactor) (*SetTransactor, error) {
	contract, err := bindSet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SetTransactor{contract: contract}, nil
}

// NewSetFilterer creates a new log filterer instance of Set, bound to a specific deployed contract.
func NewSetFilterer(address common.Address, filterer bind.ContractFilterer) (*SetFilterer, error) {
	contract, err := bindSet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SetFilterer{contract: contract}, nil
}

// bindSet binds a generic wrapper to an already deployed contract.
func bindSet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SetABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Set *SetRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Set.Contract.SetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Set *SetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Set.Contract.SetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Set *SetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Set.Contract.SetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Set *SetCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Set.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Set *SetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Set.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Set *SetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Set.Contract.contract.Transact(opts, method, params...)
}
