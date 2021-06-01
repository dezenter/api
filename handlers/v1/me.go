package v1

import (
	"github.com/dezenter/api/middlewares"
	"github.com/dezenter/api/models"
	"github.com/dezenter/api/repositories"
	"github.com/dezenter/api/validators"
	"github.com/gofiber/fiber/v2"
)

// Me
func Me(c *fiber.Ctx) error {
	a, err := middlewares.UserForContext(c)
	if err != nil {
		return c.JSON(fiber.Map{
			"status":  false,
			"message": "error",
		})
	}

	repo := repositories.NewUserRepository()
	u, err := repo.FindByID(a.UserId)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": true,
		"data":   u,
	})
}

func MeUpdate(c *fiber.Ctx) error {
	a, err := middlewares.UserForContext(c)
	if err != nil {
		return c.JSON(fiber.Map{
			"status":  false,
			"message": "error",
		})
	}

	params := models.UserUpdateMeInput{}
	c.BodyParser(&params)
	errors := validators.UpdateMe(params)
	if errors != nil {
		return c.JSON(fiber.Map{
			"status":  false,
			"message": errors,
		})
	}

	convert := models.UserUpdateInput{
		Email:     params.Email,
		FirstName: params.FirstName,
		LastName:  params.LastName,
	}

	repo := repositories.NewUserRepository()
	u, err := repo.Update(a.UserId, convert)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": true,
		"data":   u,
	})
}

func MeDelete(c *fiber.Ctx) error {
	a, err := middlewares.UserForContext(c)
	if err != nil {
		return c.JSON(fiber.Map{
			"status":  false,
			"message": "error",
		})
	}

	repo := repositories.NewUserRepository()
	u, err := repo.Delete(a.UserId)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": true,
		"data":   u,
	})
}
