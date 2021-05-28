package middlewares

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	user string
}

// UserCustomClaims ...
type UserCustomClaims struct {
	UserID    string `json:"userId"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Role      string `json:"role"`
	jwt.StandardClaims
}

func Auth() func(c *fiber.Ctx) error {
	fmt.Println()
	return nil
	// return c.Next()
	// scr := os.Getenv("JWT_SECRET")

	// token, err := jwt.ParseWithClaims(ac[7:], &UserCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
	// 	return []byte(scr), nil
	// })

	// if err != nil {
	// 	http.Error(w, `{"message": "`+err.Error()+`"}`, http.StatusUnauthorized)
	// 	return
	// }

	// claims, ok := token.Claims.(*UserCustomClaims)
	// if !ok && !token.Valid {
	// 	http.Error(w, `{"message": "`+err.Error()+`"}`, http.StatusUnauthorized)
	// 	return
	// }
}
