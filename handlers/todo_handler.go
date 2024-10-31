package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
	"todo-app/config"
	"todo-app/models"
	"todo-app/services"
)

var validate = validator.New()

func CreateTodoHandler(c *gin.Context) {
	var todo models.TodoItem
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validate.Struct(todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"validation_error": err.Error()})
		return
	}

	if err := services.CreateTodoItem(&todo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, todo)
}

func GetAllTodosHandler(c *gin.Context) {
	todos, err := services.GetAllTodoItems()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todos)
}

func GetTodoByIDHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	todo, err := services.GetTodoItemByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	c.JSON(http.StatusOK, todo)
}

func UpdateTodoHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var todo models.TodoItem
	if err := config.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.UpdateTodoItem(&todo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todo)
}

func DeleteTodoHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := services.DeleteTodoItem(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
}
