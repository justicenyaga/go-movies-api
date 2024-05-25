package repository

import (
	"database/sql"
	"vidly/internal/models"
)

type DatabaseRepo interface {
	Connection() *sql.DB

	AllMovies() ([]*models.Movie, error)
	OneMovie(id int) (*models.Movie, error)
	OneMovieForEdit(id int) (*models.Movie, []*models.Genre, error)

	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id int) (*models.User, error)
}
