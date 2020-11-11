package api

import (
	"github.com/gofiber/fiber/v2"
)

func (w *WebServices) CreateUserHandler(c *fiber.Ctx) error {
	var cmd CreateUserCMD
	err := c.BodyParser(&cmd)

	res, err := w.Services.users.SaveUser(cmd)

	if err != nil {
		return fiber.NewError(400, "cannot create user")
	}

	res.JWT = signToken(w.tokenKey, res.ID)
	return c.JSON(res)
}

func (w *WebServices) WishListHandler(c *fiber.Ctx) error {
	var cmd WishMovieCMD
	_ = c.BodyParser(&cmd)
	bearer := c.Get("Authorization")
	userID := extractUserIDFromJWT(bearer, w.tokenKey)
	err := w.users.AddWishMovie(userID, cmd.MovieID, cmd.Comment)

	if err != nil {
		return fiber.NewError(400, "cannot add to the wishlist")
	}

	return c.JSON(struct {
		R string `json:"result"`
	}{
		R: "movie added to the wishlist",
	})
}

func (w *WebServices) ServeVideo(c *fiber.Ctx) error {
	c.Set("Content-Type", "video/mp4")
	err := c.SendFile("test.MP4", false)

	if err != nil {
		return fiber.NewError(400, "cannot display video")

	}
	return nil
}

func (w *WebServices) LoginHandler(c *fiber.Ctx) error {
	var cmd LoginCMD
	err := c.BodyParser(&cmd)
	if err != nil {
		return fiber.NewError(400, "cannot parse params")
	}

	id := w.users.Login(cmd)

	if id == "" {
		return fiber.NewError(404, "user not found")
	}

	return c.JSON(struct {
		Token string `json:"token"`
	}{
		Token: signToken(w.tokenKey, id),
	})
}

type LoginCMD struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type WishMovieCMD struct {
	MovieID string `json:"movie_id"`
	Comment string `json:"comment"`
}
