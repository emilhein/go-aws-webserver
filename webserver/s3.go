package webserver

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/emilhein/go-aws-utility/util/services"
)

// Input tpo the file retriever
type Input struct {
	Bucket    string   `json:"bucket"`
	Filepaths []string `json:"filepaths"`
}

/*
Takes JSON input "Bucket" & "Filepaths"


*/
func getS3Files(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var input Input
	_ = json.NewDecoder(r.Body).Decode(&input)
	var lowercasesBucket = strings.ToLower(input.Bucket)
	fmt.Printf("Getting: %s/%s \n", lowercasesBucket, input.Filepaths)

	inputToFunc := services.FilesInput{Bucket: lowercasesBucket, FileNames: input.Filepaths}
	returnValues := services.GetS3Files(inputToFunc)
	json.NewEncoder(w).Encode(returnValues)

}
