package api

import "github.com/tomiok/movies-suggester/internal/database"

type Services struct {
	search MovieSearch
	users  UserGateway
}

func NewServices() Services {
	client := database.NewMySQLClient()
	return Services{
		search: &MovieService{client},
		users:  &UserService{client},
	}
}

type WebServices struct {
	Services
}

func start() *WebServices {
	return &WebServices{NewServices()}
}
