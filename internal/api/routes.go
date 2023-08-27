package api

import (
	"bookstore/internal/api/handlers"
	"bookstore/internal/service"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, todoService service.TodoService) {
	todoHandler := handlers.NewTodoHandler(todoService)

	v1 := r.Group("/api/v1")
	{

		v1.POST("/todos", todoHandler.CreateTodo)
		v1.GET("/todos", todoHandler.GetTodos)
	}
}
