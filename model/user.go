package model

import (
	"time"

	"gorm.io/gorm"
)

// User model
type User struct {
	ID        string         `gorm:"primaryKey" json:"id"`
	Username  string         `gorm:"unique;not null" json:"username"`
	Password  string         `gorm:"unique;not null" json:"password"`
	Email     *string        `gorm:"unique;not null" json:"email"`
	FirstName *string        `gorm:"not null" json:"firstName"`
	LastName  *string        `gorm:"not null" json:"lastName"`
	IsActive  *bool          `gorm:"not null" json:"isActive"`
	LastLogin time.Time      `json:"lastLogin"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	// Posts     []*Post        `gorm:"foreignKey:created_by" json:"posts"`
}

// UserPaginate pagination
type UserPaginate struct {
	Total    int64 `json:"total"`
	PerPage  int   `json:"perPage"`
	Page     int   `json:"page"`
	LastPage int   `json:"lastPage"`
	Users    []*User
}

// UserCreateInput create user
type UserCreateInput struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// UserUpdateInput update input
type UserUpdateInput struct {
	Email     *string `json:"email"`
	FirstName *string `json:"firstName"`
	LastName  *string `json:"lastName"`
	IsActive  *bool   `json:"isActive"`
}

// UserRegisterInput register user
type UserRegisterInput struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// UserUpdateMeInput update input
type UserUpdateMeInput struct {
	Email     *string `json:"email"`
	FirstName *string `json:"firstName"`
	LastName  *string `json:"lastName"`
}

// UserLoginInput ...
type UserLoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// UserToken jwt token
type UserToken struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}
