package response

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"kelvinmai.io/rss/internal/model"
)

func ErrorBadRequest(err error) error {
	return model.NewApiError(
		fiber.StatusBadRequest,
		"Bad request",
		err.Error(),
		nil,
	)
}

func ErrorNotFound(err error) error {
	return model.NewApiError(
		fiber.StatusNotFound,
		"Resource not found",
		err.Error(),
		nil,
	)
}

func ErrorValidation(ve []model.ValidationError) error {
	return model.NewApiError(
		fiber.StatusBadRequest,
		"Validation error",
		"Does not satisfy input requirements",
		ve,
	)
}

func ErrorUnauthorized(err error) error {
	return model.NewApiError(
		fiber.StatusUnauthorized,
		"Unauthorized error",
		err.Error(),
		nil,
	)
}

func ErrorRequireAdmin() error {
	return model.NewApiError(
		fiber.StatusUnauthorized,
		"Unauthorized error",
		"Require admin access",
		nil,
	)
}

func ErrorDatabase(err error) error {
	return model.NewApiError(
		fiber.StatusInternalServerError,
		"Database error",
		err.Error(),
		nil,
	)
}

func DefaultErrorHandler(ctx *fiber.Ctx, err error) error {
	log.Errorf("error: %v\n", err)
	meta := newMeta(ctx)
	e, ok := err.(*model.ApiError)
	if !ok {
		ef, ok := err.(*fiber.Error)
		if !ok {
			e = model.NewApiError(
				fiber.StatusInternalServerError,
				"Internal Server Error",
				"Something went wrong",
				nil,
			)
		} else {
			e = model.NewApiError(
				ef.Code,
				ef.Message,
				ef.Error(),
				nil,
			)
		}
	}
	return ctx.Status(e.Code).JSON(model.ApiResponse{
		Success: false,
		Error:   e,
		Meta:    meta,
	})
}
