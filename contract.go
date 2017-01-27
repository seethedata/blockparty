// This file is an automatically generated Go binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package main

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// ContractABI is the input ABI used to generate the binding from.
const ContractABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"sold\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBidIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"notForSale\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setAppraisalValue\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"askingPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isForSale\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"appraised\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"forSale\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"bidNumber\",\"type\":\"uint256\"}],\"name\":\"getBid\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"changedMortgageValue\",\"type\":\"uint256\"}],\"name\":\"changeMortgageValue\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"bids\",\"outputs\":[{\"name\":\"bidder\",\"type\":\"address\"},{\"name\":\"bidValue\",\"type\":\"uint256\"},{\"name\":\"accepted\",\"type\":\"bool\"},{\"name\":\"bidNumber\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"inspected\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"streetAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"appraisalValue\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"biddingPrice\",\"type\":\"uint256\"}],\"name\":\"checkBid\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"bidderAddress\",\"type\":\"address\"}],\"name\":\"rejectBid\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bidIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"houseOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"bidder\",\"type\":\"address\"}],\"name\":\"acceptBid\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"buyer\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bidValue\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"status\",\"type\":\"uint256\"}],\"name\":\"inspectionStatus\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"biddingPrice\",\"type\":\"uint256\"}],\"name\":\"placeBid\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newPrice\",\"type\":\"uint256\"}],\"name\":\"changePrice\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"rejectMortgage\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"state\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"requestedMortgage\",\"type\":\"uint256\"}],\"name\":\"applyMortgage\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"acceptedBid\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deleteAllBids\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bidder\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"soldPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"stAddress\",\"type\":\"bytes32\"}],\"payable\":false,\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddress\",\"type\":\"address\"}],\"name\":\"InitHouse\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddress\",\"type\":\"address\"}],\"name\":\"HouseInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"askingPrice\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"stAddress\",\"type\":\"bytes32\"}],\"name\":\"HouseForSale\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"stAddress\",\"type\":\"bytes32\"}],\"name\":\"HouseNotForSale\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"askingPrice\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"stAddress\",\"type\":\"bytes32\"}],\"name\":\"PriceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"biddingPrice\",\"type\":\"uint256\"}],\"name\":\"BidPlaced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"bidder\",\"type\":\"address\"}],\"name\":\"BidRejected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"acceptedBid\",\"type\":\"uint256\"}],\"name\":\"BidAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"biddingPrice\",\"type\":\"uint256\"}],\"name\":\"BidExists\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"inspector\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"status\",\"type\":\"uint256\"}],\"name\":\"Inspected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"lender\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"MortgageApplied\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"lender\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"MortgageApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"lender\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"MortgageRejected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"HouseSold\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"mortgageValue\",\"type\":\"uint256\"}],\"name\":\"ChangedMortgageValue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"biddingPrice\",\"type\":\"uint256\"}],\"name\":\"BidAlreadyExists\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"bidNumber\",\"type\":\"uint256\"}],\"name\":\"BidAlreadyAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"acceptedBid\",\"type\":\"uint256\"}],\"name\":\"NoBidsAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"bidNumber\",\"type\":\"uint256\"}],\"name\":\"FoundBid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"bidNumber\",\"type\":\"uint256\"}],\"name\":\"UpdatedBid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"AllBidsDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"appraiser\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"AppraisalValueSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"bidNumber\",\"type\":\"uint256\"}],\"name\":\"DeleteBids\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"bidNumber\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"acceptedBid\",\"type\":\"uint256\"}],\"name\":\"BidsDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"}],\"name\":\"Killed\",\"type\":\"event\"}]"

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// AcceptedBid is a free data retrieval call binding the contract method 0xd9fa1c5d.
//
// Solidity: function acceptedBid() constant returns(uint256)
func (_Contract *ContractCaller) AcceptedBid(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "acceptedBid")
	return *ret0, err
}

