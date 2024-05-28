package main

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v\n", err)
	}
	dsn := url.URL{
		User: url.UserPassword(os.Getenv("DATABASE_USER"),
			os.Getenv("DATABASE_PASSWORD")),
		Scheme: "postgres",
		Host: fmt.Sprintf("%s:%d", os.Getenv("DATABASE_HOST"),
			5432,
		),
		Path:     os.Getenv("DATABASE_NAME"),
		RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
	}
	m, err := migrate.New(
		"file://migrations",
		dsn.String(),
	)
	if err != nil {
		log.Fatalf("Error connecting to database migrations: %v\n", err)
	}
	if err := m.Up(); err != nil {
		log.Fatalf("Error running database migrations: %v\n", err)
	} else {
		log.Println("migration successful")
	}
}
