package models

import "time"

type SuccessResponse struct {
	Message   string    `json:"message"`
	Data      any       `json:"data,omitempty"`
	Timestamp time.Time `json:"timestam	p"`
}

type FailureReponse struct {
	Message   string    `json:"message"`
	Data      any       `json:"data,omitempty"`
	Timestamp time.Time `json:"timestam	p"`
}
