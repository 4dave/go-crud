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

	// only for testing - allow from localhost
	// r.Use(cors.New(cors.Config{
	// 	AllowOrigins: []string{"http://localhost:3000", "http://localhost:3001"},
	// 	AllowMethods: []string{"GET", "POST", "PATCH", "DELETE"},
	// }))

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

	// Contacts
	r.GET("/contacts", controllers.GetContacts)
	r.GET("/contacts/:id", controllers.GetContact)
	r.POST("/contacts", controllers.CreateContact)
	r.PATCH("/contacts/:id", controllers.UpdateContact)
	r.DELETE("/contacts/:id", controllers.DeleteContact)

	r.Run()
}
