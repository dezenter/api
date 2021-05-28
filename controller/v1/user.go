package v1

import (
	"strconv"

	"github.com/dezenter/api/model"
	"github.com/dezenter/api/repository"
	"github.com/dezenter/api/util"
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
	repo := repository.NewUserRepository()
	r, err := repo.Paginate(currentPage, limit)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	return c.JSON(r)
}

// UserCreate ...
func UserCreate(c *fiber.Ctx) error {
	params := model.UserCreateInput{}

	c.BodyParser(&params)
	repo := repository.NewUserRepository()
	r, err := repo.Create(params)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data": r,
	})

}

// UserShow ...
func UserShow(c *fiber.Ctx) error {
	id := c.Params("id")
	repo := repository.NewUserRepository()
	r, err := repo.FindByID(id)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data": r,
	})
}

// UserUpdate ...
func UserUpdate(c *fiber.Ctx) error {
	id := c.Params("id")
	params := model.UserUpdateInput{}
	c.BodyParser(&params)
	repo := repository.NewUserRepository()
	r, err := repo.Update(id, params)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data": r,
	})
}

// UserDelete ...
func UserDelete(c *fiber.Ctx) error {
	id := c.Params("id")

	repo := repository.NewUserRepository()
	_, err := repo.Delete(id)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": util.MsgSuccessDelete,
	})
}