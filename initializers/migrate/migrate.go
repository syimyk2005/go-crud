package main

import (
	"github.com/rubbyklein/go-crud/initializers"
	"github.com/rubbyklein/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
