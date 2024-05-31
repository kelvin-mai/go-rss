package response

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"kelvinmai.io/rss/internal/model"
)

func newMeta(ctx *fiber.Ctx) model.ApiMeta {
	return model.ApiMeta{
		Timestamp: time.Now(),
		Path:      ctx.Path(),
		Method:    ctx.Method(),
	}
}

func Response(
	ctx *fiber.Ctx,
	code int,
	data interface{},
) error {
	return ctx.Status(code).JSON(model.ApiResponse{
		Success: true,
		Data:    data,
		Meta:    newMeta(ctx),
	})
}

func Ok(ctx *fiber.Ctx, data interface{}) error {
	return Response(ctx, fiber.StatusOK, data)
}

func Created(ctx *fiber.Ctx, data interface{}) error {
	return Response(ctx, fiber.StatusCreated, data)
}
