package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"github.com/worawit4516/devcular-test-go/src/config"
	"github.com/worawit4516/devcular-test-go/src/models"
	"gorm.io/gorm"
)

// Get DB connections
var db *gorm.DB = config.ConnectDatabase()

// functions for check status server
func CheckHealth(context *gin.Context) {
	context.JSON(http.StatusOK, "Backend is Good!!")
}

// functions for create new todo
func CreateTodo(context *gin.Context) {
	var data models.TodoRequest

	//Get raw body then pass from JSON to Struct
	if err := context.ShouldBindJSON(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	todo := models.Todo{}
	todo.Name = data.Name
	todo.Description = data.Description

	//Create
	result := db.Create(&todo)

	//Check response of db
	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Error": " Can't save new todo"})
		return
	}

	//Create response
	var response models.TodoResponse
	response.ID = todo.ID
	response.Name = todo.Name
	response.Description = todo.Description

	context.JSON(http.StatusCreated, response)
}

// functions for get all todo
func GetAllTodos(context *gin.Context) {
	var todos []models.Todo

	//Search todo
	err := db.Find(&todos)
	if err.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Error": " Can't get todo data"})
		return
	}

	//Create response
	context.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
		"data":    todos,
	})
}

// functions for get todo by use Id
func GetTodoById(context *gin.Context) {
	var data models.TodoRequest

	//Get paramiter and cast it
	reqId := context.Param("id")
	id := cast.ToUint(reqId)

	if err := context.BindJSON(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	//Search
	todo := models.Todo{}
	todoById := db.Where("id = ?", id).First(&todo)

	if todoById.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Error": "Can't found todo"})
		return
	}

	//Create response
	var response models.TodoResponse
	response.ID = todo.ID
	response.Name = todo.Name
	response.Description = todo.Description

	context.JSON(http.StatusCreated, response)
}

// functions for update todo
func EditTodo(context *gin.Context) {
	var data models.TodoRequest

	//Get paramiter and cast it
	reqId := context.Param("id")
	id := cast.ToUint(reqId)

	if err := context.BindJSON(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	//Search
	todo := models.Todo{}
	todoById := db.Where("id = ?", id).First(&todo)

	if todoById.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Error": "Can't found todo"})
		return
	}

	//update new data
	todo.Name = data.Name
	todo.Description = data.Description

	//Save to DB
	result := db.Save(&todo)
	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Error": " Can't save todo"})
		return
	}

	//Create response
	var response models.TodoResponse
	response.ID = todo.ID
	response.Name = todo.Name
	response.Description = todo.Description

	context.JSON(http.StatusCreated, response)
}

// functions for delete todo
func DeleteTodo(context *gin.Context) {
	var data models.TodoRequest

	//Get paramiter and cast it
	reqId := context.Param("id")
	id := cast.ToUint(reqId)

	if err := context.BindJSON(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	//Search then delete
	todo := models.Todo{}
	result := db.Delete(&todo, id)
	print(result)

	//Create response
	context.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
		"data":    id,
	})
}
