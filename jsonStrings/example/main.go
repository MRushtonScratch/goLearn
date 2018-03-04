package main

import (
	"github.com/MRushtonScratch/goLearn/jsonStrings"
	"fmt"
)

func main () {

	movie := `{ "title": "John Wick", "released": 2017, "color": true, "actors": [ "Keanu Reeves" ] }`

	jsonStrings.AddMovie(movie)

	movies := jsonStrings.GetMovies()
	fmt.Printf("%s\n", movies)

}


