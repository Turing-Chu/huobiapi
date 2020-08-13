// Author: Turing Zhu
// Date: 2020/8/10 9:18 PM
// File: market-status.go

package models

type MarketStatus struct {
	// Market status (1=normal, 2=halted, 3=cancel-only)
	MarketStatus int32 `json:"marketStatus" example:"1"`
	// Halt start time (unix time in millisecond) , only valid for marketStatus=halted or cancel-only
	HaltStartTime int64 `json:"haltStartTime,omitempty" `
	// Estimated halt end time (unix time in millisecond) , only valid for marketStatus=halted or cancel-only; if this field is not returned during marketStatus=halted or cancel-only, it implicates the halt end time cannot be estimated at this time.
	HaltEndTime int64 `json:"haltEndTime,omitempty"`
	// Halt reason (2=emergency-maintenance, 3=scheduled-maintenance) , only valid for marketStatus=halted or cancel-only
	HaltReason int64 `json:"haltReason,omitempty"`
	// Affected symbols, separated by comma. If affect all symbols just respond with value ‘all’. Only valid for marketStatus=halted or cancel-only
	AffectedSymbols string `json:"affectedSymbols,omitempty"`
}
