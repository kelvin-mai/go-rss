package controller

import (
	"errors"
	"log"
	"os"
	"time"

	pasetoware "github.com/gofiber/contrib/paseto"
	"github.com/gofiber/fiber/v2"
	"github.com/o1egl/paseto"
	"kelvinmai.io/rss/internal/model"
	"kelvinmai.io/rss/internal/router/response"
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
		return err
	}
	password, err := util.HashPassword(input.Password)
	if err != nil {
		return err
	}
	user, err := c.s.Create(input.Username, password)
	if err != nil {
		return err
	}
	token, err := createToken(input.Username)
	if err != nil {
		return err
	}

	return response.Created(ctx, fiber.Map{
		"user":  user,
		"token": token,
	})
}

func (c AuthController) Login(ctx *fiber.Ctx) error {
	input := model.AuthInput{}
	if err := ctx.BodyParser(&input); err != nil {
		return err
	}
	user, err := c.s.GetByUsername(input.Username)
	if err != nil {
		return err
	}
	if !util.CheckPassword(input.Password, user.Password) {
		return errors.New("password: invalid password")
	}
	token, err := createToken(input.Username)
	if err != nil {
		log.Printf("%v\n", err)
	}
	return response.Ok(ctx, fiber.Map{
		"user":  user,
		"token": token,
	})
}

func (c AuthController) CurrentUser(ctx *fiber.Ctx) error {
	payload := ctx.Locals(pasetoware.DefaultContextKey).(model.AuthPayload)
	user, _ := c.s.GetByUsername(payload.Username)
	return response.Ok(ctx, fiber.Map{
		"user": user,
	})
}
