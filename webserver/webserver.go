package webserver

import (
	"encoding/json"
	"errors"
	"fmt"
	// "log"
	"github.com/gorilla/mux"

	"net/http"

	"math/rand"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

const S3_REGION = "eu-west-1"

// const S3_BUCKET = "ehe-development"

var wg sync.WaitGroup

func finder(mine []string, oreChannel chan string) {
	defer wg.Done()
	for _, value := range mine {
		if value == "ore" {
			oreChannel <- value //send item on oreChannel
		}
	}
	close(oreChannel)

}

func breaker(oreChannel <-chan string, minedOreChan chan<- string) {
	defer wg.Done()
	for elem := range oreChannel {
		fmt.Println("From Finder: ", elem)
		minedOreChan <- "minedOre" //send to minedOreChan
	}
	close(minedOreChan)

}
func smelter(minedOreChan <-chan string, name string, maxFound int) {
	defer wg.Done()
	for minedOre := range minedOreChan { //read from minedOreChan by ranging
		fmt.Println("From Miner in FUNCTION: ", minedOre)
		fmt.Printf("From Smelter (%s): Ore is smelted \n", name)
	}
}
func StartMining(w http.ResponseWriter, r *http.Request) {
	wg.Add(3)
	theMine := []string{"rock", "ore", "ore", "rock", "ore", "ore", "rock", "ore", "ore", "rock", "ore", "ore"}
	oreChannel := make(chan string)

	minedOreChan := make(chan string)
	// Finder
	go finder(theMine, oreChannel)
	// Ore Breaker
	go breaker(oreChannel, minedOreChan)

	// Smelters
	go smelter(minedOreChan, "Bob", len(theMine))

	wg.Wait()
	fmt.Println("Main: Completed")
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])

}

type Input struct {
	Bucket string `json:"bucket"`
	Key    string `json:"key"`
}

func GetS3File(w http.ResponseWriter, r *http.Request) {
	sess, err := session.NewSession(&aws.Config{Region: aws.String(S3_REGION)})
	if err != nil {
		// Handle error
	}
	w.Header().Set("Content-Type", "application/json")
	var input Input
	_ = json.NewDecoder(r.Body).Decode(&input)

	fmt.Printf("Getting: %s/%s \n", input.Bucket, input.Key)

	handler := S3Handler{
		Session: sess,
		Bucket:  input.Bucket,
	}
	fileString := input.Key // "movies.json"
	contents, err := handler.ReadFile(fileString)
	if err != nil {
		fmt.Println(err)
	}

	// toJson, error := extractMovieData(contents)
	// if err != nil {
	// 	fmt.Println(error)
	// }
	// for _, p := range toJson {
	// 	log.Printf("Name: %s , adsense: %s \n", p.Title, p.Rating)
	// }

	// fmt.Fprintf(w, contents)
	// var openReplacement interface{}
	// err := json.Marshal(contents, &openReplacement)
	var jsons interface{}
	error := json.Unmarshal(contents, &jsons)
	if error != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(jsons)
	// return contents

}
func Simple(w http.ResponseWriter, r *http.Request) {
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

	// fmt.Println(toJson) // "This is a test file"

	// fmt.Fprintf(w, "Check the log. Simple things has been written")

}
func Start() {
	r := mux.NewRouter()

	r.HandleFunc("/simple", Simple)
	r.HandleFunc("/getS3file", GetS3File).Methods("POST")
	r.HandleFunc("/interfaces", InterfaceMethod)
	r.HandleFunc("/startmining", StartMining)
	err := http.ListenAndServe(":3001", r)
	if err != nil {
		fmt.Printf("Could not start the server: %v", err)
	}

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
	} else {
		return a + b, "Everything bad ", errors.New("We dont work with small numbers")
	}
}
