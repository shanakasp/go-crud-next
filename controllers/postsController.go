package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/princesp/go-crud-next/initializers"
	"github.com/princesp/go-crud-next/models"
)

func PostsCreate(c *gin.Context) {
	// GEt Data off req body

	// Create a post
	post := models.Post{Title: "Shanaka", Body: "Hi HI"}
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