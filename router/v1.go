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
	v1.Put("/users/:id", c.UserUpdate)
	v1.Patch("/users/:id", c.UserUpdate)
	v1.Delete("/users/:id", c.UserDelete)

}
