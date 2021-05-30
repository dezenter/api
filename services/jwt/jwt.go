package jwt

import (
	"os"
	"time"

	"github.com/dezenter/api/models"
	"github.com/dezenter/api/utils"
	"github.com/form3tech-oss/jwt-go"
)

// CreateUserToken ...
func CreateUserToken(u *models.User) (*models.UserToken, error) {
	exp := time.Now().Add(time.Hour * utils.AccessTokenExpire).Unix()

	s := os.Getenv("JWT_SECRET")
	ac := jwt.MapClaims{
		"userId":    u.ID,
		"firstName": u.FirstName,
		"lastName":  u.LastName,
		"role":      "user",
		"exp":       exp,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, ac)
	gt, err := token.SignedString([]byte(s))
	if err != nil {
		return nil, err
	}
	return &models.UserToken{Token: gt, User: *u}, nil
}

// CreateAdminToken ...
// func CreateAdminToken(u *models.Admin) (*models.AdminToken, error) {
// 	exp := time.Now().Add(time.Hour * utils.AccessTokenExpire).Unix()

// 	s := os.Getenv("JWT_SECRET")
// 	ac := jwt.MapClaims{
// 		"adminId":   u.ID,
// 		"firstName": u.FirstName,
// 		"lastName":  u.LastName,
// 		"role":      u.Role,
// 		"exp":       exp,
// 	}
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, ac)
// 	gt, err := token.SignedString([]byte(s))
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &models.AdminToken{Token: gt, Admin: *u}, nil
// }
