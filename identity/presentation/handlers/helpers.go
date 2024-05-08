package handlers

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
)

type APIError struct {
	StatusCode int `json:"statusCode"`
	Message    any `json:"message"`
}

type APIFunc func(responseWriter http.ResponseWriter, request *http.Request) error

func (apiError APIError) Error() string {
	return fmt.Sprintf("api error: %d", apiError.StatusCode)
}

func NewApiError(statusCode int, err error) APIError {
	return APIError{
		StatusCode: statusCode,
		Message:    err.Error(),
	}
}

func InvalidRequestData(errors map[string]string) APIError {
	return APIError{
		StatusCode: http.StatusUnprocessableEntity,
		Message:    errors,
	}
}

func InvalidJSON() APIError {
	return NewApiError(http.StatusBadRequest, fmt.Errorf("invalid JSON request data"))
}

func Make(handler APIFunc) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		if err := handler(responseWriter, request); err != nil {
			if apiErr, ok := err.(APIError); ok {
				writeJSON(responseWriter, apiErr.StatusCode, apiErr)
			} else {
				errResp := map[string]any{
					"statusCode": http.StatusInternalServerError,
					"message":    "internal server error",
				}
				writeJSON(responseWriter, http.StatusInternalServerError, errResp)
			}
			slog.Error("HTTP API error", "err", err.Error(), "path", request.URL.Path)
		}
	}
}

func writeJSON(responseWriter http.ResponseWriter, status int, v any) error {
	responseWriter.WriteHeader(status)
	responseWriter.Header().Set("Content-Type", "application/problem+json")

	return json.NewEncoder(responseWriter).Encode(v)
}
