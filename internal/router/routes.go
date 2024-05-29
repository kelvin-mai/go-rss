package router

import (
	"github.com/gofiber/fiber/v2"
	"kelvinmai.io/rss/internal/controller"
	"kelvinmai.io/rss/internal/router/middleware"
)

type RouteControllers struct {
	AuthController *controller.AuthController
	UserController *controller.UserController
}

func healthCheck(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"status": "ok",
	})
}

func (r *Router) SetupRoutes(c RouteControllers) {
	api := r.app.Group("/api")
	api.Get("/", healthCheck)

	api.Post("/login", c.AuthController.Login)
	api.Post("/register", c.AuthController.Register)

	users := api.Group("/users")
	users.Get("/", c.UserController.GetAllUsers)
	users.Get("/me", middleware.Authenticate(), c.AuthController.CurrentUser)
	users.Get("/:id", c.UserController.GetUserById)
}
