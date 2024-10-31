package main

import (
	"github.com/gin-gonic/gin"
	"todo-app/config"
	"todo-app/routes"
)

func main() {
	// Инициализация базы данных
	config.InitDB()

	// Настройка маршрутов
	router := gin.Default()
	routes.SetupRoutes(router)

	// Запуск сервера
	err := router.Run(":8090")
	if err != nil {
		return
	}
}
