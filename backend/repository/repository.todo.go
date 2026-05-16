package repository

import (
	"backend-coding-challenge/model"
	"backend-coding-challenge/utils"
	"math"

	"gorm.io/gorm"
)


type todoRepository interface {
	Create(todo *model.Todo) error
	Update(todo *model.Todo) error
	Delete(todo *model.Todo) error
	GetByID(id string) (*model.Todo, error)
	GetAll() (utils.FormatTodoResponse, error)
}

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *TodoRepository {
	return &TodoRepository{db: db}
}

func (r *TodoRepository) Create(todo *model.Todo) error {
	return r.db.Create(todo).Error
}

func (r *TodoRepository) Update(todo *model.Todo) error {
	return r.db.Save(todo).Error
}

func (r *TodoRepository) Delete(todo *model.Todo) error {
	return r.db.Delete(todo).Error
}

func (r *TodoRepository) GetByID(id string) (*model.Todo, error) {
	var todo model.Todo
	if err := r.db.Preload("Category").First(&todo, id).Error; err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r *TodoRepository) GetAll(params utils.TodoFilterParam) (utils.FormatTodoResponse, error) {
	var todos []model.Todo
	var total int64
	
	if params.Page < 0 {
		params.Page = 1
	}
	
	if params.Limit == 0 {
		params.Limit = 10
	}
	

	query := r.db.Model(&model.Todo{})
	if params.Category != 0 {
		query = query.Where("category_id = ?", params.Category)
	}
	if params.Completed != nil {
		query = query.Where("completed = ?", *params.Completed)
	}
	if params.Priority != "" {
		query = query.Where("priority = ?", params.Priority)
	}
	if params.Search != "" {
		query = query.Where("title ILIKE ?", "%"+params.Search+"%")
	}

	query.Count(&total)

	if params.SortBy != "" {
		orderDirection := "ASC"
		if params.OrderBy == "desc" || params.OrderBy == "DESC" {
			orderDirection = "DESC"
		}
		orderString := params.SortBy + " " + orderDirection
		query = query.Order(orderString)
	} else {
		query = query.Order("created_at DESC") 
	}

	offset := (params.Page - 1) * params.Limit

	if err := query.Limit(int(params.Limit)).Offset(int(offset)).Preload("Category").Find(&todos).Error; err != nil {
		return utils.FormatTodoResponse{}, err
	}

	totalPages := int(math.Ceil(float64(total) / float64(params.Limit)))
	currentPage := int(params.Page)
	
	
	response := utils.FormatTodoResponse{
		Data: todos,
		Pagination: utils.Pagination{
			CurrentPage: currentPage,
			PerPage:  int(params.Limit),
			Total: int(total),
			TotalPages: totalPages,
		},
	}

	
	return response, nil
}



