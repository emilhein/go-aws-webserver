package webserver

import (
	"fmt"
	"net/http"
)

// Movie struct
type Movie struct {
	Title  string `json:"title"`
	Rating string `json:"rating"`
	Year   int    `json:"year"`
}

// Cinema interface
type Cinema interface {
	getMovies() string
}

func (m Movie) format() string {
	fmt.Printf("The movie %s (%v) has a rating of %s", m.Title, m.Year, m.Rating)
	return "OK"
}
func (m Movie) getMovies() string {
	return fmt.Sprintf("The movie %s (%v) has a rating of %s", m.Title, m.Year, m.Rating)
}

func printMoviesInCinema(cinema Cinema) string {
	return cinema.getMovies()
}

func interfaceMethod(w http.ResponseWriter, r *http.Request) {
	myCinema := Movie{Title: "Batman", Rating: "8.8", Year: 2017}
	logForCiname := printMoviesInCinema(myCinema)
	fmt.Fprintf(w, "Interfaces used: %s", logForCiname)

}
