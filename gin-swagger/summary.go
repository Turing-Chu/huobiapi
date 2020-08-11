// Author: Turing Zhu
// Date: 2020/8/10 9:15 PM
// File: summary.go

package ginSwagger

import (
	_ "github.com/Turing-Chu/huobiapi/models"
	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
)

// Get system status
// @summary Get system status
// @description This endpoint allows users to get system status, Incidents and planned maintenance.
// @tags Reference Data
// @router /v2/summary.json [GET]
// @success 200 {object} models.Summary
// @accept json
// @produce json
func Summary(ctx *gin.Context) {
}
