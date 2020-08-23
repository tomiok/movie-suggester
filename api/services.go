package api

type Services struct {
	search MovieSearch
}

func NewServices() Services {
	return Services{
		search: &MovieService{},
	}
}

type WebServices struct {
	s Services
}

func start() *WebServices {
	return &WebServices{s: NewServices()}
}
