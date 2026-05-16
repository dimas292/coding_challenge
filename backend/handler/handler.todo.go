package handler

import (
	"backend-coding-challenge/model"
	"backend-coding-challenge/service"
	"backend-coding-challenge/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	TodoService *service.TodoService
	CategoryService *service.CategoryService
}

func NewTodoHandler(service *service.TodoService, categoryService *service.CategoryService) *TodoHandler {
	return &TodoHandler{TodoService: service, CategoryService: categoryService}
}

func (h *TodoHandler) Create(c *gin.Context) {
	var todo model.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	
	if err := h.TodoService.Create(&todo); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse{
		Message: "Todo created successfully",
		Data:    todo,
	})
}

func (h *TodoHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	todo, err := h.TodoService.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, utils.SuccessResponse{
		Message: "Todo retrieved successfully",
		Data:    todo,
	})
}

func (h *TodoHandler) GetAll(c *gin.Context) {
	var params utils.TodoFilterParam
	if val := c.Query("search"); val != "" {
		params.Search = val
	}
	if val := c.Query("page"); val != "" {
		params.Page, _ = strconv.ParseInt(val, 10, 64)
	}

	if val := c.Query("limit"); val != "" {
		params.Limit, _ = strconv.ParseInt(val, 10, 64)
	}

	if val := c.Query("category"); val != "" {
		params.Category, _ = strconv.ParseInt(val, 10, 64)
	}

	if val := c.Query("completed"); val != "" {
		parsed, _ := strconv.ParseBool(val)
		params.Completed = &parsed
	}

	if val := c.Query("priority"); val != "" {
		params.Priority = val
	}

	if val := c.Query("sort_by"); val != "" {
		params.SortBy = val
	}

	if val := c.Query("order_by"); val != "" {
		params.OrderBy = val
	}

	response, err := h.TodoService.GetAll(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}

func (h *TodoHandler) Update(c *gin.Context) {
	var todo model.Todo

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// id 
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	category, err := h.CategoryService.GetByID(strconv.Itoa(int(todo.CategoryID)))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Category not found"})
		return
	}
	
	todo.Category.ID = category.ID
	todo.Category.Name = category.Name
	todo.Category.Color = category.Color
	todo.Category.CreatedAt = category.CreatedAt

	if err := h.TodoService.Update(idInt, &todo); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse{
		Message: "Todo updated successfully",
		Data:    todo,
	})
}

func (h *TodoHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: "ID is required",
		})
		return
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Message: "Invalid ID",
		})
		return
	}
	todo := model.Todo{ID: idInt}
	if err := h.TodoService.Delete(&todo); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, utils.SuccessResponse{
		Message: "Todo deleted successfully",
		Data:    todo,
	})
}

