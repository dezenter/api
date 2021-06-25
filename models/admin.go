package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

// Admin model
type Admin struct {
	ID        string         `gorm:"primaryKey" json:"id"`
	Username  string         `gorm:"unique;not null" json:"username"`
	Password  string         `gorm:"unique;not null" json:"-"`
	Email     string         `gorm:"unique;not null" json:"email"`
	FirstName string         `gorm:"not null" json:"firstName"`
	LastName  string         `gorm:"not null" json:"lastName"`
	IsActive  *bool          `gorm:"not null" json:"isActive"`
	Role      Role           `gorm:"not null" json:"role"`
	LastLogin *time.Time     `json:"lastLogin"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Role string

const (
	SUPER_ADMIN Role = "SUPER_ADMIN"
	ADMIN            = "ADMIN"
	EDITOR           = "EDITOR"
	SUPPORT          = "SUPPORT"
)

func (lt Role) IsValid() error {
	switch lt {
	case SUPER_ADMIN, ADMIN, EDITOR, SUPPORT:
		return nil
	}
	return errors.New("Invalid role type")
}

// AdminPaginate
type AdminPaginate struct {
	Total    int64    `json:"total"`
	PerPage  int      `json:"perPage"`
	Page     int      `json:"page"`
	LastPage int      `json:"lastPage"`
	Admins   []*Admin `json:"admins"`
}

// AdminCreateInput
type AdminCreateInput struct {
	Username  string `validate:"required" json:"username"`
	Password  string `validate:"required" json:"password"`
	Email     string `validate:"required" json:"email"`
	FirstName string `validate:"required" json:"firstName"`
	LastName  string `validate:"required" json:"lastName"`
	Role      Role   `validate:"required" json:"role"`
	IsActive  *bool  `validate:"required" json:"isActive"`
}

// AdminUpdateInput
type AdminUpdateInput struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Role      Role   `json:"role"`
	IsActive  *bool  `json:"isActive"`
}

// AdminLoginInput
type AdminLoginInput struct {
	Username string `validate:"required" json:"username"`
	Password string `validate:"required" json:"password"`
}

// AdminToken jwt token
type AdminToken struct {
	Token string `json:"token"`
	Admin Admin  `json:"admin"`
}

type AdminAuth struct {
	AdminID   string `json:"adminId"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Role      string `json:"role"`
}
