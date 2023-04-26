package errs

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

var ErrValidationError = errors.New("Validation Error")

func NewValidationError(errDetails validator.ValidationErrors) error {
	return fmt.Errorf("%w: %w", ErrValidationError, errDetails)
}
