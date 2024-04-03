package main

import (
	"example/newprojectgo/Rest_Api/controllers"
	initializers "example/newprojectgo/Rest_Api/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadVariables()
	initializers.Connectdb()
}

func main() {
	r := gin.Default()
	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts", controllers.PostIndex)
	r.GET("/posts/:id", controllers.PostShow)

	r.Run()
}
