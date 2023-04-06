package validation

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type ErrMsg struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func Errors(err error) interface{} {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		var errs = make(map[string]ErrMsg)
		for _, e := range ve {
			errs[e.Field()] = ErrMsg{
				Code:    "ERR_" + e.Tag(),
				Message: getErrMsg(e),
			}
		}
		return errs
	}
	if err.Error() == "EOF" {
		return "invalid request body"
	}
	return err.Error()
}

func getErrMsg(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return "is required."
	case "email":
		return "must be a valid email address."
	case "min":
		return fmt.Sprintf("must be at least %s characters.", err.Param())
	case "max":
		return fmt.Sprintf("must be at most %s characters.", err.Param())
	default:
		return err.Error()
	}
}
