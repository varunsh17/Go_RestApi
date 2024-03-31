package main

import (
	initializers "example/newprojectgo/Rest_Api/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadVariables()
	initializers.Connectdb()
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
