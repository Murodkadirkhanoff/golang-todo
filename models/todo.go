package models

import "gorm.io/gorm"

type TodoItem struct {
	gorm.Model
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}
