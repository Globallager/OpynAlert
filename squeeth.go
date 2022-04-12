// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// SqueethMetaData contains all meta data concerning the Squeeth contract.
var SqueethMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_shortPowerPerp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_wPowerPerp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_weth\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_quoteCurrency\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_ethQuoteCurrencyPool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_wPowerPerpPool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_uniPositionManager\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"_feeTier\",\"type\":\"uint24\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"BurnShort\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DepositCollateral\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"DepositUniPositionToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newFee\",\"type\":\"uint256\"}],\"name\":\"FeeRateUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldFeeRecipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newFeeRecipient\",\"type\":\"address\"}],\"name\":\"FeeRecipientUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"debtAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateralPaid\",\"type\":\"uint256\"}],\"name\":\"Liquidate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"MintShort\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldNormFactor\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newNormFactor\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lastModificationTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"NormalizationFactorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"}],\"name\":\"OpenVault\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pausesLeft\",\"type\":\"uint256\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wPowerPerpAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"payoutAmount\",\"type\":\"uint256\"}],\"name\":\"RedeemLong\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vauldId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateralAmount\",\"type\":\"uint256\"}],\"name\":\"RedeemShort\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ethRedeemed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wPowerPerpRedeemed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wPowerPerpBurned\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wPowerPerpExcess\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bounty\",\"type\":\"uint256\"}],\"name\":\"ReduceDebt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"indexForSettlement\",\"type\":\"uint256\"}],\"name\":\"Shutdown\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"unpauser\",\"type\":\"address\"}],\"name\":\"UnPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"UpdateOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawCollateral\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"WithdrawUniPositionToken\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"FUNDING_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TWAP_PERIOD\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"applyFunding\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_powerPerpAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_withdrawAmount\",\"type\":\"uint256\"}],\"name\":\"burnPowerPerpAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_wPowerPerpAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_withdrawAmount\",\"type\":\"uint256\"}],\"name\":\"burnWPowerPerpAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_vaultId\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_uniTokenId\",\"type\":\"uint256\"}],\"name\":\"depositUniPositionToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"donate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ethQuoteCurrencyPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeTier\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_period\",\"type\":\"uint32\"}],\"name\":\"getDenormalizedMark\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_period\",\"type\":\"uint32\"}],\"name\":\"getDenormalizedMarkForFunding\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getExpectedNormalizationFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_period\",\"type\":\"uint32\"}],\"name\":\"getIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_period\",\"type\":\"uint32\"}],\"name\":\"getUnscaledIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"indexForSettlement\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isShutDown\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isSystemPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_vaultId\",\"type\":\"uint256\"}],\"name\":\"isVaultSafe\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastFundingUpdateTimestamp\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastPauseTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxDebtAmount\",\"type\":\"uint256\"}],\"name\":\"liquidate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_powerPerpAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_uniTokenId\",\"type\":\"uint256\"}],\"name\":\"mintPowerPerpAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_wPowerPerpAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_uniTokenId\",\"type\":\"uint256\"}],\"name\":\"mintWPowerPerpAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"normalizationFactor\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pausesLeft\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quoteCurrency\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_wPerpAmount\",\"type\":\"uint256\"}],\"name\":\"redeemLong\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_vaultId\",\"type\":\"uint256\"}],\"name\":\"redeemShort\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_vaultId\",\"type\":\"uint256\"}],\"name\":\"reduceDebt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_vaultId\",\"type\":\"uint256\"}],\"name\":\"reduceDebtShutdown\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newFeeRate\",\"type\":\"uint256\"}],\"name\":\"setFeeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newFeeRecipient\",\"type\":\"address\"}],\"name\":\"setFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"shortPowerPerp\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"shutDown\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unPauseAnyone\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unPauseOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_vaultId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"updateOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"vaults\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"NftCollateralId\",\"type\":\"uint32\"},{\"internalType\":\"uint96\",\"name\":\"collateralAmount\",\"type\":\"uint96\"},{\"internalType\":\"uint128\",\"name\":\"shortAmount\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wPowerPerp\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wPowerPerpPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"weth\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_vaultId\",\"type\":\"uint256\"}],\"name\":\"withdrawUniPositionToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// SqueethABI is the input ABI used to generate the binding from.
// Deprecated: Use SqueethMetaData.ABI instead.
var SqueethABI = SqueethMetaData.ABI

// Squeeth is an auto generated Go binding around an Ethereum contract.
type Squeeth struct {
	SqueethCaller     // Read-only binding to the contract
	SqueethTransactor // Write-only binding to the contract
	SqueethFilterer   // Log filterer for contract events
}

