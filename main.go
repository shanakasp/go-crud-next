package main

import (
	"github.com/gin-gonic/gin"
	"github.com/princesp/go-crud-next/controllers"
	"github.com/princesp/go-crud-next/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
 }
func main() {

	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong 13",
	// 	})
	// })
	// r.Run() 
	
	
	r := gin.Default()
	r.POST("/ping", controllers.PostsCreate)
	r.Run() 
	
}