// Author: Turing Zhu
// Date: 2020/8/13 4:04 PM
// File: currency.go

package ginSwagger

import (
	"github.com/gin-gonic/gin"
	_ "huobiapi/models"
)

// Get all Supported Currencies
// @summary Get all Supported Currencies
// @description This endpoint returns all Huobi's supported trading currencies.
// @id SupportCurrenciesV1
// @tags Reference Data
// @router /v1/common/currencys [GET]
// @success 200 {object} models.Response{data=[]string} "supported currency"
// @accept json
// @produce json
func SupportCurrency(ctx *gin.Context) {}

// Get Currency & Chains
// @summary Currency & Chains
// @description API user could query static reference information for each currency, as well as its corresponding chain(s). (Public Endpoint)
// @id CurrencyAndChainsV2
// @tags Reference Data
// @param currency path string false "btc, ltc, etc ...(available currencies in Huobi Global)"
// @param authorizedUser path boolean false "true or false (if not filled, default value is true)"
// @router /v2/reference/currencies [GET]
// @success 200 {object} models.Response{data=models.Currency{chains=[]models.Chain}} "Currency & Chains list"
// @accept json
// @produce json
func Currency(ctx *gin.Context) {}
