package validators

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type customValidator struct {
	validator *validator.Validate
}

func NewCustomValidator() echo.Validator {
	v := validator.New()
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	return &customValidator{validator: v}
}

func (cv *customValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
