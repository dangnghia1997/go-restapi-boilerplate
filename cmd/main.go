package main

import (
	"bookstore/internal/api"
	"bookstore/internal/entity"
	"bookstore/internal/repository"
	"bookstore/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

import (
	"bookstore/internal/config"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		panic("failed to load config")
	}

	db, err := gorm.Open(mysql.Open(cfg.DBConnectionString), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Auto Migrate - Creates or updates the database tables
	err = db.AutoMigrate(&entity.Todo{})
	if err != nil {
		panic("failed to migrate database")
	}

	todoRepo := repository.NewTodoRepository(db)
	todoService := service.NewTodoService(todoRepo)

	r := gin.Default()
	api.SetupRoutes(r, todoService)

	r.Run(cfg.ServerPort)
}
