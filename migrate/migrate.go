package main

import (
	"github.com/princesp/go-crud-next/initializers"
	"github.com/princesp/go-crud-next/models"
)

func init(){
	initializers.LoadEnvVariables()
initializers.ConnectToDB()
}

func main() {
	
	 initializers.DB.AutoMigrate(&models.Post{})

}