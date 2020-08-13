// Author: Turing Zhu
// Date: 2020/8/13 4:05 PM
// File: currency.go

package models

import (
	"github.com/shopspring/decimal"
)

type Currency struct {
	Currency         string  `json:"currency"`   // Currency
	InstrumentStatus string  `json:"instStatus"` // Instrument status	normal,delisted
	Chains           []Chain `json:"chains"`
}

type Chain struct {
	// Chain name
	Chain string `json:"chain"`
	// Chain display name
	DisplayName string `json:"displayName"`
	// Base chain name
	BaseChain string `json:"baseChain,omitempty"`
	// Base chain protocol
	BaseChainProtocol string `json:"baseChainProtocol,omitempty"`
	// Is dynamic fee type or not (only applicable to withdrawFeeType = fixed)
	IsDynamic bool `json:"isDynamic"`
	// Number of confirmations required for deposit success (trading & withdrawal allowed once reached)
	NumOfConfirmations int64 `json:"numOfConfirmations"`
	// Number of confirmations required for quick success (trading allowed but withdrawal disallowed once reached)
	NumOfFastConfirmations int64 `json:"numOfFastConfirmations"`
	// Minimal deposit amount in each request
	MinDepositAmt decimal.Decimal `json:"minDepositAmt"`
	// Deposit status allowed,prohibited
	DepositStatus string `json:"deposit_status"`
	// Minimal withdraw amount in each request
	MinWithdrawAmt decimal.Decimal `json:"minWithdrawAmt"`
	// Maximum withdraw amount in each request
	MaxWithdrawAmt decimal.Decimal `json:"maxWithdrawAmt"`
	// Maximum withdraw amount in a day (Singapore timezone)
	WithdrawQuotaPerDay decimal.Decimal `json:"withdrawQuotaPerDay"`
	// Maximum withdraw amount in a year
	WithdrawQuotaPerYear decimal.Decimal `json:"withdrawQuotaPerYear"`
	// Maximum withdraw amount in total
	WithdrawQuotaTotal decimal.Decimal `json:"withdrawQuotaTotal"`
	// Withdraw amount precision
	WithdrawPrecision int64 `json:"withdrawPrecision"`
	// Type of withdraw fee (only one type can be applied to each currency) fixed,circulated,ratio
	WithdrawFeeType string `json:"withdrawFeeType"`
	// Withdraw fee in each request (only applicable to withdrawFeeType = fixed)
	TransactFeeWithdraw decimal.Decimal `json:"transactFeeWithdraw,omitempty"`
	// Minimal withdraw fee in each request (only applicable to withdrawFeeType = circulated or ratio)
	MinTransactFeeWithdraw decimal.Decimal `json:"minTransactFeeWithdraw,omitempty"`
	// Maximum withdraw fee in each request (only applicable to withdrawFeeType = circulated or ratio)
	MaxTransactFeeWithdraw decimal.Decimal `json:"maxTransactFeeWithdraw,omitempty"`
	// Withdraw fee in each request (only applicable to withdrawFeeType = ratio)
	TransactFeeRateWithdraw decimal.Decimal `json:"transactFeeRateWithdraw,omitempty"`
	// Withdraw status allowed,prohibited
	WithdrawStatus string `json:"withdrawStatus"`
}
