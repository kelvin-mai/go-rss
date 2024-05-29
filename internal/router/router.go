package router

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type Router struct {
	app *fiber.App
}

func Init() *Router {
	app := fiber.New(fiber.Config{
		ErrorHandler: DefaultErrorHandler,
	})
	return &Router{
		app: app,
	}
}

func (r *Router) Serve() {
	log.Fatal(r.app.Listen(":4000"))
}
