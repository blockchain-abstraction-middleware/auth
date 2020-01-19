// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Authorization

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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// AuthorizationABI is the input ABI used to generate the binding from.
const AuthorizationABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Subscribed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"hasWithdrawn\",\"type\":\"bool\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[],\"name\":\"subscribe\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"widthdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Authorization is an auto generated Go binding around an Ethereum contract.
type Authorization struct {
	AuthorizationCaller     // Read-only binding to the contract
	AuthorizationTransactor // Write-only binding to the contract
	AuthorizationFilterer   // Log filterer for contract events
}

// AuthorizationCaller is an auto generated read-only Go binding around an Ethereum contract.
type AuthorizationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuthorizationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AuthorizationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuthorizationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AuthorizationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuthorizationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AuthorizationSession struct {
	Contract     *Authorization    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AuthorizationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AuthorizationCallerSession struct {
	Contract *AuthorizationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// AuthorizationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AuthorizationTransactorSession struct {
	Contract     *AuthorizationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// AuthorizationRaw is an auto generated low-level Go binding around an Ethereum contract.
type AuthorizationRaw struct {
	Contract *Authorization // Generic contract binding to access the raw methods on
}

// AuthorizationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AuthorizationCallerRaw struct {
	Contract *AuthorizationCaller // Generic read-only contract binding to access the raw methods on
}

// AuthorizationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AuthorizationTransactorRaw struct {
	Contract *AuthorizationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAuthorization creates a new instance of Authorization, bound to a specific deployed contract.
func NewAuthorization(address common.Address, backend bind.ContractBackend) (*Authorization, error) {
	contract, err := bindAuthorization(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Authorization{AuthorizationCaller: AuthorizationCaller{contract: contract}, AuthorizationTransactor: AuthorizationTransactor{contract: contract}, AuthorizationFilterer: AuthorizationFilterer{contract: contract}}, nil
}

// NewAuthorizationCaller creates a new read-only instance of Authorization, bound to a specific deployed contract.
func NewAuthorizationCaller(address common.Address, caller bind.ContractCaller) (*AuthorizationCaller, error) {
	contract, err := bindAuthorization(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AuthorizationCaller{contract: contract}, nil
}

// NewAuthorizationTransactor creates a new write-only instance of Authorization, bound to a specific deployed contract.
func NewAuthorizationTransactor(address common.Address, transactor bind.ContractTransactor) (*AuthorizationTransactor, error) {
	contract, err := bindAuthorization(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AuthorizationTransactor{contract: contract}, nil
}

// NewAuthorizationFilterer creates a new log filterer instance of Authorization, bound to a specific deployed contract.
func NewAuthorizationFilterer(address common.Address, filterer bind.ContractFilterer) (*AuthorizationFilterer, error) {
	contract, err := bindAuthorization(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AuthorizationFilterer{contract: contract}, nil
}

// bindAuthorization binds a generic wrapper to an already deployed contract.
func bindAuthorization(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AuthorizationABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Authorization *AuthorizationRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Authorization.Contract.AuthorizationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Authorization *AuthorizationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Authorization.Contract.AuthorizationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Authorization *AuthorizationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Authorization.Contract.AuthorizationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Authorization *AuthorizationCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Authorization.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Authorization *AuthorizationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Authorization.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Authorization *AuthorizationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Authorization.Contract.contract.Transact(opts, method, params...)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Authorization *AuthorizationCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Authorization.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Authorization *AuthorizationSession) IsOwner() (bool, error) {
	return _Authorization.Contract.IsOwner(&_Authorization.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Authorization *AuthorizationCallerSession) IsOwner() (bool, error) {
	return _Authorization.Contract.IsOwner(&_Authorization.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Authorization *AuthorizationCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Authorization.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Authorization *AuthorizationSession) Owner() (common.Address, error) {
	return _Authorization.Contract.Owner(&_Authorization.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Authorization *AuthorizationCallerSession) Owner() (common.Address, error) {
	return _Authorization.Contract.Owner(&_Authorization.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Authorization *AuthorizationTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Authorization.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Authorization *AuthorizationSession) RenounceOwnership() (*types.Transaction, error) {
	return _Authorization.Contract.RenounceOwnership(&_Authorization.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Authorization *AuthorizationTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Authorization.Contract.RenounceOwnership(&_Authorization.TransactOpts)
}

// Subscribe is a paid mutator transaction binding the contract method 0x8f449a05.
//
// Solidity: function subscribe() returns()
func (_Authorization *AuthorizationTransactor) Subscribe(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Authorization.contract.Transact(opts, "subscribe")
}

// Subscribe is a paid mutator transaction binding the contract method 0x8f449a05.
//
// Solidity: function subscribe() returns()
func (_Authorization *AuthorizationSession) Subscribe() (*types.Transaction, error) {
	return _Authorization.Contract.Subscribe(&_Authorization.TransactOpts)
}

// Subscribe is a paid mutator transaction binding the contract method 0x8f449a05.
//
// Solidity: function subscribe() returns()
func (_Authorization *AuthorizationTransactorSession) Subscribe() (*types.Transaction, error) {
	return _Authorization.Contract.Subscribe(&_Authorization.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Authorization *AuthorizationTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Authorization.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Authorization *AuthorizationSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Authorization.Contract.TransferOwnership(&_Authorization.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Authorization *AuthorizationTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Authorization.Contract.TransferOwnership(&_Authorization.TransactOpts, newOwner)
}

// Widthdraw is a paid mutator transaction binding the contract method 0x52b50a2a.
//
// Solidity: function widthdraw() returns()
func (_Authorization *AuthorizationTransactor) Widthdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Authorization.contract.Transact(opts, "widthdraw")
}

// Widthdraw is a paid mutator transaction binding the contract method 0x52b50a2a.
//
// Solidity: function widthdraw() returns()
func (_Authorization *AuthorizationSession) Widthdraw() (*types.Transaction, error) {
	return _Authorization.Contract.Widthdraw(&_Authorization.TransactOpts)
}

// Widthdraw is a paid mutator transaction binding the contract method 0x52b50a2a.
//
// Solidity: function widthdraw() returns()
func (_Authorization *AuthorizationTransactorSession) Widthdraw() (*types.Transaction, error) {
	return _Authorization.Contract.Widthdraw(&_Authorization.TransactOpts)
}

// AuthorizationOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Authorization contract.
type AuthorizationOwnershipTransferredIterator struct {
	Event *AuthorizationOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AuthorizationOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuthorizationOwnershipTransferred)
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
		it.Event = new(AuthorizationOwnershipTransferred)
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
func (it *AuthorizationOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuthorizationOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuthorizationOwnershipTransferred represents a OwnershipTransferred event raised by the Authorization contract.
type AuthorizationOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Authorization *AuthorizationFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AuthorizationOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Authorization.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AuthorizationOwnershipTransferredIterator{contract: _Authorization.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Authorization *AuthorizationFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AuthorizationOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Authorization.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuthorizationOwnershipTransferred)
				if err := _Authorization.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Authorization *AuthorizationFilterer) ParseOwnershipTransferred(log types.Log) (*AuthorizationOwnershipTransferred, error) {
	event := new(AuthorizationOwnershipTransferred)
	if err := _Authorization.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AuthorizationSubscribedIterator is returned from FilterSubscribed and is used to iterate over the raw logs and unpacked data for Subscribed events raised by the Authorization contract.
type AuthorizationSubscribedIterator struct {
	Event *AuthorizationSubscribed // Event containing the contract specifics and raw log

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
func (it *AuthorizationSubscribedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuthorizationSubscribed)
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
		it.Event = new(AuthorizationSubscribed)
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
func (it *AuthorizationSubscribedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuthorizationSubscribedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuthorizationSubscribed represents a Subscribed event raised by the Authorization contract.
type AuthorizationSubscribed struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSubscribed is a free log retrieval operation binding the contract event 0xa88fac53823b534c72d6958345adbe6801e072750d83b490d52ebbac51473a63.
//
// Solidity: event Subscribed(address indexed account)
func (_Authorization *AuthorizationFilterer) FilterSubscribed(opts *bind.FilterOpts, account []common.Address) (*AuthorizationSubscribedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Authorization.contract.FilterLogs(opts, "Subscribed", accountRule)
	if err != nil {
		return nil, err
	}
	return &AuthorizationSubscribedIterator{contract: _Authorization.contract, event: "Subscribed", logs: logs, sub: sub}, nil
}

// WatchSubscribed is a free log subscription operation binding the contract event 0xa88fac53823b534c72d6958345adbe6801e072750d83b490d52ebbac51473a63.
//
// Solidity: event Subscribed(address indexed account)
func (_Authorization *AuthorizationFilterer) WatchSubscribed(opts *bind.WatchOpts, sink chan<- *AuthorizationSubscribed, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Authorization.contract.WatchLogs(opts, "Subscribed", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuthorizationSubscribed)
				if err := _Authorization.contract.UnpackLog(event, "Subscribed", log); err != nil {
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

// ParseSubscribed is a log parse operation binding the contract event 0xa88fac53823b534c72d6958345adbe6801e072750d83b490d52ebbac51473a63.
//
// Solidity: event Subscribed(address indexed account)
func (_Authorization *AuthorizationFilterer) ParseSubscribed(log types.Log) (*AuthorizationSubscribed, error) {
	event := new(AuthorizationSubscribed)
	if err := _Authorization.contract.UnpackLog(event, "Subscribed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AuthorizationWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Authorization contract.
type AuthorizationWithdrawIterator struct {
	Event *AuthorizationWithdraw // Event containing the contract specifics and raw log

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
func (it *AuthorizationWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuthorizationWithdraw)
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
		it.Event = new(AuthorizationWithdraw)
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
func (it *AuthorizationWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuthorizationWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuthorizationWithdraw represents a Withdraw event raised by the Authorization contract.
type AuthorizationWithdraw struct {
	HasWithdrawn bool
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xfe589245338ce477b4fb3c95d7fef59c7ae5fb37c54522f0e5e8922e56039247.
//
// Solidity: event Withdraw(bool hasWithdrawn)
func (_Authorization *AuthorizationFilterer) FilterWithdraw(opts *bind.FilterOpts) (*AuthorizationWithdrawIterator, error) {

	logs, sub, err := _Authorization.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &AuthorizationWithdrawIterator{contract: _Authorization.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xfe589245338ce477b4fb3c95d7fef59c7ae5fb37c54522f0e5e8922e56039247.
//
// Solidity: event Withdraw(bool hasWithdrawn)
func (_Authorization *AuthorizationFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *AuthorizationWithdraw) (event.Subscription, error) {

	logs, sub, err := _Authorization.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuthorizationWithdraw)
				if err := _Authorization.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xfe589245338ce477b4fb3c95d7fef59c7ae5fb37c54522f0e5e8922e56039247.
//
// Solidity: event Withdraw(bool hasWithdrawn)
func (_Authorization *AuthorizationFilterer) ParseWithdraw(log types.Log) (*AuthorizationWithdraw, error) {
	event := new(AuthorizationWithdraw)
	if err := _Authorization.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	return event, nil
}
