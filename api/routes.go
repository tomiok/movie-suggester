package api

import (
	"github.com/gofiber/fiber"
	jwtware "github.com/gofiber/jwt"
)

func SetupMoviesRoutes(app *fiber.App) {
	s := start()
	grp := app.Group("/movies")
	grp.Get("/", s.SearchMovieHandler)
}

func SetupUsersRoutes(app *fiber.App) {
	s := start()
	grp := app.Group("/users")
	grp.Post("/", s.CreateUserHandler)

	grp.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("mysecretkey-changeme"),
	})).Get("/wishlist", s.WishListHandler)
}
