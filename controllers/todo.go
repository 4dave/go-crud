package controllers

import (
	"net/http"

	"github.com/davefredkove/go-crud/initializers"
	"github.com/davefredkove/go-crud/models"
	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	var todos []models.Todo
	initializers.DB.Order("created_at desc").Find(&todos)
	c.JSON(http.StatusOK, gin.H{"data": todos})
}

func GetTodo(c *gin.Context) {
	var todo models.Todo
	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&todo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": todo})
}

func CreateTodo(c *gin.Context) {
	var input models.Todo
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	todo := models.Todo{Name: input.Name, Completed: input.Completed}
	initializers.DB.Create(&todo)
	c.JSON(http.StatusOK, gin.H{"data": todo})
}

func UpdateTodo(c *gin.Context) {
	var todo models.Todo
	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&todo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	var input models.Todo
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	initializers.DB.Model(&todo).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": todo})
}

func DeleteTodo(c *gin.Context) {
	var todo models.Todo
	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&todo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	initializers.DB.Delete(&todo)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
