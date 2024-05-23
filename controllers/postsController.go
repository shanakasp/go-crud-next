package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/princesp/go-crud-next/initializers"
	"github.com/princesp/go-crud-next/models"
)

func PostsCreate(c *gin.Context) {
	// GEt Data off req body
	var body struct{
		Body string
		Title string
	}

	c.Bind(&body)

	// Create a post
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
	    c.Status(400)
		return
	}
	// Return it

	c.JSON(200, gin.H{
		"post": post,
	})
}

func GetAllPosts(c *gin.Context) {
	// Get all records
	var posts []models.Post
	initializers.DB.Find(&posts)


	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func GetPostBYId(c *gin.Context) {

	//get id
	id := c.Param("id")
	// Get all records
	var post models.Post
	initializers.DB.First(&post,id )

	

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {

	//Get the ID of the record, url
	id := c.Param("id")
	// GEt Data off req body
	var body struct{
		Body string
		Title string
	}

	c.Bind(&body)
	//Find the post
	var post models.Post
	initializers.DB.First(&post,id )

	// Update the post
	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})
	// Return it

	c.JSON(200, gin.H{
		"post": post,
	})
}

func DeletePost(c *gin.Context) {

	//get id
	id := c.Param("id")
	// Delete the Post
	initializers.DB.Delete(&models.Post{}, id)
	
	//responce
	c.Status(200)
}