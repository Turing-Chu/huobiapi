// Author: Turing Zhu
// Date: 2020/8/13 4:04 PM
// File: currency.go

package ginSwagger

import (
	_ "github.com/Turing-Chu/huobiapi/models"
	"github.com/gin-gonic/gin"
)

// Get Currency & Chains
// @summary Currency & Chains
// @description API user could query static reference information for each currency, as well as its corresponding chain(s). (Public Endpoint)
// @id CurrencyAndChainsV2
// @tags Reference Data
// @router /v2/reference/currencies [GET]
// @success 200 {object} models.Response{data=models.Currency} "Currency & Chains list"
// @accept json
// @produce json
func Currency(ctx *gin.Context) {}
