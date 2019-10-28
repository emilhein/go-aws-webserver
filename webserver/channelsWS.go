package webserver

import (
	"fmt"
	"sync"
	"time"
)

var waitg sync.WaitGroup

type finderInput struct {
	mine            []string
	oreChannel      chan string
	brodcastChannel chan<- Event
}
type breakerInput struct {
	oreChannel      <-chan string
	minedOreChan    chan<- string
	brodcastChannel chan<- Event
}
type smelterInput struct {
	name            string
	minedOreChan    <-chan string
	shippingChannel chan<- string
	brodcastChannel chan<- Event
}
type packerInput struct {
	shippingChannel <-chan string
	doneChannel     chan<- MineShipment
	brodcastChannel chan<- Event
}

func formatTime(t time.Time) string {
	res := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d.%02d\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second(), t.Nanosecond())
	return res
}
func finderSocket(f finderInput) {
	defer waitg.Done()
	for index, value := range f.mine {
		if value == "ore" {
			oreNumber := fmt.Sprintf("We found ore (%v)", index+1)
			socketMsg := Event{Type: "Finder", Content: oreNumber, Timestamp: formatTime(time.Now())}
			f.brodcastChannel <- socketMsg
			f.oreChannel <- fmt.Sprintf("ore(%v)", index+1) //send item on oreChannel
		}
	}
	close(f.oreChannel)

}

func breakerSocket(b breakerInput) {
	defer waitg.Done()
	for elem := range b.oreChannel {
		fmt.Println("From Finder: ", elem)
		breakerMsg := fmt.Sprintf("We mined %v", elem)

		socketMsg := Event{Type: "Miner", Content: breakerMsg, Timestamp: formatTime(time.Now())}
		b.brodcastChannel <- socketMsg
		b.minedOreChan <- elem //send to minedOreChan

	}
	close(b.minedOreChan)

}
func smelterSocket(s smelterInput) {
	defer waitg.Done()
	for minedOre := range s.minedOreChan { //read from minedOreChan by ranging
		fmt.Println("From Miner in FUNCTION: ", minedOre)
		fmt.Printf("From Smelter (%s): Ore is smelted \n", s.name)
		smelterMsg := fmt.Sprintf("%v smelted %v", s.name, minedOre)

		// var readyForShipment = fmt.Sprintf("%v made by %v (%v)", minedOre, s.name, formatTime(time.Now()))
		socketMsg := Event{Type: "Smelter", Content: smelterMsg, Timestamp: formatTime(time.Now())}
		s.brodcastChannel <- socketMsg
		s.shippingChannel <- minedOre
	}
	close(s.shippingChannel)

}

func packerSocket(p packerInput) {

	var container = MineShipment{}

	for shipItem := range p.shippingChannel { //read from minedOreChan by ranging
		fmt.Printf("Packing... \n")
		container.Items = append(container.Items, shipItem)
		packerMsg := fmt.Sprintf("We packet %v", shipItem)

		socketMsg := Event{Type: "Packer", Content: packerMsg, Timestamp: formatTime(time.Now())}
		p.brodcastChannel <- socketMsg
	}
	p.doneChannel <- container
}

func startProcess(brodcastChannel chan<- Event) MineShipment {

	waitg.Add(4)
	theMine := []string{"rock", "ore", "ore", "rock", "ore", "ore", "rock", "ore", "ore", "rock", "ore", "ore"}
	oreChannel := make(chan string)
	minedOreChan := make(chan string)
	shippingChannel := make(chan string)
	donePacking := make(chan MineShipment)

	// Finder
	finderinput := finderInput{mine: theMine, oreChannel: oreChannel, brodcastChannel: brodcastChannel}
	go finderSocket(finderinput)
	// Ore Breaker
	breakerinput := breakerInput{oreChannel: oreChannel, minedOreChan: minedOreChan, brodcastChannel: brodcastChannel}
	go breakerSocket(breakerinput)

	// Smelters
	smelterinput := smelterInput{brodcastChannel: brodcastChannel, minedOreChan: minedOreChan, name: "Bob", shippingChannel: shippingChannel}
	go smelterSocket(smelterinput)

	// wg.Wait()
	// Packer
	packerinput := packerInput{brodcastChannel: brodcastChannel, shippingChannel: shippingChannel, doneChannel: donePacking}
	go packerSocket(packerinput)
	result := <-donePacking
	return result

}
