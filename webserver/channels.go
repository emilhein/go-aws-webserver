package webserver

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup

type MineShipment struct {
	Items []string
}

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
func smelter(minedOreChan <-chan string, shippingChannel chan<- string, name string) {
	defer wg.Done()
	for minedOre := range minedOreChan { //read from minedOreChan by ranging
		fmt.Println("From Miner in FUNCTION: ", minedOre)
		fmt.Printf("From Smelter (%s): Ore is smelted \n", name)
		var readyForShipment = fmt.Sprintf("%v made by %v (%v)", minedOre, name, time.Now())
		shippingChannel <- readyForShipment
	}
	close(shippingChannel)

}
func packer(shippingChannel <-chan string, doneChannel chan<- MineShipment) {

	var container = MineShipment{}

	for shipItem := range shippingChannel { //read from minedOreChan by ranging
		fmt.Printf("Packing... \n")
		container.Items = append(container.Items, shipItem)
	}
	doneChannel <- container
}

func StartMining(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	wg.Add(3)
	theMine := []string{"rock", "ore", "ore", "rock", "ore", "ore", "rock", "ore", "ore", "rock", "ore", "ore"}
	oreChannel := make(chan string)
	minedOreChan := make(chan string)
	shippingChannel := make(chan string)
	donePacking := make(chan MineShipment)

	// Finder
	go finder(theMine, oreChannel)
	// Ore Breaker
	go breaker(oreChannel, minedOreChan)

	// Smelters
	go smelter(minedOreChan, shippingChannel, "Peter")

	// wg.Wait()
	// Packer
	go packer(shippingChannel, donePacking)
	result := <-donePacking
	json.NewEncoder(w).Encode(result)

}
