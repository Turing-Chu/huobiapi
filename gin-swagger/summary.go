// Author: Turing Zhu
// Date: 2020/8/10 9:15 PM
// File: summary.go

package ginSwagger

import (
	"github.com/gin-gonic/gin"
	_ "huobiapi/models"
)

// Get system status
// @summary Get system status
// @description This endpoint allows users to get system status, Incidents and planned maintenance.
// @tags Reference Data
// @id Summary
// @host status.huobigroup.com
// @router /api/v2/summary.json [GET]
// @success 200 {object} models.Summary{page=models.Page,components=[]models.Component,incidents=[]models.Incident{incident_updates=[]models.IncidentUpdate{affected_components=[]models.AffectedComponent}},scheduled_maintenances=[]models.ScheduledMaintenance,status=models.SystemStatus}
// @accept json
// @produce json
func Summary(ctx *gin.Context) {}
