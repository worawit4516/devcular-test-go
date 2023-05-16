package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"github.com/worawit4516/devcular-test-go/src/config"
	"github.com/worawit4516/devcular-test-go/src/models"
	"gorm.io/gorm"
)

var db *gorm.DB = config.ConnectDatabase()

func CheckHealth(context *gin.Context) {
	context.JSON(http.StatusOK, "Backend is Good!!")
}

func CreateTodo(context *gin.Context) {
	var data models.TodoRequest

	if err := context.ShouldBindJSON(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	todo := models.Todo{}
	todo.Name = data.Name
	todo.Description = data.Description

	result := db.Create(&todo)

	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Error": " Can't save new todo"})
		return
	}

	var response models.TodoResponse
	response.ID = todo.ID
	response.Name = todo.Name
	response.Description = todo.Description

	context.JSON(http.StatusCreated, response)
}

func GetAllTodos(context *gin.Context) {
	var todos []models.Todo

	err := db.Find(&todos)
	if err.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Error": " Can't get todo data"})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
		"data":    todos,
	})
}
func GetTodoById(context *gin.Context) {
	var data models.TodoRequest

	reqId := context.Param("id")
	id := cast.ToUint(reqId)

	if err := context.BindJSON(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	todo := models.Todo{}
	todoById := db.Where("id = ?", id).First(&todo)

	if todoById.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Error": "Can't found todo"})
		return
	}

	var response models.TodoResponse
	response.ID = todo.ID
	response.Name = todo.Name
	response.Description = todo.Description

	context.JSON(http.StatusCreated, response)
}

func EditTodo(context *gin.Context) {
	var data models.TodoRequest

	reqId := context.Param("id")
	id := cast.ToUint(reqId)

	if err := context.BindJSON(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	todo := models.Todo{}
	todoById := db.Where("id = ?", id).First(&todo)

	if todoById.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Error": "Can't found todo"})
		return
	}
	todo.Name = data.Name
	todo.Description = data.Description

	result := db.Save(&todo)
	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Error": " Can't save todo"})
		return
	}

	var response models.TodoResponse
	response.ID = todo.ID
	response.Name = todo.Name
	response.Description = todo.Description

	context.JSON(http.StatusCreated, response)
}

func DeleteTodo(context *gin.Context) {
	var data models.TodoRequest

	reqId := context.Param("id")
	id := cast.ToUint(reqId)

	if err := context.BindJSON(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	todo := models.Todo{}
	result := db.Delete(&todo, id)
	print(result)

	context.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
		"data":    id,
	})
}
