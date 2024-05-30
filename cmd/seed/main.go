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
	password, err := util.HashPassword("password")
	if err != nil {
		log.Fatalf("Error generating password: %v\n", err)
	}
	users := []model.User{
		{
			Username: "admin",
			Password: password,
			IsAdmin:  true,
		},
	}
	_, err = db.NamedExec(
		`insert into users (username, password, is_admin)
		 values (:username, :password, :is_admin)
		`,
		users,
	)
	if err != nil {
		log.Fatalf("Error inserting users: %v\n", err)
	}
	log.Printf("Successfully inserted users: %v\n", users)
}
