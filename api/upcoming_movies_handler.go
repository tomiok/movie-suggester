package api

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/tomiok/fuego-cache/logs"
)

func (w *WebServices) AddUpcomingHandler(c *fiber.Ctx) error {
	var upcoming UpcomingCMD
	err := c.BodyParser(&upcoming)

	if err != nil {
		logs.LogError(err)
	}

	value := fmt.Sprintf("Title: %s, Date: %s", upcoming.Movie, upcoming.ReleaseDate)
	return w.cache.Insert(upcoming.Movie, value)
}

func (w *WebServices) GetUpcomingHandler(c *fiber.Ctx) error {
	movies := w.cache.List()
	return c.JSON(&movies)
}

func (w *WebServices) DeleteUpcomingHandler(c *fiber.Ctx) error {
	movieKey := c.Query("key")

	logs.Info(movieKey)
	s := w.cache.Delete(movieKey)

	if s != "ok" {
		return errors.New("cannot delete the movie " + movieKey)
	}

	return nil
}

type UpcomingSummary struct {
	Movies []string `json:"movies"`
}

type UpcomingCMD struct {
	Movie       string `json:"title"`
	ReleaseDate string `json:"release_date"`
}
