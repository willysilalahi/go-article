package repository

import (
	"go-article/model"

	"gorm.io/gorm"
)

type Repository interface {
	GetAll() ([]model.Post, error)
	// GetById(id int) (model.Post, error)
	// Create(post model.Post) (model.Post, error)
	// Update(id int, post model.Post) (model.Post, error)
	// Delete(id int) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAll() ([]model.Post, error) {
	var posts []model.Post
	err := r.db.Preload("Admin").Preload("Tags").Preload("Comments").Find(&posts).Error
	return posts, err
}
