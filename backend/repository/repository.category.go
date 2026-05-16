package repository

import (
	"backend-coding-challenge/model"

	"gorm.io/gorm"
)


type categoryRepository interface {
	Create(todo *model.Todo) error
	Update(todo *model.Todo) error
	Delete(todo *model.Todo) error
	GetByID(id string) (*model.Todo, error)
	GetAll() ([]model.Todo, error)
}

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) Create(category *model.Category) error {
	return r.db.Create(category).Error
}

func (r *CategoryRepository) Update(category *model.Category) error {
	return r.db.Save(category).Error
}

func (r *CategoryRepository) Delete(category *model.Category) error {
	return r.db.Delete(category).Error
}

func (r *CategoryRepository) GetByID(id string) (*model.Category, error) {
	var category model.Category
	if err := r.db.First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *CategoryRepository) GetAll() ([]model.Category, error) {
	var categories []model.Category
	if err := r.db.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}



