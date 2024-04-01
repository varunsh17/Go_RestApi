package controllers

import (
	"example/golang_practice/Go_RestApi/initializers"
	"example/golang_practice/Go_RestApi/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {

	post := models.Post{Title: "Jinzhu", Body: "f ewrf erwrf ere e "}

	result := initializers.DB.Create(&post) // pass pointer of data to Create
	if result.Error != nil {

		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"post": post,
	})
}
