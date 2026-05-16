package service

import (
	"backend-coding-challenge/model"
	"backend-coding-challenge/repository"
	"backend-coding-challenge/utils"
)

type TodoService struct {
	repo *repository.TodoRepository
}

func NewTodoService(repo *repository.TodoRepository) *TodoService {
	return &TodoService{repo: repo}
}

func (s *TodoService) Create(todo *model.Todo) error {
	return s.repo.Create(todo)
}

func (s *TodoService) Update(id int, todo *model.Todo) error {
	todo.ID = id
	return s.repo.Update(todo)
}

func (s *TodoService) Delete(todo *model.Todo) error {
	return s.repo.Delete(todo)
}

func (s *TodoService) GetByID(id string) (*model.Todo, error) {
	return s.repo.GetByID(id)
}

func (s *TodoService) GetAll(params utils.TodoFilterParam) (utils.FormatTodoResponse, error) {
	return s.repo.GetAll(params)
}
