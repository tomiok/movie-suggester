package main

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	"github.com/tomiok/movies-suggester/api"
	"github.com/tomiok/movies-suggester/internal"
)

func main() {
	app := fiber.New()
	key := "tokenKey"
	internal.SetErrorHandler(app)
	app.Use(middleware.Recover())
	api.SetupMoviesRoutes(app, key)
	api.SetupUsersRoutes(app, key)

	_ = app.Listen("3001")
}
