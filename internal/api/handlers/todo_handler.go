package handlers

import (
	"bookstore/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TodoHandler struct {
	todoService service.TodoService
}

func NewTodoHandler(todoService service.TodoService) *TodoHandler {
	return &TodoHandler{todoService: todoService}
}

func (th *TodoHandler) CreateTodo(c *gin.Context) {
	var request struct {
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := th.todoService.CreateTodo(request.Title, request.Completed); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.Status(http.StatusCreated)
}

func (th *TodoHandler) GetTodos(c *gin.Context) {
	todos, err := th.todoService.GetTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, todos)
}
