package webserver

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

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