// AcceptedBid is a free data retrieval call binding the contract method 0xd9fa1c5d.
//
// Solidity: function acceptedBid() constant returns(uint256)
func (_Contract *ContractSession) AcceptedBid() (*big.Int, error) {
	return _Contract.Contract.AcceptedBid(&_Contract.CallOpts)
}

// AcceptedBid is a free data retrieval call binding the contract method 0xd9fa1c5d.
//
// Solidity: function acceptedBid() constant returns(uint256)
func (_Contract *ContractCallerSession) AcceptedBid() (*big.Int, error) {
	return _Contract.Contract.AcceptedBid(&_Contract.CallOpts)
}

// AppraisalValue is a free data retrieval call binding the contract method 0x484d6df7.
//
// Solidity: function appraisalValue() constant returns(uint256)
func (_Contract *ContractCaller) AppraisalValue(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "appraisalValue")
	return *ret0, err
}

// AppraisalValue is a free data retrieval call binding the contract method 0x484d6df7.
//
// Solidity: function appraisalValue() constant returns(uint256)
func (_Contract *ContractSession) AppraisalValue() (*big.Int, error) {
	return _Contract.Contract.AppraisalValue(&_Contract.CallOpts)
}

// AppraisalValue is a free data retrieval call binding the contract method 0x484d6df7.
//
// Solidity: function appraisalValue() constant returns(uint256)
func (_Contract *ContractCallerSession) AppraisalValue() (*big.Int, error) {
	return _Contract.Contract.AppraisalValue(&_Contract.CallOpts)
}

// Appraised is a free data retrieval call binding the contract method 0x21294bf6.
//
// Solidity: function appraised() constant returns(bool)
func (_Contract *ContractCaller) Appraised(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "appraised")
	return *ret0, err
}

// Appraised is a free data retrieval call binding the contract method 0x21294bf6.
//
// Solidity: function appraised() constant returns(bool)
func (_Contract *ContractSession) Appraised() (bool, error) {
	return _Contract.Contract.Appraised(&_Contract.CallOpts)
}

// Appraised is a free data retrieval call binding the contract method 0x21294bf6.
//
// Solidity: function appraised() constant returns(bool)
func (_Contract *ContractCallerSession) Appraised() (bool, error) {
	return _Contract.Contract.Appraised(&_Contract.CallOpts)
}

// AskingPrice is a free data retrieval call binding the contract method 0x1a7ac726.
//
// Solidity: function askingPrice() constant returns(uint256)
func (_Contract *ContractCaller) AskingPrice(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "askingPrice")
	return *ret0, err
}

// AskingPrice is a free data retrieval call binding the contract method 0x1a7ac726.
//
// Solidity: function askingPrice() constant returns(uint256)
func (_Contract *ContractSession) AskingPrice() (*big.Int, error) {
	return _Contract.Contract.AskingPrice(&_Contract.CallOpts)
}

// AskingPrice is a free data retrieval call binding the contract method 0x1a7ac726.
//
// Solidity: function askingPrice() constant returns(uint256)
func (_Contract *ContractCallerSession) AskingPrice() (*big.Int, error) {
	return _Contract.Contract.AskingPrice(&_Contract.CallOpts)
}

// BidIndex is a free data retrieval call binding the contract method 0x61f04c10.
//
// Solidity: function bidIndex() constant returns(uint256)
func (_Contract *ContractCaller) BidIndex(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "bidIndex")
	return *ret0, err
}

// BidIndex is a free data retrieval call binding the contract method 0x61f04c10.
//
// Solidity: function bidIndex() constant returns(uint256)
func (_Contract *ContractSession) BidIndex() (*big.Int, error) {
	return _Contract.Contract.BidIndex(&_Contract.CallOpts)
}

// BidIndex is a free data retrieval call binding the contract method 0x61f04c10.
//
// Solidity: function bidIndex() constant returns(uint256)
func (_Contract *ContractCallerSession) BidIndex() (*big.Int, error) {
	return _Contract.Contract.BidIndex(&_Contract.CallOpts)
}

