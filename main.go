package main

import (
	"net/http"

	"github.com/davefredkove/go-crud/controllers"
	"github.com/davefredkove/go-crud/initializers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectDB()
}

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Hello World!"})
	})

	// Notes
	r.GET("/notes", controllers.GetNotes)
	r.GET("/notes/:id", controllers.GetNote)
	r.POST("/notes", controllers.CreateNote)
	r.PATCH("/notes/:id", controllers.UpdateNote)
	r.DELETE("/notes/:id", controllers.DeleteNote)

	// Todos
	r.GET("/todos", controllers.GetTodos)
	r.GET("/todos/:id", controllers.GetTodo)
	r.POST("/todos", controllers.CreateTodo)
	r.PATCH("/todos/:id", controllers.UpdateTodo)
	r.DELETE("/todos/:id", controllers.DeleteTodo)

	r.Run()
}
