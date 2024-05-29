package middleware

import (
	"encoding/json"
	"log"
	"os"

	pasetoware "github.com/gofiber/contrib/paseto"
	"github.com/gofiber/fiber/v2"
	"kelvinmai.io/rss/internal/model"
)

func validateToken(decrypted []byte) (any, error) {
	var payload model.AuthPayload
	err := json.Unmarshal(decrypted, &payload)
	return payload, err
}

func Authenticate() fiber.Handler {
	key := os.Getenv("PASETO_KEY")
	log.Printf("PASETO_KEY: %v\n", key)

	return pasetoware.New(pasetoware.Config{
		TokenPrefix:  "Bearer",
		SymmetricKey: []byte(key),
		Validate:     validateToken,
	})
}
