package repositories

import (
	"Moview/models"
)

type MovieRepository interface {
	CreateMovie(movie *models.Movie) (*models.Movie, error)
	GetAllMovies() ([]models.Movie, error)
	// UpdateMovie(id int, movie *models.Movie) (*models.Movie,)
}
