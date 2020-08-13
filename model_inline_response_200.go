/*
 * Huobi API
 *
 * Huobi API
 *
 * API version: 0.1
 * Contact: qishiwenjun@163.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package huobiapi
// InlineResponse200 struct for InlineResponse200
type InlineResponse200 struct {
	Components []ModelsComponent `json:"components,omitempty"`
	Incidents []ModelsIncident `json:"incidents,omitempty"`
	ModelsSummary ModelsSummary `json:"models.Summary,omitempty"`
	Page ModelsPage `json:"page,omitempty"`
	ScheduledMaintenances []ModelsScheduledMaintenance `json:"scheduled_maintenances,omitempty"`
	Status ModelsSystemStatus `json:"status,omitempty"`
}
