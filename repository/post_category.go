package repository

import (
	"math"

	"github.com/dezenter/api/config"
	"github.com/dezenter/api/model"
	"github.com/dezenter/api/util"
	"gorm.io/gorm"
)

// PostCategoryRepository ...
type PostCategoryRepository struct {
	db *gorm.DB
}

// NewUserRepository ...
func NewPostCategoryRepository() *PostCategoryRepository {
	return &PostCategoryRepository{db: config.DB}
}

// Paginate post category
func (u *PostCategoryRepository) Paginate(page int, limit int) (*model.PostCategoryPaginate, error) {
	var c []*model.PostCategory
	var count int64

	var skip int
	if page == 1 {
		skip = 0
	} else {
		skip = limit * (page - 1)
	}
	u.db.Model(&model.PostCategory{}).Count(&count)
	u.db.Limit(limit).Offset(skip).Order("created_at desc").Find(&c)

	lastPage := int(math.Ceil(float64(count) / float64(limit)))
	p := model.PostCategoryPaginate{
		Total:          count,
		PerPage:        limit,
		Page:           page,
		LastPage:       lastPage,
		PostCategories: c,
	}
	return &p, nil
}

// Create post category
func (u *PostCategoryRepository) Create(input model.PostCategoryInput) (*model.PostCategory, error) {

	cID := util.GenerateID()

	c := model.PostCategory{
		ID:           cID,
		CategoryName: input.CategoryName,
	}

	r := u.db.Create(&c)
	if r.Error != nil {
		return nil, r.Error
	}

	r = u.db.Where("id = ?", cID).First(&c)
	if r.Error != nil {
		return nil, r.Error
	}

	return &c, nil
}

// FindByID post category
func (u *PostCategoryRepository) FindById(id string) (*model.PostCategory, error) {
	var c model.PostCategory
	r := u.db.Where("id = ?", id).First(&c)
	if r.Error != nil {
		return nil, r.Error
	}
	return &c, nil
}

// Update post category
func (u *PostCategoryRepository) Update(id string, input model.PostCategoryInput) (*model.PostCategory, error) {

	c := model.PostCategory{
		CategoryName: input.CategoryName,
	}
	r := u.db.Where("id = ?", id).Updates(&c)
	if r.Error != nil {
		return nil, r.Error
	}

	r = u.db.Where("id = ?", id).First(&c)
	if r.Error != nil {
		return nil, r.Error
	}
	return &c, nil
}

// Delete post category
func (u *PostCategoryRepository) Delete(id string) (bool, error) {
	var c model.PostCategory
	r := u.db.Model(&model.PostCategory{}).Where("id = ?", id).First(&c)
	if r.Error != nil {
		return false, r.Error
	}
	r = u.db.Where("id = ?", id).Delete(&model.PostCategory{})
	if r.Error != nil {
		return false, r.Error
	}
	return true, nil
}
