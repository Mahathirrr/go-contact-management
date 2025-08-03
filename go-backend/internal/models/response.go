package models

type APIResponse struct {
	Data   interface{} `json:"data,omitempty"`
	Errors string      `json:"errors,omitempty"`
}

type ErrorResponse struct {
	Errors string `json:"errors"`
}

type SuccessResponse struct {
	Data interface{} `json:"data"`
}