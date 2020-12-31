package services

import (
	"Moview/models"
	"errors"
)

type MovieService interface {
	GetAllMovies() ([]models.Movie, error)
	CreateMovie(movie *models.Movie) (*models.Movie, error)
	ValidateMovie(movie *models.Movie) error
}

type movieService struct{}

func NewMovieService() MovieService {
	return &movieService
}

func (*movieService) GetAllMovies() ([]models.Movie, error) {

}

func (*movieService) CreateMovie(movie *models.Movie) (*models.Movie, error) {

}

func (*movieService) ValidateMovie(movie *models.Movie) error {
	if movie == nil {
		return errors.New("Invalid movie")
	}

	if movie.Name == "" || movie.YearOfRelease == 0 || movie.Actors == nil {
		return errors.New("Invalid movie, missing either name, year of release or actors")
	}

	return nil
}
