package repositories

import (
	"github.com/dezenter/api/configs"
	"gorm.io/gorm"
)

// PostRepository ...
type PostRepository struct {
	db *gorm.DB
}

// NewPostRepository ...
func NewPostRepository() *PostRepository {
	return &PostRepository{db: configs.DB}
}

// Paginate Post
// func (r *PostRepository) Paginate(page int, limit int) (*models.PostPaginate, error) {
// 	var posts []*models.Post
// 	var showData []*models.PostShow
// 	var count int64
// 	if page == 0 {
// 		page = 1
// 	}
// 	var skip int
// 	if page == 1 {
// 		skip = 0
// 	} else {
// 		skip = limit * (page - 1)
// 	}
// 	r.db.Model((&models.Post{})).Count(&count)
// 	r.db.Preload("User").Limit(limit).Offset(skip).Order("created_at desc").Find(&posts)

// 	for _, p := range posts {
// 		if *p.IsAnonymous {
// 			showData = append(showData, &models.PostShow{
// 				ID:          p.ID,
// 				Title:       p.Title,
// 				Description: p.Description,
// 				Latitude:    p.Latitude,
// 				Longitude:   p.Longitude,
// 				Status:      p.Status,
// 				CreatedBy:   nil,
// 				CreatedAt:   p.CreatedAt,
// 				UpdatedAt:   p.UpdatedAt,
// 			})
// 		} else {
// 			showData = append(showData, &models.PostShow{
// 				ID:          p.ID,
// 				Title:       p.Title,
// 				Description: p.Description,
// 				Latitude:    p.Latitude,
// 				Longitude:   p.Longitude,
// 				Status:      p.Status,
// 				CreatedBy:   &p.User,
// 				CreatedAt:   p.CreatedAt,
// 				UpdatedAt:   p.UpdatedAt,
// 			})
// 		}

// 	}

// 	lastPage := int(math.Ceil(float64(count) / float64(limit)))
// 	p := models.PostPaginate{
// 		Total:    count,
// 		PerPage:  limit,
// 		Page:     page,
// 		LastPage: lastPage,
// 		Data:     showData,
// 	}
// 	return &p, nil
// }

// // Create Post
// func (r *PostRepository) Create(userId string, input models.PostCreateInput) (*models.PostShow, error) {
// 	var su models.PostShow
// 	id := utils.GenerateID()

// 	rp := models.Post{
// 		ID:          id,
// 		Title:       input.Title,
// 		Description: &input.Description,
// 		Longitude:   &input.Longitude,
// 		Latitude:    &input.Latitude,
// 		Status:      &input.Status,
// 		CreatedBy:   userId,
// 		IsAnonymous: input.IsAnonymous,
// 	}
// 	rs := r.db.Create(&rp)
// 	if rs.Error != nil {
// 		return nil, rs.Error
// 	}

// 	rs = r.db.Preload("User").Where("id = ?", id).First(&rp)
// 	if rs.Error != nil {
// 		return nil, rs.Error
// 	}

// 	if *rp.IsAnonymous {
// 		su = models.PostShow{
// 			ID:          id,
// 			Title:       rp.Title,
// 			Description: rp.Description,
// 			Longitude:   rp.Longitude,
// 			Latitude:    rp.Latitude,
// 			Status:      rp.Status,
// 			CreatedBy:   nil,
// 			CreatedAt:   rp.CreatedAt,
// 			UpdatedAt:   rp.UpdatedAt,
// 		}
// 	} else {
// 		su = models.PostShow{
// 			ID:          id,
// 			Title:       rp.Title,
// 			Description: rp.Description,
// 			Longitude:   rp.Longitude,
// 			Latitude:    rp.Latitude,
// 			Status:      rp.Status,
// 			CreatedBy:   &rp.User,
// 			CreatedAt:   rp.CreatedAt,
// 			UpdatedAt:   rp.UpdatedAt,
// 		}
// 	}

