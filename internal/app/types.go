package app

import (
	"github.com/BEON-Tech-Studio/golang-live-coding-challenge/internal/models"
)

type BaseResponse struct {
	Status string      `json:"status"`
	Info   interface{} `json:"info"`
}

type StatesResponse struct {
	BaseResponse
	States []models.State `json:"data"`
}

// Timing information for API execution
type Timing struct {
	Executing int    `json:"executing"`
	Unit      string `json:"unit"`
}

// Total record count information
type Total struct {
	RecordCount int `json:"recordCount"`
}

// Detailed information about the API response
type Info struct {
	Timing         Timing `json:"timing"`
	ResultCoverage string `json:"resultCoverage"`
	Total          Total  `json:"total"`
}

// Individual report item
type Report struct {
	Num         int     `json:"num"`
	Header      string  `json:"header"`
	Description *string `json:"description"`
	Terms       string  `json:"terms"`
}

// ReportsResponse represents the complete response structure for reports
type ReportsResponse struct {
	Status  string   `json:"status"`
	Info    Info     `json:"info"`
	Message *string  `json:"message"`
	Data    []Report `json:"data"`
	Detail  *string  `json:"detail"`
	Errors  *string  `json:"errors"`
}
