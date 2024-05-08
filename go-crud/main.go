package main

import (
	"go-crud/controllers"
	"go-crud/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	r := gin.Default()

	r.GET("/ping", controllers.HandlePingRequest)

	r.POST("/posts", controllers.CreatePost)
	// route for getting post by id
	r.GET("/posts/:id", controllers.GetPost)

	r.Run()
}
