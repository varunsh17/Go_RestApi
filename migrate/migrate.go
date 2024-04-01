package main

import (
	"example/golang_practice/Go_RestApi/initializers"
	"example/golang_practice/Go_RestApi/models"
)

func init() {
	initializers.LoadVariables()
	initializers.Connectdb()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
