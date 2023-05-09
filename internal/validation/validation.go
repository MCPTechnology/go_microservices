package validation

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

func Validate(target interface{}) []error {
	v := validator.New()
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	err := v.Struct(target)
	if err != nil {
		return ExtendValidationErrors(err.(validator.ValidationErrors))
	}
	return nil
}

func ExtendValidationErrors(errs validator.ValidationErrors) []error {
	convertedErrs := make([]error, len(errs))
	for i, err := range errs {
		convertedErrs[i] = getErrorForTag(err)
	}
	return convertedErrs
}

var (
	ErrFieldRequired        = errors.New("field is required")
	ErrInvalidEmail         = errors.New("invalid email")
	ErrMinimumSizeConstrain = errors.New("field does not fit it's minimum size of")
	ErrMaximumSizeConstrain = errors.New("field does not fit it's maximum size of")
	ErrGTEConstrain         = errors.New("field value should be greater than or equal to")
	ErrLTEConstrain         = errors.New("field value should be smaller than or equal to")
	ErrGTConstrain          = errors.New("field value should be greater than")
	ErrLTConstrain          = errors.New("field value should be smaller than")
)

func getErrorForTag(err validator.FieldError) error {
	field := err.Field()
	tag := err.Tag()
	switch tag {
	case "required":
		return fmt.Errorf("%v %w", field, ErrFieldRequired)
	case "email":
		return ErrInvalidEmail
	case "min":
		param := err.Param()
		return fmt.Errorf("%v %w %v", field, ErrMinimumSizeConstrain, param)
	case "max":
		param := err.Param()
		return fmt.Errorf("%v %w %v", field, ErrMaximumSizeConstrain, param)
	case "gte":
		param := err.Param()
		return fmt.Errorf("%v %w %v", field, ErrGTEConstrain, param)
	case "lte":
		param := err.Param()
		return fmt.Errorf("%v %w %v", field, ErrLTEConstrain, param)
	case "gt":
		param := err.Param()
		return fmt.Errorf("%v %w %v", field, ErrGTConstrain, param)
	case "lt":
		param := err.Param()
		return fmt.Errorf("%v %w %v", field, ErrLTConstrain, param)
	}
	return err // default error
}
