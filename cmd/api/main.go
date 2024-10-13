package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
	// Create a new GIN router instance
	router := gin.Default()

	// Define a GET route for the homepage
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	// Run the Gin server on port 8080
	router.Run(":8080")
}