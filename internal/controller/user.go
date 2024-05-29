package controller

import (
	"github.com/gofiber/fiber/v2"
	"kelvinmai.io/rss/internal/service"
)

type UserController struct {
	s *service.UserService
}

func NewUserController(s *service.UserService) *UserController {
	return &UserController{
		s: s,
	}
}

func (c *UserController) GetAllUsers(ctx *fiber.Ctx) error {
	users, err := c.s.GetAll()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"users": users,
		},
	})
}

func (c *UserController) GetUserById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	user, err := c.s.GetById(id)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Resource not found",
			"data":    nil,
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"user": user,
		},
	})
}
