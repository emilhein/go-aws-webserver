package webserver

import "testing"

func TestSum(t *testing.T) {
	myCinema := Movie{Title: "Batman", Rating: "8.8", Year: 2017}

	functionOutputString := myCinema.getMovies()
	if functionOutputString != "The movie Batman (2017) has a rating of 8.8" {
		t.Errorf("The output was not as expected")
	}
}
