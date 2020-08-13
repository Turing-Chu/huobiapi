// Author: Turing Zhu
// Date: 2020/8/13 6:30 PM
// File: trading_symbol.go

package ginSwagger

import (
	"github.com/gin-gonic/gin"
	_ "huobiapi/models"
)

// Get all Supported Trading Symbol
// @summary Get all Supported Trading Symbol
// @description This endpoint returns all Huobi's supported trading symbol.
// @id SupportedTradingSymbolV1
// @tags Reference Data
// @router /v1/common/symbols [GET]
// @success 200 {object} models.ResponseV1{data=[]models.Symbol} "Supported Trading Symbol List"
// @accept json
// @produce json
func SupportedTradingSymbol(ctx *gin.Context) {}
