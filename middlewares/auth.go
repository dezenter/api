package middlewares

import (
	"fmt"
	"os"

	"github.com/dezenter/api/models"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

// Protected protect routes
func Auth() fiber.Handler {
	s := os.Getenv("JWT_SECRET")

	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(s),
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{
				"status":  "error",
				"message": "Missing or malformed JWT",
				"data":    nil,
			})
	}
	return c.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid or expired JWT",
			"data":    nil,
		})
}

func UserForContext(c *fiber.Ctx) (*models.UserAuth, error) {
	user := c.Locals("user").(*jwt.Token)
	if !user.Valid {
		return nil, fmt.Errorf("error")
	}

	u := user.Claims.(jwt.MapClaims)
	getUser := models.UserAuth{
		UserId:    u["userId"].(string),
		FirstName: u["firstName"].(string),
		LastName:  u["lastName"].(string),
		Role:      u["role"].(string),
	}
	return &getUser, nil
}
