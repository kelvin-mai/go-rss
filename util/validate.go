package util

import (
	"github.com/go-playground/validator/v10"
	"kelvinmai.io/rss/internal/model"
)

var validate = validator.New()

func Validate(data interface{}) ([]model.ValidationError, bool) {
	ve := []model.ValidationError{}
	errs := validate.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			ve = append(ve, model.ValidationError{
				Field: err.Field(),
				Tag:   err.Tag(),
				Value: err.Value(),
			})
		}
	}
	return ve, len(ve) == 0
}
