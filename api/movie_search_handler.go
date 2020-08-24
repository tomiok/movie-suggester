package api

import (
	"github.com/gofiber/fiber"
	"github.com/tomiok/movies-suggester/internal/logs"
)

func (w *WebServices) SearchMovieHandler(c *fiber.Ctx) {
	d := c.Query("director")
	logs.Info(d)
	res, err := w.search.Search(MovieFilter{})

	if err != nil {
		err = fiber.NewError(400, "cannot bring movies")
		c.Next(err)
		return
	}

	c.JSON(res)
}
