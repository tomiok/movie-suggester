package internal

import (
	"github.com/gofiber/fiber"
)

func SetErrorHandler(app *fiber.App) {

	app.Settings.ErrorHandler = func(ctx *fiber.Ctx, err error) {
		// Status code defaults to 500
		code := fiber.StatusInternalServerError
		var msg string
		// Retrieve the custom status code if it's an fiber.*Error
		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
			msg = e.Message
		}

		if msg == "" {
			msg = "cannot process the http call"
		}

		// Send custom error page
		err = ctx.Status(code).JSON(internalError{
			Message: msg,
		})
	}
}

type internalError struct {
	Message string `json:"message"`
}
