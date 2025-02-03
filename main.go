package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rowjay007/number-classifier-api/routes"
	"log"
	"os"
)

func main() {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Next()
	})

	routes.RegisterRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" 
	}
	log.Println("Server is running on port " + port)
	r.Run(":" + port)
}
