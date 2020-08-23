package main

import (
	"github.com/gofiber/fiber"
	"github.com/tomiok/movies-suggester/api"
)

func main() {
	app := fiber.New()
	api.SetupMoviesRoutes(app)
	_ = app.Listen("3001")
}
