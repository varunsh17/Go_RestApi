package main

import (
	"example/newprojectgo/Rest_Api/initializers"
	"example/newprojectgo/Rest_Api/models"
)

func init() {
	initializers.LoadVariables()
	initializers.Connectdb()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
