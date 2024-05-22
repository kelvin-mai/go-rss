package main

import (
	"log"

	"github.com/joho/godotenv"
	"kelvinmai.io/rss/internal/database"
	"kelvinmai.io/rss/internal/router"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	_, err := database.Connect()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	r := router.Init()
	r.Serve()
}
