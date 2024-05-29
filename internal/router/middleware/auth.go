package middleware

import (
	"encoding/json"
	"errors"
	"os"

	pasetoware "github.com/gofiber/contrib/paseto"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"kelvinmai.io/rss/internal/model"
)

func validateToken(decrypted []byte) (any, error) {
	var payload model.AuthPayload
	err := json.Unmarshal(decrypted, &payload)
	return payload, err
}

func errorHandler(ctx *fiber.Ctx, err error) error {
	log.Errorf("error: %v\n", err)
	code := fiber.StatusBadRequest
	if errors.Is(err, pasetoware.ErrDataUnmarshal) || errors.Is(err, pasetoware.ErrExpiredToken) {
		code = fiber.StatusUnauthorized
	}
	return ctx.Status(code).JSON(model.ApiResponse{
		Success: false,
		Data: fiber.Map{
			"error": err.Error(),
		},
	})
}

func Authenticate() fiber.Handler {
	key := os.Getenv("PASETO_KEY")
	return pasetoware.New(pasetoware.Config{
		TokenPrefix:  "Bearer",
		SymmetricKey: []byte(key),
		Validate:     validateToken,
		ErrorHandler: errorHandler,
	})
}
