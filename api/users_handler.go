package api

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
	"time"
)

func (w *WebServices) CreateUserHandler(c *fiber.Ctx) {
	var cmd CreateUserCMD
	err := c.BodyParser(&cmd)

	res, err := w.Services.users.SaveUser(cmd)

	if err != nil {
		err = fiber.NewError(400, "cannot create user")
		c.Next(err)
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["test-name"] = "tomasito"
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(w.tokenKey))

	res.JWT = t
	_ = c.JSON(res)
}

func (w *WebServices) WishListHandler(c *fiber.Ctx) {

	_ = c.JSON(struct {
		WishList string `json:"wish_list"`
	}{
		WishList: "some movies here",
	})
}
