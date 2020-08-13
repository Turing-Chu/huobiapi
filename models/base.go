// Author: Turing Zhu
// Date: 2020/8/13 4:05 PM
// File: base.go

package models

import "time"

type Response struct {
	Code      int64       `json:"code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data,omitempty"`
	
	Timestamp int64       `json:"timestamp,omitempty"`
	Status    int64       `json:"status,omitempty"`
	Error     string      `json:"error,omitempty"`
	Path      string      `json:"path,omitempty"`
}

type ResponseV1 struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

// basic information of huobi spot status page
type Page struct {
	Id        string    `json:"id"`         // page id
	Name      string    `json:"name"`       // page name
	Url       string    `json:"url"`        // page url
	TimeZone  string    `json:"time_zone"`  // time zone
	UpdatedAt time.Time `json:"updated_at"` // page update time
}
