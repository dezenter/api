package v1

import (
	"github.com/dezenter/api/models"
	"github.com/dezenter/api/repositories"
	"github.com/gofiber/fiber/v2"
)

// AuthLogin
func AuthLogin(c *fiber.Ctx) error {
	params := models.UserLoginInput{}

	c.BodyParser(&params)
	repo := repositories.NewUserRepository()
	r, err := repo.Login(params)

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
