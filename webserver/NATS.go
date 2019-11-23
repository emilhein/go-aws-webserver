package webserver

import (
	"encoding/json"
	"fmt"
	"net/http"

	stan "github.com/nats-io/stan.go"
)

func subscribe(w http.ResponseWriter, r *http.Request) {
	fmt.Println("We Start")
	getLast := make(chan uint64, 1)
	doneChannel := make(chan bool, 1)

	sc, err := stan.Connect("test-cluster", "clientID")
	if err != nil {
		fmt.Print(err)
	}

	// Simple Async Subscriber
	sub, _ := sc.Subscribe("animals", func(m *stan.Msg) {
		fmt.Printf("Received newest: %s \n", string(m.Data))
		getLast <- m.Sequence
	}, stan.StartWithLastReceived())

	lastSequenseNum := <-getLast

	accumulator := make(map[uint64]string)

	sc.Subscribe("animals", func(m *stan.Msg) {
		accumulator[m.Sequence] = string(m.Data)
		if m.Sequence == lastSequenseNum {
			doneChannel <- true
		}
	}, stan.DeliverAllAvailable())

	<-doneChannel
	fmt.Println("We Done")

	// Unsubscribe
	sub.Unsubscribe()

	// Close connection
	sc.Close()
	json.NewEncoder(w).Encode(accumulator)

}
