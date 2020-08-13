// Author: Turing Zhu
// Date: 2020/8/10 9:17 PM
// File: market-status.go

package ginSwagger

import (
	"github.com/gin-gonic/gin"
	_ "huobiapi/models"
)

// Get Market Status
// @summary The endpoint returns current market status
// @description The enum values of market status includes: 1 - normal (order submission & cancellation are allowed)，
// 2 - halted (order submission & cancellation are prohibited)，3 - cancel-only(order submission is prohibited but order
// cancellation is allowed). Halt reason includes: 2 - emergency maintenance，3 - schedule maintenance.
// @tags Reference Data
// @id MarketStatusV2
// @router /v2/market-status [GET]
// @success 200 {object} models.Response{data=models.MarketStatus}
// @accept json
// @produce json
func MarketStatus(ctx *gin.Context) {}
