package response

import (
	"net/http"
	"strings"
)

func StatusFromError(err error) int {
	msg := strings.ToLower(err.Error())

	switch {
	// NOT FOUND
	case strings.Contains(msg, "not found"),
		strings.Contains(msg, "failed to find"):
		return http.StatusNotFound

	// AUTH / UNAUTHORIZED
	case strings.Contains(msg, "invalid password"),
		strings.Contains(msg, "unauthorized"),
		strings.Contains(msg, "invalid token"):
		return http.StatusUnauthorized

	// CONFLICT / ALREADY EXISTS
	case strings.Contains(msg, "already"):
		return http.StatusConflict

	// BAD REQUEST
	case strings.Contains(msg, "invalid"),
		strings.Contains(msg, "mismatch"),
		strings.Contains(msg, "bad request"):
		return http.StatusBadRequest

	// RATE LIMIT
	case strings.Contains(msg, "limit exceeded"),
		strings.Contains(msg, "too many"):
		return http.StatusTooManyRequests

	// DEFAULT
	default:
		return http.StatusInternalServerError
	}
}
