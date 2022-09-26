package routes

import (
	"example/data-access/controllers"
	"example/data-access/middlewares"

	"github.com/gin-gonic/gin"
)

func RouterInit() {

	R := gin.Default()
	R.Use(middlewares.CorsMiddleware())
	R.POST("/todos", controllers.CreateTodos)
	R.GET("/todos", controllers.GetTodos)
	R.PATCH("/todos/:id", controllers.UpdateTodo)
	R.DELETE("/todos/:id", controllers.DeleteTodo)
	R.Run(":8080")

}
