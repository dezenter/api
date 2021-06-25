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

// AdminRepository
type AdminRepository struct {
	db *gorm.DB
}

// NewAdminRepository
func NewAdminRepository() *AdminRepository {
	return &AdminRepository{db: configs.DB}
}

// Paginate admin
func (u *AdminRepository) Paginate(page int, limit int) (*models.AdminPaginate, error) {
	var admins []*models.Admin
	var count int64

	var skip int
	if page == 1 {
		skip = 0
	} else {
		skip = limit * (page - 1)
	}
	u.db.Model(&models.Admin{}).Count(&count)
	u.db.Limit(limit).Offset(skip).Order("created_at desc").Find(&admins)

	lastPage := int(math.Ceil(float64(count) / float64(limit)))
	p := models.AdminPaginate{
		Total:    count,
		PerPage:  limit,
		Page:     page,
		LastPage: lastPage,
		Admins:   admins,
	}
	return &p, nil
}

// Create admin
func (u *AdminRepository) Create(input models.AdminCreateInput) (*models.Admin, error) {
	var c int64
	r := u.db.Model(&models.Admin{}).Where("email = ?", input.Email).Count(&c)
	if r.Error != nil {
		return nil, errors.New("Email has exists")
	}
	r = u.db.Model(&models.Admin{}).Where("username = ?", input.Username).Count(&c)
	if r.Error != nil {
		return nil, errors.New("Username has exists")
	}

	uID := utils.GenerateID()

	pwd, err := utils.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	if err := input.Role.IsValid(); err != nil {
		return nil, err
	}

	admin := models.Admin{
		ID:        uID,
		Username:  input.Username,
		Password:  pwd,
		Email:     input.Email,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		IsActive:  input.IsActive,
		Role:      input.Role,
	}

	r = u.db.Create(&admin)
	if r.Error != nil {
		return nil, r.Error
	}

	r = u.db.Where("id = ?", uID).First(&admin)
	if r.Error != nil {
		return nil, r.Error
	}

	return &admin, nil
}

// FindByID admin
func (u *AdminRepository) FindByID(id string) (*models.Admin, error) {
	var admin models.Admin
	r := u.db.Where("id = ?", id).First(&admin)
	if r.Error != nil {
		return nil, r.Error
	}
	return &admin, nil
}

// Update admin
func (u *AdminRepository) Update(id string, input models.AdminUpdateInput) (*models.Admin, error) {
	if err := input.Role.IsValid(); err != nil {
		return nil, err
	}
	admin := models.Admin{
		ID:        id,
		Email:     input.Email,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		IsActive:  input.IsActive,
		Role:      input.Role,
	}
	r := u.db.Updates(&admin)
	if r.Error != nil {
		return nil, r.Error
	}

	r = u.db.Where("id = ?", id).First(&admin)
	if r.Error != nil {
		return nil, r.Error
	}
	return &admin, nil
}

// Delete admin
func (u *AdminRepository) Delete(id string) (bool, error) {
	var admin models.Admin
	r := u.db.Model(&models.Admin{}).Where("id = ?", id).First(&admin)
	if r.Error != nil {
		return false, r.Error
	}
	r = u.db.Where("id = ?", id).Delete(&models.Admin{})
	if r.Error != nil {
		return false, r.Error
	}
	return true, nil
}

// FindByEmail
func (u *AdminRepository) FindByEmail(email string) (*models.Admin, error) {
	var admin models.Admin
	r := u.db.Where("email = ?", email).First(&admin)
	if r.Error != nil {
		return nil, r.Error
	}
	return &admin, nil
}

// Login admin
func (a *AdminRepository) Login(input models.AdminLoginInput) (*models.AdminToken, error) {
	var u models.Admin

	r := a.db.Where("username = ?", input.Username).Where("is_active", true).First(&u)
	if r.Error != nil {
		return nil, r.Error
	}

	c := utils.CheckPasswordHash(input.Password, u.Password)
	if c == false {
		return nil, fmt.Errorf("Password not match")
	}

	ut, err := jwt.CreateAdminToken(&u)
	if err != nil {
		return nil, err
	}

	return ut, nil
}
