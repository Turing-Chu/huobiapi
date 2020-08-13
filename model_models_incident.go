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
// ModelsIncident struct for ModelsIncident
type ModelsIncident struct {
	Components []ModelsComponent `json:"components,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	Id string `json:"id,omitempty"`
	Impact string `json:"impact,omitempty"`
	IncidentUpdates []ModelsIncidentUpdate `json:"incident_updates,omitempty"`
	MonitoringAt string `json:"monitoring_at,omitempty"`
	Name string `json:"name,omitempty"`
	PageId string `json:"page_id,omitempty"`
	ResolvedAt string `json:"resolved_at,omitempty"`
	Shortlink string `json:"shortlink,omitempty"`
	StartedAt string `json:"started_at,omitempty"`
	Status string `json:"status,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}