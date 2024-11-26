package main

import (
	"log"
	"github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"go-vue-todo/models"
	"go-vue-todo/handlers"
)


func main() {
	// db init
	db, err := gorm.Open(sqlite.Open("todos.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}	
	
    // auto-migrate the database schema
    err = db.AutoMigrate(&models.Todo{})
    if err != nil {
        log.Fatal("Failed to migrate database schema:", err)   
    }

    r := gin.Default()

    // CORS middleware
    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:8080"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
    }))

    api := r.Group("/api") 
    {
        api.GET("/todos", handlers.GetTodos(db))
        api.POST("/todos", handlers.CreateTodo(db))
        api.PUT("/todos/:id", handlers.UpdateTodo(db))
        api.DELETE("/todos/:id", handlers.DeleteTodo(db))
    }

    r.Run(":3000")

}
