package controller

import (
	"fmt"
	"log"
	"os"
	"time"

	pasetoware "github.com/gofiber/contrib/paseto"
	"github.com/gofiber/fiber/v2"
	"github.com/o1egl/paseto"
	"kelvinmai.io/rss/internal/model"
	"kelvinmai.io/rss/internal/service"
	"kelvinmai.io/rss/util"
)

type AuthController struct {
	s *service.UserService
}

func NewAuthController(s *service.UserService) *AuthController {
	return &AuthController{
		s: s,
	}
}

func createToken(username string) (string, error) {
	key := os.Getenv("PASETO_KEY")
	duration := 12 * time.Hour
	payload := &model.AuthPayload{
		Username:  username,
		ExpiresAt: time.Now().Add(duration),
	}
	return paseto.NewV2().Encrypt([]byte(key), payload, nil)
}

func (c *AuthController) Register(ctx *fiber.Ctx) error {
	input := model.AuthInput{}
	if err := ctx.BodyParser(&input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Error on register",
			"data":    err,
		})
	}
	password := util.HashPassword(input.Password)
	// user := &model.User{}
	// c.s.GetByUsername(input.Username)
	err := c.s.Create(input.Username, password)
	if err != nil {
		log.Printf("%v\n", err)
	}
	token, err := createToken(input.Username)
	if err != nil {
		log.Printf("%v\n", err)
	}
	user, _ := c.s.GetByUsername(input.Username)

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"user":  user,
			"token": token,
		},
	})
}

func (c AuthController) Login(ctx *fiber.Ctx) error {
	input := model.AuthInput{}
	if err := ctx.BodyParser(&input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Error on login",
			"data":    err,
		})
	}
	user, _ := c.s.GetByUsername(input.Username)
	if !util.CheckPassword(input.Password, user.Password) {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "Error or login",
		})
	}
	token, err := createToken(input.Username)
	if err != nil {
		log.Printf("%v\n", err)
	}
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"user":  user,
			"token": token,
		},
	})
}

func (c AuthController) CurrentUser(ctx *fiber.Ctx) error {
	payload := ctx.Locals(pasetoware.DefaultContextKey).(model.AuthPayload)
	fmt.Print(payload)
	user, _ := c.s.GetByUsername(payload.Username)
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"user":    user,
			"payload": payload,
		},
	})
}
