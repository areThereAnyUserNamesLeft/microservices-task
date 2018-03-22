package main

import (
	"microservice-task/app"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("api/v1")
	{
		v1.GET("/users", app.GetUsers)
		v1.GET("/users/:id", app.GetUser)
		v1.POST("/users", app.PostUser)
		v1.PUT("/users/:id", app.UpdateUser)
		v1.DELETE("/users/:id", app.DeleteUser)
	}

	return router
}

func main() {
	router := SetupRouter()
	router.Run(":8080")
}