// BidValue is a free data retrieval call binding the contract method 0x72e92e22.
//
// Solidity: function bidValue() constant returns(uint256)
func (_Contract *ContractCaller) BidValue(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "bidValue")
	return *ret0, err
}

// BidValue is a free data retrieval call binding the contract method 0x72e92e22.
//
// Solidity: function bidValue() constant returns(uint256)
func (_Contract *ContractSession) BidValue() (*big.Int, error) {
	return _Contract.Contract.BidValue(&_Contract.CallOpts)
}

// BidValue is a free data retrieval call binding the contract method 0x72e92e22.
//
// Solidity: function bidValue() constant returns(uint256)
func (_Contract *ContractCallerSession) BidValue() (*big.Int, error) {
	return _Contract.Contract.BidValue(&_Contract.CallOpts)
}

// Bidder is a free data retrieval call binding the contract method 0xf496d882.
//
// Solidity: function bidder() constant returns(address)
func (_Contract *ContractCaller) Bidder(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "bidder")
	return *ret0, err
}

// Bidder is a free data retrieval call binding the contract method 0xf496d882.
//
// Solidity: function bidder() constant returns(address)
func (_Contract *ContractSession) Bidder() (common.Address, error) {
	return _Contract.Contract.Bidder(&_Contract.CallOpts)
}

// Bidder is a free data retrieval call binding the contract method 0xf496d882.
//
// Solidity: function bidder() constant returns(address)
func (_Contract *ContractCallerSession) Bidder() (common.Address, error) {
	return _Contract.Contract.Bidder(&_Contract.CallOpts)
}

// Bids is a free data retrieval call binding the contract method 0x4423c5f1.
//
// Solidity: function bids( uint256) constant returns(bidder address, bidValue uint256, accepted bool, bidNumber uint256)
func (_Contract *ContractCaller) Bids(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Bidder    common.Address
	BidValue  *big.Int
	Accepted  bool
	BidNumber *big.Int
}, error) {
	ret := new(struct {
		Bidder    common.Address
		BidValue  *big.Int
		Accepted  bool
		BidNumber *big.Int
	})
	out := ret
	err := _Contract.contract.Call(opts, out, "bids", arg0)
	return *ret, err
}

// Bids is a free data retrieval call binding the contract method 0x4423c5f1.
//
// Solidity: function bids( uint256) constant returns(bidder address, bidValue uint256, accepted bool, bidNumber uint256)
func (_Contract *ContractSession) Bids(arg0 *big.Int) (struct {
	Bidder    common.Address
	BidValue  *big.Int
	Accepted  bool
	BidNumber *big.Int
}, error) {
	return _Contract.Contract.Bids(&_Contract.CallOpts, arg0)
}

// Bids is a free data retrieval call binding the contract method 0x4423c5f1.
//
// Solidity: function bids( uint256) constant returns(bidder address, bidValue uint256, accepted bool, bidNumber uint256)
func (_Contract *ContractCallerSession) Bids(arg0 *big.Int) (struct {
	Bidder    common.Address
	BidValue  *big.Int
	Accepted  bool
	BidNumber *big.Int
}, error) {
	return _Contract.Contract.Bids(&_Contract.CallOpts, arg0)
}

// Buyer is a free data retrieval call binding the contract method 0x7150d8ae.
//
// Solidity: function buyer() constant returns(address)
func (_Contract *ContractCaller) Buyer(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "buyer")
	return *ret0, err
}

// Buyer is a free data retrieval call binding the contract method 0x7150d8ae.
//
// Solidity: function buyer() constant returns(address)
func (_Contract *ContractSession) Buyer() (common.Address, error) {
	return _Contract.Contract.Buyer(&_Contract.CallOpts)
}

// Buyer is a free data retrieval call binding the contract method 0x7150d8ae.
//
// Solidity: function buyer() constant returns(address)
func (_Contract *ContractCallerSession) Buyer() (common.Address, error) {
	return _Contract.Contract.Buyer(&_Contract.CallOpts)
}

