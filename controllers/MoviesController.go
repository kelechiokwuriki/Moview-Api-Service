package controllers

import (
	"Moview/models"
	"Moview/services"
	"encoding/json"
	"net/http"
)

var (
	movieService services.MovieService = services.NewMovieService()
)

func GetAllMovies(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-type", "application/json")

	movies, err := movieService.GetAllMovies()

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

	if validatedError := movieService.ValidateMovie(&movie); validatedError != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		responseWriter.Write([]byte(`{"error": "Error found when validating movie"}`))
		return
	}

	createdMovie, err := movieService.CreateMovie(&movie)

	responseWriter.WriteHeader(http.StatusOK)
	json.NewEncoder(responseWriter).Encode(createdMovie)
}
