package webserver

import (
	"fmt"
	"net/http"
	"sync"
)

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