// CheckBid is a free data retrieval call binding the contract method 0x58ed1aae.
//
// Solidity: function checkBid(biddingPrice uint256) constant returns(bool)
func (_Contract *ContractCaller) CheckBid(opts *bind.CallOpts, biddingPrice *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "checkBid", biddingPrice)
	return *ret0, err
}

// CheckBid is a free data retrieval call binding the contract method 0x58ed1aae.
//
// Solidity: function checkBid(biddingPrice uint256) constant returns(bool)
func (_Contract *ContractSession) CheckBid(biddingPrice *big.Int) (bool, error) {
	return _Contract.Contract.CheckBid(&_Contract.CallOpts, biddingPrice)
}

// CheckBid is a free data retrieval call binding the contract method 0x58ed1aae.
//
// Solidity: function checkBid(biddingPrice uint256) constant returns(bool)
func (_Contract *ContractCallerSession) CheckBid(biddingPrice *big.Int) (bool, error) {
	return _Contract.Contract.CheckBid(&_Contract.CallOpts, biddingPrice)
}

// GetBid is a free data retrieval call binding the contract method 0x3c889e6f.
//
// Solidity: function getBid(bidNumber uint256) constant returns(address, uint256, bool, uint256)
func (_Contract *ContractCaller) GetBid(opts *bind.CallOpts, bidNumber *big.Int) (common.Address, *big.Int, bool, *big.Int, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(*big.Int)
		ret2 = new(bool)
		ret3 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
	}
	err := _Contract.contract.Call(opts, out, "getBid", bidNumber)
	return *ret0, *ret1, *ret2, *ret3, err
}

// GetBid is a free data retrieval call binding the contract method 0x3c889e6f.
//
// Solidity: function getBid(bidNumber uint256) constant returns(address, uint256, bool, uint256)
func (_Contract *ContractSession) GetBid(bidNumber *big.Int) (common.Address, *big.Int, bool, *big.Int, error) {
	return _Contract.Contract.GetBid(&_Contract.CallOpts, bidNumber)
}

// GetBid is a free data retrieval call binding the contract method 0x3c889e6f.
//
// Solidity: function getBid(bidNumber uint256) constant returns(address, uint256, bool, uint256)
func (_Contract *ContractCallerSession) GetBid(bidNumber *big.Int) (common.Address, *big.Int, bool, *big.Int, error) {
	return _Contract.Contract.GetBid(&_Contract.CallOpts, bidNumber)
}

// GetBidIndex is a free data retrieval call binding the contract method 0x05a13a68.
//
// Solidity: function getBidIndex() constant returns(uint256)
func (_Contract *ContractCaller) GetBidIndex(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "getBidIndex")
	return *ret0, err
}

// GetBidIndex is a free data retrieval call binding the contract method 0x05a13a68.
//
// Solidity: function getBidIndex() constant returns(uint256)
func (_Contract *ContractSession) GetBidIndex() (*big.Int, error) {
	return _Contract.Contract.GetBidIndex(&_Contract.CallOpts)
}

// GetBidIndex is a free data retrieval call binding the contract method 0x05a13a68.
//
// Solidity: function getBidIndex() constant returns(uint256)
func (_Contract *ContractCallerSession) GetBidIndex() (*big.Int, error) {
	return _Contract.Contract.GetBidIndex(&_Contract.CallOpts)
}

// HouseOwner is a free data retrieval call binding the contract method 0x64bb6270.
//
// Solidity: function houseOwner() constant returns(address)
func (_Contract *ContractCaller) HouseOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "houseOwner")
	return *ret0, err
}

// HouseOwner is a free data retrieval call binding the contract method 0x64bb6270.
//
// Solidity: function houseOwner() constant returns(address)
func (_Contract *ContractSession) HouseOwner() (common.Address, error) {
	return _Contract.Contract.HouseOwner(&_Contract.CallOpts)
}

