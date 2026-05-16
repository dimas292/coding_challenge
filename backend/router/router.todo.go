package router

import (
	"backend-coding-challenge/handler"

	"github.com/gin-gonic/gin"
)

type TodoRouter struct {
	Handler *handler.TodoHandler
}

func NewTodoRouter(handler *handler.TodoHandler) *TodoRouter {
	return &TodoRouter{Handler: handler}
}

func (r *TodoRouter) InitTodoRoutes(rg *gin.RouterGroup) {
	rg.POST("/", r.Handler.Create)
	rg.GET("/", r.Handler.GetAll)
	rg.GET("/:id", r.Handler.GetByID)
	rg.PUT("/:id", r.Handler.Update)
	rg.DELETE("/:id", r.Handler.Delete)
}
