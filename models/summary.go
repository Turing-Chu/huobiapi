// Author: Turing Zhu
// Date: 2020/8/10 1:09 PM
// File: summary.go

package models

import (
	"time"
)

// System components and their status
type Component struct {
	Id string `json:"id"` // component id
	// component name, including Order submission, Order cancellation, Deposit etc.
	Name string `json:"name"`
	// component status, value range: operational, degraded_performance, partial_outage, major_outage, under maintenance
	Status             string    `json:"status"`
	CreatedAt          time.Time `json:"created_at"` // component create time
	UpdatedAt          time.Time `json:"updated_at"` // component update time
	Position           int64     `json:"position"`
	Description        *string   `json:"description"`
	Showcase           bool      `json:"showcase"`
	GroupId            string    `json:"group_id"`
	PageId             string    `json:"page_id"`
	Group              bool      `json:"group"`
	OnlyShowIfDegraded bool      `json:"only_show_if_degraded"`
}

// System fault incident and their status. If there is no fault at present, it will return to null
type Incident struct {
	Id   string `json:"id"`   // incident id
	Name string `json:"name"` // incident name
	// incident staus, value range: investigating, identified, monitoring, resolved
	Status          string           `json:"status"`
	CreatedAt       time.Time        `json:"created_at"` // incident creat time
	UpdatedAt       time.Time        `json:"updated_at"` // incident update time
	MonitoringAt    *time.Time       `json:"monitoring_at"`
	ResolvedAt      *time.Time       `json:"resolved_at"`
	Impact          string           `json:"impact"`
	Shortlink       string           `json:"shortlink"`
	StartedAt       string           `json:"started_at"`
	PageId          string           `json:"page_id"`
	IncidentUpdates []IncidentUpdate `json:"incident_updates"`
	Components      []Component      `json:"components"`
}

type AffectedComponent struct {
	Code      string `json:"code"`
	Name      string `json:"name"`
	OldStatus string `json:"old_status"`
	NewStatus string `json:"new_status"`
}

type IncidentUpdate struct {
	Id                   string              `json:"id"`
	Status               string              `json:"status"`
	Body                 string              `json:"body"`
	IncidentId           string              `json:"incident_id"`
	CreatedAt            time.Time           `json:"created_at"`
	UpdatedAt            time.Time           `json:"updated_at"`
	DisplayAt            time.Time           `json:"display_at"`
	AffectedComponents   []AffectedComponent `json:"affected_components"`
	DeliverNotifications bool                `json:"deliver_notifications"`
	CustomTweet          interface{}         `json:"custom_tweet"` // TODO
	TweetId              *string             `json:"tweet_id"`
}

// System scheduled maintenance incident and status. If there is no scheduled maintenance at present, it will return to null
type ScheduledMaintenance struct {
	Id              string           `json:"id"`         // incident id
	Name            string           `json:"name"`       // incident name
	Status          string           `json:"status"`     // incident status, value range: scheduled, in progress, verifying, completed
	CreatedAt       time.Time        `json:"created_at"` // incident creat time
	UpdatedAt       time.Time        `json:"updated_at"` // incident update time
	MonitoringAt    *time.Time       `json:"monitoring_at"`
	ResolvedAt      *time.Time       `json:"resolved_at"`
	Impact          string           `json:"impact"`
	Shortlink       string           `json:"shortlink"`
	StartedAt       string           `json:"started_at"`
	PageId          string           `json:"page_id"`
	IncidentUpdates []IncidentUpdate `json:"incident_updates"`
	Components      []Component      `json:"components"`
	ScheduledFor    time.Time        `json:"scheduled_for"`   // scheduled maintenance start time
	ScheduledUntil  time.Time        `json:"scheduled_until"` // scheduled maintenance end time
}

// The overall current status of the system
type SystemStatus struct {
	// system indicator, value range: none, minor, major, critical, maintenance
	Indicator string `json:"indicator"`
	// system description, value range: All Systems Operational, Minor Service Outager, Partial System Outage, Partially
	// Degraded Service, Service Under Maintenance
	Description string `json:"description"`
}

type Summary struct {
	// basic information of huobi spot status page
	Page                  Page                   `json:"page"`
	// System components and their status
	Components            []Component            `json:"components"`
	// System fault incident and their status. If there is no fault at present, it will return to null
	Incidents             []Incident             `json:"incidents"`
	// System scheduled maintenance incident and status. If there is no scheduled maintenance at present, it will return to null
	ScheduledMaintenances []ScheduledMaintenance `json:"scheduled_maintenances"`
	// The overall current status of the system
	Status SystemStatus `json:"status"`
}
