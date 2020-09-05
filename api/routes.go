package api

import (
	"github.com/gofiber/fiber"
)

func SetupMoviesRoutes(app *fiber.App, tokenKey string) {
	s := start(tokenKey)
	grp := app.Group("/movies")
	grp.Get("/", s.SearchMovieHandler)
}

func SetupUsersRoutes(app *fiber.App, tokenKey string) {
	s := start(tokenKey)
	grp := app.Group("/users")
	grp.Post("/", s.CreateUserHandler)

	grp.Use(jwtMiddleware(tokenKey)).Get("/wishlist", s.WishListHandler)
}
