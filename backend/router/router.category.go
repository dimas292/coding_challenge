package router

import (
	"backend-coding-challenge/handler"

	"github.com/gin-gonic/gin"
)

type CategoryRouter struct {
	CategoryHandler *handler.CategoryHandler
}

func NewCategoryRouter(handler *handler.CategoryHandler) *CategoryRouter {
	return &CategoryRouter{CategoryHandler: handler}
}

func (r *CategoryRouter) InitCategoryRoutes(rg *gin.RouterGroup) {
	rg.POST("/", r.CategoryHandler.Create)
	rg.GET("/", r.CategoryHandler.GetAll)
	rg.GET("/:id", r.CategoryHandler.GetByID)
	rg.PUT("/", r.CategoryHandler.Update)
	rg.DELETE("/:id", r.CategoryHandler.Delete)
}
