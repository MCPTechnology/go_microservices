package errs

import (
	"fmt"
	"net/http"
	"strings"
)

type sentinelAPIError struct {
	StatusCode int      `json:"-"`
	Msg        string   `json:"error,omitempty"`
	Details    []string `json:"details,omitempty"`
}

func (e sentinelAPIError) Error() string {
	return e.errorMessage()
}

func (e sentinelAPIError) errorMessage() string {
	return fmt.Sprintf("%v: %v", e.Msg, strings.Join(e.Details, "; "))
}

var (
	msgAuthError           = "invalid token"
	msgNotFoundError       = "resource could not be found"
	msgDuplicateError      = "duplicate resource"
	msgInternalServerError = "the server encountered an unexpected condition that prevented it from fulfilling the request"
	msgBadRequestError     = "server cannot or will not process the request due to something that is perceived to be a client error"
	msgValidationError     = "validation error"
	msgInvalidUUID         = "the provided id is invalid"
)

var (
	ErrAuth           = &sentinelAPIError{StatusCode: http.StatusUnauthorized, Msg: msgAuthError}
	ErrNotFound       = &sentinelAPIError{StatusCode: http.StatusNotFound, Msg: msgNotFoundError}
	ErrDuplicate      = &sentinelAPIError{StatusCode: http.StatusBadRequest, Msg: msgDuplicateError}
	ErrInternalServer = &sentinelAPIError{StatusCode: http.StatusInternalServerError, Msg: msgInternalServerError}
	ErrBadRequest     = &sentinelAPIError{StatusCode: http.StatusInternalServerError, Msg: msgBadRequestError}
	ErrValidation     = &sentinelAPIError{StatusCode: http.StatusBadRequest, Msg: msgValidationError}
	ErrInvalidUUID    = &sentinelAPIError{StatusCode: http.StatusBadRequest, Msg: msgInvalidUUID}
)
