package repository

import (
	"bookstore/internal/entity"
	"gorm.io/gorm"
)

type TodoRepository interface {
	Create(todo *entity.Todo) error
	FindAll() ([]entity.Todo, error)
}

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &todoRepository{db: db}
}

func (tr *todoRepository) Create(todo *entity.Todo) error {
	return tr.db.Create(todo).Error
}

func (tr *todoRepository) FindAll() ([]entity.Todo, error) {
	var todos []entity.Todo
	err := tr.db.Find(&todos).Error
	return todos, err
}
