package repositories

import (
	"errors"
	"fmt"
	"math"

	"github.com/dezenter/api/configs"
	"github.com/dezenter/api/models"
	"github.com/dezenter/api/services/jwt"
	"github.com/dezenter/api/utils"
	"gorm.io/gorm"
)

// UserRepository ...
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository ...
func NewUserRepository() *UserRepository {
	return &UserRepository{db: configs.DB}
}

// Paginate user
func (u *UserRepository) Paginate(page int, limit int) (*models.UserPaginate, error) {
	var users []*models.User
	var count int64

	var skip int
	if page == 1 {
		skip = 0
	} else {
		skip = limit * (page - 1)
	}
	u.db.Model(&models.User{}).Count(&count)
	u.db.Limit(limit).Offset(skip).Order("created_at desc").Find(&users)

	lastPage := int(math.Ceil(float64(count) / float64(limit)))
	p := models.UserPaginate{
		Total:    count,
		PerPage:  limit,
		Page:     page,
		LastPage: lastPage,
		Users:    users,
	}
	return &p, nil
}

// Create user
func (u *UserRepository) Create(input models.UserCreateInput) (*models.User, error) {
	var c int64
	r := u.db.Model(&models.User{}).Where("email = ?", input.Email).Count(&c)
	if r.Error != nil {
		return nil, errors.New("Email has exists")
	}
	r = u.db.Model(&models.User{}).Where("username = ?", input.Username).Count(&c)
	if r.Error != nil {
		return nil, errors.New("Username has exists")
	}

	uID := utils.GenerateID()

	pwd, err := utils.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	isActive := false

	user := models.User{
		ID:        uID,
		Username:  input.Username,
		Password:  pwd,
		Email:     &input.Email,
		FirstName: &input.FirstName,
		LastName:  &input.LastName,
		IsActive:  &isActive,
	}

	r = u.db.Create(&user)
	if r.Error != nil {
		return nil, r.Error
	}

	r = u.db.Where("id = ?", uID).First(&user)
	if r.Error != nil {
		return nil, r.Error
	}

	return &user, nil
}

// FindByID user
func (u *UserRepository) FindByID(id string) (*models.User, error) {
	var user models.User
	r := u.db.Where("id = ?", id).First(&user)
	if r.Error != nil {
		return nil, r.Error
	}
	return &user, nil
}

// Update user
func (u *UserRepository) Update(id string, input models.UserUpdateInput) (*models.User, error) {

	user := models.User{
		ID:        id,
		Email:     input.Email,
		FirstName: input.FirstName,
		LastName:  input.LastName,
	}
	r := u.db.Updates(&user)
	if r.Error != nil {
		return nil, r.Error
	}

	r = u.db.Where("id = ?", id).First(&user)
	if r.Error != nil {
		return nil, r.Error
	}
	return &user, nil
}

// Delete user
func (u *UserRepository) Delete(id string) (bool, error) {
	var user models.User
	r := u.db.Model(&models.User{}).Where("id = ?", id).First(&user)
	if r.Error != nil {
		return false, r.Error
	}
	r = u.db.Where("id = ?", id).Delete(&models.User{})
	if r.Error != nil {
		return false, r.Error
	}
	return true, nil
}

// FindByEmail ...
func (u *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	r := u.db.Where("email = ?", email).First(&user)
	if r.Error != nil {
		return nil, r.Error
	}
	return &user, nil
}

// Login user
func (a *UserRepository) Login(input models.UserLoginInput) (*models.UserToken, error) {
	var u models.User

	r := a.db.Where("username = ?", input.Username).Where("is_active", true).First(&u)
	if r.Error != nil {
		return nil, r.Error
	}

	c := utils.CheckPasswordHash(input.Password, u.Password)
	if c == false {
		return nil, fmt.Errorf("Password not match")
	}

	ut, err := jwt.CreateUserToken(&u)
	if err != nil {
		return nil, err
	}

	return ut, nil
}
