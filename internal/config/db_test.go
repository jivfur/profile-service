package config

import (
	"log"
	"testing"

	"github.com/joho/godotenv"
)

func TestConnectDB(t *testing.T) {
	err := godotenv.Load("../../.env") // Adjust path if needed
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db := ConnectDB()
	sqlDB, err := db.DB()
	if err != nil {
		t.Fatalf("Failed to get generic database object: %v", err)
	}
	defer sqlDB.Close()

	err = sqlDB.Ping()
	if err != nil {
		t.Fatalf("Failed to ping database: %v", err)
	}
}
