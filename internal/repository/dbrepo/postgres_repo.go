package dbrepo

import (
	"backend/internal/models"
	"context"
	"database/sql"
	"time"
)

var dbTimeout = time.Second * 3

type PostgresDBRepo struct {
	DB *sql.DB
}

func (r *PostgresDBRepo) Connection() *sql.DB {
	return r.DB
}

func (r *PostgresDBRepo) AllMovies() ([]*models.Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		SELECT
			id, title, release_date, runtime,
			mpaa_rating, description, COALESCE(image, ''),
			created_at, updated_at
		FROM
			movies
		ORDER BY
			title
	`

	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []*models.Movie

	for rows.Next() {
		var movie models.Movie
		err := rows.Scan(
			&movie.Id,
			&movie.Title,
			&movie.ReleaseDate,
			&movie.Runtime,
			&movie.MpaaRating,
			&movie.Description,
			&movie.Image,
			&movie.CreatedAt,
			&movie.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		movies = append(movies, &movie)
	}

	return movies, nil
}
