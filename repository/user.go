package repository

import (
	"errors"
	"math"

	"github.com/dezenter/api/config"
	"github.com/dezenter/api/model"
	"github.com/dezenter/api/util"
	"gorm.io/gorm"
)

// UserRepository ...
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository ...
func NewUserRepository() *UserRepository {
	return &UserRepository{db: config.DB}
}

// Paginate user
func (u *UserRepository) Paginate(page int, limit int) (*model.UserPaginate, error) {
	var users []*model.User
	var count int64

	var skip int
	if page == 1 {
		skip = 0
	} else {
		skip = limit * (page - 1)
	}
	u.db.Model(&model.User{}).Count(&count)
	u.db.Limit(limit).Offset(skip).Order("created_at desc").Find(&users)

	lastPage := int(math.Ceil(float64(count) / float64(limit)))
	p := model.UserPaginate{
		Total:    count,
		PerPage:  limit,
		Page:     page,
		LastPage: lastPage,
		Users:    users,
	}
	return &p, nil
}

// Create user
func (u *UserRepository) Create(input model.UserCreateInput) (*model.User, error) {
	var c int64
	r := u.db.Model(&model.User{}).Where("email = ?", input.Email).Count(&c)
	if r.Error != nil {
		return nil, errors.New("Email has exists")
	}
	r = u.db.Model(&model.User{}).Where("username = ?", input.Username).Count(&c)
	if r.Error != nil {
		return nil, errors.New("Username has exists")
	}

	uID := util.GenerateID()

	pwd, err := util.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	isActive := false

	user := model.User{
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
func (u *UserRepository) FindByID(id string) (*model.User, error) {
	var user model.User
	r := u.db.Where("id = ?", id).First(&user)
	if r.Error != nil {
		return nil, r.Error
	}
	return &user, nil
}

// Update user
func (u *UserRepository) Update(id string, input model.UserUpdateInput) (*model.User, error) {

	user := model.User{
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
	var user model.User
	r := u.db.Model(&model.User{}).Where("id = ?", id).First(&user)
	if r.Error != nil {
		return false, r.Error
	}
	r = u.db.Where("id = ?", id).Delete(&model.User{})
	if r.Error != nil {
		return false, r.Error
	}
	return true, nil
}

// FindByEmail ...
func (u *UserRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	r := u.db.Where("email = ?", email).First(&user)
	if r.Error != nil {
		return nil, r.Error
	}
	return &user, nil
}
