package main

import (
	"github.com/gofiber/fiber"
	"github.com/tomiok/movies-suggester/api"
	"github.com/tomiok/movies-suggester/internal"
)

func main() {
	app := fiber.New()
	internal.SetErrorHandler(app)
	api.SetupMoviesRoutes(app)
	_ = app.Listen("3001")
}
