package models

type PostCategory struct {
	ID           string `gorm:"primaryKey" json:"id"`
	CategoryName string `gorm:"not null" json:"categoryName"`
}

// PostCategoryPaginate pagination
type PostCategoryPaginate struct {
	Total          int64 `json:"total"`
	PerPage        int   `json:"perPage"`
	Page           int   `json:"page"`
	LastPage       int   `json:"lastPage"`
	PostCategories []*PostCategory
}

type PostCategoryInput struct {
	CategoryName string `validate:"required" json:"categoryName"`
}
