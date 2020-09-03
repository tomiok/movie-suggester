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
	Login()
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

func (us *UserService) Login() {

}
