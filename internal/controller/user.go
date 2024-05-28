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

func (uc *UserController) GetAllUsers(c *fiber.Ctx) error {
	users, err := uc.s.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"users":   users,
	})
}
