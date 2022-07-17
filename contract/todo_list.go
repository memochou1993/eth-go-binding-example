// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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
)

// TodoListTask is an auto generated low-level Go binding around an user-defined struct.
type TodoListTask struct {
	Idx       *big.Int
	Content   string
	Completed bool
}

// TodoListMetaData contains all meta data concerning the TodoList contract.
var TodoListMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"completed\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structTodoList.Task\",\"name\":\"task\",\"type\":\"tuple\"}],\"name\":\"TaskCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"completed\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structTodoList.Task\",\"name\":\"task\",\"type\":\"tuple\"}],\"name\":\"TaskUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"taskCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tasks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"completed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_content\",\"type\":\"string\"}],\"name\":\"createTask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_idx\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_completed\",\"type\":\"bool\"}],\"name\":\"updateTask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TodoListABI is the input ABI used to generate the binding from.
// Deprecated: Use TodoListMetaData.ABI instead.
var TodoListABI = TodoListMetaData.ABI

// TodoList is an auto generated Go binding around an Ethereum contract.
type TodoList struct {
	TodoListCaller     // Read-only binding to the contract
	TodoListTransactor // Write-only binding to the contract
	TodoListFilterer   // Log filterer for contract events
}

// TodoListCaller is an auto generated read-only Go binding around an Ethereum contract.
type TodoListCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TodoListTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TodoListTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TodoListFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TodoListFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TodoListSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TodoListSession struct {
	Contract     *TodoList         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TodoListCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TodoListCallerSession struct {
	Contract *TodoListCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// TodoListTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TodoListTransactorSession struct {
	Contract     *TodoListTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// TodoListRaw is an auto generated low-level Go binding around an Ethereum contract.
type TodoListRaw struct {
	Contract *TodoList // Generic contract binding to access the raw methods on
}

// TodoListCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TodoListCallerRaw struct {
	Contract *TodoListCaller // Generic read-only contract binding to access the raw methods on
}

// TodoListTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TodoListTransactorRaw struct {
	Contract *TodoListTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTodoList creates a new instance of TodoList, bound to a specific deployed contract.
func NewTodoList(address common.Address, backend bind.ContractBackend) (*TodoList, error) {
	contract, err := bindTodoList(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TodoList{TodoListCaller: TodoListCaller{contract: contract}, TodoListTransactor: TodoListTransactor{contract: contract}, TodoListFilterer: TodoListFilterer{contract: contract}}, nil
}

// NewTodoListCaller creates a new read-only instance of TodoList, bound to a specific deployed contract.
func NewTodoListCaller(address common.Address, caller bind.ContractCaller) (*TodoListCaller, error) {
	contract, err := bindTodoList(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TodoListCaller{contract: contract}, nil
}

// NewTodoListTransactor creates a new write-only instance of TodoList, bound to a specific deployed contract.
func NewTodoListTransactor(address common.Address, transactor bind.ContractTransactor) (*TodoListTransactor, error) {
	contract, err := bindTodoList(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TodoListTransactor{contract: contract}, nil
}

// NewTodoListFilterer creates a new log filterer instance of TodoList, bound to a specific deployed contract.
func NewTodoListFilterer(address common.Address, filterer bind.ContractFilterer) (*TodoListFilterer, error) {
	contract, err := bindTodoList(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TodoListFilterer{contract: contract}, nil
}

// bindTodoList binds a generic wrapper to an already deployed contract.
func bindTodoList(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TodoListABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TodoList *TodoListRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TodoList.Contract.TodoListCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TodoList *TodoListRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TodoList.Contract.TodoListTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TodoList *TodoListRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TodoList.Contract.TodoListTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TodoList *TodoListCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TodoList.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TodoList *TodoListTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TodoList.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TodoList *TodoListTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TodoList.Contract.contract.Transact(opts, method, params...)
}

// TaskCount is a free data retrieval call binding the contract method 0xb6cb58a5.
//
// Solidity: function taskCount() view returns(uint256)
func (_TodoList *TodoListCaller) TaskCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TodoList.contract.Call(opts, &out, "taskCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TaskCount is a free data retrieval call binding the contract method 0xb6cb58a5.
//
// Solidity: function taskCount() view returns(uint256)
func (_TodoList *TodoListSession) TaskCount() (*big.Int, error) {
	return _TodoList.Contract.TaskCount(&_TodoList.CallOpts)
}

// TaskCount is a free data retrieval call binding the contract method 0xb6cb58a5.
//
// Solidity: function taskCount() view returns(uint256)
func (_TodoList *TodoListCallerSession) TaskCount() (*big.Int, error) {
	return _TodoList.Contract.TaskCount(&_TodoList.CallOpts)
}

// Tasks is a free data retrieval call binding the contract method 0x8d977672.
//
// Solidity: function tasks(uint256 ) view returns(uint256 idx, string content, bool completed)
func (_TodoList *TodoListCaller) Tasks(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Idx       *big.Int
	Content   string
	Completed bool
}, error) {
	var out []interface{}
	err := _TodoList.contract.Call(opts, &out, "tasks", arg0)

	outstruct := new(struct {
		Idx       *big.Int
		Content   string
		Completed bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Idx = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Content = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Completed = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// Tasks is a free data retrieval call binding the contract method 0x8d977672.
//
// Solidity: function tasks(uint256 ) view returns(uint256 idx, string content, bool completed)
func (_TodoList *TodoListSession) Tasks(arg0 *big.Int) (struct {
	Idx       *big.Int
	Content   string
	Completed bool
}, error) {
	return _TodoList.Contract.Tasks(&_TodoList.CallOpts, arg0)
}

// Tasks is a free data retrieval call binding the contract method 0x8d977672.
//
// Solidity: function tasks(uint256 ) view returns(uint256 idx, string content, bool completed)
func (_TodoList *TodoListCallerSession) Tasks(arg0 *big.Int) (struct {
	Idx       *big.Int
	Content   string
	Completed bool
}, error) {
	return _TodoList.Contract.Tasks(&_TodoList.CallOpts, arg0)
}

// CreateTask is a paid mutator transaction binding the contract method 0x111002aa.
//
// Solidity: function createTask(string _content) returns()
func (_TodoList *TodoListTransactor) CreateTask(opts *bind.TransactOpts, _content string) (*types.Transaction, error) {
	return _TodoList.contract.Transact(opts, "createTask", _content)
}

// CreateTask is a paid mutator transaction binding the contract method 0x111002aa.
//
// Solidity: function createTask(string _content) returns()
func (_TodoList *TodoListSession) CreateTask(_content string) (*types.Transaction, error) {
	return _TodoList.Contract.CreateTask(&_TodoList.TransactOpts, _content)
}

// CreateTask is a paid mutator transaction binding the contract method 0x111002aa.
//
// Solidity: function createTask(string _content) returns()
func (_TodoList *TodoListTransactorSession) CreateTask(_content string) (*types.Transaction, error) {
	return _TodoList.Contract.CreateTask(&_TodoList.TransactOpts, _content)
}

// UpdateTask is a paid mutator transaction binding the contract method 0x5e1bf4a1.
//
// Solidity: function updateTask(uint256 _idx, bool _completed) returns()
func (_TodoList *TodoListTransactor) UpdateTask(opts *bind.TransactOpts, _idx *big.Int, _completed bool) (*types.Transaction, error) {
	return _TodoList.contract.Transact(opts, "updateTask", _idx, _completed)
}

// UpdateTask is a paid mutator transaction binding the contract method 0x5e1bf4a1.
//
// Solidity: function updateTask(uint256 _idx, bool _completed) returns()
func (_TodoList *TodoListSession) UpdateTask(_idx *big.Int, _completed bool) (*types.Transaction, error) {
	return _TodoList.Contract.UpdateTask(&_TodoList.TransactOpts, _idx, _completed)
}

// UpdateTask is a paid mutator transaction binding the contract method 0x5e1bf4a1.
//
// Solidity: function updateTask(uint256 _idx, bool _completed) returns()
func (_TodoList *TodoListTransactorSession) UpdateTask(_idx *big.Int, _completed bool) (*types.Transaction, error) {
	return _TodoList.Contract.UpdateTask(&_TodoList.TransactOpts, _idx, _completed)
}

// TodoListTaskCreatedIterator is returned from FilterTaskCreated and is used to iterate over the raw logs and unpacked data for TaskCreated events raised by the TodoList contract.
type TodoListTaskCreatedIterator struct {
	Event *TodoListTaskCreated // Event containing the contract specifics and raw log

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
func (it *TodoListTaskCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TodoListTaskCreated)
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
		it.Event = new(TodoListTaskCreated)
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
func (it *TodoListTaskCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TodoListTaskCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TodoListTaskCreated represents a TaskCreated event raised by the TodoList contract.
type TodoListTaskCreated struct {
	Idx  *big.Int
	Task TodoListTask
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterTaskCreated is a free log retrieval operation binding the contract event 0xfe0b441f8ad823c57d031b1c59ca16726562a2569e9dd158a5e5dd86f1f48227.
//
// Solidity: event TaskCreated(uint256 idx, (uint256,string,bool) task)
func (_TodoList *TodoListFilterer) FilterTaskCreated(opts *bind.FilterOpts) (*TodoListTaskCreatedIterator, error) {

	logs, sub, err := _TodoList.contract.FilterLogs(opts, "TaskCreated")
	if err != nil {
		return nil, err
	}
	return &TodoListTaskCreatedIterator{contract: _TodoList.contract, event: "TaskCreated", logs: logs, sub: sub}, nil
}

// WatchTaskCreated is a free log subscription operation binding the contract event 0xfe0b441f8ad823c57d031b1c59ca16726562a2569e9dd158a5e5dd86f1f48227.
//
// Solidity: event TaskCreated(uint256 idx, (uint256,string,bool) task)
func (_TodoList *TodoListFilterer) WatchTaskCreated(opts *bind.WatchOpts, sink chan<- *TodoListTaskCreated) (event.Subscription, error) {

	logs, sub, err := _TodoList.contract.WatchLogs(opts, "TaskCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TodoListTaskCreated)
				if err := _TodoList.contract.UnpackLog(event, "TaskCreated", log); err != nil {
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

// ParseTaskCreated is a log parse operation binding the contract event 0xfe0b441f8ad823c57d031b1c59ca16726562a2569e9dd158a5e5dd86f1f48227.
//
// Solidity: event TaskCreated(uint256 idx, (uint256,string,bool) task)
func (_TodoList *TodoListFilterer) ParseTaskCreated(log types.Log) (*TodoListTaskCreated, error) {
	event := new(TodoListTaskCreated)
	if err := _TodoList.contract.UnpackLog(event, "TaskCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TodoListTaskUpdatedIterator is returned from FilterTaskUpdated and is used to iterate over the raw logs and unpacked data for TaskUpdated events raised by the TodoList contract.
type TodoListTaskUpdatedIterator struct {
	Event *TodoListTaskUpdated // Event containing the contract specifics and raw log

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
func (it *TodoListTaskUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TodoListTaskUpdated)
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
		it.Event = new(TodoListTaskUpdated)
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
func (it *TodoListTaskUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TodoListTaskUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TodoListTaskUpdated represents a TaskUpdated event raised by the TodoList contract.
type TodoListTaskUpdated struct {
	Idx  *big.Int
	Task TodoListTask
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterTaskUpdated is a free log retrieval operation binding the contract event 0x3e4c9db1f679bbc95759eb695b18cd4a4f1c66dc46015072290fe692ff768c8d.
//
// Solidity: event TaskUpdated(uint256 idx, (uint256,string,bool) task)
func (_TodoList *TodoListFilterer) FilterTaskUpdated(opts *bind.FilterOpts) (*TodoListTaskUpdatedIterator, error) {

	logs, sub, err := _TodoList.contract.FilterLogs(opts, "TaskUpdated")
	if err != nil {
		return nil, err
	}
	return &TodoListTaskUpdatedIterator{contract: _TodoList.contract, event: "TaskUpdated", logs: logs, sub: sub}, nil
}

// WatchTaskUpdated is a free log subscription operation binding the contract event 0x3e4c9db1f679bbc95759eb695b18cd4a4f1c66dc46015072290fe692ff768c8d.
//
// Solidity: event TaskUpdated(uint256 idx, (uint256,string,bool) task)
func (_TodoList *TodoListFilterer) WatchTaskUpdated(opts *bind.WatchOpts, sink chan<- *TodoListTaskUpdated) (event.Subscription, error) {

	logs, sub, err := _TodoList.contract.WatchLogs(opts, "TaskUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TodoListTaskUpdated)
				if err := _TodoList.contract.UnpackLog(event, "TaskUpdated", log); err != nil {
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

// ParseTaskUpdated is a log parse operation binding the contract event 0x3e4c9db1f679bbc95759eb695b18cd4a4f1c66dc46015072290fe692ff768c8d.
//
// Solidity: event TaskUpdated(uint256 idx, (uint256,string,bool) task)
func (_TodoList *TodoListFilterer) ParseTaskUpdated(log types.Log) (*TodoListTaskUpdated, error) {
	event := new(TodoListTaskUpdated)
	if err := _TodoList.contract.UnpackLog(event, "TaskUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
