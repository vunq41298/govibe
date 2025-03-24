package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	// Load environment variables
	pgURL := os.Getenv("PG_URL")
	if pgURL == "" {
		log.Fatal("PG_URL is not set")
	}

	// Initialize PostgreSQL connection
	db, err := sql.Open("postgres", pgURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Database is not reachable: %v", err)
	}
	fmt.Println("Connected to PostgreSQL successfully!")

	// Initialize Gin server
	r := gin.Default()

	// Health check endpoint
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	r.Run(":" + port)
}
