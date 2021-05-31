package validators

import (
	"github.com/dezenter/api/models"
	"github.com/go-playground/validator"
)

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func CreatePostStruct(post models.PostCreateInput) []*ErrorResponse {
	var errors []*ErrorResponse
	validate := validator.New()
	err := validate.Struct(post)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
