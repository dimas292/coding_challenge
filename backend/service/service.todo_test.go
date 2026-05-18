package service

import (
	"backend-coding-challenge/model"
	"backend-coding-challenge/utils"
	"errors"
	"testing"
	"time"
)

type mockTodoRepository struct {
	createFn  func(todo *model.Todo) error
	updateFn  func(todo *model.Todo) error
	deleteFn  func(todo *model.Todo) error
	getByIDFn func(id string) (*model.Todo, error)
	getAllFn   func(params utils.TodoFilterParam) (utils.FormatTodoResponse, error)
}

func (m *mockTodoRepository) Create(todo *model.Todo) error {
	if m.createFn != nil {
		return m.createFn(todo)
	}
	return nil
}

func (m *mockTodoRepository) Update(todo *model.Todo) error {
	if m.updateFn != nil {
		return m.updateFn(todo)
	}
	return nil
}

func (m *mockTodoRepository) Delete(todo *model.Todo) error {
	if m.deleteFn != nil {
		return m.deleteFn(todo)
	}
	return nil
}

func (m *mockTodoRepository) GetByID(id string) (*model.Todo, error) {
	if m.getByIDFn != nil {
		return m.getByIDFn(id)
	}
	return nil, nil
}

func (m *mockTodoRepository) GetAll(params utils.TodoFilterParam) (utils.FormatTodoResponse, error) {
	if m.getAllFn != nil {
		return m.getAllFn(params)
	}
	return utils.FormatTodoResponse{}, nil
}

func ptrBool(b bool) *bool { return &b }

func sampleTodo() *model.Todo {
	now := time.Now()
	return &model.Todo{
		ID:          1,
		Title:       "Test Todo",
		Description: "Test Description",
		CategoryID:  1,
		Category: model.Category{
			ID:   1,
			Name: "Work",
		},
		Priority:  "high",
		Completed: false,
		DueDate:   now.Add(24 * time.Hour),
		CreatedAt: &now,
	}
}

