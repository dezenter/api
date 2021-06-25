package router

import (
	c "github.com/dezenter/api/handlers/v1"
	"github.com/dezenter/api/middlewares"
	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {

	v1 := app.Group("/v1")
	// V1 Guest
	v1.Post("/register", c.Register)
	v1.Post("/forget-password", c.ForgetPassword)
	v1.Post("/reset-password", c.ResetPassword)

	// V1 Users
	v1.Get("/users", c.UserIndex)
	v1.Post("/users", c.UserCreate)
	v1.Get("/users/:id", c.UserShow)
	v1.Patch("/users/:id", c.UserUpdate)
	v1.Delete("/users/:id", c.UserDelete)

	v1.Get("/me", middlewares.Auth(), c.Me)
	v1.Patch("/me", middlewares.Auth(), c.MeUpdate)
	v1.Delete("/me", middlewares.Auth(), c.MeDelete)

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

	// V1 Admin
	v1.Get("/admins", c.AdminIndex)
	v1.Post("/admins", c.AdminCreate)
	v1.Get("/admins/:id", c.AdminShow)
	v1.Patch("/admins/:id", c.AdminUpdate)
	v1.Delete("/admins/:id", c.AdminDelete)
}
