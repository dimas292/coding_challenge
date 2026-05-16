package service

import (
	"backend-coding-challenge/model"
	"backend-coding-challenge/repository"
)

type CategoryService struct {
	repo *repository.CategoryRepository
}

func NewCategoryService(repo *repository.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) Create(category *model.Category) error {
	return s.repo.Create(category)
}

func (s *CategoryService) Update(category *model.Category) error {
	return s.repo.Update(category)
}

func (s *CategoryService) Delete(category *model.Category) error {
	return s.repo.Delete(category)
}

func (s *CategoryService) GetByID(id string) (*model.Category, error) {
	return s.repo.GetByID(id)
}

func (s *CategoryService) GetAll() ([]model.Category, error) {
	return s.repo.GetAll()
}
