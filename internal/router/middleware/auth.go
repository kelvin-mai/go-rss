package middleware

import (
	"encoding/json"
	"errors"
	"os"

	pasetoware "github.com/gofiber/contrib/paseto"
	"github.com/gofiber/fiber/v2"
	"kelvinmai.io/rss/internal/model"
	"kelvinmai.io/rss/internal/router/response"
)

func validateToken(decrypted []byte) (any, error) {
	var payload model.AuthPayload
	err := json.Unmarshal(decrypted, &payload)
	return payload, err
}

func errorHandler(ctx *fiber.Ctx, err error) error {
	if errors.Is(err, pasetoware.ErrDataUnmarshal) || errors.Is(err, pasetoware.ErrExpiredToken) {
		return response.ErrorUnauthorized(err)
	}
	return response.ErrorBadRequest(err)
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
