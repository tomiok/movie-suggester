package main

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	"github.com/tomiok/movies-suggester/api"
	"github.com/tomiok/movies-suggester/internal"
	"github.com/tomiok/movies-suggester/internal/logs"
)

func main() {
	app := fiber.New()
	key := "tokenKey"
	internal.SetErrorHandler(app)
	app.Use(middleware.Recover())
	api.SetupMoviesRoutes(app, key)
	api.SetupUsersRoutes(app, key)
	logs.Info("starting...")
	_ = app.Listen("3001")
}
