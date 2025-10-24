package main

import (
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/data/aztables"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	var _ = aztables.ServiceClient{}

	_ = godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	r.GET("/azure/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"error": "route not found"})
	})

	log.Printf("Server running on port %s", port)
	r.Run(":" + port)
}
