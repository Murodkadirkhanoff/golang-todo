package routes

import (
	"github.com/gin-gonic/gin"
	"todo-app/handlers"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.POST("/todos", handlers.CreateTodoHandler)
		api.GET("/todos", handlers.GetAllTodosHandler)
		api.GET("/todos/:id", handlers.GetTodoByIDHandler)
		api.PUT("/todos/:id", handlers.UpdateTodoHandler)
		api.DELETE("/todos/:id", handlers.DeleteTodoHandler)
	}
}
