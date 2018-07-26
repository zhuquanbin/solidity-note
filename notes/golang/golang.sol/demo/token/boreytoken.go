// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package token

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

// BoreyTokenABI is the input ABI used to generate the binding from.
const BoreyTokenABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"supply\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// BoreyTokenBin is the compiled bytecode used for deploying new contracts.
const BoreyTokenBin = `0x608060405234801561001057600080fd5b50604051602080610287833981016040908152905160008054600160a060020a031916339081178255815260016020529190912055610233806100546000396000f3006080604052600436106100565763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166370a08231811461005b5780638da5cb5b1461009b578063a9059cbb146100d9575b600080fd5b34801561006757600080fd5b5061008973ffffffffffffffffffffffffffffffffffffffff6004351661011e565b60408051918252519081900360200190f35b3480156100a757600080fd5b506100b0610130565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b3480156100e557600080fd5b5061010a73ffffffffffffffffffffffffffffffffffffffff6004351660243561014c565b604080519115158252519081900360200190f35b60016020526000908152604090205481565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b600073ffffffffffffffffffffffffffffffffffffffff8316151561017057600080fd5b3360009081526001602052604090205482111561018c57600080fd5b3360008181526001602090815260408083208054879003905573ffffffffffffffffffffffffffffffffffffffff871680845292819020805487019055805186815290519293927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929181900390910190a3506001929150505600a165627a7a7230582012e426d59d7877eba3a48ed53cd039ff66d4399e3b87c8a8cf2c48628e827eda0029`

// DeployBoreyToken deploys a new Ethereum contract, binding an instance of BoreyToken to it.
func DeployBoreyToken(auth *bind.TransactOpts, backend bind.ContractBackend, supply *big.Int) (common.Address, *types.Transaction, *BoreyToken, error) {
	parsed, err := abi.JSON(strings.NewReader(BoreyTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BoreyTokenBin), backend, supply)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BoreyToken{BoreyTokenCaller: BoreyTokenCaller{contract: contract}, BoreyTokenTransactor: BoreyTokenTransactor{contract: contract}, BoreyTokenFilterer: BoreyTokenFilterer{contract: contract}}, nil
}

// BoreyToken is an auto generated Go binding around an Ethereum contract.
type BoreyToken struct {
	BoreyTokenCaller     // Read-only binding to the contract
	BoreyTokenTransactor // Write-only binding to the contract
	BoreyTokenFilterer   // Log filterer for contract events
}

// BoreyTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type BoreyTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BoreyTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BoreyTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BoreyTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BoreyTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BoreyTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BoreyTokenSession struct {
	Contract     *BoreyToken       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BoreyTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BoreyTokenCallerSession struct {
	Contract *BoreyTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BoreyTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BoreyTokenTransactorSession struct {
	Contract     *BoreyTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BoreyTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type BoreyTokenRaw struct {
	Contract *BoreyToken // Generic contract binding to access the raw methods on
}

// BoreyTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BoreyTokenCallerRaw struct {
	Contract *BoreyTokenCaller // Generic read-only contract binding to access the raw methods on
}

// BoreyTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BoreyTokenTransactorRaw struct {
	Contract *BoreyTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBoreyToken creates a new instance of BoreyToken, bound to a specific deployed contract.
func NewBoreyToken(address common.Address, backend bind.ContractBackend) (*BoreyToken, error) {
	contract, err := bindBoreyToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BoreyToken{BoreyTokenCaller: BoreyTokenCaller{contract: contract}, BoreyTokenTransactor: BoreyTokenTransactor{contract: contract}, BoreyTokenFilterer: BoreyTokenFilterer{contract: contract}}, nil
}

// NewBoreyTokenCaller creates a new read-only instance of BoreyToken, bound to a specific deployed contract.
func NewBoreyTokenCaller(address common.Address, caller bind.ContractCaller) (*BoreyTokenCaller, error) {
	contract, err := bindBoreyToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BoreyTokenCaller{contract: contract}, nil
}

// NewBoreyTokenTransactor creates a new write-only instance of BoreyToken, bound to a specific deployed contract.
func NewBoreyTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*BoreyTokenTransactor, error) {
	contract, err := bindBoreyToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BoreyTokenTransactor{contract: contract}, nil
}

// NewBoreyTokenFilterer creates a new log filterer instance of BoreyToken, bound to a specific deployed contract.
func NewBoreyTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*BoreyTokenFilterer, error) {
	contract, err := bindBoreyToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BoreyTokenFilterer{contract: contract}, nil
}

// bindBoreyToken binds a generic wrapper to an already deployed contract.
func bindBoreyToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BoreyTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BoreyToken *BoreyTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BoreyToken.Contract.BoreyTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BoreyToken *BoreyTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BoreyToken.Contract.BoreyTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BoreyToken *BoreyTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BoreyToken.Contract.BoreyTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BoreyToken *BoreyTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BoreyToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BoreyToken *BoreyTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BoreyToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BoreyToken *BoreyTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BoreyToken.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf( address) constant returns(uint256)
func (_BoreyToken *BoreyTokenCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BoreyToken.contract.Call(opts, out, "balanceOf", arg0)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf( address) constant returns(uint256)
func (_BoreyToken *BoreyTokenSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _BoreyToken.Contract.BalanceOf(&_BoreyToken.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf( address) constant returns(uint256)
func (_BoreyToken *BoreyTokenCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _BoreyToken.Contract.BalanceOf(&_BoreyToken.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BoreyToken *BoreyTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BoreyToken.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BoreyToken *BoreyTokenSession) Owner() (common.Address, error) {
	return _BoreyToken.Contract.Owner(&_BoreyToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BoreyToken *BoreyTokenCallerSession) Owner() (common.Address, error) {
	return _BoreyToken.Contract.Owner(&_BoreyToken.CallOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_BoreyToken *BoreyTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BoreyToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_BoreyToken *BoreyTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BoreyToken.Contract.Transfer(&_BoreyToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_BoreyToken *BoreyTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BoreyToken.Contract.Transfer(&_BoreyToken.TransactOpts, _to, _value)
}

// BoreyTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BoreyToken contract.
type BoreyTokenTransferIterator struct {
	Event *BoreyTokenTransfer // Event containing the contract specifics and raw log

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
func (it *BoreyTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BoreyTokenTransfer)
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
		it.Event = new(BoreyTokenTransfer)
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
func (it *BoreyTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BoreyTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BoreyTokenTransfer represents a Transfer event raised by the BoreyToken contract.
type BoreyTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _value uint256)
func (_BoreyToken *BoreyTokenFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*BoreyTokenTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _BoreyToken.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &BoreyTokenTransferIterator{contract: _BoreyToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _value uint256)
func (_BoreyToken *BoreyTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BoreyTokenTransfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _BoreyToken.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BoreyTokenTransfer)
				if err := _BoreyToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