// HouseOwner is a free data retrieval call binding the contract method 0x64bb6270.
//
// Solidity: function houseOwner() constant returns(address)
func (_Contract *ContractCallerSession) HouseOwner() (common.Address, error) {
	return _Contract.Contract.HouseOwner(&_Contract.CallOpts)
}

// Inspected is a free data retrieval call binding the contract method 0x44ec5640.
//
// Solidity: function inspected() constant returns(bool)
func (_Contract *ContractCaller) Inspected(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "inspected")
	return *ret0, err
}

// Inspected is a free data retrieval call binding the contract method 0x44ec5640.
//
// Solidity: function inspected() constant returns(bool)
func (_Contract *ContractSession) Inspected() (bool, error) {
	return _Contract.Contract.Inspected(&_Contract.CallOpts)
}

// Inspected is a free data retrieval call binding the contract method 0x44ec5640.
//
// Solidity: function inspected() constant returns(bool)
func (_Contract *ContractCallerSession) Inspected() (bool, error) {
	return _Contract.Contract.Inspected(&_Contract.CallOpts)
}

// IsForSale is a free data retrieval call binding the contract method 0x20b08dc2.
//
// Solidity: function isForSale() constant returns(bool)
func (_Contract *ContractCaller) IsForSale(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "isForSale")
	return *ret0, err
}

// IsForSale is a free data retrieval call binding the contract method 0x20b08dc2.
//
// Solidity: function isForSale() constant returns(bool)
func (_Contract *ContractSession) IsForSale() (bool, error) {
	return _Contract.Contract.IsForSale(&_Contract.CallOpts)
}

// IsForSale is a free data retrieval call binding the contract method 0x20b08dc2.
//
// Solidity: function isForSale() constant returns(bool)
func (_Contract *ContractCallerSession) IsForSale() (bool, error) {
	return _Contract.Contract.IsForSale(&_Contract.CallOpts)
}

// SoldPrice is a free data retrieval call binding the contract method 0xf746905b.
//
// Solidity: function soldPrice() constant returns(uint256)
func (_Contract *ContractCaller) SoldPrice(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "soldPrice")
	return *ret0, err
}

// SoldPrice is a free data retrieval call binding the contract method 0xf746905b.
//
// Solidity: function soldPrice() constant returns(uint256)
func (_Contract *ContractSession) SoldPrice() (*big.Int, error) {
	return _Contract.Contract.SoldPrice(&_Contract.CallOpts)
}

