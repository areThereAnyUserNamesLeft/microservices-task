package main

import (
	"microservice-task/app"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	// Apply the middleware to the router (works with groups too
	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))
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
	router.Run(":8090")
}
