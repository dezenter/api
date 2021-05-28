package router

import (
	c "github.com/dezenter/api/handlers/v1"
	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {

	v1 := app.Group("/v1")

	// V1 Users
	v1.Get("/users", c.UserIndex)
	v1.Post("/users", c.UserCreate)
	v1.Get("/users/:id", c.UserShow)
	v1.Patch("/users/:id", c.UserUpdate)
	v1.Delete("/users/:id", c.UserDelete)

	v1.Get("/me", c.UserMe)

	// V1 Auth
	v1.Post("/auth", c.AuthLogin)

	// V1 Post Category
	v1.Get("/posts/categories", c.PostCategoryIndex)
	v1.Post("/posts/categories", c.PostCategoryCreate)
	v1.Get("/posts/categories/:id", c.PostCategoryShow)
	v1.Patch("/posts/categories/:id", c.PostCategoryUpdate)
	v1.Delete("/posts/categories/:id", c.PostCategoryDelete)

	// V1 Post
	// v1.Get("/posts", c.PostIndex)
	// v1.Post("/posts", c.PostCreate)
	// v1.Get("/posts/:id", c.PostShow)
	// v1.Patch("/posts/:id", c.PostUpdate)
	// v1.Delete("/posts/:id", c.PostDelete)
}
