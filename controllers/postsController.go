package controllers

import (
	"example/newprojectgo/Rest_Api/initializers"
	"example/newprojectgo/Rest_Api/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {

	var body struct {
		Title string
		Body  string
	}
	c.Bind(&body)
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post) // pass pointer of data to Create
	if result.Error != nil {

		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostIndex(c *gin.Context) {
	var posts []models.Post

	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"posts": posts,
	})

}

func PostShow(c *gin.Context) {
	//get id from url
	id := c.Param("id")

	var post models.Post

	initializers.DB.First(&post, id)

	c.JSON(200, gin.H{
		"post": post,
	})

}

func PostUpdate(c *gin.Context) {
	//get ID from url
	id := c.Param("id")

	//get data
	var body struct {
		Title string
		Body  string
	}
	c.Bind(&body)

	//Fetch post of that id
	var post models.Post

	initializers.DB.First(&post, id)

	//update
	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})

	c.JSON(200, gin.H{
		"post": post,
	})

}

func PostDelete(c *gin.Context) {
	//get ID from url
	id := c.Param("id")

	//get data
	var body struct {
		Title string
		Body  string
	}
	c.Bind(&body)

	//Fetch post of that id
	var post models.Post

	initializers.DB.Delete(&post, id)

	c.JSON(200, gin.H{
		"post": post,
	})

}
