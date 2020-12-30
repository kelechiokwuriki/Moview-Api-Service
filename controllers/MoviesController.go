package controllers

import (
	"Moview/models"
	"Moview/repositories"
	"encoding/json"
	"math/rand"
	"net/http"
)

var (
	movieRepository repositories.MovieRepository = repositories.NewMovieRepository()
)

func GetAllMovies(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-type", "application/json")

	movies, err := movieRepository.GetAllMovies()

	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		responseWriter.Write([]byte(`{"error": "Error retrieving movies"}`))
		return
	}

	responseWriter.WriteHeader(http.StatusOK)
	json.NewEncoder(responseWriter).Encode(movies)
}

func CreateMovie(responseWriter http.ResponseWriter, request *http.Request) {
	var movie models.Movie

	err := json.NewDecoder(request.Body).Decode(&movie)

	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		responseWriter.Write([]byte(`{"error": "Unable to decode request"}`))
		return
	}

	movie.ID = rand.Int63()
	movieRepository.CreateMovie(&movie)

	responseWriter.WriteHeader(http.StatusOK)
	json.NewEncoder(responseWriter).Encode(movie)
}
