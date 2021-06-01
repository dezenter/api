package v1

import (
	"github.com/dezenter/api/models"
	"github.com/dezenter/api/repositories"
	"github.com/dezenter/api/validators"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	params := models.UserRegisterInput{}
	c.BodyParser(&params)

	errors := validators.Register(params)
	if errors != nil {
		return c.JSON(fiber.Map{
			"status":  false,
			"message": errors,
		})
	}

	isActive := false

	convert := models.UserCreateInput{
		Username:  params.Username,
		Password:  params.Password,
		Email:     params.Email,
		FirstName: params.FirstName,
		LastName:  params.LastName,
		IsActive:  &isActive,
	}

	repo := repositories.NewUserRepository()
	r, err := repo.Create(convert)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": true,
		"data":   r,
	})
}
