// Author: Turing Zhu
// Date: 2020/8/13 6:17 PM
// File: trading_symbol.go

package models

type Symbol struct {
	// Base currency in a trading symbol
	BaseCurrency string `json:"base-currency"`
	// Quote currency in a trading symbol
	QuoteCurrency string `json:"quote-currency"`
	// Quote currency precision when quote price(decimal places)）
	PricePrecision int64 `json:"price-precision"`
	// Base currency precision when quote amount(decimal places)）
	AmountPrecision int64 `json:"amount-precision"`
	// Trading section, possible values: [main，innovation]
	SymbolPartition string `json:"symbol-partition"`
	// symbol
	Symbol string `json:"symbol"`
	// The status of the symbol；Allowable values: [online，offline,suspend].
	// "online" - Listed, available for trading,
	// "offline" - de-listed, not available for trading，
	// "suspend" - suspended for trading,
	// "pre-online" - to be online soon
	State string `json:"state"`
	// Precision of value in quote currency (value = price * amount)
	ValuePrecision int64 `json:"value-precision"`
	// Minimum order amount of limit order in base currency (to be obsoleted)
	MinOrderAmt float64 `json:"min-order-amt"`
	// Max order amount of limit order in base currency (to be obsoleted)
	MaxOrderAmt float64 `json:"max-order-amt"`
	// Minimum order amount of limit order in base currency (NEW)
	LimitOrderMinOrderAmt float64 `json:"limit-order-min-order-amt"`
	// Max order amount of limit order in base currency (NEW)
	LimitOrderMaxOrderAmt float64 `json:"limit-order-max-order-amt"`
	// Minimum order amount of sell-market order in base currency (NEW)
	SellMarketMinOrderAmt float64 `json:"sell-market-min-order-amt"`
	// Max order amount of sell-market order in base currency (NEW)
	SellMarketMaxOrderAmt float64 `json:"sell_market_max_order_amt"`
	// Max order value of buy-market order in quote currency (NEW)
	BuyMarketMaxOrderValue float64 `json:"buy-market-max-order-value"`
	// Minimum order value of limit order and buy-market order in quote currency
	MinOrderValue float64 `json:"min-order-value"`
	// Maximum order value of limit order and buy-market order in quote currency
	MaxOrderValue float64 `json:"max-order-value"`
	// The applicable leverage ratio
	// (only valid for symbols eligible for isolated margin trading, cross-margin trading, or ETP trading.)
	LeverageRatio float64 `json:"leverage_ratio"`
	// Underlying ETP code (only valid for ETP symbols)
	Underlying string `json:"underlying"`
	// Position charge rate (only valid for ETP symbols)
	MgmtFeeRate float64 `json:"mgmt-fee-rate"`
	// Position charging time (in GMT+8, in format HH:MM:SS, only valid for ETP symbols)
	ChargeTime string `json:"charge-time,omitempty"`
	// Regular position rebalance time (in GMT+8, in format HH:MM:SS, only valid for ETP symbols)
	RebalTime string `json:"rebal-time,omitempty"`
	// 	The threshold which triggers adhoc position rebalance (evaluated by actual leverage ratio, only valid for ETP symbols)
	RebalThreshold float64 `json:"rebal-threshold,omitempty"`
	// Initial NAV (only valid for ETP symbols)
	InitNav float64 `json:"init-nav,omitempty"`
}
