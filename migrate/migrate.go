package main

import (
	"github.com/davefredkove/go-crud/initializers"
	"github.com/davefredkove/go-crud/models"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Note{})
	initializers.DB.AutoMigrate(&models.Todo{})
	initializers.DB.AutoMigrate(&models.Contacts{})
}
