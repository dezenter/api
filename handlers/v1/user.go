package v1

import (
	"fmt"
	"strconv"

	"github.com/dezenter/api/models"
	"github.com/dezenter/api/repositories"
	"github.com/dezenter/api/utils"
	"github.com/gofiber/fiber/v2"
)

// UserIndex ...
func UserIndex(c *fiber.Ctx) error {
	var currentPage = 1
	getCurrentPage := c.Query("page")
	if getCurrentPage != "" {
		currentPage, _ = strconv.Atoi(getCurrentPage)
	}
	limit := 15
	repo := repositories.NewUserRepository()
	r, err := repo.Paginate(currentPage, limit)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": true,
		"data":   r,
	})
}

// UserCreate ...
func UserCreate(c *fiber.Ctx) error {
	params := models.UserCreateInput{}

	c.BodyParser(&params)
	repo := repositories.NewUserRepository()
	r, err := repo.Create(params)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": true,
		"data":   r,
	})
}

// UserShow ...
func UserShow(c *fiber.Ctx) error {
	id := c.Params("id")
	repo := repositories.NewUserRepository()
	r, err := repo.FindByID(id)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": true,
		"data":   r,
	})
}

// UserUpdate ...
func UserUpdate(c *fiber.Ctx) error {
	id := c.Params("id")
	params := models.UserUpdateInput{}
	c.BodyParser(&params)
	repo := repositories.NewUserRepository()
	r, err := repo.Update(id, params)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": true,
		"data":   r,
	})
}

// UserDelete ...
func UserDelete(c *fiber.Ctx) error {
	id := c.Params("id")

	repo := repositories.NewUserRepository()
	_, err := repo.Delete(id)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": utils.MsgSuccessDelete,
	})
}

// UserMe
func UserMe(c *fiber.Ctx) error {
	fmt.Println(c)
	return nil
}
