package api

import "github.com/gofiber/fiber/v2"

func (w *WebServices) AddUpcomingHandler(c *fiber.Ctx) error {
	var movie UpcomingCMD
	c.BodyParser(&movie)

	return nil
}

type UpcomingCMD struct {
	Movie string `json:"movie"`
}