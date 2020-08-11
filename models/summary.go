// Author: Turing Zhu
// Date: 2020/8/10 1:09 PM
// File: summary.go

package models

import (
	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
	"github.com/vmihailenco/msgpack"
	"time"
)

// basic information of huobi spot status page
type Page struct {
	Id        string    `json:"id" msgpack:"id"`                 // page id
	Name      string    `json:"name" msgpack:"name"`             // page name
	Url       string    `json:"url" msgpack:"url"`               // page url
	TimeZone  string    `json:"time_zone" msgpack:"time_zone"`   // time zone
	UpdatedAt time.Time `json:"updated_at" msgpack:"updated_at"` // page update time
}

///implement encoding.JSON.Marshaler interface
func (m *Page) MarshalJSON() ([]byte, error)    { return msgpack.Marshal(m) }
func (m *Page) UnmarshalJSON(data []byte) error { return msgpack.Unmarshal(data, m) }

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

func (m *Component) MarshalJSON() ([]byte, error)    { return msgpack.Marshal(m) }
func (m *Component) UnmarshalJSON(data []byte) error { return msgpack.Unmarshal(data, m) }

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

func (m *Incident) MarshalJSON() ([]byte, error)    { return msgpack.Marshal(m) }
func (m *Incident) UnmarshalJSON(data []byte) error { return msgpack.Unmarshal(data, m) }

type AffectedComponent struct {
	Code      string `json:"code"`
	Name      string `json:"name"`
	OldStatus string `json:"old_status"`
	NewStatus string `json:"new_status"`
}

func (m *AffectedComponent) MarshalJSON() ([]byte, error)    { return msgpack.Marshal(m) }
func (m *AffectedComponent) UnmarshalJSON(data []byte) error { return msgpack.Unmarshal(data, m) }

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

func (m *IncidentUpdate) MarshalJSON() ([]byte, error)    { return msgpack.Marshal(m) }
func (m *IncidentUpdate) UnmarshalJSON(data []byte) error { return msgpack.Unmarshal(data, m) }

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

func (m *ScheduledMaintenance) MarshalJSON() ([]byte, error)    { return msgpack.Marshal(m) }
func (m *ScheduledMaintenance) UnmarshalJSON(data []byte) error { return msgpack.Unmarshal(data, m) }

// The overall current status of the system
type Status struct {
	// system indicator, value range: none, minor, major, critical, maintenance
	Indicator string `json:"indicator"`
	// system description, value range: All Systems Operational, Minor Service Outager, Partial System Outage, Partially
	// Degraded Service, Service Under Maintenance
	Description string `json:"description"`
}

func (m *Status) MarshalJSON() ([]byte, error)    { return msgpack.Marshal(m) }
func (m *Status) UnmarshalJSON(data []byte) error { return msgpack.Unmarshal(data, m) }

type Summary struct {
	// basic information of huobi spot status page
	//Page                  Page                   `json:"page" msgpack:"page"`
	//Components            []Component            `json:"components" msgpack:"components"`
	//Incidents             []Incident             `json:"incidents" msgpack:"incidents"`
	//ScheduledMaintenances []ScheduledMaintenance `json:"scheduled_maintenances" msgpack:"scheduled_maintenances"`
	Status                Status                 `json:"status" msgpack:"status"`
}
