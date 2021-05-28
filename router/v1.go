package router

import (
	c "github.com/dezenter/api/controller/v1"
	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {

	v1 := app.Group("/v1")

	// V1 users
	v1.Get("/users", c.UserIndex)
	v1.Post("/users", c.UserCreate)
	v1.Get("/users/:id", c.UserShow)
	v1.Patch("/users/:id", c.UserUpdate)
	v1.Delete("/users/:id", c.UserDelete)

	// V1 post category
	v1.Get("/posts/categories", c.PostCategoryIndex)
	v1.Post("/posts/categories", c.PostCategoryCreate)
	v1.Get("/posts/categories/:id", c.PostCategoryShow)
	v1.Patch("/posts/categories/:id", c.PostCategoryUpdate)
	v1.Delete("/posts/categories/:id", c.PostCategoryDelete)
}
