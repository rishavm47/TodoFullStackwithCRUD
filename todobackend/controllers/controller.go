package controllers

import (
	"example/data-access/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTodos(c *gin.Context) {
	var todoinput models.CreateTodoInput
	if err := c.ShouldBindJSON(&todoinput); err != nil {
		fmt.Println("error")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(todoinput)
	todo := models.Todo{Title: todoinput.Title, Completed: todoinput.Completed}

	err := models.Db.Create(&todo).Error
	if err != nil {
		return
	}
	c.JSON(http.StatusCreated, gin.H{"title": todo.Title, "completed": todo.Completed, "id": todo.ID})
}

func GetTodos(c *gin.Context) {
	var todos []models.Todo
	err := models.Db.Find(&todos).Error
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, todos)
}

func DeleteTodo(c *gin.Context) {
	var todo models.Todo
	if err := models.Db.Where("id = ?", c.Param("id")).First(&todo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.Db.Delete(&todo)

	c.JSON(http.StatusOK, gin.H{"data": "Deleted successfully"})
}

func UpdateTodo(c *gin.Context) {
	var todo models.Todo
	if err := models.Db.Where("id = ?", c.Param("id")).First(&todo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input models.UpdateTodoInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	todo.Title = input.Title
	todo.Completed = input.Completed
	models.Db.Save(&todo)
	c.JSON(http.StatusOK, todo)
}
