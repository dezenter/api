package router

import (
	c "github.com/dezenter/api/handlers/v1"
	"github.com/dezenter/api/middlewares"
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

	v1.Get("/me", middlewares.Auth(), c.UserMe)

	// V1 Auth
	v1.Post("/auth", c.AuthLogin)

	// V1 Post Category
	v1.Get("/posts/categories", c.PostCategoryIndex)
	v1.Post("/posts/categories", middlewares.Auth(), c.PostCategoryCreate)
	v1.Get("/posts/categories/:id", c.PostCategoryShow)
	v1.Patch("/posts/categories/:id", middlewares.Auth(), c.PostCategoryUpdate)
	v1.Delete("/posts/categories/:id", middlewares.Auth(), c.PostCategoryDelete)

	// V1 Post
	v1.Get("/posts", c.PostIndex)
	v1.Post("/posts", middlewares.Auth(), c.PostCreate)
	v1.Get("/posts/:id", c.PostShow)
	v1.Patch("/posts/:id", middlewares.Auth(), c.PostUpdate)
	v1.Delete("/posts/:id", middlewares.Auth(), c.PostDelete)
}
