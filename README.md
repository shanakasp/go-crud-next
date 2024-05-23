go mod init

go get github.com/githubnemo/CompileDaemon
https://github.com/githubnemo/CompileDaemon
go install github.com/githubnemo/CompileDaemon

go get github.com/joho/godotenv
https://github.com/joho/godotenv

go get -u github.com/gin-gonic/gin
https://gin-gonic.com/docs/quickstart/

go get -u gorm.io/gorm
https://gorm.io/docs/

go get -u gorm.io/driver/postgres

Created main.go file

CompileDaemon -command="./go-crud-next"

go run migrate/migrate.go


With and without GIN,

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router
	router := gin.Default()

	// Define a route for the root path
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	// Define a route with a path parameter
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, " + name,
		})
	})

	// Start the server on port 8080
	router.Run(":8080")
}


and, 

package main

import (
	"encoding/json"
	"net/http"
	"strings"
	"fmt"
)

func main() {
	// Define a route for the root path
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			response := map[string]string{"message": "Hello, World!"}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	// Define a route with a path parameter
	http.HandleFunc("/user/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			name := strings.TrimPrefix(r.URL.Path, "/user/")
			response := map[string]string{"message": "Hello, " + name}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	// Start the server on port 8080
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
