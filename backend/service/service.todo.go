package service

import (
	"backend-coding-challenge/model"
	"backend-coding-challenge/repository"
	"backend-coding-challenge/utils"
)

type TodoRepositoryInterface interface {
	Create(todo *model.Todo) error
	Update(todo *model.Todo) error
	Delete(todo *model.Todo) error
	GetByID(id string) (*model.Todo, error)
	GetAll(params utils.TodoFilterParam) (utils.FormatTodoResponse, error)
}

type TodoService struct {
	repo TodoRepositoryInterface
}

// NewTodoService creates a new TodoService with a concrete repository.
func NewTodoService(repo *repository.TodoRepository) *TodoService {
	return &TodoService{repo: repo}
}

// NewTodoServiceWithInterface creates a new TodoService with any implementation of TodoRepositoryInterface.
func NewTodoServiceWithInterface(repo TodoRepositoryInterface) *TodoService {
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
