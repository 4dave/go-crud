package controllers

import (
	"net/http"

	"github.com/davefredkove/go-crud/initializers"
	"github.com/davefredkove/go-crud/models"
	"github.com/gin-gonic/gin"
)

func GetNotes(c *gin.Context) {
	var notes []models.Note
	initializers.DB.Order("created_at desc").Find(&notes)
	c.JSON(http.StatusOK, gin.H{"data": notes})
}

func GetNote(c *gin.Context) {
	var note models.Note
	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&note).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": note})
}

func CreateNote(c *gin.Context) {
	var input models.Note
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	note := models.Note{Title: input.Title, Body: input.Body}
	initializers.DB.Create(&note)
	c.JSON(http.StatusOK, gin.H{"data": note})
}

func UpdateNote(c *gin.Context) {
	var note models.Note
	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&note).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	var input models.Note
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	initializers.DB.Model(&note).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": note})
}

func DeleteNote(c *gin.Context) {
	var note models.Note
	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&note).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	initializers.DB.Delete(&note)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
