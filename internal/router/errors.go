package router

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"kelvinmai.io/rss/internal/model"
)

func DefaultErrorHandler(ctx *fiber.Ctx, err error) error {
	log.Errorf("error: %v\n", err)
	code := fiber.StatusInternalServerError
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}
	return ctx.Status(code).JSON(model.ApiResponse{
		Success: false,
		Data: fiber.Map{
			"error": err.Error(),
		},
	})
}