// SoldPrice is a free data retrieval call binding the contract method 0xf746905b.
//
// Solidity: function soldPrice() constant returns(uint256)
func (_Contract *ContractCallerSession) SoldPrice() (*big.Int, error) {
	return _Contract.Contract.SoldPrice(&_Contract.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() constant returns(uint8)
func (_Contract *ContractCaller) State(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "state")
	return *ret0, err
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() constant returns(uint8)
func (_Contract *ContractSession) State() (uint8, error) {
	return _Contract.Contract.State(&_Contract.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() constant returns(uint8)
func (_Contract *ContractCallerSession) State() (uint8, error) {
	return _Contract.Contract.State(&_Contract.CallOpts)
}

// StreetAddress is a free data retrieval call binding the contract method 0x47d90576.
//
// Solidity: function streetAddress() constant returns(bytes32)
func (_Contract *ContractCaller) StreetAddress(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "streetAddress")
	return *ret0, err
}

// StreetAddress is a free data retrieval call binding the contract method 0x47d90576.
//
// Solidity: function streetAddress() constant returns(bytes32)
func (_Contract *ContractSession) StreetAddress() ([32]byte, error) {
	return _Contract.Contract.StreetAddress(&_Contract.CallOpts)
}

// StreetAddress is a free data retrieval call binding the contract method 0x47d90576.
//
// Solidity: function streetAddress() constant returns(bytes32)
func (_Contract *ContractCallerSession) StreetAddress() ([32]byte, error) {
	return _Contract.Contract.StreetAddress(&_Contract.CallOpts)
}

// AcceptBid is a paid mutator transaction binding the contract method 0x6ba74f17.
//
// Solidity: function acceptBid(bidder address) returns()
func (_Contract *ContractTransactor) AcceptBid(opts *bind.TransactOpts, bidder common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "acceptBid", bidder)
}

// AcceptBid is a paid mutator transaction binding the contract method 0x6ba74f17.
//
// Solidity: function acceptBid(bidder address) returns()
func (_Contract *ContractSession) AcceptBid(bidder common.Address) (*types.Transaction, error) {
	return _Contract.Contract.AcceptBid(&_Contract.TransactOpts, bidder)
}

// AcceptBid is a paid mutator transaction binding the contract method 0x6ba74f17.
//
// Solidity: function acceptBid(bidder address) returns()
func (_Contract *ContractTransactorSession) AcceptBid(bidder common.Address) (*types.Transaction, error) {
	return _Contract.Contract.AcceptBid(&_Contract.TransactOpts, bidder)
}

// ApplyMortgage is a paid mutator transaction binding the contract method 0xc88c7211.
//
// Solidity: function applyMortgage(requestedMortgage uint256) returns()
func (_Contract *ContractTransactor) ApplyMortgage(opts *bind.TransactOpts, requestedMortgage *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "applyMortgage", requestedMortgage)
}

// ApplyMortgage is a paid mutator transaction binding the contract method 0xc88c7211.
//
// Solidity: function applyMortgage(requestedMortgage uint256) returns()
func (_Contract *ContractSession) ApplyMortgage(requestedMortgage *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ApplyMortgage(&_Contract.TransactOpts, requestedMortgage)
}

// ApplyMortgage is a paid mutator transaction binding the contract method 0xc88c7211.
//
// Solidity: function applyMortgage(requestedMortgage uint256) returns()
func (_Contract *ContractTransactorSession) ApplyMortgage(requestedMortgage *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ApplyMortgage(&_Contract.TransactOpts, requestedMortgage)
}

// ChangeMortgageValue is a paid mutator transaction binding the contract method 0x4054a1a1.
//
// Solidity: function changeMortgageValue(changedMortgageValue uint256) returns()
func (_Contract *ContractTransactor) ChangeMortgageValue(opts *bind.TransactOpts, changedMortgageValue *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "changeMortgageValue", changedMortgageValue)
}

// ChangeMortgageValue is a paid mutator transaction binding the contract method 0x4054a1a1.
//
// Solidity: function changeMortgageValue(changedMortgageValue uint256) returns()
func (_Contract *ContractSession) ChangeMortgageValue(changedMortgageValue *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ChangeMortgageValue(&_Contract.TransactOpts, changedMortgageValue)
}

// ChangeMortgageValue is a paid mutator transaction binding the contract method 0x4054a1a1.
//
// Solidity: function changeMortgageValue(changedMortgageValue uint256) returns()
func (_Contract *ContractTransactorSession) ChangeMortgageValue(changedMortgageValue *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ChangeMortgageValue(&_Contract.TransactOpts, changedMortgageValue)
}

// ChangePrice is a paid mutator transaction binding the contract method 0xa2b40d19.
//
// Solidity: function changePrice(newPrice uint256) returns()
func (_Contract *ContractTransactor) ChangePrice(opts *bind.TransactOpts, newPrice *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "changePrice", newPrice)
}

// ChangePrice is a paid mutator transaction binding the contract method 0xa2b40d19.
//
// Solidity: function changePrice(newPrice uint256) returns()
func (_Contract *ContractSession) ChangePrice(newPrice *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ChangePrice(&_Contract.TransactOpts, newPrice)
}

// ChangePrice is a paid mutator transaction binding the contract method 0xa2b40d19.
//
// Solidity: function changePrice(newPrice uint256) returns()
func (_Contract *ContractTransactorSession) ChangePrice(newPrice *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ChangePrice(&_Contract.TransactOpts, newPrice)
}

