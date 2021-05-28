package v1

import (
	"strconv"

	"github.com/dezenter/api/models"
	"github.com/dezenter/api/repositories"
	"github.com/dezenter/api/utils"
	"github.com/gofiber/fiber/v2"
)

// PostCategoryIndex ...
func PostCategoryIndex(c *fiber.Ctx) error {
	var currentPage = 1
	getCurrentPage := c.Query("page")
	if getCurrentPage != "" {
		currentPage, _ = strconv.Atoi(getCurrentPage)
	}
	limit := 15
	repo := repositories.NewPostCategoryRepository()
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

// PostCategoryCreate
func PostCategoryCreate(c *fiber.Ctx) error {
	params := models.PostCategoryInput{}

	c.BodyParser(&params)

	repo := repositories.NewPostCategoryRepository()
	r, err := repo.Create(params)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data": r,
	})
}

// PostCategoryShow
func PostCategoryShow(c *fiber.Ctx) error {
	id := c.Params("id")

	repo := repositories.NewPostCategoryRepository()
	r, err := repo.FindById(id)

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

// PostCategoryUpdate
func PostCategoryUpdate(c *fiber.Ctx) error {
	id := c.Params("id")

	params := models.PostCategoryInput{}

	c.BodyParser(&params)

	repo := repositories.NewPostCategoryRepository()
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

// PostCategoryDelete
func PostCategoryDelete(c *fiber.Ctx) error {
	id := c.Params("id")

	repo := repositories.NewPostCategoryRepository()
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
