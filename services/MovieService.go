package services

import "Moview/models"

type MovieService interface {
	GetAllMovies() ([]models.Movie, error)
	CreateMovie(movie *models.Movie) (*models.Movie, error)
	ValidateMovie(movie *models.Movie) error
}
