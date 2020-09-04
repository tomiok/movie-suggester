package api

import (
	"github.com/tomiok/movies-suggester/internal/database"
	"github.com/tomiok/movies-suggester/internal/logs"
)

type MovieFilter struct {
	Title    string `json:"title,omitempty"`
	Genre    string `json:"genre,omitempty"`
	Director string `json:"director,omitempty"`
}

type Movie struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Cast        string `json:"cast"`
	ReleaseDate string `json:"release_date"`
	Genre       string `json:"genre"`
	Director    string `json:"director,omitempty"`
}

type MovieSearch interface {
	Search(filter MovieFilter) ([]Movie, error)
}

type MovieService struct {
	*database.MySqlClient
}

func (s *MovieService) Search(filter MovieFilter) ([]Movie, error) {
	tx, err := s.Begin()

	if err != nil {
		logs.Error("cannot create transaction " + err.Error())
		return nil, err
	}

	rows, err := tx.Query(getMoviesQuery(filter))

	if err != nil {
		logs.Error("cannot read movies " + err.Error())
		_ = tx.Rollback()
		return nil, err
	}

	var _movies []Movie
	for rows.Next() {
		var movie Movie
		err := rows.Scan(&movie.Id, &movie.Title, &movie.Cast, &movie.Genre, &movie.ReleaseDate, &movie.Director)
		if err != nil {
			logs.Error("cannot read movies " + err.Error())
		}
		_movies = append(_movies, movie)
	}

	_ = tx.Commit()
	return _movies, nil
}
