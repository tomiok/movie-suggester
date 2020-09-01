package api

import (
	"github.com/gofiber/fiber"
)

func (w *WebServices) SearchMovieHandler(c *fiber.Ctx) {

	res, err := w.search.Search(MovieFilter{
		Title:    c.Query("title"),
		Genre:    c.Query("genre"),
		Director: c.Query("director"),
	})

	if err != nil {
		err = fiber.NewError(400, "cannot bring movies")
		c.Next(err)
		return
	}

	if len(res) == 0 {
		_ = c.JSON([]interface{}{})
	}

	_ = c.JSON(res)
}
