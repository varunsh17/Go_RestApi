package main

import (
	"example/golang_practice/Go_RestApi/controllers"
	initializers "example/golang_practice/Go_RestApi/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadVariables()
	initializers.Connectdb()
}

func main() {
	r := gin.Default()
	r.POST("/posts", controllers.PostsCreate)
	r.Run()
}
