package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rubbyklein/go-crud/controllers"
	"github.com/rubbyklein/go-crud/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {

	router := gin.Default()
	router.POST("/posts", controllers.PostsCreate)
	router.PUT("/posts/:id", controllers.PostsUpdate)
	router.GET("/posts", controllers.PostsIndex)
	router.GET("/posts/:id", controllers.PostsById)
	router.DELETE("/posts/:id", controllers.PostsDelete)

	router.Run() // listens on 0.0.0.0:8080 by default
}
