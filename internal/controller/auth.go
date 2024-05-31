package controller

import (
	"errors"
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

func createToken(user *model.User) (string, error) {
	key := os.Getenv("PASETO_KEY")
	duration := 12 * time.Hour
	payload := &model.AuthPayload{
		Username:  user.Username,
		UserId:    user.Id,
		IsAdmin:   user.IsAdmin,
		ExpiresAt: time.Now().Add(duration),
	}
	return paseto.NewV2().Encrypt([]byte(key), payload, nil)
}

func (c *AuthController) Register(ctx *fiber.Ctx) error {
	input := model.AuthInput{}
	if err := ctx.BodyParser(&input); err != nil {
		return response.ErrorBadRequest(err)
	}
	ve, ok := util.Validate(input)
	if !ok {
		return response.ErrorValidation(ve)
	}
	password, err := util.HashPassword(input.Password)
	if err != nil {
		return response.ErrorBadRequest(
			errors.New("invalid password"),
		)
	}
	user, err := c.s.Create(input.Username, password)
	if err != nil {
		return response.ErrorUnauthorized(
			errors.New("registration error"),
		)
	}
	token, err := createToken(user)
	if err != nil {
		return response.ErrorUnauthorized(
			errors.New("registration error"),
		)
	}

	return response.Created(ctx, fiber.Map{
		"user":  user,
		"token": token,
	})
}

func (c AuthController) Login(ctx *fiber.Ctx) error {
	input := model.AuthInput{}
	if err := ctx.BodyParser(&input); err != nil {
		return response.ErrorBadRequest(err)
	}
	ve, ok := util.Validate(input)
	if !ok {
		return response.ErrorValidation(ve)
	}
	user, err := c.s.GetByUsername(input.Username)
	if err != nil {
		return response.ErrorUnauthorized(
			errors.New("login error"),
		)
	}
	if !util.CheckPassword(input.Password, user.Password) {
		return response.ErrorUnauthorized(
			errors.New("login error"),
		)
	}
	token, err := createToken(&user)
	if err != nil {
		return response.ErrorUnauthorized(
			errors.New("login error"),
		)
	}
	return response.Ok(ctx, fiber.Map{
		"user":  user,
		"token": token,
	})
}

func (c AuthController) CurrentUser(ctx *fiber.Ctx) error {
	payload := ctx.Locals(pasetoware.DefaultContextKey).(model.AuthPayload)
	user, err := c.s.GetByUsername(payload.Username)
	if err != nil {
		return response.ErrorDatabase(err)
	}
	return response.Ok(ctx, fiber.Map{
		"user": user,
	})
}
