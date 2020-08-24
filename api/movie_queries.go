package api

func getMoviesQuery() string {
	return "select id, title, cast, genre, release_date, director from movie"
}
