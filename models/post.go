package models

import (
	"time"

	"gorm.io/gorm"
)

// Post model
type Post struct {
	ID          string         `gorm:"primaryKey" json:"id"`
	CategoryID  string         `gorm:"not null"`
	Category    PostCategory   `gorm:"foreignKey:CategoryID" json:"category"`
	Title       string         `gorm:"not null" json:"title"`
	Description string         `json:"description"`
	Longitude   *float64       `json:"longitude"`
	Latitude    *float64       `json:"latitude"`
	Status      PostStatus     `gorm:"not null" json:"status"`
	CreatedBy   string         `gorm:"not null"`
	User        User           `gorm:"foreignKey:CreatedBy" json:"createdBy"`
	IsAnonymous *bool          `gorm:"default:true; not null" json:"isAnonymous"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

// PostPaginate ...
type PostPaginate struct {
	Total    int64       `json:"total"`
	PerPage  int         `json:"perPage"`
	Page     int         `json:"page"`
	LastPage int         `json:"lastPage"`
	Data     []*PostShow `json:"data"`
}

// PostShow
type PostShow struct {
	ID          string       `json:"id"`
	Category    PostCategory `json:"category"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Latitude    *float64     `json:"latitude"`
	Longitude   *float64     `json:"longitude"`
	Status      PostStatus   `json:"status"`
	CreatedBy   *User        `json:"createdBy"`
	CreatedAt   time.Time    `json:"createdAt"`
	UpdatedAt   time.Time    `json:"updatedAt"`
}

// PostCreateInput input
type PostCreateInput struct {
	CategoryID  string   `validate:"required" json:"categoryId"`
	Title       string   `validate:"required" json:"title"`
	Description string   `validate:"required" json:"description"`
	Latitude    *float64 `json:"latitude"`
	Longitude   *float64 `json:"longitude"`
	IsAnonymous *bool    `validate:"required" json:"isAnonymous"`
}

// PostUpdateInput input
type PostUpdateInput struct {
	CategoryID  string     `json:"categoryId"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Latitude    *float64   `json:"latitude"`
	Longitude   *float64   `json:"longitude"`
	Status      PostStatus `json:"status"`
	IsAnonymous *bool      `json:"isAnonymous"`
}

// PostDeleteInput input
type PostDeleteInput struct {
	ID string `json:"id"`
}

type PostStatus string

const (
	PostStatusOpen       PostStatus = "OPEN"
	PostStatusInProgress            = "IN_PROGRESS"
	PostStatusDone                  = "DONE"
	PostStatusClose                 = "CLOSE"
)
