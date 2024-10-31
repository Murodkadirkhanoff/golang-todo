package services

import (
	"todo-app/config"
	"todo-app/models"
)

func CreateTodoItem(todo *models.TodoItem) error {
	return config.DB.Create(todo).Error
}

func GetAllTodoItems() ([]models.TodoItem, error) {
	var todos []models.TodoItem
	err := config.DB.Find(&todos).Error
	return todos, err
}

func GetTodoItemByID(id uint) (models.TodoItem, error) {
	var todo models.TodoItem
	err := config.DB.First(&todo, id).Error
	return todo, err
}

func UpdateTodoItem(todo *models.TodoItem) error {
	return config.DB.Save(todo).Error
}

func DeleteTodoItem(id uint) error {
	return config.DB.Delete(&models.TodoItem{}, id).Error
}
