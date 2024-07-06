package types

import (
	"fmt"
	"strings"
)

type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

// Error implements the error interface.
func (e *APIError) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s, Details: %s", e.Code, e.Message, e.Details)
}

// NewAPIError creates a new APIError
func NewAPIError(code int, message string, details string) *APIError {
	return &APIError{
		Code:    code,
		Message: message,
		Details: details,
	}
}

func InvalidUserId(userId string) *APIError {
	return NewAPIError(400, "invalid_user_id", fmt.Sprintf("%s is not a valid user_id", userId))
}

func InvalidPeriodError(period string, allowedPeriods []string) *APIError {
	return NewAPIError(400, "invalid_leaderboard_period", fmt.Sprintf("%s is not a valid leaderboard_period (%s)", period, strings.Join(allowedPeriods, ", ")))
}
