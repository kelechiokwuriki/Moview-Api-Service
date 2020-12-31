package repositories

import (
	"Moview/models"
	"context"
	"log"

	"cloud.google.com/go/firestore"
)

const projectId string = "moview-4862b"
const firestoreCollectionName string = "movies"

type fireStoreRepo struct{}

func NewFireStoreRepository() MovieRepository {
	return &fireStoreRepo{}
}

func (*fireStoreRepo) CreateMovie(movie *models.Movie) (*models.Movie, error) {
	//remember we added the path to the service account json file
	//in bash profile.
	//This code uses the path
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "moview-4862b")

	if err != nil {
		log.Fatalf("Failed to connect to firestore")
		return nil, err
	}

	//close client when function ends
	defer client.Close()

	_, _, err = client.Collection(firestoreCollectionName).Add(ctx, map[string]interface{}{
		"ID":            movie.ID,
		"Name":          movie.Name,
		"YearOfRelease": movie.YearOfRelease,
	})

	if err != nil {
		log.Fatalf("Failed to add create movie, error: %v", err)
		return nil, err
	}

	return movie, err
}

func (*fireStoreRepo) GetAllMovies() ([]models.Movie, error) {
	//remember we added the path to the service account json file
	//in bash profile.
	//This code uses the path
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "moview-4862b")

	if err != nil {
		log.Fatalf("Failed to connect to firestore")
		return nil, err
	}

	//close client when function ends
	defer client.Close()

	var movies []models.Movie

	iterator := client.Collection(firestoreCollectionName).Documents(ctx)
	for {
		doc, _ := iterator.Next()

		//exit if no movie item in iterator
		if doc == nil {
			break
		}

		movie := models.Movie{
			ID:            doc.Data()["ID"].(int64),
			Name:          doc.Data()["Name"].(string),
			YearOfRelease: doc.Data()["YearOfRelease"].(int64),
			Actors:        doc.Data()["Actors"].([]interface{}),
		}

		movies = append(movies, movie)

	}

	return movies, nil
}