func TestCreate_Success(t *testing.T) {
	todo := sampleTodo()

	mock := &mockTodoRepository{
		createFn: func(td *model.Todo) error {
			return nil
		},
	}

	svc := NewTodoServiceWithInterface(mock)
	err := svc.Create(todo)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestCreate_RepoError(t *testing.T) {
	todo := sampleTodo()
	expectedErr := errors.New("database connection failed")

	mock := &mockTodoRepository{
		createFn: func(td *model.Todo) error {
			return expectedErr
		},
	}

	svc := NewTodoServiceWithInterface(mock)
	err := svc.Create(todo)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if err.Error() != expectedErr.Error() {
		t.Fatalf("expected error %q, got %q", expectedErr.Error(), err.Error())
	}
}

func TestCreate_PassesCorrectTodo(t *testing.T) {
	todo := sampleTodo()
	var received *model.Todo

	mock := &mockTodoRepository{
		createFn: func(td *model.Todo) error {
			received = td
			return nil
		},
	}

	svc := NewTodoServiceWithInterface(mock)
	_ = svc.Create(todo)

	if received == nil {
		t.Fatal("expected repository to receive todo, got nil")
	}
	if received.Title != todo.Title {
		t.Fatalf("expected title %q, got %q", todo.Title, received.Title)
	}
	if received.Description != todo.Description {
		t.Fatalf("expected description %q, got %q", todo.Description, received.Description)
	}
}

func TestUpdate_Success(t *testing.T) {
	todo := sampleTodo()

	mock := &mockTodoRepository{
		updateFn: func(td *model.Todo) error {
			return nil
		},
	}

	svc := NewTodoServiceWithInterface(mock)
	err := svc.Update(5, todo)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestUpdate_SetsIDOnTodo(t *testing.T) {
	todo := sampleTodo()
	var received *model.Todo

	mock := &mockTodoRepository{
		updateFn: func(td *model.Todo) error {
			received = td
			return nil
		},
	}

	svc := NewTodoServiceWithInterface(mock)
	_ = svc.Update(99, todo)

	if received == nil {
		t.Fatal("expected repository to receive todo, got nil")
	}
	if received.ID != 99 {
		t.Fatalf("expected todo ID to be 99, got %d", received.ID)
	}
}

func TestUpdate_RepoError(t *testing.T) {
	todo := sampleTodo()
	expectedErr := errors.New("record not found")

	mock := &mockTodoRepository{
		updateFn: func(td *model.Todo) error {
			return expectedErr
		},
	}

	svc := NewTodoServiceWithInterface(mock)
	err := svc.Update(1, todo)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if err.Error() != expectedErr.Error() {
		t.Fatalf("expected error %q, got %q", expectedErr.Error(), err.Error())
	}
}

func TestUpdate_PreservesOtherFields(t *testing.T) {
	todo := sampleTodo()
	todo.Title = "Updated Title"
	todo.Completed = true
	var received *model.Todo

	mock := &mockTodoRepository{
		updateFn: func(td *model.Todo) error {
			received = td
			return nil
		},
	}

	svc := NewTodoServiceWithInterface(mock)
	_ = svc.Update(10, todo)

	if received.Title != "Updated Title" {
		t.Fatalf("expected title %q, got %q", "Updated Title", received.Title)
	}
	if !received.Completed {
		t.Fatal("expected Completed to be true")
	}
	if received.ID != 10 {
		t.Fatalf("expected ID 10, got %d", received.ID)
	}
}

func TestDelete_Success(t *testing.T) {
	todo := sampleTodo()

	mock := &mockTodoRepository{
		deleteFn: func(td *model.Todo) error {
			return nil
		},
	}

	svc := NewTodoServiceWithInterface(mock)
	err := svc.Delete(todo)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestDelete_RepoError(t *testing.T) {
	todo := sampleTodo()
	expectedErr := errors.New("cannot delete: foreign key constraint")

	mock := &mockTodoRepository{
		deleteFn: func(td *model.Todo) error {
			return expectedErr
		},
	}

	svc := NewTodoServiceWithInterface(mock)
	err := svc.Delete(todo)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if err.Error() != expectedErr.Error() {
		t.Fatalf("expected error %q, got %q", expectedErr.Error(), err.Error())
	}
}

func TestDelete_PassesCorrectTodo(t *testing.T) {
	todo := sampleTodo()
	todo.ID = 42
	var received *model.Todo

	mock := &mockTodoRepository{
		deleteFn: func(td *model.Todo) error {
			received = td
			return nil
		},
	}

	svc := NewTodoServiceWithInterface(mock)
	_ = svc.Delete(todo)

	if received == nil {
		t.Fatal("expected repository to receive todo")
	}
	if received.ID != 42 {
		t.Fatalf("expected ID 42, got %d", received.ID)
	}
}

func TestGetByID_Success(t *testing.T) {
	expected := sampleTodo()

	mock := &mockTodoRepository{
		getByIDFn: func(id string) (*model.Todo, error) {
			return expected, nil
		},
	}

	svc := NewTodoServiceWithInterface(mock)
	result, err := svc.GetByID("1")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected result, got nil")
	}
	if result.Title != expected.Title {
		t.Fatalf("expected title %q, got %q", expected.Title, result.Title)
	}
}

func TestGetByID_NotFound(t *testing.T) {
	expectedErr := errors.New("record not found")

	mock := &mockTodoRepository{
		getByIDFn: func(id string) (*model.Todo, error) {
			return nil, expectedErr
		},
	}

	svc := NewTodoServiceWithInterface(mock)
	result, err := svc.GetByID("999")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if result != nil {
		t.Fatalf("expected nil result, got %v", result)
	}
	if err.Error() != expectedErr.Error() {
		t.Fatalf("expected error %q, got %q", expectedErr.Error(), err.Error())
	}
}

func TestGetByID_PassesCorrectID(t *testing.T) {
	var receivedID string

	mock := &mockTodoRepository{
		getByIDFn: func(id string) (*model.Todo, error) {
			receivedID = id
			return sampleTodo(), nil
		},
	}

	svc := NewTodoServiceWithInterface(mock)
	_, _ = svc.GetByID("55")

	if receivedID != "55" {
		t.Fatalf("expected ID %q, got %q", "55", receivedID)
	}
}

func TestGetAll_Success(t *testing.T) {
	todos := []model.Todo{*sampleTodo()}
	expectedResp := utils.FormatTodoResponse{
		Data: todos,
		Pagination: utils.Pagination{
			CurrentPage: 1,
			PerPage:     10,
			Total:       1,
			TotalPages:  1,
		},
	}

	mock := &mockTodoRepository{
		getAllFn: func(params utils.TodoFilterParam) (utils.FormatTodoResponse, error) {
			return expectedResp, nil
		},
	}

	svc := NewTodoServiceWithInterface(mock)
	params := utils.TodoFilterParam{Page: 1, Limit: 10}
	result, err := svc.GetAll(params)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result.Pagination.Total != 1 {
		t.Fatalf("expected total 1, got %d", result.Pagination.Total)
	}
	if result.Pagination.CurrentPage != 1 {
		t.Fatalf("expected page 1, got %d", result.Pagination.CurrentPage)
	}
}

func TestGetAll_RepoError(t *testing.T) {
	expectedErr := errors.New("database timeout")

	mock := &mockTodoRepository{
		getAllFn: func(params utils.TodoFilterParam) (utils.FormatTodoResponse, error) {
			return utils.FormatTodoResponse{}, expectedErr
		},
	}

	svc := NewTodoServiceWithInterface(mock)
	_, err := svc.GetAll(utils.TodoFilterParam{})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if err.Error() != expectedErr.Error() {
		t.Fatalf("expected error %q, got %q", expectedErr.Error(), err.Error())
	}
}

func TestGetAll_EmptyResult(t *testing.T) {
	mock := &mockTodoRepository{
		getAllFn: func(params utils.TodoFilterParam) (utils.FormatTodoResponse, error) {
			return utils.FormatTodoResponse{
				Data: []model.Todo{},
				Pagination: utils.Pagination{
					CurrentPage: 1,
					PerPage:     10,
					Total:       0,
					TotalPages:  0,
				},
			}, nil
		},
	}

	svc := NewTodoServiceWithInterface(mock)
	result, err := svc.GetAll(utils.TodoFilterParam{Page: 1, Limit: 10})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result.Pagination.Total != 0 {
		t.Fatalf("expected total 0, got %d", result.Pagination.Total)
	}
}

func TestGetAll_PassesFilterParams(t *testing.T) {
	completed := true
	var receivedParams utils.TodoFilterParam

	mock := &mockTodoRepository{
		getAllFn: func(params utils.TodoFilterParam) (utils.FormatTodoResponse, error) {
			receivedParams = params
			return utils.FormatTodoResponse{}, nil
		},
	}

	svc := NewTodoServiceWithInterface(mock)
	inputParams := utils.TodoFilterParam{
		Page:      2,
		Limit:     20,
		Category:  3,
		Completed: &completed,
		Priority:  "high",
		SortBy:    "created_at",
		OrderBy:   "desc",
		Search:    "test",
	}
	_, _ = svc.GetAll(inputParams)

	if receivedParams.Page != 2 {
		t.Fatalf("expected page 2, got %d", receivedParams.Page)
	}
	if receivedParams.Limit != 20 {
		t.Fatalf("expected limit 20, got %d", receivedParams.Limit)
	}
	if receivedParams.Category != 3 {
		t.Fatalf("expected category 3, got %d", receivedParams.Category)
	}
	if receivedParams.Completed == nil || *receivedParams.Completed != true {
		t.Fatal("expected completed to be true")
	}
	if receivedParams.Priority != "high" {
		t.Fatalf("expected priority %q, got %q", "high", receivedParams.Priority)
	}
	if receivedParams.SortBy != "created_at" {
		t.Fatalf("expected sort_by %q, got %q", "created_at", receivedParams.SortBy)
	}
	if receivedParams.OrderBy != "desc" {
		t.Fatalf("expected order_by %q, got %q", "desc", receivedParams.OrderBy)
	}
	if receivedParams.Search != "test" {
		t.Fatalf("expected search %q, got %q", "test", receivedParams.Search)
	}
}

func TestGetAll_WithPagination(t *testing.T) {
	mock := &mockTodoRepository{
		getAllFn: func(params utils.TodoFilterParam) (utils.FormatTodoResponse, error) {
			return utils.FormatTodoResponse{
				Data: []model.Todo{*sampleTodo()},
				Pagination: utils.Pagination{
					CurrentPage: 3,
					PerPage:     5,
					Total:       25,
					TotalPages:  5,
				},
			}, nil
		},
	}

	svc := NewTodoServiceWithInterface(mock)
	result, err := svc.GetAll(utils.TodoFilterParam{Page: 3, Limit: 5})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result.Pagination.CurrentPage != 3 {
		t.Fatalf("expected page 3, got %d", result.Pagination.CurrentPage)
	}
	if result.Pagination.TotalPages != 5 {
		t.Fatalf("expected 5 total pages, got %d", result.Pagination.TotalPages)
	}
	if result.Pagination.Total != 25 {
		t.Fatalf("expected 25 total items, got %d", result.Pagination.Total)
	}
}


func TestNewTodoServiceWithInterface(t *testing.T) {
	mock := &mockTodoRepository{}
	svc := NewTodoServiceWithInterface(mock)
	if svc == nil {
		t.Fatal("expected non-nil service")
	}
	if svc.repo == nil {
		t.Fatal("expected non-nil repo")
	}
}