// 	return &su, nil
// }

// // Show Post
// func (r *PostRepository) FindById(id string) (*models.PostShow, error) {
// 	var rp models.Post
// 	var su models.PostShow
// 	rs := r.db.Preload("User").Where("id = ?", id).First(&rp)
// 	if rs.Error != nil {
// 		return nil, rs.Error
// 	}

// 	if *rp.IsAnonymous {
// 		su = models.PostShow{
// 			ID:          id,
// 			Title:       rp.Title,
// 			Description: rp.Description,
// 			Longitude:   rp.Longitude,
// 			Latitude:    rp.Latitude,
// 			Status:      rp.Status,
// 			CreatedBy:   nil,
// 			CreatedAt:   rp.CreatedAt,
// 			UpdatedAt:   rp.UpdatedAt,
// 		}
// 	} else {
// 		su = models.PostShow{
// 			ID:          id,
// 			Title:       rp.Title,
// 			Description: rp.Description,
// 			Longitude:   rp.Longitude,
// 			Latitude:    rp.Latitude,
// 			Status:      rp.Status,
// 			CreatedBy:   &rp.User,
// 			CreatedAt:   rp.CreatedAt,
// 			UpdatedAt:   rp.UpdatedAt,
// 		}
// 	}
// 	return &su, nil
// }

// // Update Post @todo check role is admin cant update status
// func (r *PostRepository) Update(userId string, input models.PostUpdateInput) (*models.PostShow, error) {
// 	var rp models.Post
// 	var su models.PostShow
// 	rs := r.db.Where("id = ?", input.ID).Where("created_by = ?", userId).First(&rp)
// 	if rs.Error != nil {
// 		return nil, rs.Error
// 	}

// 	rp = models.Post{
// 		Description: input.Description,
// 		Longitude:   input.Longitude,
// 		Latitude:    input.Latitude,
// 		Status:      input.Status,
// 		IsAnonymous: input.IsAnonymous,
// 	}
// 	rs = r.db.Model(&models.Post{}).Where("id = ?", input.ID).Updates(&rp)
// 	if rs.Error != nil {
// 		return nil, rs.Error
// 	}

// 	rs = r.db.Preload("User").Model(&models.Post{}).Where("id = ?", input.ID).First(&rp)
// 	if rs.Error != nil {
// 		return nil, rs.Error
// 	}

// 	if *rp.IsAnonymous {
// 		su = models.PostShow{
// 			ID:          rp.ID,
// 			Title:       rp.Title,
// 			Description: rp.Description,
// 			Longitude:   rp.Longitude,
// 			Latitude:    rp.Latitude,
// 			Status:      rp.Status,
// 			CreatedBy:   nil,
// 			CreatedAt:   rp.CreatedAt,
// 			UpdatedAt:   rp.UpdatedAt,
// 		}
// 	} else {
// 		su = models.PostShow{
// 			ID:          rp.ID,
// 			Title:       rp.Title,
// 			Description: rp.Description,
// 			Longitude:   rp.Longitude,
// 			Latitude:    rp.Latitude,
// 			Status:      rp.Status,
// 			CreatedBy:   &rp.User,
// 			CreatedAt:   rp.CreatedAt,
// 			UpdatedAt:   rp.UpdatedAt,
// 		}
// 	}
// 	return &su, nil
// }

// // Delete Post
// func (r *PostRepository) Delete(userId string, id string) (bool, error) {
// 	var rp models.Post
// 	rs := r.db.Model(&models.Post{}).Where("id = ?", id).Where("created_by = ?", userId).First(&rp)
// 	if rs.Error != nil {
// 		return false, rs.Error
// 	}
// 	rs = r.db.Where("id = ?", id).Delete(&models.Post{})
// 	if rs.Error != nil {
// 		return false, rs.Error
// 	}
// 	return true, nil
// }
