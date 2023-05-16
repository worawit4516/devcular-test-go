package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/worawit4516/devcular-test-go/src/controllers"
)

func Routes() {
	route := gin.Default()
	route.GET("/health", controllers.CheckHealth)
	route.GET("/todo", controllers.GetAllTodos)
	route.GET("/todo/:id", controllers.GetTodoById)
	route.POST("/todo", controllers.CreateTodo)
	route.PUT("/todo/:id", controllers.EditTodo)
	route.DELETE("/todo/:id", controllers.DeleteTodo)
	route.Run()
}
