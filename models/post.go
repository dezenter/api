package models

import (
	"time"

	"gorm.io/gorm"
)

// Post model
type Post struct {
	ID          string         `gorm:"primaryKey" json:"id"`
	Title       string         `gorm:"not null" json:"title"`
	Description *string        `json:"description"`
	Longitude   *float64       `json:"longitude"`
	Latitude    *float64       `json:"latitude"`
	Status      *PostStatus    `gorm:"not null" json:"status"`
	CreatedBy   string         `gorm:"not null"`
	User        User           `gorm:"foreignKey:CreatedBy" json:"createdBy"`
	IsAnonymous *bool          `gorm:"not null" json:"isAnonymous"`
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
	ID          string      `json:"id"`
	Title       string      `json:"title"`
	Description *string     `json:"description"`
	Latitude    *float64    `json:"latitude"`
	Longitude   *float64    `json:"longitude"`
	Status      *PostStatus `json:"status"`
	CreatedBy   *User       `json:"createdBy"`
	CreatedAt   time.Time   `json:"createdAt"`
	UpdatedAt   time.Time   `json:"updatedAt"`
}

// Comment model
type Comment struct {
	ID          string    `gorm:"primaryKey" json:"id"`
	Comment     string    `json:"comment"`
	CreatedBy   *string   `gorm:"not null" json:"createdBy"`
	IsAnonymous *bool     `gorm:"not null" json:"isAnonymous"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// PostFile model
type PostFile struct {
	ID     string `gorm:"primaryKey" json:"id"`
	PostID string
	FileID string
}

// PostCreateInput input
type PostCreateInput struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Latitude    float64    `json:"latitude"`
	Longitude   float64    `json:"longitude"`
	Status      PostStatus `json:"status"`
	IsAnonymous *bool      `json:"isAnonymous"`
}

// PostUpdateInput input
type PostUpdateInput struct {
	ID          string      `json:"id"`
	Title       *string     `json:"title"`
	Description *string     `json:"description"`
	Latitude    *float64    `json:"latitude"`
	Longitude   *float64    `json:"longitude"`
	Status      *PostStatus `json:"status"`
	IsAnonymous *bool       `json:"isAnonymous"`
}

// PostDeleteInput input
type PostDeleteInput struct {
	ID string `json:"id"`
}

type PostStatus string

const (
	PostStatusOpen       PostStatus = "OPEN"
	PostStatusInProgress PostStatus = "IN_PROGRESS"
	PostStatusDone       PostStatus = "DONE"
	PostStatusClose      PostStatus = "CLOSE"
)
