package models

type Movie struct {
	ID            int64
	Name          string        `json:"name"`
	YearOfRelease int64         `json:"yearOfRelease"`
	Actors        []interface{} `json:"actors"`
}
