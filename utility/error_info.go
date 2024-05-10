package utility

import "time"

type ErrorInfo struct {
	Code      int       `json:"code"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

func NewErrorInfo(code int, message string, timestamp time.Time) ErrorInfo {
	return ErrorInfo{
		Code:      code,
		Message:   message,
		Timestamp: timestamp,
	}
}
