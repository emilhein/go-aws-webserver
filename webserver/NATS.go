package webserver

import (
	"fmt"
	"net/http"

	"github.com/nats-io/stan.go"
)

func NATStart(w http.ResponseWriter, r *http.Request) {
	fmt.Println("YOLO")
	sc, err := stan.Connect("nats://demo.nats.io:4222", "Test")
	if err != nil {
		fmt.Print(err)

	}
	// Simple Synchronous Publisher
	sc.Publish("updates", []byte("Hello World")) // does not return until an ack has been received from NATS Streaming

	// Simple Async Subscriber
	sub, _ := sc.Subscribe("updates", func(m *stan.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})

	// Unsubscribe
	sub.Unsubscribe()

	// Close connection
	sc.Close()
	// nc, err := nats.Connect("demo.nats.io", nats.Name("API PublishBytes Example"))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer nc.Close()

	// if err := nc.Publish("updates", []byte("All is Well")); err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Print("MEssage sent")

	// wg := sync.WaitGroup{}
	// wg.Add(1)

	// // Subscribe
	// if _, err := nc.Subscribe("updates", func(m *nats.Msg) {
	// 	fmt.Println(m.Data)
	// 	wg.Done()
	// }); err != nil {
	// 	log.Fatal(err)
	// }

	// // Wait for a message to come in
	// wg.Wait()

}