// DeleteAllBids is a paid mutator transaction binding the contract method 0xe1dfaab2.
//
// Solidity: function deleteAllBids() returns()
func (_Contract *ContractTransactor) DeleteAllBids(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "deleteAllBids")
}

// DeleteAllBids is a paid mutator transaction binding the contract method 0xe1dfaab2.
//
// Solidity: function deleteAllBids() returns()
func (_Contract *ContractSession) DeleteAllBids() (*types.Transaction, error) {
	return _Contract.Contract.DeleteAllBids(&_Contract.TransactOpts)
}

// DeleteAllBids is a paid mutator transaction binding the contract method 0xe1dfaab2.
//
// Solidity: function deleteAllBids() returns()
func (_Contract *ContractTransactorSession) DeleteAllBids() (*types.Transaction, error) {
	return _Contract.Contract.DeleteAllBids(&_Contract.TransactOpts)
}

// ForSale is a paid mutator transaction binding the contract method 0x2831f2f4.
//
// Solidity: function forSale(price uint256) returns()
func (_Contract *ContractTransactor) ForSale(opts *bind.TransactOpts, price *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "forSale", price)
}

// ForSale is a paid mutator transaction binding the contract method 0x2831f2f4.
//
// Solidity: function forSale(price uint256) returns()
func (_Contract *ContractSession) ForSale(price *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ForSale(&_Contract.TransactOpts, price)
}

// ForSale is a paid mutator transaction binding the contract method 0x2831f2f4.
//
// Solidity: function forSale(price uint256) returns()
func (_Contract *ContractTransactorSession) ForSale(price *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ForSale(&_Contract.TransactOpts, price)
}

// InspectionStatus is a paid mutator transaction binding the contract method 0x8175e543.
//
// Solidity: function inspectionStatus(status uint256) returns()
func (_Contract *ContractTransactor) InspectionStatus(opts *bind.TransactOpts, status *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "inspectionStatus", status)
}

// InspectionStatus is a paid mutator transaction binding the contract method 0x8175e543.
//
// Solidity: function inspectionStatus(status uint256) returns()
func (_Contract *ContractSession) InspectionStatus(status *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.InspectionStatus(&_Contract.TransactOpts, status)
}

// InspectionStatus is a paid mutator transaction binding the contract method 0x8175e543.
//
// Solidity: function inspectionStatus(status uint256) returns()
func (_Contract *ContractTransactorSession) InspectionStatus(status *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.InspectionStatus(&_Contract.TransactOpts, status)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Contract *ContractTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Contract *ContractSession) Kill() (*types.Transaction, error) {
	return _Contract.Contract.Kill(&_Contract.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Contract *ContractTransactorSession) Kill() (*types.Transaction, error) {
	return _Contract.Contract.Kill(&_Contract.TransactOpts)
}

// NotForSale is a paid mutator transaction binding the contract method 0x0ba062b5.
//
// Solidity: function notForSale() returns()
func (_Contract *ContractTransactor) NotForSale(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "notForSale")
}

// NotForSale is a paid mutator transaction binding the contract method 0x0ba062b5.
//
// Solidity: function notForSale() returns()
func (_Contract *ContractSession) NotForSale() (*types.Transaction, error) {
	return _Contract.Contract.NotForSale(&_Contract.TransactOpts)
}

// NotForSale is a paid mutator transaction binding the contract method 0x0ba062b5.
//
// Solidity: function notForSale() returns()
func (_Contract *ContractTransactorSession) NotForSale() (*types.Transaction, error) {
	return _Contract.Contract.NotForSale(&_Contract.TransactOpts)
}

// PlaceBid is a paid mutator transaction binding the contract method 0x9979ef45.
//
// Solidity: function placeBid(biddingPrice uint256) returns()
func (_Contract *ContractTransactor) PlaceBid(opts *bind.TransactOpts, biddingPrice *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "placeBid", biddingPrice)
}

