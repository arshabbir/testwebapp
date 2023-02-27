package utils

type ApiError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
