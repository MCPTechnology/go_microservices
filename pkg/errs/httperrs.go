package errs

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

func NewHttpValidationError(errDetails validator.ValidationErrors) error {
	err := &HttpValidationError{
		StatusCode: http.StatusBadRequest,
	}
	if len(errDetails) > 0 {
		err.addDetails(errDetails)
	}
	return err
}

type HttpValidationError struct {
	Details    []string
	StatusCode int
}

func (v HttpValidationError) Error() string {
	detailsString := strings.Join(v.Details, ",\n")
	return fmt.Sprintf("Validation error:\n%v", detailsString)
}

func (v *HttpValidationError) addDetails(validationErrs validator.ValidationErrors) {
	errs := make([]string, 0, len(validationErrs))
	for _, err := range validationErrs {
		msg := getMsgForTag(err)
		errs = append(errs, msg)
	}
	v.Details = errs
}

func getMsgForTag(ve validator.FieldError) string {
	field := ve.Field()
	tag := ve.Tag()
	switch tag {
	case "required":
		return fmt.Sprintf("%v field is required", field)
	case "email":
		return "Invalid email"
	case "min":
		param := ve.Param()
		return fmt.Sprintf("%v field does not fit it's minimum size of %v", field, param)
	case "max":
		param := ve.Param()
		return fmt.Sprintf("%v field does not fit it's maximum size of %v", field, param)
	case "gte":
		param := ve.Param()
		return fmt.Sprintf("%v field should be greater than or equal to %v", field, param)
	case "lte":
		param := ve.Param()
		return fmt.Sprintf("%v field should be smaller than or equal to %v", field, param)
	}
	return ve.Error() // default error
}
