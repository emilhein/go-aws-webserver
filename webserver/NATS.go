package webserver

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/nats-io/nats.go"
)

func NATStart(w http.ResponseWriter, r *http.Request) {
	// sc, err := stan.Connect(
	// 	nats.DefaultURL
	// 	"test-cluster",
	// 	"client-1",
	// 	stan.Pings(1, 3),
	// stan.NatsURL(strings.Join(os.Args[1:], ",")),
	// )
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalln(err)
	}

	c, _ := nats.NewEncodedConn(nc, "json")

	defer c.Close()

	sub, err := c.Subscribe("movies", func(m *Movie) {
		fmt.Printf("Received a movie! %+v\n", m)
	})
	if err != nil {
		log.Fatalln(err)
	}
	defer sub.Unsubscribe()

	for {
		movie := &Movie{Rating: "3.3", Title: "Breaking bad", Year: 2019}
		if err := c.Publish("movies", movie); err != nil {
			log.Fatalln(err)
		}
		time.Sleep(time.Millisecond * 1000)
	}

}
