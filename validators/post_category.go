package validators

import (
	"github.com/dezenter/api/models"
	"github.com/dezenter/api/utils"
	"github.com/go-playground/validator"
)

func CreatePostCategory(pc models.PostCategoryInput) []*utils.ErrorResponse {
	var errors []*utils.ErrorResponse
	validate := validator.New()
	err := validate.Struct(pc)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element utils.ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
