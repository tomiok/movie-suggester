package api

import "strings"

func getMoviesQuery(filter MovieFilter) string {
	var (
		d, g, t string
		clause  = false
		q       = "select id, title, cast, genre, release_date, director from movie"
		b       = strings.Builder{}
	)

	b.WriteString(q)

	if filter.Director != "" {
		d = "director like '%" + filter.Director + "%'"
		clause = true
	}

	if filter.Genre != "" {
		g = "genre like '%" + filter.Genre + "%'"
		clause = true
	}

	if filter.Title != "" {
		t = "title like '%" + filter.Title + "%'"
		clause = true
	}

	if clause {
		var i int
		b.WriteString(" where ")
		if d != "" {
			b.WriteString(d)
			i = 1
		}

		if g != "" {
			if i == 1 {
				b.WriteString(" or ")
			}
			b.WriteString(g)
			i = 2
		}

		if t != "" {
			if i == 1 || i == 2 {
				b.WriteString(" or ")
			}
			b.WriteString(t)
		}

		return b.String()
	} else {
		return b.String()
	}
}
