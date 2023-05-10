package utils

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/MCPTechnology/go_microservices/internal/errs"
)

var ErrMarshalError error = errors.New("unable to parse data structure")

func ToJson(w io.Writer, target interface{}) error {
	err := json.NewEncoder(w).Encode(target)
	if err != nil {
		return ErrMarshalError
	}
	return nil
}

func FromJson(r io.Reader, target interface{}) error {
	err := json.NewDecoder(r).Decode(target)
	if err != nil {
		return errs.WrapError(errs.ErrBadRequest, err)
	}
	return nil
}

func ToHTTPResponse(rw http.ResponseWriter, statusCode int, details interface{}) {
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("X-Content-Type-Options", "nosniff")
	rw.WriteHeader(statusCode)
	if err := ToJson(rw, details); err != nil {
		unmarshalHttpError(rw)
	}
}

func ToHTTPErrorResponse(rw http.ResponseWriter, err error) {
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("X-Content-Type-Options", "nosniff")
	statusCode, details := getErrorDetails(err)
	rw.WriteHeader(statusCode)
	if err := ToJson(rw, details); err != nil {
		unmarshalHttpError(rw)
	}
}

func getErrorDetails(err error) (int, interface{}) {
	switch e := err.(type) {
	case errs.SentinelWrappedError:
		return e.Sentinel.StatusCode, e.Sentinel
	default:
		wErr := errs.WrapError(errs.ErrInternalServer, e)
		return wErr.Sentinel.StatusCode, wErr.Sentinel
	}
}

func unmarshalHttpError(rw http.ResponseWriter) {
	http.Error(rw, ErrMarshalError.Error(), http.StatusInternalServerError)
}
