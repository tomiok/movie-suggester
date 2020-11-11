package api

import (
	"github.com/gofiber/fiber/v2"
)

func (w *WebServices) SearchMovieHandler(c *fiber.Ctx) error {

	res, err := w.search.Search(MovieFilter{
		Title:    c.Query("title"),
		Genre:    c.Query("genre"),
		Director: c.Query("director"),
	})

	if err != nil {
		return fiber.NewError(400, "cannot bring movies")
	}

	if len(res) == 0 {
		return c.JSON([]interface{}{})
	}

	return c.JSON(res)
}
