package api

import (
	"github.com/gofiber/fiber"
	jwtware "github.com/gofiber/jwt"
)

func jwtMiddleware(secret string) func(*fiber.Ctx) {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(secret),
	})
}
