package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jivfur/profile-service/internal/config"
	"github.com/jivfur/profile-service/internal/model"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	// Connect to DB
	db := config.ConnectDB()

	// Run AutoMigrate to create/update tables
	if err := db.AutoMigrate(&model.Profile{}, &model.Photo{}, &model.Location{}); err != nil {
		log.Fatalf("Database migration failed: %v", err)
	}

	log.Println("Database migrated successfully")

	// Set up Gin router
	r := gin.Default()

	// Simple health check endpoint
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
