// Author: Turing Zhu
// Date: 2020/8/12 9:09 AM
// File: timestamp.go

package ginSwagger

import (
	"github.com/gin-gonic/gin"
	_ "huobiapi/models"
)

// Get Current Timestamp
// @summary Get Current Timestamp
// @description This endpoint returns the current timestamp, i.e. the number of milliseconds that have elapsed since 00:00:00 UTC on 1 January 1970.
// @id TimestampV1
// @tags Reference Data
// @router /v1/common/timestamp [GET]
// @success 200 integer models.Timestamp
// @accept json
// @produce json
func Timestamp(ctx *gin.Context) {

}
