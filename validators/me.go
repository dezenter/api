package validators

import (
	"github.com/dezenter/api/models"
	"github.com/dezenter/api/utils"
	"github.com/go-playground/validator"
)

func UpdateMe(input models.UserUpdateMeInput) []*utils.ErrorResponse {
	var errors []*utils.ErrorResponse
	validate := validator.New()
	err := validate.Struct(input)
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
