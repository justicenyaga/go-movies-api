package repository

import (
	"database/sql"
	"vidly/internal/models"
)

type DatabaseRepo interface {
	Connection() *sql.DB

	AllMovies() ([]*models.Movie, error)
	OneMovie(id int) (*models.Movie, error)

	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id int) (*models.User, error)
}
