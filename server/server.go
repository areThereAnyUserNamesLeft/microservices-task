package main

import (
	"./app"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("api/v1")
	{
		v1.GET("/users", app.Getusers)
		v1.GET("/users/:id", app.GetInstruction)
		v1.POST("/users", app.PostInstruction)
		v1.PUT("/users/:id", app.UpdateInstruction)
		v1.DELETE("/users/:id", app.DeleteInstruction)
	}

	return router
}

func main() {
	router := SetupRouter()
	router.Run(":8080")
}
