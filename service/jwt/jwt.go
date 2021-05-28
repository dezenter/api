package jwt

import (
	"os"
	"time"

	"github.com/dezenter/api/model"
	"github.com/dezenter/api/util"
	"github.com/dgrijalva/jwt-go"
)

// CreateUserToken ...
func CreateUserToken(u *model.User) (*model.UserToken, error) {
	exp := time.Now().Add(time.Hour * util.AccessTokenExpire).Unix()

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
	return &model.UserToken{Token: gt, User: *u}, nil
}

// CreateAdminToken ...
// func CreateAdminToken(u *model.Admin) (*model.AdminToken, error) {
// 	exp := time.Now().Add(time.Hour * util.AccessTokenExpire).Unix()

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
// 	return &model.AdminToken{Token: gt, Admin: *u}, nil
// }
