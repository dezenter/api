package v1

import (
	"strconv"

	"github.com/dezenter/api/models"
	"github.com/dezenter/api/repositories"
	"github.com/dezenter/api/utils"
	"github.com/dezenter/api/validators"
	"github.com/gofiber/fiber/v2"
)

// AdminIndex
func AdminIndex(c *fiber.Ctx) error {
	var currentPage = 1
	getCurrentPage := c.Query("page")
	if getCurrentPage != "" {
		currentPage, _ = strconv.Atoi(getCurrentPage)
	}
	limit := 15
	repo := repositories.NewAdminRepository()
	r, err := repo.Paginate(currentPage, limit)

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

// AdminCreate
func AdminCreate(c *fiber.Ctx) error {
	params := models.AdminCreateInput{}
	c.BodyParser(&params)

	errors := validators.CreateAdmin(params)
	if errors != nil {
		return c.JSON(fiber.Map{
			"status":  false,
			"message": errors,
		})
	}

	repo := repositories.NewAdminRepository()
	r, err := repo.Create(params)

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

// AdminShow
func AdminShow(c *fiber.Ctx) error {
	id := c.Params("id")
	repo := repositories.NewAdminRepository()
	r, err := repo.FindByID(id)

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

// AdminUpdate
func AdminUpdate(c *fiber.Ctx) error {
	id := c.Params("id")
	params := models.AdminUpdateInput{}
	c.BodyParser(&params)

	repo := repositories.NewAdminRepository()
	r, err := repo.Update(id, params)

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

// AdminDelete
func AdminDelete(c *fiber.Ctx) error {
	id := c.Params("id")

	repo := repositories.NewAdminRepository()
	_, err := repo.Delete(id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": utils.MsgSuccessDelete,
	})
}
