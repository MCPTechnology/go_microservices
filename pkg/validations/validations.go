package validations

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

func Validate(target interface{}) error {
	v := validator.New()
	v.RegisterTagNameFunc(
		func(f reflect.StructField) string {
			name := strings.SplitN(f.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	err := v.Struct(target)
	if err != nil {
		validationErr := err.(validator.ValidationErrors)
		return newValidationError(validationErr)
	}
	return nil
}

var ErrValidationError = errors.New("Validation Error")

func newValidationError(errDetails validator.ValidationErrors) error {
	errs := errors.Join(extendValidationErrors(errDetails)...)
	return fmt.Errorf("%w\n%w",ErrValidationError, errs)
}

func readConvertedErrors(validationErrs validator.ValidationErrors) <-chan error {
	ch := make(chan error)
	go func() {
		for _, err := range validationErrs {
			ch <- getErrorForTag(err)
		}
		close(ch)
	}()
	return ch
}

func extendValidationErrors(validationErrs validator.ValidationErrors) []error {
	errs := make([]error, 0, len(validationErrs))
	for _, err := range validationErrs {
		msg := getErrorForTag(err)
		errs = append(errs, msg)
	}
	return errs
}

func getErrorForTag(ve validator.FieldError) error {
	field := ve.Field()
	tag := ve.Tag()
	switch tag {
	case "required":
		return fmt.Errorf("%v field is required", field)
	case "email":
		param := ve.Param()
		return fmt.Errorf("Invalid email: %v", param)
	case "min":
		param := ve.Param()
		return fmt.Errorf("%v field does not fit it's minimum size of %v", field, param)
	case "max":
		param := ve.Param()
		return fmt.Errorf("%v field does not fit it's maximum size of %v", field, param)
	case "gte":
		param := ve.Param()
		return fmt.Errorf("%v field should be greater than or equal to %v", field, param)
	case "lte":
		param := ve.Param()
		return fmt.Errorf("%v field should be smaller than or equal to %v", field, param)
	}
	return ve // default error
}
