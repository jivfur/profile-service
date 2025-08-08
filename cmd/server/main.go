package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jivfur/profile-service/internal/config"
	"github.com/jivfur/profile-service/internal/model"
	"github.com/jivfur/profile-service/internal/repository"
	"github.com/jivfur/profile-service/internal/service"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := db.AutoMigrate(&model.Profile{}, &model.Photo{}, &model.Location{}); err != nil {
		log.Fatalf("Database migration failed: %v", err)
	}
	log.Println("Database migrated successfully")

	repo := repository.NewGormProfileRepository(db)
	svc := service.NewProfileService(repo) // svc
	_ = svc                                // This is where you would typically set up your handler
	// h := handler.NewProfileHandler(svc)

	r := gin.Default()

	// Health check
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	// ✅ Register profile route
	// r.POST("/profiles", h.Create)

	r.Run() // default :8080
}