// PlaceBid is a paid mutator transaction binding the contract method 0x9979ef45.
//
// Solidity: function placeBid(biddingPrice uint256) returns()
func (_Contract *ContractSession) PlaceBid(biddingPrice *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.PlaceBid(&_Contract.TransactOpts, biddingPrice)
}

// PlaceBid is a paid mutator transaction binding the contract method 0x9979ef45.
//
// Solidity: function placeBid(biddingPrice uint256) returns()
func (_Contract *ContractTransactorSession) PlaceBid(biddingPrice *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.PlaceBid(&_Contract.TransactOpts, biddingPrice)
}

// RejectBid is a paid mutator transaction binding the contract method 0x609468fe.
//
// Solidity: function rejectBid(bidderAddress address) returns()
func (_Contract *ContractTransactor) RejectBid(opts *bind.TransactOpts, bidderAddress common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "rejectBid", bidderAddress)
}

// RejectBid is a paid mutator transaction binding the contract method 0x609468fe.
//
// Solidity: function rejectBid(bidderAddress address) returns()
func (_Contract *ContractSession) RejectBid(bidderAddress common.Address) (*types.Transaction, error) {
	return _Contract.Contract.RejectBid(&_Contract.TransactOpts, bidderAddress)
}

// RejectBid is a paid mutator transaction binding the contract method 0x609468fe.
//
// Solidity: function rejectBid(bidderAddress address) returns()
func (_Contract *ContractTransactorSession) RejectBid(bidderAddress common.Address) (*types.Transaction, error) {
	return _Contract.Contract.RejectBid(&_Contract.TransactOpts, bidderAddress)
}

// RejectMortgage is a paid mutator transaction binding the contract method 0xbcdea56e.
//
// Solidity: function rejectMortgage(value uint256) returns()
func (_Contract *ContractTransactor) RejectMortgage(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "rejectMortgage", value)
}

// RejectMortgage is a paid mutator transaction binding the contract method 0xbcdea56e.
//
// Solidity: function rejectMortgage(value uint256) returns()
func (_Contract *ContractSession) RejectMortgage(value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.RejectMortgage(&_Contract.TransactOpts, value)
}

// RejectMortgage is a paid mutator transaction binding the contract method 0xbcdea56e.
//
// Solidity: function rejectMortgage(value uint256) returns()
func (_Contract *ContractTransactorSession) RejectMortgage(value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.RejectMortgage(&_Contract.TransactOpts, value)
}

// SetAppraisalValue is a paid mutator transaction binding the contract method 0x0cf6bbfb.
//
// Solidity: function setAppraisalValue(value uint256) returns()
func (_Contract *ContractTransactor) SetAppraisalValue(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setAppraisalValue", value)
}

// SetAppraisalValue is a paid mutator transaction binding the contract method 0x0cf6bbfb.
//
// Solidity: function setAppraisalValue(value uint256) returns()
func (_Contract *ContractSession) SetAppraisalValue(value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetAppraisalValue(&_Contract.TransactOpts, value)
}

// SetAppraisalValue is a paid mutator transaction binding the contract method 0x0cf6bbfb.
//
// Solidity: function setAppraisalValue(value uint256) returns()
func (_Contract *ContractTransactorSession) SetAppraisalValue(value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetAppraisalValue(&_Contract.TransactOpts, value)
}

// Sold is a paid mutator transaction binding the contract method 0x02c7e7af.
//
// Solidity: function sold() returns()
func (_Contract *ContractTransactor) Sold(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "sold")
}

// Sold is a paid mutator transaction binding the contract method 0x02c7e7af.
//
// Solidity: function sold() returns()
func (_Contract *ContractSession) Sold() (*types.Transaction, error) {
	return _Contract.Contract.Sold(&_Contract.TransactOpts)
}

// Sold is a paid mutator transaction binding the contract method 0x02c7e7af.
//
// Solidity: function sold() returns()
func (_Contract *ContractTransactorSession) Sold() (*types.Transaction, error) {
	return _Contract.Contract.Sold(&_Contract.TransactOpts)
}
