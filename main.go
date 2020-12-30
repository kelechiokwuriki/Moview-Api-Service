package main

import (
	"Moview/controllers"
	"context"
	"log"
	"net/http"

	firebase "firebase.google.com/go"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	const serverPort = ":8080"

	router.HandleFunc("/", controllers.GetAllMovies).Methods("GET")
	router.HandleFunc("/movies", controllers.CreateMovie).Methods("POST")
	firebase.NewApp(context.Background(), nil)

	// _, err := firebase.NewApp(context.Background(), nil)
	// if err != nil {
	// 	log.Fatalf("error initializing app: %v\n", err)
	// }

	log.Println("Server listening on port", serverPort)

	log.Fatalln(http.ListenAndServe(serverPort, router))
}
