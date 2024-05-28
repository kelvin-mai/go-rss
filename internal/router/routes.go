package router

import (
	"github.com/gofiber/fiber/v2"
	"kelvinmai.io/rss/internal/controller"
)

type RouteControllers struct {
	UserController *controller.UserController
}

func healthCheck(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status": "ok",
	})
}

func (r *Router) SetupRoutes(c RouteControllers) {
	api := r.app.Group("/api")
	api.Get("/", healthCheck)

	user := api.Group("/user")
	user.Get("/", c.UserController.GetAllUsers)
}
