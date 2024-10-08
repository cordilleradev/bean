package types

import (
	"fmt"
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

func InvalidPeriodError(period, message string) *APIError {
	return NewAPIError(400, "invalid_leaderboard_period", fmt.Sprintf("%s is not a valid leaderboard period (%s)", period, message))
}

func InvalidLimitError(limit int, allowedLimit int) *APIError {
	return NewAPIError(400, "invalid_leaderboard_limit", fmt.Sprintf("%d is greater than maximum limit (%d)", limit, allowedLimit))
}

func FailedLeaderboardCall(err error) *APIError {
	return NewAPIError(500, "failed_leaderboard_call", err.Error())
}

func FailedPositionsCall(err error) *APIError {
	return NewAPIError(500, "failed_positions_call", err.Error())
}

func NoSuchStream(userId string) *APIError {
	return NewAPIError(400, "no_such_stream", fmt.Sprintf("there is no active stream for '%s'", userId))
}

func StreamAlreadyStarted(userId string) *APIError {
	return NewAPIError(400, "stream_in_progress", fmt.Sprintf("there is already a stream in progress for '%s'", userId))
}
