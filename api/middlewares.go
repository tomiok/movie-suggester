package api

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
	jwtware "github.com/gofiber/jwt"
	"github.com/tomiok/movies-suggester/internal/logs"
	"time"
)

func jwtMiddleware(secret string) func(*fiber.Ctx) {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(secret),
	})
}

func signToken(tokenKey, id string) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	claims["sub"] = id

	t, err := token.SignedString([]byte(tokenKey))

	if err != nil {
		return ""
	}

	return t
}

func extractUserIDFromJWT(bearer, tokenKey string) string {
	//Bearer ey.....
	token := bearer[7:]
	logs.Info(token)
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(tokenKey), nil
	})

	if err != nil {
		return ""
	}

	if t.Valid {
		claims := t.Claims.(jwt.MapClaims)
		return claims["sub"].(string)
	}

	return ""
}
