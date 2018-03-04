package jsonStrings

import (
	"encoding/json"
	"log"
)

type Movie struct {
	Title  string `json:"title"`
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string `json:"actors"`
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true,
		Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true,
		Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
}

func GetMovies() []byte {
	data, err :=  json.MarshalIndent(movies, "", "	")

	if err != nil {
		log.Fatalf("JSON marshalling failed", err)
	}

	return data
}

func AddMovie(jsonMovie string) {

	rawIn := json.RawMessage(jsonMovie)
	bytes, err := rawIn.MarshalJSON()
	if err != nil {
		panic(err)
	}

	var m Movie
	err = json.Unmarshal(bytes, &m)
	if err != nil {
		log.Fatalf("JSON unmarshalling failed: %s", err)
	}

	movies = append(movies, m)
}

