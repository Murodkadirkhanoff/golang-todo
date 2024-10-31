package config

import (
	"fmt"
	"log"
	"todo-app/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error

	// Подключение к PostgreSQL
	dsn := "host=localhost user=laravel password=secret dbname=todo port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Миграция модели
	err = DB.AutoMigrate(&models.TodoItem{})
	if err != nil {
		return
	}
	fmt.Println("Database connected and migrated!")
}
