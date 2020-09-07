package api

import (
	"github.com/gofiber/utils"
	"github.com/tomiok/movies-suggester/internal/database"
	"github.com/tomiok/movies-suggester/internal/logs"
)

type CreateUserCMD struct {
	Username       string `json:"username"`
	Password       string `json:"password"`
	RepeatPassword string `json:"repeat_password"`
}

type UserSummary struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	JWT      string `json:"token"`
}

type UserGateway interface {
	SaveUser(cmd CreateUserCMD) (*UserSummary, error)
	Login(cmd LoginCMD) string
	AddWishMovie(userID, movieID, comment string) error
}

type UserService struct {
	*database.MySqlClient
}

func (us *UserService) SaveUser(cmd CreateUserCMD) (*UserSummary, error) {
	id := utils.UUID()
	_, err := us.Exec(CreateUserQuery(), id, cmd.Username, cmd.Password)

	if err != nil {
		logs.Error("cannot insert user " + err.Error())
		return nil, err
	}

	return &UserSummary{
		ID:       id,
		Username: cmd.Username,
		JWT:      "",
	}, nil
}

// returns a string for the JWT
func (us *UserService) Login(cmd LoginCMD) string {
	var id string
	err := us.QueryRow(GetLoginQuery(), cmd.Username, cmd.Password).Scan(&id)

	if err != nil {
		logs.Error(err.Error())
		return ""
	}
	return id
}

func (us *UserService) AddWishMovie(userID, movieID, comment string) error {
	_, err := us.Exec(GetAddWishMovieQuery(), userID, movieID, comment)

	if err != nil {
		return err
	}

	return nil
}
