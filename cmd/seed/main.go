package main

import (
	"log"

	"github.com/joho/godotenv"
	"kelvinmai.io/rss/internal/database"
	"kelvinmai.io/rss/internal/model"
	"kelvinmai.io/rss/util"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v\n", err)
	}
	db := database.Connect()
	password := util.HashPassword("password")
	users := []model.User{
		{
			Username: "user",
			Password: password,
		},
	}
	_, err := db.NamedExec(
		`insert into users (username, password)
		 values (:username, :password)
		`,
		users,
	)
	if err != nil {
		log.Fatalf("Error inserting users: %v\n", err)
	}
	log.Printf("Successfully inserted users: %v\n", users)
}
