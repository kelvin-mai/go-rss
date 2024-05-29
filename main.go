package main

import (
	"log"

	"github.com/joho/godotenv"
	"kelvinmai.io/rss/internal/controller"
	"kelvinmai.io/rss/internal/database"
	"kelvinmai.io/rss/internal/router"
	"kelvinmai.io/rss/internal/service"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	db := database.Connect()

	r := router.Init()

	us := service.NewUserService(db)

	ac := controller.NewAuthController(us)
	uc := controller.NewUserController(us)

	r.SetupRoutes(router.RouteControllers{
		AuthController: ac,
		UserController: uc,
	})
	r.Serve()
}
