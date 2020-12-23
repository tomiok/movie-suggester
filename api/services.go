package api

import (
	embedded "github.com/tomiok/fuego-cache/clients/inmemory"
	"github.com/tomiok/movies-suggester/internal/database"
)

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
	tokenKey string
	cache    *embedded.FuegoInMemory
}

func start(tokenKey string) *WebServices {
	inMemoryDB := embedded.NewInMemory(false, "")
	return &WebServices{NewServices(), tokenKey, inMemoryDB}
}
