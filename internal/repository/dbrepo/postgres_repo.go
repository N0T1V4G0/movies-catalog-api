package dbrepo

import (
	"backend/internal/models"
	"database/sql"
)

type PostgresDBRepo struct {
	DB *sql.DB
}

func (r *PostgresDBRepo) Connection() *sql.DB {
	return r.DB
}

func (r *PostgresDBRepo) AllMovies() ([]*models.Movie, error) {
	var movies []*models.Movie

	return movies, nil
}