// SqueethCaller is an auto generated read-only Go binding around an Ethereum contract.
type SqueethCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SqueethTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SqueethTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SqueethFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SqueethFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SqueethSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SqueethSession struct {
	Contract     *Squeeth          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SqueethCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SqueethCallerSession struct {
	Contract *SqueethCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// SqueethTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SqueethTransactorSession struct {
	Contract     *SqueethTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// SqueethRaw is an auto generated low-level Go binding around an Ethereum contract.
type SqueethRaw struct {
	Contract *Squeeth // Generic contract binding to access the raw methods on
}

// SqueethCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SqueethCallerRaw struct {
	Contract *SqueethCaller // Generic read-only contract binding to access the raw methods on
}

// SqueethTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SqueethTransactorRaw struct {
	Contract *SqueethTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSqueeth creates a new instance of Squeeth, bound to a specific deployed contract.
func NewSqueeth(address common.Address, backend bind.ContractBackend) (*Squeeth, error) {
	contract, err := bindSqueeth(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Squeeth{SqueethCaller: SqueethCaller{contract: contract}, SqueethTransactor: SqueethTransactor{contract: contract}, SqueethFilterer: SqueethFilterer{contract: contract}}, nil
}

// NewSqueethCaller creates a new read-only instance of Squeeth, bound to a specific deployed contract.
func NewSqueethCaller(address common.Address, caller bind.ContractCaller) (*SqueethCaller, error) {
	contract, err := bindSqueeth(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SqueethCaller{contract: contract}, nil
}

// NewSqueethTransactor creates a new write-only instance of Squeeth, bound to a specific deployed contract.
func NewSqueethTransactor(address common.Address, transactor bind.ContractTransactor) (*SqueethTransactor, error) {
	contract, err := bindSqueeth(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SqueethTransactor{contract: contract}, nil
}

// NewSqueethFilterer creates a new log filterer instance of Squeeth, bound to a specific deployed contract.
func NewSqueethFilterer(address common.Address, filterer bind.ContractFilterer) (*SqueethFilterer, error) {
	contract, err := bindSqueeth(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SqueethFilterer{contract: contract}, nil
}

// bindSqueeth binds a generic wrapper to an already deployed contract.
func bindSqueeth(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SqueethABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Squeeth *SqueethRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Squeeth.Contract.SqueethCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Squeeth *SqueethRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Squeeth.Contract.SqueethTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Squeeth *SqueethRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Squeeth.Contract.SqueethTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Squeeth *SqueethCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Squeeth.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Squeeth *SqueethTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Squeeth.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Squeeth *SqueethTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Squeeth.Contract.contract.Transact(opts, method, params...)
}

// FUNDINGPERIOD is a free data retrieval call binding the contract method 0xf90c3f27.
//
// Solidity: function FUNDING_PERIOD() view returns(uint256)
func (_Squeeth *SqueethCaller) FUNDINGPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Squeeth.contract.Call(opts, &out, "FUNDING_PERIOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FUNDINGPERIOD is a free data retrieval call binding the contract method 0xf90c3f27.
//
// Solidity: function FUNDING_PERIOD() view returns(uint256)
func (_Squeeth *SqueethSession) FUNDINGPERIOD() (*big.Int, error) {
	return _Squeeth.Contract.FUNDINGPERIOD(&_Squeeth.CallOpts)
}

// FUNDINGPERIOD is a free data retrieval call binding the contract method 0xf90c3f27.
//
// Solidity: function FUNDING_PERIOD() view returns(uint256)
func (_Squeeth *SqueethCallerSession) FUNDINGPERIOD() (*big.Int, error) {
	return _Squeeth.Contract.FUNDINGPERIOD(&_Squeeth.CallOpts)
}

// TWAPPERIOD is a free data retrieval call binding the contract method 0x7ca25184.
//
// Solidity: function TWAP_PERIOD() view returns(uint32)
func (_Squeeth *SqueethCaller) TWAPPERIOD(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Squeeth.contract.Call(opts, &out, "TWAP_PERIOD")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TWAPPERIOD is a free data retrieval call binding the contract method 0x7ca25184.
//
// Solidity: function TWAP_PERIOD() view returns(uint32)
func (_Squeeth *SqueethSession) TWAPPERIOD() (uint32, error) {
	return _Squeeth.Contract.TWAPPERIOD(&_Squeeth.CallOpts)
}

// TWAPPERIOD is a free data retrieval call binding the contract method 0x7ca25184.
//
// Solidity: function TWAP_PERIOD() view returns(uint32)
func (_Squeeth *SqueethCallerSession) TWAPPERIOD() (uint32, error) {
	return _Squeeth.Contract.TWAPPERIOD(&_Squeeth.CallOpts)
}

// EthQuoteCurrencyPool is a free data retrieval call binding the contract method 0x4468c022.
//
// Solidity: function ethQuoteCurrencyPool() view returns(address)
func (_Squeeth *SqueethCaller) EthQuoteCurrencyPool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Squeeth.contract.Call(opts, &out, "ethQuoteCurrencyPool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EthQuoteCurrencyPool is a free data retrieval call binding the contract method 0x4468c022.
//
// Solidity: function ethQuoteCurrencyPool() view returns(address)
func (_Squeeth *SqueethSession) EthQuoteCurrencyPool() (common.Address, error) {
	return _Squeeth.Contract.EthQuoteCurrencyPool(&_Squeeth.CallOpts)
}

// EthQuoteCurrencyPool is a free data retrieval call binding the contract method 0x4468c022.
//
// Solidity: function ethQuoteCurrencyPool() view returns(address)
func (_Squeeth *SqueethCallerSession) EthQuoteCurrencyPool() (common.Address, error) {
	return _Squeeth.Contract.EthQuoteCurrencyPool(&_Squeeth.CallOpts)
}

// FeeRate is a free data retrieval call binding the contract method 0x978bbdb9.
//
// Solidity: function feeRate() view returns(uint256)
func (_Squeeth *SqueethCaller) FeeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Squeeth.contract.Call(opts, &out, "feeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeRate is a free data retrieval call binding the contract method 0x978bbdb9.
//
// Solidity: function feeRate() view returns(uint256)
func (_Squeeth *SqueethSession) FeeRate() (*big.Int, error) {
	return _Squeeth.Contract.FeeRate(&_Squeeth.CallOpts)
}

// FeeRate is a free data retrieval call binding the contract method 0x978bbdb9.
//
// Solidity: function feeRate() view returns(uint256)
func (_Squeeth *SqueethCallerSession) FeeRate() (*big.Int, error) {
	return _Squeeth.Contract.FeeRate(&_Squeeth.CallOpts)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_Squeeth *SqueethCaller) FeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Squeeth.contract.Call(opts, &out, "feeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_Squeeth *SqueethSession) FeeRecipient() (common.Address, error) {
	return _Squeeth.Contract.FeeRecipient(&_Squeeth.CallOpts)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_Squeeth *SqueethCallerSession) FeeRecipient() (common.Address, error) {
	return _Squeeth.Contract.FeeRecipient(&_Squeeth.CallOpts)
}

// FeeTier is a free data retrieval call binding the contract method 0x72f5d98a.
//
// Solidity: function feeTier() view returns(uint24)
func (_Squeeth *SqueethCaller) FeeTier(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Squeeth.contract.Call(opts, &out, "feeTier")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeTier is a free data retrieval call binding the contract method 0x72f5d98a.
//
// Solidity: function feeTier() view returns(uint24)
func (_Squeeth *SqueethSession) FeeTier() (*big.Int, error) {
	return _Squeeth.Contract.FeeTier(&_Squeeth.CallOpts)
}

// FeeTier is a free data retrieval call binding the contract method 0x72f5d98a.
//
// Solidity: function feeTier() view returns(uint24)
func (_Squeeth *SqueethCallerSession) FeeTier() (*big.Int, error) {
	return _Squeeth.Contract.FeeTier(&_Squeeth.CallOpts)
}

// GetDenormalizedMark is a free data retrieval call binding the contract method 0xee3189ff.
//
// Solidity: function getDenormalizedMark(uint32 _period) view returns(uint256)
func (_Squeeth *SqueethCaller) GetDenormalizedMark(opts *bind.CallOpts, _period uint32) (*big.Int, error) {
	var out []interface{}
	err := _Squeeth.contract.Call(opts, &out, "getDenormalizedMark", _period)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDenormalizedMark is a free data retrieval call binding the contract method 0xee3189ff.
//
// Solidity: function getDenormalizedMark(uint32 _period) view returns(uint256)
func (_Squeeth *SqueethSession) GetDenormalizedMark(_period uint32) (*big.Int, error) {
	return _Squeeth.Contract.GetDenormalizedMark(&_Squeeth.CallOpts, _period)
}

// GetDenormalizedMark is a free data retrieval call binding the contract method 0xee3189ff.
//
// Solidity: function getDenormalizedMark(uint32 _period) view returns(uint256)
func (_Squeeth *SqueethCallerSession) GetDenormalizedMark(_period uint32) (*big.Int, error) {
	return _Squeeth.Contract.GetDenormalizedMark(&_Squeeth.CallOpts, _period)
}

// GetDenormalizedMarkForFunding is a free data retrieval call binding the contract method 0x377a1936.
//
// Solidity: function getDenormalizedMarkForFunding(uint32 _period) view returns(uint256)
func (_Squeeth *SqueethCaller) GetDenormalizedMarkForFunding(opts *bind.CallOpts, _period uint32) (*big.Int, error) {
	var out []interface{}
	err := _Squeeth.contract.Call(opts, &out, "getDenormalizedMarkForFunding", _period)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDenormalizedMarkForFunding is a free data retrieval call binding the contract method 0x377a1936.
//
// Solidity: function getDenormalizedMarkForFunding(uint32 _period) view returns(uint256)
func (_Squeeth *SqueethSession) GetDenormalizedMarkForFunding(_period uint32) (*big.Int, error) {
	return _Squeeth.Contract.GetDenormalizedMarkForFunding(&_Squeeth.CallOpts, _period)
}

// GetDenormalizedMarkForFunding is a free data retrieval call binding the contract method 0x377a1936.
//
// Solidity: function getDenormalizedMarkForFunding(uint32 _period) view returns(uint256)
func (_Squeeth *SqueethCallerSession) GetDenormalizedMarkForFunding(_period uint32) (*big.Int, error) {
	return _Squeeth.Contract.GetDenormalizedMarkForFunding(&_Squeeth.CallOpts, _period)
}

// GetExpectedNormalizationFactor is a free data retrieval call binding the contract method 0x24f5f531.
//
// Solidity: function getExpectedNormalizationFactor() view returns(uint256)
func (_Squeeth *SqueethCaller) GetExpectedNormalizationFactor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Squeeth.contract.Call(opts, &out, "getExpectedNormalizationFactor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetExpectedNormalizationFactor is a free data retrieval call binding the contract method 0x24f5f531.
//
// Solidity: function getExpectedNormalizationFactor() view returns(uint256)
func (_Squeeth *SqueethSession) GetExpectedNormalizationFactor() (*big.Int, error) {
	return _Squeeth.Contract.GetExpectedNormalizationFactor(&_Squeeth.CallOpts)
}

// GetExpectedNormalizationFactor is a free data retrieval call binding the contract method 0x24f5f531.
//
// Solidity: function getExpectedNormalizationFactor() view returns(uint256)
func (_Squeeth *SqueethCallerSession) GetExpectedNormalizationFactor() (*big.Int, error) {
	return _Squeeth.Contract.GetExpectedNormalizationFactor(&_Squeeth.CallOpts)
}

// GetIndex is a free data retrieval call binding the contract method 0x8cd21d7c.
//
// Solidity: function getIndex(uint32 _period) view returns(uint256)
func (_Squeeth *SqueethCaller) GetIndex(opts *bind.CallOpts, _period uint32) (*big.Int, error) {
	var out []interface{}
	err := _Squeeth.contract.Call(opts, &out, "getIndex", _period)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetIndex is a free data retrieval call binding the contract method 0x8cd21d7c.
//
// Solidity: function getIndex(uint32 _period) view returns(uint256)
func (_Squeeth *SqueethSession) GetIndex(_period uint32) (*big.Int, error) {
	return _Squeeth.Contract.GetIndex(&_Squeeth.CallOpts, _period)
}

// GetIndex is a free data retrieval call binding the contract method 0x8cd21d7c.
//
// Solidity: function getIndex(uint32 _period) view returns(uint256)
func (_Squeeth *SqueethCallerSession) GetIndex(_period uint32) (*big.Int, error) {
	return _Squeeth.Contract.GetIndex(&_Squeeth.CallOpts, _period)
}

// GetUnscaledIndex is a free data retrieval call binding the contract method 0x15aded83.
//
// Solidity: function getUnscaledIndex(uint32 _period) view returns(uint256)
func (_Squeeth *SqueethCaller) GetUnscaledIndex(opts *bind.CallOpts, _period uint32) (*big.Int, error) {
	var out []interface{}
	err := _Squeeth.contract.Call(opts, &out, "getUnscaledIndex", _period)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUnscaledIndex is a free data retrieval call binding the contract method 0x15aded83.
//
// Solidity: function getUnscaledIndex(uint32 _period) view returns(uint256)
func (_Squeeth *SqueethSession) GetUnscaledIndex(_period uint32) (*big.Int, error) {
	return _Squeeth.Contract.GetUnscaledIndex(&_Squeeth.CallOpts, _period)
}

// GetUnscaledIndex is a free data retrieval call binding the contract method 0x15aded83.
//
// Solidity: function getUnscaledIndex(uint32 _period) view returns(uint256)
func (_Squeeth *SqueethCallerSession) GetUnscaledIndex(_period uint32) (*big.Int, error) {
	return _Squeeth.Contract.GetUnscaledIndex(&_Squeeth.CallOpts, _period)
}

// IndexForSettlement is a free data retrieval call binding the contract method 0xde4a427a.
//
// Solidity: function indexForSettlement() view returns(uint256)
func (_Squeeth *SqueethCaller) IndexForSettlement(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Squeeth.contract.Call(opts, &out, "indexForSettlement")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IndexForSettlement is a free data retrieval call binding the contract method 0xde4a427a.
//
// Solidity: function indexForSettlement() view returns(uint256)
func (_Squeeth *SqueethSession) IndexForSettlement() (*big.Int, error) {
	return _Squeeth.Contract.IndexForSettlement(&_Squeeth.CallOpts)
}

// IndexForSettlement is a free data retrieval call binding the contract method 0xde4a427a.
//
// Solidity: function indexForSettlement() view returns(uint256)
func (_Squeeth *SqueethCallerSession) IndexForSettlement() (*big.Int, error) {
	return _Squeeth.Contract.IndexForSettlement(&_Squeeth.CallOpts)
}

// IsShutDown is a free data retrieval call binding the contract method 0xff947525.
//
// Solidity: function isShutDown() view returns(bool)
func (_Squeeth *SqueethCaller) IsShutDown(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Squeeth.contract.Call(opts, &out, "isShutDown")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsShutDown is a free data retrieval call binding the contract method 0xff947525.
//
// Solidity: function isShutDown() view returns(bool)
func (_Squeeth *SqueethSession) IsShutDown() (bool, error) {
	return _Squeeth.Contract.IsShutDown(&_Squeeth.CallOpts)
}

// IsShutDown is a free data retrieval call binding the contract method 0xff947525.
//
// Solidity: function isShutDown() view returns(bool)
func (_Squeeth *SqueethCallerSession) IsShutDown() (bool, error) {
	return _Squeeth.Contract.IsShutDown(&_Squeeth.CallOpts)
}

// IsSystemPaused is a free data retrieval call binding the contract method 0x7691c4ac.
//
// Solidity: function isSystemPaused() view returns(bool)
func (_Squeeth *SqueethCaller) IsSystemPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Squeeth.contract.Call(opts, &out, "isSystemPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSystemPaused is a free data retrieval call binding the contract method 0x7691c4ac.
//
// Solidity: function isSystemPaused() view returns(bool)
func (_Squeeth *SqueethSession) IsSystemPaused() (bool, error) {
	return _Squeeth.Contract.IsSystemPaused(&_Squeeth.CallOpts)
}

// IsSystemPaused is a free data retrieval call binding the contract method 0x7691c4ac.
//
// Solidity: function isSystemPaused() view returns(bool)
func (_Squeeth *SqueethCallerSession) IsSystemPaused() (bool, error) {
	return _Squeeth.Contract.IsSystemPaused(&_Squeeth.CallOpts)
}

// IsVaultSafe is a free data retrieval call binding the contract method 0xa847e674.
//
// Solidity: function isVaultSafe(uint256 _vaultId) view returns(bool)
func (_Squeeth *SqueethCaller) IsVaultSafe(opts *bind.CallOpts, _vaultId *big.Int) (bool, error) {
	var out []interface{}
	err := _Squeeth.contract.Call(opts, &out, "isVaultSafe", _vaultId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsVaultSafe is a free data retrieval call binding the contract method 0xa847e674.
//
// Solidity: function isVaultSafe(uint256 _vaultId) view returns(bool)
func (_Squeeth *SqueethSession) IsVaultSafe(_vaultId *big.Int) (bool, error) {
	return _Squeeth.Contract.IsVaultSafe(&_Squeeth.CallOpts, _vaultId)
}

// IsVaultSafe is a free data retrieval call binding the contract method 0xa847e674.
//
// Solidity: function isVaultSafe(uint256 _vaultId) view returns(bool)
func (_Squeeth *SqueethCallerSession) IsVaultSafe(_vaultId *big.Int) (bool, error) {
	return _Squeeth.Contract.IsVaultSafe(&_Squeeth.CallOpts, _vaultId)
}

// LastFundingUpdateTimestamp is a free data retrieval call binding the contract method 0x4be2822c.
//
// Solidity: function lastFundingUpdateTimestamp() view returns(uint128)
func (_Squeeth *SqueethCaller) LastFundingUpdateTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Squeeth.contract.Call(opts, &out, "lastFundingUpdateTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastFundingUpdateTimestamp is a free data retrieval call binding the contract method 0x4be2822c.
//
// Solidity: function lastFundingUpdateTimestamp() view returns(uint128)
func (_Squeeth *SqueethSession) LastFundingUpdateTimestamp() (*big.Int, error) {
	return _Squeeth.Contract.LastFundingUpdateTimestamp(&_Squeeth.CallOpts)
}

// LastFundingUpdateTimestamp is a free data retrieval call binding the contract method 0x4be2822c.
//
// Solidity: function lastFundingUpdateTimestamp() view returns(uint128)
func (_Squeeth *SqueethCallerSession) LastFundingUpdateTimestamp() (*big.Int, error) {
	return _Squeeth.Contract.LastFundingUpdateTimestamp(&_Squeeth.CallOpts)
}

// LastPauseTime is a free data retrieval call binding the contract method 0x91b4ded9.
//
// Solidity: function lastPauseTime() view returns(uint256)
func (_Squeeth *SqueethCaller) LastPauseTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Squeeth.contract.Call(opts, &out, "lastPauseTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastPauseTime is a free data retrieval call binding the contract method 0x91b4ded9.
//
// Solidity: function lastPauseTime() view returns(uint256)
func (_Squeeth *SqueethSession) LastPauseTime() (*big.Int, error) {
	return _Squeeth.Contract.LastPauseTime(&_Squeeth.CallOpts)
}

// LastPauseTime is a free data retrieval call binding the contract method 0x91b4ded9.
//
// Solidity: function lastPauseTime() view returns(uint256)
func (_Squeeth *SqueethCallerSession) LastPauseTime() (*big.Int, error) {
	return _Squeeth.Contract.LastPauseTime(&_Squeeth.CallOpts)
}

// NormalizationFactor is a free data retrieval call binding the contract method 0xd5272584.
//
// Solidity: function normalizationFactor() view returns(uint128)
func (_Squeeth *SqueethCaller) NormalizationFactor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Squeeth.contract.Call(opts, &out, "normalizationFactor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NormalizationFactor is a free data retrieval call binding the contract method 0xd5272584.
//
// Solidity: function normalizationFactor() view returns(uint128)
func (_Squeeth *SqueethSession) NormalizationFactor() (*big.Int, error) {
	return _Squeeth.Contract.NormalizationFactor(&_Squeeth.CallOpts)
}

// NormalizationFactor is a free data retrieval call binding the contract method 0xd5272584.
//
// Solidity: function normalizationFactor() view returns(uint128)
func (_Squeeth *SqueethCallerSession) NormalizationFactor() (*big.Int, error) {
	return _Squeeth.Contract.NormalizationFactor(&_Squeeth.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Squeeth *SqueethCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Squeeth.contract.Call(opts, &out, "oracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Squeeth *SqueethSession) Oracle() (common.Address, error) {
	return _Squeeth.Contract.Oracle(&_Squeeth.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Squeeth *SqueethCallerSession) Oracle() (common.Address, error) {
	return _Squeeth.Contract.Oracle(&_Squeeth.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Squeeth *SqueethCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Squeeth.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Squeeth *SqueethSession) Owner() (common.Address, error) {
	return _Squeeth.Contract.Owner(&_Squeeth.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Squeeth *SqueethCallerSession) Owner() (common.Address, error) {
	return _Squeeth.Contract.Owner(&_Squeeth.CallOpts)
}

// PausesLeft is a free data retrieval call binding the contract method 0x63b38ae4.
//
// Solidity: function pausesLeft() view returns(uint256)
func (_Squeeth *SqueethCaller) PausesLeft(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Squeeth.contract.Call(opts, &out, "pausesLeft")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PausesLeft is a free data retrieval call binding the contract method 0x63b38ae4.
//
// Solidity: function pausesLeft() view returns(uint256)
func (_Squeeth *SqueethSession) PausesLeft() (*big.Int, error) {
	return _Squeeth.Contract.PausesLeft(&_Squeeth.CallOpts)
}

// PausesLeft is a free data retrieval call binding the contract method 0x63b38ae4.
//
// Solidity: function pausesLeft() view returns(uint256)
func (_Squeeth *SqueethCallerSession) PausesLeft() (*big.Int, error) {
	return _Squeeth.Contract.PausesLeft(&_Squeeth.CallOpts)
}

// QuoteCurrency is a free data retrieval call binding the contract method 0x82564bca.
//
// Solidity: function quoteCurrency() view returns(address)
func (_Squeeth *SqueethCaller) QuoteCurrency(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Squeeth.contract.Call(opts, &out, "quoteCurrency")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// QuoteCurrency is a free data retrieval call binding the contract method 0x82564bca.
//
// Solidity: function quoteCurrency() view returns(address)
func (_Squeeth *SqueethSession) QuoteCurrency() (common.Address, error) {
	return _Squeeth.Contract.QuoteCurrency(&_Squeeth.CallOpts)
}

// QuoteCurrency is a free data retrieval call binding the contract method 0x82564bca.
//
// Solidity: function quoteCurrency() view returns(address)
func (_Squeeth *SqueethCallerSession) QuoteCurrency() (common.Address, error) {
	return _Squeeth.Contract.QuoteCurrency(&_Squeeth.CallOpts)
}

// ShortPowerPerp is a free data retrieval call binding the contract method 0x9d4c9442.
//
// Solidity: function shortPowerPerp() view returns(address)
func (_Squeeth *SqueethCaller) ShortPowerPerp(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Squeeth.contract.Call(opts, &out, "shortPowerPerp")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ShortPowerPerp is a free data retrieval call binding the contract method 0x9d4c9442.
//
// Solidity: function shortPowerPerp() view returns(address)
func (_Squeeth *SqueethSession) ShortPowerPerp() (common.Address, error) {
	return _Squeeth.Contract.ShortPowerPerp(&_Squeeth.CallOpts)
}

// ShortPowerPerp is a free data retrieval call binding the contract method 0x9d4c9442.
//
// Solidity: function shortPowerPerp() view returns(address)
func (_Squeeth *SqueethCallerSession) ShortPowerPerp() (common.Address, error) {
	return _Squeeth.Contract.ShortPowerPerp(&_Squeeth.CallOpts)
}

// Vaults is a free data retrieval call binding the contract method 0x8c64ea4a.
//
// Solidity: function vaults(uint256 ) view returns(address operator, uint32 NftCollateralId, uint96 collateralAmount, uint128 shortAmount)
func (_Squeeth *SqueethCaller) Vaults(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Operator         common.Address
	NftCollateralId  uint32
	CollateralAmount *big.Int
	ShortAmount      *big.Int
}, error) {
	var out []interface{}
	err := _Squeeth.contract.Call(opts, &out, "vaults", arg0)

	outstruct := new(struct {
		Operator         common.Address
		NftCollateralId  uint32
		CollateralAmount *big.Int
		ShortAmount      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Operator = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.NftCollateralId = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.CollateralAmount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.ShortAmount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Vaults is a free data retrieval call binding the contract method 0x8c64ea4a.
//
// Solidity: function vaults(uint256 ) view returns(address operator, uint32 NftCollateralId, uint96 collateralAmount, uint128 shortAmount)
func (_Squeeth *SqueethSession) Vaults(arg0 *big.Int) (struct {
	Operator         common.Address
	NftCollateralId  uint32
	CollateralAmount *big.Int
	ShortAmount      *big.Int
}, error) {
	return _Squeeth.Contract.Vaults(&_Squeeth.CallOpts, arg0)
}

// Vaults is a free data retrieval call binding the contract method 0x8c64ea4a.
//
// Solidity: function vaults(uint256 ) view returns(address operator, uint32 NftCollateralId, uint96 collateralAmount, uint128 shortAmount)
func (_Squeeth *SqueethCallerSession) Vaults(arg0 *big.Int) (struct {
	Operator         common.Address
	NftCollateralId  uint32
	CollateralAmount *big.Int
	ShortAmount      *big.Int
}, error) {
	return _Squeeth.Contract.Vaults(&_Squeeth.CallOpts, arg0)
}

// WPowerPerp is a free data retrieval call binding the contract method 0x7f07b130.
//
// Solidity: function wPowerPerp() view returns(address)
func (_Squeeth *SqueethCaller) WPowerPerp(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Squeeth.contract.Call(opts, &out, "wPowerPerp")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WPowerPerp is a free data retrieval call binding the contract method 0x7f07b130.
//
// Solidity: function wPowerPerp() view returns(address)
func (_Squeeth *SqueethSession) WPowerPerp() (common.Address, error) {
	return _Squeeth.Contract.WPowerPerp(&_Squeeth.CallOpts)
}

// WPowerPerp is a free data retrieval call binding the contract method 0x7f07b130.
//
// Solidity: function wPowerPerp() view returns(address)
func (_Squeeth *SqueethCallerSession) WPowerPerp() (common.Address, error) {
	return _Squeeth.Contract.WPowerPerp(&_Squeeth.CallOpts)
}

// WPowerPerpPool is a free data retrieval call binding the contract method 0xb707ab99.
//
// Solidity: function wPowerPerpPool() view returns(address)
func (_Squeeth *SqueethCaller) WPowerPerpPool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Squeeth.contract.Call(opts, &out, "wPowerPerpPool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WPowerPerpPool is a free data retrieval call binding the contract method 0xb707ab99.
//
// Solidity: function wPowerPerpPool() view returns(address)
func (_Squeeth *SqueethSession) WPowerPerpPool() (common.Address, error) {
	return _Squeeth.Contract.WPowerPerpPool(&_Squeeth.CallOpts)
}

// WPowerPerpPool is a free data retrieval call binding the contract method 0xb707ab99.
//
// Solidity: function wPowerPerpPool() view returns(address)
func (_Squeeth *SqueethCallerSession) WPowerPerpPool() (common.Address, error) {
	return _Squeeth.Contract.WPowerPerpPool(&_Squeeth.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_Squeeth *SqueethCaller) Weth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Squeeth.contract.Call(opts, &out, "weth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_Squeeth *SqueethSession) Weth() (common.Address, error) {
	return _Squeeth.Contract.Weth(&_Squeeth.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_Squeeth *SqueethCallerSession) Weth() (common.Address, error) {
	return _Squeeth.Contract.Weth(&_Squeeth.CallOpts)
}

// ApplyFunding is a paid mutator transaction binding the contract method 0x200f4b8d.
//
// Solidity: function applyFunding() returns()
func (_Squeeth *SqueethTransactor) ApplyFunding(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Squeeth.contract.Transact(opts, "applyFunding")
}

// ApplyFunding is a paid mutator transaction binding the contract method 0x200f4b8d.
//
// Solidity: function applyFunding() returns()
func (_Squeeth *SqueethSession) ApplyFunding() (*types.Transaction, error) {
	return _Squeeth.Contract.ApplyFunding(&_Squeeth.TransactOpts)
}

// ApplyFunding is a paid mutator transaction binding the contract method 0x200f4b8d.
//
// Solidity: function applyFunding() returns()
func (_Squeeth *SqueethTransactorSession) ApplyFunding() (*types.Transaction, error) {
	return _Squeeth.Contract.ApplyFunding(&_Squeeth.TransactOpts)
}

// BurnPowerPerpAmount is a paid mutator transaction binding the contract method 0x4394318d.
//
// Solidity: function burnPowerPerpAmount(uint256 _vaultId, uint256 _powerPerpAmount, uint256 _withdrawAmount) returns(uint256)
func (_Squeeth *SqueethTransactor) BurnPowerPerpAmount(opts *bind.TransactOpts, _vaultId *big.Int, _powerPerpAmount *big.Int, _withdrawAmount *big.Int) (*types.Transaction, error) {
	return _Squeeth.contract.Transact(opts, "burnPowerPerpAmount", _vaultId, _powerPerpAmount, _withdrawAmount)
}

// BurnPowerPerpAmount is a paid mutator transaction binding the contract method 0x4394318d.
//
// Solidity: function burnPowerPerpAmount(uint256 _vaultId, uint256 _powerPerpAmount, uint256 _withdrawAmount) returns(uint256)
func (_Squeeth *SqueethSession) BurnPowerPerpAmount(_vaultId *big.Int, _powerPerpAmount *big.Int, _withdrawAmount *big.Int) (*types.Transaction, error) {
	return _Squeeth.Contract.BurnPowerPerpAmount(&_Squeeth.TransactOpts, _vaultId, _powerPerpAmount, _withdrawAmount)
}

// BurnPowerPerpAmount is a paid mutator transaction binding the contract method 0x4394318d.
//
// Solidity: function burnPowerPerpAmount(uint256 _vaultId, uint256 _powerPerpAmount, uint256 _withdrawAmount) returns(uint256)
func (_Squeeth *SqueethTransactorSession) BurnPowerPerpAmount(_vaultId *big.Int, _powerPerpAmount *big.Int, _withdrawAmount *big.Int) (*types.Transaction, error) {
	return _Squeeth.Contract.BurnPowerPerpAmount(&_Squeeth.TransactOpts, _vaultId, _powerPerpAmount, _withdrawAmount)
}

// BurnWPowerPerpAmount is a paid mutator transaction binding the contract method 0x8632cb03.
//
// Solidity: function burnWPowerPerpAmount(uint256 _vaultId, uint256 _wPowerPerpAmount, uint256 _withdrawAmount) returns()
func (_Squeeth *SqueethTransactor) BurnWPowerPerpAmount(opts *bind.TransactOpts, _vaultId *big.Int, _wPowerPerpAmount *big.Int, _withdrawAmount *big.Int) (*types.Transaction, error) {
	return _Squeeth.contract.Transact(opts, "burnWPowerPerpAmount", _vaultId, _wPowerPerpAmount, _withdrawAmount)
}

// BurnWPowerPerpAmount is a paid mutator transaction binding the contract method 0x8632cb03.
//
// Solidity: function burnWPowerPerpAmount(uint256 _vaultId, uint256 _wPowerPerpAmount, uint256 _withdrawAmount) returns()
func (_Squeeth *SqueethSession) BurnWPowerPerpAmount(_vaultId *big.Int, _wPowerPerpAmount *big.Int, _withdrawAmount *big.Int) (*types.Transaction, error) {
	return _Squeeth.Contract.BurnWPowerPerpAmount(&_Squeeth.TransactOpts, _vaultId, _wPowerPerpAmount, _withdrawAmount)
}

// BurnWPowerPerpAmount is a paid mutator transaction binding the contract method 0x8632cb03.
//
// Solidity: function burnWPowerPerpAmount(uint256 _vaultId, uint256 _wPowerPerpAmount, uint256 _withdrawAmount) returns()
func (_Squeeth *SqueethTransactorSession) BurnWPowerPerpAmount(_vaultId *big.Int, _wPowerPerpAmount *big.Int, _withdrawAmount *big.Int) (*types.Transaction, error) {
	return _Squeeth.Contract.BurnWPowerPerpAmount(&_Squeeth.TransactOpts, _vaultId, _wPowerPerpAmount, _withdrawAmount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _vaultId) payable returns()
func (_Squeeth *SqueethTransactor) Deposit(opts *bind.TransactOpts, _vaultId *big.Int) (*types.Transaction, error) {
	return _Squeeth.contract.Transact(opts, "deposit", _vaultId)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _vaultId) payable returns()
func (_Squeeth *SqueethSession) Deposit(_vaultId *big.Int) (*types.Transaction, error) {
	return _Squeeth.Contract.Deposit(&_Squeeth.TransactOpts, _vaultId)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _vaultId) payable returns()
func (_Squeeth *SqueethTransactorSession) Deposit(_vaultId *big.Int) (*types.Transaction, error) {
	return _Squeeth.Contract.Deposit(&_Squeeth.TransactOpts, _vaultId)
}

// DepositUniPositionToken is a paid mutator transaction binding the contract method 0x91b8d34a.
//
// Solidity: function depositUniPositionToken(uint256 _vaultId, uint256 _uniTokenId) returns()
func (_Squeeth *SqueethTransactor) DepositUniPositionToken(opts *bind.TransactOpts, _vaultId *big.Int, _uniTokenId *big.Int) (*types.Transaction, error) {
	return _Squeeth.contract.Transact(opts, "depositUniPositionToken", _vaultId, _uniTokenId)
}

// DepositUniPositionToken is a paid mutator transaction binding the contract method 0x91b8d34a.
//
// Solidity: function depositUniPositionToken(uint256 _vaultId, uint256 _uniTokenId) returns()
func (_Squeeth *SqueethSession) DepositUniPositionToken(_vaultId *big.Int, _uniTokenId *big.Int) (*types.Transaction, error) {
	return _Squeeth.Contract.DepositUniPositionToken(&_Squeeth.TransactOpts, _vaultId, _uniTokenId)
}

// DepositUniPositionToken is a paid mutator transaction binding the contract method 0x91b8d34a.
//
// Solidity: function depositUniPositionToken(uint256 _vaultId, uint256 _uniTokenId) returns()
func (_Squeeth *SqueethTransactorSession) DepositUniPositionToken(_vaultId *big.Int, _uniTokenId *big.Int) (*types.Transaction, error) {
	return _Squeeth.Contract.DepositUniPositionToken(&_Squeeth.TransactOpts, _vaultId, _uniTokenId)
}

// Donate is a paid mutator transaction binding the contract method 0xed88c68e.
//
// Solidity: function donate() payable returns()
func (_Squeeth *SqueethTransactor) Donate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Squeeth.contract.Transact(opts, "donate")
}

// Donate is a paid mutator transaction binding the contract method 0xed88c68e.
//
// Solidity: function donate() payable returns()
func (_Squeeth *SqueethSession) Donate() (*types.Transaction, error) {
	return _Squeeth.Contract.Donate(&_Squeeth.TransactOpts)
}

// Donate is a paid mutator transaction binding the contract method 0xed88c68e.
//
// Solidity: function donate() payable returns()
func (_Squeeth *SqueethTransactorSession) Donate() (*types.Transaction, error) {
	return _Squeeth.Contract.Donate(&_Squeeth.TransactOpts)
}

// Liquidate is a paid mutator transaction binding the contract method 0xd296d1f1.
//
// Solidity: function liquidate(uint256 _vaultId, uint256 _maxDebtAmount) returns(uint256)
func (_Squeeth *SqueethTransactor) Liquidate(opts *bind.TransactOpts, _vaultId *big.Int, _maxDebtAmount *big.Int) (*types.Transaction, error) {
	return _Squeeth.contract.Transact(opts, "liquidate", _vaultId, _maxDebtAmount)
}

// Liquidate is a paid mutator transaction binding the contract method 0xd296d1f1.
//
// Solidity: function liquidate(uint256 _vaultId, uint256 _maxDebtAmount) returns(uint256)
func (_Squeeth *SqueethSession) Liquidate(_vaultId *big.Int, _maxDebtAmount *big.Int) (*types.Transaction, error) {
	return _Squeeth.Contract.Liquidate(&_Squeeth.TransactOpts, _vaultId, _maxDebtAmount)
}

// Liquidate is a paid mutator transaction binding the contract method 0xd296d1f1.
//
// Solidity: function liquidate(uint256 _vaultId, uint256 _maxDebtAmount) returns(uint256)
func (_Squeeth *SqueethTransactorSession) Liquidate(_vaultId *big.Int, _maxDebtAmount *big.Int) (*types.Transaction, error) {
	return _Squeeth.Contract.Liquidate(&_Squeeth.TransactOpts, _vaultId, _maxDebtAmount)
}

// MintPowerPerpAmount is a paid mutator transaction binding the contract method 0x1bf7bf6c.
//
// Solidity: function mintPowerPerpAmount(uint256 _vaultId, uint256 _powerPerpAmount, uint256 _uniTokenId) payable returns(uint256, uint256)
func (_Squeeth *SqueethTransactor) MintPowerPerpAmount(opts *bind.TransactOpts, _vaultId *big.Int, _powerPerpAmount *big.Int, _uniTokenId *big.Int) (*types.Transaction, error) {
	return _Squeeth.contract.Transact(opts, "mintPowerPerpAmount", _vaultId, _powerPerpAmount, _uniTokenId)
}

// MintPowerPerpAmount is a paid mutator transaction binding the contract method 0x1bf7bf6c.
//
// Solidity: function mintPowerPerpAmount(uint256 _vaultId, uint256 _powerPerpAmount, uint256 _uniTokenId) payable returns(uint256, uint256)
func (_Squeeth *SqueethSession) MintPowerPerpAmount(_vaultId *big.Int, _powerPerpAmount *big.Int, _uniTokenId *big.Int) (*types.Transaction, error) {
	return _Squeeth.Contract.MintPowerPerpAmount(&_Squeeth.TransactOpts, _vaultId, _powerPerpAmount, _uniTokenId)
}

// MintPowerPerpAmount is a paid mutator transaction binding the contract method 0x1bf7bf6c.
//
// Solidity: function mintPowerPerpAmount(uint256 _vaultId, uint256 _powerPerpAmount, uint256 _uniTokenId) payable returns(uint256, uint256)
func (_Squeeth *SqueethTransactorSession) MintPowerPerpAmount(_vaultId *big.Int, _powerPerpAmount *big.Int, _uniTokenId *big.Int) (*types.Transaction, error) {
	return _Squeeth.Contract.MintPowerPerpAmount(&_Squeeth.TransactOpts, _vaultId, _powerPerpAmount, _uniTokenId)
}

// MintWPowerPerpAmount is a paid mutator transaction binding the contract method 0x39467918.
//
// Solidity: function mintWPowerPerpAmount(uint256 _vaultId, uint256 _wPowerPerpAmount, uint256 _uniTokenId) payable returns(uint256)
func (_Squeeth *SqueethTransactor) MintWPowerPerpAmount(opts *bind.TransactOpts, _vaultId *big.Int, _wPowerPerpAmount *big.Int, _uniTokenId *big.Int) (*types.Transaction, error) {
	return _Squeeth.contract.Transact(opts, "mintWPowerPerpAmount", _vaultId, _wPowerPerpAmount, _uniTokenId)
}

// MintWPowerPerpAmount is a paid mutator transaction binding the contract method 0x39467918.
//
// Solidity: function mintWPowerPerpAmount(uint256 _vaultId, uint256 _wPowerPerpAmount, uint256 _uniTokenId) payable returns(uint256)
func (_Squeeth *SqueethSession) MintWPowerPerpAmount(_vaultId *big.Int, _wPowerPerpAmount *big.Int, _uniTokenId *big.Int) (*types.Transaction, error) {
	return _Squeeth.Contract.MintWPowerPerpAmount(&_Squeeth.TransactOpts, _vaultId, _wPowerPerpAmount, _uniTokenId)
}

// MintWPowerPerpAmount is a paid mutator transaction binding the contract method 0x39467918.
//
// Solidity: function mintWPowerPerpAmount(uint256 _vaultId, uint256 _wPowerPerpAmount, uint256 _uniTokenId) payable returns(uint256)
func (_Squeeth *SqueethTransactorSession) MintWPowerPerpAmount(_vaultId *big.Int, _wPowerPerpAmount *big.Int, _uniTokenId *big.Int) (*types.Transaction, error) {
	return _Squeeth.Contract.MintWPowerPerpAmount(&_Squeeth.TransactOpts, _vaultId, _wPowerPerpAmount, _uniTokenId)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Squeeth *SqueethTransactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Squeeth.contract.Transact(opts, "onERC721Received", arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Squeeth *SqueethSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Squeeth.Contract.OnERC721Received(&_Squeeth.TransactOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Squeeth *SqueethTransactorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Squeeth.Contract.OnERC721Received(&_Squeeth.TransactOpts, arg0, arg1, arg2, arg3)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Squeeth *SqueethTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Squeeth.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Squeeth *SqueethSession) Pause() (*types.Transaction, error) {
	return _Squeeth.Contract.Pause(&_Squeeth.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Squeeth *SqueethTransactorSession) Pause() (*types.Transaction, error) {
	return _Squeeth.Contract.Pause(&_Squeeth.TransactOpts)
}

// RedeemLong is a paid mutator transaction binding the contract method 0xac6cd5ef.
//
// Solidity: function redeemLong(uint256 _wPerpAmount) returns()
func (_Squeeth *SqueethTransactor) RedeemLong(opts *bind.TransactOpts, _wPerpAmount *big.Int) (*types.Transaction, error) {
	return _Squeeth.contract.Transact(opts, "redeemLong", _wPerpAmount)
}

// RedeemLong is a paid mutator transaction binding the contract method 0xac6cd5ef.
//
// Solidity: function redeemLong(uint256 _wPerpAmount) returns()
func (_Squeeth *SqueethSession) RedeemLong(_wPerpAmount *big.Int) (*types.Transaction, error) {
	return _Squeeth.Contract.RedeemLong(&_Squeeth.TransactOpts, _wPerpAmount)
}

// RedeemLong is a paid mutator transaction binding the contract method 0xac6cd5ef.
//
// Solidity: function redeemLong(uint256 _wPerpAmount) returns()
func (_Squeeth *SqueethTransactorSession) RedeemLong(_wPerpAmount *big.Int) (*types.Transaction, error) {
	return _Squeeth.Contract.RedeemLong(&_Squeeth.TransactOpts, _wPerpAmount)
}

// RedeemShort is a paid mutator transaction binding the contract method 0x97efa942.
//
// Solidity: function redeemShort(uint256 _vaultId) returns()
func (_Squeeth *SqueethTransactor) RedeemShort(opts *bind.TransactOpts, _vaultId *big.Int) (*types.Transaction, error) {
	return _Squeeth.contract.Transact(opts, "redeemShort", _vaultId)
}

// RedeemShort is a paid mutator transaction binding the contract method 0x97efa942.
//
// Solidity: function redeemShort(uint256 _vaultId) returns()
func (_Squeeth *SqueethSession) RedeemShort(_vaultId *big.Int) (*types.Transaction, error) {
	return _Squeeth.Contract.RedeemShort(&_Squeeth.TransactOpts, _vaultId)
}

// RedeemShort is a paid mutator transaction binding the contract method 0x97efa942.
//
// Solidity: function redeemShort(uint256 _vaultId) returns()
func (_Squeeth *SqueethTransactorSession) RedeemShort(_vaultId *big.Int) (*types.Transaction, error) {
	return _Squeeth.Contract.RedeemShort(&_Squeeth.TransactOpts, _vaultId)
}

// ReduceDebt is a paid mutator transaction binding the contract method 0xc9e77ee8.
//
// Solidity: function reduceDebt(uint256 _vaultId) returns()
func (_Squeeth *SqueethTransactor) ReduceDebt(opts *bind.TransactOpts, _vaultId *big.Int) (*types.Transaction, error) {
	return _Squeeth.contract.Transact(opts, "reduceDebt", _vaultId)
}

// ReduceDebt is a paid mutator transaction binding the contract method 0xc9e77ee8.
//
// Solidity: function reduceDebt(uint256 _vaultId) returns()
func (_Squeeth *SqueethSession) ReduceDebt(_vaultId *big.Int) (*types.Transaction, error) {
	return _Squeeth.Contract.ReduceDebt(&_Squeeth.TransactOpts, _vaultId)
}

// ReduceDebt is a paid mutator transaction binding the contract method 0xc9e77ee8.
//
// Solidity: function reduceDebt(uint256 _vaultId) returns()
func (_Squeeth *SqueethTransactorSession) ReduceDebt(_vaultId *big.Int) (*types.Transaction, error) {
	return _Squeeth.Contract.ReduceDebt(&_Squeeth.TransactOpts, _vaultId)
}

// ReduceDebtShutdown is a paid mutator transaction binding the contract method 0xfbfc6bc0.
//
// Solidity: function reduceDebtShutdown(uint256 _vaultId) returns()
func (_Squeeth *SqueethTransactor) ReduceDebtShutdown(opts *bind.TransactOpts, _vaultId *big.Int) (*types.Transaction, error) {
	return _Squeeth.contract.Transact(opts, "reduceDebtShutdown", _vaultId)
}

// ReduceDebtShutdown is a paid mutator transaction binding the contract method 0xfbfc6bc0.
//
// Solidity: function reduceDebtShutdown(uint256 _vaultId) returns()
func (_Squeeth *SqueethSession) ReduceDebtShutdown(_vaultId *big.Int) (*types.Transaction, error) {
	return _Squeeth.Contract.ReduceDebtShutdown(&_Squeeth.TransactOpts, _vaultId)
}

// ReduceDebtShutdown is a paid mutator transaction binding the contract method 0xfbfc6bc0.
//
// Solidity: function reduceDebtShutdown(uint256 _vaultId) returns()
func (_Squeeth *SqueethTransactorSession) ReduceDebtShutdown(_vaultId *big.Int) (*types.Transaction, error) {
	return _Squeeth.Contract.ReduceDebtShutdown(&_Squeeth.TransactOpts, _vaultId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Squeeth *SqueethTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Squeeth.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Squeeth *SqueethSession) RenounceOwnership() (*types.Transaction, error) {
	return _Squeeth.Contract.RenounceOwnership(&_Squeeth.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Squeeth *SqueethTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Squeeth.Contract.RenounceOwnership(&_Squeeth.TransactOpts)
}

// SetFeeRate is a paid mutator transaction binding the contract method 0x45596e2e.
//
// Solidity: function setFeeRate(uint256 _newFeeRate) returns()
func (_Squeeth *SqueethTransactor) SetFeeRate(opts *bind.TransactOpts, _newFeeRate *big.Int) (*types.Transaction, error) {
	return _Squeeth.contract.Transact(opts, "setFeeRate", _newFeeRate)
}

// SetFeeRate is a paid mutator transaction binding the contract method 0x45596e2e.
//
// Solidity: function setFeeRate(uint256 _newFeeRate) returns()
func (_Squeeth *SqueethSession) SetFeeRate(_newFeeRate *big.Int) (*types.Transaction, error) {
	return _Squeeth.Contract.SetFeeRate(&_Squeeth.TransactOpts, _newFeeRate)
}

// SetFeeRate is a paid mutator transaction binding the contract method 0x45596e2e.
//
// Solidity: function setFeeRate(uint256 _newFeeRate) returns()
func (_Squeeth *SqueethTransactorSession) SetFeeRate(_newFeeRate *big.Int) (*types.Transaction, error) {
	return _Squeeth.Contract.SetFeeRate(&_Squeeth.TransactOpts, _newFeeRate)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address _newFeeRecipient) returns()
func (_Squeeth *SqueethTransactor) SetFeeRecipient(opts *bind.TransactOpts, _newFeeRecipient common.Address) (*types.Transaction, error) {
	return _Squeeth.contract.Transact(opts, "setFeeRecipient", _newFeeRecipient)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address _newFeeRecipient) returns()
func (_Squeeth *SqueethSession) SetFeeRecipient(_newFeeRecipient common.Address) (*types.Transaction, error) {
	return _Squeeth.Contract.SetFeeRecipient(&_Squeeth.TransactOpts, _newFeeRecipient)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address _newFeeRecipient) returns()
func (_Squeeth *SqueethTransactorSession) SetFeeRecipient(_newFeeRecipient common.Address) (*types.Transaction, error) {
	return _Squeeth.Contract.SetFeeRecipient(&_Squeeth.TransactOpts, _newFeeRecipient)
}

// ShutDown is a paid mutator transaction binding the contract method 0x10b9e583.
//
// Solidity: function shutDown() returns()
func (_Squeeth *SqueethTransactor) ShutDown(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Squeeth.contract.Transact(opts, "shutDown")
}

// ShutDown is a paid mutator transaction binding the contract method 0x10b9e583.
//
// Solidity: function shutDown() returns()
func (_Squeeth *SqueethSession) ShutDown() (*types.Transaction, error) {
	return _Squeeth.Contract.ShutDown(&_Squeeth.TransactOpts)
}

// ShutDown is a paid mutator transaction binding the contract method 0x10b9e583.
//
// Solidity: function shutDown() returns()
func (_Squeeth *SqueethTransactorSession) ShutDown() (*types.Transaction, error) {
	return _Squeeth.Contract.ShutDown(&_Squeeth.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Squeeth *SqueethTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Squeeth.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Squeeth *SqueethSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Squeeth.Contract.TransferOwnership(&_Squeeth.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Squeeth *SqueethTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Squeeth.Contract.TransferOwnership(&_Squeeth.TransactOpts, newOwner)
}

// UnPauseAnyone is a paid mutator transaction binding the contract method 0x8146b09f.
//
// Solidity: function unPauseAnyone() returns()
func (_Squeeth *SqueethTransactor) UnPauseAnyone(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Squeeth.contract.Transact(opts, "unPauseAnyone")
}

// UnPauseAnyone is a paid mutator transaction binding the contract method 0x8146b09f.
//
// Solidity: function unPauseAnyone() returns()
func (_Squeeth *SqueethSession) UnPauseAnyone() (*types.Transaction, error) {
	return _Squeeth.Contract.UnPauseAnyone(&_Squeeth.TransactOpts)
}

// UnPauseAnyone is a paid mutator transaction binding the contract method 0x8146b09f.
//
// Solidity: function unPauseAnyone() returns()
func (_Squeeth *SqueethTransactorSession) UnPauseAnyone() (*types.Transaction, error) {
	return _Squeeth.Contract.UnPauseAnyone(&_Squeeth.TransactOpts)
}

// UnPauseOwner is a paid mutator transaction binding the contract method 0x07633669.
//
// Solidity: function unPauseOwner() returns()
func (_Squeeth *SqueethTransactor) UnPauseOwner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Squeeth.contract.Transact(opts, "unPauseOwner")
}

// UnPauseOwner is a paid mutator transaction binding the contract method 0x07633669.
//
// Solidity: function unPauseOwner() returns()
func (_Squeeth *SqueethSession) UnPauseOwner() (*types.Transaction, error) {
	return _Squeeth.Contract.UnPauseOwner(&_Squeeth.TransactOpts)
}

// UnPauseOwner is a paid mutator transaction binding the contract method 0x07633669.
//
// Solidity: function unPauseOwner() returns()
func (_Squeeth *SqueethTransactorSession) UnPauseOwner() (*types.Transaction, error) {
	return _Squeeth.Contract.UnPauseOwner(&_Squeeth.TransactOpts)
}

// UpdateOperator is a paid mutator transaction binding the contract method 0xc65a391d.
//
// Solidity: function updateOperator(uint256 _vaultId, address _operator) returns()
func (_Squeeth *SqueethTransactor) UpdateOperator(opts *bind.TransactOpts, _vaultId *big.Int, _operator common.Address) (*types.Transaction, error) {
	return _Squeeth.contract.Transact(opts, "updateOperator", _vaultId, _operator)
}

// UpdateOperator is a paid mutator transaction binding the contract method 0xc65a391d.
//
// Solidity: function updateOperator(uint256 _vaultId, address _operator) returns()
func (_Squeeth *SqueethSession) UpdateOperator(_vaultId *big.Int, _operator common.Address) (*types.Transaction, error) {
	return _Squeeth.Contract.UpdateOperator(&_Squeeth.TransactOpts, _vaultId, _operator)
}

// UpdateOperator is a paid mutator transaction binding the contract method 0xc65a391d.
//
// Solidity: function updateOperator(uint256 _vaultId, address _operator) returns()
func (_Squeeth *SqueethTransactorSession) UpdateOperator(_vaultId *big.Int, _operator common.Address) (*types.Transaction, error) {
	return _Squeeth.Contract.UpdateOperator(&_Squeeth.TransactOpts, _vaultId, _operator)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _vaultId, uint256 _amount) returns()
func (_Squeeth *SqueethTransactor) Withdraw(opts *bind.TransactOpts, _vaultId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Squeeth.contract.Transact(opts, "withdraw", _vaultId, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _vaultId, uint256 _amount) returns()
func (_Squeeth *SqueethSession) Withdraw(_vaultId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Squeeth.Contract.Withdraw(&_Squeeth.TransactOpts, _vaultId, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _vaultId, uint256 _amount) returns()
func (_Squeeth *SqueethTransactorSession) Withdraw(_vaultId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Squeeth.Contract.Withdraw(&_Squeeth.TransactOpts, _vaultId, _amount)
}

// WithdrawUniPositionToken is a paid mutator transaction binding the contract method 0x713d517f.
//
// Solidity: function withdrawUniPositionToken(uint256 _vaultId) returns()
func (_Squeeth *SqueethTransactor) WithdrawUniPositionToken(opts *bind.TransactOpts, _vaultId *big.Int) (*types.Transaction, error) {
	return _Squeeth.contract.Transact(opts, "withdrawUniPositionToken", _vaultId)
}

// WithdrawUniPositionToken is a paid mutator transaction binding the contract method 0x713d517f.
//
// Solidity: function withdrawUniPositionToken(uint256 _vaultId) returns()
func (_Squeeth *SqueethSession) WithdrawUniPositionToken(_vaultId *big.Int) (*types.Transaction, error) {
	return _Squeeth.Contract.WithdrawUniPositionToken(&_Squeeth.TransactOpts, _vaultId)
}

// WithdrawUniPositionToken is a paid mutator transaction binding the contract method 0x713d517f.
//
// Solidity: function withdrawUniPositionToken(uint256 _vaultId) returns()
func (_Squeeth *SqueethTransactorSession) WithdrawUniPositionToken(_vaultId *big.Int) (*types.Transaction, error) {
	return _Squeeth.Contract.WithdrawUniPositionToken(&_Squeeth.TransactOpts, _vaultId)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Squeeth *SqueethTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Squeeth.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Squeeth *SqueethSession) Receive() (*types.Transaction, error) {
	return _Squeeth.Contract.Receive(&_Squeeth.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Squeeth *SqueethTransactorSession) Receive() (*types.Transaction, error) {
	return _Squeeth.Contract.Receive(&_Squeeth.TransactOpts)
}

// SqueethBurnShortIterator is returned from FilterBurnShort and is used to iterate over the raw logs and unpacked data for BurnShort events raised by the Squeeth contract.
type SqueethBurnShortIterator struct {
	Event *SqueethBurnShort // Event containing the contract specifics and raw log

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
func (it *SqueethBurnShortIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SqueethBurnShort)
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
		it.Event = new(SqueethBurnShort)
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
func (it *SqueethBurnShortIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SqueethBurnShortIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SqueethBurnShort represents a BurnShort event raised by the Squeeth contract.
type SqueethBurnShort struct {
	Sender  common.Address
	Amount  *big.Int
	VaultId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBurnShort is a free log retrieval operation binding the contract event 0xea19ffc45b48de6d95594aacff7214dd24595fdb0c60e98c8f0b269058c2f792.
//
// Solidity: event BurnShort(address sender, uint256 amount, uint256 vaultId)
func (_Squeeth *SqueethFilterer) FilterBurnShort(opts *bind.FilterOpts) (*SqueethBurnShortIterator, error) {

	logs, sub, err := _Squeeth.contract.FilterLogs(opts, "BurnShort")
	if err != nil {
		return nil, err
	}
	return &SqueethBurnShortIterator{contract: _Squeeth.contract, event: "BurnShort", logs: logs, sub: sub}, nil
}

// WatchBurnShort is a free log subscription operation binding the contract event 0xea19ffc45b48de6d95594aacff7214dd24595fdb0c60e98c8f0b269058c2f792.
//
// Solidity: event BurnShort(address sender, uint256 amount, uint256 vaultId)
func (_Squeeth *SqueethFilterer) WatchBurnShort(opts *bind.WatchOpts, sink chan<- *SqueethBurnShort) (event.Subscription, error) {

	logs, sub, err := _Squeeth.contract.WatchLogs(opts, "BurnShort")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SqueethBurnShort)
				if err := _Squeeth.contract.UnpackLog(event, "BurnShort", log); err != nil {
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

// ParseBurnShort is a log parse operation binding the contract event 0xea19ffc45b48de6d95594aacff7214dd24595fdb0c60e98c8f0b269058c2f792.
//
// Solidity: event BurnShort(address sender, uint256 amount, uint256 vaultId)
func (_Squeeth *SqueethFilterer) ParseBurnShort(log types.Log) (*SqueethBurnShort, error) {
	event := new(SqueethBurnShort)
	if err := _Squeeth.contract.UnpackLog(event, "BurnShort", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SqueethDepositCollateralIterator is returned from FilterDepositCollateral and is used to iterate over the raw logs and unpacked data for DepositCollateral events raised by the Squeeth contract.
type SqueethDepositCollateralIterator struct {
	Event *SqueethDepositCollateral // Event containing the contract specifics and raw log

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
func (it *SqueethDepositCollateralIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SqueethDepositCollateral)
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
		it.Event = new(SqueethDepositCollateral)
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
func (it *SqueethDepositCollateralIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SqueethDepositCollateralIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SqueethDepositCollateral represents a DepositCollateral event raised by the Squeeth contract.
type SqueethDepositCollateral struct {
	Sender  common.Address
	VaultId *big.Int
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDepositCollateral is a free log retrieval operation binding the contract event 0x3ca13b7aab12bad7472691fe558faa6b25e99099824a0070a88bd5aa84be610f.
//
// Solidity: event DepositCollateral(address sender, uint256 vaultId, uint256 amount)
func (_Squeeth *SqueethFilterer) FilterDepositCollateral(opts *bind.FilterOpts) (*SqueethDepositCollateralIterator, error) {

	logs, sub, err := _Squeeth.contract.FilterLogs(opts, "DepositCollateral")
	if err != nil {
		return nil, err
	}
	return &SqueethDepositCollateralIterator{contract: _Squeeth.contract, event: "DepositCollateral", logs: logs, sub: sub}, nil
}

// WatchDepositCollateral is a free log subscription operation binding the contract event 0x3ca13b7aab12bad7472691fe558faa6b25e99099824a0070a88bd5aa84be610f.
//
// Solidity: event DepositCollateral(address sender, uint256 vaultId, uint256 amount)
func (_Squeeth *SqueethFilterer) WatchDepositCollateral(opts *bind.WatchOpts, sink chan<- *SqueethDepositCollateral) (event.Subscription, error) {

	logs, sub, err := _Squeeth.contract.WatchLogs(opts, "DepositCollateral")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SqueethDepositCollateral)
				if err := _Squeeth.contract.UnpackLog(event, "DepositCollateral", log); err != nil {
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

// ParseDepositCollateral is a log parse operation binding the contract event 0x3ca13b7aab12bad7472691fe558faa6b25e99099824a0070a88bd5aa84be610f.
//
// Solidity: event DepositCollateral(address sender, uint256 vaultId, uint256 amount)
func (_Squeeth *SqueethFilterer) ParseDepositCollateral(log types.Log) (*SqueethDepositCollateral, error) {
	event := new(SqueethDepositCollateral)
	if err := _Squeeth.contract.UnpackLog(event, "DepositCollateral", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SqueethDepositUniPositionTokenIterator is returned from FilterDepositUniPositionToken and is used to iterate over the raw logs and unpacked data for DepositUniPositionToken events raised by the Squeeth contract.
type SqueethDepositUniPositionTokenIterator struct {
	Event *SqueethDepositUniPositionToken // Event containing the contract specifics and raw log

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
func (it *SqueethDepositUniPositionTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SqueethDepositUniPositionToken)
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
		it.Event = new(SqueethDepositUniPositionToken)
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
func (it *SqueethDepositUniPositionTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SqueethDepositUniPositionTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SqueethDepositUniPositionToken represents a DepositUniPositionToken event raised by the Squeeth contract.
type SqueethDepositUniPositionToken struct {
	Sender  common.Address
	VaultId *big.Int
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDepositUniPositionToken is a free log retrieval operation binding the contract event 0x3917c2f26ce18614e3aedd1289da672ef6563c5c295f49e9b1697ae0ad315562.
//
// Solidity: event DepositUniPositionToken(address sender, uint256 vaultId, uint256 tokenId)
func (_Squeeth *SqueethFilterer) FilterDepositUniPositionToken(opts *bind.FilterOpts) (*SqueethDepositUniPositionTokenIterator, error) {

	logs, sub, err := _Squeeth.contract.FilterLogs(opts, "DepositUniPositionToken")
	if err != nil {
		return nil, err
	}
	return &SqueethDepositUniPositionTokenIterator{contract: _Squeeth.contract, event: "DepositUniPositionToken", logs: logs, sub: sub}, nil
}

// WatchDepositUniPositionToken is a free log subscription operation binding the contract event 0x3917c2f26ce18614e3aedd1289da672ef6563c5c295f49e9b1697ae0ad315562.
//
// Solidity: event DepositUniPositionToken(address sender, uint256 vaultId, uint256 tokenId)
func (_Squeeth *SqueethFilterer) WatchDepositUniPositionToken(opts *bind.WatchOpts, sink chan<- *SqueethDepositUniPositionToken) (event.Subscription, error) {

	logs, sub, err := _Squeeth.contract.WatchLogs(opts, "DepositUniPositionToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SqueethDepositUniPositionToken)
				if err := _Squeeth.contract.UnpackLog(event, "DepositUniPositionToken", log); err != nil {
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

// ParseDepositUniPositionToken is a log parse operation binding the contract event 0x3917c2f26ce18614e3aedd1289da672ef6563c5c295f49e9b1697ae0ad315562.
//
// Solidity: event DepositUniPositionToken(address sender, uint256 vaultId, uint256 tokenId)
func (_Squeeth *SqueethFilterer) ParseDepositUniPositionToken(log types.Log) (*SqueethDepositUniPositionToken, error) {
	event := new(SqueethDepositUniPositionToken)
	if err := _Squeeth.contract.UnpackLog(event, "DepositUniPositionToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SqueethFeeRateUpdatedIterator is returned from FilterFeeRateUpdated and is used to iterate over the raw logs and unpacked data for FeeRateUpdated events raised by the Squeeth contract.
type SqueethFeeRateUpdatedIterator struct {
	Event *SqueethFeeRateUpdated // Event containing the contract specifics and raw log

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
func (it *SqueethFeeRateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SqueethFeeRateUpdated)
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
		it.Event = new(SqueethFeeRateUpdated)
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
func (it *SqueethFeeRateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SqueethFeeRateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SqueethFeeRateUpdated represents a FeeRateUpdated event raised by the Squeeth contract.
type SqueethFeeRateUpdated struct {
	OldFee *big.Int
	NewFee *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFeeRateUpdated is a free log retrieval operation binding the contract event 0x14914da2bf76024616fbe1859783fcd4dbddcb179b1f3a854949fbf920dcb957.
//
// Solidity: event FeeRateUpdated(uint256 oldFee, uint256 newFee)
func (_Squeeth *SqueethFilterer) FilterFeeRateUpdated(opts *bind.FilterOpts) (*SqueethFeeRateUpdatedIterator, error) {

	logs, sub, err := _Squeeth.contract.FilterLogs(opts, "FeeRateUpdated")
	if err != nil {
		return nil, err
	}
	return &SqueethFeeRateUpdatedIterator{contract: _Squeeth.contract, event: "FeeRateUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeRateUpdated is a free log subscription operation binding the contract event 0x14914da2bf76024616fbe1859783fcd4dbddcb179b1f3a854949fbf920dcb957.
//
// Solidity: event FeeRateUpdated(uint256 oldFee, uint256 newFee)
func (_Squeeth *SqueethFilterer) WatchFeeRateUpdated(opts *bind.WatchOpts, sink chan<- *SqueethFeeRateUpdated) (event.Subscription, error) {

	logs, sub, err := _Squeeth.contract.WatchLogs(opts, "FeeRateUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SqueethFeeRateUpdated)
				if err := _Squeeth.contract.UnpackLog(event, "FeeRateUpdated", log); err != nil {
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

// ParseFeeRateUpdated is a log parse operation binding the contract event 0x14914da2bf76024616fbe1859783fcd4dbddcb179b1f3a854949fbf920dcb957.
//
// Solidity: event FeeRateUpdated(uint256 oldFee, uint256 newFee)
func (_Squeeth *SqueethFilterer) ParseFeeRateUpdated(log types.Log) (*SqueethFeeRateUpdated, error) {
	event := new(SqueethFeeRateUpdated)
	if err := _Squeeth.contract.UnpackLog(event, "FeeRateUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SqueethFeeRecipientUpdatedIterator is returned from FilterFeeRecipientUpdated and is used to iterate over the raw logs and unpacked data for FeeRecipientUpdated events raised by the Squeeth contract.
type SqueethFeeRecipientUpdatedIterator struct {
	Event *SqueethFeeRecipientUpdated // Event containing the contract specifics and raw log

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
func (it *SqueethFeeRecipientUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SqueethFeeRecipientUpdated)
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
		it.Event = new(SqueethFeeRecipientUpdated)
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
func (it *SqueethFeeRecipientUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SqueethFeeRecipientUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SqueethFeeRecipientUpdated represents a FeeRecipientUpdated event raised by the Squeeth contract.
type SqueethFeeRecipientUpdated struct {
	OldFeeRecipient common.Address
	NewFeeRecipient common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterFeeRecipientUpdated is a free log retrieval operation binding the contract event 0xaaebcf1bfa00580e41d966056b48521fa9f202645c86d4ddf28113e617c1b1d3.
//
// Solidity: event FeeRecipientUpdated(address oldFeeRecipient, address newFeeRecipient)
func (_Squeeth *SqueethFilterer) FilterFeeRecipientUpdated(opts *bind.FilterOpts) (*SqueethFeeRecipientUpdatedIterator, error) {

	logs, sub, err := _Squeeth.contract.FilterLogs(opts, "FeeRecipientUpdated")
	if err != nil {
		return nil, err
	}
	return &SqueethFeeRecipientUpdatedIterator{contract: _Squeeth.contract, event: "FeeRecipientUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeRecipientUpdated is a free log subscription operation binding the contract event 0xaaebcf1bfa00580e41d966056b48521fa9f202645c86d4ddf28113e617c1b1d3.
//
// Solidity: event FeeRecipientUpdated(address oldFeeRecipient, address newFeeRecipient)
func (_Squeeth *SqueethFilterer) WatchFeeRecipientUpdated(opts *bind.WatchOpts, sink chan<- *SqueethFeeRecipientUpdated) (event.Subscription, error) {

	logs, sub, err := _Squeeth.contract.WatchLogs(opts, "FeeRecipientUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SqueethFeeRecipientUpdated)
				if err := _Squeeth.contract.UnpackLog(event, "FeeRecipientUpdated", log); err != nil {
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

// ParseFeeRecipientUpdated is a log parse operation binding the contract event 0xaaebcf1bfa00580e41d966056b48521fa9f202645c86d4ddf28113e617c1b1d3.
//
// Solidity: event FeeRecipientUpdated(address oldFeeRecipient, address newFeeRecipient)
func (_Squeeth *SqueethFilterer) ParseFeeRecipientUpdated(log types.Log) (*SqueethFeeRecipientUpdated, error) {
	event := new(SqueethFeeRecipientUpdated)
	if err := _Squeeth.contract.UnpackLog(event, "FeeRecipientUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SqueethLiquidateIterator is returned from FilterLiquidate and is used to iterate over the raw logs and unpacked data for Liquidate events raised by the Squeeth contract.
type SqueethLiquidateIterator struct {
	Event *SqueethLiquidate // Event containing the contract specifics and raw log

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
func (it *SqueethLiquidateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SqueethLiquidate)
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
		it.Event = new(SqueethLiquidate)
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
func (it *SqueethLiquidateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SqueethLiquidateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SqueethLiquidate represents a Liquidate event raised by the Squeeth contract.
type SqueethLiquidate struct {
	Liquidator     common.Address
	VaultId        *big.Int
	DebtAmount     *big.Int
	CollateralPaid *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLiquidate is a free log retrieval operation binding the contract event 0x158ba9ab7bbbd08eeffa4753bad41f4d450e24831d293427308badf3eadd8c76.
//
// Solidity: event Liquidate(address liquidator, uint256 vaultId, uint256 debtAmount, uint256 collateralPaid)
func (_Squeeth *SqueethFilterer) FilterLiquidate(opts *bind.FilterOpts) (*SqueethLiquidateIterator, error) {

	logs, sub, err := _Squeeth.contract.FilterLogs(opts, "Liquidate")
	if err != nil {
		return nil, err
	}
	return &SqueethLiquidateIterator{contract: _Squeeth.contract, event: "Liquidate", logs: logs, sub: sub}, nil
}

// WatchLiquidate is a free log subscription operation binding the contract event 0x158ba9ab7bbbd08eeffa4753bad41f4d450e24831d293427308badf3eadd8c76.
//
// Solidity: event Liquidate(address liquidator, uint256 vaultId, uint256 debtAmount, uint256 collateralPaid)
func (_Squeeth *SqueethFilterer) WatchLiquidate(opts *bind.WatchOpts, sink chan<- *SqueethLiquidate) (event.Subscription, error) {

	logs, sub, err := _Squeeth.contract.WatchLogs(opts, "Liquidate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SqueethLiquidate)
				if err := _Squeeth.contract.UnpackLog(event, "Liquidate", log); err != nil {
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

// ParseLiquidate is a log parse operation binding the contract event 0x158ba9ab7bbbd08eeffa4753bad41f4d450e24831d293427308badf3eadd8c76.
//
// Solidity: event Liquidate(address liquidator, uint256 vaultId, uint256 debtAmount, uint256 collateralPaid)
func (_Squeeth *SqueethFilterer) ParseLiquidate(log types.Log) (*SqueethLiquidate, error) {
	event := new(SqueethLiquidate)
	if err := _Squeeth.contract.UnpackLog(event, "Liquidate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SqueethMintShortIterator is returned from FilterMintShort and is used to iterate over the raw logs and unpacked data for MintShort events raised by the Squeeth contract.
type SqueethMintShortIterator struct {
	Event *SqueethMintShort // Event containing the contract specifics and raw log

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
func (it *SqueethMintShortIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SqueethMintShort)
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
		it.Event = new(SqueethMintShort)
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
func (it *SqueethMintShortIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SqueethMintShortIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SqueethMintShort represents a MintShort event raised by the Squeeth contract.
type SqueethMintShort struct {
	Sender  common.Address
	Amount  *big.Int
	VaultId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMintShort is a free log retrieval operation binding the contract event 0xb19fa182730a088464dad0e9e0badeb470d0d8d937d854f5caf15c6ad1992c36.
//
// Solidity: event MintShort(address sender, uint256 amount, uint256 vaultId)
func (_Squeeth *SqueethFilterer) FilterMintShort(opts *bind.FilterOpts) (*SqueethMintShortIterator, error) {

	logs, sub, err := _Squeeth.contract.FilterLogs(opts, "MintShort")
	if err != nil {
		return nil, err
	}
	return &SqueethMintShortIterator{contract: _Squeeth.contract, event: "MintShort", logs: logs, sub: sub}, nil
}

// WatchMintShort is a free log subscription operation binding the contract event 0xb19fa182730a088464dad0e9e0badeb470d0d8d937d854f5caf15c6ad1992c36.
//
// Solidity: event MintShort(address sender, uint256 amount, uint256 vaultId)
func (_Squeeth *SqueethFilterer) WatchMintShort(opts *bind.WatchOpts, sink chan<- *SqueethMintShort) (event.Subscription, error) {

	logs, sub, err := _Squeeth.contract.WatchLogs(opts, "MintShort")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SqueethMintShort)
				if err := _Squeeth.contract.UnpackLog(event, "MintShort", log); err != nil {
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

// ParseMintShort is a log parse operation binding the contract event 0xb19fa182730a088464dad0e9e0badeb470d0d8d937d854f5caf15c6ad1992c36.
//
// Solidity: event MintShort(address sender, uint256 amount, uint256 vaultId)
func (_Squeeth *SqueethFilterer) ParseMintShort(log types.Log) (*SqueethMintShort, error) {
	event := new(SqueethMintShort)
	if err := _Squeeth.contract.UnpackLog(event, "MintShort", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SqueethNormalizationFactorUpdatedIterator is returned from FilterNormalizationFactorUpdated and is used to iterate over the raw logs and unpacked data for NormalizationFactorUpdated events raised by the Squeeth contract.
type SqueethNormalizationFactorUpdatedIterator struct {
	Event *SqueethNormalizationFactorUpdated // Event containing the contract specifics and raw log

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
func (it *SqueethNormalizationFactorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SqueethNormalizationFactorUpdated)
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
		it.Event = new(SqueethNormalizationFactorUpdated)
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
func (it *SqueethNormalizationFactorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SqueethNormalizationFactorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SqueethNormalizationFactorUpdated represents a NormalizationFactorUpdated event raised by the Squeeth contract.
type SqueethNormalizationFactorUpdated struct {
	OldNormFactor             *big.Int
	NewNormFactor             *big.Int
	LastModificationTimestamp *big.Int
	Timestamp                 *big.Int
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterNormalizationFactorUpdated is a free log retrieval operation binding the contract event 0x339e53729b0447795ff69e70a74fed98fc7fef6fe94b7521099b32f0f8de4822.
//
// Solidity: event NormalizationFactorUpdated(uint256 oldNormFactor, uint256 newNormFactor, uint256 lastModificationTimestamp, uint256 timestamp)
func (_Squeeth *SqueethFilterer) FilterNormalizationFactorUpdated(opts *bind.FilterOpts) (*SqueethNormalizationFactorUpdatedIterator, error) {

	logs, sub, err := _Squeeth.contract.FilterLogs(opts, "NormalizationFactorUpdated")
	if err != nil {
		return nil, err
	}
	return &SqueethNormalizationFactorUpdatedIterator{contract: _Squeeth.contract, event: "NormalizationFactorUpdated", logs: logs, sub: sub}, nil
}

// WatchNormalizationFactorUpdated is a free log subscription operation binding the contract event 0x339e53729b0447795ff69e70a74fed98fc7fef6fe94b7521099b32f0f8de4822.
//
// Solidity: event NormalizationFactorUpdated(uint256 oldNormFactor, uint256 newNormFactor, uint256 lastModificationTimestamp, uint256 timestamp)
func (_Squeeth *SqueethFilterer) WatchNormalizationFactorUpdated(opts *bind.WatchOpts, sink chan<- *SqueethNormalizationFactorUpdated) (event.Subscription, error) {

	logs, sub, err := _Squeeth.contract.WatchLogs(opts, "NormalizationFactorUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SqueethNormalizationFactorUpdated)
				if err := _Squeeth.contract.UnpackLog(event, "NormalizationFactorUpdated", log); err != nil {
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

// ParseNormalizationFactorUpdated is a log parse operation binding the contract event 0x339e53729b0447795ff69e70a74fed98fc7fef6fe94b7521099b32f0f8de4822.
//
// Solidity: event NormalizationFactorUpdated(uint256 oldNormFactor, uint256 newNormFactor, uint256 lastModificationTimestamp, uint256 timestamp)
func (_Squeeth *SqueethFilterer) ParseNormalizationFactorUpdated(log types.Log) (*SqueethNormalizationFactorUpdated, error) {
	event := new(SqueethNormalizationFactorUpdated)
	if err := _Squeeth.contract.UnpackLog(event, "NormalizationFactorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SqueethOpenVaultIterator is returned from FilterOpenVault and is used to iterate over the raw logs and unpacked data for OpenVault events raised by the Squeeth contract.
type SqueethOpenVaultIterator struct {
	Event *SqueethOpenVault // Event containing the contract specifics and raw log

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
func (it *SqueethOpenVaultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SqueethOpenVault)
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
		it.Event = new(SqueethOpenVault)
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
func (it *SqueethOpenVaultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SqueethOpenVaultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SqueethOpenVault represents a OpenVault event raised by the Squeeth contract.
type SqueethOpenVault struct {
	Sender  common.Address
	VaultId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOpenVault is a free log retrieval operation binding the contract event 0x25ff1e0131762176a9084e92880f880f39d6d0e62134607f37e631efe1bad871.
//
// Solidity: event OpenVault(address sender, uint256 vaultId)
func (_Squeeth *SqueethFilterer) FilterOpenVault(opts *bind.FilterOpts) (*SqueethOpenVaultIterator, error) {

	logs, sub, err := _Squeeth.contract.FilterLogs(opts, "OpenVault")
	if err != nil {
		return nil, err
	}
	return &SqueethOpenVaultIterator{contract: _Squeeth.contract, event: "OpenVault", logs: logs, sub: sub}, nil
}

// WatchOpenVault is a free log subscription operation binding the contract event 0x25ff1e0131762176a9084e92880f880f39d6d0e62134607f37e631efe1bad871.
//
// Solidity: event OpenVault(address sender, uint256 vaultId)
func (_Squeeth *SqueethFilterer) WatchOpenVault(opts *bind.WatchOpts, sink chan<- *SqueethOpenVault) (event.Subscription, error) {

	logs, sub, err := _Squeeth.contract.WatchLogs(opts, "OpenVault")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SqueethOpenVault)
				if err := _Squeeth.contract.UnpackLog(event, "OpenVault", log); err != nil {
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

// ParseOpenVault is a log parse operation binding the contract event 0x25ff1e0131762176a9084e92880f880f39d6d0e62134607f37e631efe1bad871.
//
// Solidity: event OpenVault(address sender, uint256 vaultId)
func (_Squeeth *SqueethFilterer) ParseOpenVault(log types.Log) (*SqueethOpenVault, error) {
	event := new(SqueethOpenVault)
	if err := _Squeeth.contract.UnpackLog(event, "OpenVault", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SqueethOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Squeeth contract.
type SqueethOwnershipTransferredIterator struct {
	Event *SqueethOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SqueethOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SqueethOwnershipTransferred)
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
		it.Event = new(SqueethOwnershipTransferred)
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
func (it *SqueethOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SqueethOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SqueethOwnershipTransferred represents a OwnershipTransferred event raised by the Squeeth contract.
type SqueethOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Squeeth *SqueethFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SqueethOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Squeeth.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SqueethOwnershipTransferredIterator{contract: _Squeeth.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Squeeth *SqueethFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SqueethOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Squeeth.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SqueethOwnershipTransferred)
				if err := _Squeeth.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Squeeth *SqueethFilterer) ParseOwnershipTransferred(log types.Log) (*SqueethOwnershipTransferred, error) {
	event := new(SqueethOwnershipTransferred)
	if err := _Squeeth.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SqueethPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Squeeth contract.
type SqueethPausedIterator struct {
	Event *SqueethPaused // Event containing the contract specifics and raw log

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
func (it *SqueethPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SqueethPaused)
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
		it.Event = new(SqueethPaused)
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
func (it *SqueethPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SqueethPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SqueethPaused represents a Paused event raised by the Squeeth contract.
type SqueethPaused struct {
	PausesLeft *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x32fb7c9891bc4f963c7de9f1186d2a7755c7d6e9f4604dabe1d8bb3027c2f49e.
//
// Solidity: event Paused(uint256 pausesLeft)
func (_Squeeth *SqueethFilterer) FilterPaused(opts *bind.FilterOpts) (*SqueethPausedIterator, error) {

	logs, sub, err := _Squeeth.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &SqueethPausedIterator{contract: _Squeeth.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x32fb7c9891bc4f963c7de9f1186d2a7755c7d6e9f4604dabe1d8bb3027c2f49e.
//
// Solidity: event Paused(uint256 pausesLeft)
func (_Squeeth *SqueethFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *SqueethPaused) (event.Subscription, error) {

	logs, sub, err := _Squeeth.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SqueethPaused)
				if err := _Squeeth.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x32fb7c9891bc4f963c7de9f1186d2a7755c7d6e9f4604dabe1d8bb3027c2f49e.
//
// Solidity: event Paused(uint256 pausesLeft)
func (_Squeeth *SqueethFilterer) ParsePaused(log types.Log) (*SqueethPaused, error) {
	event := new(SqueethPaused)
	if err := _Squeeth.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SqueethRedeemLongIterator is returned from FilterRedeemLong and is used to iterate over the raw logs and unpacked data for RedeemLong events raised by the Squeeth contract.
type SqueethRedeemLongIterator struct {
	Event *SqueethRedeemLong // Event containing the contract specifics and raw log

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
func (it *SqueethRedeemLongIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SqueethRedeemLong)
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
		it.Event = new(SqueethRedeemLong)
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
func (it *SqueethRedeemLongIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SqueethRedeemLongIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SqueethRedeemLong represents a RedeemLong event raised by the Squeeth contract.
type SqueethRedeemLong struct {
	Sender           common.Address
	WPowerPerpAmount *big.Int
	PayoutAmount     *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterRedeemLong is a free log retrieval operation binding the contract event 0x2131ef4f2f82ca75fe7d2e646ebfa45b6be25e53510c829629c76b641500ec67.
//
// Solidity: event RedeemLong(address sender, uint256 wPowerPerpAmount, uint256 payoutAmount)
func (_Squeeth *SqueethFilterer) FilterRedeemLong(opts *bind.FilterOpts) (*SqueethRedeemLongIterator, error) {

	logs, sub, err := _Squeeth.contract.FilterLogs(opts, "RedeemLong")
	if err != nil {
		return nil, err
	}
	return &SqueethRedeemLongIterator{contract: _Squeeth.contract, event: "RedeemLong", logs: logs, sub: sub}, nil
}

// WatchRedeemLong is a free log subscription operation binding the contract event 0x2131ef4f2f82ca75fe7d2e646ebfa45b6be25e53510c829629c76b641500ec67.
//
// Solidity: event RedeemLong(address sender, uint256 wPowerPerpAmount, uint256 payoutAmount)
func (_Squeeth *SqueethFilterer) WatchRedeemLong(opts *bind.WatchOpts, sink chan<- *SqueethRedeemLong) (event.Subscription, error) {

	logs, sub, err := _Squeeth.contract.WatchLogs(opts, "RedeemLong")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SqueethRedeemLong)
				if err := _Squeeth.contract.UnpackLog(event, "RedeemLong", log); err != nil {
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

// ParseRedeemLong is a log parse operation binding the contract event 0x2131ef4f2f82ca75fe7d2e646ebfa45b6be25e53510c829629c76b641500ec67.
//
// Solidity: event RedeemLong(address sender, uint256 wPowerPerpAmount, uint256 payoutAmount)
func (_Squeeth *SqueethFilterer) ParseRedeemLong(log types.Log) (*SqueethRedeemLong, error) {
	event := new(SqueethRedeemLong)
	if err := _Squeeth.contract.UnpackLog(event, "RedeemLong", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SqueethRedeemShortIterator is returned from FilterRedeemShort and is used to iterate over the raw logs and unpacked data for RedeemShort events raised by the Squeeth contract.
type SqueethRedeemShortIterator struct {
	Event *SqueethRedeemShort // Event containing the contract specifics and raw log

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
func (it *SqueethRedeemShortIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SqueethRedeemShort)
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
		it.Event = new(SqueethRedeemShort)
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
func (it *SqueethRedeemShortIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SqueethRedeemShortIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SqueethRedeemShort represents a RedeemShort event raised by the Squeeth contract.
type SqueethRedeemShort struct {
	Sender           common.Address
	VauldId          *big.Int
	CollateralAmount *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterRedeemShort is a free log retrieval operation binding the contract event 0x7dff8cdaec6a8d4d1ad32d3c947ed0f0281c3d6456621ef928defae96ec6cddb.
//
// Solidity: event RedeemShort(address sender, uint256 vauldId, uint256 collateralAmount)
func (_Squeeth *SqueethFilterer) FilterRedeemShort(opts *bind.FilterOpts) (*SqueethRedeemShortIterator, error) {

	logs, sub, err := _Squeeth.contract.FilterLogs(opts, "RedeemShort")
	if err != nil {
		return nil, err
	}
	return &SqueethRedeemShortIterator{contract: _Squeeth.contract, event: "RedeemShort", logs: logs, sub: sub}, nil
}

// WatchRedeemShort is a free log subscription operation binding the contract event 0x7dff8cdaec6a8d4d1ad32d3c947ed0f0281c3d6456621ef928defae96ec6cddb.
//
// Solidity: event RedeemShort(address sender, uint256 vauldId, uint256 collateralAmount)
func (_Squeeth *SqueethFilterer) WatchRedeemShort(opts *bind.WatchOpts, sink chan<- *SqueethRedeemShort) (event.Subscription, error) {

	logs, sub, err := _Squeeth.contract.WatchLogs(opts, "RedeemShort")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SqueethRedeemShort)
				if err := _Squeeth.contract.UnpackLog(event, "RedeemShort", log); err != nil {
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

// ParseRedeemShort is a log parse operation binding the contract event 0x7dff8cdaec6a8d4d1ad32d3c947ed0f0281c3d6456621ef928defae96ec6cddb.
//
// Solidity: event RedeemShort(address sender, uint256 vauldId, uint256 collateralAmount)
func (_Squeeth *SqueethFilterer) ParseRedeemShort(log types.Log) (*SqueethRedeemShort, error) {
	event := new(SqueethRedeemShort)
	if err := _Squeeth.contract.UnpackLog(event, "RedeemShort", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SqueethReduceDebtIterator is returned from FilterReduceDebt and is used to iterate over the raw logs and unpacked data for ReduceDebt events raised by the Squeeth contract.
type SqueethReduceDebtIterator struct {
	Event *SqueethReduceDebt // Event containing the contract specifics and raw log

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
func (it *SqueethReduceDebtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SqueethReduceDebt)
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
		it.Event = new(SqueethReduceDebt)
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
func (it *SqueethReduceDebtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SqueethReduceDebtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SqueethReduceDebt represents a ReduceDebt event raised by the Squeeth contract.
type SqueethReduceDebt struct {
	Sender             common.Address
	VaultId            *big.Int
	EthRedeemed        *big.Int
	WPowerPerpRedeemed *big.Int
	WPowerPerpBurned   *big.Int
	WPowerPerpExcess   *big.Int
	Bounty             *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterReduceDebt is a free log retrieval operation binding the contract event 0xfd0ae2fd36bd955810ae42615bc5ff277c0d0dfcb930f06c9f1777c0ef0752e3.
//
// Solidity: event ReduceDebt(address sender, uint256 vaultId, uint256 ethRedeemed, uint256 wPowerPerpRedeemed, uint256 wPowerPerpBurned, uint256 wPowerPerpExcess, uint256 bounty)
func (_Squeeth *SqueethFilterer) FilterReduceDebt(opts *bind.FilterOpts) (*SqueethReduceDebtIterator, error) {

	logs, sub, err := _Squeeth.contract.FilterLogs(opts, "ReduceDebt")
	if err != nil {
		return nil, err
	}
	return &SqueethReduceDebtIterator{contract: _Squeeth.contract, event: "ReduceDebt", logs: logs, sub: sub}, nil
}

// WatchReduceDebt is a free log subscription operation binding the contract event 0xfd0ae2fd36bd955810ae42615bc5ff277c0d0dfcb930f06c9f1777c0ef0752e3.
//
// Solidity: event ReduceDebt(address sender, uint256 vaultId, uint256 ethRedeemed, uint256 wPowerPerpRedeemed, uint256 wPowerPerpBurned, uint256 wPowerPerpExcess, uint256 bounty)
func (_Squeeth *SqueethFilterer) WatchReduceDebt(opts *bind.WatchOpts, sink chan<- *SqueethReduceDebt) (event.Subscription, error) {

	logs, sub, err := _Squeeth.contract.WatchLogs(opts, "ReduceDebt")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SqueethReduceDebt)
				if err := _Squeeth.contract.UnpackLog(event, "ReduceDebt", log); err != nil {
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

// ParseReduceDebt is a log parse operation binding the contract event 0xfd0ae2fd36bd955810ae42615bc5ff277c0d0dfcb930f06c9f1777c0ef0752e3.
//
// Solidity: event ReduceDebt(address sender, uint256 vaultId, uint256 ethRedeemed, uint256 wPowerPerpRedeemed, uint256 wPowerPerpBurned, uint256 wPowerPerpExcess, uint256 bounty)
func (_Squeeth *SqueethFilterer) ParseReduceDebt(log types.Log) (*SqueethReduceDebt, error) {
	event := new(SqueethReduceDebt)
	if err := _Squeeth.contract.UnpackLog(event, "ReduceDebt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SqueethShutdownIterator is returned from FilterShutdown and is used to iterate over the raw logs and unpacked data for Shutdown events raised by the Squeeth contract.
type SqueethShutdownIterator struct {
	Event *SqueethShutdown // Event containing the contract specifics and raw log

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
func (it *SqueethShutdownIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SqueethShutdown)
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
		it.Event = new(SqueethShutdown)
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
func (it *SqueethShutdownIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SqueethShutdownIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SqueethShutdown represents a Shutdown event raised by the Squeeth contract.
type SqueethShutdown struct {
	IndexForSettlement *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterShutdown is a free log retrieval operation binding the contract event 0x574214b195bf5273a95bb4498e35cf1fde0ce327c727a95ec2ab359f7ba4e11a.
//
// Solidity: event Shutdown(uint256 indexForSettlement)
func (_Squeeth *SqueethFilterer) FilterShutdown(opts *bind.FilterOpts) (*SqueethShutdownIterator, error) {

	logs, sub, err := _Squeeth.contract.FilterLogs(opts, "Shutdown")
	if err != nil {
		return nil, err
	}
	return &SqueethShutdownIterator{contract: _Squeeth.contract, event: "Shutdown", logs: logs, sub: sub}, nil
}

// WatchShutdown is a free log subscription operation binding the contract event 0x574214b195bf5273a95bb4498e35cf1fde0ce327c727a95ec2ab359f7ba4e11a.
//
// Solidity: event Shutdown(uint256 indexForSettlement)
func (_Squeeth *SqueethFilterer) WatchShutdown(opts *bind.WatchOpts, sink chan<- *SqueethShutdown) (event.Subscription, error) {

	logs, sub, err := _Squeeth.contract.WatchLogs(opts, "Shutdown")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SqueethShutdown)
				if err := _Squeeth.contract.UnpackLog(event, "Shutdown", log); err != nil {
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

// ParseShutdown is a log parse operation binding the contract event 0x574214b195bf5273a95bb4498e35cf1fde0ce327c727a95ec2ab359f7ba4e11a.
//
// Solidity: event Shutdown(uint256 indexForSettlement)
func (_Squeeth *SqueethFilterer) ParseShutdown(log types.Log) (*SqueethShutdown, error) {
	event := new(SqueethShutdown)
	if err := _Squeeth.contract.UnpackLog(event, "Shutdown", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SqueethUnPausedIterator is returned from FilterUnPaused and is used to iterate over the raw logs and unpacked data for UnPaused events raised by the Squeeth contract.
type SqueethUnPausedIterator struct {
	Event *SqueethUnPaused // Event containing the contract specifics and raw log

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
func (it *SqueethUnPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SqueethUnPaused)
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
		it.Event = new(SqueethUnPaused)
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
func (it *SqueethUnPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SqueethUnPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SqueethUnPaused represents a UnPaused event raised by the Squeeth contract.
type SqueethUnPaused struct {
	Unpauser common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUnPaused is a free log retrieval operation binding the contract event 0xff2b959f2bcdb44c7ecb4b16dae055431019d7350607125cfc2b74a06632c90e.
//
// Solidity: event UnPaused(address unpauser)
func (_Squeeth *SqueethFilterer) FilterUnPaused(opts *bind.FilterOpts) (*SqueethUnPausedIterator, error) {

	logs, sub, err := _Squeeth.contract.FilterLogs(opts, "UnPaused")
	if err != nil {
		return nil, err
	}
	return &SqueethUnPausedIterator{contract: _Squeeth.contract, event: "UnPaused", logs: logs, sub: sub}, nil
}

// WatchUnPaused is a free log subscription operation binding the contract event 0xff2b959f2bcdb44c7ecb4b16dae055431019d7350607125cfc2b74a06632c90e.
//
// Solidity: event UnPaused(address unpauser)
func (_Squeeth *SqueethFilterer) WatchUnPaused(opts *bind.WatchOpts, sink chan<- *SqueethUnPaused) (event.Subscription, error) {

	logs, sub, err := _Squeeth.contract.WatchLogs(opts, "UnPaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SqueethUnPaused)
				if err := _Squeeth.contract.UnpackLog(event, "UnPaused", log); err != nil {
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

// ParseUnPaused is a log parse operation binding the contract event 0xff2b959f2bcdb44c7ecb4b16dae055431019d7350607125cfc2b74a06632c90e.
//
// Solidity: event UnPaused(address unpauser)
func (_Squeeth *SqueethFilterer) ParseUnPaused(log types.Log) (*SqueethUnPaused, error) {
	event := new(SqueethUnPaused)
	if err := _Squeeth.contract.UnpackLog(event, "UnPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SqueethUpdateOperatorIterator is returned from FilterUpdateOperator and is used to iterate over the raw logs and unpacked data for UpdateOperator events raised by the Squeeth contract.
type SqueethUpdateOperatorIterator struct {
	Event *SqueethUpdateOperator // Event containing the contract specifics and raw log

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
func (it *SqueethUpdateOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SqueethUpdateOperator)
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
		it.Event = new(SqueethUpdateOperator)
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
func (it *SqueethUpdateOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SqueethUpdateOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SqueethUpdateOperator represents a UpdateOperator event raised by the Squeeth contract.
type SqueethUpdateOperator struct {
	Sender   common.Address
	VaultId  *big.Int
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdateOperator is a free log retrieval operation binding the contract event 0x3137fc9cd2e33c34f86e29c24d81f3c75b0bce639d3c4ed0d31eeff1160a7ff5.
//
// Solidity: event UpdateOperator(address sender, uint256 vaultId, address operator)
func (_Squeeth *SqueethFilterer) FilterUpdateOperator(opts *bind.FilterOpts) (*SqueethUpdateOperatorIterator, error) {

	logs, sub, err := _Squeeth.contract.FilterLogs(opts, "UpdateOperator")
	if err != nil {
		return nil, err
	}
	return &SqueethUpdateOperatorIterator{contract: _Squeeth.contract, event: "UpdateOperator", logs: logs, sub: sub}, nil
}

// WatchUpdateOperator is a free log subscription operation binding the contract event 0x3137fc9cd2e33c34f86e29c24d81f3c75b0bce639d3c4ed0d31eeff1160a7ff5.
//
// Solidity: event UpdateOperator(address sender, uint256 vaultId, address operator)
func (_Squeeth *SqueethFilterer) WatchUpdateOperator(opts *bind.WatchOpts, sink chan<- *SqueethUpdateOperator) (event.Subscription, error) {

	logs, sub, err := _Squeeth.contract.WatchLogs(opts, "UpdateOperator")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SqueethUpdateOperator)
				if err := _Squeeth.contract.UnpackLog(event, "UpdateOperator", log); err != nil {
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

// ParseUpdateOperator is a log parse operation binding the contract event 0x3137fc9cd2e33c34f86e29c24d81f3c75b0bce639d3c4ed0d31eeff1160a7ff5.
//
// Solidity: event UpdateOperator(address sender, uint256 vaultId, address operator)
func (_Squeeth *SqueethFilterer) ParseUpdateOperator(log types.Log) (*SqueethUpdateOperator, error) {
	event := new(SqueethUpdateOperator)
	if err := _Squeeth.contract.UnpackLog(event, "UpdateOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SqueethWithdrawCollateralIterator is returned from FilterWithdrawCollateral and is used to iterate over the raw logs and unpacked data for WithdrawCollateral events raised by the Squeeth contract.
type SqueethWithdrawCollateralIterator struct {
	Event *SqueethWithdrawCollateral // Event containing the contract specifics and raw log

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
func (it *SqueethWithdrawCollateralIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SqueethWithdrawCollateral)
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
		it.Event = new(SqueethWithdrawCollateral)
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
func (it *SqueethWithdrawCollateralIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SqueethWithdrawCollateralIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SqueethWithdrawCollateral represents a WithdrawCollateral event raised by the Squeeth contract.
type SqueethWithdrawCollateral struct {
	Sender  common.Address
	VaultId *big.Int
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawCollateral is a free log retrieval operation binding the contract event 0x627a692d5a03ab34732c0d2aa319f3ecdebdc4528f383eabcb25441dc0a70cfb.
//
// Solidity: event WithdrawCollateral(address sender, uint256 vaultId, uint256 amount)
func (_Squeeth *SqueethFilterer) FilterWithdrawCollateral(opts *bind.FilterOpts) (*SqueethWithdrawCollateralIterator, error) {

	logs, sub, err := _Squeeth.contract.FilterLogs(opts, "WithdrawCollateral")
	if err != nil {
		return nil, err
	}
	return &SqueethWithdrawCollateralIterator{contract: _Squeeth.contract, event: "WithdrawCollateral", logs: logs, sub: sub}, nil
}

// WatchWithdrawCollateral is a free log subscription operation binding the contract event 0x627a692d5a03ab34732c0d2aa319f3ecdebdc4528f383eabcb25441dc0a70cfb.
//
// Solidity: event WithdrawCollateral(address sender, uint256 vaultId, uint256 amount)
func (_Squeeth *SqueethFilterer) WatchWithdrawCollateral(opts *bind.WatchOpts, sink chan<- *SqueethWithdrawCollateral) (event.Subscription, error) {

	logs, sub, err := _Squeeth.contract.WatchLogs(opts, "WithdrawCollateral")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SqueethWithdrawCollateral)
				if err := _Squeeth.contract.UnpackLog(event, "WithdrawCollateral", log); err != nil {
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

// ParseWithdrawCollateral is a log parse operation binding the contract event 0x627a692d5a03ab34732c0d2aa319f3ecdebdc4528f383eabcb25441dc0a70cfb.
//
// Solidity: event WithdrawCollateral(address sender, uint256 vaultId, uint256 amount)
func (_Squeeth *SqueethFilterer) ParseWithdrawCollateral(log types.Log) (*SqueethWithdrawCollateral, error) {
	event := new(SqueethWithdrawCollateral)
	if err := _Squeeth.contract.UnpackLog(event, "WithdrawCollateral", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SqueethWithdrawUniPositionTokenIterator is returned from FilterWithdrawUniPositionToken and is used to iterate over the raw logs and unpacked data for WithdrawUniPositionToken events raised by the Squeeth contract.
type SqueethWithdrawUniPositionTokenIterator struct {
	Event *SqueethWithdrawUniPositionToken // Event containing the contract specifics and raw log

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
func (it *SqueethWithdrawUniPositionTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SqueethWithdrawUniPositionToken)
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
		it.Event = new(SqueethWithdrawUniPositionToken)
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
func (it *SqueethWithdrawUniPositionTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SqueethWithdrawUniPositionTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SqueethWithdrawUniPositionToken represents a WithdrawUniPositionToken event raised by the Squeeth contract.
type SqueethWithdrawUniPositionToken struct {
	Sender  common.Address
	VaultId *big.Int
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawUniPositionToken is a free log retrieval operation binding the contract event 0xe59f38fa1264fc25c9f0185eee136eaf810d90b8e7293b342e4037c68720177a.
//
// Solidity: event WithdrawUniPositionToken(address sender, uint256 vaultId, uint256 tokenId)
func (_Squeeth *SqueethFilterer) FilterWithdrawUniPositionToken(opts *bind.FilterOpts) (*SqueethWithdrawUniPositionTokenIterator, error) {

	logs, sub, err := _Squeeth.contract.FilterLogs(opts, "WithdrawUniPositionToken")
	if err != nil {
		return nil, err
	}
	return &SqueethWithdrawUniPositionTokenIterator{contract: _Squeeth.contract, event: "WithdrawUniPositionToken", logs: logs, sub: sub}, nil
}

// WatchWithdrawUniPositionToken is a free log subscription operation binding the contract event 0xe59f38fa1264fc25c9f0185eee136eaf810d90b8e7293b342e4037c68720177a.
//
// Solidity: event WithdrawUniPositionToken(address sender, uint256 vaultId, uint256 tokenId)
func (_Squeeth *SqueethFilterer) WatchWithdrawUniPositionToken(opts *bind.WatchOpts, sink chan<- *SqueethWithdrawUniPositionToken) (event.Subscription, error) {

	logs, sub, err := _Squeeth.contract.WatchLogs(opts, "WithdrawUniPositionToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SqueethWithdrawUniPositionToken)
				if err := _Squeeth.contract.UnpackLog(event, "WithdrawUniPositionToken", log); err != nil {
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

// ParseWithdrawUniPositionToken is a log parse operation binding the contract event 0xe59f38fa1264fc25c9f0185eee136eaf810d90b8e7293b342e4037c68720177a.
//
// Solidity: event WithdrawUniPositionToken(address sender, uint256 vaultId, uint256 tokenId)
func (_Squeeth *SqueethFilterer) ParseWithdrawUniPositionToken(log types.Log) (*SqueethWithdrawUniPositionToken, error) {
	event := new(SqueethWithdrawUniPositionToken)
	if err := _Squeeth.contract.UnpackLog(event, "WithdrawUniPositionToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
