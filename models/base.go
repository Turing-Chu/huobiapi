// Author: Turing Zhu
// Date: 2020/8/13 4:05 PM
// File: base.go

package models

type Response struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseV1 struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}
