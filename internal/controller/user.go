package controller

import (
	"github.com/gofiber/fiber/v2"
	"kelvinmai.io/rss/internal/router/response"
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
		return err
	}
	return response.Ok(ctx, fiber.Map{
		"users": users,
	})
}

func (c *UserController) GetUserById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	user, err := c.s.GetById(id)
	if err != nil {
		return err
	}
	return response.Ok(ctx, fiber.Map{
		"user": user,
	})
}
