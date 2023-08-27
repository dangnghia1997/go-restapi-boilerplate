package service

import (
	"bookstore/internal/entity"
	"bookstore/internal/repository"
)

type TodoService interface {
	CreateTodo(title string, completed bool) error
	GetTodos() ([]entity.Todo, error)
}

type todoService struct {
	todoRepo repository.TodoRepository
}

func NewTodoService(todoRepo repository.TodoRepository) TodoService {
	return &todoService{todoRepo: todoRepo}
}

func (ts *todoService) CreateTodo(title string, completed bool) error {
	todo := &entity.Todo{Title: title, Completed: completed}
	return ts.todoRepo.Create(todo)
}

func (ts *todoService) GetTodos() ([]entity.Todo, error) {
	return ts.todoRepo.FindAll()
}
