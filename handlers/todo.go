package handlers

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"go-vue-todo/models"
)

func GetTodos(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var todos []models.Todo
		db.Find(&todos)
		c.JSON(http.StatusOK, todos)
	}
}

func CreateTodo(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newTodo models.Todo
		if err := c.ShouldBindJSON(&newTodo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		db.Create(&newTodo)
		c.JSON(http.StatusCreated, newTodo)
	}
}

func UpdateTodo(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		var todo models.Todo

		if err := db.First(&todo, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
			return
		}

		if err := c.ShouldBindJSON(&todo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		db.Save(&todo)
		c.JSON(http.StatusOK, todo)
	}
}

func DeleteTodo(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		var todo models.Todo

		if err := db.First(&todo, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
			return
		}

		db.Delete(&todo)
		c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
	}
}
