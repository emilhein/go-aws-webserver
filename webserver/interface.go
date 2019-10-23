package webserver

import (
	"fmt"
	"net/http"
)

type Movie struct {
	Title  string `json:"title"`
	Rating string `json:"rating"`
	Year   int    `json:"year"`
}

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

func PrintMoviesInCinema(cinema Cinema) string {
	return cinema.getMovies()
}

func InterfaceMethod(w http.ResponseWriter, r *http.Request) {
	myCinema := Movie{Title: "Batman", Rating: "8.8", Year: 2017}
	logForCiname := PrintMoviesInCinema(myCinema)
	fmt.Fprintf(w, "Interfaces used: %s", logForCiname)

}
