package errs

import (
	"errors"
)

type SentinelWrappedError struct {
	error
	Sentinel *sentinelAPIError
}

func appendDetailErrors(e *sentinelAPIError, errs []error) {
	e.Details = make([]string, len(errs))
	for i, err := range errs {
		e.Details[i] = err.Error()
	}
}

func WrapError(sentinel *sentinelAPIError, err ...error) SentinelWrappedError {
	switch e := err[0].(type) {
	case SentinelWrappedError:
		return e
	default:
		sentinelError := SentinelWrappedError{error: errors.Join(err...), Sentinel: sentinel}
		appendDetailErrors(sentinelError.Sentinel, err)
		return sentinelError
	}
}

func (swe SentinelWrappedError) Is(err error) bool {
	if errors.Is(swe.error, err) {
		return true
	}
	if errors.Is(swe.Sentinel, err) {
		return true
	}
	if swe.Sentinel == err {
		return true
	}
	return false
}
