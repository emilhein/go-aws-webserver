package webserver

import (
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func simple(w http.ResponseWriter, r *http.Request) {
	someMap := map[string]int{"Food": 1, "music": 2}
	printKeysAndValues(someMap)

	if result, message, err := computeTotal(5, 10); err != nil {
		fmt.Println("We got a big problem")
	} else {
		fmt.Println("We are perfect")
		fmt.Printf("Sum is %v and message is %s \n", result, message)

	}

	myMovie := Movie{Title: "Avengers", Year: 2018, Rating: "7.1"}
	myMovie.format()

}

func printKeysAndValues(themap map[string]int) string {
	for index, num := range themap {
		fmt.Printf("Index is %s, value is %v \n", index, num)
	}
	return "ok"
}

func computeTotal(a int, b int) (int, string, error) {
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := 10
	randomN := rand.Intn(max-min) + min
	fmt.Printf("The random number is :- %v \n", randomN)
	if randomN >= 5 {
		return a + b, "Everything ok we have a high number ", nil
	}
	return a + b, "Everything bad ", errors.New("We dont work with small numbers")
}
