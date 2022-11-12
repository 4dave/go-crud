package controllers

import (
	"net/http"

	"github.com/davefredkove/go-crud/initializers"
	"github.com/davefredkove/go-crud/models"
	"github.com/gin-gonic/gin"
)

func GetContacts(c *gin.Context) {
	var Contacts []models.Contacts
	initializers.DB.Order("created_at desc").Find(&Contacts)
	c.JSON(http.StatusOK, gin.H{"data": Contacts})
}

func GetContact(c *gin.Context) {
	var contact models.Contacts
	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&contact).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": contact})
}

func CreateContact(c *gin.Context) {
	var input models.Contacts
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	contact := models.Contacts{FirstName: input.FirstName, LastName: input.LastName, Email: input.Email, Phone: input.Phone}
	initializers.DB.Create(&contact)
	c.JSON(http.StatusOK, gin.H{"data": contact})
}

func UpdateContact(c *gin.Context) {
	var contact models.Contacts
	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&contact).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	var input models.Contacts
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	initializers.DB.Model(&contact).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": contact})
}

func DeleteContact(c *gin.Context) {
	var contact models.Contacts
	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&contact).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	initializers.DB.Delete(&contact)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
